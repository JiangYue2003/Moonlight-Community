package app

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestSetupRuntimeLoggingAppliesLevelBeforeServiceContext(t *testing.T) {
	cfg := Config{}
	cfg.Log.Level = "error"
	cfg.Log.Stat = false

	setupRuntimeLogging(cfg)
	var output bytes.Buffer
	logx.SetWriter(logx.NewWriter(&output))
	logx.Info("runtime-info-must-be-suppressed")
	logx.Error("runtime-error-must-remain")

	require.NotContains(t, output.String(), "runtime-info-must-be-suppressed")
	require.Contains(t, output.String(), "runtime-error-must-remain")
}

func TestSuperviseListenersPropagatesExitAndCancelsPeers(t *testing.T) {
	peerCanceled := make(chan struct{})
	runners := []namedListenerRunner{
		{name: "counter", run: func(context.Context) error { return errors.New("counter stopped") }},
		{name: "relation-epoch", run: func(ctx context.Context) error {
			<-ctx.Done()
			close(peerCanceled)
			return nil
		}},
	}

	err := superviseListeners(context.Background(), runners)

	require.ErrorContains(t, err, "counter")
	select {
	case <-peerCanceled:
	case <-time.After(time.Second):
		t.Fatal("peer listener was not canceled")
	}
}

func TestSuperviseListenersReturnsCleanlyOnContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := superviseListeners(ctx, []namedListenerRunner{{name: "listener", run: func(ctx context.Context) error {
		<-ctx.Done()
		return nil
	}}})
	require.NoError(t, err)
}
