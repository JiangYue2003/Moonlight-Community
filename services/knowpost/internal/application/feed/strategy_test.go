package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestParseStrategy(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want Strategy
	}{
		{name: "empty defaults to hybrid", raw: "", want: StrategyHybrid},
		{name: "push", raw: "push", want: StrategyPush},
		{name: "pull", raw: "pull", want: StrategyPull},
		{name: "hybrid", raw: "hybrid", want: StrategyHybrid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStrategy(tt.raw)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParseStrategyRejectsUnknownValue(t *testing.T) {
	_, err := ParseStrategy("random")

	require.ErrorContains(t, err, "invalid feed strategy")
}

func TestFeedWriterForcedPushIgnoresFollowerThreshold(t *testing.T) {
	redis := NewMockRedisClient()
	kafka := NewMockKafkaProducer()
	writer := NewFeedWriterWithStrategy(
		redis,
		kafka,
		logx.WithContext(context.Background()),
		StrategyPush,
	)

	err := writer.OnPostPublished(context.Background(), 101, 202, BIGV_THRESHOLD+1)

	require.NoError(t, err)
	require.Len(t, kafka.messages, 1)
	require.Len(t, redis.zaddCalls, 1)
	require.Equal(t, "feed:inbox:202", redis.zaddCalls[0].Key)
}

func TestFeedWriterForcedPullIgnoresFollowerThreshold(t *testing.T) {
	redis := NewMockRedisClient()
	kafka := NewMockKafkaProducer()
	writer := NewFeedWriterWithStrategy(
		redis,
		kafka,
		logx.WithContext(context.Background()),
		StrategyPull,
	)

	err := writer.OnPostPublished(context.Background(), 101, 202, 0)

	require.NoError(t, err)
	require.Empty(t, kafka.messages)
	require.Len(t, redis.zaddCalls, 2)
	require.Equal(t, "feed:inbox:202", redis.zaddCalls[0].Key)
	require.Equal(t, "feed:bigv:202", redis.zaddCalls[1].Key)
}
