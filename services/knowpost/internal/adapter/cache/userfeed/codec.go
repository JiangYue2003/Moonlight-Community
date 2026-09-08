package userfeed

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

const (
	recordMagic      = "UFP1"
	recordHeaderSize = 4 + 8 + 8 + 8
)

type cacheRecord struct {
	page       *pb.FeedPage
	freshUntil time.Time
	refreshAt  time.Time
	staleUntil time.Time
}

func encodeRecord(record cacheRecord) ([]byte, error) {
	if err := validateRecord(record); err != nil {
		return nil, err
	}
	body, err := proto.Marshal(record.page)
	if err != nil {
		return nil, err
	}
	raw := make([]byte, recordHeaderSize+len(body))
	copy(raw, recordMagic)
	binary.BigEndian.PutUint64(raw[4:12], uint64(record.freshUntil.UnixNano()))
	binary.BigEndian.PutUint64(raw[12:20], uint64(record.refreshAt.UnixNano()))
	binary.BigEndian.PutUint64(raw[20:28], uint64(record.staleUntil.UnixNano()))
	copy(raw[recordHeaderSize:], body)
	return raw, nil
}

func decodeRecord(raw []byte) (cacheRecord, error) {
	if len(raw) <= recordHeaderSize {
		return cacheRecord{}, errors.New("user feed page record is truncated")
	}
	if string(raw[:4]) != recordMagic {
		return cacheRecord{}, errors.New("user feed page record magic is invalid")
	}
	record := cacheRecord{
		freshUntil: time.Unix(0, int64(binary.BigEndian.Uint64(raw[4:12]))),
		refreshAt:  time.Unix(0, int64(binary.BigEndian.Uint64(raw[12:20]))),
		staleUntil: time.Unix(0, int64(binary.BigEndian.Uint64(raw[20:28]))),
		page:       &pb.FeedPage{},
	}
	if err := proto.Unmarshal(raw[recordHeaderSize:], record.page); err != nil {
		return cacheRecord{}, err
	}
	if err := validateRecord(record); err != nil {
		return cacheRecord{}, err
	}
	return record, nil
}

func validateRecord(record cacheRecord) error {
	if record.page == nil {
		return errors.New("user feed page is nil")
	}
	if record.freshUntil.IsZero() || record.refreshAt.IsZero() || record.staleUntil.IsZero() {
		return errors.New("user feed page record time is zero")
	}
	if record.refreshAt.After(record.freshUntil) {
		return fmt.Errorf("user feed page refresh time %s is after fresh time %s", record.refreshAt, record.freshUntil)
	}
	if record.freshUntil.After(record.staleUntil) {
		return fmt.Errorf("user feed page fresh time %s is after stale time %s", record.freshUntil, record.staleUntil)
	}
	return nil
}
