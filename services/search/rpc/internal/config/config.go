package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	CounterRpc zrpc.RpcClientConf
	Es         EsConf

	ContentIndex string `json:",default=zhiguang_content_index"`
}

type EsConf struct {
	Addrs     []string
	Username  string `json:",optional"`
	Password  string `json:",optional"`
	TimeoutMs int    `json:",default=10000"`
}
