package logic

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	token  string
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) WithToken(t string) *MeLogic { l.token = t; return l }

// Me 通过 auth-rpc.VerifyToken 校验 access token 并返回当前用户身份。
// 阶段1只返回 id + nickname；其它字段后续接 user-rpc.GetById 补全。
func (l *MeLogic) Me() (*types.AuthUser, error) {
	v, err := l.svcCtx.AuthRpc.VerifyToken(l.ctx, &userpb.VerifyTokenReq{AccessToken: l.token})
	if err != nil {
		return nil, err
	}
	if !v.Valid {
		return nil, errorx.New(errorx.CodeUnauthorized, "invalid access token")
	}
	g, err := l.svcCtx.UserRpc.GetById(l.ctx, &userpb.GetByIdReq{Id: v.UserId})
	if err != nil || g.User == nil {
		// 兜底：最少返回 token 中可用信息，避免前端会话直接失效
		return &types.AuthUser{
			Id:       v.UserId,
			Nickname: v.Nickname,
		}, nil
	}
	u := g.User
	out := &types.AuthUser{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
		Email:    u.Email,
		ZgId:     u.ZgId,
		ZhId:     u.ZgId,
		Birthday: u.Birthday,
		School:   u.School,
		Bio:      u.Bio,
		Gender:   u.Gender,
		TagsJson: u.TagsJson,
		TagJson:  u.TagsJson,
	}
	if strings.TrimSpace(u.TagsJson) != "" {
		out.Skills = nil
	}
	return out, nil
}
