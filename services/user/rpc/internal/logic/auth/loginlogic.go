package authlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.AuthResp, error) {
	id := normalizeIdentifier(in.Identifier)
	if err := validateIdentifier(id); err != nil {
		return nil, err
	}

	u, err := l.svcCtx.UsersModel.FindOneByIdentifier(l.ctx, id)
	if err != nil || u == nil {
		recordLoginLog(l.ctx, l.svcCtx, 0, id, in.Channel, in.Ip, in.UserAgent, "FAILED")
		return nil, errorx.New(errorx.CodeNotFound, "user not found")
	}
	uid := int64(u.Id)

	switch in.Channel {
	case "PASSWORD":
		if !u.PasswordHash.Valid || u.PasswordHash.String == "" {
			recordLoginLog(l.ctx, l.svcCtx, uid, id, "PASSWORD", in.Ip, in.UserAgent, "FAILED")
			return nil, errorx.New(errorx.CodeInvalidCredentials, "invalid credentials")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash.String), []byte(in.Password)); err != nil {
			recordLoginLog(l.ctx, l.svcCtx, uid, id, "PASSWORD", in.Ip, in.UserAgent, "FAILED")
			return nil, errorx.New(errorx.CodeInvalidCredentials, "invalid credentials")
		}
	case "CODE":
		if err := l.svcCtx.Verifier.Verify(l.ctx, "LOGIN", id, in.Code); err != nil {
			recordLoginLog(l.ctx, l.svcCtx, uid, id, "CODE", in.Ip, in.UserAgent, "FAILED")
			return nil, err
		}
	default:
		return nil, errorx.New(errorx.CodeBadRequest, "invalid channel")
	}

	pair, err := issueAndPersist(l.ctx, l.svcCtx, uid, u.Nickname)
	if err != nil {
		return nil, err
	}
	recordLoginLog(l.ctx, l.svcCtx, uid, id, in.Channel, in.Ip, in.UserAgent, "SUCCESS")
	return &user.AuthResp{User: modelToAuthUser(u), Token: pair}, nil
}
