package app

import (
	"context"
	"errors"
	"testing"
	"time"
)

type testComponent struct {
	name string
	run  func(ctx context.Context) error
}

func (t testComponent) Name() string { return t.name }
func (t testComponent) Run(ctx context.Context) error {
	return t.run(ctx)
}

func TestRun_ReturnsComponentErrorAndCancelsOthers(t *testing.T) {
	canceled := make(chan struct{})
	errBoom := errors.New("boom")

	components := []Component{
		testComponent{name: "ok", run: func(ctx context.Context) error {
			<-ctx.Done()
			close(canceled)
			return nil
		}},
		testComponent{name: "bad", run: func(ctx context.Context) error {
			return errBoom
		}},
	}

	err := Run(context.Background(), components)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, errBoom) {
		t.Fatalf("expected boom, got %v", err)
	}

	select {
	case <-canceled:
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("other components were not canceled")
	}
}

func TestRun_ReturnsNilOnParentCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	components := []Component{
		testComponent{name: "wait", run: func(ctx context.Context) error {
			<-ctx.Done()
			return nil
		}},
	}

	if err := Run(ctx, components); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}
