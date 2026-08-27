package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	relationmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

const backfillPromotionReason = "mysql_backfill_threshold"

type activeFollowerCountPager interface {
	PageActiveCounts(ctx context.Context, afterAuthorID int64, limit int) ([]relationmodel.ActiveFollowerCount, error)
}

type backfillTierStore interface {
	BatchBigVs(ctx context.Context, authorIDs []int64) (map[int64]struct{}, error)
	PromoteBigV(ctx context.Context, authorID int64, reason string, observedFollowers int64) error
}

type tierBackfillOptions struct {
	AfterAuthorID int64
	BatchSize     int
	MaxPages      int
	Apply         bool
	PageInterval  time.Duration
}

type tierBackfillReport struct {
	Scanned          int
	Eligible         int
	Promoted         int
	AlreadyPromoted  int
	DryRunCandidates int
	Failed           int
	NextAfterID      int64
	Complete         bool
}

func runTierBackfill(
	ctx context.Context,
	source activeFollowerCountPager,
	tiers backfillTierStore,
	options tierBackfillOptions,
) (tierBackfillReport, error) {
	if options.BatchSize <= 0 || options.BatchSize > 10_000 {
		return tierBackfillReport{}, fmt.Errorf("batch size %d must be between 1 and 10000", options.BatchSize)
	}
	if options.MaxPages <= 0 {
		return tierBackfillReport{}, fmt.Errorf("max pages %d must be positive", options.MaxPages)
	}
	if options.PageInterval < 0 {
		return tierBackfillReport{}, fmt.Errorf("page interval %s must not be negative", options.PageInterval)
	}

	report := tierBackfillReport{NextAfterID: options.AfterAuthorID}
	for page := 0; page < options.MaxPages; page++ {
		rows, err := source.PageActiveCounts(ctx, report.NextAfterID, options.BatchSize)
		if err != nil {
			return report, fmt.Errorf("page active follower counts after %d: %w", report.NextAfterID, err)
		}
		if len(rows) == 0 {
			report.Complete = true
			break
		}

		report.Scanned += len(rows)
		pageLastAuthorID := rows[len(rows)-1].AuthorID
		eligibleIDs := make([]int64, 0, len(rows))
		counts := make(map[int64]int64, len(rows))
		for _, row := range rows {
			if row.Count <= feed.BIGV_THRESHOLD {
				continue
			}
			report.Eligible++
			eligibleIDs = append(eligibleIDs, row.AuthorID)
			counts[row.AuthorID] = row.Count
		}

		persisted, err := tiers.BatchBigVs(ctx, eligibleIDs)
		if err != nil {
			return report, fmt.Errorf("batch read persisted tiers: %w", err)
		}
		var pagePromotionErrors []error
		for _, authorID := range eligibleIDs {
			if _, ok := persisted[authorID]; ok {
				report.AlreadyPromoted++
				continue
			}
			if !options.Apply {
				report.DryRunCandidates++
				continue
			}
			if err := tiers.PromoteBigV(ctx, authorID, backfillPromotionReason, counts[authorID]); err != nil {
				report.Failed++
				pagePromotionErrors = append(pagePromotionErrors, fmt.Errorf("promote author %d: %w", authorID, err))
				continue
			}
			report.Promoted++
		}
		if len(pagePromotionErrors) > 0 {
			return report, errors.Join(pagePromotionErrors...)
		}
		report.NextAfterID = pageLastAuthorID

		if len(rows) < options.BatchSize {
			report.Complete = true
			break
		}
		if options.PageInterval > 0 && page+1 < options.MaxPages {
			timer := time.NewTimer(options.PageInterval)
			select {
			case <-ctx.Done():
				if !timer.Stop() {
					<-timer.C
				}
				return report, ctx.Err()
			case <-timer.C:
			}
		}
	}
	return report, nil
}
