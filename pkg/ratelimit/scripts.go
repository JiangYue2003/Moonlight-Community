// Package ratelimit 提供基于 Lua 令牌桶的分布式限流。
//
// 与计数模块的 counterlua 同款 //go:embed 模式，避免 Lua 字符串散落在业务代码里。
package ratelimit

import _ "embed"

//go:embed follow.lua
var TokenBucketScript string
