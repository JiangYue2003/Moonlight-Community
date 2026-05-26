package app

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/processor"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/svc"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	logx.MustSetup(cfg.Log)

	sc := svc.NewServiceContext(cfg)
	defer sc.Close()

	proc := processor.New(sc)
	logx.Info("relation-syncer started")

	err := kafkax.RunConsumer(ctx, kafkax.ConsumerConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topic,
		GroupId: cfg.Kafka.GroupId,
	}, func(ctx context.Context, m kafka.Message) error {
		return proc.Handle(ctx, m.Value)
	})
	if err != nil {
		return fmt.Errorf("consume kafka: %w", err)
	}
	return nil
}
