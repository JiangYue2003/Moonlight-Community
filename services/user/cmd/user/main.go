package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/user/internal/bootstrap"
)

var configFile = flag.String("f", "services/user/cmd/user/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c bootstrap.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := bootstrap.Run(ctx, c); err != nil {
		logx.Errorf("user service exit: %v", err)
		os.Exit(1)
	}
}
