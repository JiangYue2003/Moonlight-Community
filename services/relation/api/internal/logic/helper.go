package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

func summariesFromPb(items []*pb.UserSummary) []types.UserSummary {
	out := make([]types.UserSummary, 0, len(items))
	for _, u := range items {
		out = append(out, types.UserSummary{
			Id: u.Id, Nickname: u.Nickname, Avatar: u.Avatar, ZgId: u.ZgId, Bio: u.Bio,
		})
	}
	return out
}

func hydrateProfiles(ctx context.Context, uc userpb.UserClient, base []types.UserSummary) ([]types.UserSummary, error) {
	if len(base) == 0 {
		return base, nil
	}
	ids := make([]int64, 0, len(base))
	for _, b := range base {
		ids = append(ids, b.Id)
	}

	resp, err := uc.FindByIds(ctx, &userpb.FindByIdsReq{Ids: ids})
	if err != nil {
		return base, nil
	}

	idx := make(map[int64]*userpb.UserInfo, len(resp.Users))
	for _, u := range resp.Users {
		idx[u.Id] = u
	}
	for i := range base {
		u, ok := idx[base[i].Id]
		if !ok {
			continue
		}
		base[i].Gender = u.Gender
		base[i].Birthday = u.Birthday
		base[i].School = u.School
		base[i].Phone = u.Phone
		base[i].Email = u.Email
		base[i].TagJson = u.TagsJson
	}
	return base, nil
}
