package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublicFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicFeedLogic {
	return &GetPublicFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublicFeedLogic) GetPublicFeed(in *knowpost.GetPublicFeedReq) (*knowpost.FeedPage, error) {
	// todo: add your logic here and delete this line

	return &knowpost.FeedPage{}, nil
}
