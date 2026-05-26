package app

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/config"
	rpcserver "github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/server/relation"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	svcCtx := svc.NewServiceContext(cfg)
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		relation.RegisterRelationServer(grpcServer, rpcserver.NewRelationServer(svcCtx))
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

	logx.Infof("relation-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("relation-rpc server exited unexpectedly")
	}
}
