package svc

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/counterlua"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/pkg/lockx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Redis goredis.UniversalClient
	Kafka *kafkax.Producer

	// 预编译 Lua 脚本：客户端会先尝试 EVALSHA，失败再 EVAL 上传。
	ToggleScript    *goredis.Script
	IncrFieldScript *goredis.Script
	DecrFieldScript *goredis.Script

	// Locks 用于 GetCounts 在 SDS 缺失时的重建保护，防止多副本同时 BITCOUNT 雪崩。
	Locks *lockx.Mutexer
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	return &ServiceContext{
		Config:          c,
		Redis:           rdb,
		Kafka:           kafkax.NewProducerWithOptions(c.Kafka.Brokers, kafkax.ProducerOptions{LingerMs: c.Kafka.Producer.LingerMs}),
		ToggleScript:    goredis.NewScript(counterlua.Toggle),
		IncrFieldScript: goredis.NewScript(counterlua.IncrField),
		DecrFieldScript: goredis.NewScript(counterlua.DecrField),
		Locks:           lockx.New(rdb),
	}
}
