package feedepoch

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
)

func TestMissingEpochsReturnZeroAndRelationUsesIndependentL1(t *testing.T) {
	server := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	l1 := newTestL1(t)
	store := mustTestStore(t, client, l1, Config{KeyPrefix: "test:run-1:feed", RelationL1TTL: 20 * time.Millisecond})

	relation, err := store.Relation(context.Background(), 42)
	require.NoError(t, err)
	require.Zero(t, relation)
	safety, err := store.Safety(context.Background())
	require.NoError(t, err)
	require.Zero(t, safety)

	l1.Wait()
	require.NoError(t, client.Set(context.Background(), relationKey("test:run-1:feed", 42), 9, 0).Err())
	relation, err = store.Relation(context.Background(), 42)
	require.NoError(t, err)
	require.Zero(t, relation, "the relation L1 should hide later Redis changes until expiry")
	time.Sleep(30 * time.Millisecond)
	relation, err = store.Relation(context.Background(), 42)
	require.NoError(t, err)
	require.Equal(t, uint64(9), relation)
}

func TestBumpRelationIsAtomicConcurrentAndImmediatelyVisible(t *testing.T) {
	server := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: "test:atomic:feed", RelationL1TTL: time.Second})

	const workers = 64
	var wg sync.WaitGroup
	errorsCh := make(chan error, workers)
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			_, err := store.BumpRelation(context.Background(), 7)
			errorsCh <- err
		}()
	}
	wg.Wait()
	close(errorsCh)
	for err := range errorsCh {
		require.NoError(t, err)
	}

	value, err := store.Relation(context.Background(), 7)
	require.NoError(t, err)
	require.Equal(t, uint64(workers), value)
	redisValue, err := client.Get(context.Background(), relationKey("test:atomic:feed", 7)).Uint64()
	require.NoError(t, err)
	require.Equal(t, uint64(workers), redisValue)
	require.Zero(t, server.TTL(relationKey("test:atomic:feed", 7)), "epoch keys must remain persistent")
}

func TestConcurrentColdRelationMissesUseOneRedisGet(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	prefix := "test:singleflight:feed"
	client := &blockingGetRedis{
		UniversalClient: base,
		blockedKey:      relationKey(prefix, 7),
		started:         make(chan struct{}),
		release:         make(chan struct{}),
	}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: prefix})

	const workers = 32
	results := make(chan error, workers)
	for range workers {
		go func() {
			_, err := store.Relation(context.Background(), 7)
			results <- err
		}()
	}
	<-client.started
	time.Sleep(20 * time.Millisecond)
	close(client.release)
	for range workers {
		require.NoError(t, <-results)
	}
	require.Equal(t, int64(1), client.getCalls.Load())
}

func TestRelationForAnotherShardDoesNotWaitBehindBump(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	client := &blockingIncrRedis{
		UniversalClient: base,
		started:         make(chan struct{}),
		release:         make(chan struct{}),
	}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: "test:shards:feed"})
	bumpDone := make(chan error, 1)
	go func() {
		_, err := store.BumpRelation(context.Background(), 1)
		bumpDone <- err
	}()
	<-client.started

	relationDone := make(chan error, 1)
	go func() {
		_, err := store.Relation(context.Background(), 2)
		relationDone <- err
	}()
	var relationErr error
	completedBeforeRelease := false
	select {
	case relationErr = <-relationDone:
		completedBeforeRelease = true
	case <-time.After(500 * time.Millisecond):
	}
	close(client.release)
	require.NoError(t, <-bumpDone)
	if !completedBeforeRelease {
		relationErr = <-relationDone
	}
	require.NoError(t, relationErr)
	require.True(t, completedBeforeRelease, "an unrelated user must not wait behind another user's Redis INCR")
}

