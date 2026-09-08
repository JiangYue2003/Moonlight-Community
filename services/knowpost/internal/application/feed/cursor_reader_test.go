package feed

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/internal/application/feedcursor"
)

type scriptedCursorRedis struct {
	*MockRedisClient
	batches [][]ZRevRangeResult
	calls   int
}

type cancelingCursorRedis struct {
	*MockRedisClient
	cancel context.CancelFunc
	calls  int
}

func (r *cancelingCursorRedis) ZRevRangeByScoreWithScoresBatch(
	_ context.Context,
	requests []ZRevRangeByScoreRequest,
) []ZRevRangeResult {
	r.calls++
	if r.calls == 1 {
		r.cancel()
		return []ZRevRangeResult{
			{},
			{Scores: []ZScore{
				{Member: 9, Score: 99},
				{Member: 8, Score: 98},
				{Member: 7, Score: 97},
			}},
		}
	}
	return make([]ZRevRangeResult, len(requests))
}

func (r *scriptedCursorRedis) ZRevRangeByScoreWithScoresBatch(
	_ context.Context,
	requests []ZRevRangeByScoreRequest,
) []ZRevRangeResult {
	r.calls++
	if r.calls > len(r.batches) {
		return make([]ZRevRangeResult, len(requests))
	}
	return append([]ZRevRangeResult(nil), r.batches[r.calls-1]...)
}

func TestFeedReadSnapshotGetFeedPostsMatchesLegacyPage(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()

	const userID int64 = 81
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)
	for id, score := range map[int64]float64{1: 101, 2: 105, 3: 103, 4: 102} {
		require.NoError(t, client.ZAdd(ctx, key, goredis.Z{Score: score, Member: id}).Err())
	}

	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{7})
	reader := NewFeedReader(
		NewRedisAdapter(client),
		relation,
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	posts, postsHaveMore, err := snapshot.GetFeedPosts(ctx, 1, 3)
	require.NoError(t, err)
	ids, idsHaveMore, err := snapshot.GetFeed(ctx, 1, 3)
	require.NoError(t, err)

	require.Equal(t, []Post{
		{ID: 2, CreateTime: 105},
		{ID: 3, CreateTime: 103},
		{ID: 4, CreateTime: 102},
	}, posts)
	require.Equal(t, []int64{2, 3, 4}, ids)
	require.True(t, postsHaveMore)
	require.Equal(t, postsHaveMore, idsHaveMore)
}

func TestFeedReadSnapshotPageRecordsComparablePaginationWork(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()
	const userID int64 = 91
	for id := int64(1); id <= 4; id++ {
		require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_INBOX_KEY, userID), goredis.Z{
			Score: float64(100 + id), Member: id,
		}).Err())
	}
	observer := &recordingFeedObserver{}
	reader := NewFeedReaderWithOptions(
		NewRedisAdapter(client),
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(ctx),
		FeedReaderOptions{Strategy: StrategyHybrid, Observer: observer},
	)
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	posts, _, err := snapshot.GetFeedPosts(ctx, 1, 3)

	require.NoError(t, err)
	require.Len(t, posts, 3)
	require.Equal(t, 4, observerWorkTotal(observer.calls, CursorWorkRedisMembers))
	require.Equal(t, 1, observerWorkTotal(observer.calls, CursorWorkRedisCommands))
	require.Equal(t, 1, observerWorkTotal(observer.calls, CursorWorkRedisRoundTrips))
	require.Equal(t, 4, observerWorkTotal(observer.calls, CursorWorkMergeCandidates))
}

func observerWorkTotal(calls []observerCall, work FeedCursorWork) int {
	total := 0
	for _, call := range calls {
		if call.kind == "cursor_work" && call.work == work {
			total += call.value
		}
	}
	return total
}

