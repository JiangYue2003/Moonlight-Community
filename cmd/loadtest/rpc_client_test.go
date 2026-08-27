package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type stubRPCFeed struct {
	request *knowpostpb.GetUserFeedReq
}

func (s *stubRPCFeed) GetUserFeed(
	_ context.Context,
	in *knowpostpb.GetUserFeedReq,
	_ ...grpc.CallOption,
) (*knowpostpb.FeedPage, error) {
	s.request = in
	return &knowpostpb.FeedPage{
		Items: []*knowpostpb.FeedItem{{Id: "7", CreatorId: 9}}, NextCursor: "next-cursor",
	}, nil
}

func TestRPCFeedClientUsesReaderUserID(t *testing.T) {
	stub := &stubRPCFeed{}
	client := newRPCFeedClient(stub)

	page, err := client.GetUserFeed(context.Background(), readerIdentity{UserID: 42}, feedReadRequest{
		Page: 3, Size: 15, Cursor: "current-cursor",
	})

	require.NoError(t, err)
	require.Equal(t, int64(42), stub.request.UserId)
	require.Equal(t, int32(3), stub.request.Page)
	require.Equal(t, int32(15), stub.request.Size)
	require.Equal(t, "current-cursor", stub.request.Cursor)
	require.Equal(t, "7", page.Items[0].ID)
	require.Equal(t, "next-cursor", page.NextCursor)
}
