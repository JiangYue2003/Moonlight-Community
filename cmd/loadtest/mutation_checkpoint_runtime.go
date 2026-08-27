package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type mutationRestoreObservation struct {
	CheckpointID     string        `json:"checkpoint_id"`
	Duration         time.Duration `json:"duration_ns"`
	RemovedPosts     int           `json:"removed_posts"`
	RemovedOutbox    int           `json:"removed_outbox"`
	RemovedPostIDs   []int64       `json:"removed_post_ids,omitempty"`
	RemovedOutboxIDs []int64       `json:"removed_outbox_ids,omitempty"`
	SafetyEpoch      int64         `json:"safety_epoch"`
	Complete         bool          `json:"complete"`
	Error            string        `json:"error,omitempty"`
}

type mutationCheckpointMetadata struct {
	ID                    string    `json:"checkpoint_id"`
	BaselineFingerprint   string    `json:"baseline_fingerprint"`
	CreatedAt             time.Time `json:"created_at"`
	ToolVersion           string    `json:"tool_version"`
	SubjectBinaryIdentity string    `json:"subject_binary_identity"`
}

type mutationCheckpointRestorer interface {
	Metadata() mutationCheckpointMetadata
	Restore(context.Context) (mutationRestoreObservation, error)
	Close() error
}

type mutationTrialObservation struct {
	CheckpointID          string                      `json:"checkpoint_id"`
	BaselineFingerprint   string                      `json:"baseline_fingerprint"`
	CheckpointCreatedAt   time.Time                   `json:"checkpoint_created_at"`
	ToolVersion           string                      `json:"tool_version"`
	SubjectBinaryIdentity string                      `json:"subject_binary_identity"`
	PublishedPostIDs      []int64                     `json:"published_post_ids,omitempty"`
	InitialRestore        *mutationRestoreObservation `json:"initial_restore,omitempty"`
	AfterWarmupRestore    *mutationRestoreObservation `json:"after_warmup_restore,omitempty"`
	AfterTrialRestore     *mutationRestoreObservation `json:"after_trial_restore,omitempty"`
}

type mutationRestoreOperations struct {
	waitKafka          func(context.Context) error
	restoreMySQL       func(context.Context) (mutationMySQLRestoreResult, error)
	deleteDerivedRedis func(context.Context, mutationMySQLRestoreResult) error
	restoreRedis       func(context.Context) (int64, error)
	waitSafetyEpoch    func(context.Context) error
	verifyMySQL        func(context.Context) error
	verifyRedis        func(context.Context) error
}

type mutationCheckpointController struct {
	cfg        benchmarkConfig
	manifest   datasetManifest
	checkpoint mutationCheckpoint
	db         *sql.DB
	redis      *redis.Client
}

const (
	mutationKafkaStableWindow = 2 * time.Second
	mutationKafkaPollInterval = 250 * time.Millisecond
)

type kafkaLagCollector func(context.Context) (map[string]float64, error)

func waitForKafkaStableZero(
	ctx context.Context,
	timeout time.Duration,
	stableWindow time.Duration,
	pollInterval time.Duration,
	collect kafkaLagCollector,
) (map[string]float64, error) {
	if timeout <= 0 || stableWindow <= 0 || pollInterval <= 0 || collect == nil {
		return nil, fmt.Errorf("Kafka stable-zero wait requires positive timing and a collector")
	}
	waitCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var stableSince time.Time
	var stableLogEnd float64
	var last map[string]float64
	var lastErr error
	for {
		values, err := collect(waitCtx)
		now := time.Now()
		if err == nil {
			last = values
			logEnd, hasLogEnd := values["log_end_offset_total"]
			if values["lag_total"] == 0 && hasLogEnd {
				if stableSince.IsZero() || logEnd != stableLogEnd {
					stableSince = now
					stableLogEnd = logEnd
				} else if now.Sub(stableSince) >= stableWindow {
					return values, nil
				}
			} else {
				stableSince = time.Time{}
			}
		} else {
			lastErr = err
			stableSince = time.Time{}
		}

		timer := time.NewTimer(pollInterval)
		select {
		case <-waitCtx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			return nil, fmt.Errorf("Kafka lag did not remain at stable zero: last=%v last_error=%v", last, lastErr)
		case <-timer.C:
		}
	}
}

