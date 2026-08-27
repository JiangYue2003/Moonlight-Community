package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type containerFeedFeatureFlags struct {
	ObservabilityEnabled    bool
	RouteSnapshotEnabled    bool
	RelationEpochEnabled    bool
	SafetyEpochEnabled      bool
	CombinedPipelineEnabled bool
	CursorPaginationEnabled bool
	PageCacheMode           string
}

type feedRuntimeState struct {
	Strategy string
	Features containerFeedFeatureFlags
	Logging  runtimeLoggingState
	Source   string
}

type runtimeLoggingState struct {
	Preset            string
	Level             string
	Stat              bool
	GatewayAccessLog  bool
	RPCStatMiddleware bool
	SQLStatementInfo  bool
	Persistence       string
}

func readFeedRuntimeState(ctx context.Context, execution executionConfig, monitor monitorConfig) (feedRuntimeState, error) {
	if execution.Topology == "compose-middleware-local-services" {
		return readLocalFeedRuntimeState(execution.LocalStateFile)
	}
	strategy, err := readContainerFeedStrategy(ctx, monitor.StrategyContainer)
	if err != nil {
		return feedRuntimeState{}, err
	}
	features, err := readContainerFeedFeatureFlags(ctx, monitor.StrategyContainer)
	if err != nil {
		return feedRuntimeState{}, err
	}
	return feedRuntimeState{
		Strategy: strategy, Features: features, Source: "docker:" + monitor.StrategyContainer,
	}, nil
}

func readLocalFeedRuntimeState(path string) (feedRuntimeState, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return feedRuntimeState{}, fmt.Errorf("local service state file is required")
	}
	body, err := os.ReadFile(path)
	if err != nil {
		return feedRuntimeState{}, fmt.Errorf("read local service state %s: %w", path, err)
	}
	var state struct {
		Topology     string `json:"topology"`
		Strategy     string `json:"strategy"`
		FeatureFlags struct {
			Observability    string `json:"observability"`
			RelationEpoch    string `json:"relation_epoch"`
			SafetyEpoch      string `json:"safety_epoch"`
			RouteSnapshot    string `json:"route_snapshot"`
			CombinedPipeline string `json:"combined_pipeline"`
			CursorPagination string `json:"cursor_pagination"`
			PageCacheMode    string `json:"page_cache_mode"`
		} `json:"feature_flags"`
		Logging struct {
			Preset            string `json:"preset"`
			Level             string `json:"level"`
			Stat              bool   `json:"stat"`
			GatewayAccessLog  bool   `json:"gateway_access_log"`
			RPCStatMiddleware bool   `json:"rpc_stat_middleware"`
			SQLStatementInfo  bool   `json:"sql_statement_info"`
			Persistence       string `json:"persistence"`
		} `json:"logging"`
	}
	if err := json.Unmarshal(body, &state); err != nil {
		return feedRuntimeState{}, fmt.Errorf("decode local service state %s: %w", path, err)
	}
	if state.Topology != "compose-middleware-local-services" {
		return feedRuntimeState{}, fmt.Errorf("local service state %s has topology %q", path, state.Topology)
	}
	strategy, err := parseContainerFeedStrategy("FEED_STRATEGY=" + state.Strategy)
	if err != nil {
		return feedRuntimeState{}, fmt.Errorf("decode local Feed strategy: %w", err)
	}
	rawFlags := strings.Join([]string{
		"FEED_OBSERVABILITY=" + state.FeatureFlags.Observability,
		"FEED_ROUTE_SNAPSHOT_ENABLED=" + state.FeatureFlags.RouteSnapshot,
		"FEED_RELATION_EPOCH_ENABLED=" + state.FeatureFlags.RelationEpoch,
		"FEED_CONTENT_SAFETY_EPOCH_ENABLED=" + state.FeatureFlags.SafetyEpoch,
		"FEED_COMBINED_PIPELINE_ENABLED=" + state.FeatureFlags.CombinedPipeline,
		"FEED_CURSOR_PAGINATION_ENABLED=" + state.FeatureFlags.CursorPagination,
		"FEED_PAGE_CACHE_MODE=" + state.FeatureFlags.PageCacheMode,
	}, "\n")
	features, err := parseContainerFeedFeatureFlags(rawFlags)
	if err != nil {
		return feedRuntimeState{}, fmt.Errorf("decode local Feed feature flags: %w", err)
	}
	logging := runtimeLoggingState{
		Preset:            strings.ToLower(strings.TrimSpace(state.Logging.Preset)),
		Level:             strings.ToLower(strings.TrimSpace(state.Logging.Level)),
		Stat:              state.Logging.Stat,
		GatewayAccessLog:  state.Logging.GatewayAccessLog,
		RPCStatMiddleware: state.Logging.RPCStatMiddleware,
		SQLStatementInfo:  state.Logging.SQLStatementInfo,
		Persistence:       strings.ToLower(strings.TrimSpace(state.Logging.Persistence)),
	}
	if err := validateLocalRuntimeLogging(logging); err != nil {
		return feedRuntimeState{}, fmt.Errorf("decode local logging state: %w", err)
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		absolute = path
	}
	return feedRuntimeState{Strategy: strategy, Features: features, Logging: logging, Source: "local-state:" + absolute}, nil
}

