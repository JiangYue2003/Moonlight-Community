package main

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type publisherClient interface {
	CreateDraft(context.Context, *knowpostpb.CreateDraftReq, ...grpc.CallOption) (*knowpostpb.CreateDraftResp, error)
	PatchMetadata(context.Context, *knowpostpb.PatchMetadataReq, ...grpc.CallOption) (*knowpostpb.KnowPostDetail, error)
	ConfirmContent(context.Context, *knowpostpb.ConfirmContentReq, ...grpc.CallOption) (*knowpostpb.Empty, error)
	Publish(context.Context, *knowpostpb.PublishReq, ...grpc.CallOption) (*knowpostpb.KnowPostDetail, error)
}

type publishSpec struct {
	Requests       int
	Duration       time.Duration
	Concurrency    int
	RequestTimeout time.Duration
	RunID          string
	Seed           int64
}

func (s publishSpec) validate(authors []int64) error {
	if s.Concurrency <= 0 {
		return fmt.Errorf("concurrency must be positive")
	}
	if s.Requests <= 0 && s.Duration <= 0 {
		return fmt.Errorf("requests or duration must be positive")
	}
	if s.RequestTimeout <= 0 {
		return fmt.Errorf("request timeout must be positive")
	}
	if len(authors) == 0 {
		return fmt.Errorf("at least one author is required")
	}
	if s.RunID == "" {
		return fmt.Errorf("run id is required")
	}
	return nil
}

func runPublishStage(
	ctx context.Context,
	client publisherClient,
	authors []int64,
	spec publishSpec,
) ([]stageResult, []int64, error) {
	if err := spec.validate(authors); err != nil {
		return nil, nil, err
	}

	runCtx := ctx
	cancel := func() {}
	if spec.Duration > 0 {
		runCtx, cancel = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancel()

	total := newStageRecorder("publish_total")
	draft := newStageRecorder("publish_draft")
	metadata := newStageRecorder("publish_metadata")
	confirm := newStageRecorder("publish_confirm")
	publish := newStageRecorder("publish_commit")
	recorders := []*stageRecorder{total, draft, metadata, confirm, publish}

	started := time.Now()
	var sequence atomic.Int64
	var postIDsMu sync.Mutex
	postIDs := make([]int64, 0, spec.Requests)
	var workers sync.WaitGroup
	workers.Add(spec.Concurrency)

	for workerID := 0; workerID < spec.Concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for {
				if err := runCtx.Err(); err != nil {
					return
				}
				requestNumber := sequence.Add(1)
				if spec.Requests > 0 && requestNumber > int64(spec.Requests) {
					return
				}

				authorID := authors[deterministicIndex(spec.Seed, requestNumber, len(authors))]
				requestCtx, requestCancel := context.WithTimeout(ctx, spec.RequestTimeout)
				operationStarted := time.Now()
				postID, err := publishOne(
					requestCtx,
					client,
					authorID,
					requestNumber,
					spec.RunID,
					draft,
					metadata,
					confirm,
					publish,
				)
				requestCancel()
				total.Record(time.Since(operationStarted), err)
				if err == nil {
					postIDsMu.Lock()
					postIDs = append(postIDs, postID)
					postIDsMu.Unlock()
				}
			}
		}()
	}
	workers.Wait()
	elapsed := time.Since(started)

	results := make([]stageResult, 0, len(recorders))
	for _, recorder := range recorders {
		results = append(results, recorder.Result(elapsed))
	}
	sort.Slice(postIDs, func(i, j int) bool { return postIDs[i] < postIDs[j] })
	return results, postIDs, nil
}

func publishOne(
	ctx context.Context,
	client publisherClient,
	authorID, requestNumber int64,
	runID string,
	draft, metadata, confirm, publish *stageRecorder,
) (int64, error) {
	started := time.Now()
	draftResp, err := client.CreateDraft(ctx, &knowpostpb.CreateDraftReq{CreatorId: authorID})
	draft.Record(time.Since(started), err)
	if err != nil {
		return 0, fmt.Errorf("create draft: %w", err)
	}
	postID, err := strconv.ParseInt(draftResp.Id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse draft id %q: %w", draftResp.Id, err)
	}

	started = time.Now()
	_, err = client.PatchMetadata(ctx, &knowpostpb.PatchMetadataReq{
		Id:        postID,
		CreatorId: authorID,
		Title:     fmt.Sprintf("feed-loadtest %s #%d", runID, requestNumber),
		TitleSet:  true,
	})
	metadata.Record(time.Since(started), err)
	if err != nil {
		return 0, fmt.Errorf("patch metadata: %w", err)
	}

	started = time.Now()
	_, err = client.ConfirmContent(ctx, &knowpostpb.ConfirmContentReq{
		Id:        postID,
		CreatorId: authorID,
		ObjectKey: fmt.Sprintf("loadtest/%s/%d/content.md", runID, postID),
		Size:      1,
	})
	confirm.Record(time.Since(started), err)
	if err != nil {
		return 0, fmt.Errorf("confirm content: %w", err)
	}

	started = time.Now()
	_, err = client.Publish(ctx, &knowpostpb.PublishReq{Id: postID, CreatorId: authorID})
	publish.Record(time.Since(started), err)
	if err != nil {
		return 0, fmt.Errorf("publish: %w", err)
	}
	return postID, nil
}

func publishAuthorsOnce(
	ctx context.Context,
	client publisherClient,
	authors []int64,
	runID string,
	requestTimeout time.Duration,
) ([]int64, error) {
	if len(authors) == 0 {
		return nil, fmt.Errorf("at least one author is required")
	}
	if requestTimeout <= 0 {
		return nil, fmt.Errorf("request timeout must be positive")
	}
	draft := newStageRecorder("smoke_draft")
	metadata := newStageRecorder("smoke_metadata")
	confirm := newStageRecorder("smoke_confirm")
	publish := newStageRecorder("smoke_publish")
	postIDs := make([]int64, 0, len(authors))
	for i, authorID := range authors {
		requestCtx, cancel := context.WithTimeout(ctx, requestTimeout)
		postID, err := publishOne(
			requestCtx,
			client,
			authorID,
			int64(i+1),
			runID,
			draft,
			metadata,
			confirm,
			publish,
		)
		cancel()
		if err != nil {
			return nil, fmt.Errorf("publish smoke author %d: %w", authorID, err)
		}
		postIDs = append(postIDs, postID)
	}
	return postIDs, nil
}

func seedFeed(
	ctx context.Context,
	client publisherClient,
	manifest datasetManifest,
	strategy string,
	postsPerAuthor int,
	requestTimeout time.Duration,
) ([]int64, error) {
	if postsPerAuthor <= 0 {
		return nil, fmt.Errorf("seed posts per author must be positive")
	}
	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	postIDs := make([]int64, 0, len(authors)*postsPerAuthor)
	for round := 1; round <= postsPerAuthor; round++ {
		ids, err := publishAuthorsOnce(
			ctx,
			client,
			authors,
			fmt.Sprintf("%s-seed-%s-%d", manifest.RunID, strategy, round),
			requestTimeout,
		)
		if err != nil {
			return nil, fmt.Errorf("seed round %d: %w", round, err)
		}
		postIDs = append(postIDs, ids...)
	}
	return postIDs, nil
}
