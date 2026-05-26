package providers

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval"
)

type ESVectorProvider struct {
	es    *esx.Client
	index string
}

func NewESVectorProvider(es *esx.Client, index string) *ESVectorProvider {
	return &ESVectorProvider{es: es, index: index}
}

func (p *ESVectorProvider) Name() string { return "es_vector" }

func (p *ESVectorProvider) Search(ctx context.Context, q retrieval.Query) ([]retrieval.ScoredItem, error) {
	filter, _ := json.Marshal(map[string]any{"bool": map[string]any{"filter": []any{
		map[string]any{"term": map[string]any{"user_id": q.UserID}},
		map[string]any{"term": map[string]any{"status": "ready"}},
	}}})
	res, err := p.es.KnnSearch(ctx, p.index, "embedding", q.Vector, q.TopK, q.TopK*2, filter)
	if err != nil {
		return nil, err
	}
	items := make([]retrieval.ScoredItem, 0, len(res.Hits.Hits))
	for _, h := range res.Hits.Hits {
		var s struct {
			UserID  int64  `json:"user_id"`
			PostID  int64  `json:"post_id"`
			ChunkID string `json:"chunk_id"`
			Text    string `json:"text"`
		}
		if err := json.Unmarshal(h.Source, &s); err != nil {
			continue
		}
		if s.UserID != q.UserID || strings.TrimSpace(s.Text) == "" {
			continue
		}
		items = append(items, retrieval.ScoredItem{DocID: h.Id, PostID: s.PostID, ChunkID: s.ChunkID, Text: s.Text, Source: p.Name()})
	}
	return items, nil
}

type ESKeywordProvider struct {
	es    *esx.Client
	index string
}

func NewESKeywordProvider(es *esx.Client, index string) *ESKeywordProvider {
	return &ESKeywordProvider{es: es, index: index}
}

func (p *ESKeywordProvider) Name() string { return "es_bm25" }

func (p *ESKeywordProvider) Search(ctx context.Context, q retrieval.Query) ([]retrieval.ScoredItem, error) {
	body := map[string]any{
		"size": q.TopK,
		"query": map[string]any{
			"bool": map[string]any{
				"must": []any{map[string]any{"match": map[string]any{"text": q.Text}}},
				"filter": []any{
					map[string]any{"term": map[string]any{"user_id": q.UserID}},
					map[string]any{"term": map[string]any{"status": "ready"}},
				},
			},
		},
	}
	b, _ := json.Marshal(body)
	res, err := p.es.Search(ctx, p.index, b)
	if err != nil {
		return nil, err
	}
	items := make([]retrieval.ScoredItem, 0, len(res.Hits.Hits))
	for _, h := range res.Hits.Hits {
		var s struct {
			UserID  int64  `json:"user_id"`
			PostID  int64  `json:"post_id"`
			ChunkID string `json:"chunk_id"`
			Text    string `json:"text"`
		}
		if err := json.Unmarshal(h.Source, &s); err != nil {
			continue
		}
		if s.UserID != q.UserID || strings.TrimSpace(s.Text) == "" {
			continue
		}
		items = append(items, retrieval.ScoredItem{DocID: h.Id, PostID: s.PostID, ChunkID: s.ChunkID, Text: s.Text, Source: p.Name()})
	}
	return items, nil
}
