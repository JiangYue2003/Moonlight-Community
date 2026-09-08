package config

import (
	"github.com/zhiguang/zhiguang-go/pkg/debughttp"
	knowpostapiapp "github.com/zhiguang/zhiguang-go/services/knowpost/api/app"
	knowpostrpcapp "github.com/zhiguang/zhiguang-go/services/knowpost/internal/bootstrap"
)

type Config struct {
	DisableAPI bool `json:",default=false"`
	DebugHTTP  debughttp.Config
	Api        knowpostapiapp.Config
	Rpc        knowpostrpcapp.Config
}
