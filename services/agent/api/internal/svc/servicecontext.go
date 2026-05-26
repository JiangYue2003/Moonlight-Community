package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	einodashscope "github.com/cloudwego/eino-ext/components/embedding/dashscope"
	einodeepseek "github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/model"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/observability"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
	memoryproviders "github.com/zhiguang/zhiguang-go/services/agent/shared/memory/providers"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	AuthRpc userpb.AuthClient

	Chat     model.ChatModel
	ChatLite model.ChatModel
	ChatPro  model.ChatModel
	Router   *ModelRouter
	Embed    embedding.Embedder
	Es       *esx.Client
	Milvus   *milvusclient.Client
	Db       sqlx.SqlConn
	Redis    goredis.UniversalClient

	MemoryFacts   memory.FactStore
	MemoryVectors memory.VectorStore
	Obs           *observability.AgentObservability

	RateLimit *ratelimit.TokenBucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()

	liteModel := strings.TrimSpace(c.DeepSeek.LiteModel)
	if liteModel == "" {
		liteModel = strings.TrimSpace(c.DeepSeek.Model)
	}
	if liteModel == "" {
		liteModel = "deepseek-chat"
	}
	proModel := strings.TrimSpace(c.DeepSeek.ProModel)
	if proModel == "" {
		proModel = liteModel
	}
	liteTimeout := c.DeepSeek.LiteTimeoutMs
	if liteTimeout <= 0 {
		liteTimeout = c.DeepSeek.TimeoutMs
	}
	proTimeout := c.DeepSeek.ProTimeoutMs
	if proTimeout <= 0 {
		proTimeout = c.DeepSeek.TimeoutMs
	}

	chatLite, err := einodeepseek.NewChatModel(ctx, &einodeepseek.ChatModelConfig{
		APIKey:      c.DeepSeek.ApiKey,
		Model:       liteModel,
		BaseURL:     c.DeepSeek.BaseUrl,
		Timeout:     time.Duration(liteTimeout) * time.Millisecond,
		Temperature: c.DeepSeek.Temperature,
	})
	if err != nil {
		log.Fatalf("agent-api: deepseek lite chat model: %v", err)
	}
	chatPro, err := einodeepseek.NewChatModel(ctx, &einodeepseek.ChatModelConfig{
		APIKey:      c.DeepSeek.ApiKey,
		Model:       proModel,
		BaseURL:     c.DeepSeek.BaseUrl,
		Timeout:     time.Duration(proTimeout) * time.Millisecond,
		Temperature: c.DeepSeek.Temperature,
	})
	if err != nil {
		log.Fatalf("agent-api: deepseek pro chat model: %v", err)
	}

	dims := c.Tongyi.EffectiveDimensions()
	emb, err := einodashscope.NewEmbedder(ctx, &einodashscope.EmbeddingConfig{
		APIKey:     c.Tongyi.ApiKey,
		Model:      c.Tongyi.Model,
		Dimensions: &dims,
		Timeout:    time.Duration(c.Tongyi.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("agent-api: dashscope embedder: %v", err)
	}

	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("agent-api: esx: %v", err)
	}
	if err := es.EnsureIndex(ctx, c.KnowledgeIndex, esx.AgentKnowledgeMapping()); err != nil {
		log.Fatalf("agent-api: ensure index: %v", err)
	}

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})

	var mcli *milvusclient.Client
	if c.Agent.EnableMilvus {
		mctx, cancel := context.WithTimeout(ctx, time.Duration(c.Milvus.TimeoutMs)*time.Millisecond)
		defer cancel()
		mc, merr := milvusclient.New(mctx, &milvusclient.ClientConfig{
			Address: c.Milvus.Address,
			APIKey:  c.Milvus.ApiKey,
		})
		if merr != nil {
			log.Fatalf("agent-api: milvus connect: %v", merr)
		}
		if err := ensureKnowledgeCollection(mctx, mc, c.Milvus.Collection, c.Milvus.VectorDim, c.Milvus.VectorField); err != nil {
			log.Fatalf("agent-api: milvus ensure knowledge collection: %v", err)
		}
		if err := ensureMemoryFactCollection(mctx, mc, c.Milvus.MemoryCollection, c.Milvus.VectorDim, c.Milvus.VectorField); err != nil {
			log.Fatalf("agent-api: milvus ensure memory collection: %v", err)
		}
		mcli = mc
	}

	sc := &ServiceContext{
		Config:    c,
		AuthRpc:   userpb.NewAuthClient(zrpc.MustNewClient(c.AuthRpc).Conn()),
		Chat:      chatLite,
		ChatLite:  chatLite,
		ChatPro:   chatPro,
		Embed:     emb,
		Es:        es,
		Milvus:    mcli,
		Db:        sqlx.NewMysql(c.Mysql.DataSource),
		Redis:     rdb,
		RateLimit: ratelimit.New(rdb),
	}

	if err := sc.ensureTables(ctx); err != nil {
		log.Fatalf("agent-api: ensure tables: %v", err)
	}
	sc.MemoryFacts = memoryproviders.NewMySQLFactStore(sc.Db)
	sc.MemoryVectors = memoryproviders.NewMilvusFactVectorStore(c.Agent.EnableMilvus, sc.Milvus, c.Milvus.MemoryCollection, c.Milvus.VectorField, c.Milvus.VectorDim)
	sc.Obs = observability.NewAgentObservability(c.Agent.Observability, c.Agent.ModelCost)
	sc.Router = NewModelRouter(c.Agent.ModelRoute, sc.Redis, sc.ChatLite, sc.ChatPro, sc.Obs)
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
	if strings.TrimSpace(collection) == "" {
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
		`CREATE TABLE IF NOT EXISTS agent_sessions (
			session_id VARCHAR(64) PRIMARY KEY,
			user_id BIGINT NOT NULL,
			title VARCHAR(255) NOT NULL DEFAULT '',
			created_at BIGINT NOT NULL,
			updated_at BIGINT NOT NULL,
			INDEX idx_user_updated (user_id, updated_at)
		)`,
		`CREATE TABLE IF NOT EXISTS agent_feedback (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			session_id VARCHAR(64) NOT NULL,
			user_id BIGINT NOT NULL,
			trace_id VARCHAR(128) NOT NULL,
			score INT NOT NULL,
			comment TEXT,
			created_at BIGINT NOT NULL,
			INDEX idx_user_session (user_id, session_id)
		)`,
		`CREATE TABLE IF NOT EXISTS agent_memory_pin (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL,
			session_id VARCHAR(64) NOT NULL,
			tag VARCHAR(64) NOT NULL DEFAULT '',
			content TEXT NOT NULL,
			created_at BIGINT NOT NULL,
			INDEX idx_user_created (user_id, created_at)
		)`,
		`CREATE TABLE IF NOT EXISTS agent_tool_audit (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			session_id VARCHAR(64) NOT NULL,
			trace_id VARCHAR(128) NOT NULL,
			user_id BIGINT NOT NULL,
			tool_name VARCHAR(64) NOT NULL,
			params_hash VARCHAR(128) NOT NULL,
			latency_ms BIGINT NOT NULL,
			status VARCHAR(16) NOT NULL,
			err_msg VARCHAR(255) NOT NULL DEFAULT '',
			created_at BIGINT NOT NULL,
			INDEX idx_user_created (user_id, created_at),
			INDEX idx_trace (trace_id)
		)`,
		`CREATE TABLE IF NOT EXISTS agent_memory_facts (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL,
			fact_id VARCHAR(64) NOT NULL,
			subject VARCHAR(255) NOT NULL,
			predicate VARCHAR(255) NOT NULL,
			object_value TEXT NOT NULL,
			source_ref VARCHAR(255) NOT NULL,
			confidence DOUBLE NOT NULL,
			version VARCHAR(128) NOT NULL,
			status VARCHAR(16) NOT NULL DEFAULT 'active',
			created_at BIGINT NOT NULL,
			updated_at BIGINT NOT NULL,
			UNIQUE KEY uk_fact (user_id, subject, predicate, version, fact_id),
			INDEX idx_user_created (user_id, created_at)
		)`,
	}
	for _, q := range stmts {
		if _, err := s.Db.ExecCtx(ctx, q); err != nil {
			return err
		}
	}
	return nil
}

func SessionMessagesKey(userID int64, sessionID string) string {
	return fmt.Sprintf("agent:session:%d:%s:msgs", userID, sessionID)
}

func SessionSummaryKey(userID int64, sessionID string) string {
	return fmt.Sprintf("agent:session:%d:%s:summary", userID, sessionID)
}

func MarshalHistory(role, content string, ts int64) string {
	m := map[string]any{"role": role, "content": content, "createdAt": ts}
	b, _ := json.Marshal(m)
	return string(b)
}

func TrimContent(s string, n int) string {
	r := []rune(strings.TrimSpace(s))
	if n <= 0 || len(r) <= n {
		return string(r)
	}
	return string(r[:n])
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
	if s.Milvus != nil {
		_ = s.Milvus.Close(context.Background())
	}
}
