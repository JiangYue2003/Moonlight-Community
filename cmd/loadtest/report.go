package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type resourceSample struct {
	Timestamp time.Time          `json:"timestamp"`
	Source    string             `json:"source"`
	Identity  string             `json:"identity,omitempty"`
	Values    map[string]float64 `json:"values"`
}

type benchmarkReport struct {
	RunID          string                    `json:"run_id"`
	Strategy       string                    `json:"strategy"`
	Entry          string                    `json:"entry"`
	Scenario       string                    `json:"scenario"`
	StartedAt      time.Time                 `json:"started_at"`
	Duration       time.Duration             `json:"duration_ns"`
	Concurrency    int                       `json:"concurrency"`
	Trial          int                       `json:"trial,omitempty"`
	ExpectedTrials int                       `json:"expected_trials,omitempty"`
	Complete       bool                      `json:"complete"`
	MissingMetrics []string                  `json:"missing_metrics,omitempty"`
	Environment    map[string]string         `json:"environment,omitempty"`
	Stages         []stageResult             `json:"stages"`
	FeedMetrics    *feedMetricsSummary       `json:"feed_metrics,omitempty"`
	RedisMetrics   *redisMetricsSummary      `json:"redis_metrics,omitempty"`
	Resources      []resourceSample          `json:"resources,omitempty"`
	Recovery       *recoveryObservation      `json:"recovery,omitempty"`
	Mutation       *mutationTrialObservation `json:"mutation,omitempty"`
	Notes          []string                  `json:"notes,omitempty"`
}

type redisMetricsSummary struct {
	CommandsDelta            float64 `json:"commands_delta"`
	NetInputBytesDelta       float64 `json:"net_input_bytes_delta"`
	NetOutputBytesDelta      float64 `json:"net_output_bytes_delta"`
	KeyspaceHitsDelta        float64 `json:"keyspace_hits_delta"`
	KeyspaceMissesDelta      float64 `json:"keyspace_misses_delta"`
	RunHitRate               float64 `json:"run_hit_rate"`
	EvictedKeysDelta         float64 `json:"evicted_keys_delta"`
	RejectedConnectionsDelta float64 `json:"rejected_connections_delta"`
	OpsPerSecondMax          float64 `json:"ops_per_sec_max"`
	SafetyEpochStart         float64 `json:"safety_epoch_start"`
	SafetyEpochEnd           float64 `json:"safety_epoch_end"`
}

type feedMetricsSummary struct {
	StageCalls                     map[string]float64            `json:"stage_calls,omitempty"`
	StageMeanMilliseconds          map[string]float64            `json:"stage_mean_ms,omitempty"`
	StageCallsByOutcome            map[string]map[string]float64 `json:"stage_calls_by_outcome,omitempty"`
	StageMeanMillisecondsByOutcome map[string]map[string]float64 `json:"stage_mean_ms_by_outcome,omitempty"`
	DependencyCalls                map[string]float64            `json:"dependency_calls,omitempty"`
	DependencyCallsPerSuccess      map[string]float64            `json:"dependency_calls_per_success,omitempty"`
	DependencyCallsByOutcome       map[string]map[string]float64 `json:"dependency_calls_by_outcome,omitempty"`
	DependencyOperations           map[string]float64            `json:"dependency_operations,omitempty"`
	DependencyOperationsPerSuccess map[string]float64            `json:"dependency_operations_per_success,omitempty"`
	DependencyOperationsByOutcome  map[string]map[string]float64 `json:"dependency_operations_by_outcome,omitempty"`
	ColdComputes                   float64                       `json:"cold_computes"`
	ColdComputesPerSuccess         float64                       `json:"cold_computes_per_success"`
	ColdComputesByOutcome          map[string]float64            `json:"cold_computes_by_outcome,omitempty"`
	PageCacheSources               map[string]float64            `json:"page_cache_sources,omitempty"`
	PageCacheSourcesByOutcome      map[string]map[string]float64 `json:"page_cache_sources_by_outcome,omitempty"`
	PageCacheFreshHitRatio         float64                       `json:"page_cache_fresh_hit_ratio"`
	PaginationModes                map[string]float64            `json:"pagination_modes,omitempty"`
	PaginationModesByResult        map[string]map[string]float64 `json:"pagination_modes_by_result,omitempty"`
	PaginationMeanMilliseconds     map[string]float64            `json:"pagination_mean_ms,omitempty"`
	CursorWork                     map[string]float64            `json:"cursor_work,omitempty"`
	CursorWorkPerSuccess           map[string]float64            `json:"cursor_work_per_success,omitempty"`
	RefreshQueueDepthMax           float64                       `json:"refresh_queue_depth_max"`
	RefreshActiveWorkersMax        float64                       `json:"refresh_active_workers_max"`
	RefreshPendingKeysMax          float64                       `json:"refresh_pending_keys_max"`
	Degraded                       bool                          `json:"degraded"`
}

