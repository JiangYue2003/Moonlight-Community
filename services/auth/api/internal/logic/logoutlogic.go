package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Logout 调 auth-rpc.RevokeRefresh：RPC 内部解析 refresh token 拿 uid+jti 撤销白名单，
// 前端无需提交 uid。即使 token 已损坏，RPC 也会返回 revoked=false 而非错误，
// HTTP 接口对外保持幂等。
func (l *LogoutLogic) Logout(req *types.LogoutReq) (*types.Empty, error) {
	if _, err := l.svcCtx.AuthRpc.RevokeRefresh(l.ctx,
		&userpb.RevokeRefreshReq{RefreshToken: req.RefreshToken}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
