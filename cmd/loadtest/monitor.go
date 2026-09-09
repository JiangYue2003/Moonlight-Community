package main

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type monitorConfig struct {
	Enabled             bool          `json:",default=true"`
	Interval            time.Duration `json:",default=1s"`
	DockerContainers    []string
	StrategyContainer   string `json:",default=zg-knowpost"`
	MySQLDSN            string
	KafkaContainer      string `json:",default=zg-kafka"`
	KafkaBootstrap      string `json:",default=kafka:29092"`
	KafkaGroup          string `json:",default=feed-fanout-group"`
	PrometheusContainer string `json:",default=zg-knowpost"`
	PrometheusURL       string `json:",default=http://127.0.0.1:9104/metrics"`
	LocalStateFile      string `json:",optional"`
}

type collectorResult struct {
	name    string
	samples []resourceSample
	err     error
}

type dockerContainerState struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Running      bool   `json:"running"`
	Health       string `json:"health"`
	RestartCount int    `json:"restart_count"`
}

type monitorCollectionStatus struct {
	feedMetrics      bool
	dockerState      bool
	localProcess     bool
	processStartTime float64
}

type resourceMonitor struct {
	cfg       monitorConfig
	redisAddr string
	cancel    context.CancelFunc
	done      chan struct{}

	mu                 sync.Mutex
	samples            []resourceSample
	seen               map[string]bool
	measurementStarted time.Time
	initial            monitorCollectionStatus
	final              monitorCollectionStatus
}

func (m *resourceMonitor) ResetCounterBaselines(parent context.Context) (time.Duration, error) {
	started := time.Now()
	if !m.cfg.Enabled {
		return 0, fmt.Errorf("resource monitor is disabled")
	}
	type boundaryResult struct {
		source  string
		samples []resourceSample
		err     error
	}
	collectorCount := 2
	if strings.TrimSpace(m.cfg.LocalStateFile) != "" {
		collectorCount++
	}
	results := make(chan boundaryResult, collectorCount)
	go func() {
		values, err := collectKnowPostMetrics(parent, m.cfg)
		results <- boundaryResult{source: "prometheus:knowpost", samples: valueSample("prometheus:knowpost", values, started, err), err: err}
	}()
	go func() {
		values, runID, err := collectRedisMetrics(parent, m.redisAddr)
		samples := valueSample("redis", values, started, err)
		if len(samples) == 1 {
			samples[0].Identity = runID
		}
		results <- boundaryResult{source: "redis", samples: samples, err: err}
	}()
	if strings.TrimSpace(m.cfg.LocalStateFile) != "" {
		go func() {
			samples, err := collectLocalProcessStats(m.cfg.LocalStateFile, started)
			results <- boundaryResult{source: "local_process", samples: samples, err: err}
		}()
	}
	collected := make([]resourceSample, 0, collectorCount)
	for range collectorCount {
		result := <-results
		if result.err != nil {
			return time.Since(started), fmt.Errorf("reset %s baseline: %w", result.source, result.err)
		}
		collected = append(collected, result.samples...)
	}
	m.mu.Lock()
	m.measurementStarted = started
	filtered := m.samples[:0]
	for _, sample := range m.samples {
		if strings.HasPrefix(sample.Source, "docker-state:") {
			filtered = append(filtered, sample)
		}
	}
	m.samples = append(filtered, collected...)
	m.mu.Unlock()
	return time.Since(started), nil
}

func startResourceMonitor(parent context.Context, cfg monitorConfig, redisAddr string) *resourceMonitor {
	monitor := &resourceMonitor{cfg: cfg, redisAddr: redisAddr, done: make(chan struct{}), seen: make(map[string]bool)}
	if !cfg.Enabled {
		close(monitor.done)
		return monitor
	}
	ctx, cancel := context.WithCancel(parent)
	monitor.cancel = cancel
	// Capture cumulative application counters before issuing load so the final
	// report can use exact run deltas instead of process-lifetime totals.
	initialCtx, initialCancel := context.WithTimeout(ctx, 8*time.Second)
	monitor.initial = monitor.collect(initialCtx)
	initialCancel()
	go func() {
		defer close(monitor.done)
		ticker := time.NewTicker(cfg.Interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				monitor.collect(ctx)
			}
		}
	}()
	return monitor
}

