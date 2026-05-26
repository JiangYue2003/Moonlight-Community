package tooling

import (
	"context"
	"testing"
)

type mockAudit struct{ n int }

func (m *mockAudit) LogToolCall(context.Context, AuditRecord) { m.n++ }

func TestRegistryExecute(t *testing.T) {
	a := &mockAudit{}
	r := NewRegistry([]string{"t1"}, a)
	r.Register(Tool{
		Name: "t1",
		Schema: Schema{
			"q": {Type: FieldString, Required: true, MaxLen: 10},
		},
		Run: func(ctx context.Context, c Call) (any, error) {
			return "ok", nil
		},
	})
	ret, err := r.Execute(context.Background(), Call{Tool: "t1", UserID: 1, Params: map[string]any{"q": "abc"}})
	if err != nil || ret.(string) != "ok" {
		t.Fatalf("ret=%v err=%v", ret, err)
	}
	if a.n != 1 {
		t.Fatalf("audit n=%d", a.n)
	}
}

func TestRegistryRejectUnknownParam(t *testing.T) {
	r := NewRegistry([]string{"t1"}, nil)
	r.Register(Tool{Name: "t1", Schema: Schema{"q": {Type: FieldString}}, Run: func(ctx context.Context, c Call) (any, error) { return nil, nil }})
	_, err := r.Execute(context.Background(), Call{Tool: "t1", UserID: 1, Params: map[string]any{"x": "1"}})
	if err == nil {
		t.Fatal("expected error")
	}
}
