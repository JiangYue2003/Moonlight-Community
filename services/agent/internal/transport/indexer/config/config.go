package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Es             EsConf
	KnowledgeIndex string `json:",default=zhiguang_agent_knowledge"`

	Kafka KafkaConf

	KnowPostRpc zrpc.RpcClientConf
	CounterRpc  zrpc.RpcClientConf

	Tongyi TongyiConf
	Chunk  ChunkConf

	Mysql  MysqlConf
	Redis  redis.RedisConf
	Dedup  DedupConf
	Milvus MilvusConf

	HttpFetchTimeoutMs int `json:",default=8000"`
	CompensateMinutes  int `json:",default=10"`
	RetryBackoffSec    int `json:",default=120"`

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
}

type TongyiConf struct {
	BaseUrl   string `json:",default=https://dashscope.aliyuncs.com/compatible-mode/v1"`
	ApiKey    string
	Model     string `json:",default=text-embedding-v3"`
	// Dimensions 对齐 Java spring.ai.openai.embedding.options.dimensions
	Dimensions int `json:",default=1536"`
	TimeoutMs int    `json:",default=30000"`
}

func (c TongyiConf) EffectiveDimensions() int {
	if c.Dimensions > 0 {
		return c.Dimensions
	}
	return 1536
}

type ChunkConf struct {
	Size    int `json:",default=800"`
	Overlap int `json:",default=100"`
}

type MysqlConf struct {
	DataSource string
}

type DedupConf struct {
	TtlSeconds int `json:",default=600"`
}

type MilvusConf struct {
	Enabled          bool   `json:",default=false"`
	Address          string `json:",default=127.0.0.1:19530"`
	ApiKey           string `json:",optional"`
	Collection       string `json:",default=agent_knowledge"`
	MemoryCollection string `json:",default=agent_memory_vectors"`
	VectorDim        int    `json:",default=1536"`
	VectorField      string `json:",default=embedding"`
	TimeoutMs        int    `json:",default=5000"`
}
