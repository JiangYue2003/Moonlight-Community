package svc

import (
	"log"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc userpb.AuthClient
	CounterRpc counterpb.CounterClient
	Es      *esx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("search-api: init esx: %v", err)
	}
	return &ServiceContext{
		Config:     c,
		AuthRpc:    userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		CounterRpc: counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn()),
		Es:         es,
	}
}
