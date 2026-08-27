package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/config"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	knowpostclient "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/client/knowpost"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	llmclient "github.com/zhiguang/zhiguang-go/services/llm/rpc/client/llm"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
	searchclient "github.com/zhiguang/zhiguang-go/services/search/rpc/client/search"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
	authclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/auth"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestNewEngineOmitsGinAccessLoggerWhenDisabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewEngine(&srv.ServiceContext{Config: config.Config{HTTPAccessLog: false}})

	if got := len(router.Handlers); got != 1 {
		t.Fatalf("global middleware count=%d, want recovery only", got)
	}
}

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

type stubFollowingFeedAuth struct {
	authclient.Auth
}

func (stubFollowingFeedAuth) VerifyToken(
	_ context.Context,
	in *authclient.VerifyTokenReq,
	_ ...grpc.CallOption,
) (*authclient.VerifyTokenResp, error) {
	if in.AccessToken != "valid-token" {
		return &userpb.VerifyTokenResp{Valid: false}, nil
	}
	return &userpb.VerifyTokenResp{Valid: true, UserId: 42}, nil
}

type stubFollowingFeedKnowPost struct {
	knowpostclient.KnowPost
	request  *knowpostclient.GetUserFeedReq
	response *knowpostclient.FeedPage
	err      error
}

func (s *stubFollowingFeedKnowPost) GetUserFeed(
	_ context.Context,
	in *knowpostclient.GetUserFeedReq,
	_ ...grpc.CallOption,
) (*knowpostclient.FeedPage, error) {
	s.request = in
	if s.err != nil {
		return nil, s.err
	}
	if s.response != nil {
		return s.response, nil
	}
	return &knowpostpb.FeedPage{
		Items: []*knowpostpb.FeedItem{{Id: "99", CreatorId: 7, Title: "feed item"}},
		Page:  in.Page,
		Size:  in.Size,
	}, nil
}

func TestFollowingFeedRouteRequiresAuthAndUsesAuthenticatedUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	knowPost := &stubFollowingFeedKnowPost{}
	router := NewEngine(&srv.ServiceContext{
		AuthRpc:     stubFollowingFeedAuth{},
		KnowPostRpc: knowPost,
	})

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?page=2&size=10", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request != nil {
			t.Fatal("GetUserFeed called without authentication")
		}
	})

	t.Run("authenticated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?page=2&size=10", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request == nil {
			t.Fatal("GetUserFeed was not called")
		}
		if knowPost.request.UserId != 42 || knowPost.request.Page != 2 || knowPost.request.Size != 10 {
			t.Fatalf("request=%+v", knowPost.request)
		}
		if !strings.Contains(rec.Body.String(), `"id":"99"`) {
			t.Fatalf("body=%s", rec.Body.String())
		}
	})

	t.Run("rejects invalid pagination before rpc", func(t *testing.T) {
		knowPost.request = nil
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?page=0&size=101", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request != nil {
			t.Fatal("GetUserFeed called with invalid pagination")
		}
	})

	t.Run("rejects non-numeric pagination before rpc", func(t *testing.T) {
		knowPost.request = nil
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?page=abc&size=20", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request != nil {
			t.Fatal("GetUserFeed called with non-numeric pagination")
		}
	})

	t.Run("cursor mode ignores page and forwards cursor", func(t *testing.T) {
		knowPost.request = nil
		knowPost.response = &knowpostpb.FeedPage{
			Items:      []*knowpostpb.FeedItem{{Id: "100"}},
			HasMore:    true,
			Size:       10,
			NextCursor: "next-opaque-cursor",
		}
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?cursor=current-opaque-cursor&page=not-a-number&size=10", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request == nil || knowPost.request.UserId != 42 || knowPost.request.Page != 0 ||
			knowPost.request.Size != 10 || knowPost.request.Cursor != "current-opaque-cursor" {
			t.Fatalf("request=%+v", knowPost.request)
		}
		if !strings.Contains(rec.Body.String(), `"nextCursor":"next-opaque-cursor"`) {
			t.Fatalf("body=%s", rec.Body.String())
		}
	})

	t.Run("rejects oversized cursor before rpc", func(t *testing.T) {
		knowPost.request = nil
		knowPost.response = nil
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?cursor="+strings.Repeat("a", 257), nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if knowPost.request != nil {
			t.Fatal("GetUserFeed called with oversized cursor")
		}
	})

	t.Run("maps invalid cursor from rpc to bad request", func(t *testing.T) {
		knowPost.request = nil
		knowPost.response = nil
		knowPost.err = status.Error(codes.InvalidArgument, "invalid feed cursor")
		t.Cleanup(func() { knowPost.err = nil })
		req := httptest.NewRequest(http.MethodGet, "/api/v1/knowposts/following-feed?cursor=invalid&size=20", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
		}
		if !strings.Contains(rec.Body.String(), `"code":"BAD_REQUEST"`) {
			t.Fatalf("body=%s", rec.Body.String())
		}
	})
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
