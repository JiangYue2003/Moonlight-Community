package feedcursor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const (
	Version          = 1
	MaxEncodedLength = 256
)

var ErrInvalidCursor = errors.New("invalid feed cursor")

type Cursor struct {
	SortTime int64
	PostID   int64
}

type payload struct {
	Version  int   `json:"v"`
	SortTime int64 `json:"t"`
	PostID   int64 `json:"p"`
}

func Encode(cursor Cursor) (string, error) {
	if err := validate(cursor); err != nil {
		return "", err
	}
	raw, err := json.Marshal(payload{
		Version:  Version,
		SortTime: cursor.SortTime,
		PostID:   cursor.PostID,
	})
	if err != nil {
		return "", fmt.Errorf("%w: encode payload: %v", ErrInvalidCursor, err)
	}
	encoded := base64.RawURLEncoding.EncodeToString(raw)
	if len(encoded) > MaxEncodedLength {
		return "", fmt.Errorf("%w: encoded length exceeds %d", ErrInvalidCursor, MaxEncodedLength)
	}
	return encoded, nil
}

func Decode(encoded string) (Cursor, error) {
	if encoded == "" || len(encoded) > MaxEncodedLength {
		return Cursor{}, fmt.Errorf("%w: encoded length must be between 1 and %d", ErrInvalidCursor, MaxEncodedLength)
	}
	raw, err := base64.RawURLEncoding.Strict().DecodeString(encoded)
	if err != nil {
		return Cursor{}, fmt.Errorf("%w: decode base64url: %v", ErrInvalidCursor, err)
	}

	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var value payload
	if err := decoder.Decode(&value); err != nil {
		return Cursor{}, fmt.Errorf("%w: decode payload: %v", ErrInvalidCursor, err)
	}
	if err := rejectTrailingJSON(decoder); err != nil {
		return Cursor{}, err
	}
	if value.Version != Version {
		return Cursor{}, fmt.Errorf("%w: unsupported version %d", ErrInvalidCursor, value.Version)
	}

	cursor := Cursor{SortTime: value.SortTime, PostID: value.PostID}
	canonical, err := Encode(cursor)
	if err != nil {
		return Cursor{}, err
	}
	if canonical != encoded {
		return Cursor{}, fmt.Errorf("%w: non-canonical encoding", ErrInvalidCursor)
	}
	return cursor, nil
}

func validate(cursor Cursor) error {
	if cursor.SortTime <= 0 {
		return fmt.Errorf("%w: sort time must be positive", ErrInvalidCursor)
	}
	if cursor.PostID <= 0 {
		return fmt.Errorf("%w: post id must be positive", ErrInvalidCursor)
	}
	return nil
}

func rejectTrailingJSON(decoder *json.Decoder) error {
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return fmt.Errorf("%w: trailing payload", ErrInvalidCursor)
		}
		return fmt.Errorf("%w: trailing payload: %v", ErrInvalidCursor, err)
	}
	return nil
}
