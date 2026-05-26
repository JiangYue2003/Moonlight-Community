package knowpostlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
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
	invalidateKnowPostCaches(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	row, err := findOwnedRow(l.ctx, l.svcCtx, in.Id, in.CreatorId)
	if err != nil {
		return nil, err
	}
	if in.IsTop {
		row.IsTop = 1
	} else {
		row.IsTop = 0
	}
	if err := updateAndEmitOutbox(l.ctx, l.svcCtx, row); err != nil {
		return nil, err
	}
	invalidateKnowPostCaches(l.ctx, l.svcCtx, int64(row.Id), in.CreatorId)
	return &knowpost.Empty{}, nil
}
