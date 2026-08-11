package model

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/zeromicro/go-zero/core/stores/cache"
	zeroredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestFindPublishedFeedByIDsUsesOneFilteredQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })

	columns := make([]string, len(knowPostsFieldNames))
	for i, name := range knowPostsFieldNames {
		columns[i] = strings.Trim(name, "`")
	}
	now := time.Now()
	rows := sqlmock.NewRows(columns).AddRow(
		uint64(2), sql.NullInt64{}, `["go"]`, "two", "desc", "url", "key", "etag",
		int64(10), "sha", uint64(9), int64(0), "image_text", "public", `["img"]`, "",
		"published", now, now, now,
	)
	query := "select " + knowPostsRows + " from `know_posts` " +
		"where `id` in (?,?) and status='published' and visible in ('public','followers')"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(uint64(1), uint64(2)).
		WillReturnRows(rows)

	mr := miniredis.RunT(t)
	m := NewKnowPostsModel(sqlx.NewSqlConnFromDB(db), cache.CacheConf{{
		RedisConf: zeroredis.RedisConf{Host: mr.Addr(), Type: "node"},
		Weight:    100,
	}})
	got, err := m.FindPublishedFeedByIDs(context.Background(), []uint64{1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Id != 2 {
		t.Fatalf("unexpected rows: %+v", got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
