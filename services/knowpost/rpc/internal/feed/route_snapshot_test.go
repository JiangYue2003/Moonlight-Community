package feed

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
)

type mutableRelationEpoch struct {
	mu    sync.RWMutex
	value uint64
	err   error
	calls int
}

type cancelAwareRelationEpoch struct {
	started chan<- struct{}
}

func (e cancelAwareRelationEpoch) Relation(ctx context.Context, _ int64) (uint64, error) {
	e.started <- struct{}{}
	<-ctx.Done()
	return 0, ctx.Err()
}

func (e *mutableRelationEpoch) Relation(context.Context, int64) (uint64, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.calls++
	return e.value, e.err
}

func (e *mutableRelationEpoch) set(value uint64, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.value = value
	e.err = err
}

type countingRouteRelation struct {
	mu         sync.Mutex
	followings []int64
	err        error
	calls      int
}

func (r *countingRouteRelation) GetFollowings(context.Context, int64) ([]int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls++
	return append([]int64(nil), r.followings...), r.err
}

func (r *countingRouteRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func (r *countingRouteRelation) set(followings []int64, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.followings = append([]int64(nil), followings...)
	r.err = err
}

func (r *countingRouteRelation) callCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.calls
}

type atomicRouteCounter struct {
	counts map[int64]int64
	err    error
	calls  atomic.Int32
}

func (c *atomicRouteCounter) BatchGetFollowerCounts(context.Context, []int64) (map[int64]int64, error) {
	c.calls.Add(1)
	result := make(map[int64]int64, len(c.counts))
	for userID, count := range c.counts {
		result[userID] = count
	}
	return result, c.err
}

type blockingRouteRelation struct {
	calls    atomic.Int32
	started  chan<- struct{}
	release  <-chan struct{}
	canceled chan<- struct{}
}

type errorCountingLogger struct {
	logx.Logger
	calls atomic.Int32
}

func (l *errorCountingLogger) Errorf(string, ...interface{}) {
	l.calls.Add(1)
}

func (r *blockingRouteRelation) GetFollowings(ctx context.Context, _ int64) ([]int64, error) {
	r.calls.Add(1)
	r.started <- struct{}{}
	select {
	case <-r.release:
		return []int64{201}, nil
	case <-ctx.Done():
		r.canceled <- struct{}{}
		return nil, ctx.Err()
	}
}

func (r *blockingRouteRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func newRouteSnapshotReader(
	strategy Strategy,
	epoch RelationEpochReader,
	cache FeedRouteCache,
	ttl time.Duration,
	relation RelationClient,
	counter CounterClient,
) *FeedReader {
	return NewFeedReaderWithOptions(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:             strategy,
			RouteCache:           cache,
			RouteCacheTTL:        ttl,
			RouteSnapshotEnabled: true,
			RelationEpochs:       epoch,
		},
	)
}

