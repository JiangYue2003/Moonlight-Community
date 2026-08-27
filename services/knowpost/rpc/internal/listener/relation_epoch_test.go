package listener

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	relationevent "github.com/zhiguang/zhiguang-go/services/relation/shared/event"
)

func TestRelationEpochProcessorBumpsCreatedCanceledAndEveryOutboxRow(t *testing.T) {
	store := &recordingEpochStore{}
	processor := newRelationEpochProcessor(store, newLogThrottle(time.Hour))
	message := relationCanalMessage(t,
		relationOutboxRow(t, 1, "following", relationevent.RelationEvent{Type: relationevent.TypeFollowCreated, FromUserId: 11, ToUserId: 21}),
		relationOutboxRow(t, 2, "following", relationevent.RelationEvent{Type: relationevent.TypeFollowCanceled, FromUserId: 12, ToUserId: 22}),
	)

	err := processor.Handle(context.Background(), message)

	require.NoError(t, err)
	require.Equal(t, []int64{11, 12}, store.UserIDs())
}

func TestRelationEpochProcessorAllowsDuplicateAndReverseOrderExtraBumps(t *testing.T) {
	store := &recordingEpochStore{}
	processor := newRelationEpochProcessor(store, newLogThrottle(time.Hour))
	created := relationCanalMessage(t, relationOutboxRow(t, 1, "following", relationevent.RelationEvent{
		Type: relationevent.TypeFollowCreated, FromUserId: 42, ToUserId: 7,
	}))
	canceled := relationCanalMessage(t, relationOutboxRow(t, 2, "following", relationevent.RelationEvent{
		Type: relationevent.TypeFollowCanceled, FromUserId: 42, ToUserId: 7,
	}))

	require.NoError(t, processor.Handle(context.Background(), canceled))
	require.NoError(t, processor.Handle(context.Background(), created))
	require.NoError(t, processor.Handle(context.Background(), created))
	require.Equal(t, []int64{42, 42, 42}, store.UserIDs())
}

func TestRelationEpochProcessorSkipsIrrelevantAndPoisonMessages(t *testing.T) {
	store := &recordingEpochStore{}
	processor := newRelationEpochProcessor(store, newLogThrottle(time.Hour))

	for _, message := range [][]byte{
		[]byte("not-json"),
		relationCanalMessage(t, map[string]string{
			"id": "1", "aggregate_type": "following", "aggregate_id": "1", "type": "FollowCreated", "payload": "{bad",
		}),
		relationCanalMessage(t, relationOutboxRow(t, 2, "knowpost", relationevent.RelationEvent{
			Type: relationevent.TypeFollowCreated, FromUserId: 9,
		})),
		relationCanalMessage(t, relationOutboxRow(t, 3, "following", relationevent.RelationEvent{
			Type: "Unknown", FromUserId: 9,
		})),
		relationCanalMessage(t, relationOutboxRow(t, 4, "following", relationevent.RelationEvent{
			Type: relationevent.TypeFollowCreated, FromUserId: 0,
		})),
	} {
		require.NoError(t, processor.Handle(context.Background(), message))
	}
	require.Empty(t, store.UserIDs())
}

func TestRelationEpochProcessorReturnsStoreErrorForKafkaRetry(t *testing.T) {
	store := &recordingEpochStore{err: errors.New("redis unavailable")}
	processor := newRelationEpochProcessor(store, newLogThrottle(time.Hour))
	message := relationCanalMessage(t, relationOutboxRow(t, 1, "following", relationevent.RelationEvent{
		Type: relationevent.TypeFollowCreated, FromUserId: 42,
	}))

	err := processor.Handle(context.Background(), message)

	require.ErrorContains(t, err, "redis unavailable")
}

func TestRelationEpochProcessorHonorsCanceledContext(t *testing.T) {
	store := &recordingEpochStore{}
	processor := newRelationEpochProcessor(store, newLogThrottle(time.Hour))
	message := relationCanalMessage(t, relationOutboxRow(t, 1, "following", relationevent.RelationEvent{
		Type: relationevent.TypeFollowCreated, FromUserId: 42,
	}))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := processor.Handle(ctx, message)

	require.ErrorIs(t, err, context.Canceled)
	require.Empty(t, store.UserIDs())
}

func TestRelationEpochConsumerUsesIndependentGroup(t *testing.T) {
	cfg := relationEpochConsumerConfig(config.Config{Kafka: config.KafkaConf{
		Brokers:              []string{"kafka:29092"},
		CanalOutboxTopic:     "canal-outbox",
		RelationEpochGroupId: "knowpost-feed-relation-epoch",
	}})

	require.Equal(t, []string{"kafka:29092"}, cfg.Brokers)
	require.Equal(t, "canal-outbox", cfg.Topic)
	require.Equal(t, "knowpost-feed-relation-epoch", cfg.GroupId)
	require.Equal(t, "latest", cfg.StartOffset)
}

type recordingEpochStore struct {
	mu      sync.Mutex
	userIDs []int64
	err     error
}

func (s *recordingEpochStore) BumpRelation(_ context.Context, userID int64) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return 0, s.err
	}
	s.userIDs = append(s.userIDs, userID)
	return uint64(len(s.userIDs)), nil
}

func (s *recordingEpochStore) UserIDs() []int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]int64(nil), s.userIDs...)
}

func relationOutboxRow(t *testing.T, id int64, aggregateType string, event relationevent.RelationEvent) map[string]string {
	t.Helper()
	payload, err := json.Marshal(event)
	require.NoError(t, err)
	return map[string]string{
		"id":             jsonNumber(id),
		"aggregate_type": aggregateType,
		"aggregate_id":   jsonNumber(event.FromUserId),
		"type":           event.Type,
		"payload":        string(payload),
	}
}

func relationCanalMessage(t *testing.T, rows ...map[string]string) []byte {
	t.Helper()
	body, err := json.Marshal(map[string]any{
		"database": "zhiguang", "table": "outbox", "type": "INSERT", "isDdl": false, "data": rows,
	})
	require.NoError(t, err)
	return body
}

func jsonNumber(value int64) string {
	return strconv.FormatInt(value, 10)
}
