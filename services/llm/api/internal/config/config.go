package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuthRpc zrpc.RpcClientConf

	DeepSeek DeepSeekConf
	Tongyi   TongyiConf

	Es       EsConf
	RagIndex string `json:",default=zhiguang_rag_vector"`

	LlmRateLimit LlmRateLimitConf
}

type DeepSeekConf struct {
	BaseUrl   string `json:",default=https://api.deepseek.com"`
	ApiKey    string
	Model     string `json:",default=deepseek-chat"`
	TimeoutMs int    `json:",default=60000"`
	// Temperature 对齐 Java spring.ai.deepseek.chat.options.temperature
	Temperature float32 `json:",default=0.8"`
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

type EsConf struct {
	Addrs     []string
	Username  string `json:",optional"`
	Password  string `json:",optional"`
	TimeoutMs int    `json:",default=10000"`
}

type LlmRateLimitConf struct {
	RedisHost string `json:",default=127.0.0.1:6379"`
	// Describe 每用户令牌桶
	DescribeCapacity     int64 `json:",default=5"`
	DescribeRefillPerSec int64 `json:",default=1"`
	// QaStream 每用户令牌桶
	QaCapacity     int64 `json:",default=3"`
	QaRefillPerSec int64 `json:",default=1"`
}
