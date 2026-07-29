package feed

import (
	"context"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

// CounterClientAdapter 实现 CounterClient 接口，适配 Counter RPC
type CounterClientAdapter struct {
	client counterpb.UserCounterClient
}

func NewCounterClientAdapter(client counterpb.UserCounterClient) *CounterClientAdapter {
	return &CounterClientAdapter{client: client}
}

func (c *CounterClientAdapter) GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	resp, err := c.client.GetUserSnapshot(ctx, &counterpb.GetUserSnapshotReq{
		UserId: userID,
	})
	if err != nil {
		return 0, err
	}

	if resp == nil || resp.Snapshot == nil {
		return 0, nil
	}

	return resp.Snapshot.Followers, nil
}
