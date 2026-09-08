package knowpostlogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/internal/adapter/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/adapter/cache/userfeed"
	svc "github.com/zhiguang/zhiguang-go/services/knowpost/internal/application"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/feed"
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

type recordingPageCache struct {
	calls        int
	request      userfeed.Request
	page         *knowpost.FeedPage
	source       userfeed.Source
	err          error
	invokeLoader bool
	loaderCtx    context.Context
	beforeReturn func()
}

func (c *recordingPageCache) GetOrLoad(
	_ context.Context,
	req userfeed.Request,
	loader userfeed.Loader,
) (*knowpost.FeedPage, userfeed.Source, error) {
	c.calls++
	c.request = req
	if c.invokeLoader {
		page, err := loader(c.loaderCtx)
		return page, userfeed.SourceLoad, err
	}
	if c.beforeReturn != nil {
		c.beforeReturn()
	}
	return c.page, c.source, c.err
}

func (*recordingPageCache) Close() error { return nil }

type stubFeedEpochs struct {
	relation      uint64
	safety        uint64
	relationErr   error
	safetyErr     error
	flushErr      error
	pending       bool
	relationCalls int
	safetyCalls   int
	flushCalls    int
}

func (s *stubFeedEpochs) Relation(context.Context, int64) (uint64, error) {
	s.relationCalls++
	return s.relation, s.relationErr
}

func (s *stubFeedEpochs) BumpRelation(context.Context, int64) (uint64, error) {
	s.relation++
	return s.relation, nil
}

func (s *stubFeedEpochs) Safety(context.Context) (uint64, error) {
	s.safetyCalls++
	return s.safety, s.safetyErr
}

func (s *stubFeedEpochs) BumpSafety(context.Context) (uint64, error) {
	s.safety++
	return s.safety, nil
}
func (s *stubFeedEpochs) MarkSafetyPending()  { s.pending = true }
func (s *stubFeedEpochs) SafetyPending() bool { return s.pending }
func (s *stubFeedEpochs) FlushPendingSafety(context.Context) error {
	s.flushCalls++
	if s.flushErr != nil {
		return s.flushErr
	}
	if s.pending {
		s.safety++
	}
	s.pending = false
	return nil
}

func TestGetUserFeedPageCacheHitUsesEpochKeyWithoutColdCompute(t *testing.T) {
	observer := &logicRecordingObserver{}
	pageCache := &recordingPageCache{
		page:   &knowpost.FeedPage{Page: 1, Size: 20, Items: []*knowpost.FeedItem{{Id: "cached"}}},
		source: userfeed.SourceL1Fresh,
	}
	epochs := &stubFeedEpochs{relation: 7, safety: 9}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			Strategy:  "hybrid",
			PageCache: config.FeedPageCacheConf{Mode: "l1-l2"},
		}},
		FeedEpochs:    epochs,
		FeedPageCache: pageCache,
		FeedObserver:  observer,
	}

	got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   1,
		Size:   20,
	})
	requireNoError(t, err)
	if len(got.Items) != 1 || got.Items[0].Id != "cached" {
		t.Fatalf("unexpected cached response: %+v", got)
	}
	if pageCache.calls != 1 {
		t.Fatalf("page cache calls = %d, want 1", pageCache.calls)
	}
	wantReq := userfeed.Request{
		UserID:          42,
		StrategyVersion: userfeed.StrategyHybridV1,
		RelationEpoch:   7,
		SafetyEpoch:     9,
		Page:            1,
		Size:            20,
	}
	if pageCache.request != wantReq {
		t.Fatalf("page cache request = %+v, want %+v", pageCache.request, wantReq)
	}
	if epochs.relationCalls != 2 || epochs.safetyCalls != 2 {
		t.Fatalf("epoch calls relation=%d safety=%d, want 2 each for pre/post lookup validation", epochs.relationCalls, epochs.safetyCalls)
	}
	if !logicHasPageCache(observer.calls, feed.PageCacheL1Fresh, feed.OutcomeSuccess) {
		t.Fatalf("missing page cache hit metric: %+v", observer.calls)
	}
}

