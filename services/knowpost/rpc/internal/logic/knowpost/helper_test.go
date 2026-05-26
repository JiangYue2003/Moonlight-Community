package knowpostlogic

import (
	"database/sql"
	"testing"
	"time"

	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

func TestParseStringList_Empty(t *testing.T) {
	if got := parseStringList(""); got != nil {
		t.Fatalf("empty should be nil, got %v", got)
	}
	if got := parseStringList("not json"); got != nil {
		t.Fatalf("invalid json should be nil, got %v", got)
	}
}

func TestParseStringList_RoundTrip(t *testing.T) {
	in := []string{"java", "编程", "ai"}
	s, err := encodeStringList(in)
	if err != nil {
		t.Fatal(err)
	}
	got := parseStringList(s)
	if len(got) != 3 || got[1] != "编程" {
		t.Fatalf("round trip lost data: %+v", got)
	}
}

func TestEncodeStringList_NilStaysEmpty(t *testing.T) {
	s, err := encodeStringList(nil)
	if err != nil || s != "" {
		t.Fatalf("nil → empty string; got %q err=%v", s, err)
	}
}

func TestRowToDetail_PopulatesAllFields(t *testing.T) {
	now := time.Now()
	row := &model.KnowPosts{
		Id: 42, CreatorId: 7,
		Title:            sql.NullString{String: "T", Valid: true},
		Description:      sql.NullString{String: "D", Valid: true},
		ContentObjectKey: sql.NullString{String: "k", Valid: true},
		ContentEtag:      sql.NullString{String: "et", Valid: true},
		ContentSize:      sql.NullInt64{Int64: 100, Valid: true},
		ContentSha256:    sql.NullString{String: "sha", Valid: true},
		Tags:             sql.NullString{String: `["a","b"]`, Valid: true},
		ImgUrls:          sql.NullString{String: `["u1","u2"]`, Valid: true},
		Type:             "image_text",
		Visible:          "public",
		Status:           "published",
		IsTop:            1,
		CreateTime:       now, UpdateTime: now,
		PublishTime: sql.NullTime{Time: now, Valid: true},
	}
	d := rowToDetail(row)
	if d.Id != "42" || d.CreatorId != 7 || d.Title != "T" || d.IsTop != true ||
		len(d.Tags) != 2 || d.Tags[0] != "a" || len(d.ImgUrls) != 2 ||
		d.Status != "published" || d.PublishTime == 0 {
		t.Fatalf("detail mapping drift: %+v", d)
	}
}

func TestRowToFeedItem_StripsContentMeta(t *testing.T) {
	row := &model.KnowPosts{
		Id: 1, CreatorId: 2,
		Title:         sql.NullString{String: "T", Valid: true},
		ContentEtag:   sql.NullString{String: "should-not-leak", Valid: true},
		ContentSha256: sql.NullString{String: "ditto", Valid: true},
		Visible:       "public",
		Status:        "published",
	}
	it := rowToFeedItem(row)
	if it.Title != "T" || it.Visible != "public" {
		t.Fatalf("feed item mapping wrong: %+v", it)
	}
	// 不应有 etag/sha256/size 字段（pb 定义里就没有，断言成立即 OK）。
}

func TestSetNS_EmptyAndNonEmpty(t *testing.T) {
	if got := setNS(""); got.Valid {
		t.Fatal("empty should be invalid NullString")
	}
	if got := setNS("x"); !got.Valid || got.String != "x" {
		t.Fatalf("non-empty drift: %+v", got)
	}
}

func TestNormalizePage_Bounds(t *testing.T) {
	cases := []struct {
		page, size   int32
		wantP, wantS int
	}{
		{0, 0, 1, 20},
		{-1, 99, 1, 50},
		{2, 30, 2, 30},
	}
	for _, c := range cases {
		p, s := normalizePage(c.page, c.size)
		if p != c.wantP || s != c.wantS {
			t.Errorf("normalize(%d,%d) = (%d,%d), want (%d,%d)",
				c.page, c.size, p, s, c.wantP, c.wantS)
		}
	}
}
