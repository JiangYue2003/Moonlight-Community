package logic

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval/providers"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/security"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/tooling"
)

type Citation struct {
	PostID  int64  `json:"postId"`
	ChunkID string `json:"chunkId"`
	Source  string `json:"source"`
}

type ChatPlan struct {
	Prompt             string
	Citations          []Citation
	TraceID            string
	Planned            bool
	QuestionType       string
	SubQueryCount      int
	PlanModel          string
	PlanFallbackReason string
}

type RetrievalBundle struct {
	Question string
	TopK     int
	Intent   string
	Milvus   []retrieval.ScoredItem
	VectorES []retrieval.ScoredItem
	Keyword  []retrieval.ScoredItem
	Memory   []retrieval.ScoredItem
	Graph    []retrieval.ScoredItem
	Merged   []retrieval.ScoredItem
}

type Orchestrator struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	tools *tooling.Registry

	esVector *providers.ESVectorProvider
	esBM25   *providers.ESKeywordProvider
	milvus   *providers.MilvusProvider
	neo4j    *providers.Neo4jProvider
}

type RetrievalPlan struct {
	QuestionType   string         `json:"question_type"`
	Topics         []string       `json:"topics"`
	Constraints    []string       `json:"constraints"`
	SubQueries     []PlanSubQuery `json:"sub_queries"`
	AnswerStrategy string         `json:"answer_strategy"`
}

type PlanSubQuery struct {
	Goal     string   `json:"goal"`
	Query    string   `json:"query"`
	Channels []string `json:"channels"`
}

func NewOrchestrator(ctx context.Context, svcCtx *svc.ServiceContext) *Orchestrator {
	toolWL := svcCtx.Config.Agent.ToolWhitelist
	if len(toolWL) == 0 {
		toolWL = []string{"hybrid_retrieve", "search_vector_milvus", "search_vector_es", "search_bm25", "search_memory", "search_graph", "fuse_rrf"}
	}
	o := &Orchestrator{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		esVector: providers.NewESVectorProvider(svcCtx.Es, svcCtx.Config.KnowledgeIndex),
		esBM25:   providers.NewESKeywordProvider(svcCtx.Es, svcCtx.Config.KnowledgeIndex),
		milvus:   providers.NewMilvusProvider(svcCtx.Config.Agent.EnableMilvus, svcCtx.Milvus, svcCtx.Config.Milvus.Collection, svcCtx.Config.Milvus.VectorField),
		neo4j:    providers.NewNeo4jProvider(svcCtx.Config.Agent.EnableGraph),
	}
	o.tools = tooling.NewRegistry(toolWL, NewToolAuditStore(svcCtx))
	o.registerTools()
	return o
}

