package feed

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
)

type concurrencyTrackingRedis struct {
	*MockRedisClient
	active    int32
	maxActive int32
	started   chan<- struct{}
	release   <-chan struct{}
}

type batchTrackingRedis struct {
	*MockRedisClient
	batchCalls    int
	requests      []ZRevRangeRequest
	batches       [][]ZRevRangeRequest
	results       []ZRevRangeResult
	resultBatches [][]ZRevRangeResult
	singleResults map[string][]ZScore
	singleCalls   []string
}

type scriptedSingleRedis struct {
	*MockRedisClient
	results map[string][]ZScore
	calls   []string
}

type cancelingBatchRedis struct {
	*MockRedisClient
	batchCalls int
	cancel     context.CancelFunc
}

type lockedRouteCache struct {
	mu     sync.RWMutex
	values map[string]any
}

type scriptedCounterResult struct {
	counts map[int64]int64
	err    error
}

type scriptedCounterClient struct {
	mu      sync.Mutex
	results []scriptedCounterResult
	calls   int
}

type blockingCounterClient struct {
	calls   int32
	started chan<- struct{}
	release <-chan struct{}
	counts  map[int64]int64
}

type contextAwareBlockingCounterClient struct {
	calls    int32
	started  chan<- struct{}
	release  <-chan struct{}
	canceled chan<- struct{}
	counts   map[int64]int64
}

type logLevelTrackingLogger struct {
	logx.Logger
	infoCalls  int
	debugCalls int
}

func (l *logLevelTrackingLogger) Infof(string, ...interface{}) {
	l.infoCalls++
}

func (l *logLevelTrackingLogger) Debugf(string, ...interface{}) {
	l.debugCalls++
}

func (m *batchTrackingRedis) ZRevRangeWithScoresBatch(_ context.Context, requests []ZRevRangeRequest) []ZRevRangeResult {
	m.batchCalls++
	m.requests = append([]ZRevRangeRequest(nil), requests...)
	m.batches = append(m.batches, append([]ZRevRangeRequest(nil), requests...))
	if index := m.batchCalls - 1; index < len(m.resultBatches) {
		return append([]ZRevRangeResult(nil), m.resultBatches[index]...)
	}
	if m.results == nil {
		return make([]ZRevRangeResult, len(requests))
	}
	return append([]ZRevRangeResult(nil), m.results...)
}

func (m *batchTrackingRedis) ZRevRangeWithScores(
	_ context.Context,
	key string,
	_, _ int64,
) ([]ZScore, error) {
	m.singleCalls = append(m.singleCalls, key)
	return append([]ZScore(nil), m.singleResults[key]...), nil
}

func (m *scriptedSingleRedis) ZRevRangeWithScores(
	_ context.Context,
	key string,
	_, _ int64,
) ([]ZScore, error) {
	m.calls = append(m.calls, key)
	return append([]ZScore(nil), m.results[key]...), nil
}

func (m *cancelingBatchRedis) ZRevRangeWithScoresBatch(
	_ context.Context,
	requests []ZRevRangeRequest,
) []ZRevRangeResult {
	m.batchCalls++
	if m.batchCalls == 1 {
		m.cancel()
	}
	results := make([]ZRevRangeResult, len(requests))
	for i := range results {
		results[i].Err = context.Canceled
	}
	return results
}

func newCombinedPipelineTestReader(redis RedisClient, observer FeedObserver) *FeedReader {
	return NewFeedReaderWithOptions(
		redis,
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:                  StrategyHybrid,
			CombinedPipelineEnabled:   true,
			CombinedPipelineBatchSize: maxBigVOutboxPipelineSize,
			Observer:                  observer,
		},
	)
}

func newLockedRouteCache() *lockedRouteCache {
	return &lockedRouteCache{values: make(map[string]any)}
}

func (c *lockedRouteCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.values[key]
	return value, ok
}

func (c *lockedRouteCache) SetWithTTL(key string, value any, _ int64, _ time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
	return true
}

func (c *scriptedCounterClient) BatchGetFollowerCounts(_ context.Context, _ []int64) (map[int64]int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.calls++
	index := c.calls - 1
	if index >= len(c.results) {
		index = len(c.results) - 1
	}
	result := c.results[index]
	return result.counts, result.err
}

func (c *blockingCounterClient) BatchGetFollowerCounts(_ context.Context, _ []int64) (map[int64]int64, error) {
	atomic.AddInt32(&c.calls, 1)
	c.started <- struct{}{}
	<-c.release
	return c.counts, nil
}

