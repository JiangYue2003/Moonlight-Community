package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	knowmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"github.com/zhiguang/zhiguang-go/services/storage/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	KnowPostsModel knowmodel.KnowPostsModel
	Oss            *ossx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	cli, err := ossx.New(c.Oss)
	if err != nil {
		panic(err)
	}
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		KnowPostsModel: knowmodel.NewKnowPostsModel(conn, c.CacheRedis),
		Oss:            cli,
	}
}
