// Package cachex 提供三级缓存抽象（L1 ristretto + L2 Redis + L3 业务回源）。
//
// 关键设计：
//   - GetOrLoad 沿 L1→L2→L3 链查；L3 由调用方提供 loader（不让 SQL 渗入 pkg）
//   - 单飞（SingleFlight）防击穿：同 sfKey 的并发只回源一次
//   - 空值哨兵防穿透：L3 返回 ErrNotFound 时缓存哨兵，TTL 较短
//   - TTL 抖动 + 热点延长：避免雪崩并自适应热点 key
package cachex

import (
	"context"
	"errors"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	"github.com/zhiguang/zhiguang-go/pkg/sfx"
)

// ErrNotFound 业务回源若返回此错误，cachex 会写入 null sentinel。
var ErrNotFound = errors.New("cachex: not found")

// Loader 业务侧 L3 回源函数；返回 ErrNotFound 触发空值缓存。
type Loader[V any] func(ctx context.Context) (V, error)

// Cache 三级缓存对外接口。
type Cache[V any] interface {
	GetOrLoad(ctx context.Context, key, sfKey string, loader Loader[V]) (V, error)
	Invalidate(ctx context.Context, keys ...string) error
	PutNullSentinel(ctx context.Context, key string) error
	// SetWithExtension 命中后根据 hot level 延长 TTL（写 L1 与 L2）。
	SetWithExtension(ctx context.Context, key string, val V, level hotkey.Level) error
}

// Options Cache 构造参数（业务方各自实例化）。
type Options[V any] struct {
	L1 *L1
	L2 *L2

	BaseTTL       time.Duration
	JitterMax     time.Duration
	NullTTL       time.Duration
	NullJitterMax time.Duration

	Hot         *hotkey.Detector
	HotKeyFor   func(cacheKey string) string
	TTLExtender func(base time.Duration, level hotkey.Level) time.Duration

	Marshal   func(V) ([]byte, error)
	Unmarshal func([]byte, *V) error
	// CostFn 估算 ristretto 缓存项的 cost；通常为 len(jsonBytes)。
	CostFn func(val V, raw []byte) int64
}

// New 构造一个 Cache 实例。
func New[V any](opts Options[V]) Cache[V] {
	if opts.BaseTTL <= 0 {
		opts.BaseTTL = 60 * time.Second
	}
	if opts.NullTTL <= 0 {
		opts.NullTTL = 30 * time.Second
	}
	if opts.TTLExtender == nil {
		opts.TTLExtender = hotkey.TTLForPublic
	}
	if opts.HotKeyFor == nil {
		opts.HotKeyFor = func(k string) string { return k }
	}
	if opts.CostFn == nil {
		opts.CostFn = func(_ V, raw []byte) int64 {
			if len(raw) == 0 {
				return 1
			}
			return int64(len(raw))
		}
	}
	return &cache[V]{opts: opts, sf: sfx.New[V]()}
}

type cache[V any] struct {
	opts Options[V]
	sf   *sfx.Group[V]
}

func (c *cache[V]) GetOrLoad(ctx context.Context, key, sfKey string, loader Loader[V]) (V, error) {
	var zero V

	// L1
	if v, ok := c.opts.L1.Get(key); ok {
		if isNullSentinel(v) {
			return zero, ErrNotFound
		}
		var out V
		if raw, ok := v.([]byte); ok {
			if err := c.opts.Unmarshal(raw, &out); err == nil {
				return out, nil
			}
		}
	}

	// L2
	raw, hit, err := c.opts.L2.Get(ctx, key)
	if err != nil {
		return zero, err
	}
	if hit {
		if IsNullSentinelBytes(raw) {
			c.opts.L1.SetWithTTL(key, NullSentinel(), 1, c.nullTTL())
			return zero, ErrNotFound
		}
		var out V
		if err := c.opts.Unmarshal(raw, &out); err == nil {
			c.opts.L1.SetWithTTL(key, raw, c.opts.CostFn(out, raw), c.opts.BaseTTL)
			return out, nil
		}
	}

	// L3 via singleflight
	return c.sf.DoCtx(ctx, sfKey, func(ctx context.Context) (V, error) {
		// 双检（同 sfKey 的等待者拿到锁后再次确认 L1/L2）
		if v, ok := c.opts.L1.Get(key); ok {
			if isNullSentinel(v) {
				return zero, ErrNotFound
			}
			if rawBytes, ok := v.([]byte); ok {
				var out V
				if err := c.opts.Unmarshal(rawBytes, &out); err == nil {
					return out, nil
				}
			}
		}
		if raw2, hit2, err := c.opts.L2.Get(ctx, key); err == nil && hit2 {
			if IsNullSentinelBytes(raw2) {
				return zero, ErrNotFound
			}
			var out V
			if err := c.opts.Unmarshal(raw2, &out); err == nil {
				c.opts.L1.SetWithTTL(key, raw2, c.opts.CostFn(out, raw2), c.opts.BaseTTL)
				return out, nil
			}
		}

		val, err := loader(ctx)
		if errors.Is(err, ErrNotFound) {
			_ = c.PutNullSentinel(ctx, key)
			return zero, ErrNotFound
		}
		if err != nil {
			return zero, err
		}
		raw, err := c.opts.Marshal(val)
		if err != nil {
			return zero, err
		}
		ttl := c.ttlWithJitter()
		_ = c.opts.L2.Set(ctx, key, raw, ttl)
		c.opts.L1.SetWithTTL(key, raw, c.opts.CostFn(val, raw), ttl)
		return val, nil
	})
}

func (c *cache[V]) Invalidate(ctx context.Context, keys ...string) error {
	// 写路径双删：先删 L2，再删 L1。
	if err := c.opts.L2.Del(ctx, keys...); err != nil {
		return err
	}
	for _, k := range keys {
		c.opts.L1.Del(k)
	}
	return nil
}

func (c *cache[V]) PutNullSentinel(ctx context.Context, key string) error {
	ttl := c.nullTTL()
	if err := c.opts.L2.Set(ctx, key, NullSentinel(), ttl); err != nil {
		return err
	}
	c.opts.L1.SetWithTTL(key, NullSentinel(), 1, ttl)
	return nil
}

func (c *cache[V]) SetWithExtension(ctx context.Context, key string, val V, level hotkey.Level) error {
	raw, err := c.opts.Marshal(val)
	if err != nil {
		return err
	}
	ttl := c.opts.TTLExtender(c.opts.BaseTTL, level)
	if err := c.opts.L2.Set(ctx, key, raw, ttl); err != nil {
		return err
	}
	c.opts.L1.SetWithTTL(key, raw, c.opts.CostFn(val, raw), ttl)
	return nil
}

func (c *cache[V]) ttlWithJitter() time.Duration {
	return Jitter(c.opts.BaseTTL, c.opts.JitterMax)
}

func (c *cache[V]) nullTTL() time.Duration {
	return Jitter(c.opts.NullTTL, c.opts.NullJitterMax)
}