func (c *contextAwareBlockingCounterClient) BatchGetFollowerCounts(ctx context.Context, _ []int64) (map[int64]int64, error) {
	atomic.AddInt32(&c.calls, 1)
	c.started <- struct{}{}
	select {
	case <-c.release:
		return c.counts, nil
	case <-ctx.Done():
		c.canceled <- struct{}{}
		return nil, ctx.Err()
	}
}

func (m *concurrencyTrackingRedis) ZRevRangeWithScores(context.Context, string, int64, int64) ([]ZScore, error) {
	active := atomic.AddInt32(&m.active, 1)
	for {
		maxActive := atomic.LoadInt32(&m.maxActive)
		if active <= maxActive || atomic.CompareAndSwapInt32(&m.maxActive, maxActive, active) {
			break
		}
	}
	m.started <- struct{}{}
	<-m.release
	atomic.AddInt32(&m.active, -1)
	return nil, nil
}

// ===== 测试归并、去重等纯函数 =====

func TestMergePosts(t *testing.T) {
	// 测试归并排序（按时间倒序）
	list1 := []Post{
		{ID: 1, CreateTime: 1000},
		{ID: 2, CreateTime: 800},
	}

	list2 := []Post{
		{ID: 3, CreateTime: 900},
		{ID: 4, CreateTime: 700},
	}

	result := MergePosts(list1, list2)

	// 验证：按时间倒序排列
	assert.Equal(t, 4, len(result))
	assert.Equal(t, int64(1000), result[0].CreateTime) // 最新
	assert.Equal(t, int64(900), result[1].CreateTime)
	assert.Equal(t, int64(800), result[2].CreateTime)
	assert.Equal(t, int64(700), result[3].CreateTime) // 最旧
}

func TestMergePostsBreaksTimestampTiesByPostID(t *testing.T) {
	posts := MergePosts(
		[]Post{{ID: 10, CreateTime: 1000}},
		[]Post{{ID: 12, CreateTime: 1000}, {ID: 11, CreateTime: 1000}},
	)

	require.Equal(t, []int64{12, 11, 10}, []int64{posts[0].ID, posts[1].ID, posts[2].ID})
}

func TestDeduplicate(t *testing.T) {
	// 测试去重（保留第一次出现的）
	posts := []Post{
		{ID: 1, CreateTime: 1000},
		{ID: 2, CreateTime: 900},
		{ID: 1, CreateTime: 800}, // 重复
		{ID: 3, CreateTime: 700},
		{ID: 2, CreateTime: 600}, // 重复
	}

	result := Deduplicate(posts)

	// 验证：只保留3个唯一的帖子
	assert.Equal(t, 3, len(result))
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, int64(2), result[1].ID)
	assert.Equal(t, int64(3), result[2].ID)

	// 验证：保留的是第一次出现的（时间戳更早）
	assert.Equal(t, int64(1000), result[0].CreateTime)
	assert.Equal(t, int64(900), result[1].CreateTime)
}

func TestBalanceFeed(t *testing.T) {
	// 当前是简化版本，直接返回原数组
	posts := []Post{
		{ID: 1, CreateTime: 1000},
		{ID: 2, CreateTime: 900},
	}

	result := BalanceFeed(posts, 0.5)

	assert.Equal(t, len(posts), len(result))
}

// ===== Mock RelationClient =====

type MockRelationClient struct {
	followings map[int64][]int64
	followers  map[int64][]int64
}

func NewMockRelationClient() *MockRelationClient {
	return &MockRelationClient{
		followings: make(map[int64][]int64),
		followers:  make(map[int64][]int64),
	}
}

func (m *MockRelationClient) SetFollowings(userID int64, followings []int64) {
	m.followings[userID] = followings
}

func (m *MockRelationClient) SetFollowers(userID int64, followers []int64) {
	m.followers[userID] = followers
}

func (m *MockRelationClient) GetFollowings(ctx context.Context, userID int64) ([]int64, error) {
	if followings, ok := m.followings[userID]; ok {
		return followings, nil
	}
	return []int64{}, nil
}

func (m *MockRelationClient) GetFollowers(ctx context.Context, userID int64) ([]int64, error) {
	if followers, ok := m.followers[userID]; ok {
		return followers, nil
	}
	return []int64{}, nil
}

// ===== 扩展 MockRedisClient 支持 ZRevRangeWithScores =====

func (m *MockRedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZScore, error) {
	// 简单实现：返回空
	// 实际测试中可以根据需要填充数据
	return []ZScore{}, nil
}

// ===== Mock CounterClient =====

type MockCounterClient struct {
	followerCounts map[int64]int64
	batchCalls     int
}

