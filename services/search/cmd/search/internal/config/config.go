package config

import (
	searchapiapp "github.com/zhiguang/zhiguang-go/services/search/api/app"
	searchindexerapp "github.com/zhiguang/zhiguang-go/services/search/indexer/app"
)

// Config merges search-api and search-indexer configurations.
type Config struct {
	DisableAPI bool `json:",default=false"`
	Api     searchapiapp.Config
	Indexer searchindexerapp.Config
}
