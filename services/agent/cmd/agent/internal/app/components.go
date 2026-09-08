package app

import (
	"context"

	agentbootstrap "github.com/zhiguang/zhiguang-go/services/agent/internal/bootstrap"
)

type apiComponent struct{ cfg agentbootstrap.APIConfig }
type indexerComponent struct{ cfg agentbootstrap.IndexerConfig }

func NewAPIComponent(cfg agentbootstrap.APIConfig) Component {
	return &apiComponent{cfg: cfg}
}

func NewIndexerComponent(cfg agentbootstrap.IndexerConfig) Component {
	return &indexerComponent{cfg: cfg}
}

func (c *apiComponent) Name() string     { return "agent-api" }
func (c *indexerComponent) Name() string { return "agent-indexer" }

func (c *apiComponent) Run(ctx context.Context) error {
	return agentbootstrap.RunAPI(ctx, c.cfg)
}

func (c *indexerComponent) Run(ctx context.Context) error {
	return agentbootstrap.RunIndexer(ctx, c.cfg)
}
