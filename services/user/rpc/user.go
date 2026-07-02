package main

import (
	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/app"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/config"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := app.Run(context.Background(), c); err != nil {
		panic(err)
	}
}
