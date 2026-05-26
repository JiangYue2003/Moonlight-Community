package processor

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Dedup 用 SETNX 实现的事件去重。
//
// 与 Java RelationEventProcessor 一致：
//   - 同一 outbox.id 在 TTL 窗口内只允许处理一次
//   - SETNX 失败说明已被处理（或正在被另一副本处理），直接跳过
type Dedup struct {
	rdb goredis.UniversalClient
}

func NewDedup(rdb goredis.UniversalClient) *Dedup { return &Dedup{rdb: rdb} }

// Acquire 返回 true 表示当前调用是首次处理，应继续执行 handler；
// false 表示已被其它处理过/正在处理中，应跳过。
func (d *Dedup) Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return d.rdb.SetNX(ctx, key, "1", ttl).Result()
}
