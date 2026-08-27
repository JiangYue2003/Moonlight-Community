package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const cursorDeepDatasetVersion = 1

type cursorOraclePost struct {
	PostID    int64 `json:"post_id"`
	CreatorID int64 `json:"creator_id"`
	SortTime  int64 `json:"sort_time"`
}

type cursorDeepDataset struct {
	Ready                         bool               `json:"ready"`
	Version                       int                `json:"version"`
	NormalPostsPerAuthor          int                `json:"normal_posts_per_author"`
	BigVPostsPerAuthor            int                `json:"bigv_posts_per_author"`
	InboxCandidatesPerReader      int                `json:"inbox_candidates_per_reader"`
	BigVOutboxCandidatesPerAuthor int                `json:"bigv_outbox_candidates_per_author"`
	Oracle                        []cursorOraclePost `json:"oracle"`
	OracleHash                    string             `json:"oracle_hash"`
	RedisFingerprint              string             `json:"redis_fingerprint"`
	MySQLFingerprint              string             `json:"mysql_fingerprint"`
	SameSecondTieMembers          int                `json:"same_second_tie_members,omitempty"`
	CrossSourceDuplicatePostID    int64              `json:"cross_source_duplicate_post_id,omitempty"`
	PostAuthors                   map[int64]int64    `json:"post_authors"`
}

type cursorSeedJob struct {
	RequestNumber int64
	AuthorID      int64
}

type cursorPublishedPost struct {
	RequestNumber int64
	PostID        int64
	AuthorID      int64
}

type cursorMySQLRow struct {
	PostID      int64
	CreatorID   int64
	Status      string
	Visible     string
	PublishUnix int64
}

type cursorRedisEvidence struct {
	Oracle                     []cursorOraclePost
	OracleHash                 string
	RedisFingerprint           string
	SameSecondTieMembers       int
	CrossSourceDuplicatePostID int64
}

type cursorRedisKeySnapshot struct {
	key   string
	posts []cursorOraclePost
}

func buildCursorSeedJobs(normalAuthors, bigVAuthors []int64, existingPostAuthors map[int64]int64) ([]cursorSeedJob, error) {
	targets := make(map[int64]int, len(normalAuthors)+len(bigVAuthors))
	for _, authorID := range normalAuthors {
		if authorID <= 0 {
			return nil, fmt.Errorf("normal author ids must be positive")
		}
		if _, exists := targets[authorID]; exists {
			return nil, fmt.Errorf("duplicate cursor seed author %d", authorID)
		}
		targets[authorID] = 50
	}
	for _, authorID := range bigVAuthors {
		if authorID <= 0 {
			return nil, fmt.Errorf("big-v author ids must be positive")
		}
		if _, exists := targets[authorID]; exists {
			return nil, fmt.Errorf("duplicate cursor seed author %d", authorID)
		}
		targets[authorID] = 100
	}
	if len(normalAuthors) != 20 || len(bigVAuthors) != 5 {
		return nil, fmt.Errorf("cursor-deep seed requires 20 normal and 5 big-v authors")
	}

	counts := make(map[int64]int, len(targets))
	for postID, authorID := range existingPostAuthors {
		if postID <= 0 || authorID <= 0 {
			return nil, fmt.Errorf("existing cursor seed post and author ids must be positive")
		}
		if _, ok := targets[authorID]; !ok {
			return nil, fmt.Errorf("existing cursor seed post %d has unknown author %d", postID, authorID)
		}
		counts[authorID]++
		if counts[authorID] > targets[authorID] {
			return nil, fmt.Errorf("cursor seed author %d already exceeds target %d", authorID, targets[authorID])
		}
	}

	authors := append(append([]int64(nil), normalAuthors...), bigVAuthors...)
	jobs := make([]cursorSeedJob, 0, 1500-len(existingPostAuthors))
	nextRequestNumber := int64(len(existingPostAuthors) + 1)
	for _, authorID := range authors {
		for count := counts[authorID]; count < targets[authorID]; count++ {
			jobs = append(jobs, cursorSeedJob{RequestNumber: nextRequestNumber, AuthorID: authorID})
			nextRequestNumber++
		}
	}
	return jobs, nil
}

