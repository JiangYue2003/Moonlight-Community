package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Es       EsConf
	RagIndex string `json:",default=zhiguang_rag_vector"`

	Kafka KafkaConf

	KnowPostRpc zrpc.RpcClientConf

	Tongyi TongyiConf

	Chunk ChunkConf

	Redis redis.RedisConf
	Dedup DedupConf

	HttpFetchTimeoutMs int `json:",default=8000"`

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

type TongyiConf struct {
	BaseUrl   string `json:",default=https://dashscope.aliyuncs.com/compatible-mode/v1"`
	ApiKey    string
	Model     string `json:",default=text-embedding-v3"`
	// Dims 兼容历史字段；新配置优先使用 Dimensions。
	Dims int `json:",optional"`
	// Dimensions 对齐 Java spring.ai.openai.embedding.options.dimensions
	Dimensions int `json:",default=1536"`
	BatchSize int    `json:",default=25"`
	TimeoutMs int    `json:",default=30000"`
}

func (c TongyiConf) EffectiveDimensions() int {
	if c.Dimensions > 0 {
		return c.Dimensions
	}
	if c.Dims > 0 {
		return c.Dims
	}
	return 1536
}

type ChunkConf struct {
	Size    int `json:",default=800"`
	Overlap int `json:",default=100"`
}

type DedupConf struct {
	TtlSeconds int `json:",default=600"`
}
