package listener

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/event"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

func TestHandleCounterEvent_RebuildL1FromFragments(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatalf("new l1: %v", err)
	}
	sc := &svc.ServiceContext{
		Redis:        rdb,
		L1FeedPublic: l1,
	}
	ctx := context.Background()

	eid := int64(101)
	hour := cachekeys.HourSlot(time.Now())
	idsKey := cachekeys.FeedPublicIdsKey(20, 1, hour)
	ridx := cachekeys.FeedReverseIndexKey(eid, hour)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	pageKey := cachekeys.FeedPublicL1Key(20, 1)

	_ = rdb.SAdd(ctx, ridx, pageKey).Err()
	_ = rdb.RPush(ctx, idsKey, "101", "202").Err()
	_ = rdb.Set(ctx, pageKey, `{"items":[],"hasMore":false,"size":20,"page":1}`, 50*time.Second).Err()
	_ = rdb.Set(ctx, hasMoreKey, "1", 50*time.Second).Err()
	_ = rdb.Expire(ctx, idsKey, 50*time.Second).Err()

	item1, _ := json.Marshal(&pb.FeedItem{Id: "101", Title: "a"})
	item2, _ := json.Marshal(&pb.FeedItem{Id: "202", Title: "b"})
	_ = rdb.Set(ctx, cachekeys.FeedItemKey(101), item1, 50*time.Second).Err()
	_ = rdb.Set(ctx, cachekeys.FeedItemKey(202), item2, 50*time.Second).Err()

	handleCounterEvent(ctx, sc, event.CounterEvent{
		EntityType: "knowpost",
		EntityId:   "101",
		Metric:     "like",
		Delta:      1,
	})

	sc.L1FeedPublic.Wait()
	raw, ok := sc.L1FeedPublic.Get(pageKey)
	if !ok {
		t.Fatalf("expected l1 page cached")
	}
	pageRaw, ok := raw.([]byte)
	if !ok || len(pageRaw) == 0 {
		t.Fatalf("expected l1 page raw bytes")
	}
	var page pb.FeedPage
	if err := json.Unmarshal(pageRaw, &page); err != nil {
		t.Fatalf("unmarshal page: %v", err)
	}
	if len(page.Items) != 2 || page.Items[0].GetId() != "101" {
		t.Fatalf("unexpected page items: %+v", page.Items)
	}
	if page.GetHasMore() != true {
		t.Fatalf("expected hasMore=true")
	}
	redisPageRaw, err := rdb.Get(ctx, pageKey).Bytes()
	if err != nil || len(redisPageRaw) == 0 {
		t.Fatalf("expected redis page preserved and rewritten")
	}
}

func TestHandleCounterEvent_FallbackInvalidateWhenFragmentBroken(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatalf("new l1: %v", err)
	}
	sc := &svc.ServiceContext{
		Redis:        rdb,
		L1FeedPublic: l1,
	}
	ctx := context.Background()

	eid := int64(303)
	hour := cachekeys.HourSlot(time.Now())
	idsKey := cachekeys.FeedPublicIdsKey(20, 2, hour)
	ridx := cachekeys.FeedReverseIndexKey(eid, hour)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	pageKey := cachekeys.FeedPublicL1Key(20, 2)

	_ = rdb.SAdd(ctx, ridx, pageKey).Err()
	_ = rdb.SAdd(ctx, cachekeys.FeedAllPagesKey, pageKey).Err()
	_ = rdb.RPush(ctx, idsKey, "303", "404").Err()
	_ = rdb.Set(ctx, pageKey, `{"items":[],"hasMore":false,"size":20,"page":2}`, time.Minute).Err()
	_ = rdb.Set(ctx, hasMoreKey, "0", time.Minute).Err()
	// 只放一个 feed:item，触发重建失败。
	item1, _ := json.Marshal(&pb.FeedItem{Id: "303"})
	_ = rdb.Set(ctx, cachekeys.FeedItemKey(303), item1, time.Minute).Err()
	sc.L1FeedPublic.SetWithTTL(pageKey, []byte("old"), 3, time.Minute)
	sc.L1FeedPublic.Wait()

	handleCounterEvent(ctx, sc, event.CounterEvent{
		EntityType: "knowpost",
		EntityId:   "303",
		Metric:     "fav",
		Delta:      1,
	})

	if rdb.Exists(ctx, idsKey).Val() != 0 || rdb.Exists(ctx, hasMoreKey).Val() != 0 || rdb.Exists(ctx, pageKey).Val() != 0 {
		t.Fatalf("fallback should invalidate page fragments")
	}
	if rdb.SIsMember(ctx, cachekeys.FeedAllPagesKey, pageKey).Val() {
		t.Fatalf("page key should be removed from all pages set")
	}
	if rdb.SIsMember(ctx, ridx, pageKey).Val() {
		t.Fatalf("page key should be removed from reverse index set")
	}
	if _, ok := sc.L1FeedPublic.Get(pageKey); ok {
		t.Fatalf("l1 page should be invalidated on fallback")
	}
}

func TestHandleCounterEvent_IgnoreNonKnowpostOrMetric(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatalf("new l1: %v", err)
	}
	sc := &svc.ServiceContext{
		Redis:        rdb,
		L1FeedPublic: l1,
	}
	ctx := context.Background()

	eid := int64(505)
	hour := cachekeys.HourSlot(time.Now())
	idsKey := cachekeys.FeedPublicIdsKey(20, 3, hour)
	ridx := cachekeys.FeedReverseIndexKey(eid, hour)
	pageKey := cachekeys.FeedPublicL1Key(20, 3)
	_ = rdb.SAdd(ctx, ridx, pageKey).Err()
	_ = rdb.RPush(ctx, idsKey, "505").Err()
	_ = rdb.Set(ctx, cachekeys.FeedPublicHasMoreKey(idsKey), "0", time.Minute).Err()

	handleCounterEvent(ctx, sc, event.CounterEvent{
		EntityType: "user",
		EntityId:   "505",
		Metric:     "like",
		Delta:      1,
	})
	handleCounterEvent(ctx, sc, event.CounterEvent{
		EntityType: "knowpost",
		EntityId:   "505",
		Metric:     "view",
		Delta:      1,
	})

	if rdb.Exists(ctx, idsKey).Val() != 1 {
		t.Fatalf("ignored events should not touch feed fragments")
	}
}
