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
