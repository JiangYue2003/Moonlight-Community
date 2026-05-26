package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
)

type Config struct {
	Name    string `json:",default=gateway"`
	Host    string `json:",default=0.0.0.0"`
	Port    int    `json:",default=8080"`
	Mode    string `json:",default=dev"`
	Oss     ossx.Config
	AuthRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf

	StorageRpc  zrpc.RpcClientConf
	KnowPostRpc zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
	CounterRpc  zrpc.RpcClientConf
	UserCounterRpc zrpc.RpcClientConf
	SearchRpc   zrpc.RpcClientConf
	LlmRpc      zrpc.RpcClientConf
}
