//go:build integration
// +build integration

package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

// 集成测试：测试推拉结合Feed架构的完整流程
// 运行方式：go test -tags=integration -v ./services/knowpost/rpc/internal/feed/

const (
	testRedisAddr   = "127.0.0.1:6379"
	testKafkaBroker = "127.0.0.1:9092"
)

// TestPushPullFeedIntegration 测试推拉结合Feed的完整流程
func TestPushPullFeedIntegration(t *testing.T) {
	// 跳过条件：如果环境变量设置了SKIP_INTEGRATION
	// if os.Getenv("SKIP_INTEGRATION") != "" {
	// 	t.Skip("Skipping integration test")
	// }

	ctx := context.Background()

	// 1. 初始化Redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr: testRedisAddr,
	})
	testKeys := []string{
		fmt.Sprintf(FEED_INBOX_KEY, 1001),
		fmt.Sprintf(FEED_INBOX_KEY, 1002),
		fmt.Sprintf(FEED_INBOX_KEY, 1003),
		fmt.Sprintf(FEED_INBOX_KEY, 1004),
		fmt.Sprintf(FEED_INBOX_KEY, 1005),
		fmt.Sprintf(FEED_INBOX_KEY, 3001),
		fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 1001),
		fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 1002),
		fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 1003),
		fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 1004),
	}
	t.Cleanup(func() {
		if err := deleteFeedIntegrationKeys(redisClient, ctx, testKeys...); err != nil {
			t.Logf("Warning: failed to delete scoped integration keys: %v", err)
		}
		_ = redisClient.Close()
	})

	// 测试Redis连接
	err := redisClient.Ping(ctx).Err()
	require.NoError(t, err, "Redis should be available")

	// Fail closed against residue from an interrupted previous run.
	require.NoError(t, deleteFeedIntegrationKeys(redisClient, ctx, testKeys...))

	// 2. 初始化适配器
	redisAdapter := NewRedisAdapter(redisClient)
	mockCounter := &mockCounterClient{}
	mockRelation := &mockRelationClient{}
	mockKafka := &mockKafkaProducer{}

	// 创建logger
	logger := logx.WithContext(ctx)

	// 3. 创建FeedWriter
	writer := NewFeedWriter(redisAdapter, mockKafka, logger)

	// 4. 创建FeedReader (用于其他非MixedRead测试)
	_ = NewFeedReader(redisAdapter, mockRelation, mockCounter, logger)

	// === 场景1：普通用户发帖（推模式）===
	t.Run("NormalUser_PushMode", func(t *testing.T) {
		creatorID := int64(1001)
		postID := int64(2001)
		followerCount := int64(500) // ≤ 1000，推模式
		timestamp := time.Now().Unix()

		// 发帖
		err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)
		require.NoError(t, err)

		// 验证：应该发送Kafka消息
		assert.True(t, mockKafka.sent, "Should send Kafka message for normal user")
		assert.Equal(t, postID, mockKafka.lastEvent.PostID)
		assert.Equal(t, creatorID, mockKafka.lastEvent.CreatorID)

		// 验证：作者自己的收件箱应该有这条帖子
		inbox, err := redisAdapter.ZRevRangeWithScores(ctx, fmt.Sprintf(FEED_INBOX_KEY, creatorID), 0, 100)
		require.NoError(t, err)
		inboxIDs := make([]int64, len(inbox))
		for i, z := range inbox {
			inboxIDs[i] = z.Member
		}
		assert.Contains(t, inboxIDs, postID, "Creator should see their own post")

		// 验证：不应该写入大V发件箱
		bigvOutbox, err := redisAdapter.ZRevRangeWithScores(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID), 0, 100)
		require.NoError(t, err)
		bigvIDs := make([]int64, len(bigvOutbox))
		for i, z := range bigvOutbox {
			bigvIDs[i] = z.Member
		}
		assert.NotContains(t, bigvIDs, postID, "Should not write to bigv outbox")

		t.Logf("✓ Normal user post: postID=%d, creatorID=%d, followerCount=%d, timestamp=%d",
			postID, creatorID, followerCount, timestamp)
	})

	// === 场景2：大V发帖（拉模式）===
	t.Run("BigV_PullMode", func(t *testing.T) {
		creatorID := int64(1002)
		postID := int64(2002)
		followerCount := int64(5000) // > 1000，拉模式
		timestamp := time.Now().Unix()

		// 重置mock状态
		mockKafka.sent = false

		// 发帖
		err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)
		require.NoError(t, err)

		// 验证：不应该发送Kafka消息
		assert.False(t, mockKafka.sent, "Should NOT send Kafka message for big V")

		// 验证：应该写入大V发件箱
		bigvOutbox, err := redisAdapter.ZRevRangeWithScores(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID), 0, 100)
		require.NoError(t, err)
		bigvIDs := make([]int64, len(bigvOutbox))
		for i, z := range bigvOutbox {
			bigvIDs[i] = z.Member
		}
		assert.Contains(t, bigvIDs, postID, "Should write to bigv outbox")

		// 验证：作者自己的收件箱应该有这条帖子
		inbox, err := redisAdapter.ZRevRangeWithScores(ctx, fmt.Sprintf(FEED_INBOX_KEY, creatorID), 0, 100)
		require.NoError(t, err)
		inboxIDs := make([]int64, len(inbox))
		for i, z := range inbox {
			inboxIDs[i] = z.Member
		}
		assert.Contains(t, inboxIDs, postID, "Creator should see their own post")

		t.Logf("✓ Big V post: postID=%d, creatorID=%d, followerCount=%d, timestamp=%d",
			postID, creatorID, followerCount, timestamp)
	})

	// === 场景3：阈值边界测试===
	t.Run("Threshold_Boundary", func(t *testing.T) {
		// 刚好1000粉丝，应该走推模式
		creatorID := int64(1003)
		postID := int64(2003)
		followerCount := int64(1000)

		mockKafka.sent = false
		err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)
		require.NoError(t, err)
		assert.True(t, mockKafka.sent, "1000 followers should use push mode")

		// 1001粉丝，应该走拉模式
		creatorID = int64(1004)
		postID = int64(2004)
		followerCount = int64(1001)

		mockKafka.sent = false
		err = writer.OnPostPublished(ctx, postID, creatorID, followerCount)
		require.NoError(t, err)
		assert.False(t, mockKafka.sent, "1001 followers should use pull mode")

		t.Logf("✓ Threshold boundary test passed")
	})

	// === 场景4：推拉混合读取===
	t.Run("PushPull_MixedRead", func(t *testing.T) {
		// 为这个子测试创建独立的mock实例
		subMockCounter := &mockCounterClient{}
		subMockRelation := &mockRelationClient{}
		subReader := NewFeedReader(redisAdapter, subMockRelation, subMockCounter, logger)

		userID := int64(3001)
		normalCreator := int64(1001) // 普通用户
		bigVCreator := int64(1002)   // 大V

		// 模拟关注列表：1个普通用户 + 1个大V
		subMockRelation.followings = []int64{normalCreator, bigVCreator}

		// 模拟粉丝数：用于分类大V
		subMockCounter.followerCounts = map[int64]int64{
			normalCreator: 500,
			bigVCreator:   5000,
		}

		// 模拟收件箱：普通用户的帖子已经在里面
		inboxKey := fmt.Sprintf(FEED_INBOX_KEY, userID)
		timestamp := float64(time.Now().Unix())
		err := redisAdapter.ZAdd(ctx, inboxKey, timestamp, int64(2001))
		require.NoError(t, err)

		// 验证写入
		count, _ := redisClient.ZCard(ctx, inboxKey).Result()
		t.Logf("Inbox %s has %d members after ZAdd", inboxKey, count)

		// 模拟大V发件箱：大V的帖子
		bigvOutboxKey := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, bigVCreator)
		err = redisAdapter.ZAdd(ctx, bigvOutboxKey, timestamp, int64(2002))
		require.NoError(t, err)

		// 验证写入
		count, _ = redisClient.ZCard(ctx, bigvOutboxKey).Result()
		t.Logf("BigV outbox %s has %d members after ZAdd", bigvOutboxKey, count)

		// 读取Feed
		posts, hasMore, err := subReader.GetFeed(ctx, userID, 1, 20)
		require.NoError(t, err)

		// 验证：应该包含两条帖子
		assert.Len(t, posts, 2, "Should have 2 posts from push and pull")
		assert.False(t, hasMore, "Should not have more posts")

		// 验证：应该包含普通用户和大V的帖子
		postIDs := make(map[int64]bool)
		for _, p := range posts {
			postIDs[p] = true
		}
		assert.True(t, postIDs[2001], "Should contain normal user's post")
		assert.True(t, postIDs[2002], "Should contain big V's post")

		t.Logf("✓ Mixed push-pull read: got %d posts", len(posts))
	})

	// === 场景5：容量限制测试===
	t.Run("Capacity_Limit", func(t *testing.T) {
		creatorID := int64(1005)

		// 写入超过容量限制的帖子
		for i := 0; i < INBOX_MAX_SIZE+100; i++ {
			postID := int64(3000 + i)
			timestamp := time.Now().Unix() + int64(i) // 递增时间戳

			inboxKey := fmt.Sprintf(FEED_INBOX_KEY, creatorID)
			err := redisAdapter.ZAdd(ctx, inboxKey, float64(timestamp), postID)
			require.NoError(t, err)

			// 保留最新的N条
			err = redisAdapter.ZRemRangeByRank(ctx, inboxKey, 0, -INBOX_MAX_SIZE-1)
			require.NoError(t, err)
		}

		// 验证：收件箱只保留了最大容量
		inboxKey := fmt.Sprintf(FEED_INBOX_KEY, creatorID)
		count, err := redisClient.ZCard(ctx, inboxKey).Result()
		require.NoError(t, err)
		assert.LessOrEqual(t, int(count), INBOX_MAX_SIZE,
			"Inbox should not exceed max size")

		t.Logf("✓ Capacity limit: inbox size=%d, max=%d", count, INBOX_MAX_SIZE)
	})
}

