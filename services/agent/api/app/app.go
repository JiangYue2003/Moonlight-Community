package app

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
)

type Config = config.Config

func Run(ctx context.Context, cfg Config) error {
	server := rest.MustNewServer(cfg.RestConf)
	defer server.Stop()

	svcCtx := svc.NewServiceContext(cfg)
	defer svcCtx.Close()
	handler.RegisterHandlers(server, svcCtx)

	done := make(chan struct{})
	go func() {
		defer close(done)
		server.Start()
	}()

	logx.Infof("agent-api listening at %s:%d", cfg.Host, cfg.Port)

	select {
	case <-ctx.Done():
		server.Stop()
		<-done
		return nil
	case <-done:
		return errors.New("agent-api server exited unexpectedly")
	}
}
