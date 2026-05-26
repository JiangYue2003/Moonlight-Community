package hotkey

import (
	"testing"
	"time"
)

func TestExtension_AllLevels(t *testing.T) {
	cases := []struct {
		level Level
		want  time.Duration
	}{
		{LevelNone, 0},
		{LevelLow, 20 * time.Second},
		{LevelMedium, 60 * time.Second},
		{LevelHigh, 120 * time.Second},
	}
	for _, c := range cases {
		if got := Extension(c.level); got != c.want {
			t.Errorf("Extension(%v) = %v, want %v", c.level, got, c.want)
		}
	}
}

func TestTTLForPublic_AddsExtension(t *testing.T) {
	base := 60 * time.Second
	cases := []struct {
		level Level
		want  time.Duration
	}{
		{LevelNone, 60 * time.Second},
		{LevelLow, 80 * time.Second},
		{LevelMedium, 120 * time.Second},
		{LevelHigh, 180 * time.Second},
	}
	for _, c := range cases {
		if got := TTLForPublic(base, c.level); got != c.want {
			t.Errorf("TTLForPublic(60s, %v) = %v, want %v", c.level, got, c.want)
		}
	}
}

func TestTTLForMine_DifferentBaseSameExtension(t *testing.T) {
	base := 30 * time.Second
	if got := TTLForMine(base, LevelMedium); got != 90*time.Second {
		t.Fatalf("TTLForMine(30s, MEDIUM) = %v, want 90s", got)
	}
	if got := TTLForMine(base, LevelNone); got != 30*time.Second {
		t.Fatalf("TTLForMine(30s, NONE) should equal base, got %v", got)
	}
}
