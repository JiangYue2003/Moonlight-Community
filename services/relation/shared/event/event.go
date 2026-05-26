// Package event 跨进程的关系领域事件结构。
//
// 阶段3 主要被两边使用：
//   - relation-rpc：写 outbox 时把 RelationEvent 序列化为 payload
//   - relation-syncer：消费 canal-outbox 后从 payload 反序列化
package event

const (
	TypeFollowCreated  = "FollowCreated"
	TypeFollowCanceled = "FollowCanceled"
)

// RelationEvent 与原 Java RelationEvent record 字段一致，序列化后写入 outbox.payload。
type RelationEvent struct {
	Type       string `json:"type"`
	FromUserId int64  `json:"fromUserId"`
	ToUserId   int64  `json:"toUserId"`
	Id         int64  `json:"id"`
}
