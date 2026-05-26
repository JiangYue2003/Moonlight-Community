package logic

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func TestSuggestDescriptionCompat_ForwardsToLLM(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/llm/describe" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Fatalf("auth header not forwarded")
		}
		b, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(b), "content") {
			t.Fatalf("unexpected body: %s", string(b))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"description":"ok"}`))
	}))
	defer up.Close()

	sc := &svc.ServiceContext{Config: config.Config{LlmProxy: config.LlmProxyConf{BaseURL: up.URL}}, LlmClient: up.Client()}
	logic := NewSuggestDescriptionCompatLogic(context.Background(), sc)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/knowposts/description/suggest", strings.NewReader(`{"content":"hello"}`))
	req.Header.Set("Authorization", "Bearer token")
	rr := httptest.NewRecorder()

	if err := logic.Run(rr, req); err != nil {
		t.Fatalf("run error: %v", err)
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	if !strings.Contains(rr.Body.String(), "description") {
		t.Fatalf("unexpected body: %s", rr.Body.String())
	}
}
