// Package hotkey 实现滑动窗口热点探测。
//
// 设计：将一段窗口（windowSeconds）分为若干 segment（segmentSeconds）并轮转计数。
// 每次 Hit 累加当前段；Score 返回所有段之和（即过去 window 内的命中数）。
// 三级阈值（low / medium / high）映射到 TTL 延长策略，避免热点反复回源。
package hotkey

import (
	"sync"
	"sync/atomic"
	"time"
)

type Level int

const (
	LevelNone Level = iota
	LevelLow
	LevelMedium
	LevelHigh
)

// Config 热点探测参数。
type Config struct {
	WindowSeconds  int
	SegmentSeconds int
	LevelLow       int64
	LevelMedium    int64
	LevelHigh      int64
}

// Detector 简易热点探测器；并发安全。
type Detector struct {
	cfg      Config
	segments int
	stripe   int

	mu     sync.RWMutex
	keys   map[string]*ring // 按 key 维度独立计数
	rotate chan struct{}
}

type ring struct {
	idx   atomic.Int32
	slots []atomic.Int64
}

// New 创建一个 Detector，并启动后台 segment 轮转 goroutine。
// 调用方需要在退出时 close stop chan 释放资源（这里简化未实现）。
func New(cfg Config) *Detector {
	if cfg.WindowSeconds <= 0 {
		cfg.WindowSeconds = 60
	}
	if cfg.SegmentSeconds <= 0 {
		cfg.SegmentSeconds = 10
	}
	segs := cfg.WindowSeconds / cfg.SegmentSeconds
	if segs < 2 {
		segs = 2
	}
	d := &Detector{
		cfg:      cfg,
		segments: segs,
		keys:     make(map[string]*ring),
	}
	go d.rotateLoop()
	return d
}

func (d *Detector) ringFor(key string) *ring {
	d.mu.RLock()
	r, ok := d.keys[key]
	d.mu.RUnlock()
	if ok {
		return r
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if r, ok = d.keys[key]; ok {
		return r
	}
	r = &ring{slots: make([]atomic.Int64, d.segments)}
	d.keys[key] = r
	return r
}

// Hit 累加一次命中。
func (d *Detector) Hit(key string) { d.HitN(key, 1) }

// HitN 累加 n 次命中。
func (d *Detector) HitN(key string, n int64) {
	r := d.ringFor(key)
	idx := int(r.idx.Load())
	r.slots[idx%len(r.slots)].Add(n)
}

// Score 返回当前 key 在窗口内的总命中数。
func (d *Detector) Score(key string) int64 {
	d.mu.RLock()
	r, ok := d.keys[key]
	d.mu.RUnlock()
	if !ok {
		return 0
	}
	var sum int64
	for i := range r.slots {
		sum += r.slots[i].Load()
	}
	return sum
}

// Level 根据 Score 与配置阈值映射到等级。
func (d *Detector) Level(key string) Level {
	s := d.Score(key)
	switch {
	case s >= d.cfg.LevelHigh && d.cfg.LevelHigh > 0:
		return LevelHigh
	case s >= d.cfg.LevelMedium && d.cfg.LevelMedium > 0:
		return LevelMedium
	case s >= d.cfg.LevelLow && d.cfg.LevelLow > 0:
		return LevelLow
	default:
		return LevelNone
	}
}

func (d *Detector) rotateLoop() {
	t := time.NewTicker(time.Duration(d.cfg.SegmentSeconds) * time.Second)
	defer t.Stop()
	for range t.C {
		d.mu.RLock()
		for _, r := range d.keys {
			next := (int(r.idx.Add(1))) % len(r.slots)
			r.slots[next].Store(0)
		}
		d.mu.RUnlock()
	}
}
