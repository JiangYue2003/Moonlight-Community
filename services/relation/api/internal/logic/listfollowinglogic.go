package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
)

type ListFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowingLogic {
	return &ListFollowingLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// ListFollowing UserId=0 时表示当前登录用户。
func (l *ListFollowingLogic) ListFollowing(req *types.ListReq) ([]types.UserSummary, error) {
	uid := req.UserId
	if uid <= 0 {
		uid, _ = ctxdata.GetUserId(l.ctx)
	}
	r, err := l.svcCtx.RelationRpc.ListFollowing(l.ctx, &pb.ListReq{
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