func TestGetUserFeedDiscardsCachedPageWhenEpochStateChangesDuringLookup(t *testing.T) {
	for _, tc := range []struct {
		name   string
		source userfeed.Source
		change func(*stubFeedEpochs)
	}{
		{name: "safety pending fresh", source: userfeed.SourceL1Fresh, change: func(epochs *stubFeedEpochs) { epochs.MarkSafetyPending() }},
		{name: "safety pending stale", source: userfeed.SourceL2Stale, change: func(epochs *stubFeedEpochs) { epochs.MarkSafetyPending() }},
		{name: "relation bump", source: userfeed.SourceL1Fresh, change: func(epochs *stubFeedEpochs) { _, _ = epochs.BumpRelation(context.Background(), 42) }},
		{name: "safety bump", source: userfeed.SourceL1Fresh, change: func(epochs *stubFeedEpochs) { _, _ = epochs.BumpSafety(context.Background()) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			epochs := &stubFeedEpochs{relation: 7, safety: 9}
			observer := &logicRecordingObserver{}
			pageCache := &recordingPageCache{
				page:   &knowpost.FeedPage{Page: 1, Size: 20, Items: []*knowpost.FeedItem{{Id: "stale-sentinel"}}},
				source: tc.source,
			}
			pageCache.beforeReturn = func() { tc.change(epochs) }
			serviceCtx := emptyFeedServiceContext(t)
			serviceCtx.Config.Feed.Strategy = "hybrid"
			serviceCtx.Config.Feed.PageCache = config.FeedPageCacheConf{Mode: "l1-l2"}
			serviceCtx.FeedEpochs = epochs
			serviceCtx.FeedPageCache = pageCache
			serviceCtx.FeedObserver = observer

			got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
				UserId: 42, Page: 1, Size: 20,
			})
			requireNoError(t, err)
			if got == nil || len(got.Items) != 0 {
				t.Fatalf("stale cached page was returned after %s: %+v", tc.name, got)
			}
			if pageCache.calls != 1 {
				t.Fatalf("page cache calls = %d, want 1", pageCache.calls)
			}
			if !logicHasPageCache(observer.calls, feed.PageCacheBypass, feed.OutcomeSuccess) {
				t.Fatalf("missing post-lookup bypass metric: %+v", observer.calls)
			}
			if logicHasPageCache(observer.calls, pageCacheMetricSource(tc.source), feed.OutcomeSuccess) {
				t.Fatalf("discarded hit was recorded as served: %+v", observer.calls)
			}
		})
	}
}

func TestGetUserFeedBypassesPageCacheWhenEpochCannotBeTrusted(t *testing.T) {
	for _, tc := range []struct {
		name              string
		epochs            *stubFeedEpochs
		wantRelationCalls int
		wantSafetyCalls   int
		wantFlushCalls    int
	}{
		{name: "safety pending", epochs: &stubFeedEpochs{pending: true, flushErr: errors.New("safety flush")}, wantRelationCalls: 0, wantSafetyCalls: 0, wantFlushCalls: 1},
		{name: "relation error", epochs: &stubFeedEpochs{relationErr: errors.New("relation epoch")}, wantRelationCalls: 1, wantSafetyCalls: 0},
		{name: "safety error", epochs: &stubFeedEpochs{relation: 1, safetyErr: errors.New("safety epoch")}, wantRelationCalls: 1, wantSafetyCalls: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pageCache := &recordingPageCache{page: &knowpost.FeedPage{Page: 1, Size: 20}}
			serviceCtx := emptyFeedServiceContext(t)
			serviceCtx.Config.Feed.Strategy = "hybrid"
			serviceCtx.Config.Feed.PageCache = config.FeedPageCacheConf{Mode: "l2"}
			serviceCtx.FeedEpochs = tc.epochs
			serviceCtx.FeedPageCache = pageCache

			got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
				UserId: 42,
				Page:   1,
				Size:   20,
			})
			requireNoError(t, err)
			if got == nil || len(got.Items) != 0 {
				t.Fatalf("cold fallback response = %+v", got)
			}
			if pageCache.calls != 0 {
				t.Fatalf("untrusted epoch reached page cache: calls=%d", pageCache.calls)
			}
			if tc.epochs.relationCalls != tc.wantRelationCalls || tc.epochs.safetyCalls != tc.wantSafetyCalls {
				t.Fatalf("epoch calls relation=%d safety=%d, want %d/%d",
					tc.epochs.relationCalls, tc.epochs.safetyCalls, tc.wantRelationCalls, tc.wantSafetyCalls)
			}
			if tc.epochs.flushCalls != tc.wantFlushCalls {
				t.Fatalf("flush calls=%d, want %d", tc.epochs.flushCalls, tc.wantFlushCalls)
			}
		})
	}
}