func NewMockCounterClient() *MockCounterClient {
	return &MockCounterClient{
		followerCounts: make(map[int64]int64),
	}
}

func (m *MockCounterClient) SetFollowerCount(userID int64, count int64) {
	m.followerCounts[userID] = count
}

func (m *MockCounterClient) BatchGetFollowerCounts(_ context.Context, userIDs []int64) (map[int64]int64, error) {
	m.batchCalls++
	counts := make(map[int64]int64, len(userIDs))
	for _, userID := range userIDs {
		counts[userID] = m.followerCounts[userID]
	}
	return counts, nil
}

// ===== 测试 FeedReader =====

func TestFeedReader_NoFollowings(t *testing.T) {
	// 场景：用户没有关注任何人

	redis := NewMockRedisClient()
	relation := NewMockRelationClient()
	counter := NewMockCounterClient()
	reader := NewFeedReader(redis, relation, counter, logx.WithContext(context.Background()))

	ctx := context.Background()
	userID := int64(123)

	// 没有设置关注列表
	result, hasMore, err := reader.GetFeed(ctx, userID, 1, 20)

	// 验证
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
	assert.False(t, hasMore)
}

func TestFeedReaderPrepareDoesNotLogSuccessfulClassificationPerRequest(t *testing.T) {
	ctx := context.Background()
	logger := &logLevelTrackingLogger{Logger: logx.WithContext(ctx)}
	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	reader := NewFeedReader(NewMockRedisClient(), relation, NewMockCounterClient(), logger)

	_, err := reader.Prepare(ctx, 123)

	require.NoError(t, err)
	require.Zero(t, logger.infoCalls)
	require.Zero(t, logger.debugCalls)
}

func TestFeedReader_NoFollowingsStillReadsOwnInbox(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	ctx := context.Background()
	const userID int64 = 123
	require.NoError(t, rdb.ZAdd(ctx, fmt.Sprintf(FEED_INBOX_KEY, userID), goredis.Z{
		Score:  1000,
		Member: 99,
	}).Err())

	reader := NewFeedReader(
		NewRedisAdapter(rdb),
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)

	ids, hasMore, err := reader.GetFeed(ctx, userID, 1, 20)
	require.NoError(t, err)
	require.Equal(t, []int64{99}, ids)
	require.False(t, hasMore)
}

func TestFeedReader_WithFollowings(t *testing.T) {
	// 场景：用户关注了一些人（都是普通用户，不是大V）

	redis := NewMockRedisClient()
	relation := NewMockRelationClient()
	counter := NewMockCounterClient()
	reader := NewFeedReader(redis, relation, counter, logx.WithContext(context.Background()))

	ctx := context.Background()
	userID := int64(123)

	// 设置关注列表
	relation.SetFollowings(userID, []int64{201, 202, 203})

	// 设置粉丝数（都是普通用户）
	counter.SetFollowerCount(201, 100)
	counter.SetFollowerCount(202, 200)
	counter.SetFollowerCount(203, 300)

	result, hasMore, err := reader.GetFeed(ctx, userID, 1, 20)

	// 验证：不报错（即使收件箱和大V发件箱都是空的）
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, hasMore)
}

func TestFeedReaderRouteCacheReusesClassificationForUnchangedFollowings(t *testing.T) {
	ctx := context.Background()
	routeCache, err := cachex.NewL1(cachex.L1Config{
		NumCounters: 100,
		MaxCost:     1 << 20,
	})
	require.NoError(t, err)

	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201, 202})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, BIGV_THRESHOLD+1)
	counter.SetFollowerCount(202, 10)
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		routeCache,
		time.Minute,
	)

	first, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)
	routeCache.Wait()
	second, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)

	require.Equal(t, []int64{201}, first.bigVs)
	require.Equal(t, []int64{201}, second.bigVs)
	require.Equal(t, 1, counter.batchCalls)
}

func TestFeedReaderRouteCacheReclassifiesChangedFollowings(t *testing.T) {
	ctx := context.Background()
	routeCache, err := cachex.NewL1(cachex.L1Config{
		NumCounters: 100,
		MaxCost:     1 << 20,
	})
	require.NoError(t, err)

	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, BIGV_THRESHOLD+1)
	counter.SetFollowerCount(202, BIGV_THRESHOLD+1)
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		routeCache,
		time.Minute,
	)

	_, err = reader.Prepare(ctx, 123)
	require.NoError(t, err)
	routeCache.Wait()
	relation.SetFollowings(123, []int64{202})
	second, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)

	require.Equal(t, []int64{202}, second.bigVs)
	require.False(t, second.AllowsCreator(201))
	require.True(t, second.AllowsCreator(202))
	require.Equal(t, 2, counter.batchCalls)
}

