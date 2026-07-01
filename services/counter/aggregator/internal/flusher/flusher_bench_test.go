package flusher

import (
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/counterlua"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

const flushBenchLatencySampleLimit = 20000

func benchRedis(b *testing.B) *goredis.Client {
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

	rdb := goredis.NewClient(&goredis.Options{
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

func BenchmarkFlushKeyLikeDelta(b *testing.B) {
	ctx := context.Background()
	rdb := benchRedis(b)
	sc := &svc.ServiceContext{
		Redis:           rdb,
		IncrFieldScript: goredis.NewScript(counterlua.IncrField),
		DecrFieldScript: goredis.NewScript(counterlua.DecrField),
	}

	deltas := []int64{1, 100, 1000, 10000, 100000}
	for _, delta := range deltas {
		delta := delta
		b.Run(fmt.Sprintf("delta=%d", delta), func(b *testing.B) {
			samples := make([]int64, min(b.N, flushBenchLatencySampleLimit))
			postIDPrefix := fmt.Sprintf("bench-flush-%d-%d", delta, time.Now().UnixNano())
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				postID := fmt.Sprintf("%s-%d", postIDPrefix, i)
				aggKey := schema.AggKey("knowpost", postID)
				sdsKey := schema.SdsKey("knowpost", postID)

				if err := rdb.Del(ctx, aggKey, sdsKey).Err(); err != nil {
					b.Fatalf("cleanup keys: %v", err)
				}
				if err := rdb.HSet(ctx, aggKey, strconv.Itoa(schema.IdxLike), delta).Err(); err != nil {
					b.Fatalf("seed agg key: %v", err)
				}

				t0 := time.Now()
				if err := flushKey(ctx, sc, aggKey); err != nil {
					b.Fatalf("flushKey: %v", err)
				}
				dur := time.Since(t0).Nanoseconds()
				if i < len(samples) {
					samples[i] = dur
				}

				raw, err := rdb.Get(ctx, sdsKey).Bytes()
				if err != nil {
					b.Fatalf("read sds: %v", err)
				}
				if got := readBenchFieldUint32(b, raw, schema.IdxLike); got != uint32(delta) {
					b.Fatalf("like field mismatch: got %d want %d", got, delta)
				}

				remaining, err := rdb.HGet(ctx, aggKey, strconv.Itoa(schema.IdxLike)).Result()
				if err != goredis.Nil && err != nil {
					b.Fatalf("read agg remainder: %v", err)
				}
				if remaining != "" && remaining != "0" {
					b.Fatalf("agg key not drained, field=%q", remaining)
				}
			}
			b.StopTimer()

			elapsed := b.Elapsed()
			reportLatencyMetrics(b, samples, min(b.N, len(samples)), elapsed, int64(b.N), "flush_ops/s")
			b.ReportMetric(float64(delta*int64(b.N))/elapsed.Seconds(), "likes_flushed/s")
		})
	}
}
