package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	"google.golang.org/grpc"
)

type relationFollowClient interface {
	Follow(context.Context, *relationpb.FollowReq, ...grpc.CallOption) (*relationpb.FollowResp, error)
}

type userCounterSnapshotClient interface {
	GetUserSnapshot(context.Context, *counterpb.GetUserSnapshotReq, ...grpc.CallOption) (*counterpb.GetUserSnapshotResp, error)
}

func effectiveCounterSyncTimeout(cardinality string, configured time.Duration) time.Duration {
	if cardinality == readerCardinalityHigh && configured < 15*time.Minute {
		return 15 * time.Minute
	}
	return configured
}

func validateResumableManifest(cfg benchmarkConfig, manifest datasetManifest, strategy string) error {
	if manifest.RunID != cfg.RunID {
		return fmt.Errorf("manifest run id %q does not match configured run id %q", manifest.RunID, cfg.RunID)
	}
	if err := validateManifestStrategy(manifest, strategy); err != nil {
		return err
	}
	if err := validateManifestSeed(manifest, cfg.Topology.Seed); err != nil {
		return err
	}
	if manifest.ReaderCardinality != cfg.ReaderCardinality {
		return fmt.Errorf("manifest reader cardinality %q does not match configured %q", manifest.ReaderCardinality, cfg.ReaderCardinality)
	}
	if len(manifest.Users) != cfg.Topology.RequiredUsers() {
		return fmt.Errorf("manifest users=%d, want %d", len(manifest.Users), cfg.Topology.RequiredUsers())
	}
	if len(manifest.NormalAuthors) != cfg.Topology.NormalAuthors || len(manifest.BigVAuthors) != cfg.Topology.BigVAuthors {
		return fmt.Errorf("manifest authors=%d/%d, want %d/%d", len(manifest.NormalAuthors), len(manifest.BigVAuthors), cfg.Topology.NormalAuthors, cfg.Topology.BigVAuthors)
	}
	if len(manifest.Readers) != cfg.Topology.Readers {
		return fmt.Errorf("manifest readers=%d, want %d", len(manifest.Readers), cfg.Topology.Readers)
	}
	for _, authorID := range append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...) {
		if manifest.FollowerCounts[authorID] <= 0 {
			return fmt.Errorf("manifest author %d has no expected follower count", authorID)
		}
	}
	return nil
}

func finalizeDatasetSetup(
	ctx context.Context,
	cfg benchmarkConfig,
	counter userCounterSnapshotClient,
	publisher publisherClient,
	manifest datasetManifest,
	manifestPath string,
) (datasetManifest, error) {
	syncTimeout := effectiveCounterSyncTimeout(cfg.ReaderCardinality, cfg.Setup.CounterSyncTimeout)
	syncCtx, cancel := context.WithTimeout(ctx, syncTimeout)
	defer cancel()
	if err := waitForFollowerCounts(syncCtx, counter, manifest.FollowerCounts, cfg.Setup.CounterPollInterval); err != nil {
		return manifest, err
	}

	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	wantPosts := len(authors) * cfg.Setup.WarmupPosts
	if cfg.Setup.WarmupPosts > 0 && len(manifest.Posts) == 0 {
		stages, posts, err := runPublishStage(ctx, publisher, authors, publishSpec{
			Requests:       wantPosts,
			Concurrency:    minInt(cfg.Setup.WarmupConcurrency, len(authors)),
			RequestTimeout: cfg.Load.RequestTimeout,
			RunID:          cfg.RunID + "-warmup",
			Seed:           cfg.Topology.Seed,
		})
		if err != nil {
			return manifest, err
		}
		if stages[0].Failed > 0 {
			return manifest, fmt.Errorf("warmup publish failures: %d", stages[0].Failed)
		}
		manifest.Posts = posts
	} else if cfg.Setup.WarmupPosts > 0 && len(manifest.Posts) != wantPosts {
		return manifest, fmt.Errorf("manifest warmup posts=%d, want 0 or %d", len(manifest.Posts), wantPosts)
	}
	if err := saveManifest(manifestPath, manifest); err != nil {
		return manifest, err
	}
	return manifest, nil
}

func applyTopology(
	ctx context.Context,
	client relationFollowClient,
	plan *topologyPlan,
	concurrency int,
) (int, error) {
	if plan == nil {
		return 0, fmt.Errorf("topology plan is required")
	}
	if concurrency <= 0 {
		return 0, fmt.Errorf("relation concurrency must be positive")
	}

	type edge struct {
		from int64
		to   int64
	}
	jobs := make(chan edge)
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	var successful atomic.Int64
	var firstErr error
	var errMu sync.Mutex
	var workers sync.WaitGroup
	workers.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer workers.Done()
			for item := range jobs {
				_, err := client.Follow(runCtx, &relationpb.FollowReq{
					FromUserId: item.from,
					ToUserId:   item.to,
				})
				if err != nil {
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("follow %d -> %d: %w", item.from, item.to, err)
						cancel()
					}
					errMu.Unlock()
					continue
				}
				successful.Add(1)
			}
		}()
	}

sendEdges:
	for from, authors := range plan.FollowingByUser {
		for _, to := range authors {
			select {
			case <-runCtx.Done():
				break sendEdges
			case jobs <- edge{from: from, to: to}:
			}
		}
	}
	close(jobs)
	workers.Wait()
	if firstErr != nil {
		return int(successful.Load()), firstErr
	}
	return int(successful.Load()), nil
}

func waitForFollowerCounts(
	ctx context.Context,
	client userCounterSnapshotClient,
	expected map[int64]int,
	pollInterval time.Duration,
) error {
	if pollInterval <= 0 {
		return fmt.Errorf("poll interval must be positive")
	}
	if len(expected) == 0 {
		return nil
	}

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	lastState := "counter state not checked"
	for {
		ready := true
		for userID, want := range expected {
			resp, err := client.GetUserSnapshot(ctx, &counterpb.GetUserSnapshotReq{UserId: userID})
			if err != nil {
				ready = false
				lastState = fmt.Sprintf("user %d counter error: %v", userID, err)
				break
			}
			got := int64(0)
			if resp != nil && resp.Snapshot != nil {
				got = resp.Snapshot.Followers
			}
			if got < int64(want) {
				ready = false
				lastState = fmt.Sprintf("user %d followers=%d want>=%d", userID, got, want)
				break
			}
		}
		if ready {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for follower counters: %s: %w", lastState, ctx.Err())
		case <-ticker.C:
		}
	}
}