type recoveryObservation struct {
	KafkaDrain   time.Duration      `json:"kafka_drain_ns"`
	FinalMetrics map[string]float64 `json:"final_metrics,omitempty"`
	Complete     bool               `json:"complete"`
	Error        string             `json:"error,omitempty"`
}

type reportPaths struct {
	JSON     string
	CSV      string
	Markdown string
}

func writeReportFiles(root string, report benchmarkReport) (reportPaths, error) {
	dir := filepath.Join(
		root,
		safePathSegment(report.RunID),
		safePathSegment(report.Strategy+"-"+report.Entry+"-"+report.Scenario),
	)
	if report.Trial > 0 {
		dir = filepath.Join(dir, fmt.Sprintf("trial-%d", report.Trial))
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return reportPaths{}, fmt.Errorf("create report directory: %w", err)
	}

	paths := reportPaths{
		JSON:     filepath.Join(dir, "report.json"),
		CSV:      filepath.Join(dir, "stages.csv"),
		Markdown: filepath.Join(dir, "report.md"),
	}
	jsonBody, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return reportPaths{}, fmt.Errorf("marshal JSON report: %w", err)
	}
	if err := os.WriteFile(paths.JSON, append(jsonBody, '\n'), 0o644); err != nil {
		return reportPaths{}, fmt.Errorf("write JSON report: %w", err)
	}
	if err := writeCSVReport(paths.CSV, report); err != nil {
		return reportPaths{}, err
	}
	if err := os.WriteFile(paths.Markdown, []byte(markdownReport(report)), 0o644); err != nil {
		return reportPaths{}, fmt.Errorf("write Markdown report: %w", err)
	}
	return paths, nil
}

