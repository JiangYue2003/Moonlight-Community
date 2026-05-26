package esx

import (
	"context"
	"encoding/json"
	"net/http"
)

// KnnSearch 调用 ES 9 的 knn 段；filter 可选（传 nil）。
//
// numCandidates 一般取 k*2~k*4，控制召回 vs. 时延。
func (c *Client) KnnSearch(
	ctx context.Context,
	index, field string,
	queryVector []float32,
	k, numCandidates int,
	filter json.RawMessage,
) (*SearchResult, error) {
	if k <= 0 {
		k = 10
	}
	if numCandidates < k {
		numCandidates = k * 2
	}
	knn := map[string]any{
		"field":          field,
		"query_vector":   queryVector,
		"k":              k,
		"num_candidates": numCandidates,
	}
	if len(filter) > 0 {
		knn["filter"] = filter
	}
	body := map[string]any{
		"size": k,
		"knn":  knn,
	}
	resp, err := c.do(ctx, http.MethodPost, "/"+escapePath(index)+"/_search", body)
	if err != nil {
		return nil, err
	}
	var out SearchResult
	if err := c.readJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
