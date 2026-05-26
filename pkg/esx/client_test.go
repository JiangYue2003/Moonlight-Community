package esx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// helper：起 mock ES，路由器。
func mockES(t *testing.T, h func(w http.ResponseWriter, r *http.Request)) (*Client, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(h))
	t.Cleanup(srv.Close)
	c, err := New(Config{Addrs: []string{srv.URL}})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c, srv
}

func TestEnsureIndex_Idempotent(t *testing.T) {
	calls := map[string]int{}
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		calls[r.Method+" "+r.URL.Path]++
		switch {
		case r.Method == "HEAD" && r.URL.Path == "/foo":
			// 第一次 404，第二次 200
			if calls["HEAD /foo"] == 1 {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		case r.Method == "PUT" && r.URL.Path == "/foo":
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, `{"acknowledged":true}`)
		default:
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
	})
	ctx := context.Background()
	if err := c.EnsureIndex(ctx, "foo", json.RawMessage(`{}`)); err != nil {
		t.Fatalf("first EnsureIndex: %v", err)
	}
	if err := c.EnsureIndex(ctx, "foo", json.RawMessage(`{}`)); err != nil {
		t.Fatalf("second EnsureIndex: %v", err)
	}
	if calls["PUT /foo"] != 1 {
		t.Errorf("PUT /foo expected once, got %d", calls["PUT /foo"])
	}
}

func TestSearch_ParsesHits(t *testing.T) {
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/idx/_search" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = io.WriteString(w, `{
			"took":3,
			"hits":{"total":{"value":1},"hits":[
				{"_id":"123","_score":1.5,"_source":{"title":"hello"},"highlight":{"title":["he<em>llo</em>"]},"sort":[1.5,1700000000000,123]}
			]}
		}`)
	})
	res, err := c.Search(context.Background(), "idx", json.RawMessage(`{}`))
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if len(res.Hits.Hits) != 1 || res.Hits.Hits[0].Id != "123" {
		t.Fatalf("hits: %+v", res.Hits.Hits)
	}
	if res.Hits.Hits[0].Highlight["title"][0] != "he<em>llo</em>" {
		t.Fatalf("highlight: %+v", res.Hits.Hits[0].Highlight)
	}
	if len(res.Hits.Hits[0].Sort) != 3 {
		t.Fatalf("sort len: %d", len(res.Hits.Hits[0].Sort))
	}
}

func TestSuggest_Dedupe(t *testing.T) {
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{
			"suggest":{"s":[{"text":"go","options":[
				{"text":"go-zero","_score":1},
				{"text":"go语言","_score":0.9},
				{"text":"go-zero","_score":0.8}
			]}]}
		}`)
	})
	got, err := c.Suggest(context.Background(), "idx", "title_suggest", "go", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "go-zero" || got[1] != "go语言" {
		t.Fatalf("got: %v", got)
	}
}

func TestKnnSearch_BuildsBody(t *testing.T) {
	var gotBody string
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		_, _ = io.WriteString(w, `{"hits":{"total":{"value":0},"hits":[]}}`)
	})
	_, err := c.KnnSearch(context.Background(), "idx", "embedding", []float32{1, 2, 3}, 5, 20, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(gotBody, `"knn"`) || !strings.Contains(gotBody, `"k":5`) || !strings.Contains(gotBody, `"num_candidates":20`) {
		t.Fatalf("body: %s", gotBody)
	}
}

func TestError_4xxPropagated(t *testing.T) {
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, `{"error":"bad"}`)
	})
	_, err := c.Search(context.Background(), "idx", json.RawMessage(`{}`))
	if err == nil {
		t.Fatal("expected error")
	}
	var e *Error
	if !asErr(err, &e) || e.Status != 400 {
		t.Fatalf("got: %v", err)
	}
}

func TestDelete_NotFoundIsNoop(t *testing.T) {
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"result":"not_found"}`)
	})
	if err := c.Delete(context.Background(), "idx", "x"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
}

func TestCount_NotFoundReturnsZero(t *testing.T) {
	c, _ := mockES(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{}`)
	})
	n, err := c.Count(context.Background(), "missing")
	if err != nil || n != 0 {
		t.Fatalf("Count: n=%d err=%v", n, err)
	}
}

func TestMappingsEmbedded(t *testing.T) {
	if len(ContentMapping()) < 100 {
		t.Fatal("content mapping empty")
	}
	if len(RagMapping()) < 50 {
		t.Fatal("rag mapping empty")
	}
}

// asErr 是 errors.As 的薄包装。
func asErr(err error, target any) bool { return errors.As(err, target) }
