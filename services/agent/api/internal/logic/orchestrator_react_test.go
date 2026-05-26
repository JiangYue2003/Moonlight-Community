package logic

import (
	"context"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/retrieval"
)

func newReactTestOrchestrator() *Orchestrator {
	return &Orchestrator{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Config: config.Config{
				Agent: config.AgentConf{
					DefaultTopK: 20,
					MaxTopK:     50,
					React: config.ReactConf{
						Enable:                true,
						MaxSteps:              3,
						MaxElapsedMs:          5000,
						MaxSameQueryRepeats:   2,
						MaxRewriteWithoutGain: 1,
						MinNewEvidencePerStep: 1,
						MaxEvidencePool:       20,
						DefaultStepTopK:       5,
						StrictJSON:            true,
					},
				},
			},
		},
	}
}

func TestShouldUseReAct(t *testing.T) {
	o := newReactTestOrchestrator()
	if o.shouldUseReAct("你好") {
		t.Fatal("simple question should not use react")
	}
	if !o.shouldUseReAct("请先对比 Redis 和 Caffeine 的缓存一致性方案，再分别说明在 agent 模块中的适用场景和限制") {
		t.Fatal("complex compare question should use react")
	}
}

func TestParseReActActionInvalid(t *testing.T) {
	o := newReactTestOrchestrator()
	_, err := o.parseReActAction(`{"action":"hack","query":"x"}`)
	if err == nil {
		t.Fatal("invalid action should fail")
	}
}

func TestDetectQueryLoop(t *testing.T) {
	o := newReactTestOrchestrator()
	state := &ReActState{
		VisitedQueries: map[string]int{
			"redis cache": 2,
		},
	}
	if !o.detectQueryLoop(state, "redis cache") {
		t.Fatal("expected loop detection")
	}
}

func TestEvaluateEvidenceCoverageNoGain(t *testing.T) {
	o := newReactTestOrchestrator()
	state := &ReActState{
		EvidencePool: []retrieval.ScoredItem{
			{ChunkID: "c1", Text: "same", Score: 0.8},
		},
		StagnationCount: 1,
	}
	obs := ReActObservation{
		NewEvidence: []retrieval.ScoredItem{
			{ChunkID: "c1", Text: "same", Score: 0.8},
		},
	}
	result := o.evaluateEvidenceCoverage(state, "请对比 A 和 B", obs)
	if !result.NeedStop {
		t.Fatal("expected stagnation stop")
	}
}
