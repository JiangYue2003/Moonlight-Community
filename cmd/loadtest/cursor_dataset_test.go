package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type failingCursorPublisher struct {
	recordingPublisher
	failAfter int
}

func (p *failingCursorPublisher) CreateDraft(
	ctx context.Context,
	in *knowpostpb.CreateDraftReq,
	options ...grpc.CallOption,
) (*knowpostpb.CreateDraftResp, error) {
	p.mu.Lock()
	shouldFail := int(p.nextID) >= p.failAfter
	p.mu.Unlock()
	if shouldFail {
		return nil, errors.New("injected cursor seed failure")
	}
	return p.recordingPublisher.CreateDraft(ctx, in, options...)
}

func TestBuildCursorOracleUsesNumericTotalOrderAndDeduplicatesSources(t *testing.T) {
	posts := []cursorOraclePost{
		{PostID: 9, CreatorID: 1, SortTime: 100},
		{PostID: 100, CreatorID: 2, SortTime: 100},
		{PostID: 10, CreatorID: 3, SortTime: 100},
		{PostID: 9, CreatorID: 1, SortTime: 100},
		{PostID: 200, CreatorID: 2, SortTime: 99},
	}

	oracle, hash, err := buildCursorOracle(posts)

	require.NoError(t, err)
	require.Equal(t, []int64{100, 10, 9, 200}, cursorOraclePostIDs(oracle))
	require.Len(t, hash, 64)

	_, changedHash, err := buildCursorOracle(append(posts, cursorOraclePost{PostID: 201, CreatorID: 2, SortTime: 98}))
	require.NoError(t, err)
	require.NotEqual(t, hash, changedHash)
}

func TestCaptureCursorRedisEvidenceVerifiesEveryReaderAndFindsCrossSourceDuplicate(t *testing.T) {
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()
	for _, readerID := range []int64{10, 11} {
		require.NoError(t, client.ZAdd(ctx, feedInboxKey(readerID),
			redis.Z{Score: 100, Member: 1}, redis.Z{Score: 99, Member: 2}, redis.Z{Score: 100, Member: 5},
		).Err())
	}
	require.NoError(t, client.ZAdd(ctx, feedBigVKey(20),
		redis.Z{Score: 100, Member: 5}, redis.Z{Score: 98, Member: 6},
	).Err())

	evidence, err := captureCursorRedisEvidence(ctx, client, []int64{10, 11}, []int64{20}, map[int64]int64{
		1: 3, 2: 3, 5: 20, 6: 20,
	}, 3, 2)

	require.NoError(t, err)
	require.Equal(t, []int64{5, 1, 2, 6}, cursorOraclePostIDs(evidence.Oracle))
	require.Equal(t, int64(5), evidence.CrossSourceDuplicatePostID)
	require.Equal(t, 2, evidence.SameSecondTieMembers)
	require.Len(t, evidence.OracleHash, 64)
	require.Len(t, evidence.RedisFingerprint, 64)
}

func TestNormalizeCursorRedisScoresBuildsTieGroupAndCrossSourceDuplicate(t *testing.T) {
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	ctx := context.Background()
	for _, readerID := range []int64{10, 11} {
		require.NoError(t, client.ZAdd(ctx, feedInboxKey(readerID),
			redis.Z{Score: 10, Member: 1}, redis.Z{Score: 11, Member: 2},
		).Err())
	}
	require.NoError(t, client.ZAdd(ctx, feedBigVKey(20),
		redis.Z{Score: 12, Member: 5}, redis.Z{Score: 13, Member: 6},
	).Err())

	duplicateID, err := normalizeCursorRedisScores(ctx, client, []int64{10, 11},
		map[int64][]int64{3: {1, 2}}, map[int64][]int64{20: {5, 6}}, 2)

	require.NoError(t, err)
	require.Equal(t, int64(6), duplicateID)
	evidence, err := captureCursorRedisEvidence(ctx, client, []int64{10, 11}, []int64{20}, map[int64]int64{
		1: 3, 2: 3, 5: 20, 6: 20,
	}, 2, 2)
	require.NoError(t, err)
	require.Equal(t, duplicateID, evidence.CrossSourceDuplicatePostID)
	require.GreaterOrEqual(t, evidence.SameSecondTieMembers, 2)

	secondDuplicateID, err := normalizeCursorRedisScores(ctx, client, []int64{10, 11},
		map[int64][]int64{3: {1, 2}}, map[int64][]int64{20: {5, 6}}, 2)
	require.NoError(t, err)
	require.Equal(t, duplicateID, secondDuplicateID)
}

func TestValidateCursorDeepManifestRequiresAuditableShapeAndOracle(t *testing.T) {
	manifest := validCursorDeepManifestForTest(t)

	require.NoError(t, validateCursorDeepManifest(manifest))
	manifest.CursorDeep.Ready = false
	require.ErrorContains(t, validateCursorDeepManifest(manifest), "not ready")
	manifest.CursorDeep.Ready = true
	manifest.CursorDeep.OracleHash = ""
	require.ErrorContains(t, validateCursorDeepManifest(manifest), "oracle hash")
}