func (o *Orchestrator) registerTools() {
	commonSchema := tooling.Schema{
		"question": {Type: tooling.FieldString, Required: true, MaxLen: o.svcCtx.Config.Agent.MaxQuestionRunes},
		"topK":     {Type: tooling.FieldInt, Required: true, MinInt: 1, MaxInt: o.svcCtx.Config.Agent.MaxTopK},
	}
	o.tools.Register(tooling.Tool{
		Name:   "search_vector_milvus",
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			q := c.Params["question"].(string)
			topK, _ := asInt(c.Params["topK"])
			vecs, err := llmx.EmbedFloat32(ctx, o.svcCtx.Embed, []string{q})
			if err != nil || len(vecs) == 0 {
				return nil, fmt.Errorf("embed failed")
			}
			return o.milvus.Search(ctx, retrieval.Query{UserID: c.UserID, TopK: topK, Text: q, Vector: vecs[0]})
		},
	})
	o.tools.Register(tooling.Tool{
		Name:   "search_vector_es",
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			q := c.Params["question"].(string)
			topK, _ := asInt(c.Params["topK"])
			vecs, err := llmx.EmbedFloat32(ctx, o.svcCtx.Embed, []string{q})
			if err != nil || len(vecs) == 0 {
				return nil, fmt.Errorf("embed failed")
			}
			return o.esVector.Search(ctx, retrieval.Query{UserID: c.UserID, TopK: topK, Text: q, Vector: vecs[0]})
		},
	})
	o.tools.Register(tooling.Tool{
		Name:   "search_bm25",
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			q := c.Params["question"].(string)
			topK, _ := asInt(c.Params["topK"])
			return o.esBM25.Search(ctx, retrieval.Query{UserID: c.UserID, TopK: topK, Text: q})
		},
	})
	o.tools.Register(tooling.Tool{
		Name:   "search_memory",
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			if o.svcCtx.MemoryFacts == nil {
				return []retrieval.ScoredItem{}, nil
			}
			q := c.Params["question"].(string)
			topK, _ := asInt(c.Params["topK"])
			vecs, err := llmx.EmbedFloat32(ctx, o.svcCtx.Embed, []string{q})
			if err != nil || len(vecs) == 0 {
				return nil, fmt.Errorf("embed failed")
			}
			items, err := o.svcCtx.MemoryFacts.SearchFacts(ctx, memory.Query{
				UserID: c.UserID,
				TopK:   topK,
				Text:   q,
				Vector: vecs[0],
			})
			if err != nil {
				return nil, err
			}
			return mapMemoryFacts(items), nil
		},
	})
	o.tools.Register(tooling.Tool{
		Name: "search_graph",
		Enabled: func() bool {
			return o.neo4j.Enabled()
		},
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			q := c.Params["question"].(string)
			topK, _ := asInt(c.Params["topK"])
			return o.neo4j.Search(ctx, retrieval.Query{UserID: c.UserID, TopK: topK, Text: q})
		},
	})
	o.tools.Register(tooling.Tool{
		Name: "fuse_rrf",
		Schema: tooling.Schema{
			"k": {Type: tooling.FieldInt, Required: true, MinInt: 1, MaxInt: 200},
		},
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			k, _ := asInt(c.Params["k"])
			return k, nil
		},
	})
	o.tools.Register(tooling.Tool{
		Name:   "hybrid_retrieve",
		Schema: commonSchema,
		Run: func(ctx context.Context, c tooling.Call) (any, error) {
			q := strings.TrimSpace(c.Params["question"].(string))
			topK, _ := asInt(c.Params["topK"])
			return o.hybridRetrieve(ctx, c.UserID, c.SessionID, c.TraceID, q, topK)
		},
	})
}

func (o *Orchestrator) Build(userID int64, sessionID, question string, topK int) (*ChatPlan, error) {
	if err := security.EnsureUserScope(userID); err != nil {
		return nil, errorx.New(errorx.CodeUnauthorized, "unauthorized")
	}
	if err := security.ValidateQueryInput(question, o.svcCtx.Config.Agent.MaxQuestionRunes); err != nil {
		return nil, errorx.New(errorx.CodeBadRequest, err.Error())
	}
	topK = security.ClampTopK(topK, o.svcCtx.Config.Agent.DefaultTopK, 1, o.svcCtx.Config.Agent.MaxTopK)
	traceID := uuid.NewString()

	plan := &ChatPlan{TraceID: traceID}
	var bundle *RetrievalBundle
	var err error
	if o.shouldPlan(question) {
		var retrievalPlan *RetrievalPlan
		retrievalPlan, err = o.buildRetrievalPlan(o.ctx, strings.TrimSpace(question))
		if err == nil && retrievalPlan != nil {
			bundle, err = o.executePlannedRetrieval(o.ctx, userID, sessionID, traceID, topK, retrievalPlan)
			if err == nil {
				plan.Planned = true
				plan.QuestionType = retrievalPlan.QuestionType
				plan.SubQueryCount = len(retrievalPlan.SubQueries)
				plan.PlanModel = o.planModelName(question)
			}
		}
		if err != nil {
			plan.PlanFallbackReason = err.Error()
		}
	}
	if bundle == nil {
		ret, rerr := o.tools.Execute(o.ctx, tooling.Call{
			Tool:      "hybrid_retrieve",
			UserID:    userID,
			SessionID: sessionID,
			TraceID:   traceID,
			Params: map[string]any{
				"question": strings.TrimSpace(question),
				"topK":     topK,
			},
		})
		if rerr != nil {
			return nil, errorx.Wrap(errorx.CodeInternalError, "hybrid retrieve failed", rerr)
		}
		var ok bool
		bundle, ok = ret.(*RetrievalBundle)
		if !ok || bundle == nil {
			return nil, errorx.New(errorx.CodeInternalError, "invalid retrieval bundle")
		}
	}

	if len(bundle.Merged) == 0 {
		plan.Prompt = fmt.Sprintf("用户问题：%s\n\n未检索到可用知识，请明确说明未找到答案并建议用户补充收藏内容。", strings.TrimSpace(question))
		return plan, nil
	}

	var b strings.Builder
	b.WriteString("你是个人知识助手。只基于给定上下文回答，不要编造。\n")
	if plan.Planned {
		b.WriteString("问题类型：")
		b.WriteString(plan.QuestionType)
		b.WriteString("\n")
	}
	b.WriteString("问题：")
	b.WriteString(strings.TrimSpace(question))
	b.WriteString("\n\n上下文：\n")
	for i, it := range bundle.Merged {
		fmt.Fprintf(&b, "[%d] (post=%d chunk=%s source=%s) %s\n", i+1, it.PostID, it.ChunkID, it.Source, strings.TrimSpace(it.Text))
		plan.Citations = append(plan.Citations, Citation{PostID: it.PostID, ChunkID: it.ChunkID, Source: it.Source})
	}
	plan.Prompt = b.String()
	return plan, nil
}

