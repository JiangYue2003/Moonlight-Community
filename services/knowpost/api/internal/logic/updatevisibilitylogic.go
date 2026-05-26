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

type UpdateVisibilityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewUpdateVisibilityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVisibilityLogic {
	return &UpdateVisibilityLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}
func (l *UpdateVisibilityLogic) WithId(id int64) *UpdateVisibilityLogic { l.id = id; return l }

func (l *UpdateVisibilityLogic) UpdateVisibility(req *types.UpdateVisibilityReq) (*types.Empty, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	if _, err := l.svcCtx.KnowPostRpc.UpdateVisibility(l.ctx, &pb.UpdateVisibilityReq{
		Id: l.id, CreatorId: uid, Visible: req.Visible,
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