func (m *resourceMonitor) Stop(parent context.Context) ([]resourceSample, []string) {
	if m.cancel != nil {
		m.cancel()
	}
	<-m.done
	if m.cfg.Enabled {
		finalCtx, cancel := context.WithTimeout(parent, 8*time.Second)
		m.final = m.collect(finalCtx)
		cancel()
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	samples := append([]resourceSample(nil), m.samples...)
	sort.Slice(samples, func(i, j int) bool { return samples[i].Timestamp.Before(samples[j].Timestamp) })
	processMissing := addProcessCPUPercent(samples)
	missing := make([]string, 0, 4)
	requiredSources := []string{"docker", "docker_state", "feed_metrics", "kafka", "mysql", "redis"}
	if strings.TrimSpace(m.cfg.LocalStateFile) != "" {
		requiredSources = append(requiredSources, "local_process")
		if !m.final.localProcess {
			missing = append(missing, "local_process_final")
		}
	}
	for _, source := range requiredSources {
		if !m.seen[source] {
			missing = append(missing, source)
		}
	}
	missing = append(missing, processMissing...)
	missing = append(missing, validateMonitorEvidence(samples, m.initial, m.final)...)
	return samples, missing
}

func (m *resourceMonitor) collect(parent context.Context) monitorCollectionStatus {
	var status monitorCollectionStatus
	if parent.Err() != nil {
		return status
	}
	timestamp := time.Now()
	collectorCount := 6
	if strings.TrimSpace(m.cfg.LocalStateFile) != "" {
		collectorCount++
	}
	results := make(chan collectorResult, collectorCount)
	go func() {
		samples, err := collectDockerStats(parent, m.cfg.DockerContainers, timestamp)
		results <- collectorResult{name: "docker", samples: samples, err: err}
	}()
	go func() {
		samples, _, err := collectDockerState(parent, m.cfg.DockerContainers, timestamp)
		results <- collectorResult{name: "docker_state", samples: samples, err: err}
	}()
	go func() {
		values, err := collectKafkaLag(parent, m.cfg)
		results <- collectorResult{name: "kafka", samples: valueSample("kafka", values, timestamp, err), err: err}
	}()
	go func() {
		values, runID, err := collectRedisMetrics(parent, m.redisAddr)
		samples := valueSample("redis", values, timestamp, err)
		if len(samples) == 1 {
			samples[0].Identity = runID
		}
		results <- collectorResult{name: "redis", samples: samples, err: err}
	}()
	go func() {
		values, err := collectMySQLMetrics(parent, m.cfg.MySQLDSN)
		results <- collectorResult{name: "mysql", samples: valueSample("mysql", values, timestamp, err), err: err}
	}()
	go func() {
		values, err := collectKnowPostMetrics(parent, m.cfg)
		results <- collectorResult{
			name:    "feed_metrics",
			samples: valueSample("prometheus:knowpost", values, timestamp, err),
			err:     err,
		}
	}()
	if strings.TrimSpace(m.cfg.LocalStateFile) != "" {
		go func() {
			samples, err := collectLocalProcessStats(m.cfg.LocalStateFile, timestamp)
			results <- collectorResult{name: "local_process", samples: samples, err: err}
		}()
	}

	for range collectorCount {
		result := <-results
		if result.err == nil && len(result.samples) > 0 {
			switch result.name {
			case "feed_metrics":
				status.feedMetrics = true
				status.processStartTime = result.samples[0].Values["process_start_time_seconds"]
			case "docker_state":
				status.dockerState = true
			case "local_process":
				status.localProcess = true
			}
		}
		m.mu.Lock()
		if result.err == nil && len(result.samples) > 0 {
			m.seen[result.name] = true
			for _, sample := range result.samples {
				if m.measurementStarted.IsZero() || !sample.Timestamp.Before(m.measurementStarted) {
					m.samples = append(m.samples, sample)
				}
			}
		}
		m.mu.Unlock()
	}
	return status
}

func collectKnowPostMetrics(ctx context.Context, cfg monitorConfig) (map[string]float64, error) {
	if strings.TrimSpace(cfg.PrometheusURL) == "" {
		return nil, fmt.Errorf("KnowPost Prometheus collector is not configured")
	}
	if strings.TrimSpace(cfg.PrometheusContainer) == "" {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.PrometheusURL, nil)
		if err != nil {
			return nil, fmt.Errorf("create KnowPost Prometheus request: %w", err)
		}
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, fmt.Errorf("collect local KnowPost Prometheus metrics: %w", err)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("collect local KnowPost Prometheus metrics: status %s", response.Status)
		}
		raw, err := io.ReadAll(io.LimitReader(response.Body, 8<<20))
		if err != nil {
			return nil, fmt.Errorf("read local KnowPost Prometheus metrics: %w", err)
		}
		return parseKnowPostFeedMetrics(raw)
	}
	command := exec.CommandContext(
		ctx,
		"docker",
		"exec",
		cfg.PrometheusContainer,
		"wget",
		"-qO-",
		cfg.PrometheusURL,
	)
	raw, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("collect KnowPost Prometheus metrics: %w", err)
	}
	return parseKnowPostFeedMetrics(raw)
}

