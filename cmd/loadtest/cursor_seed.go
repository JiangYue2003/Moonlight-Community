package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/redis/go-redis/v9"
)

const cursorDeepTieMembers = 120

func mergeCursorPublishedPosts(manifest *datasetManifest, published []cursorPublishedPost) error {
	if manifest == nil || manifest.CursorDeep == nil {
		return fmt.Errorf("cursor-deep manifest is required")
	}
	if manifest.CursorDeep.PostAuthors == nil {
		manifest.CursorDeep.PostAuthors = make(map[int64]int64)
	}
	seen := make(map[int64]struct{}, len(manifest.Posts)+len(published))
	for _, postID := range manifest.Posts {
		if postID <= 0 {
			return fmt.Errorf("cursor-deep manifest contains non-positive post id")
		}
		if _, duplicate := seen[postID]; duplicate {
			return fmt.Errorf("cursor-deep manifest contains duplicate post %d", postID)
		}
		seen[postID] = struct{}{}
		if manifest.CursorDeep.PostAuthors[postID] <= 0 {
			return fmt.Errorf("cursor-deep manifest post %d has no author", postID)
		}
	}
	ordered := append([]cursorPublishedPost(nil), published...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].RequestNumber < ordered[j].RequestNumber })
	for _, post := range ordered {
		if post.PostID <= 0 || post.AuthorID <= 0 {
			return fmt.Errorf("cursor-deep published post and author ids must be positive")
		}
		if _, duplicate := seen[post.PostID]; duplicate {
			return fmt.Errorf("cursor-deep published duplicate post %d", post.PostID)
		}
		seen[post.PostID] = struct{}{}
		manifest.Posts = append(manifest.Posts, post.PostID)
		manifest.CursorDeep.PostAuthors[post.PostID] = post.AuthorID
	}
	return nil
}