func TestFeedReaderRouteCacheReclassifiesThresholdTransitionAfterTTL(t *testing.T) {
	ctx := context.Background()
	routeCache, err := cachex.NewL1(cachex.L1Config{
		NumCounters: 100,
		MaxCost:     1 << 20,
	})
	require.NoError(t, err)

	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, BIGV_THRESHOLD)
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		routeCache,
		20*time.Millisecond,
	)

	first, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)
	routeCache.Wait()
	counter.SetFollowerCount(201, BIGV_THRESHOLD+1)
	withinTTL, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)
	time.Sleep(30 * time.Millisecond)
	afterTTL, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)

	require.Empty(t, first.bigVs)
	require.Empty(t, withinTTL.bigVs, "阈值切换在 TTL 内按最终一致性处理")
	require.Equal(t, []int64{201}, afterTTL.bigVs)
	require.Equal(t, 2, counter.batchCalls)
}

func TestFeedReaderRouteCacheDoesNotCacheCounterFailureFallback(t *testing.T) {
	ctx := context.Background()
	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	counter := &scriptedCounterClient{results: []scriptedCounterResult{
		{err: errors.New("counter unavailable")},
		{counts: map[int64]int64{201: BIGV_THRESHOLD + 1}},
	}}
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		newLockedRouteCache(),
		time.Minute,
	)

	first, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)
	second, err := reader.Prepare(ctx, 123)
	require.NoError(t, err)

	require.Empty(t, first.bigVs, "Counter 失败时应保持原有的全普通用户降级")
	require.Equal(t, []int64{201}, second.bigVs, "Counter 恢复后必须重新分类")
	require.Equal(t, 2, counter.calls)
}

func TestFeedReaderRouteCacheCoalescesConcurrentColdMisses(t *testing.T) {
	const requestCount = 32
	ctx := context.Background()
	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	started := make(chan struct{}, requestCount)
	release := make(chan struct{})
	counter := &blockingCounterClient{
		started: started,
		release: release,
		counts:  map[int64]int64{201: BIGV_THRESHOLD + 1},
	}
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		newLockedRouteCache(),
		time.Minute,
	)

	var wg sync.WaitGroup
	type prepareResult struct {
		bigVs []int64
		err   error
	}
	results := make(chan prepareResult, requestCount)
	wg.Add(requestCount)
	for i := 0; i < requestCount; i++ {
		go func() {
			defer wg.Done()
			snapshot, err := reader.Prepare(ctx, 123)
			if err != nil {
				results <- prepareResult{err: err}
				return
			}
			results <- prepareResult{bigVs: snapshot.bigVs}
		}()
	}

	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("Counter RPC did not start")
	}
	time.Sleep(50 * time.Millisecond)
	require.Equal(t, int32(1), atomic.LoadInt32(&counter.calls), "同一用户的冷缓存请求应合并为一次 Counter RPC")
	close(release)
	wg.Wait()
	close(results)
	for result := range results {
		require.NoError(t, result.err)
		require.Equal(t, []int64{201}, result.bigVs)
	}
}

func TestFeedReaderRouteCacheDoesNotBindSharedLookupToFirstCallerContext(t *testing.T) {
	ctx := context.Background()
	firstCtx, cancelFirst := context.WithCancel(ctx)
	relation := NewMockRelationClient()
	relation.SetFollowings(123, []int64{201})
	started := make(chan struct{}, 2)
	release := make(chan struct{})
	canceled := make(chan struct{}, 1)
	counter := &contextAwareBlockingCounterClient{
		started:  started,
		release:  release,
		canceled: canceled,
		counts:   map[int64]int64{201: BIGV_THRESHOLD + 1},
	}
	reader := NewFeedReaderWithStrategyAndRouteCache(
		NewMockRedisClient(),
		relation,
		counter,
		logx.WithContext(ctx),
		StrategyHybrid,
		newLockedRouteCache(),
		time.Minute,
	)

	type prepareResult struct {
		snapshot *FeedReadSnapshot
		err      error
	}
	firstDone := make(chan prepareResult, 1)
	secondDone := make(chan prepareResult, 1)
	go func() {
		snapshot, err := reader.Prepare(firstCtx, 123)
		firstDone <- prepareResult{snapshot: snapshot, err: err}
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("shared Counter lookup did not start")
	}
	go func() {
		snapshot, err := reader.Prepare(ctx, 123)
		secondDone <- prepareResult{snapshot: snapshot, err: err}
	}()

	cancelFirst()
	select {
	case result := <-firstDone:
		require.NoError(t, result.err)
		require.Empty(t, result.snapshot.bigVs)
	case <-time.After(200 * time.Millisecond):
		t.Fatal("canceled waiter did not exit promptly")
	}
	select {
	case <-canceled:
		t.Fatal("first caller cancellation propagated into shared Counter lookup")
	case <-time.After(50 * time.Millisecond):
	}

	close(release)
	select {
	case result := <-secondDone:
		require.NoError(t, result.err)
		require.Equal(t, []int64{201}, result.snapshot.bigVs)
	case <-time.After(time.Second):
		t.Fatal("valid waiter did not receive shared classification")
	}
	require.Equal(t, int32(1), atomic.LoadInt32(&counter.calls))
}

