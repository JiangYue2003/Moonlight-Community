package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadLocalFeedRuntimeState(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	require.NoError(t, os.WriteFile(path, []byte(`{
  "topology": "compose-middleware-local-services",
  "strategy": "hybrid",
  "feature_flags": {
    "observability": "true",
    "relation_epoch": "true",
    "safety_epoch": "true",
    "route_snapshot": "true",
    "combined_pipeline": "true",
	"cursor_pagination": "true",
    "page_cache_mode": "l1-l2"
  },
  "logging": {
    "preset": "benchmark-error-only",
    "level": "error",
    "stat": false,
    "gateway_access_log": false,
	"rpc_stat_middleware": false,
	"sql_statement_info": false,
    "persistence": "error-only-bounded"
  }
}`), 0o600))

	state, err := readLocalFeedRuntimeState(path)

	require.NoError(t, err)
	require.Equal(t, "hybrid", state.Strategy)
	require.True(t, state.Features.RelationEpochEnabled)
	require.True(t, state.Features.SafetyEpochEnabled)
	require.True(t, state.Features.RouteSnapshotEnabled)
	require.True(t, state.Features.CombinedPipelineEnabled)
	require.True(t, state.Features.CursorPaginationEnabled)
	require.Equal(t, "l1-l2", state.Features.PageCacheMode)
	require.Equal(t, "benchmark-error-only", state.Logging.Preset)
	require.Equal(t, "error", state.Logging.Level)
	require.False(t, state.Logging.Stat)
	require.False(t, state.Logging.GatewayAccessLog)
	require.False(t, state.Logging.RPCStatMiddleware)
	require.False(t, state.Logging.SQLStatementInfo)
	require.Equal(t, "error-only-bounded", state.Logging.Persistence)
	require.Contains(t, state.Source, "local-state:")
}

func TestReadLocalFeedRuntimeStateRequiresLoggingEvidence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	require.NoError(t, os.WriteFile(path, []byte(`{
  "topology": "compose-middleware-local-services",
  "strategy": "hybrid",
  "feature_flags": {
    "observability": "true",
    "relation_epoch": "true",
    "safety_epoch": "true",
    "route_snapshot": "true",
    "combined_pipeline": "true",
	"cursor_pagination": "true",
    "page_cache_mode": "l1-l2"
  }
}`), 0o600))

	_, err := readLocalFeedRuntimeState(path)

	require.ErrorContains(t, err, "logging")
}

func TestParseContainerFeedStrategy(t *testing.T) {
	strategy, err := parseContainerFeedStrategy("TZ=Asia/Shanghai\nFEED_STRATEGY=hybrid\nPOD_IP=127.0.0.1\n")

	require.NoError(t, err)
	require.Equal(t, "hybrid", strategy)
}

func TestValidateEffectiveStrategyRejectsMislabeledRun(t *testing.T) {
	err := validateEffectiveStrategy("pull", "hybrid")

	require.ErrorContains(t, err, "strategy mismatch")
}

func TestSetupPhasesRequireEffectiveStrategyVerification(t *testing.T) {
	require.True(t, phaseRequiresEffectiveStrategy("setup"))
	require.True(t, phaseRequiresEffectiveStrategy("resume-setup"))
	require.True(t, phaseRequiresEffectiveStrategy("setup-cursor-deep"))
	require.True(t, phaseRequiresEffectiveStrategy("seed-cursor-deep"))
	require.True(t, phaseRequiresEffectiveStrategy("create-mutation-checkpoint"))
	require.True(t, phaseRequiresEffectiveStrategy("verify-mutation-checkpoint"))
	require.True(t, phaseRequiresEffectiveStrategy("restore-mutation-checkpoint"))
}

func TestValidateRuntimeFeaturePreset(t *testing.T) {
	treatment := containerFeedFeatureFlags{
		ObservabilityEnabled: true, RouteSnapshotEnabled: true, RelationEpochEnabled: true,
		SafetyEpochEnabled: true, CombinedPipelineEnabled: true, CursorPaginationEnabled: true, PageCacheMode: "l1-l2",
	}
	require.NoError(t, validateRuntimeFeaturePreset("treatment", treatment))
	cursorAB := treatment
	cursorAB.PageCacheMode = "off"
	require.NoError(t, validateRuntimeFeaturePreset("cursor-ab", cursorAB))

	control := containerFeedFeatureFlags{ObservabilityEnabled: true, PageCacheMode: "off"}
	require.NoError(t, validateRuntimeFeaturePreset("control", control))

	control.RouteSnapshotEnabled = true
	require.ErrorContains(t, validateRuntimeFeaturePreset("control", control), "feature preset mismatch")
	require.ErrorContains(t, validateRuntimeFeaturePreset("cursor-ab", treatment), "feature preset mismatch")
}

