// Package schema 定义计数模块的 SDS 结构与 metric 索引（与 Java 版严格对齐）。
package schema

const (
	SchemaId  = "v1"
	FieldSize = 4 // 每字段 4 字节 BigEndian Int32
	SchemaLen = 5 // 字段数

	IdxRead    = 0 // 预留
	IdxLike    = 1
	IdxFav     = 2
	IdxComment = 3 // 预留
	IdxRepost  = 4 // 预留

	MetricLike = "like"
	MetricFav  = "fav"
)

// IdxOf 返回 metric 在 SDS 中的字段索引；未支持返回 -1。
func IdxOf(metric string) int {
	switch metric {
	case MetricLike:
		return IdxLike
	case MetricFav:
		return IdxFav
	default:
		return -1
	}
}

// Supported metric 列表（用于 GetCounts 默认返回）。
var Supported = []string{MetricLike, MetricFav}
