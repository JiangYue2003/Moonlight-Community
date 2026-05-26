package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetPublicFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicFeedLogic {
	return &GetPublicFeedLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *GetPublicFeedLogic) GetPublicFeed(req *types.GetPublicFeedReq) (*types.FeedPage, error) {
	r, err := l.svcCtx.KnowPostRpc.GetPublicFeed(l.ctx, &pb.GetPublicFeedReq{
		Page: req.Page, Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return feedPageFromPb(l.ctx, l.svcCtx, r), nil
}
