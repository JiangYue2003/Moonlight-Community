// Code scaffolded by goctl. Safe to edit.

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/config"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc userpb.AuthClient
	UserRpc userpb.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.AuthRpc).Conn()
	return &ServiceContext{
		Config:  c,
		AuthRpc: userpb.NewAuthClient(conn),
		UserRpc: userpb.NewUserClient(conn),
	}
}
