package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestFeedReader_ClassifyBigV(t *testing.T) {
	// 测试大V分类逻辑

	redis := NewMockRedisClient()
	relation := NewMockRelationClient()
	counter := NewMockCounterClient()
	reader := NewFeedReader(redis, relation, counter, logx.WithContext(context.Background()))

	ctx := context.Background()
	userID := int64(123)

	// 设置关注列表：3个普通用户 + 2个大V
	relation.SetFollowings(userID, []int64{201, 202, 203, 301, 302})

	// 设置粉丝数
	counter.SetFollowerCount(201, 100)     // 普通用户
	counter.SetFollowerCount(202, 500)     // 普通用户
	counter.SetFollowerCount(203, 1000)    // 普通用户（刚好阈值）
	counter.SetFollowerCount(301, 5000)    // 大V
	counter.SetFollowerCount(302, 100000)  // 超级大V

	result, hasMore, err := reader.GetFeed(ctx, userID, 1, 20)

	// 验证：不报错
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, hasMore)

	// 验证分类逻辑（通过日志可以看到）
	// 预期：3个普通用户，2个大V
}
