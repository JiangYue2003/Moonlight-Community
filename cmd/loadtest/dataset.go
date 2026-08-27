package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	authpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
	"google.golang.org/grpc"
)

type tokenPairIssuer interface {
	IssuePair(uid int64, nickname string) (*jwtx.TokenPair, error)
}

func issueReaderTokens(ctx context.Context, issuer tokenPairIssuer, users []benchmarkUser, concurrency int) ([]readerIdentity, error) {
	if len(users) == 0 {
		return nil, nil
	}
	if issuer == nil || concurrency <= 0 {
		return nil, fmt.Errorf("token issuer and positive concurrency are required")
	}
	if concurrency > len(users) {
		concurrency = len(users)
	}
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobs := make(chan int)
	readers := make([]readerIdentity, len(users))
	var firstErr error
	var errMu sync.Mutex
	var workers sync.WaitGroup
	workers.Add(concurrency)
	for workerID := 0; workerID < concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for index := range jobs {
				select {
				case <-runCtx.Done():
					return
				default:
				}
				user := users[index]
				pair, err := issuer.IssuePair(user.ID, "")
				if err == nil && (pair == nil || pair.AccessToken == "") {
					err = fmt.Errorf("missing access token")
				}
				if err != nil {
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("issue token for reader %d: %w", user.ID, err)
						cancel()
					}
					errMu.Unlock()
					continue
				}
				readers[index] = readerIdentity{UserID: user.ID, AccessToken: pair.AccessToken}
			}
		}()
	}
sendJobs:
	for index := range users {
		select {
		case <-runCtx.Done():
			break sendJobs
		case jobs <- index:
		}
	}
	close(jobs)
	workers.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	return readers, nil
}

type benchmarkUser struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type datasetManifest struct {
	RunID             string             `json:"run_id"`
	Seed              int64              `json:"seed,omitempty"`
	Strategy          string             `json:"strategy,omitempty"`
	ReaderCardinality string             `json:"reader_cardinality,omitempty"`
	Users             []benchmarkUser    `json:"users"`
	NormalAuthors     []int64            `json:"normal_authors"`
	BigVAuthors       []int64            `json:"bigv_authors"`
	Readers           []int64            `json:"readers"`
	FollowerCounts    map[int64]int      `json:"follower_counts"`
	Posts             []int64            `json:"posts"`
	CursorDeep        *cursorDeepDataset `json:"cursor_deep,omitempty"`
}

func resolveManifestReaderCardinality(manifest datasetManifest, requested string) (string, error) {
	actual := strings.ToLower(strings.TrimSpace(manifest.ReaderCardinality))
	if actual == "" {
		switch len(manifest.Readers) {
		case 20:
			actual = readerCardinalityHot
		case 1200:
			actual = readerCardinalityDistributed
		case 20000:
			actual = readerCardinalityHigh
		default:
			return "", fmt.Errorf("cannot infer reader cardinality from %d manifest readers", len(manifest.Readers))
		}
	}
	wanted := strings.ToLower(strings.TrimSpace(requested))
	if wanted == "" {
		wanted = actual
	}
	var preset topologyConfig
	if err := applyReaderCardinalityPreset(wanted, &preset); err != nil {
		return "", err
	}
	if actual != wanted || len(manifest.Readers) != preset.Readers {
		return "", fmt.Errorf(
			"manifest reader cardinality %q/%d does not match requested %q/%d",
			actual, len(manifest.Readers), wanted, preset.Readers,
		)
	}
	return wanted, nil
}

func validateManifestSeed(manifest datasetManifest, configured int64) error {
	if manifest.Seed == 0 {
		return fmt.Errorf("manifest %q has no recorded seed; create a new auditable dataset", manifest.RunID)
	}
	if configured != manifest.Seed {
		return fmt.Errorf("manifest seed %d does not match configured seed %d", manifest.Seed, configured)
	}
	return nil
}

func validateManifestStrategy(manifest datasetManifest, requested string) error {
	actual := strings.ToLower(strings.TrimSpace(manifest.Strategy))
	requested = strings.ToLower(strings.TrimSpace(requested))
	if actual == "" {
		return fmt.Errorf("manifest %q has no recorded Feed strategy; create a new auditable dataset", manifest.RunID)
	}
	if actual != requested {
		return fmt.Errorf("manifest Feed strategy %q does not match requested strategy %q", actual, requested)
	}
	return nil
}

