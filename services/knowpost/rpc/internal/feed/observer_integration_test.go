package feed

import (
	"context"
	"errors"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestFeedReaderRecordsStagesAndDependencyCalls(t *testing.T) {
	ctx := context.Background()
	const userID int64 = 123
	relation := NewMockRelationClient()
	relation.SetFollowings(userID, []int64{201})
	counter := NewMockCounterClient()
	counter.SetFollowerCount(201, BIGV_THRESHOLD+1)
	redis := &batchTrackingRedis{
		MockRedisClient: NewMockRedisClient(),
		results: []ZRevRangeResult{
			{Scores: []ZScore{}},
			{Scores: []ZScore{}},
		},
	}
	observer := &recordingFeedObserver{}
	reader := NewFeedReaderWithOptions(
		redis,
		relation,
		counter,
		logx.WithContext(ctx),
		FeedReaderOptions{
			Strategy:                  StrategyHybrid,
			CombinedPipelineEnabled:   true,
			CombinedPipelineBatchSize: maxBigVOutboxPipelineSize,
			Observer:                  observer,
		},
	)

	snapshot, err := reader.Prepare(ctx, userID)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if _, _, err := snapshot.GetFeed(ctx, 1, 20); err != nil {
		t.Fatalf("GetFeed: %v", err)
	}

	for _, stage := range []FeedStage{
		StageRelation,
		StageCounter,
		StageRoute,
		StageInbox,
		StageBigVPipeline,
		StageMergeDedup,
	} {
		if !hasStageCall(observer.calls, stage, OutcomeSuccess) {
			t.Errorf("missing successful stage observation %q: %+v", stage, observer.calls)
		}
	}
	for _, expected := range []struct {
		dependency FeedDependency
		operation  FeedOperation
	}{
		{DependencyRelation, OperationListFollowings},
		{DependencyCounter, OperationBatchFollowerCounts},
		{DependencyRedis, OperationInboxBigVPipeline},
	} {
		if !hasDependencyCall(observer.calls, expected.dependency, expected.operation, OutcomeSuccess) {
			t.Errorf("missing dependency observation %q/%q: %+v",
				expected.dependency, expected.operation, observer.calls)
		}
	}
	for _, operation := range []FeedOperation{OperationInboxRead, OperationBigVPipeline} {
		if hasDependencyCall(observer.calls, DependencyRedis, operation, OutcomeSuccess) {
			t.Errorf("combined first pipeline must not record a second Redis round trip %q: %+v",
				operation, observer.calls)
		}
	}
}

func TestFeedReaderRecordsRelationFailure(t *testing.T) {
	ctx := context.Background()
	observer := &recordingFeedObserver{}
	reader := NewFeedReaderWithOptions(
		NewMockRedisClient(),
		failingRelationClient{err: errors.New("relation unavailable")},
		NewMockCounterClient(),
		logx.WithContext(ctx),
		FeedReaderOptions{Strategy: StrategyHybrid, Observer: observer},
	)

	if _, err := reader.Prepare(ctx, 123); err == nil {
		t.Fatal("Prepare error = nil, want relation failure")
	}
	if !hasStageCall(observer.calls, StageRelation, OutcomeError) {
		t.Fatalf("missing relation error stage: %+v", observer.calls)
	}
	if !hasDependencyCall(observer.calls, DependencyRelation, OperationListFollowings, OutcomeError) {
		t.Fatalf("missing relation error dependency call: %+v", observer.calls)
	}
}

type failingRelationClient struct {
	err error
}

func (c failingRelationClient) GetFollowings(context.Context, int64) ([]int64, error) {
	return nil, c.err
}

func (c failingRelationClient) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, c.err
}

func hasStageCall(calls []observerCall, stage FeedStage, outcome FeedOutcome) bool {
	for _, call := range calls {
		if call.kind == "stage" && call.stage == stage && call.outcome == outcome {
			return true
		}
	}
	return false
}

func hasDependencyCall(
	calls []observerCall,
	dependency FeedDependency,
	operation FeedOperation,
	outcome FeedOutcome,
) bool {
	for _, call := range calls {
		if call.kind == "dependency" && call.dependency == dependency &&
			call.operation == operation && call.outcome == outcome {
			return true
		}
	}
	return false
}
