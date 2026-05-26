package config

import (
	agentapiapp "github.com/zhiguang/zhiguang-go/services/agent/api/app"
	agentindexerapp "github.com/zhiguang/zhiguang-go/services/agent/indexer/app"
)

type Config struct {
	Api     agentapiapp.Config
	Indexer agentindexerapp.Config
}
