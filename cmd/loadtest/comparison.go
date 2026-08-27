package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type comparisonRow struct {
	Strategy          string        `json:"strategy"`
	Entry             string        `json:"entry"`
	Scenario          string        `json:"scenario"`
	ReaderCardinality string        `json:"reader_cardinality"`
	CacheState        string        `json:"cache_state"`
	Topology          string        `json:"topology"`
	Stage             string        `json:"stage"`
	Concurrency       int           `json:"concurrency"`
	Trial             int           `json:"trial,omitempty"`
	ExpectedTrials    int           `json:"expected_trials,omitempty"`
	Total             int64         `json:"total"`
	Success           int64         `json:"success"`
	Failed            int64         `json:"failed"`
	Timeouts          int64         `json:"timeouts"`
	SuccessQPS        float64       `json:"success_qps"`
	P90               time.Duration `json:"p90_ns"`
	P95               time.Duration `json:"p95_ns"`
	P99               time.Duration `json:"p99_ns"`
	Max               time.Duration `json:"max_ns"`
	Complete          bool          `json:"complete"`
	MaxKafkaLag       float64       `json:"max_kafka_lag"`
	MaxKnowCPU        float64       `json:"max_knowpost_cpu_percent"`
	MaxClientCPU      float64       `json:"max_client_cpu_percent_normalized"`
	MaxRedisOps       float64       `json:"max_redis_ops_per_sec"`
	MaxMySQLRuns      float64       `json:"max_mysql_threads_running"`
}

type capacitySummary struct {
	Strategy                string        `json:"strategy"`
	Entry                   string        `json:"entry"`
	Scenario                string        `json:"scenario"`
	ReaderCardinality       string        `json:"reader_cardinality"`
	CacheState              string        `json:"cache_state"`
	Topology                string        `json:"topology"`
	Stage                   string        `json:"stage"`
	PeakSuccessQPS          float64       `json:"peak_success_qps"`
	PeakP95                 time.Duration `json:"peak_p95_ns"`
	StableConcurrency       int           `json:"stable_concurrency"`
	KneeConcurrency         int           `json:"knee_concurrency"`
	NoErrorUpperConcurrency int           `json:"no_error_upper_concurrency"`
	FirstErrorConcurrency   int           `json:"first_error_concurrency"`
}

type comparisonReport struct {
	RunID      string            `json:"run_id"`
	Rows       []comparisonRow   `json:"rows"`
	Capacities []capacitySummary `json:"capacities"`
	Trials     []trialSummary    `json:"trial_summaries"`
}

type trialSummary struct {
	Strategy          string        `json:"strategy"`
	Entry             string        `json:"entry"`
	Scenario          string        `json:"scenario"`
	ReaderCardinality string        `json:"reader_cardinality"`
	CacheState        string        `json:"cache_state"`
	Topology          string        `json:"topology"`
	Stage             string        `json:"stage"`
	Concurrency       int           `json:"concurrency"`
	Runs              int           `json:"runs"`
	CompleteRuns      int           `json:"complete_runs"`
	QPSMedian         float64       `json:"qps_median"`
	QPSMin            float64       `json:"qps_min"`
	QPSMax            float64       `json:"qps_max"`
	QPSSpreadPercent  float64       `json:"qps_spread_percent"`
	P95Median         time.Duration `json:"p95_median_ns"`
	P95Min            time.Duration `json:"p95_min_ns"`
	P95Max            time.Duration `json:"p95_max_ns"`
}

type comparisonPaths struct {
	JSON     string
	CSV      string
	Markdown string
}

