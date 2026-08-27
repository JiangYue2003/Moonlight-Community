package main

import (
	"context"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type userFeedRPC interface {
	GetUserFeed(context.Context, *knowpostpb.GetUserFeedReq, ...grpc.CallOption) (*knowpostpb.FeedPage, error)
}

type rpcFeedClient struct {
	client userFeedRPC
}

func newRPCFeedClient(client userFeedRPC) *rpcFeedClient {
	return &rpcFeedClient{client: client}
}

func (c *rpcFeedClient) GetUserFeed(
	ctx context.Context,
	reader readerIdentity,
	read feedReadRequest,
) (*feedPageResponse, error) {
	resp, err := c.client.GetUserFeed(ctx, &knowpostpb.GetUserFeedReq{
		UserId: reader.UserID,
		Page:   read.Page,
		Size:   read.Size,
		Cursor: read.Cursor,
	})
	if err != nil {
		return nil, err
	}
	result := &feedPageResponse{
		Items:      make([]feedItemResponse, 0, len(resp.Items)),
		HasMore:    resp.HasMore,
		Page:       resp.Page,
		Size:       resp.Size,
		NextCursor: resp.NextCursor,
	}
	for _, item := range resp.Items {
		result.Items = append(result.Items, feedItemResponse{
			ID:        item.Id,
			CreatorID: item.CreatorId,
			Title:     item.Title,
		})
	}
	return result, nil
}
