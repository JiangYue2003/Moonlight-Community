package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfirmContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmContentLogic {
	return &ConfirmContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfirmContentLogic) ConfirmContent(in *knowpost.ConfirmContentReq) (*knowpost.Empty, error) {
	// todo: add your logic here and delete this line

	return &knowpost.Empty{}, nil
}
