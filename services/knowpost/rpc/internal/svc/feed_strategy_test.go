package svc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
)

func TestFeedPageCacheConfigTagsDefaultToJitterSafeTTLs(t *testing.T) {
	var decoded struct {
		PageCache config.FeedPageCacheConf
	}
	require.NoError(t, conf.LoadFromYamlBytes([]byte("PageCache:\n  Mode: \"off\"\n"), &decoded))
	require.Equal(t, 800*time.Millisecond, decoded.PageCache.L1FreshTTL)
	require.Equal(t, 4*time.Second, decoded.PageCache.L2FreshTTL)
	require.NoError(t, applyFeedPageCacheDefaults(&decoded.PageCache))
}

func TestFeedRouteCacheTTLBoundsRoutingTransitionWindow(t *testing.T) {
	require.LessOrEqual(t, feedRouteCacheTTL, 5*time.Second)
}

func TestFeedRouteSnapshotDefaultsRemainOptInAndBounded(t *testing.T) {
	cfg := config.FeedRouteSnapshotConf{}
	require.NoError(t, applyFeedRouteSnapshotDefaults(&cfg))

	require.False(t, cfg.Enabled)
	require.Equal(t, 5*time.Second, cfg.TTL)
}

func TestFeedRouteSnapshotRejectsTTLAboveConsistencyBound(t *testing.T) {
	cfg := config.FeedRouteSnapshotConf{Enabled: true, TTL: 5*time.Second + time.Nanosecond}

	err := applyFeedRouteSnapshotDefaults(&cfg)

	require.ErrorContains(t, err, "must not exceed 5s")
}

func TestFeedEpochConfigDefaultsStayShortAndIndependentlyBudgeted(t *testing.T) {
	cfg := config.FeedEpochConf{}
	applyFeedEpochDefaults(&cfg)

	require.Equal(t, "feed", cfg.KeyPrefix)
	require.Equal(t, time.Second, cfg.RelationL1TTL)
	require.Equal(t, time.Second, cfg.SafetyL1TTL)
}

func TestFeedCombinedPipelineDefaultsRemainOptInAndBounded(t *testing.T) {
	cfg := config.FeedCombinedPipelineConf{}
	require.NoError(t, applyFeedCombinedPipelineDefaults(&cfg))

	require.False(t, cfg.Enabled)
	require.Equal(t, 128, cfg.BatchSize)
}

func TestFeedCombinedPipelineRejectsBatchOutsideBounds(t *testing.T) {
	for _, batchSize := range []int{1, 129} {
		cfg := config.FeedCombinedPipelineConf{Enabled: true, BatchSize: batchSize}

		err := applyFeedCombinedPipelineDefaults(&cfg)

		require.ErrorContains(t, err, "must be between 2 and 128")
	}
}

func TestFeedCursorPaginationDefaultsRemainOptIn(t *testing.T) {
	var decoded struct {
		CursorPagination config.FeedCursorPaginationConf
	}
	require.NoError(t, conf.LoadFromYamlBytes([]byte("CursorPagination:\n  Enabled: false\n"), &decoded))
	require.False(t, decoded.CursorPagination.Enabled)
}

func TestFeedPageCacheDefaultsRemainOptInAndBounded(t *testing.T) {
	cfg := config.FeedPageCacheConf{}
	require.NoError(t, applyFeedPageCacheDefaults(&cfg))

	require.Equal(t, "off", cfg.Mode)
	require.Equal(t, "feed", cfg.KeyPrefix)
	require.Equal(t, int32(1), cfg.Page)
	require.Equal(t, int32(20), cfg.Size)
	require.Equal(t, 800*time.Millisecond, cfg.L1FreshTTL)
	require.Equal(t, 4*time.Second, cfg.L2FreshTTL)
	require.Equal(t, 10*time.Second, cfg.StaleTTL)
	require.Equal(t, 20, cfg.JitterPercent)
	require.Equal(t, 32, cfg.RefreshWorkers)
	require.Equal(t, 1024, cfg.RefreshQueue)
	require.Equal(t, 2*time.Second, cfg.LoaderTimeout)
}

