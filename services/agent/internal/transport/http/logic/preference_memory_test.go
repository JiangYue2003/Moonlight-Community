package logic

import (
	"strings"
	"testing"

	"github.com/cloudwego/eino/schema"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

func TestParsePreferenceCards(t *testing.T) {
	raw := `[
		{"kind":"response_style","content":"回答偏扁平叙述，少分点","confidence":0.9,"source":"explicit_pin"},
		{"kind":"focus_area","content":"当前主要关注 agent 模块","confidence":0.8,"source":"inferred_from_dialog"}
	]`
	items, err := parsePreferenceCards(raw)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items got %d", len(items))
	}
}

func TestValidPreferenceRejectsTemporaryInstruction(t *testing.T) {
	item := memory.Preference{
		Kind:       "response_style",
		Content:    "这次简单点",
		Confidence: 0.9,
		Source:     "inferred_from_dialog",
	}
	if validPreference(item) {
		t.Fatal("temporary instruction should not be persisted")
	}
}

func TestComposeMessagesInjectsPreferences(t *testing.T) {
	prefs := []memory.Preference{
		{Kind: "response_style", Content: "回答偏扁平叙述，少分点", Confidence: 0.9},
		{Kind: "working_preference", Content: "优先给代码实现，不要只讲产品描述", Confidence: 0.8},
	}
	msgs := composeMessages("summary", "prompt", prefs)
	if len(msgs) < 4 {
		t.Fatalf("want preference system messages, got %d", len(msgs))
	}
	var systemTexts []string
	for _, msg := range msgs {
		if msg.Role == schema.System {
			systemTexts = append(systemTexts, msg.Content)
		}
	}
	joined := strings.Join(systemTexts, "\n")
	if !strings.Contains(joined, "扁平叙述") {
		t.Fatal("missing response_style preference")
	}
	if !strings.Contains(joined, "代码实现") {
		t.Fatal("missing working_preference preference")
	}
}
