package retrieval

import "sort"

// ScoredItem 是检索融合后的统一结果结构。
type ScoredItem struct {
	DocID   string
	PostID  int64
	ChunkID string
	Text    string
	Source  string
	Score   float64
	Rank    int
}

// FuseRRF 对多路结果做 RRF 融合，k 建议为 60。
func FuseRRF(k int, channels ...[]ScoredItem) []ScoredItem {
	if k <= 0 {
		k = 60
	}
	merged := make(map[string]ScoredItem, 128)
	for _, ch := range channels {
		for i, it := range ch {
			key := it.DocID
			if key == "" {
				key = it.ChunkID
			}
			if key == "" {
				continue
			}
			rrf := 1.0 / float64(k+i+1)
			cur, ok := merged[key]
			if !ok {
				it.Score = rrf
				it.Rank = i + 1
				merged[key] = it
				continue
			}
			cur.Score += rrf
			if cur.Text == "" && it.Text != "" {
				cur.Text = it.Text
			}
			if cur.Source == "" {
				cur.Source = it.Source
			}
			merged[key] = cur
		}
	}

	out := make([]ScoredItem, 0, len(merged))
	for _, it := range merged {
		out = append(out, it)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].DocID < out[j].DocID
		}
		return out[i].Score > out[j].Score
	})
	for i := range out {
		out[i].Rank = i + 1
	}
	return out
}
