package authlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *SendCodeLogic) SendCode(in *user.SendCodeReq) (*user.SendCodeResp, error) {
	id := normalizeIdentifier(in.Identifier)
	if err := validateIdentifier(id); err != nil {
		return nil, err
	}
	switch in.Scene {
	case "REGISTER":
		existing, _ := l.svcCtx.UsersModel.FindOneByIdentifier(l.ctx, id)
		if existing != nil {
			return nil, errorx.New(errorx.CodeIdentifierExists, "identifier already registered")
		}
	case "LOGIN", "RESET_PASSWORD":
		existing, _ := l.svcCtx.UsersModel.FindOneByIdentifier(l.ctx, id)
		if existing == nil {
			return nil, errorx.New(errorx.CodeIdentifierNotFound, "identifier not registered")
		}
	default:
		return nil, errorx.New(errorx.CodeBadRequest, "invalid scene")
	}

	code, err := l.svcCtx.Verifier.Send(l.ctx, in.Scene, id)
	if err != nil {
		return nil, err
	}
	logx.WithContext(l.ctx).Infof("[verification] scene=%s identifier=%s code=%s", in.Scene, id, code)

	return &user.SendCodeResp{
		CooldownSeconds: int32(l.svcCtx.Verifier.CooldownSeconds()),
		ExpireSeconds:   int32(l.svcCtx.Verifier.ExpireSeconds()),
	}, nil
}