func parseKnowPostFeedMetrics(raw []byte) (map[string]float64, error) {
	values := make(map[string]float64)
	feedMetricCount := 0
	scanner := bufio.NewScanner(bytes.NewReader(raw))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "process_start_time_seconds" {
			value, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return nil, fmt.Errorf("parse process_start_time_seconds value %q: %w", fields[1], err)
			}
			values["process_start_time_seconds"] = value
			continue
		}
		if len(fields) >= 2 {
			runtimeKey := ""
			switch fields[0] {
			case "go_gc_duration_seconds_sum",
				"go_gc_duration_seconds_count",
				"go_memstats_alloc_bytes_total",
				"go_memstats_heap_alloc_bytes":
				runtimeKey = fields[0]
			}
			if runtimeKey != "" {
				value, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					return nil, fmt.Errorf("parse Prometheus runtime metric %s value %q: %w", runtimeKey, fields[1], err)
				}
				values[runtimeKey] = value
				continue
			}
		}
		if len(fields) >= 2 {
			gaugeKey := ""
			switch fields[0] {
			case "zhiguang_knowpost_feed_page_refresh_queue_depth":
				gaugeKey = "refresh_queue_depth"
			case "zhiguang_knowpost_feed_page_refresh_active_workers":
				gaugeKey = "refresh_active_workers"
			case "zhiguang_knowpost_feed_page_refresh_pending_keys":
				gaugeKey = "refresh_pending_keys"
			}
			if gaugeKey != "" {
				value, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					return nil, fmt.Errorf("parse Prometheus gauge %s value %q: %w", fields[0], fields[1], err)
				}
				values[gaugeKey] = value
				feedMetricCount++
				continue
			}
		}
		open := strings.IndexByte(line, '{')
		close := strings.LastIndexByte(line, '}')
		if open <= 0 || close <= open {
			continue
		}
		name := line[:open]
		kind := ""
		switch name {
		case "zhiguang_knowpost_feed_stage_total":
			kind = "stage_total"
		case "zhiguang_knowpost_feed_stage_duration_ms_sum":
			kind = "stage_duration_sum"
		case "zhiguang_knowpost_feed_stage_duration_ms_count":
			kind = "stage_duration_count"
		case "zhiguang_knowpost_feed_dependency_calls_total":
			kind = "dependency_calls"
		case "zhiguang_knowpost_feed_cold_compute_total":
			kind = "cold_compute"
		case "zhiguang_knowpost_feed_page_cache_total":
			kind = "page_cache"
		case "zhiguang_knowpost_feed_pagination_total":
			kind = "pagination"
		case "zhiguang_knowpost_feed_pagination_duration_ms_sum":
			kind = "pagination_duration_sum"
		case "zhiguang_knowpost_feed_pagination_duration_ms_count":
			kind = "pagination_duration_count"
		case "zhiguang_knowpost_feed_cursor_work_total":
			kind = "cursor_work"
		default:
			continue
		}

		metricFields := strings.Fields(line[close+1:])
		if len(metricFields) == 0 {
			return nil, fmt.Errorf("Prometheus metric %s has no value", name)
		}
		value, err := strconv.ParseFloat(metricFields[0], 64)
		if err != nil {
			return nil, fmt.Errorf("parse Prometheus metric %s value %q: %w", name, metricFields[0], err)
		}
		labels := parsePrometheusLabels(line[open+1 : close])
		var key string
		switch kind {
		case "stage_total", "stage_duration_sum", "stage_duration_count":
			key = strings.Join([]string{kind, labels["stage"], labels["outcome"]}, ":")
		case "dependency_calls":
			key = strings.Join([]string{kind, labels["dependency"], labels["operation"], labels["outcome"]}, ":")
		case "cold_compute":
			key = strings.Join([]string{kind, labels["outcome"]}, ":")
		case "page_cache":
			key = strings.Join([]string{kind, labels["source"], labels["outcome"]}, ":")
		case "pagination", "pagination_duration_sum", "pagination_duration_count":
			key = strings.Join([]string{kind, labels["mode"], labels["result"]}, ":")
		case "cursor_work":
			key = strings.Join([]string{kind, labels["work"]}, ":")
		}
		if strings.Contains(key, "::") || strings.HasSuffix(key, ":") {
			return nil, fmt.Errorf("Prometheus metric %s is missing required labels", name)
		}
		values[key] = value
		feedMetricCount++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan KnowPost Prometheus metrics: %w", err)
	}
	if feedMetricCount == 0 {
		return nil, fmt.Errorf("KnowPost Prometheus output contains no Feed metrics")
	}
	return values, nil
}

