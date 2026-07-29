package feed

import (
	"context"

	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

// RelationClientAdapter 实现 RelationClient 接口，适配 Relation RPC
type RelationClientAdapter struct {
	client relationpb.RelationClient
}

func NewRelationClientAdapter(client relationpb.RelationClient) *RelationClientAdapter {
	return &RelationClientAdapter{client: client}
}

func (r *RelationClientAdapter) GetFollowers(ctx context.Context, userID int64) ([]int64, error) {
	resp, err := r.client.ListFollowers(ctx, &relationpb.ListReq{
		UserId: userID,
		// 暂时不分页，获取所有粉丝
		// TODO: 如果粉丝太多，需要分页处理
	})
	if err != nil {
		return nil, err
	}

	if resp == nil || len(resp.Items) == 0 {
		return []int64{}, nil
	}

	// 从 UserSummary 中提取 Id
	userIDs := make([]int64, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			userIDs = append(userIDs, item.Id)
		}
	}

	return userIDs, nil
}

func (r *RelationClientAdapter) GetFollowings(ctx context.Context, userID int64) ([]int64, error) {
	resp, err := r.client.ListFollowing(ctx, &relationpb.ListReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	if resp == nil || len(resp.Items) == 0 {
		return []int64{}, nil
	}

	// 从 UserSummary 中提取 Id
	userIDs := make([]int64, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item != nil {
			userIDs = append(userIDs, item.Id)
		}
	}

	return userIDs, nil
}
