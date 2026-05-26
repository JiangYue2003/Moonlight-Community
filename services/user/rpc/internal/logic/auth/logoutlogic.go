package authlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Logout 撤销单个 refresh jti；access token 在 TTL 内仍然有效（无状态 JWT 限制）。
func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.Empty, error) {
	if in.UserId > 0 && in.RefreshTokenId != "" {
		_ = l.svcCtx.Tokens.Revoke(l.ctx, in.UserId, in.RefreshTokenId)
	}
	return &user.Empty{}, nil
}
