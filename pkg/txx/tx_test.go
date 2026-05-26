package txx

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 用 sqlmock 套出一个 sqlx.SqlConn 实例，方便断言事务行为。
func newConn(t *testing.T) (sqlx.SqlConn, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	conn := sqlx.NewSqlConnFromDB(db)
	return conn, mock
}

func TestWithTx_CommitsOnNilError(t *testing.T) {
	conn, mock := newConn(t)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := WithTx(context.Background(), conn, func(ctx context.Context, sess sqlx.Session) error {
		_, err := sess.ExecCtx(ctx, "INSERT INTO t VALUES (?)", 1)
		return err
	})
	if err != nil {
		t.Fatalf("WithTx: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations: %v", err)
	}
}

func TestWithTx_RollsBackOnError(t *testing.T) {
	conn, mock := newConn(t)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	want := errors.New("biz failure")
	err := WithTx(context.Background(), conn, func(ctx context.Context, sess sqlx.Session) error {
		_, _ = sess.ExecCtx(ctx, "INSERT INTO t VALUES (?)", 1)
		return want
	})
	if !errors.Is(err, want) {
		t.Fatalf("err not propagated: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations: %v", err)
	}
}

func TestWithTx_NilFnError_StillCommit(t *testing.T) {
	// fn 完全没操作也应该 BEGIN+COMMIT；不会跳过事务。
	conn, mock := newConn(t)
	mock.ExpectBegin()
	mock.ExpectCommit()
	if err := WithTx(context.Background(), conn, func(ctx context.Context, sess sqlx.Session) error {
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations: %v", err)
	}
}
