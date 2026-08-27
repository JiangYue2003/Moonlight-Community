package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigExpandsFeedStrategyEnvironment(t *testing.T) {
	t.Setenv("FEED_STRATEGY", "pull")
	t.Setenv("FEED_OBSERVABILITY", "true")
	t.Setenv("FEED_RELATION_EPOCH_ENABLED", "true")
	t.Setenv("FEED_CONTENT_SAFETY_EPOCH_ENABLED", "true")
	t.Setenv("FEED_ROUTE_SNAPSHOT_ENABLED", "true")
	t.Setenv("FEED_COMBINED_PIPELINE_ENABLED", "true")
	t.Setenv("FEED_CURSOR_PAGINATION_ENABLED", "true")
	t.Setenv("FEED_PAGE_CACHE_MODE", "l2")

	cfg, err := loadConfig("etc/knowpost-docker.yaml")

	require.NoError(t, err)
	require.Equal(t, "pull", cfg.Rpc.Feed.Strategy)
	require.True(t, cfg.Rpc.Feed.Observability.Enabled)
	require.Equal(t, "feed", cfg.Rpc.Feed.Epoch.KeyPrefix)
	require.Equal(t, time.Second, cfg.Rpc.Feed.Epoch.RelationL1TTL)
	require.Equal(t, time.Second, cfg.Rpc.Feed.Epoch.SafetyL1TTL)
	require.True(t, cfg.Rpc.Feed.Epoch.RelationConsumerEnabled)
	require.True(t, cfg.Rpc.Feed.Epoch.SafetyConsumerEnabled)
	require.True(t, cfg.Rpc.Feed.RouteSnapshot.Enabled)
	require.Equal(t, 5*time.Second, cfg.Rpc.Feed.RouteSnapshot.TTL)
	require.True(t, cfg.Rpc.Feed.CombinedPipeline.Enabled)
	require.Equal(t, 128, cfg.Rpc.Feed.CombinedPipeline.BatchSize)
	require.True(t, cfg.Rpc.Feed.CursorPagination.Enabled)
	require.Equal(t, "l2", cfg.Rpc.Feed.PageCache.Mode)
	require.Equal(t, "feed", cfg.Rpc.Feed.PageCache.KeyPrefix)
	require.Equal(t, int32(1), cfg.Rpc.Feed.PageCache.Page)
	require.Equal(t, int32(20), cfg.Rpc.Feed.PageCache.Size)
	require.Equal(t, 800*time.Millisecond, cfg.Rpc.Feed.PageCache.L1FreshTTL)
	require.Equal(t, 4*time.Second, cfg.Rpc.Feed.PageCache.L2FreshTTL)
	require.Equal(t, 10*time.Second, cfg.Rpc.Feed.PageCache.StaleTTL)
	require.Equal(t, 20, cfg.Rpc.Feed.PageCache.JitterPercent)
	require.Equal(t, 32, cfg.Rpc.Feed.PageCache.RefreshWorkers)
	require.Equal(t, 1024, cfg.Rpc.Feed.PageCache.RefreshQueue)
	require.Equal(t, 2*time.Second, cfg.Rpc.Feed.PageCache.LoaderTimeout)
	require.Equal(t, "canal-outbox", cfg.Rpc.Kafka.CanalOutboxTopic)
	require.Equal(t, "knowpost-feed-relation-epoch", cfg.Rpc.Kafka.RelationEpochGroupId)
	require.Equal(t, "knowpost-feed-content-safety-epoch", cfg.Rpc.Kafka.SafetyEpochGroupId)
	require.Equal(t, int64(200_000), cfg.Rpc.L1.FeedEpochNumCounters)
	require.Equal(t, int64(4), cfg.Rpc.L1.FeedEpochMaxCostMB)
	require.Equal(t, int64(100_000), cfg.Rpc.L1.FeedRouteNumCounters)
	require.Equal(t, int64(32), cfg.Rpc.L1.FeedRouteMaxCostMB)
	require.Equal(t, int64(100_000), cfg.Rpc.L1.FeedPageNumCounters)
	require.Equal(t, int64(128), cfg.Rpc.L1.FeedPageMaxCostMB)
	require.True(t, cfg.DebugHTTP.Enabled)
	require.Equal(t, "0.0.0.0:6064", cfg.DebugHTTP.ListenOn)
}
