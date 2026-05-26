package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/profile/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type GetMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMeLogic) GetMe() (*types.ProfileResp, error) {
	uid, ok := ctxdata.GetUserId(l.ctx)
	if !ok {
		return nil, errorx.New(errorx.CodeUnauthorized, "missing user id")
	}
	r, err := l.svcCtx.UserRpc.GetById(l.ctx, &userpb.GetByIdReq{Id: uid})
	if err != nil {
		return nil, err
	}
	return toProfileResp(r.User), nil
}