// TestFanoutWorkerIntegration 测试扇出Worker的完整流程
func TestFanoutWorkerIntegration(t *testing.T) {
	ctx := context.Background()

	// 初始化Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: testRedisAddr,
	})
	testKeys := []string{
		fmt.Sprintf(FEED_INBOX_KEY, 2001),
		fmt.Sprintf(FEED_INBOX_KEY, 2002),
		fmt.Sprintf(FEED_INBOX_KEY, 2003),
		fmt.Sprintf(FEED_INBOX_KEY, 2004),
		fmt.Sprintf(FEED_INBOX_KEY, 2005),
		fmt.Sprintf(FANOUT_PROCESSING_KEY, 3001),
	}
	t.Cleanup(func() {
		if err := deleteFeedIntegrationKeys(redisClient, ctx, testKeys...); err != nil {
			t.Logf("Warning: failed to delete scoped integration keys: %v", err)
		}
		_ = redisClient.Close()
	})

	err := redisClient.Ping(ctx).Err()
	require.NoError(t, err, "Redis should be available")

	require.NoError(t, deleteFeedIntegrationKeys(redisClient, ctx, testKeys...))

	redisAdapter := NewRedisAdapter(redisClient)
	mockRelation := &mockRelationClient{}

	// 模拟粉丝列表
	creatorID := int64(1001)
	followers := []int64{2001, 2002, 2003, 2004, 2005}
	mockRelation.followers = followers

	// 模拟扇出事件
	event := FeedEvent{
		PostID:     3001,
		CreatorID:  creatorID,
		CreateTime: time.Now().Unix(),
	}

	// 手动执行扇出逻辑（不依赖Kafka）
	processKey := fmt.Sprintf(FANOUT_PROCESSING_KEY, event.PostID)

	// 幂等性检查
	exists, err := redisAdapter.Exists(ctx, processKey)
	require.NoError(t, err)
	assert.False(t, exists, "Should not be processed yet")

	// 获取粉丝列表
	followerList, err := mockRelation.GetFollowers(ctx, event.CreatorID)
	require.NoError(t, err)
	assert.Equal(t, followers, followerList)

	// 推送到粉丝收件箱
	successCount := 0
	for _, followerID := range followerList {
		inboxKey := fmt.Sprintf(FEED_INBOX_KEY, followerID)
		err := redisAdapter.ZAdd(ctx, inboxKey, float64(event.CreateTime), event.PostID)
		if err == nil {
			successCount++
		}
	}

	// 标记已处理
	err = redisAdapter.SetEx(ctx, processKey, "1", FANOUT_PROCESSING_TTL_SECONDS)
	require.NoError(t, err)

	// 验证
	assert.Equal(t, len(followers), successCount, "All followers should receive the post")

	// 验证每个粉丝的收件箱
	for _, followerID := range followers {
		inboxKey := fmt.Sprintf(FEED_INBOX_KEY, followerID)
		posts, err := redisAdapter.ZRevRangeWithScores(ctx, inboxKey, 0, 10)
		require.NoError(t, err)
		postIDs := make([]int64, len(posts))
		for i, z := range posts {
			postIDs[i] = z.Member
		}
		assert.Contains(t, postIDs, event.PostID,
			fmt.Sprintf("Follower %d should have the post", followerID))
	}

	// 验证幂等性
	exists, err = redisAdapter.Exists(ctx, processKey)
	require.NoError(t, err)
	assert.True(t, exists, "Should be marked as processed")

	t.Logf("✓ Fanout worker: pushed to %d/%d followers", successCount, len(followers))
}

