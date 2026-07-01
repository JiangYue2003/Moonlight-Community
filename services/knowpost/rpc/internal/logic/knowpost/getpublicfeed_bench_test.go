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
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

func BenchmarkGetPublicFeed_L1PageHit(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	page, size := 1, 20
	l1Key := cachekeys.FeedPublicL1Key(size, page)
	pageResp := &pb.FeedPage{
		Items:   feedItemsFromRows(makeBenchFeedRows(size, 0)),
		HasMore: true,
		Size:    int32(size),
		Page:    int32(page),
	}
	raw, _ := json.Marshal(pageResp)
	fixture.l1FeedPublic.SetWithTTL(l1Key, raw, int64(len(raw)), time.Minute)
	fixture.l1FeedPublic.Wait()

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(context.Context, int, int) ([]*model.KnowPosts, error) {
			b.Fatal("ListPublicFeed should not run on L1 hit")
			return nil, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(page), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) != size {
				b.Fatalf("public feed l1 hit failed: resp=%+v err=%v", resp, err)
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

func BenchmarkGetPublicFeed_RedisPageHit(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	ctx := context.Background()
	page, size := 1, 20
	redisPageKey := cachekeys.FeedPublicPageKey(size, page)
	pageResp := &pb.FeedPage{
		Items:   feedItemsFromRows(makeBenchFeedRows(size, 0)),
		HasMore: true,
		Size:    int32(size),
		Page:    int32(page),
	}
	raw, _ := json.Marshal(pageResp)
	if err := fixture.rdb.Set(ctx, redisPageKey, raw, time.Minute).Err(); err != nil {
		b.Fatalf("seed redis page: %v", err)
	}

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(context.Context, int, int) ([]*model.KnowPosts, error) {
			b.Fatal("ListPublicFeed should not run on redis page hit")
			return nil, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			fixture.l1FeedPublic.Del(cachekeys.FeedPublicL1Key(size, page))
			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(page), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) != size {
				b.Fatalf("public feed redis page hit failed: resp=%+v err=%v", resp, err)
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

func BenchmarkGetPublicFeed_RedisFragmentsHit(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	ctx := context.Background()
	page, size := 1, 20
	rows := makeBenchFeedRows(size, 1000)
	hourSlot := cachekeys.HourSlot(time.Now())
	idsKey := cachekeys.FeedPublicIdsKey(size, page, hourSlot)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	pageKey := cachekeys.FeedPublicPageKey(size, page)

	idVals := make([]any, 0, len(rows))
	for _, row := range rows {
		idVals = append(idVals, row.Id)
		raw, _ := json.Marshal(rowToFeedItem(row))
		if err := fixture.rdb.Set(ctx, cachekeys.FeedItemKey(int64(row.Id)), raw, time.Minute).Err(); err != nil {
			b.Fatalf("seed feed item: %v", err)
		}
	}
	if err := fixture.rdb.Del(ctx, idsKey).Err(); err != nil {
		b.Fatalf("del ids key: %v", err)
	}
	if err := fixture.rdb.RPush(ctx, idsKey, idVals...).Err(); err != nil {
		b.Fatalf("seed ids key: %v", err)
	}
	if err := fixture.rdb.Expire(ctx, idsKey, time.Minute).Err(); err != nil {
		b.Fatalf("expire ids key: %v", err)
	}
	if err := fixture.rdb.Set(ctx, hasMoreKey, "1", time.Minute).Err(); err != nil {
		b.Fatalf("seed hasMore: %v", err)
	}

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(context.Context, int, int) ([]*model.KnowPosts, error) {
			b.Fatal("ListPublicFeed should not run on fragment hit")
			return nil, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			fixture.l1FeedPublic.Del(cachekeys.FeedPublicL1Key(size, page))
			_ = fixture.rdb.Del(ctx, pageKey).Err()
			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(page), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) != size {
				b.Fatalf("public feed fragment hit failed: resp=%+v err=%v", resp, err)
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

func BenchmarkGetPublicFeed_DBMissWriteBack(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	ctx := context.Background()
	size := 20
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	var loaderCalls atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(_ context.Context, limit, offset int) ([]*model.KnowPosts, error) {
			loaderCalls.Add(1)
			return makeBenchFeedRows(limit, uint64(offset*1000)), nil
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			seq := totalOps.Add(1)
			pageNow := int(seq)

			pageKey := cachekeys.FeedPublicPageKey(size, pageNow)
			l1Key := cachekeys.FeedPublicL1Key(size, pageNow)
			hourSlot := cachekeys.HourSlot(time.Now())
			idsKey := cachekeys.FeedPublicIdsKey(size, pageNow, hourSlot)
			hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
			fixture.l1FeedPublic.Del(l1Key)
			_ = fixture.rdb.Del(ctx, pageKey, idsKey, hasMoreKey).Err()

			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(pageNow), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) != size {
				b.Fatalf("public feed db miss failed: resp=%+v err=%v", resp, err)
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

func BenchmarkGetPublicFeed_ParallelSamePage(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	ctx := context.Background()
	page, size := 1, 20
	postRows := makeBenchFeedRows(size+1, 5000)
	pageKey := cachekeys.FeedPublicPageKey(size, page)
	l1Key := cachekeys.FeedPublicL1Key(size, page)
	hourSlot := cachekeys.HourSlot(time.Now())
	idsKey := cachekeys.FeedPublicIdsKey(size, page, hourSlot)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	var loaderCalls atomic.Int64
	var firstErr error
	var firstErrMu sync.Mutex

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(context.Context, int, int) ([]*model.KnowPosts, error) {
			loaderCalls.Add(1)
			time.Sleep(500 * time.Microsecond)
			return postRows, nil
		},
	}

	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			fixture.l1FeedPublic.Del(l1Key)
			_ = fixture.rdb.Del(ctx, pageKey, idsKey, hasMoreKey).Err()
			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(page), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) != size {
				setFirstErr(&firstErrMu, &firstErr, fmt.Errorf("public feed same page failed: resp=%+v err=%v", resp, err))
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

func BenchmarkGetPublicFeed_ParallelZipfPages(b *testing.B) {
	fixture := newPublicFeedBenchFixture(b)
	size := 20
	var loaderCalls atomic.Int64
	var seq atomic.Int64
	var totalOps atomic.Int64
	var sampleIdx atomic.Int64
	samples := make([]int64, min(b.N, knowpostBenchLatencySampleLimit))
	var firstErr error
	var firstErrMu sync.Mutex

	fixture.sc.KnowPostsModel = &benchKnowPostsModel{
		listPublicFeedFn: func(context.Context, int, int) ([]*model.KnowPosts, error) {
			loaderCalls.Add(1)
			return makeBenchFeedRows(size+1, 10000), nil
		},
	}

	b.SetParallelism(8)
	b.ResetTimer()
	b.RunParallel(func(pbIt *testing.PB) {
		for pbIt.Next() {
			n := seq.Add(1)
			page := zipfLikePage(n)
			t0 := time.Now()
			resp, err := fixture.logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: int32(page), Size: int32(size)})
			dur := time.Since(t0).Nanoseconds()
			if err != nil || resp == nil || len(resp.Items) == 0 {
				setFirstErr(&firstErrMu, &firstErr, fmt.Errorf("public feed zipf page failed: page=%d resp=%+v err=%v", page, resp, err))
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

func feedItemsFromRows(rows []*model.KnowPosts) []*pb.FeedItem {
	items := make([]*pb.FeedItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, rowToFeedItem(row))
	}
	return items
}

func zipfLikePage(n int64) int {
	switch {
	case n%10 < 6:
		return 1
	case n%10 < 8:
		return 2
	case n%10 == 8:
		return 3
	default:
		return int((n % 20) + 1)
	}
}
