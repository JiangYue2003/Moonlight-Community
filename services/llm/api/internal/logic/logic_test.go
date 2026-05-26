package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	einodeepseek "github.com/cloudwego/eino-ext/components/model/deepseek"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/types"
)

// mockEmbedder 实现 embedding.Embedder 接口，直接返回预设向量。
type mockEmbedder struct {
	vecs [][]float64
	err  error
}

func (m *mockEmbedder) EmbedStrings(_ context.Context, texts []string, _ ...embedding.Option) ([][]float64, error) {
	if m.err != nil {
		return nil, m.err
	}
	out := make([][]float64, len(texts))
	for i := range texts {
		if i < len(m.vecs) {
			out[i] = m.vecs[i]
		} else {
			out[i] = []float64{0.1}
		}
	}
	return out, nil
}

// newSvc 构造测试用 ServiceContext：DeepSeek 用 httptest mock，Embedder 用 mockEmbedder，ES 用 httptest mock。
// RateLimit 用 miniredis 支撑，容量设为 9999 确保测试不被限流。
func newSvc(t *testing.T, deepseekHandler http.HandlerFunc, emb embedding.Embedder, esHandler http.HandlerFunc) *svc.ServiceContext {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/chat/completions", deepseekHandler)
	ds := httptest.NewServer(mux)
	t.Cleanup(ds.Close)
	es := httptest.NewServer(esHandler)
	t.Cleanup(es.Close)

	chat, err := einodeepseek.NewChatModel(context.Background(), &einodeepseek.ChatModelConfig{
		APIKey:  "k",
		Model:   "deepseek-chat",
		BaseURL: ds.URL,
		Timeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	esc, err := esx.New(esx.Config{Addrs: []string{es.URL}, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}

	mr, _ := miniredis.Run()
	t.Cleanup(mr.Close)
	rl := ratelimit.New(goredis.NewUniversalClient(&goredis.UniversalOptions{Addrs: []string{mr.Addr()}}))

	return &svc.ServiceContext{
		Config: config.Config{
			RagIndex: "rag",
			LlmRateLimit: config.LlmRateLimitConf{
				DescribeCapacity:     9999,
				DescribeRefillPerSec: 9999,
				QaCapacity:           9999,
				QaRefillPerSec:       9999,
			},
		},
		Chat:      chat,
		Embed:     emb,
		Es:        esc,
		RateLimit: rl,
	}
}

func defaultEmb() embedding.Embedder {
	return &mockEmbedder{vecs: [][]float64{{0.1, 0.2}}}
}

func TestDescribe_PostProcessTrims(t *testing.T) {
	sc := newSvc(t,
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, `{"id":"1","object":"chat.completion","model":"deepseek-chat","choices":[{"index":0,"message":{"role":"assistant","content":"  \"这是一个非常吸引人的描述介绍文案。\" \n\n第二段不要"},"finish_reason":"stop"}]}`)
		},
		&mockEmbedder{},
		func(w http.ResponseWriter, r *http.Request) { t.Fatal("es should not be called") },
	)
	l := NewDescribeLogic(context.Background(), sc)
	resp, err := l.Describe(&types.DescribeReq{Body: "正文"})
	if err != nil {
		t.Fatal(err)
	}
	want := "这是一个非常吸引人的描述介绍文案"
	if resp.Description != want {
		t.Errorf("got %q want %q", resp.Description, want)
	}
}

func TestDescribe_EmptyBody(t *testing.T) {
	sc := newSvc(t,
		func(w http.ResponseWriter, r *http.Request) { t.Fatal("chat should not be called") },
		&mockEmbedder{},
		func(w http.ResponseWriter, r *http.Request) {},
	)
	l := NewDescribeLogic(context.Background(), sc)
	if _, err := l.Describe(&types.DescribeReq{Body: "  "}); err == nil {
		t.Fatal("expect error")
	}
}

// fakeFlusher 让 httptest.NewRecorder 实现 http.Flusher。
type fakeFlusher struct {
	*httptest.ResponseRecorder
}

func (f *fakeFlusher) Flush() {}

func TestQaStream_NoMatchReplies(t *testing.T) {
	sc := newSvc(t,
		func(w http.ResponseWriter, r *http.Request) { t.Fatal("chat should not be called") },
		defaultEmb(),
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, `{"hits":{"total":{"value":0},"hits":[]}}`)
		},
	)
	rec := &fakeFlusher{httptest.NewRecorder()}
	l := NewQaStreamLogic(context.Background(), sc)
	if err := l.Run(rec, &types.QaReq{PostId: 1, Question: "hi"}); err != nil {
		t.Fatal(err)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "未在该知文中找到相关内容") || !strings.Contains(body, "data: [DONE]") {
		t.Fatalf("body=%s", body)
	}
}

