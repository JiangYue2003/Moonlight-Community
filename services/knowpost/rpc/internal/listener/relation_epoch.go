package listener

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	relationevent "github.com/zhiguang/zhiguang-go/services/relation/shared/event"
)

const badRelationEventLogInterval = 30 * time.Second

type relationEpochStore interface {
	BumpRelation(ctx context.Context, userID int64) (uint64, error)
}

type relationEpochProcessor struct {
	store relationEpochStore
	logs  *logThrottle
}

func newRelationEpochProcessor(store relationEpochStore, logs *logThrottle) *relationEpochProcessor {
	if logs == nil {
		logs = newLogThrottle(badRelationEventLogInterval)
	}
	return &relationEpochProcessor{store: store, logs: logs}
}

// Handle processes every following outbox row. Duplicate and out-of-order
// deliveries intentionally produce additional monotonic bumps; no dedup state
// is needed for correctness.
func (p *relationEpochProcessor) Handle(ctx context.Context, value []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	flat, err := canalx.ParseFlat(value)
	if err != nil {
		p.logs.Errorf(ctx, "bad-canal", "skip invalid relation epoch Canal message: %v", err)
		return nil
	}
	for _, row := range canalx.ExtractOutboxRows(flat) {
		if err := ctx.Err(); err != nil {
			return err
		}
		if row.AggregateType != "following" {
			continue
		}
		var event relationevent.RelationEvent
		if err := json.Unmarshal([]byte(row.Payload), &event); err != nil {
			p.logs.Errorf(ctx, "bad-payload", "skip invalid relation epoch payload: %v", err)
			continue
		}
		if event.Type != relationevent.TypeFollowCreated && event.Type != relationevent.TypeFollowCanceled {
			p.logs.Errorf(ctx, "bad-type", "skip unsupported relation epoch event type")
			continue
		}
		if event.FromUserId <= 0 {
			p.logs.Errorf(ctx, "bad-user", "skip relation epoch event with invalid from-user")
			continue
		}
		if _, err := p.store.BumpRelation(ctx, event.FromUserId); err != nil {
			return fmt.Errorf("bump relation epoch: %w", err)
		}
	}
	return nil
}

func relationEpochConsumerConfig(c config.Config) kafkax.ConsumerConfig {
	return kafkax.ConsumerConfig{
		Brokers:     c.Kafka.Brokers,
		Topic:       c.Kafka.CanalOutboxTopic,
		GroupId:     c.Kafka.RelationEpochGroupId,
		StartOffset: "latest",
	}
}

// RunRelationEpoch blocks until ctx is canceled or the Kafka consumer exits.
// A non-nil handler error keeps the offset uncommitted so kafkax retries it.
func RunRelationEpoch(ctx context.Context, sc *svc.ServiceContext) error {
	if sc == nil || sc.FeedEpochs == nil {
		return errors.New("relation epoch listener requires FeedEpochStore")
	}
	processor := newRelationEpochProcessor(sc.FeedEpochs, newLogThrottle(badRelationEventLogInterval))
	return kafkax.RunConsumer(ctx, relationEpochConsumerConfig(sc.Config), func(ctx context.Context, message kafka.Message) error {
		return processor.Handle(ctx, message.Value)
	})
}

type logThrottle struct {
	mu       sync.Mutex
	interval time.Duration
	last     map[string]time.Time
}

func newLogThrottle(interval time.Duration) *logThrottle {
	if interval <= 0 {
		interval = badRelationEventLogInterval
	}
	return &logThrottle{interval: interval, last: make(map[string]time.Time)}
}

func (l *logThrottle) Errorf(ctx context.Context, key, format string, args ...any) {
	now := time.Now()
	l.mu.Lock()
	last := l.last[key]
	if !last.IsZero() && now.Sub(last) < l.interval {
		l.mu.Unlock()
		return
	}
	l.last[key] = now
	l.mu.Unlock()
	logx.WithContext(ctx).Errorf(format, args...)
}
