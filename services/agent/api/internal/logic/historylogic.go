package logic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
)

type HistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryLogic {
	return &HistoryLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *HistoryLogic) Get(req *types.HistoryReq) (*types.HistoryResp, error) {
	userID, _ := ctxdata.GetUserId(l.ctx)
	if userID <= 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	if strings.TrimSpace(req.SessionID) == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "sessionId required")
	}
	if err := l.mustOwn(userID, req.SessionID); err != nil {
		return nil, err
	}
	limit := req.Limit
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	vals, err := l.svcCtx.Redis.LRange(l.ctx, svc.SessionMessagesKey(userID, req.SessionID), int64(-limit), -1).Result()
	if err != nil {
		return nil, err
	}
	items := make([]types.MessageItem, 0, len(vals))
	for _, v := range vals {
		var m types.MessageItem
		if json.Unmarshal([]byte(v), &m) == nil {
			items = append(items, m)
		}
	}
	return &types.HistoryResp{Items: items}, nil
}

func (l *HistoryLogic) mustOwn(userID int64, sessionID string) error {
	var uid int64
	if err := l.svcCtx.Db.QueryRowCtx(l.ctx, &uid, "SELECT user_id FROM agent_sessions WHERE session_id=? LIMIT 1", sessionID); err != nil {
		return errorx.New(errorx.CodeNotFound, "session not found")
	}
	if uid != userID {
		return errorx.New(errorx.CodeForbidden, "forbidden")
	}
	return nil
}
