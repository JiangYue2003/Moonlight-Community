package app

import (
	"context"

	counteraggregatorapp "github.com/zhiguang/zhiguang-go/services/counter/aggregator/app"
	counterapiapp "github.com/zhiguang/zhiguang-go/services/counter/api/app"
	counterrpcapp "github.com/zhiguang/zhiguang-go/services/counter/rpc/app"
)

type apiComponent struct{ cfg counterapiapp.Config }

type rpcComponent struct{ cfg counterrpcapp.Config }

type aggregatorComponent struct{ cfg counteraggregatorapp.Config }

func NewAPIComponent(cfg counterapiapp.Config) Component { return &apiComponent{cfg: cfg} }
func NewRPCComponent(cfg counterrpcapp.Config) Component { return &rpcComponent{cfg: cfg} }
func NewAggregatorComponent(cfg counteraggregatorapp.Config) Component {
	return &aggregatorComponent{cfg: cfg}
}

func (c *apiComponent) Name() string        { return "counter-api" }
func (c *rpcComponent) Name() string        { return "counter-rpc" }
func (c *aggregatorComponent) Name() string { return "counter-aggregator" }

func (c *apiComponent) Run(ctx context.Context) error { return counterapiapp.Run(ctx, c.cfg) }
func (c *rpcComponent) Run(ctx context.Context) error { return counterrpcapp.Run(ctx, c.cfg) }
func (c *aggregatorComponent) Run(ctx context.Context) error {
	return counteraggregatorapp.Run(ctx, c.cfg)
}
