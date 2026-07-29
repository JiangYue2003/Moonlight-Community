package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// FeedEvent Feed扇出事件
type FeedEvent struct {
	PostID     int64 `json:"post_id"`
	CreatorID  int64 `json:"creator_id"`
	CreateTime int64 `json:"create_time"` // Unix时间戳
}

// FeedWriter 负责发帖时的推送逻辑
type FeedWriter struct {
	redis  RedisClient
	kafka  KafkaProducer
	logger logx.Logger
}

// RedisClient Redis客户端接口（方便测试）
type RedisClient interface {
	ZAdd(ctx context.Context, key string, members ...interface{}) error
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) error
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZScore, error)
	Expire(ctx context.Context, key string, seconds int) error
	Exists(ctx context.Context, key string) (bool, error)
	SetEx(ctx context.Context, key string, value interface{}, seconds int) error
	SCard(ctx context.Context, key string) (int64, error)
}

// ZScore Redis ZSet 的成员和分数
type ZScore struct {
	Member int64
	Score  float64
}

// KafkaProducer Kafka生产者接口
type KafkaProducer interface {
	SendMessage(ctx context.Context, topic string, key string, value []byte) error
}

func NewFeedWriter(redis RedisClient, kafka KafkaProducer, logger logx.Logger) *FeedWriter {
	return &FeedWriter{
		redis:  redis,
		kafka:  kafka,
		logger: logger,
	}
}

// OnPostPublished 发帖时调用：根据粉丝数决定推模式或拉模式
func (w *FeedWriter) OnPostPublished(ctx context.Context, postID, creatorID int64, followerCount int64) error {
	now := time.Now().Unix()

	// 1. 写入自己的收件箱（同步，保证自己能立刻看到）
	if err := w.pushToInbox(ctx, creatorID, postID, now); err != nil {
		w.logger.Errorf("failed to push to creator's inbox: %v", err)
		// 不返回错误，不影响发帖成功
	}

	// 2. 判断是否大V
	if followerCount <= BIGV_THRESHOLD {
		// 推模式：发送Kafka消息，由Worker异步推送给粉丝
		w.logger.Infof("post %d: push mode (followers=%d)", postID, followerCount)
		return w.sendFanoutEvent(ctx, postID, creatorID, now)
	} else {
		// 拉模式：写入大V发件箱
		w.logger.Infof("post %d: pull mode (followers=%d, bigv)", postID, followerCount)
		return w.pushToBigVOutbox(ctx, creatorID, postID, now)
	}
}

// pushToInbox 推送到用户收件箱
func (w *FeedWriter) pushToInbox(ctx context.Context, userID, postID, timestamp int64) error {
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)

	// 1. 添加到ZSet（score=时间戳，member=postID）
	if err := w.redis.ZAdd(ctx, key, float64(timestamp), postID); err != nil {
		return fmt.Errorf("zadd failed: %w", err)
	}

	// 2. 保留最新的N条（删除排名在后面的）
	// ZRemRangeByRank 保留 rank [0, INBOX_MAX_SIZE-1]，删除 [INBOX_MAX_SIZE, -1]
	if err := w.redis.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1); err != nil {
		w.logger.Errorf("zremrangebyrank failed: %v", err)
		// 不返回错误，容量控制失败不影响主流程
	}

	// 3. 设置过期时间
	if err := w.redis.Expire(ctx, key, INBOX_TTL_SECONDS); err != nil {
		w.logger.Errorf("expire failed: %v", err)
	}

	return nil
}

// pushToBigVOutbox 推送到大V发件箱
func (w *FeedWriter) pushToBigVOutbox(ctx context.Context, creatorID, postID, timestamp int64) error {
	key := fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, creatorID)

	// 1. 添加到ZSet
	if err := w.redis.ZAdd(ctx, key, float64(timestamp), postID); err != nil {
		return fmt.Errorf("zadd bigv outbox failed: %w", err)
	}

	// 2. 保留最新的N条
	if err := w.redis.ZRemRangeByRank(ctx, key, 0, -BIGV_OUTBOX_MAX_SIZE-1); err != nil {
		w.logger.Errorf("zremrangebyrank bigv outbox failed: %v", err)
	}

	// 3. 设置过期时间
	if err := w.redis.Expire(ctx, key, BIGV_OUTBOX_TTL_SECONDS); err != nil {
		w.logger.Errorf("expire bigv outbox failed: %v", err)
	}

	return nil
}

// sendFanoutEvent 发送扇出事件到Kafka
func (w *FeedWriter) sendFanoutEvent(ctx context.Context, postID, creatorID, createTime int64) error {
	if w.kafka == nil {
		return fmt.Errorf("kafka producer is not configured")
	}

	event := FeedEvent{
		PostID:     postID,
		CreatorID:  creatorID,
		CreateTime: createTime,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event failed: %w", err)
	}

	// 使用 postID 作为 key，保证同一个帖子的事件顺序
	key := fmt.Sprintf("%d", postID)

	if err := w.kafka.SendMessage(ctx, FEED_FANOUT_TOPIC, key, payload); err != nil {
		return fmt.Errorf("send kafka message failed: %w", err)
	}

	w.logger.Infof("sent fanout event: post=%d, creator=%d", postID, creatorID)
	return nil
}
