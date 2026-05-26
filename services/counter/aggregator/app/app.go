package app

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/config"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/flusher"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/svc"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	logx.MustSetup(cfg.Log)

	sc := svc.NewServiceContext(cfg)
	defer sc.Close()

	go flusher.RunFlusher(ctx, sc)
	if err := flusher.RunConsumer(ctx, sc); err != nil {
		return fmt.Errorf("run consumer: %w", err)
	}
	return nil
}
