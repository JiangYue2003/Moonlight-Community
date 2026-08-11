package knowpostlogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

type oneFollowingRelation struct{}

func (oneFollowingRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func (oneFollowingRelation) GetFollowings(context.Context, int64) ([]int64, error) {
	return []int64{7}, nil
}

type zeroFollowerCounter struct{}

func (zeroFollowerCounter) BatchGetFollowerCounts(_ context.Context, userIDs []int64) (map[int64]int64, error) {
	counts := make(map[int64]int64, len(userIDs))
	for _, userID := range userIDs {
		counts[userID] = 0
	}
	return counts, nil
}

func TestGetUserFeedNormalizesPagination(t *testing.T) {
	redisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf("start miniredis: %v", err)
	}
	t.Cleanup(redisServer.Close)

	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() {
		_ = redisClient.Close()
	})

	ctx := context.Background()
	serviceCtx := &svc.ServiceContext{
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			oneFollowingRelation{},
			zeroFollowerCounter{},
			logx.WithContext(ctx),
		),
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Errorf("GetUserFeed panicked for invalid pagination: %v", recovered)
		}
	}()

	resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   0,
		Size:   100,
	})
	if err != nil {
		t.Fatalf("GetUserFeed failed: %v", err)
	}
	if resp.Page != 1 {
		t.Fatalf("page = %d, want 1", resp.Page)
	}
	if resp.Size != 50 {
		t.Fatalf("size = %d, want 50", resp.Size)
	}
}

type recordingFeedPostLoader struct {
	rows      map[uint64]*model.KnowPosts
	calls     int
	requested []uint64
}

func (l *recordingFeedPostLoader) FindPublishedFeedByIDs(_ context.Context, ids []uint64) ([]*model.KnowPosts, error) {
	l.calls++
	l.requested = append([]uint64(nil), ids...)
	rows := make([]*model.KnowPosts, 0, len(ids))
	for i := len(ids) - 1; i >= 0; i-- {
		if row := l.rows[ids[i]]; row != nil {
			rows = append(rows, row)
		}
	}
	return rows, nil
}

