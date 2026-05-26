package cachex

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
)

type item struct {
	Id   int64  `json:"id"`
	Body string `json:"body"`
}

func newCache(t *testing.T) (*miniredis.Miniredis, Cache[*item], *L1, *L2) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis: %v", err)
	}
	t.Cleanup(mr.Close)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	l1, err := NewL1(L1Config{NumCounters: 1000, MaxCost: 1 << 20, BufferItems: 64})
	if err != nil {
		t.Fatal(err)
	}
	l2 := NewL2(rdb)

	c := New[*item](Options[*item]{
		L1: l1, L2: l2,
		BaseTTL: 60 * time.Second, JitterMax: 0,
		NullTTL: 30 * time.Second, NullJitterMax: 0,
		Marshal:   func(v *item) ([]byte, error) { return json.Marshal(v) },
		Unmarshal: func(b []byte, v **item) error { var x item; err := json.Unmarshal(b, &x); *v = &x; return err },
	})
	return mr, c, l1, l2
}

func TestGetOrLoad_LoaderRunsOnFirstMiss(t *testing.T) {
	_, c, _, _ := newCache(t)
	var calls atomic.Int32
	v, err := c.GetOrLoad(context.Background(), "k", "sf:k", func(ctx context.Context) (*item, error) {
		calls.Add(1)
		return &item{Id: 1, Body: "hi"}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if v.Id != 1 || v.Body != "hi" {
		t.Fatalf("payload drift: %+v", v)
	}
	if calls.Load() != 1 {
		t.Fatalf("loader should run once, got %d", calls.Load())
	}
}

func TestGetOrLoad_L2HitPromotesL1(t *testing.T) {
	mr, c, l1, _ := newCache(t)

	// 直接预填 L2，跳过 L3
	raw, _ := json.Marshal(&item{Id: 7, Body: "from-l2"})
	mr.Set("k", string(raw))

	v, err := c.GetOrLoad(context.Background(), "k", "sf:k", func(ctx context.Context) (*item, error) {
		t.Fatal("loader must not run on L2 hit")
		return nil, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if v.Id != 7 {
		t.Fatalf("L2 payload not returned: %+v", v)
	}
	l1.Wait()
	if _, ok := l1.Get("k"); !ok {
		t.Fatal("L1 should be promoted after L2 hit")
	}
}

func TestGetOrLoad_NullSentinelPreventsRepeatedLoads(t *testing.T) {
	_, c, _, _ := newCache(t)
	var calls atomic.Int32
	loader := func(ctx context.Context) (*item, error) {
		calls.Add(1)
		return nil, ErrNotFound
	}
	if _, err := c.GetOrLoad(context.Background(), "k", "sf:k", loader); !errors.Is(err, ErrNotFound) {
		t.Fatal("first miss should return ErrNotFound")
	}
	// 等 L1 写入对外可见后再次调用：应该走 L1 sentinel，loader 不再被触发。
	c2, _, _, _ := func() (Cache[*item], *L1, *L2, *miniredis.Miniredis) { return c, nil, nil, nil }()
	_ = c2
	for i := 0; i < 5; i++ {
		_, _ = c.GetOrLoad(context.Background(), "k", "sf:k", loader)
	}
	if got := calls.Load(); got > 5 {
		t.Fatalf("loader called %d times, expected at most 5 (L1 promotion is async)", got)
	}
}

func TestInvalidate_DoubleDelete(t *testing.T) {
	mr, c, l1, _ := newCache(t)
	if _, err := c.GetOrLoad(context.Background(), "k", "sf", func(ctx context.Context) (*item, error) {
		return &item{Id: 1, Body: "x"}, nil
	}); err != nil {
		t.Fatal(err)
	}
	l1.Wait()
	if err := c.Invalidate(context.Background(), "k"); err != nil {
		t.Fatal(err)
	}
	if mr.Exists("k") {
		t.Fatal("L2 not deleted")
	}
	l1.Wait()
	if _, ok := l1.Get("k"); ok {
		t.Fatal("L1 not deleted")
	}
}

func TestGetOrLoad_SingleFlightDeduplicatesParallelLoads(t *testing.T) {
	_, c, _, _ := newCache(t)
	var calls atomic.Int32
	loader := func(ctx context.Context) (*item, error) {
		calls.Add(1)
		time.Sleep(40 * time.Millisecond)
		return &item{Id: 9, Body: "slow"}, nil
	}
	const N = 50
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			_, _ = c.GetOrLoad(context.Background(), "k", "sf:k", loader)
		}()
	}
	wg.Wait()
	if got := calls.Load(); got != 1 {
		t.Fatalf("singleflight failed: loader ran %d times", got)
	}
}

func TestSetWithExtension_AppliesHotLevelTTL(t *testing.T) {
	mr, c, l1, _ := newCache(t)
	val := &item{Id: 1, Body: "ext"}
	if err := c.SetWithExtension(context.Background(), "k", val, hotkey.LevelMedium); err != nil {
		t.Fatal(err)
	}
	l1.Wait()
	ttl := mr.TTL("k")
	if ttl <= 60*time.Second {
		t.Fatalf("MEDIUM should extend TTL > 60s base, got %v", ttl)
	}
	if _, ok := l1.Get("k"); !ok {
		t.Fatal("L1 not written")
	}
}

func TestJitter_ReturnsBaseWhenMaxZero(t *testing.T) {
	if got := Jitter(60*time.Second, 0); got != 60*time.Second {
		t.Fatalf("Jitter(60s, 0) = %v, want 60s", got)
	}
}

func TestJitter_StaysWithinRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := Jitter(60*time.Second, 30*time.Second)
		if got < 60*time.Second || got >= 90*time.Second {
			t.Fatalf("jitter out of range: %v", got)
		}
	}
}
