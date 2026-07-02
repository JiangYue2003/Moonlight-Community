package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/user/cmd/user/internal/app"
	"github.com/zhiguang/zhiguang-go/services/user/cmd/user/internal/config"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	components := []app.Component{
		app.NewUserComponent(c.User),
		app.NewStorageComponent(c.Storage),
	}

	if err := app.Run(ctx, components); err != nil {
		logx.Errorf("user merged service exit: %v", err)
		os.Exit(1)
	}
}
