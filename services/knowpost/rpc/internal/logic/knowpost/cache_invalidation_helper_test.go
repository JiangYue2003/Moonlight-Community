package knowpostlogic

import (
	"context"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/mine"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
)

func TestParseFeedPublicIDsKey(t *testing.T) {
	size, page, ok := parseFeedPublicIDsKey("feed:public:ids:20:123:2")
	if !ok || size != 20 || page != 2 {
		t.Fatalf("unexpected parse result: ok=%v size=%d page=%d", ok, size, page)
	}
	if _, _, ok = parseFeedPublicIDsKey("feed:public:20:1:v1"); ok {
		t.Fatalf("unexpected parse success for invalid key")
	}
}

func TestInvalidatePublicFeedPagesByPost(t *testing.T) {
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
	postID := int64(123)
	hour := cachekeys.HourSlot(time.Now())
	ridx := cachekeys.FeedReverseIndexKey(postID, hour)
	idsKey := cachekeys.FeedPublicIdsKey(20, 2, hour)
	hasMoreKey := cachekeys.FeedPublicHasMoreKey(idsKey)
	l1PageKey := cachekeys.FeedPublicL1Key(20, 2)
	pageKey := cachekeys.FeedPublicPageKey(20, 2)

	if err := rdb.SAdd(ctx, ridx, pageKey).Err(); err != nil {
		t.Fatalf("sadd ridx: %v", err)
	}
	if err := rdb.RPush(ctx, idsKey, "1", "2").Err(); err != nil {
		t.Fatalf("rpush ids: %v", err)
	}
	if err := rdb.Set(ctx, hasMoreKey, "1", time.Minute).Err(); err != nil {
		t.Fatalf("set hasMore: %v", err)
	}
	if err := rdb.Set(ctx, pageKey, "page", time.Minute).Err(); err != nil {
		t.Fatalf("set page: %v", err)
	}
	if err := rdb.SAdd(ctx, cachekeys.FeedAllPagesKey, pageKey).Err(); err != nil {
		t.Fatalf("sadd all pages: %v", err)
	}
	sc.L1FeedPublic.SetWithTTL(l1PageKey, []byte("x"), 1, time.Minute)
	sc.L1FeedPublic.Wait()

	invalidatePublicFeedPagesByPost(ctx, sc, postID)

	if rdb.Exists(ctx, idsKey).Val() != 0 || rdb.Exists(ctx, hasMoreKey).Val() != 0 || rdb.Exists(ctx, pageKey).Val() != 0 || rdb.Exists(ctx, ridx).Val() != 0 {
		t.Fatalf("public feed cache keys should be deleted")
	}
	if rdb.SIsMember(ctx, cachekeys.FeedAllPagesKey, pageKey).Val() {
		t.Fatalf("page key should be removed from feed:public:pages")
	}
	if _, ok := sc.L1FeedPublic.Get(l1PageKey); ok {
		t.Fatalf("l1 feed public page should be deleted")
	}
}

func TestInvalidateMineFeedPagesByCreator(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	l1Mine, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatalf("new l1 mine: %v", err)
	}
	l2 := cachex.NewL2(rdb)
	feedMineCache := mine.New(l1Mine, l2, nil)
	sc := &svc.ServiceContext{
		Redis:        rdb,
		FeedMineCache: feedMineCache,
	}

	ctx := context.Background()
	k1 := cachekeys.FeedMineKey(7, 20, 1)
	k2 := cachekeys.FeedMineKey(7, 20, 2)
	kOther := cachekeys.FeedMineKey(8, 20, 1)
	_ = rdb.Set(ctx, k1, "x", time.Minute).Err()
	_ = rdb.Set(ctx, k2, "y", time.Minute).Err()
	_ = rdb.Set(ctx, kOther, "z", time.Minute).Err()

	invalidateMineFeedPagesByCreator(ctx, sc, 7)

	if rdb.Exists(ctx, k1).Val() != 0 || rdb.Exists(ctx, k2).Val() != 0 {
		t.Fatalf("creator mine feed keys should be deleted")
	}
	if rdb.Exists(ctx, kOther).Val() != 1 {
		t.Fatalf("other creator key should remain")
	}
}
