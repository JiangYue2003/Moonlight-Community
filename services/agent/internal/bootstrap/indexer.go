package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/agent/internal/transport/indexer/config"
	"github.com/zhiguang/zhiguang-go/services/agent/internal/transport/indexer/processor"
	"github.com/zhiguang/zhiguang-go/services/agent/internal/transport/indexer/svc"
)

type IndexerConfig = config.Config

func RunIndexer(ctx context.Context, cfg IndexerConfig) error {
	logx.MustSetup(cfg.Log)

	sc := svc.NewServiceContext(cfg)
	defer sc.Close()

	proc := processor.New(sc)
	go processor.RunCompensation(ctx, proc, time.Duration(cfg.CompensateMinutes)*time.Minute)

	logx.Info("agent-indexer started")
	err := kafkax.RunConsumer(ctx, kafkax.ConsumerConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topic,
		GroupId: cfg.Kafka.GroupId,
	}, func(ctx context.Context, m kafka.Message) error {
		return proc.Handle(ctx, m)
	})
	if err != nil {
		return fmt.Errorf("consume kafka: %w", err)
	}
	return nil
}
