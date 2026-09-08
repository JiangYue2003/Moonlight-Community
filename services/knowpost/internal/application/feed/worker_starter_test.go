package feed

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
)

var _ func([]string, RedisClient, RelationClient, logx.Logger) queue.MessageQueue = NewFeedFanoutWorker
var _ func([]string, RedisClient, RelationClient, logx.Logger) error = StartFeedFanoutWorker

type fanoutStateStoreStub struct {
	exists    bool
	existsErr error
	setErr    error
	setCalls  int
}

func (s *fanoutStateStoreStub) Exists(context.Context, string) (bool, error) {
	return s.exists, s.existsErr
}

func (s *fanoutStateStoreStub) SetEx(context.Context, string, interface{}, int) error {
	s.setCalls++
	return s.setErr
}

type inboxBatchWriterStub struct {
	batches  [][]int64
	failCall int
}

func (w *inboxBatchWriterStub) PushToInboxes(
	_ context.Context,
	userIDs []int64,
	_, _ int64,
) error {
	w.batches = append(w.batches, append([]int64(nil), userIDs...))
	if w.failCall > 0 && len(w.batches) == w.failCall {
		return errors.New("pipeline failed")
	}
	return nil
}

type redisWithNativeBatchStub struct {
	*MockRedisClient
	batchCalls int
}

func (r *redisWithNativeBatchStub) PushToInboxes(context.Context, []int64, int64, int64) error {
	r.batchCalls++
	return nil
}

type fanoutRelationStub struct {
	followers []int64
	err       error
}

func (r *fanoutRelationStub) GetFollowers(context.Context, int64) ([]int64, error) {
	return append([]int64(nil), r.followers...), r.err
}

func (r *fanoutRelationStub) GetFollowings(context.Context, int64) ([]int64, error) {
	return nil, nil
}

type messageQueueStub struct {
	stopped bool
}

func (*messageQueueStub) Start() {}
func (q *messageQueueStub) Stop() {
	q.stopped = true
}

func TestFeedFanoutQueueConfigStartsConsumers(t *testing.T) {
	config := newFeedFanoutQueueConfig([]string{"127.0.0.1:9092"})

	require.Equal(t, FEED_FANOUT_TOPIC, config.Topic)
	require.Equal(t, "feed-fanout-group", config.Group)
	require.Positive(t, config.Conns)
	require.Positive(t, config.Consumers)
	require.Equal(t, 1, config.Processors, "serial processing prevents later offsets from skipping a failed event")
	require.Positive(t, config.MinBytes)
	require.GreaterOrEqual(t, config.MaxBytes, config.MinBytes)
	require.False(t, config.ForceCommit, "handler errors must leave the Kafka offset uncommitted")
}

func TestInboxBatchWriterForUsesNativeBatchImplementation(t *testing.T) {
	redis := &redisWithNativeBatchStub{MockRedisClient: NewMockRedisClient()}

	err := inboxBatchWriterFor(redis).PushToInboxes(context.Background(), []int64{1, 2}, 11, 33)

	require.NoError(t, err)
	require.Equal(t, 1, redis.batchCalls)
	require.Empty(t, redis.zaddCalls)
}

func TestInboxBatchWriterForSupportsLegacyRedisClient(t *testing.T) {
	redis := NewMockRedisClient()

	err := inboxBatchWriterFor(redis).PushToInboxes(context.Background(), []int64{1, 2}, 11, 33)

	require.NoError(t, err)
	require.Len(t, redis.zaddCalls, 2)
	require.Len(t, redis.expireCalls, 2)
}

func TestFeedFanoutHandlerWritesFollowersInBoundedBatches(t *testing.T) {
	followers := make([]int64, 205)
	for i := range followers {
		followers[i] = int64(i + 1)
	}
	state := &fanoutStateStoreStub{}
	inboxes := &inboxBatchWriterStub{}
	handler := newFeedFanoutHandler(
		state,
		inboxes,
		&fanoutRelationStub{followers: followers},
		logx.WithContext(context.Background()),
	)

	err := handler.Handle(context.Background(), FeedEvent{PostID: 11, CreatorID: 22, CreateTime: 33})

	require.NoError(t, err)
	require.Len(t, inboxes.batches, 3)
	require.Len(t, inboxes.batches[0], 100)
	require.Len(t, inboxes.batches[1], 100)
	require.Len(t, inboxes.batches[2], 5)
	require.Equal(t, 1, state.setCalls)
}

