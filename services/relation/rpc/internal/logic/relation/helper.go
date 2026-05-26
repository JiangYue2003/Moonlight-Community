package relationlogic

import (
	"context"
	"strconv"

	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

// hydrateUsers 调 user-rpc.FindByIds 把 id 列表转为 UserSummary 列表，保持原顺序。
func hydrateUsers(ctx context.Context, sc *svc.ServiceContext, ids []int64) ([]*pb.UserSummary, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	resp, err := sc.UserRpc.FindByIds(ctx, &userpb.FindByIdsReq{Ids: ids})
	if err != nil {
		return nil, err
	}
	idx := make(map[int64]*userpb.UserInfo, len(resp.Users))
	for _, u := range resp.Users {
		idx[u.Id] = u
	}
	out := make([]*pb.UserSummary, 0, len(ids))
	for _, id := range ids {
		u, ok := idx[id]
		if !ok {
			continue
		}
		out = append(out, &pb.UserSummary{
			Id: u.Id, Nickname: u.Nickname, Avatar: u.Avatar, ZgId: u.ZgId, Bio: u.Bio,
		})
	}
	return out, nil
}

func clampLimit(n int32) int {
	if n <= 0 {
		return 20
	}
	if n > 100 {
		return 100
	}
	return int(n)
}

// 防 unused 报错（debug 时偶尔用 strconv）。
var _ = strconv.Itoa
