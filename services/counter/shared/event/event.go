// Package event 定义计数模块跨服务的 Kafka 事件结构。
package event

// CounterEvent Kafka 消息体；与 Java CounterEvent 字段严格一致。
type CounterEvent struct {
	EntityType string `json:"entityType"`
	EntityId   string `json:"entityId"`
	Metric     string `json:"metric"`
	Idx        int    `json:"idx"`
	UserId     int64  `json:"userId"`
	Delta      int32  `json:"delta"`
}

// TopicEvents 默认 topic 名（与 Java CounterTopics.EVENTS 一致）。
const TopicEvents = "counter-events"
