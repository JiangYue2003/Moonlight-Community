package app

import (
	"context"

	agentapiapp "github.com/zhiguang/zhiguang-go/services/agent/api/app"
	agentindexerapp "github.com/zhiguang/zhiguang-go/services/agent/indexer/app"
)

type apiComponent struct{ cfg agentapiapp.Config }
type indexerComponent struct{ cfg agentindexerapp.Config }

func NewAPIComponent(cfg agentapiapp.Config) Component {
	return &apiComponent{cfg: cfg}
}

func NewIndexerComponent(cfg agentindexerapp.Config) Component {
	return &indexerComponent{cfg: cfg}
}

func (c *apiComponent) Name() string     { return "agent-api" }
func (c *indexerComponent) Name() string { return "agent-indexer" }

func (c *apiComponent) Run(ctx context.Context) error {
	return agentapiapp.Run(ctx, c.cfg)
}

func (c *indexerComponent) Run(ctx context.Context) error {
	return agentindexerapp.Run(ctx, c.cfg)
}