func TestFeedReader_Pagination(t *testing.T) {
	// 测试分页逻辑

	// 模拟 MergePosts 返回大量帖子
	posts := []Post{}
	for i := 0; i < 100; i++ {
		posts = append(posts, Post{
			ID:         int64(i + 1),
			CreateTime: int64(1000 - i),
		})
	}

	// 测试第一页
	page := 1
	size := 20
	start := (page - 1) * size
	end := start + size
	hasMore := len(posts) > end

	assert.Equal(t, 0, start)
	assert.Equal(t, 20, end)
	assert.True(t, hasMore)

	// 测试最后一页
	page = 5
	start = (page - 1) * size
	end = start + size
	if end > len(posts) {
		end = len(posts)
	}
	hasMore = len(posts) > end

	assert.Equal(t, 80, start)
	assert.Equal(t, 100, end)
	assert.False(t, hasMore)
}

func TestFeedReader_DeepPageFetchesEnoughCandidates(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	ctx := context.Background()
	const userID int64 = 456
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)
	for id := int64(1); id <= 61; id++ {
		require.NoError(t, rdb.ZAdd(ctx, key, goredis.Z{
			Score:  float64(1000 - id),
			Member: id,
		}).Err())
	}

	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{7})
	reader := NewFeedReader(
		NewRedisAdapter(rdb),
		relation,
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)

	ids, hasMore, err := reader.GetFeed(ctx, userID, 3, 20)
	require.NoError(t, err)
	require.Len(t, ids, 20)
	require.Equal(t, int64(41), ids[0])
	require.Equal(t, int64(60), ids[19])
	require.True(t, hasMore)
}

func TestFeedReader_DeepPageKeepsNewestPostsFromSkewedBigV(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	ctx := context.Background()
	const userID int64 = 789
	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{7, 8})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(7, BIGV_THRESHOLD+1)
	counter.SetFollowerCount(8, BIGV_THRESHOLD+1)

	for id := int64(1); id <= 61; id++ {
		require.NoError(t, rdb.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 7), goredis.Z{
			Score:  float64(10000 - id),
			Member: id,
		}).Err())
	}
	for offset := int64(1); offset <= 61; offset++ {
		require.NoError(t, rdb.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 8), goredis.Z{
			Score:  float64(1000 - offset),
			Member: 1000 + offset,
		}).Err())
	}

	reader := NewFeedReader(NewRedisAdapter(rdb), relation, counter, logx.WithContext(ctx))
	ids, hasMore, err := reader.GetFeed(ctx, userID, 3, 20)
	require.NoError(t, err)
	require.Len(t, ids, 20)
	require.Equal(t, int64(41), ids[0])
	require.Equal(t, int64(60), ids[19])
	require.True(t, hasMore)
}

func TestFeedReader_PullFromBigVsBoundsConcurrency(t *testing.T) {
	const (
		bigVCount      = 32
		maxConcurrency = 16
		candidateLimit = 20
	)

	started := make(chan struct{}, bigVCount)
	release := make(chan struct{})
	redis := &concurrencyTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		started:         started,
		release:         release,
	}
	reader := NewFeedReader(redis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))
	bigVs := make([]int64, bigVCount)
	for i := range bigVs {
		bigVs[i] = int64(i + 1)
	}

	done := make(chan struct{})
	go func() {
		reader.pullFromBigVs(context.Background(), bigVs, candidateLimit)
		close(done)
	}()

	for i := 0; i < maxConcurrency; i++ {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatalf("only %d pulls started before timeout", i)
		}
	}

	overflowed := false
	select {
	case <-started:
		overflowed = true
	case <-time.After(100 * time.Millisecond):
	}
	close(release)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("pullFromBigVs did not finish after releasing workers")
	}

	require.False(t, overflowed, "more than %d Redis pulls ran concurrently", maxConcurrency)
	require.LessOrEqual(t, atomic.LoadInt32(&redis.maxActive), int32(maxConcurrency))
}

