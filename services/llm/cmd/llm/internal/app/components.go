package app

import (
	"context"

	llmapiapp "github.com/zhiguang/zhiguang-go/services/llm/api/app"
	llmragindexerapp "github.com/zhiguang/zhiguang-go/services/llm/ragindexer/app"
)

type apiComponent struct{ cfg llmapiapp.Config }
type ragIndexerComponent struct{ cfg llmragindexerapp.Config }

func NewAPIComponent(cfg llmapiapp.Config) Component {
	return &apiComponent{cfg: cfg}
}

func NewRagIndexerComponent(cfg llmragindexerapp.Config) Component {
	return &ragIndexerComponent{cfg: cfg}
}

func (c *apiComponent) Name() string        { return "llm-api" }
func (c *ragIndexerComponent) Name() string { return "llm-ragindexer" }

func (c *apiComponent) Run(ctx context.Context) error {
	return llmapiapp.Run(ctx, c.cfg)
}

func (c *ragIndexerComponent) Run(ctx context.Context) error {
	return llmragindexerapp.Run(ctx, c.cfg)
}
