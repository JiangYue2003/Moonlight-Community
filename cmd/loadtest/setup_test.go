package main

import (
	"context"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"google.golang.org/grpc"
)

type recordingRelationClient struct {
	mu       sync.Mutex
	requests []*relationpb.FollowReq
}

func (c *recordingRelationClient) Follow(
	_ context.Context,
	req *relationpb.FollowReq,
	_ ...grpc.CallOption,
) (*relationpb.FollowResp, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.requests = append(c.requests, req)
	return &relationpb.FollowResp{Changed: true}, nil
}

func TestApplyTopologyCreatesEveryPlannedEdge(t *testing.T) {
	client := &recordingRelationClient{}
	plan := &topologyPlan{FollowingByUser: map[int64][]int64{
		10: {1, 2},
		11: {2},
	}}

	created, err := applyTopology(context.Background(), client, plan, 2)

	require.NoError(t, err)
	require.Equal(t, 3, created)
	require.Len(t, client.requests, 3)
}

type eventuallyConsistentCounter struct {
	mu       sync.Mutex
	counts   map[int64]int64
	requests int
}

func (c *eventuallyConsistentCounter) GetUserSnapshot(
	_ context.Context,
	req *counterpb.GetUserSnapshotReq,
	_ ...grpc.CallOption,
) (*counterpb.GetUserSnapshotResp, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.requests++
	count := int64(0)
	if c.requests > 1 {
		count = c.counts[req.UserId]
	}
	return &counterpb.GetUserSnapshotResp{Snapshot: &counterpb.UserSnapshot{Followers: count}}, nil
}

func TestWaitForFollowerCountsPollsVerifiedState(t *testing.T) {
	client := &eventuallyConsistentCounter{counts: map[int64]int64{1: 1001}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := waitForFollowerCounts(ctx, client, map[int64]int{1: 1001}, time.Millisecond)

	require.NoError(t, err)
	require.GreaterOrEqual(t, client.requests, 2)
}

func TestEffectiveCounterSyncTimeoutExtendsHighCardinality(t *testing.T) {
	require.Equal(t, 15*time.Minute, effectiveCounterSyncTimeout(readerCardinalityHigh, 3*time.Minute))
	require.Equal(t, 20*time.Minute, effectiveCounterSyncTimeout(readerCardinalityHigh, 20*time.Minute))
	require.Equal(t, 3*time.Minute, effectiveCounterSyncTimeout(readerCardinalityDistributed, 3*time.Minute))
}

func TestValidateResumableManifestRequiresCompleteOwnedTopology(t *testing.T) {
	cfg := benchmarkConfig{
		RunID: "resume-run", ReaderCardinality: readerCardinalityHigh,
		Topology: topologyConfig{FollowerUsers: 2, NormalAuthors: 1, BigVAuthors: 1, Readers: 2, Seed: 42},
	}
	manifest := datasetManifest{
		RunID: "resume-run", Seed: 42, Strategy: "hybrid", ReaderCardinality: readerCardinalityHigh,
		Users:         []benchmarkUser{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
		NormalAuthors: []int64{1}, BigVAuthors: []int64{2}, Readers: []int64{3, 4},
		FollowerCounts: map[int64]int{1: 1, 2: 2},
	}

	require.NoError(t, validateResumableManifest(cfg, manifest, "hybrid"))
	manifest.Users = manifest.Users[:3]
	require.ErrorContains(t, validateResumableManifest(cfg, manifest, "hybrid"), "users")
}

func TestFinalizeDatasetSetupPublishesAndPersistsMissingWarmup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.json")
	manifest := datasetManifest{
		RunID: "resume-run", Seed: 42, Strategy: "hybrid", ReaderCardinality: readerCardinalityHigh,
		Users:         []benchmarkUser{{ID: 7}, {ID: 8}},
		NormalAuthors: []int64{7}, BigVAuthors: []int64{8},
		FollowerCounts: map[int64]int{7: 100, 8: 1001},
	}
	counter := &eventuallyConsistentCounter{counts: map[int64]int64{7: 100, 8: 1001}, requests: 1}
	publisher := &recordingPublisher{nextID: 300}
	cfg := benchmarkConfig{
		ReaderCardinality: readerCardinalityHigh,
		Setup:             setupConfig{CounterSyncTimeout: time.Second, CounterPollInterval: time.Millisecond, WarmupPosts: 2, WarmupConcurrency: 1},
		Load:              benchmarkLoadConfig{RequestTimeout: time.Second},
		Topology:          topologyConfig{Seed: 42},
	}

	got, err := finalizeDatasetSetup(context.Background(), cfg, counter, publisher, manifest, path)

	require.NoError(t, err)
	require.Len(t, got.Posts, 4)
	require.Len(t, publisher.published, 4)
	saved, err := loadManifest(path)
	require.NoError(t, err)
	require.Equal(t, got.Posts, saved.Posts)
}

func TestFinalizeDatasetSetupDoesNotRepublishCompletedWarmup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.json")
	manifest := datasetManifest{
		RunID: "resume-run", Seed: 42, Strategy: "hybrid", ReaderCardinality: readerCardinalityHigh,
		Users:         []benchmarkUser{{ID: 7}, {ID: 8}},
		NormalAuthors: []int64{7}, BigVAuthors: []int64{8}, Posts: []int64{96, 97, 98, 99},
		FollowerCounts: map[int64]int{7: 100, 8: 1001},
	}
	counter := &eventuallyConsistentCounter{counts: map[int64]int64{7: 100, 8: 1001}, requests: 1}
	publisher := &recordingPublisher{nextID: 300}
	cfg := benchmarkConfig{
		ReaderCardinality: readerCardinalityHigh,
		Setup:             setupConfig{CounterSyncTimeout: time.Second, CounterPollInterval: time.Millisecond, WarmupPosts: 2, WarmupConcurrency: 1},
		Load:              benchmarkLoadConfig{RequestTimeout: time.Second},
	}

	got, err := finalizeDatasetSetup(context.Background(), cfg, counter, publisher, manifest, path)

	require.NoError(t, err)
	require.Equal(t, []int64{96, 97, 98, 99}, got.Posts)
	require.Empty(t, publisher.published)
}
