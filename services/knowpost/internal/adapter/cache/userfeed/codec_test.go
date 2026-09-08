package userfeed

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestL2RecordUsesBinaryTimeHeaderAndProtobufBody(t *testing.T) {
	page := testPage(1)
	record := cacheRecord{
		page:       page,
		freshUntil: time.Unix(1_700_000_005, 123),
		refreshAt:  time.Unix(1_700_000_004, 123),
		staleUntil: time.Unix(1_700_000_015, 123),
	}
	raw, err := encodeRecord(record)
	require.NoError(t, err)
	require.True(t, bytes.HasPrefix(raw, []byte("UFP1")))
	require.NotContains(t, string(raw), `"fresh_until"`)

	decoded, err := decodeRecord(raw)
	require.NoError(t, err)
	require.Equal(t, record.freshUntil, decoded.freshUntil)
	require.Equal(t, record.refreshAt, decoded.refreshAt)
	require.Equal(t, record.staleUntil, decoded.staleUntil)
	require.True(t, proto.Equal(page, decoded.page))
}

func TestL2RecordRejectsMalformedOrInvalidTimeHeader(t *testing.T) {
	for _, raw := range [][]byte{
		nil,
		[]byte("UFP1"),
		append([]byte("BAD!"), make([]byte, 32)...),
	} {
		_, err := decodeRecord(raw)
		require.Error(t, err)
	}

	record := cacheRecord{
		page:       testPage(1),
		freshUntil: time.Unix(10, 0),
		refreshAt:  time.Unix(11, 0),
		staleUntil: time.Unix(12, 0),
	}
	_, err := encodeRecord(record)
	require.Error(t, err)
}
