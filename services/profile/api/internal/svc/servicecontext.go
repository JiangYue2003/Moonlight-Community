package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/config"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc userpb.AuthClient
	UserRpc userpb.UserClient
	Oss     *ossx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	cli, err := ossx.New(c.Oss)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		AuthRpc: userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		UserRpc: userpb.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		Oss:     cli,
	}
}
