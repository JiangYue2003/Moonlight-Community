package knowpostlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	svc "github.com/zhiguang/zhiguang-go/services/knowpost/internal/application"
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
