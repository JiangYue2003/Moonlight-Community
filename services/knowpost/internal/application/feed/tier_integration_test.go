package feed

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestFeedReaderHybridEnforceDoesNotCacheUnresolvedTierRoute(t *testing.T) {
	relation := NewMockRelationClient()
	relation.SetFollowings(7, []int64{42})
	resolver := NewAuthorTierResolver(
		&memoryTierStore{},
		&scriptedTierCounter{err: errors.New("counter down")},
		&scriptedActiveFollowerCounter{err: errors.New("mysql down")},
		nil,
	)
	epoch := &mutableRelationEpoch{value: 3}
	routeCache := newLockedRouteCache()
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		relation,
		NewMockCounterClient(),
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:             StrategyHybrid,
			TierMode:             AuthorTierModeEnforce,
			TierResolver:         resolver,
			RouteCache:           routeCache,
			RouteCacheTTL:        time.Second,
			RouteSnapshotEnabled: true,
			RelationEpochs:       epoch,
		},
	)

	_, err := reader.Prepare(context.Background(), 7)

	require.Error(t, err)
	require.True(t, IsTierResolutionError(err))
	_, cached := routeCache.Get(routeSnapshotKey(7, 3))
	require.False(t, cached)
}

func TestFeedWriterAndReaderEnforceSharePersistedBigVClassification(t *testing.T) {
	store := &memoryTierStore{bigVs: map[int64]struct{}{42: {}}}
	counter := &scriptedTierCounter{err: errors.New("must not be called")}
	followers := &scriptedActiveFollowerCounter{err: errors.New("must not be called")}
	resolver := NewAuthorTierResolver(store, counter, followers, nil)

	redis := NewMockRedisClient()
	writer := NewFeedWriterWithOptions(redis, NewMockKafkaProducer(), logx.WithContext(context.Background()), FeedWriterOptions{
		Strategy:     StrategyHybrid,
		TierMode:     AuthorTierModeEnforce,
		TierResolver: resolver,
	})
	require.NoError(t, writer.OnPostPublished(context.Background(), 99, 42, 0))
	require.Len(t, redis.zaddCalls, 2)
	require.Equal(t, "feed:bigv:42", redis.zaddCalls[1].Key)

	reader := NewFeedReaderWithOptions(
		redis,
		NewMockRelationClient(),
		NewMockCounterClient(),
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:     StrategyHybrid,
			TierMode:     AuthorTierModeEnforce,
			TierResolver: resolver,
		},
	)
	bigVs, normal, cacheable, err := reader.classifyFollowingsResult(context.Background(), []int64{42})
	require.NoError(t, err)
	require.Equal(t, []int64{42}, bigVs)
	require.Empty(t, normal)
	require.True(t, cacheable)
	require.Zero(t, counter.calls)
	require.Zero(t, followers.calls)
}

func TestFeedReaderShadowRecordsMismatchWithoutChangingLegacyRoute(t *testing.T) {
	observer := &recordingFeedObserver{}
	store := &memoryTierStore{bigVs: map[int64]struct{}{42: {}}}
	resolver := NewAuthorTierResolver(
		store,
		&scriptedTierCounter{err: errors.New("must not be called")},
		&scriptedActiveFollowerCounter{err: errors.New("must not be called")},
		observer,
	)
	legacyCounter := NewMockCounterClient()
	legacyCounter.SetFollowerCount(42, 0)
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		NewMockRelationClient(),
		legacyCounter,
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:     StrategyHybrid,
			TierMode:     AuthorTierModeShadow,
			TierResolver: resolver,
			Observer:     observer,
		},
	)

	bigVs, normal, cacheable, err := reader.classifyFollowingsResult(context.Background(), []int64{42})

	require.NoError(t, err)
	require.True(t, cacheable)
	require.Empty(t, bigVs)
	require.Equal(t, []int64{42}, normal)
	require.Equal(t, 1, store.batchReadCalls)
	require.Contains(t, observer.calls, observerCall{
		kind: "tier", evidence: TierEvidencePersisted, outcome: OutcomeSuccess, mismatch: true,
	})
}

