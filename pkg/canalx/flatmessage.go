// Package canalx 解析 Canal Server 在 flatMessage=true 模式下输出到 Kafka 的 JSON。
//
// Canal flatMessage 顶层字段（v1.1.7）：
//
//	database / table / type / isDdl / data[] / old[] / ts / es / sql / sqlType / mysqlType / pkNames
//
// 注意：Canal 把所有字段值都序列化为字符串（即使 BIGINT 也是 "12345"），
// 所以 data/old 的值类型用 string，由 outbox.go 等上层做类型转换。
package canalx

import (
	"encoding/json"
	"errors"
)

// FlatMessage Canal 扁平消息体。
type FlatMessage struct {
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Type     string              `json:"type"` // INSERT / UPDATE / DELETE
	IsDdl    bool                `json:"isDdl"`
	Data     []map[string]string `json:"data"`
	Old      []map[string]string `json:"old"`
	Ts       int64               `json:"ts"`
	Es       int64               `json:"es"`
	Sql      string              `json:"sql"`
	PkNames  []string            `json:"pkNames"`
}

// ErrInvalidJson 输入字节既不是合法 JSON 也不是 Canal flatMessage。
var ErrInvalidJson = errors.New("canalx: invalid flat message json")

// ParseFlat 反序列化字节流为 FlatMessage。
// DDL 事件（isDdl=true）由调用方决定是否处理；本函数不过滤。
func ParseFlat(raw []byte) (*FlatMessage, error) {
	if len(raw) == 0 {
		return nil, ErrInvalidJson
	}
	var m FlatMessage
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, ErrInvalidJson
	}
	return &m, nil
}
