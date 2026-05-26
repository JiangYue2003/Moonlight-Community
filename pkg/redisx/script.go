package redisx

import "github.com/redis/go-redis/v9"

// Script 是 *redis.Script 的别名；上层用 NewScript 创建后由 svc 持有，
// 客户端会自动 EVALSHA 缓存命中、不命中时回退 EVAL。
type Script = redis.Script

// NewScript 构造预编译脚本；调用 .Run / .Eval 时由 go-redis 自动缓存 SHA1。
func NewScript(src string) *Script {
	return redis.NewScript(src)
}