func publishCursorSeedJobs(
	ctx context.Context,
	client publisherClient,
	jobs []cursorSeedJob,
	concurrency int,
	requestTimeout time.Duration,
	runID string,
) ([]cursorPublishedPost, error) {
	if len(jobs) == 0 {
		return nil, nil
	}
	if client == nil || concurrency <= 0 || requestTimeout <= 0 || strings.TrimSpace(runID) == "" {
		return nil, fmt.Errorf("cursor seed publisher, positive concurrency/timeout, and run id are required")
	}
	if concurrency > len(jobs) {
		concurrency = len(jobs)
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobCh := make(chan cursorSeedJob)
	published := make([]cursorPublishedPost, 0, len(jobs))
	var publishedMu sync.Mutex
	var firstErr error
	var errMu sync.Mutex
	draft := newStageRecorder("cursor_seed_draft")
	metadata := newStageRecorder("cursor_seed_metadata")
	confirm := newStageRecorder("cursor_seed_confirm")
	publish := newStageRecorder("cursor_seed_publish")

	var workers sync.WaitGroup
	workers.Add(concurrency)
	for workerID := 0; workerID < concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for job := range jobCh {
				requestCtx, requestCancel := context.WithTimeout(runCtx, requestTimeout)
				postID, err := publishOne(requestCtx, client, job.AuthorID, job.RequestNumber, runID, draft, metadata, confirm, publish)
				requestCancel()
				if err != nil {
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("cursor seed request %d author %d: %w", job.RequestNumber, job.AuthorID, err)
						cancel()
					}
					errMu.Unlock()
					continue
				}
				publishedMu.Lock()
				published = append(published, cursorPublishedPost{
					RequestNumber: job.RequestNumber,
					PostID:        postID,
					AuthorID:      job.AuthorID,
				})
				publishedMu.Unlock()
			}
		}()
	}

sendJobs:
	for _, job := range jobs {
		select {
		case <-runCtx.Done():
			break sendJobs
		case jobCh <- job:
		}
	}
	close(jobCh)
	workers.Wait()
	sort.Slice(published, func(i, j int) bool { return published[i].RequestNumber < published[j].RequestNumber })
	return published, firstErr
}

func fingerprintCursorDeepMySQL(
	ctx context.Context,
	dsn string,
	expectedPostAuthors map[int64]int64,
) (string, error) {
	if strings.TrimSpace(dsn) == "" || len(expectedPostAuthors) == 0 {
		return "", fmt.Errorf("cursor-deep MySQL DSN and expected posts are required")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return "", fmt.Errorf("open cursor-deep MySQL: %w", err)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		return "", fmt.Errorf("connect cursor-deep MySQL: %w", err)
	}

	postIDs := make([]int64, 0, len(expectedPostAuthors))
	for postID := range expectedPostAuthors {
		postIDs = append(postIDs, postID)
	}
	sort.Slice(postIDs, func(i, j int) bool { return postIDs[i] < postIDs[j] })
	rows := make([]cursorMySQLRow, 0, len(postIDs))
	const chunkSize = 500
	for start := 0; start < len(postIDs); start += chunkSize {
		end := start + chunkSize
		if end > len(postIDs) {
			end = len(postIDs)
		}
		placeholders := strings.TrimSuffix(strings.Repeat("?,", end-start), ",")
		args := make([]any, 0, end-start)
		for _, postID := range postIDs[start:end] {
			args = append(args, postID)
		}
		query := "SELECT id, creator_id, status, visible, UNIX_TIMESTAMP(publish_time) FROM know_posts WHERE id IN (" + placeholders + ")"
		result, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			return "", fmt.Errorf("query cursor-deep MySQL posts: %w", err)
		}
		for result.Next() {
			var row cursorMySQLRow
			var publishUnix sql.NullInt64
			if err := result.Scan(&row.PostID, &row.CreatorID, &row.Status, &row.Visible, &publishUnix); err != nil {
				_ = result.Close()
				return "", fmt.Errorf("scan cursor-deep MySQL post: %w", err)
			}
			if publishUnix.Valid {
				row.PublishUnix = publishUnix.Int64
			}
			rows = append(rows, row)
		}
		if err := result.Err(); err != nil {
			_ = result.Close()
			return "", fmt.Errorf("iterate cursor-deep MySQL posts: %w", err)
		}
		_ = result.Close()
	}
	return hashCursorMySQLRows(rows, expectedPostAuthors)
}

