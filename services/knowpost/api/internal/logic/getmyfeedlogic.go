package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetMyFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFeedLogic {
	return &GetMyFeedLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *GetMyFeedLogic) GetMyFeed(req *types.GetMyFeedReq) (*types.FeedPage, error) {
	uid, _ := ctxdata.GetUserId(l.ctx)
	r, err := l.svcCtx.KnowPostRpc.GetMyFeed(l.ctx, &pb.GetMyFeedReq{
		CreatorId: uid, Page: req.Page, Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return feedPageFromPb(l.ctx, l.svcCtx, r), nil
}
