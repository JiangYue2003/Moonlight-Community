package feed

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type memoryTierStore struct {
	bigVs          map[int64]struct{}
	readErr        error
	promotionErr   error
	promotions     []tierPromotion
	batchReadCalls int
	singleCalls    int
}

type tierPromotion struct {
	authorID int64
	reason   string
	count    int64
}

func (s *memoryTierStore) IsBigV(_ context.Context, authorID int64) (bool, error) {
	s.singleCalls++
	if s.readErr != nil {
		return false, s.readErr
	}
	_, ok := s.bigVs[authorID]
	return ok, nil
}

func (s *memoryTierStore) BatchBigVs(_ context.Context, authorIDs []int64) (map[int64]struct{}, error) {
	s.batchReadCalls++
	if s.readErr != nil {
		return nil, s.readErr
	}
	result := make(map[int64]struct{})
	for _, authorID := range authorIDs {
		if _, ok := s.bigVs[authorID]; ok {
			result[authorID] = struct{}{}
		}
	}
	return result, nil
}

func (s *memoryTierStore) PromoteBigV(_ context.Context, authorID int64, reason string, count int64) error {
	if s.promotionErr != nil {
		return s.promotionErr
	}
	s.promotions = append(s.promotions, tierPromotion{authorID: authorID, reason: reason, count: count})
	if s.bigVs == nil {
		s.bigVs = make(map[int64]struct{})
	}
	s.bigVs[authorID] = struct{}{}
	return nil
}

type scriptedTierCounter struct {
	counts  map[int64]int64
	err     error
	calls   int
	userIDs []int64
}

func (c *scriptedTierCounter) BatchGetFollowerCounts(_ context.Context, userIDs []int64) (map[int64]int64, error) {
	c.calls++
	c.userIDs = append([]int64(nil), userIDs...)
	if c.err != nil {
		return nil, c.err
	}
	return c.counts, nil
}

type scriptedActiveFollowerCounter struct {
	counts  map[int64]int64
	err     error
	calls   int
	userIDs []int64
}

func (c *scriptedActiveFollowerCounter) BatchCountActive(_ context.Context, userIDs []int64) (map[int64]int64, error) {
	c.calls++
	c.userIDs = append([]int64(nil), userIDs...)
	if c.err != nil {
		return nil, c.err
	}
	return c.counts, nil
}

func TestAuthorTierResolverPersistedBigVBypassesCountDependencies(t *testing.T) {
	store := &memoryTierStore{bigVs: map[int64]struct{}{42: {}}}
	counter := &scriptedTierCounter{err: errors.New("must not be called")}
	followers := &scriptedActiveFollowerCounter{err: errors.New("must not be called")}
	resolver := NewAuthorTierResolver(store, counter, followers, nil)

	got, err := resolver.Resolve(context.Background(), 42)

	require.NoError(t, err)
	require.True(t, got.BigV)
	require.Equal(t, TierEvidencePersisted, got.Evidence)
	require.Zero(t, counter.calls)
	require.Zero(t, followers.calls)
	require.Empty(t, store.promotions)
}

func TestAuthorTierResolverPromotesOnlyAboveCounterThreshold(t *testing.T) {
	for _, test := range []struct {
		name      string
		count     int64
		wantBigV  bool
		wantWrite bool
	}{
		{name: "below", count: 999},
		{name: "equal", count: BIGV_THRESHOLD},
		{name: "above", count: BIGV_THRESHOLD + 1, wantBigV: true, wantWrite: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			store := &memoryTierStore{}
			counter := &scriptedTierCounter{counts: map[int64]int64{42: test.count}}
			followers := &scriptedActiveFollowerCounter{err: errors.New("must not be called")}
			resolver := NewAuthorTierResolver(store, counter, followers, nil)

			got, err := resolver.Resolve(context.Background(), 42)

			require.NoError(t, err)
			require.Equal(t, test.wantBigV, got.BigV)
			require.Equal(t, TierEvidenceCounter, got.Evidence)
			require.Equal(t, test.count, got.Followers)
			require.Equal(t, 1, counter.calls)
			require.Zero(t, followers.calls)
			if test.wantWrite {
				require.Equal(t, []tierPromotion{{
					authorID: 42,
					reason:   "counter_threshold",
					count:    BIGV_THRESHOLD + 1,
				}}, store.promotions)
			} else {
				require.Empty(t, store.promotions)
			}
		})
	}
}

