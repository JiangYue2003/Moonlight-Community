package knowpostlogic

import (
	"context"
	"database/sql"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/detail"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

type stubKnowPostsModel struct {
	row *model.KnowPosts
}

func (s *stubKnowPostsModel) Insert(context.Context, *model.KnowPosts) (sql.Result, error) {
	panic("not implemented")
}
func (s *stubKnowPostsModel) FindOne(context.Context, uint64) (*model.KnowPosts, error) {
	if s.row == nil {
		return nil, model.ErrNotFound
	}
	return s.row, nil
}
func (s *stubKnowPostsModel) Update(context.Context, *model.KnowPosts) error { panic("not implemented") }
func (s *stubKnowPostsModel) Delete(context.Context, uint64) error            { panic("not implemented") }
func (s *stubKnowPostsModel) ListPublicFeed(context.Context, int, int) ([]*model.KnowPosts, error) {
	panic("not implemented")
}
func (s *stubKnowPostsModel) ListMyFeed(context.Context, uint64, int, int) ([]*model.KnowPosts, error) {
	panic("not implemented")
}
func (s *stubKnowPostsModel) UpdateInTx(context.Context, sqlx.Session, *model.KnowPosts) error {
	panic("not implemented")
}

func newDetailFixture(t *testing.T, row *model.KnowPosts) (*GetDetailLogic, *svc.ServiceContext, *miniredis.Miniredis) {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	if err != nil {
		t.Fatal(err)
	}
	l2 := cachex.NewL2(rdb)
	hotCfg := hotkey.Config{WindowSeconds: 60, SegmentSeconds: 10, LevelLow: 1, LevelMedium: 5, LevelHigh: 10}
	hot := hotkey.New(hotCfg)
	detailCache := detail.New(l1, l2, hot)
	sc := &svc.ServiceContext{
		KnowPostsModel: &stubKnowPostsModel{row: row},
		Redis:          rdb,
		DetailCache:    detailCache,
		L1FeedItem:     l1,
		HotDetail:      hot,
	}
	return NewGetDetailLogic(context.Background(), sc), sc, mr
}

func newKnowpostRow(visible string) *model.KnowPosts {
	now := time.Now()
	return &model.KnowPosts{
		Id:          9,
		CreatorId:   7,
		Visible:     visible,
		Status:      "published",
		Type:        "image_text",
		Title:       sql.NullString{String: "t", Valid: true},
		Description: sql.NullString{String: "d", Valid: true},
		CreateTime:  now,
		UpdateTime:  now,
		PublishTime: sql.NullTime{Time: now, Valid: true},
	}
}

func TestGetDetail_DeniesNonPublicForAnonymous(t *testing.T) {
	for _, visible := range []string{"followers", "school", "unlisted", "private"} {
		logic, _, _ := newDetailFixture(t, newKnowpostRow(visible))
		if _, err := logic.GetDetail(&knowpost.GetDetailReq{Id: 9}); err == nil {
			t.Fatalf("visible=%s should be denied for anonymous", visible)
		}
	}
}

func TestGetDetail_AllowsOwnerForNonPublic(t *testing.T) {
	for _, visible := range []string{"followers", "school", "unlisted", "private"} {
		logic, _, _ := newDetailFixture(t, newKnowpostRow(visible))
		got, err := logic.GetDetail(&knowpost.GetDetailReq{Id: 9, ViewerId: 7})
		if err != nil || got == nil || got.Id != "9" {
			t.Fatalf("visible=%s owner should access got=%+v err=%v", visible, got, err)
		}
	}
}

func TestGetDetail_HotHitExtendsFeedItemTTL(t *testing.T) {
	logic, sc, mr := newDetailFixture(t, newKnowpostRow("public"))
	ctx := context.Background()
	itemKey := cachekeys.FeedItemKey(9)
	if err := sc.Redis.Set(ctx, itemKey, "x", 5*time.Second).Err(); err != nil {
		t.Fatal(err)
	}
	if _, err := logic.GetDetail(&knowpost.GetDetailReq{Id: 9}); err != nil {
		t.Fatal(err)
	}
	if _, err := logic.GetDetail(&knowpost.GetDetailReq{Id: 9}); err != nil {
		t.Fatal(err)
	}
	if ttl := mr.TTL(itemKey); ttl <= 5*time.Second {
		t.Fatalf("feed:item ttl should be extended, got %v", ttl)
	}
}
