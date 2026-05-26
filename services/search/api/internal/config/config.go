package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// AuthRpc 可选 JWT；NonBlock 允许在 auth-rpc 短暂不可用时仍能匿名搜索。
	AuthRpc zrpc.RpcClientConf
	CounterRpc zrpc.RpcClientConf

	// Es 客户端配置（多 addr 即 round-robin）。
	Es EsConf

	// 索引名（默认值见 svcCtx）。
	ContentIndex string `json:",default=zhiguang_content_index"`
}

type EsConf struct {
	Addrs     []string
	Username  string `json:",optional"`
	Password  string `json:",optional"`
	TimeoutMs int    `json:",default=10000"`
}
