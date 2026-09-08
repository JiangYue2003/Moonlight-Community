package feed

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	coreprometheus "github.com/zeromicro/go-zero/core/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type observerCall struct {
	kind       string
	stage      FeedStage
	dependency FeedDependency
	operation  FeedOperation
	cache      FeedPageCacheSource
	outcome    FeedOutcome
	mode       FeedPaginationMode
	pageResult FeedPaginationResult
	work       FeedCursorWork
	evidence   TierEvidence
	fallback   bool
	mismatch   bool
	value      int
	duration   time.Duration
}

type recordingFeedObserver struct {
	calls         []observerCall
	queueDepth    int
	activeWorkers int
	pendingKeys   int
}

func (o *recordingFeedObserver) ObserveStage(stage FeedStage, outcome FeedOutcome, duration time.Duration) {
	o.calls = append(o.calls, observerCall{kind: "stage", stage: stage, outcome: outcome, duration: duration})
}

func (o *recordingFeedObserver) ObserveDependency(
	dependency FeedDependency,
	operation FeedOperation,
	outcome FeedOutcome,
) {
	o.calls = append(o.calls, observerCall{
		kind:       "dependency",
		dependency: dependency,
		operation:  operation,
		outcome:    outcome,
	})
}

func (o *recordingFeedObserver) ObserveColdCompute(outcome FeedOutcome) {
	o.calls = append(o.calls, observerCall{kind: "cold", outcome: outcome})
}

func (o *recordingFeedObserver) ObservePageCache(source FeedPageCacheSource, outcome FeedOutcome) {
	o.calls = append(o.calls, observerCall{kind: "page_cache", cache: source, outcome: outcome})
}

func (o *recordingFeedObserver) ObservePageRefreshState(queueDepth, activeWorkers, pendingKeys int) {
	o.queueDepth = queueDepth
	o.activeWorkers = activeWorkers
	o.pendingKeys = pendingKeys
}

func (o *recordingFeedObserver) ObservePagination(
	mode FeedPaginationMode,
	result FeedPaginationResult,
	duration time.Duration,
) {
	o.calls = append(o.calls, observerCall{
		kind: "pagination", mode: mode, pageResult: result, duration: duration,
	})
}

func (o *recordingFeedObserver) ObserveCursorWork(work FeedCursorWork, value int) {
	o.calls = append(o.calls, observerCall{kind: "cursor_work", work: work, value: value})
}

func (o *recordingFeedObserver) ObserveTierResolution(
	evidence TierEvidence,
	outcome FeedOutcome,
	fallback bool,
	mismatch bool,
) {
	o.calls = append(o.calls, observerCall{
		kind: "tier", evidence: evidence, outcome: outcome, fallback: fallback, mismatch: mismatch,
	})
}

type panickingFeedObserver struct{}

func (panickingFeedObserver) ObserveStage(FeedStage, FeedOutcome, time.Duration) {
	panic("stage observer failed")
}

func (panickingFeedObserver) ObserveDependency(FeedDependency, FeedOperation, FeedOutcome) {
	panic("dependency observer failed")
}

func (panickingFeedObserver) ObserveColdCompute(FeedOutcome) {
	panic("cold observer failed")
}

func (panickingFeedObserver) ObservePageCache(FeedPageCacheSource, FeedOutcome) {
	panic("page cache observer failed")
}

func (panickingFeedObserver) ObservePageRefreshState(int, int, int) {
	panic("refresh queue observer failed")
}

func (panickingFeedObserver) ObserveTierResolution(TierEvidence, FeedOutcome, bool, bool) {
	panic("tier observer failed")
}

func TestFeedOutcomeFromError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want FeedOutcome
	}{
		{name: "success", want: OutcomeSuccess},
		{name: "error", err: errors.New("boom"), want: OutcomeError},
		{name: "deadline", err: context.DeadlineExceeded, want: OutcomeTimeout},
		{name: "grpc deadline", err: status.Error(codes.DeadlineExceeded, "late"), want: OutcomeTimeout},
		{name: "canceled", err: context.Canceled, want: OutcomeCanceled},
		{name: "grpc canceled", err: status.Error(codes.Canceled, "gone"), want: OutcomeCanceled},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := OutcomeFromError(test.err); got != test.want {
				t.Fatalf("OutcomeFromError(%v) = %q, want %q", test.err, got, test.want)
			}
		})
	}
}

