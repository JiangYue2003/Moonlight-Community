// Package svc 持有 counter-aggregator 运行时依赖。
package svc

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/counterlua"
	"github.com/zhiguang/zhiguang-go/pkg/lockx"
	"github.com/zhiguang/zhiguang-go/services/counter/aggregator/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Redis           goredis.UniversalClient
	IncrFieldScript *goredis.Script
	DecrFieldScript *goredis.Script

	Locks *lockx.Mutexer
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Flush.IntervalMs == 0 {
		c.Flush.IntervalMs = 1000
	}
	if c.Flush.BatchSize == 0 {
		c.Flush.BatchSize = 200
	}
	if c.Lock.Key == "" {
		c.Lock.Key = "aggr:flush:lock"
	}
	if c.Lock.TtlMs == 0 && c.LegacyLockTtlMs > 0 {
		c.Lock.TtlMs = c.LegacyLockTtlMs
	}
	if c.Lock.TtlMs == 0 {
		c.Lock.TtlMs = c.Flush.IntervalMs + 5000
	}
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	return &ServiceContext{
		Config:          c,
		Redis:           rdb,
		IncrFieldScript: goredis.NewScript(counterlua.IncrField),
		DecrFieldScript: goredis.NewScript(counterlua.DecrField),
		Locks:           lockx.New(rdb),
	}
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
}
