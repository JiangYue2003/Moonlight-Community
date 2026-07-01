package counterlua

import (
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

const benchLatencySampleLimit = 200000

// These benchmarks require a real Redis instance. By default they target
// 127.0.0.1:6379 DB 15; override with REDIS_ADDR / REDIS_PASSWORD / REDIS_BENCH_DB.
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

func reportLatencyMetrics(b *testing.B, samples []int64, sampleCount int, elapsed time.Duration, totalOps int64, opMetric string) {
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

func readBenchFieldUint32(b *testing.B, raw []byte, idx int) uint32 {
	b.Helper()
	off := idx * schema.FieldSize
	if off+schema.FieldSize > len(raw) {
		b.Fatalf("buffer too short: len=%d need %d", len(raw), off+schema.FieldSize)
	}
	return binary.BigEndian.Uint32(raw[off : off+schema.FieldSize])
}

func BenchmarkToggleLikeHotspot(b *testing.B) {
	ctx := context.Background()
	rdb := benchRedis(b)
	script := redis.NewScript(Toggle)

	postID := fmt.Sprintf("bench-toggle-%d", time.Now().UnixNano())
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, benchLatencySampleLimit))
	var firstErr error
	var firstErrMu sync.Mutex

	setFirstErr := func(err error) {
		if err == nil {
			return
		}
		firstErrMu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		firstErrMu.Unlock()
	}

	b.Cleanup(func() {
		chunks := int(schema.ChunkOf(totalOps.Load()+1)) + 1
		keys := make([]string, 0, chunks)
		for i := 0; i < chunks; i++ {
			keys = append(keys, schema.BitmapKey(schema.MetricLike, "knowpost", postID, int64(i)))
		}
		if len(keys) > 0 {
			_ = rdb.Del(ctx, keys...).Err()
		}
	})

	b.SetParallelism(4)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uid := totalOps.Add(1)
			key := schema.BitmapKey(schema.MetricLike, "knowpost", postID, schema.ChunkOf(uid))

			t0 := time.Now()
			changed, err := script.Run(ctx, rdb, []string{key}, schema.BitOf(uid), "add").Int64()
			dur := time.Since(t0).Nanoseconds()
			if err != nil {
				setFirstErr(fmt.Errorf("toggle like: %w", err))
				return
			}
			if changed != 1 {
				setFirstErr(fmt.Errorf("toggle like expected changed=1, got %d", changed))
				return
			}

			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()

	if firstErr != nil {
		b.Fatal(firstErr)
	}
	reportLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "toggle_ops/s")
}

func BenchmarkIncrFieldLikeHotspot(b *testing.B) {
	ctx := context.Background()
	rdb := benchRedis(b)
	script := redis.NewScript(IncrField)

	postID := fmt.Sprintf("bench-incr-%d", time.Now().UnixNano())
	sdsKey := schema.SdsKey("knowpost", postID)
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, benchLatencySampleLimit))
	var firstErr error
	var firstErrMu sync.Mutex

	setFirstErr := func(err error) {
		if err == nil {
			return
		}
		firstErrMu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		firstErrMu.Unlock()
	}

	_ = rdb.Del(ctx, sdsKey).Err()
	b.Cleanup(func() { _ = rdb.Del(ctx, sdsKey).Err() })

	b.SetParallelism(4)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t0 := time.Now()
			_, err := script.Run(ctx, rdb, []string{sdsKey},
				schema.IdxLike, 1, schema.SchemaLen, schema.FieldSize).Int64()
			dur := time.Since(t0).Nanoseconds()
			if err != nil {
				setFirstErr(fmt.Errorf("incr like field: %w", err))
				return
			}

			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()

	if firstErr != nil {
		b.Fatal(firstErr)
	}

	raw, err := rdb.Get(ctx, sdsKey).Bytes()
	if err != nil {
		b.Fatalf("read sds after benchmark: %v", err)
	}
	if got := readBenchFieldUint32(b, raw, schema.IdxLike); got != uint32(totalOps.Load()) {
		b.Fatalf("like field mismatch: got %d want %d", got, totalOps.Load())
	}

	reportLatencyMetrics(b, samples, int(min(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "sds_ops/s")
}
