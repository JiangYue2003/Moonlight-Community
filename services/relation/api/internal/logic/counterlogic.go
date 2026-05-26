package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
)

type CounterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CounterLogic {
	return &CounterLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// Counter 返回用户维度计数（关注数/粉丝数/发文数/获赞数）。
// 与 Java GET /api/v1/relation/counter 语义一致。
// userId=0 时取当前登录用户；否则取指定用户（公开数据，无需鉴权）。
func (l *CounterLogic) Counter(req *types.CounterReq) (*types.CounterResp, error) {
	targetId := req.UserId
	if targetId == 0 {
		uid, ok := ctxdata.GetUserId(l.ctx)
		if !ok || uid == 0 {
			return &types.CounterResp{}, nil
		}
		targetId = uid
	}

	resp, err := l.svcCtx.UserCounterRpc.GetUserSnapshot(l.ctx, &counterpb.GetUserSnapshotReq{
		UserId: targetId,
	})
	if err != nil {
		return nil, err
	}
	if resp.Snapshot == nil {
		return &types.CounterResp{}, nil
	}
	s := resp.Snapshot
	return &types.CounterResp{
		Followings: s.Followings,
		Followers:  s.Followers,
		Posts:      s.Posts,
		// 先按现有 user-counter 能力映射，保证前端字段可用
		LikedPosts: s.LikesReceived,
		FavedPosts: 0,
	}, nil
}
