package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type recordingKafkaPusher struct {
	key   string
	value string
}

func (p *recordingKafkaPusher) PushWithKey(_ context.Context, key, value string) error {
	p.key = key
	p.value = value
	return nil
}

func TestKafkaProducerAdapterPreservesMessageKey(t *testing.T) {
	pusher := &recordingKafkaPusher{}
	adapter := &KafkaProducerAdapter{pusher: pusher}

	err := adapter.SendMessage(
		context.Background(),
		FEED_FANOUT_TOPIC,
		"post-123",
		[]byte(`{"post_id":123}`),
	)

	require.NoError(t, err)
	require.Equal(t, "post-123", pusher.key)
	require.Equal(t, `{"post_id":123}`, pusher.value)
}
