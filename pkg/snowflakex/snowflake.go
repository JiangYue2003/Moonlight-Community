// Package snowflakex 生成 64 位分布式唯一 ID。
//
// 位布局（与原 Java 实现对齐）：
//
//	1 (sign) | 41 (timestamp ms since EPOCH) | 5 (datacenter) | 5 (worker) | 12 (sequence)
//
// EPOCH = 2024-01-01 00:00:00 UTC（与 Java 严格一致），便于历史 ID 双写过渡。
package snowflakex

import (
	"errors"
	"sync"
	"time"
)

const (
	// Epoch 起始时间戳，2024-01-01T00:00:00Z 的毫秒值。
	Epoch int64 = 1704067200000

	workerIdBits     uint8 = 5
	datacenterIdBits uint8 = 5
	sequenceBits     uint8 = 12

	maxWorkerId     int64 = -1 ^ (-1 << workerIdBits)
	maxDatacenterId int64 = -1 ^ (-1 << datacenterIdBits)
	sequenceMask    int64 = -1 ^ (-1 << sequenceBits)

	workerIdShift     = sequenceBits
	datacenterIdShift = sequenceBits + workerIdBits
	timestampShift    = sequenceBits + workerIdBits + datacenterIdBits

	// 时钟回拨容忍上限（毫秒）。
	maxBackwardMs int64 = 5
)

// ErrClockBackward 当系统时钟回拨且超过容忍上限时返回。
var ErrClockBackward = errors.New("snowflakex: clock moved backward beyond tolerance")

// Generator 是线程安全的雪花 ID 生成器。零值不可用，请用 New 构造。
type Generator struct {
	mu            sync.Mutex
	workerId      int64
	datacenterId  int64
	sequence      int64
	lastTimestamp int64
}

// New 创建一个新生成器；workerId / datacenterId 取值范围 [0, 31]。
func New(datacenterId, workerId int64) (*Generator, error) {
	if workerId < 0 || workerId > maxWorkerId {
		return nil, errors.New("snowflakex: workerId out of range [0,31]")
	}
	if datacenterId < 0 || datacenterId > maxDatacenterId {
		return nil, errors.New("snowflakex: datacenterId out of range [0,31]")
	}
	return &Generator{
		workerId:     workerId,
		datacenterId: datacenterId,
	}, nil
}

// MustNew 等价 New 但失败时 panic。
func MustNew(datacenterId, workerId int64) *Generator {
	g, err := New(datacenterId, workerId)
	if err != nil {
		panic(err)
	}
	return g
}

// NextId 返回下一个 ID。并发安全。
func (g *Generator) NextId() (int64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := nowMs()

	if now < g.lastTimestamp {
		offset := g.lastTimestamp - now
		if offset > maxBackwardMs {
			return 0, ErrClockBackward
		}
		// 小幅回拨：等待时钟追上。
		time.Sleep(time.Duration(offset) * time.Millisecond)
		now = nowMs()
		if now < g.lastTimestamp {
			return 0, ErrClockBackward
		}
	}

	if now == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & sequenceMask
		if g.sequence == 0 {
			now = waitNextMs(g.lastTimestamp)
		}
	} else {
		g.sequence = 0
	}
	g.lastTimestamp = now

	id := ((now - Epoch) << timestampShift) |
		(g.datacenterId << datacenterIdShift) |
		(g.workerId << workerIdShift) |
		g.sequence
	return id, nil
}

// MustNextId 等价 NextId，失败 panic。
func (g *Generator) MustNextId() int64 {
	id, err := g.NextId()
	if err != nil {
		panic(err)
	}
	return id
}

func nowMs() int64 {
	return time.Now().UnixMilli()
}

func waitNextMs(last int64) int64 {
	ts := nowMs()
	for ts <= last {
		time.Sleep(100 * time.Microsecond)
		ts = nowMs()
	}
	return ts
}
