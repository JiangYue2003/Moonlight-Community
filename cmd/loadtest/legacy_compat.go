package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
)

// LoadTestConfig and LoadTestStats are retained for compatibility with the
// original validation helpers while the CLI uses benchmarkConfig.
type LoadTestConfig struct {
	PostCount int
	ReadCount int
}

type LoadTestStats struct {
	PostTotal   int64
	PostSuccess int64
	PostFail    int64
	ReadTotal   int64
	ReadSuccess int64
	ReadFail    int64
	mu          sync.Mutex
}

type userCreator interface {
	Create(context.Context, *userpb.CreateReq, ...grpc.CallOption) (*userpb.CreateResp, error)
}

func createTestUsers(ctx context.Context, client userCreator, count int, runID string) ([]int64, error) {
	userIDs := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		resp, err := client.Create(ctx, &userpb.CreateReq{
			Email:    fmt.Sprintf("loadtest+%s-%d@example.invalid", runID, i),
			Nickname: fmt.Sprintf("Load Test User %d", i),
		})
		if err != nil {
			return nil, fmt.Errorf("create load-test user %d: %w", i, err)
		}
		if resp == nil || resp.Id <= 0 {
			return nil, fmt.Errorf("create load-test user %d: invalid user id", i)
		}
		userIDs = append(userIDs, resp.Id)
	}
	return userIDs, nil
}

func newConfirmContentRequest(postID, creatorID int64) *knowpostpb.ConfirmContentReq {
	return &knowpostpb.ConfirmContentReq{
		Id:        postID,
		CreatorId: creatorID,
		ObjectKey: fmt.Sprintf("loadtest/%d/content.md", postID),
		Size:      1,
	}
}

func validateLoadTestResults(config *LoadTestConfig, stats *LoadTestStats) error {
	postTotal := atomic.LoadInt64(&stats.PostTotal)
	postSuccess := atomic.LoadInt64(&stats.PostSuccess)
	postFail := atomic.LoadInt64(&stats.PostFail)
	readTotal := atomic.LoadInt64(&stats.ReadTotal)
	readSuccess := atomic.LoadInt64(&stats.ReadSuccess)
	readFail := atomic.LoadInt64(&stats.ReadFail)
	if postTotal != int64(config.PostCount) {
		return fmt.Errorf("post requests: got %d, want %d", postTotal, config.PostCount)
	}
	if postFail > 0 {
		return fmt.Errorf("post failures: %d", postFail)
	}
	if postSuccess != int64(config.PostCount) {
		return fmt.Errorf("post successes: got %d, want %d", postSuccess, config.PostCount)
	}
	if readTotal != int64(config.ReadCount) {
		return fmt.Errorf("read requests: got %d, want %d", readTotal, config.ReadCount)
	}
	if readFail > 0 {
		return fmt.Errorf("read failures: %d", readFail)
	}
	if readSuccess != int64(config.ReadCount) {
		return fmt.Errorf("read successes: got %d, want %d", readSuccess, config.ReadCount)
	}
	return nil
}

func waitForLoadTest(done <-chan struct{}, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-done:
		return nil
	case <-timer.C:
		return fmt.Errorf("load test timed out after %s", timeout)
	}
}
