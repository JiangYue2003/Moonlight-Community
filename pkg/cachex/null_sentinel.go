package cachex

import "bytes"

// 哨兵值用于防止缓存穿透：DB 返回 ErrNotFound 时把 sentinel 写入 L1+L2。
var nullSentinelBytes = []byte("__CACHEX_NULL__")

// NullSentinel 返回哨兵字节切片的副本。
func NullSentinel() []byte {
	out := make([]byte, len(nullSentinelBytes))
	copy(out, nullSentinelBytes)
	return out
}

// IsNullSentinelBytes 判断字节是否是哨兵。
func IsNullSentinelBytes(b []byte) bool {
	return bytes.Equal(b, nullSentinelBytes)
}

// isNullSentinel 兼容 L1 中可能存的 []byte 或字符串。
func isNullSentinel(v any) bool {
	switch x := v.(type) {
	case []byte:
		return IsNullSentinelBytes(x)
	case string:
		return x == string(nullSentinelBytes)
	default:
		return false
	}
}
