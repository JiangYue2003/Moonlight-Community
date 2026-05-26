package schema

import "testing"

func TestIdxOf_KnownAndUnknownMetrics(t *testing.T) {
	cases := map[string]int{
		"like":    IdxLike,
		"fav":     IdxFav,
		"":        -1,
		"unknown": -1,
		"comment": -1, // 已在 schema 中预留索引但未在 SUPPORTED 中暴露
	}
	for metric, want := range cases {
		if got := IdxOf(metric); got != want {
			t.Errorf("IdxOf(%q) = %d, want %d", metric, got, want)
		}
	}
}

func TestSupported_OnlyExposesActiveMetrics(t *testing.T) {
	if len(Supported) != 2 {
		t.Fatalf("Supported should be {like, fav}, got %v", Supported)
	}
	for _, m := range Supported {
		if IdxOf(m) < 0 {
			t.Fatalf("Supported metric %q must map to a valid idx", m)
		}
	}
}

func TestSchemaConstants_StayCompatibleWithJava(t *testing.T) {
	// 这些常量是与 Java 端 SDS 二进制布局对齐的契约，任何修改都会破坏跨语言兼容性。
	if SchemaLen != 5 || FieldSize != 4 {
		t.Fatalf("SDS layout drifted from Java contract: len=%d size=%d", SchemaLen, FieldSize)
	}
	if SchemaId != "v1" {
		t.Fatalf("SchemaId drifted: %q", SchemaId)
	}
}

func TestSdsKey_Format(t *testing.T) {
	got := SdsKey("knowpost", "123")
	want := "cnt:v1:knowpost:123"
	if got != want {
		t.Fatalf("SdsKey = %q, want %q", got, want)
	}
}

func TestBitmapKey_Format(t *testing.T) {
	got := BitmapKey("like", "knowpost", "123", 7)
	want := "bm:like:knowpost:123:7"
	if got != want {
		t.Fatalf("BitmapKey = %q, want %q", got, want)
	}
}

func TestAggKey_Format(t *testing.T) {
	got := AggKey("knowpost", "abc")
	want := "agg:v1:knowpost:abc"
	if got != want {
		t.Fatalf("AggKey = %q, want %q", got, want)
	}
}

func TestUserSdsKey_Format(t *testing.T) {
	if got := UserSdsKey(42); got != "ucnt:42" {
		t.Fatalf("UserSdsKey = %q", got)
	}
}

func TestChunkOfAndBitOf_Boundaries(t *testing.T) {
	cases := []struct {
		uid           int64
		wantChunk     int64
		wantBitOffset int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{ChunkSize - 1, 0, ChunkSize - 1}, // 32767 → chunk0 末位
		{ChunkSize, 1, 0},                 // 32768 → chunk1 首位
		{ChunkSize + 5, 1, 5},
		{ChunkSize*3 + 17, 3, 17},
		{ChunkSize * 1234, 1234, 0},
	}
	for _, c := range cases {
		if got := ChunkOf(c.uid); got != c.wantChunk {
			t.Errorf("ChunkOf(%d) = %d, want %d", c.uid, got, c.wantChunk)
		}
		if got := BitOf(c.uid); got != c.wantBitOffset {
			t.Errorf("BitOf(%d) = %d, want %d", c.uid, got, c.wantBitOffset)
		}
	}
}

func TestChunkSize_MatchesJavaContract(t *testing.T) {
	// 32768 位 = 4KB / shard，Java 版同值；改动需配合迁移历史位图数据。
	if ChunkSize != 32768 {
		t.Fatalf("ChunkSize drifted from Java contract: %d", ChunkSize)
	}
}