func TestFeedFanoutHandlerLeavesEventUnprocessedWhenBatchFails(t *testing.T) {
	followers := make([]int64, 205)
	state := &fanoutStateStoreStub{}
	inboxes := &inboxBatchWriterStub{failCall: 2}
	handler := newFeedFanoutHandler(
		state,
		inboxes,
		&fanoutRelationStub{followers: followers},
		logx.WithContext(context.Background()),
	)

	err := handler.Handle(context.Background(), FeedEvent{PostID: 11, CreatorID: 22, CreateTime: 33})

	require.ErrorContains(t, err, "pipeline failed")
	require.Len(t, inboxes.batches, 2)
	require.Zero(t, state.setCalls)
}

func TestFeedFanoutHandlerReturnsErrorWhenMarkProcessedFails(t *testing.T) {
	state := &fanoutStateStoreStub{setErr: errors.New("mark failed")}
	inboxes := &inboxBatchWriterStub{}
	handler := newFeedFanoutHandler(
		state,
		inboxes,
		&fanoutRelationStub{followers: []int64{1}},
		logx.WithContext(context.Background()),
	)

	err := handler.Handle(context.Background(), FeedEvent{PostID: 11, CreatorID: 22, CreateTime: 33})

	require.ErrorContains(t, err, "mark failed")
	require.Len(t, inboxes.batches, 1)
	require.Equal(t, 1, state.setCalls)
}

func TestFeedFanoutHandlerMarksEmptyFanoutProcessed(t *testing.T) {
	state := &fanoutStateStoreStub{}
	inboxes := &inboxBatchWriterStub{}
	handler := newFeedFanoutHandler(
		state,
		inboxes,
		&fanoutRelationStub{},
		logx.WithContext(context.Background()),
	)

	err := handler.Handle(context.Background(), FeedEvent{PostID: 11, CreatorID: 22, CreateTime: 33})

	require.NoError(t, err)
	require.Empty(t, inboxes.batches)
	require.Equal(t, 1, state.setCalls)
}

func TestRetryFeedFanoutRetriesWholeEventUntilSuccess(t *testing.T) {
	attempts := 0
	event := FeedEvent{PostID: 11, CreatorID: 22, CreateTime: 33}

	err := retryFeedFanout(context.Background(), event, func(context.Context, FeedEvent) error {
		attempts++
		if attempts == 1 {
			return errors.New("temporary failure")
		}
		return nil
	}, time.Millisecond)

	require.NoError(t, err)
	require.Equal(t, 2, attempts)
}

func TestRetryFeedFanoutStopsBeforeAttemptWhenCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	attempts := 0

	err := retryFeedFanout(ctx, FeedEvent{}, func(context.Context, FeedEvent) error {
		attempts++
		return errors.New("should not run")
	}, time.Millisecond)

	require.ErrorIs(t, err, context.Canceled)
	require.Zero(t, attempts)
}

func TestConsumeFeedFanoutMessageRejectsPrefetchedMessageAfterStop(t *testing.T) {
	workerCtx, cancelWorker := context.WithCancel(context.Background())
	cancelWorker()
	handleCalls := 0

	err := consumeFeedFanoutMessage(
		workerCtx,
		context.Background(),
		"{malformed-json",
		func(context.Context, FeedEvent) error {
			handleCalls++
			return nil
		},
		logx.WithContext(context.Background()),
		time.Millisecond,
	)

	require.ErrorIs(t, err, context.Canceled)
	require.Zero(t, handleCalls)
}

func TestCancelableMessageQueueStopsRetriesBeforeInnerQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	inner := &messageQueueStub{}
	worker := &cancelableMessageQueue{MessageQueue: inner, cancel: cancel}

	worker.Stop()

	require.ErrorIs(t, ctx.Err(), context.Canceled)
	require.True(t, inner.stopped)
}
