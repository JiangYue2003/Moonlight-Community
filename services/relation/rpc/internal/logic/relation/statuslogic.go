package relationlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type StatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatusLogic {
	return &StatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Status 三态查询：直接查 following 表两次（小查询，无需缓存）。
//
//	following = ExistsActive(from→to)
//	followedBy = ExistsActive(to→from)
//	mutual = following && followedBy
func (l *StatusLogic) Status(in *relation.StatusReq) (*relation.StatusResp, error) {
	if in.FromUserId <= 0 || in.ToUserId <= 0 {
		return nil, errorx.New(errorx.CodeBadRequest, "user id required")
	}
	if in.FromUserId == in.ToUserId {
		return &relation.StatusResp{}, nil
	}
	following, err := l.svcCtx.FollowingModel.ExistsActive(l.ctx, in.FromUserId, in.ToUserId)
	if err != nil {
		return nil, err
	}
	followedBy, err := l.svcCtx.FollowingModel.ExistsActive(l.ctx, in.ToUserId, in.FromUserId)
	if err != nil {
		return nil, err
	}
	return &relation.StatusResp{
		Following:  following,
		FollowedBy: followedBy,
		Mutual:     following && followedBy,
	}, nil
}
