package feed

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisAdapterPushToInboxesWritesPostAndTTL(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	err := adapter.PushToInboxes(ctx, []int64{7, 8}, 99, 1234)

	require.NoError(t, err)
	for _, userID := range []int64{7, 8} {
		key := fmt.Sprintf(FEED_INBOX_KEY, userID)
		score, err := client.ZScore(ctx, key, "99").Result()
		require.NoError(t, err)
		require.Equal(t, float64(1234), score)
		require.Positive(t, client.TTL(ctx, key).Val())
	}
}

func TestRedisAdapterPushToInboxesKeepsNewestInboxEntries(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()
	key := fmt.Sprintf(FEED_INBOX_KEY, 7)
	oldPosts := make([]goredis.Z, INBOX_MAX_SIZE)
	for i := range oldPosts {
		oldPosts[i] = goredis.Z{Score: float64(i + 1), Member: int64(i + 1)}
	}
	require.NoError(t, client.ZAdd(ctx, key, oldPosts...).Err())

	err := adapter.PushToInboxes(ctx, []int64{7}, INBOX_MAX_SIZE+1, INBOX_MAX_SIZE+1)

	require.NoError(t, err)
	require.Equal(t, int64(INBOX_MAX_SIZE), client.ZCard(ctx, key).Val())
	require.ErrorIs(t, client.ZScore(ctx, key, "1").Err(), goredis.Nil)
	require.Equal(t, float64(INBOX_MAX_SIZE+1), client.ZScore(ctx, key, fmt.Sprint(INBOX_MAX_SIZE+1)).Val())
}

func TestRedisAdapterPushToInboxesIsIdempotent(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	require.NoError(t, adapter.PushToInboxes(ctx, []int64{7}, 99, 1234))
	require.NoError(t, adapter.PushToInboxes(ctx, []int64{7}, 99, 1234))

	require.Equal(t, int64(1), client.ZCard(ctx, fmt.Sprintf(FEED_INBOX_KEY, 7)).Val())
}

func TestRedisAdapterZRevRangeWithScoresBatchReturnsEveryRequestedRange(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	require.NoError(t, client.ZAdd(ctx, "outbox:1",
		goredis.Z{Score: 30, Member: 103},
		goredis.Z{Score: 20, Member: 102},
	).Err())
	require.NoError(t, client.ZAdd(ctx, "outbox:2",
		goredis.Z{Score: 40, Member: 204},
	).Err())

	results := adapter.ZRevRangeWithScoresBatch(ctx, []ZRevRangeRequest{
		{Key: "outbox:1", Start: 0, Stop: 0},
		{Key: "outbox:2", Start: 0, Stop: 0},
	})

	require.Len(t, results, 2)
	require.NoError(t, results[0].Err)
	require.Equal(t, []ZScore{{Member: 103, Score: 30}}, results[0].Scores)
	require.NoError(t, results[1].Err)
	require.Equal(t, []ZScore{{Member: 204, Score: 40}}, results[1].Scores)
}

func TestRedisAdapterZRevRangeWithScoresBatchKeepsSuccessfulRangesOnPartialFailure(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	require.NoError(t, client.ZAdd(ctx, "outbox:1", goredis.Z{Score: 30, Member: 103}).Err())
	require.NoError(t, client.Set(ctx, "outbox:broken", "not-a-zset", 0).Err())
	require.NoError(t, client.ZAdd(ctx, "outbox:2", goredis.Z{Score: 40, Member: 204}).Err())

	results := adapter.ZRevRangeWithScoresBatch(ctx, []ZRevRangeRequest{
		{Key: "outbox:1", Start: 0, Stop: 0},
		{Key: "outbox:broken", Start: 0, Stop: 0},
		{Key: "outbox:2", Start: 0, Stop: 0},
	})

	require.Len(t, results, 3)
	require.NoError(t, results[0].Err)
	require.Equal(t, []ZScore{{Member: 103, Score: 30}}, results[0].Scores)
	require.Error(t, results[1].Err)
	require.NoError(t, results[2].Err)
	require.Equal(t, []ZScore{{Member: 204, Score: 40}}, results[2].Scores)
}

func TestRedisAdapterZRevRangeByScoreWithScoresBatchReturnsExactAndExclusiveRanges(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	require.NoError(t, client.ZAdd(ctx, "feed:cursor",
		goredis.Z{Score: 100, Member: 9},
		goredis.Z{Score: 100, Member: 10},
		goredis.Z{Score: 99, Member: 11},
		goredis.Z{Score: 98, Member: 12},
		goredis.Z{Score: 97, Member: 13},
	).Err())

	results := adapter.ZRevRangeByScoreWithScoresBatch(ctx, []ZRevRangeByScoreRequest{
		{Key: "feed:cursor", Min: "100", Max: "100"},
		{Key: "feed:cursor", Min: "-inf", Max: "(100", Offset: 0, Count: 2},
	})

	require.Len(t, results, 2)
	require.NoError(t, results[0].Err)
	require.ElementsMatch(t, []ZScore{{Member: 9, Score: 100}, {Member: 10, Score: 100}}, results[0].Scores)
	require.NoError(t, results[1].Err)
	require.Equal(t, []ZScore{{Member: 11, Score: 99}, {Member: 12, Score: 98}}, results[1].Scores)
}

func TestRedisAdapterZRevRangeByScoreWithScoresBatchKeepsResultOffsetsOnFailure(t *testing.T) {
	mr := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	adapter := NewRedisAdapter(client)
	ctx := context.Background()

	require.NoError(t, client.ZAdd(ctx, "feed:one", goredis.Z{Score: 100, Member: 1}).Err())
	require.NoError(t, client.Set(ctx, "feed:broken", "not-a-zset", 0).Err())
	require.NoError(t, client.ZAdd(ctx, "feed:two", goredis.Z{Score: 90, Member: 2}).Err())

	results := adapter.ZRevRangeByScoreWithScoresBatch(ctx, []ZRevRangeByScoreRequest{
		{Key: "feed:one", Min: "-inf", Max: "+inf"},
		{Key: "feed:broken", Min: "-inf", Max: "+inf"},
		{Key: "feed:two", Min: "-inf", Max: "+inf"},
	})

	require.Len(t, results, 3)
	require.NoError(t, results[0].Err)
	require.Equal(t, []ZScore{{Member: 1, Score: 100}}, results[0].Scores)
	require.Error(t, results[1].Err)
	require.NoError(t, results[2].Err)
	require.Equal(t, []ZScore{{Member: 2, Score: 90}}, results[2].Scores)
}
