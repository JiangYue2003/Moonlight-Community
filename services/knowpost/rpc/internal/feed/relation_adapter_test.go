package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type pagedRelationRPC struct {
	relationpb.RelationClient
	items             []*relationpb.UserSummary
	followerRequests  []*relationpb.ListReq
	followingRequests []*relationpb.ListReq
}

func (c *pagedRelationRPC) ListFollowers(
	_ context.Context,
	req *relationpb.ListReq,
	_ ...grpc.CallOption,
) (*relationpb.ListResp, error) {
	c.followerRequests = append(c.followerRequests, req)
	return relationPage(c.items, req), nil
}

func (c *pagedRelationRPC) ListFollowing(
	_ context.Context,
	req *relationpb.ListReq,
	_ ...grpc.CallOption,
) (*relationpb.ListResp, error) {
	c.followingRequests = append(c.followingRequests, req)
	return relationPage(c.items, req), nil
}

func relationPage(items []*relationpb.UserSummary, req *relationpb.ListReq) *relationpb.ListResp {
	limit := int(req.Limit)
	if limit <= 0 {
		limit = 20
	}
	start := int(req.Offset)
	if start >= len(items) {
		return &relationpb.ListResp{}
	}
	end := start + limit
	if end > len(items) {
		end = len(items)
	}
	return &relationpb.ListResp{
		Items:   items[start:end],
		HasMore: end < len(items),
	}
}

func TestRelationClientAdapterPaginatesFollowersAndFollowings(t *testing.T) {
	items := make([]*relationpb.UserSummary, 125)
	want := make([]int64, 125)
	for i := range items {
		id := int64(i + 1)
		items[i] = &relationpb.UserSummary{Id: id}
		want[i] = id
	}

	client := &pagedRelationRPC{items: items}
	adapter := NewRelationClientAdapter(client)

	followers, err := adapter.GetFollowers(context.Background(), 42)
	require.NoError(t, err)
	require.Equal(t, want, followers)
	require.Len(t, client.followerRequests, 2)
	require.Equal(t, int32(100), client.followerRequests[0].Limit)
	require.Equal(t, int32(0), client.followerRequests[0].Offset)
	require.Equal(t, int32(100), client.followerRequests[1].Offset)
	require.True(t, client.followerRequests[0].IdsOnly)
	require.True(t, client.followerRequests[1].IdsOnly)

	followings, err := adapter.GetFollowings(context.Background(), 42)
	require.NoError(t, err)
	require.Equal(t, want, followings)
	require.Len(t, client.followingRequests, 2)
	require.Equal(t, int32(100), client.followingRequests[0].Limit)
	require.Equal(t, int32(100), client.followingRequests[1].Offset)
	require.True(t, client.followingRequests[0].IdsOnly)
	require.True(t, client.followingRequests[1].IdsOnly)
}
