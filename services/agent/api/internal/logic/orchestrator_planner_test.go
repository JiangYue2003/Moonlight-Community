package logic

import (
	"context"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval"
)

func newPlannerTestOrchestrator() *Orchestrator {
	return &Orchestrator{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Config: config.Config{
				Agent: config.AgentConf{
					DefaultTopK: 20,
					MaxTopK:     50,
					Planner: config.PlannerConf{
						Enable:                   true,
						MaxSubQueries:            3,
						MaxTopicTerms:            8,
						QuestionRunesThreshold:   30,
						CompareKeywordThreshold:  1,
						ForcePlanQuestionRunes:   60,
						DefaultSubQueryTopK:      6,
						MaxSubQueryTopK:          12,
						UseProOnComplex:          true,
						StrictJSON:               true,
						ConstraintKeywordTrigger: 1,
					},
				},
			},
		},
	}
}

func TestShouldPlan(t *testing.T) {
	o := newPlannerTestOrchestrator()
	if o.shouldPlan("今天天气如何") {
		t.Fatal("simple question should not plan")
	}
	if !o.shouldPlan("请分别对比 A 和 B 的优缺点，并按性能、易用性分类说明") {
		t.Fatal("compare question should plan")
	}
}

func TestParsePlannerJSONFallback(t *testing.T) {
	o := newPlannerTestOrchestrator()
	_, err := o.parseRetrievalPlanFromJSON("not-json", "q")
	if err == nil {
		t.Fatal("invalid json should fail")
	}
}

func TestParsePlannerJSONValid(t *testing.T) {
	o := newPlannerTestOrchestrator()
	raw := `{"question_type":"compare","topics":["A","B"],"constraints":[],"sub_queries":[{"goal":"A","query":"A 的特点","channels":["vector","keyword"]},{"goal":"B","query":"B 的特点","channels":["vector"]}],"answer_strategy":"compare_and_classify"}`
	p, err := o.parseRetrievalPlanFromJSON(raw, "对比 A 和 B")
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if len(p.SubQueries) != 2 {
		t.Fatalf("want 2 subqueries got %d", len(p.SubQueries))
	}
}

func TestCompressPlannedEvidenceDedup(t *testing.T) {
	o := newPlannerTestOrchestrator()
	items := []retrieval.ScoredItem{
		{PostID: 1, ChunkID: "c1", Source: "s", Text: "t1", Score: 0.9},
		{PostID: 1, ChunkID: "c1", Source: "s", Text: "t1", Score: 0.8},
		{PostID: 1, ChunkID: "c2", Source: "s", Text: "t2", Score: 0.7},
	}
	out := o.compressPlannedEvidence(items, 10)
	if len(out) != 2 {
		t.Fatalf("want dedup=2 got=%d", len(out))
	}
}