func waitForMutationKafkaDrain(
	ctx context.Context,
	cfg monitorConfig,
	timeout time.Duration,
) (map[string]float64, error) {
	return waitForKafkaStableZero(
		ctx, timeout, mutationKafkaStableWindow, mutationKafkaPollInterval,
		func(ctx context.Context) (map[string]float64, error) {
			return collectKafkaLag(ctx, cfg)
		},
	)
}

func (controller *mutationCheckpointController) Metadata() mutationCheckpointMetadata {
	return mutationCheckpointMetadata{
		ID:                    controller.checkpoint.ID,
		BaselineFingerprint:   controller.checkpoint.BaselineFingerprint,
		CreatedAt:             controller.checkpoint.CreatedAt,
		ToolVersion:           controller.checkpoint.ToolVersion,
		SubjectBinaryIdentity: controller.checkpoint.SubjectBinaryIdentity,
	}
}

func prepareMutationTrial(
	ctx context.Context,
	restorer mutationCheckpointRestorer,
	warmup func() error,
) (mutationTrialObservation, error) {
	metadata := restorer.Metadata()
	observation := mutationTrialObservation{
		CheckpointID: metadata.ID, BaselineFingerprint: metadata.BaselineFingerprint,
		CheckpointCreatedAt: metadata.CreatedAt, ToolVersion: metadata.ToolVersion,
		SubjectBinaryIdentity: metadata.SubjectBinaryIdentity,
	}
	initial, err := restorer.Restore(ctx)
	observation.InitialRestore = &initial
	if err != nil {
		return observation, fmt.Errorf("restore mutation checkpoint before trial: %w", err)
	}
	if warmup == nil {
		return observation, nil
	}
	if err := warmup(); err != nil {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 4*time.Minute)
		cleanup, cleanupErr := restorer.Restore(cleanupCtx)
		cancelCleanup()
		observation.AfterWarmupRestore = &cleanup
		return observation, errors.Join(
			fmt.Errorf("mutation warmup: %w", err),
			wrapMutationRestoreError("restore after failed mutation warmup", cleanupErr),
		)
	}
	afterWarmup, err := restorer.Restore(ctx)
	observation.AfterWarmupRestore = &afterWarmup
	if err != nil {
		return observation, fmt.Errorf("restore after mutation warmup: %w", err)
	}
	return observation, nil
}

func wrapMutationRestoreError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", operation, err)
}

func writeMutationReportWithRestore(
	ctx context.Context,
	root string,
	report *benchmarkReport,
	restorer mutationCheckpointRestorer,
) (reportPaths, error) {
	return writeMutationReportWithRestoreUsingWriter(ctx, root, report, restorer, writeReportFiles)
}

