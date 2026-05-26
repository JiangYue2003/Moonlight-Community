package logic

import (
	"context"
	"time"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/tooling"
)

type ToolAuditStore struct {
	dbSvc *svc.ServiceContext
}

func NewToolAuditStore(sc *svc.ServiceContext) *ToolAuditStore {
	return &ToolAuditStore{dbSvc: sc}
}

func (s *ToolAuditStore) LogToolCall(ctx context.Context, rec tooling.AuditRecord) {
	now := time.Now().UnixMilli()
	_, _ = s.dbSvc.Db.ExecCtx(ctx,
		"INSERT INTO agent_tool_audit (session_id,trace_id,user_id,tool_name,params_hash,latency_ms,status,err_msg,created_at) VALUES (?,?,?,?,?,?,?,?,?)",
		rec.SessionID, rec.TraceID, rec.UserID, rec.Tool, rec.ParamsHash, rec.LatencyMs, rec.Status, svc.TrimContent(rec.ErrMsg, 255), now,
	)
}
