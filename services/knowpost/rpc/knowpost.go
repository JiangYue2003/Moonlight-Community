package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/listener"
	knowpostServer "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/server/knowpost"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

var configFile = flag.String("f", "etc/knowpost.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	sc := svc.NewServiceContext(c)

	// 启动 Kafka cache invalidation listener（与 gRPC server 并行运行）
	listenerCtx, cancelListener := context.WithCancel(context.Background())
	defer cancelListener()
	go func() {
		if err := listener.Run(listenerCtx, sc); err != nil {
			logx.Errorf("cache invalidation listener exit: %v", err)
		}
	}()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		knowpost.RegisterKnowPostServer(grpcServer, knowpostServer.NewKnowPostServer(sc))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
