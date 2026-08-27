package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFeedLogic {
	return &GetMyFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyFeedLogic) GetMyFeed(in *knowpost.GetMyFeedReq) (*knowpost.FeedPage, error) {
	// todo: add your logic here and delete this line

	return &knowpost.FeedPage{}, nil
}
