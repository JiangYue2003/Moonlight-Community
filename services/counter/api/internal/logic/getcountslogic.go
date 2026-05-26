package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/types"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type GetCountsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountsLogic {
	return &GetCountsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCountsLogic) GetCounts(req *types.GetCountsPath) (*types.CountsResp, error) {
	resp, err := l.svcCtx.CounterRpc.GetCounts(l.ctx, &counter.GetCountsReq{
		EntityType: req.EntityType,
		EntityId:   req.EntityId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CountsResp{
		EntityType: req.EntityType,
		EntityId:   req.EntityId,
		Counts:     resp.Counts,
	}, nil
}
