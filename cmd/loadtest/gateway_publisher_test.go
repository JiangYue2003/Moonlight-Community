package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

func TestGatewayPublisherUsesAuthenticatedFullPublishWorkflow(t *testing.T) {
	var mu sync.Mutex
	var calls []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer token-7", r.Header.Get("Authorization"))
		mu.Lock()
		calls = append(calls, r.Method+" "+r.URL.Path)
		mu.Unlock()
		switch r.Method + " " + r.URL.Path {
		case "POST /api/v1/knowposts/drafts":
			_ = json.NewEncoder(w).Encode(map[string]string{"id": "123"})
		case "PATCH /api/v1/knowposts/123":
			var body map[string]any
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			require.Equal(t, "title", body["title"])
			_ = json.NewEncoder(w).Encode(map[string]any{"id": "123"})
		case "POST /api/v1/knowposts/123/content/confirm":
			w.WriteHeader(http.StatusNoContent)
		case "POST /api/v1/knowposts/123/publish":
			_ = json.NewEncoder(w).Encode(map[string]any{"id": "123"})
		default:
			http.Error(w, fmt.Sprintf("unexpected route %s %s", r.Method, r.URL.Path), http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)
	client := newGatewayPublisherClient(server.URL, server.Client(), map[int64]string{7: "token-7"})
	ctx := context.Background()

	draft, err := client.CreateDraft(ctx, &knowpostpb.CreateDraftReq{CreatorId: 7})
	require.NoError(t, err)
	require.Equal(t, "123", draft.Id)
	_, err = client.PatchMetadata(ctx, &knowpostpb.PatchMetadataReq{Id: 123, CreatorId: 7, Title: "title", TitleSet: true})
	require.NoError(t, err)
	_, err = client.ConfirmContent(ctx, &knowpostpb.ConfirmContentReq{Id: 123, CreatorId: 7, ObjectKey: "loadtest/key", Size: 1})
	require.NoError(t, err)
	_, err = client.Publish(ctx, &knowpostpb.PublishReq{Id: 123, CreatorId: 7})
	require.NoError(t, err)

	require.Equal(t, []string{
		"POST /api/v1/knowposts/drafts",
		"PATCH /api/v1/knowposts/123",
		"POST /api/v1/knowposts/123/content/confirm",
		"POST /api/v1/knowposts/123/publish",
	}, calls)
}

func TestGatewayPublisherRejectsMissingAuthorToken(t *testing.T) {
	client := newGatewayPublisherClient("http://127.0.0.1", nil, nil)

	_, err := client.CreateDraft(context.Background(), &knowpostpb.CreateDraftReq{CreatorId: 7})

	require.ErrorContains(t, err, "access token")
}
