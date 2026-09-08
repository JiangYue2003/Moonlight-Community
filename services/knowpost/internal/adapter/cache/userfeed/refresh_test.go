package userfeed

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type manualClock struct {
	mu  sync.Mutex
	now time.Time
}

func newManualClock() *manualClock {
	return &manualClock{now: time.Unix(1_700_000_000, 0)}
}

func (c *manualClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *manualClock) Advance(d time.Duration) {
	c.mu.Lock()
	c.now = c.now.Add(d)
	c.mu.Unlock()
}

func TestConcurrentMissSameFullKeyRunsOneLoader(t *testing.T) {
	l2 := &fakeL2{}
	cache := mustCache(t, Config{
		Mode: ModeL2, KeyPrefix: "feed", L2FreshTTL: 5 * time.Second,
		StaleTTL: 10 * time.Second, RefreshWorkers: 2, RefreshQueue: 16,
		LoaderTimeout: time.Second, RandomFloat64: func() float64 { return 0.5 },
	}, nil, l2)
	req := cacheableRequest(nil)
	started := make(chan struct{})
	release := make(chan struct{})
	var loads atomic.Int32
	loader := func(context.Context) (*pb.FeedPage, error) {
		if loads.Add(1) == 1 {
			close(started)
		}
		<-release
		return testPage(1), nil
	}

	const callers = 64
	var wg sync.WaitGroup
	errs := make(chan error, callers)
	wg.Add(callers)
	for i := 0; i < callers; i++ {
		go func() {
			defer wg.Done()
			page, source, err := cache.GetOrLoad(context.Background(), req, loader)
			if err != nil {
				errs <- err
				return
			}
			if page == nil || source != SourceLoad {
				errs <- errors.New("unexpected coalesced result")
			}
		}()
	}
	<-started
	time.Sleep(10 * time.Millisecond)
	close(release)
	wg.Wait()
	close(errs)
	for err := range errs {
		require.NoError(t, err)
	}
	require.Equal(t, int32(1), loads.Load())
	require.Equal(t, 1, l2.gets, "the cold burst must share the L2 lookup as well as the loader")
	require.Equal(t, 1, l2.sets)
}

func TestConcurrentMissDifferentFullKeysRemainIndependent(t *testing.T) {
	cache := newWP9TestCache(t, nil)
	started := make(chan int64, 2)
	release := make(chan struct{})
	loader := func(userID int64) Loader {
		return func(context.Context) (*pb.FeedPage, error) {
			started <- userID
			<-release
			return testPage(1), nil
		}
	}

	var wg sync.WaitGroup
	for _, userID := range []int64{42, 43} {
		userID := userID
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := cacheableRequest(func(r *Request) { r.UserID = userID })
			_, _, _ = cache.GetOrLoad(context.Background(), req, loader(userID))
		}()
	}
	seen := map[int64]bool{}
	for len(seen) < 2 {
		select {
		case userID := <-started:
			seen[userID] = true
		case <-time.After(time.Second):
			t.Fatal("different keys were serialized")
		}
	}
	close(release)
	wg.Wait()
}

func TestCanceledWaiterDoesNotCancelSharedLoader(t *testing.T) {
	cache := newWP9TestCache(t, nil)
	req := cacheableRequest(nil)
	started := make(chan struct{})
	release := make(chan struct{})
	var loads atomic.Int32
	loader := func(ctx context.Context) (*pb.FeedPage, error) {
		require.NoError(t, ctx.Err())
		if loads.Add(1) == 1 {
			close(started)
		}
		select {
		case <-release:
			return testPage(1), nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	firstCtx, cancelFirst := context.WithCancel(context.Background())
	firstDone := make(chan error, 1)
	go func() {
		_, _, err := cache.GetOrLoad(firstCtx, req, loader)
		firstDone <- err
	}()
	<-started

	secondDone := make(chan error, 1)
	go func() {
		page, source, err := cache.GetOrLoad(context.Background(), req, loader)
		if err == nil && (page == nil || source == SourceBypass) {
			err = errors.New("unexpected second waiter result")
		}
		secondDone <- err
	}()
	cancelFirst()
	require.ErrorIs(t, <-firstDone, context.Canceled)
	close(release)
	require.NoError(t, <-secondDone)
	require.Equal(t, int32(1), loads.Load())
}

func TestFreshFinalWindowQueuesOnlyOneRefresh(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.L1FreshTTL = time.Second
		cfg.L2FreshTTL = 5 * time.Second
		cfg.RefreshWorkers = 1
	})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(3300 * time.Millisecond)

	started := make(chan struct{})
	release := make(chan struct{})
	var refreshes atomic.Int32
	loader := func(context.Context) (*pb.FeedPage, error) {
		if refreshes.Add(1) == 1 {
			close(started)
		}
		<-release
		return testPage(2), nil
	}
	for i := 0; i < 32; i++ {
		page, source, getErr := cache.GetOrLoad(context.Background(), req, loader)
		require.NoError(t, getErr)
		require.Equal(t, SourceL2Fresh, source)
		require.Equal(t, int32(1), page.Page)
	}
	<-started
	require.Equal(t, int32(1), refreshes.Load())
	close(release)
}

