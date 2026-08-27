package knowpostlogic

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feedcursor"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUserFeedRejectsCursorBeforeDependenciesWhenDisabledOrInvalid(t *testing.T) {
	for _, test := range []struct {
		name    string
		enabled bool
		cursor  string
	}{
		{name: "disabled", enabled: false, cursor: "opaque-cursor"},
		{name: "invalid", enabled: true, cursor: "not-base64url-json"},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			redisServer := miniredis.RunT(t)
			redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
			t.Cleanup(func() { _ = redisClient.Close() })
			relation := &countingRelation{}
			serviceCtx := &svc.ServiceContext{
				Config: config.Config{Feed: config.FeedConf{
					CursorPagination: config.FeedCursorPaginationConf{Enabled: test.enabled},
				}},
				FeedReader: feed.NewFeedReader(
					feed.NewRedisAdapter(redisClient),
					relation,
					&countingCounter{},
					logx.WithContext(ctx),
				),
			}

			response, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
				UserId: 42,
				Page:   9,
				Size:   20,
				Cursor: test.cursor,
			})

			require.Nil(t, response)
			require.Equal(t, codes.InvalidArgument, status.Code(err))
			require.Zero(t, relation.calls)
		})
	}
}

func TestGetUserFeedReturnsAndConsumesExactCursor(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()

	const userID int64 = 90
	rows := make(map[uint64]*model.KnowPosts)
	for id := int64(1); id <= 6; id++ {
		require.NoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
			Score:  float64(1000 - id),
			Member: id,
		}).Err())
		rows[uint64(id)] = &model.KnowPosts{
			Id:          uint64(id),
			CreatorId:   7,
			Title:       sql.NullString{String: fmt.Sprintf("post-%d", id), Valid: true},
			Visible:     "public",
			Status:      "published",
			PublishTime: sql.NullTime{Time: time.Unix(1000-id, 0), Valid: true},
		}
	}
	relation := &countingRelation{followings: []int64{7}}
	observer := &logicRecordingObserver{}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			CursorPagination: config.FeedCursorPaginationConf{Enabled: true},
		}},
		Redis:          redisClient,
		FeedPostLoader: &recordingFeedPostLoader{rows: rows},
		FeedObserver:   observer,
		FeedReader: feed.NewFeedReaderWithOptions(
			feed.NewRedisAdapter(redisClient),
			relation,
			&countingCounter{},
			logx.WithContext(ctx),
			feed.FeedReaderOptions{Strategy: feed.StrategyHybrid, Observer: observer},
		),
	}
	logic := NewGetUserFeedLogic(ctx, serviceCtx)

	first, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Page: 1, Size: 2})
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2"}, feedItemIDs(first.Items))
	require.True(t, first.HasMore)
	require.NotEmpty(t, first.NextCursor)
	firstCursor, err := feedcursor.Decode(first.NextCursor)
	require.NoError(t, err)
	require.Equal(t, feedcursor.Cursor{SortTime: 998, PostID: 2}, firstCursor)

	second, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Page:   99,
		Size:   2,
		Cursor: first.NextCursor,
	})
	require.NoError(t, err)
	require.Equal(t, []string{"3", "4"}, feedItemIDs(second.Items))
	require.True(t, second.HasMore)
	require.Equal(t, int32(0), second.Page)
	require.Equal(t, int32(2), second.Size)
	secondCursor, err := feedcursor.Decode(second.NextCursor)
	require.NoError(t, err)
	require.Equal(t, feedcursor.Cursor{SortTime: 996, PostID: 4}, secondCursor)
	require.Equal(t, 2, relation.calls, "each RPC request should prepare its route once")
	require.Equal(t, 1, countPaginationCalls(observer.calls, feed.PaginationModePage, feed.PaginationResultSuccess))
	require.Equal(t, 1, countPaginationCalls(observer.calls, feed.PaginationModeCursor, feed.PaginationResultSuccess))
	require.Positive(t, sumCursorWork(observer.calls, feed.CursorWorkRedisMembers))
	require.Positive(t, sumCursorWork(observer.calls, feed.CursorWorkHydrateIDs))
}

func TestGetUserFeedCursorBackfillsFilteredRowsAndUsesLastVisiblePost(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()

	const userID int64 = 91
	for id := int64(1); id <= 8; id++ {
		require.NoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
			Score:  float64(1000 - id),
			Member: id,
		}).Err())
	}
	rows := make(map[uint64]*model.KnowPosts)
	for id := int64(5); id <= 7; id++ {
		rows[uint64(id)] = &model.KnowPosts{
			Id:          uint64(id),
			CreatorId:   7,
			Title:       sql.NullString{String: fmt.Sprintf("post-%d", id), Valid: true},
			Visible:     "public",
			Status:      "published",
			PublishTime: sql.NullTime{Time: time.Unix(1000-id, 0), Valid: true},
		}
	}
	relation := &countingRelation{followings: []int64{7}}
	counter := &countingCounter{}
	loader := &recordingFeedPostLoader{rows: rows}
	pageCache := &recordingPageCache{}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			CursorPagination: config.FeedCursorPaginationConf{Enabled: true},
			PageCache:        config.FeedPageCacheConf{Mode: "l2"},
		}},
		Redis:          redisClient,
		FeedPostLoader: loader,
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			relation,
			counter,
			logx.WithContext(ctx),
		),
		FeedPageCache: pageCache,
	}
	start, err := feedcursor.Encode(feedcursor.Cursor{SortTime: 1001, PostID: 999})
	require.NoError(t, err)

	first, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Size:   2,
		Cursor: start,
	})
	require.NoError(t, err)
	require.Equal(t, []string{"5", "6"}, feedItemIDs(first.Items))
	require.True(t, first.HasMore)
	firstCursor, err := feedcursor.Decode(first.NextCursor)
	require.NoError(t, err)
	require.Equal(t, feedcursor.Cursor{SortTime: 994, PostID: 6}, firstCursor)
	require.Equal(t, 1, relation.calls, "backfill must reuse one route snapshot")
	require.Equal(t, 1, counter.calls, "backfill must not reclassify followings")
	require.Equal(t, 3, loader.calls, "only newly discovered candidates should be hydrated")

	last, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Size:   2,
		Cursor: first.NextCursor,
	})
	require.NoError(t, err)
	require.Equal(t, []string{"7"}, feedItemIDs(last.Items))
	require.False(t, last.HasMore)
	require.Empty(t, last.NextCursor)
	require.Zero(t, pageCache.calls, "cursor requests must bypass the full-page cache")
}