func TestRouteSnapshotSecondPrepareSkipsRelationAndCounter(t *testing.T) {
	epoch := &mutableRelationEpoch{value: 7}
	relation := &countingRouteRelation{followings: []int64{201, 202}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

	first, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)
	second, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)

	require.Equal(t, []int64{201}, first.bigVs)
	require.Equal(t, []int64{201}, second.bigVs)
	require.True(t, second.AllowsCreator(202))
	require.Equal(t, 1, relation.callCount())
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotRealL1IsVisibleBeforeColdPrepareReturns(t *testing.T) {
	cache, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	require.NoError(t, err)
	epoch := &mutableRelationEpoch{value: 7}
	relation := &countingRouteRelation{followings: []int64{201}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, cache, time.Minute, relation, counter)

	_, firstErr := reader.Prepare(context.Background(), 123)
	_, secondErr := reader.Prepare(context.Background(), 123)

	require.NoError(t, firstErr)
	require.NoError(t, secondErr)
	require.Equal(t, 1, relation.callCount())
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotRealL1PreventsPostFlightBurstReload(t *testing.T) {
	const requests = 32
	cache, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	require.NoError(t, err)
	epoch := &mutableRelationEpoch{value: 9}
	started := make(chan struct{}, requests+1)
	release := make(chan struct{})
	relation := &blockingRouteRelation{started: started, release: release, canceled: make(chan struct{}, 1)}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, cache, time.Minute, relation, counter)

	firstDone := make(chan error, 1)
	go func() {
		_, prepareErr := reader.Prepare(context.Background(), 123)
		firstDone <- prepareErr
	}()
	<-started
	close(release)
	require.NoError(t, <-firstDone)

	results := make(chan error, requests)
	for range requests {
		go func() {
			_, prepareErr := reader.Prepare(context.Background(), 123)
			results <- prepareErr
		}()
	}
	for range requests {
		require.NoError(t, <-results)
	}
	require.Equal(t, int32(1), relation.calls.Load())
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotConcurrentCachedReadersOnlyReadSharedState(t *testing.T) {
	const readers = 64
	cache, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	require.NoError(t, err)
	epoch := &mutableRelationEpoch{value: 12}
	relation := &countingRouteRelation{followings: []int64{201, 202}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, cache, time.Minute, relation, counter)

	_, err = reader.Prepare(context.Background(), 123)
	require.NoError(t, err)
	results := make(chan error, readers)
	for range readers {
		go func() {
			snapshot, prepareErr := reader.Prepare(context.Background(), 123)
			if prepareErr == nil && (!snapshot.AllowsCreator(201) || !snapshot.AllowsCreator(202) || snapshot.AllowsCreator(203)) {
				prepareErr = errors.New("shared route snapshot changed")
			}
			results <- prepareErr
		}()
	}
	for range readers {
		require.NoError(t, <-results)
	}
	require.Equal(t, 1, relation.callCount())
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotFeatureFlagOffPreservesLegacyRelationPath(t *testing.T) {
	epoch := &mutableRelationEpoch{value: 7}
	relation := &countingRouteRelation{followings: []int64{201}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:             StrategyHybrid,
			RouteCache:           newLockedRouteCache(),
			RouteCacheTTL:        time.Minute,
			RouteSnapshotEnabled: false,
			RelationEpochs:       epoch,
		},
	)

	_, firstErr := reader.Prepare(context.Background(), 123)
	_, secondErr := reader.Prepare(context.Background(), 123)

	require.NoError(t, firstErr)
	require.NoError(t, secondErr)
	require.Equal(t, 2, relation.callCount(), "legacy partial cache still reads Relation on every request")
	require.Equal(t, int32(1), counter.calls.Load(), "legacy partial cache still reuses classification")
	require.Zero(t, epoch.calls)
}

func TestRouteSnapshotKeyIncludesUserAndEpoch(t *testing.T) {
	require.Equal(t, "feed:route:v2:123:e7", routeSnapshotKey(123, 7))
	require.NotEqual(t, routeSnapshotKey(123, 7), routeSnapshotKey(123, 8))
	require.NotEqual(t, routeSnapshotKey(123, 7), routeSnapshotKey(124, 7))
}

func TestRouteSnapshotEpochChangeForcesCompleteReload(t *testing.T) {
	epoch := &mutableRelationEpoch{value: 1}
	relation := &countingRouteRelation{followings: []int64{201}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1, 202: 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

	first, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)
	require.True(t, first.AllowsCreator(201))

	relation.set([]int64{202}, nil)
	epoch.set(2, nil)
	second, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)

	require.False(t, second.AllowsCreator(201))
	require.True(t, second.AllowsCreator(202))
	require.Equal(t, 2, relation.callCount())
	require.Equal(t, int32(2), counter.calls.Load())
}

func TestRouteSnapshotTTLReclassifiesBigVThreshold(t *testing.T) {
	cache := newExpiringRouteCache()
	epoch := &mutableRelationEpoch{value: 3}
	relation := &countingRouteRelation{followings: []int64{201}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, cache, 20*time.Millisecond, relation, counter)

	first, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)
	counter.counts[201] = BIGV_THRESHOLD + 1
	withinTTL, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)
	time.Sleep(30 * time.Millisecond)
	afterTTL, err := reader.Prepare(context.Background(), 123)
	require.NoError(t, err)

	require.Empty(t, first.bigVs)
	require.Empty(t, withinTTL.bigVs)
	require.Equal(t, []int64{201}, afterTTL.bigVs)
	require.Equal(t, 2, relation.callCount())
	require.Equal(t, int32(2), counter.calls.Load())
}

