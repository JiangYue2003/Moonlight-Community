package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	searchclient "github.com/zhiguang/zhiguang-go/services/search/rpc/client/search"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
	llmclient "github.com/zhiguang/zhiguang-go/services/llm/rpc/client/llm"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type stubSearch struct{}

func (stubSearch) Search(ctx context.Context, in *searchclient.SearchReq, opts ...grpc.CallOption) (*searchclient.SearchResp, error) {
	return &searchpb.SearchResp{
		Items: []*searchpb.Hit{{Id: "1", Title: "t1", Description: "d1"}},
	}, nil
}

func (stubSearch) Suggest(ctx context.Context, in *searchclient.SuggestReq, opts ...grpc.CallOption) (*searchclient.SuggestResp, error) {
	return &searchpb.SuggestResp{Items: []string{"a", "b"}}, nil
}

type stubQaStream struct {
	items []*llmpb.QaChunk
	idx   int
}

func (s *stubQaStream) Header() (metadata.MD, error) { return nil, nil }
func (s *stubQaStream) Trailer() metadata.MD         { return nil }
func (s *stubQaStream) CloseSend() error             { return nil }
func (s *stubQaStream) Context() context.Context     { return context.Background() }
func (s *stubQaStream) SendMsg(m any) error          { return nil }
func (s *stubQaStream) RecvMsg(m any) error          { return nil }
func (s *stubQaStream) Recv() (*llmpb.QaChunk, error) {
	if s.idx >= len(s.items) {
		return nil, io.EOF
	}
	item := s.items[s.idx]
	s.idx++
	return item, nil
}

type stubLlm struct{}

func (stubLlm) Describe(ctx context.Context, in *llmclient.DescribeReq, opts ...grpc.CallOption) (*llmclient.DescribeResp, error) {
	return &llmpb.DescribeResp{Description: "desc"}, nil
}

func (stubLlm) QaStream(ctx context.Context, in *llmclient.QaStreamReq, opts ...grpc.CallOption) (llmpb.Llm_QaStreamClient, error) {
	return &stubQaStream{items: []*llmpb.QaChunk{
		{Data: "hello"},
		{Data: "[DONE]", Done: true},
	}}, nil
}

func TestSearchRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sc := &srv.ServiceContext{SearchRpc: stubSearch{}}
	r := gin.New()
	r.GET("/api/v1/search/", searchPosts(sc))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/search/?q=go", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), `"contentId":"1"`) {
		t.Fatalf("body=%s", rec.Body.String())
	}
}

func TestQaCompatStreamRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sc := &srv.ServiceContext{LlmRpc: stubLlm{}}
	r := gin.New()
	r.GET("/api/v1/knowposts/:id/qa/stream", qaCompatStream(sc))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/123/qa/stream?question=hi", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	body := rec.Body.String()
	if !strings.Contains(body, "data: hello") || !strings.Contains(body, "data: [DONE]") {
		t.Fatalf("body=%s", body)
	}
}
