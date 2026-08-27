package main

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCaptureMutationMySQLSnapshotUsesOwnedAuthorsAndStableRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users: []benchmarkUser{
			{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"},
			{ID: 8, Email: "feed-loadtest+run-42-00001@example.invalid"},
		},
		NormalAuthors: []int64{7}, BigVAuthors: []int64{8}, Readers: []int64{7},
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email FROM users WHERE id IN (?,?)")).
		WithArgs(int64(7), int64(8)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(7, manifest.Users[0].Email).
			AddRow(8, manifest.Users[1].Email))
	mock.ExpectQuery("SELECT id, creator_id, .* FROM know_posts WHERE creator_id IN \\(\\?,\\?\\) ORDER BY id").
		WithArgs(int64(7), int64(8)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "status"}).
			AddRow(10, 7, "published").
			AddRow(20, 8, "published"))
	mock.ExpectQuery("SELECT id, aggregate_type, aggregate_id, type, payload, created_at FROM outbox WHERE aggregate_type = \\? AND aggregate_id IN \\(\\?,\\?\\) ORDER BY id").
		WithArgs("knowpost", int64(10), int64(20)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "aggregate_type", "aggregate_id", "type", "payload", "created_at"}).
			AddRow(100, "knowpost", 10, "KnowPostPublished", `{"post_id":10}`, "2026-08-18 00:00:00").
			AddRow(200, "knowpost", 20, "KnowPostPublished", `{"post_id":20}`, "2026-08-18 00:00:01"))
	mock.ExpectRollback()

	snapshot, err := captureMutationMySQLSnapshot(context.Background(), db, manifest)

	require.NoError(t, err)
	require.Equal(t, []int64{7, 8}, snapshot.AuthorIDs)
	require.Equal(t, []int64{10, 20}, snapshot.PostIDs)
	require.Equal(t, []int64{100, 200}, snapshot.OutboxIDs)
	require.NotEmpty(t, snapshot.PostFingerprint)
	require.NotEmpty(t, snapshot.OutboxFingerprint)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestFingerprintMutationRowsIncludesColumnNamesNullsAndValues(t *testing.T) {
	first := fingerprintMutationRowsForTest(t, []string{"id", "title"}, [][]any{{int64(10), nil}})
	same := fingerprintMutationRowsForTest(t, []string{"id", "title"}, [][]any{{int64(10), nil}})
	changedNull := fingerprintMutationRowsForTest(t, []string{"id", "title"}, [][]any{{int64(10), ""}})
	changedColumn := fingerprintMutationRowsForTest(t, []string{"id", "description"}, [][]any{{int64(10), nil}})

	require.Equal(t, first.fingerprint, same.fingerprint)
	require.Equal(t, []int64{10}, first.ids)
	require.NotEqual(t, first.fingerprint, changedNull.fingerprint)
	require.NotEqual(t, first.fingerprint, changedColumn.fingerprint)
}

func TestRestoreMutationMySQLSnapshotDeletesOnlyPostsAfterBaseline(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users:         []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
		NormalAuthors: []int64{7}, Readers: []int64{7},
	}
	postBaseline := fingerprintMutationRowsForTest(t,
		[]string{"id", "creator_id", "status"}, [][]any{{int64(10), int64(7), "published"}},
	)
	outboxBaseline := fingerprintMutationRowsForTest(t,
		[]string{"id", "aggregate_type", "aggregate_id", "type", "payload", "created_at"},
		[][]any{{int64(100), "knowpost", int64(10), "KnowPostPublished", `{"post_id":10}`, "2026-08-18 00:00:00"}},
	)
	snapshot := mutationMySQLSnapshot{
		AuthorIDs: []int64{7}, PostIDs: postBaseline.ids, PostFingerprint: postBaseline.fingerprint,
		OutboxIDs: outboxBaseline.ids, OutboxFingerprint: outboxBaseline.fingerprint,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email FROM users WHERE id IN (?)")).
		WithArgs(int64(7)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(7, manifest.Users[0].Email))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM know_posts WHERE creator_id IN (?) ORDER BY id")).
		WithArgs(int64(7)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10).AddRow(30))
	mock.ExpectQuery("SELECT id, creator_id, .* FROM know_posts WHERE id IN \\(\\?\\) ORDER BY id").
		WithArgs(int64(10)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "creator_id", "status"}).AddRow(10, 7, "published"))
	mock.ExpectQuery("SELECT id, aggregate_type, aggregate_id, type, payload, created_at FROM outbox WHERE id IN \\(\\?\\) ORDER BY id").
		WithArgs(int64(100)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "aggregate_type", "aggregate_id", "type", "payload", "created_at"}).
			AddRow(100, "knowpost", 10, "KnowPostPublished", `{"post_id":10}`, "2026-08-18 00:00:00"))
	mock.ExpectQuery("SELECT id FROM outbox WHERE aggregate_type = \\? AND aggregate_id IN \\(\\?\\) ORDER BY id").
		WithArgs("knowpost", int64(30)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(300))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM outbox WHERE id IN (?)")).
		WithArgs(int64(300)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM know_posts WHERE id IN (?)")).
		WithArgs(int64(30)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	result, err := restoreMutationMySQLSnapshot(context.Background(), db, manifest, snapshot)

	require.NoError(t, err)
	require.Equal(t, []int64{30}, result.PostIDs)
	require.Equal(t, []int64{300}, result.OutboxIDs)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMutationIDsAfterBaselineRejectsMissingBaselineAndIsIdempotent(t *testing.T) {
	ids, err := mutationIDsAfterBaseline([]int64{10, 20, 30}, []int64{10, 20})
	require.NoError(t, err)
	require.Equal(t, []int64{30}, ids)

	ids, err = mutationIDsAfterBaseline([]int64{10, 20}, []int64{10, 20})
	require.NoError(t, err)
	require.Empty(t, ids)

	_, err = mutationIDsAfterBaseline([]int64{10, 30}, []int64{10, 20})
	require.ErrorContains(t, err, "missing baseline post 20")
}

func TestFingerprintMutationRowsRejectsNonPositiveIDWithoutFormattingArtifact(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectQuery("SELECT invalid_row").
		WillReturnRows(sqlmock.NewRows([]string{"id", "value"}).AddRow(0, "invalid"))
	rows, err := db.QueryContext(context.Background(), "SELECT invalid_row")
	require.NoError(t, err)
	defer rows.Close()

	_, err = fingerprintMutationRows(rows, "id")

	require.ErrorContains(t, err, "must be positive")
	require.NotContains(t, err.Error(), "%!w")
}

func fingerprintMutationRowsForTest(t *testing.T, columns []string, values [][]any) mutationRowsFingerprint {
	t.Helper()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	rows := sqlmock.NewRows(columns)
	for _, row := range values {
		driverRow := make([]driver.Value, len(row))
		for index := range row {
			driverRow[index] = row[index]
		}
		rows.AddRow(driverRow...)
	}
	mock.ExpectQuery("SELECT test_rows").WillReturnRows(rows)
	resultRows, err := db.QueryContext(context.Background(), "SELECT test_rows")
	require.NoError(t, err)
	result, err := fingerprintMutationRows(resultRows, "id")
	require.NoError(t, err)
	require.NoError(t, resultRows.Close())
	require.NoError(t, mock.ExpectationsWereMet())
	return result
}
