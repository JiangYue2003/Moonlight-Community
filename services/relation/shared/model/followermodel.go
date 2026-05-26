package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowerModel = (*customFollowerModel)(nil)

type (
	FollowerModel interface {
		followerModel

		UpsertActive(ctx context.Context, sess sqlx.Session, id, toUserId, fromUserId int64) error
		MarkInactive(ctx context.Context, sess sqlx.Session, toUserId, fromUserId int64) (int64, error)
		PageActive(ctx context.Context, toUserId int64, limit, offset int, cursor int64) ([]*Follower, error)
		CountActive(ctx context.Context, toUserId int64) (int64, error)
	}

	customFollowerModel struct {
		*defaultFollowerModel
	}
)

func NewFollowerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowerModel {
	return &customFollowerModel{
		defaultFollowerModel: newFollowerModel(conn, c, opts...),
	}
}

func (m *customFollowerModel) UpsertActive(ctx context.Context, sess sqlx.Session, id, toUserId, fromUserId int64) error {
	q := fmt.Sprintf(
		"insert into %s (id, to_user_id, from_user_id, rel_status, created_at, updated_at) "+
			"values (?, ?, ?, 1, NOW(3), NOW(3)) "+
			"on duplicate key update rel_status=1, updated_at=NOW(3)",
		m.table)
	_, err := sess.ExecCtx(ctx, q, id, toUserId, fromUserId)
	return err
}

func (m *customFollowerModel) MarkInactive(ctx context.Context, sess sqlx.Session, toUserId, fromUserId int64) (int64, error) {
	q := fmt.Sprintf("update %s set rel_status=0, updated_at=NOW(3) where to_user_id=? and from_user_id=? and rel_status=1", m.table)
	res, err := sess.ExecCtx(ctx, q, toUserId, fromUserId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (m *customFollowerModel) PageActive(ctx context.Context, toUserId int64, limit, offset int, cursor int64) ([]*Follower, error) {
	var rows []*Follower
	if cursor > 0 {
		q := fmt.Sprintf(
			"select %s from %s where to_user_id=? and rel_status=1 and UNIX_TIMESTAMP(created_at)*1000 < ? "+
				"order by created_at desc limit ?",
			followerRows, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &rows, q, toUserId, cursor, limit)
		return rows, err
	}
	q := fmt.Sprintf(
		"select %s from %s where to_user_id=? and rel_status=1 order by created_at desc limit ? offset ?",
		followerRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &rows, q, toUserId, limit, offset)
	return rows, err
}

func (m *customFollowerModel) CountActive(ctx context.Context, toUserId int64) (int64, error) {
	q := fmt.Sprintf("select count(1) from %s where to_user_id=? and rel_status=1", m.table)
	var n int64
	err := m.QueryRowNoCacheCtx(ctx, &n, q, toUserId)
	if err == sqlx.ErrNotFound {
		return 0, nil
	}
	return n, err
}
