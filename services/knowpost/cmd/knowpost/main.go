package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/debughttp"
	"github.com/zhiguang/zhiguang-go/services/knowpost/cmd/knowpost/internal/app"
	"github.com/zhiguang/zhiguang-go/services/knowpost/cmd/knowpost/internal/config"
)

var configFile = flag.String("f", "etc/knowpost.yaml", "the config file")

func main() {
	flag.Parse()

	c, err := loadConfig(*configFile)
	if err != nil {
		logx.Errorf("load knowpost config: %v", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	components := []app.Component{
		app.NewRPCComponent(c.Rpc),
		debughttp.New(c.DebugHTTP),
	}
	if !c.DisableAPI {
		components = append([]app.Component{app.NewAPIComponent(c.Api)}, components...)
	}

	if err := app.Run(ctx, components); err != nil {
		logx.Errorf("knowpost merged service exit: %v", err)
		os.Exit(1)
	}
}

func loadConfig(path string) (config.Config, error) {
	var c config.Config
	err := conf.Load(path, &c, conf.UseEnv())
	return c, err
}
