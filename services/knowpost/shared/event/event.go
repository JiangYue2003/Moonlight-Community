// Package event 定义 knowpost 领域事件结构（写入 outbox.payload，被 search-indexer / rag-indexer 等下游消费）。
package event

// 事件类型常量。注意：与 outbox.aggregate_type=knowpost 配套。
const (
	TypeKnowPostCreated   = "KnowPostCreated"   // 新建草稿（一般不广播）
	TypeKnowPostPublished = "KnowPostPublished" // 发布
	TypeKnowPostUpdated   = "KnowPostUpdated"   // 已发布的内容/元信息更新
	TypeKnowPostDeleted   = "KnowPostDeleted"   // 删除
)

// AggregateType 给 outbox 用的聚合根类型。
const AggregateType = "knowpost"

// KnowPostEvent 是 outbox.payload 的 JSON 结构。
//
// 字段尽量精简：下游需要更多字段时统一回拉 knowpost-rpc.GetById 拿全量。
type KnowPostEvent struct {
	Type   string `json:"type"`             // 上面的常量之一
	PostId int64  `json:"postId"`           // 受影响的知文 id
	Author int64  `json:"author,omitempty"` // 作者 id（删除事件可能没有完整 row 时也尽量带上）
}