func TestFeedReaderPullFromBigVsUsesBatchReader(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Scores: []ZScore{{Member: 101, Score: 1000}}},
			{Scores: []ZScore{{Member: 202, Score: 1100}}},
		},
	}
	reader := NewFeedReader(redis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))

	posts := reader.pullFromBigVs(context.Background(), []int64{1, 2}, 20)

	require.Equal(t, 1, redis.batchCalls)
	require.Equal(t, []ZRevRangeRequest{
		{Key: fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 1), Start: 0, Stop: 19},
		{Key: fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 2), Start: 0, Stop: 19},
	}, redis.requests)
	require.Equal(t, []Post{
		{ID: 202, CreatorID: 2, CreateTime: 1100},
		{ID: 101, CreatorID: 1, CreateTime: 1000},
	}, posts)
}

func TestFeedReadSnapshotCombinesInboxAndFirstBigVBatch(t *testing.T) {
	const userID int64 = 99
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Scores: []ZScore{{Member: 100, Score: 1200}}},
			{Scores: []ZScore{{Member: 201, Score: 1100}}},
			{Scores: []ZScore{{Member: 301, Score: 1000}}},
		},
	}
	reader := newCombinedPipelineTestReader(redis, nil)
	snapshot := &FeedReadSnapshot{
		reader: reader,
		userID: userID,
		bigVs:  []int64{2, 3},
	}

	ids, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{100, 201, 301}, ids)
	require.Equal(t, 1, redis.batchCalls)
	require.Equal(t, []ZRevRangeRequest{
		{Key: fmt.Sprintf(FEED_INBOX_KEY, userID), Start: 0, Stop: 20},
		{Key: fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 2), Start: 0, Stop: 20},
		{Key: fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 3), Start: 0, Stop: 20},
	}, redis.batches[0])
}

func TestFeedReadSnapshotCombinedPipelineRemainsDisabledByDefault(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results:         []ZRevRangeResult{{Scores: []ZScore{{Member: 201, Score: 1100}}}},
	}
	reader := NewFeedReader(redis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: []int64{2}}

	_, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, 1, redis.batchCalls)
	require.Equal(t, []ZRevRangeRequest{
		{Key: fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 2), Start: 0, Stop: 20},
	}, redis.batches[0])
}

func TestFeedReadSnapshotCombinedPipelineWorksWithRedisAdapter(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()
	const userID int64 = 99
	require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_INBOX_KEY, userID),
		goredis.Z{Score: 1200, Member: 100},
	).Err())
	require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 2),
		goredis.Z{Score: 1100, Member: 201},
	).Err())

	reader := newCombinedPipelineTestReader(NewRedisAdapter(client), nil)
	snapshot := &FeedReadSnapshot{reader: reader, userID: userID, bigVs: []int64{2}}

	ids, hasMore, err := snapshot.GetFeed(ctx, 1, 20)

	require.NoError(t, err)
	require.False(t, hasMore)
	require.Equal(t, []int64{100, 201}, ids)
}

func TestFeedReadSnapshotFirstPipelineCountsInboxAgainstBatchLimit(t *testing.T) {
	const userID int64 = 99
	redis := &batchTrackingRedis{MockRedisClient: NewMockRedisClient()}
	reader := newCombinedPipelineTestReader(redis, nil)
	bigVs := make([]int64, maxBigVOutboxPipelineSize)
	for i := range bigVs {
		bigVs[i] = int64(i + 1)
	}
	snapshot := &FeedReadSnapshot{reader: reader, userID: userID, bigVs: bigVs}

	_, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, 2, redis.batchCalls)
	require.Len(t, redis.batches[0], maxBigVOutboxPipelineSize)
	require.Equal(t, fmt.Sprintf(FEED_INBOX_KEY, userID), redis.batches[0][0].Key)
	require.Len(t, redis.batches[1], 1)
	for _, batch := range redis.batches {
		require.LessOrEqual(t, len(batch), maxBigVOutboxPipelineSize)
	}
}

func TestFeedReadSnapshotCombinedPipelineHonorsConfiguredBatchSize(t *testing.T) {
	redis := &batchTrackingRedis{MockRedisClient: NewMockRedisClient()}
	reader := NewFeedReaderWithOptions(
		redis,
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:                  StrategyHybrid,
			CombinedPipelineEnabled:   true,
			CombinedPipelineBatchSize: 4,
		},
	)
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: []int64{1, 2, 3, 4}}

	_, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, 2, redis.batchCalls)
	require.Len(t, redis.batches[0], 4, "Inbox plus three BigV commands")
	require.Len(t, redis.batches[1], 1)
}

