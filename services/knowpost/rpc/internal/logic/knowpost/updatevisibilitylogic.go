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
	committed, updateErr := updateAndEmitOutbox(l.ctx, l.svcCtx, row)
	if !committed {
		return nil, updateErr
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	bumpFeedPageSafety(l.ctx, l.svcCtx)
	if updateErr != nil {
		return nil, updateErr
	}
	return &knowpost.Empty{}, nil
}
