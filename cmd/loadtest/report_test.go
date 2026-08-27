package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWriteReportFilesCreatesJSONCSVAndMarkdown(t *testing.T) {
	report := benchmarkReport{
		RunID:       "run-42",
		Strategy:    "hybrid",
		Entry:       "rpc",
		Scenario:    "hot-read",
		StartedAt:   time.Date(2026, 8, 14, 12, 0, 0, 0, time.UTC),
		Duration:    2 * time.Second,
		Concurrency: 8,
		Complete:    true,
		Stages: []stageResult{{
			Name:       "read",
			Total:      10,
			Success:    10,
			SuccessQPS: 5,
			Latency: latencySummary{
				P50: 5 * time.Millisecond,
				P90: 8 * time.Millisecond,
				P95: 9 * time.Millisecond,
				P99: 10 * time.Millisecond,
				Max: 12 * time.Millisecond,
			},
			Pagination: &paginationObservation{
				Mode: "cursor", TargetPage: 20, PreparationRequests: 19,
				PreparationDuration: 25 * time.Millisecond, OracleHash: "oracle-a", ObservedHash: "oracle-a",
			},
		}},
		Resources: []resourceSample{{
			Timestamp: time.Date(2026, 8, 14, 12, 0, 1, 0, time.UTC),
			Source:    "kafka",
			Values:    map[string]float64{"lag_total": 7},
		}},
		Recovery: &recoveryObservation{
			KafkaDrain:   1500 * time.Millisecond,
			FinalMetrics: map[string]float64{"lag_total": 0},
			Complete:     true,
		},
		Mutation: &mutationTrialObservation{
			CheckpointID: "checkpoint-1", BaselineFingerprint: "baseline-1",
			CheckpointCreatedAt: time.Date(2026, 8, 14, 11, 0, 0, 0, time.UTC),
			ToolVersion:         "mutation-v1", PublishedPostIDs: []int64{101, 102},
			AfterTrialRestore: &mutationRestoreObservation{
				CheckpointID: "checkpoint-1", Complete: true, RemovedPosts: 2,
				RemovedPostIDs: []int64{101, 102}, SafetyEpoch: 9,
			},
		},
		FeedMetrics: &feedMetricsSummary{
			CursorWork:           map[string]float64{"redis_members": 240},
			CursorWorkPerSuccess: map[string]float64{"redis_members": 24},
		},
	}

	paths, err := writeReportFiles(t.TempDir(), report)

	require.NoError(t, err)
	require.FileExists(t, paths.JSON)
	require.FileExists(t, paths.CSV)
	require.FileExists(t, paths.Markdown)
	jsonBody, err := os.ReadFile(paths.JSON)
	require.NoError(t, err)
	require.Contains(t, string(jsonBody), `"run_id": "run-42"`)
	csvBody, err := os.ReadFile(paths.CSV)
	require.NoError(t, err)
	require.Contains(t, string(csvBody), "stage,total,success,failed,timeouts,success_qps,p50_ms,p90_ms,p95_ms,p99_ms,max_ms")
	markdownBody, err := os.ReadFile(paths.Markdown)
	require.NoError(t, err)
	require.True(t, strings.Contains(string(markdownBody), "| read | 10 | 10 | 0 |"))
	require.Contains(t, string(markdownBody), "| kafka | lag_total | 7.000 |")
	require.Contains(t, string(markdownBody), "Kafka drain：1.5s")
	require.Contains(t, string(markdownBody), "目标页：20")
	require.Contains(t, string(markdownBody), "准备请求：19")
	require.Contains(t, string(markdownBody), "分页工作量")
	require.Contains(t, string(markdownBody), "| redis_members | 240 | 24.000 |")
	require.Contains(t, string(markdownBody), "Mutation checkpoint")
	require.Contains(t, string(markdownBody), "checkpoint-1")
	require.Contains(t, string(markdownBody), "成功发布帖子：2")
	require.Contains(t, string(markdownBody), "测量后恢复：complete=true")
	require.Equal(t, "report.json", filepath.Base(paths.JSON))
}

