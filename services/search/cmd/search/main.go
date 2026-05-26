package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/search/cmd/search/internal/app"
	"github.com/zhiguang/zhiguang-go/services/search/cmd/search/internal/config"
)

var configFile = flag.String("f", "etc/search.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	components := []app.Component{
		app.NewIndexerComponent(c.Indexer),
	}
	if !c.DisableAPI {
		components = append([]app.Component{app.NewAPIComponent(c.Api)}, components...)
	}

	if err := app.Run(ctx, components); err != nil {
		logx.Errorf("search merged service exit: %v", err)
		os.Exit(1)
	}
}
