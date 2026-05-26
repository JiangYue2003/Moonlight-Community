package knowpostlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type UpdateVisibilityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVisibilityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVisibilityLogic {
	return &UpdateVisibilityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVisibilityLogic) UpdateVisibility(in *knowpost.UpdateVisibilityReq) (*knowpost.Empty, error) {
	if !validVisible(in.Visible) {
		return nil, errorx.New(errorx.CodeBadRequest, "invalid visible value")
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	row.Visible = in.Visible
	if err := updateAndEmitOutbox(l.ctx, l.svcCtx, row); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	return &knowpost.Empty{}, nil
}
