//go:build integration

package kafka

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/zset"
)

func TestRelationEpochConsumerIntegration(t *testing.T) {
	if os.Getenv("FEED_RELATION_EPOCH_INTEGRATION") != "1" {
		t.Skip("set FEED_RELATION_EPOCH_INTEGRATION=1 to exercise the dev stack")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		envOrDefault("RELATION_RPC_ADDR", "127.0.0.1:9006"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer conn.Close()
	relations := relationpb.NewRelationClient(conn)

	redisClient := goredis.NewClient(&goredis.Options{Addr: envOrDefault("REDIS_ADDR", "127.0.0.1:6379")})
	defer redisClient.Close()
	require.NoError(t, redisClient.Ping(ctx).Err())

	fromUserID := envInt64(t, "RELATION_FROM_USER_ID", 3124)
	toUserID := envInt64(t, "RELATION_TO_USER_ID", 3125)
	require.NotEqual(t, fromUserID, toUserID)

	status, err := relations.Status(ctx, &relationpb.StatusReq{FromUserId: fromUserID, ToUserId: toUserID})
	require.NoError(t, err)
	initialFollowing := status.GetFollowing()

	// Restore the original relation even if an intermediate assertion fails.
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		current, statusErr := relations.Status(cleanupCtx, &relationpb.StatusReq{FromUserId: fromUserID, ToUserId: toUserID})
		if statusErr != nil || current.GetFollowing() == initialFollowing {
			return
		}
		if initialFollowing {
			_, _ = relations.Follow(cleanupCtx, &relationpb.FollowReq{FromUserId: fromUserID, ToUserId: toUserID})
			return
		}
		_, _ = relations.Unfollow(cleanupCtx, &relationpb.UnfollowReq{FromUserId: fromUserID, ToUserId: toUserID})
	})

	epochKey := fmt.Sprintf("feed:relation:epoch:%d", fromUserID)
	before := readEpoch(t, ctx, redisClient, epochKey)

	if initialFollowing {
		resp, callErr := relations.Unfollow(ctx, &relationpb.UnfollowReq{FromUserId: fromUserID, ToUserId: toUserID})
		require.NoError(t, callErr)
		require.True(t, resp.GetChanged())
		waitForEpochAndZSet(t, ctx, redisClient, epochKey, before, fromUserID, toUserID, false)
	} else {
		resp, callErr := relations.Follow(ctx, &relationpb.FollowReq{FromUserId: fromUserID, ToUserId: toUserID})
		require.NoError(t, callErr)
		require.True(t, resp.GetChanged())
		waitForEpochAndZSet(t, ctx, redisClient, epochKey, before, fromUserID, toUserID, true)
	}

	middle := readEpoch(t, ctx, redisClient, epochKey)
	if initialFollowing {
		resp, callErr := relations.Follow(ctx, &relationpb.FollowReq{FromUserId: fromUserID, ToUserId: toUserID})
		require.NoError(t, callErr)
		require.True(t, resp.GetChanged())
		waitForEpochAndZSet(t, ctx, redisClient, epochKey, middle, fromUserID, toUserID, true)
	} else {
		resp, callErr := relations.Unfollow(ctx, &relationpb.UnfollowReq{FromUserId: fromUserID, ToUserId: toUserID})
		require.NoError(t, callErr)
		require.True(t, resp.GetChanged())
		waitForEpochAndZSet(t, ctx, redisClient, epochKey, middle, fromUserID, toUserID, false)
	}
}

func waitForEpochAndZSet(
	t *testing.T,
	ctx context.Context,
	client *goredis.Client,
	epochKey string,
	previous uint64,
	fromUserID, toUserID int64,
	wantFollowing bool,
) {
	t.Helper()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		epoch, epochErr := redisEpoch(ctx, client, epochKey)
		_, scoreErr := client.ZScore(ctx, zset.FollowingKey(fromUserID), strconv.FormatInt(toUserID, 10)).Result()
		following := scoreErr == nil
		if epochErr == nil && epoch > previous && following == wantFollowing {
			return
		}
		if scoreErr != nil && !errors.Is(scoreErr, goredis.Nil) {
			require.NoError(t, scoreErr)
		}
		select {
		case <-ctx.Done():
			t.Fatalf("wait for relation convergence: epoch=%d previous=%d following=%t want=%t: %v", epoch, previous, following, wantFollowing, ctx.Err())
		case <-ticker.C:
		}
	}
}

func readEpoch(t *testing.T, ctx context.Context, client *goredis.Client, key string) uint64 {
	t.Helper()
	epoch, err := redisEpoch(ctx, client, key)
	require.NoError(t, err)
	return epoch
}

func redisEpoch(ctx context.Context, client *goredis.Client, key string) (uint64, error) {
	value, err := client.Get(ctx, key).Result()
	if errors.Is(err, goredis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(value, 10, 64)
}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

func envInt64(t *testing.T, name string, fallback int64) int64 {
	t.Helper()
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	require.NoError(t, err)
	return parsed
}
