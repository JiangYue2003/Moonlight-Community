package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zhiguang/zhiguang-go/services/gateway/internal/config"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
)

func Run(ctx context.Context, c config.Config) error {
	sc := srv.NewServiceContext(c)
	engine := handler.NewEngine(sc)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler: engine,
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}
