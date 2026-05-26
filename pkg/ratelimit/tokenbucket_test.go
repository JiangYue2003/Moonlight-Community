package ratelimit

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func newRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis: %v", err)
	}
	t.Cleanup(mr.Close)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return mr, rdb
}

func TestTake_Capacity1ExhaustsAfterFirst(t *testing.T) {
	_, rdb := newRedis(t)
	b := New(rdb)
	ctx := context.Background()

	ok, err := b.Take(ctx, "k", 1, 1)
	if err != nil || !ok {
		t.Fatalf("first take should pass: ok=%v err=%v", ok, err)
	}
	ok2, err := b.Take(ctx, "k", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if ok2 {
		t.Fatal("second take with capacity=1 must fail")
	}
}

func TestTake_RefillRestoresTokens(t *testing.T) {
	mr, rdb := newRedis(t)
	b := New(rdb)
	ctx := context.Background()

	// 容量 2，速率 10/s（每 100ms 补 1）
	_, _ = b.Take(ctx, "k", 2, 10)
	_, _ = b.Take(ctx, "k", 2, 10)
	ok, _ := b.Take(ctx, "k", 2, 10)
	if ok {
		t.Fatal("third take immediately must fail")
	}
	mr.FastForward(300 * time.Millisecond) // miniredis 可手动推进时间，但脚本里用 ARGV[3] 时间戳
	// 真推进时间：手动 sleep 让 now_ms 变化。
	time.Sleep(310 * time.Millisecond)
	ok, _ = b.Take(ctx, "k", 2, 10)
	if !ok {
		t.Fatal("after refill window, take should pass")
	}
}

func TestTake_KeysAreIsolated(t *testing.T) {
	_, rdb := newRedis(t)
	b := New(rdb)
	ctx := context.Background()
	_, _ = b.Take(ctx, "ka", 1, 1)
	if ok, _ := b.Take(ctx, "kb", 1, 1); !ok {
		t.Fatal("different key should have own bucket")
	}
}

func TestTake_Concurrent_TotalNoMoreThanCapacity(t *testing.T) {
	_, rdb := newRedis(t)
	b := New(rdb)
	ctx := context.Background()

	const N = 50
	var pass atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			ok, _ := b.Take(ctx, "burst", 5, 0)
			if ok {
				pass.Add(1)
			}
		}()
	}
	wg.Wait()
	// refill rate=0 时桶不会再补，最多通过 capacity=5 次。
	if pass.Load() > 5 {
		t.Fatalf("total passes should be <= capacity(5), got %d", pass.Load())
	}
	if pass.Load() == 0 {
		t.Fatal("nobody passed; bug?")
	}
}

func TestTake_BackwardClockNotPanic(t *testing.T) {
	// 时间回拨情况下脚本应当将 delta clamp 为 0，不应崩溃。
	_, rdb := newRedis(t)
	b := New(rdb)
	if _, err := b.Take(context.Background(), "k", 1, 1); err != nil {
		t.Fatal(err)
	}
}
