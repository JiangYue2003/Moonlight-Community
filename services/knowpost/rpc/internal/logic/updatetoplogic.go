package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopLogic {
	return &UpdateTopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTopLogic) UpdateTop(in *knowpost.UpdateTopReq) (*knowpost.Empty, error) {
	// todo: add your logic here and delete this line

	return &knowpost.Empty{}, nil
}
