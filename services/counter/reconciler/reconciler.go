package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/counter/reconciler/internal/config"
	"github.com/zhiguang/zhiguang-go/services/counter/reconciler/internal/worker"
)

func main() {
	var cfg config.Config
	conf.MustLoad(os.Args[1], &cfg)
	logx.MustSetup(logx.LogConf{ServiceName: cfg.Name})

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{cfg.Redis.Host},
		Password: cfg.Redis.Pass,
	})
	defer rdb.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logx.Infof("counter-reconciler started, interval=%dh threshold_abs=%d threshold_pct=%.1f%%",
		cfg.Scan.IntervalHours, cfg.Scan.ThresholdAbsolute, cfg.Scan.ThresholdPercent)

	worker.New(cfg, rdb).Run(ctx)
	logx.Info("counter-reconciler stopped")
}
