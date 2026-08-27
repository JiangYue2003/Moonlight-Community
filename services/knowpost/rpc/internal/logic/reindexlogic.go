package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReindexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReindexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReindexLogic {
	return &ReindexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReindexLogic) Reindex(in *knowpost.ReindexReq) (*knowpost.Empty, error) {
	// todo: add your logic here and delete this line

	return &knowpost.Empty{}, nil
}
