package feed

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/metric"
	coreprometheus "github.com/zeromicro/go-zero/core/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FeedStage is a bounded stage label for the personal Feed read path.
type FeedStage string

const (
	StageRelation     FeedStage = "relation"
	StageCounter      FeedStage = "counter"
	StageRoute        FeedStage = "route"
	StageInbox        FeedStage = "inbox"
	StageBigVPipeline FeedStage = "bigv_pipeline"
	StageMergeDedup   FeedStage = "merge_dedup"
	StageHydrate      FeedStage = "hydrate"
	StageCursorDecode FeedStage = "cursor_decode"
	StageCursorSeek   FeedStage = "cursor_seek"
	StageTotal        FeedStage = "total"
	StageUnknown      FeedStage = "unknown"
)

// FeedOutcome is a bounded result label shared by all Feed metrics.
type FeedOutcome string

const (
	OutcomeSuccess  FeedOutcome = "success"
	OutcomeError    FeedOutcome = "error"
	OutcomeTimeout  FeedOutcome = "timeout"
	OutcomeCanceled FeedOutcome = "canceled"
	OutcomeUnknown  FeedOutcome = "unknown"
)

// FeedDependency identifies an external dependency without request IDs.
type FeedDependency string

const (
	DependencyRelation FeedDependency = "relation"
	DependencyCounter  FeedDependency = "counter"
	DependencyRedis    FeedDependency = "redis"
	DependencyMySQL    FeedDependency = "mysql"
	DependencyCache    FeedDependency = "cache"
	DependencyUnknown  FeedDependency = "unknown"
)

// FeedOperation identifies a bounded dependency operation.
type FeedOperation string

const (
	OperationListFollowings          FeedOperation = "list_followings"
	OperationBatchFollowerCounts     FeedOperation = "batch_follower_counts"
	OperationInboxRead               FeedOperation = "inbox_read"
	OperationInboxBigVPipeline       FeedOperation = "inbox_bigv_pipeline"
	OperationBigVPipeline            FeedOperation = "bigv_pipeline"
	OperationFeedItemMGet            FeedOperation = "feed_item_mget"
	OperationFeedItemDB              FeedOperation = "feed_item_db"
	OperationFeedItemCacheWrite      FeedOperation = "feed_item_cache_write"
	OperationPageCacheL2Get          FeedOperation = "page_cache_l2_get"
	OperationPageCacheL2Decode       FeedOperation = "page_cache_l2_decode"
	OperationPageCacheL2Encode       FeedOperation = "page_cache_l2_encode"
	OperationPageCacheL2Set          FeedOperation = "page_cache_l2_set"
	OperationPageCacheSFShared       FeedOperation = "page_cache_singleflight_shared"
	OperationPageCacheRefreshEnqueue FeedOperation = "page_cache_refresh_enqueue"
	OperationPageCacheRefreshLoad    FeedOperation = "page_cache_refresh_load"
	OperationTierRead                FeedOperation = "tier_read"
	OperationTierPromote             FeedOperation = "tier_promote"
	OperationActiveFollowerCounts    FeedOperation = "active_follower_counts"
	OperationUnknown                 FeedOperation = "unknown"
)

// FeedPageCacheSource is a bounded outcome for the fully materialized first
// page cache. It intentionally separates bypass from a cache miss/cold load.
type FeedPageCacheSource string

const (
	PageCacheBypass  FeedPageCacheSource = "bypass"
	PageCacheL1Fresh FeedPageCacheSource = "l1_fresh"
	PageCacheL2Fresh FeedPageCacheSource = "l2_fresh"
	PageCacheL1Stale FeedPageCacheSource = "l1_stale"
	PageCacheL2Stale FeedPageCacheSource = "l2_stale"
	PageCacheMiss    FeedPageCacheSource = "miss"
	PageCacheUnknown FeedPageCacheSource = "unknown"
)

// FeedObserver records bounded Feed read signals. Implementations must not put
// user IDs, post IDs, run IDs, or arbitrary errors into labels.
type FeedObserver interface {
	ObserveStage(stage FeedStage, outcome FeedOutcome, duration time.Duration)
	ObserveDependency(dependency FeedDependency, operation FeedOperation, outcome FeedOutcome)
	ObserveColdCompute(outcome FeedOutcome)
	ObservePageCache(source FeedPageCacheSource, outcome FeedOutcome)
	ObservePageRefreshState(queueDepth, activeWorkers, pendingKeys int)
}

