package query

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// EncodeCursor 把 ES 返回的 sort 数组转成 base64url 字符串。
//
// sort 是 []json.RawMessage（每元素已是合法 JSON），我们直接合成 JSON 数组再 base64。
func EncodeCursor(sort []json.RawMessage) string {
	if len(sort) == 0 {
		return ""
	}
	buf, err := json.Marshal(sort)
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}

// DecodeCursor base64url → []any（保持原数字类型）。空串返回 nil（首页）。
func DecodeCursor(s string) ([]any, error) {
	if s == "" {
		return nil, nil
	}
	buf, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var out []any
	dec := json.NewDecoder(strings.NewReader(string(buf)))
	dec.UseNumber()
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, errors.New("cursor empty")
	}
	return out, nil
}
