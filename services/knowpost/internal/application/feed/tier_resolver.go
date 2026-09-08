package feed

import (
	"context"
	"errors"
	"fmt"
)

type TierEvidence string

const (
	TierEvidencePersisted TierEvidence = "persisted"
	TierEvidenceCounter   TierEvidence = "counter"
	TierEvidenceMySQL     TierEvidence = "mysql"
	TierEvidenceUnknown   TierEvidence = "unknown"

	promotionReasonCounterThreshold = "counter_threshold"
	promotionReasonMySQLFallback    = "mysql_fallback_threshold"
)

type TierResolution struct {
	BigV      bool
	Evidence  TierEvidence
	Followers int64
	Fallback  bool
}

type TierResolutionError struct {
	stage string
	cause error
}

func (e *TierResolutionError) Error() string {
	return fmt.Sprintf("tier-resolution %s failed: %v", e.stage, e.cause)
}

func (e *TierResolutionError) Unwrap() error {
	return e.cause
}

func (e *TierResolutionError) Retryable() bool {
	return true
}

func IsTierResolutionError(err error) bool {
	var target *TierResolutionError
	return errors.As(err, &target)
}

func newTierResolutionError(stage string, cause error) error {
	return &TierResolutionError{stage: stage, cause: cause}
}

type AuthorTierStore interface {
	IsBigV(ctx context.Context, authorID int64) (bool, error)
	BatchBigVs(ctx context.Context, authorIDs []int64) (map[int64]struct{}, error)
	PromoteBigV(ctx context.Context, authorID int64, reason string, observedFollowers int64) error
}

type ActiveFollowerCounter interface {
	BatchCountActive(ctx context.Context, authorIDs []int64) (map[int64]int64, error)
}

type AuthorTierResolver struct {
	store     AuthorTierStore
	counter   CounterClient
	followers ActiveFollowerCounter
	observer  FeedObserver
}

func NewAuthorTierResolver(
	store AuthorTierStore,
	counter CounterClient,
	followers ActiveFollowerCounter,
	observer FeedObserver,
) *AuthorTierResolver {
	return &AuthorTierResolver{
		store:     store,
		counter:   counter,
		followers: followers,
		observer:  observer,
	}
}

func (r *AuthorTierResolver) Resolve(ctx context.Context, authorID int64) (TierResolution, error) {
	bigV, err := r.store.IsBigV(ctx, authorID)
	RecordDependency(r.observer, DependencyMySQL, OperationTierRead, OutcomeFromError(err))
	if err != nil {
		return TierResolution{}, newTierResolutionError("tier-read", err)
	}
	if bigV {
		return TierResolution{BigV: true, Evidence: TierEvidencePersisted}, nil
	}

	counts, counterErr := r.counter.BatchGetFollowerCounts(ctx, []int64{authorID})
	RecordDependency(r.observer, DependencyCounter, OperationBatchFollowerCounts, OutcomeFromError(counterErr))
	if counterErr == nil {
		if followers, ok := counts[authorID]; ok {
			resolution := TierResolution{Evidence: TierEvidenceCounter, Followers: followers}
			if followers <= BIGV_THRESHOLD {
				return resolution, nil
			}
			if err := r.promoteBigV(ctx, authorID, promotionReasonCounterThreshold, followers); err != nil {
				return TierResolution{}, newTierResolutionError("tier-promote", err)
			}
			resolution.BigV = true
			return resolution, nil
		}
	}

	mysqlCounts, err := r.followers.BatchCountActive(ctx, []int64{authorID})
	RecordDependency(r.observer, DependencyMySQL, OperationActiveFollowerCounts, OutcomeFromError(err))
	if err != nil {
		if counterErr != nil {
			err = errors.Join(counterErr, err)
		}
		return TierResolution{}, newTierResolutionError("evidence", err)
	}
	followers, ok := mysqlCounts[authorID]
	if !ok {
		return TierResolution{}, newTierResolutionError(
			"mysql-evidence",
			fmt.Errorf("active follower count missing for author %d", authorID),
		)
	}
	resolution := TierResolution{
		Evidence:  TierEvidenceMySQL,
		Followers: followers,
		Fallback:  true,
	}
	if followers <= BIGV_THRESHOLD {
		return resolution, nil
	}
	if err := r.promoteBigV(ctx, authorID, promotionReasonMySQLFallback, followers); err != nil {
		return TierResolution{}, newTierResolutionError("tier-promote", err)
	}
	resolution.BigV = true
	return resolution, nil
}

