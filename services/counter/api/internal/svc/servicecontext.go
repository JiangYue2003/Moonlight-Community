package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/config"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config     config.Config
	CounterRpc counterpb.CounterClient
	AuthRpc    userpb.AuthClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CounterRpc: counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn()),
		AuthRpc:    userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
	}
}
