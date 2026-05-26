package authlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
)

type RevokeRefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevokeRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevokeRefreshLogic {
	return &RevokeRefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RevokeRefresh 解析 refresh token（RS256 验签）→ 撤销白名单 jti。
// HTTP 网关 logout 时直接调用，前端无需提交 uid+jti。
func (l *RevokeRefreshLogic) RevokeRefresh(in *user.RevokeRefreshReq) (*user.RevokeRefreshResp, error) {
	if in.RefreshToken == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "refresh token required")
	}
	claims, err := l.svcCtx.JwtSigner.ParseRefresh(in.RefreshToken)
	if err != nil {
		// token 无法解析视为已失效；返回 revoked=false 而非错误，避免暴露细节。
		l.Logger.Infof("revoke refresh: token invalid: %v", err)
		return &user.RevokeRefreshResp{Revoked: false}, nil
	}
	if err := l.svcCtx.Tokens.Revoke(l.ctx, claims.Uid, claims.ID); err != nil {
		return nil, err
	}
	return &user.RevokeRefreshResp{Revoked: true}, nil
}