func TestStaleEntryReturnsImmediatelyAndRefreshesAsynchronously(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.L1FreshTTL = time.Second
		cfg.L2FreshTTL = 5 * time.Second
		cfg.StaleTTL = 10 * time.Second
		cfg.RefreshWorkers = 1
	})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(5100 * time.Millisecond)

	started := make(chan struct{})
	release := make(chan struct{})
	page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		close(started)
		<-release
		return testPage(2), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Stale, source)
	require.Equal(t, int32(1), page.Page)
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("stale hit did not schedule refresh")
	}
	close(release)
}

func TestStaleEntryIsIsolatedByEveryEpochDimension(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
	})
	oldReq := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), oldReq, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(6 * time.Second)

	for _, req := range []Request{
		withRequest(oldReq, func(r *Request) { r.RelationEpoch++ }),
		withRequest(oldReq, func(r *Request) { r.SafetyEpoch++ }),
	} {
		loads := 0
		page, source, getErr := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			loads++
			return testPage(2), nil
		})
		require.NoError(t, getErr)
		require.Equal(t, SourceLoad, source)
		require.Equal(t, int32(2), page.Page)
		require.Equal(t, 1, loads)
	}
}

func TestExpiredRequestJoinsInFlightRefreshWithoutOlderOverwrite(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 1
	})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(3300 * time.Millisecond)

	refreshStarted := make(chan struct{})
	releaseRefresh := make(chan struct{})
	page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		close(refreshStarted)
		<-releaseRefresh
		return testPage(2), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Fresh, source)
	require.Equal(t, int32(1), page.Page)
	<-refreshStarted

	// Cross the old record's stale boundary while the refresh is active. The
	// cold request must join that full-key work instead of starting a second
	// loader that could race its write-back.
	clock.Advance(11 * time.Second)
	var coldLoads atomic.Int32
	coldDone := make(chan *pb.FeedPage, 1)
	coldErr := make(chan error, 1)
	go func() {
		loaded, _, loadErr := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			coldLoads.Add(1)
			return testPage(3), nil
		})
		coldDone <- loaded
		coldErr <- loadErr
	}()
	time.Sleep(20 * time.Millisecond)
	require.Equal(t, int32(0), coldLoads.Load())
	select {
	case <-coldDone:
		t.Fatal("expired request returned before the in-flight full-key refresh")
	default:
	}
	close(releaseRefresh)
	require.NoError(t, <-coldErr)
	require.Equal(t, int32(2), (<-coldDone).Page)
	require.Equal(t, int32(0), coldLoads.Load())

	current, currentSource, err := cache.GetOrLoad(context.Background(), req, loaderMustNotRun(t))
	require.NoError(t, err)
	require.Equal(t, SourceL2Fresh, currentSource)
	require.Equal(t, int32(2), current.Page)
}

