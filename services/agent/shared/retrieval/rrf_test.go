package retrieval

import "testing"

func TestFuseRRF(t *testing.T) {
	c1 := []ScoredItem{{DocID: "a"}, {DocID: "b"}, {DocID: "c"}}
	c2 := []ScoredItem{{DocID: "b"}, {DocID: "d"}, {DocID: "a"}}
	out := FuseRRF(60, c1, c2)
	if len(out) != 4 {
		t.Fatalf("len=%d", len(out))
	}
	if out[0].DocID != "b" && out[0].DocID != "a" {
		t.Fatalf("unexpected top doc=%s", out[0].DocID)
	}
}