func writeCSVReport(path string, report benchmarkReport) error {
	var body bytes.Buffer
	writer := csv.NewWriter(&body)
	if err := writer.Write([]string{
		"stage", "total", "success", "failed", "timeouts", "success_qps",
		"p50_ms", "p90_ms", "p95_ms", "p99_ms", "max_ms",
		"pagination_mode", "target_page", "preparation_requests", "preparation_ms",
		"sequences", "pages_traversed", "duplicate_items", "oracle_mismatches", "cursor_loops", "early_terminations",
	}); err != nil {
		return fmt.Errorf("write CSV header: %w", err)
	}
	for _, stage := range report.Stages {
		pagination := paginationObservation{}
		if stage.Pagination != nil {
			pagination = *stage.Pagination
		}
		if err := writer.Write([]string{
			stage.Name,
			strconv.FormatInt(stage.Total, 10),
			strconv.FormatInt(stage.Success, 10),
			strconv.FormatInt(stage.Failed, 10),
			strconv.FormatInt(stage.Timeouts, 10),
			fmt.Sprintf("%.3f", stage.SuccessQPS),
			fmt.Sprintf("%.3f", durationMillis(stage.Latency.P50)),
			fmt.Sprintf("%.3f", durationMillis(stage.Latency.P90)),
			fmt.Sprintf("%.3f", durationMillis(stage.Latency.P95)),
			fmt.Sprintf("%.3f", durationMillis(stage.Latency.P99)),
			fmt.Sprintf("%.3f", durationMillis(stage.Latency.Max)),
			pagination.Mode,
			strconv.Itoa(pagination.TargetPage),
			strconv.FormatInt(pagination.PreparationRequests, 10),
			fmt.Sprintf("%.3f", durationMillis(pagination.PreparationDuration)),
			strconv.FormatInt(pagination.Sequences, 10),
			strconv.FormatInt(pagination.PagesTraversed, 10),
			strconv.FormatInt(pagination.DuplicateItems, 10),
			strconv.FormatInt(pagination.OracleMismatches, 10),
			strconv.FormatInt(pagination.CursorLoops, 10),
			strconv.FormatInt(pagination.EarlyTerminations, 10),
		}); err != nil {
			return fmt.Errorf("write CSV row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("flush CSV report: %w", err)
	}
	if err := os.WriteFile(path, body.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write CSV report: %w", err)
	}
	return nil
}

type resourceMaximumRow struct {
	source string
	metric string
	value  float64
}

func summarizeFeedMetrics(samples []resourceSample, successfulReads int64) (*feedMetricsSummary, bool) {
	var snapshots []resourceSample
	for _, sample := range samples {
		if sample.Source == "prometheus:knowpost" {
			snapshots = append(snapshots, sample)
		}
	}
	if len(snapshots) < 2 {
		return nil, false
	}
	sort.Slice(snapshots, func(i, j int) bool { return snapshots[i].Timestamp.Before(snapshots[j].Timestamp) })
	first := snapshots[0].Values
	last := snapshots[len(snapshots)-1].Values
	for _, gauge := range []string{"refresh_queue_depth", "refresh_active_workers", "refresh_pending_keys"} {
		if _, ok := first[gauge]; !ok {
			return nil, false
		}
		if _, ok := last[gauge]; !ok {
			return nil, false
		}
	}
	deltas := make(map[string]float64, len(last))
	for key, lastValue := range last {
		if isNonMonotonicMetricKey(key) {
			continue
		}
		delta := lastValue - first[key]
		if delta < 0 {
			// The application restarted or a counter reset during the run. Mixing
			// values from two process lifetimes would make per-request ratios false.
			return nil, false
		}
		deltas[key] = delta
	}

	summary := &feedMetricsSummary{
		StageCalls:                     make(map[string]float64),
		StageMeanMilliseconds:          make(map[string]float64),
		StageCallsByOutcome:            make(map[string]map[string]float64),
		StageMeanMillisecondsByOutcome: make(map[string]map[string]float64),
		DependencyCalls:                make(map[string]float64),
		DependencyCallsPerSuccess:      make(map[string]float64),
		DependencyCallsByOutcome:       make(map[string]map[string]float64),
		DependencyOperations:           make(map[string]float64),
		DependencyOperationsPerSuccess: make(map[string]float64),
		DependencyOperationsByOutcome:  make(map[string]map[string]float64),
		ColdComputesByOutcome:          make(map[string]float64),
		PageCacheSources:               make(map[string]float64),
		PageCacheSourcesByOutcome:      make(map[string]map[string]float64),
		PaginationModes:                make(map[string]float64),
		PaginationModesByResult:        make(map[string]map[string]float64),
		PaginationMeanMilliseconds:     make(map[string]float64),
		CursorWork:                     make(map[string]float64),
		CursorWorkPerSuccess:           make(map[string]float64),
	}
	for _, snapshot := range snapshots {
		if value := snapshot.Values["refresh_queue_depth"]; value > summary.RefreshQueueDepthMax {
			summary.RefreshQueueDepthMax = value
		}
		if value := snapshot.Values["refresh_active_workers"]; value > summary.RefreshActiveWorkersMax {
			summary.RefreshActiveWorkersMax = value
		}
		if value := snapshot.Values["refresh_pending_keys"]; value > summary.RefreshPendingKeysMax {
			summary.RefreshPendingKeysMax = value
		}
	}
	stageDurationSums := make(map[string]map[string]float64)
	stageDurationCounts := make(map[string]map[string]float64)
	paginationDurationSums := make(map[string]map[string]float64)
	paginationDurationCounts := make(map[string]map[string]float64)
	for key, delta := range deltas {
		parts := strings.Split(key, ":")
		switch {
		case len(parts) == 3 && parts[0] == "stage_total":
			summary.StageCalls[parts[1]] += delta
			addNestedMetric(summary.StageCallsByOutcome, parts[1], parts[2], delta)
			if parts[2] != "success" && delta > 0 {
				summary.Degraded = true
			}
		case len(parts) == 3 && parts[0] == "stage_duration_sum":
			addNestedMetric(stageDurationSums, parts[1], parts[2], delta)
		case len(parts) == 3 && parts[0] == "stage_duration_count":
			addNestedMetric(stageDurationCounts, parts[1], parts[2], delta)
		case len(parts) == 4 && parts[0] == "dependency_calls":
			operation := parts[1] + "." + parts[2]
			addNestedMetric(summary.DependencyCallsByOutcome, parts[1], parts[3], delta)
			addNestedMetric(summary.DependencyOperationsByOutcome, operation, parts[3], delta)
			if parts[3] == "success" {
				summary.DependencyCalls[parts[1]] += delta
				summary.DependencyOperations[operation] += delta
			}
			if parts[3] != "success" && delta > 0 {
				summary.Degraded = true
			}
		case len(parts) == 2 && parts[0] == "cold_compute":
			summary.ColdComputesByOutcome[parts[1]] += delta
			if parts[1] == "success" {
				summary.ColdComputes += delta
			}
			if parts[1] != "success" && delta > 0 {
				summary.Degraded = true
			}
		case len(parts) == 3 && parts[0] == "page_cache":
			addNestedMetric(summary.PageCacheSourcesByOutcome, parts[1], parts[2], delta)
			if parts[2] == "success" {
				summary.PageCacheSources[parts[1]] += delta
			}
			if parts[2] != "success" && delta > 0 {
				summary.Degraded = true
			}
		case len(parts) == 3 && parts[0] == "pagination":
			addNestedMetric(summary.PaginationModesByResult, parts[1], parts[2], delta)
			if parts[2] == "success" {
				summary.PaginationModes[parts[1]] += delta
			} else if delta > 0 {
				summary.Degraded = true
			}
		case len(parts) == 3 && parts[0] == "pagination_duration_sum":
			addNestedMetric(paginationDurationSums, parts[1], parts[2], delta)
		case len(parts) == 3 && parts[0] == "pagination_duration_count":
			addNestedMetric(paginationDurationCounts, parts[1], parts[2], delta)
		case len(parts) == 2 && parts[0] == "cursor_work":
			summary.CursorWork[parts[1]] += delta
		}
	}
	pageCacheTotal := 0.0
	for _, requests := range summary.PageCacheSources {
		pageCacheTotal += requests
	}
	if pageCacheTotal > 0 {
		summary.PageCacheFreshHitRatio = (summary.PageCacheSources["l1_fresh"] + summary.PageCacheSources["l2_fresh"]) / pageCacheTotal
	}
	for stage, outcomes := range stageDurationSums {
		for outcome, sum := range outcomes {
			if count := stageDurationCounts[stage][outcome]; count > 0 {
				mean := sum / count
				addNestedMetric(summary.StageMeanMillisecondsByOutcome, stage, outcome, mean)
				if outcome == "success" {
					summary.StageMeanMilliseconds[stage] = mean
				}
			}
		}
	}
	for mode, outcomes := range paginationDurationSums {
		if count := paginationDurationCounts[mode]["success"]; count > 0 {
			summary.PaginationMeanMilliseconds[mode] = outcomes["success"] / count
		}
	}
	if successfulReads > 0 {
		denominator := float64(successfulReads)
		pageCacheSuccess := 0.0
		for _, count := range summary.PageCacheSources {
			pageCacheSuccess += count
		}
		paginationSuccess := 0.0
		for _, count := range summary.PaginationModes {
			paginationSuccess += count
		}
		if pageCacheSuccess != denominator || summary.StageCallsByOutcome["total"]["success"] != denominator ||
			paginationSuccess != denominator {
			return nil, false
		}
		for dependency, calls := range summary.DependencyCalls {
			summary.DependencyCallsPerSuccess[dependency] = calls / denominator
		}
		for operation, calls := range summary.DependencyOperations {
			summary.DependencyOperationsPerSuccess[operation] = calls / denominator
		}
		for work, value := range summary.CursorWork {
			summary.CursorWorkPerSuccess[work] = value / denominator
		}
		summary.ColdComputesPerSuccess = summary.ColdComputes / denominator
	}
	return summary, true
}

func summarizeRedisMetrics(samples []resourceSample) (*redisMetricsSummary, bool) {
	redisSamples := make([]resourceSample, 0)
	for _, sample := range samples {
		if sample.Source == "redis" {
			redisSamples = append(redisSamples, sample)
		}
	}
	if len(redisSamples) < 2 {
		return nil, false
	}
	sort.Slice(redisSamples, func(i, j int) bool { return redisSamples[i].Timestamp.Before(redisSamples[j].Timestamp) })
	first, last := redisSamples[0].Values, redisSamples[len(redisSamples)-1].Values
	if redisSamples[0].Identity == "" || redisSamples[0].Identity != redisSamples[len(redisSamples)-1].Identity {
		return nil, false
	}
	if _, ok := first["uptime_seconds"]; !ok {
		return nil, false
	}
	if _, ok := last["uptime_seconds"]; !ok || last["uptime_seconds"] < first["uptime_seconds"] {
		return nil, false
	}
	counterDelta := func(key string) (float64, bool) {
		firstValue, firstOK := first[key]
		lastValue, lastOK := last[key]
		if !firstOK || !lastOK || lastValue < firstValue {
			return 0, false
		}
		return lastValue - firstValue, true
	}
	commands, ok := counterDelta("commands_total")
	if !ok {
		return nil, false
	}
	input, ok := counterDelta("net_input_bytes")
	if !ok {
		return nil, false
	}
	output, ok := counterDelta("net_output_bytes")
	if !ok {
		return nil, false
	}
	hits, ok := counterDelta("keyspace_hits")
	if !ok {
		return nil, false
	}
	misses, ok := counterDelta("keyspace_misses")
	if !ok {
		return nil, false
	}
	evictions, ok := counterDelta("evicted_keys")
	if !ok {
		return nil, false
	}
	rejections, ok := counterDelta("rejected_connections")
	if !ok {
		return nil, false
	}
	summary := &redisMetricsSummary{
		CommandsDelta: commands, NetInputBytesDelta: input, NetOutputBytesDelta: output,
		KeyspaceHitsDelta: hits, KeyspaceMissesDelta: misses, EvictedKeysDelta: evictions,
		RejectedConnectionsDelta: rejections, SafetyEpochStart: first["feed_safety_epoch"],
		SafetyEpochEnd: last["feed_safety_epoch"],
	}
	if hits+misses > 0 {
		summary.RunHitRate = hits / (hits + misses)
	}
	for _, sample := range redisSamples {
		if sample.Values["ops_per_sec"] > summary.OpsPerSecondMax {
			summary.OpsPerSecondMax = sample.Values["ops_per_sec"]
		}
	}
	return summary, true
}

func isNonMonotonicMetricKey(key string) bool {
	switch key {
	case "refresh_queue_depth", "refresh_active_workers", "refresh_pending_keys",
		"go_memstats_heap_alloc_bytes":
		return true
	default:
		return false
	}
}

func addNestedMetric(target map[string]map[string]float64, first, second string, value float64) {
	if target[first] == nil {
		target[first] = make(map[string]float64)
	}
	target[first][second] += value
}

func resourceMaximumRows(samples []resourceSample) []resourceMaximumRow {
	maximums := make(map[string]map[string]float64)
	for _, sample := range samples {
		if sample.Source == "prometheus:knowpost" {
			// Feed counters are cumulative process-lifetime values. Their maximum
			// is misleading; the run-delta section above is the authoritative view.
			continue
		}
		if maximums[sample.Source] == nil {
			maximums[sample.Source] = make(map[string]float64)
		}
		for metric, value := range sample.Values {
			current, exists := maximums[sample.Source][metric]
			if !exists || value > current {
				maximums[sample.Source][metric] = value
			}
		}
	}
	rows := make([]resourceMaximumRow, 0)
	for source, metrics := range maximums {
		for metric, value := range metrics {
			rows = append(rows, resourceMaximumRow{source: source, metric: metric, value: value})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].source == rows[j].source {
			return rows[i].metric < rows[j].metric
		}
		return rows[i].source < rows[j].source
	})
	return rows
}

func markdownReport(report benchmarkReport) string {
	var body strings.Builder
	fmt.Fprintf(&body, "# Feed 压测报告：%s / %s / %s\n\n", report.Strategy, report.Entry, report.Scenario)
	fmt.Fprintf(&body, "- Run ID：`%s`\n", report.RunID)
	fmt.Fprintf(&body, "- 开始时间：%s\n", report.StartedAt.Format(time.RFC3339))
	fmt.Fprintf(&body, "- 采样时长：%s\n", report.Duration)
	fmt.Fprintf(&body, "- 并发：%d\n", report.Concurrency)
	if report.Trial > 0 {
		fmt.Fprintf(&body, "- 重复轮次：%d / %d\n", report.Trial, normalizedExpectedTrials(report.ExpectedTrials))
	}
	fmt.Fprintf(&body, "- 报告完整：%t\n\n", report.Complete)
	body.WriteString("| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |\n")
	body.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, stage := range report.Stages {
		fmt.Fprintf(&body, "| %s | %d | %d | %d | %d | %.2f | %.3f | %.3f | %.3f | %.3f | %.3f |\n",
			stage.Name, stage.Total, stage.Success, stage.Failed, stage.Timeouts, stage.SuccessQPS,
			durationMillis(stage.Latency.P50), durationMillis(stage.Latency.P90), durationMillis(stage.Latency.P95),
			durationMillis(stage.Latency.P99), durationMillis(stage.Latency.Max))
	}
	for _, stage := range report.Stages {
		if stage.Pagination == nil {
			continue
		}
		observation := stage.Pagination
		body.WriteString("\n## 分页正确性与准备成本\n\n")
		fmt.Fprintf(&body, "- 模式：`%s`；目标页：%d\n", observation.Mode, observation.TargetPage)
		fmt.Fprintf(&body, "- 准备请求：%d；准备耗时：%s（不计目标页 stage 延迟）\n",
			observation.PreparationRequests, observation.PreparationDuration)
		fmt.Fprintf(&body, "- 完成序列/读取页：%d/%d\n", observation.Sequences, observation.PagesTraversed)
		fmt.Fprintf(&body, "- 重复/Oracle不匹配/游标循环/提前结束：%d/%d/%d/%d\n",
			observation.DuplicateItems, observation.OracleMismatches,
			observation.CursorLoops, observation.EarlyTerminations)
		if observation.OracleHash != "" {
			fmt.Fprintf(&body, "- Oracle/Observed hash：`%s` / `%s`\n", observation.OracleHash, observation.ObservedHash)
		}
	}
	if report.FeedMetrics != nil {
		body.WriteString("\n## Feed 冷路径指标（本次运行增量）\n\n")
		body.WriteString("| dependency | calls | calls / successful read |\n")
		body.WriteString("|---|---:|---:|\n")
		dependencies := make([]string, 0, len(report.FeedMetrics.DependencyCalls))
		for dependency := range report.FeedMetrics.DependencyCalls {
			dependencies = append(dependencies, dependency)
		}
		sort.Strings(dependencies)
		for _, dependency := range dependencies {
			fmt.Fprintf(&body, "| %s | %.0f | %.3f |\n", dependency,
				report.FeedMetrics.DependencyCalls[dependency],
				report.FeedMetrics.DependencyCallsPerSuccess[dependency])
		}
		fmt.Fprintf(&body, "\n- Cold compute：%.0f（%.3f / successful read）\n",
			report.FeedMetrics.ColdComputes, report.FeedMetrics.ColdComputesPerSuccess)
		if len(report.FeedMetrics.CursorWorkPerSuccess) > 0 {
			body.WriteString("\n### 分页工作量\n\n")
			body.WriteString("| pagination work | total | per successful read |\n")
			body.WriteString("|---|---:|---:|\n")
			workKinds := make([]string, 0, len(report.FeedMetrics.CursorWorkPerSuccess))
			for work := range report.FeedMetrics.CursorWorkPerSuccess {
				workKinds = append(workKinds, work)
			}
			sort.Strings(workKinds)
			for _, work := range workKinds {
				fmt.Fprintf(&body, "| %s | %.0f | %.3f |\n", work,
					report.FeedMetrics.CursorWork[work], report.FeedMetrics.CursorWorkPerSuccess[work])
			}
		}
		if len(report.FeedMetrics.PageCacheSources) > 0 {
			body.WriteString("\n| page cache source | requests |\n")
			body.WriteString("|---|---:|\n")
			sources := make([]string, 0, len(report.FeedMetrics.PageCacheSources))
			for source := range report.FeedMetrics.PageCacheSources {
				sources = append(sources, source)
			}
			sort.Strings(sources)
			for _, source := range sources {
				fmt.Fprintf(&body, "| %s | %.0f |\n", source, report.FeedMetrics.PageCacheSources[source])
			}
			fmt.Fprintf(&body, "\n- L1+L2 Fresh ratio：%.2f%%\n", report.FeedMetrics.PageCacheFreshHitRatio*100)
		}
		fmt.Fprintf(&body, "- Refresh max：queue=%.0f active=%.0f pending=%.0f\n",
			report.FeedMetrics.RefreshQueueDepthMax,
			report.FeedMetrics.RefreshActiveWorkersMax,
			report.FeedMetrics.RefreshPendingKeysMax)
		stages := make([]string, 0, len(report.FeedMetrics.StageMeanMilliseconds))
		for stage := range report.FeedMetrics.StageMeanMilliseconds {
			stages = append(stages, stage)
		}
		sort.Strings(stages)
		if len(stages) > 0 {
			body.WriteString("\n| feed stage | calls(all outcomes) | success mean(ms) |\n")
			body.WriteString("|---|---:|---:|\n")
			for _, stage := range stages {
				fmt.Fprintf(&body, "| %s | %.0f | %.3f |\n", stage,
					report.FeedMetrics.StageCalls[stage], report.FeedMetrics.StageMeanMilliseconds[stage])
			}
		}
		if report.FeedMetrics.Degraded {
			body.WriteString("\n### 非成功 outcome（本轮已降级）\n\n")
			body.WriteString("| kind | name | outcome | calls | mean(ms) |\n")
			body.WriteString("|---|---|---|---:|---:|\n")
			for _, row := range feedOutcomeRows(report.FeedMetrics) {
				fmt.Fprintf(&body, "| %s | %s | %s | %.0f | %.3f |\n",
					row.kind, row.name, row.outcome, row.calls, row.meanMilliseconds)
			}
		}
	}
	if report.RedisMetrics != nil {
		body.WriteString("\n## Redis 本轮边界增量\n\n")
		fmt.Fprintf(&body, "- Commands：%.0f；input：%.0f bytes；output：%.0f bytes\n",
			report.RedisMetrics.CommandsDelta, report.RedisMetrics.NetInputBytesDelta, report.RedisMetrics.NetOutputBytesDelta)
		fmt.Fprintf(&body, "- Hits/Misses：%.0f/%.0f；run hit rate：%.2f%%\n",
			report.RedisMetrics.KeyspaceHitsDelta, report.RedisMetrics.KeyspaceMissesDelta, report.RedisMetrics.RunHitRate*100)
		fmt.Fprintf(&body, "- Evicted/Rejected：%.0f/%.0f；ops/s max：%.0f；safety epoch：%.0f -> %.0f\n",
			report.RedisMetrics.EvictedKeysDelta, report.RedisMetrics.RejectedConnectionsDelta,
			report.RedisMetrics.OpsPerSecondMax, report.RedisMetrics.SafetyEpochStart, report.RedisMetrics.SafetyEpochEnd)
	}
	if len(report.Resources) > 0 {
		body.WriteString("\n## 资源采样峰值\n\n")
		body.WriteString("| source | metric | max |\n")
		body.WriteString("|---|---|---:|\n")
		for _, row := range resourceMaximumRows(report.Resources) {
			fmt.Fprintf(&body, "| %s | %s | %.3f |\n", row.source, row.metric, row.value)
		}
	}
	if report.Recovery != nil {
		body.WriteString("\n## 停止施压后的恢复\n\n")
		fmt.Fprintf(&body, "- Kafka drain：%s\n", report.Recovery.KafkaDrain)
		fmt.Fprintf(&body, "- 完成：%t\n", report.Recovery.Complete)
		if lag, ok := report.Recovery.FinalMetrics["lag_total"]; ok {
			fmt.Fprintf(&body, "- 最终 lag：%.0f\n", lag)
		}
		if report.Recovery.Error != "" {
			fmt.Fprintf(&body, "- 错误：%s\n", report.Recovery.Error)
		}
	}
	if report.Mutation != nil {
		body.WriteString("\n## Mutation checkpoint\n\n")
		fmt.Fprintf(&body, "- Checkpoint：`%s`；baseline：`%s`\n",
			report.Mutation.CheckpointID, report.Mutation.BaselineFingerprint)
		fmt.Fprintf(&body, "- 创建时间：%s；工具：`%s`\n",
			report.Mutation.CheckpointCreatedAt.Format(time.RFC3339Nano), report.Mutation.ToolVersion)
		fmt.Fprintf(&body, "- 成功发布帖子：%d\n", len(report.Mutation.PublishedPostIDs))
		writeMutationRestoreMarkdown(&body, "测量前恢复", report.Mutation.InitialRestore)
		writeMutationRestoreMarkdown(&body, "预热后恢复", report.Mutation.AfterWarmupRestore)
		writeMutationRestoreMarkdown(&body, "测量后恢复", report.Mutation.AfterTrialRestore)
	}
	if len(report.MissingMetrics) > 0 {
		body.WriteString("\n## 缺失指标\n\n")
		for _, item := range report.MissingMetrics {
			fmt.Fprintf(&body, "- %s\n", item)
		}
	}
	if len(report.Notes) > 0 {
		body.WriteString("\n## 说明\n\n")
		for _, item := range report.Notes {
			fmt.Fprintf(&body, "- %s\n", item)
		}
	}
	return body.String()
}

func writeMutationRestoreMarkdown(body *strings.Builder, label string, observation *mutationRestoreObservation) {
	if observation == nil {
		return
	}
	fmt.Fprintf(body, "- %s：complete=%t；耗时=%s；删除帖子/Outbox=%d/%d；safety epoch=%d\n",
		label, observation.Complete, observation.Duration,
		observation.RemovedPosts, observation.RemovedOutbox, observation.SafetyEpoch)
	if observation.Error != "" {
		fmt.Fprintf(body, "  - 错误：%s\n", observation.Error)
	}
}

type feedOutcomeRow struct {
	kind             string
	name             string
	outcome          string
	calls            float64
	meanMilliseconds float64
}

func feedOutcomeRows(summary *feedMetricsSummary) []feedOutcomeRow {
	if summary == nil {
		return nil
	}
	rows := make([]feedOutcomeRow, 0)
	for stage, outcomes := range summary.StageCallsByOutcome {
		for outcome, calls := range outcomes {
			if outcome == "success" || calls <= 0 {
				continue
			}
			rows = append(rows, feedOutcomeRow{
				kind: "stage", name: stage, outcome: outcome, calls: calls,
				meanMilliseconds: summary.StageMeanMillisecondsByOutcome[stage][outcome],
			})
		}
	}
	for operation, outcomes := range summary.DependencyOperationsByOutcome {
		for outcome, calls := range outcomes {
			if outcome == "success" || calls <= 0 {
				continue
			}
			rows = append(rows, feedOutcomeRow{kind: "dependency", name: operation, outcome: outcome, calls: calls})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		left := strings.Join([]string{rows[i].kind, rows[i].name, rows[i].outcome}, "\x00")
		right := strings.Join([]string{rows[j].kind, rows[j].name, rows[j].outcome}, "\x00")
		return left < right
	})
	return rows
}

func durationMillis(value time.Duration) float64 {
	return float64(value) / float64(time.Millisecond)
}

func safePathSegment(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "unknown"
	}
	replacer := strings.NewReplacer("/", "_", "\\", "_", ":", "_", "..", "_")
	return replacer.Replace(value)
}
