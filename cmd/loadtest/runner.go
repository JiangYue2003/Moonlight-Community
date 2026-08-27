package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type feedReader interface {
	GetUserFeed(context.Context, readerIdentity, feedReadRequest) (*feedPageResponse, error)
}

type feedReadRequest struct {
	Page   int32
	Size   int32
	Cursor string
}

type loadSpec struct {
	Requests        int
	Duration        time.Duration
	Concurrency     int
	RequestTimeout  time.Duration
	Page            int32
	Size            int32
	Seed            int64
	RequireItems    bool
	AllowedCreators map[int64]struct{}
}

func (s loadSpec) validate(readers []readerIdentity) error {
	if s.Concurrency <= 0 {
		return fmt.Errorf("concurrency must be positive")
	}
	if s.Requests <= 0 && s.Duration <= 0 {
		return fmt.Errorf("requests or duration must be positive")
	}
	if len(readers) == 0 {
		return fmt.Errorf("at least one reader is required")
	}
	if s.RequestTimeout <= 0 {
		return fmt.Errorf("request timeout must be positive")
	}
	if s.Page <= 0 || s.Size <= 0 {
		return fmt.Errorf("page and size must be positive")
	}
	return nil
}

func runReadStage(
	ctx context.Context,
	client feedReader,
	readers []readerIdentity,
	spec loadSpec,
) (stageResult, error) {
	if err := spec.validate(readers); err != nil {
		return stageResult{}, err
	}

	runCtx := ctx
	cancel := func() {}
	if spec.Duration > 0 {
		runCtx, cancel = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancel()

	recorder := newStageRecorder("read")
	started := time.Now()
	var sequence atomic.Int64
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

				reader := readers[deterministicIndex(spec.Seed, requestNumber, len(readers))]
				requestCtx, requestCancel := context.WithTimeout(ctx, spec.RequestTimeout)
				requestStarted := time.Now()
				page, err := client.GetUserFeed(requestCtx, reader, feedReadRequest{Page: spec.Page, Size: spec.Size})
				latency := time.Since(requestStarted)
				requestCancel()
				if err == nil {
					err = validateFeedPage(page, spec)
				}
				recorder.Record(latency, err)
			}
		}()
	}
	workers.Wait()
	return recorder.Result(time.Since(started)), nil
}

// deterministicIndex makes a fixed request ordinal select the same target
// regardless of which worker happens to claim it.
func deterministicIndex(seed, ordinal int64, size int) int {
	if size <= 0 {
		return 0
	}
	x := uint64(seed) + uint64(ordinal)*0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	x ^= x >> 31
	return int(x % uint64(size))
}

func validateFeedPage(page *feedPageResponse, spec loadSpec) error {
	if page == nil {
		return fmt.Errorf("feed response is nil")
	}
	if spec.RequireItems && len(page.Items) == 0 {
		return fmt.Errorf("feed response has no items")
	}
	seen := make(map[string]struct{}, len(page.Items))
	for _, item := range page.Items {
		if item.ID == "" {
			return fmt.Errorf("feed item id is empty")
		}
		if len(spec.AllowedCreators) > 0 {
			if _, ok := spec.AllowedCreators[item.CreatorID]; !ok {
				return fmt.Errorf("feed item %s has unexpected creator %d", item.ID, item.CreatorID)
			}
		}
		if _, ok := seen[item.ID]; ok {
			return fmt.Errorf("duplicate feed item id %s", item.ID)
		}
		seen[item.ID] = struct{}{}
	}
	return nil
}
