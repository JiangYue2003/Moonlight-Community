package bootstrap

import (
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestConfigFilesPreserveRuntimeContract(t *testing.T) {
	for _, name := range []string{"user.yaml", "user-docker.yaml"} {
		t.Run(name, func(t *testing.T) {
			var cfg Config
			path := filepath.Join("..", "..", "cmd", "user", "etc", name)
			if err := conf.Load(path, &cfg); err != nil {
				t.Fatalf("load %s: %v", path, err)
			}

			if cfg.Name != "user.rpc" {
				t.Fatalf("Name = %q, want user.rpc", cfg.Name)
			}
			if cfg.ListenOn != "0.0.0.0:9002" {
				t.Fatalf("ListenOn = %q, want 0.0.0.0:9002", cfg.ListenOn)
			}
			if cfg.Etcd.Key != "user.rpc" {
				t.Fatalf("Etcd.Key = %q, want user.rpc", cfg.Etcd.Key)
			}
			if cfg.Prometheus.Port != 9102 {
				t.Fatalf("Prometheus.Port = %d, want 9102", cfg.Prometheus.Port)
			}
			if cfg.Redis.Key != "user-rpc" {
				t.Fatalf("Redis.Key = %q, want user-rpc", cfg.Redis.Key)
			}
		})
	}
}
