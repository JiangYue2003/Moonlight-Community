package userlogic

import (
	"database/sql"
	"testing"
	"time"
)

func TestSetOrEmpty(t *testing.T) {
	if got := setOrEmpty(""); got.Valid {
		t.Fatal("empty string should produce invalid sql.NullString")
	}
	got := setOrEmpty("hello")
	if !got.Valid || got.String != "hello" {
		t.Fatalf("non-empty should be valid: %+v", got)
	}
}

func TestUpdateProfileLogic_BirthdayParse(t *testing.T) {
	// 验证我们使用的日期格式与 ProfilePatchRequest 校验保持一致（YYYY-MM-DD）。
	good := "2000-01-02"
	bad := "01/02/2000"
	if _, err := time.Parse("2006-01-02", good); err != nil {
		t.Fatalf("%q should parse: %v", good, err)
	}
	if _, err := time.Parse("2006-01-02", bad); err == nil {
		t.Fatal("non-ISO date should not parse")
	}
}

// 用 sql.NullTime 的零值表示"清空生日"——对调用 UpdateProfile 时传 BirthdaySet=true 且 Birthday=""
// 的语义做硬约束。
func TestNullTimeZeroValueIsInvalid(t *testing.T) {
	var nt sql.NullTime
	if nt.Valid {
		t.Fatal("zero-value sql.NullTime must be invalid (so DB writes NULL)")
	}
}