func TestParseContainerFeedFeatureFlags(t *testing.T) {
	flags, err := parseContainerFeedFeatureFlags("TZ=Asia/Shanghai\nFEED_OBSERVABILITY=true\nFEED_ROUTE_SNAPSHOT_ENABLED=true\nFEED_RELATION_EPOCH_ENABLED=false\nFEED_CONTENT_SAFETY_EPOCH_ENABLED=true\nFEED_COMBINED_PIPELINE_ENABLED=true\nFEED_CURSOR_PAGINATION_ENABLED=true\nFEED_PAGE_CACHE_MODE=l2\n")

	require.NoError(t, err)
	require.True(t, flags.ObservabilityEnabled)
	require.True(t, flags.RouteSnapshotEnabled)
	require.False(t, flags.RelationEpochEnabled)
	require.True(t, flags.SafetyEpochEnabled)
	require.True(t, flags.CombinedPipelineEnabled)
	require.True(t, flags.CursorPaginationEnabled)
	require.Equal(t, "l2", flags.PageCacheMode)
}

func TestParseContainerFeedFeatureFlagsRequiresExplicitValues(t *testing.T) {
	_, err := parseContainerFeedFeatureFlags("FEED_STRATEGY=hybrid\n")

	require.ErrorContains(t, err, "FEED_OBSERVABILITY")
}

func TestParseContainerFeedFeatureFlagsRejectsInvalidBoolean(t *testing.T) {
	_, err := parseContainerFeedFeatureFlags("FEED_OBSERVABILITY=true\nFEED_ROUTE_SNAPSHOT_ENABLED=sometimes\nFEED_RELATION_EPOCH_ENABLED=true\nFEED_CONTENT_SAFETY_EPOCH_ENABLED=true\nFEED_COMBINED_PIPELINE_ENABLED=false\nFEED_CURSOR_PAGINATION_ENABLED=false\nFEED_PAGE_CACHE_MODE=off\n")

	require.ErrorContains(t, err, "invalid FEED_ROUTE_SNAPSHOT_ENABLED")
}

func TestParseContainerFeedFeatureFlagsRequiresCombinedPipelineValue(t *testing.T) {
	_, err := parseContainerFeedFeatureFlags("FEED_OBSERVABILITY=true\nFEED_ROUTE_SNAPSHOT_ENABLED=true\nFEED_RELATION_EPOCH_ENABLED=true\nFEED_CONTENT_SAFETY_EPOCH_ENABLED=true\nFEED_CURSOR_PAGINATION_ENABLED=true\nFEED_PAGE_CACHE_MODE=off\n")

	require.ErrorContains(t, err, "FEED_COMBINED_PIPELINE_ENABLED")
}

func TestParseContainerFeedFeatureFlagsRequiresPageCacheMode(t *testing.T) {
	_, err := parseContainerFeedFeatureFlags("FEED_OBSERVABILITY=true\nFEED_ROUTE_SNAPSHOT_ENABLED=true\nFEED_RELATION_EPOCH_ENABLED=true\nFEED_CONTENT_SAFETY_EPOCH_ENABLED=true\nFEED_COMBINED_PIPELINE_ENABLED=true\nFEED_CURSOR_PAGINATION_ENABLED=true\n")

	require.ErrorContains(t, err, "FEED_PAGE_CACHE_MODE")
}

func TestParseContainerFeedFeatureFlagsRejectsInvalidPageCacheMode(t *testing.T) {
	_, err := parseContainerFeedFeatureFlags("FEED_OBSERVABILITY=true\nFEED_ROUTE_SNAPSHOT_ENABLED=true\nFEED_RELATION_EPOCH_ENABLED=true\nFEED_CONTENT_SAFETY_EPOCH_ENABLED=true\nFEED_COMBINED_PIPELINE_ENABLED=true\nFEED_CURSOR_PAGINATION_ENABLED=true\nFEED_PAGE_CACHE_MODE=l1\n")

	require.ErrorContains(t, err, "invalid FEED_PAGE_CACHE_MODE")
}
