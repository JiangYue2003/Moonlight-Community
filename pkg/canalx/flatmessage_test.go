package canalx

import (
	"encoding/json"
	"strings"
	"testing"
)

// 来自 Canal 1.1.7 实测的 flatMessage 样本（INSERT outbox 一行）
const sampleInsertOutbox = `{
  "database": "zhiguang",
  "table": "outbox",
  "type": "INSERT",
  "isDdl": false,
  "data": [{
    "id": "9876543210",
    "aggregate_type": "following",
    "aggregate_id": "42",
    "type": "FollowCreated",
    "payload": "{\"type\":\"FollowCreated\",\"fromUserId\":42,\"toUserId\":7,\"id\":9876543210}",
    "created_at": "2026-05-15 10:00:00.123"
  }],
  "old": null,
  "ts": 1715760000123,
  "es": 1715760000100,
  "sql": "",
  "pkNames": ["id"]
}`

const sampleUpdateOtherTable = `{
  "database": "zhiguang", "table": "users", "type": "UPDATE", "isDdl": false,
  "data": [{"id": "1", "nickname": "x"}],
  "old": [{"nickname": "y"}],
  "ts": 1, "es": 1
}`

const sampleDdl = `{
  "database": "zhiguang", "table": "outbox", "type": "QUERY", "isDdl": true,
  "data": [], "ts": 1, "es": 1, "sql": "ALTER TABLE outbox ADD COLUMN x INT"
}`

const sampleDelete = `{
  "database": "zhiguang", "table": "outbox", "type": "DELETE", "isDdl": false,
  "data": [{"id": "1", "aggregate_type": "x", "aggregate_id": "1", "type": "y", "payload": "{}"}],
  "ts": 1, "es": 1
}`

func TestParseFlat_RejectsEmpty(t *testing.T) {
	if _, err := ParseFlat(nil); err == nil {
		t.Fatal("nil input should error")
	}
	if _, err := ParseFlat([]byte("")); err == nil {
		t.Fatal("empty input should error")
	}
}

func TestParseFlat_RejectsBadJson(t *testing.T) {
	if _, err := ParseFlat([]byte("not json")); err == nil {
		t.Fatal("bad json should error")
	}
	if _, err := ParseFlat([]byte("{invalid")); err == nil {
		t.Fatal("malformed json should error")
	}
}

func TestParseFlat_OutboxInsertSample(t *testing.T) {
	m, err := ParseFlat([]byte(sampleInsertOutbox))
	if err != nil {
		t.Fatal(err)
	}
	if m.Database != "zhiguang" || m.Table != "outbox" || m.Type != "INSERT" {
		t.Fatalf("header drift: %+v", m)
	}
	if m.IsDdl {
		t.Fatal("isDdl should be false")
	}
	if len(m.Data) != 1 {
		t.Fatalf("expect 1 data row, got %d", len(m.Data))
	}
	row := m.Data[0]
	if row["aggregate_type"] != "following" || row["type"] != "FollowCreated" {
		t.Fatalf("row drift: %+v", row)
	}
	if !strings.Contains(row["payload"], "FollowCreated") {
		t.Fatalf("payload should contain event type: %q", row["payload"])
	}
}

func TestExtractOutboxRows_OnlyOutboxInsertUpdate(t *testing.T) {
	m, _ := ParseFlat([]byte(sampleInsertOutbox))
	rows := ExtractOutboxRows(m)
	if len(rows) != 1 {
		t.Fatalf("expect 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Id != 9876543210 || r.AggregateId != 42 || r.AggregateType != "following" || r.Type != "FollowCreated" {
		t.Fatalf("row mapping drift: %+v", r)
	}

	m2, _ := ParseFlat([]byte(sampleUpdateOtherTable))
	if rs := ExtractOutboxRows(m2); len(rs) != 0 {
		t.Fatalf("non-outbox table should produce 0 rows, got %d", len(rs))
	}

	m3, _ := ParseFlat([]byte(sampleDdl))
	if rs := ExtractOutboxRows(m3); len(rs) != 0 {
		t.Fatalf("DDL should produce 0 rows, got %d", len(rs))
	}

	m4, _ := ParseFlat([]byte(sampleDelete))
	if rs := ExtractOutboxRows(m4); len(rs) != 0 {
		t.Fatalf("DELETE should produce 0 rows, got %d", len(rs))
	}
}

func TestExtractOutboxRows_NilSafe(t *testing.T) {
	if rs := ExtractOutboxRows(nil); len(rs) != 0 {
		t.Fatalf("nil should be safe, got %d rows", len(rs))
	}
}

func TestExtractOutboxRows_SkipsRowsWithBadId(t *testing.T) {
	raw := []byte(`{
      "database":"zhiguang","table":"outbox","type":"INSERT","isDdl":false,
      "data":[
        {"id":"abc","aggregate_type":"x","aggregate_id":"1","type":"y","payload":"{}"},
        {"id":"100","aggregate_type":"following","aggregate_id":"7","type":"FollowCreated","payload":"{}"}
      ],"ts":1,"es":1
    }`)
	m, _ := ParseFlat(raw)
	rows := ExtractOutboxRows(m)
	if len(rows) != 1 {
		t.Fatalf("bad id should be skipped, leaving 1 valid row, got %d", len(rows))
	}
	if rows[0].Id != 100 {
		t.Fatalf("kept wrong row: %+v", rows[0])
	}
}

func TestParseFlat_PreservesAllTopLevelFields(t *testing.T) {
	m, _ := ParseFlat([]byte(sampleInsertOutbox))
	if m.Ts == 0 || m.Es == 0 {
		t.Fatalf("ts/es lost: %+v", m)
	}
	if len(m.PkNames) != 1 || m.PkNames[0] != "id" {
		t.Fatalf("pkNames drift: %+v", m.PkNames)
	}
	if _, err := json.Marshal(m); err != nil {
		t.Fatal(err)
	}
}