func hashCursorMySQLRows(rows []cursorMySQLRow, expectedPostAuthors map[int64]int64) (string, error) {
	if len(rows) != len(expectedPostAuthors) {
		return "", fmt.Errorf("cursor-deep MySQL rows=%d, want %d", len(rows), len(expectedPostAuthors))
	}
	ordered := append([]cursorMySQLRow(nil), rows...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].PostID < ordered[j].PostID })
	seen := make(map[int64]struct{}, len(ordered))
	digest := sha256.New()
	var encoded [24]byte
	for _, row := range ordered {
		expectedAuthor, ok := expectedPostAuthors[row.PostID]
		if !ok {
			return "", fmt.Errorf("cursor-deep MySQL contains unexpected post %d", row.PostID)
		}
		if _, duplicate := seen[row.PostID]; duplicate {
			return "", fmt.Errorf("cursor-deep MySQL contains duplicate post %d", row.PostID)
		}
		seen[row.PostID] = struct{}{}
		if row.CreatorID != expectedAuthor {
			return "", fmt.Errorf("cursor-deep MySQL post %d creator=%d, want %d", row.PostID, row.CreatorID, expectedAuthor)
		}
		if row.Status != "published" || row.Visible != "public" || row.PublishUnix <= 0 {
			return "", fmt.Errorf("cursor-deep MySQL post %d is not published/public with publish_time", row.PostID)
		}
		binary.LittleEndian.PutUint64(encoded[0:8], uint64(row.PostID))
		binary.LittleEndian.PutUint64(encoded[8:16], uint64(row.CreatorID))
		binary.LittleEndian.PutUint64(encoded[16:24], uint64(row.PublishUnix))
		_, _ = digest.Write(encoded[:])
		_, _ = digest.Write([]byte(row.Status))
		_, _ = digest.Write([]byte{0})
		_, _ = digest.Write([]byte(row.Visible))
		_, _ = digest.Write([]byte{0})
	}
	return hex.EncodeToString(digest.Sum(nil)), nil
}

func feedInboxKey(userID int64) string {
	return fmt.Sprintf("feed:inbox:%d", userID)
}

func feedBigVKey(creatorID int64) string {
	return fmt.Sprintf("feed:bigv:%d", creatorID)
}

func captureCursorRedisEvidence(
	ctx context.Context,
	client redis.Cmdable,
	readers []int64,
	bigVAuthors []int64,
	postAuthors map[int64]int64,
	expectedInbox int,
	expectedOutbox int,
) (cursorRedisEvidence, error) {
	if client == nil || len(readers) == 0 || len(bigVAuthors) == 0 {
		return cursorRedisEvidence{}, fmt.Errorf("Redis client, readers, and big-v authors are required")
	}
	outboxPosts := make([]cursorOraclePost, 0, len(bigVAuthors)*expectedOutbox)
	snapshots := make([]cursorRedisKeySnapshot, 0, len(readers)+len(bigVAuthors))
	for _, creatorID := range bigVAuthors {
		key := feedBigVKey(creatorID)
		posts, err := readCursorRedisKey(ctx, client, key, postAuthors)
		if err != nil {
			return cursorRedisEvidence{}, err
		}
		if len(posts) != expectedOutbox {
			return cursorRedisEvidence{}, fmt.Errorf("cursor-deep outbox %s members=%d, want %d", key, len(posts), expectedOutbox)
		}
		for _, post := range posts {
			if post.CreatorID != creatorID {
				return cursorRedisEvidence{}, fmt.Errorf("cursor-deep outbox %s contains creator %d", key, post.CreatorID)
			}
		}
		outboxPosts = append(outboxPosts, posts...)
		snapshots = append(snapshots, cursorRedisKeySnapshot{key: key, posts: posts})
	}

	var evidence cursorRedisEvidence
	for index, readerID := range readers {
		key := feedInboxKey(readerID)
		inboxPosts, err := readCursorRedisKey(ctx, client, key, postAuthors)
		if err != nil {
			return cursorRedisEvidence{}, err
		}
		if len(inboxPosts) != expectedInbox {
			return cursorRedisEvidence{}, fmt.Errorf("cursor-deep inbox %s members=%d, want %d", key, len(inboxPosts), expectedInbox)
		}
		allPosts := append(append([]cursorOraclePost(nil), inboxPosts...), outboxPosts...)
		oracle, oracleHash, err := buildCursorOracle(allPosts)
		if err != nil {
			return cursorRedisEvidence{}, fmt.Errorf("cursor-deep reader %d oracle: %w", readerID, err)
		}
		if index == 0 {
			evidence.Oracle = oracle
			evidence.OracleHash = oracleHash
			evidence.CrossSourceDuplicatePostID = firstCursorDuplicatePostID(allPosts)
			evidence.SameSecondTieMembers = maximumCursorTieMembers(oracle)
		} else if oracleHash != evidence.OracleHash {
			return cursorRedisEvidence{}, fmt.Errorf("cursor-deep reader %d oracle differs from reader %d", readerID, readers[0])
		}
		snapshots = append(snapshots, cursorRedisKeySnapshot{key: key, posts: inboxPosts})
	}
	evidence.RedisFingerprint = hashCursorRedisSnapshots(snapshots)
	return evidence, nil
}

