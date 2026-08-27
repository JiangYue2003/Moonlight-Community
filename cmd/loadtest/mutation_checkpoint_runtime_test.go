package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func loadBenchmarkReport(path string) (benchmarkReport, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return benchmarkReport{}, err
	}
	var report benchmarkReport
	err = json.Unmarshal(body, &report)
	return report, err
}

type fakeMutationCheckpointRestorer struct {
	metadata mutationCheckpointMetadata
	results  []mutationRestoreObservation
	errors   []error
	calls    int
	restore  func(context.Context, int) (mutationRestoreObservation, error)
}

func (fake *fakeMutationCheckpointRestorer) Metadata() mutationCheckpointMetadata {
	return fake.metadata
}

func (fake *fakeMutationCheckpointRestorer) Restore(ctx context.Context) (mutationRestoreObservation, error) {
	index := fake.calls
	fake.calls++
	if fake.restore != nil {
		return fake.restore(ctx, index)
	}
	if index >= len(fake.results) {
		return mutationRestoreObservation{}, errors.New("unexpected restore")
	}
	return fake.results[index], fake.errors[index]
}

func (fake *fakeMutationCheckpointRestorer) Close() error { return nil }

func TestExecuteMutationRestoreRunsFailClosedSequence(t *testing.T) {
	order := make([]string, 0, 7)
	ops := mutationRestoreOperations{
		waitKafka: func(context.Context) error {
			order = append(order, "kafka")
			return nil
		},
		restoreMySQL: func(context.Context) (mutationMySQLRestoreResult, error) {
			order = append(order, "mysql-restore")
			return mutationMySQLRestoreResult{PostIDs: []int64{30}, OutboxIDs: []int64{300}}, nil
		},
		deleteDerivedRedis: func(_ context.Context, result mutationMySQLRestoreResult) error {
			order = append(order, "redis-derived-delete")
			require.Equal(t, []int64{30}, result.PostIDs)
			return nil
		},
		restoreRedis: func(context.Context) (int64, error) {
			order = append(order, "redis-restore")
			return 9, nil
		},
		waitSafetyEpoch: func(context.Context) error {
			order = append(order, "safety-wait")
			return nil
		},
		verifyMySQL: func(context.Context) error {
			order = append(order, "mysql-verify")
			return nil
		},
		verifyRedis: func(context.Context) error {
			order = append(order, "redis-verify")
			return nil
		},
	}

	observation, err := executeMutationRestore(context.Background(), "checkpoint-1", ops)

	require.NoError(t, err)
	require.True(t, observation.Complete)
	require.Equal(t, int64(9), observation.SafetyEpoch)
	require.Equal(t, 1, observation.RemovedPosts)
	require.Equal(t, 1, observation.RemovedOutbox)
	require.Equal(t, []string{
		"kafka", "mysql-restore", "redis-derived-delete", "redis-restore",
		"safety-wait", "mysql-verify", "redis-verify",
	}, order)
}

func TestExecuteMutationRestoreStopsAfterFirstFailure(t *testing.T) {
	order := make([]string, 0, 4)
	ops := mutationRestoreOperations{
		waitKafka: func(context.Context) error {
			order = append(order, "kafka")
			return nil
		},
		restoreMySQL: func(context.Context) (mutationMySQLRestoreResult, error) {
			order = append(order, "mysql-restore")
			return mutationMySQLRestoreResult{}, nil
		},
		deleteDerivedRedis: func(context.Context, mutationMySQLRestoreResult) error {
			order = append(order, "redis-derived-delete")
			return errors.New("redis unavailable")
		},
		restoreRedis: func(context.Context) (int64, error) {
			order = append(order, "must-not-run")
			return 0, nil
		},
	}

	observation, err := executeMutationRestore(context.Background(), "checkpoint-1", ops)

	require.ErrorContains(t, err, "delete derived Redis state")
	require.False(t, observation.Complete)
	require.Contains(t, observation.Error, "redis unavailable")
	require.Equal(t, []string{"kafka", "mysql-restore", "redis-derived-delete"}, order)
}

func TestPrepareMutationTrialRestoresBeforeAndAfterWarmup(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1", BaselineFingerprint: "baseline-1"},
		results: []mutationRestoreObservation{
			{CheckpointID: "checkpoint-1", Complete: true},
			{CheckpointID: "checkpoint-1", Complete: true},
		},
		errors: make([]error, 2),
	}
	warmups := 0

	observation, err := prepareMutationTrial(context.Background(), fake, func() error {
		warmups++
		return nil
	})

	require.NoError(t, err)
	require.Equal(t, 1, warmups)
	require.Equal(t, 2, fake.calls)
	require.Equal(t, "checkpoint-1", observation.CheckpointID)
	require.Equal(t, "baseline-1", observation.BaselineFingerprint)
	require.NotNil(t, observation.InitialRestore)
	require.NotNil(t, observation.AfterWarmupRestore)
}

