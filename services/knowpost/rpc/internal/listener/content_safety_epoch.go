package listener

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	kafka "github.com/segmentio/kafka-go"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	knowpostevent "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
)

type contentSafetyEpochStore interface {
	BumpSafety(ctx context.Context) (uint64, error)
}

type contentSafetyEpochProcessor struct {
	store contentSafetyEpochStore
	logs  *logThrottle
}

func newContentSafetyEpochProcessor(store contentSafetyEpochStore, logs *logThrottle) *contentSafetyEpochProcessor {
	if logs == nil {
		logs = newLogThrottle(badRelationEventLogInterval)
	}
	return &contentSafetyEpochProcessor{store: store, logs: logs}
}

// Handle provides durable compensation for the synchronous safety bump in the
// mutation path. At-least-once delivery may increment the epoch again; that is
// intentional and only causes an extra cache invalidation.
func (p *contentSafetyEpochProcessor) Handle(ctx context.Context, value []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	flat, err := canalx.ParseFlat(value)
	if err != nil {
		p.logs.Errorf(ctx, "bad-safety-canal", "skip invalid content safety Canal message: %v", err)
		return nil
	}
	for _, row := range canalx.ExtractOutboxRows(flat) {
		if err := ctx.Err(); err != nil {
			return err
		}
		if row.AggregateType != knowpostevent.AggregateType {
			continue
		}
		var event knowpostevent.KnowPostEvent
		if err := json.Unmarshal([]byte(row.Payload), &event); err != nil {
			p.logs.Errorf(ctx, "bad-safety-payload", "skip invalid content safety payload: %v", err)
			continue
		}
		switch event.Type {
		case knowpostevent.TypeKnowPostUpdated, knowpostevent.TypeKnowPostDeleted:
			if _, err := p.store.BumpSafety(ctx); err != nil {
				return fmt.Errorf("bump content safety epoch from outbox: %w", err)
			}
		case knowpostevent.TypeKnowPostPublished:
			// New content relies on the short Fresh TTL and must not globally
			// invalidate every personal Feed page.
		default:
			p.logs.Errorf(ctx, "bad-safety-type", "skip unsupported content safety event type")
		}
	}
	return nil
}

func contentSafetyEpochConsumerConfig(c config.Config) kafkax.ConsumerConfig {
	return kafkax.ConsumerConfig{
		Brokers:     c.Kafka.Brokers,
		Topic:       c.Kafka.CanalOutboxTopic,
		GroupId:     c.Kafka.SafetyEpochGroupId,
		StartOffset: "latest",
	}
}

func RunContentSafetyEpoch(ctx context.Context, sc *svc.ServiceContext) error {
	if sc == nil || sc.FeedEpochs == nil {
		return errors.New("content safety epoch listener requires FeedEpochStore")
	}
	processor := newContentSafetyEpochProcessor(sc.FeedEpochs, newLogThrottle(badRelationEventLogInterval))
	return kafkax.RunConsumer(ctx, contentSafetyEpochConsumerConfig(sc.Config), func(ctx context.Context, message kafka.Message) error {
		return processor.Handle(ctx, message.Value)
	})
}
