//go:build integration

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

const contentSafetyEpochKey = "feed:content:safety:epoch"

func TestContentSafetyMutationsInvalidateWarmPersonalFeedPage(t *testing.T) {
	if os.Getenv("FEED_CONTENT_SAFETY_INTEGRATION") != "1" {
		t.Skip("set FEED_CONTENT_SAFETY_INTEGRATION=1 to exercise the dev stack")
	}

	repoRoot := filepath.Clean(filepath.Join("..", ".."))
	var cfg benchmarkConfig
	configPath := filepath.Join(repoRoot, "cmd", "loadtest", "load_test.yaml")
	if err := conf.Load(configPath, &cfg, conf.UseEnv()); err != nil {
		t.Fatalf("load config %s: %v", configPath, err)
	}
	applyConfigDefaults(&cfg)
	manifest, err := loadManifest(filepath.Join(repoRoot, "results", "feed-loadtest", "manifests", "feedbench-20260814c.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(manifest.BigVAuthors) == 0 || len(manifest.Readers) == 0 {
		t.Fatal("integration manifest requires at least one big-V author and reader")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	flags, err := readContainerFeedFeatureFlags(ctx, cfg.Monitor.StrategyContainer)
	if err != nil {
		t.Fatal(err)
	}
	if flags.PageCacheMode != "l1-l2" {
		t.Fatalf("content-safety integration requires FEED_PAGE_CACHE_MODE=l1-l2, got %q", flags.PageCacheMode)
	}
	if !flags.SafetyEpochEnabled {
		t.Fatal("content-safety integration requires FEED_CONTENT_SAFETY_EPOCH_ENABLED=true")
	}
	strategy, err := readContainerFeedStrategy(ctx, cfg.Monitor.StrategyContainer)
	if err != nil {
		t.Fatal(err)
	}
	if strategy != "hybrid" {
		t.Fatalf("content-safety integration requires hybrid strategy, got %q", strategy)
	}

	logx.DisableStat()
	clients := connectRPCs(cfg)
	rdb := goredis.NewClient(&goredis.Options{Addr: cfg.RedisAddr})
	t.Cleanup(func() { _ = rdb.Close() })
	authorID := manifest.BigVAuthors[0]
	readerID := manifest.Readers[0]
	postID := publishContentSafetyPost(t, ctx, clients.knowpost, authorID)
	deleted := false
	t.Cleanup(func() {
		if deleted {
			return
		}
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cleanupCancel()
		_, _ = clients.knowpost.Delete(cleanupCtx, &knowpostpb.DeleteReq{Id: postID, CreatorId: authorID})
	})

	if _, err := waitForPostState(ctx, clients.knowpost, readerID, postID, true, 12*time.Second); err != nil {
		t.Fatal(err)
	}
	warmPersonalFeedPage(t, ctx, clients.knowpost, readerID, postID)

	beforePrivate := readContentSafetyEpoch(t, ctx, rdb)
	started := time.Now()
	if _, err := clients.knowpost.UpdateVisibility(ctx, &knowpostpb.UpdateVisibilityReq{
		Id: postID, CreatorId: authorID, Visible: "private",
	}); err != nil {
		t.Fatalf("make post private: %v", err)
	}
	if _, err := waitForPostState(ctx, clients.knowpost, readerID, postID, false, time.Second); err != nil {
		t.Fatal(err)
	}
	privateDelay := time.Since(started)
	afterPrivate, err := waitForContentSafetyEpoch(ctx, rdb, beforePrivate+2, 3*time.Second)
	if err != nil {
		t.Fatalf("wait for durable private safety compensation: %v", err)
	}
	t.Logf("private convergence=%s safety_epoch=%d->%d", privateDelay, beforePrivate, afterPrivate)

	if _, err := clients.knowpost.UpdateVisibility(ctx, &knowpostpb.UpdateVisibilityReq{
		Id: postID, CreatorId: authorID, Visible: "public",
	}); err != nil {
		t.Fatalf("restore post visibility: %v", err)
	}
	if _, err := waitForPostState(ctx, clients.knowpost, readerID, postID, true, 3*time.Second); err != nil {
		t.Fatal(err)
	}
	warmPersonalFeedPage(t, ctx, clients.knowpost, readerID, postID)

	beforeDelete := readContentSafetyEpoch(t, ctx, rdb)
	started = time.Now()
	if _, err := clients.knowpost.Delete(ctx, &knowpostpb.DeleteReq{Id: postID, CreatorId: authorID}); err != nil {
		t.Fatalf("delete post: %v", err)
	}
	deleted = true
	if _, err := waitForPostState(ctx, clients.knowpost, readerID, postID, false, time.Second); err != nil {
		t.Fatal(err)
	}
	deleteDelay := time.Since(started)
	afterDelete, err := waitForContentSafetyEpoch(ctx, rdb, beforeDelete+2, 3*time.Second)
	if err != nil {
		t.Fatalf("wait for durable delete safety compensation: %v", err)
	}
	t.Logf("delete convergence=%s safety_epoch=%d->%d", deleteDelay, beforeDelete, afterDelete)
}

func publishContentSafetyPost(
	t *testing.T,
	ctx context.Context,
	client knowpostpb.KnowPostClient,
	authorID int64,
) int64 {
	t.Helper()
	draft, err := client.CreateDraft(ctx, &knowpostpb.CreateDraftReq{CreatorId: authorID})
	if err != nil {
		t.Fatalf("create content-safety draft: %v", err)
	}
	postID, err := strconv.ParseInt(draft.Id, 10, 64)
	if err != nil {
		t.Fatalf("parse draft id %q: %v", draft.Id, err)
	}
	stamp := time.Now().Format("20060102-150405.000")
	if _, err := client.PatchMetadata(ctx, &knowpostpb.PatchMetadataReq{
		Id: postID, CreatorId: authorID, TitleSet: true, Title: "feed-safety-" + stamp,
	}); err != nil {
		t.Fatalf("patch content-safety metadata: %v", err)
	}
	if _, err := client.ConfirmContent(ctx, &knowpostpb.ConfirmContentReq{
		Id: postID, CreatorId: authorID, ObjectKey: fmt.Sprintf("loadtest/feed-safety/%d/content.md", postID), Size: 1,
	}); err != nil {
		t.Fatalf("confirm content-safety content: %v", err)
	}
	if _, err := client.Publish(ctx, &knowpostpb.PublishReq{Id: postID, CreatorId: authorID}); err != nil {
		t.Fatalf("publish content-safety post: %v", err)
	}
	return postID
}

func warmPersonalFeedPage(
	t *testing.T,
	ctx context.Context,
	client knowpostpb.KnowPostClient,
	readerID, postID int64,
) {
	t.Helper()
	for range 2 {
		page, err := client.GetUserFeed(ctx, &knowpostpb.GetUserFeedReq{UserId: readerID, Page: 1, Size: 20})
		if err != nil {
			t.Fatalf("warm personal feed page: %v", err)
		}
		if !feedPageContainsPost(page, postID) {
			t.Fatalf("warm personal feed page does not contain post %d", postID)
		}
	}
}

func waitForPostState(
	ctx context.Context,
	client knowpostpb.KnowPostClient,
	readerID, postID int64,
	wantPresent bool,
	timeout time.Duration,
) (time.Duration, error) {
	started := time.Now()
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for {
		page, err := client.GetUserFeed(ctx, &knowpostpb.GetUserFeedReq{UserId: readerID, Page: 1, Size: 20})
		if err == nil && feedPageContainsPost(page, postID) == wantPresent {
			return time.Since(started), nil
		}
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-deadline.C:
			return 0, fmt.Errorf("post %d presence did not become %t within %s (last error: %v)", postID, wantPresent, timeout, err)
		case <-ticker.C:
		}
	}
}

func feedPageContainsPost(page *knowpostpb.FeedPage, postID int64) bool {
	if page == nil {
		return false
	}
	want := strconv.FormatInt(postID, 10)
	for _, item := range page.Items {
		if item != nil && item.Id == want {
			return true
		}
	}
	return false
}

func readContentSafetyEpoch(t *testing.T, ctx context.Context, rdb *goredis.Client) uint64 {
	t.Helper()
	epoch, err := rdb.Get(ctx, contentSafetyEpochKey).Uint64()
	if errors.Is(err, goredis.Nil) {
		return 0
	}
	if err != nil {
		t.Fatalf("read content safety epoch: %v", err)
	}
	return epoch
}

func waitForContentSafetyEpoch(
	ctx context.Context,
	rdb *goredis.Client,
	want uint64,
	timeout time.Duration,
) (uint64, error) {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for {
		epoch, err := rdb.Get(ctx, contentSafetyEpochKey).Uint64()
		if err == nil && epoch >= want {
			return epoch, nil
		}
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-deadline.C:
			return 0, fmt.Errorf("safety epoch did not reach %d within %s (last=%d, err=%v)", want, timeout, epoch, err)
		case <-ticker.C:
		}
	}
}
