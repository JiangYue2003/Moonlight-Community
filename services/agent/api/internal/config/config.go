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

	Es             EsConf
	KnowledgeIndex string `json:",default=zhiguang_agent_knowledge"`

	Mysql MysqlConf
	Redis RedisConf

	Agent  AgentConf
	Milvus MilvusConf
}

type DeepSeekConf struct {
	BaseUrl   string `json:",default=https://api.deepseek.com"`
	ApiKey    string
	Model     string `json:",default=deepseek-chat"`
	TimeoutMs int    `json:",default=60000"`
	// Temperature 对齐 Java spring.ai.deepseek.chat.options.temperature
	Temperature float32 `json:",default=0.8"`

	LiteModel     string `json:",default=deepseek-chat"`
	ProModel      string `json:",default=deepseek-reasoner"`
	LiteTimeoutMs int    `json:",default=45000"`
	ProTimeoutMs  int    `json:",default=90000"`
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

type EsConf struct {
	Addrs     []string
	Username  string `json:",optional"`
	Password  string `json:",optional"`
	TimeoutMs int    `json:",default=10000"`
}

type MysqlConf struct {
	DataSource string
}

type RedisConf struct {
	Host string `json:",default=127.0.0.1:6379"`
	Pass string `json:",optional"`
}

type AgentConf struct {
	MaxSteps         int `json:",default=3"`
	DefaultTopK      int `json:",default=20"`
	MaxTopK          int `json:",default=50"`
	RRFK             int `json:",default=60"`
	MaxQuestionRunes int `json:",default=2000"`
	SessionTTLHours  int `json:",default=24"`
	SummaryTTLHours  int `json:",default=72"`
	ChatRateLimit    RateLimitConf

	ToolWhitelist []string

	PrimaryVector string `json:",default=es"`
	EnableMilvus  bool   `json:",default=false"`
	EnableGraph   bool   `json:",default=false"`

	ModelRoute    ModelRouteConf
	Observability ObservabilityConf
	ModelCost     ModelCostConf
}

type MilvusConf struct {
	Address          string `json:",default=127.0.0.1:19530"`
	ApiKey           string `json:",optional"`
	Collection       string `json:",default=agent_knowledge"`
	MemoryCollection string `json:",default=agent_memory_vectors"`
	VectorDim        int    `json:",default=1536"`
	VectorField      string `json:",default=embedding"`
	TimeoutMs        int    `json:",default=5000"`
}

type RateLimitConf struct {
	Capacity     int64 `json:",default=4"`
	RefillPerSec int64 `json:",default=1"`
}

type ModelRouteConf struct {
	Enable         bool `json:",default=true"`
	EmitRouteEvent bool `json:",default=false"`
	RetryOnProFail bool `json:",default=true"`

	ProFailWindowSec int `json:",default=120"`
	ProFailThreshold int `json:",default=5"`

	QuestionRunesPro   int `json:",default=120"`
	PromptRunesPro     int `json:",default=4000"`
	RecallCountPro     int `json:",default=10"`
	SummaryRunesPro    int `json:",default=400"`
	SessionMsgsPro     int `json:",default=16"`
	PinContentRunesPro int `json:",default=1200"`
}

type ObservabilityConf struct {
	Enable     bool `json:",default=true"`
	SampleRate int  `json:",default=100"`
}

type ModelCostConf struct {
	Currency              string             `json:",default=USD"`
	EstimateCharsPerToken float64            `json:",default=4"`
	DefaultInputPer1K     float64            `json:",default=0"`
	DefaultOutputPer1K    float64            `json:",default=0"`
	InputPer1K            map[string]float64 `json:",optional"`
	OutputPer1K           map[string]float64 `json:",optional"`
}
