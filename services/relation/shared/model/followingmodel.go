package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowingModel = (*customFollowingModel)(nil)

type (
	// FollowingModel 业务定制接口：在 goctl 生成接口之上加自定义查询。
	FollowingModel interface {
		followingModel

		// ExistsActive 当前 from→to 的关注关系是否处于 active（rel_status=1）。
		ExistsActive(ctx context.Context, fromUserId, toUserId int64) (bool, error)

		// UpsertActive 关注：INSERT 或把 rel_status 置 1（取消后再次关注的语义）。
		// 该操作必须在事务里调用，与 outbox 同步生效。
		UpsertActive(ctx context.Context, sess sqlx.Session, id, fromUserId, toUserId int64) error

		// MarkInactive 取关：UPDATE rel_status=0；返回受影响行数（=0 视为本就未关注）。
		MarkInactive(ctx context.Context, sess sqlx.Session, fromUserId, toUserId int64) (int64, error)

		// PageActive 关注列表分页：按 created_at desc。offset>=0 时使用 offset；否则用 cursor 走"created_at < cursor"。
		PageActive(ctx context.Context, fromUserId int64, limit, offset int, cursor int64) ([]*Following, error)
	}

	customFollowingModel struct {
		*defaultFollowingModel
	}
)

// NewFollowingModel returns a model for the database table.
func NewFollowingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowingModel {
	return &customFollowingModel{
		defaultFollowingModel: newFollowingModel(conn, c, opts...),
	}
}

func (m *customFollowingModel) ExistsActive(ctx context.Context, fromUserId, toUserId int64) (bool, error) {
	q := fmt.Sprintf("select 1 from %s where from_user_id=? and to_user_id=? and rel_status=1 limit 1", m.table)
	var one int
	err := m.QueryRowNoCacheCtx(ctx, &one, q, fromUserId, toUserId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *customFollowingModel) UpsertActive(ctx context.Context, sess sqlx.Session, id, fromUserId, toUserId int64) error {
	// ON DUPLICATE KEY UPDATE：把已存在但被取消的关系重新置为 1，刷新 updated_at。
	q := fmt.Sprintf(
		"insert into %s (id, from_user_id, to_user_id, rel_status, created_at, updated_at) "+
			"values (?, ?, ?, 1, NOW(3), NOW(3)) "+
			"on duplicate key update rel_status=1, updated_at=NOW(3)",
		m.table)
	_, err := sess.ExecCtx(ctx, q, id, fromUserId, toUserId)
	return err
}

func (m *customFollowingModel) MarkInactive(ctx context.Context, sess sqlx.Session, fromUserId, toUserId int64) (int64, error) {
	q := fmt.Sprintf("update %s set rel_status=0, updated_at=NOW(3) where from_user_id=? and to_user_id=? and rel_status=1", m.table)
	res, err := sess.ExecCtx(ctx, q, fromUserId, toUserId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (m *customFollowingModel) PageActive(ctx context.Context, fromUserId int64, limit, offset int, cursor int64) ([]*Following, error) {
	var rows []*Following
	if cursor > 0 {
		q := fmt.Sprintf(
			"select %s from %s where from_user_id=? and rel_status=1 and UNIX_TIMESTAMP(created_at)*1000 < ? "+
				"order by created_at desc limit ?",
			followingRows, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &rows, q, fromUserId, cursor, limit)
		return rows, err
	}
	q := fmt.Sprintf(
		"select %s from %s where from_user_id=? and rel_status=1 order by created_at desc limit ? offset ?",
		followingRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &rows, q, fromUserId, limit, offset)
	return rows, err
}
