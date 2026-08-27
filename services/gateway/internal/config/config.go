package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
)

type Config struct {
	Name string `json:",default=gateway"`
	Host string `json:",default=0.0.0.0"`
	Port int    `json:",default=8080"`
	Mode string `json:",default=dev"`
	Log  logx.LogConf
	// HTTPAccessLog controls Gin's synchronous per-request access logger.
	// It remains enabled by default for development and is explicitly disabled
	// by the benchmark-only generated config.
	HTTPAccessLog bool `json:",default=true"`
	Oss           ossx.Config
	AuthRpc       zrpc.RpcClientConf
	UserRpc       zrpc.RpcClientConf

	StorageRpc     zrpc.RpcClientConf
	KnowPostRpc    zrpc.RpcClientConf
	RelationRpc    zrpc.RpcClientConf
	CounterRpc     zrpc.RpcClientConf
	UserCounterRpc zrpc.RpcClientConf
	SearchRpc      zrpc.RpcClientConf
	LlmRpc         zrpc.RpcClientConf
}
