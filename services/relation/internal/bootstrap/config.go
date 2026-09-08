package bootstrap

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/services/relation/internal/application"
)

// Config relation-rpc 配置：MySQL（写 following + outbox）+ Redis（限流 + ZSet 列表缓存）
// + 下游 user-rpc（聚合 UserSummary）+ 雪花 ID（生成 outbox.id 与 following.id）。
type Config struct {
	zrpc.RpcServerConf

	Mysql      application.MysqlConfig
	CacheRedis cache.CacheConf

	UserRpc zrpc.RpcClientConf

	RateLimit application.RateLimitConfig
	Snowflake application.SnowflakeConfig
}
