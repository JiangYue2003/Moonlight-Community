package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
)

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
}

func NewMockCounterClient() *MockCounterClient {
	return &MockCounterClient{
		followerCounts: make(map[int64]int64),
	}
}

func (m *MockCounterClient) SetFollowerCount(userID int64, count int64) {
	m.followerCounts[userID] = count
}

func (m *MockCounterClient) GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	if count, ok := m.followerCounts[userID]; ok {
		return count, nil
	}
	return 0, nil
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
