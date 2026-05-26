// Package config 提供 relation-syncer 配置。
package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Mysql      MysqlConf
	CacheRedis cache.CacheConf
	Redis      RedisConf

	Kafka KafkaConf

	UserCounterRpc zrpc.RpcClientConf

	Dedup DedupConf
	ZSet  ZSetConf

	Snowflake SnowflakeConf

	Log logx.LogConf
}

type MysqlConf struct {
	DataSource string
}

type RedisConf struct {
	Host string
	Type string `json:",default=node"`
	Pass string `json:",optional"`
}

type KafkaConf struct {
	Brokers []string
	Topic   string `json:",default=canal-outbox"`
	GroupId string `json:",default=relation-syncer"`
	// AutoOffsetReset 对齐 Java auto-offset-reset（latest/earliest）
	AutoOffsetReset string `json:",default=latest"`
}

type DedupConf struct {
	TtlSeconds int `json:",default=600"`
}

// ZSet 容量上限：写入 follower/following 列表时只保留最近 N 条；超出旧元素被 ZREMRANGEBYRANK 截断。
type ZSetConf struct {
	MaxMembers int `json:",default=2000"`
	TtlSeconds int `json:",default=7200"`
}

type SnowflakeConf struct {
	WorkerId     int64 `json:",default=2"`
	DatacenterId int64 `json:",default=3"`
}