func TestSummarizeFeedMetricsUsesRunDeltas(t *testing.T) {
	samples := []resourceSample{
		{
			Timestamp: time.Unix(1, 0),
			Source:    "prometheus:knowpost",
			Values: map[string]float64{
				"dependency_calls:relation:list_followings:success":      100,
				"dependency_calls:counter:batch_follower_counts:success": 80,
				"dependency_calls:redis:inbox_read:success":              100,
				"dependency_calls:redis:bigv_pipeline:success":           100,
				"cold_compute:success":                                   100,
				"stage_duration_sum:total:success":                       1000,
				"stage_duration_count:total:success":                     100,
				"stage_total:total:success":                              100,
				"dependency_calls:redis:inbox_read:error":                2,
				"stage_total:inbox:error":                                2,
				"stage_duration_sum:inbox:error":                         20,
				"stage_duration_count:inbox:error":                       2,
				"page_cache:l1_fresh:success":                            80,
				"page_cache:miss:success":                                20,
				"pagination:cursor:success":                              100,
				"pagination_duration_sum:cursor:success":                 400,
				"pagination_duration_count:cursor:success":               100,
				"cursor_work:redis_members":                              2000,
				"go_gc_duration_seconds_count":                           100,
				"go_memstats_alloc_bytes_total":                          10000,
				"go_memstats_heap_alloc_bytes":                           9000,
				"refresh_queue_depth":                                    0,
				"refresh_active_workers":                                 0,
				"refresh_pending_keys":                                   0,
			},
		},
		{
			Timestamp: time.Unix(2, 0),
			Source:    "prometheus:knowpost",
			Values: map[string]float64{
				"dependency_calls:relation:list_followings:success":      1100,
				"dependency_calls:counter:batch_follower_counts:success": 580,
				"dependency_calls:redis:inbox_read:success":              1100,
				"dependency_calls:redis:bigv_pipeline:success":           1100,
				"cold_compute:success":                                   1100,
				"stage_duration_sum:total:success":                       51000,
				"stage_duration_count:total:success":                     1100,
				"stage_total:total:success":                              1100,
				"dependency_calls:redis:inbox_read:error":                4,
				"stage_total:inbox:error":                                4,
				"stage_duration_sum:inbox:error":                         60,
				"stage_duration_count:inbox:error":                       4,
				"page_cache:l1_fresh:success":                            1050,
				"page_cache:miss:success":                                50,
				"pagination:cursor:success":                              1100,
				"pagination_duration_sum:cursor:success":                 5400,
				"pagination_duration_count:cursor:success":               1100,
				"cursor_work:redis_members":                              22000,
				"go_gc_duration_seconds_count":                           120,
				"go_memstats_alloc_bytes_total":                          20000,
				"go_memstats_heap_alloc_bytes":                           5000,
				"refresh_queue_depth":                                    7,
				"refresh_active_workers":                                 3,
				"refresh_pending_keys":                                   10,
			},
		},
	}

	summary, ok := summarizeFeedMetrics(samples, 1000)

	require.True(t, ok)
	require.Equal(t, 1.0, summary.DependencyCallsPerSuccess["relation"])
	require.Equal(t, 0.5, summary.DependencyCallsPerSuccess["counter"])
	require.Equal(t, 2.0, summary.DependencyCallsPerSuccess["redis"])
	require.Equal(t, 1.0, summary.ColdComputesPerSuccess)
	require.Equal(t, 50.0, summary.StageMeanMilliseconds["total"])
	require.Equal(t, 2.0, summary.DependencyCallsByOutcome["redis"]["error"])
	require.Equal(t, 2.0, summary.StageCallsByOutcome["inbox"]["error"])
	require.Equal(t, 20.0, summary.StageMeanMillisecondsByOutcome["inbox"]["error"])
	require.Equal(t, 970.0, summary.PageCacheSources["l1_fresh"])
	require.Equal(t, 30.0, summary.PageCacheSources["miss"])
	require.InDelta(t, 0.97, summary.PageCacheFreshHitRatio, 0.000001)
	require.Equal(t, 7.0, summary.RefreshQueueDepthMax)
	require.Equal(t, 3.0, summary.RefreshActiveWorkersMax)
	require.Equal(t, 10.0, summary.RefreshPendingKeysMax)
	require.Equal(t, 1000.0, summary.PaginationModes["cursor"])
	require.Equal(t, 5.0, summary.PaginationMeanMilliseconds["cursor"])
	require.Equal(t, 20000.0, summary.CursorWork["redis_members"])
	require.Equal(t, 20.0, summary.CursorWorkPerSuccess["redis_members"])
	require.True(t, summary.Degraded)
}

