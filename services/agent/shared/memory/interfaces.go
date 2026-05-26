package memory

import "context"

type Fact struct {
	FactID      string
	Subject     string
	Predicate   string
	ObjectValue string
	SourceRef   string
	Confidence  float64
	Version     string
	Status      string
	CreatedAt   int64
	UpdatedAt   int64
}

type Preference struct {
	PreferenceID string
	Kind         string
	Content      string
	Confidence   float64
	Source       string
	Status       string
	LastSeenAt   int64
	CreatedAt    int64
	UpdatedAt    int64
}

type FactVector struct {
	Fact   Fact
	Vector []float32
}

type Query struct {
	UserID int64
	TopK   int
	Text   string
	Vector []float32
}

type ScoredFact struct {
	Fact   Fact
	Score  float64
	Source string
	Rank   int
}

type FactStore interface {
	UpsertFacts(ctx context.Context, userID int64, facts []Fact) error
	SearchFacts(ctx context.Context, q Query) ([]ScoredFact, error)
}

type VectorStore interface {
	UpsertFactVectors(ctx context.Context, userID int64, vectors []FactVector) error
	SearchFactVectors(ctx context.Context, q Query) ([]ScoredFact, error)
}

type PreferenceStore interface {
	UpsertPreferences(ctx context.Context, userID int64, prefs []Preference) error
	ListActivePreferences(ctx context.Context, userID int64, limit int) ([]Preference, error)
}
