package logic

import (
	"context"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
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

func TestParseReActActionRewriteSameQueryInvalid(t *testing.T) {
	o := newReactTestOrchestrator()
	_, err := o.validateReActAction(&ReActAction{
		Action: "rewrite_query",
		Query:  "redis cache",
	}, "redis cache")
	if err == nil {
		t.Fatal("rewrite to same query should fail")
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

func TestFallbackHeuristicAction(t *testing.T) {
	o := newReactTestOrchestrator()
	action := o.fallbackHeuristicAction(&ReActState{}, "请对比 Redis 和 Caffeine", 5)
	if action == nil || action.Action != "search_knowledge" {
		t.Fatal("expected heuristic search action")
	}
}

func TestExecuteReactActionMemoryPreferences(t *testing.T) {
	o := newReactTestOrchestrator()
	o.svcCtx.Preferences = &stubPreferenceStore{
		items: []memory.Preference{
			{PreferenceID: "p1", Kind: "response_style", Content: "偏扁平叙述", Confidence: 0.9, Status: "active"},
		},
	}
	bundle, err := o.executeReactAction(context.Background(), 1, "s1", "t1", &ReActAction{
		Action: "search_memory_preferences",
		Query:  "回答偏好",
		TopK:   3,
	})
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	if len(bundle.Merged) != 1 || bundle.Merged[0].Source != "memory_preference" {
		t.Fatal("expected preference retrieval bundle")
	}
}

func TestMapMemoryPreferencesShouldNotBecomeCitations(t *testing.T) {
	o := newReactTestOrchestrator()
	o.svcCtx.Preferences = &stubPreferenceStore{
		items: []memory.Preference{
			{PreferenceID: "p1", Kind: "response_style", Content: "偏扁平叙述", Confidence: 0.9, Status: "active"},
		},
	}
	items, err := o.searchMemoryPreferences(context.Background(), 1, 3)
	if err != nil {
		t.Fatalf("search failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("want 1 item got %d", len(items))
	}
	if items[0].PostID != 0 {
		t.Fatal("preference memory should not map to article citation post ids")
	}
}

func TestEvaluateEvidenceCoverageShouldNotStopOnFirstUsefulStep(t *testing.T) {
	o := newReactTestOrchestrator()
	state := &ReActState{
		MaxSteps: 3,
	}
	obs := ReActObservation{
		NewEvidence: []retrieval.ScoredItem{
			{ChunkID: "c1", Text: "new", Score: 0.9},
		},
	}
	result := o.evaluateEvidenceCoverage(state, "请对比 A 和 B", obs)
	if result.NeedStop {
		t.Fatal("should not stop after first useful step")
	}
}

func TestFallbackHeuristicActionOnControllerFailure(t *testing.T) {
	o := newReactTestOrchestrator()
	state := &ReActState{CurrentQuery: "redis cache"}
	action := o.fallbackHeuristicAction(state, "redis cache", 5)
	if action.Query != "redis cache" || action.Action != "search_knowledge" {
		t.Fatal("unexpected fallback heuristic action")
	}
}

type stubPreferenceStore struct {
	items []memory.Preference
}

func (s *stubPreferenceStore) UpsertPreferences(ctx context.Context, userID int64, prefs []memory.Preference) error {
	return nil
}

func (s *stubPreferenceStore) ListActivePreferences(ctx context.Context, userID int64, limit int) ([]memory.Preference, error) {
	if limit > 0 && len(s.items) > limit {
		return s.items[:limit], nil
	}
	return s.items, nil
}
