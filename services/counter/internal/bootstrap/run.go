package bootstrap

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/services/counter/internal/application"
	grpcserver "github.com/zhiguang/zhiguang-go/services/counter/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

func Run(ctx context.Context, cfg Config) error {
	svcCtx := application.NewServiceContext(application.Config{
		Redis:   cfg.Redis,
		Kafka:   cfg.Kafka,
		Rebuild: cfg.Rebuild,
	})
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		counter.RegisterCounterServer(grpcServer, grpcserver.NewCounterServer(svcCtx))
		counter.RegisterUserCounterServer(grpcServer, grpcserver.NewUserCounterServer(svcCtx))
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
