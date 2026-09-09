package app

import (
	"context"

	relationbootstrap "github.com/zhiguang/zhiguang-go/services/relation/internal/bootstrap"
	relationsyncerapp "github.com/zhiguang/zhiguang-go/services/relation/syncer/app"
)

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
