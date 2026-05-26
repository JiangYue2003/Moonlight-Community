// outbox.go 提供 outbox 表行的提取，对应 Java 端 OutboxMessageUtil。
package canalx

import (
	"strconv"
	"strings"
)

// OutboxRow 一行 outbox 表的领域事件。
//
// payload 是字符串：因 Canal flatMessage 把 JSON 字段透传为字符串，
// 由下游 syncer 的 handler 自己 json.Unmarshal 成具体事件结构（RelationEvent 等）。
type OutboxRow struct {
	Id            int64
	AggregateType string
	AggregateId   int64
	Type          string
	Payload       string
}

// ExtractOutboxRows 从 FlatMessage 中过滤出 outbox 表的 INSERT/UPDATE 行。
//
//   - 非 outbox 表 → 返回空
//   - DELETE 类型 → 跳过（outbox 是只追加表，DELETE 通常是 GC worker 清理产生）
//   - DDL → 跳过
//   - 字段缺失或 id 解析失败 → 跳过该行（不抛错，便于消费者继续处理其它 row）
func ExtractOutboxRows(m *FlatMessage) []OutboxRow {
	if m == nil || m.IsDdl || !strings.EqualFold(m.Table, "outbox") {
		return nil
	}
	if m.Type != "INSERT" && m.Type != "UPDATE" {
		return nil
	}
	out := make([]OutboxRow, 0, len(m.Data))
	for _, row := range m.Data {
		id, err := strconv.ParseInt(row["id"], 10, 64)
		if err != nil {
			continue
		}
		aggId, _ := strconv.ParseInt(row["aggregate_id"], 10, 64)
		out = append(out, OutboxRow{
			Id:            id,
			AggregateType: row["aggregate_type"],
			AggregateId:   aggId,
			Type:          row["type"],
			Payload:       row["payload"],
		})
	}
	return out
}
