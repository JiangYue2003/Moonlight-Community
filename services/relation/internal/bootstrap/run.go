package bootstrap

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	application "github.com/zhiguang/zhiguang-go/services/relation/internal/application"
	rpcserver "github.com/zhiguang/zhiguang-go/services/relation/internal/transport/grpc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

func Run(ctx context.Context, cfg Config) error {
	svcCtx := application.NewServiceContext(application.Config{
		Mysql:      cfg.Mysql,
		CacheRedis: cfg.CacheRedis,
		Redis:      cfg.Redis,
		UserRpc:    cfg.UserRpc,
		RateLimit:  cfg.RateLimit,
		Snowflake:  cfg.Snowflake,
	})
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