func TestBumpForAnotherShardDoesNotWaitBehindColdGet(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	prefix := "test:shard-get:feed"
	client := &blockingGetRedis{
		UniversalClient: base,
		blockedKey:      relationKey(prefix, 1),
		started:         make(chan struct{}),
		release:         make(chan struct{}),
	}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: prefix})
	readDone := make(chan error, 1)
	go func() {
		_, err := store.Relation(context.Background(), 1)
		readDone <- err
	}()
	<-client.started

	bumpDone := make(chan error, 1)
	go func() {
		_, err := store.BumpRelation(context.Background(), 2)
		bumpDone <- err
	}()
	var bumpErr error
	bumpedBeforeRelease := false
	select {
	case bumpErr = <-bumpDone:
		bumpedBeforeRelease = true
	case <-time.After(500 * time.Millisecond):
	}
	close(client.release)
	require.NoError(t, <-readDone)
	if !bumpedBeforeRelease {
		bumpErr = <-bumpDone
	}
	require.NoError(t, bumpErr)
	require.True(t, bumpedBeforeRelease, "an unrelated user bump must not wait behind another user's Redis GET")
}

func TestStaleRelationFillCannotOverwriteConcurrentBump(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	prefix := "test:fill-race:feed"
	key := relationKey(prefix, 9)
	require.NoError(t, base.Set(context.Background(), key, 5, 0).Err())
	client := &blockingGetRedis{
		UniversalClient:    base,
		blockedKey:         key,
		captureBeforeBlock: true,
		started:            make(chan struct{}),
		release:            make(chan struct{}),
	}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: prefix})
	readDone := make(chan error, 1)
	go func() {
		_, err := store.Relation(context.Background(), 9)
		readDone <- err
	}()
	<-client.started
	bumpDone := make(chan error, 1)
	go func() {
		_, err := store.BumpRelation(context.Background(), 9)
		bumpDone <- err
	}()
	select {
	case err := <-bumpDone:
		require.NoError(t, err)
		t.Fatal("same-key bump completed while the stale fill still held its shard lock")
	case <-time.After(100 * time.Millisecond):
	}
	close(client.release)
	require.NoError(t, <-readDone)
	require.NoError(t, <-bumpDone)

	value, err := store.Relation(context.Background(), 9)
	require.NoError(t, err)
	require.Equal(t, uint64(6), value)
}

func TestBumpRelationRejectsInvalidUserID(t *testing.T) {
	server := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	store := mustTestStore(t, client, newTestL1(t), Config{})

	_, err := store.Relation(context.Background(), 0)
	require.ErrorContains(t, err, "userID")
	_, err = store.BumpRelation(context.Background(), -1)
	require.ErrorContains(t, err, "userID")
}

func TestSafetyFailureSetsPendingAndFlushCoalescesFailures(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	client := &switchableRedis{UniversalClient: base}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: "test:safety:feed"})
	client.failIncr.Store(true)

	for range 3 {
		_, err := store.BumpSafety(context.Background())
		require.Error(t, err)
		require.True(t, store.SafetyPending())
	}

	client.failIncr.Store(false)
	require.NoError(t, store.FlushPendingSafety(context.Background()))
	require.False(t, store.SafetyPending())
	value, err := base.Get(context.Background(), safetyKey("test:safety:feed")).Uint64()
	require.NoError(t, err)
	require.Equal(t, uint64(1), value, "one compensating bump invalidates every old page")
	require.Zero(t, server.TTL(safetyKey("test:safety:feed")), "epoch keys must remain persistent")
	require.NoError(t, store.FlushPendingSafety(context.Background()))
	value, err = base.Get(context.Background(), safetyKey("test:safety:feed")).Uint64()
	require.NoError(t, err)
	require.Equal(t, uint64(1), value, "flushing without pending work must be a no-op")
}

func TestMarkSafetyPendingCannotBeLostBehindConcurrentSuccessfulBump(t *testing.T) {
	server := miniredis.RunT(t)
	base := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	client := &blockingIncrRedis{
		UniversalClient: base,
		started:         make(chan struct{}),
		release:         make(chan struct{}),
	}
	store := mustTestStore(t, client, newTestL1(t), Config{KeyPrefix: "test:safety-race:feed"})
	bumpDone := make(chan error, 1)
	go func() {
		_, err := store.BumpSafety(context.Background())
		bumpDone <- err
	}()
	<-client.started

	store.MarkSafetyPending()
	close(client.release)
	require.NoError(t, <-bumpDone)
	require.True(t, store.SafetyPending())
}

