package ossx

import (
	"strings"
	"testing"
)

func TestExtFromContentType_ExplicitExtPreferred(t *testing.T) {
	got := ExtFromContentType("text/markdown", SceneKnowPostContent, "ZIP")
	if got != ".zip" {
		t.Fatalf("explicit ext should win: got %q", got)
	}
	got = ExtFromContentType("text/markdown", SceneKnowPostContent, ".PNG")
	if got != ".png" {
		t.Fatalf("dotted explicit ext: got %q", got)
	}
}

func TestExtFromContentType_ContentScene(t *testing.T) {
	cases := map[string]string{
		"text/markdown":    ".md",
		"text/html":        ".html",
		"text/plain":       ".txt",
		"application/json": ".json",
		"application/xyz":  ".bin",
		"":                 ".bin",
	}
	for ct, want := range cases {
		if got := ExtFromContentType(ct, SceneKnowPostContent, ""); got != want {
			t.Errorf("content %q: got %q, want %q", ct, got, want)
		}
	}
}

func TestExtFromContentType_ImageScene(t *testing.T) {
	cases := map[string]string{
		"image/jpeg": ".jpg",
		"image/jpg":  ".jpg",
		"image/png":  ".png",
		"image/webp": ".webp",
		"image/heic": ".img",
		"":           ".img",
		"text/plain": ".img",
	}
	for ct, want := range cases {
		if got := ExtFromContentType(ct, SceneKnowPostImage, ""); got != want {
			t.Errorf("image %q: got %q, want %q", ct, got, want)
		}
	}
}

func TestObjectKeyFor_KnowPostContent(t *testing.T) {
	got := ObjectKeyFor(SceneKnowPostContent, "12345", 0, "", ".md")
	if got != "posts/12345/content.md" {
		t.Fatalf("content key drift: %q", got)
	}
}

func TestObjectKeyFor_KnowPostImageHasDateAndRand(t *testing.T) {
	got := ObjectKeyFor(SceneKnowPostImage, "12345", 0, "", ".png")
	// posts/12345/images/{8 digits}/{8 hex}.png
	if !strings.HasPrefix(got, "posts/12345/images/") {
		t.Fatalf("prefix wrong: %q", got)
	}
	if !strings.HasSuffix(got, ".png") {
		t.Fatalf("suffix wrong: %q", got)
	}
	parts := strings.Split(got, "/")
	if len(parts) != 5 {
		t.Fatalf("expected 5 segments, got %d (%q)", len(parts), got)
	}
	if len(parts[3]) != 8 {
		t.Fatalf("date segment should be 8 chars: %q", parts[3])
	}
	tail := strings.TrimSuffix(parts[4], ".png")
	if len(tail) != 8 {
		t.Fatalf("rand segment should be 8 hex chars: %q", parts[4])
	}
}

func TestObjectKeyFor_AvatarUsesFolderUserIdAndEpoch(t *testing.T) {
	got := ObjectKeyFor(SceneAvatar, "", 42, "avatars", ".jpg")
	if !strings.HasPrefix(got, "avatars/42-") {
		t.Fatalf("avatar prefix wrong: %q", got)
	}
	if !strings.HasSuffix(got, ".jpg") {
		t.Fatalf("avatar suffix wrong: %q", got)
	}
}

func TestObjectKeyFor_AvatarDefaultFolder(t *testing.T) {
	got := ObjectKeyFor(SceneAvatar, "", 1, "", ".png")
	if !strings.HasPrefix(got, "avatars/") {
		t.Fatalf("default folder should be 'avatars': %q", got)
	}
}

func TestBuildContentUrl_PublicDomainPrioritized(t *testing.T) {
	c := &Client{cfg: Config{
		Endpoint: "oss-cn-shanghai.aliyuncs.com", Bucket: "b",
		PublicDomain: "https://cdn.example.com",
	}}
	if got := c.BuildContentUrl("posts/1/content.md"); got != "https://cdn.example.com/posts/1/content.md" {
		t.Fatalf("PublicDomain should win: %q", got)
	}
}

func TestBuildContentUrl_FallsBackToEndpoint(t *testing.T) {
	c := &Client{cfg: Config{
		Endpoint: "oss-cn-shanghai.aliyuncs.com", Bucket: "zhiguang",
	}}
	want := "https://zhiguang.oss-cn-shanghai.aliyuncs.com/posts/1/content.md"
	if got := c.BuildContentUrl("posts/1/content.md"); got != want {
		t.Fatalf("endpoint url: got %q want %q", got, want)
	}
}

func TestNew_RejectsMissingConfig(t *testing.T) {
	if _, err := New(Config{}); err == nil {
		t.Fatal("empty cfg should error")
	}
	if _, err := New(Config{Endpoint: "x", Bucket: "b"}); err == nil {
		t.Fatal("missing creds should error")
	}
}
