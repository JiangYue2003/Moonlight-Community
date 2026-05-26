package config

import (
	counteraggregatorapp "github.com/zhiguang/zhiguang-go/services/counter/aggregator/app"
	counterapiapp "github.com/zhiguang/zhiguang-go/services/counter/api/app"
	counterrpcapp "github.com/zhiguang/zhiguang-go/services/counter/rpc/app"
)

type Config struct {
	DisableAPI bool `json:",default=false"`
	Api        counterapiapp.Config
	Rpc        counterrpcapp.Config
	Aggregator counteraggregatorapp.Config
}
