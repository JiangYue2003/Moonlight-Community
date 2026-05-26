package sds

import (
	"encoding/binary"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

func TestDecodeN_VariableLength(t *testing.T) {
	buf := make([]byte, 3*schema.FieldSize)
	binary.BigEndian.PutUint32(buf[0:4], 1)
	binary.BigEndian.PutUint32(buf[4:8], 2)
	binary.BigEndian.PutUint32(buf[8:12], 3)

	out := DecodeN(buf, 3)
	if len(out) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(out))
	}
	for i, v := range []int64{1, 2, 3} {
		if out[i] != v {
			t.Errorf("idx %d: want %d, got %d", i, v, out[i])
		}
	}
}

func TestDecodeN_RequestMoreThanBuffer_PadsZero(t *testing.T) {
	buf := make([]byte, 2*schema.FieldSize)
	binary.BigEndian.PutUint32(buf[0:4], 9)
	binary.BigEndian.PutUint32(buf[4:8], 8)

	out := DecodeN(buf, 5)
	if len(out) != 5 {
		t.Fatalf("expected 5 slots, got %d", len(out))
	}
	if out[0] != 9 || out[1] != 8 {
		t.Errorf("first two should be 9,8 got %v", out[:2])
	}
	for i := 2; i < 5; i++ {
		if out[i] != 0 {
			t.Errorf("idx %d expected 0 (pad), got %d", i, out[i])
		}
	}
}

func TestDecodeN_ZeroN_ReturnsEmpty(t *testing.T) {
	out := DecodeN(make([]byte, 100), 0)
	if len(out) != 0 {
		t.Fatalf("DecodeN(_, 0) should return empty slice, got %d", len(out))
	}
}

func TestDecode_StillCallableForLegacyCallers(t *testing.T) {
	// 老接口必须保持原行为：固定 SchemaLen 数组。
	buf := make([]byte, schema.SchemaLen*schema.FieldSize)
	binary.BigEndian.PutUint32(buf[1*schema.FieldSize:2*schema.FieldSize], 7)
	out := Decode(buf)
	if out[1] != 7 {
		t.Fatalf("legacy Decode lost data: %v", out)
	}
}
