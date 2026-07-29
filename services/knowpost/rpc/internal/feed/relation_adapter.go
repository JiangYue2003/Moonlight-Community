package feed

import (
	"context"

	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

const relationPageSize int32 = 100

// RelationClientAdapter 实现 RelationClient 接口，适配 Relation RPC
type RelationClientAdapter struct {
	client relationpb.RelationClient
}

func NewRelationClientAdapter(client relationpb.RelationClient) *RelationClientAdapter {
	return &RelationClientAdapter{client: client}
}

func (r *RelationClientAdapter) GetFollowers(ctx context.Context, userID int64) ([]int64, error) {
	return listRelationUserIDs(ctx, userID, func(ctx context.Context, req *relationpb.ListReq) (*relationpb.ListResp, error) {
		return r.client.ListFollowers(ctx, req)
	})
}

func (r *RelationClientAdapter) GetFollowings(ctx context.Context, userID int64) ([]int64, error) {
	return listRelationUserIDs(ctx, userID, func(ctx context.Context, req *relationpb.ListReq) (*relationpb.ListResp, error) {
		return r.client.ListFollowing(ctx, req)
	})
}

func listRelationUserIDs(
	ctx context.Context,
	userID int64,
	list func(context.Context, *relationpb.ListReq) (*relationpb.ListResp, error),
) ([]int64, error) {
	userIDs := make([]int64, 0, relationPageSize)
	var offset int32

	for {
		resp, err := list(ctx, &relationpb.ListReq{
			UserId: userID,
			Offset: offset,
			Limit:  relationPageSize,
		})
		if err != nil {
			return nil, err
		}
		if resp == nil || len(resp.Items) == 0 {
			return userIDs, nil
		}

		for _, item := range resp.Items {
			if item != nil {
				userIDs = append(userIDs, item.Id)
			}
		}
		if !resp.HasMore {
			return userIDs, nil
		}

		offset += int32(len(resp.Items))
	}
}
