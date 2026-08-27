package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	authpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
)

type recordingTokenIssuer struct{}

func (recordingTokenIssuer) IssuePair(uid int64, _ string) (*jwtx.TokenPair, error) {
	return &jwtx.TokenPair{AccessToken: fmt.Sprintf("signed-%d", uid)}, nil
}

func TestIssueReaderTokensAvoidsLoginAndPreservesManifestOrder(t *testing.T) {
	users := []benchmarkUser{{ID: 9}, {ID: 7}}

	readers, err := issueReaderTokens(context.Background(), recordingTokenIssuer{}, users, 2)

	require.NoError(t, err)
	require.Equal(t, []readerIdentity{{UserID: 9, AccessToken: "signed-9"}, {UserID: 7, AccessToken: "signed-7"}}, readers)
}

type failingUserCreator struct {
	concurrentUserCreator
	failAt int64
}

func (c *failingUserCreator) Create(
	ctx context.Context,
	req *userpb.CreateReq,
	options ...grpc.CallOption,
) (*userpb.CreateResp, error) {
	c.mu.Lock()
	next := c.nextID + 1
	c.mu.Unlock()
	if next == c.failAt {
		return nil, errors.New("injected create failure")
	}
	return c.concurrentUserCreator.Create(ctx, req, options...)
}

type concurrentUserCreator struct {
	mu       sync.Mutex
	nextID   int64
	requests []*userpb.CreateReq
}

func (c *concurrentUserCreator) Create(
	_ context.Context,
	req *userpb.CreateReq,
	_ ...grpc.CallOption,
) (*userpb.CreateResp, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nextID++
	c.requests = append(c.requests, req)
	return &userpb.CreateResp{Id: c.nextID}, nil
}

func TestCreateBenchmarkUsersIncludesRunIDAndPasswordHash(t *testing.T) {
	client := &concurrentUserCreator{nextID: 100}

	users, err := createBenchmarkUsers(context.Background(), client, 4, 2, "run-42", "bcrypt-hash")

	require.NoError(t, err)
	require.Len(t, users, 4)
	require.Equal(t, int64(101), users[0].ID)
	require.Contains(t, users[0].Email, "run-42")
	require.Len(t, client.requests, 4)
	for _, req := range client.requests {
		require.Equal(t, "bcrypt-hash", req.PasswordHash)
	}
}

func TestCreateBenchmarkUsersReturnsCreatedUsersWhenSetupFails(t *testing.T) {
	client := &failingUserCreator{concurrentUserCreator: concurrentUserCreator{nextID: 100}, failAt: 103}

	users, err := createBenchmarkUsers(context.Background(), client, 5, 1, "run-42", "bcrypt-hash")

	require.ErrorContains(t, err, "injected create failure")
	require.Len(t, users, 2)
	require.Equal(t, int64(101), users[0].ID)
	require.Equal(t, int64(102), users[1].ID)
}

type recordingAuthLogin struct {
	requests []*authpb.LoginReq
}

func (c *recordingAuthLogin) Login(
	_ context.Context,
	req *authpb.LoginReq,
	_ ...grpc.CallOption,
) (*authpb.AuthResp, error) {
	c.requests = append(c.requests, req)
	return &authpb.AuthResp{Token: &authpb.TokenPair{AccessToken: "token-for-" + req.Identifier}}, nil
}

func TestLoginReadersReturnsRealAccessTokens(t *testing.T) {
	client := &recordingAuthLogin{}
	users := []benchmarkUser{{ID: 7, Email: "reader@example.invalid"}}

	readers, err := loginReaders(context.Background(), client, users, "Password!2026", 1)

	require.NoError(t, err)
	require.Equal(t, []readerIdentity{{
		UserID:      7,
		AccessToken: "token-for-reader@example.invalid",
	}}, readers)
	require.Equal(t, "PASSWORD", client.requests[0].Channel)
}

func TestParseConcurrencyListPreservesRequestedCompressedLevels(t *testing.T) {
	levels, err := parseConcurrencyList("16,32,64")

	require.NoError(t, err)
	require.Equal(t, []int{16, 32, 64}, levels)
	require.Error(t, func() error {
		_, err := parseConcurrencyList("16,0,64")
		return err
	}())
}

func TestDatasetManifestRoundTrip(t *testing.T) {
	path := t.TempDir() + "/manifest.json"
	want := datasetManifest{
		RunID:    "run-42",
		Strategy: "hybrid",
		Users:    []benchmarkUser{{ID: 1, Email: "u@example.invalid", Role: "reader"}},
		Readers:  []int64{1},
		Posts:    []int64{9},
	}

	require.NoError(t, saveManifest(path, want))
	got, err := loadManifest(path)

	require.NoError(t, err)
	require.Equal(t, want.RunID, got.RunID)
	require.Equal(t, want.Strategy, got.Strategy)
	require.Equal(t, want.Users, got.Users)
	require.Equal(t, want.Readers, got.Readers)
	require.Equal(t, want.Posts, got.Posts)
}

func TestResolveManifestReaderCardinalitySupportsPresetsAndRejectsMismatch(t *testing.T) {
	manifest := datasetManifest{ReaderCardinality: "distributed", Readers: make([]int64, 1200)}

	got, err := resolveManifestReaderCardinality(manifest, "distributed")
	require.NoError(t, err)
	require.Equal(t, "distributed", got)

	_, err = resolveManifestReaderCardinality(manifest, "hot")
	require.ErrorContains(t, err, "cardinality")
}

func TestResolveManifestReaderCardinalityInfersLegacyHotManifest(t *testing.T) {
	got, err := resolveManifestReaderCardinality(datasetManifest{Readers: make([]int64, 20)}, "hot")
	require.NoError(t, err)
	require.Equal(t, "hot", got)
}

func TestValidateManifestSeedRejectsLegacyAndMismatch(t *testing.T) {
	require.ErrorContains(t, validateManifestSeed(datasetManifest{RunID: "legacy"}, 42), "no recorded seed")
	require.ErrorContains(t, validateManifestSeed(datasetManifest{RunID: "new", Seed: 42}, 7), "does not match")
	require.NoError(t, validateManifestSeed(datasetManifest{RunID: "new", Seed: 42}, 42))
}

func TestValidateManifestStrategyRejectsLegacyAndMismatch(t *testing.T) {
	require.ErrorContains(t, validateManifestStrategy(datasetManifest{RunID: "legacy"}, "hybrid"), "no recorded Feed strategy")
	require.ErrorContains(t, validateManifestStrategy(datasetManifest{RunID: "new", Strategy: "push"}, "hybrid"), "does not match")
	require.NoError(t, validateManifestStrategy(datasetManifest{RunID: "new", Strategy: "hybrid"}, "hybrid"))
}
