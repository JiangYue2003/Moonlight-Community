package esx

import (
	"context"
	"encoding/json"
	"net/http"
)

// Hit 单条搜索命中。
type Hit struct {
	Id        string              `json:"_id"`
	Score     float64             `json:"_score"`
	Source    json.RawMessage     `json:"_source"`
	Highlight map[string][]string `json:"highlight,omitempty"`
	Sort      []json.RawMessage   `json:"sort,omitempty"`
}

// SearchResult 检索响应（裁剪到我们需要的字段）。
type SearchResult struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

// Search POST /{index}/_search
func (c *Client) Search(ctx context.Context, index string, query json.RawMessage) (*SearchResult, error) {
	path := "/" + escapePath(index) + "/_search"
	resp, err := c.do(ctx, http.MethodPost, path, query)
	if err != nil {
		return nil, err
	}
	var out SearchResult
	if err := c.readJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
