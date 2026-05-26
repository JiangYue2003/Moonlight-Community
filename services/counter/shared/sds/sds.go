// Package sds 解析 SDS 二进制：每字段 4 字节 BigEndian Int32（无符号读取）。
//
// 提供两类入口：
//   - DecodeN(raw, n)：通用变长解码，user counter / 未来扩展用
//   - Decode(raw)    ：兼容旧调用，固定按实体计数 SchemaLen 返回数组
package sds

import (
	"encoding/binary"

	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

// DecodeN 通用解码：读取前 n 个 4 字节字段。输入长度不足时高位补零；超出忽略。
// n<=0 返回空切片。
func DecodeN(raw []byte, n int) []int64 {
	if n <= 0 {
		return []int64{}
	}
	out := make([]int64, n)
	for i := 0; i < n; i++ {
		off := i * schema.FieldSize
		if off+schema.FieldSize > len(raw) {
			break
		}
		out[i] = int64(binary.BigEndian.Uint32(raw[off : off+schema.FieldSize]))
	}
	return out
}

// Decode 兼容旧调用方：固定按实体计数 SchemaLen 解码（5 字段）。
func Decode(raw []byte) [schema.SchemaLen]int64 {
	var out [schema.SchemaLen]int64
	dec := DecodeN(raw, schema.SchemaLen)
	copy(out[:], dec)
	return out
}
