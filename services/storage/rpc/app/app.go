package app

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/config"
	storageServer "github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/server/storage"
	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/svc"
	storagepb "github.com/zhiguang/zhiguang-go/services/storage/rpc/storage"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	svcCtx := svc.NewServiceContext(cfg)
	s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
		storagepb.RegisterStorageServer(grpcServer, storageServer.NewStorageServer(svcCtx))
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

	logx.Infof("storage-rpc listening at %s", cfg.ListenOn)

	select {
	case <-ctx.Done():
		s.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("storage-rpc server exited unexpectedly")
	}
}