func TestQaStream_StreamsTokensAndDone(t *testing.T) {
	sc := newSvc(t,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			f, _ := w.(http.Flusher)
			writeSSEData(w, f, `{"id":"1","object":"chat.completion.chunk","model":"deepseek-chat","choices":[{"index":0,"delta":{"content":"片"},"finish_reason":null}]}`)
			writeSSEData(w, f, `{"id":"1","object":"chat.completion.chunk","model":"deepseek-chat","choices":[{"index":0,"delta":{"content":"段"},"finish_reason":null}]}`)
			writeSSEData(w, f, `[DONE]`)
		},
		defaultEmb(),
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, `{"hits":{"total":{"value":1},"hits":[
				{"_id":"c1","_score":1,"_source":{"post_id":1,"text":"context片段"}}
			]}}`)
		},
	)
	rec := &fakeFlusher{httptest.NewRecorder()}
	l := NewQaStreamLogic(context.Background(), sc)
	if err := l.Run(rec, &types.QaReq{PostId: 1, Question: "hi", TopK: 5}); err != nil {
		t.Fatal(err)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "data: 片") || !strings.Contains(body, "data: 段") {
		t.Fatalf("missing tokens: %s", body)
	}
	if !strings.HasSuffix(strings.TrimRight(body, "\n"), "data: [DONE]") {
		t.Fatalf("body did not end with [DONE]: %s", body)
	}
}

func TestQaStream_FilterByPostId(t *testing.T) {
	res := &esx.SearchResult{}
	res.Hits.Hits = []esx.Hit{
		{Source: json.RawMessage(`{"post_id":2,"text":"a"}`)},
		{Source: json.RawMessage(`{"post_id":1,"text":"b"}`)},
		{Source: json.RawMessage(`{"post_id":1,"text":"c"}`)},
		{Source: json.RawMessage(`{"post_id":3,"text":"d"}`)},
	}
	got := filterByPostId(res, 1, 5)
	if !equalSlices(got, []string{"b", "c"}) {
		t.Errorf("got %v", got)
	}
}

func TestQaStream_BadParams(t *testing.T) {
	sc := newSvc(t,
		func(w http.ResponseWriter, r *http.Request) {},
		&mockEmbedder{},
		func(w http.ResponseWriter, r *http.Request) {},
	)
	rec := &fakeFlusher{httptest.NewRecorder()}
	l := NewQaStreamLogic(context.Background(), sc)
	err := l.Run(rec, &types.QaReq{PostId: 0, Question: ""})
	if err == nil || !strings.Contains(err.Error(), "required") {
		t.Fatalf("err=%v", err)
	}
}

// helpers
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func writeSSEData(w io.Writer, f http.Flusher, payload string) {
	_, _ = io.WriteString(w, "data: "+payload+"\n\n")
	if f != nil {
		f.Flush()
	}
}

// 防止 go vet 抱怨未用 import
var _ = errors.New
var _ = fmt.Errorf
var _ model.ChatModel = nil
var _ schema.Message = schema.Message{}
