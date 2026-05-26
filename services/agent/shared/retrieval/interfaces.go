package retrieval

import "context"

type Query struct {
	UserID int64
	TopK   int
	Text   string
	Vector []float32
}

type VectorProvider interface {
	Search(ctx context.Context, q Query) ([]ScoredItem, error)
	Name() string
}

type KeywordProvider interface {
	Search(ctx context.Context, q Query) ([]ScoredItem, error)
	Name() string
}

type GraphProvider interface {
	Search(ctx context.Context, q Query) ([]ScoredItem, error)
	Name() string
	Enabled() bool
}
