package observability

import (
	"testing"

	"github.com/cloudwego/eino/schema"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
)

func TestResolveUsageUsageSource(t *testing.T) {
	a := &AgentObservability{enable: true, cfg: normalizeCost(config.ModelCostConf{Currency: "USD", EstimateCharsPerToken: 4})}
	u, source := a.resolveUsage(ModelTelemetry{
		InputText:  "ignored",
		OutputText: "ignored",
		Usage: &Usage{
			PromptTokens:     10,
			CompletionTokens: 5,
			TotalTokens:      15,
		},
	})
	if source != "usage" || u.TotalTokens != 15 {
		t.Fatalf("source=%s usage=%+v", source, u)
	}
}

func TestResolveUsageEstimateSource(t *testing.T) {
	a := &AgentObservability{enable: true, cfg: normalizeCost(config.ModelCostConf{Currency: "USD", EstimateCharsPerToken: 4})}
	u, source := a.resolveUsage(ModelTelemetry{
		InputText:  "12345678",
		OutputText: "1234",
	})
	if source != "estimate" {
		t.Fatalf("source=%s", source)
	}
	if u.PromptTokens != 2 || u.CompletionTokens != 1 || u.TotalTokens != 3 {
		t.Fatalf("usage=%+v", u)
	}
}

func TestExtractUsageFromMessage(t *testing.T) {
	msg := &schema.Message{
		ResponseMeta: &schema.ResponseMeta{
			Usage: &schema.TokenUsage{
				PromptTokens:     7,
				CompletionTokens: 3,
				TotalTokens:      10,
			},
		},
	}
	u := ExtractUsageFromMessage(msg)
	if u == nil || u.TotalTokens != 10 {
		t.Fatalf("usage=%+v", u)
	}
}
