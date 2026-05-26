// Package ctxdata 在 context 中传递跨服务的运行时数据（如 userId）。
//
// 约定 metadata key 为 "x-user-id"，HTTP middleware 与 gRPC interceptor
// 都基于该 key 完成 access token → context.Context 的写入。
package ctxdata

import (
	"context"
	"strconv"
)

const (
	// MetadataKeyUserId gRPC metadata 中携带 userId 的 key（小写）。
	MetadataKeyUserId = "x-user-id"
)

type ctxKey struct{ name string }

var userIdKey = ctxKey{"userId"}

// WithUserId 将 userId 写入 context。
func WithUserId(ctx context.Context, uid int64) context.Context {
	return context.WithValue(ctx, userIdKey, uid)
}

// GetUserId 从 context 中读取 userId；未设置时返回 0, false。
func GetUserId(ctx context.Context) (int64, bool) {
	v := ctx.Value(userIdKey)
	if v == nil {
		return 0, false
	}
	uid, ok := v.(int64)
	return uid, ok
}

// MustGetUserId 读取 userId；未设置 panic。仅用于已经强制鉴权的路径。
func MustGetUserId(ctx context.Context) int64 {
	uid, ok := GetUserId(ctx)
	if !ok {
		panic("ctxdata: userId not in context")
	}
	return uid
}

// ParseUserId 从字符串解析 userId（用于 metadata 中的字符串值）。
func ParseUserId(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }

// FormatUserId 将 userId 序列化为字符串（用于 metadata 注入）。
func FormatUserId(uid int64) string { return strconv.FormatInt(uid, 10) }
