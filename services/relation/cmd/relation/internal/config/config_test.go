package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestDecodeRelationMergedConfig(t *testing.T) {
	p := filepath.Join("..", "..", "etc", "relation.yaml")
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
	if !c.DisableAPI {
		t.Fatalf("merged config should disable the legacy API")
	}
	if c.Rpc.ListenOn == "" {
		t.Fatalf("rpc listen address should not be empty")
	}
	if c.Syncer.Kafka.Topic == "" {
		t.Fatalf("syncer kafka topic should not be empty")
	}
	if !c.DebugHTTP.Enabled {
		t.Fatalf("debug HTTP should be enabled in local development config")
	}
	if c.DebugHTTP.ListenOn != "127.0.0.1:6066" {
		t.Fatalf("debug HTTP listen address = %q, want 127.0.0.1:6066", c.DebugHTTP.ListenOn)
	}
}

func TestDecodeRelationDockerConfig(t *testing.T) {
	p := filepath.Join("..", "..", "etc", "relation-docker.yaml")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read config: %v", err)
	}

	var c Config
	if err := conf.LoadFromYamlBytes(b, &c); err != nil {
		t.Fatalf("load config: %v", err)
	}

	if !c.DebugHTTP.Enabled {
		t.Fatal("debug HTTP should be enabled in Docker development config")
	}
	if c.DebugHTTP.ListenOn != "0.0.0.0:6066" {
		t.Fatalf("debug HTTP listen address = %q, want 0.0.0.0:6066", c.DebugHTTP.ListenOn)
	}
}
