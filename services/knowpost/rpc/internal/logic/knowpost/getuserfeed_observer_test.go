package knowpostlogic

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

type logicObserverCall struct {
	kind       string
	stage      feed.FeedStage
	dependency feed.FeedDependency
	operation  feed.FeedOperation
	cache      feed.FeedPageCacheSource
	outcome    feed.FeedOutcome
	mode       feed.FeedPaginationMode
	pageResult feed.FeedPaginationResult
	work       feed.FeedCursorWork
	value      int
}

type logicRecordingObserver struct {
	calls []logicObserverCall
}

func (o *logicRecordingObserver) ObserveStage(
	stage feed.FeedStage,
	outcome feed.FeedOutcome,
	_ time.Duration,
) {
	o.calls = append(o.calls, logicObserverCall{kind: "stage", stage: stage, outcome: outcome})
}

func (o *logicRecordingObserver) ObserveDependency(
	dependency feed.FeedDependency,
	operation feed.FeedOperation,
	outcome feed.FeedOutcome,
) {
	o.calls = append(o.calls, logicObserverCall{
		kind:       "dependency",
		dependency: dependency,
		operation:  operation,
		outcome:    outcome,
	})
}

func (o *logicRecordingObserver) ObserveColdCompute(outcome feed.FeedOutcome) {
	o.calls = append(o.calls, logicObserverCall{kind: "cold", outcome: outcome})
}

func (o *logicRecordingObserver) ObservePageCache(source feed.FeedPageCacheSource, outcome feed.FeedOutcome) {
	o.calls = append(o.calls, logicObserverCall{kind: "page_cache", cache: source, outcome: outcome})
}

func (o *logicRecordingObserver) ObservePageRefreshState(int, int, int) {}

func (o *logicRecordingObserver) ObservePagination(
	mode feed.FeedPaginationMode,
	result feed.FeedPaginationResult,
	_ time.Duration,
) {
	o.calls = append(o.calls, logicObserverCall{kind: "pagination", mode: mode, pageResult: result})
}

func (o *logicRecordingObserver) ObserveCursorWork(work feed.FeedCursorWork, value int) {
	o.calls = append(o.calls, logicObserverCall{kind: "cursor_work", work: work, value: value})
}

func TestGetUserFeedRecordsHydrationTotalAndColdCompute(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })

	ctx := context.Background()
	const userID int64 = 42
	if err := redisClient.ZAdd(ctx, fmt.Sprintf(feed.FEED_INBOX_KEY, userID), redis.Z{
		Score:  1000,
		Member: 1,
	}).Err(); err != nil {
		t.Fatalf("seed inbox: %v", err)
	}
	observer := &logicRecordingObserver{}
	loader := &recordingFeedPostLoader{rows: map[uint64]*model.KnowPosts{
		1: {
			Id:          1,
			CreatorId:   7,
			Visible:     "public",
			Status:      "published",
			PublishTime: sql.NullTime{Time: time.Now(), Valid: true},
		},
	}}
	reader := feed.NewFeedReaderWithOptions(
		feed.NewRedisAdapter(redisClient),
		oneFollowingRelation{},
		zeroFollowerCounter{},
		logx.WithContext(ctx),
		feed.FeedReaderOptions{Strategy: feed.StrategyHybrid, Observer: observer},
	)
	serviceCtx := &svc.ServiceContext{
		Redis:          redisClient,
		FeedPostLoader: loader,
		FeedReader:     reader,
		FeedObserver:   observer,
	}

	resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: userID,
		Page:   1,
		Size:   20,
	})
	if err != nil {
		t.Fatalf("GetUserFeed: %v", err)
	}
	if len(resp.Items) != 1 {
		t.Fatalf("items = %d, want 1", len(resp.Items))
	}

	for _, stage := range []feed.FeedStage{feed.StageHydrate, feed.StageTotal} {
		if !logicHasStage(observer.calls, stage, feed.OutcomeSuccess) {
			t.Errorf("missing successful stage %q: %+v", stage, observer.calls)
		}
	}
	if !logicHasCold(observer.calls, feed.OutcomeSuccess) {
		t.Errorf("missing successful cold compute: %+v", observer.calls)
	}
	if !logicHasPageCache(observer.calls, feed.PageCacheBypass, feed.OutcomeSuccess) {
		t.Errorf("missing successful page cache bypass while cache is disabled: %+v", observer.calls)
	}
	if got := sumCursorWork(observer.calls, feed.CursorWorkHydrateIDs); got != 1 {
		t.Errorf("page hydrate work = %d, want 1: %+v", got, observer.calls)
	}
	for _, expected := range []struct {
		dependency feed.FeedDependency
		operation  feed.FeedOperation
	}{
		{feed.DependencyRedis, feed.OperationFeedItemMGet},
		{feed.DependencyMySQL, feed.OperationFeedItemDB},
	} {
		if !logicHasDependency(observer.calls, expected.dependency, expected.operation, feed.OutcomeSuccess) {
			t.Errorf("missing dependency %q/%q: %+v", expected.dependency, expected.operation, observer.calls)
		}
	}
}

func logicHasStage(calls []logicObserverCall, stage feed.FeedStage, outcome feed.FeedOutcome) bool {
	for _, call := range calls {
		if call.kind == "stage" && call.stage == stage && call.outcome == outcome {
			return true
		}
	}
	return false
}

func logicHasDependency(
	calls []logicObserverCall,
	dependency feed.FeedDependency,
	operation feed.FeedOperation,
	outcome feed.FeedOutcome,
) bool {
	for _, call := range calls {
		if call.kind == "dependency" && call.dependency == dependency &&
			call.operation == operation && call.outcome == outcome {
			return true
		}
	}
	return false
}

func logicHasCold(calls []logicObserverCall, outcome feed.FeedOutcome) bool {
	for _, call := range calls {
		if call.kind == "cold" && call.outcome == outcome {
			return true
		}
	}
	return false
}

func logicHasPageCache(
	calls []logicObserverCall,
	source feed.FeedPageCacheSource,
	outcome feed.FeedOutcome,
) bool {
	for _, call := range calls {
		if call.kind == "page_cache" && call.cache == source && call.outcome == outcome {
			return true
		}
	}
	return false
}
