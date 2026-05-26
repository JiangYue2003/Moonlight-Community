package providers

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

type MySQLPreferenceStore struct {
	db sqlx.SqlConn
}

func NewMySQLPreferenceStore(db sqlx.SqlConn) *MySQLPreferenceStore {
	return &MySQLPreferenceStore{db: db}
}

func (s *MySQLPreferenceStore) UpsertPreferences(ctx context.Context, userID int64, prefs []memory.Preference) error {
	if s == nil || s.db == nil || userID <= 0 || len(prefs) == 0 {
		return nil
	}
	for _, p := range prefs {
		if strings.TrimSpace(p.PreferenceID) == "" || strings.TrimSpace(p.Kind) == "" || strings.TrimSpace(p.Content) == "" {
			continue
		}
		status := strings.TrimSpace(p.Status)
		if status == "" {
			status = "active"
		}
		_, err := s.db.ExecCtx(ctx,
			"INSERT INTO agent_memory_preferences (user_id,preference_id,kind,content,confidence,source,status,last_seen_at,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE content=VALUES(content),confidence=VALUES(confidence),source=VALUES(source),status=VALUES(status),last_seen_at=VALUES(last_seen_at),updated_at=VALUES(updated_at)",
			userID, p.PreferenceID, p.Kind, p.Content, p.Confidence, p.Source, status, p.LastSeenAt, p.CreatedAt, p.UpdatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQLPreferenceStore) ListActivePreferences(ctx context.Context, userID int64, limit int) ([]memory.Preference, error) {
	if s == nil || s.db == nil || userID <= 0 || limit <= 0 {
		return nil, nil
	}
	var rows []struct {
		PreferenceID string  `db:"preference_id"`
		Kind         string  `db:"kind"`
		Content      string  `db:"content"`
		Confidence   float64 `db:"confidence"`
		Source       string  `db:"source"`
		Status       string  `db:"status"`
		LastSeenAt   int64   `db:"last_seen_at"`
		CreatedAt    int64   `db:"created_at"`
		UpdatedAt    int64   `db:"updated_at"`
	}
	if err := s.db.QueryRowsCtx(ctx, &rows,
		"SELECT preference_id,kind,content,confidence,source,status,last_seen_at,created_at,updated_at FROM agent_memory_preferences WHERE user_id=? AND status='active' ORDER BY confidence DESC, updated_at DESC LIMIT ?",
		userID, limit,
	); err != nil {
		return nil, err
	}
	out := make([]memory.Preference, 0, len(rows))
	for _, r := range rows {
		out = append(out, memory.Preference{
			PreferenceID: r.PreferenceID,
			Kind:         r.Kind,
			Content:      r.Content,
			Confidence:   r.Confidence,
			Source:       r.Source,
			Status:       r.Status,
			LastSeenAt:   r.LastSeenAt,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		})
	}
	return out, nil
}

var _ memory.PreferenceStore = (*MySQLPreferenceStore)(nil)