func normalizeCursorRedisScores(
	ctx context.Context,
	client redis.UniversalClient,
	readers []int64,
	normalPostsByAuthor map[int64][]int64,
	bigVPostsByAuthor map[int64][]int64,
	tieMembers int,
) (int64, error) {
	if client == nil || len(readers) == 0 || len(normalPostsByAuthor) == 0 || len(bigVPostsByAuthor) == 0 {
		return 0, fmt.Errorf("cursor-deep normalization requires Redis, readers, normal posts, and big-v posts")
	}
	normalPostIDs := flattenCursorPostIDs(normalPostsByAuthor)
	bigVPostIDs := flattenCursorPostIDs(bigVPostsByAuthor)
	if len(normalPostIDs) == 0 || len(bigVPostIDs) == 0 {
		return 0, fmt.Errorf("cursor-deep normalization post sets are empty")
	}
	duplicatePostID := bigVPostIDs[0]
	removedNormalPostID := normalPostIDs[len(normalPostIDs)-1]
	normalizedInboxPostIDs := make([]int64, 0, len(normalPostIDs))
	for _, postID := range normalPostIDs {
		if postID != removedNormalPostID {
			normalizedInboxPostIDs = append(normalizedInboxPostIDs, postID)
		}
	}
	normalizedInboxPostIDs = append(normalizedInboxPostIDs, duplicatePostID)
	for _, readerID := range readers {
		if err := verifyCursorRedisMembersAny(ctx, client, feedInboxKey(readerID), normalPostIDs, normalizedInboxPostIDs); err != nil {
			return 0, err
		}
	}
	for creatorID, postIDs := range bigVPostsByAuthor {
		if err := verifyCursorRedisMembers(ctx, client, feedBigVKey(creatorID), postIDs); err != nil {
			return 0, err
		}
	}

	allPostIDs := append(append([]int64(nil), normalPostIDs...), bigVPostIDs...)
	sort.Slice(allPostIDs, func(i, j int) bool { return allPostIDs[i] > allPostIDs[j] })
	if tieMembers < 1 {
		tieMembers = 1
	}
	if tieMembers > len(allPostIDs) {
		tieMembers = len(allPostIDs)
	}
	const baseSortTime int64 = 1_700_000_000
	scores := make(map[int64]int64, len(allPostIDs))
	for index, postID := range allPostIDs {
		score := baseSortTime
		if index >= tieMembers {
			score -= int64(index - tieMembers + 1)
		}
		scores[postID] = score
	}

	pipe := client.Pipeline()
	for _, readerID := range readers {
		key := feedInboxKey(readerID)
		for _, postID := range normalPostIDs {
			pipe.ZAdd(ctx, key, redis.Z{Score: float64(scores[postID]), Member: postID})
		}
	}
	for creatorID, postIDs := range bigVPostsByAuthor {
		key := feedBigVKey(creatorID)
		for _, postID := range postIDs {
			pipe.ZAdd(ctx, key, redis.Z{Score: float64(scores[postID]), Member: postID})
		}
	}
	for _, readerID := range readers {
		key := feedInboxKey(readerID)
		pipe.ZRem(ctx, key, removedNormalPostID)
		pipe.ZAdd(ctx, key, redis.Z{Score: float64(scores[duplicatePostID]), Member: duplicatePostID})
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return 0, fmt.Errorf("normalize cursor-deep Redis scores: %w", err)
	}
	return duplicatePostID, nil
}

