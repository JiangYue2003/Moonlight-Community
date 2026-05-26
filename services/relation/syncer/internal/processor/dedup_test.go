package processor

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/services/relation/shared/event"
)

func newDedup(t *testing.T) *Dedup {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return NewDedup(rdb)
}

func TestDedup_FirstAcquireWinsSubsequentSkip(t *testing.T) {
	d := newDedup(t)
	ctx := context.Background()
	first, _ := d.Acquire(ctx, "k", time.Second)
	second, _ := d.Acquire(ctx, "k", time.Second)
	if !first || second {
		t.Fatalf("first should win, second skip; got %v/%v", first, second)
	}
}

func TestDedup_DifferentKeysIndependent(t *testing.T) {
	d := newDedup(t)
	ctx := context.Background()
	_, _ = d.Acquire(ctx, "ka", time.Second)
	if ok, _ := d.Acquire(ctx, "kb", time.Second); !ok {
		t.Fatal("different key should be independent")
	}
}

// 模拟"同一 outbox.id 在 retry 场景下被多次喂进来"
func TestDedup_RetryLoopOnlyProcessedOnce(t *testing.T) {
	d := newDedup(t)
	ctx := context.Background()
	const N = 50
	wins := 0
	for i := 0; i < N; i++ {
		ok, _ := d.Acquire(ctx, "rel:FollowCreated:42", time.Second)
		if ok {
			wins++
		}
	}
	if wins != 1 {
		t.Fatalf("retry loop should produce exactly 1 win, got %d", wins)
	}
}

// 防 unused：event 包至少 import 一次让 IDE 友好
var _ = event.TypeFollowCreated
