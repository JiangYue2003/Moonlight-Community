package config

import (
	"github.com/zhiguang/zhiguang-go/pkg/debughttp"
	knowpostrpcapp "github.com/zhiguang/zhiguang-go/services/knowpost/internal/bootstrap"
)

type Config struct {
	DebugHTTP debughttp.Config
	Rpc       knowpostrpcapp.Config
}
