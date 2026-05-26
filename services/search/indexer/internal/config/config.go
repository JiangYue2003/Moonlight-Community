package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Es           EsConf
	ContentIndex string `json:",default=zhiguang_content_index"`

	Kafka KafkaConf

	KnowPostRpc zrpc.RpcClientConf

	Redis redis.RedisConf

	Dedup DedupConf

	// 索引正文最长截断（rune）
	ContentMaxRunes int `json:",default=4000"`
	// HTTP 拉 contentUrl 的超时
	HttpFetchTimeoutMs int `json:",default=5000"`

	Log logx.LogConf
}

type EsConf struct {
	Addrs     []string
	Username  string `json:",optional"`
	Password  string `json:",optional"`
	TimeoutMs int    `json:",default=10000"`
}

type KafkaConf struct {
	Brokers []string
	Topic   string
	GroupId string
	AutoOffsetReset string `json:",default=latest"`
}

type DedupConf struct {
	TtlSeconds int `json:",default=600"`
}
