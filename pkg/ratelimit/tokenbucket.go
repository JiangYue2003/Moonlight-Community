package ratelimit

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TokenBucket 令牌桶限流器，多副本共享一份 Redis 桶状态。
type TokenBucket struct {
	rdb    redis.UniversalClient
	script *redis.Script
}

// New 构造 TokenBucket，预编译 Lua（首次执行 EVALSHA 失败时 go-redis 会自动 fallback EVAL）。
func New(rdb redis.UniversalClient) *TokenBucket {
	return &TokenBucket{
		rdb:    rdb,
		script: redis.NewScript(TokenBucketScript),
	}
}

// Take 尝试扣减 1 token；返回是否允许放行。
//
//	key            桶 key（推荐带 hash tag 防 cluster 跨槽，例如 "rl:follow:{userId}"）
//	capacity       桶容量
//	refillPerSec   每秒补充令牌数
//
// 任何 Redis 错误（连接、Lua 异常）都会被原样返回；调用方需根据策略选择"拒绝放行"或"放行降级"。
func (b *TokenBucket) Take(ctx context.Context, key string, capacity, refillPerSec int64) (bool, error) {
	now := time.Now().UnixMilli()
	res, err := b.script.Run(ctx, b.rdb,
		[]string{key}, capacity, refillPerSec, now).Int64()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}
