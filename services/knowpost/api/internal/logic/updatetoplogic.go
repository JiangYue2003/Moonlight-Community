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

type UpdateTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewUpdateTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopLogic {
	return &UpdateTopLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}
func (l *UpdateTopLogic) WithId(id int64) *UpdateTopLogic { l.id = id; return l }

func (l *UpdateTopLogic) UpdateTop(req *types.UpdateTopReq) (*types.Empty, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	if _, err := l.svcCtx.KnowPostRpc.UpdateTop(l.ctx, &pb.UpdateTopReq{
		Id: l.id, CreatorId: uid, IsTop: req.IsTop,
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
