package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config relation-rpc 配置：MySQL（写 following + outbox）+ Redis（限流 + ZSet 列表缓存）
// + 下游 user-rpc（聚合 UserSummary）+ 雪花 ID（生成 outbox.id 与 following.id）。
type Config struct {
	zrpc.RpcServerConf

	Mysql      MysqlConf
	CacheRedis cache.CacheConf

	UserRpc zrpc.RpcClientConf

	RateLimit RateLimitConf
	Snowflake SnowflakeConf
}

type MysqlConf struct {
	DataSource string
}

type RateLimitConf struct {
	FollowCapacity     int64 `json:",default=100"`
	FollowRefillPerSec int64 `json:",default=1"`
}

type SnowflakeConf struct {
	WorkerId     int64 `json:",default=1"`
	DatacenterId int64 `json:",default=3"`
}