func TestRouteSnapshotCoalescesCompleteConcurrentMiss(t *testing.T) {
	const requests = 32
	epoch := &mutableRelationEpoch{value: 9}
	started := make(chan struct{}, requests)
	release := make(chan struct{})
	canceled := make(chan struct{}, 1)
	relation := &blockingRouteRelation{started: started, release: release, canceled: canceled}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

	results := make(chan error, requests)
	for range requests {
		go func() {
			snapshot, err := reader.Prepare(context.Background(), 123)
			if err == nil && len(snapshot.bigVs) != 1 {
				err = errors.New("missing big V")
			}
			results <- err
		}()
	}
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("Relation RPC did not start")
	}
	time.Sleep(50 * time.Millisecond)
	require.Equal(t, int32(1), relation.calls.Load())
	close(release)
	for range requests {
		require.NoError(t, <-results)
	}
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotFirstWaiterCancellationDoesNotCancelSharedLookup(t *testing.T) {
	epoch := &mutableRelationEpoch{value: 11}
	started := make(chan struct{}, 2)
	release := make(chan struct{})
	canceled := make(chan struct{}, 1)
	relation := &blockingRouteRelation{started: started, release: release, canceled: canceled}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

	firstCtx, cancelFirst := context.WithCancel(context.Background())
	firstDone := make(chan error, 1)
	secondDone := make(chan error, 1)
	go func() {
		_, err := reader.Prepare(firstCtx, 123)
		firstDone <- err
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("Relation RPC did not start")
	}
	go func() {
		_, err := reader.Prepare(context.Background(), 123)
		secondDone <- err
	}()
	cancelFirst()
	require.ErrorIs(t, <-firstDone, context.Canceled)
	select {
	case <-canceled:
		t.Fatal("first waiter cancellation reached shared Relation RPC")
	case <-time.After(50 * time.Millisecond):
	}
	close(release)
	require.NoError(t, <-secondDone)
	require.Equal(t, int32(1), relation.calls.Load())
	require.Equal(t, int32(1), counter.calls.Load())
}

func TestRouteSnapshotDependencyErrorsAreNotCached(t *testing.T) {
	t.Run("relation", func(t *testing.T) {
		epoch := &mutableRelationEpoch{value: 1}
		relation := &countingRouteRelation{err: errors.New("relation unavailable")}
		reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, &atomicRouteCounter{})

		_, firstErr := reader.Prepare(context.Background(), 123)
		_, secondErr := reader.Prepare(context.Background(), 123)

		require.Error(t, firstErr)
		require.Error(t, secondErr)
		require.Equal(t, 2, relation.callCount())
	})

	t.Run("counter fallback", func(t *testing.T) {
		epoch := &mutableRelationEpoch{value: 1}
		relation := &countingRouteRelation{followings: []int64{201}}
		counter := &scriptedCounterClient{results: []scriptedCounterResult{
			{err: errors.New("counter unavailable")},
			{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}},
		}}
		reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

		first, firstErr := reader.Prepare(context.Background(), 123)
		second, secondErr := reader.Prepare(context.Background(), 123)

		require.NoError(t, firstErr)
		require.NoError(t, secondErr)
		require.Empty(t, first.bigVs)
		require.Equal(t, []int64{201}, second.bigVs)
		require.Equal(t, 2, relation.callCount())
		require.Equal(t, 2, counter.calls)
	})
}

func TestRouteSnapshotEpochFailureBypassesUnverifiableCache(t *testing.T) {
	epoch := &mutableRelationEpoch{err: errors.New("redis unavailable")}
	relation := &countingRouteRelation{followings: []int64{201}}
	counter := &atomicRouteCounter{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}}
	reader := newRouteSnapshotReader(StrategyHybrid, epoch, newLockedRouteCache(), time.Minute, relation, counter)

	first, firstErr := reader.Prepare(context.Background(), 123)
	second, secondErr := reader.Prepare(context.Background(), 123)

	require.NoError(t, firstErr)
	require.NoError(t, secondErr)
	require.Equal(t, []int64{201}, first.bigVs)
	require.Equal(t, []int64{201}, second.bigVs)
	require.Equal(t, 2, relation.callCount())
	require.Equal(t, int32(2), counter.calls.Load())
}

