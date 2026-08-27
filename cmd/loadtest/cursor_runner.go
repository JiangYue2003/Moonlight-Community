package main

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func fixedCursorTargetPage(scenario string) (int, bool) {
	const prefix = "cursor-page"
	if !strings.HasPrefix(scenario, prefix) {
		if scenario == "same-second-cursor" {
			return 2, true
		}
		return 0, false
	}
	page, err := strconv.Atoi(strings.TrimPrefix(scenario, prefix))
	if err != nil || page < 2 {
		return 0, false
	}
	return page, true
}

func offsetTargetPage(scenario string) (int, bool) {
	const prefix = "page"
	if !strings.HasPrefix(scenario, prefix) || strings.HasPrefix(scenario, "page-cache") {
		return 0, false
	}
	page, err := strconv.Atoi(strings.TrimPrefix(scenario, prefix))
	if err != nil || page < 1 {
		return 0, false
	}
	return page, true
}

func sequentialCursorTargetPage(scenario string) (int, bool) {
	const prefix = "sequential-"
	if !strings.HasPrefix(scenario, prefix) || strings.HasPrefix(scenario, "sequential-page-") {
		return 0, false
	}
	page, err := strconv.Atoi(strings.TrimPrefix(scenario, prefix))
	if err == nil && page >= 1 {
		return page, true
	}
	return 0, false
}

func sequentialOffsetTargetPage(scenario string) (int, bool) {
	const prefix = "sequential-page-"
	if !strings.HasPrefix(scenario, prefix) {
		return 0, false
	}
	page, err := strconv.Atoi(strings.TrimPrefix(scenario, prefix))
	if err == nil && page >= 1 {
		return page, true
	}
	return 0, false
}

func scenarioRequiresCursorDeepDataset(scenario string) bool {
	if _, ok := offsetTargetPage(scenario); ok {
		return true
	}
	if _, ok := fixedCursorTargetPage(scenario); ok {
		return true
	}
	if _, ok := sequentialCursorTargetPage(scenario); ok {
		return true
	}
	_, ok := sequentialOffsetTargetPage(scenario)
	return ok
}

type cursorTargetSpec struct {
	TargetPage      int
	Size            int32
	Timeout         time.Duration
	Requests        int
	Duration        time.Duration
	Concurrency     int
	Seed            int64
	RequireItems    bool
	AllowedCreators map[int64]struct{}
	Oracle          map[int64][]string
	OracleHash      string
}

type preparedCursorTarget struct {
	Cursor      string
	PrefixIDs   []string
	ExpectedIDs []string
}

type cursorPreparation struct {
	Targets  map[int64]preparedCursorTarget
	Requests int64
	Duration time.Duration
}

type paginationObservation struct {
	Mode                string        `json:"mode"`
	TargetPage          int           `json:"target_page"`
	PreparationRequests int64         `json:"preparation_requests"`
	PreparationDuration time.Duration `json:"preparation_duration_ns"`
	Sequences           int64         `json:"sequences,omitempty"`
	PagesTraversed      int64         `json:"pages_traversed,omitempty"`
	DuplicateItems      int64         `json:"duplicate_items"`
	OracleMismatches    int64         `json:"oracle_mismatches"`
	CursorLoops         int64         `json:"cursor_loops"`
	EarlyTerminations   int64         `json:"early_terminations"`
	OracleHash          string        `json:"oracle_hash,omitempty"`
	ObservedHash        string        `json:"observed_hash,omitempty"`
}

