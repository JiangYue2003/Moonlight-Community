package hotkey

import (
	"sync/atomic"
	"time"
)

var (
	extLowSeconds    int64 = 20
	extMediumSeconds int64 = 60
	extHighSeconds   int64 = 120
)

// ConfigureExtensions 允许按业务配置覆盖不同热度等级的 TTL 延长秒数。
func ConfigureExtensions(lowSeconds, mediumSeconds, highSeconds int) {
	if lowSeconds > 0 {
		atomic.StoreInt64(&extLowSeconds, int64(lowSeconds))
	}
	if mediumSeconds > 0 {
		atomic.StoreInt64(&extMediumSeconds, int64(mediumSeconds))
	}
	if highSeconds > 0 {
		atomic.StoreInt64(&extHighSeconds, int64(highSeconds))
	}
}

// Extension 根据热度等级返回 TTL 延长量（与 Java HotKeyDetector 严格一致）：
//
//	NONE   → +0
//	LOW    → +20s
//	MEDIUM → +60s
//	HIGH   → +120s
func Extension(l Level) time.Duration {
	switch l {
	case LevelLow:
		return time.Duration(atomic.LoadInt64(&extLowSeconds)) * time.Second
	case LevelMedium:
		return time.Duration(atomic.LoadInt64(&extMediumSeconds)) * time.Second
	case LevelHigh:
		return time.Duration(atomic.LoadInt64(&extHighSeconds)) * time.Second
	default:
		return 0
	}
}

// TTLForPublic 公共读路径（详情、公共 feed、feed item）的 TTL = base + Extension(l)
func TTLForPublic(base time.Duration, l Level) time.Duration {
	return base + Extension(l)
}

// TTLForMine 个人读路径（我的 feed）的 TTL = base + Extension(l)；
// 当前与 Public 共用扩展规则，独立函数便于未来分化。
func TTLForMine(base time.Duration, l Level) time.Duration {
	return base + Extension(l)
}