func TestAuthorTierResolverFallsBackToMySQLForCounterFailureOrMissingResult(t *testing.T) {
	for _, test := range []struct {
		name       string
		counterErr error
		counts     map[int64]int64
		mysqlCount int64
		wantBigV   bool
	}{
		{name: "counter error below", counterErr: errors.New("counter down"), mysqlCount: 999},
		{name: "counter error equal", counterErr: errors.New("counter down"), mysqlCount: BIGV_THRESHOLD},
		{name: "counter error above", counterErr: errors.New("counter down"), mysqlCount: BIGV_THRESHOLD + 1, wantBigV: true},
		{name: "counter missing below", counts: map[int64]int64{}, mysqlCount: 999},
		{name: "counter missing equal", counts: map[int64]int64{}, mysqlCount: BIGV_THRESHOLD},
		{name: "counter missing above", counts: map[int64]int64{}, mysqlCount: BIGV_THRESHOLD + 1, wantBigV: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			store := &memoryTierStore{}
			counter := &scriptedTierCounter{counts: test.counts, err: test.counterErr}
			followers := &scriptedActiveFollowerCounter{counts: map[int64]int64{42: test.mysqlCount}}
			resolver := NewAuthorTierResolver(store, counter, followers, nil)

			got, err := resolver.Resolve(context.Background(), 42)

			require.NoError(t, err)
			require.Equal(t, test.wantBigV, got.BigV)
			require.Equal(t, TierEvidenceMySQL, got.Evidence)
			require.Equal(t, test.mysqlCount, got.Followers)
			require.True(t, got.Fallback)
			require.Equal(t, 1, counter.calls)
			require.Equal(t, 1, followers.calls)
			if test.wantBigV {
				require.Equal(t, []tierPromotion{{
					authorID: 42,
					reason:   "mysql_fallback_threshold",
					count:    BIGV_THRESHOLD + 1,
				}}, store.promotions)
			} else {
				require.Empty(t, store.promotions)
			}
		})
	}
}

func TestAuthorTierResolverReturnsRetryableErrorForUnresolvedTier(t *testing.T) {
	tests := []struct {
		name      string
		store     *memoryTierStore
		counter   *scriptedTierCounter
		followers *scriptedActiveFollowerCounter
	}{
		{
			name:      "tier read",
			store:     &memoryTierStore{readErr: errors.New("tier db down")},
			counter:   &scriptedTierCounter{},
			followers: &scriptedActiveFollowerCounter{},
		},
		{
			name:      "dual source failure",
			store:     &memoryTierStore{},
			counter:   &scriptedTierCounter{err: errors.New("counter down")},
			followers: &scriptedActiveFollowerCounter{err: errors.New("mysql down")},
		},
		{
			name:      "mysql result missing",
			store:     &memoryTierStore{},
			counter:   &scriptedTierCounter{counts: map[int64]int64{}},
			followers: &scriptedActiveFollowerCounter{counts: map[int64]int64{}},
		},
		{
			name:      "counter promotion",
			store:     &memoryTierStore{promotionErr: errors.New("tier write down")},
			counter:   &scriptedTierCounter{counts: map[int64]int64{42: BIGV_THRESHOLD + 1}},
			followers: &scriptedActiveFollowerCounter{},
		},
		{
			name:      "mysql promotion",
			store:     &memoryTierStore{promotionErr: errors.New("tier write down")},
			counter:   &scriptedTierCounter{err: errors.New("counter down")},
			followers: &scriptedActiveFollowerCounter{counts: map[int64]int64{42: BIGV_THRESHOLD + 1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resolver := NewAuthorTierResolver(test.store, test.counter, test.followers, nil)

			got, err := resolver.Resolve(context.Background(), 42)

			require.Equal(t, TierResolution{}, got)
			require.ErrorContains(t, err, "tier-resolution")
			var retryable interface{ Retryable() bool }
			require.ErrorAs(t, err, &retryable)
			require.True(t, retryable.Retryable())
			require.Empty(t, test.store.promotions)
		})
	}
}

func TestAuthorTierResolverBatchFallsBackOnlyForMissingCounterAuthors(t *testing.T) {
	store := &memoryTierStore{bigVs: map[int64]struct{}{11: {}}}
	counter := &scriptedTierCounter{counts: map[int64]int64{
		12: BIGV_THRESHOLD,
		13: BIGV_THRESHOLD + 1,
	}}
	followers := &scriptedActiveFollowerCounter{counts: map[int64]int64{14: BIGV_THRESHOLD + 2}}
	resolver := NewAuthorTierResolver(store, counter, followers, nil)

	got, err := resolver.ResolveBatch(context.Background(), []int64{11, 12, 13, 14})

	require.NoError(t, err)
	require.Equal(t, map[int64]TierResolution{
		11: {BigV: true, Evidence: TierEvidencePersisted},
		12: {Evidence: TierEvidenceCounter, Followers: BIGV_THRESHOLD},
		13: {BigV: true, Evidence: TierEvidenceCounter, Followers: BIGV_THRESHOLD + 1},
		14: {BigV: true, Evidence: TierEvidenceMySQL, Followers: BIGV_THRESHOLD + 2, Fallback: true},
	}, got)
	require.Equal(t, []int64{12, 13, 14}, counter.userIDs)
	require.Equal(t, []int64{14}, followers.userIDs)
	require.Equal(t, []tierPromotion{
		{authorID: 13, reason: "counter_threshold", count: BIGV_THRESHOLD + 1},
		{authorID: 14, reason: "mysql_fallback_threshold", count: BIGV_THRESHOLD + 2},
	}, store.promotions)
}
