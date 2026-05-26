package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
)

var configFile = flag.String("f", "etc/agent-api.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	defer ctx.Close()
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("agent-api listening at %s:%d\n", c.Host, c.Port)
	server.Start()
}
