package logic

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/observability"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

type MemoryPinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type factCard struct {
	Subject    string  `json:"subject"`
	Predicate  string  `json:"predicate"`
	Object     string  `json:"object"`
	Confidence float64 `json:"confidence"`
	SourceRef  string  `json:"sourceRef"`
}

type preferenceCard struct {
	Kind       string  `json:"kind"`
	Content    string  `json:"content"`
	Confidence float64 `json:"confidence"`
	Source     string  `json:"source"`
}

func NewMemoryPinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemoryPinLogic {
	return &MemoryPinLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *MemoryPinLogic) Pin(req *types.MemoryPinReq) (*types.MemoryPinResp, error) {
	userID, _ := ctxdata.GetUserId(l.ctx)
	if userID <= 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	if strings.TrimSpace(req.SessionID) == "" || strings.TrimSpace(req.Content) == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "sessionId/content required")
	}
	now := time.Now().UnixMilli()
	content := svc.TrimContent(req.Content, 4000)
	_, err := l.svcCtx.Db.ExecCtx(l.ctx,
		"INSERT INTO agent_memory_pin (user_id,session_id,tag,content,created_at) VALUES (?,?,?,?,?)",
		userID, req.SessionID, svc.TrimContent(req.Tag, 64), content, now,
	)
	if err != nil {
		return nil, err
	}
	go l.extractFactsAsync(context.Background(), userID, req.SessionID, content, now)
	return &types.MemoryPinResp{Pinned: true}, nil
}

func (l *MemoryPinLogic) extractFactsAsync(ctx context.Context, userID int64, sessionID, content string, ts int64) {
	tctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	prompt := fmt.Sprintf("请将以下文本抽取为JSON对象，包含 facts 和 preferences 两个数组。facts 每项字段: subject,predicate,object,confidence(0-1),sourceRef。preferences 每项字段: kind,content,confidence(0-1),source，kind 仅允许 response_style、focus_area、working_preference；只提取稳定长期偏好，不要提取一次性要求；只输出JSON。文本：%s", content)
	msgs := []*schema.Message{
		schema.SystemMessage("你是知识与偏好抽取器。只输出合法JSON对象。"),
		schema.UserMessage(prompt),
	}
	decision := l.svcCtx.Router.Decide(tctx, svc.RouteScenarioFactExtract, svc.RouteInput{PinContent: content})
	l.Logger.Infof("agent model route %s", svc.RouteLogFields(sessionID, decision, false))
	callStart := time.Now()
	usedModel := decision.ModelName
	callStatus := "ok"
	resp, err := decision.Model.Generate(tctx, msgs)
	if err != nil || strings.TrimSpace(resp.Content) == "" {
		callStatus = "error"
		l.svcCtx.Router.RecordFailure(tctx, decision.ModelName)
		if decision.ModelName == "pro" && l.svcCtx.Router.RetryOnProFail() && l.svcCtx.ChatLite != nil {
			fb := svc.RouteDecision{Model: l.svcCtx.ChatLite, ModelName: "lite", Reason: "pro_generate_failed"}
			l.Logger.Infof("agent model route %s", svc.RouteLogFields(sessionID, fb, true))
			if l.svcCtx.Obs != nil && l.svcCtx.Obs.Enabled() {
				l.svcCtx.Obs.RecordFallback(string(svc.RouteScenarioFactExtract), decision.ModelName, fb.ModelName, fb.Reason)
			}
			callStart = time.Now()
			resp, err = fb.Model.Generate(tctx, msgs)
			usedModel = fb.ModelName
			if err == nil && strings.TrimSpace(resp.Content) != "" {
				callStatus = "fallback_ok"
			} else {
				callStatus = "fallback_error"
			}
		}
	}
	if err != nil || strings.TrimSpace(resp.Content) == "" {
		l.recordGenerateTelemetry(sessionID, content, "", usedModel, callStatus, callStart, observability.ExtractUsageFromMessage(resp))
		l.Logger.Errorf("extractFacts chat err=%v", err)
		return
	}
	l.recordGenerateTelemetry(sessionID, content, resp.Content, usedModel, callStatus, callStart, observability.ExtractUsageFromMessage(resp))
	cards, prefs, err := parseMemoryExtraction(resp.Content)
	if err != nil {
		l.Logger.Errorf("extractFacts parse err=%v", err)
		return
	}
	facts := make([]memory.Fact, 0, len(cards))
	embedTexts := make([]string, 0, len(cards))
	for _, c := range cards {
		if !validFact(c) {
			continue
		}
		now := time.Now().UnixMilli()
		f := memory.Fact{
			FactID:      factID(userID, c.Subject, c.Predicate, c.Object),
			Subject:     svc.TrimContent(c.Subject, 255),
			Predicate:   svc.TrimContent(c.Predicate, 255),
			ObjectValue: svc.TrimContent(c.Object, 2000),
			SourceRef:   svc.TrimContent(c.SourceRef, 255),
			Confidence:  c.Confidence,
			Version:     sessionID,
			Status:      "active",
			CreatedAt:   ts,
			UpdatedAt:   now,
		}
		facts = append(facts, f)
		embedTexts = append(embedTexts, formatFactForEmbedding(f))
	}
	if len(facts) == 0 {
		// continue to preference persistence if any
	}
	if len(facts) > 0 && l.svcCtx.MemoryFacts != nil {
		if err := l.svcCtx.MemoryFacts.UpsertFacts(tctx, userID, facts); err != nil {
			l.Logger.Errorf("extractFacts upsert mysql err=%v", err)
			return
		}
	}
	if len(facts) > 0 && l.svcCtx.MemoryVectors != nil && l.svcCtx.Milvus != nil && l.svcCtx.Config.Agent.EnableMilvus {
		vecs, err := llmx.EmbedFloat32(tctx, l.svcCtx.Embed, embedTexts)
		if err != nil || len(vecs) != len(facts) {
			if err == nil {
				err = fmt.Errorf("fact embed size mismatch")
			}
			l.Logger.Errorf("extractFacts embed err=%v", err)
			return
		}
		fvs := make([]memory.FactVector, 0, len(facts))
		for i := range facts {
			fvs = append(fvs, memory.FactVector{Fact: facts[i], Vector: vecs[i]})
		}
		if err := l.svcCtx.MemoryVectors.UpsertFactVectors(tctx, userID, fvs); err != nil {
			l.Logger.Errorf("extractFacts upsert milvus err=%v", err)
		}
	}
	preferences := make([]memory.Preference, 0, len(prefs))
	for _, p := range prefs {
		pref := memory.Preference{
			PreferenceID: preferenceID(userID, p.Kind, p.Content),
			Kind:         svc.TrimContent(p.Kind, 64),
			Content:      svc.TrimContent(p.Content, 512),
			Confidence:   p.Confidence,
			Source:       svc.TrimContent(p.Source, 64),
			Status:       "active",
			LastSeenAt:   ts,
			CreatedAt:    ts,
			UpdatedAt:    time.Now().UnixMilli(),
		}
		if !validPreference(pref) {
			continue
		}
		preferences = append(preferences, pref)
	}
	if len(preferences) > 0 && l.svcCtx.Preferences != nil {
		if err := l.svcCtx.Preferences.UpsertPreferences(tctx, userID, preferences); err != nil {
			l.Logger.Errorf("extractPreferences upsert mysql err=%v", err)
		}
	}
}

