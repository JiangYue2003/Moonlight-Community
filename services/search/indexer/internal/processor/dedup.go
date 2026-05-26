package processor

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Dedup SETNX 去重，与 relation-syncer 同套思路。
type Dedup struct{ rdb goredis.UniversalClient }

func NewDedup(rdb goredis.UniversalClient) *Dedup { return &Dedup{rdb: rdb} }

func (d *Dedup) Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return d.rdb.SetNX(ctx, key, "1", ttl).Result()
}
