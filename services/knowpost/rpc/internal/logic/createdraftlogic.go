package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDraftLogic {
	return &CreateDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDraftLogic) CreateDraft(in *knowpost.CreateDraftReq) (*knowpost.CreateDraftResp, error) {
	// todo: add your logic here and delete this line

	return &knowpost.CreateDraftResp{}, nil
}
