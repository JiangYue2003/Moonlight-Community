package processor

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/svc"
)

// fakeKnowPostClient 实现 knowpostpb.KnowPostClient（仅 GetDetail / GetPublicFeed）。
type fakeKnowPostClient struct {
	knowpostpb.KnowPostClient
	detailFn func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error)
}

func (f *fakeKnowPostClient) GetDetail(ctx context.Context, in *knowpostpb.GetDetailReq, opts ...grpc.CallOption) (*knowpostpb.KnowPostDetail, error) {
	return f.detailFn(ctx, in)
}

// 起一个 ES mock + miniredis + content-host mock。
func setup(t *testing.T, esHandler http.HandlerFunc, contentHandler http.HandlerFunc, detailFn func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error)) (*Processor, *svc.ServiceContext) {
	t.Helper()
	es := httptest.NewServer(esHandler)
	t.Cleanup(es.Close)
	cnt := httptest.NewServer(contentHandler)
	t.Cleanup(cnt.Close)
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)

	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{Addrs: []string{mr.Addr()}})
	esc, _ := esx.New(esx.Config{Addrs: []string{es.URL}, Timeout: 2 * time.Second})

	sc := &svc.ServiceContext{
		Config: config.Config{
			ContentIndex:    "idx",
			Dedup:           config.DedupConf{TtlSeconds: 600},
			ContentMaxRunes: 4000,
		},
		Es:          esc,
		KnowPostRpc: &fakeKnowPostClient{detailFn: detailFn},
		Redis:       rdb,
		HttpClient:  &http.Client{Timeout: 2 * time.Second},
	}
	// 把测试用 contentUrl host 替换进 detailFn 已经处理；返回 sc + Processor 即可
	_ = cnt
	return New(sc), sc
}

// canalEnvelope 构造 Canal flatMessage 包裹一条 outbox 行。
func canalEnvelope(rowId int64, evType string, ev event.KnowPostEvent) []byte {
	payload, _ := json.Marshal(ev)
	row := map[string]string{
		"id":             "1001",
		"aggregate_type": event.AggregateType,
		"aggregate_id":   "1",
		"type":           evType,
		"payload":        string(payload),
		"created_at":     "2025-01-01 00:00:00",
	}
	_ = rowId
	flat := map[string]any{
		"database": "zhiguang",
		"table":    "outbox",
		"type":     "INSERT",
		"isDdl":    false,
		"data":     []map[string]string{row},
	}
	buf, _ := json.Marshal(flat)
	return buf
}

func TestUpsert_PublishedPublic(t *testing.T) {
	var indexedDoc string
	esCalls := atomic.Int32{}
	contentSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "# 标题\n这是正文")
	}))
	defer contentSrv.Close()
	esHandler := func(w http.ResponseWriter, r *http.Request) {
		esCalls.Add(1)
		if r.Method == "PUT" && strings.Contains(r.URL.Path, "/idx/_doc/123") {
			buf, _ := io.ReadAll(r.Body)
			indexedDoc = string(buf)
			_, _ = io.WriteString(w, `{"result":"created"}`)
			return
		}
		t.Fatalf("unexpected ES call: %s %s", r.Method, r.URL.Path)
	}
	proc, _ := setup(t, esHandler, nil,
		func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error) {
			return &knowpostpb.KnowPostDetail{
				Id: "123", Title: "T", Description: "D", Status: "published", Visible: "public",
				ContentUrl: contentSrv.URL, Tags: []string{"go"}, CreatorId: 7, PublishTime: 1700000000,
			}, nil
		},
	)
	msg := canalEnvelope(1001, "INSERT", event.KnowPostEvent{
		Type: event.TypeKnowPostPublished, PostId: 123, Author: 7,
	})
	if err := proc.Handle(context.Background(), msg); err != nil {
		t.Fatal(err)
	}
	if esCalls.Load() != 1 {
		t.Fatalf("ES calls = %d", esCalls.Load())
	}
	if !strings.Contains(indexedDoc, `"title":"T"`) {
		t.Fatalf("doc: %s", indexedDoc)
	}
	if !strings.Contains(indexedDoc, `"这是正文"`) && !strings.Contains(indexedDoc, "这是正文") {
		t.Fatalf("doc body missing: %s", indexedDoc)
	}
}

func TestUpsert_DraftFallsbackToSoftDelete(t *testing.T) {
	var gotPath string
	esHandler := func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.Method + " " + r.URL.Path
		_, _ = io.WriteString(w, `{"result":"updated"}`)
	}
	proc, _ := setup(t, esHandler, nil,
		func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error) {
			return &knowpostpb.KnowPostDetail{Id: "200", Status: "draft", Visible: "private"}, nil
		},
	)
	msg := canalEnvelope(1002, "INSERT", event.KnowPostEvent{
		Type: event.TypeKnowPostUpdated, PostId: 200,
	})
	if err := proc.Handle(context.Background(), msg); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(gotPath, "POST /idx/_update/200") {
		t.Fatalf("expected SoftDelete path, got %s", gotPath)
	}
}

func TestDedup_DuplicateSkipped(t *testing.T) {
	calls := atomic.Int32{}
	esHandler := func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		_, _ = io.WriteString(w, `{}`)
	}
	contentSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "x")
	}))
	defer contentSrv.Close()
	proc, _ := setup(t, esHandler, nil,
		func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error) {
			return &knowpostpb.KnowPostDetail{Id: "1", Status: "published", Visible: "public", ContentUrl: contentSrv.URL}, nil
		},
	)
	msg := canalEnvelope(1003, "INSERT", event.KnowPostEvent{
		Type: event.TypeKnowPostPublished, PostId: 1,
	})
	if err := proc.Handle(context.Background(), msg); err != nil {
		t.Fatal(err)
	}
	if err := proc.Handle(context.Background(), msg); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 {
		t.Errorf("ES called %d times, want 1", calls.Load())
	}
}

func TestNonKnowpost_Skipped(t *testing.T) {
	esHandler := func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not call ES")
	}
	proc, _ := setup(t, esHandler, nil,
		func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error) {
			t.Fatal("should not call rpc")
			return nil, nil
		},
	)
	flat := map[string]any{
		"database": "zhiguang", "table": "outbox", "type": "INSERT", "isDdl": false,
		"data": []map[string]string{{
			"id": "1", "aggregate_type": "following", "aggregate_id": "1", "type": "FollowCreated", "payload": "{}",
		}},
	}
	buf, _ := json.Marshal(flat)
	if err := proc.Handle(context.Background(), buf); err != nil {
		t.Fatal(err)
	}
}

func TestParseFlatBadJson(t *testing.T) {
	// 防御坏 JSON 不应阻塞 group 进度
	contentSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer contentSrv.Close()
	proc, _ := setup(t, func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not call ES")
	}, nil, func(ctx context.Context, req *knowpostpb.GetDetailReq) (*knowpostpb.KnowPostDetail, error) {
		t.Fatal("should not call rpc")
		return nil, nil
	})
	if err := proc.Handle(context.Background(), []byte("not json")); err != nil {
		t.Fatalf("expect nil, got %v", err)
	}
	_ = canalx.OutboxRow{}
}
