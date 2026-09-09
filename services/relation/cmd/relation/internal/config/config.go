package config

import (
	"github.com/zhiguang/zhiguang-go/pkg/debughttp"
	relationbootstrap "github.com/zhiguang/zhiguang-go/services/relation/internal/bootstrap"
	relationsyncerapp "github.com/zhiguang/zhiguang-go/services/relation/syncer/app"
)

// Config merges relation-rpc and relation-syncer.
type Config struct {
	DebugHTTP debughttp.Config
	Rpc       relationbootstrap.Config
	Syncer    relationsyncerapp.Config
}
