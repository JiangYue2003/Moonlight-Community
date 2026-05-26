package worker

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/outbox/gc/internal/config"
)

func newTestWorker(t *testing.T, db sqlx.SqlConn, retainDays, batchSize int) *Worker {
	t.Helper()
	return New(config.Config{
		RetainDays:      retainDays,
		BatchSize:       batchSize,
		BatchIntervalMs: 0,
	}, db)
}

func TestGC_DeletesOldRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 第一批删 2 行，第二批删 0 行（结束）
	mock.ExpectExec("DELETE FROM outbox").
		WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectExec("DELETE FROM outbox").
		WillReturnResult(sqlmock.NewResult(0, 0))

	w := newTestWorker(t, sqlx.NewSqlConnFromDB(db), 7, 500)
	if err := w.gc(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGC_PreservesRecentRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 没有超期记录，第一批就返回 0
	mock.ExpectExec("DELETE FROM outbox").
		WillReturnResult(sqlmock.NewResult(0, 0))

	w := newTestWorker(t, sqlx.NewSqlConnFromDB(db), 7, 500)
	if err := w.gc(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGC_BatchesMultipleRounds(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// batchSize=2，第一批删 2 行，第二批删 1 行，第三批删 0 行
	mock.ExpectExec("DELETE FROM outbox").WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectExec("DELETE FROM outbox").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("DELETE FROM outbox").WillReturnResult(sqlmock.NewResult(0, 0))

	w := newTestWorker(t, sqlx.NewSqlConnFromDB(db), 7, 2)
	if err := w.gc(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

// 防止 go vet 抱怨未用 import
var _ = time.Now
var _ = sql.ErrNoRows
