package query

import (
	"encoding/json"
	"strings"
)

func BuildSearchBody(q string, tags []string, size int, after []any) json.RawMessage {
	must := []any{
		map[string]any{
			"multi_match": map[string]any{
				"query":  q,
				"fields": []string{"title^3", "body"},
				"type":   "best_fields",
			},
		},
	}
	filter := []any{
		map[string]any{"term": map[string]any{"status": "published"}},
	}
	if len(tags) > 0 {
		filter = append(filter, map[string]any{"terms": map[string]any{"tags": tags}})
	}
	body := map[string]any{
		"size": size,
		"query": map[string]any{
			"function_score": map[string]any{
				"query": map[string]any{
					"bool": map[string]any{"must": must, "filter": filter},
				},
				"functions": []any{
					map[string]any{
						"field_value_factor": map[string]any{"field": "like_count", "modifier": "log1p", "missing": 0},
						"weight":             2.0,
					},
					map[string]any{
						"field_value_factor": map[string]any{"field": "view_count", "modifier": "log1p", "missing": 0},
						"weight":             1.0,
					},
				},
				"score_mode": "sum",
				"boost_mode": "sum",
			},
		},
		"highlight": map[string]any{
			"pre_tags":  []string{"<em>"},
			"post_tags": []string{"</em>"},
			"fields": map[string]any{
				"title": map[string]any{},
				"body":  map[string]any{"fragment_size": 160, "number_of_fragments": 1},
			},
		},
		"sort": []any{
			map[string]any{"_score": "desc"},
			map[string]any{"publish_time": "desc"},
			map[string]any{"like_count": "desc"},
			map[string]any{"view_count": "desc"},
			map[string]any{"content_id": "desc"},
		},
	}
	if len(after) > 0 {
		body["search_after"] = after
	}
	out, _ := json.Marshal(body)
	return out
}

func ParseTags(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
