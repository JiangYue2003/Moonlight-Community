package logic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/types"
	"google.golang.org/grpc"
)

type stubCounterClient struct {
	counts map[string]map[string]int64
	marks  map[string]bool
}

func (s stubCounterClient) Toggle(context.Context, *counterpb.ToggleReq, ...grpc.CallOption) (*counterpb.ToggleResp, error) {
	panic("not implemented")
}
func (s stubCounterClient) GetCounts(_ context.Context, in *counterpb.GetCountsReq, _ ...grpc.CallOption) (*counterpb.GetCountsResp, error) {
	return &counterpb.GetCountsResp{Counts: s.counts[in.EntityId]}, nil
}
func (s stubCounterClient) IsMarked(_ context.Context, in *counterpb.IsMarkedReq, _ ...grpc.CallOption) (*counterpb.IsMarkedResp, error) {
	key := in.EntityId + ":" + in.Metric + ":" + ctxdata.FormatUserId(in.UserId)
	return &counterpb.IsMarkedResp{Marked: s.marks[key]}, nil
}
func (s stubCounterClient) BatchGetCounts(context.Context, *counterpb.BatchGetCountsReq, ...grpc.CallOption) (*counterpb.BatchGetCountsResp, error) {
	panic("not implemented")
}

func newTestSvc(t *testing.T, h http.HandlerFunc) (*svc.ServiceContext, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c, err := esx.New(esx.Config{Addrs: []string{srv.URL}, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	return &svc.ServiceContext{
		Config:     config.Config{ContentIndex: "idx"},
		Es:         c,
		CounterRpc: stubCounterClient{},
	}, srv
}

func TestSearch_EmptyQReturnsEmpty(t *testing.T) {
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not call ES")
	})
	l := NewSearchLogic(context.Background(), sc)
	resp, err := l.Search(&types.SearchReq{Q: "  ", Size: 10})
	if err != nil || len(resp.Items) != 0 || resp.HasMore {
		t.Fatalf("resp=%+v err=%v", resp, err)
	}
}

func TestSearch_DescriptionUsesHighlight(t *testing.T) {
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{
			"hits":{"total":{"value":1},"hits":[
				{"_id":"1","_score":1,
				 "_source":{"content_id":"1","title":"t","description":"desc","tags":["go"],"author_id":7},
				 "highlight":{"title":["<em>go</em>"],"body":["body <em>go</em> snippet"]},
				 "sort":[1.0, 1700000000000, 1]}
			]}
		}`)
	})
	sc.CounterRpc = stubCounterClient{
		counts: map[string]map[string]int64{"1": {"like": 12, "fav": 4}},
		marks:  map[string]bool{"1:like:99": true, "1:fav:99": true},
	}
	l := NewSearchLogic(ctxdata.WithUserId(context.Background(), 99), sc)
	resp, err := l.Search(&types.SearchReq{Q: "go", Size: 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 1 {
		t.Fatalf("items=%d", len(resp.Items))
	}
	got := resp.Items[0]
	if !strings.Contains(got.Description, "<em>go</em>") {
		t.Errorf("description=%q", got.Description)
	}
	if got.AuthorId != 7 {
		t.Errorf("authorId=%d", got.AuthorId)
	}
	if got.LikeCount != 12 || got.FavoriteCount != 4 {
		t.Errorf("counts should be refreshed from counter rpc: %+v", got)
	}
	if !got.Liked || !got.Faved {
		t.Errorf("liked/faved should be hydrated: %+v", got)
	}
	// size==len(items) → has more
	if !resp.HasMore || resp.NextAfter == "" {
		t.Errorf("hasMore=%v nextAfter=%q", resp.HasMore, resp.NextAfter)
	}
}

func TestSearch_DescriptionFallbackToDescription(t *testing.T) {
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{
			"hits":{"total":{"value":1},"hits":[
				{"_id":"1","_score":1,
				 "_source":{"content_id":"1","description":"小描述"},
				 "sort":[1,1,1]}
			]}
		}`)
	})
	l := NewSearchLogic(context.Background(), sc)
	resp, err := l.Search(&types.SearchReq{Q: "x", Size: 5})
	if err != nil || len(resp.Items) != 1 {
		t.Fatalf("err=%v items=%v", err, resp.Items)
	}
	if resp.Items[0].Description != "小描述" {
		t.Errorf("description=%q", resp.Items[0].Description)
	}
	// items < size → no more
	if resp.HasMore {
		t.Errorf("should not have more")
	}
}

func TestSearch_BadCursorFallback(t *testing.T) {
	called := 0
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		buf, _ := io.ReadAll(r.Body)
		// 期待 fallback 后 body 不含 search_after
		if strings.Contains(string(buf), "search_after") {
			t.Errorf("should drop bad cursor: %s", buf)
		}
		called++
		_, _ = io.WriteString(w, `{"hits":{"total":{"value":0},"hits":[]}}`)
	})
	l := NewSearchLogic(context.Background(), sc)
	_, err := l.Search(&types.SearchReq{Q: "x", Size: 5, After: "!!!bad"})
	if err != nil || called != 1 {
		t.Fatalf("err=%v called=%d", err, called)
	}
}

func TestSuggest_EmptyPrefix(t *testing.T) {
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not call ES")
	})
	l := NewSuggestLogic(context.Background(), sc)
	resp, err := l.Suggest(&types.SuggestReq{Prefix: "  ", Size: 10})
	if err != nil || len(resp.Items) != 0 {
		t.Fatalf("resp=%+v err=%v", resp, err)
	}
}

func TestSuggest_OK(t *testing.T) {
	sc, _ := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"suggest":{"s":[{"text":"go","options":[
			{"text":"go-zero"},{"text":"go语言"}
		]}]}}`)
	})
	l := NewSuggestLogic(context.Background(), sc)
	resp, err := l.Suggest(&types.SuggestReq{Prefix: "go", Size: 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 2 || resp.Items[0] != "go-zero" {
		t.Errorf("items=%v", resp.Items)
	}
}

// 防止 go vet 抱怨没用到 json
var _ = json.RawMessage{}
