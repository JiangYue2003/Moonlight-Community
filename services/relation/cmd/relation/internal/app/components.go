package app

import (
	"context"

	relationapiapp "github.com/zhiguang/zhiguang-go/services/relation/api/app"
	relationrpcapp "github.com/zhiguang/zhiguang-go/services/relation/rpc/app"
	relationsyncerapp "github.com/zhiguang/zhiguang-go/services/relation/syncer/app"
)

type apiComponent struct {
	cfg relationapiapp.Config
}

func NewAPIComponent(cfg relationapiapp.Config) Component {
	return &apiComponent{cfg: cfg}
}

func (c *apiComponent) Name() string { return "relation-api" }

func (c *apiComponent) Run(ctx context.Context) error {
	return relationapiapp.Run(ctx, c.cfg)
}

type rpcComponent struct {
	cfg relationrpcapp.Config
}

func NewRPCComponent(cfg relationrpcapp.Config) Component {
	return &rpcComponent{cfg: cfg}
}

func (c *rpcComponent) Name() string { return "relation-rpc" }

func (c *rpcComponent) Run(ctx context.Context) error {
	return relationrpcapp.Run(ctx, c.cfg)
}

type syncerComponent struct {
	cfg relationsyncerapp.Config
}

func NewSyncerComponent(cfg relationsyncerapp.Config) Component {
	return &syncerComponent{cfg: cfg}
}

func (c *syncerComponent) Name() string { return "relation-syncer" }

func (c *syncerComponent) Run(ctx context.Context) error {
	return relationsyncerapp.Run(ctx, c.cfg)
}
