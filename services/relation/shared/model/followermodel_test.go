package model

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestActiveFollowerCountModelBatchReturnsCompleteCounts(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	query := "select to_user_id, count(1) as active_count from `follower` " +
		"where rel_status=1 and to_user_id in (?,?,?) group by to_user_id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(11), int64(12), int64(13)).
		WillReturnRows(sqlmock.NewRows([]string{"to_user_id", "active_count"}).
			AddRow(int64(11), int64(1000)).
			AddRow(int64(13), int64(1001)))

	m := NewActiveFollowerCountModel(sqlx.NewSqlConnFromDB(db))
	got, err := m.BatchCountActive(context.Background(), []int64{11, 12, 13})
	require.NoError(t, err)
	require.Equal(t, map[int64]int64{11: 1000, 12: 0, 13: 1001}, got)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestActiveFollowerCountModelPagesByStableAuthorCursor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	query := "select to_user_id, count(1) as active_count from `follower` " +
		"where rel_status=1 and to_user_id>? group by to_user_id order by to_user_id asc limit ?"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(100), 2).
		WillReturnRows(sqlmock.NewRows([]string{"to_user_id", "active_count"}).
			AddRow(int64(101), int64(1000)).
			AddRow(int64(105), int64(1001)))

	m := NewActiveFollowerCountModel(sqlx.NewSqlConnFromDB(db))
	got, err := m.PageActiveCounts(context.Background(), 100, 2)
	require.NoError(t, err)
	require.Equal(t, []ActiveFollowerCount{
		{AuthorID: 101, Count: 1000},
		{AuthorID: 105, Count: 1001},
	}, got)
	require.NoError(t, mock.ExpectationsWereMet())
}