func TestFeedReadSnapshotCombinedPipelineMultiBatchMatchesLegacyTopN(t *testing.T) {
	const (
		userID    int64 = 99
		bigVCount       = maxBigVOutboxPipelineSize + 1
	)
	bigVs := make([]int64, bigVCount)
	creatorResult := func(creatorID int64) ZRevRangeResult {
		return ZRevRangeResult{Scores: []ZScore{{
			Member: 1_000 + creatorID,
			Score:  float64(1_000 + creatorID),
		}}}
	}
	combinedFirst := make([]ZRevRangeResult, maxBigVOutboxPipelineSize)
	combinedFirst[0] = ZRevRangeResult{Scores: []ZScore{
		{Member: 9_001, Score: 1_200},
		{Member: 9_002, Score: 800},
	}}
	legacyFirst := make([]ZRevRangeResult, maxBigVOutboxPipelineSize)
	for i := range bigVs {
		creatorID := int64(i + 1)
		bigVs[i] = creatorID
		if i < maxBigVOutboxPipelineSize-1 {
			combinedFirst[i+1] = creatorResult(creatorID)
		}
		if i < maxBigVOutboxPipelineSize {
			legacyFirst[i] = creatorResult(creatorID)
		}
	}
	combinedSecond := []ZRevRangeResult{
		creatorResult(bigVs[maxBigVOutboxPipelineSize-1]),
		creatorResult(bigVs[maxBigVOutboxPipelineSize]),
	}
	legacySecond := []ZRevRangeResult{creatorResult(bigVs[maxBigVOutboxPipelineSize])}
	// The newest creators cross the batch boundary and share a timestamp, so
	// ID descending remains the deterministic tie-break after heap eviction.
	combinedSecond[0].Scores[0].Score = combinedSecond[1].Scores[0].Score
	legacyFirst[maxBigVOutboxPipelineSize-1].Scores[0].Score = combinedSecond[1].Scores[0].Score
	inboxKey := fmt.Sprintf(FEED_INBOX_KEY, userID)

	combinedRedis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		resultBatches:   [][]ZRevRangeResult{combinedFirst, combinedSecond},
	}
	observer := &recordingFeedObserver{}
	combinedReader := newCombinedPipelineTestReader(combinedRedis, observer)
	combinedSnapshot := &FeedReadSnapshot{reader: combinedReader, userID: userID, bigVs: bigVs}

	legacyRedis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		resultBatches:   [][]ZRevRangeResult{legacyFirst, legacySecond},
		singleResults: map[string][]ZScore{
			inboxKey: combinedFirst[0].Scores,
		},
	}
	legacyReader := NewFeedReader(legacyRedis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))
	legacySnapshot := &FeedReadSnapshot{reader: legacyReader, userID: userID, bigVs: bigVs}

	combinedIDs, combinedHasMore, err := combinedSnapshot.GetFeed(context.Background(), 1, 5)
	require.NoError(t, err)
	legacyIDs, legacyHasMore, err := legacySnapshot.GetFeed(context.Background(), 1, 5)
	require.NoError(t, err)

	require.Equal(t, legacyIDs, combinedIDs)
	require.Equal(t, []int64{9_001, 1_129, 1_128, 1_127, 1_126}, combinedIDs)
	require.True(t, combinedHasMore)
	require.Equal(t, legacyHasMore, combinedHasMore)
	require.Equal(t, 2, combinedRedis.batchCalls)
	require.Equal(t, 2, legacyRedis.batchCalls)
	require.Equal(t, []string{inboxKey}, legacyRedis.singleCalls)

	dependencyCounts := make(map[FeedOperation]int)
	for _, call := range observer.calls {
		if call.kind == "dependency" && call.dependency == DependencyRedis {
			dependencyCounts[call.operation]++
		}
	}
	require.Equal(t, 1, dependencyCounts[OperationInboxBigVPipeline])
	require.Equal(t, 1, dependencyCounts[OperationBigVPipeline])
	require.Zero(t, dependencyCounts[OperationInboxRead])
}

