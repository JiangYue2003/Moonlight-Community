package cachex

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

const cacheBenchLatencySampleLimit = 200000

func benchRedis(b *testing.B) *redis.Client {
	b.Helper()

	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	password := os.Getenv("REDIS_PASSWORD")
	db := 15
	if raw := os.Getenv("REDIS_BENCH_DB"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil {
			b.Fatalf("invalid REDIS_BENCH_DB %q: %v", raw, err)
		}
		db = n
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		b.Fatalf("ping redis at %s db=%d: %v", addr, db, err)
	}
	b.Cleanup(func() { _ = rdb.Close() })
	return rdb
}

func reportBenchLatencyMetrics(b *testing.B, samples []int64, sampleCount int, elapsed time.Duration, totalOps int64, opMetric string) {
	b.Helper()
	if elapsed <= 0 || totalOps <= 0 {
		return
	}
	b.ReportMetric(float64(totalOps)/elapsed.Seconds(), opMetric)
	if sampleCount == 0 {
		return
	}
	cp := append([]int64(nil), samples[:sampleCount]...)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	b.ReportMetric(float64(percentileNanos(cp, 0.50))/1e6, "p50_ms")
	b.ReportMetric(float64(percentileNanos(cp, 0.95))/1e6, "p95_ms")
	b.ReportMetric(float64(percentileNanos(cp, 0.99))/1e6, "p99_ms")
}

func percentileNanos(sorted []int64, q float64) int64 {
	if len(sorted) == 0 {
		return 0
	}
	if q <= 0 {
		return sorted[0]
	}
	if q >= 1 {
		return sorted[len(sorted)-1]
	}
	idx := int(float64(len(sorted)-1) * q)
	return sorted[idx]
}

func newBenchCache(b *testing.B) (Cache[*item], *L1, *redis.Client) {
	b.Helper()
	rdb := benchRedis(b)
	l1, err := NewL1(L1Config{NumCounters: 50000, MaxCost: 32 << 20, BufferItems: 64})
	if err != nil {
		b.Fatalf("new l1: %v", err)
	}
	l2 := NewL2(rdb)
	c := New[*item](Options[*item]{
		L1: l1, L2: l2,
		BaseTTL: 60 * time.Second, JitterMax: 0,
		NullTTL: 30 * time.Second, NullJitterMax: 0,
		Marshal:   func(v *item) ([]byte, error) { return json.Marshal(v) },
		Unmarshal: func(b []byte, v **item) error { var x item; err := json.Unmarshal(b, &x); *v = &x; return err },
	})
	return c, l1, rdb
}

func BenchmarkCacheGetOrLoad_L1Hit(b *testing.B) {
	ctx := context.Background()
	c, l1, rdb := newBenchCache(b)
	key := "bench:cachex:l1"
	sfKey := "sf:" + key

	raw, _ := json.Marshal(&item{Id: 1, Body: "warm-l1"})
	l1.SetWithTTL(key, raw, int64(len(raw)), time.Minute)
	l1.Wait()
	_ = rdb.Del(ctx, key).Err()

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, cacheBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t0 := time.Now()
			v, err := c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				b.Fatal("loader should not run on l1 hit")
				return nil, nil
			})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || v == nil || v.Id != 1 {
				b.Fatalf("l1 hit failed: v=%+v err=%v", v, err)
			}
			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
}

func BenchmarkCacheGetOrLoad_L2Hit(b *testing.B) {
	ctx := context.Background()
	c, l1, rdb := newBenchCache(b)
	key := "bench:cachex:l2"
	sfKey := "sf:" + key

	raw, _ := json.Marshal(&item{Id: 2, Body: "warm-l2"})
	if err := rdb.Set(ctx, key, raw, time.Minute).Err(); err != nil {
		b.Fatalf("seed l2: %v", err)
	}
	l1.Del(key)

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, cacheBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l1.Del(key)
			t0 := time.Now()
			v, err := c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				b.Fatal("loader should not run on l2 hit")
				return nil, nil
			})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || v == nil || v.Id != 2 {
				b.Fatalf("l2 hit failed: v=%+v err=%v", v, err)
			}
			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
}

func BenchmarkCacheGetOrLoad_L3Miss(b *testing.B) {
	ctx := context.Background()
	c, l1, rdb := newBenchCache(b)
	keyPrefix := "bench:cachex:l3:"

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, cacheBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			seq := totalOps.Add(1)
			key := keyPrefix + strconv.FormatInt(seq, 10)
			sfKey := "sf:" + key
			l1.Del(key)
			_ = rdb.Del(ctx, key).Err()

			t0 := time.Now()
			v, err := c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				loaderCalls.Add(1)
				return &item{Id: seq, Body: "from-loader"}, nil
			})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || v == nil || v.Id != seq {
				b.Fatalf("l3 miss failed: v=%+v err=%v", v, err)
			}
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()

	reportBenchLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}

func BenchmarkCacheGetOrLoad_NotFoundSentinel(b *testing.B) {
	ctx := context.Background()
	c, l1, rdb := newBenchCache(b)
	keyPrefix := "bench:cachex:null:"

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, cacheBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			seq := totalOps.Add(1)
			key := keyPrefix + strconv.FormatInt(seq, 10)
			sfKey := "sf:" + key
			l1.Del(key)
			_ = rdb.Del(ctx, key).Err()

			t0 := time.Now()
			_, err := c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				loaderCalls.Add(1)
				return nil, ErrNotFound
			})
			dur := time.Since(t0).Nanoseconds()
			if !errors.Is(err, ErrNotFound) {
				b.Fatalf("expected ErrNotFound, got %v", err)
			}

			_, err = c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				loaderCalls.Add(1)
				return nil, ErrNotFound
			})
			if !errors.Is(err, ErrNotFound) {
				b.Fatalf("sentinel second read expected ErrNotFound, got %v", err)
			}

			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()

	reportBenchLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}

func BenchmarkCacheGetOrLoad_SingleFlightSameKey(b *testing.B) {
	ctx := context.Background()
	c, l1, rdb := newBenchCache(b)
	key := "bench:cachex:sf"
	sfKey := "sf:" + key
	_ = rdb.Del(ctx, key).Err()
	l1.Del(key)

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, cacheBenchLatencySampleLimit))
	var firstErr error
	var firstErrMu sync.Mutex

	setErr := func(err error) {
		if err == nil {
			return
		}
		firstErrMu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		firstErrMu.Unlock()
	}

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t0 := time.Now()
			v, err := c.GetOrLoad(ctx, key, sfKey, func(context.Context) (*item, error) {
				loaderCalls.Add(1)
				time.Sleep(500 * time.Microsecond)
				return &item{Id: 9, Body: "singleflight"}, nil
			})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || v == nil || v.Id != 9 {
				setErr(err)
				return
			}
			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
			l1.Del(key)
			_ = rdb.Del(ctx, key).Err()
		}
	})
	b.StopTimer()

	if firstErr != nil {
		b.Fatal(firstErr)
	}
	reportBenchLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
