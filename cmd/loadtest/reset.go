package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/redis/go-redis/v9"
)

const redisDeleteBatchSize = 500

func feedKeysForManifest(manifest datasetManifest) []string {
	userIDs := make(map[int64]struct{}, len(manifest.Users)+len(manifest.NormalAuthors)+len(manifest.BigVAuthors))
	authorIDs := make(map[int64]struct{}, len(manifest.NormalAuthors)+len(manifest.BigVAuthors))
	keys := make(map[string]struct{})
	for _, user := range manifest.Users {
		userIDs[user.ID] = struct{}{}
	}
	for _, authors := range [][]int64{manifest.NormalAuthors, manifest.BigVAuthors} {
		for _, id := range authors {
			userIDs[id] = struct{}{}
			authorIDs[id] = struct{}{}
		}
	}
	for id := range userIDs {
		keys[fmt.Sprintf("feed:inbox:%d", id)] = struct{}{}
	}
	for id := range authorIDs {
		keys[fmt.Sprintf("feed:bigv:%d", id)] = struct{}{}
	}
	for _, id := range manifest.Posts {
		keys[fmt.Sprintf("feed:fanout:processing:%d", id)] = struct{}{}
	}

	result := make([]string, 0, len(keys))
	for key := range keys {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}

func resetFeedKeys(ctx context.Context, redisAddr string, manifest datasetManifest) (int64, error) {
	client := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer client.Close()
	if err := client.Ping(ctx).Err(); err != nil {
		return 0, fmt.Errorf("connect feed redis %s: %w", redisAddr, err)
	}

	keys := feedKeysForManifest(manifest)
	var deleted int64
	for start := 0; start < len(keys); start += redisDeleteBatchSize {
		end := minInt(start+redisDeleteBatchSize, len(keys))
		count, err := client.Del(ctx, keys[start:end]...).Result()
		if err != nil {
			return deleted, fmt.Errorf("delete feed key batch [%d:%d]: %w", start, end, err)
		}
		deleted += count
	}
	return deleted, nil
}
