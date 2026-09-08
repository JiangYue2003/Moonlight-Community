package feed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

// FanoutWorker 消费Kafka消息，将帖子推送到粉丝收件箱
type FanoutWorker struct {
	redis          RedisClient
	relationClient RelationClient
	kafka          KafkaConsumer
	logger         logx.Logger
	queue          interface{} // kq.Queue
}

// RelationClient 关系服务客户端接口
type RelationClient interface {
	// GetFollowers 获取用户的粉丝列表
	GetFollowers(ctx context.Context, userID int64) ([]int64, error)
	// GetFollowings 获取用户关注的人
	GetFollowings(ctx context.Context, userID int64) ([]int64, error)
}

// KafkaConsumer Kafka消费者接口
type KafkaConsumer interface {
	ConsumeMessages(ctx context.Context, topic, groupID string, handler func(key, value []byte) error) error
}

func NewFanoutWorker(redis RedisClient, relationClient RelationClient, kafka KafkaConsumer, logger logx.Logger) *FanoutWorker {
	return &FanoutWorker{
		redis:          redis,
		relationClient: relationClient,
		kafka:          kafka,
		logger:         logger,
	}
}

// Start 启动Worker，消费Kafka消息
func (w *FanoutWorker) Start(ctx context.Context) error {
	w.logger.Info("fanout worker started")

	return w.kafka.ConsumeMessages(ctx, FEED_FANOUT_TOPIC, "feed-fanout-group", func(key, value []byte) error {
		return w.handleFanoutEvent(ctx, value)
	})
}

// handleFanoutEvent 处理单个扇出事件
func (w *FanoutWorker) handleFanoutEvent(ctx context.Context, payload []byte) error {
	var event FeedEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		w.logger.Errorf("unmarshal event failed: %v", err)
		return nil // 解析失败，跳过这条消息
	}

	w.logger.Infof("processing fanout event: post=%d, creator=%d", event.PostID, event.CreatorID)

	// 1. 幂等性检查（防止重复处理）
	processKey := fmt.Sprintf(FANOUT_PROCESSING_KEY, event.PostID)
	exists, err := w.redis.Exists(ctx, processKey)
	if err != nil {
		w.logger.Errorf("check processing key failed: %v", err)
		// 继续处理，不因为检查失败而跳过
	} else if exists {
		w.logger.Infof("event already processed: post=%d", event.PostID)
		return nil // 已处理过，跳过
	}

	// 2. 获取粉丝列表
	followers, err := w.relationClient.GetFollowers(ctx, event.CreatorID)
	if err != nil {
		w.logger.Errorf("get followers failed: creator=%d, err=%v", event.CreatorID, err)
		return err // 返回错误，Kafka会重试
	}

	if len(followers) == 0 {
		w.logger.Infof("no followers for creator=%d", event.CreatorID)
		// 标记已处理
		w.markProcessed(ctx, processKey)
		return nil
	}

	w.logger.Infof("fanout to %d followers", len(followers))

	// 3. 批量推送到粉丝收件箱
	successCount := 0
	for _, followerID := range followers {
		if err := w.pushToInbox(ctx, followerID, event.PostID, event.CreateTime); err != nil {
			w.logger.Errorf("push to inbox failed: follower=%d, post=%d, err=%v",
				followerID, event.PostID, err)
			// 继续推送其他粉丝，不中断
		} else {
			successCount++
		}
	}

	w.logger.Infof("fanout completed: post=%d, success=%d/%d",
		event.PostID, successCount, len(followers))

	// 4. 标记已处理
	w.markProcessed(ctx, processKey)

	return nil
}

// pushToInbox 推送到单个粉丝的收件箱
func (w *FanoutWorker) pushToInbox(ctx context.Context, userID, postID, timestamp int64) error {
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)

	// 1. 添加到ZSet
	if err := w.redis.ZAdd(ctx, key, float64(timestamp), postID); err != nil {
		return err
	}

	// 2. 保留最新的N条
	if err := w.redis.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1); err != nil {
		w.logger.Errorf("zremrangebyrank failed: user=%d, err=%v", userID, err)
	}

	// 3. 设置过期时间
	if err := w.redis.Expire(ctx, key, INBOX_TTL_SECONDS); err != nil {
		w.logger.Errorf("expire failed: user=%d, err=%v", userID, err)
	}

	return nil
}

// markProcessed 标记事件已处理（幂等性）
func (w *FanoutWorker) markProcessed(ctx context.Context, processKey string) {
	if err := w.redis.SetEx(ctx, processKey, "1", FANOUT_PROCESSING_TTL_SECONDS); err != nil {
		w.logger.Errorf("mark processed failed: key=%s, err=%v", processKey, err)
	}
}
