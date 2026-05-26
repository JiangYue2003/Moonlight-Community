package relationlogic

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/txx"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"github.com/zhiguang/zhiguang-go/services/relation/shared/event"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Unfollow 与 Follow 对称：UPDATE rel_status=0 + INSERT outbox FollowCanceled。
// 关注关系本就不存在时返回 changed=false，但仍写 outbox 让下游 syncer 幂等清理。
func (l *UnfollowLogic) Unfollow(in *relation.UnfollowReq) (*relation.UnfollowResp, error) {
	if in.FromUserId <= 0 || in.ToUserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "user id required")
	}
	if in.FromUserId == in.ToUserId {
		return nil, errorx.New(errorx.CodeBadRequest, "cannot unfollow self")
	}

	outboxId, err := l.svcCtx.Snowflake.NextId()
	if err != nil {
		return nil, err
	}

	payloadBytes, _ := json.Marshal(event.RelationEvent{
		Type:       event.TypeFollowCanceled,
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Id:         outboxId,
	})
	payload := string(payloadBytes)

	var changed int64
	err = txx.WithTx(l.ctx, l.svcCtx.Db, func(ctx context.Context, sess sqlx.Session) error {
		n, err := l.svcCtx.FollowingModel.MarkInactive(ctx, sess, in.FromUserId, in.ToUserId)
		if err != nil {
			return err
		}
		changed = n
		if changed == 0 {
			return nil
		}
		return l.svcCtx.OutboxModel.InsertInTx(ctx, sess, outboxId, "following", in.FromUserId,
			event.TypeFollowCanceled, payload)
	})
	if err != nil {
		return nil, err
	}
	return &relation.UnfollowResp{Changed: changed > 0}, nil
}
