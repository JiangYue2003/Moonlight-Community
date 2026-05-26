package main

import (
	"flag"
	"fmt"

	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/config"
	counterServer "github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/server/counter"
	usercounterServer "github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/server/usercounter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/counter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		counter.RegisterCounterServer(grpcServer, counterServer.NewCounterServer(ctx))
		counter.RegisterUserCounterServer(grpcServer, usercounterServer.NewUserCounterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
