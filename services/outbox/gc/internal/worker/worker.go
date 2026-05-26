package worker

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/outbox/gc/internal/config"
)

// Worker 定时分批删除 outbox 表中超过 RetainDays 天的记录。
type Worker struct {
	cfg config.Config
	db  sqlx.SqlConn
}

func New(cfg config.Config, db sqlx.SqlConn) *Worker {
	return &Worker{cfg: cfg, db: db}
}

func (w *Worker) Run(ctx context.Context) {
	interval := time.Duration(w.cfg.IntervalHours) * time.Hour
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := w.gc(ctx); err != nil {
				logx.Errorf("outbox-gc: error: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) gc(ctx context.Context) error {
	cutoff := time.Now().Add(-time.Duration(w.cfg.RetainDays) * 24 * time.Hour)
	batchInterval := time.Duration(w.cfg.BatchIntervalMs) * time.Millisecond
	total := int64(0)

	for {
		result, err := w.db.ExecCtx(ctx,
			"DELETE FROM outbox WHERE created_at < ? ORDER BY id LIMIT ?",
			cutoff, w.cfg.BatchSize)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		total += affected
		if affected == 0 {
			break
		}
		if batchInterval > 0 {
			select {
			case <-time.After(batchInterval):
			case <-ctx.Done():
				logx.Infof("outbox-gc: interrupted, deleted %d rows", total)
				return nil
			}
		}
	}
	if total > 0 {
		logx.Infof("outbox-gc: deleted %d rows older than %d days", total, w.cfg.RetainDays)
	}
	return nil
}
