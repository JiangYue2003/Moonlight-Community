// Code scaffolded by goctl. Safe to edit.

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
)

type Config struct {
	rest.RestConf
	AuthRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
	Oss     ossx.Config
}
