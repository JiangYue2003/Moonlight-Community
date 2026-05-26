// Package config 提供 counter-aggregator 配置。
package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Redis redis.RedisKeyConf
	Kafka KafkaConf
	Flush FlushConf
	Lock  LockConf
	// LegacyLockTtlMs 兼容旧配置 lock_ttl_ms（并入 Lock.TtlMs）
	LegacyLockTtlMs int `json:"lock_ttl_ms,optional"`
	Log   logx.LogConf
}

type KafkaConf struct {
	Brokers []string
	Topic   string
	GroupId string
}

type FlushConf struct {
	IntervalMs int // 默认 1000
	BatchSize  int // 单次 SCAN 批大小，默认 200
}

// LockConf 多副本环境下用 redsync 选主，只让一个 leader 周期性刷写。
type LockConf struct {
	Key   string `json:",default=aggr:flush:lock"`
	TtlMs int    `json:",default=11000"` // 略大于 IntervalMs，避免续期前过期
}