func TestFeedPageCacheRejectsUnsupportedModeOrFreshnessWindow(t *testing.T) {
	for _, cfg := range []config.FeedPageCacheConf{
		{Mode: "l1"},
		{Mode: "unknown"},
		{Mode: "l2", L2FreshTTL: 5 * time.Second},
		{Mode: "l2", L2FreshTTL: 2 * time.Second},
		{Mode: "l1-l2", L1FreshTTL: time.Second, L2FreshTTL: 4 * time.Second},
		{Mode: "l1-l2", L1FreshTTL: 400 * time.Millisecond, L2FreshTTL: 4 * time.Second},
		{Mode: "l2", Page: 2, Size: 20},
		{Mode: "l2", Page: 1, Size: 10},
		{Mode: "l2", StaleTTL: -time.Second},
		{Mode: "l2", StaleTTL: 10*time.Second + time.Nanosecond},
		{Mode: "l2", JitterPercent: -1},
		{Mode: "l2", JitterPercent: 21},
		{Mode: "l2", RefreshWorkers: -1},
		{Mode: "l2", RefreshQueue: -1},
		{Mode: "l2", LoaderTimeout: -time.Second},
	} {
		err := applyFeedPageCacheDefaults(&cfg)
		require.Error(t, err, "cfg=%+v", cfg)
	}
}

func TestFeedPageCacheRequiresDurableSafetyEpochConsumer(t *testing.T) {
	err := validateFeedSafetyConfiguration(config.FeedConf{
		PageCache: config.FeedPageCacheConf{Mode: "l2"},
		Epoch:     config.FeedEpochConf{SafetyConsumerEnabled: false},
	})
	require.ErrorContains(t, err, "SafetyConsumerEnabled")

	require.NoError(t, validateFeedSafetyConfiguration(config.FeedConf{
		PageCache: config.FeedPageCacheConf{Mode: "l1-l2"},
		Epoch:     config.FeedEpochConf{SafetyConsumerEnabled: true},
	}))
	require.NoError(t, validateFeedSafetyConfiguration(config.FeedConf{
		PageCache: config.FeedPageCacheConf{Mode: "off"},
	}))
}

func TestResolveFeedStrategy(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want feed.Strategy
	}{
		{name: "default", raw: "", want: feed.StrategyHybrid},
		{name: "push", raw: "push", want: feed.StrategyPush},
		{name: "pull", raw: "pull", want: feed.StrategyPull},
		{name: "hybrid", raw: "hybrid", want: feed.StrategyHybrid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveFeedStrategy(config.FeedConf{Strategy: tt.raw})
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestResolveFeedStrategyRejectsInvalidConfig(t *testing.T) {
	_, err := resolveFeedStrategy(config.FeedConf{Strategy: "invalid"})

	require.ErrorContains(t, err, "invalid feed strategy")
}

func TestResolveAuthorTierMode(t *testing.T) {
	for _, test := range []struct {
		raw  string
		want feed.AuthorTierMode
	}{
		{raw: "", want: feed.AuthorTierModeOff},
		{raw: "off", want: feed.AuthorTierModeOff},
		{raw: "shadow", want: feed.AuthorTierModeShadow},
		{raw: "enforce", want: feed.AuthorTierModeEnforce},
	} {
		got, err := resolveAuthorTierMode(config.FeedConf{
			AuthorTier: config.FeedAuthorTierConf{Mode: test.raw},
		})
		require.NoError(t, err)
		require.Equal(t, test.want, got)
	}
}

func TestResolveAuthorTierModeRejectsInvalidConfig(t *testing.T) {
	_, err := resolveAuthorTierMode(config.FeedConf{
		AuthorTier: config.FeedAuthorTierConf{Mode: "automatic"},
	})

	require.ErrorContains(t, err, "invalid feed author tier mode")
}