func TestRouteSnapshotEpochFailureLogIsRateLimited(t *testing.T) {
	epoch := &mutableRelationEpoch{err: errors.New("redis unavailable")}
	relation := &countingRouteRelation{followings: []int64{201}}
	logger := &errorCountingLogger{Logger: logx.WithContext(context.Background())}
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		relation,
		&atomicRouteCounter{},
		logger,
		FeedReaderOptions{
			Strategy:             StrategyHybrid,
			RouteCache:           newLockedRouteCache(),
			RouteCacheTTL:        time.Minute,
			RouteSnapshotEnabled: true,
			RelationEpochs:       epoch,
		},
	)

	_, firstErr := reader.Prepare(context.Background(), 123)
	_, secondErr := reader.Prepare(context.Background(), 123)

	require.NoError(t, firstErr)
	require.NoError(t, secondErr)
	require.Equal(t, int32(1), logger.calls.Load())
}

func TestRouteSnapshotRequestCancellationDoesNotEnterEpochBypass(t *testing.T) {
	started := make(chan struct{}, 1)
	relation := &countingRouteRelation{followings: []int64{201}}
	logger := &errorCountingLogger{Logger: logx.WithContext(context.Background())}
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		relation,
		&atomicRouteCounter{},
		logger,
		FeedReaderOptions{
			Strategy:             StrategyHybrid,
			RouteCache:           newLockedRouteCache(),
			RouteCacheTTL:        time.Minute,
			RouteSnapshotEnabled: true,
			RelationEpochs:       cancelAwareRelationEpoch{started: started},
		},
	)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, err := reader.Prepare(ctx, 123)
		done <- err
	}()
	<-started
	cancel()

	require.ErrorIs(t, <-done, context.Canceled)
	require.Zero(t, relation.callCount())
	require.Zero(t, logger.calls.Load())
}

func TestRouteSnapshotPreservesPushAndPullClassification(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		strategy Strategy
		wantBigV []int64
	}{
		{name: "push", strategy: StrategyPush, wantBigV: []int64{}},
		{name: "pull", strategy: StrategyPull, wantBigV: []int64{201, 202}},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			epoch := &mutableRelationEpoch{value: 1}
			relation := &countingRouteRelation{followings: []int64{201, 202}}
			counter := &atomicRouteCounter{}
			reader := newRouteSnapshotReader(testCase.strategy, epoch, newLockedRouteCache(), time.Minute, relation, counter)

			first, err := reader.Prepare(context.Background(), 123)
			require.NoError(t, err)
			second, err := reader.Prepare(context.Background(), 123)
			require.NoError(t, err)

			require.Equal(t, testCase.wantBigV, first.bigVs)
			require.Equal(t, testCase.wantBigV, second.bigVs)
			require.Equal(t, 1, relation.callCount())
			require.Zero(t, counter.calls.Load())
		})
	}
}

type expiringRouteCacheEntry struct {
	value     any
	expiresAt time.Time
}

type expiringRouteCache struct {
	mu     sync.Mutex
	values map[string]expiringRouteCacheEntry
}

type firstSetRejectedRouteCache struct {
	sets  int
	waits int
}

func (c *firstSetRejectedRouteCache) Get(string) (any, bool) { return nil, false }

func (c *firstSetRejectedRouteCache) SetWithTTL(string, any, int64, time.Duration) bool {
	c.sets++
	return c.sets > 1
}

func (c *firstSetRejectedRouteCache) Wait() { c.waits++ }

func TestRouteSnapshotSetBufferRejectionRetriesOnlyOnce(t *testing.T) {
	cache := &firstSetRejectedRouteCache{}
	reader := &FeedReader{routeCache: cache, routeCacheTTL: time.Second}

	reader.publishRouteSnapshot("route", &feedRouteSnapshot{userID: 123})

	require.Equal(t, 2, cache.sets)
	require.Equal(t, 2, cache.waits)
}

func newExpiringRouteCache() *expiringRouteCache {
	return &expiringRouteCache{values: make(map[string]expiringRouteCacheEntry)}
}

func (c *expiringRouteCache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.values[key]
	if !ok || time.Now().After(entry.expiresAt) {
		delete(c.values, key)
		return nil, false
	}
	return entry.value, true
}

func (c *expiringRouteCache) SetWithTTL(key string, value any, _ int64, ttl time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = expiringRouteCacheEntry{value: value, expiresAt: time.Now().Add(ttl)}
	return true
}
