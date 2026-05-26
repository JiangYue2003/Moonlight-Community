package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/types"
)

type FavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavLogic {
	return &FavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavLogic) Fav(req *types.ActionReq) (*types.ActionResp, error) {
	return dispatchToggle(l.ctx, l.svcCtx, req, "fav", true)
}
