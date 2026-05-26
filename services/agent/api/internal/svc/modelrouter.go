package svc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
)

type RouteScenario string

const (
	RouteScenarioChat        RouteScenario = "chat"
	RouteScenarioFactExtract RouteScenario = "fact_extract"
)

type RouteInput struct {
	Question        string
	Prompt          string
	Summary         string
	RecallCount     int
	SessionMsgCount int
	PinContent      string
}

type RouteDecision struct {
	Model     model.ChatModel
	ModelName string
	Reason    string
}

type ModelRouter struct {
	cfg  config.ModelRouteConf
	rdb  redis.UniversalClient
	lite model.ChatModel
	pro  model.ChatModel
	obs  RouteObserver
}

type RouteObserver interface {
	RecordRoute(scenario, model, reason string)
	RecordProCircuitOpen(reason string)
}

func NewModelRouter(cfg config.ModelRouteConf, rdb redis.UniversalClient, lite, pro model.ChatModel, obs RouteObserver) *ModelRouter {
	return &ModelRouter{
		cfg:  cfg,
		rdb:  rdb,
		lite: lite,
		pro:  pro,
		obs:  obs,
	}
}

func (r *ModelRouter) Decide(ctx context.Context, scenario RouteScenario, in RouteInput) RouteDecision {
	record := func(d RouteDecision) RouteDecision {
		if r != nil && r.obs != nil {
			r.obs.RecordRoute(string(scenario), d.ModelName, d.Reason)
			if d.Reason == "pro_unhealthy" {
				r.obs.RecordProCircuitOpen(d.Reason)
			}
		}
		return d
	}
	if r == nil || r.lite == nil {
		return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "lite_missing"})
	}
	if !r.cfg.Enable || r.pro == nil {
		return record(RouteDecision{Model: r.lite, ModelName: "lite", Reason: "route_disabled_or_pro_missing"})
	}
	if r.isProUnhealthy(ctx) {
		return record(RouteDecision{Model: r.lite, ModelName: "lite", Reason: "pro_unhealthy"})
	}
	switch scenario {
	case RouteScenarioFactExtract:
		if runeLen(in.PinContent) >= r.cfg.PinContentRunesPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "pin_content_long"})
		}
	case RouteScenarioChat:
		if runeLen(in.Question) >= r.cfg.QuestionRunesPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "question_long"})
		}
		if runeLen(in.Prompt) >= r.cfg.PromptRunesPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "prompt_long"})
		}
		if in.RecallCount >= r.cfg.RecallCountPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "recall_many"})
		}
		if runeLen(in.Summary) >= r.cfg.SummaryRunesPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "summary_long"})
		}
		if in.SessionMsgCount >= r.cfg.SessionMsgsPro {
			return record(RouteDecision{Model: r.pro, ModelName: "pro", Reason: "session_history_long"})
		}
	}
	return record(RouteDecision{Model: r.lite, ModelName: "lite", Reason: "cost_first_default"})
}

func (r *ModelRouter) RecordFailure(ctx context.Context, modelName string) {
	if r == nil || r.rdb == nil || strings.TrimSpace(modelName) != "pro" {
		return
	}
	windowSec := r.cfg.ProFailWindowSec
	if windowSec <= 0 {
		windowSec = 120
	}
	key := r.proFailKey()
	pipe := r.rdb.TxPipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Duration(windowSec)*time.Second)
	_, _ = pipe.Exec(ctx)
}

func (r *ModelRouter) RetryOnProFail() bool {
	if r == nil {
		return false
	}
	return r.cfg.RetryOnProFail
}

func (r *ModelRouter) EmitRouteEvent() bool {
	if r == nil {
		return false
	}
	return r.cfg.EmitRouteEvent
}

func (r *ModelRouter) isProUnhealthy(ctx context.Context) bool {
	if r == nil || r.rdb == nil {
		return false
	}
	threshold := r.cfg.ProFailThreshold
	if threshold <= 0 {
		return false
	}
	n, err := r.rdb.Get(ctx, r.proFailKey()).Int()
	if err != nil {
		return false
	}
	return n >= threshold
}

func (r *ModelRouter) proFailKey() string {
	return "agent:modelroute:pro:fail"
}

func runeLen(s string) int {
	return len([]rune(strings.TrimSpace(s)))
}

func routeLogFields(traceID string, d RouteDecision, fallback bool) string {
	return fmt.Sprintf("trace=%s model=%s reason=%s fallback=%v", traceID, d.ModelName, d.Reason, fallback)
}

func RouteLogFields(traceID string, d RouteDecision, fallback bool) string {
	return routeLogFields(traceID, d, fallback)
}