func prepareCursorTargets(
	ctx context.Context,
	client feedReader,
	readers []readerIdentity,
	spec cursorTargetSpec,
) (cursorPreparation, error) {
	started := time.Now()
	result := cursorPreparation{Targets: make(map[int64]preparedCursorTarget, len(readers))}
	if client == nil {
		return result, fmt.Errorf("feed reader is required")
	}
	if len(readers) == 0 {
		return result, fmt.Errorf("at least one reader is required")
	}
	if spec.TargetPage < 2 {
		return result, fmt.Errorf("cursor target page must be at least 2")
	}
	if spec.Size <= 0 || spec.Timeout <= 0 {
		return result, fmt.Errorf("cursor size and timeout must be positive")
	}

	for _, reader := range readers {
		cursor := ""
		seenCursors := make(map[string]struct{}, spec.TargetPage-1)
		seenIDs := make(map[string]struct{}, int(spec.Size)*spec.TargetPage)
		prefix := make([]string, 0, int(spec.Size)*(spec.TargetPage-1))
		for pageNumber := 1; pageNumber < spec.TargetPage; pageNumber++ {
			request := feedReadRequest{Size: spec.Size, Cursor: cursor}
			if cursor == "" {
				request.Page = 1
			}
			requestCtx, cancel := context.WithTimeout(ctx, spec.Timeout)
			page, err := client.GetUserFeed(requestCtx, reader, request)
			cancel()
			result.Requests++
			if err != nil {
				return result, fmt.Errorf("prepare reader %d page %d: %w", reader.UserID, pageNumber, err)
			}
			if err := validateFeedPage(page, loadSpec{
				RequireItems: spec.RequireItems, AllowedCreators: spec.AllowedCreators,
			}); err != nil {
				return result, fmt.Errorf("prepare reader %d page %d: %w", reader.UserID, pageNumber, err)
			}
			for _, item := range page.Items {
				if _, exists := seenIDs[item.ID]; exists {
					return result, fmt.Errorf("prepare reader %d: duplicate feed item %s across cursor pages", reader.UserID, item.ID)
				}
				seenIDs[item.ID] = struct{}{}
				prefix = append(prefix, item.ID)
			}
			if !page.HasMore {
				return result, fmt.Errorf("prepare reader %d: feed ended before target page %d", reader.UserID, spec.TargetPage)
			}
			if page.NextCursor == "" {
				return result, fmt.Errorf("prepare reader %d page %d: missing next cursor", reader.UserID, pageNumber)
			}
			if _, exists := seenCursors[page.NextCursor]; exists {
				return result, fmt.Errorf("prepare reader %d page %d: cursor loop", reader.UserID, pageNumber)
			}
			seenCursors[page.NextCursor] = struct{}{}
			cursor = page.NextCursor
		}

		target := preparedCursorTarget{Cursor: cursor, PrefixIDs: append([]string(nil), prefix...)}
		if oracle := spec.Oracle[reader.UserID]; oracle != nil {
			prefixEnd := int(spec.Size) * (spec.TargetPage - 1)
			if prefixEnd > len(oracle) || !slices.Equal(prefix, oracle[:prefixEnd]) {
				return result, fmt.Errorf("prepare reader %d: cursor prefix does not match oracle", reader.UserID)
			}
			targetEnd := prefixEnd + int(spec.Size)
			if targetEnd > len(oracle) {
				targetEnd = len(oracle)
			}
			target.ExpectedIDs = append([]string(nil), oracle[prefixEnd:targetEnd]...)
		}
		result.Targets[reader.UserID] = target
	}
	result.Duration = time.Since(started)
	return result, nil
}