// ZCard helper for RedisAdapter
func (r *RedisAdapter) ZCard(ctx context.Context, key string) (int64, error) {
	return r.client.ZCard(ctx, key).Result()
}

// deleteFeedIntegrationKeys only deletes the fixed keys owned by this test.
func deleteFeedIntegrationKeys(client *redis.Client, ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return client.Del(ctx, keys...).Err()
}

// Mock implementations

type mockCounterClient struct {
	followerCounts map[int64]int64
}

func (m *mockCounterClient) GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	if m.followerCounts == nil {
		return 0, nil
	}
	return m.followerCounts[userID], nil
}

func (m *mockCounterClient) BatchGetFollowerCounts(ctx context.Context, userIDs []int64) (map[int64]int64, error) {
	if m.followerCounts == nil {
		return make(map[int64]int64), nil
	}
	result := make(map[int64]int64)
	for _, uid := range userIDs {
		result[uid] = m.followerCounts[uid]
	}
	return result, nil
}

type mockRelationClient struct {
	followings []int64
	followers  []int64
}

func (m *mockRelationClient) GetFollowers(ctx context.Context, userID int64) ([]int64, error) {
	if m.followers == nil {
		return []int64{}, nil
	}
	return m.followers, nil
}

func (m *mockRelationClient) GetFollowings(ctx context.Context, userID int64) ([]int64, error) {
	if m.followings == nil {
		return []int64{}, nil
	}
	return m.followings, nil
}

type mockKafkaProducer struct {
	sent      bool
	lastEvent FeedEvent
}

func (m *mockKafkaProducer) SendMessage(ctx context.Context, topic, key string, value []byte) error {
	m.sent = true
	// 解析事件
	var event FeedEvent
	if err := json.Unmarshal(value, &event); err == nil {
		m.lastEvent = event
	}
	return nil
}