func validateLocalRuntimeLogging(logging runtimeLoggingState) error {
	switch logging.Preset {
	case "development":
		if logging.Level == "" || logging.Persistence == "" {
			return fmt.Errorf("development logging requires level and persistence evidence")
		}
		return nil
	case "benchmark-error-only":
		if logging.Level != "error" || logging.Stat || logging.GatewayAccessLog || logging.RPCStatMiddleware || logging.SQLStatementInfo || logging.Persistence != "error-only-bounded" {
			return fmt.Errorf(
				"benchmark-error-only requires level=error stat=false gateway_access_log=false rpc_stat_middleware=false sql_statement_info=false persistence=error-only-bounded",
			)
		}
		return nil
	default:
		return fmt.Errorf("logging preset %q must be development or benchmark-error-only", logging.Preset)
	}
}

func phaseRequiresEffectiveStrategy(phase string) bool {
	switch phase {
	case "setup", "resume-setup", "setup-cursor-deep", "seed-cursor-deep", "smoke", "seed-feed", "run", "all",
		"create-mutation-checkpoint", "verify-mutation-checkpoint", "restore-mutation-checkpoint":
		return true
	default:
		return false
	}
}

func readContainerFeedStrategy(ctx context.Context, container string) (string, error) {
	if strings.TrimSpace(container) == "" {
		return "", fmt.Errorf("strategy container is required")
	}
	raw, err := exec.CommandContext(
		ctx,
		"docker", "inspect", container,
		"--format", "{{range .Config.Env}}{{println .}}{{end}}",
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("inspect feed strategy container %s: %w: %s", container, err, strings.TrimSpace(string(raw)))
	}
	return parseContainerFeedStrategy(string(raw))
}

func readContainerFeedFeatureFlags(ctx context.Context, container string) (containerFeedFeatureFlags, error) {
	if strings.TrimSpace(container) == "" {
		return containerFeedFeatureFlags{}, fmt.Errorf("strategy container is required")
	}
	raw, err := exec.CommandContext(
		ctx,
		"docker", "inspect", container,
		"--format", "{{range .Config.Env}}{{println .}}{{end}}",
	).CombinedOutput()
	if err != nil {
		return containerFeedFeatureFlags{}, fmt.Errorf("inspect feed feature flags for container %s: %w: %s", container, err, strings.TrimSpace(string(raw)))
	}
	return parseContainerFeedFeatureFlags(string(raw))
}

func parseContainerFeedStrategy(raw string) (string, error) {
	for _, line := range strings.Split(raw, "\n") {
		key, value, ok := strings.Cut(strings.TrimSpace(line), "=")
		if ok && key == "FEED_STRATEGY" {
			strategy := strings.ToLower(strings.TrimSpace(value))
			if strategy == "push" || strategy == "pull" || strategy == "hybrid" {
				return strategy, nil
			}
			return "", fmt.Errorf("container has invalid FEED_STRATEGY=%q", value)
		}
	}
	return "", fmt.Errorf("container does not expose FEED_STRATEGY")
}

