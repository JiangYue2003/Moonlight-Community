package query

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBuildSearchBody_NoTagsNoCursor(t *testing.T) {
	body := BuildSearchBody("go-zero", nil, 20, nil)
	s := string(body)
	for _, want := range []string{`"size":20`, `"multi_match"`, `"title^3"`, `"function_score"`, `"score_mode":"sum"`, `"boost_mode":"sum"`} {
		if !strings.Contains(s, want) {
			t.Errorf("missing %q\nbody=%s", want, s)
		}
	}
	if strings.Contains(s, `"search_after"`) {
		t.Errorf("should not have search_after on first page")
	}
	if strings.Contains(s, `"terms":{"tags"`) {
		t.Errorf("should not have tags filter when empty")
	}
	if strings.Contains(s, `"operator":"and"`) {
		t.Errorf("java alignment should not force operator=and: %s", s)
	}
}

func TestBuildSearchBody_WithTagsAndAfter(t *testing.T) {
	body := BuildSearchBody("hi", []string{"后端", "Go"}, 10, []any{json.Number("1.5"), json.Number("100"), "abc"})
	s := string(body)
	if !strings.Contains(s, `"tags":["后端","Go"]`) {
		t.Errorf("tags missing: %s", s)
	}
	if !strings.Contains(s, `"search_after"`) {
		t.Errorf("search_after missing: %s", s)
	}
}

func TestEncodeDecodeCursor(t *testing.T) {
	in := []json.RawMessage{
		json.RawMessage(`1.5`),
		json.RawMessage(`1700000000000`),
		json.RawMessage(`"abc"`),
	}
	enc := EncodeCursor(in)
	if enc == "" {
		t.Fatal("encode empty")
	}
	dec, err := DecodeCursor(enc)
	if err != nil {
		t.Fatal(err)
	}
	if len(dec) != 3 {
		t.Fatalf("len=%d", len(dec))
	}
	// 数字保持为 json.Number
	if _, ok := dec[0].(json.Number); !ok {
		t.Errorf("dec[0] not json.Number: %T", dec[0])
	}
	if dec[2] != "abc" {
		t.Errorf("dec[2] = %v", dec[2])
	}
}

func TestDecodeCursor_EmptyReturnsNil(t *testing.T) {
	dec, err := DecodeCursor("")
	if err != nil || dec != nil {
		t.Fatalf("err=%v dec=%v", err, dec)
	}
}

func TestDecodeCursor_Bad(t *testing.T) {
	if _, err := DecodeCursor("!!!notbase64"); err == nil {
		t.Fatal("expect error")
	}
}

func TestParseTags(t *testing.T) {
	cases := map[string][]string{
		"":         nil,
		"a":        {"a"},
		"a,b, c":   {"a", "b", "c"},
		" , a , ,": {"a"},
	}
	for in, want := range cases {
		got := ParseTags(in)
		if !slicesEqual(got, want) {
			t.Errorf("in=%q got=%v want=%v", in, got, want)
		}
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
