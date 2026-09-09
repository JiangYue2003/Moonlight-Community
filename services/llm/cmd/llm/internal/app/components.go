package app

import (
	"context"

	llmragindexerapp "github.com/zhiguang/zhiguang-go/services/llm/ragindexer/app"
)

type ragIndexerComponent struct{ cfg llmragindexerapp.Config }

func NewRagIndexerComponent(cfg llmragindexerapp.Config) Component {
	return &ragIndexerComponent{cfg: cfg}
}

func (c *ragIndexerComponent) Name() string { return "llm-ragindexer" }

func (c *ragIndexerComponent) Run(ctx context.Context) error {
	return llmragindexerapp.Run(ctx, c.cfg)
}