type authLoginClient interface {
	Login(context.Context, *authpb.LoginReq, ...grpc.CallOption) (*authpb.AuthResp, error)
}

func createBenchmarkUsers(
	ctx context.Context,
	client userCreator,
	count, concurrency int,
	runID, passwordHash string,
) ([]benchmarkUser, error) {
	if count <= 0 {
		return nil, fmt.Errorf("user count must be positive")
	}
	if concurrency <= 0 {
		return nil, fmt.Errorf("user creation concurrency must be positive")
	}
	if runID == "" || passwordHash == "" {
		return nil, fmt.Errorf("run id and password hash are required")
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	type job struct {
		index int
	}
	jobs := make(chan job)
	users := make([]benchmarkUser, count)
	var firstErr error
	var errMu sync.Mutex
	var workers sync.WaitGroup
	workers.Add(concurrency)
	for workerID := 0; workerID < concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for work := range jobs {
				email := fmt.Sprintf("feed-loadtest+%s-%05d@example.invalid", runID, work.index)
				resp, err := client.Create(runCtx, &userpb.CreateReq{
					Email:        email,
					Nickname:     fmt.Sprintf("Feed Loadtest %s %05d", runID, work.index),
					PasswordHash: passwordHash,
				})
				if err != nil || resp == nil || resp.Id <= 0 {
					if err == nil {
						err = fmt.Errorf("invalid user id")
					}
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("create benchmark user %d: %w", work.index, err)
						cancel()
					}
					errMu.Unlock()
					continue
				}
				users[work.index] = benchmarkUser{ID: resp.Id, Email: email}
			}
		}()
	}
sendJobs:
	for i := 0; i < count; i++ {
		select {
		case <-runCtx.Done():
			break sendJobs
		case jobs <- job{index: i}:
		}
	}
	close(jobs)
	workers.Wait()
	if firstErr != nil {
		created := users[:0]
		for _, user := range users {
			if user.ID > 0 {
				created = append(created, user)
			}
		}
		sort.Slice(created, func(i, j int) bool { return created[i].ID < created[j].ID })
		return created, firstErr
	}
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })
	return users, nil
}

func loginReaders(
	ctx context.Context,
	client authLoginClient,
	users []benchmarkUser,
	password string,
	concurrency int,
) ([]readerIdentity, error) {
	if len(users) == 0 {
		return nil, nil
	}
	if concurrency <= 0 {
		return nil, fmt.Errorf("login concurrency must be positive")
	}
	if concurrency > len(users) {
		concurrency = len(users)
	}
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	type job struct{ index int }
	jobs := make(chan job)
	readers := make([]readerIdentity, len(users))
	var firstErr error
	var errMu sync.Mutex
	var workers sync.WaitGroup
	workers.Add(concurrency)
	for workerID := 0; workerID < concurrency; workerID++ {
		go func() {
			defer workers.Done()
			for work := range jobs {
				user := users[work.index]
				resp, err := client.Login(runCtx, &authpb.LoginReq{
					Identifier: user.Email,
					Password:   password,
					Channel:    "PASSWORD",
				})
				if err == nil && (resp == nil || resp.Token == nil || resp.Token.AccessToken == "") {
					err = fmt.Errorf("missing access token")
				}
				if err != nil {
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("login reader %d: %w", user.ID, err)
						cancel()
					}
					errMu.Unlock()
					continue
				}
				readers[work.index] = readerIdentity{UserID: user.ID, AccessToken: resp.Token.AccessToken}
			}
		}()
	}
sendJobs:
	for index := range users {
		select {
		case <-runCtx.Done():
			break sendJobs
		case jobs <- job{index: index}:
		}
	}
	close(jobs)
	workers.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	return readers, nil
}

func saveManifest(path string, manifest datasetManifest) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create manifest directory: %w", err)
	}
	body, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal manifest: %w", err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o600); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}
	return nil
}

func loadManifest(path string) (datasetManifest, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return datasetManifest{}, fmt.Errorf("read manifest: %w", err)
	}
	var manifest datasetManifest
	if err := json.Unmarshal(body, &manifest); err != nil {
		return datasetManifest{}, fmt.Errorf("decode manifest: %w", err)
	}
	if manifest.RunID == "" || len(manifest.Users) == 0 {
		return datasetManifest{}, fmt.Errorf("manifest is incomplete")
	}
	return manifest, nil
}