func parsePrometheusLabels(raw string) map[string]string {
	labels := make(map[string]string)
	for _, item := range strings.Split(raw, ",") {
		parts := strings.SplitN(strings.TrimSpace(item), "=", 2)
		if len(parts) != 2 {
			continue
		}
		labels[parts[0]] = strings.Trim(parts[1], "\"")
	}
	return labels
}

func collectDockerState(ctx context.Context, containers []string, timestamp time.Time) ([]resourceSample, []dockerContainerState, error) {
	if len(containers) == 0 {
		return nil, nil, fmt.Errorf("no Docker containers configured")
	}
	args := append([]string{"inspect"}, containers...)
	raw, err := exec.CommandContext(ctx, "docker", args...).Output()
	if err != nil {
		return nil, nil, fmt.Errorf("docker inspect: %w", err)
	}
	return parseDockerState(raw, timestamp)
}

func parseDockerState(raw []byte, timestamp time.Time) ([]resourceSample, []dockerContainerState, error) {
	var inspected []struct {
		ID           string `json:"Id"`
		Name         string `json:"Name"`
		RestartCount int    `json:"RestartCount"`
		State        struct {
			Running bool `json:"Running"`
			Health  *struct {
				Status string `json:"Status"`
			} `json:"Health"`
		} `json:"State"`
	}
	if err := json.Unmarshal(raw, &inspected); err != nil {
		return nil, nil, fmt.Errorf("decode docker inspect: %w", err)
	}
	if len(inspected) == 0 {
		return nil, nil, fmt.Errorf("Docker returned no container state")
	}
	samples := make([]resourceSample, 0, len(inspected))
	states := make([]dockerContainerState, 0, len(inspected))
	for _, item := range inspected {
		name := strings.TrimPrefix(item.Name, "/")
		health := "not_configured"
		healthValue := 0.0
		if item.State.Health != nil {
			health = item.State.Health.Status
			if health == "healthy" {
				healthValue = 1
			}
		}
		runningValue := 0.0
		if item.State.Running {
			runningValue = 1
		}
		states = append(states, dockerContainerState{
			ID: item.ID, Name: name, Running: item.State.Running, Health: health, RestartCount: item.RestartCount,
		})
		samples = append(samples, resourceSample{
			Timestamp: timestamp,
			Source:    "docker-state:" + name,
			Identity:  item.ID,
			Values: map[string]float64{
				"running": runningValue,
				"healthy": healthValue,
				"health_configured": func() float64 {
					if item.State.Health != nil {
						return 1
					}
					return 0
				}(),
				"restart_count": float64(item.RestartCount),
			},
		})
	}
	return samples, states, nil
}

