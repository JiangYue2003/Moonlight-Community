package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AuthRpc        zrpc.RpcClientConf
	RelationRpc    zrpc.RpcClientConf
	UserCounterRpc zrpc.RpcClientConf
}
