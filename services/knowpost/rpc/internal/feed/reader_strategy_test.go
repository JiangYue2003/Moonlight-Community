package feed

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestFeedReaderPullStrategyReadsEveryFollowedAuthorOutbox(t *testing.T) {
	ctx := context.Background()
	server := miniredis.RunT(t)
	client := goredis.NewClient(&goredis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	redis := NewRedisAdapter(client)
	relation := NewMockRelationClient()
	relation.SetFollowings(100, []int64{201})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, 10)
	require.NoError(t, redis.ZAdd(ctx, "feed:bigv:201", float64(1234), int64(9001)))
	reader := NewFeedReaderWithStrategy(redis, relation, counter, logx.WithContext(ctx), StrategyPull)

	posts, _, err := reader.GetFeed(ctx, 100, 1, 20)

	require.NoError(t, err)
	require.Equal(t, []int64{9001}, posts)
	require.Zero(t, counter.batchCalls, "forced pull must not call Counter to reclassify authors")
}

func TestFeedReaderPushStrategySkipsOutboxClassification(t *testing.T) {
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, BIGV_THRESHOLD+1)
	reader := NewFeedReaderWithStrategy(
		NewMockRedisClient(),
		NewMockRelationClient(),
		counter,
		logx.WithContext(context.Background()),
		StrategyPush,
	)

	bigVs, normalUsers := reader.classifyFollowings(context.Background(), []int64{201})

	require.Empty(t, bigVs)
	require.Equal(t, []int64{201}, normalUsers)
	require.Zero(t, counter.batchCalls, "forced push must not call Counter or pull author outboxes")
}
