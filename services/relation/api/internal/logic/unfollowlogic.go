package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *UnfollowLogic) Unfollow(req *types.UnfollowReq) (bool, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return false, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	r, err := l.svcCtx.RelationRpc.Unfollow(l.ctx, &pb.UnfollowReq{
		FromUserId: uid, ToUserId: req.ToUserId,
	})
	if err != nil {
		return false, err
	}
	return r.GetChanged(), nil
}