func TestGetUserFeedCursorWindowIgnoresNewerInsertAndRereadSeesIt(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()
	const userID int64 = 92

	rows := make(map[uint64]*model.KnowPosts)
	for id := int64(1); id <= 5; id++ {
		require.NoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
			Score: float64(1000 - id), Member: id,
		}).Err())
		rows[uint64(id)] = publicCursorTestRow(id, 7, 1000-id)
	}
	serviceCtx := cursorTestServiceContext(ctx, redisClient, userID, []int64{7}, rows)
	logic := NewGetUserFeedLogic(ctx, serviceCtx)

	first, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Page: 1, Size: 2})
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2"}, feedItemIDs(first.Items))

	rows[100] = publicCursorTestRow(100, 7, 1100)
	require.NoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
		Score: 1100, Member: 100,
	}).Err())

	next, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Size: 2, Cursor: first.NextCursor})
	require.NoError(t, err)
	require.Equal(t, []string{"3", "4"}, feedItemIDs(next.Items), "newer insert must not enter an existing cursor window")

	refreshed, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Page: 1, Size: 2})
	require.NoError(t, err)
	require.Equal(t, []string{"100", "1"}, feedItemIDs(refreshed.Items), "a new first-page read must see the insert")
}

func TestGetUserFeedCursorUsesCurrentFollowingRelationship(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	ctx := context.Background()
	const userID int64 = 93

	rows := make(map[uint64]*model.KnowPosts)
	for id := int64(1); id <= 4; id++ {
		require.NoError(t, redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
			Score: float64(1000 - id), Member: id,
		}).Err())
		rows[uint64(id)] = publicCursorTestRow(id, 7, 1000-id)
	}
	relation := &countingRelation{followings: []int64{7}}
	serviceCtx := &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			CursorPagination: config.FeedCursorPaginationConf{Enabled: true},
		}},
		Redis:          redisClient,
		FeedPostLoader: &recordingFeedPostLoader{rows: rows},
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient), relation, &countingCounter{}, logx.WithContext(ctx),
		),
	}
	logic := NewGetUserFeedLogic(ctx, serviceCtx)

	first, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Page: 1, Size: 1})
	require.NoError(t, err)
	require.Equal(t, []string{"1"}, feedItemIDs(first.Items))
	relation.followings = nil

	next, err := logic.GetUserFeed(&knowpost.GetUserFeedReq{UserId: userID, Size: 2, Cursor: first.NextCursor})
	require.NoError(t, err)
	require.Empty(t, next.Items, "an old cursor must not retain access to an unfollowed author")
	require.False(t, next.HasMore)
	require.Empty(t, next.NextCursor)
}

func cursorTestServiceContext(
	ctx context.Context,
	redisClient *redis.Client,
	userID int64,
	followings []int64,
	rows map[uint64]*model.KnowPosts,
) *svc.ServiceContext {
	return &svc.ServiceContext{
		Config: config.Config{Feed: config.FeedConf{
			CursorPagination: config.FeedCursorPaginationConf{Enabled: true},
		}},
		Redis:          redisClient,
		FeedPostLoader: &recordingFeedPostLoader{rows: rows},
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			&countingRelation{followings: followings},
			&countingCounter{},
			logx.WithContext(ctx),
		),
	}
}

func publicCursorTestRow(id, creatorID, publishTime int64) *model.KnowPosts {
	return &model.KnowPosts{
		Id:          uint64(id),
		CreatorId:   uint64(creatorID),
		Title:       sql.NullString{String: fmt.Sprintf("post-%d", id), Valid: true},
		Visible:     "public",
		Status:      "published",
		PublishTime: sql.NullTime{Time: time.Unix(publishTime, 0), Valid: true},
	}
}

func feedItemIDs(items []*knowpost.FeedItem) []string {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.Id
	}
	return ids
}

func countPaginationCalls(
	calls []logicObserverCall,
	mode feed.FeedPaginationMode,
	result feed.FeedPaginationResult,
) int {
	count := 0
	for _, call := range calls {
		if call.kind == "pagination" && call.mode == mode && call.pageResult == result {
			count++
		}
	}
	return count
}

func sumCursorWork(calls []logicObserverCall, work feed.FeedCursorWork) int {
	total := 0
	for _, call := range calls {
		if call.kind == "cursor_work" && call.work == work {
			total += call.value
		}
	}
	return total
}
