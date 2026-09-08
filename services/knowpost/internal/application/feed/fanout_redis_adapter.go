package feed

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// PushToInboxes 在一次 Pipeline 往返中维护一批粉丝收件箱。
func (r *RedisAdapter) PushToInboxes(ctx context.Context, userIDs []int64, postID, timestamp int64) error {
	if len(userIDs) == 0 {
		return nil
	}

	pipe := r.client.Pipeline()
	for _, userID := range userIDs {
		key := fmt.Sprintf(FEED_INBOX_KEY, userID)
		pipe.ZAdd(ctx, key, redis.Z{Score: float64(timestamp), Member: postID})
		pipe.ZRemRangeByRank(ctx, key, 0, -INBOX_MAX_SIZE-1)
		pipe.Expire(ctx, key, INBOX_TTL_SECONDS*time.Second)
	}
	_, err := pipe.Exec(ctx)
	return err
}