func (o *Orchestrator) shouldPlan(question string) bool {
	cfg := o.svcCtx.Config.Agent.Planner
	if !cfg.Enable {
		return false
	}
	q := strings.TrimSpace(question)
	runes := len([]rune(q))
	if runes >= cfg.ForcePlanQuestionRunes {
		return true
	}
	compareHits := keywordHits(q, []string{"对比", "比较", "分别", "哪些", "分类", "同时", "异步", "同步"})
	if compareHits >= cfg.CompareKeywordThreshold {
		return true
	}
	constraintHits := keywordHits(q, []string{"只看", "仅看", "限定", "范围", "时间", "收藏"})
	if constraintHits >= cfg.ConstraintKeywordTrigger {
		return true
	}
	return runes >= cfg.QuestionRunesThreshold
}

func (o *Orchestrator) planModelName(question string) string {
	if o.svcCtx.Config.Agent.Planner.UseProOnComplex && o.shouldPlan(question) {
		return "pro"
	}
	return "lite"
}

func (o *Orchestrator) buildRetrievalPlan(ctx context.Context, question string) (*RetrievalPlan, error) {
	_ = ctx
	// v1: heuristic JSON plan, keeps path deterministic and fallback-safe.
	p := &RetrievalPlan{
		QuestionType:   "fact",
		Topics:         []string{question},
		Constraints:    nil,
		AnswerStrategy: "direct",
		SubQueries: []PlanSubQuery{
			{Goal: "answer", Query: question, Channels: []string{"vector", "keyword", "memory"}},
		},
	}
	if keywordHits(question, []string{"对比", "比较", "分别", "分类"}) > 0 {
		p.QuestionType = "compare"
		p.AnswerStrategy = "compare_and_classify"
	}
	return o.validateRetrievalPlan(p, question)
}

func (o *Orchestrator) parseRetrievalPlanFromJSON(raw string, question string) (*RetrievalPlan, error) {
	var p RetrievalPlan
	if err := json.Unmarshal([]byte(strings.TrimSpace(raw)), &p); err != nil {
		return nil, fmt.Errorf("planner_json_parse_failed")
	}
	return o.validateRetrievalPlan(&p, question)
}

