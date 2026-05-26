package providers

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

func TestMySQLPreferenceStoreUpsertPreferences(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	conn := sqlx.NewSqlConnFromDB(db)
	store := NewMySQLPreferenceStore(conn)
	prefs := []memory.Preference{
		{
			PreferenceID: "p1",
			Kind:         "response_style",
			Content:      "回答偏扁平叙述",
			Confidence:   0.9,
			Source:       "explicit_pin",
			Status:       "active",
			CreatedAt:    1,
			UpdatedAt:    2,
			LastSeenAt:   2,
		},
	}

	mock.ExpectExec("INSERT INTO agent_memory_preferences").
		WithArgs(int64(1), "p1", "response_style", "回答偏扁平叙述", 0.9, "explicit_pin", "active", int64(2), int64(1), int64(2)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := store.UpsertPreferences(context.Background(), 1, prefs); err != nil {
		t.Fatalf("upsert failed: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestMySQLPreferenceStoreListActivePreferences(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	conn := sqlx.NewSqlConnFromDB(db)
	store := NewMySQLPreferenceStore(conn)
	rows := sqlmock.NewRows([]string{"preference_id", "kind", "content", "confidence", "source", "status", "last_seen_at", "created_at", "updated_at"}).
		AddRow("p1", "response_style", "回答偏扁平叙述", 0.9, "explicit_pin", "active", 2, 1, 2)

	mock.ExpectQuery("SELECT preference_id,kind,content,confidence,source,status,last_seen_at,created_at,updated_at FROM agent_memory_preferences").
		WithArgs(int64(1), 3).
		WillReturnRows(rows)

	items, err := store.ListActivePreferences(context.Background(), 1, 3)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("want 1 item got %d", len(items))
	}
	if items[0].Content != "回答偏扁平叙述" {
		t.Fatalf("unexpected content: %s", items[0].Content)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
