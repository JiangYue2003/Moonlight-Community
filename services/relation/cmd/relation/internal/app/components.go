package app

import (
	"context"

	relationapiapp "github.com/zhiguang/zhiguang-go/services/relation/api/app"
	relationbootstrap "github.com/zhiguang/zhiguang-go/services/relation/internal/bootstrap"
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
	cfg relationbootstrap.Config
}

func NewRPCComponent(cfg relationbootstrap.Config) Component {
	return &rpcComponent{cfg: cfg}
}

func (c *rpcComponent) Name() string { return "relation-rpc" }

func (c *rpcComponent) Run(ctx context.Context) error {
	return relationbootstrap.Run(ctx, c.cfg)
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
