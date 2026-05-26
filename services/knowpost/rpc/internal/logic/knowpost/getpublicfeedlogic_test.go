package knowpostlogic

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

func TestGetPublicFeed_PrefersRedisPageCache(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatalf("new l1: %v", err)
	}

	page := &pb.FeedPage{
		Items: []*pb.FeedItem{
			{Id: "11", Title: "cached-page"},
		},
		HasMore: true,
		Size:    20,
		Page:    1,
	}
	raw, _ := json.Marshal(page)

	pageKey := cachekeys.FeedPublicL1Key(20, 1)
	if err := rdb.Set(context.Background(), pageKey, raw, time.Minute).Err(); err != nil {
		t.Fatalf("set redis page: %v", err)
	}

	logic := NewGetPublicFeedLogic(context.Background(), &svc.ServiceContext{
		Redis:         rdb,
		L1FeedPublic:  l1,
		HotFeedPublic: hotkey.New(hotkey.Config{WindowSeconds: 60, SegmentSeconds: 10, LevelLow: 5, LevelMedium: 10, LevelHigh: 20}),
		HotFeedItem:   hotkey.New(hotkey.Config{WindowSeconds: 60, SegmentSeconds: 10, LevelLow: 5, LevelMedium: 10, LevelHigh: 20}),
	})

	resp, err := logic.GetPublicFeed(&pb.GetPublicFeedReq{Page: 1, Size: 20})
	if err != nil {
		t.Fatalf("get public feed: %v", err)
	}
	if len(resp.Items) != 1 || resp.Items[0].GetTitle() != "cached-page" {
		t.Fatalf("expected redis page cache hit, got %+v", resp.Items)
	}

	l1.Wait()
	cached, ok := l1.Get(pageKey)
	if !ok {
		t.Fatalf("expected l1 page to be backfilled")
	}
	pageRaw, ok := cached.([]byte)
	if !ok || len(pageRaw) == 0 {
		t.Fatalf("expected l1 raw page bytes")
	}
}
