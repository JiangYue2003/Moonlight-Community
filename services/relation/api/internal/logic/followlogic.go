// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp bool, err error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return false, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	r, err := l.svcCtx.RelationRpc.Follow(l.ctx, &pb.FollowReq{
		FromUserId: uid,
		ToUserId:   req.ToUserId,
	})
	if err != nil {
		return false, err
	}
	return r.GetChanged(), nil
}
