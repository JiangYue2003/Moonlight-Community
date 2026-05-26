package app

import (
	"context"

	knowpostapiapp "github.com/zhiguang/zhiguang-go/services/knowpost/api/app"
	knowpostrpcapp "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/app"
)

type apiComponent struct{ cfg knowpostapiapp.Config }
type rpcComponent struct{ cfg knowpostrpcapp.Config }

func NewAPIComponent(cfg knowpostapiapp.Config) Component { return &apiComponent{cfg: cfg} }
func NewRPCComponent(cfg knowpostrpcapp.Config) Component { return &rpcComponent{cfg: cfg} }

func (c *apiComponent) Name() string { return "knowpost-api" }
func (c *rpcComponent) Name() string { return "knowpost-rpc" }

func (c *apiComponent) Run(ctx context.Context) error { return knowpostapiapp.Run(ctx, c.cfg) }
func (c *rpcComponent) Run(ctx context.Context) error { return knowpostrpcapp.Run(ctx, c.cfg) }
