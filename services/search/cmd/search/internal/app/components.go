package app

import (
	"context"

	searchapiapp "github.com/zhiguang/zhiguang-go/services/search/api/app"
	searchindexerapp "github.com/zhiguang/zhiguang-go/services/search/indexer/app"
	searchrpcapp "github.com/zhiguang/zhiguang-go/services/search/rpc/app"
)

type apiComponent struct {
	cfg searchapiapp.Config
}

func NewAPIComponent(cfg searchapiapp.Config) Component {
	return &apiComponent{cfg: cfg}
}

func (c *apiComponent) Name() string { return "search-api" }

func (c *apiComponent) Run(ctx context.Context) error {
	return searchapiapp.Run(ctx, c.cfg)
}

type rpcComponent struct {
	cfg searchrpcapp.Config
}

func NewRPCComponent(cfg searchrpcapp.Config) Component {
	return &rpcComponent{cfg: cfg}
}

func (c *rpcComponent) Name() string { return "search-rpc" }

func (c *rpcComponent) Run(ctx context.Context) error {
	return searchrpcapp.Run(ctx, c.cfg)
}

type indexerComponent struct {
	cfg searchindexerapp.Config
}

func NewIndexerComponent(cfg searchindexerapp.Config) Component {
	return &indexerComponent{cfg: cfg}
}

func (c *indexerComponent) Name() string { return "search-indexer" }

func (c *indexerComponent) Run(ctx context.Context) error {
	return searchindexerapp.Run(ctx, c.cfg)
}
