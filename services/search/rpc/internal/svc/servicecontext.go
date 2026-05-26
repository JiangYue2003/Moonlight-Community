package svc

import (
	"log"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/services/search/rpc/internal/config"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type ServiceContext struct {
	Config     config.Config
	CounterRpc counterpb.CounterClient
	Es         *esx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("search-rpc: init esx: %v", err)
	}
	return &ServiceContext{
		Config:     c,
		CounterRpc: counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn()),
		Es:         es,
	}
}
