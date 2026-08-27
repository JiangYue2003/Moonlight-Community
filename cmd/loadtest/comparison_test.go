package main

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadComparisonRowsRejectsMixedExecutionTopologies(t *testing.T) {
	root := t.TempDir()
	for index, topology := range []string{"compose-full", "compose-middleware-local-services"} {
		report := benchmarkReport{
			RunID: "mixed-topology", Strategy: "hybrid", Entry: "rpc",
			Scenario: "hot-read-c32-" + string(rune('a'+index)), Trial: 1,
			Environment: map[string]string{"topology": topology, "reader_cardinality": "hot", "cache_state": "l1-warm"},
			Stages:      []stageResult{{Name: "read", Total: 1, Success: 1}},
		}
		_, err := writeReportFiles(root, report)
		require.NoError(t, err)
	}

	_, err := loadComparisonRows(root, "mixed-topology")
	require.ErrorContains(t, err, "topology")
	require.DirExists(t, filepath.Join(root, "mixed-topology"))
}

func TestLoadComparisonRowsRejectsMismatchedCompatibilityEvidence(t *testing.T) {
	root := t.TempDir()
	for trial, fingerprint := range []string{"same-config-a", "different-config-b"} {
		report := benchmarkReport{
			RunID: "evidence-mismatch", Strategy: "hybrid", Entry: "rpc", Scenario: "hot-read-c32",
			Trial: trial + 1, ExpectedTrials: 2,
			Environment: map[string]string{
				"topology": "compose-full", "reader_cardinality": "hot", "cache_state": "l1-warm",
				"compatibility_fingerprint": fingerprint,
			},
			Stages: []stageResult{{Name: "read", Total: 1, Success: 1}},
		}
		_, err := writeReportFiles(root, report)
		require.NoError(t, err)
	}

	_, err := loadComparisonRows(root, "evidence-mismatch")
	require.ErrorContains(t, err, "evidence mismatch")
}

func TestFindCapacityDetectsKneeAndStableStep(t *testing.T) {
	rows := []comparisonRow{
		{Concurrency: 1, SuccessQPS: 100, P95: 5 * time.Millisecond, Complete: true},
		{Concurrency: 8, SuccessQPS: 500, P95: 10 * time.Millisecond, Complete: true},
		{Concurrency: 32, SuccessQPS: 600, P95: 30 * time.Millisecond, Complete: true},
		{Concurrency: 128, SuccessQPS: 610, P95: 150 * time.Millisecond, Complete: true},
	}

	capacity := findCapacity(rows)

	require.Equal(t, 32, capacity.StableConcurrency)
	require.Equal(t, 128, capacity.KneeConcurrency)
	require.Equal(t, 610.0, capacity.PeakSuccessQPS)
}

func TestFindCapacityTreatsFirstFailureAsHardBoundary(t *testing.T) {
	rows := []comparisonRow{
		{Concurrency: 32, SuccessQPS: 10, P95: 3 * time.Second, Complete: true},
		{Concurrency: 40, SuccessQPS: 9.5, P95: 4 * time.Second, Complete: true},
		{Concurrency: 44, SuccessQPS: 2, P95: 4 * time.Second, Failed: 100},
		{Concurrency: 64, SuccessQPS: 20, P95: 2 * time.Second, Complete: true},
	}

	capacity := findCapacity(rows)

	require.Equal(t, 32, capacity.StableConcurrency)
	require.Equal(t, 40, capacity.NoErrorUpperConcurrency)
	require.Equal(t, 44, capacity.FirstErrorConcurrency)
}

func TestFindCapacityExcludesIncompleteNoErrorReports(t *testing.T) {
	rows := []comparisonRow{
		{Concurrency: 32, SuccessQPS: 10, P95: 3 * time.Second, Complete: true},
		{Concurrency: 40, SuccessQPS: 20, P95: 2 * time.Second},
	}

	capacity := findCapacity(rows)

	require.Equal(t, 32, capacity.NoErrorUpperConcurrency)
	require.Equal(t, 10.0, capacity.PeakSuccessQPS)
}

