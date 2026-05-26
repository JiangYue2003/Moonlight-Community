package svc

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/config"
	model "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

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

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	return &ServiceContext{
		Config:         c,
		Db:             conn,
		FollowingModel: model.NewFollowingModel(conn, c.CacheRedis),
		FollowerModel:  model.NewFollowerModel(conn, c.CacheRedis),
		OutboxModel:    model.NewOutboxModel(conn, c.CacheRedis),
		Redis:          rdb,
		UserRpc:        userpb.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		RateLimiter:    ratelimit.New(rdb),
		Snowflake:      snowflakex.MustNew(c.Snowflake.DatacenterId, c.Snowflake.WorkerId),
		FollowingTopCache: make(map[int64][]int64),
		FollowerTopCache:  make(map[int64][]int64),
	}
}
