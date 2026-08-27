package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	knowpostmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	relationmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

type memoryFollowerCountPager struct {
	rows []relationmodel.ActiveFollowerCount
}

func (p *memoryFollowerCountPager) PageActiveCounts(
	_ context.Context,
	afterAuthorID int64,
	limit int,
) ([]relationmodel.ActiveFollowerCount, error) {
	page := make([]relationmodel.ActiveFollowerCount, 0, limit)
	for _, row := range p.rows {
		if row.AuthorID <= afterAuthorID {
			continue
		}
		page = append(page, row)
		if len(page) == limit {
			break
		}
	}
	return page, nil
}

type memoryBackfillTierStore struct {
	bigVs      map[int64]struct{}
	promotions []backfillPromotion
	promoteErr error
}

type backfillPromotion struct {
	authorID int64
	reason   string
	count    int64
}

func (s *memoryBackfillTierStore) BatchBigVs(
	_ context.Context,
	authorIDs []int64,
) (map[int64]struct{}, error) {
	result := make(map[int64]struct{})
	for _, authorID := range authorIDs {
		if _, ok := s.bigVs[authorID]; ok {
			result[authorID] = struct{}{}
		}
	}
	return result, nil
}

func (s *memoryBackfillTierStore) PromoteBigV(
	_ context.Context,
	authorID int64,
	reason string,
	count int64,
) error {
	if s.promoteErr != nil {
		return s.promoteErr
	}
	s.promotions = append(s.promotions, backfillPromotion{authorID: authorID, reason: reason, count: count})
	if s.bigVs == nil {
		s.bigVs = make(map[int64]struct{})
	}
	s.bigVs[authorID] = struct{}{}
	return nil
}

func TestRunTierBackfillDryRunDoesNotPersist(t *testing.T) {
	source := &memoryFollowerCountPager{rows: []relationmodel.ActiveFollowerCount{
		{AuthorID: 7, Count: 1001},
	}}
	tiers := &memoryBackfillTierStore{}

	report, err := runTierBackfill(context.Background(), source, tiers, tierBackfillOptions{
		BatchSize: 10,
		MaxPages:  1,
	})

	require.NoError(t, err)
	require.Equal(t, 1, report.Scanned)
	require.Equal(t, 1, report.Eligible)
	require.Equal(t, 1, report.DryRunCandidates)
	require.True(t, report.Complete)
	require.Empty(t, tiers.promotions)
}

func TestRunTierBackfillReportsPromotionFailures(t *testing.T) {
	source := &memoryFollowerCountPager{rows: []relationmodel.ActiveFollowerCount{
		{AuthorID: 7, Count: 1001},
	}}
	tiers := &memoryBackfillTierStore{promoteErr: errors.New("tier db down")}

	report, err := runTierBackfill(context.Background(), source, tiers, tierBackfillOptions{
		BatchSize: 10,
		MaxPages:  1,
		Apply:     true,
	})

	require.ErrorContains(t, err, "promote author 7")
	require.Equal(t, 1, report.Failed)
	require.Zero(t, report.Promoted)
	require.Zero(t, report.NextAfterID)
	require.False(t, report.Complete)
}

var _ knowpostmodel.FeedAuthorDeliveryTierModel = (*memoryBackfillTierStore)(nil)

func (s *memoryBackfillTierStore) IsBigV(_ context.Context, authorID int64) (bool, error) {
	_, ok := s.bigVs[authorID]
	return ok, nil
}

func TestRunTierBackfillIsBoundedAndPromotesOnlyNewAuthorsAboveThreshold(t *testing.T) {
	source := &memoryFollowerCountPager{rows: []relationmodel.ActiveFollowerCount{
		{AuthorID: 1, Count: 1000},
		{AuthorID: 2, Count: 1001},
		{AuthorID: 3, Count: 5000},
		{AuthorID: 4, Count: 20},
		{AuthorID: 5, Count: 9000},
	}}
	tiers := &memoryBackfillTierStore{bigVs: map[int64]struct{}{3: {}}}

	report, err := runTierBackfill(context.Background(), source, tiers, tierBackfillOptions{
		BatchSize: 2,
		MaxPages:  2,
		Apply:     true,
	})

	require.NoError(t, err)
	require.Equal(t, tierBackfillReport{
		Scanned:         4,
		Eligible:        2,
		Promoted:        1,
		AlreadyPromoted: 1,
		NextAfterID:     4,
	}, report)
	require.Equal(t, []backfillPromotion{{
		authorID: 2,
		reason:   "mysql_backfill_threshold",
		count:    1001,
	}}, tiers.promotions)
}