func seedCursorDeepDataset(
	ctx context.Context,
	cfg benchmarkConfig,
	publisher publisherClient,
	manifest datasetManifest,
	manifestPath string,
) (datasetManifest, error) {
	if cfg.ReaderCardinality != readerCardinalityCursorDeep || manifest.ReaderCardinality != readerCardinalityCursorDeep {
		return manifest, fmt.Errorf("seed-cursor-deep requires reader cardinality %q", readerCardinalityCursorDeep)
	}
	if manifest.RunID != cfg.RunID {
		return manifest, fmt.Errorf("manifest run id %q does not match configured run id %q", manifest.RunID, cfg.RunID)
	}
	if err := validateManifestSeed(manifest, cfg.Topology.Seed); err != nil {
		return manifest, err
	}
	if len(manifest.NormalAuthors) != 20 || len(manifest.BigVAuthors) != 5 || len(manifest.Readers) != 20 {
		return manifest, fmt.Errorf("cursor-deep topology requires 20 normal authors, 5 big-v authors, and 20 readers")
	}
	if manifest.CursorDeep == nil {
		manifest.CursorDeep = &cursorDeepDataset{
			Version:                       cursorDeepDatasetVersion,
			NormalPostsPerAuthor:          50,
			BigVPostsPerAuthor:            100,
			InboxCandidatesPerReader:      1000,
			BigVOutboxCandidatesPerAuthor: 100,
			PostAuthors:                   make(map[int64]int64, 1500),
		}
	} else if manifest.CursorDeep.Version != cursorDeepDatasetVersion {
		return manifest, fmt.Errorf("cursor-deep version=%d, want %d", manifest.CursorDeep.Version, cursorDeepDatasetVersion)
	}
	if len(manifest.Posts) != len(manifest.CursorDeep.PostAuthors) {
		return manifest, fmt.Errorf(
			"cursor-deep checkpoint posts=%d ownership=%d; cannot resume an ambiguous partial publish",
			len(manifest.Posts), len(manifest.CursorDeep.PostAuthors),
		)
	}
	manifest.CursorDeep.Ready = false

	jobs, err := buildCursorSeedJobs(manifest.NormalAuthors, manifest.BigVAuthors, manifest.CursorDeep.PostAuthors)
	if err != nil {
		return manifest, err
	}
	if len(jobs) > 0 {
		published, publishErr := publishCursorSeedJobs(
			ctx,
			publisher,
			jobs,
			minInt(cfg.Setup.CursorSeedConcurrency, len(jobs)),
			cfg.Load.RequestTimeout,
			manifest.RunID+"-cursor-deep",
		)
		if err := mergeCursorPublishedPosts(&manifest, published); err != nil {
			return manifest, err
		}
		if err := saveManifest(manifestPath, manifest); err != nil {
			return manifest, err
		}
		if publishErr != nil {
			return manifest, fmt.Errorf("cursor-deep publish checkpointed %d/%d completed jobs: %w", len(published), len(jobs), publishErr)
		}
	}
	if remaining, err := buildCursorSeedJobs(manifest.NormalAuthors, manifest.BigVAuthors, manifest.CursorDeep.PostAuthors); err != nil {
		return manifest, err
	} else if len(remaining) != 0 || len(manifest.Posts) != 1500 {
		return manifest, fmt.Errorf("cursor-deep publish is incomplete: posts=%d remaining=%d", len(manifest.Posts), len(remaining))
	}

	if _, err := waitForKafkaZero(ctx, cfg.Monitor, 5*time.Minute); err != nil {
		_ = saveManifest(manifestPath, manifest)
		return manifest, fmt.Errorf("wait for cursor-deep feed materialization: %w", err)
	}
	redisClient := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	defer redisClient.Close()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return manifest, fmt.Errorf("connect cursor-deep Redis %s: %w", cfg.RedisAddr, err)
	}
	normalPosts, bigVPosts, err := cursorPostsByTier(manifest)
	if err != nil {
		return manifest, err
	}
	duplicatePostID, err := normalizeCursorRedisScores(
		ctx, redisClient, manifest.Readers, normalPosts, bigVPosts, cursorDeepTieMembers,
	)
	if err != nil {
		return manifest, err
	}
	evidence, err := captureCursorRedisEvidence(
		ctx,
		redisClient,
		manifest.Readers,
		manifest.BigVAuthors,
		manifest.CursorDeep.PostAuthors,
		1000,
		100,
	)
	if err != nil {
		return manifest, err
	}
	if evidence.CrossSourceDuplicatePostID != duplicatePostID || evidence.SameSecondTieMembers < 100 {
		return manifest, fmt.Errorf(
			"cursor-deep Redis evidence mismatch: duplicate=%d/%d tie=%d",
			evidence.CrossSourceDuplicatePostID, duplicatePostID, evidence.SameSecondTieMembers,
		)
	}
	mysqlFingerprint, err := fingerprintCursorDeepMySQL(ctx, cfg.Monitor.MySQLDSN, manifest.CursorDeep.PostAuthors)
	if err != nil {
		return manifest, err
	}

	manifest.CursorDeep.Oracle = evidence.Oracle
	manifest.CursorDeep.OracleHash = evidence.OracleHash
	manifest.CursorDeep.RedisFingerprint = evidence.RedisFingerprint
	manifest.CursorDeep.MySQLFingerprint = mysqlFingerprint
	manifest.CursorDeep.SameSecondTieMembers = evidence.SameSecondTieMembers
	manifest.CursorDeep.CrossSourceDuplicatePostID = evidence.CrossSourceDuplicatePostID
	manifest.CursorDeep.Ready = true
	if err := validateCursorDeepManifest(manifest); err != nil {
		manifest.CursorDeep.Ready = false
		return manifest, err
	}
	if err := saveManifest(manifestPath, manifest); err != nil {
		return manifest, err
	}
	return manifest, nil
}

func cursorPostsByTier(manifest datasetManifest) (map[int64][]int64, map[int64][]int64, error) {
	normalSet := make(map[int64]struct{}, len(manifest.NormalAuthors))
	bigVSet := make(map[int64]struct{}, len(manifest.BigVAuthors))
	normal := make(map[int64][]int64, len(manifest.NormalAuthors))
	bigV := make(map[int64][]int64, len(manifest.BigVAuthors))
	for _, authorID := range manifest.NormalAuthors {
		normalSet[authorID] = struct{}{}
		normal[authorID] = nil
	}
	for _, authorID := range manifest.BigVAuthors {
		bigVSet[authorID] = struct{}{}
		bigV[authorID] = nil
	}
	for postID, authorID := range manifest.CursorDeep.PostAuthors {
		if _, ok := normalSet[authorID]; ok {
			normal[authorID] = append(normal[authorID], postID)
			continue
		}
		if _, ok := bigVSet[authorID]; ok {
			bigV[authorID] = append(bigV[authorID], postID)
			continue
		}
		return nil, nil, fmt.Errorf("cursor-deep post %d has unknown author %d", postID, authorID)
	}
	for authorID, postIDs := range normal {
		if len(postIDs) != 50 {
			return nil, nil, fmt.Errorf("cursor-deep normal author %d posts=%d, want 50", authorID, len(postIDs))
		}
	}
	for authorID, postIDs := range bigV {
		if len(postIDs) != 100 {
			return nil, nil, fmt.Errorf("cursor-deep big-v author %d posts=%d, want 100", authorID, len(postIDs))
		}
	}
	return normal, bigV, nil
}
