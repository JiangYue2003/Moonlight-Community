package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/outbox/gc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/outbox/gc/internal/worker"
)

func main() {
	var cfg config.Config
	conf.MustLoad(os.Args[1], &cfg)
	logx.MustSetup(logx.LogConf{ServiceName: cfg.Name})

	db := sqlx.NewMysql(cfg.DataSource)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logx.Infof("outbox-gc started, interval=%dh retain=%dd batch=%d",
		cfg.IntervalHours, cfg.RetainDays, cfg.BatchSize)

	worker.New(cfg, db).Run(ctx)
	logx.Info("outbox-gc stopped")
}