func verifyCursorRedisMembersAny(ctx context.Context, client redis.Cmdable, key string, alternatives ...[]int64) error {
	members, err := client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return fmt.Errorf("verify cursor-deep Redis key %s: %w", key, err)
	}
	actual := make(map[int64]struct{}, len(members))
	for _, member := range members {
		postID, err := strconv.ParseInt(member, 10, 64)
		if err != nil {
			return fmt.Errorf("cursor-deep Redis key %s has invalid member %q", key, member)
		}
		actual[postID] = struct{}{}
	}
	for _, expected := range alternatives {
		want := make(map[int64]struct{}, len(expected))
		for _, postID := range expected {
			want[postID] = struct{}{}
		}
		if len(actual) != len(want) {
			continue
		}
		matches := true
		for postID := range actual {
			if _, ok := want[postID]; !ok {
				matches = false
				break
			}
		}
		if matches {
			return nil
		}
	}
	return fmt.Errorf("cursor-deep Redis key %s members do not match original or normalized dataset", key)
}

func flattenCursorPostIDs(postsByAuthor map[int64][]int64) []int64 {
	result := make([]int64, 0)
	for _, postIDs := range postsByAuthor {
		result = append(result, postIDs...)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })
	return result
}

func verifyCursorRedisMembers(
	ctx context.Context,
	client redis.Cmdable,
	key string,
	expected []int64,
) error {
	members, err := client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return fmt.Errorf("verify cursor-deep Redis key %s: %w", key, err)
	}
	want := make(map[int64]struct{}, len(expected))
	for _, postID := range expected {
		want[postID] = struct{}{}
	}
	if len(members) != len(want) {
		return fmt.Errorf("cursor-deep Redis key %s members=%d, want %d", key, len(members), len(want))
	}
	for _, member := range members {
		postID, err := strconv.ParseInt(member, 10, 64)
		if err != nil {
			return fmt.Errorf("cursor-deep Redis key %s has invalid member %q", key, member)
		}
		if _, ok := want[postID]; !ok {
			return fmt.Errorf("cursor-deep Redis key %s has unexpected member %d", key, postID)
		}
	}
	return nil
}

func readCursorRedisKey(
	ctx context.Context,
	client redis.Cmdable,
	key string,
	postAuthors map[int64]int64,
) ([]cursorOraclePost, error) {
	values, err := client.ZRevRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("read cursor-deep Redis key %s: %w", key, err)
	}
	posts := make([]cursorOraclePost, 0, len(values))
	for _, value := range values {
		postID, err := strconv.ParseInt(fmt.Sprint(value.Member), 10, 64)
		if err != nil || postID <= 0 {
			return nil, fmt.Errorf("cursor-deep Redis key %s has invalid member %v", key, value.Member)
		}
		creatorID := postAuthors[postID]
		if creatorID <= 0 {
			return nil, fmt.Errorf("cursor-deep Redis key %s has unknown post %d", key, postID)
		}
		if value.Score <= 0 || math.Trunc(value.Score) != value.Score {
			return nil, fmt.Errorf("cursor-deep Redis key %s post %d has invalid score %v", key, postID, value.Score)
		}
		posts = append(posts, cursorOraclePost{PostID: postID, CreatorID: creatorID, SortTime: int64(value.Score)})
	}
	return posts, nil
}