// FeedPaginationMode is deliberately bounded; page numbers and Cursor values
// must never appear in metric labels.
type FeedPaginationMode string

const (
	PaginationModePage    FeedPaginationMode = "page"
	PaginationModeCursor  FeedPaginationMode = "cursor"
	PaginationModeUnknown FeedPaginationMode = "unknown"
)

// FeedPaginationResult separates invalid/unavailable Cursor requests from
// dependency failures without exposing error messages as labels.
type FeedPaginationResult string

const (
	PaginationResultSuccess           FeedPaginationResult = "success"
	PaginationResultInvalid           FeedPaginationResult = "invalid"
	PaginationResultDisabled          FeedPaginationResult = "disabled"
	PaginationResultReaderUnavailable FeedPaginationResult = "reader_unavailable"
	PaginationResultError             FeedPaginationResult = "error"
	PaginationResultTimeout           FeedPaginationResult = "timeout"
	PaginationResultCanceled          FeedPaginationResult = "canceled"
	PaginationResultUnknown           FeedPaginationResult = "unknown"
)

// FeedCursorWork identifies additive per-request pagination work. The legacy
// type and metric names are retained for report compatibility; both page and
// Cursor paths record the common work kinds used by A/B comparisons.
type FeedCursorWork string

const (
	CursorWorkRedisMembers    FeedCursorWork = "redis_members"
	CursorWorkTieMembers      FeedCursorWork = "tie_members"
	CursorWorkMergeCandidates FeedCursorWork = "merge_candidates"
	CursorWorkHydrateIDs      FeedCursorWork = "hydrate_ids"
	CursorWorkBackfillRounds  FeedCursorWork = "backfill_rounds"
	CursorWorkRedisCommands   FeedCursorWork = "redis_commands"
	CursorWorkRedisRoundTrips FeedCursorWork = "redis_roundtrips"
	CursorWorkUnknown         FeedCursorWork = "unknown"
)

// FeedPaginationObserver is optional so existing FeedObserver implementations
// and test fakes remain source compatible.
type FeedPaginationObserver interface {
	ObservePagination(mode FeedPaginationMode, result FeedPaginationResult, duration time.Duration)
	ObserveCursorWork(work FeedCursorWork, value int)
}

type FeedTierObserver interface {
	ObserveTierResolution(evidence TierEvidence, outcome FeedOutcome, fallback bool, mismatch bool)
}

type metricsObserver struct {
	stageTotal               metric.CounterVec
	stageDuration            metric.HistogramVec
	dependencyTotal          metric.CounterVec
	coldTotal                metric.CounterVec
	pageCacheTotal           metric.CounterVec
	paginationTotal          metric.CounterVec
	paginationDuration       metric.HistogramVec
	cursorWorkTotal          metric.CounterVec
	tierResolutionTotal      metric.CounterVec
	pageRefreshQueueDepth    metric.GaugeVec
	pageRefreshActiveWorkers metric.GaugeVec
	pageRefreshPendingKeys   metric.GaugeVec
	pageRefreshMetricsOnce   sync.Once
}

var (
	metricsOnce     sync.Once
	processObserver *metricsObserver
)

