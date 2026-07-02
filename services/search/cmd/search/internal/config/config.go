package config

import (
	searchapiapp "github.com/zhiguang/zhiguang-go/services/search/api/app"
	searchindexerapp "github.com/zhiguang/zhiguang-go/services/search/indexer/app"
	searchrpcapp "github.com/zhiguang/zhiguang-go/services/search/rpc/app"
)

// Config merges search-api, search-rpc and search-indexer configurations.
type Config struct {
	DisableAPI bool `json:",default=false"`
	Api        searchapiapp.Config
	Rpc        searchrpcapp.Config
	Indexer    searchindexerapp.Config
}
