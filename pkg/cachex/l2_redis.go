package cachex

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// L2 是 Redis String 类型的薄包装。
// 与 ristretto 不同，L2 的 Set/Get/Del 都需要 ctx 与 error。
type L2 struct {
	rdb redis.UniversalClient
}

// NewL2 创建 L2 适配。
func NewL2(rdb redis.UniversalClient) *L2 {
	return &L2{rdb: rdb}
}

// Get 返回 (raw, hit, err)。键不存在时 hit=false 且 err=nil。
func (l *L2) Get(ctx context.Context, key string) ([]byte, bool, error) {
	raw, err := l.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return raw, true, nil
}

// Set 写入字节数据并设置 TTL。
func (l *L2) Set(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	return l.rdb.Set(ctx, key, val, ttl).Err()
}

// Del 删除一组 key（双删用）。
func (l *L2) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return l.rdb.Del(ctx, keys...).Err()
}
