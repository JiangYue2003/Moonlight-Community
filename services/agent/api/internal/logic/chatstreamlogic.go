package logic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/observability"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

type ChatStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatStreamLogic {
	return &ChatStreamLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *ChatStreamLogic) Run(w http.ResponseWriter, req *types.ChatStreamReq) error {
	if strings.TrimSpace(req.SessionID) == "" {
		return errorx.New(errorx.CodeBadRequest, "sessionId required")
	}
	userID, _ := ctxdata.GetUserId(l.ctx)
	if userID <= 0 {
		return errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	if err := l.mustOwnSession(userID, req.SessionID); err != nil {
		return err
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		return errorx.New(errorx.CodeInternalError, "streaming unsupported")
	}

	rl := l.svcCtx.Config.Agent.ChatRateLimit
	allowed, err := l.svcCtx.RateLimit.Take(l.ctx, "rl:agent:chat:"+strconv.FormatInt(userID, 10), rl.Capacity, rl.RefillPerSec)
	if err != nil {
		l.Logger.Errorf("agent ratelimit err=%v", err)
		allowed = true
	}

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	if !allowed {
		l.sendEvent(w, flusher, "final", map[string]any{"traceId": "", "answer": "请求过于频繁，请稍后再试。"})
		l.sendDone(w, flusher)
		return nil
	}

	question := strings.TrimSpace(req.Question)
	if question == "" {
		l.sendEvent(w, flusher, "final", map[string]any{"traceId": "", "answer": "问题不能为空。"})
		l.sendDone(w, flusher)
		return nil
	}

	l.sendEvent(w, flusher, "tool_call", map[string]any{"tool": "hybrid_retrieve", "sessionId": req.SessionID})
	orch := NewOrchestrator(l.ctx, l.svcCtx)
	plan, err := orch.Build(userID, req.SessionID, question, req.TopK)
	if err != nil {
		l.sendEvent(w, flusher, "final", map[string]any{"traceId": "", "answer": "检索失败，请稍后再试。"})
		l.sendDone(w, flusher)
		return nil
	}
	l.sendEvent(w, flusher, "tool_result", map[string]any{"tool": "hybrid_retrieve", "citations": len(plan.Citations), "traceId": plan.TraceID})

	summary, _ := l.svcCtx.Redis.Get(l.ctx, svc.SessionSummaryKey(userID, req.SessionID)).Result()
	preferences := l.activePreferences(userID)
	msgs := composeMessages(summary, plan.Prompt, preferences)
	sessionMsgCount := l.sessionMessageCount(userID, req.SessionID)
	decision := l.svcCtx.Router.Decide(l.ctx, svc.RouteScenarioChat, svc.RouteInput{
		Question:        question,
		Prompt:          plan.Prompt,
		Summary:         summary,
		RecallCount:     len(plan.Citations),
		SessionMsgCount: sessionMsgCount,
	})
	if l.svcCtx.Router.EmitRouteEvent() {
		l.sendEvent(w, flusher, "route", map[string]any{
			"traceId":   plan.TraceID,
			"model":     decision.ModelName,
			"reason":    decision.Reason,
			"scenario":  string(svc.RouteScenarioChat),
			"fallback":  false,
			"recallCnt": len(plan.Citations),
		})
	}
	l.Logger.Infof("agent model route %s", svc.RouteLogFields(plan.TraceID, decision, false))

	callStart := time.Now()
	reader, err := decision.Model.Stream(l.ctx, msgs)
	usedModel := decision.ModelName
	callStatus := "ok"
	if err != nil {
		callStatus = "error"
		l.svcCtx.Router.RecordFailure(l.ctx, decision.ModelName)
		if decision.ModelName == "pro" && l.svcCtx.Router.RetryOnProFail() && l.svcCtx.ChatLite != nil {
			fallbackDecision := svc.RouteDecision{Model: l.svcCtx.ChatLite, ModelName: "lite", Reason: "pro_stream_failed"}
			l.Logger.Errorf("agent pro stream failed, fallback to lite trace=%s err=%v", plan.TraceID, err)
			l.Logger.Infof("agent model route %s", svc.RouteLogFields(plan.TraceID, fallbackDecision, true))
			if l.svcCtx.Obs != nil && l.svcCtx.Obs.Enabled() {
				l.svcCtx.Obs.RecordFallback(string(svc.RouteScenarioChat), decision.ModelName, fallbackDecision.ModelName, fallbackDecision.Reason)
			}
			if l.svcCtx.Router.EmitRouteEvent() {
				l.sendEvent(w, flusher, "route", map[string]any{
					"traceId":  plan.TraceID,
					"model":    fallbackDecision.ModelName,
					"reason":   fallbackDecision.Reason,
					"scenario": string(svc.RouteScenarioChat),
					"fallback": true,
				})
			}
			callStart = time.Now()
			reader, err = fallbackDecision.Model.Stream(l.ctx, msgs)
			usedModel = fallbackDecision.ModelName
			if err == nil {
				callStatus = "fallback_ok"
			} else {
				callStatus = "fallback_error"
			}
		}
	}
	if err != nil {
		l.recordStreamTelemetry(plan.TraceID, question, "", usedModel, callStatus, callStart, nil)
		l.sendEvent(w, flusher, "final", map[string]any{"traceId": plan.TraceID, "answer": "模型暂不可用，请稍后再试。"})
		l.sendDone(w, flusher)
		return nil
	}
	defer reader.Close()

	var answerBuilder strings.Builder
	var lastMsg *schema.Message
	for {
		msg, recvErr := reader.Recv()
		if recvErr != nil {
			if !errorsIsEOF(recvErr) {
				callStatus = "error"
			}
			break
		}
		lastMsg = msg
		if err != nil {
			break
		}
		if msg == nil || msg.Content == "" {
			continue
		}
		answerBuilder.WriteString(msg.Content)
		l.sendEvent(w, flusher, "token", map[string]any{"traceId": plan.TraceID, "content": msg.Content})
	}

	answer := strings.TrimSpace(answerBuilder.String())
	if answer == "" {
		answer = "未生成有效回答，请稍后重试。"
	}
	for _, c := range plan.Citations {
		l.sendEvent(w, flusher, "citation", c)
	}
	l.recordStreamTelemetry(plan.TraceID, question, answer, usedModel, callStatus, callStart, observability.ExtractUsageFromMessage(lastMsg))
	l.sendEvent(w, flusher, "final", map[string]any{"traceId": plan.TraceID, "answer": answer})
	l.sendDone(w, flusher)

	now := time.Now().UnixMilli()
	_ = l.appendHistory(userID, req.SessionID, "user", question, now)
	_ = l.appendHistory(userID, req.SessionID, "assistant", answer, now)
	_ = l.compactSession(userID, req.SessionID)
	_, _ = l.svcCtx.Db.ExecCtx(l.ctx, "UPDATE agent_sessions SET updated_at=? WHERE session_id=? AND user_id=?", now, req.SessionID, userID)
	return nil
}

func (l *ChatStreamLogic) activePreferences(userID int64) []memory.Preference {
	if l.svcCtx == nil || l.svcCtx.Preferences == nil || userID <= 0 {
		return nil
	}
	items, err := l.svcCtx.Preferences.ListActivePreferences(l.ctx, userID, 3)
	if err != nil {
		l.Logger.Errorf("agent preference load failed user=%d err=%v", userID, err)
		return nil
	}
	return items
}

func (l *ChatStreamLogic) mustOwnSession(userID int64, sessionID string) error {
	var uid int64
	if err := l.svcCtx.Db.QueryRowCtx(l.ctx, &uid, "SELECT user_id FROM agent_sessions WHERE session_id=? LIMIT 1", sessionID); err != nil {
		return errorx.New(errorx.CodeNotFound, "session not found")
	}
	if uid != userID {
		return errorx.New(errorx.CodeForbidden, "forbidden")
	}
	return nil
}

func (l *ChatStreamLogic) appendHistory(userID int64, sessionID, role, content string, ts int64) error {
	key := svc.SessionMessagesKey(userID, sessionID)
	entry := svc.MarshalHistory(role, svc.TrimContent(content, 2000), ts)
	if err := l.svcCtx.Redis.RPush(l.ctx, key, entry).Err(); err != nil {
		return err
	}
	ttl := time.Duration(l.svcCtx.Config.Agent.SessionTTLHours) * time.Hour
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return l.svcCtx.Redis.Expire(l.ctx, key, ttl).Err()
}

func (l *ChatStreamLogic) compactSession(userID int64, sessionID string) error {
	key := svc.SessionMessagesKey(userID, sessionID)
	maxItems := int64(24)
	if err := l.svcCtx.Redis.LTrim(l.ctx, key, -maxItems, -1).Err(); err != nil {
		return err
	}
	vals, err := l.svcCtx.Redis.LRange(l.ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}
	if len(vals) < int(maxItems) {
		return nil
	}
	var msgs []map[string]any
	for _, v := range vals {
		var m map[string]any
		if json.Unmarshal([]byte(v), &m) == nil {
			msgs = append(msgs, m)
		}
	}
	if len(msgs) == 0 {
		return nil
	}
	last := msgs[len(msgs)-1]
	summary := "会话摘要：最近讨论了 " + toString(last["content"])
	sKey := svc.SessionSummaryKey(userID, sessionID)
	ttl := time.Duration(l.svcCtx.Config.Agent.SummaryTTLHours) * time.Hour
	if ttl <= 0 {
		ttl = 72 * time.Hour
	}
	if err := l.svcCtx.Redis.Set(l.ctx, sKey, svc.TrimContent(summary, 500), ttl).Err(); err != nil {
		return err
	}
	return nil
}

func toString(v any) string {
	s, _ := v.(string)
	return strings.TrimSpace(s)
}

func (l *ChatStreamLogic) sendEvent(w http.ResponseWriter, f http.Flusher, name string, payload any) {
	b, _ := json.Marshal(payload)
	w.Write([]byte("event: " + name + "\n"))
	w.Write([]byte("data: "))
	w.Write(b)
	w.Write([]byte("\n\n"))
	f.Flush()
}

func (l *ChatStreamLogic) sendDone(w http.ResponseWriter, f http.Flusher) {
	w.Write([]byte("data: [DONE]\n\n"))
	f.Flush()
}

func (l *ChatStreamLogic) sessionMessageCount(userID int64, sessionID string) int {
	vals, err := l.svcCtx.Redis.LRange(l.ctx, svc.SessionMessagesKey(userID, sessionID), 0, -1).Result()
	if err != nil {
		return 0
	}
	return len(vals)
}

func (l *ChatStreamLogic) recordStreamTelemetry(traceID, in, out, modelName, status string, startedAt time.Time, usage *observability.Usage) {
	if l.svcCtx == nil || l.svcCtx.Obs == nil || !l.svcCtx.Obs.Enabled() {
		return
	}
	u, cost, source := l.svcCtx.Obs.RecordModelCall(observability.ModelTelemetry{
		Scenario:   string(svc.RouteScenarioChat),
		Model:      modelName,
		Method:     "stream",
		Status:     status,
		InputText:  in,
		OutputText: out,
		Usage:      usage,
	}, startedAt)
	l.Logger.Infof("agent model telemetry trace=%s scenario=chat model=%s status=%s in_tokens=%d out_tokens=%d total_tokens=%d token_source=%s cost=%.8f",
		traceID, modelName, status, u.PromptTokens, u.CompletionTokens, u.TotalTokens, source, cost)
}

func errorsIsEOF(err error) bool { return err == io.EOF }

var _ = schema.AssistantMessage
