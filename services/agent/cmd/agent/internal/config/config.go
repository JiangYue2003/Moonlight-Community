package config

import agentbootstrap "github.com/zhiguang/zhiguang-go/services/agent/internal/bootstrap"

type Config struct {
	Api     agentbootstrap.APIConfig
	Indexer agentbootstrap.IndexerConfig
}