// NewMetricsObserver returns nil when disabled so hot paths can skip timing and
// allocation work with a single branch.
func NewMetricsObserver(enabled bool) FeedObserver {
	if !enabled {
		return nil
	}
	metricsOnce.Do(func() {
		processObserver = &metricsObserver{
			stageTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_stage_total",
				Help:      "Personal Feed stage executions by bounded stage and outcome.",
				Labels:    []string{"stage", "outcome"},
			}),
			stageDuration: metric.NewHistogramVec(&metric.HistogramVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_stage_duration_ms",
				Help:      "Personal Feed stage latency in milliseconds.",
				Labels:    []string{"stage", "outcome"},
				Buckets:   []float64{0.1, 0.25, 0.5, 1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000},
			}),
			dependencyTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_dependency_calls_total",
				Help:      "Personal Feed external dependency calls by bounded dependency, operation, and outcome.",
				Labels:    []string{"dependency", "operation", "outcome"},
			}),
			coldTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_cold_compute_total",
				Help:      "Personal Feed requests that execute the uncached computation path.",
				Labels:    []string{"outcome"},
			}),
			pageCacheTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_page_cache_total",
				Help:      "Personal Feed page cache requests by bounded source and outcome.",
				Labels:    []string{"source", "outcome"},
			}),
			paginationTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_pagination_total",
				Help:      "Personal Feed requests by bounded pagination mode and result.",
				Labels:    []string{"mode", "result"},
			}),
			paginationDuration: metric.NewHistogramVec(&metric.HistogramVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_pagination_duration_ms",
				Help:      "Personal Feed request latency by bounded pagination mode and result.",
				Labels:    []string{"mode", "result"},
				Buckets:   []float64{0.1, 0.25, 0.5, 1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000},
			}),
			cursorWorkTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_cursor_work_total",
				Help:      "Additive page and Cursor read work by bounded work kind.",
				Labels:    []string{"work"},
			}),
			tierResolutionTotal: metric.NewCounterVec(&metric.CounterVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_author_tier_resolution_total",
				Help:      "Author tier resolutions by bounded evidence and result labels.",
				Labels:    []string{"evidence", "outcome", "fallback", "mismatch"},
			}),
			pageRefreshQueueDepth: metric.NewGaugeVec(&metric.GaugeVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_page_refresh_queue_depth",
				Help:      "Current number of buffered Personal Feed page refresh jobs.",
			}),
			pageRefreshActiveWorkers: metric.NewGaugeVec(&metric.GaugeVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_page_refresh_active_workers",
				Help:      "Current number of active Personal Feed page refresh workers.",
			}),
			pageRefreshPendingKeys: metric.NewGaugeVec(&metric.GaugeVecOpts{
				Namespace: "zhiguang",
				Subsystem: "knowpost",
				Name:      "feed_page_refresh_pending_keys",
				Help:      "Current number of Personal Feed page keys with queued or active refresh.",
			}),
		}
	})
	return processObserver
}

func (o *metricsObserver) ObserveStage(stage FeedStage, outcome FeedOutcome, duration time.Duration) {
	stage = normalizeStage(stage)
	outcome = normalizeOutcome(outcome)
	o.stageTotal.Inc(string(stage), string(outcome))
	o.stageDuration.ObserveFloat(float64(duration)/float64(time.Millisecond), string(stage), string(outcome))
}

func (o *metricsObserver) ObserveDependency(
	dependency FeedDependency,
	operation FeedOperation,
	outcome FeedOutcome,
) {
	dependency = normalizeDependency(dependency)
	operation = normalizeOperation(operation)
	outcome = normalizeOutcome(outcome)
	o.dependencyTotal.Inc(string(dependency), string(operation), string(outcome))
}

func (o *metricsObserver) ObserveColdCompute(outcome FeedOutcome) {
	o.coldTotal.Inc(string(normalizeOutcome(outcome)))
}

func (o *metricsObserver) ObservePageCache(source FeedPageCacheSource, outcome FeedOutcome) {
	o.ensurePageRefreshMetrics()
	o.pageCacheTotal.Inc(string(normalizePageCacheSource(source)), string(normalizeOutcome(outcome)))
}

func (o *metricsObserver) ObservePageRefreshState(queueDepth, activeWorkers, pendingKeys int) {
	o.ensurePageRefreshMetrics()
	if queueDepth < 0 {
		queueDepth = 0
	}
	if activeWorkers < 0 {
		activeWorkers = 0
	}
	if pendingKeys < 0 {
		pendingKeys = 0
	}
	o.pageRefreshQueueDepth.Set(float64(queueDepth))
	o.pageRefreshActiveWorkers.Set(float64(activeWorkers))
	o.pageRefreshPendingKeys.Set(float64(pendingKeys))
}

func (o *metricsObserver) ObservePagination(
	mode FeedPaginationMode,
	result FeedPaginationResult,
	duration time.Duration,
) {
	mode = normalizePaginationMode(mode)
	result = normalizePaginationResult(result)
	o.paginationTotal.Inc(string(mode), string(result))
	o.paginationDuration.ObserveFloat(float64(duration)/float64(time.Millisecond), string(mode), string(result))
}

