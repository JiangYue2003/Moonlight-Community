package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type CreateDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDraftLogic {
	return &CreateDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDraftLogic) CreateDraft() (*types.CreateDraftResp, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	r, err := l.svcCtx.KnowPostRpc.CreateDraft(l.ctx, &pb.CreateDraftReq{CreatorId: uid})
	if err != nil {
		return nil, err
	}
	return &types.CreateDraftResp{Id: r.Id}, nil
}
