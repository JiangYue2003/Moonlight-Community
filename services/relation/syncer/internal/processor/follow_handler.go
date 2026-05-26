package processor

import (
	"context"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/pkg/txx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/event"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/zset"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/svc"

	goredis "github.com/redis/go-redis/v9"
)

// FollowHandler 处理 FollowCreated / FollowCanceled 事件。
//
// 步骤（与原 Java 一致）：
//  1. 写 follower 反查表（INSERT...ON DUPLICATE KEY 或 UPDATE rel_status=0）
//  2. 维护 ZSet 双向：uf:flws:{from} 与 uf:fans:{to}
//  3. 调 usercounter.Increment 更新双方计数
//
// 任意一步失败 → 返回 error 触发 Kafka 重试。SETNX 去重保证重试不会重复 INCR。
type FollowHandler struct {
	sc *svc.ServiceContext
}

func NewFollowHandler(sc *svc.ServiceContext) *FollowHandler { return &FollowHandler{sc: sc} }

func (h *FollowHandler) HandleCreated(ctx context.Context, ev event.RelationEvent) error {
	now := time.Now()
	score := float64(now.UnixMilli())

	// 1. follower 反查表
	id, err := h.sc.Snowflake.NextId()
	if err != nil {
		return err
	}
	if err := txx.WithTx(ctx, h.sc.Db, func(ctx context.Context, sess sqlx.Session) error {
		return h.sc.FollowerModel.UpsertActive(ctx, sess, id, ev.ToUserId, ev.FromUserId)
	}); err != nil {
		logx.WithContext(ctx).Errorf("FollowerModel.UpsertActive: %v", err)
		return err
	}

	// 2. ZSet 双向写
	if err := h.zsetAdd(ctx, zset.FollowingKey(ev.FromUserId), score, ev.ToUserId); err != nil {
		return err
	}
	if err := h.zsetAdd(ctx, zset.FollowerKey(ev.ToUserId), score, ev.FromUserId); err != nil {
		return err
	}

	// 3. usercounter 双向 +1
	if _, err := h.sc.UserCounterRpc.UserIncrement(ctx, &counterpb.UserIncrementReq{
		UserId: ev.FromUserId, Field: "followings", Delta: 1,
	}); err != nil {
		logx.WithContext(ctx).Errorf("usercounter followings +1: %v", err)
		return err
	}
	if _, err := h.sc.UserCounterRpc.UserIncrement(ctx, &counterpb.UserIncrementReq{
		UserId: ev.ToUserId, Field: "followers", Delta: 1,
	}); err != nil {
		logx.WithContext(ctx).Errorf("usercounter followers +1: %v", err)
		return err
	}
	return nil
}

func (h *FollowHandler) HandleCanceled(ctx context.Context, ev event.RelationEvent) error {
	// 1. follower 反查表 inactive
	if err := txx.WithTx(ctx, h.sc.Db, func(ctx context.Context, sess sqlx.Session) error {
		_, err := h.sc.FollowerModel.MarkInactive(ctx, sess, ev.ToUserId, ev.FromUserId)
		return err
	}); err != nil {
		logx.WithContext(ctx).Errorf("FollowerModel.MarkInactive: %v", err)
		return err
	}

	// 2. ZSet 双向 ZREM
	if err := h.sc.Redis.ZRem(ctx, zset.FollowingKey(ev.FromUserId), strconv.FormatInt(ev.ToUserId, 10)).Err(); err != nil {
		return err
	}
	if err := h.sc.Redis.ZRem(ctx, zset.FollowerKey(ev.ToUserId), strconv.FormatInt(ev.FromUserId, 10)).Err(); err != nil {
		return err
	}

	// 3. usercounter 双向 -1
	if _, err := h.sc.UserCounterRpc.UserIncrement(ctx, &counterpb.UserIncrementReq{
		UserId: ev.FromUserId, Field: "followings", Delta: -1,
	}); err != nil {
		logx.WithContext(ctx).Errorf("usercounter followings -1: %v", err)
		return err
	}
	if _, err := h.sc.UserCounterRpc.UserIncrement(ctx, &counterpb.UserIncrementReq{
		UserId: ev.ToUserId, Field: "followers", Delta: -1,
	}); err != nil {
		logx.WithContext(ctx).Errorf("usercounter followers -1: %v", err)
		return err
	}
	return nil
}

// zsetAdd 添加成员、刷新 TTL，并截断到容量上限（保留最近 N 条）。
func (h *FollowHandler) zsetAdd(ctx context.Context, key string, score float64, member int64) error {
	pipe := h.sc.Redis.Pipeline()
	pipe.ZAdd(ctx, key, goredis.Z{Score: score, Member: strconv.FormatInt(member, 10)})
	pipe.Expire(ctx, key, time.Duration(h.sc.Config.ZSet.TtlSeconds)*time.Second)
	// ZREMRANGEBYRANK 保留 score 最大的前 N 个：删 [0, -(MaxMembers+1)]
	if h.sc.Config.ZSet.MaxMembers > 0 {
		pipe.ZRemRangeByRank(ctx, key, 0, int64(-h.sc.Config.ZSet.MaxMembers-1))
	}
	_, err := pipe.Exec(ctx)
	return err
}