func TestPrepareMutationTrialStopsWhenWarmupRestoreFails(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		results: []mutationRestoreObservation{
			{CheckpointID: "checkpoint-1", Complete: true},
			{CheckpointID: "checkpoint-1", Error: "redis unavailable"},
		},
		errors: []error{nil, errors.New("redis unavailable")},
	}

	observation, err := prepareMutationTrial(context.Background(), fake, func() error { return nil })

	require.ErrorContains(t, err, "restore after mutation warmup")
	require.NotNil(t, observation.AfterWarmupRestore)
	require.False(t, observation.AfterWarmupRestore.Complete)
}

func TestPrepareMutationTrialAttemptsCleanupAfterWarmupCancelsParent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		restore: func(restoreCtx context.Context, call int) (mutationRestoreObservation, error) {
			if call == 1 && restoreCtx.Err() != nil {
				return mutationRestoreObservation{CheckpointID: "checkpoint-1"}, restoreCtx.Err()
			}
			return mutationRestoreObservation{CheckpointID: "checkpoint-1", Complete: true}, nil
		},
	}

	observation, err := prepareMutationTrial(ctx, fake, func() error {
		cancel()
		return errors.New("warmup failed")
	})

	require.ErrorContains(t, err, "warmup failed")
	require.Equal(t, 2, fake.calls)
	require.NotNil(t, observation.AfterWarmupRestore)
	require.True(t, observation.AfterWarmupRestore.Complete)
}

func TestWriteMutationReportRestoresAfterProvisionalReport(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1", BaselineFingerprint: "baseline-1"},
		results: []mutationRestoreObservation{{
			CheckpointID: "checkpoint-1", Complete: true, RemovedPostIDs: []int64{101},
		}},
		errors: []error{nil},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages: []stageResult{{Name: "publish_total", Total: 1, Success: 1}},
		Mutation: &mutationTrialObservation{
			CheckpointID: "checkpoint-1", BaselineFingerprint: "baseline-1", PublishedPostIDs: []int64{101},
		},
	}
	root := t.TempDir()

	paths, err := writeMutationReportWithRestore(context.Background(), root, &report, fake)

	require.NoError(t, err)
	require.Equal(t, 1, fake.calls)
	stored, err := loadBenchmarkReport(filepath.Clean(paths.JSON))
	require.NoError(t, err)
	require.True(t, stored.Complete)
	require.NotNil(t, stored.Mutation.AfterTrialRestore)
	require.True(t, stored.Mutation.AfterTrialRestore.Complete)
	require.NotContains(t, stored.MissingMetrics, "mutation_restore_pending")
}

func TestWriteMutationReportCannotBecomeCompleteWhenProvisionalWriteFails(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		results: []mutationRestoreObservation{{
			CheckpointID: "checkpoint-1", Complete: true, RemovedPostIDs: []int64{101},
		}},
		errors: []error{nil},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages: []stageResult{{Name: "publish_total", Total: 1, Success: 1}},
		Mutation: &mutationTrialObservation{
			CheckpointID: "checkpoint-1", PublishedPostIDs: []int64{101},
		},
	}
	root := t.TempDir()
	writes := 0
	writer := func(root string, report benchmarkReport) (reportPaths, error) {
		writes++
		if writes == 1 {
			return reportPaths{}, errors.New("disk full before pending report")
		}
		return writeReportFiles(root, report)
	}

	paths, err := writeMutationReportWithRestoreUsingWriter(context.Background(), root, &report, fake, writer)

	require.ErrorContains(t, err, "disk full before pending report")
	require.Equal(t, 1, fake.calls)
	stored, loadErr := loadBenchmarkReport(paths.JSON)
	require.NoError(t, loadErr)
	require.False(t, stored.Complete)
	require.Contains(t, stored.MissingMetrics, "mutation_provisional_report")
}

func TestWriteMutationPreparationReportPersistsPendingAndFailure(t *testing.T) {
	root := t.TempDir()
	manifest := datasetManifest{RunID: "run-1"}
	options := cliOptions{
		reportRunID: "report-1", strategy: "hybrid", entry: "rpc", scenario: "mixed-90-10", trials: 3,
	}
	observation := &mutationTrialObservation{CheckpointID: "checkpoint-1"}

	paths, err := writeMutationPreparationReport(root, manifest, options, 32, 2, observation, nil)
	require.NoError(t, err)
	pending, err := loadBenchmarkReport(paths.JSON)
	require.NoError(t, err)
	require.False(t, pending.Complete)
	require.Contains(t, pending.MissingMetrics, "mutation_preparation_pending")

	paths, err = writeMutationPreparationReport(
		root, manifest, options, 32, 2, observation, errors.New("warmup restore failed"),
	)
	require.NoError(t, err)
	failed, err := loadBenchmarkReport(paths.JSON)
	require.NoError(t, err)
	require.False(t, failed.Complete)
	require.Contains(t, failed.MissingMetrics, "mutation_preparation")
	require.Contains(t, failed.Notes, "Mutation preparation error: warmup restore failed")
}