func (o *metricsObserver) ObserveCursorWork(work FeedCursorWork, value int) {
	if value <= 0 {
		return
	}
	o.cursorWorkTotal.Add(float64(value), string(normalizeCursorWork(work)))
}

func (o *metricsObserver) ObserveTierResolution(
	evidence TierEvidence,
	outcome FeedOutcome,
	fallback bool,
	mismatch bool,
) {
	o.tierResolutionTotal.Inc(
		string(normalizeTierEvidence(evidence)),
		string(normalizeOutcome(outcome)),
		strconv.FormatBool(fallback),
		strconv.FormatBool(mismatch),
	)
}

func (o *metricsObserver) ensurePageRefreshMetrics() {
	if !coreprometheus.Enabled() {
		return
	}
	o.pageRefreshMetricsOnce.Do(func() {
		o.pageRefreshQueueDepth.Set(0)
		o.pageRefreshActiveWorkers.Set(0)
		o.pageRefreshPendingKeys.Set(0)
	})
}

// RecordStage isolates observer failures from the Feed request path.
func RecordStage(observer FeedObserver, stage FeedStage, outcome FeedOutcome, duration time.Duration) {
	if observer == nil {
		return
	}
	defer ignoreObserverPanic()
	observer.ObserveStage(normalizeStage(stage), normalizeOutcome(outcome), duration)
}

// RecordDependency isolates observer failures from the Feed request path.
func RecordDependency(
	observer FeedObserver,
	dependency FeedDependency,
	operation FeedOperation,
	outcome FeedOutcome,
) {
	if observer == nil {
		return
	}
	defer ignoreObserverPanic()
	observer.ObserveDependency(
		normalizeDependency(dependency),
		normalizeOperation(operation),
		normalizeOutcome(outcome),
	)
}

// RecordColdCompute isolates observer failures from the Feed request path.
func RecordColdCompute(observer FeedObserver, outcome FeedOutcome) {
	if observer == nil {
		return
	}
	defer ignoreObserverPanic()
	observer.ObserveColdCompute(normalizeOutcome(outcome))
}

// RecordPageCache isolates observer failures from the Feed request path.
func RecordPageCache(observer FeedObserver, source FeedPageCacheSource, outcome FeedOutcome) {
	if observer == nil {
		return
	}
	defer ignoreObserverPanic()
	observer.ObservePageCache(normalizePageCacheSource(source), normalizeOutcome(outcome))
}

// RecordPageRefreshState isolates observer failures from refresh workers.
func RecordPageRefreshState(observer FeedObserver, queueDepth, activeWorkers, pendingKeys int) {
	if observer == nil {
		return
	}
	defer ignoreObserverPanic()
	observer.ObservePageRefreshState(queueDepth, activeWorkers, pendingKeys)
}

// RecordPagination forwards only when the observer opts into Cursor metrics.
func RecordPagination(
	observer FeedObserver,
	mode FeedPaginationMode,
	result FeedPaginationResult,
	duration time.Duration,
) {
	extended, ok := observer.(FeedPaginationObserver)
	if !ok || extended == nil {
		return
	}
	defer ignoreObserverPanic()
	extended.ObservePagination(normalizePaginationMode(mode), normalizePaginationResult(result), duration)
}

// RecordCursorWork forwards a bounded work kind and ignores non-positive values.
func RecordCursorWork(observer FeedObserver, work FeedCursorWork, value int) {
	if value <= 0 {
		return
	}
	extended, ok := observer.(FeedPaginationObserver)
	if !ok || extended == nil {
		return
	}
	defer ignoreObserverPanic()
	extended.ObserveCursorWork(normalizeCursorWork(work), value)
}

func RecordTierResolution(
	observer FeedObserver,
	evidence TierEvidence,
	outcome FeedOutcome,
	fallback bool,
	mismatch bool,
) {
	extended, ok := observer.(FeedTierObserver)
	if !ok || extended == nil {
		return
	}
	defer ignoreObserverPanic()
	extended.ObserveTierResolution(
		normalizeTierEvidence(evidence),
		normalizeOutcome(outcome),
		fallback,
		mismatch,
	)
}

func observationStarted(observer FeedObserver) time.Time {
	if observer == nil {
		return time.Time{}
	}
	return time.Now()
}

