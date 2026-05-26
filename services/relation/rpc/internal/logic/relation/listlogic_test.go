package relationlogic

import (
	"context"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	model "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
)

type stubUserClient struct{}

func (stubUserClient) GetById(context.Context, *userpb.GetByIdReq, ...grpc.CallOption) (*userpb.GetByIdResp, error) {
	return nil, nil
}
func (stubUserClient) GetByIdentifier(context.Context, *userpb.GetByIdentifierReq, ...grpc.CallOption) (*userpb.GetByIdentifierResp, error) {
	return nil, nil
}
func (stubUserClient) FindByIds(_ context.Context, in *userpb.FindByIdsReq, _ ...grpc.CallOption) (*userpb.FindByIdsResp, error) {
	out := make([]*userpb.UserInfo, 0, len(in.Ids))
	for _, id := range in.Ids {
		out = append(out, &userpb.UserInfo{Id: id, Nickname: "u"})
	}
	return &userpb.FindByIdsResp{Users: out}, nil
}
func (stubUserClient) Create(context.Context, *userpb.CreateReq, ...grpc.CallOption) (*userpb.CreateResp, error) {
	return nil, nil
}
func (stubUserClient) ExistsByIdentifier(context.Context, *userpb.ExistsByIdentifierReq, ...grpc.CallOption) (*userpb.ExistsByIdentifierResp, error) {
	return nil, nil
}
func (stubUserClient) UpdatePassword(context.Context, *userpb.UpdatePasswordReq, ...grpc.CallOption) (*userpb.UpdatePasswordResp, error) {
	return nil, nil
}
func (stubUserClient) ExistsByZgIdExceptId(context.Context, *userpb.ExistsByZgIdExceptIdReq, ...grpc.CallOption) (*userpb.ExistsByZgIdExceptIdResp, error) {
	return nil, nil
}
func (stubUserClient) UpdateProfile(context.Context, *userpb.UpdateProfileReq, ...grpc.CallOption) (*userpb.UpdateProfileResp, error) {
	return nil, nil
}

type panicFollowingModel struct{ model.FollowingModel }
type panicFollowerModel struct{ model.FollowerModel }

func (panicFollowingModel) PageActive(context.Context, int64, int, int, int64) ([]*model.Following, error) {
	panic("db should not be hit when top cache covers request")
}
func (panicFollowerModel) PageActive(context.Context, int64, int, int, int64) ([]*model.Follower, error) {
	panic("db should not be hit when top cache covers request")
}

func TestListFollowing_UsesTopCacheBeforeRedisAndDB(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	sc := &svc.ServiceContext{
		Redis:          rdb,
		UserRpc:        stubUserClient{},
		FollowingModel: panicFollowingModel{},
	}
	sc.FollowingTopCache = map[int64][]int64{
		7: {101, 102, 103},
	}

	resp, err := NewListFollowingLogic(context.Background(), sc).ListFollowing(&relation.ListReq{
		UserId: 7,
		Offset: 1,
		Limit:  2,
	})
	if err != nil {
		t.Fatalf("list following: %v", err)
	}
	if len(resp.Items) != 2 || resp.Items[0].GetId() != 102 || resp.Items[1].GetId() != 103 {
		t.Fatalf("unexpected top cache items: %+v", resp.Items)
	}
}

func TestListFollowers_UsesTopCacheBeforeRedisAndDB(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	sc := &svc.ServiceContext{
		Redis:         rdb,
		UserRpc:       stubUserClient{},
		FollowerModel: panicFollowerModel{},
	}
	sc.FollowerTopCache = map[int64][]int64{
		8: {201, 202, 203},
	}

	resp, err := NewListFollowersLogic(context.Background(), sc).ListFollowers(&relation.ListReq{
		UserId: 8,
		Offset: 0,
		Limit:  2,
	})
	if err != nil {
		t.Fatalf("list followers: %v", err)
	}
	if len(resp.Items) != 2 || resp.Items[0].GetId() != 201 || resp.Items[1].GetId() != 202 {
		t.Fatalf("unexpected top cache items: %+v", resp.Items)
	}
}