func validCursorDeepManifestForTest(t *testing.T) datasetManifest {
	t.Helper()
	manifest := datasetManifest{
		RunID:         "feed-cursor-deep-v1",
		NormalAuthors: make([]int64, 20),
		BigVAuthors:   make([]int64, 5),
		Readers:       make([]int64, 20),
		Posts:         make([]int64, 1500),
	}
	for i := range manifest.NormalAuthors {
		manifest.NormalAuthors[i] = int64(i + 1)
	}
	for i := range manifest.BigVAuthors {
		manifest.BigVAuthors[i] = int64(i + 101)
	}
	for i := range manifest.Readers {
		manifest.Readers[i] = int64(i + 201)
	}
	postAuthors := make(map[int64]int64, 1500)
	oracleInput := make([]cursorOraclePost, 0, 1499)
	postIndex := 0
	for _, authorID := range manifest.NormalAuthors {
		for count := 0; count < 50; count++ {
			postID := int64(postIndex + 1)
			manifest.Posts[postIndex] = postID
			postAuthors[postID] = authorID
			if postID != 1000 {
				oracleInput = append(oracleInput, cursorOraclePost{PostID: postID, CreatorID: authorID, SortTime: cursorTestSortTime(postIndex)})
			}
			postIndex++
		}
	}
	for _, authorID := range manifest.BigVAuthors {
		for count := 0; count < 100; count++ {
			postID := int64(postIndex + 1)
			manifest.Posts[postIndex] = postID
			postAuthors[postID] = authorID
			oracleInput = append(oracleInput, cursorOraclePost{PostID: postID, CreatorID: authorID, SortTime: cursorTestSortTime(postIndex)})
			postIndex++
		}
	}
	oracle, oracleHash, err := buildCursorOracle(oracleInput)
	require.NoError(t, err)
	manifest.CursorDeep = &cursorDeepDataset{
		Ready: true, Version: 1, NormalPostsPerAuthor: 50, BigVPostsPerAuthor: 100,
		InboxCandidatesPerReader: 1000, BigVOutboxCandidatesPerAuthor: 100,
		Oracle: oracle, OracleHash: oracleHash, RedisFingerprint: "redis", MySQLFingerprint: "mysql",
		SameSecondTieMembers: 120, CrossSourceDuplicatePostID: 1001, PostAuthors: postAuthors,
	}
	return manifest
}

func cursorTestSortTime(postIndex int) int64 {
	if postIndex < 120 {
		return 2000
	}
	return 2000 - int64(postIndex-119)
}

func TestBuildCursorSeedJobsCreatesExactPerAuthorDepthAndResumesDeficits(t *testing.T) {
	normalAuthors := make([]int64, 20)
	bigVAuthors := make([]int64, 5)
	for i := range normalAuthors {
		normalAuthors[i] = int64(i + 1)
	}
	for i := range bigVAuthors {
		bigVAuthors[i] = int64(i + 101)
	}
	existing := map[int64]int64{
		9001: normalAuthors[0],
		9002: bigVAuthors[0],
	}

	jobs, err := buildCursorSeedJobs(normalAuthors, bigVAuthors, existing)

	require.NoError(t, err)
	require.Len(t, jobs, 1498)
	counts := map[int64]int{}
	for _, authorID := range existing {
		counts[authorID]++
	}
	for _, job := range jobs {
		require.Positive(t, job.RequestNumber)
		counts[job.AuthorID]++
	}
	for _, authorID := range normalAuthors {
		require.Equal(t, 50, counts[authorID], "normal author %d", authorID)
	}
	for _, authorID := range bigVAuthors {
		require.Equal(t, 100, counts[authorID], "big-v author %d", authorID)
	}
}

func TestPublishCursorSeedJobsReturnsSuccessfulSubsetForResume(t *testing.T) {
	client := &failingCursorPublisher{failAfter: 2}
	jobs := []cursorSeedJob{
		{RequestNumber: 1, AuthorID: 7},
		{RequestNumber: 2, AuthorID: 7},
		{RequestNumber: 3, AuthorID: 8},
	}

	published, err := publishCursorSeedJobs(
		context.Background(), client, jobs, 1, time.Second, "cursor-seed",
	)

	require.ErrorContains(t, err, "injected cursor seed failure")
	require.Len(t, published, 2)
	require.Equal(t, []int64{7, 7}, []int64{published[0].AuthorID, published[1].AuthorID})
	for _, post := range published {
		require.Positive(t, post.PostID)
	}
}

func TestHashCursorMySQLRowsIsOrderIndependentAndRejectsWrongOwnership(t *testing.T) {
	expected := map[int64]int64{11: 7, 12: 8}
	rows := []cursorMySQLRow{
		{PostID: 12, CreatorID: 8, Status: "published", Visible: "public", PublishUnix: 101},
		{PostID: 11, CreatorID: 7, Status: "published", Visible: "public", PublishUnix: 100},
	}

	first, err := hashCursorMySQLRows(rows, expected)
	require.NoError(t, err)
	second, err := hashCursorMySQLRows([]cursorMySQLRow{rows[1], rows[0]}, expected)
	require.NoError(t, err)
	require.Equal(t, first, second)
	require.Len(t, first, 64)

	rows[0].CreatorID = 999
	_, err = hashCursorMySQLRows(rows, expected)
	require.ErrorContains(t, err, "creator")
}

func TestMergeCursorPublishedPostsCheckpointsOnlyUniqueSuccessfulPosts(t *testing.T) {
	manifest := datasetManifest{
		Posts: []int64{10},
		CursorDeep: &cursorDeepDataset{
			PostAuthors: map[int64]int64{10: 7},
		},
	}

	err := mergeCursorPublishedPosts(&manifest, []cursorPublishedPost{
		{RequestNumber: 2, PostID: 12, AuthorID: 8},
		{RequestNumber: 1, PostID: 11, AuthorID: 7},
	})

	require.NoError(t, err)
	require.Equal(t, []int64{10, 11, 12}, manifest.Posts)
	require.Equal(t, int64(8), manifest.CursorDeep.PostAuthors[12])

	err = mergeCursorPublishedPosts(&manifest, []cursorPublishedPost{{PostID: 12, AuthorID: 8}})
	require.ErrorContains(t, err, "duplicate")
}
