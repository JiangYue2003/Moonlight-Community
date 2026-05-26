package sfx

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoCtx_DeduplicatesConcurrentCalls(t *testing.T) {
	g := New[int]()
	var calls atomic.Int32

	const N = 100
	loader := func(ctx context.Context) (int, error) {
		calls.Add(1)
		time.Sleep(50 * time.Millisecond) // 拖长执行时间确保并发命中
		return 42, nil
	}

	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			v, err := g.DoCtx(context.Background(), "k", loader)
			if err != nil {
				t.Errorf("err: %v", err)
				return
			}
			if v != 42 {
				t.Errorf("got %d, want 42", v)
			}
		}()
	}
	wg.Wait()
	if got := calls.Load(); got != 1 {
		t.Fatalf("loader should be called exactly once, got %d", got)
	}
}

func TestDoCtx_PropagatesError(t *testing.T) {
	g := New[string]()
	want := errors.New("boom")
	v, err := g.DoCtx(context.Background(), "k", func(ctx context.Context) (string, error) {
		return "", want
	})
	if !errors.Is(err, want) {
		t.Fatalf("error not propagated: got %v", err)
	}
	if v != "" {
		t.Fatalf("on error v should be zero value, got %q", v)
	}
}

func TestDoCtx_DifferentKeysAreIndependent(t *testing.T) {
	g := New[int]()
	a, _ := g.DoCtx(context.Background(), "a", func(ctx context.Context) (int, error) { return 1, nil })
	b, _ := g.DoCtx(context.Background(), "b", func(ctx context.Context) (int, error) { return 2, nil })
	if a != 1 || b != 2 {
		t.Fatalf("got a=%d b=%d, want 1,2", a, b)
	}
}

func TestDoCtx_RespectsContextCancel(t *testing.T) {
	g := New[int]()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := g.DoCtx(ctx, "k", func(ctx context.Context) (int, error) {
		return 0, ctx.Err()
	})
	if err == nil {
		t.Fatal("cancelled ctx should produce error")
	}
}