func generateComparison(root, runID string) (comparisonPaths, error) {
	rows, err := loadComparisonRows(root, runID)
	if err != nil {
		return comparisonPaths{}, err
	}
	if len(rows) == 0 {
		return comparisonPaths{}, fmt.Errorf("no benchmark reports found for run %s", runID)
	}

	report := comparisonReport{
		RunID:      runID,
		Rows:       rows,
		Capacities: summarizeCapacities(rows),
		Trials:     summarizeTrials(rows),
	}
	dir := filepath.Join(root, safePathSegment(runID))
	paths := comparisonPaths{
		JSON:     filepath.Join(dir, "comparison.json"),
		CSV:      filepath.Join(dir, "comparison.csv"),
		Markdown: filepath.Join(dir, "comparison.md"),
	}
	jsonBody, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return comparisonPaths{}, fmt.Errorf("marshal comparison JSON: %w", err)
	}
	if err := os.WriteFile(paths.JSON, append(jsonBody, '\n'), 0o644); err != nil {
		return comparisonPaths{}, fmt.Errorf("write comparison JSON: %w", err)
	}
	if err := writeComparisonCSV(paths.CSV, rows); err != nil {
		return comparisonPaths{}, err
	}
	if err := os.WriteFile(paths.Markdown, []byte(comparisonMarkdown(runID, rows)), 0o644); err != nil {
		return comparisonPaths{}, fmt.Errorf("write comparison Markdown: %w", err)
	}
	return paths, nil
}

func loadComparisonRows(root, runID string) ([]comparisonRow, error) {
	dir := filepath.Join(root, safePathSegment(runID))
	rows := make([]comparisonRow, 0)
	executionTopology := ""
	compatibilityByGroup := make(map[string]string)
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() != "report.json" {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read report %s: %w", path, err)
		}
		var report benchmarkReport
		if err := json.Unmarshal(body, &report); err != nil {
			return fmt.Errorf("decode report %s: %w", path, err)
		}
		if report.RunID != runID {
			return nil
		}
		topology := strings.TrimSpace(report.Environment["topology"])
		if topology == "" {
			topology = "legacy-unknown"
		}
		if executionTopology == "" {
			executionTopology = topology
		} else if executionTopology != topology {
			return fmt.Errorf("comparison topology mismatch: %q versus %q", executionTopology, topology)
		}
		compatibility := strings.TrimSpace(report.Environment["compatibility_fingerprint"])
		if compatibility == "" {
			compatibility = "legacy-unknown"
		}
		compatibilityGroup := strings.Join([]string{
			report.Strategy, report.Entry, baseScenario(report.Scenario),
			report.Environment["reader_cardinality"], report.Environment["cache_state"], topology,
		}, "\x00")
		if previous := compatibilityByGroup[compatibilityGroup]; previous == "" {
			compatibilityByGroup[compatibilityGroup] = compatibility
		} else if previous != compatibility {
			return fmt.Errorf("comparison evidence mismatch for %s: %q versus %q", compatibilityGroup, previous, compatibility)
		}
		resources := comparisonResourceMaximums(report.Resources)
		knowPostCPU := resources["docker:zg-knowpost/cpu_percent"]
		if topology == "compose-middleware-local-services" {
			knowPostCPU = resources["process:knowpost/cpu_percent"]
		}
		for _, stage := range report.Stages {
			if stage.Name != "read" && stage.Name != "publish_total" {
				continue
			}
			rows = append(rows, comparisonRow{
				Strategy: report.Strategy, Entry: report.Entry,
				Scenario: baseScenario(report.Scenario), ReaderCardinality: report.Environment["reader_cardinality"],
				CacheState: report.Environment["cache_state"], Topology: topology, Stage: stage.Name,
				Concurrency: report.Concurrency, Trial: report.Trial, ExpectedTrials: normalizedExpectedTrials(report.ExpectedTrials), Total: stage.Total, Success: stage.Success,
				Failed: stage.Failed, Timeouts: stage.Timeouts, SuccessQPS: stage.SuccessQPS,
				P90: stage.Latency.P90, P95: stage.Latency.P95, P99: stage.Latency.P99, Max: stage.Latency.Max, Complete: report.Complete,
				MaxKafkaLag: resources["kafka/lag_max"], MaxKnowCPU: knowPostCPU,
				MaxClientCPU: resources["client:loadtest/cpu_percent_normalized"],
				MaxRedisOps:  resources["redis/ops_per_sec"], MaxMySQLRuns: resources["mysql/threads_running"],
			})
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("load comparison reports: %w", err)
	}
	sortComparisonRows(rows)
	return rows, nil
}