func TestSafetyL1ExpiresAndBumpIsImmediatelyVisible(t *testing.T) {
	server := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	l1 := newTestL1(t)
	prefix := "test:safety-l1:feed"
	store := mustTestStore(t, client, l1, Config{KeyPrefix: prefix, SafetyL1TTL: 20 * time.Millisecond})

	value, err := store.Safety(context.Background())
	require.NoError(t, err)
	require.Zero(t, value)
	l1.Wait()
	require.NoError(t, client.Set(context.Background(), safetyKey(prefix), 9, 0).Err())
	value, err = store.Safety(context.Background())
	require.NoError(t, err)
	require.Zero(t, value)
	time.Sleep(30 * time.Millisecond)
	value, err = store.Safety(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(9), value)

	bumped, err := store.BumpSafety(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(10), bumped)
	value, err = store.Safety(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(10), value)
}

func TestRedisFailuresNeverFabricateEpochs(t *testing.T) {
	client := &switchableRedis{}
	client.failGet.Store(true)
	client.failIncr.Store(true)
	store := mustTestStore(t, client, newTestL1(t), Config{})

	relation, err := store.Relation(context.Background(), 1)
	require.Error(t, err)
	require.Zero(t, relation)
	bumped, err := store.BumpRelation(context.Background(), 1)
	require.Error(t, err)
	require.Zero(t, bumped)
	safety, err := store.Safety(context.Background())
	require.Error(t, err)
	require.Zero(t, safety)
}

func TestEpochKeysHonorCleanupPrefix(t *testing.T) {
	require.Equal(t, "test:run-42:feed:relation:epoch:99", relationKey("test:run-42:feed", 99))
	require.Equal(t, "test:run-42:feed:content:safety:epoch", safetyKey("test:run-42:feed"))
}

func newTestL1(t *testing.T) *cachex.L1 {
	t.Helper()
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1_000, MaxCost: 1 << 20})
	require.NoError(t, err)
	return l1
}

func mustTestStore(t *testing.T, client RedisClient, l1 *cachex.L1, cfg Config) Store {
	t.Helper()
	store, err := NewStore(client, l1, cfg)
	require.NoError(t, err)
	return store
}

type switchableRedis struct {
	goredis.UniversalClient
	failGet  atomic.Bool
	failIncr atomic.Bool
}

type blockingIncrRedis struct {
	goredis.UniversalClient
	started   chan struct{}
	startOnce sync.Once
	release   chan struct{}
}

type blockingGetRedis struct {
	goredis.UniversalClient
	blockedKey         string
	captureBeforeBlock bool
	started            chan struct{}
	startOnce          sync.Once
	release            chan struct{}
	getCalls           atomic.Int64
}

func (r *blockingGetRedis) Get(ctx context.Context, key string) *goredis.StringCmd {
	r.getCalls.Add(1)
	if key == r.blockedKey {
		var captured *goredis.StringCmd
		if r.captureBeforeBlock {
			captured = r.UniversalClient.Get(ctx, key)
		}
		r.startOnce.Do(func() { close(r.started) })
		select {
		case <-r.release:
		case <-ctx.Done():
			return goredis.NewStringResult("", ctx.Err())
		}
		if captured != nil {
			return captured
		}
	}
	return r.UniversalClient.Get(ctx, key)
}

func (r *blockingIncrRedis) Incr(ctx context.Context, key string) *goredis.IntCmd {
	r.startOnce.Do(func() { close(r.started) })
	select {
	case <-r.release:
	case <-ctx.Done():
		return goredis.NewIntResult(0, ctx.Err())
	}
	return r.UniversalClient.Incr(ctx, key)
}

func (r *switchableRedis) Get(ctx context.Context, key string) *goredis.StringCmd {
	if r.failGet.Load() || r.UniversalClient == nil {
		return goredis.NewStringResult("", errors.New("redis get unavailable"))
	}
	return r.UniversalClient.Get(ctx, key)
}

func (r *switchableRedis) Incr(ctx context.Context, key string) *goredis.IntCmd {
	if r.failIncr.Load() || r.UniversalClient == nil {
		return goredis.NewIntResult(0, errors.New("redis incr unavailable"))
	}
	return r.UniversalClient.Incr(ctx, key)
}
