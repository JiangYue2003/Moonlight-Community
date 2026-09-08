package feed

import (
	"context"
	"fmt"
	"slices"
	"time"
)

const (
	routeSnapshotKeyVersion = "v2"
	routeEpochLogInterval   = 30 * time.Second
)

// RelationEpochReader exposes only the version read required by the Feed read
// path. feedepoch.Store satisfies this interface without coupling packages.
type RelationEpochReader interface {
	Relation(ctx context.Context, userID int64) (uint64, error)
}

// feedRouteSnapshot is immutable after construction. FeedReadSnapshot borrows
// its bigVAuthors and followingSet for the lifetime of a request. Keeping the
// type and collections private prevents callers from mutating cached state.
type feedRouteSnapshot struct {
	userID        int64
	relationEpoch uint64
	followings    []int64
	bigVAuthors   []int64
	loadedAt      time.Time
	followingSet  map[int64]struct{}
}

func (r *FeedReader) prepareVersionedRoute(ctx context.Context, userID int64) (*FeedReadSnapshot, error) {
	started := observationStarted(r.observer)
	epoch, err := r.relationEpochs.Relation(ctx, userID)
	if err != nil {
		if ctx.Err() != nil {
			recordStageSince(r.observer, StageRoute, OutcomeFromError(ctx.Err()), started)
			return nil, ctx.Err()
		}
		// Epoch is a cache-validity proof, not a request dependency. If it is
		// unavailable, bypass every unverifiable snapshot and use the source.
		r.logRouteEpochBypass(err)
		snapshot, _, sourceErr := r.loadRouteSnapshot(ctx, userID, 0)
		recordStageSince(r.observer, StageRoute, OutcomeFromError(sourceErr), started)
		if sourceErr != nil {
			return nil, sourceErr
		}
		return r.readSnapshot(snapshot), nil
	}

	key := routeSnapshotKey(userID, epoch)
	if snapshot, ok := r.cachedRouteSnapshot(key, userID, epoch); ok {
		recordStageSince(r.observer, StageRoute, OutcomeSuccess, started)
		return r.readSnapshot(snapshot), nil
	}

	resultC := r.routeGroup.DoChan(key, func() (any, error) {
		if snapshot, ok := r.cachedRouteSnapshot(key, userID, epoch); ok {
			return snapshot, nil
		}

		lookupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), feedRouteLookupTimeout)
		defer cancel()
		snapshot, cacheable, loadErr := r.loadRouteSnapshot(lookupCtx, userID, epoch)
		if loadErr != nil {
			return nil, loadErr
		}
		if cacheable {
			r.publishRouteSnapshot(key, snapshot)
		}
		return snapshot, nil
	})

	select {
	case <-ctx.Done():
		recordStageSince(r.observer, StageRoute, OutcomeFromError(ctx.Err()), started)
		return nil, ctx.Err()
	case result := <-resultC:
		recordStageSince(r.observer, StageRoute, OutcomeFromError(result.Err), started)
		if result.Err != nil {
			return nil, result.Err
		}
		snapshot, ok := result.Val.(*feedRouteSnapshot)
		if !ok || snapshot == nil {
			return nil, fmt.Errorf("route snapshot lookup returned %T", result.Val)
		}
		return r.readSnapshot(snapshot), nil
	}
}

type feedRouteCacheVisibility interface {
	Wait()
}

func (r *FeedReader) publishRouteSnapshot(key string, snapshot *feedRouteSnapshot) {
	accepted := r.routeCache.SetWithTTL(key, snapshot, routeSnapshotCost(snapshot), r.routeCacheTTL)
	visibility, canWait := r.routeCache.(feedRouteCacheVisibility)
	if !canWait {
		return
	}
	// RouteSnapshot owns a dedicated L1. Waiting here drains only its cold
	// writes and closes the gap between singleflight completion and Ristretto
	// visibility; hot hits never pay this cost.
	visibility.Wait()
	if !accepted {
		// A full Set buffer is transient. Waiting drains it, so give this cold
		// fill one bounded retry before allowing the cache admission policy to
		// reject the item normally.
		r.routeCache.SetWithTTL(key, snapshot, routeSnapshotCost(snapshot), r.routeCacheTTL)
		visibility.Wait()
	}
}

func (r *FeedReader) logRouteEpochBypass(err error) {
	now := time.Now().UnixNano()
	for {
		last := r.routeEpochLog.Load()
		if last != 0 && time.Duration(now-last) < routeEpochLogInterval {
			return
		}
		if r.routeEpochLog.CompareAndSwap(last, now) {
			r.logger.Errorf("bypass route snapshot after relation epoch failure: %v", err)
			return
		}
	}
}

func (r *FeedReader) loadRouteSnapshot(
	ctx context.Context,
	userID int64,
	epoch uint64,
) (*feedRouteSnapshot, bool, error) {
	relationStarted := observationStarted(r.observer)
	followings, err := r.relationClient.GetFollowings(ctx, userID)
	relationOutcome := OutcomeFromError(err)
	recordStageSince(r.observer, StageRelation, relationOutcome, relationStarted)
	RecordDependency(r.observer, DependencyRelation, OperationListFollowings, relationOutcome)
	if err != nil {
		r.logger.Errorf("route snapshot Relation lookup failed: %v", err)
		return nil, false, fmt.Errorf("get followings failed: %w", err)
	}

	// Own all collections before publishing the snapshot into the shared L1.
	followings = slices.Clone(followings)
	bigVs, _, cacheable, err := r.classifyFollowingsResult(ctx, followings)
	if err != nil {
		return nil, false, err
	}
	bigVs = slices.Clone(bigVs)
	followingSet := make(map[int64]struct{}, len(followings))
	for _, followingID := range followings {
		followingSet[followingID] = struct{}{}
	}
	return &feedRouteSnapshot{
		userID:        userID,
		relationEpoch: epoch,
		followings:    followings,
		bigVAuthors:   bigVs,
		loadedAt:      time.Now(),
		followingSet:  followingSet,
	}, cacheable, nil
}

func (r *FeedReader) cachedRouteSnapshot(key string, userID int64, epoch uint64) (*feedRouteSnapshot, bool) {
	value, ok := r.routeCache.Get(key)
	if !ok {
		return nil, false
	}
	snapshot, ok := value.(*feedRouteSnapshot)
	if !ok || snapshot == nil || snapshot.userID != userID || snapshot.relationEpoch != epoch || snapshot.followingSet == nil {
		return nil, false
	}
	return snapshot, true
}

func (r *FeedReader) readSnapshot(snapshot *feedRouteSnapshot) *FeedReadSnapshot {
	return &FeedReadSnapshot{
		reader:     r,
		userID:     snapshot.userID,
		bigVs:      snapshot.bigVAuthors,
		followings: snapshot.followingSet,
	}
}

func routeSnapshotKey(userID int64, epoch uint64) string {
	return fmt.Sprintf("feed:route:%s:%d:e%d", routeSnapshotKeyVersion, userID, epoch)
}

func routeSnapshotCost(snapshot *feedRouteSnapshot) int64 {
	if snapshot == nil {
		return 1
	}
	// Slices plus a conservative allowance for each map bucket/entry and key.
	return int64(128 + 32*len(snapshot.followings) + 8*len(snapshot.bigVAuthors))
}
