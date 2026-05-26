package app

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/services/search/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	server := rest.MustNewServer(cfg.RestConf)
	defer server.Stop()

	svcCtx := svc.NewServiceContext(cfg)
	handler.RegisterHandlers(server, svcCtx)

	done := make(chan struct{})
	go func() {
		defer close(done)
		server.Start()
	}()

	logx.Infof("search-api listening at %s:%d", cfg.Host, cfg.Port)

	select {
	case <-ctx.Done():
		server.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("search-api server exited unexpectedly")
	}
}
