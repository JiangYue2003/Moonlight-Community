package knowpostlogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
	"testing"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/detail"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

const knowpostBenchLatencySampleLimit = 200000

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

type benchKnowPostsModel struct {
	findOneFn        func(context.Context, uint64) (*model.KnowPosts, error)
	listPublicFeedFn func(context.Context, int, int) ([]*model.KnowPosts, error)
}

func (m *benchKnowPostsModel) Insert(context.Context, *model.KnowPosts) (sql.Result, error) {
	panic("not implemented")
}

func (m *benchKnowPostsModel) FindOne(ctx context.Context, id uint64) (*model.KnowPosts, error) {
	if m.findOneFn == nil {
		return nil, model.ErrNotFound
	}
	return m.findOneFn(ctx, id)
}

func (m *benchKnowPostsModel) Update(context.Context, *model.KnowPosts) error { panic("not implemented") }
func (m *benchKnowPostsModel) Delete(context.Context, uint64) error            { panic("not implemented") }

func (m *benchKnowPostsModel) ListPublicFeed(ctx context.Context, limit, offset int) ([]*model.KnowPosts, error) {
	if m.listPublicFeedFn == nil {
		return nil, nil
	}
	return m.listPublicFeedFn(ctx, limit, offset)
}

func (m *benchKnowPostsModel) ListMyFeed(context.Context, uint64, int, int) ([]*model.KnowPosts, error) {
	panic("not implemented")
}

func (m *benchKnowPostsModel) UpdateInTx(context.Context, sqlx.Session, *model.KnowPosts) error {
	panic("not implemented")
}

type detailBenchFixture struct {
	sc       *svc.ServiceContext
	logic    *GetDetailLogic
	rdb      *goredis.Client
	detailL1 *cachex.L1
}

func newDetailBenchFixture(b *testing.B) *detailBenchFixture {
	b.Helper()
	rdb := benchRedis(b)
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 50000, MaxCost: 64 << 20, BufferItems: 64})
	if err != nil {
		b.Fatalf("new detail l1: %v", err)
	}
	l2 := cachex.NewL2(rdb)
	hotCfg := hotkey.Config{WindowSeconds: 60, SegmentSeconds: 10, LevelLow: 1000000, LevelMedium: 2000000, LevelHigh: 3000000}
	hot := hotkey.New(hotCfg)

	sc := &svc.ServiceContext{
		Redis:       rdb,
		DetailCache: detail.New(l1, l2, hot),
		L1FeedItem:  l1,
		HotDetail:   hot,
	}
	return &detailBenchFixture{
		sc:       sc,
		logic:    NewGetDetailLogic(context.Background(), sc),
		rdb:      rdb,
		detailL1: l1,
	}
}

type publicFeedBenchFixture struct {
	sc           *svc.ServiceContext
	logic        *GetPublicFeedLogic
	rdb          *goredis.Client
	l1FeedPublic *cachex.L1
	l1FeedItem   *cachex.L1
}

func newPublicFeedBenchFixture(b *testing.B) *publicFeedBenchFixture {
	b.Helper()
	rdb := benchRedis(b)
	l1FeedPublic, err := cachex.NewL1(cachex.L1Config{NumCounters: 20000, MaxCost: 64 << 20, BufferItems: 64})
	if err != nil {
		b.Fatalf("new feed public l1: %v", err)
	}
	l1FeedItem, err := cachex.NewL1(cachex.L1Config{NumCounters: 50000, MaxCost: 64 << 20, BufferItems: 64})
	if err != nil {
		b.Fatalf("new feed item l1: %v", err)
	}
	hotCfg := hotkey.Config{WindowSeconds: 60, SegmentSeconds: 10, LevelLow: 1000000, LevelMedium: 2000000, LevelHigh: 3000000}
	sc := &svc.ServiceContext{
		Redis:         rdb,
		L1FeedPublic:  l1FeedPublic,
		L1FeedItem:    l1FeedItem,
		HotFeedPublic: hotkey.New(hotCfg),
		HotFeedItem:   hotkey.New(hotCfg),
	}
	return &publicFeedBenchFixture{
		sc:           sc,
		logic:        NewGetPublicFeedLogic(context.Background(), sc),
		rdb:          rdb,
		l1FeedPublic: l1FeedPublic,
		l1FeedItem:   l1FeedItem,
	}
}

func makeBenchKnowpostRow(id uint64) *model.KnowPosts {
	now := time.Now()
	tags, _ := json.Marshal([]string{"cache", "bench", "golang"})
	imgs, _ := json.Marshal([]string{
		fmt.Sprintf("https://cdn.example.com/post/%d/1.png", id),
		fmt.Sprintf("https://cdn.example.com/post/%d/2.png", id),
	})
	return &model.KnowPosts{
		Id:               id,
		CreatorId:        7,
		TagId:            sql.NullInt64{Int64: 3, Valid: true},
		Tags:             sql.NullString{String: string(tags), Valid: true},
		Title:            sql.NullString{String: fmt.Sprintf("bench-title-%d", id), Valid: true},
		Description:      sql.NullString{String: "bench description for cache performance test", Valid: true},
		ContentUrl:       sql.NullString{String: fmt.Sprintf("https://cdn.example.com/content/%d.md", id), Valid: true},
		ContentObjectKey: sql.NullString{String: fmt.Sprintf("content/%d.md", id), Valid: true},
		ContentEtag:      sql.NullString{String: fmt.Sprintf("etag-%d", id), Valid: true},
		ContentSize:      sql.NullInt64{Int64: 2048, Valid: true},
		ContentSha256:    sql.NullString{String: "sha256-bench", Valid: true},
		Visible:          "public",
		Status:           "published",
		Type:             "image_text",
		ImgUrls:          sql.NullString{String: string(imgs), Valid: true},
		IsTop:            0,
		CreateTime:       now,
		UpdateTime:       now,
		PublishTime:      sql.NullTime{Time: now, Valid: true},
	}
}

func makeBenchFeedRows(count int, offset uint64) []*model.KnowPosts {
	rows := make([]*model.KnowPosts, 0, count)
	for i := 0; i < count; i++ {
		rows = append(rows, makeBenchKnowpostRow(offset+uint64(i)+1))
	}
	return rows
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func setFirstErr(mu *sync.Mutex, firstErr *error, err error) {
	if err == nil {
		return
	}
	mu.Lock()
	if *firstErr == nil {
		*firstErr = err
	}
	mu.Unlock()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
