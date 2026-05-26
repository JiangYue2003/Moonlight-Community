package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestDecodeCounterMergedConfig(t *testing.T) {
	p := filepath.Join("..", "..", "etc", "counter.yaml")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read config: %v", err)
	}

	var c Config
	if err := conf.LoadFromYamlBytes(b, &c); err != nil {
		t.Fatalf("load config: %v", err)
	}

	if c.Api.Port == 0 {
		t.Fatalf("api port should not be zero")
	}
	if c.DisableAPI {
		t.Fatalf("disableAPI should default to false")
	}
	if c.Rpc.ListenOn == "" {
		t.Fatalf("rpc listen address should not be empty")
	}
	if c.Aggregator.Kafka.Topic == "" {
		t.Fatalf("aggregator kafka topic should not be empty")
	}
}
