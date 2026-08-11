package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/service"
)

const fanoutBatchSize = 100

const fanoutRetryDelay = time.Second

type fanoutStateStore interface {
	Exists(ctx context.Context, key string) (bool, error)
	SetEx(ctx context.Context, key string, value interface{}, seconds int) error
}

// InboxBatchWriter 批量维护粉丝收件箱；实现负责合并 Redis 网络往返。
type InboxBatchWriter interface {
	PushToInboxes(ctx context.Context, userIDs []int64, postID, timestamp int64) error
}

type sequentialInboxBatchWriter struct {
	redis RedisClient
}

func (w *sequentialInboxBatchWriter) PushToInboxes(
	ctx context.Context,
	userIDs []int64,
	postID, timestamp int64,
) error {
	for _, userID := range userIDs {
		key := fmt.Sprintf(FEED_INBOX_KEY, userID)
		if err := w.redis.ZAdd(ctx, key, float64(timestamp), postID); err != nil {
			return fmt.Errorf("zadd inbox failed: user=%d: %w", userID, err)
		}
		if err := w.redis.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1); err != nil {
			return fmt.Errorf("trim inbox failed: user=%d: %w", userID, err)
		}
		if err := w.redis.Expire(ctx, key, INBOX_TTL_SECONDS); err != nil {
			return fmt.Errorf("expire inbox failed: user=%d: %w", userID, err)
		}
	}
	return nil
}

func inboxBatchWriterFor(redis RedisClient) InboxBatchWriter {
	if writer, ok := redis.(InboxBatchWriter); ok {
		return writer
	}
	return &sequentialInboxBatchWriter{redis: redis}
}

type cancelableMessageQueue struct {
	queue.MessageQueue
	cancel context.CancelFunc
}

func (q *cancelableMessageQueue) Stop() {
	q.cancel()
	q.MessageQueue.Stop()
}

type feedFanoutHandler struct {
	state          fanoutStateStore
	inboxes        InboxBatchWriter
	relationClient RelationClient
	logger         logx.Logger
}

func newFeedFanoutHandler(
	state fanoutStateStore,
	inboxes InboxBatchWriter,
	relationClient RelationClient,
	logger logx.Logger,
) *feedFanoutHandler {
	return &feedFanoutHandler{
		state:          state,
		inboxes:        inboxes,
		relationClient: relationClient,
		logger:         logger,
	}
}

func (h *feedFanoutHandler) Handle(ctx context.Context, event FeedEvent) error {
	processKey := fmt.Sprintf(FANOUT_PROCESSING_KEY, event.PostID)
	exists, err := h.state.Exists(ctx, processKey)
	if err != nil {
		h.logger.Errorf("check processing key failed: %v", err)
	} else if exists {
		h.logger.Infof("event already processed: post=%d", event.PostID)
		return nil
	}

	followers, err := h.relationClient.GetFollowers(ctx, event.CreatorID)
	if err != nil {
		return fmt.Errorf("get followers failed: creator=%d: %w", event.CreatorID, err)
	}

	for start := 0; start < len(followers); start += fanoutBatchSize {
		end := start + fanoutBatchSize
		if end > len(followers) {
			end = len(followers)
		}
		if err := h.inboxes.PushToInboxes(ctx, followers[start:end], event.PostID, event.CreateTime); err != nil {
			return fmt.Errorf("push fanout batch [%d:%d] failed: %w", start, end, err)
		}
	}

	if err := h.state.SetEx(ctx, processKey, "1", FANOUT_PROCESSING_TTL_SECONDS); err != nil {
		return fmt.Errorf("mark processed failed: key=%s: %w", processKey, err)
	}
	h.logger.Infof("fanout completed: post=%d, success=%d/%d", event.PostID, len(followers), len(followers))
	return nil
}

func retryFeedFanout(
	ctx context.Context,
	event FeedEvent,
	handle func(context.Context, FeedEvent) error,
	retryDelay time.Duration,
) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := handle(ctx, event); err == nil {
			return nil
		}

		timer := time.NewTimer(retryDelay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return ctx.Err()
		case <-timer.C:
		}
	}
}

func consumeFeedFanoutMessage(
	workerCtx context.Context,
	messageCtx context.Context,
	rawEvent string,
	handle func(context.Context, FeedEvent) error,
	logger logx.Logger,
	retryDelay time.Duration,
) error {
	// Stop may leave one message prefetched behind the currently failing offset.
	// Reject it before parsing so it cannot advance Kafka's committed offset.
	if err := workerCtx.Err(); err != nil {
		return err
	}

	var event FeedEvent
	if err := json.Unmarshal([]byte(rawEvent), &event); err != nil {
		logger.Errorf("unmarshal event failed: %v", err)
		return nil
	}

	logger.Infof("processing fanout event: post=%d, creator=%d", event.PostID, event.CreatorID)
	eventCtx, cancelEvent := context.WithCancel(messageCtx)
	stopCancelPropagation := context.AfterFunc(workerCtx, cancelEvent)
	if err := workerCtx.Err(); err != nil {
		cancelEvent()
	}
	defer func() {
		stopCancelPropagation()
		cancelEvent()
	}()

	return retryFeedFanout(eventCtx, event, func(attemptCtx context.Context, attempt FeedEvent) error {
		if err := workerCtx.Err(); err != nil {
			return err
		}
		err := handle(attemptCtx, attempt)
		if err != nil {
			logger.Errorf("fanout attempt failed, retrying: post=%d, err=%v", attempt.PostID, err)
		}
		return err
	}, retryDelay)
}

func newFeedFanoutQueueConfig(brokers []string) kq.KqConf {
	return kq.KqConf{
		ServiceConf: service.ServiceConf{Name: "feed-fanout-consumer"},
		Brokers:     brokers,
		Group:       "feed-fanout-group",
		Topic:       FEED_FANOUT_TOPIC,
		Offset:      "last",
		Conns:       1,
		Consumers:   1,
		Processors:  1,
		MinBytes:    1,
		MaxBytes:    10 * 1024 * 1024,
		ForceCommit: false,
	}
}

// NewFeedFanoutWorker 创建 Feed 扇出 Worker，由服务生命周期负责启动和停止。
func NewFeedFanoutWorker(
	brokers []string,
	redis RedisClient,
	relationClient RelationClient,
	logger logx.Logger,
) queue.MessageQueue {
	handler := newFeedFanoutHandler(redis, inboxBatchWriterFor(redis), relationClient, logger)
	workerCtx, cancelWorker := context.WithCancel(context.Background())

	// 创建 Kafka consumer
	worker := kq.MustNewQueue(newFeedFanoutQueueConfig(brokers), kq.WithHandle(func(ctx context.Context, k, v string) error {
		return consumeFeedFanoutMessage(workerCtx, ctx, v, handler.Handle, logger, fanoutRetryDelay)
	}))
	return &cancelableMessageQueue{MessageQueue: worker, cancel: cancelWorker}
}

// StartFeedFanoutWorker 启动并阻塞运行 Feed 扇出 Worker。
func StartFeedFanoutWorker(
	brokers []string,
	redis RedisClient,
	relationClient RelationClient,
	logger logx.Logger,
) error {
	worker := NewFeedFanoutWorker(brokers, redis, relationClient, logger)
	logger.Info("feed fanout worker started")
	worker.Start()
	return nil
}
