// Package txx 提供 go-zero sqlx 事务的薄封装。
//
// 与 conn.TransactCtx 完全等价，但有两点改进：
//  1. 包装 fn 的签名让 logic 层不必直接 import sqlx
//  2. 留有 hook 点便于后续接入 logx 慢事务统计 / metrics
package txx

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// Fn 事务体函数；返回非 nil 错误自动 ROLLBACK。
type Fn func(ctx context.Context, sess sqlx.Session) error

// WithTx 在一个 BEGIN/COMMIT 边界内执行 fn。
// 调用方在 fn 内必须使用传入的 sess 而不是外层 conn，否则数据库写入不在事务内。
func WithTx(ctx context.Context, conn sqlx.SqlConn, fn Fn) error {
	return conn.TransactCtx(ctx, func(ctx context.Context, sess sqlx.Session) error {
		return fn(ctx, sess)
	})
}
