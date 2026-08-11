package feed

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

type concurrencyTrackingRedis struct {
	*MockRedisClient
	active    int32
	maxActive int32
	started   chan<- struct{}
	release   <-chan struct{}
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
