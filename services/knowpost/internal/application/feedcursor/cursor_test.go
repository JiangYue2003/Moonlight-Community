package feedcursor

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestCursorRoundTripUsesCanonicalEncoding(t *testing.T) {
	want := Cursor{SortTime: 1_723_456_789, PostID: 9_001}

	encoded, err := Encode(want)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	if strings.Contains(encoded, "=") {
		t.Fatalf("cursor must use raw base64url without padding: %q", encoded)
	}

	got, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got != want {
		t.Fatalf("Decode(Encode(cursor)) = %+v, want %+v", got, want)
	}

	reencoded, err := Encode(got)
	if err != nil {
		t.Fatalf("re-Encode: %v", err)
	}
	if reencoded != encoded {
		t.Fatalf("cursor encoding is not deterministic: first=%q second=%q", encoded, reencoded)
	}
}

func TestCursorRejectsInvalidOrNonCanonicalInput(t *testing.T) {
	tests := map[string]string{
		"empty":            "",
		"invalid base64":   "%%%",
		"too long":         strings.Repeat("a", MaxEncodedLength+1),
		"truncated json":   rawCursor(`{"v":1,"t":2`),
		"trailing value":   rawCursor(`{"v":1,"t":2,"p":3} {}`),
		"unknown field":    rawCursor(`{"v":1,"t":2,"p":3,"x":4}`),
		"unknown version":  rawCursor(`{"v":2,"t":2,"p":3}`),
		"zero sort time":   rawCursor(`{"v":1,"t":0,"p":3}`),
		"negative time":    rawCursor(`{"v":1,"t":-2,"p":3}`),
		"zero post id":     rawCursor(`{"v":1,"t":2,"p":0}`),
		"negative post id": rawCursor(`{"v":1,"t":2,"p":-3}`),
		"non canonical":    rawCursor(`{ "v": 1, "t": 2, "p": 3 }`),
	}

	for name, encoded := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Decode(encoded); err == nil {
				t.Fatalf("Decode(%q) succeeded, want error", encoded)
			}
		})
	}
}

func TestCursorEncodeRejectsInvalidPosition(t *testing.T) {
	for name, cursor := range map[string]Cursor{
		"zero sort time": {SortTime: 0, PostID: 1},
		"negative time":  {SortTime: -1, PostID: 1},
		"zero post id":   {SortTime: 1, PostID: 0},
		"negative id":    {SortTime: 1, PostID: -1},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Encode(cursor); err == nil {
				t.Fatalf("Encode(%+v) succeeded, want error", cursor)
			}
		})
	}
}

func rawCursor(payload string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(payload))
}