func TestFeedReaderShadowStillObservesTierWhenLegacyCounterFails(t *testing.T) {
	observer := &recordingFeedObserver{}
	store := &memoryTierStore{bigVs: map[int64]struct{}{42: {}}}
	resolver := NewAuthorTierResolver(
		store,
		&scriptedTierCounter{err: errors.New("must not be called")},
		&scriptedActiveFollowerCounter{err: errors.New("must not be called")},
		observer,
	)
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		NewMockRelationClient(),
		&scriptedTierCounter{err: errors.New("legacy counter down")},
		logx.WithContext(context.Background()),
		FeedReaderOptions{
			Strategy:     StrategyHybrid,
			TierMode:     AuthorTierModeShadow,
			TierResolver: resolver,
			Observer:     observer,
		},
	)

	bigVs, normal, cacheable, err := reader.classifyFollowingsResult(context.Background(), []int64{42})

	require.NoError(t, err)
	require.False(t, cacheable)
	require.Empty(t, bigVs)
	require.Equal(t, []int64{42}, normal)
	require.Contains(t, observer.calls, observerCall{
		kind: "tier", evidence: TierEvidencePersisted, outcome: OutcomeSuccess, mismatch: true,
	})
}

func TestForcedWriterStrategiesBypassTierEnforce(t *testing.T) {
	for _, test := range []struct {
		name          string
		strategy      Strategy
		wantKafka     int
		wantBigVWrite bool
	}{
		{name: "push", strategy: StrategyPush, wantKafka: 1},
		{name: "pull", strategy: StrategyPull, wantBigVWrite: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			store := &memoryTierStore{readErr: errors.New("must not be called")}
			resolver := NewAuthorTierResolver(
				store,
				&scriptedTierCounter{err: errors.New("must not be called")},
				&scriptedActiveFollowerCounter{err: errors.New("must not be called")},
				nil,
			)
			redis := NewMockRedisClient()
			kafka := NewMockKafkaProducer()
			writer := NewFeedWriterWithOptions(redis, kafka, logx.WithContext(context.Background()), FeedWriterOptions{
				Strategy:     test.strategy,
				TierMode:     AuthorTierModeEnforce,
				TierResolver: resolver,
			})

			require.NoError(t, writer.OnPostPublished(context.Background(), 99, 42, 0))
			require.Equal(t, test.wantKafka, len(kafka.messages))
			require.Equal(t, test.wantBigVWrite, len(redis.zaddCalls) == 2)
			require.Zero(t, store.singleCalls)
		})
	}
}

func TestForcedReaderStrategiesBypassTierEnforce(t *testing.T) {
	for _, test := range []struct {
		name       string
		strategy   Strategy
		wantBigVs  []int64
		wantNormal []int64
	}{
		{name: "push", strategy: StrategyPush, wantBigVs: []int64{}, wantNormal: []int64{42}},
		{name: "pull", strategy: StrategyPull, wantBigVs: []int64{42}, wantNormal: []int64{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			store := &memoryTierStore{readErr: errors.New("must not be called")}
			resolver := NewAuthorTierResolver(
				store,
				&scriptedTierCounter{err: errors.New("must not be called")},
				&scriptedActiveFollowerCounter{err: errors.New("must not be called")},
				nil,
			)
			reader := NewFeedReaderWithOptions(
				NewMockRedisClient(),
				NewMockRelationClient(),
				NewMockCounterClient(),
				logx.WithContext(context.Background()),
				FeedReaderOptions{
					Strategy:     test.strategy,
					TierMode:     AuthorTierModeEnforce,
					TierResolver: resolver,
				},
			)

			bigVs, normal, cacheable, err := reader.classifyFollowingsResult(context.Background(), []int64{42})

			require.NoError(t, err)
			require.True(t, cacheable)
			require.Equal(t, test.wantBigVs, bigVs)
			require.Equal(t, test.wantNormal, normal)
			require.Zero(t, store.batchReadCalls)
		})
	}
}
