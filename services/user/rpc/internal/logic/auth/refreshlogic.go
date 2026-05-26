package authlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *RefreshLogic) Refresh(in *user.RefreshReq) (*user.AuthResp, error) {
	claims, err := l.svcCtx.JwtSigner.ParseRefresh(in.RefreshToken)
	if err != nil {
		return nil, errorx.Wrap(errorx.CodeRefreshTokenInvalid, "refresh token invalid", err)
	}
	ok, err := l.svcCtx.Tokens.Valid(l.ctx, claims.Uid, claims.ID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errorx.New(errorx.CodeRefreshTokenInvalid, "refresh token revoked")
	}
	_ = l.svcCtx.Tokens.Revoke(l.ctx, claims.Uid, claims.ID)

	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(claims.Uid))
	if err != nil {
		return nil, err
	}
	pair, err := issueAndPersist(l.ctx, l.svcCtx, claims.Uid, u.Nickname)
	if err != nil {
		return nil, err
	}
	return &user.AuthResp{User: modelToAuthUser(u), Token: pair}, nil
}
