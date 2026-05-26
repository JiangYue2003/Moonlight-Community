package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/types"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.ActionReq) (*types.ActionResp, error) {
	return dispatchToggle(l.ctx, l.svcCtx, req, "like", true)
}
