package app

import (
	"context"
	"errors"

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

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	sc := svc.NewServiceContext(cfg)

	go func() {
		if err := listener.Run(ctx, sc); err != nil {
			logx.Errorf("cache invalidation listener exit: %v", err)
		}
	}()

	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		knowpost.RegisterKnowPostServer(grpcServer, knowpostServer.NewKnowPostServer(sc))
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

	logx.Infof("knowpost-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("knowpost-rpc server exited unexpectedly")
	}
}
