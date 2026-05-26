package cachex

import (
	"math/rand"
	"sync"
	"time"
)

var (
	jitterMu  sync.Mutex
	jitterRng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Jitter 返回 base + [0, max) 的随机时长，用于防止缓存雪崩。
// max <= 0 时直接返回 base（无抖动）。
func Jitter(base, max time.Duration) time.Duration {
	if max <= 0 {
		return base
	}
	jitterMu.Lock()
	defer jitterMu.Unlock()
	return base + time.Duration(jitterRng.Int63n(int64(max)))
}
