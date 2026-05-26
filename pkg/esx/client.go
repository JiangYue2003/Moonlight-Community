// Package esx 是 Elasticsearch 9.x REST 客户端的薄封装。
//
// 设计取舍：不引官方 elastic SDK，直接 net/http + json，原因：
//   - 我们只用到 HEAD/PUT/POST/GET/DELETE 几个固定路径
//   - 测试用 httptest 起 mock 最直接
//   - 减少 go.mod 体积与依赖面
package esx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"
)

// Config 客户端构造参数。
type Config struct {
	Addrs    []string      // 至少 1 个：http://host:9200；多个则 round-robin
	Username string        // 可选 basic auth
	Password string        // 可选 basic auth
	Timeout  time.Duration // 单请求超时；0 → 默认 10s
}

// Client 暴露给业务侧的 ES REST 客户端。
type Client struct {
	addrs    []string
	username string
	password string
	hc       *http.Client
	rr       uint32 // 简易 round-robin 计数器（atomic）
}

// New 构造 Client；不主动探活，留给上层在启动时调 Ping。
func New(cfg Config) (*Client, error) {
	if len(cfg.Addrs) == 0 {
		return nil, errors.New("esx: at least one addr required")
	}
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	c := &Client{
		addrs:    append([]string{}, cfg.Addrs...),
		username: cfg.Username,
		password: cfg.Password,
		hc:       &http.Client{Timeout: timeout},
	}
	return c, nil
}

// Ping 探活，调用 GET /。
func (c *Client) Ping(ctx context.Context) error {
	resp, err := c.do(ctx, http.MethodGet, "/", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("esx ping: status=%d", resp.StatusCode)
	}
	return nil
}

// pickAddr 简易 round-robin。
func (c *Client) pickAddr() string {
	n := atomic.AddUint32(&c.rr, 1)
	return c.addrs[int(n)%len(c.addrs)]
}

// do 发请求并返回原始 response。调用方负责 Close。非 2xx 不返回 err，让上层结合 path 决定。
func (c *Client) do(ctx context.Context, method, path string, body any) (*http.Response, error) {
	addr := strings.TrimRight(c.pickAddr(), "/")
	u := addr + path

	var rd io.Reader
	if body != nil {
		switch v := body.(type) {
		case []byte:
			rd = bytes.NewReader(v)
		case json.RawMessage:
			rd = bytes.NewReader(v)
		case string:
			rd = strings.NewReader(v)
		default:
			buf, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("esx marshal: %w", err)
			}
			rd = bytes.NewReader(buf)
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, u, rd)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.username != "" {
		req.SetBasicAuth(c.username, c.password)
	}
	return c.hc.Do(req)
}

// readJSON 读 body 反序列化；非 2xx 返回 *Error。
func (c *Client) readJSON(resp *http.Response, out any) error {
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return &Error{Status: resp.StatusCode, Body: string(buf)}
	}
	if out == nil || len(buf) == 0 {
		return nil
	}
	return json.Unmarshal(buf, out)
}

// Error 非 2xx 上游错误。
type Error struct {
	Status int
	Body   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("esx http %d: %s", e.Status, truncate(e.Body, 256))
}

// IsNotFound 便于 caller 判断 404。
func IsNotFound(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Status == http.StatusNotFound
	}
	return false
}

// escapePath 为路径段做 URL 转义；ES index/id 允许字符多但安全起见走 PathEscape。
func escapePath(s string) string { return url.PathEscape(s) }

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "...(truncated)"
}