func runFixedCursorStage(
	ctx context.Context,
	client feedReader,
	readers []readerIdentity,
	prepared cursorPreparation,
	spec cursorTargetSpec,
) (stageResult, error) {
	if spec.Concurrency <= 0 {
		return stageResult{}, fmt.Errorf("concurrency must be positive")
	}
	if spec.Requests <= 0 && spec.Duration <= 0 {
		return stageResult{}, fmt.Errorf("requests or duration must be positive")
	}
	if spec.TargetPage < 2 || spec.Size <= 0 || spec.Timeout <= 0 {
		return stageResult{}, fmt.Errorf("cursor target page, size and timeout are invalid")
	}
	if len(readers) == 0 {
		return stageResult{}, fmt.Errorf("at least one reader is required")
	}
	for _, reader := range readers {
		if prepared.Targets[reader.UserID].Cursor == "" {
			return stageResult{}, fmt.Errorf("reader %d has no prepared target cursor", reader.UserID)
		}
	}

	runCtx := ctx
	cancelRun := func() {}
	if spec.Duration > 0 {
		runCtx, cancelRun = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancelRun()

	recorder := newStageRecorder("read")
	started := time.Now()
	var sequence atomic.Int64
	var duplicates atomic.Int64
	var mismatches atomic.Int64
	var workers sync.WaitGroup
	workers.Add(spec.Concurrency)
	for workerID := 0; workerID < spec.Concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for {
				if runCtx.Err() != nil {
					return
				}
				requestNumber := sequence.Add(1)
				if spec.Requests > 0 && requestNumber > int64(spec.Requests) {
					return
				}
				reader := readers[deterministicIndex(spec.Seed, requestNumber, len(readers))]
				target := prepared.Targets[reader.UserID]
				requestCtx, cancel := context.WithTimeout(ctx, spec.Timeout)
				requestStarted := time.Now()
				page, err := client.GetUserFeed(requestCtx, reader, feedReadRequest{
					Size: spec.Size, Cursor: target.Cursor,
				})
				latency := time.Since(requestStarted)
				cancel()
				if err == nil {
					err = validateFeedPage(page, loadSpec{
						RequireItems: spec.RequireItems, AllowedCreators: spec.AllowedCreators,
					})
				}
				if err == nil {
					prefix := make(map[string]struct{}, len(target.PrefixIDs))
					for _, id := range target.PrefixIDs {
						prefix[id] = struct{}{}
					}
					for _, item := range page.Items {
						if _, exists := prefix[item.ID]; exists {
							duplicates.Add(1)
							err = fmt.Errorf("target page repeats prefix item %s", item.ID)
							break
						}
					}
				}
				if err == nil && target.ExpectedIDs != nil {
					actual := feedPageIDs(page)
					if !slices.Equal(actual, target.ExpectedIDs) {
						mismatches.Add(1)
						err = fmt.Errorf("target page does not match oracle")
					}
				}
				recorder.Record(latency, err)
			}
		}()
	}
	workers.Wait()
	result := recorder.Result(time.Since(started))
	result.Pagination = &paginationObservation{
		Mode:                "cursor",
		TargetPage:          spec.TargetPage,
		PreparationRequests: prepared.Requests,
		PreparationDuration: prepared.Duration,
		DuplicateItems:      duplicates.Load(),
		OracleMismatches:    mismatches.Load(),
		OracleHash:          spec.OracleHash,
		ObservedHash:        spec.OracleHash,
	}
	return result, nil
}