func TestSummarizeFeedMetricsRejectsGaugeOnlyOrMissingPageCoverage(t *testing.T) {
	gauges := map[string]float64{
		"refresh_queue_depth": 0, "refresh_active_workers": 0, "refresh_pending_keys": 0,
	}
	_, ok := summarizeFeedMetrics([]resourceSample{
		{Timestamp: time.Unix(1, 0), Source: "prometheus:knowpost", Values: gauges},
		{Timestamp: time.Unix(2, 0), Source: "prometheus:knowpost", Values: gauges},
	}, 10)
	require.False(t, ok)

	_, ok = summarizeFeedMetrics([]resourceSample{
		{Timestamp: time.Unix(1, 0), Source: "prometheus:knowpost", Values: map[string]float64{
			"refresh_queue_depth": 0, "refresh_active_workers": 0, "refresh_pending_keys": 0,
			"stage_total:total:success": 100,
		}},
		{Timestamp: time.Unix(2, 0), Source: "prometheus:knowpost", Values: map[string]float64{
			"refresh_queue_depth": 0, "refresh_active_workers": 0, "refresh_pending_keys": 0,
			"stage_total:total:success": 110,
		}},
	}, 10)
	require.False(t, ok)
}

func TestSummarizeRedisMetricsUsesBoundaryDeltas(t *testing.T) {
	samples := []resourceSample{
		{Timestamp: time.Unix(1, 0), Source: "redis", Identity: "redis-a", Values: map[string]float64{
			"uptime_seconds": 100, "commands_total": 1000, "net_input_bytes": 2000,
			"net_output_bytes": 3000, "keyspace_hits": 400, "keyspace_misses": 100,
			"evicted_keys": 2, "rejected_connections": 1, "ops_per_sec": 10, "feed_safety_epoch": 7,
		}},
		{Timestamp: time.Unix(2, 0), Source: "redis", Identity: "redis-a", Values: map[string]float64{
			"uptime_seconds": 101, "commands_total": 1120, "net_input_bytes": 2400,
			"net_output_bytes": 3700, "keyspace_hits": 460, "keyspace_misses": 110,
			"evicted_keys": 3, "rejected_connections": 1, "ops_per_sec": 80, "feed_safety_epoch": 8,
		}},
	}

	summary, ok := summarizeRedisMetrics(samples)

	require.True(t, ok)
	require.Equal(t, 120.0, summary.CommandsDelta)
	require.Equal(t, 400.0, summary.NetInputBytesDelta)
	require.Equal(t, 700.0, summary.NetOutputBytesDelta)
	require.Equal(t, 60.0, summary.KeyspaceHitsDelta)
	require.Equal(t, 10.0, summary.KeyspaceMissesDelta)
	require.Equal(t, 1.0, summary.EvictedKeysDelta)
	require.Equal(t, 0.0, summary.RejectedConnectionsDelta)
	require.Equal(t, 80.0, summary.OpsPerSecondMax)
	require.Equal(t, 7.0, summary.SafetyEpochStart)
	require.Equal(t, 8.0, summary.SafetyEpochEnd)
}

func TestSummarizeRedisMetricsRejectsRestartOrCounterReset(t *testing.T) {
	samples := []resourceSample{
		{Timestamp: time.Unix(1, 0), Source: "redis", Identity: "redis-a", Values: map[string]float64{"uptime_seconds": 100, "commands_total": 100}},
		{Timestamp: time.Unix(2, 0), Source: "redis", Identity: "redis-b", Values: map[string]float64{"uptime_seconds": 101, "commands_total": 101}},
	}

	_, ok := summarizeRedisMetrics(samples)
	require.False(t, ok)
}

func TestRedisSafetyEpochChangeInvalidatesReadButNotMutationScenario(t *testing.T) {
	metrics := &redisMetricsSummary{SafetyEpochStart: 7, SafetyEpochEnd: 8}

	require.True(t, redisSafetyEpochChanged("hot-read", metrics))
	require.True(t, redisSafetyEpochChanged("distributed-read", metrics))
	require.False(t, redisSafetyEpochChanged("publish", metrics))
	require.False(t, redisSafetyEpochChanged("mixed-90-10", metrics))
	require.False(t, redisSafetyEpochChanged("hot-read", &redisMetricsSummary{SafetyEpochStart: 7, SafetyEpochEnd: 7}))
}

func TestWriteReportFilesKeepsTrialsSeparate(t *testing.T) {
	root := t.TempDir()
	base := benchmarkReport{
		RunID: "run-trials", Strategy: "hybrid", Entry: "rpc",
		Scenario: "distributed-read-c128", Trial: 1,
	}

	first, err := writeReportFiles(root, base)
	require.NoError(t, err)
	base.Trial = 2
	second, err := writeReportFiles(root, base)
	require.NoError(t, err)

	require.NotEqual(t, first.JSON, second.JSON)
	require.Equal(t, "trial-1", filepath.Base(filepath.Dir(first.JSON)))
	require.Equal(t, "trial-2", filepath.Base(filepath.Dir(second.JSON)))
}
