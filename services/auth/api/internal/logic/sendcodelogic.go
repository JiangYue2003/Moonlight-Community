package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (*types.SendCodeResp, error) {
	r, err := l.svcCtx.AuthRpc.SendCode(l.ctx, &userpb.SendCodeReq{
		Scene:      req.Scene,
		Identifier: req.Identifier,
	})
	if err != nil {
		return nil, err
	}
	return &types.SendCodeResp{
		CooldownSeconds: r.CooldownSeconds,
		ExpireSeconds:   r.ExpireSeconds,
	}, nil
}
