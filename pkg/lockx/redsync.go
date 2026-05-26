// Package lockx 提供基于 redsync 的分布式锁封装。
//
// 主要使用场景：
//   - counter-aggregator 多副本时只让一个 leader 周期性 flush
//   - 未来 counter-rpc.GetCounts 的 SDS 缺失重建时防止并发雪崩
package lockx

import (
	"context"
	"errors"
	"time"

	"github.com/go-redsync/redsync/v4"
	rsredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// Mutexer 由 New 构造；底层共享一个 redsync 实例。
type Mutexer struct {
	rs *redsync.Redsync
}

// New 用 go-redis/v9 客户端创建 Mutexer。
func New(rdb redis.UniversalClient) *Mutexer {
	pool := rsredis.NewPool(rdb)
	return &Mutexer{rs: redsync.New(pool)}
}

// TryAcquire 尝试获取锁，立即返回。
// ok=true 时返回的 release 必须被调用以释放锁；release 内部错误（锁已过期等）会被原样返回。
// ok=false 表示锁被他人持有，未发生错误。
func (m *Mutexer) TryAcquire(ctx context.Context, key string, ttl time.Duration) (release func() error, ok bool, err error) {
	mu := m.rs.NewMutex(
		key,
		redsync.WithExpiry(ttl),
		redsync.WithTries(1), // 不重试，立即让步
	)
	if err := mu.LockContext(ctx); err != nil {
		// redsync 在拿不到锁时返回 *ErrTaken / ErrFailed；不当作错误传出。
		var taken *redsync.ErrTaken
		if errors.As(err, &taken) || errors.Is(err, redsync.ErrFailed) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return func() error {
		_, err := mu.UnlockContext(ctx)
		return err
	}, true, nil
}
