package config

import (
	"github.com/zhiguang/zhiguang-go/pkg/debughttp"
	relationapiapp "github.com/zhiguang/zhiguang-go/services/relation/api/app"
	relationbootstrap "github.com/zhiguang/zhiguang-go/services/relation/internal/bootstrap"
	relationsyncerapp "github.com/zhiguang/zhiguang-go/services/relation/syncer/app"
)

// Config merges relation-api, relation-rpc and relation-syncer.
type Config struct {
	DisableAPI bool `json:",default=false"`
	DebugHTTP  debughttp.Config
	Api        relationapiapp.Config
	Rpc        relationbootstrap.Config
	Syncer     relationsyncerapp.Config
}
