// Package redisx 封装 go-redis 客户端构造与 Lua 脚本预编译。
package redisx

import (
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config Redis 连接参数（与 go-zero RedisCache 字段尽量对齐）。
type Config struct {
	Host     string `json:",default=127.0.0.1:6379"`
	Pass     string `json:",optional"`
	Type     string `json:",default=node,options=[node,cluster]"`
	DB       int    `json:",optional"`
	Tls      bool   `json:",optional"`
	PoolSize int    `json:",optional"`
}

// New 根据 Config 构造一个 Cmdable 客户端（统一 node / cluster 两种形态）。
// 上层对 cluster 支持 PFCOUNT / SCAN 类跨槽命令时需注意分片规约。
func New(c Config) redis.UniversalClient {
	opts := &redis.UniversalOptions{
		Addrs:        []string{c.Host},
		Password:     c.Pass,
		DB:           c.DB,
		PoolSize:     c.PoolSize,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	if c.Tls {
		opts.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	if c.Type == "cluster" {
		opts.Addrs = splitAddrs(c.Host)
	}
	return redis.NewUniversalClient(opts)
}

// splitAddrs 将 "h1:6379,h2:6379" 这类多节点字符串切片化。
func splitAddrs(s string) []string {
	out := []string{}
	cur := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == ',' || ch == ';' || ch == ' ' {
			if cur != "" {
				out = append(out, cur)
				cur = ""
			}
			continue
		}
		cur += string(ch)
	}
	if cur != "" {
		out = append(out, cur)
	}
	return out
}
