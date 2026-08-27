package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWriteEnvironmentSnapshotPersistsContainerStateAndGaps(t *testing.T) {
	snapshot := environmentSnapshot{
		RunID:                        "run-42",
		CapturedAt:                   time.Unix(42, 0),
		EffectiveStrategy:            "hybrid",
		RouteSnapshotEnabled:         true,
		RelationEpochConsumerEnabled: true,
		SafetyEpochConsumerEnabled:   true,
		CombinedPipelineEnabled:      true,
		CursorPaginationEnabled:      true,
		PageCacheMode:                "l2",
		Containers:                   []dockerContainerState{{Name: "zg-knowpost", Running: true, Health: "healthy"}},
		Kafka:                        map[string]float64{"lag_total": 0},
		MissingData:                  []string{"prometheus_application_metrics"},
	}

	path, err := writeEnvironmentSnapshot(t.TempDir(), snapshot)
	require.NoError(t, err)
	body, err := os.ReadFile(path)
	require.NoError(t, err)

	var decoded environmentSnapshot
	require.NoError(t, json.Unmarshal(body, &decoded))
	require.Equal(t, snapshot, decoded)
}
