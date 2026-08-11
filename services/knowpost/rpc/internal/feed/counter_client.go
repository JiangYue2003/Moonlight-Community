package feed

import (
	"context"
)

// CounterClient Counter服务客户端接口
type CounterClient interface {
	// BatchGetFollowerCounts 批量获取用户的粉丝数，避免按关注用户逐个 RPC。
	BatchGetFollowerCounts(ctx context.Context, userIDs []int64) (map[int64]int64, error)
}