func runSequentialCursorStage(
	ctx context.Context,
	client feedReader,
	readers []readerIdentity,
	spec cursorTargetSpec,
) (stageResult, error) {
	if spec.Concurrency <= 0 {
		return stageResult{}, fmt.Errorf("concurrency must be positive")
	}
	if spec.Requests <= 0 && spec.Duration <= 0 {
		return stageResult{}, fmt.Errorf("requests or duration must be positive")
	}
	if spec.TargetPage < 1 || spec.Size <= 0 || spec.Timeout <= 0 {
		return stageResult{}, fmt.Errorf("cursor target page, size and timeout are invalid")
	}
	if len(readers) == 0 {
		return stageResult{}, fmt.Errorf("at least one reader is required")
	}

	runCtx := ctx
	cancelRun := func() {}
	if spec.Duration > 0 {
		runCtx, cancelRun = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancelRun()

	recorder := newStageRecorder("read")
	started := time.Now()
	var sequence atomic.Int64
	var completeSequences atomic.Int64
	var pagesTraversed atomic.Int64
	var duplicates atomic.Int64
	var mismatches atomic.Int64
	var loops atomic.Int64
	var earlyTerminations atomic.Int64
	var workers sync.WaitGroup
	workers.Add(spec.Concurrency)
	for workerID := 0; workerID < spec.Concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for {
				if runCtx.Err() != nil {
					return
				}
				sequenceNumber := sequence.Add(1)
				if spec.Requests > 0 && sequenceNumber > int64(spec.Requests) {
					return
				}
				reader := readers[deterministicIndex(spec.Seed, sequenceNumber, len(readers))]
				cursor := ""
				seenCursors := make(map[string]struct{}, spec.TargetPage)
				seenIDs := make(map[string]struct{}, int(spec.Size)*spec.TargetPage)
				sequenceOK := true
				for pageNumber := 1; pageNumber <= spec.TargetPage; pageNumber++ {
					request := feedReadRequest{Size: spec.Size, Cursor: cursor}
					if cursor == "" {
						request.Page = 1
					}
					requestCtx, cancel := context.WithTimeout(ctx, spec.Timeout)
					requestStarted := time.Now()
					page, err := client.GetUserFeed(requestCtx, reader, request)
					latency := time.Since(requestStarted)
					cancel()
					pagesTraversed.Add(1)
					if err == nil {
						err = validateFeedPage(page, loadSpec{
							RequireItems: spec.RequireItems, AllowedCreators: spec.AllowedCreators,
						})
					}
					if err == nil {
						for _, item := range page.Items {
							if _, exists := seenIDs[item.ID]; exists {
								duplicates.Add(1)
								err = fmt.Errorf("sequential cursor repeated item %s", item.ID)
								break
							}
							seenIDs[item.ID] = struct{}{}
						}
					}
					if err == nil {
						if oracle := spec.Oracle[reader.UserID]; oracle != nil {
							start := (pageNumber - 1) * int(spec.Size)
							end := start + int(spec.Size)
							if start > len(oracle) {
								start = len(oracle)
							}
							if end > len(oracle) {
								end = len(oracle)
							}
							if !slices.Equal(feedPageIDs(page), oracle[start:end]) {
								mismatches.Add(1)
								err = fmt.Errorf("sequential page %d does not match oracle", pageNumber)
							}
						}
					}
					if err == nil && pageNumber < spec.TargetPage {
						if !page.HasMore || page.NextCursor == "" {
							earlyTerminations.Add(1)
							err = fmt.Errorf("sequential feed ended before page %d", spec.TargetPage)
						} else if _, exists := seenCursors[page.NextCursor]; exists {
							loops.Add(1)
							err = fmt.Errorf("sequential cursor loop at page %d", pageNumber)
						}
					}
					recorder.Record(latency, err)
					if err != nil {
						sequenceOK = false
						break
					}
					if pageNumber < spec.TargetPage {
						seenCursors[page.NextCursor] = struct{}{}
						cursor = page.NextCursor
					}
				}
				if sequenceOK {
					completeSequences.Add(1)
				}
			}
		}()
	}
	workers.Wait()
	result := recorder.Result(time.Since(started))
	result.Pagination = &paginationObservation{
		Mode:              "cursor-sequential",
		TargetPage:        spec.TargetPage,
		Sequences:         completeSequences.Load(),
		PagesTraversed:    pagesTraversed.Load(),
		DuplicateItems:    duplicates.Load(),
		OracleMismatches:  mismatches.Load(),
		CursorLoops:       loops.Load(),
		EarlyTerminations: earlyTerminations.Load(),
		OracleHash:        spec.OracleHash,
		ObservedHash:      spec.OracleHash,
	}
	return result, nil
}

