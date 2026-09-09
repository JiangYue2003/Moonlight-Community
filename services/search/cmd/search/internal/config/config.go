package config

import (
	searchindexerapp "github.com/zhiguang/zhiguang-go/services/search/indexer/app"
	searchrpcapp "github.com/zhiguang/zhiguang-go/services/search/internal/bootstrap"
)

// Config merges search-rpc and search-indexer configurations.
type Config struct {
	Rpc     searchrpcapp.Config
	Indexer searchindexerapp.Config
}
