package cachex

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

// L1 是 ristretto 的薄包装。
type L1 struct {
	c *ristretto.Cache
}

// L1Config ristretto 实例参数。
type L1Config struct {
	NumCounters int64 // 通常 = 10x 预期 entries
	MaxCost     int64 // 容量上限（字节）
	BufferItems int64 // 默认 64
}

// NewL1 构造一个 ristretto 实例。
func NewL1(cfg L1Config) (*L1, error) {
	if cfg.NumCounters == 0 {
		cfg.NumCounters = 50000
	}
	if cfg.MaxCost == 0 {
		cfg.MaxCost = 100 << 20
	}
	if cfg.BufferItems == 0 {
		cfg.BufferItems = 64
	}
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: cfg.NumCounters,
		MaxCost:     cfg.MaxCost,
		BufferItems: cfg.BufferItems,
	})
	if err != nil {
		return nil, err
	}
	return &L1{c: c}, nil
}

// Get 直接查 ristretto。
func (l *L1) Get(key string) (any, bool) {
	if l == nil {
		return nil, false
	}
	return l.c.Get(key)
}

// SetWithTTL 写入并指定 TTL。cost 应为该项的字节大小估算。
// 注意：ristretto 写入是异步的，调用后立即 Get 不一定命中；本项目读路径
// 总是先经 L1 不命中再查 L2，对 Set→Get 的瞬时一致性不敏感。
func (l *L1) SetWithTTL(key string, val any, cost int64, ttl time.Duration) bool {
	if l == nil {
		return false
	}
	return l.c.SetWithTTL(key, val, cost, ttl)
}

// Del 主动失效（双删用）。
func (l *L1) Del(key string) {
	if l == nil {
		return
	}
	l.c.Del(key)
}

// Wait 等待异步写入对外可见（仅测试用）。
func (l *L1) Wait() {
	if l == nil {
		return
	}
	l.c.Wait()
}
