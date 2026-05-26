package svc

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/model"
	einodashscope "github.com/cloudwego/eino-ext/components/embedding/dashscope"
	einodeepseek "github.com/cloudwego/eino-ext/components/model/deepseek"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Chat      model.ChatModel
	Embed     embedding.Embedder
	Es        *esx.Client
	RateLimit *ratelimit.TokenBucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()

	chat, err := einodeepseek.NewChatModel(ctx, &einodeepseek.ChatModelConfig{
		APIKey:      c.DeepSeek.ApiKey,
		Model:       c.DeepSeek.Model,
		BaseURL:     c.DeepSeek.BaseUrl,
		Timeout:     time.Duration(c.DeepSeek.TimeoutMs) * time.Millisecond,
		Temperature: c.DeepSeek.Temperature,
	})
	if err != nil {
		log.Fatalf("llm-rpc: deepseek chat model: %v", err)
	}

	dims := c.Tongyi.EffectiveDimensions()
	emb, err := einodashscope.NewEmbedder(ctx, &einodashscope.EmbeddingConfig{
		APIKey:     c.Tongyi.ApiKey,
		Model:      c.Tongyi.Model,
		Dimensions: &dims,
		Timeout:    time.Duration(c.Tongyi.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("llm-rpc: dashscope embedder: %v", err)
	}

	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("llm-rpc: esx: %v", err)
	}

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs: []string{c.LlmRateLimit.RedisHost},
	})

	return &ServiceContext{
		Config:    c,
		Chat:      chat,
		Embed:     emb,
		Es:        es,
		RateLimit: ratelimit.New(rdb),
	}
}