func firstCursorDuplicatePostID(posts []cursorOraclePost) int64 {
	seen := make(map[int64]struct{}, len(posts))
	for _, post := range posts {
		if _, ok := seen[post.PostID]; ok {
			return post.PostID
		}
		seen[post.PostID] = struct{}{}
	}
	return 0
}

func maximumCursorTieMembers(posts []cursorOraclePost) int {
	maximum := 0
	for start := 0; start < len(posts); {
		end := start + 1
		for end < len(posts) && posts[end].SortTime == posts[start].SortTime {
			end++
		}
		if end-start > maximum {
			maximum = end - start
		}
		start = end
	}
	return maximum
}

func hashCursorRedisSnapshots(snapshots []cursorRedisKeySnapshot) string {
	sort.Slice(snapshots, func(i, j int) bool { return snapshots[i].key < snapshots[j].key })
	digest := sha256.New()
	var encoded [24]byte
	for _, snapshot := range snapshots {
		_, _ = digest.Write([]byte(snapshot.key))
		_, _ = digest.Write([]byte{0})
		for _, post := range snapshot.posts {
			binary.LittleEndian.PutUint64(encoded[0:8], uint64(post.SortTime))
			binary.LittleEndian.PutUint64(encoded[8:16], uint64(post.PostID))
			binary.LittleEndian.PutUint64(encoded[16:24], uint64(post.CreatorID))
			_, _ = digest.Write(encoded[:])
		}
	}
	return hex.EncodeToString(digest.Sum(nil))
}

func buildCursorOracle(posts []cursorOraclePost) ([]cursorOraclePost, string, error) {
	unique := make(map[int64]cursorOraclePost, len(posts))
	for _, post := range posts {
		if post.PostID <= 0 || post.CreatorID <= 0 || post.SortTime <= 0 {
			return nil, "", fmt.Errorf("cursor oracle contains non-positive post, creator, or sort time")
		}
		if existing, ok := unique[post.PostID]; ok {
			if existing != post {
				return nil, "", fmt.Errorf("cursor oracle post %d has conflicting source metadata", post.PostID)
			}
			continue
		}
		unique[post.PostID] = post
	}
	oracle := make([]cursorOraclePost, 0, len(unique))
	for _, post := range unique {
		oracle = append(oracle, post)
	}
	sort.Slice(oracle, func(i, j int) bool {
		if oracle[i].SortTime != oracle[j].SortTime {
			return oracle[i].SortTime > oracle[j].SortTime
		}
		return oracle[i].PostID > oracle[j].PostID
	})

	digest := sha256.New()
	var encoded [24]byte
	for _, post := range oracle {
		binary.LittleEndian.PutUint64(encoded[0:8], uint64(post.SortTime))
		binary.LittleEndian.PutUint64(encoded[8:16], uint64(post.PostID))
		binary.LittleEndian.PutUint64(encoded[16:24], uint64(post.CreatorID))
		_, _ = digest.Write(encoded[:])
	}
	return oracle, hex.EncodeToString(digest.Sum(nil)), nil
}

func cursorOraclePostIDs(posts []cursorOraclePost) []int64 {
	ids := make([]int64, len(posts))
	for i, post := range posts {
		ids[i] = post.PostID
	}
	return ids
}

