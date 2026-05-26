// Code scaffolded by goctl. Safe to edit.

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CounterRpc zrpc.RpcClientConf
	AuthRpc    zrpc.RpcClientConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
