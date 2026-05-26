package cache

import (
	"strings"
	"testing"
	"time"
)

func TestKeysHaveExpectedShape(t *testing.T) {
	if got := DetailKey(123); got != "knowpost:detail:123:v1" {
		t.Fatalf("DetailKey: %q", got)
	}
	if got := FeedPublicL1Key(20, 1); got != "feed:public:20:1:v1" {
		t.Fatalf("FeedPublicL1Key: %q", got)
	}
	if got := FeedPublicIdsKey(20, 2, 7); got != "feed:public:ids:20:7:2" {
		t.Fatalf("FeedPublicIdsKey: %q", got)
	}
	if got := FeedPublicHasMoreKey("k"); got != "k:hasMore" {
		t.Fatalf("FeedPublicHasMoreKey: %q", got)
	}
	if got := FeedItemKey(99); got != "feed:item:99" {
		t.Fatalf("FeedItemKey: %q", got)
	}
	if got := FeedReverseIndexKey(99, 12345); got != "feed:public:index:99:12345" {
		t.Fatalf("FeedReverseIndexKey: %q", got)
	}
	if got := FeedMineKey(7, 20, 1); got != "feed:mine:7:20:1" {
		t.Fatalf("FeedMineKey: %q", got)
	}
}

func TestHourSlot_AdvancesEveryHour(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2024-06-01T00:30:00Z")
	t2, _ := time.Parse(time.RFC3339, "2024-06-01T01:00:00Z")
	if HourSlot(t1) == HourSlot(t2) {
		t.Fatal("HourSlot must change at hour boundary")
	}
	t1b, _ := time.Parse(time.RFC3339, "2024-06-01T00:00:00Z")
	t1c, _ := time.Parse(time.RFC3339, "2024-06-01T00:59:59Z")
	if HourSlot(t1b) != HourSlot(t1c) {
		t.Fatal("HourSlot must be stable within the same hour")
	}
}

func TestDetailLayoutVerInKey(t *testing.T) {
	got := DetailKey(1)
	if !strings.HasSuffix(got, ":v1") {
		t.Fatalf("layout version expected to be embedded as :v1, got %q", got)
	}
}
