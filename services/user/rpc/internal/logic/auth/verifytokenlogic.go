package authlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"github.com/zhiguang/zhiguang-go/services/user/rpc/internal/svc"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VerifyToken 提供给其它 RPC 服务做 access token 校验（无 db 依赖，仅 RS256 验签）。
func (l *VerifyTokenLogic) VerifyToken(in *user.VerifyTokenReq) (*user.VerifyTokenResp, error) {
	claims, err := l.svcCtx.JwtSigner.ParseAccess(in.AccessToken)
	if err != nil {
		return &user.VerifyTokenResp{Valid: false}, nil
	}
	return &user.VerifyTokenResp{
		Valid:    true,
		UserId:   claims.Uid,
		Nickname: claims.Nickname,
	}, nil
}