func writeMutationReportWithRestoreUsingWriter(
	ctx context.Context,
	root string,
	report *benchmarkReport,
	restorer mutationCheckpointRestorer,
	writer func(string, benchmarkReport) (reportPaths, error),
) (reportPaths, error) {
	if report == nil || report.Mutation == nil {
		return reportPaths{}, fmt.Errorf("mutation report is missing checkpoint evidence")
	}
	if writer == nil {
		return reportPaths{}, fmt.Errorf("mutation report writer is required")
	}
	postIDErr := mutationPublishedPostIDError(*report)
	if postIDErr != nil {
		report.MissingMetrics = appendUniqueString(report.MissingMetrics, "mutation_post_ids")
	}
	report.MissingMetrics = appendUniqueString(report.MissingMetrics, "mutation_restore_pending")
	report.Complete = false
	paths, provisionalErr := writer(root, *report)
	if provisionalErr != nil {
		report.MissingMetrics = appendUniqueString(report.MissingMetrics, "mutation_provisional_report")
	}

	afterTrial, restoreErr := restorer.Restore(ctx)
	report.Mutation.AfterTrialRestore = &afterTrial
	report.MissingMetrics = removeString(report.MissingMetrics, "mutation_restore_pending")
	var restoreEvidenceErr error
	if restoreErr == nil && !afterTrial.Complete {
		restoreEvidenceErr = fmt.Errorf("mutation restore reported incomplete without an error")
	}
	if restoreErr != nil || restoreEvidenceErr != nil {
		report.MissingMetrics = appendUniqueString(report.MissingMetrics, "mutation_restore")
	}
	scopeErr := mutationRestoreScopeError(report.Mutation.PublishedPostIDs, afterTrial.RemovedPostIDs)
	if scopeErr != nil {
		report.MissingMetrics = appendUniqueString(report.MissingMetrics, "mutation_restore_scope")
	}
	report.Complete = reportIsComplete(report.Stages, report.MissingMetrics)
	finalPaths, finalWriteErr := writer(root, *report)
	if finalWriteErr == nil {
		paths = finalPaths
	}
	return paths, errors.Join(
		provisionalErr,
		postIDErr,
		wrapMutationRestoreError("restore mutation checkpoint after measured trial", restoreErr),
		restoreEvidenceErr,
		scopeErr,
		finalWriteErr,
	)
}

func writeMutationPreparationReport(
	root string,
	manifest datasetManifest,
	options cliOptions,
	concurrency int,
	trial int,
	observation *mutationTrialObservation,
	cause error,
) (reportPaths, error) {
	missing := []string{"mutation_preparation_pending"}
	notes := []string{"Mutation preparation has not completed."}
	if cause != nil {
		missing = []string{"mutation_preparation"}
		notes = []string{"Mutation preparation error: " + cause.Error()}
	}
	report := benchmarkReport{
		RunID:          effectiveReportRunID(manifest.RunID, options.reportRunID),
		Strategy:       options.strategy,
		Entry:          options.entry,
		Scenario:       fmt.Sprintf("%s-c%d", options.scenario, concurrency),
		StartedAt:      time.Now(),
		Concurrency:    concurrency,
		Trial:          trial,
		ExpectedTrials: normalizedExpectedTrials(options.trials),
		Complete:       false,
		MissingMetrics: missing,
		Mutation:       observation,
		Notes:          notes,
	}
	return writeReportFiles(root, report)
}

func mutationPublishedPostIDError(report benchmarkReport) error {
	if report.Mutation == nil {
		return fmt.Errorf("mutation report has no checkpoint observation")
	}
	var successfulPublishes int64
	for _, stage := range report.Stages {
		if stage.Name == "publish_total" {
			successfulPublishes += stage.Success
		}
	}
	if successfulPublishes != int64(len(report.Mutation.PublishedPostIDs)) {
		return fmt.Errorf("successful publishes=%d recorded post ids=%d", successfulPublishes, len(report.Mutation.PublishedPostIDs))
	}
	seen := make(map[int64]struct{}, len(report.Mutation.PublishedPostIDs))
	for _, postID := range report.Mutation.PublishedPostIDs {
		if postID <= 0 {
			return fmt.Errorf("recorded mutation post id %d must be positive", postID)
		}
		if _, ok := seen[postID]; ok {
			return fmt.Errorf("recorded mutation post id %d is duplicated", postID)
		}
		seen[postID] = struct{}{}
	}
	return nil
}

func mutationRestoreScopeError(published, removed []int64) error {
	removedSet := make(map[int64]struct{}, len(removed))
	for _, postID := range removed {
		removedSet[postID] = struct{}{}
	}
	for _, postID := range published {
		if _, ok := removedSet[postID]; !ok {
			return fmt.Errorf("mutation restore did not remove measured post %d", postID)
		}
	}
	return nil
}

func appendUniqueString(values []string, value string) []string {
	if containsString(values, value) {
		return values
	}
	return append(values, value)
}

func removeString(values []string, target string) []string {
	result := values[:0]
	for _, value := range values {
		if value != target {
			result = append(result, value)
		}
	}
	return result
}