func TestGetUserFeedFlushesPendingSafetyBeforeUsingPageCache(t *testing.T) {
	epochs := &stubFeedEpochs{relation: 7, safety: 9, pending: true}
	pageCache := &recordingPageCache{
		page:   &knowpost.FeedPage{Page: 1, Size: 20, Items: []*knowpost.FeedItem{{Id: "fresh-after-flush"}}},
		source: userfeed.SourceL2Fresh,
	}
	serviceCtx := emptyFeedServiceContext(t)
	serviceCtx.Config.Feed.Strategy = "hybrid"
	serviceCtx.Config.Feed.PageCache = config.FeedPageCacheConf{Mode: "l2"}
	serviceCtx.FeedEpochs = epochs
	serviceCtx.FeedPageCache = pageCache

	got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   1,
		Size:   20,
	})
	requireNoError(t, err)
	if len(got.Items) != 1 || got.Items[0].Id != "fresh-after-flush" {
		t.Fatalf("unexpected response after safety flush: %+v", got)
	}
	if epochs.flushCalls != 1 || epochs.pending {
		t.Fatalf("flush calls=%d pending=%t, want 1/false", epochs.flushCalls, epochs.pending)
	}
	if pageCache.calls != 1 || pageCache.request.SafetyEpoch != 10 {
		t.Fatalf("page cache calls=%d safety epoch=%d, want 1/10", pageCache.calls, pageCache.request.SafetyEpoch)
	}
}

func TestGetUserFeedSafetyEpochDoesNotReuseUnversionedFeedItem(t *testing.T) {
	serviceCtx, epochs, pageCache := staleFeedItemSafetyFixture(t)

	got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   1,
		Size:   20,
	})
	requireNoError(t, err)
	if len(got.Items) != 0 {
		t.Fatalf("old unversioned FeedItem leaked into new safety generation: %+v", got.Items)
	}
	if pageCache.calls != 1 || pageCache.request.SafetyEpoch != epochs.safety {
		t.Fatalf("page cache calls=%d safety=%d, want 1/%d", pageCache.calls, pageCache.request.SafetyEpoch, epochs.safety)
	}
}

func TestGetUserFeedSafetyPendingBypassDoesNotTrustFeedItemCache(t *testing.T) {
	serviceCtx, epochs, pageCache := staleFeedItemSafetyFixture(t)
	epochs.pending = true
	epochs.flushErr = errors.New("redis unavailable")
	oldItem, err := json.Marshal(&knowpost.FeedItem{Id: "1", CreatorId: 7, Visible: "public"})
	requireNoError(t, err)
	requireNoError(t, serviceCtx.Redis.Set(
		context.Background(), cachekeys.PersonalFeedItemKey(1, epochs.safety), oldItem, time.Minute,
	).Err())

	got, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   1,
		Size:   20,
	})
	requireNoError(t, err)
	if len(got.Items) != 0 {
		t.Fatalf("old FeedItem leaked while safety state was pending: %+v", got.Items)
	}
	if epochs.flushCalls != 1 || pageCache.calls != 0 {
		t.Fatalf("flush calls=%d page cache calls=%d, want 1/0", epochs.flushCalls, pageCache.calls)
	}
}

