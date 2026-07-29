package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type recordingUserCreator struct {
	requests []*userpb.CreateReq
	nextID   int64
	err      error
}

func (c *recordingUserCreator) Create(
	_ context.Context,
	req *userpb.CreateReq,
	_ ...grpc.CallOption,
) (*userpb.CreateResp, error) {
	c.requests = append(c.requests, req)
	if c.err != nil {
		return nil, c.err
	}
	c.nextID++
	return &userpb.CreateResp{Id: c.nextID}, nil
}

func TestCreateTestUsersReturnsRealUniqueUserIDs(t *testing.T) {
	client := &recordingUserCreator{nextID: 100}

	ids, err := createTestUsers(context.Background(), client, 3, "run-42")

	require.NoError(t, err)
	require.Equal(t, []int64{101, 102, 103}, ids)
	require.Len(t, client.requests, 3)
	require.NotEmpty(t, client.requests[0].Email)
	require.NotEqual(t, client.requests[0].Email, client.requests[1].Email)
	require.NotEmpty(t, client.requests[0].Nickname)
}

func TestCreateTestUsersStopsOnRPCFailure(t *testing.T) {
	client := &recordingUserCreator{err: errors.New("user rpc unavailable")}

	ids, err := createTestUsers(context.Background(), client, 3, "run-42")

	require.ErrorContains(t, err, "create load-test user 0")
	require.Nil(t, ids)
	require.Len(t, client.requests, 1)
}

func TestNewConfirmContentRequestIncludesObjectKey(t *testing.T) {
	req := newConfirmContentRequest(123, 456)

	require.Equal(t, int64(123), req.Id)
	require.Equal(t, int64(456), req.CreatorId)
	require.Equal(t, "loadtest/123/content.md", req.ObjectKey)
	require.Positive(t, req.Size)
}

func TestValidateLoadTestResults(t *testing.T) {
	config := &LoadTestConfig{PostCount: 2, ReadCount: 3}

	require.NoError(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   2,
		PostSuccess: 2,
		ReadTotal:   3,
		ReadSuccess: 3,
	}))

	require.ErrorContains(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   2,
		PostSuccess: 1,
		PostFail:    1,
		ReadTotal:   3,
		ReadSuccess: 3,
	}), "post failures: 1")

	require.ErrorContains(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   2,
		PostSuccess: 2,
		ReadTotal:   3,
		ReadSuccess: 2,
		ReadFail:    1,
	}), "read failures: 1")

	require.ErrorContains(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   1,
		PostSuccess: 1,
		ReadTotal:   3,
		ReadSuccess: 3,
	}), "post requests: got 1, want 2")

	require.ErrorContains(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   2,
		PostSuccess: 1,
		ReadTotal:   3,
		ReadSuccess: 3,
	}), "post successes: got 1, want 2")

	require.ErrorContains(t, validateLoadTestResults(config, &LoadTestStats{
		PostTotal:   2,
		PostSuccess: 2,
		ReadTotal:   3,
		ReadSuccess: 2,
	}), "read successes: got 2, want 3")
}

func TestWaitForLoadTest(t *testing.T) {
	done := make(chan struct{})
	close(done)
	require.NoError(t, waitForLoadTest(done, time.Second))

	require.ErrorContains(t, waitForLoadTest(make(chan struct{}), time.Millisecond), "timed out")
}
