package logic

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/types"
)

type fakeFlusher struct{ *httptest.ResponseRecorder }

func (f *fakeFlusher) Flush() {}

func TestChatStream_RequireSessionID(t *testing.T) {
	l := &ChatStreamLogic{ctx: context.Background(), svcCtx: &svc.ServiceContext{Config: config.Config{}}}
	rec := &fakeFlusher{httptest.NewRecorder()}
	err := l.Run(rec, &types.ChatStreamReq{SessionID: "", Question: "x"})
	if err == nil || !strings.Contains(err.Error(), "sessionId") {
		t.Fatalf("err=%v", err)
	}
}

func TestCompactSessionSummary(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{Addrs: []string{mr.Addr()}})
	sctx := &svc.ServiceContext{Config: config.Config{Agent: config.AgentConf{SummaryTTLHours: 1}}, Redis: rdb}
	l := &ChatStreamLogic{ctx: context.Background(), svcCtx: sctx}
	uid := int64(1)
	sid := "s1"
	for i := 0; i < 30; i++ {
		_ = l.appendHistory(uid, sid, "user", "hello", nowMs())
	}
	if err := l.compactSession(uid, sid); err != nil {
		t.Fatal(err)
	}
	v, err := rdb.Get(context.Background(), svc.SessionSummaryKey(uid, sid)).Result()
	if err != nil || v == "" {
		t.Fatalf("summary missing err=%v", err)
	}
}

var _ = http.StatusOK
