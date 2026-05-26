package searchlogic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/search/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/rpc/internal/svc"
	searchpb "github.com/zhiguang/zhiguang-go/services/search/rpc/search"
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
	key := in.EntityId + ":" + in.Metric + ":" + strings.TrimSpace(string(rune(in.UserId)))
	if in.UserId == 99 {
		key = in.EntityId + ":" + in.Metric + ":99"
	}
	return &counterpb.IsMarkedResp{Marked: s.marks[key]}, nil
}
func (s stubCounterClient) BatchGetCounts(context.Context, *counterpb.BatchGetCountsReq, ...grpc.CallOption) (*counterpb.BatchGetCountsResp, error) {
	panic("not implemented")
}

func newTestSvc(t *testing.T, h http.HandlerFunc) *svc.ServiceContext {
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
	}
}

func TestSearch_EmptyQueryReturnsEmpty(t *testing.T) {
	sc := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not call ES")
	})
	resp, err := NewSearchLogic(context.Background(), sc).Search(&searchpb.SearchReq{Q: "   "})
	if err != nil || len(resp.Items) != 0 || resp.HasMore {
		t.Fatalf("resp=%+v err=%v", resp, err)
	}
}

func TestSearch_UsesHighlightAndCounterHydration(t *testing.T) {
	sc := newTestSvc(t, func(w http.ResponseWriter, r *http.Request) {
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
	resp, err := NewSearchLogic(context.Background(), sc).Search(&searchpb.SearchReq{Q: "go", Size: 1, ViewerId: 99})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 1 {
		t.Fatalf("items=%d", len(resp.Items))
	}
	got := resp.Items[0]
	if !strings.Contains(got.Description, "<em>go</em>") || got.LikeCount != 12 || got.FavoriteCount != 4 || !got.Liked || !got.Faved {
		t.Fatalf("unexpected item: %+v", got)
	}
}

var _ = json.RawMessage{}
