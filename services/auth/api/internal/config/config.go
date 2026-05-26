// Code scaffolded by goctl. Safe to edit.

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AuthRpc zrpc.RpcClientConf
	Auth    struct {
		// 仅占位以满足 go-zero 内置 jwt 中间件依赖；
		// 真实校验改为自定义中间件调 auth-rpc.VerifyToken（见 mehandler）。
		AccessSecret string
		AccessExpire int64
	}
}
