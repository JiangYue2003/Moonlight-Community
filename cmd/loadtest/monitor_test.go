package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCollectKnowPostMetricsUsesDirectHTTPForLocalTopology(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, `zhiguang_knowpost_feed_cold_compute_total{outcome="success"} 7`)
		fmt.Fprintln(w, "process_start_time_seconds 123")
	}))
	defer server.Close()

	values, err := collectKnowPostMetrics(t.Context(), monitorConfig{PrometheusURL: server.URL})

	require.NoError(t, err)
	require.Equal(t, 7.0, values["cold_compute:success"])
	require.Equal(t, 123.0, values["process_start_time_seconds"])
}

func TestParseDockerStats(t *testing.T) {
	raw := []byte("{\"Name\":\"zg-knowpost\",\"CPUPerc\":\"12.50%\",\"MemPerc\":\"3.25%\",\"PIDs\":\"17\"}\n")

	samples, err := parseDockerStats(raw, time.Unix(1, 0))

	require.NoError(t, err)
	require.Len(t, samples, 1)
	require.Equal(t, "docker:zg-knowpost", samples[0].Source)
	require.Equal(t, 12.5, samples[0].Values["cpu_percent"])
	require.Equal(t, 3.25, samples[0].Values["memory_percent"])
	require.Equal(t, 17.0, samples[0].Values["pids"])
}

func TestParseDockerStatePersistsHealthAndRestartCount(t *testing.T) {
	raw := []byte(`[{"Name":"/zg-knowpost","RestartCount":0,"State":{"Running":true,"Health":{"Status":"healthy"}}}]`)

	samples, states, err := parseDockerState(raw, time.Unix(1, 0))

	require.NoError(t, err)
	require.Len(t, samples, 1)
	require.Equal(t, "docker-state:zg-knowpost", samples[0].Source)
	require.Equal(t, 1.0, samples[0].Values["running"])
	require.Equal(t, 1.0, samples[0].Values["healthy"])
	require.Equal(t, 0.0, samples[0].Values["restart_count"])
	require.Equal(t, "healthy", states[0].Health)
}

func TestParseKafkaLagSumsPartitions(t *testing.T) {
	raw := []byte(`GROUP TOPIC PARTITION CURRENT-OFFSET LOG-END-OFFSET LAG CONSUMER-ID HOST CLIENT-ID
feed-fanout-group feed-fanout 0 100 106 6 - - -
feed-fanout-group feed-fanout 1 200 203 3 - - -`)

	values, err := parseKafkaLag(raw)

	require.NoError(t, err)
	require.Equal(t, 9.0, values["lag_total"])
	require.Equal(t, 6.0, values["lag_max"])
	require.Equal(t, 300.0, values["current_offset_total"])
	require.Equal(t, 309.0, values["log_end_offset_total"])
}

func TestParseKafkaLagHandlesUncommittedPartition(t *testing.T) {
	tests := []struct {
		name   string
		logEnd string
		lag    float64
	}{
		{name: "empty topic", logEnd: "0", lag: 0},
		{name: "existing messages", logEnd: "7", lag: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := []byte("GROUP TOPIC PARTITION CURRENT-OFFSET LOG-END-OFFSET LAG CONSUMER-ID HOST CLIENT-ID\n" +
				"feed-fanout-group feed-fanout 0 - " + tt.logEnd + " - consumer host client")

			values, err := parseKafkaLag(raw)

			require.NoError(t, err)
			require.Equal(t, tt.lag, values["lag_total"])
			require.Equal(t, 0.0, values["current_offset_total"])
			require.Equal(t, 1.0, values["partitions"])
		})
	}
}

func TestParseRedisInfoIncludesKeyCountAndHitRate(t *testing.T) {
	raw := []byte("used_memory:1024\r\ntotal_commands_processed:100\r\nkeyspace_hits:80\r\nkeyspace_misses:20\r\ndb0:keys=12,expires=2,avg_ttl=100\r\n")

	values, err := parseRedisInfo(raw)

	require.NoError(t, err)
	require.Equal(t, 12.0, values["keys"])
	require.Equal(t, 0.8, values["hit_rate"])
}

func TestReportCompleteRequiresRequestsAndKafkaMetrics(t *testing.T) {
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Failed: 1}}, nil))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"kafka"}))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"feed_metrics"}))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"client_cpu"}))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"client_cpu_saturated"}))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"local_process"}))
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"local_process_changed"}))
	for _, required := range []string{"docker", "docker_state", "redis", "mysql"} {
		require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{required}))
	}
	for _, required := range []string{"kafka_recovery", "feed_runtime_flags", "redis_evictions", "redis_rejections"} {
		require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{required}))
	}
	require.False(t, reportIsComplete([]stageResult{{Total: 1, Success: 1}}, []string{"redis_safety_epoch_changed"}))
}

func TestBenchmarkReportCompletionErrorStopsMatrixAfterIncompleteTrial(t *testing.T) {
	require.NoError(t, benchmarkReportCompletionError(benchmarkReport{Complete: true}))

	err := benchmarkReportCompletionError(benchmarkReport{
		Complete: false, MissingMetrics: []string{"client_cpu_saturated", "mutation_restore"},
	})
	require.ErrorContains(t, err, "client_cpu_saturated")
	require.ErrorContains(t, err, "mutation_restore")
}

