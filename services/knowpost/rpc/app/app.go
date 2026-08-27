package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	setupRuntimeLogging(cfg)
	sc := svc.NewServiceContext(cfg)
	sc.StartFeedFanoutWorker()
	defer sc.Close()

	runtimeCtx, cancelRuntime := context.WithCancel(ctx)
	defer cancelRuntime()
	runners := []namedListenerRunner{{name: "counter-cache-invalidation", run: func(ctx context.Context) error {
		return listener.Run(ctx, sc)
	}}}
	if cfg.Feed.Epoch.RelationConsumerEnabled {
		runners = append(runners, namedListenerRunner{name: "relation-epoch", run: func(ctx context.Context) error {
			return listener.RunRelationEpoch(ctx, sc)
		}})
	}
	if cfg.Feed.Epoch.SafetyConsumerEnabled {
		runners = append(runners, namedListenerRunner{name: "content-safety-epoch", run: func(ctx context.Context) error {
			return listener.RunContentSafetyEpoch(ctx, sc)
		}})
	}
	listenersDone := make(chan error, 1)
	go func() { listenersDone <- superviseListeners(runtimeCtx, runners) }()

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
		cancelRuntime()
		s.Stop()
		<-done
		<-listenersDone
		return nil
	case err := <-listenersDone:
		cancelRuntime()
		s.Stop()
		<-done
		if ctx.Err() != nil {
			return nil
		}
		return err
	case <-done:
		cancelRuntime()
		<-listenersDone
		return errors.New("knowpost-rpc server exited unexpectedly")
	}
}

// setupRuntimeLogging must run before service construction. Several go-zero
// dependencies log while ServiceContext is being built, and logx setup is
// process-global/one-shot; configuring it later in zrpc cannot undo that first
// initialization. Error-only mode also disables statement info logging so SQL
// formatting is not paid on the benchmark hot path; errors and slow SQL remain.
func setupRuntimeLogging(cfg Config) {
	logx.MustSetup(cfg.Log)
	switch strings.ToLower(strings.TrimSpace(cfg.Log.Level)) {
	case "error", "severe":
		sqlx.DisableStmtLog()
	}
}

type namedListenerRunner struct {
	name string
	run  func(context.Context) error
}

type listenerResult struct {
	name string
	err  error
}

func superviseListeners(ctx context.Context, runners []namedListenerRunner) error {
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan listenerResult, len(runners))
	var wg sync.WaitGroup
	wg.Add(len(runners))
	for _, runner := range runners {
		runner := runner
		go func() {
			defer wg.Done()
			results <- listenerResult{name: runner.name, err: runner.run(runCtx)}
		}()
	}
	select {
	case <-ctx.Done():
		cancel()
		wg.Wait()
		return nil
	case result := <-results:
		cancel()
		wg.Wait()
		if ctx.Err() != nil {
			return nil
		}
		if result.err == nil {
			return fmt.Errorf("%s listener exited unexpectedly", result.name)
		}
		return fmt.Errorf("%s listener exited: %w", result.name, result.err)
	}
}