func comparisonResourceMaximums(samples []resourceSample) map[string]float64 {
	maximums := make(map[string]float64)
	for _, sample := range samples {
		for metric, value := range sample.Values {
			key := sample.Source + "/" + metric
			if current, ok := maximums[key]; !ok || value > current {
				maximums[key] = value
			}
		}
	}
	return maximums
}

func baseScenario(value string) string {
	lastDash := strings.LastIndex(value, "-c")
	if lastDash < 0 || lastDash+2 >= len(value) {
		return value
	}
	if _, err := strconv.Atoi(value[lastDash+2:]); err != nil {
		return value
	}
	return value[:lastDash]
}

func findCapacity(rows []comparisonRow) capacitySummary {
	if len(rows) == 0 {
		return capacitySummary{}
	}
	ordered := append([]comparisonRow(nil), rows...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Concurrency < ordered[j].Concurrency })
	result := capacitySummary{
		Strategy: ordered[0].Strategy, Entry: ordered[0].Entry,
		Scenario: ordered[0].Scenario, ReaderCardinality: ordered[0].ReaderCardinality,
		CacheState: ordered[0].CacheState, Topology: ordered[0].Topology, Stage: ordered[0].Stage,
	}
	noError := make([]comparisonRow, 0, len(ordered))
	for _, row := range ordered {
		if row.Failed > 0 {
			if result.FirstErrorConcurrency == 0 {
				result.FirstErrorConcurrency = row.Concurrency
			}
			break
		}
		if !row.Complete {
			break
		}
		noError = append(noError, row)
		if row.Concurrency > result.NoErrorUpperConcurrency {
			result.NoErrorUpperConcurrency = row.Concurrency
		}
		if row.SuccessQPS > result.PeakSuccessQPS {
			result.PeakSuccessQPS = row.SuccessQPS
			result.PeakP95 = row.P95
		}
	}
	if len(noError) == 0 {
		result.KneeConcurrency = result.FirstErrorConcurrency
		return result
	}
	for index := 1; index < len(noError); index++ {
		previous, current := noError[index-1], noError[index]
		gain := 0.0
		if previous.SuccessQPS > 0 {
			gain = (current.SuccessQPS - previous.SuccessQPS) / previous.SuccessQPS
		}
		if gain < 0.10 && current.P95 > previous.P95 {
			result.StableConcurrency = previous.Concurrency
			result.KneeConcurrency = current.Concurrency
			break
		}
	}
	if result.StableConcurrency == 0 {
		result.StableConcurrency = result.NoErrorUpperConcurrency
	}
	if result.KneeConcurrency == 0 && result.FirstErrorConcurrency > 0 {
		result.KneeConcurrency = result.FirstErrorConcurrency
	}
	return result
}

func summarizeCapacities(rows []comparisonRow) []capacitySummary {
	rows = collapseTrialRows(rows)
	groups := make(map[string][]comparisonRow)
	for _, row := range rows {
		key := strings.Join([]string{row.Strategy, row.Entry, row.Scenario, row.ReaderCardinality, row.CacheState, row.Topology, row.Stage}, "\x00")
		groups[key] = append(groups[key], row)
	}
	result := make([]capacitySummary, 0, len(groups))
	for _, group := range groups {
		result = append(result, findCapacity(group))
	}
	sort.Slice(result, func(i, j int) bool {
		left := strings.Join([]string{result[i].Strategy, result[i].Entry, result[i].Scenario, result[i].ReaderCardinality, result[i].CacheState, result[i].Topology, result[i].Stage}, "\x00")
		right := strings.Join([]string{result[j].Strategy, result[j].Entry, result[j].Scenario, result[j].ReaderCardinality, result[j].CacheState, result[j].Topology, result[j].Stage}, "\x00")
		return left < right
	})
	return result
}

