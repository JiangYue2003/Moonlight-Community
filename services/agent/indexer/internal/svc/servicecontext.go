package svc

import (
	"context"
	"log"
	"net/http"
	"time"

	einodashscope "github.com/cloudwego/eino-ext/components/embedding/dashscope"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/services/agent/indexer/internal/config"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type ServiceContext struct {
	Config config.Config

	Es          *esx.Client
	Milvus      *milvusclient.Client
	KnowPostRpc knowpostpb.KnowPostClient
	CounterRpc  counterpb.CounterClient
	Embed       embedding.Embedder
	Redis       goredis.UniversalClient
	Db          sqlx.SqlConn
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
		log.Fatalf("agent-indexer: esx: %v", err)
	}
	if err := es.EnsureIndex(ctx, c.KnowledgeIndex, esx.AgentKnowledgeMapping()); err != nil {
		log.Fatalf("agent-indexer: ensure index: %v", err)
	}

	dims := c.Tongyi.EffectiveDimensions()
	emb, err := einodashscope.NewEmbedder(ctx, &einodashscope.EmbeddingConfig{
		APIKey:     c.Tongyi.ApiKey,
		Model:      c.Tongyi.Model,
		Dimensions: &dims,
		Timeout:    time.Duration(c.Tongyi.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("agent-indexer: dashscope embedder: %v", err)
	}

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})

	var mcli *milvusclient.Client
	if c.Milvus.Enabled {
		mctx, cancel := context.WithTimeout(ctx, time.Duration(c.Milvus.TimeoutMs)*time.Millisecond)
		defer cancel()
		mc, merr := milvusclient.New(mctx, &milvusclient.ClientConfig{
			Address: c.Milvus.Address,
			APIKey:  c.Milvus.ApiKey,
		})
		if merr != nil {
			log.Fatalf("agent-indexer: milvus connect: %v", merr)
		}
		if err := ensureKnowledgeCollection(mctx, mc, c.Milvus.Collection, c.Milvus.VectorDim, c.Milvus.VectorField); err != nil {
			log.Fatalf("agent-indexer: milvus ensure knowledge collection: %v", err)
		}
		if err := ensureMemoryFactCollection(mctx, mc, c.Milvus.MemoryCollection, c.Milvus.VectorDim, c.Milvus.VectorField); err != nil {
			log.Fatalf("agent-indexer: milvus ensure memory collection: %v", err)
		}
		mcli = mc
	}

	sc := &ServiceContext{
		Config:      c,
		Es:          es,
		Milvus:      mcli,
		KnowPostRpc: knowpostpb.NewKnowPostClient(zrpc.MustNewClient(c.KnowPostRpc).Conn()),
		CounterRpc:  counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn()),
		Embed:       emb,
		Redis:       rdb,
		Db:          sqlx.NewMysql(c.Mysql.DataSource),
		HttpClient:  &http.Client{Timeout: time.Duration(c.HttpFetchTimeoutMs) * time.Millisecond},
	}
	if err := sc.ensureTables(ctx); err != nil {
		log.Fatalf("agent-indexer: ensure tables: %v", err)
	}
	return sc
}

func ensureKnowledgeCollection(ctx context.Context, cli *milvusclient.Client, collection string, dim int, vectorField string) error {
	has, err := cli.HasCollection(ctx, milvusclient.NewHasCollectionOption(collection))
	if err != nil {
		return err
	}
	if !has {
		schema := entity.NewSchema().
			WithDynamicFieldEnabled(false).
			WithField(entity.NewField().WithName("id").WithDataType(entity.FieldTypeVarChar).WithIsPrimaryKey(true).WithIsAutoID(false).WithMaxLength(128)).
			WithField(entity.NewField().WithName("user_id").WithDataType(entity.FieldTypeInt64)).
			WithField(entity.NewField().WithName("post_id").WithDataType(entity.FieldTypeInt64)).
			WithField(entity.NewField().WithName("chunk_id").WithDataType(entity.FieldTypeVarChar).WithMaxLength(64)).
			WithField(entity.NewField().WithName("title_path").WithDataType(entity.FieldTypeVarChar).WithMaxLength(1024)).
			WithField(entity.NewField().WithName("text").WithDataType(entity.FieldTypeVarChar).WithMaxLength(32768)).
			WithField(entity.NewField().WithName("version").WithDataType(entity.FieldTypeVarChar).WithMaxLength(128)).
			WithField(entity.NewField().WithName("status").WithDataType(entity.FieldTypeVarChar).WithMaxLength(16)).
			WithField(entity.NewField().WithName(vectorField).WithDataType(entity.FieldTypeFloatVector).WithDim(int64(dim)))
		opt := milvusclient.NewCreateCollectionOption(collection, schema).WithIndexOptions(
			milvusclient.NewCreateIndexOption(collection, vectorField, index.NewHNSWIndex(entity.COSINE, 16, 200)),
		)
		if err := cli.CreateCollection(ctx, opt); err != nil {
			return err
		}
	}
	loadTask, err := cli.LoadCollection(ctx, milvusclient.NewLoadCollectionOption(collection))
	if err != nil {
		return err
	}
	return loadTask.Await(ctx)
}