func staleFeedItemSafetyFixture(t *testing.T) (*svc.ServiceContext, *stubFeedEpochs, *recordingPageCache) {
	t.Helper()
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })

	ctx := context.Background()
	const (
		userID int64 = 42
		postID int64 = 1
	)
	requireNoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
		Score:  1000,
		Member: postID,
	}).Err())
	oldItem, err := json.Marshal(&knowpost.FeedItem{Id: "1", CreatorId: 7, Visible: "public"})
	requireNoError(t, err)
	requireNoError(t, redisClient.Set(ctx, cachekeys.FeedItemKey(postID), oldItem, time.Minute).Err())
	requireNoError(t, redisClient.Set(ctx, cachekeys.PersonalFeedItemKey(postID, 8), oldItem, time.Minute).Err())

	epochs := &stubFeedEpochs{relation: 7, safety: 9}
	pageCache := &recordingPageCache{invokeLoader: true, loaderCtx: ctx}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			Strategy:  "hybrid",
			PageCache: config.FeedPageCacheConf{Mode: "l2"},
		}},
		Redis:          redisClient,
		FeedPostLoader: &recordingFeedPostLoader{rows: map[uint64]*model.KnowPosts{}},
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			oneFollowingRelation{},
			zeroFollowerCounter{},
			logx.WithContext(ctx),
		),
		FeedEpochs:    epochs,
		FeedPageCache: pageCache,
	}
	return serviceCtx, epochs, pageCache
}

func TestGetUserFeedIneligibleRequestSkipsEpochAndPageCache(t *testing.T) {
	for _, tc := range []struct {
		name     string
		strategy string
		page     int32
		size     int32
	}{
		{name: "pull", strategy: "pull", page: 1, size: 20},
		{name: "page two", strategy: "hybrid", page: 2, size: 20},
		{name: "size ten", strategy: "hybrid", page: 1, size: 10},
	} {
		t.Run(tc.name, func(t *testing.T) {
			epochs := &stubFeedEpochs{}
			pageCache := &recordingPageCache{}
			serviceCtx := emptyFeedServiceContext(t)
			serviceCtx.Config.Feed.Strategy = tc.strategy
			serviceCtx.Config.Feed.PageCache = config.FeedPageCacheConf{Mode: "l1-l2"}
			serviceCtx.FeedEpochs = epochs
			serviceCtx.FeedPageCache = pageCache

			_, err := NewGetUserFeedLogic(context.Background(), serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
				UserId: 42,
				Page:   tc.page,
				Size:   tc.size,
			})
			requireNoError(t, err)
			if pageCache.calls != 0 || epochs.relationCalls != 0 || epochs.safetyCalls != 0 {
				t.Fatalf("ineligible request touched cache: page=%d relation=%d safety=%d",
					pageCache.calls, epochs.relationCalls, epochs.safetyCalls)
			}
		})
	}
}

type contextKey string

type contextRecordingFeedPostLoader struct {
	row       *model.KnowPosts
	seenValue any
}

func (l *contextRecordingFeedPostLoader) FindPublishedFeedByIDs(ctx context.Context, _ []uint64) ([]*model.KnowPosts, error) {
	l.seenValue = ctx.Value(contextKey("page-cache-loader"))
	return []*model.KnowPosts{l.row}, nil
}

func TestGetUserFeedColdLoaderUsesContextProvidedByPageCache(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()
	requireNoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, int64(42)), redis.Z{
		Score: 1000, Member: 1,
	}).Err())

	loaderCtx := context.WithValue(context.Background(), contextKey("page-cache-loader"), "cache-context")
	pageCache := &recordingPageCache{invokeLoader: true, loaderCtx: loaderCtx}
	epochs := &stubFeedEpochs{relation: 7, safety: 9}
	postLoader := &contextRecordingFeedPostLoader{row: &model.KnowPosts{
		Id: 1, CreatorId: 7, Visible: "public", Status: "published",
	}}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			Strategy:  "hybrid",
			PageCache: config.FeedPageCacheConf{Mode: "l1-l2"},
		}},
		Redis:          redisClient,
		FeedPostLoader: postLoader,
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient), oneFollowingRelation{}, zeroFollowerCounter{}, logx.WithContext(ctx),
		),
		FeedEpochs:    epochs,
		FeedPageCache: pageCache,
	}

	got, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{UserId: 42, Page: 1, Size: 20})
	requireNoError(t, err)
	if len(got.Items) != 1 {
		t.Fatalf("items = %d, want 1", len(got.Items))
	}
	if postLoader.seenValue != "cache-context" {
		t.Fatalf("post loader context value = %v, want cache-context", postLoader.seenValue)
	}
}

func emptyFeedServiceContext(t *testing.T) *svc.ServiceContext {
	t.Helper()
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()
	return &svc.ServiceContext{
		Redis: redisClient,
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient), oneFollowingRelation{}, zeroFollowerCounter{}, logx.WithContext(ctx),
		),
	}
}