func TestReportCompleteRejectsCursorWithoutOracleOrWithCorrectnessFailure(t *testing.T) {
	base := stageResult{Total: 1, Success: 1, Pagination: &paginationObservation{
		Mode: "cursor", TargetPage: 20, OracleHash: "oracle-a", ObservedHash: "oracle-a",
	}}
	require.True(t, reportIsComplete([]stageResult{base}, nil))

	missingOracle := base
	missingOracle.Pagination = &paginationObservation{Mode: "cursor", TargetPage: 20}
	require.False(t, reportIsComplete([]stageResult{missingOracle}, nil))

	duplicate := base
	copyObservation := *base.Pagination
	copyObservation.DuplicateItems = 1
	duplicate.Pagination = &copyObservation
	require.False(t, reportIsComplete([]stageResult{duplicate}, nil))
}

func TestParseKnowPostFeedMetrics(t *testing.T) {
	raw := []byte(`# HELP zhiguang_knowpost_feed_stage_total stage calls
# TYPE zhiguang_knowpost_feed_stage_total counter
zhiguang_knowpost_feed_stage_total{outcome="success",stage="relation"} 12
zhiguang_knowpost_feed_stage_duration_ms_sum{outcome="success",stage="relation"} 30.5
zhiguang_knowpost_feed_stage_duration_ms_count{outcome="success",stage="relation"} 12
zhiguang_knowpost_feed_dependency_calls_total{dependency="redis",operation="inbox_read",outcome="success"} 11
zhiguang_knowpost_feed_cold_compute_total{outcome="success"} 10
zhiguang_knowpost_feed_page_cache_total{outcome="success",source="l1_fresh"} 9
zhiguang_knowpost_feed_pagination_total{mode="cursor",result="success"} 8
zhiguang_knowpost_feed_pagination_duration_ms_sum{mode="cursor",result="success"} 24
zhiguang_knowpost_feed_pagination_duration_ms_count{mode="cursor",result="success"} 8
zhiguang_knowpost_feed_cursor_work_total{work="redis_members"} 160
zhiguang_knowpost_feed_page_refresh_queue_depth 7
zhiguang_knowpost_feed_page_refresh_active_workers 3
zhiguang_knowpost_feed_page_refresh_pending_keys 10
process_start_time_seconds 12345
go_gc_duration_seconds_sum 1.25
go_gc_duration_seconds_count 42
go_memstats_alloc_bytes_total 123456
go_memstats_heap_alloc_bytes 654321
go_goroutines 99
`)

	values, err := parseKnowPostFeedMetrics(raw)

	require.NoError(t, err)
	require.Equal(t, 12.0, values["stage_total:relation:success"])
	require.Equal(t, 30.5, values["stage_duration_sum:relation:success"])
	require.Equal(t, 12.0, values["stage_duration_count:relation:success"])
	require.Equal(t, 11.0, values["dependency_calls:redis:inbox_read:success"])
	require.Equal(t, 10.0, values["cold_compute:success"])
	require.Equal(t, 9.0, values["page_cache:l1_fresh:success"])
	require.Equal(t, 8.0, values["pagination:cursor:success"])
	require.Equal(t, 24.0, values["pagination_duration_sum:cursor:success"])
	require.Equal(t, 8.0, values["pagination_duration_count:cursor:success"])
	require.Equal(t, 160.0, values["cursor_work:redis_members"])
	require.Equal(t, 7.0, values["refresh_queue_depth"])
	require.Equal(t, 3.0, values["refresh_active_workers"])
	require.Equal(t, 10.0, values["refresh_pending_keys"])
	require.Equal(t, 12345.0, values["process_start_time_seconds"])
	require.Equal(t, 1.25, values["go_gc_duration_seconds_sum"])
	require.Equal(t, 42.0, values["go_gc_duration_seconds_count"])
	require.Equal(t, 123456.0, values["go_memstats_alloc_bytes_total"])
	require.Equal(t, 654321.0, values["go_memstats_heap_alloc_bytes"])
	require.NotContains(t, values, "go_goroutines")
}

func TestValidateMonitorEvidenceRejectsMissingBoundariesAndProcessChange(t *testing.T) {
	missing := validateMonitorEvidence(nil,
		monitorCollectionStatus{},
		monitorCollectionStatus{feedMetrics: true, dockerState: true, processStartTime: 200},
	)
	require.Contains(t, missing, "feed_metrics_initial")
	require.Contains(t, missing, "docker_state_initial")

	missing = validateMonitorEvidence(nil,
		monitorCollectionStatus{feedMetrics: true, dockerState: true, processStartTime: 100},
		monitorCollectionStatus{},
	)
	require.Contains(t, missing, "feed_metrics_final")
	require.Contains(t, missing, "docker_state_final")

	missing = validateMonitorEvidence(nil,
		monitorCollectionStatus{feedMetrics: true, dockerState: true, processStartTime: 100},
		monitorCollectionStatus{feedMetrics: true, dockerState: true, processStartTime: 200},
	)
	require.Contains(t, missing, "feed_process_changed")
}

func TestValidateMonitorEvidenceRejectsContainerIdentityAndRestartChanges(t *testing.T) {
	samples := []resourceSample{
		{Timestamp: time.Unix(1, 0), Source: "docker-state:zg-knowpost", Identity: "old", Values: map[string]float64{"running": 1, "healthy": 1, "restart_count": 0}},
		{Timestamp: time.Unix(2, 0), Source: "docker-state:zg-knowpost", Identity: "new", Values: map[string]float64{"running": 1, "healthy": 1, "restart_count": 1}},
	}
	boundary := monitorCollectionStatus{feedMetrics: true, dockerState: true, processStartTime: 100}

	missing := validateMonitorEvidence(samples, boundary, boundary)

	require.Contains(t, missing, "container_identity_changed")
	require.Contains(t, missing, "container_restart")
}