func executeMutationRestore(
	ctx context.Context,
	checkpointID string,
	ops mutationRestoreOperations,
) (mutationRestoreObservation, error) {
	started := time.Now()
	observation := mutationRestoreObservation{CheckpointID: checkpointID}
	fail := func(operation string, err error) (mutationRestoreObservation, error) {
		wrapped := fmt.Errorf("%s: %w", operation, err)
		observation.Duration = time.Since(started)
		observation.Error = wrapped.Error()
		return observation, wrapped
	}
	if ops.waitKafka == nil {
		return fail("wait for Kafka drain", fmt.Errorf("operation is not configured"))
	}
	if err := ops.waitKafka(ctx); err != nil {
		return fail("wait for Kafka drain", err)
	}
	if ops.restoreMySQL == nil {
		return fail("restore MySQL baseline", fmt.Errorf("operation is not configured"))
	}
	mysqlResult, err := ops.restoreMySQL(ctx)
	if err != nil {
		return fail("restore MySQL baseline", err)
	}
	observation.RemovedPosts = len(mysqlResult.PostIDs)
	observation.RemovedOutbox = len(mysqlResult.OutboxIDs)
	observation.RemovedPostIDs = append([]int64(nil), mysqlResult.PostIDs...)
	observation.RemovedOutboxIDs = append([]int64(nil), mysqlResult.OutboxIDs...)
	if ops.deleteDerivedRedis == nil {
		return fail("delete derived Redis state", fmt.Errorf("operation is not configured"))
	}
	if err := ops.deleteDerivedRedis(ctx, mysqlResult); err != nil {
		return fail("delete derived Redis state", err)
	}
	if ops.restoreRedis == nil {
		return fail("restore Redis Feed baseline", fmt.Errorf("operation is not configured"))
	}
	observation.SafetyEpoch, err = ops.restoreRedis(ctx)
	if err != nil {
		return fail("restore Redis Feed baseline", err)
	}
	if ops.waitSafetyEpoch == nil {
		return fail("wait for safety epoch convergence", fmt.Errorf("operation is not configured"))
	}
	if err := ops.waitSafetyEpoch(ctx); err != nil {
		return fail("wait for safety epoch convergence", err)
	}
	if ops.verifyMySQL == nil {
		return fail("verify MySQL baseline", fmt.Errorf("operation is not configured"))
	}
	if err := ops.verifyMySQL(ctx); err != nil {
		return fail("verify MySQL baseline", err)
	}
	if ops.verifyRedis == nil {
		return fail("verify Redis Feed baseline", fmt.Errorf("operation is not configured"))
	}
	if err := ops.verifyRedis(ctx); err != nil {
		return fail("verify Redis Feed baseline", err)
	}
	observation.Duration = time.Since(started)
	observation.Complete = true
	return observation, nil
}

func defaultMutationCheckpointPath(reportDir, runID string) string {
	return filepath.Join(reportDir, "checkpoints", safePathSegment(runID)+".json")
}

func mutationDerivedRedisKeys(result mutationMySQLRestoreResult) []string {
	keys := make(map[string]struct{}, len(result.PostIDs)*4+len(result.OutboxIDs))
	for _, postID := range result.PostIDs {
		for _, key := range []string{
			fmt.Sprintf("agg:v1:knowpost:%d", postID),
			fmt.Sprintf("cache:knowPosts:id:%d", postID),
			fmt.Sprintf("cnt:v1:knowpost:%d", postID),
			fmt.Sprintf("feed:fanout:processing:%d", postID),
		} {
			keys[key] = struct{}{}
		}
	}
	for _, outboxID := range result.OutboxIDs {
		keys[fmt.Sprintf("cache:outbox:id:%d", outboxID)] = struct{}{}
	}
	resultKeys := make([]string, 0, len(keys))
	for key := range keys {
		resultKeys = append(resultKeys, key)
	}
	sort.Strings(resultKeys)
	return resultKeys
}

