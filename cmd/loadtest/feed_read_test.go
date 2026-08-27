package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errTestFailure = errors.New("request failed")

func TestSummarizeLatenciesUsesRecordedSamples(t *testing.T) {
	samples := []time.Duration{
		9 * time.Millisecond,
		1 * time.Millisecond,
		5 * time.Millisecond,
		3 * time.Millisecond,
		7 * time.Millisecond,
	}

	got := summarizeLatencies(samples)

	require.Equal(t, time.Millisecond, got.Min)
	require.Equal(t, 5*time.Millisecond, got.Mean)
	require.Equal(t, 5*time.Millisecond, got.P50)
	require.Equal(t, 9*time.Millisecond, got.P90)
	require.Equal(t, 9*time.Millisecond, got.P95)
	require.Equal(t, 9*time.Millisecond, got.P99)
	require.Equal(t, 9*time.Millisecond, got.Max)
}

func TestStageRecorderSeparatesFailuresAndTimeouts(t *testing.T) {
	recorder := newStageRecorder("read")
	recorder.Record(10*time.Millisecond, nil)
	recorder.Record(20*time.Millisecond, context.DeadlineExceeded)
	recorder.Record(30*time.Millisecond, errTestFailure)

	result := recorder.Result(time.Second)

	require.Equal(t, int64(3), result.Total)
	require.Equal(t, int64(1), result.Success)
	require.Equal(t, int64(2), result.Failed)
	require.Equal(t, int64(1), result.Timeouts)
	require.Equal(t, 1.0, result.SuccessQPS)
	require.Equal(t, 10*time.Millisecond, result.Latency.Min)
	require.Equal(t, 10*time.Millisecond, result.Latency.Max)
}

func TestErrorClassRecognizesCircuitBreaker(t *testing.T) {
	err := fmt.Errorf("patch metadata: rpc error: code = Unavailable desc = circuit breaker is open")

	require.Equal(t, "circuit_breaker", errorClass(err))
}

func TestStageRecorderRecognizesGRPCDeadlineExceeded(t *testing.T) {
	recorder := newStageRecorder("publish")

	recorder.Record(time.Second, status.Error(codes.DeadlineExceeded, "deadline"))
	result := recorder.Result(time.Second)

	require.Equal(t, int64(1), result.Timeouts)
	require.Equal(t, int64(1), result.Errors["timeout"])
}

func TestGatewayFeedClientUsesFollowingFeedAndBearerToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/v1/knowposts/following-feed", r.URL.Path)
		require.Empty(t, r.URL.Query().Get("page"))
		require.Equal(t, "10", r.URL.Query().Get("size"))
		require.Equal(t, "current-cursor", r.URL.Query().Get("cursor"))
		require.Equal(t, "Bearer real-token", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[{"id":"99"}],"hasMore":true,"page":0,"size":10,"nextCursor":"next-cursor"}`))
	}))
	defer server.Close()

	client := newGatewayFeedClient(server.URL, server.Client())
	page, err := client.GetUserFeed(context.Background(), readerIdentity{
		UserID:      42,
		AccessToken: "real-token",
	}, feedReadRequest{Page: 2, Size: 10, Cursor: "current-cursor"})

	require.NoError(t, err)
	require.Len(t, page.Items, 1)
	require.Equal(t, "99", page.Items[0].ID)
	require.Equal(t, "next-cursor", page.NextCursor)
}
