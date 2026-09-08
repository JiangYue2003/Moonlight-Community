package bootstrap

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/search/internal/application"
	searchServer "github.com/zhiguang/zhiguang-go/services/search/internal/transport/grpc"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
)

func Run(ctx context.Context, cfg Config) error {
	svcCtx := application.NewServiceContext(cfg)
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		searchpb.RegisterSearchServer(grpcServer, searchServer.NewSearchServer(svcCtx))
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

	logx.Infof("search-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("search-rpc server exited unexpectedly")
	}
}
