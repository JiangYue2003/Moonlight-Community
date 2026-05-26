package sds

import (
	"encoding/binary"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

func TestDecode_ZeroLength_ReturnsAllZero(t *testing.T) {
	out := Decode(nil)
	for i, v := range out {
		if v != 0 {
			t.Errorf("idx %d should be zero, got %d", i, v)
		}
	}
}

func TestDecode_ShortBuffer_FillsWhatItCanThenZero(t *testing.T) {
	// 只填前两个字段（8 字节）；剩余字段必须保持 0。
	buf := make([]byte, 2*schema.FieldSize)
	binary.BigEndian.PutUint32(buf[0:4], 7)  // idx=0 (read 预留位)
	binary.BigEndian.PutUint32(buf[4:8], 11) // idx=1 (like)

	out := Decode(buf)
	if out[0] != 7 {
		t.Errorf("idx0 expected 7, got %d", out[0])
	}
	if out[1] != 11 {
		t.Errorf("idx1 expected 11, got %d", out[1])
	}
	for i := 2; i < schema.SchemaLen; i++ {
		if out[i] != 0 {
			t.Errorf("idx %d expected 0 (out-of-buffer), got %d", i, out[i])
		}
	}
}

func TestDecode_FullBuffer_AllFieldsPopulated(t *testing.T) {
	buf := make([]byte, schema.SchemaLen*schema.FieldSize)
	values := []uint32{1, 1234, 5678, 9012, 3456}
	for i, v := range values {
		binary.BigEndian.PutUint32(buf[i*schema.FieldSize:(i+1)*schema.FieldSize], v)
	}
	out := Decode(buf)
	for i, want := range values {
		if out[i] != int64(want) {
			t.Errorf("idx %d: want %d, got %d", i, want, out[i])
		}
	}
}

func TestDecode_BigEndianContract(t *testing.T) {
	// Lua 脚本写入 SDS 时使用 BigEndian；Decode 必须与之严格一致，
	// 否则 like/fav 计数会被解释为大于 2^24 的离谱值。
	buf := []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05}
	out := Decode(buf)
	if out[0] != 1 {
		t.Fatalf("BigEndian decoding broken: idx0 = %d (raw 00 00 00 01 should be 1)", out[0])
	}
	if out[1] != 5 {
		t.Fatalf("BigEndian decoding broken: idx1 = %d (raw 00 00 00 05 should be 5)", out[1])
	}
}

func TestDecode_OversizedBufferIgnoresExtraBytes(t *testing.T) {
	// 多余字节不应该被读入；解码后应返回一个稳定的 SchemaLen 数组。
	buf := make([]byte, schema.SchemaLen*schema.FieldSize+16)
	for i := 0; i < schema.SchemaLen; i++ {
		binary.BigEndian.PutUint32(buf[i*schema.FieldSize:(i+1)*schema.FieldSize], uint32(i+1))
	}
	for i := schema.SchemaLen * schema.FieldSize; i < len(buf); i++ {
		buf[i] = 0xFF
	}
	out := Decode(buf)
	for i := 0; i < schema.SchemaLen; i++ {
		if out[i] != int64(i+1) {
			t.Errorf("idx %d: want %d, got %d", i, i+1, out[i])
		}
	}
}

func TestDecode_LargeUint32_ReadsAsUnsigned(t *testing.T) {
	// SDS 字段以无符号 32 位写入；Decode 用 int64 接住，
	// 保证 4_000_000_000 这类大值不会被误判为负数。
	buf := make([]byte, schema.SchemaLen*schema.FieldSize)
	binary.BigEndian.PutUint32(buf[schema.FieldSize:2*schema.FieldSize], 4_000_000_000)
	out := Decode(buf)
	if out[1] != 4_000_000_000 {
		t.Fatalf("unsigned read failed: got %d", out[1])
	}
}
