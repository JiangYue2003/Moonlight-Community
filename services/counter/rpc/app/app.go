package app

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/config"
	counterServer "github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/server/counter"
	usercounterServer "github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/server/usercounter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	svcCtx := svc.NewServiceContext(cfg)
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		counter.RegisterCounterServer(grpcServer, counterServer.NewCounterServer(svcCtx))
		counter.RegisterUserCounterServer(grpcServer, usercounterServer.NewUserCounterServer(svcCtx))
		if cfg.Mode == service.DevMode || cfg.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		s.Start()
	}()

	logx.Infof("counter-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("counter-rpc server exited unexpectedly")
	}
}
