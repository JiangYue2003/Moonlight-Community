package knowpostlogic

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type oneFollowingRelation struct{}

func (oneFollowingRelation) GetFollowers(context.Context, int64) ([]int64, error) {
	return nil, nil
}

func (oneFollowingRelation) GetFollowings(context.Context, int64) ([]int64, error) {
	return []int64{7}, nil
}

type zeroFollowerCounter struct{}

func (zeroFollowerCounter) GetFollowerCount(context.Context, int64) (int64, error) {
	return 0, nil
}

func TestGetUserFeedNormalizesPagination(t *testing.T) {
	redisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf("start miniredis: %v", err)
	}
	t.Cleanup(redisServer.Close)

	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() {
		_ = redisClient.Close()
	})

	ctx := context.Background()
	serviceCtx := &svc.ServiceContext{
		FeedReader: feed.NewFeedReader(
			feed.NewRedisAdapter(redisClient),
			oneFollowingRelation{},
			zeroFollowerCounter{},
			logx.WithContext(ctx),
		),
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Errorf("GetUserFeed panicked for invalid pagination: %v", recovered)
		}
	}()

	resp, err := NewGetUserFeedLogic(ctx, serviceCtx).GetUserFeed(&knowpost.GetUserFeedReq{
		UserId: 42,
		Page:   0,
		Size:   100,
	})
	if err != nil {
		t.Fatalf("GetUserFeed failed: %v", err)
	}
	if resp.Page != 1 {
		t.Fatalf("page = %d, want 1", resp.Page)
	}
	if resp.Size != 50 {
		t.Fatalf("size = %d, want 50", resp.Size)
	}
}