func validateMonitorEvidence(
	samples []resourceSample,
	initial monitorCollectionStatus,
	final monitorCollectionStatus,
) []string {
	missing := make([]string, 0, 8)
	add := func(value string) {
		if !containsString(missing, value) {
			missing = append(missing, value)
		}
	}
	if !initial.feedMetrics {
		add("feed_metrics_initial")
	}
	if !final.feedMetrics {
		add("feed_metrics_final")
	}
	if initial.feedMetrics && final.feedMetrics {
		if initial.processStartTime <= 0 || final.processStartTime <= 0 {
			add("feed_process_identity")
		} else if initial.processStartTime != final.processStartTime {
			add("feed_process_changed")
		}
	}
	if !initial.dockerState {
		add("docker_state_initial")
	}
	if !final.dockerState {
		add("docker_state_final")
	}

	type state struct {
		identity string
		restarts float64
	}
	first := make(map[string]state)
	for _, sample := range samples {
		if !strings.HasPrefix(sample.Source, "docker-state:") {
			continue
		}
		current, ok := first[sample.Source]
		if !ok {
			first[sample.Source] = state{identity: sample.Identity, restarts: sample.Values["restart_count"]}
			current = first[sample.Source]
		}
		if current.identity != "" && sample.Identity != "" && current.identity != sample.Identity {
			add("container_identity_changed")
		}
		if sample.Values["restart_count"] > current.restarts {
			add("container_restart")
		}
		if sample.Values["running"] < 1 ||
			(sample.Values["health_configured"] > 0 && sample.Values["healthy"] < 1) {
			add("container_unhealthy")
		}
	}
	return missing
}

func valueSample(source string, values map[string]float64, timestamp time.Time, err error) []resourceSample {
	if err != nil {
		return nil
	}
	return []resourceSample{{Timestamp: timestamp, Source: source, Values: values}}
}

func collectDockerStats(ctx context.Context, containers []string, timestamp time.Time) ([]resourceSample, error) {
	if len(containers) == 0 {
		return nil, fmt.Errorf("no Docker containers configured")
	}
	args := []string{"stats", "--no-stream", "--format", "{{json .}}"}
	args = append(args, containers...)
	raw, err := exec.CommandContext(ctx, "docker", args...).Output()
	if err != nil {
		return nil, fmt.Errorf("docker stats: %w", err)
	}
	return parseDockerStats(raw, timestamp)
}

func parseDockerStats(raw []byte, timestamp time.Time) ([]resourceSample, error) {
	type dockerLine struct {
		Name    string `json:"Name"`
		CPUPerc string `json:"CPUPerc"`
		MemPerc string `json:"MemPerc"`
		PIDs    string `json:"PIDs"`
	}
	var samples []resourceSample
	scanner := bufio.NewScanner(strings.NewReader(string(raw)))
	for scanner.Scan() {
		var line dockerLine
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			return nil, fmt.Errorf("decode docker stats: %w", err)
		}
		cpu, err := parsePercent(line.CPUPerc)
		if err != nil {
			return nil, fmt.Errorf("parse Docker CPU for %s: %w", line.Name, err)
		}
		memory, err := parsePercent(line.MemPerc)
		if err != nil {
			return nil, fmt.Errorf("parse Docker memory for %s: %w", line.Name, err)
		}
		pids, err := strconv.ParseFloat(strings.TrimSpace(line.PIDs), 64)
		if err != nil {
			return nil, fmt.Errorf("parse Docker PIDs for %s: %w", line.Name, err)
		}
		samples = append(samples, resourceSample{
			Timestamp: timestamp,
			Source:    "docker:" + line.Name,
			Values: map[string]float64{
				"cpu_percent":    cpu,
				"memory_percent": memory,
				"pids":           pids,
			},
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(samples) == 0 {
		return nil, fmt.Errorf("Docker returned no stats")
	}
	return samples, nil
}

func parsePercent(value string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(value, "%")), 64)
}

