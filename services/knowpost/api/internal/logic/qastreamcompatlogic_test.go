package logic

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func TestQaStreamCompat_ForwardsSSE(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/llm/qa/stream" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("postId") != "123" {
			t.Fatalf("postId not forwarded")
		}
		if r.URL.Query().Get("question") != "q" {
			t.Fatalf("question not forwarded")
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Fatalf("auth header not forwarded")
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("data: hi\\n\\n"))
		_, _ = w.Write([]byte("data: [DONE]\\n\\n"))
	}))
	defer up.Close()

	sc := &svc.ServiceContext{Config: config.Config{LlmProxy: config.LlmProxyConf{BaseURL: up.URL}}, LlmClient: up.Client()}
	logic := NewQaStreamCompatLogic(context.Background(), sc).WithPostID(123)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/123/qa/stream?question=q&topK=5&maxTokens=1024", nil)
	req.Header.Set("Authorization", "Bearer token")
	rr := httptest.NewRecorder()

	if err := logic.Run(rr, req); err != nil {
		t.Fatalf("run error: %v", err)
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	if !strings.Contains(rr.Body.String(), "data: hi") || !strings.Contains(rr.Body.String(), "[DONE]") {
		t.Fatalf("unexpected sse body: %s", rr.Body.String())
	}
}