func TestFeedReadSnapshotCombinedPipelineFallsBackWithoutBatchReader(t *testing.T) {
	const userID int64 = 99
	inboxKey := fmt.Sprintf(FEED_INBOX_KEY, userID)
	bigVKey := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 2)
	redis := &scriptedSingleRedis{
		MockRedisClient: NewMockRedisClient(),
		results: map[string][]ZScore{
			inboxKey: {{Member: 100, Score: 1200}},
			bigVKey:  {{Member: 201, Score: 1100}},
		},
	}
	reader := newCombinedPipelineTestReader(redis, nil)
	snapshot := &FeedReadSnapshot{reader: reader, userID: userID, bigVs: []int64{2}}

	ids, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{100, 201}, ids)
	require.ElementsMatch(t, []string{inboxKey, bigVKey}, redis.calls)
}

func TestFeedReadSnapshotCombinedPipelineKeepsBigVWhenInboxFails(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Err: errors.New("inbox wrong type")},
			{Scores: []ZScore{{Member: 201, Score: 1100}}},
		},
	}
	reader := newCombinedPipelineTestReader(redis, nil)
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: []int64{2}}

	ids, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{201}, ids)
}

func TestFeedReadSnapshotCombinedPipelineKeepsInboxWhenBigVFails(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Scores: []ZScore{{Member: 100, Score: 1200}}},
			{Err: errors.New("bigv wrong type")},
		},
	}
	reader := newCombinedPipelineTestReader(redis, nil)
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: []int64{2}}

	ids, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{100}, ids)
}

func TestFeedReadSnapshotCombinedPipelineKeepsLaterBigVAfterMiddleFailure(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Scores: []ZScore{{Member: 100, Score: 1200}}},
			{Err: errors.New("first bigv wrong type")},
			{Scores: []ZScore{{Member: 301, Score: 1000}}},
		},
	}
	reader := newCombinedPipelineTestReader(redis, nil)
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: []int64{2, 3}}

	ids, _, err := snapshot.GetFeed(context.Background(), 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{100, 301}, ids)
}

func TestFeedReadSnapshotCombinedPipelineStopsAfterContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	redis := &cancelingBatchRedis{MockRedisClient: NewMockRedisClient(), cancel: cancel}
	reader := newCombinedPipelineTestReader(redis, nil)
	bigVs := make([]int64, maxBigVOutboxPipelineSize+1)
	for i := range bigVs {
		bigVs[i] = int64(i + 1)
	}
	snapshot := &FeedReadSnapshot{reader: reader, userID: 99, bigVs: bigVs}

	_, _, err := snapshot.GetFeed(ctx, 1, 20)

	require.NoError(t, err)
	require.Equal(t, 1, redis.batchCalls)
}

func TestFeedReaderPullFromBigVsBoundsPipelineBatchSize(t *testing.T) {
	redis := &batchTrackingRedis{MockRedisClient: NewMockRedisClient()}
	reader := NewFeedReader(redis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))
	bigVs := make([]int64, maxBigVOutboxPipelineSize+1)
	for i := range bigVs {
		bigVs[i] = int64(i + 1)
	}

	reader.pullFromBigVs(context.Background(), bigVs, 20)

	require.Equal(t, 2, redis.batchCalls)
	for _, batch := range redis.batches {
		require.LessOrEqual(t, len(batch), maxBigVOutboxPipelineSize)
	}
}

func TestFeedReaderPullFromBigVsBatchKeepsSuccessfulOutboxOnPartialFailure(t *testing.T) {
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Err: errors.New("wrong type")},
			{Scores: []ZScore{{Member: 202, Score: 1100}}},
		},
	}
	reader := NewFeedReader(redis, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(context.Background()))

	posts := reader.pullFromBigVs(context.Background(), []int64{1, 2}, 20)

	require.Equal(t, []Post{{ID: 202, CreatorID: 2, CreateTime: 1100}}, posts)
}

func TestFeedReader_PullFromBigVsKeepsOnlyGlobalTopN(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	ctx := context.Background()
	for creatorID := int64(1); creatorID <= 3; creatorID++ {
		for offset := int64(1); offset <= 20; offset++ {
			require.NoError(t, rdb.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID), goredis.Z{
				Score:  float64(10_000 - creatorID*100 - offset),
				Member: creatorID*1_000 + offset,
			}).Err())
		}
	}

	reader := NewFeedReader(NewRedisAdapter(rdb), NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(ctx))
	posts := reader.pullFromBigVs(ctx, []int64{1, 2, 3}, 15)

	require.Len(t, posts, 15)
	for i := 1; i < len(posts); i++ {
		require.True(t, posts[i-1].CreateTime > posts[i].CreateTime ||
			(posts[i-1].CreateTime == posts[i].CreateTime && posts[i-1].ID > posts[i].ID))
	}
	require.Equal(t, int64(1001), posts[0].ID)
	require.Equal(t, int64(1015), posts[14].ID)
}