func validateCursorDeepManifest(manifest datasetManifest) error {
	deep := manifest.CursorDeep
	if deep == nil {
		return fmt.Errorf("manifest has no cursor-deep contract")
	}
	if deep.Version != cursorDeepDatasetVersion {
		return fmt.Errorf("cursor-deep version=%d, want %d", deep.Version, cursorDeepDatasetVersion)
	}
	if !deep.Ready {
		return fmt.Errorf("cursor-deep dataset is not ready; resume seed-cursor-deep")
	}
	if len(manifest.NormalAuthors) != 20 || len(manifest.BigVAuthors) != 5 || len(manifest.Readers) != 20 {
		return fmt.Errorf("cursor-deep authors/readers=%d/%d/%d, want 20/5/20",
			len(manifest.NormalAuthors), len(manifest.BigVAuthors), len(manifest.Readers))
	}
	if deep.NormalPostsPerAuthor != 50 || deep.BigVPostsPerAuthor != 100 || len(manifest.Posts) != 1500 {
		return fmt.Errorf("cursor-deep post shape is not 20x50 + 5x100")
	}
	if err := validateCursorPostOwnership(manifest, deep.PostAuthors); err != nil {
		return err
	}
	if deep.InboxCandidatesPerReader != 1000 || deep.BigVOutboxCandidatesPerAuthor != 100 {
		return fmt.Errorf("cursor-deep Redis capacity is not inbox=1000/outbox=100")
	}
	if len(deep.Oracle) != 1499 {
		return fmt.Errorf("cursor-deep oracle posts=%d, want 1499 after cross-source deduplication", len(deep.Oracle))
	}
	if deep.OracleHash == "" {
		return fmt.Errorf("cursor-deep oracle hash is empty")
	}
	if deep.RedisFingerprint == "" || deep.MySQLFingerprint == "" {
		return fmt.Errorf("cursor-deep Redis/MySQL fingerprints are required")
	}
	if deep.SameSecondTieMembers < 100 {
		return fmt.Errorf("cursor-deep same-second tie members=%d, want at least 100", deep.SameSecondTieMembers)
	}
	if actual := maximumCursorTieMembers(deep.Oracle); actual != deep.SameSecondTieMembers {
		return fmt.Errorf("cursor-deep same-second tie evidence=%d, oracle=%d", deep.SameSecondTieMembers, actual)
	}
	if deep.CrossSourceDuplicatePostID <= 0 {
		return fmt.Errorf("cursor-deep cross-source duplicate is required")
	}
	bigVSet := make(map[int64]struct{}, len(manifest.BigVAuthors))
	for _, authorID := range manifest.BigVAuthors {
		bigVSet[authorID] = struct{}{}
	}
	if _, ok := bigVSet[deep.PostAuthors[deep.CrossSourceDuplicatePostID]]; !ok {
		return fmt.Errorf("cursor-deep cross-source duplicate %d is not a big-v post", deep.CrossSourceDuplicatePostID)
	}
	for _, post := range deep.Oracle {
		if deep.PostAuthors[post.PostID] != post.CreatorID {
			return fmt.Errorf("cursor-deep oracle post %d ownership mismatch", post.PostID)
		}
	}
	_, hash, err := buildCursorOracle(deep.Oracle)
	if err != nil {
		return err
	}
	if hash != deep.OracleHash {
		return fmt.Errorf("cursor-deep oracle hash mismatch")
	}
	return nil
}

func validateCursorPostOwnership(manifest datasetManifest, postAuthors map[int64]int64) error {
	if len(postAuthors) != len(manifest.Posts) {
		return fmt.Errorf("cursor-deep post ownership entries=%d, want %d", len(postAuthors), len(manifest.Posts))
	}
	targets := make(map[int64]int, len(manifest.NormalAuthors)+len(manifest.BigVAuthors))
	for _, authorID := range manifest.NormalAuthors {
		targets[authorID] = 50
	}
	for _, authorID := range manifest.BigVAuthors {
		if _, duplicate := targets[authorID]; duplicate {
			return fmt.Errorf("cursor-deep author %d is present in both tiers", authorID)
		}
		targets[authorID] = 100
	}
	seen := make(map[int64]struct{}, len(manifest.Posts))
	counts := make(map[int64]int, len(targets))
	for _, postID := range manifest.Posts {
		if postID <= 0 {
			return fmt.Errorf("cursor-deep manifest contains non-positive post id")
		}
		if _, duplicate := seen[postID]; duplicate {
			return fmt.Errorf("cursor-deep manifest contains duplicate post %d", postID)
		}
		seen[postID] = struct{}{}
		authorID, ok := postAuthors[postID]
		if !ok {
			return fmt.Errorf("cursor-deep manifest post %d has no owner", postID)
		}
		if _, ok := targets[authorID]; !ok {
			return fmt.Errorf("cursor-deep manifest post %d has unknown author %d", postID, authorID)
		}
		counts[authorID]++
	}
	for authorID, want := range targets {
		if counts[authorID] != want {
			return fmt.Errorf("cursor-deep author %d posts=%d, want %d", authorID, counts[authorID], want)
		}
	}
	return nil
}
