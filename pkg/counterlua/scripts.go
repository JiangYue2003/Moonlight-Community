// Package counterlua 提供计数模块的预编译 Lua 脚本。
// 三份脚本为计数系统的核心原子操作：
//   - Toggle：原子切换位图位（点赞/收藏 set/unset）
//   - IncrField：原子增减 SDS 中的 4 字节 BigEndian 字段
//   - DecrField：从聚合 Hash 扣减字段（聚合刷写后清理）
package counterlua

import _ "embed"

//go:embed toggle.lua
var Toggle string

//go:embed incr_field.lua
var IncrField string

//go:embed decr_field.lua
var DecrField string
