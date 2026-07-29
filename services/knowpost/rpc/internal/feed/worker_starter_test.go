package feed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeedFanoutQueueConfigStartsConsumers(t *testing.T) {
	config := newFeedFanoutQueueConfig([]string{"127.0.0.1:9092"})

	require.Equal(t, FEED_FANOUT_TOPIC, config.Topic)
	require.Equal(t, "feed-fanout-group", config.Group)
	require.Positive(t, config.Conns)
	require.Positive(t, config.Consumers)
	require.Positive(t, config.Processors)
	require.Positive(t, config.MinBytes)
	require.GreaterOrEqual(t, config.MaxBytes, config.MinBytes)
}
