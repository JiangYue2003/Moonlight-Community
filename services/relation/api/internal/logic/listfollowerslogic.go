package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type ListFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowersLogic {
	return &ListFollowersLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *ListFollowersLogic) ListFollowers(req *types.ListReq) ([]types.UserSummary, error) {
	uid := req.UserId
	if uid <= 0 {
		uid, _ = ctxdata.GetUserId(l.ctx)
	}
	r, err := l.svcCtx.RelationRpc.ListFollowers(l.ctx, &pb.ListReq{
		UserId: uid, Limit: req.Limit, Offset: req.Offset, Cursor: req.Cursor,
	})
	if err != nil {
		return nil, err
	}
	items, err := hydrateProfiles(l.ctx, l.svcCtx.UserRpc, summariesFromPb(r.Items))
	if err != nil {
		return nil, err
	}
	return items, nil
}