func TestComparisonMarkdownIncludesCapacityTable(t *testing.T) {
	body := comparisonMarkdown("run-42", []comparisonRow{{
		Strategy: "hybrid", Entry: "rpc", Scenario: "hot-read", Stage: "read",
		ReaderCardinality: "hot", CacheState: "l1-warm", Topology: "compose-full",
		Concurrency: 32, Total: 10, Success: 10, SuccessQPS: 500, P95: 10 * time.Millisecond, Complete: true,
	}})

	require.Contains(t, body, "Feed 三策略综合对比")
	require.Contains(t, body, "| hybrid | rpc | hot-read | hot | l1-warm | compose-full | read |")
}

func TestSummarizeCapacitiesUsesTrialMedians(t *testing.T) {
	rows := []comparisonRow{
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 32, Trial: 1, ExpectedTrials: 3, SuccessQPS: 100, P95: 10 * time.Millisecond, Complete: true},
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 32, Trial: 2, ExpectedTrials: 3, SuccessQPS: 110, P95: 11 * time.Millisecond, Complete: true},
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 32, Trial: 3, ExpectedTrials: 3, SuccessQPS: 1000, P95: 100 * time.Millisecond, Complete: true},
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 64, Trial: 1, ExpectedTrials: 3, SuccessQPS: 120, P95: 20 * time.Millisecond, Complete: true},
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 64, Trial: 2, ExpectedTrials: 3, SuccessQPS: 130, P95: 21 * time.Millisecond, Complete: true},
		{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 64, Trial: 3, ExpectedTrials: 3, SuccessQPS: 140, P95: 22 * time.Millisecond, Complete: true},
	}

	capacities := summarizeCapacities(rows)

	require.Len(t, capacities, 1)
	require.Equal(t, 130.0, capacities[0].PeakSuccessQPS)
	require.Equal(t, 21*time.Millisecond, capacities[0].PeakP95)
}

func TestCollapseTrialRowsRequiresExpectedUniqueCompleteTrials(t *testing.T) {
	base := comparisonRow{Strategy: "hybrid", Entry: "rpc", Scenario: "distributed-read", Stage: "read", Concurrency: 128, ExpectedTrials: 3, Complete: true}

	missing := collapseTrialRows([]comparisonRow{
		withTrial(base, 1, 100),
		withTrial(base, 2, 110),
	})
	require.Len(t, missing, 1)
	require.False(t, missing[0].Complete)

	duplicate := collapseTrialRows([]comparisonRow{
		withTrial(base, 1, 100),
		withTrial(base, 1, 105),
		withTrial(base, 2, 110),
	})
	require.Len(t, duplicate, 1)
	require.False(t, duplicate[0].Complete)

	outOfRange := collapseTrialRows([]comparisonRow{
		withTrial(base, 2, 100),
		withTrial(base, 3, 105),
		withTrial(base, 4, 110),
	})
	require.Len(t, outOfRange, 1)
	require.False(t, outOfRange[0].Complete)
}

func TestFindCapacityStopsAtIncompleteEvidenceGap(t *testing.T) {
	rows := []comparisonRow{
		{Concurrency: 32, SuccessQPS: 100, P95: 10 * time.Millisecond, Complete: true},
		{Concurrency: 64, SuccessQPS: 0, P95: 0, Complete: false},
		{Concurrency: 128, SuccessQPS: 1000, P95: 20 * time.Millisecond, Complete: true},
	}

	capacity := findCapacity(rows)

	require.Equal(t, 100.0, capacity.PeakSuccessQPS)
	require.Equal(t, 32, capacity.NoErrorUpperConcurrency)
}

func withTrial(base comparisonRow, trial int, qps float64) comparisonRow {
	base.Trial = trial
	base.SuccessQPS = qps
	return base
}
