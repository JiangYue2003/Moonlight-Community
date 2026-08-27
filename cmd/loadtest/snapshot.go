package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type environmentSnapshot struct {
	RunID                        string                 `json:"run_id"`
	CapturedAt                   time.Time              `json:"captured_at"`
	EffectiveStrategy            string                 `json:"effective_strategy,omitempty"`
	ObservabilityEnabled         bool                   `json:"observability_enabled"`
	RouteSnapshotEnabled         bool                   `json:"route_snapshot_enabled"`
	RelationEpochConsumerEnabled bool                   `json:"relation_epoch_consumer_enabled"`
	SafetyEpochConsumerEnabled   bool                   `json:"safety_epoch_consumer_enabled"`
	CombinedPipelineEnabled      bool                   `json:"combined_pipeline_enabled"`
	CursorPaginationEnabled      bool                   `json:"cursor_pagination_enabled"`
	PageCacheMode                string                 `json:"page_cache_mode"`
	Topology                     string                 `json:"topology"`
	StrategySource               string                 `json:"strategy_source"`
	Containers                   []dockerContainerState `json:"containers"`
	Kafka                        map[string]float64     `json:"kafka"`
	MissingData                  []string               `json:"missing_data,omitempty"`
}

func captureEnvironmentSnapshot(ctx context.Context, cfg monitorConfig, execution executionConfig, runID string) (environmentSnapshot, error) {
	_, containers, err := collectDockerState(ctx, cfg.DockerContainers, time.Now())
	if err != nil {
		return environmentSnapshot{}, err
	}
	kafka, err := collectKafkaLag(ctx, cfg)
	if err != nil {
		return environmentSnapshot{}, err
	}
	runtimeState, err := readFeedRuntimeState(ctx, execution, cfg)
	if err != nil {
		return environmentSnapshot{}, err
	}
	strategy := runtimeState.Strategy
	features := runtimeState.Features
	return environmentSnapshot{
		RunID:                        runID,
		CapturedAt:                   time.Now(),
		EffectiveStrategy:            strategy,
		ObservabilityEnabled:         features.ObservabilityEnabled,
		RouteSnapshotEnabled:         features.RouteSnapshotEnabled,
		RelationEpochConsumerEnabled: features.RelationEpochEnabled,
		SafetyEpochConsumerEnabled:   features.SafetyEpochEnabled,
		CombinedPipelineEnabled:      features.CombinedPipelineEnabled,
		CursorPaginationEnabled:      features.CursorPaginationEnabled,
		PageCacheMode:                features.PageCacheMode,
		Topology:                     execution.Topology,
		StrategySource:               runtimeState.Source,
		Containers:                   containers,
		Kafka:                        kafka,
		MissingData: []string{
			"prometheus_application_metrics",
			"correlated_timeout_panic_retry_logs",
		},
	}, nil
}

func writeEnvironmentSnapshot(root string, snapshot environmentSnapshot) (string, error) {
	dir := filepath.Join(root, safePathSegment(snapshot.RunID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create environment snapshot directory: %w", err)
	}
	body, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal environment snapshot: %w", err)
	}
	path := filepath.Join(dir, "environment-snapshot.json")
	if err := os.WriteFile(path, append(body, '\n'), 0o644); err != nil {
		return "", fmt.Errorf("write environment snapshot: %w", err)
	}
	return path, nil
}
