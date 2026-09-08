package bootstrap

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	grpcserver "github.com/zhiguang/zhiguang-go/services/user/internal/transport/grpc"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, cfg Config) error {
	svcCtx := NewServiceContext(cfg)
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		userpb.RegisterUserServer(grpcServer, grpcserver.NewUserServer(svcCtx))
		userpb.RegisterAuthServer(grpcServer, grpcserver.NewAuthServer(svcCtx))
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

	logx.Infof("user-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("user-rpc server exited unexpectedly")
	}
}
