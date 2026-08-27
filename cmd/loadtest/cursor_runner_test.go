package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type scriptedCursorFeedReader struct {
	mu        sync.Mutex
	responses map[int64]map[string]*feedPageResponse
	calls     []scriptedCursorCall
}

type scriptedCursorCall struct {
	reader  int64
	request feedReadRequest
}

type scriptedOffsetFeedReader struct {
	mu        sync.Mutex
	responses map[int64]map[int32]*feedPageResponse
	calls     []scriptedCursorCall
}

func (r *scriptedOffsetFeedReader) GetUserFeed(
	_ context.Context,
	reader readerIdentity,
	request feedReadRequest,
) (*feedPageResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, scriptedCursorCall{reader: reader.UserID, request: request})
	page := r.responses[reader.UserID][request.Page]
	if page == nil {
		return nil, fmt.Errorf("no scripted response for reader=%d page=%d", reader.UserID, request.Page)
	}
	copyPage := *page
	copyPage.Items = append([]feedItemResponse(nil), page.Items...)
	return &copyPage, nil
}

func (r *scriptedCursorFeedReader) GetUserFeed(
	_ context.Context,
	reader readerIdentity,
	request feedReadRequest,
) (*feedPageResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, scriptedCursorCall{reader: reader.UserID, request: request})
	page := r.responses[reader.UserID][request.Cursor]
	if page == nil {
		return nil, fmt.Errorf("no scripted response for reader=%d cursor=%q", reader.UserID, request.Cursor)
	}
	copyPage := *page
	copyPage.Items = append([]feedItemResponse(nil), page.Items...)
	return &copyPage, nil
}

func TestPrepareCursorTargetsBuildsVerifiedPrefixOutsideMeasuredStage(t *testing.T) {
	client := newScriptedCursorReader([]int64{11, 12})
	readers := []readerIdentity{{UserID: 11}, {UserID: 12}}
	oracle := map[int64][]string{
		11: {"11-1", "11-2", "11-3", "11-4", "11-5", "11-6"},
		12: {"12-1", "12-2", "12-3", "12-4", "12-5", "12-6"},
	}

	prepared, err := prepareCursorTargets(context.Background(), client, readers, cursorTargetSpec{
		TargetPage: 3,
		Size:       2,
		Timeout:    time.Second,
		Oracle:     oracle,
	})

	require.NoError(t, err)
	require.Equal(t, "11-c2", prepared.Targets[11].Cursor)
	require.Equal(t, []string{"11-1", "11-2", "11-3", "11-4"}, prepared.Targets[11].PrefixIDs)
	require.Equal(t, int64(4), prepared.Requests)
	require.GreaterOrEqual(t, prepared.Duration, time.Duration(0))
	require.Len(t, client.calls, 4, "only pages before the measured target should be prepared")

	stage, err := runFixedCursorStage(context.Background(), client, readers, prepared, cursorTargetSpec{
		TargetPage:      3,
		Size:            2,
		Timeout:         time.Second,
		Requests:        4,
		Concurrency:     2,
		Seed:            42,
		RequireItems:    true,
		AllowedCreators: map[int64]struct{}{7: {}},
		Oracle:          oracle,
	})
	require.NoError(t, err)
	require.Equal(t, int64(4), stage.Total)
	require.Equal(t, int64(4), stage.Success)
	require.NotNil(t, stage.Pagination)
	require.Equal(t, "cursor", stage.Pagination.Mode)
	require.Equal(t, 3, stage.Pagination.TargetPage)
	require.Equal(t, int64(4), stage.Pagination.PreparationRequests)
	require.Equal(t, int64(0), stage.Pagination.DuplicateItems)
	require.Equal(t, int64(0), stage.Pagination.OracleMismatches)
}

func TestPrepareCursorTargetsRejectsMissingLoopingAndDuplicateCursorChains(t *testing.T) {
	for _, test := range []struct {
		name  string
		first *feedPageResponse
		c1    *feedPageResponse
		want  string
	}{
		{
			name:  "missing next cursor",
			first: cursorPage([]string{"1", "2"}, true, ""),
			want:  "missing next cursor",
		},
		{
			name:  "cursor loop",
			first: cursorPage([]string{"1", "2"}, true, "c1"),
			c1:    cursorPage([]string{"3", "4"}, true, "c1"),
			want:  "cursor loop",
		},
		{
			name:  "duplicate item",
			first: cursorPage([]string{"1", "2"}, true, "c1"),
			c1:    cursorPage([]string{"2", "3"}, true, "c2"),
			want:  "duplicate feed item",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			client := &scriptedCursorFeedReader{responses: map[int64]map[string]*feedPageResponse{
				1: {"": test.first, "c1": test.c1},
			}}
			_, err := prepareCursorTargets(context.Background(), client, []readerIdentity{{UserID: 1}}, cursorTargetSpec{
				TargetPage: 3, Size: 2, Timeout: time.Second,
			})
			require.ErrorContains(t, err, test.want)
		})
	}
}