func collapseTrialRows(rows []comparisonRow) []comparisonRow {
	groups := make(map[string][]comparisonRow)
	for _, row := range rows {
		key := strings.Join([]string{
			row.Strategy, row.Entry, row.Scenario, row.ReaderCardinality, row.CacheState, row.Topology, row.Stage, strconv.Itoa(row.Concurrency),
		}, "\x00")
		groups[key] = append(groups[key], row)
	}
	result := make([]comparisonRow, 0, len(groups))
	for _, group := range groups {
		collapsed := group[0]
		collapsed.Trial = 0
		collapsed.Complete = true
		qps := make([]float64, 0, len(group))
		p95 := make([]time.Duration, 0, len(group))
		p90 := make([]time.Duration, 0, len(group))
		p99 := make([]time.Duration, 0, len(group))
		maxLatency := make([]time.Duration, 0, len(group))
		knowCPU := make([]float64, 0, len(group))
		clientCPU := make([]float64, 0, len(group))
		redisOps := make([]float64, 0, len(group))
		mysqlRuns := make([]float64, 0, len(group))
		kafkaLag := make([]float64, 0, len(group))
		expectedTrials := normalizedExpectedTrials(group[0].ExpectedTrials)
		trials := make(map[int]struct{}, len(group))
		for _, row := range group {
			qps = append(qps, row.SuccessQPS)
			p90 = append(p90, row.P90)
			p95 = append(p95, row.P95)
			p99 = append(p99, row.P99)
			maxLatency = append(maxLatency, row.Max)
			knowCPU = append(knowCPU, row.MaxKnowCPU)
			clientCPU = append(clientCPU, row.MaxClientCPU)
			redisOps = append(redisOps, row.MaxRedisOps)
			mysqlRuns = append(mysqlRuns, row.MaxMySQLRuns)
			kafkaLag = append(kafkaLag, row.MaxKafkaLag)
			collapsed.Failed += row.Failed
			collapsed.Timeouts += row.Timeouts
			if !row.Complete {
				collapsed.Complete = false
			}
			if normalizedExpectedTrials(row.ExpectedTrials) != expectedTrials || row.Trial <= 0 {
				collapsed.Complete = false
			}
			if _, exists := trials[row.Trial]; exists {
				collapsed.Complete = false
			}
			trials[row.Trial] = struct{}{}
		}
		if len(group) != expectedTrials || len(trials) != expectedTrials {
			collapsed.Complete = false
		}
		for trial := 1; trial <= expectedTrials; trial++ {
			if _, exists := trials[trial]; !exists {
				collapsed.Complete = false
			}
		}
		// group[0] was already included above; reset cumulative fields to avoid
		// double-counting it while retaining strict any-trial failure semantics.
		collapsed.Failed -= group[0].Failed
		collapsed.Timeouts -= group[0].Timeouts
		collapsed.SuccessQPS = medianFloat64(qps)
		collapsed.P90 = medianDuration(p90)
		collapsed.P95 = medianDuration(p95)
		collapsed.P99 = medianDuration(p99)
		collapsed.Max = medianDuration(maxLatency)
		collapsed.MaxKnowCPU = medianFloat64(knowCPU)
		collapsed.MaxClientCPU = medianFloat64(clientCPU)
		collapsed.MaxRedisOps = medianFloat64(redisOps)
		collapsed.MaxMySQLRuns = medianFloat64(mysqlRuns)
		collapsed.MaxKafkaLag = medianFloat64(kafkaLag)
		result = append(result, collapsed)
	}
	sortComparisonRows(result)
	return result
}

func medianFloat64(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	ordered := append([]float64(nil), values...)
	sort.Float64s(ordered)
	middle := len(ordered) / 2
	if len(ordered)%2 == 1 {
		return ordered[middle]
	}
	return (ordered[middle-1] + ordered[middle]) / 2
}

func medianDuration(values []time.Duration) time.Duration {
	if len(values) == 0 {
		return 0
	}
	ordered := append([]time.Duration(nil), values...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i] < ordered[j] })
	middle := len(ordered) / 2
	if len(ordered)%2 == 1 {
		return ordered[middle]
	}
	return ordered[middle-1] + (ordered[middle]-ordered[middle-1])/2
}