func parseFactCards(raw string) ([]factCard, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)
	var cards []factCard
	if err := json.Unmarshal([]byte(raw), &cards); err == nil {
		return cards, nil
	}
	var wrapped struct {
		Items []factCard `json:"items"`
	}
	if err := json.Unmarshal([]byte(raw), &wrapped); err == nil {
		return wrapped.Items, nil
	}
	return nil, fmt.Errorf("invalid fact json")
}

func parsePreferenceCards(raw string) ([]preferenceCard, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)
	var cards []preferenceCard
	if err := json.Unmarshal([]byte(raw), &cards); err == nil {
		return cards, nil
	}
	var wrapped struct {
		Items []preferenceCard `json:"items"`
	}
	if err := json.Unmarshal([]byte(raw), &wrapped); err == nil {
		return wrapped.Items, nil
	}
	return nil, fmt.Errorf("invalid preference json")
}

func parseMemoryExtraction(raw string) ([]factCard, []preferenceCard, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)
	var payload struct {
		Facts       []factCard       `json:"facts"`
		Preferences []preferenceCard `json:"preferences"`
	}
	if err := json.Unmarshal([]byte(raw), &payload); err == nil {
		return payload.Facts, payload.Preferences, nil
	}
	facts, ferr := parseFactCards(raw)
	if ferr == nil {
		return facts, nil, nil
	}
	return nil, nil, fmt.Errorf("invalid memory extraction json")
}

func validFact(c factCard) bool {
	if strings.TrimSpace(c.Subject) == "" || strings.TrimSpace(c.Predicate) == "" || strings.TrimSpace(c.Object) == "" {
		return false
	}
	if c.Confidence < 0.5 {
		return false
	}
	return true
}

func validPreference(p memory.Preference) bool {
	if strings.TrimSpace(p.Kind) == "" || strings.TrimSpace(p.Content) == "" {
		return false
	}
	if p.Confidence < 0.7 {
		return false
	}
	switch strings.TrimSpace(p.Kind) {
	case "response_style", "focus_area", "working_preference":
	default:
		return false
	}
	lower := strings.ToLower(strings.TrimSpace(p.Content))
	for _, bad := range []string{"这次", "本次", "先", "暂时", "临时"} {
		if strings.Contains(lower, bad) {
			return false
		}
	}
	return true
}

func factID(userID int64, s, p, o string) string {
	h := sha1.Sum([]byte(fmt.Sprintf("%d|%s|%s|%s", userID, s, p, o)))
	return hex.EncodeToString(h[:])
}

func preferenceID(userID int64, kind, content string) string {
	h := sha1.Sum([]byte(fmt.Sprintf("%d|%s|%s", userID, strings.TrimSpace(kind), strings.TrimSpace(content))))
	return hex.EncodeToString(h[:])
}

func formatFactForEmbedding(f memory.Fact) string {
	return strings.TrimSpace(f.Subject + " " + f.Predicate + " " + f.ObjectValue)
}

func (l *MemoryPinLogic) recordGenerateTelemetry(traceID, in, out, modelName, status string, startedAt time.Time, usage *observability.Usage) {
	if l.svcCtx == nil || l.svcCtx.Obs == nil || !l.svcCtx.Obs.Enabled() {
		return
	}
	u, cost, source := l.svcCtx.Obs.RecordModelCall(observability.ModelTelemetry{
		Scenario:   string(svc.RouteScenarioFactExtract),
		Model:      modelName,
		Method:     "generate",
		Status:     status,
		InputText:  in,
		OutputText: out,
		Usage:      usage,
	}, startedAt)
	l.Logger.Infof("agent model telemetry trace=%s scenario=fact_extract model=%s status=%s in_tokens=%d out_tokens=%d total_tokens=%d token_source=%s cost=%.8f",
		traceID, modelName, status, u.PromptTokens, u.CompletionTokens, u.TotalTokens, source, cost)
}
