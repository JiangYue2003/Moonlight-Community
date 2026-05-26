package esx

import (
	"context"
	"encoding/json"
	"net/http"
)

// SuggestOption 来自 completion suggester。
type SuggestOption struct {
	Text   string          `json:"text"`
	Score  float64         `json:"_score"`
	Source json.RawMessage `json:"_source"`
}

type suggestResp struct {
	Suggest map[string][]struct {
		Text    string          `json:"text"`
		Options []SuggestOption `json:"options"`
	} `json:"suggest"`
}

// Suggest 调 completion suggester；返回去重后的 text 列表。
func (c *Client) Suggest(ctx context.Context, index, field, prefix string, size int) ([]string, error) {
	if size <= 0 {
		size = 10
	}
	body := map[string]any{
		"_source": false,
		"suggest": map[string]any{
			"s": map[string]any{
				"prefix": prefix,
				"completion": map[string]any{
					"field":           field,
					"size":            size,
					"skip_duplicates": true,
				},
			},
		},
	}
	resp, err := c.do(ctx, http.MethodPost, "/"+escapePath(index)+"/_search", body)
	if err != nil {
		return nil, err
	}
	var out suggestResp
	if err := c.readJSON(resp, &out); err != nil {
		return nil, err
	}
	seen := make(map[string]struct{}, size)
	var result []string
	for _, items := range out.Suggest {
		for _, it := range items {
			for _, o := range it.Options {
				if _, ok := seen[o.Text]; ok {
					continue
				}
				seen[o.Text] = struct{}{}
				result = append(result, o.Text)
				if len(result) >= size {
					return result, nil
				}
			}
		}
	}
	return result, nil
}
