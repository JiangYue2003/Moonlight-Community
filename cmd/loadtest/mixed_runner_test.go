package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mixedFeedReader struct{}

func (mixedFeedReader) GetUserFeed(
	context.Context,
	readerIdentity,
	feedReadRequest,
) (*feedPageResponse, error) {
	return &feedPageResponse{Items: []feedItemResponse{{ID: "post-1", CreatorID: 7}}}, nil
}

func TestRunMixedStageKeepsExactRequestMixAndGlobalConcurrency(t *testing.T) {
	publisher := &recordingPublisher{nextID: 100}

	stages, postIDs, err := runMixedStage(
		context.Background(), publisher, mixedFeedReader{},
		[]int64{7}, []readerIdentity{{UserID: 1}},
		mixedSpec{
			Requests: 100, Concurrency: 8, ReadPercent: 90,
			RequestTimeout: time.Second, Page: 1, Size: 20, RunID: "run-1", Seed: 42,
			AllowedCreators: map[int64]struct{}{7: {}}, RequireItems: true,
		},
	)

	require.NoError(t, err)
	require.Len(t, stages, 6)
	require.Equal(t, "read", stages[0].Name)
	require.Equal(t, int64(90), stages[0].Total)
	require.Equal(t, "publish_total", stages[1].Name)
	require.Equal(t, int64(10), stages[1].Total)
	require.Len(t, postIDs, 10)
}

func TestRunMixedStageSupportsDurationMode(t *testing.T) {
	publisher := &recordingPublisher{nextID: 100}

	stages, postIDs, err := runMixedStage(
		context.Background(), publisher, mixedFeedReader{},
		[]int64{7}, []readerIdentity{{UserID: 1}},
		mixedSpec{
			Duration: 20 * time.Millisecond, Concurrency: 4, ReadPercent: 80,
			RequestTimeout: time.Second, Page: 1, Size: 20, RunID: "run-1", Seed: 42,
			AllowedCreators: map[int64]struct{}{7: {}}, RequireItems: true,
		},
	)

	require.NoError(t, err)
	require.Greater(t, stages[0].Total, int64(0))
	require.Greater(t, stages[1].Total, int64(0))
	require.Equal(t, int(stages[1].Success), len(postIDs))
	total := stages[0].Total + stages[1].Total
	require.InDelta(t, 0.8, float64(stages[0].Total)/float64(total), 0.02)
}

func TestRunMixedStageRejectsInvalidPercentage(t *testing.T) {
	_, _, err := runMixedStage(
		context.Background(), &recordingPublisher{}, mixedFeedReader{},
		[]int64{7}, []readerIdentity{{UserID: 1}},
		mixedSpec{Requests: 1, Concurrency: 1, ReadPercent: 100, RequestTimeout: time.Second, Page: 1, Size: 20, RunID: "run-1"},
	)

	require.ErrorContains(t, err, "read percent")
}

func TestMixedRequestScheduleInterleavesWritesWithoutChangingExactRatio(t *testing.T) {
	writes90 := make([]int64, 0)
	for ordinal := int64(1); ordinal <= 20; ordinal++ {
		if !mixedRequestIsRead(ordinal, 90) {
			writes90 = append(writes90, ordinal)
		}
	}
	require.Equal(t, []int64{10, 20}, writes90)

	writes80 := make([]int64, 0)
	for ordinal := int64(1); ordinal <= 10; ordinal++ {
		if !mixedRequestIsRead(ordinal, 80) {
			writes80 = append(writes80, ordinal)
		}
	}
	require.Equal(t, []int64{5, 10}, writes80)
}
