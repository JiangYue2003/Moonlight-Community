package providers

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

type MySQLFactStore struct {
	db sqlx.SqlConn
}

func NewMySQLFactStore(db sqlx.SqlConn) *MySQLFactStore {
	return &MySQLFactStore{db: db}
}

func (s *MySQLFactStore) UpsertFacts(ctx context.Context, userID int64, facts []memory.Fact) error {
	if s == nil || s.db == nil || userID <= 0 || len(facts) == 0 {
		return nil
	}
	for _, f := range facts {
		if strings.TrimSpace(f.FactID) == "" || strings.TrimSpace(f.Subject) == "" || strings.TrimSpace(f.Predicate) == "" || strings.TrimSpace(f.ObjectValue) == "" {
			continue
		}
		status := strings.TrimSpace(f.Status)
		if status == "" {
			status = "active"
		}
		_, err := s.db.ExecCtx(ctx,
			"INSERT INTO agent_memory_facts (user_id,fact_id,subject,predicate,object_value,source_ref,confidence,version,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE confidence=VALUES(confidence),source_ref=VALUES(source_ref),status=VALUES(status),updated_at=VALUES(updated_at)",
			userID, f.FactID, f.Subject, f.Predicate, f.ObjectValue, f.SourceRef, f.Confidence, f.Version, status, f.CreatedAt, f.UpdatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQLFactStore) SearchFacts(ctx context.Context, q memory.Query) ([]memory.ScoredFact, error) {
	if s == nil || s.db == nil || q.UserID <= 0 || strings.TrimSpace(q.Text) == "" || q.TopK <= 0 {
		return nil, nil
	}
	like := "%" + strings.TrimSpace(q.Text) + "%"
	var rows []struct {
		FactID      string  `db:"fact_id"`
		Subject     string  `db:"subject"`
		Predicate   string  `db:"predicate"`
		ObjectValue string  `db:"object_value"`
		SourceRef   string  `db:"source_ref"`
		Confidence  float64 `db:"confidence"`
		Version     string  `db:"version"`
		Status      string  `db:"status"`
		CreatedAt   int64   `db:"created_at"`
		UpdatedAt   int64   `db:"updated_at"`
	}
	if err := s.db.QueryRowsCtx(ctx, &rows,
		"SELECT fact_id,subject,predicate,object_value,source_ref,confidence,version,status,created_at,updated_at FROM agent_memory_facts WHERE user_id=? AND status='active' AND (subject LIKE ? OR predicate LIKE ? OR object_value LIKE ?) ORDER BY confidence DESC, updated_at DESC LIMIT ?",
		q.UserID, like, like, like, q.TopK,
	); err != nil {
		return nil, err
	}
	out := make([]memory.ScoredFact, 0, len(rows))
	for i, r := range rows {
		out = append(out, memory.ScoredFact{
			Fact: memory.Fact{
				FactID:      r.FactID,
				Subject:     r.Subject,
				Predicate:   r.Predicate,
				ObjectValue: r.ObjectValue,
				SourceRef:   r.SourceRef,
				Confidence:  r.Confidence,
				Version:     r.Version,
				Status:      r.Status,
				CreatedAt:   r.CreatedAt,
				UpdatedAt:   r.UpdatedAt,
			},
			Score:  r.Confidence,
			Source: "memory_mysql",
			Rank:   i + 1,
		})
	}
	return out, nil
}

var _ memory.FactStore = (*MySQLFactStore)(nil)
