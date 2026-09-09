package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestDecodeKnowpostMergedConfig(t *testing.T) {
	p := filepath.Join("..", "..", "etc", "knowpost.yaml")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read config: %v", err)
	}

	var c Config
	if err := conf.LoadFromYamlBytes(b, &c); err != nil {
		t.Fatalf("load config: %v", err)
	}

	if c.Rpc.ListenOn == "" {
		t.Fatalf("rpc listen address should not be empty")
	}
	if len(c.Rpc.Kafka.Brokers) == 0 {
		t.Fatalf("rpc kafka brokers should not be empty")
	}
	if !c.DebugHTTP.Enabled {
		t.Fatalf("debug HTTP should be enabled in local development config")
	}
	if c.DebugHTTP.ListenOn != "127.0.0.1:6064" {
		t.Fatalf("debug HTTP listen address = %q, want 127.0.0.1:6064", c.DebugHTTP.ListenOn)
	}
	if c.Rpc.Feed.Epoch.KeyPrefix != "feed" || c.Rpc.Feed.Epoch.RelationL1TTL != time.Second || c.Rpc.Feed.Epoch.SafetyL1TTL != time.Second {
		t.Fatalf("unexpected feed epoch config: %+v", c.Rpc.Feed.Epoch)
	}
	if c.Rpc.L1.FeedEpochNumCounters != 200_000 || c.Rpc.L1.FeedEpochMaxCostMB != 4 {
		t.Fatalf("unexpected independent feed epoch L1 budget: %+v", c.Rpc.L1)
	}
	if c.Rpc.Feed.RouteSnapshot.Enabled || c.Rpc.Feed.RouteSnapshot.TTL != 5*time.Second {
		t.Fatalf("unexpected opt-in route snapshot config: %+v", c.Rpc.Feed.RouteSnapshot)
	}
	if c.Rpc.Feed.CombinedPipeline.Enabled || c.Rpc.Feed.CombinedPipeline.BatchSize != 128 {
		t.Fatalf("unexpected opt-in combined pipeline config: %+v", c.Rpc.Feed.CombinedPipeline)
	}
	if c.Rpc.Feed.CursorPagination.Enabled {
		t.Fatalf("cursor pagination must remain disabled by default: %+v", c.Rpc.Feed.CursorPagination)
	}
	if c.Rpc.Feed.PageCache.Mode != "off" || c.Rpc.Feed.PageCache.KeyPrefix != "feed" ||
		c.Rpc.Feed.PageCache.Page != 1 || c.Rpc.Feed.PageCache.Size != 20 ||
		c.Rpc.Feed.PageCache.L1FreshTTL != 800*time.Millisecond || c.Rpc.Feed.PageCache.L2FreshTTL != 4*time.Second ||
		c.Rpc.Feed.PageCache.StaleTTL != 10*time.Second || c.Rpc.Feed.PageCache.JitterPercent != 20 ||
		c.Rpc.Feed.PageCache.RefreshWorkers != 32 || c.Rpc.Feed.PageCache.RefreshQueue != 1024 ||
		c.Rpc.Feed.PageCache.LoaderTimeout != 2*time.Second {
		t.Fatalf("unexpected opt-in page cache config: %+v", c.Rpc.Feed.PageCache)
	}
	if c.Rpc.L1.FeedRouteNumCounters != 100_000 || c.Rpc.L1.FeedRouteMaxCostMB != 32 {
		t.Fatalf("unexpected independent feed route L1 budget: %+v", c.Rpc.L1)
	}
	if c.Rpc.L1.FeedPageNumCounters != 100_000 || c.Rpc.L1.FeedPageMaxCostMB != 128 {
		t.Fatalf("unexpected independent feed page L1 budget: %+v", c.Rpc.L1)
	}
	if c.Rpc.Feed.Epoch.RelationConsumerEnabled {
		t.Fatal("relation epoch consumer must remain disabled by default")
	}
	if c.Rpc.Feed.Epoch.SafetyConsumerEnabled {
		t.Fatal("content safety epoch consumer must remain disabled by default")
	}
	if c.Rpc.Kafka.CanalOutboxTopic != "canal-outbox" ||
		c.Rpc.Kafka.RelationEpochGroupId != "knowpost-feed-relation-epoch" ||
		c.Rpc.Kafka.SafetyEpochGroupId != "knowpost-feed-content-safety-epoch" {
		t.Fatalf("unexpected relation epoch consumer routing: %+v", c.Rpc.Kafka)
	}
}