func TestWriteMutationReportLeavesIncompleteEvidenceWhenRestoreFails(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		results:  []mutationRestoreObservation{{CheckpointID: "checkpoint-1", Error: "redis unavailable"}},
		errors:   []error{errors.New("redis unavailable")},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages:   []stageResult{{Name: "publish_total", Total: 1, Success: 1}},
		Mutation: &mutationTrialObservation{CheckpointID: "checkpoint-1", PublishedPostIDs: []int64{101}},
	}
	root := t.TempDir()

	paths, err := writeMutationReportWithRestore(context.Background(), root, &report, fake)

	require.ErrorContains(t, err, "restore mutation checkpoint after measured trial")
	stored, loadErr := loadBenchmarkReport(filepath.Clean(paths.JSON))
	require.NoError(t, loadErr)
	require.False(t, stored.Complete)
	require.Contains(t, stored.MissingMetrics, "mutation_restore")
	require.NotNil(t, stored.Mutation.AfterTrialRestore)
}

func TestWriteMutationReportRejectsIncompleteRestoreWithoutExplicitError(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		results:  []mutationRestoreObservation{{CheckpointID: "checkpoint-1", Complete: false}},
		errors:   []error{nil},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages:   []stageResult{{Name: "publish_total", Total: 1, Success: 1}},
		Mutation: &mutationTrialObservation{CheckpointID: "checkpoint-1", PublishedPostIDs: []int64{101}},
	}

	paths, err := writeMutationReportWithRestore(context.Background(), t.TempDir(), &report, fake)

	require.ErrorContains(t, err, "reported incomplete")
	stored, loadErr := loadBenchmarkReport(paths.JSON)
	require.NoError(t, loadErr)
	require.False(t, stored.Complete)
	require.Contains(t, stored.MissingMetrics, "mutation_restore")
}

func TestWriteMutationReportRejectsRestoreThatMissesPublishedPost(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		results: []mutationRestoreObservation{{
			CheckpointID: "checkpoint-1", Complete: true,
			RemovedPostIDs: []int64{101},
		}},
		errors: []error{nil},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages: []stageResult{{Name: "publish_total", Total: 2, Success: 2}},
		Mutation: &mutationTrialObservation{
			CheckpointID: "checkpoint-1", PublishedPostIDs: []int64{101, 102},
		},
	}

	paths, err := writeMutationReportWithRestore(context.Background(), t.TempDir(), &report, fake)

	require.ErrorContains(t, err, "did not remove measured post 102")
	stored, loadErr := loadBenchmarkReport(paths.JSON)
	require.NoError(t, loadErr)
	require.False(t, stored.Complete)
	require.Contains(t, stored.MissingMetrics, "mutation_restore_scope")
}

func TestWriteMutationReportRequiresOneRecordedIDPerSuccessfulPublish(t *testing.T) {
	fake := &fakeMutationCheckpointRestorer{
		metadata: mutationCheckpointMetadata{ID: "checkpoint-1"},
		results: []mutationRestoreObservation{{
			CheckpointID: "checkpoint-1", Complete: true, RemovedPostIDs: []int64{101},
		}},
		errors: []error{nil},
	}
	report := benchmarkReport{
		RunID: "run-1", Strategy: "hybrid", Entry: "rpc", Scenario: "publish-c32",
		StartedAt: time.Now(), Duration: time.Second, Concurrency: 32,
		Stages: []stageResult{{Name: "publish_total", Total: 2, Success: 2}},
		Mutation: &mutationTrialObservation{
			CheckpointID: "checkpoint-1", PublishedPostIDs: []int64{101},
		},
	}

	paths, err := writeMutationReportWithRestore(context.Background(), t.TempDir(), &report, fake)

	require.ErrorContains(t, err, "successful publishes=2 recorded post ids=1")
	stored, loadErr := loadBenchmarkReport(paths.JSON)
	require.NoError(t, loadErr)
	require.False(t, stored.Complete)
	require.Contains(t, stored.MissingMetrics, "mutation_post_ids")
}

func TestWaitForKafkaStableZeroResetsWhenAsyncProducerAdvancesLogEnd(t *testing.T) {
	calls := 0
	collector := func(context.Context) (map[string]float64, error) {
		calls++
		logEnd := float64(10)
		if calls >= 3 {
			logEnd = 11
		}
		return map[string]float64{
			"lag_total": 0, "log_end_offset_total": logEnd, "current_offset_total": logEnd,
		}, nil
	}

	metrics, err := waitForKafkaStableZero(
		context.Background(), 200*time.Millisecond, 15*time.Millisecond, 5*time.Millisecond, collector,
	)

	require.NoError(t, err)
	require.Equal(t, float64(11), metrics["log_end_offset_total"])
	require.GreaterOrEqual(t, calls, 6)
}

func TestWaitForKafkaStableZeroTimesOutWhenLagNeverClears(t *testing.T) {
	collector := func(context.Context) (map[string]float64, error) {
		return map[string]float64{"lag_total": 1, "log_end_offset_total": 11}, nil
	}

	_, err := waitForKafkaStableZero(
		context.Background(), 25*time.Millisecond, 10*time.Millisecond, 5*time.Millisecond, collector,
	)

	require.ErrorContains(t, err, "stable zero")
}
