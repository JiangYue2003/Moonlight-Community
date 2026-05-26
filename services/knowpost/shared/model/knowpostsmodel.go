package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ KnowPostsModel = (*customKnowPostsModel)(nil)

type (
	// KnowPostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customKnowPostsModel.
	KnowPostsModel interface {
		knowPostsModel
		// 业务专用方法（不使用缓存层，直接走 sqlx 原生连接）。
		ListPublicFeed(ctx context.Context, limit, offset int) ([]*KnowPosts, error)
		ListMyFeed(ctx context.Context, creatorId uint64, limit, offset int) ([]*KnowPosts, error)
		// UpdateInTx 在外部事务里执行 update，**不更新缓存**，调用方写完 outbox 后自行 Invalidate。
		UpdateInTx(ctx context.Context, sess sqlx.Session, data *KnowPosts) error
	}

	customKnowPostsModel struct {
		*defaultKnowPostsModel
	}
)

// NewKnowPostsModel returns a model for the database table.
func NewKnowPostsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) KnowPostsModel {
	return &customKnowPostsModel{
		defaultKnowPostsModel: newKnowPostsModel(conn, c, opts...),
	}
}

// ListPublicFeed 列出 status='published' AND visible='public' 的帖子，按 publish_time DESC。
// 业务侧通常 limit=size+1 用以判定 hasMore。
func (m *customKnowPostsModel) ListPublicFeed(ctx context.Context, limit, offset int) ([]*KnowPosts, error) {
	query := fmt.Sprintf(
		"select %s from %s where status='published' and visible='public' order by publish_time desc limit ? offset ?",
		knowPostsRows, m.table)
	var rows []*KnowPosts
	err := m.QueryRowsNoCacheCtx(ctx, &rows, query, limit, offset)
	return rows, err
}

// ListMyFeed 当前用户的 published+draft（不含 deleted），按 update_time DESC。
func (m *customKnowPostsModel) ListMyFeed(ctx context.Context, creatorId uint64, limit, offset int) ([]*KnowPosts, error) {
	query := fmt.Sprintf(
		"select %s from %s where creator_id=? and status<>'deleted' order by update_time desc limit ? offset ?",
		knowPostsRows, m.table)
	var rows []*KnowPosts
	err := m.QueryRowsNoCacheCtx(ctx, &rows, query, creatorId, limit, offset)
	return rows, err
}

// UpdateInTx 与 defaultModel.Update 等价，但走外部 sess（同事务）。
// 缓存失效由调用方在事务提交后处理（避免事务回滚后仍清掉缓存）。
func (m *customKnowPostsModel) UpdateInTx(ctx context.Context, sess sqlx.Session, data *KnowPosts) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, knowPostsRowsWithPlaceHolder)
	_, err := sess.ExecCtx(ctx, query,
		data.TagId, data.Tags, data.Title, data.Description, data.ContentUrl,
		data.ContentObjectKey, data.ContentEtag, data.ContentSize, data.ContentSha256,
		data.CreatorId, data.IsTop, data.Type, data.Visible, data.ImgUrls, data.VideoUrl,
		data.Status, data.PublishTime, data.Id)
	return err
}
