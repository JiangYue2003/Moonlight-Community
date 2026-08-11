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

func (c *CounterClientAdapter) BatchGetFollowerCounts(ctx context.Context, userIDs []int64) (map[int64]int64, error) {
	counts := make(map[int64]int64, len(userIDs))
	if len(userIDs) == 0 {
		return counts, nil
	}

	resp, err := c.client.BatchGetUserSnapshot(ctx, &counterpb.BatchGetUserSnapshotReq{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return counts, nil
	}

	for _, userID := range userIDs {
		if snapshot := resp.Result[userID]; snapshot != nil {
			counts[userID] = snapshot.Followers
		}
	}
	return counts, nil
}