func collectKafkaLag(ctx context.Context, cfg monitorConfig) (map[string]float64, error) {
	raw, err := exec.CommandContext(ctx,
		"docker", "exec", cfg.KafkaContainer,
		"/opt/bitnami/kafka/bin/kafka-consumer-groups.sh",
		"--bootstrap-server", cfg.KafkaBootstrap,
		"--group", cfg.KafkaGroup,
		"--describe",
	).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Kafka lag: %w: %s", err, strings.TrimSpace(string(raw)))
	}
	return parseKafkaLag(raw)
}

func parseKafkaLag(raw []byte) (map[string]float64, error) {
	currentOffsetIndex := -1
	logEndOffsetIndex := -1
	lagIndex := -1
	var total, maxLag, currentOffsetTotal, logEndOffsetTotal float64
	partitions := 0
	scanner := bufio.NewScanner(strings.NewReader(string(raw)))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		if lagIndex < 0 {
			for i, field := range fields {
				switch field {
				case "CURRENT-OFFSET":
					currentOffsetIndex = i
				case "LOG-END-OFFSET":
					logEndOffsetIndex = i
				case "LAG":
					lagIndex = i
				}
			}
			continue
		}
		if len(fields) <= lagIndex || len(fields) <= currentOffsetIndex || len(fields) <= logEndOffsetIndex {
			continue
		}
		logEndOffset, err := strconv.ParseFloat(fields[logEndOffsetIndex], 64)
		if err != nil {
			continue
		}
		currentOffset, currentErr := strconv.ParseFloat(fields[currentOffsetIndex], 64)
		if currentErr != nil {
			if fields[currentOffsetIndex] != "-" {
				continue
			}
			currentOffset = 0
		}
		lag, lagErr := strconv.ParseFloat(fields[lagIndex], 64)
		if lagErr != nil {
			if fields[lagIndex] != "-" || fields[currentOffsetIndex] != "-" {
				continue
			}
			lag = logEndOffset
		}
		total += lag
		currentOffsetTotal += currentOffset
		logEndOffsetTotal += logEndOffset
		if lag > maxLag {
			maxLag = lag
		}
		partitions++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if currentOffsetIndex < 0 || logEndOffsetIndex < 0 || lagIndex < 0 || partitions == 0 {
		return nil, fmt.Errorf("Kafka lag output has no partition rows")
	}
	return map[string]float64{
		"lag_total": total, "lag_max": maxLag, "partitions": float64(partitions),
		"current_offset_total": currentOffsetTotal, "log_end_offset_total": logEndOffsetTotal,
	}, nil
}

func collectRedisMetrics(ctx context.Context, addr string) (map[string]float64, string, error) {
	client := redis.NewClient(&redis.Options{Addr: addr})
	defer client.Close()
	raw, err := client.Info(ctx).Result()
	if err != nil {
		return nil, "", err
	}
	values, err := parseRedisInfo([]byte(raw))
	if err != nil {
		return nil, "", err
	}
	runID, err := parseRedisRunID(raw)
	if err != nil {
		return nil, "", err
	}
	epoch, err := client.Get(ctx, "feed:content:safety:epoch").Float64()
	if err != nil && err != redis.Nil {
		return nil, "", fmt.Errorf("read Feed safety epoch: %w", err)
	}
	values["feed_safety_epoch"] = epoch
	return values, runID, nil
}

