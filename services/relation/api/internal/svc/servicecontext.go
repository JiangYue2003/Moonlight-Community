package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/config"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	AuthRpc        userpb.AuthClient
	UserRpc        userpb.UserClient
	RelationRpc    relationpb.RelationClient
	UserCounterRpc counterpb.UserCounterClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthRpc:        userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		UserRpc:        userpb.NewUserClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		RelationRpc:    relationpb.NewRelationClient(zrpc.MustNewClient(c.RelationRpc).Conn()),
		UserCounterRpc: counterpb.NewUserCounterClient(zrpc.MustNewClient(c.UserCounterRpc).Conn()),
	}
}
