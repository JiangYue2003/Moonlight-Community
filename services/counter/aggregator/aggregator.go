package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/config"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/flusher"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/svc"
)

var configFile = flag.String("f", "etc/aggregator.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	sc := svc.NewServiceContext(c)
	defer sc.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go flusher.RunFlusher(ctx, sc)
	if err := flusher.RunConsumer(ctx, sc); err != nil {
		logx.Errorf("counter-aggregator consumer exit: %v", err)
		os.Exit(1)
	}
}