func TestExpiredRequestFallsBackToColdLoaderAfterRefreshFailure(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 1
	})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(3300 * time.Millisecond)

	refreshStarted := make(chan struct{})
	releaseRefresh := make(chan struct{})
	releaseNow := closeChannelOnce(releaseRefresh)
	t.Cleanup(releaseNow)
	page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		close(refreshStarted)
		<-releaseRefresh
		return nil, errors.New("refresh failed")
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Fresh, source)
	require.Equal(t, int32(1), page.Page)
	<-refreshStarted
	clock.Advance(11 * time.Second)

	var coldLoads atomic.Int32
	coldDone := make(chan struct{})
	var coldPage *pb.FeedPage
	var coldSource Source
	var coldErr error
	go func() {
		defer close(coldDone)
		coldPage, coldSource, coldErr = cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			coldLoads.Add(1)
			return testPage(3), nil
		})
	}()
	require.Never(t, func() bool { return coldLoads.Load() != 0 }, 20*time.Millisecond, time.Millisecond)
	releaseNow()
	select {
	case <-coldDone:
	case <-time.After(time.Second):
		t.Fatal("expired request did not fall back after refresh failure")
	}
	require.NoError(t, coldErr)
	require.Equal(t, SourceLoad, coldSource)
	require.Equal(t, int32(3), coldPage.Page)
	require.Equal(t, int32(1), coldLoads.Load())
}

func TestExpiredRequestTakesOverQueuedRefreshWithoutWaitingForWorker(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 1
		cfg.RefreshQueue = 2
	})
	workerReq := cacheableRequest(func(r *Request) { r.UserID = 301 })
	queuedReq := cacheableRequest(func(r *Request) { r.UserID = 302 })
	for _, req := range []Request{workerReq, queuedReq} {
		_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			return testPage(1), nil
		})
		require.NoError(t, err)
	}
	clock.Advance(6 * time.Second)

	workerStarted := make(chan struct{})
	releaseWorker := make(chan struct{})
	releaseNow := closeChannelOnce(releaseWorker)
	t.Cleanup(releaseNow)
	_, workerSource, err := cache.GetOrLoad(context.Background(), workerReq, func(context.Context) (*pb.FeedPage, error) {
		close(workerStarted)
		<-releaseWorker
		return testPage(2), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Stale, workerSource)
	<-workerStarted

	var obsoleteRefreshLoads atomic.Int32
	_, queuedSource, err := cache.GetOrLoad(context.Background(), queuedReq, func(context.Context) (*pb.FeedPage, error) {
		obsoleteRefreshLoads.Add(1)
		return testPage(2), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Stale, queuedSource)
	clock.Advance(10 * time.Second)

	var coldLoads atomic.Int32
	started := time.Now()
	page, source, err := cache.GetOrLoad(context.Background(), queuedReq, func(context.Context) (*pb.FeedPage, error) {
		coldLoads.Add(1)
		return testPage(3), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.Equal(t, int32(3), page.Page)
	require.Equal(t, int32(1), coldLoads.Load())
	require.Less(t, time.Since(started), 250*time.Millisecond, "foreground cold load must not wait behind the refresh queue")
	require.Zero(t, obsoleteRefreshLoads.Load())

	releaseNow()
	require.Never(t, func() bool { return obsoleteRefreshLoads.Load() != 0 }, 50*time.Millisecond, time.Millisecond)
}

func TestRefreshQueueFullDoesNotStartUnboundedLoaders(t *testing.T) {
	clock := newManualClock()
	type stateSnapshot struct {
		latest RefreshState
		max    RefreshState
	}
	var stateMu sync.Mutex
	states := stateSnapshot{}
	var reportMu sync.Mutex
	var reports []Operation
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 1
		cfg.RefreshQueue = 1
		cfg.ErrorLogInterval = time.Hour
		cfg.ObserveRefreshState = func(state RefreshState) {
			stateMu.Lock()
			defer stateMu.Unlock()
			states.latest = state
			if state.QueueDepth > states.max.QueueDepth {
				states.max.QueueDepth = state.QueueDepth
			}
			if state.ActiveWorkers > states.max.ActiveWorkers {
				states.max.ActiveWorkers = state.ActiveWorkers
			}
			if state.PendingKeys > states.max.PendingKeys {
				states.max.PendingKeys = state.PendingKeys
			}
		}
		cfg.ReportError = func(operation Operation, _ error) {
			reportMu.Lock()
			reports = append(reports, operation)
			reportMu.Unlock()
		}
	})
	requests := []Request{
		cacheableRequest(func(r *Request) { r.UserID = 41 }),
		cacheableRequest(func(r *Request) { r.UserID = 42 }),
		cacheableRequest(func(r *Request) { r.UserID = 43 }),
		cacheableRequest(func(r *Request) { r.UserID = 44 }),
		cacheableRequest(func(r *Request) { r.UserID = 45 }),
	}
	for _, req := range requests {
		_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			return testPage(1), nil
		})
		require.NoError(t, err)
	}
	clock.Advance(6 * time.Second)

	started := make(chan int64, len(requests))
	release := make(chan struct{})
	releaseNow := closeChannelOnce(release)
	t.Cleanup(releaseNow)
	refresh := func(req Request) {
		page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			started <- req.UserID
			<-release
			return testPage(2), nil
		})
		require.NoError(t, err)
		require.Equal(t, SourceL2Stale, source)
		require.Equal(t, int32(1), page.Page)
	}
	refresh(requests[0])
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("refresh worker did not start")
	}
	for _, req := range requests[1:] {
		refresh(req)
	}
	time.Sleep(20 * time.Millisecond)
	require.Len(t, started, 0, "queue overflow must not spawn another loader")
	stateMu.Lock()
	maximum := states.max
	stateMu.Unlock()
	require.Equal(t, RefreshState{QueueDepth: 1, ActiveWorkers: 1, PendingKeys: 2}, maximum)
	reportMu.Lock()
	reported := append([]Operation(nil), reports...)
	reportMu.Unlock()
	require.Equal(t, []Operation{OperationRefreshEnqueue}, reported, "repeated queue overflow must be rate limited")

	coldReq := cacheableRequest(func(r *Request) { r.UserID = 46 })
	cold, coldSource, coldErr := cache.GetOrLoad(context.Background(), coldReq, func(context.Context) (*pb.FeedPage, error) {
		return testPage(3), nil
	})
	require.NoError(t, coldErr)
	require.Equal(t, SourceLoad, coldSource)
	require.Equal(t, int32(3), cold.Page, "a full refresh queue must not block a true cold miss")
	releaseNow()
	require.Eventually(t, func() bool {
		stateMu.Lock()
		defer stateMu.Unlock()
		return states.latest == (RefreshState{})
	}, time.Second, time.Millisecond, "refresh gauges must converge to zero")
}

