package config

import (
	relationapiapp "github.com/zhiguang/zhiguang-go/services/relation/api/app"
	relationrpcapp "github.com/zhiguang/zhiguang-go/services/relation/rpc/app"
	relationsyncerapp "github.com/zhiguang/zhiguang-go/services/relation/syncer/app"
)

// Config merges relation-api, relation-rpc and relation-syncer.
type Config struct {
	DisableAPI bool `json:",default=false"`
	Api    relationapiapp.Config
	Rpc    relationrpcapp.Config
	Syncer relationsyncerapp.Config
}
