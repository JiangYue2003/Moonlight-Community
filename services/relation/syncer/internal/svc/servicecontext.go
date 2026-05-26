// Package svc relation-syncer 运行时依赖。
package svc

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	model "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Db            sqlx.SqlConn
	FollowerModel model.FollowerModel

	Redis goredis.UniversalClient

	UserCounterRpc counterpb.UserCounterClient

	Snowflake *snowflakex.Generator
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
		FollowerModel:  model.NewFollowerModel(conn, c.CacheRedis),
		Redis:          rdb,
		UserCounterRpc: counterpb.NewUserCounterClient(zrpc.MustNewClient(c.UserCounterRpc).Conn()),
		Snowflake:      snowflakex.MustNew(c.Snowflake.DatacenterId, c.Snowflake.WorkerId),
	}
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
}
