package svc

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cloudwego/eino/components/embedding"
	einodashscope "github.com/cloudwego/eino-ext/components/embedding/dashscope"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Es          *esx.Client
	KnowPostRpc knowpostpb.KnowPostClient
	Embed       embedding.Embedder
	Redis       goredis.UniversalClient
	HttpClient  *http.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()

	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("rag-indexer: esx: %v", err)
	}
	if err := es.EnsureIndex(ctx, c.RagIndex, esx.RagMapping()); err != nil {
		log.Fatalf("rag-indexer: ensure index: %v", err)
	}

	dims := c.Tongyi.EffectiveDimensions()
	emb, err := einodashscope.NewEmbedder(ctx, &einodashscope.EmbeddingConfig{
		APIKey:     c.Tongyi.ApiKey,
		Model:      c.Tongyi.Model,
		Dimensions: &dims,
		Timeout:    time.Duration(c.Tongyi.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("rag-indexer: dashscope embedder: %v", err)
	}

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})

	return &ServiceContext{
		Config:      c,
		Es:          es,
		KnowPostRpc: knowpostpb.NewKnowPostClient(zrpc.MustNewClient(c.KnowPostRpc).Conn()),
		Embed:       emb,
		Redis:       rdb,
		HttpClient:  &http.Client{Timeout: time.Duration(c.HttpFetchTimeoutMs) * time.Millisecond},
	}
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
}
