package flusher

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/lockx"
)

// 验证两个并发 worker 用同一个 redsync 锁时，同一时刻只有一个进入 critical section。
func TestRedsyncLock_OnlyOneEntersAtATime(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	mx := lockx.New(rdb)
	const N = 8
	const Iter = 10

	var inCritical atomic.Int32
	var maxConcurrent atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < Iter; j++ {
				release, ok, err := mx.TryAcquire(context.Background(),
					"aggr:flush:lock", 200*time.Millisecond)
				if err != nil {
					t.Errorf("acquire: %v", err)
					return
				}
				if !ok {
					continue
				}
				cur := inCritical.Add(1)
				for {
					m := maxConcurrent.Load()
					if cur > m {
						if maxConcurrent.CompareAndSwap(m, cur) {
							break
						}
						continue
					}
					break
				}
				time.Sleep(5 * time.Millisecond)
				inCritical.Add(-1)
				_ = release()
			}
		}()
	}
	wg.Wait()
	if got := maxConcurrent.Load(); got > 1 {
		t.Fatalf("more than one goroutine in critical section: %d", got)
	}
}
