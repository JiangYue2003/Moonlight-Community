package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval"
)

type MilvusProvider struct {
	enabled     bool
	cli         *milvusclient.Client
	collection  string
	vectorField string
}

func NewMilvusProvider(enabled bool, cli *milvusclient.Client, collection, vectorField string) *MilvusProvider {
	return &MilvusProvider{enabled: enabled, cli: cli, collection: collection, vectorField: vectorField}
}

func (p *MilvusProvider) Name() string { return "milvus" }

func (p *MilvusProvider) Search(ctx context.Context, q retrieval.Query) ([]retrieval.ScoredItem, error) {
	if !p.enabled || p.cli == nil || p.collection == "" {
		return nil, nil
	}
	filter := fmt.Sprintf("user_id == %d and status == \"ready\"", q.UserID)
	resultSets, err := p.cli.Search(ctx,
		milvusclient.NewSearchOption(p.collection, q.TopK, []entity.Vector{entity.FloatVector(q.Vector)}).
			WithANNSField(p.vectorField).
			WithFilter(filter).
			WithOutputFields("post_id", "chunk_id", "text"),
	)
	if err != nil {
		return nil, err
	}
	items := make([]retrieval.ScoredItem, 0, q.TopK)
	for _, rs := range resultSets {
		postCol := rs.GetColumn("post_id")
		chunkCol := rs.GetColumn("chunk_id")
		textCol := rs.GetColumn("text")
		if postCol == nil || chunkCol == nil || textCol == nil {
			continue
		}
		for i := 0; i < rs.ResultCount; i++ {
			postID, err1 := postCol.GetAsInt64(i)
			chunkID, err2 := chunkCol.GetAsString(i)
			text, err3 := textCol.GetAsString(i)
			if err1 != nil || err2 != nil || err3 != nil {
				continue
			}
			if strings.TrimSpace(text) == "" {
				continue
			}
			docID := fmt.Sprintf("u%d-p%d-%s", q.UserID, postID, chunkID)
			items = append(items, retrieval.ScoredItem{
				DocID:   docID,
				PostID:  postID,
				ChunkID: chunkID,
				Text:    text,
				Source:  p.Name(),
			})
		}
	}
	return items, nil
}

type Neo4jProvider struct {
	enabled bool
}

func NewNeo4jProvider(enabled bool) *Neo4jProvider { return &Neo4jProvider{enabled: enabled} }

func (p *Neo4jProvider) Name() string { return "neo4j" }

func (p *Neo4jProvider) Enabled() bool { return p.enabled }

func (p *Neo4jProvider) Search(ctx context.Context, q retrieval.Query) ([]retrieval.ScoredItem, error) {
	if !p.enabled {
		return nil, nil
	}
	// Phase B: 仅保留接口占位，默认不返回图谱结果。
	return nil, nil
}