func (o *Orchestrator) validateRetrievalPlan(p *RetrievalPlan, question string) (*RetrievalPlan, error) {
	if p == nil {
		return nil, fmt.Errorf("planner_nil")
	}
	cfg := o.svcCtx.Config.Agent.Planner
	if strings.TrimSpace(p.QuestionType) == "" {
		p.QuestionType = "fact"
	}
	switch p.QuestionType {
	case "fact", "compare", "classify", "aggregate", "multi_constraint":
	default:
		return nil, fmt.Errorf("planner_invalid_question_type")
	}
	switch p.AnswerStrategy {
	case "direct", "compare_and_classify", "group_and_summarize":
	default:
		return nil, fmt.Errorf("planner_invalid_answer_strategy")
	}
	if len(p.Topics) > cfg.MaxTopicTerms {
		p.Topics = p.Topics[:cfg.MaxTopicTerms]
	}
	if len(p.SubQueries) == 0 {
		p.SubQueries = []PlanSubQuery{{Goal: "answer", Query: question, Channels: []string{"vector", "keyword", "memory"}}}
	}
	if len(p.SubQueries) > cfg.MaxSubQueries {
		return nil, fmt.Errorf("planner_subqueries_exceed")
	}
	for i := range p.SubQueries {
		p.SubQueries[i].Query = strings.TrimSpace(p.SubQueries[i].Query)
		if p.SubQueries[i].Query == "" {
			return nil, fmt.Errorf("planner_empty_subquery")
		}
		if len(p.SubQueries[i].Channels) == 0 {
			p.SubQueries[i].Channels = []string{"vector", "keyword", "memory"}
		}
		for _, ch := range p.SubQueries[i].Channels {
			if !isAllowedPlannerChannel(ch) {
				return nil, fmt.Errorf("planner_invalid_channel")
			}
		}
	}
	return p, nil
}

func (o *Orchestrator) executePlannedRetrieval(ctx context.Context, userID int64, sessionID, traceID string, topK int, p *RetrievalPlan) (*RetrievalBundle, error) {
	subTopK := o.svcCtx.Config.Agent.Planner.DefaultSubQueryTopK
	if subTopK <= 0 {
		subTopK = (topK + 1) / 2
	}
	if subTopK > o.svcCtx.Config.Agent.Planner.MaxSubQueryTopK {
		subTopK = o.svcCtx.Config.Agent.Planner.MaxSubQueryTopK
	}
	if subTopK <= 0 {
		subTopK = 1
	}
	var all []retrieval.ScoredItem
	for _, sq := range p.SubQueries {
		b, err := o.hybridRetrieve(ctx, userID, sessionID, traceID, sq.Query, subTopK)
		if err != nil {
			o.Logger.Errorf("planner subquery retrieve failed: %v", err)
			continue
		}
		all = append(all, b.Merged...)
	}
	merged := o.compressPlannedEvidence(all, topK)
	return &RetrievalBundle{
		Question: p.SubQueries[0].Query,
		TopK:     topK,
		Intent:   security.GuessIntent(p.SubQueries[0].Query),
		Merged:   merged,
	}, nil
}

func (o *Orchestrator) compressPlannedEvidence(items []retrieval.ScoredItem, limit int) []retrieval.ScoredItem {
	if len(items) == 0 {
		return nil
	}
	sort.SliceStable(items, func(i, j int) bool { return items[i].Score > items[j].Score })
	seen := make(map[string]struct{}, len(items))
	postQuota := map[int64]int{}
	out := make([]retrieval.ScoredItem, 0, min(limit, len(items)))
	maxPerPost := 3
	for _, it := range items {
		key := dedupKey(it)
		if _, ok := seen[key]; ok {
			continue
		}
		if postQuota[it.PostID] >= maxPerPost {
			continue
		}
		seen[key] = struct{}{}
		postQuota[it.PostID]++
		out = append(out, it)
		if len(out) >= limit {
			break
		}
	}
	return out
}

func dedupKey(it retrieval.ScoredItem) string {
	chunk := strings.TrimSpace(it.ChunkID)
	if chunk != "" {
		return "chunk:" + chunk
	}
	h := sha1.Sum([]byte(strings.TrimSpace(it.Source) + "|" + fmt.Sprintf("%d", it.PostID) + "|" + strings.TrimSpace(it.Text)))
	return "fallback:" + hex.EncodeToString(h[:])
}

func keywordHits(q string, kws []string) int {
	n := 0
	for _, kw := range kws {
		if strings.Contains(q, kw) {
			n++
		}
	}
	return n
}

