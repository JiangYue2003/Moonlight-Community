package application

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	model "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type Config struct {
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
	Redis      redis.RedisKeyConf
	UserRpc    zrpc.RpcClientConf
	RateLimit  RateLimitConfig
	Snowflake  SnowflakeConfig
}

type MysqlConfig struct {
	DataSource string
}

type RateLimitConfig struct {
	FollowCapacity     int64 `json:",default=100"`
	FollowRefillPerSec int64 `json:",default=1"`
}

type SnowflakeConfig struct {
	WorkerId     int64 `json:",default=1"`
	DatacenterId int64 `json:",default=3"`
}

type ServiceContext struct {
	Config Config

	Db             sqlx.SqlConn
	FollowingModel model.FollowingModel
	FollowerModel  model.FollowerModel
	OutboxModel    model.OutboxModel

	Redis goredis.UniversalClient

	UserRpc userpb.UserClient

	RateLimiter *ratelimit.TokenBucket
	Snowflake   *snowflakex.Generator

	FollowingTopCache map[int64][]int64
	FollowerTopCache  map[int64][]int64
}

func NewServiceContext(c Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	return &ServiceContext{
		Config:            c,
		Db:                conn,
		FollowingModel:    model.NewFollowingModel(conn, c.CacheRedis),
		FollowerModel:     model.NewFollowerModel(conn, c.CacheRedis),
		OutboxModel:       model.NewOutboxModel(conn, c.CacheRedis),
		Redis:             rdb,
		UserRpc:           userpb.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		RateLimiter:       ratelimit.New(rdb),
		Snowflake:         snowflakex.MustNew(c.Snowflake.DatacenterId, c.Snowflake.WorkerId),
		FollowingTopCache: make(map[int64][]int64),
		FollowerTopCache:  make(map[int64][]int64),
	}
}
