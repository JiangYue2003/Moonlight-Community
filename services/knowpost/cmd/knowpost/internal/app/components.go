package app

import (
	"context"

	knowpostrpcapp "github.com/zhiguang/zhiguang-go/services/knowpost/internal/bootstrap"
)

type rpcComponent struct{ cfg knowpostrpcapp.Config }

func NewRPCComponent(cfg knowpostrpcapp.Config) Component { return &rpcComponent{cfg: cfg} }

func (c *rpcComponent) Name() string { return "knowpost-rpc" }

func (c *rpcComponent) Run(ctx context.Context) error { return knowpostrpcapp.Run(ctx, c.cfg) }
