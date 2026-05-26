package logic

import (
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
)

type FeedbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedbackLogic {
	return &FeedbackLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *FeedbackLogic) Submit(req *types.FeedbackReq) (*types.FeedbackResp, error) {
	userID, _ := ctxdata.GetUserId(l.ctx)
	if userID <= 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	if strings.TrimSpace(req.SessionID) == "" || strings.TrimSpace(req.TraceID) == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "sessionId/traceId required")
	}
	if req.Score < -1 || req.Score > 1 {
		return nil, errorx.New(errorx.CodeBadRequest, "score must be -1/0/1")
	}
	now := time.Now().UnixMilli()
	_, err := l.svcCtx.Db.ExecCtx(l.ctx,
		"INSERT INTO agent_feedback (session_id,user_id,trace_id,score,comment,created_at) VALUES (?,?,?,?,?,?)",
		req.SessionID, userID, req.TraceID, req.Score, svc.TrimContent(req.Comment, 1000), now,
	)
	if err != nil {
		return nil, err
	}
	return &types.FeedbackResp{Accepted: true}, nil
}