func parseContainerFeedFeatureFlags(raw string) (containerFeedFeatureFlags, error) {
	var flags containerFeedFeatureFlags
	var observabilitySeen, routeSeen, epochSeen, safetySeen, pipelineSeen, cursorSeen, pageCacheSeen bool
	for _, line := range strings.Split(raw, "\n") {
		key, value, ok := strings.Cut(strings.TrimSpace(line), "=")
		if !ok {
			continue
		}
		switch key {
		case "FEED_OBSERVABILITY":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_OBSERVABILITY=%q: %w", value, err)
			}
			flags.ObservabilityEnabled = parsed
			observabilitySeen = true
		case "FEED_ROUTE_SNAPSHOT_ENABLED":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_ROUTE_SNAPSHOT_ENABLED=%q: %w", value, err)
			}
			flags.RouteSnapshotEnabled = parsed
			routeSeen = true
		case "FEED_RELATION_EPOCH_ENABLED":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_RELATION_EPOCH_ENABLED=%q: %w", value, err)
			}
			flags.RelationEpochEnabled = parsed
			epochSeen = true
		case "FEED_CONTENT_SAFETY_EPOCH_ENABLED":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_CONTENT_SAFETY_EPOCH_ENABLED=%q: %w", value, err)
			}
			flags.SafetyEpochEnabled = parsed
			safetySeen = true
		case "FEED_COMBINED_PIPELINE_ENABLED":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_COMBINED_PIPELINE_ENABLED=%q: %w", value, err)
			}
			flags.CombinedPipelineEnabled = parsed
			pipelineSeen = true
		case "FEED_CURSOR_PAGINATION_ENABLED":
			parsed, err := strconv.ParseBool(strings.TrimSpace(value))
			if err != nil {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_CURSOR_PAGINATION_ENABLED=%q: %w", value, err)
			}
			flags.CursorPaginationEnabled = parsed
			cursorSeen = true
		case "FEED_PAGE_CACHE_MODE":
			mode := strings.ToLower(strings.TrimSpace(value))
			if mode != "off" && mode != "l2" && mode != "l1-l2" {
				return containerFeedFeatureFlags{}, fmt.Errorf("invalid FEED_PAGE_CACHE_MODE=%q: expected off, l2, or l1-l2", value)
			}
			flags.PageCacheMode = mode
			pageCacheSeen = true
		}
	}
	if !observabilitySeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_OBSERVABILITY")
	}
	if !routeSeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_ROUTE_SNAPSHOT_ENABLED")
	}
	if !epochSeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_RELATION_EPOCH_ENABLED")
	}
	if !safetySeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_CONTENT_SAFETY_EPOCH_ENABLED")
	}
	if !pipelineSeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_COMBINED_PIPELINE_ENABLED")
	}
	if !cursorSeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_CURSOR_PAGINATION_ENABLED")
	}
	if !pageCacheSeen {
		return containerFeedFeatureFlags{}, fmt.Errorf("container does not expose FEED_PAGE_CACHE_MODE")
	}
	return flags, nil
}

func validateRuntimeFeaturePreset(preset string, flags containerFeedFeatureFlags) error {
	preset = strings.ToLower(strings.TrimSpace(preset))
	if preset != "control" && preset != "treatment" && preset != "cursor-ab" {
		return fmt.Errorf("feature preset %q must be control, treatment, or cursor-ab", preset)
	}
	if !flags.ObservabilityEnabled {
		return fmt.Errorf("feature preset mismatch: %s requires FEED_OBSERVABILITY=true", preset)
	}
	control := !flags.RouteSnapshotEnabled && !flags.RelationEpochEnabled && !flags.SafetyEpochEnabled &&
		!flags.CombinedPipelineEnabled && !flags.CursorPaginationEnabled && flags.PageCacheMode == "off"
	treatment := flags.RouteSnapshotEnabled && flags.RelationEpochEnabled && flags.SafetyEpochEnabled &&
		flags.CombinedPipelineEnabled && flags.CursorPaginationEnabled && flags.PageCacheMode == "l1-l2"
	cursorAB := flags.RouteSnapshotEnabled && flags.RelationEpochEnabled && flags.SafetyEpochEnabled &&
		flags.CombinedPipelineEnabled && flags.CursorPaginationEnabled && flags.PageCacheMode == "off"
	if (preset == "control" && !control) || (preset == "treatment" && !treatment) || (preset == "cursor-ab" && !cursorAB) {
		return fmt.Errorf(
			"feature preset mismatch: preset=%s route=%t relation_epoch=%t safety_epoch=%t pipeline=%t cursor=%t page_cache=%s",
			preset, flags.RouteSnapshotEnabled, flags.RelationEpochEnabled, flags.SafetyEpochEnabled,
			flags.CombinedPipelineEnabled, flags.CursorPaginationEnabled, flags.PageCacheMode,
		)
	}
	return nil
}

func validateEffectiveStrategy(label, effective string) error {
	label = strings.ToLower(strings.TrimSpace(label))
	effective = strings.ToLower(strings.TrimSpace(effective))
	if label != effective {
		return fmt.Errorf("feed strategy mismatch: benchmark label=%q container=%q", label, effective)
	}
	return nil
}