func TestRunSequentialCursorStageMeasuresEveryPageAndChecksOracle(t *testing.T) {
	client := newScriptedCursorReader([]int64{11, 12})
	readers := []readerIdentity{{UserID: 11}, {UserID: 12}}
	oracle := map[int64][]string{
		11: {"11-1", "11-2", "11-3", "11-4", "11-5", "11-6"},
		12: {"12-1", "12-2", "12-3", "12-4", "12-5", "12-6"},
	}

	stage, err := runSequentialCursorStage(context.Background(), client, readers, cursorTargetSpec{
		TargetPage: 3, Size: 2, Timeout: time.Second,
		Requests: 2, Concurrency: 2, Seed: 9,
		RequireItems: true, AllowedCreators: map[int64]struct{}{7: {}}, Oracle: oracle,
	})

	require.NoError(t, err)
	require.Equal(t, int64(6), stage.Total)
	require.Equal(t, int64(6), stage.Success)
	require.Equal(t, int64(2), stage.Pagination.Sequences)
	require.Equal(t, int64(6), stage.Pagination.PagesTraversed)
	require.Zero(t, stage.Pagination.DuplicateItems)
	require.Zero(t, stage.Pagination.OracleMismatches)
	require.Zero(t, stage.Pagination.CursorLoops)
	require.Zero(t, stage.Pagination.EarlyTerminations)
}

func TestSequentialPageScenarioUsesOffsetControl(t *testing.T) {
	target, ok := sequentialOffsetTargetPage("sequential-page-50")
	require.True(t, ok)
	require.Equal(t, 50, target)
	_, cursorOK := sequentialCursorTargetPage("sequential-page-50")
	require.False(t, cursorOK)

	client := &scriptedOffsetFeedReader{responses: map[int64]map[int32]*feedPageResponse{
		11: {
			1: cursorPage([]string{"11-1", "11-2"}, true, "ignored-c1"),
			2: cursorPage([]string{"11-3", "11-4"}, true, "ignored-c2"),
			3: cursorPage([]string{"11-5", "11-6"}, false, ""),
		},
	}}
	stage, err := runSequentialOffsetStage(context.Background(), client, []readerIdentity{{UserID: 11}}, cursorTargetSpec{
		TargetPage: 3, Size: 2, Timeout: time.Second,
		Requests: 1, Concurrency: 1, Seed: 9,
		RequireItems: true, AllowedCreators: map[int64]struct{}{7: {}},
		Oracle: map[int64][]string{11: {"11-1", "11-2", "11-3", "11-4", "11-5", "11-6"}},
	})

	require.NoError(t, err)
	require.Equal(t, int64(3), stage.Total)
	require.Equal(t, int64(3), stage.Success)
	require.Equal(t, "page-sequential", stage.Pagination.Mode)
	require.Equal(t, int64(1), stage.Pagination.Sequences)
	require.Equal(t, int64(3), stage.Pagination.PagesTraversed)
	require.Zero(t, stage.Pagination.DuplicateItems)
	require.Zero(t, stage.Pagination.OracleMismatches)
	require.Len(t, client.calls, 3)
	require.Equal(t, []int32{1, 2, 3}, []int32{
		client.calls[0].request.Page,
		client.calls[1].request.Page,
		client.calls[2].request.Page,
	})
	for _, call := range client.calls {
		require.Empty(t, call.request.Cursor)
	}
}

func newScriptedCursorReader(readers []int64) *scriptedCursorFeedReader {
	responses := make(map[int64]map[string]*feedPageResponse, len(readers))
	for _, reader := range readers {
		prefix := fmt.Sprintf("%d-", reader)
		responses[reader] = map[string]*feedPageResponse{
			"": cursorPage([]string{prefix + "1", prefix + "2"}, true, fmt.Sprintf("%d-c1", reader)),
			fmt.Sprintf("%d-c1", reader): cursorPage(
				[]string{prefix + "3", prefix + "4"}, true, fmt.Sprintf("%d-c2", reader),
			),
			fmt.Sprintf("%d-c2", reader): cursorPage([]string{prefix + "5", prefix + "6"}, false, ""),
		}
	}
	return &scriptedCursorFeedReader{responses: responses}
}

func cursorPage(ids []string, hasMore bool, next string) *feedPageResponse {
	items := make([]feedItemResponse, 0, len(ids))
	for _, id := range ids {
		items = append(items, feedItemResponse{ID: id, CreatorID: 7})
	}
	return &feedPageResponse{Items: items, HasMore: hasMore, Size: int32(len(items)), NextCursor: next}
}
