// Code scaffolded by goctl. Safe to edit.

package svc

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/config"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config      config.Config
	AuthRpc     userpb.AuthClient
	UserRpc     userpb.UserClient
	CounterRpc  counterpb.CounterClient
	KnowPostRpc knowpostpb.KnowPostClient
	LlmClient   *http.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	authConn := zrpc.MustNewClient(c.AuthRpc).Conn()
	counterConn := zrpc.MustNewClient(c.CounterRpc).Conn()
	return &ServiceContext{
		Config:      c,
		AuthRpc:     userpb.NewAuthClient(authConn),
		UserRpc:     userpb.NewUserClient(authConn),
		CounterRpc:  counterpb.NewCounterClient(counterConn),
		KnowPostRpc: knowpostpb.NewKnowPostClient(zrpc.MustNewClient(c.KnowPostRpc).Conn()),
		LlmClient:   &http.Client{Timeout: time.Duration(c.LlmProxy.TimeoutMs) * time.Millisecond},
	}
}
