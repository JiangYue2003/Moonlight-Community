package feed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

// StartFeedFanoutWorker 启动 Feed 扇出 Worker
func StartFeedFanoutWorker(
	brokers []string,
	redis RedisClient,
	relationClient RelationClient,
	logger logx.Logger,
) error {

	// 创建 Kafka consumer
	_ = kq.MustNewQueue(kq.KqConf{
		Brokers: brokers,
		Group:   "feed-fanout-group",
		Topic:   FEED_FANOUT_TOPIC,
	}, kq.WithHandle(func(ctx context.Context, k, v string) error {
		// 处理消息
		var event FeedEvent
		if err := json.Unmarshal([]byte(v), &event); err != nil {
			logger.Errorf("unmarshal event failed: %v", err)
			return nil // 解析失败，跳过
		}

		logger.Infof("processing fanout event: post=%d, creator=%d", event.PostID, event.CreatorID)

		// 幂等性检查
		processKey := fmt.Sprintf(FANOUT_PROCESSING_KEY, event.PostID)
		exists, err := redis.Exists(ctx, processKey)
		if err != nil {
			logger.Errorf("check processing key failed: %v", err)
		} else if exists {
			logger.Infof("event already processed: post=%d", event.PostID)
			return nil
		}

		// 获取粉丝列表
		followers, err := relationClient.GetFollowers(ctx, event.CreatorID)
		if err != nil {
			logger.Errorf("get followers failed: creator=%d, err=%v", event.CreatorID, err)
			return err // 返回错误，Kafka会重试
		}

		if len(followers) == 0 {
			logger.Infof("no followers for creator=%d", event.CreatorID)
			markProcessed(ctx, redis, processKey, logger)
			return nil
		}

		logger.Infof("fanout to %d followers", len(followers))

		// 批量推送到粉丝收件箱
		successCount := 0
		for _, followerID := range followers {
			if err := pushToInbox(ctx, redis, followerID, event.PostID, event.CreateTime, logger); err != nil {
				logger.Errorf("push to inbox failed: follower=%d, post=%d, err=%v",
					followerID, event.PostID, err)
			} else {
				successCount++
			}
		}

		logger.Infof("fanout completed: post=%d, success=%d/%d",
			event.PostID, successCount, len(followers))

		// 标记已处理
		markProcessed(ctx, redis, processKey, logger)

		return nil
	}))

	// queue.Start() 会在 MustNewQueue 内部自动调用

	logger.Info("feed fanout worker started")

	return nil
}

// 辅助函数
func pushToInbox(ctx context.Context, redis RedisClient, userID, postID, timestamp int64, logger logx.Logger) error {
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)

	if err := redis.ZAdd(ctx, key, float64(timestamp), postID); err != nil {
		return err
	}

	if err := redis.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1); err != nil {
		logger.Errorf("zremrangebyrank failed: user=%d, err=%v", userID, err)
	}

	if err := redis.Expire(ctx, key, INBOX_TTL_SECONDS); err != nil {
		logger.Errorf("expire failed: user=%d, err=%v", userID, err)
	}

	return nil
}

func markProcessed(ctx context.Context, redis RedisClient, processKey string, logger logx.Logger) {
	if err := redis.SetEx(ctx, processKey, "1", FANOUT_PROCESSING_TTL_SECONDS); err != nil {
		logger.Errorf("mark processed failed: key=%s, err=%v", processKey, err)
	}
}