func (r *AuthorTierResolver) ResolveBatch(
	ctx context.Context,
	authorIDs []int64,
) (map[int64]TierResolution, error) {
	result := make(map[int64]TierResolution, len(authorIDs))
	if len(authorIDs) == 0 {
		return result, nil
	}

	uniqueIDs := uniqueInt64s(authorIDs)
	persisted, err := r.store.BatchBigVs(ctx, uniqueIDs)
	RecordDependency(r.observer, DependencyMySQL, OperationTierRead, OutcomeFromError(err))
	if err != nil {
		return nil, newTierResolutionError("tier-batch-read", err)
	}
	unresolved := make([]int64, 0, len(uniqueIDs))
	for _, authorID := range uniqueIDs {
		if _, ok := persisted[authorID]; ok {
			result[authorID] = TierResolution{BigV: true, Evidence: TierEvidencePersisted}
			continue
		}
		unresolved = append(unresolved, authorID)
	}
	if len(unresolved) == 0 {
		return result, nil
	}

	counterCounts, counterErr := r.counter.BatchGetFollowerCounts(ctx, unresolved)
	RecordDependency(r.observer, DependencyCounter, OperationBatchFollowerCounts, OutcomeFromError(counterErr))
	fallbackIDs := make([]int64, 0, len(unresolved))
	if counterErr != nil {
		fallbackIDs = append(fallbackIDs, unresolved...)
	} else {
		for _, authorID := range unresolved {
			followers, ok := counterCounts[authorID]
			if !ok {
				fallbackIDs = append(fallbackIDs, authorID)
				continue
			}
			resolution := TierResolution{Evidence: TierEvidenceCounter, Followers: followers}
			if followers > BIGV_THRESHOLD {
				if err := r.promoteBigV(ctx, authorID, promotionReasonCounterThreshold, followers); err != nil {
					return nil, newTierResolutionError("tier-promote", err)
				}
				resolution.BigV = true
			}
			result[authorID] = resolution
		}
	}

	if len(fallbackIDs) == 0 {
		return result, nil
	}
	mysqlCounts, err := r.followers.BatchCountActive(ctx, fallbackIDs)
	RecordDependency(r.observer, DependencyMySQL, OperationActiveFollowerCounts, OutcomeFromError(err))
	if err != nil {
		if counterErr != nil {
			err = errors.Join(counterErr, err)
		}
		return nil, newTierResolutionError("evidence", err)
	}
	for _, authorID := range fallbackIDs {
		followers, ok := mysqlCounts[authorID]
		if !ok {
			return nil, newTierResolutionError(
				"mysql-evidence",
				fmt.Errorf("active follower count missing for author %d", authorID),
			)
		}
		resolution := TierResolution{
			Evidence:  TierEvidenceMySQL,
			Followers: followers,
			Fallback:  true,
		}
		if followers > BIGV_THRESHOLD {
			if err := r.promoteBigV(ctx, authorID, promotionReasonMySQLFallback, followers); err != nil {
				return nil, newTierResolutionError("tier-promote", err)
			}
			resolution.BigV = true
		}
		result[authorID] = resolution
	}
	return result, nil
}

func (r *AuthorTierResolver) promoteBigV(
	ctx context.Context,
	authorID int64,
	reason string,
	followers int64,
) error {
	err := r.store.PromoteBigV(ctx, authorID, reason, followers)
	RecordDependency(r.observer, DependencyMySQL, OperationTierPromote, OutcomeFromError(err))
	return err
}

func uniqueInt64s(values []int64) []int64 {
	seen := make(map[int64]struct{}, len(values))
	unique := make([]int64, 0, len(values))
	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		unique = append(unique, value)
	}
	return unique
}
