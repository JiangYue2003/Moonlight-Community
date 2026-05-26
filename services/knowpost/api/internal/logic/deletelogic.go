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

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}
func (l *DeleteLogic) WithId(id int64) *DeleteLogic { l.id = id; return l }

func (l *DeleteLogic) Delete() (*types.Empty, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	if _, err := l.svcCtx.KnowPostRpc.Delete(l.ctx, &pb.DeleteReq{Id: l.id, CreatorId: uid}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
