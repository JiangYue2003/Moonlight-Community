package query

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

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
