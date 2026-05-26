package relationlogic

import (
	"context"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	gore "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/cache"
	gzredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	model "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

func newFollowFixture(t *testing.T) (*svc.ServiceContext, sqlmock.Sqlmock, *miniredis.Miniredis) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)
	rdb := gore.NewClient(&gore.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	conn := sqlx.NewSqlConnFromDB(db)
	cc := cache.CacheConf{cache.NodeConf{
		RedisConf: gzredis.RedisConf{Host: mr.Addr(), Type: "node"},
		Weight:    100,
	}}
	sc := &svc.ServiceContext{
		Config: config.Config{RateLimit: config.RateLimitConf{
			FollowCapacity: 100, FollowRefillPerSec: 1,
		}},
		Db:             conn,
		FollowingModel: model.NewFollowingModel(conn, cc),
		OutboxModel:    model.NewOutboxModel(conn, cc),
		Redis:          rdb,
		RateLimiter:    ratelimit.New(rdb),
		Snowflake:      snowflakex.MustNew(1, 3),
	}
	return sc, mock, mr
}

func TestFollow_RejectsSelfFollow(t *testing.T) {
	sc, _, _ := newFollowFixture(t)
	_, err := NewFollowLogic(context.Background(), sc).Follow(&relation.FollowReq{FromUserId: 1, ToUserId: 1})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeBadRequest {
		t.Fatalf("want BadRequest, got %v", err)
	}
}

func TestFollow_RejectsZeroId(t *testing.T) {
	sc, _, _ := newFollowFixture(t)
	for _, in := range []*relation.FollowReq{
		{FromUserId: 0, ToUserId: 1},
		{FromUserId: 1, ToUserId: 0},
	} {
		if _, err := NewFollowLogic(context.Background(), sc).Follow(in); err == nil {
			t.Fatalf("zero id must be rejected: %+v", in)
		}
	}
}

func TestFollow_RateLimitedAfterCapacity(t *testing.T) {
	sc, _, _ := newFollowFixture(t)
	sc.Config.RateLimit.FollowCapacity = 1
	_, _ = NewFollowLogic(context.Background(), sc).Follow(&relation.FollowReq{FromUserId: 5, ToUserId: 6})
	_, err := NewFollowLogic(context.Background(), sc).Follow(&relation.FollowReq{FromUserId: 5, ToUserId: 7})
	be, _ := errorx.As(err)
	if be == nil || be.Code != errorx.CodeRateLimited {
		t.Fatalf("second call should be RateLimited; got %v", err)
	}
}

func TestFollow_ConcurrentRateLimit(t *testing.T) {
	sc, _, _ := newFollowFixture(t)
	sc.Config.RateLimit.FollowCapacity = 5
	var rateLimited int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := NewFollowLogic(context.Background(), sc).Follow(&relation.FollowReq{FromUserId: 9, ToUserId: 10})
			if be, ok := errorx.As(err); ok && be.Code == errorx.CodeRateLimited {
				mu.Lock()
				rateLimited++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if rateLimited < 10 {
		t.Fatalf("expected ≥10 rate-limited; got %d / 20", rateLimited)
	}
}

func TestFollow_DuplicateFollowReturnsChangedFalse(t *testing.T) {
	sc, mock, _ := newFollowFixture(t)
	mock.ExpectBegin()
	mock.ExpectQuery("select 1 from `following`").
		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectCommit()

	resp, err := NewFollowLogic(context.Background(), sc).Follow(&relation.FollowReq{FromUserId: 11, ToUserId: 12})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Changed {
		t.Fatalf("duplicate follow should not report changed=true")
	}
}
