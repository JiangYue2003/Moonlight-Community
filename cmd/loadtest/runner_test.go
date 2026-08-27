package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type deadlineAwareFeedReader struct{}

func (deadlineAwareFeedReader) GetUserFeed(
	ctx context.Context,
	_ readerIdentity,
	_ feedReadRequest,
) (*feedPageResponse, error) {
	select {
	case <-time.After(20 * time.Millisecond):
		return &feedPageResponse{Items: []feedItemResponse{{ID: "1"}}}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type recordingFeedReader struct {
	mu      sync.Mutex
	readers []readerIdentity
	active  int
	maxSeen int
}

func (r *recordingFeedReader) GetUserFeed(
	_ context.Context,
	reader readerIdentity,
	_ feedReadRequest,
) (*feedPageResponse, error) {
	r.mu.Lock()
	r.readers = append(r.readers, reader)
	r.active++
	if r.active > r.maxSeen {
		r.maxSeen = r.active
	}
	r.mu.Unlock()

	time.Sleep(time.Millisecond)

	r.mu.Lock()
	r.active--
	r.mu.Unlock()
	return &feedPageResponse{Items: []feedItemResponse{{ID: "1"}}}, nil
}

func TestRunReadStageUsesBoundedWorkersAndRecordsEveryRequest(t *testing.T) {
	client := &recordingFeedReader{}
	readers := []readerIdentity{{UserID: 1}, {UserID: 2}, {UserID: 3}}

	result, err := runReadStage(context.Background(), client, readers, loadSpec{
		Requests:       25,
		Concurrency:    4,
		RequestTimeout: time.Second,
		Page:           1,
		Size:           20,
		Seed:           42,
	})

	require.NoError(t, err)
	require.Equal(t, int64(25), result.Total)
	require.Equal(t, int64(25), result.Success)
	require.Zero(t, result.Failed)
	require.LessOrEqual(t, client.maxSeen, 4)
	require.Len(t, client.readers, 25)
}

func TestDeterministicIndexDependsOnSeedAndOrdinalNotWorkerScheduling(t *testing.T) {
	first := make([]int, 100)
	second := make([]int, 100)
	for ordinal := 1; ordinal <= 100; ordinal++ {
		first[ordinal-1] = deterministicIndex(42, int64(ordinal), 7)
	}
	for ordinal := 100; ordinal >= 1; ordinal-- {
		second[ordinal-1] = deterministicIndex(42, int64(ordinal), 7)
	}

	require.Equal(t, first, second)
	require.NotEqual(t, first, func() []int {
		changed := make([]int, 100)
		for ordinal := 1; ordinal <= 100; ordinal++ {
			changed[ordinal-1] = deterministicIndex(43, int64(ordinal), 7)
		}
		return changed
	}())
}

func TestRunReadStageRejectsInvalidLoad(t *testing.T) {
	_, err := runReadStage(context.Background(), &recordingFeedReader{}, nil, loadSpec{
		Requests:    1,
		Concurrency: 0,
	})

	require.ErrorContains(t, err, "concurrency")
}

func TestPrimeFeedMetricsOnlyForPublishOnlyScenario(t *testing.T) {
	reader := &recordingFeedReader{}
	identities := []readerIdentity{{UserID: 1}}
	allowed := map[int64]struct{}{0: {}}

	primed, err := primeFeedMetricsForScenario(
		context.Background(), "publish", reader, identities, allowed, time.Second, 20,
	)

	require.NoError(t, err)
	require.True(t, primed)
	require.Len(t, reader.readers, 1)

	primed, err = primeFeedMetricsForScenario(
		context.Background(), "mixed-90-10", reader, identities, allowed, time.Second, 20,
	)
	require.NoError(t, err)
	require.False(t, primed)
	require.Len(t, reader.readers, 1)
}

type duplicateFeedReader struct{}

func (duplicateFeedReader) GetUserFeed(
	context.Context,
	readerIdentity,
	feedReadRequest,
) (*feedPageResponse, error) {
	return &feedPageResponse{Items: []feedItemResponse{{ID: "1"}, {ID: "1"}}}, nil
}

type invalidContentFeedReader struct {
	items []feedItemResponse
}

func (r invalidContentFeedReader) GetUserFeed(context.Context, readerIdentity, feedReadRequest) (*feedPageResponse, error) {
	return &feedPageResponse{Items: r.items}, nil
}

func TestRunReadStageRejectsEmptyAndUnexpectedCreatorsWhenRequired(t *testing.T) {
	spec := loadSpec{
		Requests: 1, Concurrency: 1, RequestTimeout: time.Second, Page: 1, Size: 20,
		RequireItems: true, AllowedCreators: map[int64]struct{}{10: {}},
	}
	readers := []readerIdentity{{UserID: 1}}

	empty, err := runReadStage(context.Background(), invalidContentFeedReader{}, readers, spec)
	require.NoError(t, err)
	require.Equal(t, int64(1), empty.Failed)

	wrong, err := runReadStage(context.Background(), invalidContentFeedReader{
		items: []feedItemResponse{{ID: "1", CreatorID: 99}},
	}, readers, spec)
	require.NoError(t, err)
	require.Equal(t, int64(1), wrong.Failed)
}

func TestRunReadStageCountsDuplicateFeedItemsAsCorrectnessFailure(t *testing.T) {
	result, err := runReadStage(context.Background(), duplicateFeedReader{}, []readerIdentity{{UserID: 1}}, loadSpec{
		Requests:       1,
		Concurrency:    1,
		RequestTimeout: time.Second,
		Page:           1,
		Size:           20,
	})

	require.NoError(t, err)
	require.Equal(t, int64(1), result.Failed)
	require.Equal(t, int64(1), result.Errors["request_error"])
}

func TestRunReadStageDurationStopsNewWorkWithoutCancellingInflightRequest(t *testing.T) {
	result, err := runReadStage(context.Background(), deadlineAwareFeedReader{}, []readerIdentity{{UserID: 1}}, loadSpec{
		Duration:       10 * time.Millisecond,
		Concurrency:    1,
		RequestTimeout: 100 * time.Millisecond,
		Page:           1,
		Size:           20,
	})

	require.NoError(t, err)
	require.Equal(t, int64(1), result.Total)
	require.Equal(t, int64(1), result.Success)
	require.Zero(t, result.Failed)
}
