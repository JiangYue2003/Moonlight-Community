//go:build integration

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestMutationRedisSnapshotRestoresExactZSetsAndBumpsSafetyEpoch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	client := redis.NewClient(&redis.Options{Addr: addr})
	t.Cleanup(func() { _ = client.Close() })
	require.NoError(t, client.Ping(ctx).Err())

	base := time.Now().UnixNano()%1_000_000_000 + 8_000_000_000
	authorID, readerID := base, base+1
	manifest := datasetManifest{
		Users:         []benchmarkUser{{ID: authorID}, {ID: readerID}},
		NormalAuthors: []int64{authorID}, Readers: []int64{readerID},
	}
	keys := mutationFeedKeys(manifest)
	t.Cleanup(func() { _ = client.Del(context.Background(), keys...).Err() })
	inboxKey := fmt.Sprintf("feed:inbox:%d", readerID)
	bigVKey := fmt.Sprintf("feed:bigv:%d", authorID)
	absentInboxKey := fmt.Sprintf("feed:inbox:%d", authorID)

	require.NoError(t, client.ZAdd(ctx, inboxKey,
		redis.Z{Score: 100, Member: "10"}, redis.Z{Score: 200, Member: "20"},
	).Err())
	require.NoError(t, client.PExpire(ctx, inboxKey, 10*time.Minute).Err())
	require.NoError(t, client.ZAdd(ctx, bigVKey, redis.Z{Score: 300, Member: "30"}).Err())

	snapshot, err := captureMutationRedisSnapshot(ctx, client, manifest)
	require.NoError(t, err)
	require.NotEmpty(t, snapshot.Identity)
	require.Len(t, snapshot.Keys, len(keys))

	require.NoError(t, client.ZAdd(ctx, inboxKey, redis.Z{Score: 400, Member: "40"}).Err())
	require.NoError(t, client.Del(ctx, bigVKey).Err())
	require.NoError(t, client.ZAdd(ctx, absentInboxKey, redis.Z{Score: 500, Member: "50"}).Err())
	beforeEpoch, err := readRedisInt64OrZero(ctx, client, "feed:content:safety:epoch")
	require.NoError(t, err)

	newEpoch, err := restoreMutationRedisSnapshot(ctx, client, snapshot)
	require.NoError(t, err)
	require.Greater(t, newEpoch, beforeEpoch)
	require.NoError(t, verifyMutationRedisSnapshot(ctx, client, snapshot, 5*time.Second))

	inbox, err := client.ZRangeWithScores(ctx, inboxKey, 0, -1).Result()
	require.NoError(t, err)
	require.Equal(t, []redis.Z{{Score: 100, Member: "10"}, {Score: 200, Member: "20"}}, inbox)
	bigV, err := client.ZRangeWithScores(ctx, bigVKey, 0, -1).Result()
	require.NoError(t, err)
	require.Equal(t, []redis.Z{{Score: 300, Member: "30"}}, bigV)
	require.Zero(t, client.Exists(ctx, absentInboxKey).Val())
	ttl := client.PTTL(ctx, inboxKey).Val()
	require.Greater(t, ttl, 9*time.Minute)
	require.LessOrEqual(t, ttl, 10*time.Minute)

	stored, err := client.Get(ctx, "feed:content:safety:epoch").Result()
	require.NoError(t, err)
	require.Equal(t, strconv.FormatInt(newEpoch, 10), stored)
}
