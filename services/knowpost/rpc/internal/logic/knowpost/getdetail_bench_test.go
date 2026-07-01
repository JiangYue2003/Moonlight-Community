package knowpostlogic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

func BenchmarkGetDetail_L1Hit(b *testing.B) {
	fixture := newDetailBenchFixture(b)
	postID := int64(9001)
	key := cachekeys.DetailKey(postID)
	row := makeBenchKnowpostRow(uint64(postID))
	raw, _ := json.Marshal(rowToDetail(row))
	fixture.detailL1.SetWithTTL(key, raw, int64(len(raw)), time.Minute)
	fixture.detailL1.Wait()

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		findOneFn: func(context.Context, uint64) (*model.KnowPosts, error) {
			b.Fatal("FindOne should not run on L1 hit")
			return nil, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t0 := time.Now()
			resp, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || resp.Id != "9001" {
				b.Fatalf("detail l1 hit failed: resp=%+v err=%v", resp, err)
			}
			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min64(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
}

func BenchmarkGetDetail_L2Hit(b *testing.B) {
	fixture := newDetailBenchFixture(b)
	ctx := context.Background()
	postID := int64(9002)
	key := cachekeys.DetailKey(postID)
	row := makeBenchKnowpostRow(uint64(postID))
	raw, _ := json.Marshal(rowToDetail(row))
	if err := fixture.rdb.Set(ctx, key, raw, time.Minute).Err(); err != nil {
		b.Fatalf("seed redis detail: %v", err)
	}

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		findOneFn: func(context.Context, uint64) (*model.KnowPosts, error) {
			b.Fatal("FindOne should not run on L2 hit")
			return nil, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fixture.detailL1.Del(key)
			t0 := time.Now()
			resp, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || resp.Id != "9002" {
				b.Fatalf("detail l2 hit failed: resp=%+v err=%v", resp, err)
			}
			totalOps.Add(1)
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min64(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
}

func BenchmarkGetDetail_L3Miss(b *testing.B) {
	fixture := newDetailBenchFixture(b)
	ctx := context.Background()
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		findOneFn: func(_ context.Context, id uint64) (*model.KnowPosts, error) {
			loaderCalls.Add(1)
			return makeBenchKnowpostRow(id), nil
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			seq := totalOps.Add(1)
			postID := int64(100000 + seq)
			key := cachekeys.DetailKey(postID)
			fixture.detailL1.Del(key)
			_ = fixture.rdb.Del(ctx, key).Err()

			t0 := time.Now()
			resp, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || resp.Id != fmt.Sprintf("%d", postID) {
				b.Fatalf("detail l3 miss failed: resp=%+v err=%v", resp, err)
			}
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min64(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}

func BenchmarkGetDetail_ParallelSameKey(b *testing.B) {
	fixture := newDetailBenchFixture(b)
	ctx := context.Background()
	postID := int64(9003)
	key := cachekeys.DetailKey(postID)
	var loaderCalls atomic.Int64

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		findOneFn: func(context.Context, uint64) (*model.KnowPosts, error) {
			loaderCalls.Add(1)
			time.Sleep(500 * time.Microsecond)
			return makeBenchKnowpostRow(uint64(postID)), nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))
	var firstErr error
	var firstErrMu sync.Mutex

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fixture.detailL1.Del(key)
			_ = fixture.rdb.Del(ctx, key).Err()
			t0 := time.Now()
			resp, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || resp.Id != "9003" {
				setFirstErr(&firstErrMu, &firstErr, fmt.Errorf("detail same key failed: resp=%+v err=%v", resp, err))
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
	reportBenchLatencyMetrics(b, samples, int(min64(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}

func BenchmarkGetDetail_NotFoundSentinel(b *testing.B) {
	fixture := newDetailBenchFixture(b)
	ctx := context.Background()
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		findOneFn: func(context.Context, uint64) (*model.KnowPosts, error) {
			loaderCalls.Add(1)
			return nil, model.ErrNotFound
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			seq := totalOps.Add(1)
			postID := int64(200000 + seq)
			key := cachekeys.DetailKey(postID)
			fixture.detailL1.Del(key)
			_ = fixture.rdb.Del(ctx, key).Err()

			t0 := time.Now()
			if _, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID}); err == nil {
				b.Fatal("expected not found error")
			}
			dur := time.Since(t0).Nanoseconds()
			if _, err := fixture.logic.GetDetail(&knowpost.GetDetailReq{Id: postID}); err == nil {
				b.Fatal("expected sentinel-backed not found error")
			}
			idx := sampleIdx.Add(1) - 1
			if idx < int64(len(samples)) {
				samples[idx] = dur
			}
		}
	})
	b.StopTimer()
	reportBenchLatencyMetrics(b, samples, int(min64(sampleIdx.Load(), int64(len(samples)))), b.Elapsed(), totalOps.Load(), "ops/s")
	b.ReportMetric(float64(loaderCalls.Load())/float64(maxInt64(1, totalOps.Load())), "loader_calls/op")
}
