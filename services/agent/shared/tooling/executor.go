package tooling

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/security"
)

type FieldType string

const (
	FieldString FieldType = "string"
	FieldInt    FieldType = "int"
	FieldBool   FieldType = "bool"
)

type FieldRule struct {
	Type     FieldType
	Required bool
	MinInt   int
	MaxInt   int
	MaxLen   int
	Enum     []string
}

type Schema map[string]FieldRule

type Call struct {
	Tool      string
	UserID    int64
	SessionID string
	TraceID   string
	Params    map[string]any
}

type Tool struct {
	Name    string
	Schema  Schema
	Enabled func() bool
	Run     func(ctx context.Context, c Call) (any, error)
}

type AuditLogger interface {
	LogToolCall(ctx context.Context, rec AuditRecord)
}

type AuditRecord struct {
	SessionID  string
	TraceID    string
	UserID     int64
	Tool       string
	ParamsHash string
	LatencyMs  int64
	Status     string
	ErrMsg     string
}

type Registry struct {
	tools     map[string]Tool
	whitelist map[string]struct{}
	audit     AuditLogger
}

func NewRegistry(whitelist []string, audit AuditLogger) *Registry {
	wl := make(map[string]struct{}, len(whitelist))
	for _, t := range whitelist {
		if s := strings.TrimSpace(t); s != "" {
			wl[s] = struct{}{}
		}
	}
	return &Registry{tools: make(map[string]Tool, len(wl)), whitelist: wl, audit: audit}
}

func (r *Registry) Register(t Tool) {
	r.tools[t.Name] = t
}

func (r *Registry) Execute(ctx context.Context, c Call) (any, error) {
	start := time.Now()
	if err := security.EnsureUserScope(c.UserID); err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	if _, ok := r.whitelist[c.Tool]; !ok {
		return nil, fmt.Errorf("tool not allowed")
	}
	t, ok := r.tools[c.Tool]
	if !ok {
		return nil, fmt.Errorf("tool not registered")
	}
	if t.Enabled != nil && !t.Enabled() {
		return nil, fmt.Errorf("tool disabled")
	}
	if err := validateAgainstSchema(c.Params, t.Schema); err != nil {
		return nil, err
	}
	ret, err := t.Run(ctx, c)
	r.logAudit(ctx, c, elapsedMs(start), err)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (r *Registry) logAudit(ctx context.Context, c Call, latency int64, err error) {
	if r.audit == nil {
		return
	}
	status := "ok"
	errMsg := ""
	if err != nil {
		status = "error"
		errMsg = err.Error()
		if len(errMsg) > 200 {
			errMsg = errMsg[:200]
		}
	}
	r.audit.LogToolCall(ctx, AuditRecord{
		SessionID:  c.SessionID,
		TraceID:    c.TraceID,
		UserID:     c.UserID,
		Tool:       c.Tool,
		ParamsHash: hashParams(c.Params),
		LatencyMs:  latency,
		Status:     status,
		ErrMsg:     errMsg,
	})
}

func validateAgainstSchema(params map[string]any, schema Schema) error {
	if params == nil {
		params = map[string]any{}
	}
	for name, rule := range schema {
		v, ok := params[name]
		if !ok {
			if rule.Required {
				return fmt.Errorf("missing param: %s", name)
			}
			continue
		}
		switch rule.Type {
		case FieldString:
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("param %s type invalid", name)
			}
			if rule.MaxLen > 0 && len([]rune(s)) > rule.MaxLen {
				return fmt.Errorf("param %s too long", name)
			}
			if len(rule.Enum) > 0 {
				if !inEnum(s, rule.Enum) {
					return fmt.Errorf("param %s out of enum", name)
				}
			}
		case FieldInt:
			n, ok := asInt(v)
			if !ok {
				return fmt.Errorf("param %s type invalid", name)
			}
			if rule.MinInt != 0 && n < rule.MinInt {
				return fmt.Errorf("param %s too small", name)
			}
			if rule.MaxInt != 0 && n > rule.MaxInt {
				return fmt.Errorf("param %s too large", name)
			}
		case FieldBool:
			if _, ok := v.(bool); !ok {
				return fmt.Errorf("param %s type invalid", name)
			}
		default:
			return fmt.Errorf("param %s unsupported schema type", name)
		}
	}
	for k := range params {
		if _, ok := schema[k]; !ok {
			return fmt.Errorf("param %s not allowed", k)
		}
	}
	return nil
}

func inEnum(v string, enums []string) bool {
	for _, e := range enums {
		if v == e {
			return true
		}
	}
	return false
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

func hashParams(params map[string]any) string {
	if len(params) == 0 {
		return ""
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ordered := make(map[string]any, len(keys))
	for _, k := range keys {
		ordered[k] = params[k]
	}
	b, _ := json.Marshal(ordered)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func elapsedMs(t time.Time) int64 { return time.Since(t).Milliseconds() }
