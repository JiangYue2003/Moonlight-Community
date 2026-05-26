package snowflakex

import (
	"sync"
	"testing"
)

func TestEpochConstant(t *testing.T) {
	// 与原 Java 实现严格一致：2024-01-01T00:00:00Z 的毫秒数。
	if Epoch != 1704067200000 {
		t.Fatalf("EPOCH mismatch: %d", Epoch)
	}
}

func TestNextIdMonotonic(t *testing.T) {
	g := MustNew(1, 1)
	var prev int64
	for i := 0; i < 10000; i++ {
		id, err := g.NextId()
		if err != nil {
			t.Fatal(err)
		}
		if id <= prev {
			t.Fatalf("id not monotonic: prev=%d cur=%d", prev, id)
		}
		prev = id
	}
}

func TestNextIdConcurrentUnique(t *testing.T) {
	g := MustNew(1, 1)
	const N = 8
	const M = 5000
	seen := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(N)
	dup := make(chan int64, 1)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < M; j++ {
				id, err := g.NextId()
				if err != nil {
					t.Error(err)
					return
				}
				if _, loaded := seen.LoadOrStore(id, struct{}{}); loaded {
					select {
					case dup <- id:
					default:
					}
					return
				}
			}
		}()
	}
	wg.Wait()
	select {
	case id := <-dup:
		t.Fatalf("duplicate id: %d", id)
	default:
	}
}

func TestRangeGuards(t *testing.T) {
	if _, err := New(-1, 0); err == nil {
		t.Fatal("expected error for negative datacenter")
	}
	if _, err := New(32, 0); err == nil {
		t.Fatal("expected error for datacenter > 31")
	}
	if _, err := New(0, 32); err == nil {
		t.Fatal("expected error for worker > 31")
	}
}
