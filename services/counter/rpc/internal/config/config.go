package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

// Config counter-rpc 配置；Redis 用于位图与 SDS，Kafka 用于事件发布。
type Config struct {
	zrpc.RpcServerConf
	Kafka KafkaConf
	Rebuild RebuildConf `json:",optional"`
}

type KafkaConf struct {
	Brokers []string
	Topic   string
	Producer ProducerConf `json:",optional"`
}

type ProducerConf struct {
	// LingerMs 对齐 Java producer.linger（默认 10ms）。
	LingerMs int `json:",default=10"`
}

type RebuildConf struct {
	Enabled bool `json:",optional"`
	Lock    RebuildLockConf   `json:",optional"`
	Rate    RebuildRateConf   `json:",optional"`
	Backoff RebuildBackoffConf `json:",optional"`
}

type RebuildLockConf struct {
	TtlMs int `json:",default=5000"`
}

type RebuildRateConf struct {
	Permits int `json:",default=3"`
	WindowSeconds int `json:",default=10"`
}

type RebuildBackoffConf struct {
	BaseMs int `json:",default=500"`
	MaxMs  int `json:",default=30000"`
}
