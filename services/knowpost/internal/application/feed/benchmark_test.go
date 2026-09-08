//go:build integration
// +build integration

package feed

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

// 性能测试：测试推拉结合Feed架构的性能
// 运行方式：go test -tags=integration -bench=. -benchmem -benchtime=10s ./services/knowpost/rpc/internal/feed/

// BenchmarkFeedWriter_PushMode 测试推模式写入性能
func BenchmarkFeedWriter_PushMode(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)
	mockKafka := &mockKafkaProducer{}
	logger := logx.WithContext(ctx)
	writer := NewFeedWriter(redisAdapter, mockKafka, logger)

	creatorID := int64(10001)
	followerCount := int64(500) // 推模式

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		postID := int64(20000)
		for pb.Next() {
			postID++
			err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)
			if err != nil {
				b.Errorf("OnPostPublished failed: %v", err)
			}
		}
	})
}

// BenchmarkFeedWriter_PullMode 测试拉模式写入性能
func BenchmarkFeedWriter_PullMode(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)
	mockKafka := &mockKafkaProducer{}
	logger := logx.WithContext(ctx)
	writer := NewFeedWriter(redisAdapter, mockKafka, logger)

	creatorID := int64(10002)
	followerCount := int64(5000) // 拉模式

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		postID := int64(30000)
		for pb.Next() {
			postID++
			err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)
			if err != nil {
				b.Errorf("OnPostPublished failed: %v", err)
			}
		}
	})
}

// BenchmarkFeedReader_SmallFollowings 测试小规模关注列表的读取性能
func BenchmarkFeedReader_SmallFollowings(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)

	// 准备测试数据
	userID := int64(40001)
	setupFeedReaderBenchmark(b, redisAdapter, ctx, userID, 10) // 10个关注

	mockCounter := &mockCounterClient{
		followerCounts: map[int64]int64{
			50001: 500,
			50002: 500,
			50003: 600,
			50004: 700,
			50005: 800,
			50006: 900,
			50007: 5000, // 大V
			50008: 6000, // 大V
			50009: 7000, // 大V
			50010: 8000, // 大V
		},
	}

	mockRelation := &mockRelationClient{
		followings: []int64{50001, 50002, 50003, 50004, 50005, 50006, 50007, 50008, 50009, 50010},
	}

	logger := logx.WithContext(ctx)
	reader := NewFeedReader(redisAdapter, mockRelation, mockCounter, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, err := reader.GetFeed(ctx, userID, 1, 20)
			if err != nil {
				b.Errorf("GetFeed failed: %v", err)
			}
		}
	})
}

// BenchmarkFeedReader_MediumFollowings 测试中等规模关注列表的读取性能
func BenchmarkFeedReader_MediumFollowings(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)

	// 准备测试数据：50个关注
	userID := int64(40002)
	followingCount := 50
	setupFeedReaderBenchmark(b, redisAdapter, ctx, userID, followingCount)

	// 构造粉丝数映射：前30个普通用户，后20个大V
	followerCounts := make(map[int64]int64)
	followings := make([]int64, followingCount)
	for i := 0; i < followingCount; i++ {
		uid := int64(60001 + i)
		followings[i] = uid
		if i < 30 {
			followerCounts[uid] = 500 // 普通用户
		} else {
			followerCounts[uid] = 5000 // 大V
		}
	}

	mockCounter := &mockCounterClient{followerCounts: followerCounts}
	mockRelation := &mockRelationClient{followings: followings}
	logger := logx.WithContext(ctx)
	reader := NewFeedReader(redisAdapter, mockRelation, mockCounter, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, err := reader.GetFeed(ctx, userID, 1, 20)
			if err != nil {
				b.Errorf("GetFeed failed: %v", err)
			}
		}
	})
}

// BenchmarkFeedReader_LargeFollowings 测试大规模关注列表的读取性能
func BenchmarkFeedReader_LargeFollowings(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)

	// 准备测试数据：100个关注
	userID := int64(40003)
	followingCount := 100
	setupFeedReaderBenchmark(b, redisAdapter, ctx, userID, followingCount)

	// 构造粉丝数映射：前60个普通用户，后40个大V
	followerCounts := make(map[int64]int64)
	followings := make([]int64, followingCount)
	for i := 0; i < followingCount; i++ {
		uid := int64(70001 + i)
		followings[i] = uid
		if i < 60 {
			followerCounts[uid] = 500 // 普通用户
		} else {
			followerCounts[uid] = 5000 // 大V
		}
	}

	mockCounter := &mockCounterClient{followerCounts: followerCounts}
	mockRelation := &mockRelationClient{followings: followings}
	logger := logx.WithContext(ctx)
	reader := NewFeedReader(redisAdapter, mockRelation, mockCounter, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, err := reader.GetFeed(ctx, userID, 1, 20)
			if err != nil {
				b.Errorf("GetFeed failed: %v", err)
			}
		}
	})
}

// BenchmarkFanout_BatchWrite 测试批量扇出写入性能
func BenchmarkFanout_BatchWrite(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)

	// 模拟1000个粉丝
	followerCount := 1000
	followers := make([]int64, followerCount)
	for i := 0; i < followerCount; i++ {
		followers[i] = int64(80001 + i)
	}

	postID := int64(90000)
	timestamp := time.Now().Unix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postID++
		// 批量写入粉丝收件箱
		for _, followerID := range followers {
			inboxKey := fmt.Sprintf(FEED_INBOX_KEY, followerID)
			_ = redisAdapter.ZAdd(ctx, inboxKey, float64(timestamp), postID)
		}
	}
}

// BenchmarkRedis_ZAddAndTrim 测试Redis ZAdd + Trim操作性能
func BenchmarkRedis_ZAddAndTrim(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	redisAdapter := NewRedisAdapter(redisClient)

	key := "feed:inbox:benchmark"
	postID := int64(100000)
	timestamp := time.Now().Unix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postID++
		// ZAdd
		_ = redisAdapter.ZAdd(ctx, key, float64(timestamp), postID)
		// Trim
		_ = redisAdapter.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1)
	}
}

// BenchmarkRedis_Pipeline 测试Redis Pipeline性能
func BenchmarkRedis_Pipeline(b *testing.B) {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer redisClient.Close()

	postID := int64(110000)
	timestamp := float64(time.Now().Unix())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postID++

		// 使用Pipeline批量写入
		pipe := redisClient.Pipeline()
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("feed:inbox:%d", 90001+j)
			pipe.ZAdd(ctx, key, redis.Z{Score: timestamp, Member: postID})
		}
		_, _ = pipe.Exec(ctx)
	}
}

// setupFeedReaderBenchmark 准备Feed读取测试数据
func setupFeedReaderBenchmark(b *testing.B, redis RedisClient, ctx context.Context, userID int64, followingCount int) {
	// 填充用户收件箱
	inboxKey := fmt.Sprintf(FEED_INBOX_KEY, userID)
	baseTime := time.Now().Unix()
	for i := 0; i < 50; i++ {
		postID := int64(200000 + i)
		err := redis.ZAdd(ctx, inboxKey, float64(baseTime+int64(i)), postID)
		require.NoError(b, err)
	}

	// 填充大V发件箱（假设后半部分是大V）
	for i := followingCount / 2; i < followingCount; i++ {
		creatorID := int64(50001 + i)
		outboxKey := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID)
		for j := 0; j < 20; j++ {
			postID := int64(300000 + i*100 + j)
			err := redis.ZAdd(ctx, outboxKey, float64(baseTime+int64(j)), postID)
			require.NoError(b, err)
		}
	}
}
