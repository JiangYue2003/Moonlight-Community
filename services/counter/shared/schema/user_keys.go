// Package schema 中针对用户维度计数（ucnt:{userId}）的字段索引常量。
//
// 与原 Java UserCounterServiceImpl 的字段编号严格对齐：
//
//	idx 0  read           预留位
//	idx 1  followings     关注数
//	idx 2  followers      粉丝数
//	idx 3  posts          已发布帖子数
//	idx 4  likes_received 累计获赞数
//
// SchemaLen / FieldSize 与实体计数（cnt:v1:*）共用，
// 因此 incr_field.lua 可被两套 SDS 共用，不需要维护第二份脚本。
//
// UserSdsKey 已在 keys.go 定义（ucnt:{userId}），此处不重复。
package schema

const (
	UserSchemaId  = "v1"
	UserSchemaLen = SchemaLen // 复用 5
	UserFieldSize = FieldSize // 复用 4

	UserIdxRead          = 0
	UserIdxFollowings    = 1
	UserIdxFollowers     = 2
	UserIdxPosts         = 3
	UserIdxLikesReceived = 4

	UserMetricFollowings    = "followings"
	UserMetricFollowers     = "followers"
	UserMetricPosts         = "posts"
	UserMetricLikesReceived = "likes_received"
)

// UserIdxOf 把 metric 名映射到 SDS 字段索引。
// 未识别的字段返回 -1。read 字段不可写入（返回 -1）。
func UserIdxOf(metric string) int {
	switch metric {
	case UserMetricFollowings:
		return UserIdxFollowings
	case UserMetricFollowers:
		return UserIdxFollowers
	case UserMetricPosts:
		return UserIdxPosts
	case UserMetricLikesReceived:
		return UserIdxLikesReceived
	default:
		return -1
	}
}
