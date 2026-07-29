package feed

import (
	"context"
)

// CounterClient Counter服务客户端接口
type CounterClient interface {
	// GetFollowerCount 获取用户的粉丝数
	GetFollowerCount(ctx context.Context, userID int64) (int64, error)
}
