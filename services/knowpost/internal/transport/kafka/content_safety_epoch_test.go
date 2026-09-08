package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/config"
	knowpostevent "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
)

func TestContentSafetyEpochProcessorBumpsUpdatedAndDeletedRows(t *testing.T) {
	store := &recordingSafetyEpochStore{}
	processor := newContentSafetyEpochProcessor(store, newLogThrottle(time.Hour))
	message := relationCanalMessage(t,
		knowpostOutboxRow(t, 1, knowpostevent.TypeKnowPostUpdated),
		knowpostOutboxRow(t, 2, knowpostevent.TypeKnowPostDeleted),
		knowpostOutboxRow(t, 3, knowpostevent.TypeKnowPostPublished),
		map[string]string{
			"id": "4", "aggregate_type": "following", "aggregate_id": "9", "type": "FollowCreated", "payload": `{"type":"FollowCreated"}`,
		},
	)

	require.NoError(t, processor.Handle(context.Background(), message))
	require.Equal(t, 2, store.Calls())
}

func TestContentSafetyEpochProcessorSkipsPoisonMessages(t *testing.T) {
	store := &recordingSafetyEpochStore{}
	processor := newContentSafetyEpochProcessor(store, newLogThrottle(time.Hour))

	require.NoError(t, processor.Handle(context.Background(), []byte("not-json")))
	require.NoError(t, processor.Handle(context.Background(), relationCanalMessage(t, map[string]string{
		"id": "1", "aggregate_type": knowpostevent.AggregateType, "aggregate_id": "9", "type": knowpostevent.TypeKnowPostDeleted, "payload": "{bad",
	})))
	require.Equal(t, 0, store.Calls())
}

func TestContentSafetyEpochProcessorReturnsStoreErrorForKafkaRetry(t *testing.T) {
	store := &recordingSafetyEpochStore{err: errors.New("redis unavailable")}
	processor := newContentSafetyEpochProcessor(store, newLogThrottle(time.Hour))
	message := relationCanalMessage(t, knowpostOutboxRow(t, 1, knowpostevent.TypeKnowPostDeleted))

	err := processor.Handle(context.Background(), message)

	require.ErrorContains(t, err, "redis unavailable")
}

func TestContentSafetyEpochProcessorHonorsCanceledContext(t *testing.T) {
	store := &recordingSafetyEpochStore{}
	processor := newContentSafetyEpochProcessor(store, newLogThrottle(time.Hour))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := processor.Handle(ctx, relationCanalMessage(t, knowpostOutboxRow(t, 1, knowpostevent.TypeKnowPostDeleted)))

	require.ErrorIs(t, err, context.Canceled)
	require.Equal(t, 0, store.Calls())
}

func TestContentSafetyEpochConsumerUsesIndependentGroup(t *testing.T) {
	cfg := contentSafetyEpochConsumerConfig(config.Config{Kafka: config.KafkaConf{
		Brokers:            []string{"kafka:29092"},
		CanalOutboxTopic:   "canal-outbox",
		SafetyEpochGroupId: "knowpost-feed-content-safety-epoch",
	}})

	require.Equal(t, []string{"kafka:29092"}, cfg.Brokers)
	require.Equal(t, "canal-outbox", cfg.Topic)
	require.Equal(t, "knowpost-feed-content-safety-epoch", cfg.GroupId)
	require.Equal(t, "latest", cfg.StartOffset)
}

type recordingSafetyEpochStore struct {
	mu    sync.Mutex
	calls int
	err   error
}

func (s *recordingSafetyEpochStore) BumpSafety(context.Context) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return 0, s.err
	}
	s.calls++
	return uint64(s.calls), nil
}

func (s *recordingSafetyEpochStore) Calls() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.calls
}

func knowpostOutboxRow(t *testing.T, id int64, eventType string) map[string]string {
	t.Helper()
	payload := knowpostevent.KnowPostEvent{Type: eventType, PostId: 9, Author: 7}
	raw, err := json.Marshal(payload)
	require.NoError(t, err)
	return map[string]string{
		"id":             jsonNumber(id),
		"aggregate_type": knowpostevent.AggregateType,
		"aggregate_id":   "9",
		"type":           eventType,
		"payload":        string(raw),
	}
}