func recordStageSince(observer FeedObserver, stage FeedStage, outcome FeedOutcome, started time.Time) {
	if observer == nil {
		return
	}
	RecordStage(observer, stage, outcome, time.Since(started))
}

func ignoreObserverPanic() {
	_ = recover()
}

// OutcomeFromError maps errors to a bounded result label.
func OutcomeFromError(err error) FeedOutcome {
	if err == nil {
		return OutcomeSuccess
	}
	if errors.Is(err, context.DeadlineExceeded) || status.Code(err) == codes.DeadlineExceeded {
		return OutcomeTimeout
	}
	if errors.Is(err, context.Canceled) || status.Code(err) == codes.Canceled {
		return OutcomeCanceled
	}
	return OutcomeError
}

func normalizeStage(stage FeedStage) FeedStage {
	switch stage {
	case StageRelation, StageCounter, StageRoute, StageInbox, StageBigVPipeline,
		StageMergeDedup, StageHydrate, StageCursorDecode, StageCursorSeek, StageTotal, StageUnknown:
		return stage
	default:
		return StageUnknown
	}
}

func normalizePaginationMode(mode FeedPaginationMode) FeedPaginationMode {
	switch mode {
	case PaginationModePage, PaginationModeCursor, PaginationModeUnknown:
		return mode
	default:
		return PaginationModeUnknown
	}
}

func normalizePaginationResult(result FeedPaginationResult) FeedPaginationResult {
	switch result {
	case PaginationResultSuccess, PaginationResultInvalid, PaginationResultDisabled,
		PaginationResultReaderUnavailable, PaginationResultError, PaginationResultTimeout,
		PaginationResultCanceled, PaginationResultUnknown:
		return result
	default:
		return PaginationResultUnknown
	}
}

func normalizeCursorWork(work FeedCursorWork) FeedCursorWork {
	switch work {
	case CursorWorkRedisMembers, CursorWorkTieMembers, CursorWorkMergeCandidates,
		CursorWorkHydrateIDs, CursorWorkBackfillRounds, CursorWorkRedisCommands,
		CursorWorkRedisRoundTrips, CursorWorkUnknown:
		return work
	default:
		return CursorWorkUnknown
	}
}

func normalizeOutcome(outcome FeedOutcome) FeedOutcome {
	switch outcome {
	case OutcomeSuccess, OutcomeError, OutcomeTimeout, OutcomeCanceled, OutcomeUnknown:
		return outcome
	default:
		return OutcomeUnknown
	}
}

func normalizeDependency(dependency FeedDependency) FeedDependency {
	switch dependency {
	case DependencyRelation, DependencyCounter, DependencyRedis, DependencyMySQL, DependencyCache, DependencyUnknown:
		return dependency
	default:
		return DependencyUnknown
	}
}

func normalizeOperation(operation FeedOperation) FeedOperation {
	switch operation {
	case OperationListFollowings, OperationBatchFollowerCounts, OperationInboxRead,
		OperationInboxBigVPipeline, OperationBigVPipeline, OperationFeedItemMGet, OperationFeedItemDB,
		OperationFeedItemCacheWrite, OperationPageCacheL2Get, OperationPageCacheL2Decode,
		OperationPageCacheL2Encode, OperationPageCacheL2Set, OperationPageCacheSFShared,
		OperationPageCacheRefreshEnqueue, OperationPageCacheRefreshLoad, OperationTierRead,
		OperationTierPromote, OperationActiveFollowerCounts, OperationUnknown:
		return operation
	default:
		return OperationUnknown
	}
}

func normalizeTierEvidence(evidence TierEvidence) TierEvidence {
	switch evidence {
	case TierEvidencePersisted, TierEvidenceCounter, TierEvidenceMySQL, TierEvidenceUnknown:
		return evidence
	default:
		return TierEvidenceUnknown
	}
}

func normalizePageCacheSource(source FeedPageCacheSource) FeedPageCacheSource {
	switch source {
	case PageCacheBypass, PageCacheL1Fresh, PageCacheL2Fresh, PageCacheL1Stale, PageCacheL2Stale, PageCacheMiss, PageCacheUnknown:
		return source
	default:
		return PageCacheUnknown
	}
}