func createMutationCheckpointRuntime(
	ctx context.Context,
	cfg benchmarkConfig,
	manifest datasetManifest,
	strategy, confirmation, path string,
) (mutationCheckpoint, error) {
	if err := validateMutationManifest(manifest, confirmation); err != nil {
		return mutationCheckpoint{}, err
	}
	if err := validateManifestStrategy(manifest, strategy); err != nil {
		return mutationCheckpoint{}, err
	}
	if strings.TrimSpace(cfg.Monitor.MySQLDSN) == "" {
		return mutationCheckpoint{}, fmt.Errorf("mutation checkpoint requires Monitor.MySQLDSN")
	}
	if _, err := waitForMutationKafkaDrain(ctx, cfg.Monitor, 3*time.Minute); err != nil {
		return mutationCheckpoint{}, fmt.Errorf("create mutation checkpoint requires drained Kafka: %w", err)
	}
	db, redisClient, err := openMutationDataStores(ctx, cfg)
	if err != nil {
		return mutationCheckpoint{}, err
	}
	defer db.Close()
	defer redisClient.Close()
	evidenceCtx, cancelEvidence := context.WithTimeout(ctx, 8*time.Second)
	evidence, evidenceErr := captureRunEnvironmentEvidence(evidenceCtx, cfg, manifest)
	cancelEvidence()
	if evidenceErr != nil {
		return mutationCheckpoint{}, fmt.Errorf("capture mutation checkpoint environment: %w", evidenceErr)
	}
	manifestFingerprint, err := datasetManifestFingerprint(manifest)
	if err != nil {
		return mutationCheckpoint{}, err
	}
	captureStarted := time.Now()
	mysqlSnapshot, err := captureMutationMySQLSnapshot(ctx, db, manifest)
	if err != nil {
		return mutationCheckpoint{}, err
	}
	redisSnapshot, err := captureMutationRedisSnapshot(ctx, redisClient, manifest)
	if err != nil {
		return mutationCheckpoint{}, err
	}
	if _, err := waitForMutationKafkaDrain(ctx, cfg.Monitor, 3*time.Minute); err != nil {
		return mutationCheckpoint{}, fmt.Errorf("verify mutation checkpoint requires drained Kafka: %w", err)
	}
	if err := verifyMutationMySQLSnapshot(ctx, db, manifest, mysqlSnapshot); err != nil {
		return mutationCheckpoint{}, err
	}
	if err := verifyMutationRedisSnapshot(ctx, redisClient, redisSnapshot, time.Since(captureStarted)+5*time.Second); err != nil {
		return mutationCheckpoint{}, err
	}
	checkpoint, err := finalizeMutationCheckpoint(mutationCheckpoint{
		Version: mutationCheckpointVersion, RunID: manifest.RunID, CreatedAt: time.Now().UTC(),
		Strategy: strings.ToLower(strings.TrimSpace(strategy)), Seed: manifest.Seed,
		ManifestFingerprint: manifestFingerprint, ToolVersion: mutationCheckpointToolVersion,
		SubjectBinaryIdentity: evidence["subject_binary_identity"], MySQL: mysqlSnapshot, Redis: redisSnapshot,
	})
	if err != nil {
		return mutationCheckpoint{}, err
	}
	if err := saveMutationCheckpoint(path, checkpoint); err != nil {
		return mutationCheckpoint{}, err
	}
	return checkpoint, nil
}

func openMutationCheckpointController(
	ctx context.Context,
	cfg benchmarkConfig,
	manifest datasetManifest,
	strategy, confirmation, path string,
) (*mutationCheckpointController, error) {
	if err := validateMutationManifest(manifest, confirmation); err != nil {
		return nil, err
	}
	checkpoint, err := loadMutationCheckpoint(path)
	if err != nil {
		return nil, err
	}
	evidenceCtx, cancelEvidence := context.WithTimeout(ctx, 8*time.Second)
	evidence, evidenceErr := captureRunEnvironmentEvidence(evidenceCtx, cfg, manifest)
	cancelEvidence()
	if evidenceErr != nil {
		return nil, fmt.Errorf("capture mutation restore environment: %w", evidenceErr)
	}
	if err := validateMutationCheckpoint(manifest, strategy, evidence["subject_binary_identity"], checkpoint); err != nil {
		return nil, err
	}
	db, redisClient, err := openMutationDataStores(ctx, cfg)
	if err != nil {
		return nil, err
	}
	identity, err := mutationRedisIdentity(ctx, redisClient)
	if err != nil {
		db.Close()
		redisClient.Close()
		return nil, err
	}
	if identity != checkpoint.Redis.Identity {
		db.Close()
		redisClient.Close()
		return nil, fmt.Errorf("mutation checkpoint Redis identity changed: %q != %q", identity, checkpoint.Redis.Identity)
	}
	return &mutationCheckpointController{
		cfg: cfg, manifest: manifest, checkpoint: checkpoint, db: db, redis: redisClient,
	}, nil
}

