package logic

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) WithRequest(r *http.Request) *LoginLogic { l.r = r; return l }

func (l *LoginLogic) Login(req *types.LoginReq) (*types.AuthResp, error) {
	channel := strings.ToUpper(strings.TrimSpace(req.Channel))
	if channel == "" {
		switch {
		case strings.TrimSpace(req.Code) != "":
			channel = "CODE"
		case strings.TrimSpace(req.Password) != "":
			channel = "PASSWORD"
		}
	}
	resp, err := l.svcCtx.AuthRpc.Login(l.ctx, &userpb.LoginReq{
		Identifier: req.Identifier,
		Password:   req.Password,
		Code:       req.Code,
		Channel:    channel,
		Ip:         clientIp(l.r),
		UserAgent:  clientUa(l.r),
	})
	if err != nil {
		return nil, err
	}
	return pbToAuthResp(resp), nil
}
