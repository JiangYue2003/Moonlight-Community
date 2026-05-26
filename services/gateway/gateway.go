package main

import (
	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zhiguang/zhiguang-go/services/gateway/app"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/config"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := app.Run(context.Background(), c); err != nil {
		panic(err)
	}
}
