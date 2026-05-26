package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestDecodeLlmMergedConfig(t *testing.T) {
	p := filepath.Join("..", "..", "etc", "llm.yaml")
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
	if len(c.RagIndexer.Kafka.Brokers) == 0 {
		t.Fatalf("ragindexer kafka brokers should not be empty")
	}
	if c.RagIndexer.RagIndex == "" {
		t.Fatalf("ragindexer rag index should not be empty")
	}
}