func TestFeedObserverRecordFunctionsForwardFixedSignals(t *testing.T) {
	observer := &recordingFeedObserver{}

	RecordStage(observer, StageRelation, OutcomeSuccess, 5*time.Millisecond)
	RecordDependency(observer, DependencyRedis, OperationInboxRead, OutcomeError)
	RecordColdCompute(observer, OutcomeSuccess)
	RecordPageCache(observer, PageCacheL1Fresh, OutcomeSuccess)
	RecordPageRefreshState(observer, 7, 3, 9)
	RecordPagination(observer, PaginationModeCursor, PaginationResultSuccess, 6*time.Millisecond)
	RecordCursorWork(observer, CursorWorkRedisMembers, 23)
	RecordTierResolution(observer, TierEvidenceMySQL, OutcomeSuccess, true, true)

	if len(observer.calls) != 7 {
		t.Fatalf("calls = %d, want 7", len(observer.calls))
	}
	if got := observer.calls[0]; got.kind != "stage" || got.stage != StageRelation ||
		got.outcome != OutcomeSuccess || got.duration != 5*time.Millisecond {
		t.Fatalf("stage call = %+v", got)
	}
	if got := observer.calls[1]; got.kind != "dependency" || got.dependency != DependencyRedis ||
		got.operation != OperationInboxRead || got.outcome != OutcomeError {
		t.Fatalf("dependency call = %+v", got)
	}
	if got := observer.calls[2]; got.kind != "cold" || got.outcome != OutcomeSuccess {
		t.Fatalf("cold call = %+v", got)
	}
	if got := observer.calls[3]; got.kind != "page_cache" || got.cache != PageCacheL1Fresh || got.outcome != OutcomeSuccess {
		t.Fatalf("page cache call = %+v", got)
	}
	if got := observer.calls[4]; got.kind != "pagination" || got.mode != PaginationModeCursor ||
		got.pageResult != PaginationResultSuccess || got.duration != 6*time.Millisecond {
		t.Fatalf("pagination call = %+v", got)
	}
	if got := observer.calls[5]; got.kind != "cursor_work" || got.work != CursorWorkRedisMembers || got.value != 23 {
		t.Fatalf("cursor work call = %+v", got)
	}
	if got := observer.calls[6]; got.kind != "tier" || got.evidence != TierEvidenceMySQL ||
		got.outcome != OutcomeSuccess || !got.fallback || !got.mismatch {
		t.Fatalf("tier call = %+v", got)
	}
	if observer.queueDepth != 7 {
		t.Fatalf("refresh queue depth = %d, want 7", observer.queueDepth)
	}
	if observer.activeWorkers != 3 || observer.pendingKeys != 9 {
		t.Fatalf("refresh state active/pending = %d/%d, want 3/9", observer.activeWorkers, observer.pendingKeys)
	}
}

func TestFeedObserverRecordFunctionsAreNilSafeAndPanicSafe(t *testing.T) {
	for _, observer := range []FeedObserver{nil, panickingFeedObserver{}} {
		RecordStage(observer, StageTotal, OutcomeSuccess, time.Millisecond)
		RecordDependency(observer, DependencyRelation, OperationListFollowings, OutcomeSuccess)
		RecordColdCompute(observer, OutcomeSuccess)
		RecordPageCache(observer, PageCacheBypass, OutcomeSuccess)
		RecordPageRefreshState(observer, 1, 1, 1)
		RecordPagination(observer, PaginationModeCursor, PaginationResultSuccess, time.Millisecond)
		RecordCursorWork(observer, CursorWorkRedisMembers, 1)
		RecordTierResolution(observer, TierEvidenceCounter, OutcomeSuccess, false, false)
	}
}

func TestDisabledMetricsObserverIsNoop(t *testing.T) {
	if observer := NewMetricsObserver(false); observer != nil {
		t.Fatalf("disabled metrics observer = %T, want nil", observer)
	}
}

func TestMetricsObserverExportsZeroRefreshGaugesOnFirstRequestAfterAgentStarts(t *testing.T) {
	observer := NewMetricsObserver(true)
	if observer == nil {
		t.Fatal("enabled metrics observer is nil")
	}
	// ServiceContext builds the observer before go-zero starts its Prometheus
	// agent. The first Feed request must therefore initialize zero-valued gauges.
	coreprometheus.Enable()
	observer.ObservePageCache(PageCacheBypass, OutcomeSuccess)
	families, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}
	want := map[string]bool{
		"zhiguang_knowpost_feed_page_refresh_queue_depth":    false,
		"zhiguang_knowpost_feed_page_refresh_active_workers": false,
		"zhiguang_knowpost_feed_page_refresh_pending_keys":   false,
	}
	for _, family := range families {
		name := family.GetName()
		if _, ok := want[name]; !ok {
			continue
		}
		if len(family.Metric) == 1 && family.Metric[0].GetGauge().GetValue() == 0 {
			want[name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("zero refresh gauge %s was not exported", name)
		}
	}
}

func TestMetricLabelsNormalizeUnknownValues(t *testing.T) {
	if got := normalizeStage(FeedStage("user-123")); got != StageUnknown {
		t.Fatalf("normalizeStage = %q, want %q", got, StageUnknown)
	}
	if got := normalizeOutcome(FeedOutcome("post-456")); got != OutcomeUnknown {
		t.Fatalf("normalizeOutcome = %q, want %q", got, OutcomeUnknown)
	}
	if got := normalizePageCacheSource(PageCacheL2Stale); got != PageCacheL2Stale {
		t.Fatalf("normalizePageCacheSource = %q, want %q", got, PageCacheL2Stale)
	}
	if got := normalizeDependency(FeedDependency("run-789")); got != DependencyUnknown {
		t.Fatalf("normalizeDependency = %q, want %q", got, DependencyUnknown)
	}
	if got := normalizeOperation(FeedOperation("reader-42")); got != OperationUnknown {
		t.Fatalf("normalizeOperation = %q, want %q", got, OperationUnknown)
	}
	if got := normalizePageCacheSource(FeedPageCacheSource("user-123")); got != PageCacheUnknown {
		t.Fatalf("normalizePageCacheSource = %q, want %q", got, PageCacheUnknown)
	}
	if got := normalizePaginationMode(FeedPaginationMode("user-123")); got != PaginationModeUnknown {
		t.Fatalf("normalizePaginationMode = %q, want %q", got, PaginationModeUnknown)
	}
	if got := normalizePaginationResult(FeedPaginationResult("post-456")); got != PaginationResultUnknown {
		t.Fatalf("normalizePaginationResult = %q, want %q", got, PaginationResultUnknown)
	}
	if got := normalizeCursorWork(FeedCursorWork("run-789")); got != CursorWorkUnknown {
		t.Fatalf("normalizeCursorWork = %q, want %q", got, CursorWorkUnknown)
	}
	if got := normalizeTierEvidence(TierEvidence("author-123")); got != TierEvidenceUnknown {
		t.Fatalf("normalizeTierEvidence = %q, want %q", got, TierEvidenceUnknown)
	}
}
