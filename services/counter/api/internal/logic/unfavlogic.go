package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/types"
)

type UnfavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfavLogic {
	return &UnfavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfavLogic) Unfav(req *types.ActionReq) (*types.ActionResp, error) {
	return dispatchToggle(l.ctx, l.svcCtx, req, "fav", false)
}