func isAllowedPlannerChannel(ch string) bool {
	switch strings.TrimSpace(ch) {
	case "vector", "keyword", "memory", "graph":
		return true
	default:
		return false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (o *Orchestrator) hybridRetrieve(ctx context.Context, userID int64, sessionID, traceID, question string, topK int) (*RetrievalBundle, error) {
	intent := security.GuessIntent(question)
	type ret struct {
		name string
		hits []retrieval.ScoredItem
		err  error
	}
	ch := make(chan ret, 4)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		r, e := o.tools.Execute(ctx, tooling.Call{Tool: "search_vector_milvus", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"question": question, "topK": topK}})
		h, _ := r.([]retrieval.ScoredItem)
		ch <- ret{name: "milvus", hits: h, err: e}
	}()
	go func() {
		defer wg.Done()
		r, e := o.tools.Execute(ctx, tooling.Call{Tool: "search_vector_es", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"question": question, "topK": topK}})
		h, _ := r.([]retrieval.ScoredItem)
		ch <- ret{name: "es", hits: h, err: e}
	}()
	go func() {
		defer wg.Done()
		r, e := o.tools.Execute(ctx, tooling.Call{Tool: "search_bm25", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"question": question, "topK": topK}})
		h, _ := r.([]retrieval.ScoredItem)
		ch <- ret{name: "bm25", hits: h, err: e}
	}()
	go func() {
		defer wg.Done()
		r, e := o.tools.Execute(ctx, tooling.Call{Tool: "search_memory", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"question": question, "topK": topK}})
		h, _ := r.([]retrieval.ScoredItem)
		ch <- ret{name: "memory", hits: h, err: e}
	}()
	wg.Wait()
	close(ch)

	var milvusHits, vectorHits, keywordHits, memoryHits []retrieval.ScoredItem
	for r := range ch {
		if r.err != nil {
			o.Logger.Errorf("hybrid retrieve tool %s failed: %v", r.name, r.err)
			continue
		}
		switch r.name {
		case "milvus":
			milvusHits = r.hits
		case "es":
			vectorHits = r.hits
		case "bm25":
			keywordHits = r.hits
		case "memory":
			memoryHits = r.hits
		}
	}

	var graphHits []retrieval.ScoredItem
	if intent == "relation" && o.neo4j.Enabled() {
		if gr, gerr := o.tools.Execute(ctx, tooling.Call{Tool: "search_graph", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"question": question, "topK": topK}}); gerr == nil {
			graphHits, _ = gr.([]retrieval.ScoredItem)
		}
	}

	_, _ = o.tools.Execute(ctx, tooling.Call{Tool: "fuse_rrf", UserID: userID, SessionID: sessionID, TraceID: traceID, Params: map[string]any{"k": o.svcCtx.Config.Agent.RRFK}})
	merged := retrieval.FuseRRF(o.svcCtx.Config.Agent.RRFK, milvusHits, vectorHits, keywordHits, memoryHits, graphHits)
	if len(merged) > topK {
		merged = merged[:topK]
	}
	return &RetrievalBundle{
		Question: question,
		TopK:     topK,
		Intent:   intent,
		Milvus:   milvusHits,
		VectorES: vectorHits,
		Keyword:  keywordHits,
		Memory:   memoryHits,
		Graph:    graphHits,
		Merged:   merged,
	}, nil
}

func userIDFromCtx(ctx context.Context) int64 {
	uid, _ := ctxdata.GetUserId(ctx)
	return uid
}

func nowMs() int64 { return time.Now().UnixMilli() }

func composeMessages(summary, prompt string) []*schema.Message {
	msgs := []*schema.Message{schema.SystemMessage("你是中文个人知识助手，回答必须简洁并引用上下文。")}
	if strings.TrimSpace(summary) != "" {
		msgs = append(msgs, schema.SystemMessage("会话摘要："+summary))
	}
	msgs = append(msgs, schema.UserMessage(prompt))
	return msgs
}

func asInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	case float64:
		return int(n), true
	default:
		return 0, false
	}
}

func mapMemoryFacts(items []memory.ScoredFact) []retrieval.ScoredItem {
	if len(items) == 0 {
		return nil
	}
	out := make([]retrieval.ScoredItem, 0, len(items))
	for i, it := range items {
		text := strings.TrimSpace(it.Fact.Subject + " " + it.Fact.Predicate + " " + it.Fact.ObjectValue)
		if text == "" {
			continue
		}
		out = append(out, retrieval.ScoredItem{
			DocID:   "fact:" + it.Fact.FactID,
			PostID:  0,
			ChunkID: it.Fact.FactID,
			Text:    text,
			Source:  "memory_fact",
			Score:   it.Score,
			Rank:    i + 1,
		})
	}
	return out
}
