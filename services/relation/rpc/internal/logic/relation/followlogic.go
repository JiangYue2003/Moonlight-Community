package relationlogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/txx"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/event"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Follow 写路径核心：限流 → 同事务（INSERT following + INSERT outbox）。
//
// 不维护 follower 反查表 / ZSet / usercounter——这些都是 syncer 异步消费 outbox 后做的伪从。
func (l *FollowLogic) Follow(in *relation.FollowReq) (*relation.FollowResp, error) {
	if in.FromUserId <= 0 || in.ToUserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "user id required")
	}
	if in.FromUserId == in.ToUserId {
		return nil, errorx.New(errorx.CodeBadRequest, "cannot follow self")
	}

	// 1. 限流：每用户每分钟最多 100 次（容量 100，速率 1/s 持续补充）
	allowed, err := l.svcCtx.RateLimiter.Take(l.ctx,
		fmt.Sprintf("rl:follow:{%d}", in.FromUserId),
		l.svcCtx.Config.RateLimit.FollowCapacity,
		l.svcCtx.Config.RateLimit.FollowRefillPerSec)
	if err != nil {
		l.Logger.Errorf("ratelimit take: %v", err)
		// 限流系统挂了应该放行而不是阻塞核心写路径；记日志即可。
	} else if !allowed {
		return nil, errorx.New(errorx.CodeRateLimited, "follow rate limited")
	}

	outboxId, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}
	relRowId, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}

	payloadBytes, _ := json.Marshal(event.RelationEvent{
		Type:       event.TypeFollowCreated,
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Id:         outboxId,
	})
	payload := string(payloadBytes)
	changed := false

	err = txx.WithTx(l.ctx, l.svcCtx.Db, func(ctx context.Context, sess sqlx.Session) error {
		exists, err := l.svcCtx.FollowingModel.ExistsActive(ctx, in.FromUserId, in.ToUserId)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}
		if err := l.svcCtx.FollowingModel.UpsertActive(ctx, sess, relRowId, in.FromUserId, in.ToUserId); err != nil {
			return err
		}
		changed = true
		return l.svcCtx.OutboxModel.InsertInTx(ctx, sess, outboxId, "following", in.FromUserId,
			event.TypeFollowCreated, payload)
	})
	if err != nil {
		return nil, err
	}
	return &relation.FollowResp{Changed: changed}, nil
}
