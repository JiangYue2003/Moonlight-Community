package model

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestFeedAuthorDeliveryTierModelBatchBigVs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	query := "select author_id from `feed_author_delivery_tier` where tier=1 and author_id in (?,?)"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(11), int64(12)).
		WillReturnRows(sqlmock.NewRows([]string{"author_id"}).AddRow(int64(12)))

	m := NewFeedAuthorDeliveryTierModel(sqlx.NewSqlConnFromDB(db))
	got, err := m.BatchBigVs(context.Background(), []int64{11, 12})
	require.NoError(t, err)
	require.Equal(t, map[int64]struct{}{12: {}}, got)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestFeedAuthorDeliveryTierModelPromoteBigVIsMonotonicAndIdempotent(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	query := "insert into `feed_author_delivery_tier` " +
		"(author_id,tier,promotion_reason,observed_followers,promoted_at,updated_at) " +
		"values (?,1,?,?,NOW(3),NOW(3)) on duplicate key update tier=GREATEST(tier,1)"
	for range 2 {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(int64(12), "counter_threshold", int64(1001)).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	m := NewFeedAuthorDeliveryTierModel(sqlx.NewSqlConnFromDB(db))
	require.NoError(t, m.PromoteBigV(context.Background(), 12, "counter_threshold", 1001))
	require.NoError(t, m.PromoteBigV(context.Background(), 12, "counter_threshold", 1001))
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestFeedAuthorDeliveryTierModelIsBigVTreatsMissingRowAsNormal(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	query := "select author_id from `feed_author_delivery_tier` where tier=1 and author_id=? limit 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(11)).
		WillReturnRows(sqlmock.NewRows([]string{"author_id"}))
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(12)).
		WillReturnRows(sqlmock.NewRows([]string{"author_id"}).AddRow(int64(12)))

	m := NewFeedAuthorDeliveryTierModel(sqlx.NewSqlConnFromDB(db))
	bigV, err := m.IsBigV(context.Background(), 11)
	require.NoError(t, err)
	require.False(t, bigV)
	bigV, err = m.IsBigV(context.Background(), 12)
	require.NoError(t, err)
	require.True(t, bigV)
	require.NoError(t, mock.ExpectationsWereMet())
}
