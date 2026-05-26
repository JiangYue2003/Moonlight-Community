package observability

import (
	"math"
	"strings"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/metric"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
)

type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type ModelTelemetry struct {
	Scenario      string
	Model         string
	Method        string
	Status        string
	InputText     string
	OutputText    string
	Usage         *Usage
	Reason        string
	Fallback      bool
	FallbackFrom  string
	FallbackTo    string
	CircuitOpened bool
}

type AgentObservability struct {
	enable bool

	cfg config.ModelCostConf

	routeTotal       metric.CounterVec
	fallbackTotal    metric.CounterVec
	circuitOpenTotal metric.CounterVec
	callDuration     metric.HistogramVec
	tokenTotal       metric.CounterVec
	costTotal        metric.CounterVec
}

func NewAgentObservability(obs config.ObservabilityConf, cost config.ModelCostConf) *AgentObservability {
	a := &AgentObservability{
		enable: obs.Enable,
		cfg:    normalizeCost(cost),
	}
	if !a.enable {
		return a
	}

	a.routeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_route_total",
		Help:      "model route hits by scenario/model/reason",
		Labels:    []string{"scenario", "model", "reason"},
	})
	a.fallbackTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_fallback_total",
		Help:      "model fallback count",
		Labels:    []string{"scenario", "from_model", "to_model", "reason"},
	})
	a.circuitOpenTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_pro_circuit_open_total",
		Help:      "model pro circuit open count",
		Labels:    []string{"reason"},
	})
	a.callDuration = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_call_duration_ms",
		Help:      "model call duration in milliseconds",
		Labels:    []string{"scenario", "model", "method", "status"},
		Buckets:   []float64{10, 50, 100, 200, 500, 1000, 3000, 5000, 10000, 30000, 60000},
	})
	a.tokenTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_tokens_total",
		Help:      "model tokens by scenario/model/type/source",
		Labels:    []string{"scenario", "model", "token_type", "source"},
	})
	a.costTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "zhiguang",
		Subsystem: "agent",
		Name:      "model_cost_total",
		Help:      "model cost by scenario/model/currency/source",
		Labels:    []string{"scenario", "model", "currency", "source"},
	})
	return a
}

func (a *AgentObservability) Enabled() bool {
	return a != nil && a.enable
}

func (a *AgentObservability) RecordRoute(scenario, model, reason string) {
	if !a.Enabled() {
		return
	}
	a.routeTotal.Inc(nz(scenario), nz(model), nz(reason))
}

func (a *AgentObservability) RecordFallback(scenario, fromModel, toModel, reason string) {
	if !a.Enabled() {
		return
	}
	a.fallbackTotal.Inc(nz(scenario), nz(fromModel), nz(toModel), nz(reason))
}

func (a *AgentObservability) RecordProCircuitOpen(reason string) {
	if !a.Enabled() {
		return
	}
	a.circuitOpenTotal.Inc(nz(reason))
}

func (a *AgentObservability) RecordModelCall(te ModelTelemetry, startedAt time.Time) (Usage, float64, string) {
	if !a.Enabled() {
		return Usage{}, 0, "disabled"
	}
	latencyMs := time.Since(startedAt).Milliseconds()
	if latencyMs < 0 {
		latencyMs = 0
	}
	a.callDuration.Observe(latencyMs, nz(te.Scenario), nz(te.Model), nz(te.Method), nz(te.Status))

	u, source := a.resolveUsage(te)
	a.tokenTotal.Add(float64(u.PromptTokens), nz(te.Scenario), nz(te.Model), "in", source)
	a.tokenTotal.Add(float64(u.CompletionTokens), nz(te.Scenario), nz(te.Model), "out", source)
	a.tokenTotal.Add(float64(u.TotalTokens), nz(te.Scenario), nz(te.Model), "total", source)

	cost := a.calcCost(te.Model, u)
	if cost > 0 {
		a.costTotal.Add(cost, nz(te.Scenario), nz(te.Model), nz(a.cfg.Currency), source)
	}
	return u, cost, source
}

func ExtractUsageFromMessage(msg *schema.Message) *Usage {
	if msg == nil || msg.ResponseMeta == nil || msg.ResponseMeta.Usage == nil {
		return nil
	}
	u := msg.ResponseMeta.Usage
	return &Usage{
		PromptTokens:     max0(u.PromptTokens),
		CompletionTokens: max0(u.CompletionTokens),
		TotalTokens:      max0(u.TotalTokens),
	}
}

func (a *AgentObservability) resolveUsage(te ModelTelemetry) (Usage, string) {
	if te.Usage != nil && te.Usage.TotalTokens > 0 {
		u := *te.Usage
		if u.TotalTokens <= 0 {
			u.TotalTokens = max0(u.PromptTokens) + max0(u.CompletionTokens)
		}
		return normalizeUsage(u), "usage"
	}
	cpt := a.cfg.EstimateCharsPerToken
	if cpt <= 0 {
		cpt = 4
	}
	in := estimateTokens(te.InputText, cpt)
	out := estimateTokens(te.OutputText, cpt)
	return Usage{
		PromptTokens:     in,
		CompletionTokens: out,
		TotalTokens:      in + out,
	}, "estimate"
}

func (a *AgentObservability) calcCost(model string, u Usage) float64 {
	if u.TotalTokens <= 0 {
		return 0
	}
	inPrice := a.cfg.DefaultInputPer1K
	outPrice := a.cfg.DefaultOutputPer1K
	if p, ok := a.cfg.InputPer1K[model]; ok {
		inPrice = p
	}
	if p, ok := a.cfg.OutputPer1K[model]; ok {
		outPrice = p
	}
	return (float64(u.PromptTokens)/1000.0)*inPrice + (float64(u.CompletionTokens)/1000.0)*outPrice
}

func normalizeCost(c config.ModelCostConf) config.ModelCostConf {
	if strings.TrimSpace(c.Currency) == "" {
		c.Currency = "USD"
	}
	if c.EstimateCharsPerToken <= 0 {
		c.EstimateCharsPerToken = 4
	}
	if c.InputPer1K == nil {
		c.InputPer1K = map[string]float64{}
	}
	if c.OutputPer1K == nil {
		c.OutputPer1K = map[string]float64{}
	}
	return c
}

func estimateTokens(text string, charsPerToken float64) int {
	if charsPerToken <= 0 {
		charsPerToken = 4
	}
	r := []rune(strings.TrimSpace(text))
	if len(r) == 0 {
		return 0
	}
	return int(math.Ceil(float64(len(r)) / charsPerToken))
}

func normalizeUsage(u Usage) Usage {
	u.PromptTokens = max0(u.PromptTokens)
	u.CompletionTokens = max0(u.CompletionTokens)
	if u.TotalTokens <= 0 {
		u.TotalTokens = u.PromptTokens + u.CompletionTokens
	}
	u.TotalTokens = max0(u.TotalTokens)
	return u
}

func max0(v int) int {
	if v < 0 {
		return 0
	}
	return v
}

func nz(s string) string {
	v := strings.TrimSpace(s)
	if v == "" {
		return "unknown"
	}
	return v
}