func (controller *mutationCheckpointController) Restore(ctx context.Context) (mutationRestoreObservation, error) {
	return executeMutationRestore(ctx, controller.checkpoint.ID, mutationRestoreOperations{
		waitKafka: func(ctx context.Context) error {
			_, err := waitForMutationKafkaDrain(ctx, controller.cfg.Monitor, 3*time.Minute)
			return err
		},
		restoreMySQL: func(ctx context.Context) (mutationMySQLRestoreResult, error) {
			return restoreMutationMySQLSnapshot(ctx, controller.db, controller.manifest, controller.checkpoint.MySQL)
		},
		deleteDerivedRedis: func(ctx context.Context, result mutationMySQLRestoreResult) error {
			keys := mutationDerivedRedisKeys(result)
			for start := 0; start < len(keys); start += redisDeleteBatchSize {
				end := minInt(start+redisDeleteBatchSize, len(keys))
				if err := controller.redis.Del(ctx, keys[start:end]...).Err(); err != nil {
					return err
				}
			}
			return nil
		},
		restoreRedis: func(ctx context.Context) (int64, error) {
			return restoreMutationRedisSnapshot(ctx, controller.redis, controller.checkpoint.Redis)
		},
		waitSafetyEpoch: waitForMutationSafetyEpoch,
		verifyMySQL: func(ctx context.Context) error {
			return verifyMutationMySQLSnapshot(ctx, controller.db, controller.manifest, controller.checkpoint.MySQL)
		},
		verifyRedis: func(ctx context.Context) error {
			return verifyMutationRedisSnapshot(ctx, controller.redis, controller.checkpoint.Redis, 5*time.Second)
		},
	})
}

func (controller *mutationCheckpointController) Verify(ctx context.Context) error {
	if _, err := waitForMutationKafkaDrain(ctx, controller.cfg.Monitor, 3*time.Minute); err != nil {
		return err
	}
	if err := verifyMutationMySQLSnapshot(ctx, controller.db, controller.manifest, controller.checkpoint.MySQL); err != nil {
		return err
	}
	return verifyAgedMutationRedisSnapshot(ctx, controller.redis, controller.checkpoint.Redis, 5*time.Second)
}

func (controller *mutationCheckpointController) Close() error {
	var first error
	if controller.db != nil {
		first = controller.db.Close()
	}
	if controller.redis != nil {
		if err := controller.redis.Close(); first == nil {
			first = err
		}
	}
	return first
}

func openMutationDataStores(ctx context.Context, cfg benchmarkConfig) (*sql.DB, *redis.Client, error) {
	if strings.TrimSpace(cfg.Monitor.MySQLDSN) == "" {
		return nil, nil, fmt.Errorf("mutation checkpoint requires Monitor.MySQLDSN")
	}
	db, err := sql.Open("mysql", cfg.Monitor.MySQLDSN)
	if err != nil {
		return nil, nil, fmt.Errorf("open mutation checkpoint MySQL: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("connect mutation checkpoint MySQL: %w", err)
	}
	redisClient := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		db.Close()
		redisClient.Close()
		return nil, nil, fmt.Errorf("connect mutation checkpoint Redis: %w", err)
	}
	return db, redisClient, nil
}

func waitForMutationSafetyEpoch(ctx context.Context) error {
	timer := time.NewTimer(1100 * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
