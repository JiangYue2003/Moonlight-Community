// Package sfx 是 golang.org/x/sync/singleflight 的轻量泛型包装。
//
// 用途：在 cachex.GetOrLoad 与 feed/detail loader 入口去重并发回源，
// 避免缓存击穿时多 goroutine 同时打 DB。
package sfx

import (
	"context"

	"golang.org/x/sync/singleflight"
)

// Group 按 key 去重并发调用；同一 key 的并发 DoCtx 只会执行一次 fn，
// 其它 goroutine 等待并复用结果。
type Group[V any] struct {
	g singleflight.Group
}

// New 构造空 Group。
func New[V any]() *Group[V] { return &Group[V]{} }

// DoCtx 执行 fn 并按 key 去重。错误会被传播给所有等待者。
// 注意：fn 内部若用 ctx，是发起者的 ctx；其它等待者的 ctx 取消不会取消正在执行的 fn。
func (g *Group[V]) DoCtx(ctx context.Context, key string, fn func(ctx context.Context) (V, error)) (V, error) {
	var zero V
	if err := ctx.Err(); err != nil {
		return zero, err
	}
	v, err, _ := g.g.Do(key, func() (any, error) {
		return fn(ctx)
	})
	if err != nil {
		return zero, err
	}
	return v.(V), nil
}

// Forget 主动清除 key（用于失败后允许立即重试，而非等待 in-flight 释放）。
func (g *Group[V]) Forget(key string) { g.g.Forget(key) }
