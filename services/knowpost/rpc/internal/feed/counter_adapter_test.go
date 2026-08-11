package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type batchSnapshotRPC struct {
	counterpb.UserCounterClient
	calls   int
	userIDs []int64
}

func (c *batchSnapshotRPC) BatchGetUserSnapshot(
	_ context.Context,
	req *counterpb.BatchGetUserSnapshotReq,
	_ ...grpc.CallOption,
) (*counterpb.BatchGetUserSnapshotResp, error) {
	c.calls++
	c.userIDs = append([]int64(nil), req.UserIds...)
	return &counterpb.BatchGetUserSnapshotResp{
		Result: map[int64]*counterpb.UserSnapshot{
			11: {UserId: 11, Followers: 25},
			12: {UserId: 12, Followers: 2500},
		},
	}, nil
}

func TestCounterClientAdapterGetsFollowerCountsInOneBatch(t *testing.T) {
	rpc := &batchSnapshotRPC{}
	adapter := NewCounterClientAdapter(rpc)

	counts, err := adapter.BatchGetFollowerCounts(context.Background(), []int64{11, 12})

	require.NoError(t, err)
	require.Equal(t, map[int64]int64{11: 25, 12: 2500}, counts)
	require.Equal(t, 1, rpc.calls)
	require.Equal(t, []int64{11, 12}, rpc.userIDs)
}