func comparisonMarkdown(runID string, rows []comparisonRow) string {
	var body strings.Builder
	fmt.Fprintf(&body, "# Feed 三策略综合对比：%s\n\n", runID)
	body.WriteString("容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。\n\n")
	body.WriteString("| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |\n")
	body.WriteString("|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|\n")
	for _, capacity := range summarizeCapacities(rows) {
		fmt.Fprintf(&body, "| %s | %s | %s | %s | %s | %s | %s | %.2f | %.3f | %s | %s | %s | %s |\n",
			capacity.Strategy, capacity.Entry, capacity.Scenario,
			capacity.ReaderCardinality, capacity.CacheState, capacity.Topology, capacity.Stage,
			capacity.PeakSuccessQPS, durationMillis(capacity.PeakP95),
			optionalInt(capacity.StableConcurrency), optionalInt(capacity.KneeConcurrency),
			optionalInt(capacity.NoErrorUpperConcurrency), optionalInt(capacity.FirstErrorConcurrency))
	}
	body.WriteString("\n## 重复轮次中位数与波动\n\n")
	body.WriteString("| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |\n")
	body.WriteString("|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, summary := range summarizeTrials(rows) {
		fmt.Fprintf(&body, "| %s | %s | %s | %s | %s | %s | %s | %d | %d/%d | %.2f | %.2f | %.2f | %.2f%% | %.3f | %.3f | %.3f |\n",
			summary.Strategy, summary.Entry, summary.Scenario, summary.ReaderCardinality, summary.CacheState, summary.Topology, summary.Stage, summary.Concurrency,
			summary.CompleteRuns, summary.Runs, summary.QPSMedian, summary.QPSMin, summary.QPSMax,
			summary.QPSSpreadPercent, durationMillis(summary.P95Median), durationMillis(summary.P95Min),
			durationMillis(summary.P95Max))
	}
	body.WriteString("\n## 原始指标\n\n")
	body.WriteString("| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |\n")
	body.WriteString("|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|\n")
	ordered := append([]comparisonRow(nil), rows...)
	sortComparisonRows(ordered)
	for _, row := range ordered {
		fmt.Fprintf(&body, "| %s | %s | %s | %s | %s | %s | %s | %d | %d/%d | %d/%d | %d | %d | %.2f | %.3f | %.3f | %.3f | %.3f | %t | %.0f | %.2f | %.2f | %.0f | %.0f |\n",
			row.Strategy, row.Entry, row.Scenario, row.ReaderCardinality, row.CacheState, row.Topology, row.Stage, row.Concurrency, row.Trial, normalizedExpectedTrials(row.ExpectedTrials),
			row.Success, row.Total, row.Failed, row.Timeouts, row.SuccessQPS,
			durationMillis(row.P90), durationMillis(row.P95), durationMillis(row.P99), durationMillis(row.Max), row.Complete,
			row.MaxKafkaLag, row.MaxKnowCPU, row.MaxClientCPU, row.MaxRedisOps, row.MaxMySQLRuns)
	}
	return body.String()
}

func summarizeTrials(rows []comparisonRow) []trialSummary {
	groups := make(map[string][]comparisonRow)
	for _, row := range rows {
		key := strings.Join([]string{
			row.Strategy, row.Entry, row.Scenario, row.ReaderCardinality, row.CacheState, row.Topology, row.Stage, strconv.Itoa(row.Concurrency),
		}, "\x00")
		groups[key] = append(groups[key], row)
	}
	result := make([]trialSummary, 0, len(groups))
	for _, group := range groups {
		qps := make([]float64, 0, len(group))
		p95 := make([]time.Duration, 0, len(group))
		complete := 0
		for _, row := range group {
			qps = append(qps, row.SuccessQPS)
			p95 = append(p95, row.P95)
			if row.Complete {
				complete++
			}
		}
		sort.Float64s(qps)
		sort.Slice(p95, func(i, j int) bool { return p95[i] < p95[j] })
		median := medianFloat64(qps)
		spread := 0.0
		if median > 0 {
			spread = (qps[len(qps)-1] - qps[0]) / median * 100
		}
		result = append(result, trialSummary{
			Strategy: group[0].Strategy, Entry: group[0].Entry, Scenario: group[0].Scenario,
			ReaderCardinality: group[0].ReaderCardinality, CacheState: group[0].CacheState, Topology: group[0].Topology,
			Stage: group[0].Stage, Concurrency: group[0].Concurrency, Runs: len(group), CompleteRuns: complete,
			QPSMedian: median, QPSMin: qps[0], QPSMax: qps[len(qps)-1], QPSSpreadPercent: spread,
			P95Median: medianDuration(p95), P95Min: p95[0], P95Max: p95[len(p95)-1],
		})
	}
	sort.Slice(result, func(i, j int) bool {
		left := strings.Join([]string{result[i].Strategy, result[i].Entry, result[i].Scenario, result[i].ReaderCardinality, result[i].CacheState, result[i].Topology, result[i].Stage}, "\x00")
		right := strings.Join([]string{result[j].Strategy, result[j].Entry, result[j].Scenario, result[j].ReaderCardinality, result[j].CacheState, result[j].Topology, result[j].Stage}, "\x00")
		if left == right {
			return result[i].Concurrency < result[j].Concurrency
		}
		return left < right
	})
	return result
}

