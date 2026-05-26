package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/agent/cmd/agent/internal/app"
	"github.com/zhiguang/zhiguang-go/services/agent/cmd/agent/internal/config"
)

var configFile = flag.String("f", "etc/agent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	components := []app.Component{
		app.NewAPIComponent(c.Api),
		app.NewIndexerComponent(c.Indexer),
	}

	if err := app.Run(ctx, components); err != nil {
		logx.Errorf("agent merged service exit: %v", err)
		os.Exit(1)
	}
}
