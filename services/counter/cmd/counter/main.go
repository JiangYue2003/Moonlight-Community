package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/counter/internal/bootstrap"
)

var configFile = flag.String("f", "services/counter/cmd/counter/etc/counter.yaml", "the config file")

func main() {
	flag.Parse()

	var c bootstrap.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := bootstrap.Run(ctx, c); err != nil {
		logx.Errorf("counter service exit: %v", err)
		os.Exit(1)
	}
}
