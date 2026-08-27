package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteCleanupResultSeparatesDryRunFromExecutedCleanup(t *testing.T) {
	root := t.TempDir()

	preview, err := writeCleanupResult(root, cleanupResult{RunID: "run-42", DryRun: true})
	require.NoError(t, err)
	require.Equal(t, "cleanup-preview.json", filepath.Base(preview))

	executed, err := writeCleanupResult(root, cleanupResult{RunID: "run-42"})
	require.NoError(t, err)
	require.Equal(t, "cleanup.json", filepath.Base(executed))
}

func TestValidateCleanupManifestRejectsNonBenchmarkIdentity(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42",
		Users: []benchmarkUser{{ID: 7, Email: "real-user@example.com"}},
	}

	err := validateCleanupManifest(manifest, "run-42")

	require.ErrorContains(t, err, "not owned")
}

func TestValidateCleanupManifestRequiresExactConfirmation(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42",
		Users: []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
	}

	require.ErrorContains(t, validateCleanupManifest(manifest, "run-41"), "confirmation")
	require.NoError(t, validateCleanupManifest(manifest, "run-42"))
}

func TestValidateVerifiedManifestUsersRequiresEveryManifestUser(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42",
		Users: []benchmarkUser{
			{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"},
			{ID: 8, Email: "feed-loadtest+run-42-00001@example.invalid"},
		},
	}

	err := validateVerifiedManifestUsers(manifest, map[int64]string{7: manifest.Users[0].Email})

	require.ErrorContains(t, err, "missing manifest user 8")
	require.NoError(t, validateVerifiedManifestUsers(manifest, map[int64]string{
		7: manifest.Users[0].Email,
		8: manifest.Users[1].Email,
	}))
}

func TestCleanupRedisKeysStayScopedToManifestRecords(t *testing.T) {
	manifest := datasetManifest{
		RunID:         "run-42",
		Users:         []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
		NormalAuthors: []int64{7},
	}
	records := cleanupDatabaseRecords{
		PostIDs:   []int64{99},
		Following: []cleanupRelationRow{{ID: 101, FromUserID: 7, ToUserID: 8}},
		Follower:  []cleanupRelationRow{{ID: 102, FromUserID: 7, ToUserID: 8}},
		Outbox:    []cleanupOutboxRow{{ID: 103, Type: "FollowCreated"}},
	}

	keys := cleanupRedisKeys(manifest, records)

	require.Contains(t, keys, "feed:inbox:7")
	require.Contains(t, keys, "feed:bigv:7")
	require.Contains(t, keys, "cache:users:id:7")
	require.Contains(t, keys, "cache:users:email:feed-loadtest+run-42-00000@example.invalid")
	require.Contains(t, keys, "cache:knowPosts:id:99")
	require.Contains(t, keys, "cache:following:fromUserId:toUserId:7:8")
	require.Contains(t, keys, "cache:follower:toUserId:fromUserId:8:7")
	require.Contains(t, keys, "dedup:rel:FollowCreated:103")
	require.NotContains(t, keys, "feed:inbox:8")
}