func TestRefreshLoadErrorReportingIsRateLimitedPerOperation(t *testing.T) {
	clock := newManualClock()
	var reportMu sync.Mutex
	var reports []Operation
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 1
		cfg.RefreshQueue = 4
		cfg.ErrorLogInterval = time.Hour
		cfg.ReportError = func(operation Operation, _ error) {
			reportMu.Lock()
			reports = append(reports, operation)
			reportMu.Unlock()
		}
	})
	requests := make([]Request, 3)
	for i := range requests {
		userID := int64(200 + i)
		requests[i] = cacheableRequest(func(r *Request) { r.UserID = userID })
		_, _, err := cache.GetOrLoad(context.Background(), requests[i], func(context.Context) (*pb.FeedPage, error) {
			return testPage(1), nil
		})
		require.NoError(t, err)
	}
	clock.Advance(6 * time.Second)

	var loads atomic.Int32
	for _, req := range requests {
		page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			loads.Add(1)
			return nil, errors.New("refresh failed")
		})
		require.NoError(t, err)
		require.Equal(t, SourceL2Stale, source)
		require.Equal(t, int32(1), page.Page)
	}
	require.Eventually(t, func() bool { return loads.Load() == int32(len(requests)) }, time.Second, time.Millisecond)
	reportMu.Lock()
	reported := append([]Operation(nil), reports...)
	reportMu.Unlock()
	require.Equal(t, []Operation{OperationRefreshLoad}, reported)
}

func TestRefreshWorkerConcurrencyIsBounded(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
		cfg.RefreshWorkers = 2
		cfg.RefreshQueue = 8
	})
	requests := make([]Request, 6)
	for i := range requests {
		requests[i] = cacheableRequest(func(r *Request) { r.UserID = int64(100 + i) })
		_, _, err := cache.GetOrLoad(context.Background(), requests[i], func(context.Context) (*pb.FeedPage, error) {
			return testPage(1), nil
		})
		require.NoError(t, err)
	}
	clock.Advance(6 * time.Second)

	release := make(chan struct{})
	var active atomic.Int32
	var maximum atomic.Int32
	for _, req := range requests {
		req := req
		_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
			current := active.Add(1)
			for current > maximum.Load() && !maximum.CompareAndSwap(maximum.Load(), current) {
			}
			<-release
			active.Add(-1)
			return testPage(2), nil
		})
		require.NoError(t, err)
	}
	require.Eventually(t, func() bool { return active.Load() == 2 }, time.Second, time.Millisecond)
	require.LessOrEqual(t, maximum.Load(), int32(2))
	close(release)
}

