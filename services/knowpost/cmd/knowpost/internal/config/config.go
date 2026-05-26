package config

import (
	knowpostapiapp "github.com/zhiguang/zhiguang-go/services/knowpost/api/app"
	knowpostrpcapp "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/app"
)

type Config struct {
	DisableAPI bool `json:",default=false"`
	Api knowpostapiapp.Config
	Rpc knowpostrpcapp.Config
}
