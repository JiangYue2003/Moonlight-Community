package schema

import "fmt"

// 与 Java 版 CounterKeys / UserCounterKeys / BitmapShard 严格对齐。

const ChunkSize = 32768 // 每分片 32768 位 = 4KB

// SdsKey 实体计数 SDS：cnt:v1:{etype}:{eid}
func SdsKey(etype, eid string) string {
	return fmt.Sprintf("cnt:%s:%s:%s", SchemaId, etype, eid)
}

// BitmapKey 用户行为位图分片：bm:{metric}:{etype}:{eid}:{chunk}
func BitmapKey(metric, etype, eid string, chunk int64) string {
	return fmt.Sprintf("bm:%s:%s:%s:%d", metric, etype, eid, chunk)
}

// AggKey 聚合桶：agg:v1:{etype}:{eid}
func AggKey(etype, eid string) string {
	return fmt.Sprintf("agg:%s:%s:%s", SchemaId, etype, eid)
}

// UserSdsKey 用户维度 SDS：ucnt:{userId}
func UserSdsKey(userId int64) string { return fmt.Sprintf("ucnt:%d", userId) }

// ChunkOf 用户 id 的分片 idx。
func ChunkOf(userId int64) int64 { return userId / ChunkSize }

// BitOf 用户 id 在分片内的 bit offset。
func BitOf(userId int64) int64 { return userId % ChunkSize }
