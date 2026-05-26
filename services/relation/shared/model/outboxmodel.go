package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OutboxModel = (*customOutboxModel)(nil)

type (
	OutboxModel interface {
		outboxModel

		// InsertInTx 在事务内写一行 outbox（与 following 表的 INSERT/UPDATE 同事务）。
		InsertInTx(ctx context.Context, sess sqlx.Session, id int64, aggregateType string, aggregateId int64, eventType, payload string) error
	}

	customOutboxModel struct {
		*defaultOutboxModel
	}
)

func NewOutboxModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OutboxModel {
	return &customOutboxModel{
		defaultOutboxModel: newOutboxModel(conn, c, opts...),
	}
}

func (m *customOutboxModel) InsertInTx(ctx context.Context, sess sqlx.Session, id int64, aggregateType string, aggregateId int64, eventType, payload string) error {
	q := fmt.Sprintf(
		"insert into %s (id, aggregate_type, aggregate_id, type, payload, created_at) values (?, ?, ?, ?, ?, NOW(3))",
		m.table)
	_, err := sess.ExecCtx(ctx, q, id, aggregateType, aggregateId, eventType, payload)
	return err
}
