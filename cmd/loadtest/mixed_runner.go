package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type mixedSpec struct {
	Requests        int
	Duration        time.Duration
	Concurrency     int
	ReadPercent     int
	RequestTimeout  time.Duration
	Page            int32
	Size            int32
	RunID           string
	Seed            int64
	RequireItems    bool
	AllowedCreators map[int64]struct{}
}

func (spec mixedSpec) validate(authors []int64, readers []readerIdentity) error {
	if spec.Concurrency <= 0 {
		return fmt.Errorf("concurrency must be positive")
	}
	if spec.Requests <= 0 && spec.Duration <= 0 {
		return fmt.Errorf("requests or duration must be positive")
	}
	if spec.ReadPercent <= 0 || spec.ReadPercent >= 100 {
		return fmt.Errorf("read percent must be between 1 and 99")
	}
	if spec.RequestTimeout <= 0 {
		return fmt.Errorf("request timeout must be positive")
	}
	if spec.Page <= 0 || spec.Size <= 0 {
		return fmt.Errorf("page and size must be positive")
	}
	if spec.RunID == "" {
		return fmt.Errorf("run id is required")
	}
	if len(authors) == 0 {
		return fmt.Errorf("at least one author is required")
	}
	if len(readers) == 0 {
		return fmt.Errorf("at least one reader is required")
	}
	return nil
}

func runMixedStage(
	ctx context.Context,
	publisher publisherClient,
	reader feedReader,
	authors []int64,
	readers []readerIdentity,
	spec mixedSpec,
) ([]stageResult, []int64, error) {
	if err := spec.validate(authors, readers); err != nil {
		return nil, nil, err
	}

	runCtx := ctx
	cancel := func() {}
	if spec.Duration > 0 {
		runCtx, cancel = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancel()

	read := newStageRecorder("read")
	publishTotal := newStageRecorder("publish_total")
	draft := newStageRecorder("publish_draft")
	metadata := newStageRecorder("publish_metadata")
	confirm := newStageRecorder("publish_confirm")
	publish := newStageRecorder("publish_commit")
	publishRecorders := []*stageRecorder{publishTotal, draft, metadata, confirm, publish}

	started := time.Now()
	var sequence atomic.Int64
	var postIDsMu sync.Mutex
	postIDs := make([]int64, 0, maxInt(0, spec.Requests*(100-spec.ReadPercent)/100))
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
				requestCtx, requestCancel := context.WithTimeout(ctx, spec.RequestTimeout)
				if mixedRequestIsRead(requestNumber, spec.ReadPercent) {
					requestStarted := time.Now()
					identity := readers[deterministicIndex(spec.Seed, requestNumber, len(readers))]
					page, err := reader.GetUserFeed(requestCtx, identity, feedReadRequest{Page: spec.Page, Size: spec.Size})
					if err == nil {
						err = validateFeedPage(page, loadSpec{
							Page: spec.Page, Size: spec.Size, RequireItems: spec.RequireItems,
							AllowedCreators: spec.AllowedCreators,
						})
					}
					read.Record(time.Since(requestStarted), err)
					requestCancel()
					continue
				}

				requestStarted := time.Now()
				authorID := authors[deterministicIndex(spec.Seed, requestNumber, len(authors))]
				postID, err := publishOne(
					requestCtx, publisher, authorID, requestNumber, spec.RunID,
					draft, metadata, confirm, publish,
				)
				requestCancel()
				publishTotal.Record(time.Since(requestStarted), err)
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

	results := make([]stageResult, 0, 1+len(publishRecorders))
	results = append(results, read.Result(elapsed))
	for _, recorder := range publishRecorders {
		results = append(results, recorder.Result(elapsed))
	}
	sort.Slice(postIDs, func(i, j int) bool { return postIDs[i] < postIDs[j] })
	return results, postIDs, nil
}

func mixedRequestIsRead(requestNumber int64, readPercent int) bool {
	writePercent := int64(100 - readPercent)
	writesBefore := (requestNumber - 1) * writePercent / 100
	writesThroughCurrent := requestNumber * writePercent / 100
	return writesThroughCurrent == writesBefore
}
