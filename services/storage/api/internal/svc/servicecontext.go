// Code scaffolded by goctl. Safe to edit.

package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	knowmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"github.com/zhiguang/zhiguang-go/services/storage/api/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	AuthRpc        userpb.AuthClient
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
		AuthRpc:        userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		KnowPostsModel: knowmodel.NewKnowPostsModel(conn, c.CacheRedis),
		Oss:            cli,
	}
}