func TestFeedReadSnapshotGetFeedAfterUsesNumericTimestampTieBreak(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()

	const userID int64 = 82
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)
	for _, post := range []goredis.Z{
		{Score: 100, Member: 11},
		{Score: 100, Member: 10},
		{Score: 100, Member: 9},
		{Score: 100, Member: 2},
		{Score: 99, Member: 20},
		{Score: 98, Member: 19},
	} {
		require.NoError(t, client.ZAdd(ctx, key, post).Err())
	}

	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{7})
	reader := NewFeedReader(
		NewRedisAdapter(client),
		relation,
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{
		SortTime: 100,
		PostID:   10,
	}, 3)

	require.NoError(t, err)
	require.Equal(t, []Post{
		{ID: 9, CreateTime: 100},
		{ID: 2, CreateTime: 100},
		{ID: 20, CreateTime: 99},
	}, posts)
	require.True(t, hasMore)
}

func TestFeedReadSnapshotGetFeedAfterCompletesOlderBoundaryTieGroup(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()

	const userID int64 = 83
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)
	for id := int64(1); id <= 30; id++ {
		require.NoError(t, client.ZAdd(ctx, key, goredis.Z{Score: 100, Member: id}).Err())
	}

	reader := NewFeedReader(
		NewRedisAdapter(client),
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{
		SortTime: 101,
		PostID:   999,
	}, 5)

	require.NoError(t, err)
	require.Equal(t, []Post{
		{ID: 30, CreateTime: 100},
		{ID: 29, CreateTime: 100},
		{ID: 28, CreateTime: 100},
		{ID: 27, CreateTime: 100},
		{ID: 26, CreateTime: 100},
	}, posts)
	require.True(t, hasMore)
}

func TestFeedReadSnapshotGetFeedAfterMergesAndDeduplicatesInboxAndBigVs(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()

	const userID int64 = 84
	for _, post := range []goredis.Z{
		{Score: 100, Member: 100},
		{Score: 95, Member: 95},
		{Score: 94, Member: 50},
	} {
		require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_INBOX_KEY, userID), post).Err())
	}
	for _, post := range []goredis.Z{
		{Score: 99, Member: 99},
		{Score: 94, Member: 94},
		{Score: 90, Member: 100},
	} {
		require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 7), post).Err())
	}
	for _, post := range []goredis.Z{
		{Score: 98, Member: 98},
		{Score: 97, Member: 97},
		{Score: 96, Member: 96},
	} {
		require.NoError(t, client.ZAdd(ctx, fmt.Sprintf(FEED_BIGV_OUTBOX_KEY, 8), post).Err())
	}

	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{7, 8})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(7, BIGV_THRESHOLD+1)
	counter.SetFollowerCount(8, BIGV_THRESHOLD+1)
	reader := NewFeedReader(NewRedisAdapter(client), relation, counter, logx.WithContext(ctx))
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{SortTime: 101, PostID: 999}, 5)

	require.NoError(t, err)
	require.Equal(t, []Post{
		{ID: 100, CreateTime: 100},
		{ID: 99, CreatorID: 7, CreateTime: 99},
		{ID: 98, CreatorID: 8, CreateTime: 98},
		{ID: 97, CreatorID: 8, CreateTime: 97},
		{ID: 96, CreatorID: 8, CreateTime: 96},
	}, posts)
	require.True(t, hasMore)
}

func TestFeedReadSnapshotGetFeedAfterFailsWhenCursorReaderIsUnavailable(t *testing.T) {
	ctx := context.Background()
	reader := NewFeedReader(
		NewMockRedisClient(),
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(ctx),
	)
	snapshot, err := reader.Prepare(ctx, 85)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{SortTime: 100, PostID: 10}, 5)

	require.ErrorIs(t, err, ErrCursorReaderUnavailable)
	require.Nil(t, posts)
	require.False(t, hasMore)
	_, _, pageErr := snapshot.GetFeed(ctx, 1, 5)
	require.NoError(t, pageErr)
}

