package logic

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) WithRequest(r *http.Request) *RegisterLogic { l.r = r; return l }

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.AuthResp, error) {
	nickname := strings.TrimSpace(req.Nickname)
	if nickname == "" {
		// 前端当前不传 nickname，后端兜底生成默认值
		nickname = "知光用户"
	}
	resp, err := l.svcCtx.AuthRpc.Register(l.ctx, &userpb.RegisterReq{
		Identifier: req.Identifier,
		Password:   req.Password,
		Code:       req.Code,
		Nickname:   nickname,
		AgreeTerms: req.AgreeTerms,
		Ip:         clientIp(l.r),
		UserAgent:  clientUa(l.r),
	})
	if err != nil {
		return nil, err
	}
	return pbToAuthResp(resp), nil
}

func clientIp(r *http.Request) string {
	if r == nil {
		return ""
	}
	if v := r.Header.Get("X-Forwarded-For"); v != "" {
		return v
	}
	if v := r.Header.Get("X-Real-IP"); v != "" {
		return v
	}
	return httpx.GetRemoteAddr(r)
}

func clientUa(r *http.Request) string {
	if r == nil {
		return ""
	}
	return r.UserAgent()
}
