package processor

import (
	"context"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
)

func TestDedupAcquire(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{Addrs: []string{mr.Addr()}})
	d := NewDedup(rdb)
	ok, err := d.Acquire(context.Background(), "k", time.Minute)
	if err != nil || !ok {
		t.Fatalf("first acquire err=%v ok=%v", err, ok)
	}
	ok, err = d.Acquire(context.Background(), "k", time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("second acquire should be false")
	}
}

func TestShortErr(t *testing.T) {
	e := shortErr(context.DeadlineExceeded)
	if e == "" {
		t.Fatal("shortErr should not be empty")
	}
}
