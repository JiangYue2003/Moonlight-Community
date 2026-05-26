package textx

import (
	"strings"
	"testing"
)

func TestSplitByHeader_Basic(t *testing.T) {
	md := "# A\nhello\n## B\nworld\n### C\n!\n"
	got := SplitByHeader(md)
	if len(got) != 3 {
		t.Fatalf("len=%d", len(got))
	}
	if got[0].Title != "A" || got[0].Body != "hello" {
		t.Errorf("section0: %+v", got[0])
	}
	if got[1].Title != "B" || got[1].Body != "world" {
		t.Errorf("section1: %+v", got[1])
	}
	if got[2].Title != "C" || got[2].Body != "!" {
		t.Errorf("section2: %+v", got[2])
	}
}

func TestSplitByHeader_NoHeader(t *testing.T) {
	got := SplitByHeader("just plain text\nline 2")
	if len(got) != 1 || got[0].Title != "" || !strings.Contains(got[0].Body, "plain") {
		t.Fatalf("got: %+v", got)
	}
}

func TestSplitByHeader_HashInBody(t *testing.T) {
	md := "# A\nthis #is not a header\n"
	got := SplitByHeader(md)
	if len(got) != 1 {
		t.Fatalf("len=%d %+v", len(got), got)
	}
}

func TestChunk_RuneSafe(t *testing.T) {
	text := strings.Repeat("中", 1000)
	got := Chunk(text, 800, 100)
	if len(got) != 2 {
		t.Fatalf("len=%d", len(got))
	}
	for _, c := range got {
		if !strings.HasPrefix(c, "中") {
			t.Errorf("not utf8 boundary: %q", c[:6])
		}
	}
	// 第二块起点 = 800-100 = 700，长度应是 1000-700 = 300
	r := []rune(got[1])
	if len(r) != 300 {
		t.Errorf("second block len=%d want 300", len(r))
	}
}

func TestChunk_Smaller(t *testing.T) {
	if got := Chunk("短", 800, 100); len(got) != 1 || got[0] != "短" {
		t.Fatalf("got %v", got)
	}
}

func TestNormalizeNFKC(t *testing.T) {
	// 全角数字 ABC → 半角
	got := NormalizeNFKC("ＡＢＣ１")
	if got != "ABC1" {
		t.Errorf("got %q", got)
	}
}

func TestCollapseWhitespace(t *testing.T) {
	got := CollapseWhitespace("  hello　​\tworld  \n!  ")
	if got != "hello world !" {
		t.Errorf("got %q", got)
	}
}

func TestStripWrappingQuotes(t *testing.T) {
	cases := map[string]string{
		`"hello"`: "hello",
		`“你好”`:    "你好",
		`「文章」`:    "文章",
		`「『嵌套』」`:  "嵌套",
		"plain":   "plain",
		`'a'`:     "a",
	}
	for in, want := range cases {
		if got := StripWrappingQuotes(in); got != want {
			t.Errorf("in=%q got=%q want=%q", in, got, want)
		}
	}
}

func TestStripTrailingPunct(t *testing.T) {
	if got := StripTrailingPunct("hello。 "); got != "hello" {
		t.Errorf("got %q", got)
	}
	if got := StripTrailingPunct("done!!?。"); got != "done" {
		t.Errorf("got %q", got)
	}
}

func TestTruncateRunes(t *testing.T) {
	if got := TruncateRunes("中国人民万岁", 3); got != "中国人" {
		t.Errorf("got %q", got)
	}
	if got := TruncateRunes("abc", 0); got != "abc" {
		t.Errorf("got %q", got)
	}
}

func TestDescriptionPostProcess(t *testing.T) {
	in := "  “这是一个 关于 知文 的 测试   段落，  说明 文档。” \n第二段不要了"
	got := DescriptionPostProcess{MaxRunes: 50}.Apply(in)
	want := "这是一个 关于 知文 的 测试 段落, 说明 文档"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
