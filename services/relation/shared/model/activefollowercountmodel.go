package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ActiveFollowerCount struct {
	AuthorID int64 `db:"to_user_id"`
	Count    int64 `db:"active_count"`
}

type ActiveFollowerCountModel interface {
	BatchCountActive(ctx context.Context, authorIDs []int64) (map[int64]int64, error)
	PageActiveCounts(ctx context.Context, afterAuthorID int64, limit int) ([]ActiveFollowerCount, error)
}

func (m *activeFollowerCountModel) PageActiveCounts(
	ctx context.Context,
	afterAuthorID int64,
	limit int,
) ([]ActiveFollowerCount, error) {
	query := "select to_user_id, count(1) as active_count from `follower` " +
		"where rel_status=1 and to_user_id>? group by to_user_id order by to_user_id asc limit ?"
	var rows []ActiveFollowerCount
	if err := m.conn.QueryRowsCtx(ctx, &rows, query, afterAuthorID, limit); err != nil {
		return nil, err
	}
	return rows, nil
}

type activeFollowerCountModel struct {
	conn sqlx.SqlConn
}

func NewActiveFollowerCountModel(conn sqlx.SqlConn) ActiveFollowerCountModel {
	return &activeFollowerCountModel{conn: conn}
}

func (m *activeFollowerCountModel) BatchCountActive(
	ctx context.Context,
	authorIDs []int64,
) (map[int64]int64, error) {
	counts := make(map[int64]int64, len(authorIDs))
	for _, authorID := range authorIDs {
		counts[authorID] = 0
	}
	if len(authorIDs) == 0 {
		return counts, nil
	}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(authorIDs)), ",")
	query := fmt.Sprintf(
		"select to_user_id, count(1) as active_count from `follower` "+
			"where rel_status=1 and to_user_id in (%s) group by to_user_id",
		placeholders,
	)
	args := make([]any, len(authorIDs))
	for i, authorID := range authorIDs {
		args[i] = authorID
	}

	var rows []ActiveFollowerCount
	if err := m.conn.QueryRowsCtx(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	for _, row := range rows {
		counts[row.AuthorID] = row.Count
	}
	return counts, nil
}