func ensureMemoryFactCollection(ctx context.Context, cli *milvusclient.Client, collection string, dim int, vectorField string) error {
	if collection == "" {
		return nil
	}
	has, err := cli.HasCollection(ctx, milvusclient.NewHasCollectionOption(collection))
	if err != nil {
		return err
	}
	if !has {
		schema := entity.NewSchema().
			WithDynamicFieldEnabled(false).
			WithField(entity.NewField().WithName("id").WithDataType(entity.FieldTypeVarChar).WithIsPrimaryKey(true).WithIsAutoID(false).WithMaxLength(128)).
			WithField(entity.NewField().WithName("user_id").WithDataType(entity.FieldTypeInt64)).
			WithField(entity.NewField().WithName("fact_id").WithDataType(entity.FieldTypeVarChar).WithMaxLength(64)).
			WithField(entity.NewField().WithName("subject").WithDataType(entity.FieldTypeVarChar).WithMaxLength(255)).
			WithField(entity.NewField().WithName("predicate").WithDataType(entity.FieldTypeVarChar).WithMaxLength(255)).
			WithField(entity.NewField().WithName("object_value").WithDataType(entity.FieldTypeVarChar).WithMaxLength(2000)).
			WithField(entity.NewField().WithName("source_ref").WithDataType(entity.FieldTypeVarChar).WithMaxLength(255)).
			WithField(entity.NewField().WithName("version").WithDataType(entity.FieldTypeVarChar).WithMaxLength(128)).
			WithField(entity.NewField().WithName("status").WithDataType(entity.FieldTypeVarChar).WithMaxLength(16)).
			WithField(entity.NewField().WithName("confidence").WithDataType(entity.FieldTypeFloat)).
			WithField(entity.NewField().WithName(vectorField).WithDataType(entity.FieldTypeFloatVector).WithDim(int64(dim)))
		opt := milvusclient.NewCreateCollectionOption(collection, schema).WithIndexOptions(
			milvusclient.NewCreateIndexOption(collection, vectorField, index.NewHNSWIndex(entity.COSINE, 16, 200)),
		)
		if err := cli.CreateCollection(ctx, opt); err != nil {
			return err
		}
	}
	loadTask, err := cli.LoadCollection(ctx, milvusclient.NewLoadCollectionOption(collection))
	if err != nil {
		return err
	}
	return loadTask.Await(ctx)
}

func (s *ServiceContext) ensureTables(ctx context.Context) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS agent_knowledge_index (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL,
			post_id BIGINT NOT NULL,
			version_hash VARCHAR(128) NOT NULL,
			status VARCHAR(16) NOT NULL,
			chunk_count INT NOT NULL DEFAULT 0,
			last_error VARCHAR(255) NOT NULL DEFAULT '',
			retry_count INT NOT NULL DEFAULT 0,
			next_retry_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL,
			created_at BIGINT NOT NULL,
			UNIQUE KEY uk_user_post_hash (user_id, post_id, version_hash),
			INDEX idx_status_retry (status, next_retry_at),
			INDEX idx_status_updated (status, updated_at)
		)`}
	for _, q := range stmts {
		if _, err := s.Db.ExecCtx(ctx, q); err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
	if s.Milvus != nil {
		_ = s.Milvus.Close(context.Background())
	}
}
