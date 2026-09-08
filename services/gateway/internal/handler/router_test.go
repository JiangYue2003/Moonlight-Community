package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
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

func TestNewEngineRouteParity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewEngine(&srv.ServiceContext{})

	got := make([]string, 0, len(router.Routes()))
	for _, route := range router.Routes() {
		got = append(got, route.Method+" "+route.Path)
	}
	sort.Strings(got)

	want := []string{
		"DELETE /api/v1/knowposts/:id",
		"GET /api/v1/auth/me",
		"GET /api/v1/counter/:etype/:eid",
		"GET /api/v1/knowposts/:id/qa/stream",
		"GET /api/v1/knowposts/detail/:id",
		"GET /api/v1/knowposts/feed",
		"GET /api/v1/knowposts/following-feed",
		"GET /api/v1/knowposts/mine",
		"GET /api/v1/llm/qa/stream",
		"GET /api/v1/profile/me",
		"GET /api/v1/relation/counter",
		"GET /api/v1/relation/followers",
		"GET /api/v1/relation/following",
		"GET /api/v1/relation/status",
		"GET /api/v1/search/",
		"GET /api/v1/search/suggest",
		"PATCH /api/v1/knowposts/:id",
		"PATCH /api/v1/knowposts/:id/top",
		"PATCH /api/v1/knowposts/:id/visibility",
		"PATCH /api/v1/profile/",
		"POST /api/v1/action/fav",
		"POST /api/v1/action/like",
		"POST /api/v1/action/unfav",
		"POST /api/v1/action/unlike",
		"POST /api/v1/auth/login",
		"POST /api/v1/auth/logout",
		"POST /api/v1/auth/password/reset",
		"POST /api/v1/auth/register",
		"POST /api/v1/auth/send-code",
		"POST /api/v1/auth/token/refresh",
		"POST /api/v1/knowposts/:id/content/confirm",
		"POST /api/v1/knowposts/:id/publish",
		"POST /api/v1/knowposts/:id/rag/reindex",
		"POST /api/v1/knowposts/:id/reindex",
		"POST /api/v1/knowposts/description/suggest",
		"POST /api/v1/knowposts/drafts",
		"POST /api/v1/llm/describe",
		"POST /api/v1/profile/avatar",
		"POST /api/v1/relation/follow",
		"POST /api/v1/relation/unfollow",
		"POST /api/v1/storage/presign",
	}
	sort.Strings(want)

	if len(got) != len(want) {
		t.Fatalf("route count=%d, want %d\ngot: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("route[%d]=%q, want %q\ngot: %v", i, got[i], want[i], got)
		}
	}
}

func TestNewEngineRequiredAuthRouteParity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewEngine(&srv.ServiceContext{})

	protectedRoutes := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/api/v1/auth/logout"},
		{http.MethodGet, "/api/v1/profile/me"},
		{http.MethodPatch, "/api/v1/profile/"},
		{http.MethodPost, "/api/v1/profile/avatar"},
		{http.MethodPost, "/api/v1/storage/presign"},
		{http.MethodPost, "/api/v1/knowposts/drafts"},
		{http.MethodPost, "/api/v1/knowposts/1/content/confirm"},
		{http.MethodPatch, "/api/v1/knowposts/1"},
		{http.MethodPost, "/api/v1/knowposts/1/publish"},
		{http.MethodPatch, "/api/v1/knowposts/1/top"},
		{http.MethodPatch, "/api/v1/knowposts/1/visibility"},
		{http.MethodDelete, "/api/v1/knowposts/1"},
		{http.MethodGet, "/api/v1/knowposts/mine"},
		{http.MethodGet, "/api/v1/knowposts/following-feed"},
		{http.MethodPost, "/api/v1/knowposts/1/reindex"},
		{http.MethodPost, "/api/v1/knowposts/1/rag/reindex"},
		{http.MethodPost, "/api/v1/relation/follow"},
		{http.MethodPost, "/api/v1/relation/unfollow"},
		{http.MethodGet, "/api/v1/relation/following"},
		{http.MethodGet, "/api/v1/relation/followers"},
		{http.MethodGet, "/api/v1/relation/counter"},
		{http.MethodPost, "/api/v1/action/like"},
		{http.MethodPost, "/api/v1/action/unlike"},
		{http.MethodPost, "/api/v1/action/fav"},
		{http.MethodPost, "/api/v1/action/unfav"},
		{http.MethodPost, "/api/v1/llm/describe"},
	}

	for _, route := range protectedRoutes {
		route := route
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			req := httptest.NewRequest(route.method, route.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusUnauthorized {
				t.Fatalf("status=%d, want %d; body=%s", rec.Code, http.StatusUnauthorized, rec.Body.String())
			}
		})
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