func TestDetachedLoaderUsesBoundedContext(t *testing.T) {
	cache := newWP9TestCache(t, func(cfg *Config) { cfg.LoaderTimeout = 20 * time.Millisecond })
	started := time.Now()
	_, source, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), func(ctx context.Context) (*pb.FeedPage, error) {
		<-ctx.Done()
		return nil, ctx.Err()
	})
	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Equal(t, SourceLoad, source)
	require.Less(t, time.Since(started), time.Second)
}

func TestExpiredStaleWindowFallsBackToSynchronousLoad(t *testing.T) {
	clock := newManualClock()
	cache := newWP9TestCache(t, func(cfg *Config) {
		cfg.Clock = clock
		cfg.Mode = ModeL2
	})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(16 * time.Second)

	loads := 0
	page, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		loads++
		return testPage(2), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.Equal(t, int32(2), page.Page)
	require.Equal(t, 1, loads)
}

func TestCloseCancelsInFlightRefreshAndStopsWorkers(t *testing.T) {
	clock := newManualClock()
	cache := mustCache(t, Config{
		Mode: ModeL2, KeyPrefix: "feed", L2FreshTTL: 5 * time.Second, StaleTTL: 10 * time.Second,
		RefreshWorkers: 1, RefreshQueue: 1, LoaderTimeout: time.Second, Clock: clock,
		RandomFloat64: func() float64 { return 0.5 },
	}, nil, &fakeL2{})
	req := cacheableRequest(nil)
	_, _, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	clock.Advance(6 * time.Second)

	started := make(chan struct{})
	canceled := make(chan struct{})
	page, source, err := cache.GetOrLoad(context.Background(), req, func(ctx context.Context) (*pb.FeedPage, error) {
		close(started)
		<-ctx.Done()
		close(canceled)
		return nil, ctx.Err()
	})
	require.NoError(t, err)
	require.Equal(t, SourceL2Stale, source)
	require.Equal(t, int32(1), page.Page)
	<-started

	done := make(chan error, 1)
	go func() { done <- cache.Close() }()
	select {
	case <-canceled:
	case <-time.After(time.Second):
		t.Fatal("Close did not cancel the refresh loader")
	}
	select {
	case closeErr := <-done:
		require.NoError(t, closeErr)
	case <-time.After(time.Second):
		t.Fatal("Close did not stop refresh workers")
	}
}

func TestJitterUsesConfiguredSymmetricBounds(t *testing.T) {
	cache := &pageCache{cfg: Config{JitterPercent: 20, RandomFloat64: func() float64 { return 0 }}}
	require.Equal(t, 3200*time.Millisecond, cache.jitter(4*time.Second))
	cache.cfg.RandomFloat64 = func() float64 { return 1 }
	require.Equal(t, 4800*time.Millisecond, cache.jitter(4*time.Second))
}

func closeChannelOnce(ch chan struct{}) func() {
	var once sync.Once
	return func() {
		once.Do(func() { close(ch) })
	}
}

func newWP9TestCache(t *testing.T, mutate func(*Config)) Cache {
	t.Helper()
	cfg := Config{
		Mode:           ModeL1L2,
		KeyPrefix:      "feed",
		L1FreshTTL:     time.Second,
		L2FreshTTL:     5 * time.Second,
		StaleTTL:       10 * time.Second,
		JitterPercent:  20,
		RefreshWorkers: 2,
		RefreshQueue:   16,
		LoaderTimeout:  time.Second,
		RandomFloat64:  func() float64 { return 0.5 },
	}
	if mutate != nil {
		mutate(&cfg)
	}
	cache := mustCache(t, cfg, &fakeL1{}, &fakeL2{})
	t.Cleanup(func() { require.NoError(t, cache.Close()) })
	return cache
}