func TestGetUserFeedBatchLoadsCacheMissesAndBackfillsFilteredRows(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })

	ctx := context.Background()
	const userID int64 = 42
	inboxKey := fmt.Sprintf(feed.FEED_INBOX_KEY, userID)
	for id := int64(1); id <= 6; id++ {
		requireNoError(t, redisClient.ZAdd(ctx, inboxKey, redis.Z{
			Score:  float64(1000 - id),
			Member: id,
		}).Err())
	}

	cached := &knowpost.FeedItem{Id: "2", CreatorId: 7, Title: "cached", Visible: "public"}
	raw, err := json.Marshal(cached)
	requireNoError(t, err)
	requireNoError(t, redisClient.Set(ctx, cachekeys.FeedItemKey(2), raw, time.Minute).Err())

	now := time.Now()
	loader := &recordingFeedPostLoader{rows: map[uint64]*model.KnowPosts{}}
	for id := uint64(3); id <= 6; id++ {
		loader.rows[id] = &model.KnowPosts{
			Id:          id,
			CreatorId:   7,
			Title:       sql.NullString{String: fmt.Sprintf("post-%d", id), Valid: true},
			ContentUrl:  sql.NullString{String: fmt.Sprintf("https://example/%d", id), Valid: true},
			Tags:        sql.NullString{String: `["go"]`, Valid: true},
			Visible:     "public",
			Status:      "published",
			PublishTime: sql.NullTime{Time: now, Valid: true},
		}
	}
	// ID 1 deliberately has no visible/published row and must be filtered out.

	serviceCtx := &svc.ServiceContext{
		Redis:          redisClient,
		FeedPostLoader: loader,
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			oneFollowingRelation{},
			zeroFollowerCounter{},
			logx.WithContext(ctx),
		),
	}

	resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Page:   2,
		Size:   2,
	})
	requireNoError(t, err)
	if len(resp.Items) != 2 || resp.Items[0].Id != "4" || resp.Items[1].Id != "5" {
		t.Fatalf("unexpected filtered page: %+v", resp.Items)
	}
	if !resp.HasMore {
		t.Fatal("expected another visible item after page 2")
	}
	if loader.calls != 1 {
		t.Fatalf("batch loader calls = %d, want 1", loader.calls)
	}
	wantMisses := []uint64{1, 3, 4, 5, 6}
	if fmt.Sprint(loader.requested) != fmt.Sprint(wantMisses) {
		t.Fatalf("batch misses = %v, want %v", loader.requested, wantMisses)
	}
	if resp.Items[0].ContentUrl != "https://example/4" || len(resp.Items[0].Tags) != 1 {
		t.Fatalf("feed item mapping incomplete: %+v", resp.Items[0])
	}
	if exists, err := redisClient.Exists(ctx, cachekeys.FeedItemKey(4)).Result(); err != nil || exists != 1 {
		t.Fatalf("DB-loaded feed item was not written back to Redis: exists=%d err=%v", exists, err)
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

type countingRelation struct {
	followings []int64
	calls      int
}

func (r *countingRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func (r *countingRelation) GetFollowings(context.Context, int64) ([]int64, error) {
	r.calls++
	return append([]int64(nil), r.followings...), nil
}

type countingCounter struct {
	calls int
}

func (c *countingCounter) BatchGetFollowerCounts(_ context.Context, userIDs []int64) (map[int64]int64, error) {
	c.calls++
	return make(map[int64]int64, len(userIDs)), nil
}

func TestGetUserFeedBackfillClassifiesFollowingsOnce(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })

	ctx := context.Background()
	const userID int64 = 52
	for id := int64(1); id <= 10; id++ {
		requireNoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
			Score:  float64(1000 - id),
			Member: id,
		}).Err())
	}
	now := time.Now()
	loader := &recordingFeedPostLoader{rows: map[uint64]*model.KnowPosts{
		9:  {Id: 9, CreatorId: 7, Visible: "public", Status: "published", PublishTime: sql.NullTime{Time: now, Valid: true}},
		10: {Id: 10, CreatorId: 7, Visible: "public", Status: "published", PublishTime: sql.NullTime{Time: now, Valid: true}},
	}}
	relation := &countingRelation{followings: []int64{7}}
	counter := &countingCounter{}
	serviceCtx := &svc.ServiceContext{
		Redis:          redisClient,
		FeedPostLoader: loader,
		FeedReader:     feed.NewFeedReader(feed.NewRedisAdapter(redisClient), relation, counter, logx.WithContext(ctx)),
	}

	resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Page:   1,
		Size:   2,
	})
	requireNoError(t, err)
	if len(resp.Items) != 2 {
		t.Fatalf("backfilled items = %d, want 2", len(resp.Items))
	}
	if relation.calls != 1 || counter.calls != 1 {
		t.Fatalf("classification calls relation=%d counter=%d, want 1 each", relation.calls, counter.calls)
	}
}

func TestGetUserFeedRejectsRevokedFollowerContentFromCacheAndDB(t *testing.T) {
	for _, cacheHit := range []bool{false, true} {
		t.Run(fmt.Sprintf("cache_hit_%t", cacheHit), func(t *testing.T) {
			redisServer := miniredis.RunT(t)
			redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
			t.Cleanup(func() { _ = redisClient.Close() })

			ctx := context.Background()
			const userID int64 = 62
			requireNoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
				Score:  1000,
				Member: 1,
			}).Err())

			loader := &recordingFeedPostLoader{rows: map[uint64]*model.KnowPosts{
				1: {Id: 1, CreatorId: 7, Visible: "followers", Status: "published"},
			}}
			if cacheHit {
				raw, err := json.Marshal(&knowpost.FeedItem{Id: "1", CreatorId: 7, Visible: "followers"})
				requireNoError(t, err)
				requireNoError(t, redisClient.Set(ctx, cachekeys.FeedItemKey(1), raw, time.Minute).Err())
			}
			serviceCtx := &svc.ServiceContext{
				Redis:          redisClient,
				FeedPostLoader: loader,
				FeedReader: feed.NewFeedReader(
					feed.NewRedisAdapter(redisClient),
					&countingRelation{}, // user no longer follows creator 7
					&countingCounter{},
					logx.WithContext(ctx),
				),
			}

			resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
				UserId: userID,
				Page:   1,
				Size:   1,
			})
			requireNoError(t, err)
			if len(resp.Items) != 0 {
				t.Fatalf("revoked follower content leaked: %+v", resp.Items)
			}
		})
	}
}

func TestPersonalFeedWindowRejectsUnboundedPage(t *testing.T) {
	if _, _, _, ok := personalFeedWindow(math.MaxInt32, 50); ok {
		t.Fatal("unbounded page should be rejected before allocating candidate maps")
	}
}
