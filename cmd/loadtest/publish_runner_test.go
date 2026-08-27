package main

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type recordingPublisher struct {
	mu        sync.Mutex
	nextID    int64
	confirmed []*knowpostpb.ConfirmContentReq
	published []*knowpostpb.PublishReq
}

type deadlineAwarePublisher struct {
	recordingPublisher
}

func (p *deadlineAwarePublisher) CreateDraft(
	ctx context.Context,
	in *knowpostpb.CreateDraftReq,
	options ...grpc.CallOption,
) (*knowpostpb.CreateDraftResp, error) {
	select {
	case <-time.After(20 * time.Millisecond):
		return p.recordingPublisher.CreateDraft(ctx, in, options...)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (p *recordingPublisher) CreateDraft(
	context.Context,
	*knowpostpb.CreateDraftReq,
	...grpc.CallOption,
) (*knowpostpb.CreateDraftResp, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.nextID++
	return &knowpostpb.CreateDraftResp{Id: strconv.FormatInt(p.nextID, 10)}, nil
}

func (p *recordingPublisher) PatchMetadata(
	context.Context,
	*knowpostpb.PatchMetadataReq,
	...grpc.CallOption,
) (*knowpostpb.KnowPostDetail, error) {
	return &knowpostpb.KnowPostDetail{}, nil
}

func (p *recordingPublisher) ConfirmContent(
	_ context.Context,
	in *knowpostpb.ConfirmContentReq,
	_ ...grpc.CallOption,
) (*knowpostpb.Empty, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.confirmed = append(p.confirmed, in)
	return &knowpostpb.Empty{}, nil
}

func (p *recordingPublisher) Publish(
	_ context.Context,
	in *knowpostpb.PublishReq,
	_ ...grpc.CallOption,
) (*knowpostpb.KnowPostDetail, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.published = append(p.published, in)
	return &knowpostpb.KnowPostDetail{}, nil
}

func TestRunPublishStageRecordsEveryRPCStage(t *testing.T) {
	client := &recordingPublisher{nextID: 100}

	result, postIDs, err := runPublishStage(context.Background(), client, []int64{7, 8}, publishSpec{
		Requests:       6,
		Concurrency:    2,
		RequestTimeout: time.Second,
		RunID:          "run-42",
		Seed:           9,
	})

	require.NoError(t, err)
	require.Len(t, result, 5)
	for _, stage := range result {
		require.Equal(t, int64(6), stage.Total, stage.Name)
		require.Equal(t, int64(6), stage.Success, stage.Name)
	}
	require.Len(t, postIDs, 6)
	require.Len(t, client.confirmed, 6)
	require.Contains(t, client.confirmed[0].ObjectKey, "run-42")
	require.Len(t, client.published, 6)
}

func TestRunPublishStageDurationStopsNewWorkWithoutCancellingInflightPublish(t *testing.T) {
	client := &deadlineAwarePublisher{}

	result, postIDs, err := runPublishStage(context.Background(), client, []int64{7}, publishSpec{
		Duration:       10 * time.Millisecond,
		Concurrency:    1,
		RequestTimeout: 100 * time.Millisecond,
		RunID:          "run-42",
		Seed:           9,
	})

	require.NoError(t, err)
	require.Equal(t, int64(1), result[0].Total)
	require.Equal(t, int64(1), result[0].Success)
	require.Zero(t, result[0].Failed)
	require.Len(t, postIDs, 1)
}

func TestPublishAuthorsOncePublishesEverySpecifiedAuthor(t *testing.T) {
	client := &recordingPublisher{nextID: 200}

	postIDs, err := publishAuthorsOnce(
		context.Background(),
		client,
		[]int64{7, 8},
		"run-42-smoke",
		time.Second,
	)

	require.NoError(t, err)
	require.Len(t, postIDs, 2)
	require.Len(t, client.published, 2)
	require.Equal(t, int64(7), client.published[0].CreatorId)
	require.Equal(t, int64(8), client.published[1].CreatorId)
}

func TestSeedFeedPublishesConfiguredDepthForEveryAuthor(t *testing.T) {
	client := &recordingPublisher{nextID: 300}
	manifest := datasetManifest{NormalAuthors: []int64{7}, BigVAuthors: []int64{8}}

	postIDs, err := seedFeed(context.Background(), client, manifest, "hybrid", 2, time.Second)

	require.NoError(t, err)
	require.Len(t, postIDs, 4)
	require.Equal(t, []int64{7, 8, 7, 8}, []int64{
		client.published[0].CreatorId,
		client.published[1].CreatorId,
		client.published[2].CreatorId,
		client.published[3].CreatorId,
	})
}