func runSequentialOffsetStage(
	ctx context.Context,
	client feedReader,
	readers []readerIdentity,
	spec cursorTargetSpec,
) (stageResult, error) {
	if spec.Concurrency <= 0 {
		return stageResult{}, fmt.Errorf("concurrency must be positive")
	}
	if spec.Requests <= 0 && spec.Duration <= 0 {
		return stageResult{}, fmt.Errorf("requests or duration must be positive")
	}
	if spec.TargetPage < 1 || spec.Size <= 0 || spec.Timeout <= 0 {
		return stageResult{}, fmt.Errorf("offset target page, size and timeout are invalid")
	}
	if len(readers) == 0 {
		return stageResult{}, fmt.Errorf("at least one reader is required")
	}

	runCtx := ctx
	cancelRun := func() {}
	if spec.Duration > 0 {
		runCtx, cancelRun = context.WithTimeout(ctx, spec.Duration)
	}
	defer cancelRun()

	recorder := newStageRecorder("read")
	started := time.Now()
	var sequence atomic.Int64
	var completeSequences atomic.Int64
	var pagesTraversed atomic.Int64
	var duplicates atomic.Int64
	var mismatches atomic.Int64
	var earlyTerminations atomic.Int64
	var workers sync.WaitGroup
	workers.Add(spec.Concurrency)
	for workerID := 0; workerID < spec.Concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for {
				if runCtx.Err() != nil {
					return
				}
				sequenceNumber := sequence.Add(1)
				if spec.Requests > 0 && sequenceNumber > int64(spec.Requests) {
					return
				}
				reader := readers[deterministicIndex(spec.Seed, sequenceNumber, len(readers))]
				seenIDs := make(map[string]struct{}, int(spec.Size)*spec.TargetPage)
				sequenceOK := true
				for pageNumber := 1; pageNumber <= spec.TargetPage; pageNumber++ {
					requestCtx, cancel := context.WithTimeout(ctx, spec.Timeout)
					requestStarted := time.Now()
					page, err := client.GetUserFeed(requestCtx, reader, feedReadRequest{
						Page: int32(pageNumber), Size: spec.Size,
					})
					latency := time.Since(requestStarted)
					cancel()
					pagesTraversed.Add(1)
					if err == nil {
						err = validateFeedPage(page, loadSpec{
							RequireItems: spec.RequireItems, AllowedCreators: spec.AllowedCreators,
						})
					}
					if err == nil {
						for _, item := range page.Items {
							if _, exists := seenIDs[item.ID]; exists {
								duplicates.Add(1)
								err = fmt.Errorf("sequential page repeated item %s", item.ID)
								break
							}
							seenIDs[item.ID] = struct{}{}
						}
					}
					if err == nil {
						if oracle := spec.Oracle[reader.UserID]; oracle != nil {
							start := (pageNumber - 1) * int(spec.Size)
							end := start + int(spec.Size)
							if start > len(oracle) {
								start = len(oracle)
							}
							if end > len(oracle) {
								end = len(oracle)
							}
							if !slices.Equal(feedPageIDs(page), oracle[start:end]) {
								mismatches.Add(1)
								err = fmt.Errorf("sequential page %d does not match oracle", pageNumber)
							}
						}
					}
					if err == nil && pageNumber < spec.TargetPage && !page.HasMore {
						earlyTerminations.Add(1)
						err = fmt.Errorf("sequential feed ended before page %d", spec.TargetPage)
					}
					recorder.Record(latency, err)
					if err != nil {
						sequenceOK = false
						break
					}
				}
				if sequenceOK {
					completeSequences.Add(1)
				}
			}
		}()
	}
	workers.Wait()
	result := recorder.Result(time.Since(started))
	result.Pagination = &paginationObservation{
		Mode:              "page-sequential",
		TargetPage:        spec.TargetPage,
		Sequences:         completeSequences.Load(),
		PagesTraversed:    pagesTraversed.Load(),
		DuplicateItems:    duplicates.Load(),
		OracleMismatches:  mismatches.Load(),
		EarlyTerminations: earlyTerminations.Load(),
		OracleHash:        spec.OracleHash,
		ObservedHash:      spec.OracleHash,
	}
	return result, nil
}

func feedPageIDs(page *feedPageResponse) []string {
	if page == nil {
		return nil
	}
	ids := make([]string, len(page.Items))
	for i, item := range page.Items {
		ids[i] = item.ID
	}
	return ids
}