func parseRedisRunID(raw string) (string, error) {
	for _, line := range strings.Split(raw, "\n") {
		key, value, ok := strings.Cut(strings.TrimSpace(line), ":")
		if ok && key == "run_id" && strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value), nil
		}
	}
	return "", fmt.Errorf("Redis INFO contained no run_id")
}

func parseRedisInfo(raw []byte) (map[string]float64, error) {
	values := make(map[string]float64)
	wanted := map[string]string{
		"uptime_in_seconds":         "uptime_seconds",
		"used_memory":               "used_memory_bytes",
		"instantaneous_ops_per_sec": "ops_per_sec",
		"total_commands_processed":  "commands_total",
		"total_net_input_bytes":     "net_input_bytes",
		"total_net_output_bytes":    "net_output_bytes",
		"keyspace_hits":             "keyspace_hits",
		"keyspace_misses":           "keyspace_misses",
		"connected_clients":         "connected_clients",
		"blocked_clients":           "blocked_clients",
		"evicted_keys":              "evicted_keys",
		"rejected_connections":      "rejected_connections",
	}
	scanner := bufio.NewScanner(strings.NewReader(string(raw)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "db") && strings.Contains(line, ":keys=") {
			parts := strings.SplitN(line, ":keys=", 2)
			count := strings.SplitN(parts[1], ",", 2)[0]
			parsed, err := strconv.ParseFloat(count, 64)
			if err == nil {
				values["keys"] += parsed
			}
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		outputName, ok := wanted[parts[0]]
		if !ok {
			continue
		}
		parsed, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err == nil {
			values[outputName] = parsed
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	hits, misses := values["keyspace_hits"], values["keyspace_misses"]
	if hits+misses > 0 {
		values["hit_rate"] = hits / (hits + misses)
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("Redis INFO contained no supported metrics")
	}
	return values, nil
}

func collectMySQLMetrics(ctx context.Context, dsn string) (map[string]float64, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, fmt.Errorf("MySQL DSN is empty")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.QueryContext(ctx, `SHOW GLOBAL STATUS WHERE Variable_name IN ('Questions','Threads_connected','Threads_running','Slow_queries')`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	values := make(map[string]float64)
	for rows.Next() {
		var name, raw string
		if err := rows.Scan(&name, &raw); err != nil {
			return nil, err
		}
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			continue
		}
		values[strings.ToLower(name)] = value
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("MySQL returned no status metrics")
	}
	return values, nil
}

func reportIsComplete(stages []stageResult, missing []string) bool {
	if len(stages) == 0 {
		return false
	}
	for _, stage := range stages {
		if stage.Failed > 0 || stage.Total == 0 {
			return false
		}
		if observation := stage.Pagination; observation != nil && strings.Contains(observation.Mode, "cursor") {
			if observation.OracleHash == "" || observation.ObservedHash != observation.OracleHash ||
				observation.DuplicateItems > 0 || observation.OracleMismatches > 0 ||
				observation.CursorLoops > 0 || observation.EarlyTerminations > 0 {
				return false
			}
		}
	}
	for _, metric := range missing {
		switch metric {
		case "docker", "docker_state", "kafka", "kafka_recovery", "mysql", "redis", "redis_metrics_delta",
			"redis_evictions", "redis_rejections", "redis_safety_epoch_changed", "feed_runtime_flags",
			"feed_metrics", "feed_metrics_initial", "feed_metrics_final",
			"feed_process_identity", "feed_process_changed", "docker_state_initial",
			"docker_state_final", "container_identity_changed", "container_restart",
			"container_unhealthy", "feed_degraded", "client_cpu", "client_cpu_saturated",
			"local_process", "local_process_final", "local_process_boundary", "local_process_changed",
			"environment_evidence", "cursor_oracle", "cursor_correctness", "run_error",
			"mutation_checkpoint", "mutation_post_ids", "mutation_restore_pending", "mutation_restore", "mutation_restore_scope",
			"mutation_provisional_report", "mutation_preparation_pending", "mutation_preparation":
			return false
		}
	}
	return true
}
