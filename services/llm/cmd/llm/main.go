package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/llm/cmd/llm/internal/app"
	"github.com/zhiguang/zhiguang-go/services/llm/cmd/llm/internal/config"
)

var configFile = flag.String("f", "etc/llm.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	components := []app.Component{
		app.NewRagIndexerComponent(c.RagIndexer),
	}
	if !c.DisableAPI {
		components = append([]app.Component{app.NewAPIComponent(c.Api)}, components...)
	}

	if err := app.Run(ctx, components); err != nil {
		logx.Errorf("llm merged service exit: %v", err)
		os.Exit(1)
	}
}