func TestFeedReadSnapshotGetFeedAfterFailsWholePageOnRangeError(t *testing.T) {
	dependencyErr := errors.New("redis range failed")
	redisClient := &scriptedCursorRedis{
		MockRedisClient: NewMockRedisClient(),
		batches: [][]ZRevRangeResult{{
			{Scores: []ZScore{{Member: 9, Score: 100}}},
			{Err: dependencyErr},
		}},
	}
	ctx := context.Background()
	reader := NewFeedReader(redisClient, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(ctx))
	snapshot, err := reader.Prepare(ctx, 86)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{SortTime: 100, PostID: 10}, 5)

	require.ErrorIs(t, err, dependencyErr)
	require.Nil(t, posts)
	require.False(t, hasMore)
}

func TestFeedReadSnapshotGetFeedAfterFailsWholePageOnBoundaryRangeError(t *testing.T) {
	dependencyErr := errors.New("redis boundary failed")
	redisClient := &scriptedCursorRedis{
		MockRedisClient: NewMockRedisClient(),
		batches: [][]ZRevRangeResult{
			{
				{},
				{Scores: []ZScore{
					{Member: 9, Score: 99},
					{Member: 8, Score: 98},
					{Member: 7, Score: 97},
				}},
			},
			{{Err: dependencyErr}},
		},
	}
	ctx := context.Background()
	reader := NewFeedReader(redisClient, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(ctx))
	snapshot, err := reader.Prepare(ctx, 87)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{SortTime: 100, PostID: 10}, 2)

	require.ErrorIs(t, err, dependencyErr)
	require.Nil(t, posts)
	require.False(t, hasMore)
}

func TestFeedReadSnapshotGetFeedAfterStopsBeforeBoundaryBatchWhenCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	redisClient := &cancelingCursorRedis{MockRedisClient: NewMockRedisClient(), cancel: cancel}
	reader := NewFeedReader(redisClient, NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(ctx))
	snapshot, err := reader.Prepare(ctx, 88)
	require.NoError(t, err)

	posts, hasMore, err := snapshot.GetFeedAfter(ctx, feedcursor.Cursor{SortTime: 100, PostID: 10}, 2)

	require.ErrorIs(t, err, context.Canceled)
	require.Nil(t, posts)
	require.False(t, hasMore)
	require.Equal(t, 1, redisClient.calls)
}

func TestFeedReadSnapshotGetFeedAfterTraversesThousandPostTieGroupWithoutGaps(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()

	const userID int64 = 89
	key := fmt.Sprintf(FEED_INBOX_KEY, userID)
	entries := make([]goredis.Z, 0, INBOX_MAX_SIZE)
	for id := int64(1); id <= INBOX_MAX_SIZE; id++ {
		entries = append(entries, goredis.Z{Score: 100, Member: id})
	}
	require.NoError(t, client.ZAdd(ctx, key, entries...).Err())

	reader := NewFeedReader(NewRedisAdapter(client), NewMockRelationClient(), NewMockCounterClient(), logx.WithContext(ctx))
	snapshot, err := reader.Prepare(ctx, userID)
	require.NoError(t, err)

	after := feedcursor.Cursor{SortTime: 101, PostID: INBOX_MAX_SIZE + 1}
	got := make([]int64, 0, INBOX_MAX_SIZE)
	for {
		posts, hasMore, err := snapshot.GetFeedAfter(ctx, after, 20)
		require.NoError(t, err)
		for _, post := range posts {
			got = append(got, post.ID)
		}
		if !hasMore {
			break
		}
		require.NotEmpty(t, posts)
		last := posts[len(posts)-1]
		after = feedcursor.Cursor{SortTime: last.CreateTime, PostID: last.ID}
	}

	want := make([]int64, 0, INBOX_MAX_SIZE)
	for id := int64(INBOX_MAX_SIZE); id >= 1; id-- {
		want = append(want, id)
	}
	require.Equal(t, want, got)
}
