package model

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const feedAuthorDeliveryTierTable = "`feed_author_delivery_tier`"

type FeedAuthorDeliveryTierModel interface {
	IsBigV(ctx context.Context, authorID int64) (bool, error)
	BatchBigVs(ctx context.Context, authorIDs []int64) (map[int64]struct{}, error)
	PromoteBigV(ctx context.Context, authorID int64, reason string, observedFollowers int64) error
}

func (m *feedAuthorDeliveryTierModel) IsBigV(ctx context.Context, authorID int64) (bool, error) {
	query := "select author_id from `feed_author_delivery_tier` where tier=1 and author_id=? limit 1"
	var row feedAuthorDeliveryTierAuthor
	err := m.conn.QueryRowCtx(ctx, &row, query, authorID)
	if errors.Is(err, sqlx.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *feedAuthorDeliveryTierModel) PromoteBigV(
	ctx context.Context,
	authorID int64,
	reason string,
	observedFollowers int64,
) error {
	query := "insert into `feed_author_delivery_tier` " +
		"(author_id,tier,promotion_reason,observed_followers,promoted_at,updated_at) " +
		"values (?,1,?,?,NOW(3),NOW(3)) on duplicate key update tier=GREATEST(tier,1)"
	_, err := m.conn.ExecCtx(ctx, query, authorID, reason, observedFollowers)
	return err
}

type feedAuthorDeliveryTierModel struct {
	conn sqlx.SqlConn
}

type feedAuthorDeliveryTierAuthor struct {
	AuthorID int64 `db:"author_id"`
}

func NewFeedAuthorDeliveryTierModel(conn sqlx.SqlConn) FeedAuthorDeliveryTierModel {
	return &feedAuthorDeliveryTierModel{conn: conn}
}

func (m *feedAuthorDeliveryTierModel) BatchBigVs(ctx context.Context, authorIDs []int64) (map[int64]struct{}, error) {
	bigVs := make(map[int64]struct{}, len(authorIDs))
	if len(authorIDs) == 0 {
		return bigVs, nil
	}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(authorIDs)), ",")
	query := fmt.Sprintf(
		"select author_id from %s where tier=1 and author_id in (%s)",
		feedAuthorDeliveryTierTable,
		placeholders,
	)
	args := make([]any, len(authorIDs))
	for i, authorID := range authorIDs {
		args[i] = authorID
	}

	var rows []feedAuthorDeliveryTierAuthor
	if err := m.conn.QueryRowsCtx(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	for _, row := range rows {
		bigVs[row.AuthorID] = struct{}{}
	}
	return bigVs, nil
}
