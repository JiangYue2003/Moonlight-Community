package logic

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
)

type CreateSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSessionLogic {
	return &CreateSessionLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CreateSessionLogic) Create(req *types.CreateSessionReq) (*types.CreateSessionResp, error) {
	userID, _ := ctxdata.GetUserId(l.ctx)
	if userID <= 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	now := time.Now().UnixMilli()
	sid := "as_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	title := strings.TrimSpace(req.Title)
	if title == "" {
		title = "新会话"
	}
	if _, err := l.svcCtx.Db.ExecCtx(l.ctx,
		"INSERT INTO agent_sessions (session_id,user_id,title,created_at,updated_at) VALUES (?,?,?,?,?)",
		sid, userID, svc.TrimContent(title, 255), now, now,
	); err != nil {
		return nil, err
	}
	return &types.CreateSessionResp{SessionID: sid}, nil
}