func writeComparisonCSV(path string, rows []comparisonRow) error {
	var body bytes.Buffer
	writer := csv.NewWriter(&body)
	if err := writer.Write([]string{
		"strategy", "entry", "scenario", "reader_cardinality", "cache_state", "topology", "stage", "concurrency", "trial", "expected_trials", "total", "success", "failed", "timeouts",
		"success_qps", "p90_ms", "p95_ms", "p99_ms", "max_ms", "complete", "max_kafka_lag", "max_knowpost_cpu_percent",
		"max_client_cpu_percent_normalized", "max_redis_ops_per_sec", "max_mysql_threads_running",
	}); err != nil {
		return fmt.Errorf("write comparison CSV header: %w", err)
	}
	for _, row := range rows {
		if err := writer.Write([]string{
			row.Strategy, row.Entry, row.Scenario, row.ReaderCardinality, row.CacheState, row.Topology, row.Stage, strconv.Itoa(row.Concurrency), strconv.Itoa(row.Trial), strconv.Itoa(normalizedExpectedTrials(row.ExpectedTrials)),
			strconv.FormatInt(row.Total, 10), strconv.FormatInt(row.Success, 10), strconv.FormatInt(row.Failed, 10), strconv.FormatInt(row.Timeouts, 10),
			fmt.Sprintf("%.3f", row.SuccessQPS), fmt.Sprintf("%.3f", durationMillis(row.P90)), fmt.Sprintf("%.3f", durationMillis(row.P95)), fmt.Sprintf("%.3f", durationMillis(row.P99)), fmt.Sprintf("%.3f", durationMillis(row.Max)),
			strconv.FormatBool(row.Complete), fmt.Sprintf("%.3f", row.MaxKafkaLag), fmt.Sprintf("%.3f", row.MaxKnowCPU), fmt.Sprintf("%.3f", row.MaxClientCPU),
			fmt.Sprintf("%.3f", row.MaxRedisOps), fmt.Sprintf("%.3f", row.MaxMySQLRuns),
		}); err != nil {
			return fmt.Errorf("write comparison CSV row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("flush comparison CSV: %w", err)
	}
	if err := os.WriteFile(path, body.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write comparison CSV: %w", err)
	}
	return nil
}

func sortComparisonRows(rows []comparisonRow) {
	sort.Slice(rows, func(i, j int) bool {
		left := strings.Join([]string{rows[i].Strategy, rows[i].Entry, rows[i].Scenario, rows[i].ReaderCardinality, rows[i].CacheState, rows[i].Topology, rows[i].Stage}, "\x00")
		right := strings.Join([]string{rows[j].Strategy, rows[j].Entry, rows[j].Scenario, rows[j].ReaderCardinality, rows[j].CacheState, rows[j].Topology, rows[j].Stage}, "\x00")
		if left == right {
			if rows[i].Concurrency == rows[j].Concurrency {
				return rows[i].Trial < rows[j].Trial
			}
			return rows[i].Concurrency < rows[j].Concurrency
		}
		return left < right
	})
}

func optionalInt(value int) string {
	if value == 0 {
		return "-"
	}
	return strconv.Itoa(value)
}
