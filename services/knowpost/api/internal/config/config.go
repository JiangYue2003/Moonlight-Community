// Code scaffolded by goctl. Safe to edit.

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AuthRpc     zrpc.RpcClientConf
	CounterRpc  zrpc.RpcClientConf
	KnowPostRpc zrpc.RpcClientConf
	LlmProxy    LlmProxyConf
}

type LlmProxyConf struct {
	BaseURL   string `json:",default=http://127.0.0.1:8008"`
	TimeoutMs int    `json:",default=120000"`
}
