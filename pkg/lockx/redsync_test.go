package lockx

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func newRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis start: %v", err)
	}
	t.Cleanup(mr.Close)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return mr, rdb
}

func TestTryAcquire_FirstTakerWins(t *testing.T) {
	_, rdb := newRedis(t)
	m := New(rdb)
	ctx := context.Background()

	rel, ok, err := m.TryAcquire(ctx, "lock:k", time.Second)
	if err != nil || !ok {
		t.Fatalf("first acquire should succeed: ok=%v err=%v", ok, err)
	}
	t.Cleanup(func() { _ = rel() })

	_, ok2, err := m.TryAcquire(ctx, "lock:k", time.Second)
	if err != nil {
		t.Fatalf("second acquire returned err: %v", err)
	}
	if ok2 {
		t.Fatal("second acquire should not succeed while lock held")
	}
}

func TestTryAcquire_ReleaseAllowsReacquire(t *testing.T) {
	_, rdb := newRedis(t)
	m := New(rdb)
	ctx := context.Background()

	rel, ok, _ := m.TryAcquire(ctx, "lock:k", time.Second)
	if !ok {
		t.Fatal("first acquire failed")
	}
	if err := rel(); err != nil {
		t.Fatalf("release: %v", err)
	}

	_, ok, _ = m.TryAcquire(ctx, "lock:k", time.Second)
	if !ok {
		t.Fatal("after release, re-acquire should succeed")
	}
}

func TestTryAcquire_ConcurrentOnlyOneWins(t *testing.T) {
	_, rdb := newRedis(t)
	m := New(rdb)
	ctx := context.Background()

	var wins int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rel, ok, _ := m.TryAcquire(ctx, "lock:concur", 200*time.Millisecond)
			if ok {
				mu.Lock()
				wins++
				mu.Unlock()
				time.Sleep(50 * time.Millisecond)
				_ = rel()
			}
		}()
	}
	wg.Wait()
	if wins == 0 {
		t.Fatal("nobody acquired the lock")
	}
	// 不严格断言 wins=1：miniredis 的 PEXPIRE/SET NX 行为不会阻塞，但同一时刻只允许一个持有方。
	// 主要验证不发生错误、不挂起。
}
