package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFinalizeMutationCheckpointProducesCanonicalStableIdentity(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users: []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
	}
	manifestFingerprint, err := datasetManifestFingerprint(manifest)
	require.NoError(t, err)

	checkpoint := mutationCheckpoint{
		Version: mutationCheckpointVersion, RunID: manifest.RunID, CreatedAt: time.Unix(100, 0).UTC(),
		Strategy: manifest.Strategy, Seed: manifest.Seed, ManifestFingerprint: manifestFingerprint,
		ToolVersion: mutationCheckpointToolVersion, SubjectBinaryIdentity: "knowpost|sha256=abc",
		MySQL: mutationMySQLSnapshot{
			AuthorIDs: []int64{9, 7}, PostIDs: []int64{30, 10}, PostFingerprint: "posts-v1",
			OutboxIDs: []int64{300, 100}, OutboxFingerprint: "outbox-v1",
		},
		Redis: mutationRedisSnapshot{
			Identity: "redis-run-id",
			Keys: []mutationRedisKeySnapshot{
				{Key: "feed:inbox:9", Exists: true, PTTLMillis: -1, Dump: []byte("inbox-9")},
				{Key: "feed:inbox:7", Exists: false, PTTLMillis: -2},
			},
		},
	}

	first, err := finalizeMutationCheckpoint(checkpoint)
	require.NoError(t, err)
	require.NotEmpty(t, first.ID)
	require.NotEmpty(t, first.BaselineFingerprint)
	require.Equal(t, []int64{7, 9}, first.MySQL.AuthorIDs)
	require.Equal(t, []int64{10, 30}, first.MySQL.PostIDs)
	require.Equal(t, "feed:inbox:7", first.Redis.Keys[0].Key)
	require.Empty(t, first.Redis.Keys[0].DumpSHA256)
	require.NotEmpty(t, first.Redis.Keys[1].DumpSHA256)

	reordered := checkpoint
	reordered.CreatedAt = time.Unix(200, 0).UTC()
	reordered.MySQL.AuthorIDs = []int64{7, 9}
	reordered.MySQL.PostIDs = []int64{10, 30}
	reordered.MySQL.OutboxIDs = []int64{100, 300}
	reordered.Redis.Keys = []mutationRedisKeySnapshot{checkpoint.Redis.Keys[1], checkpoint.Redis.Keys[0]}
	second, err := finalizeMutationCheckpoint(reordered)
	require.NoError(t, err)
	require.Equal(t, first.ID, second.ID, "creation time and input ordering must not change logical identity")
	require.Equal(t, first.BaselineFingerprint, second.BaselineFingerprint)

	changed := checkpoint
	changed.Redis.Keys = append([]mutationRedisKeySnapshot(nil), checkpoint.Redis.Keys...)
	changed.Redis.Keys[0].Dump = []byte("changed")
	third, err := finalizeMutationCheckpoint(changed)
	require.NoError(t, err)
	require.NotEqual(t, first.ID, third.ID)
}

func TestValidateMutationCheckpointRequiresExactManifestAndBinary(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users: []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
	}
	fingerprint, err := datasetManifestFingerprint(manifest)
	require.NoError(t, err)
	checkpoint, err := finalizeMutationCheckpoint(mutationCheckpoint{
		Version: mutationCheckpointVersion, RunID: manifest.RunID, Strategy: manifest.Strategy,
		Seed: manifest.Seed, ManifestFingerprint: fingerprint, ToolVersion: mutationCheckpointToolVersion,
		SubjectBinaryIdentity: "binary-a",
		MySQL:                 mutationMySQLSnapshot{AuthorIDs: []int64{7}, PostFingerprint: "posts", OutboxFingerprint: "outbox"},
		Redis:                 mutationRedisSnapshot{Identity: "redis-a"},
	})
	require.NoError(t, err)

	require.NoError(t, validateMutationCheckpoint(manifest, "hybrid", "binary-a", checkpoint))
	require.ErrorContains(t, validateMutationCheckpoint(manifest, "push", "binary-a", checkpoint), "strategy")
	require.ErrorContains(t, validateMutationCheckpoint(manifest, "hybrid", "binary-b", checkpoint), "binary")
	missingTool := checkpoint
	missingTool.ToolVersion = ""
	_, err = finalizeMutationCheckpoint(missingTool)
	require.ErrorContains(t, err, "tool version")

	changedManifest := manifest
	changedManifest.Readers = []int64{99}
	require.ErrorContains(t, validateMutationCheckpoint(changedManifest, "hybrid", "binary-a", checkpoint), "manifest fingerprint")
}

func TestMutationFeedKeysAreExactAndManifestScoped(t *testing.T) {
	manifest := datasetManifest{
		Users:         []benchmarkUser{{ID: 1}, {ID: 3}},
		NormalAuthors: []int64{7},
		BigVAuthors:   []int64{8},
		Posts:         []int64{99},
	}

	keys := mutationFeedKeys(manifest)

	require.Equal(t, []string{
		"feed:bigv:7", "feed:bigv:8",
		"feed:inbox:1", "feed:inbox:3", "feed:inbox:7", "feed:inbox:8",
	}, keys)
	require.NotContains(t, keys, "feed:fanout:processing:99")
}

func TestValidateMutationManifestRequiresExactOwnershipAndKnownAuthors(t *testing.T) {
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users: []benchmarkUser{
			{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"},
			{ID: 8, Email: "feed-loadtest+run-42-00001@example.invalid"},
		},
		NormalAuthors: []int64{7}, BigVAuthors: []int64{8}, Readers: []int64{7},
	}

	require.NoError(t, validateMutationManifest(manifest, "run-42"))
	require.ErrorContains(t, validateMutationManifest(manifest, "wrong-run"), "confirmation")

	unknownAuthor := manifest
	unknownAuthor.BigVAuthors = []int64{99}
	require.ErrorContains(t, validateMutationManifest(unknownAuthor, "run-42"), "author 99")
}

func TestMutationCheckpointFileRefusesOverwriteAndDetectsTampering(t *testing.T) {
	checkpoint := validMutationCheckpointForTest(t)
	path := filepath.Join(t.TempDir(), "checkpoint.json")

	require.NoError(t, saveMutationCheckpoint(path, checkpoint))
	info, err := os.Stat(path)
	require.NoError(t, err)
	if runtime.GOOS != "windows" {
		require.Equal(t, os.FileMode(0o600), info.Mode().Perm())
	}
	require.ErrorContains(t, saveMutationCheckpoint(path, checkpoint), "already exists")

	loaded, err := loadMutationCheckpoint(path)
	require.NoError(t, err)
	require.Equal(t, checkpoint.ID, loaded.ID)

	body, err := os.ReadFile(path)
	require.NoError(t, err)
	body[len(body)-3] ^= 1
	require.NoError(t, os.WriteFile(path, body, 0o600))
	_, err = loadMutationCheckpoint(path)
	require.Error(t, err)
}

func TestMutationDerivedRedisKeysAreExact(t *testing.T) {
	keys := mutationDerivedRedisKeys(mutationMySQLRestoreResult{
		PostIDs: []int64{30}, OutboxIDs: []int64{300},
	})

	require.Equal(t, []string{
		"agg:v1:knowpost:30",
		"cache:knowPosts:id:30",
		"cache:outbox:id:300",
		"cnt:v1:knowpost:30",
		"feed:fanout:processing:30",
	}, keys)
}

func TestDefaultMutationCheckpointPathIsOutsideReportNamespace(t *testing.T) {
	path := defaultMutationCheckpointPath("results/feed-loadtest", "run-42")
	require.Equal(t, filepath.FromSlash("results/feed-loadtest/checkpoints/run-42.json"), path)
}

func TestMutationRedisTTLValidationSeparatesStrictRestoreFromAgedVerify(t *testing.T) {
	require.True(t, mutationRedisTTLMatches(1000, 995, 10, false))
	require.False(t, mutationRedisTTLMatches(1000, 500, 10, false))

	require.True(t, mutationRedisTTLMatches(1000, 500, 10, true))
	require.True(t, mutationRedisTTLMatches(1000, 1010, 10, true))
	require.False(t, mutationRedisTTLMatches(1000, 1011, 10, true))
	require.False(t, mutationRedisTTLMatches(1000, -2, 10, true))

	require.True(t, mutationRedisTTLMatches(-1, -1, 10, true))
	require.False(t, mutationRedisTTLMatches(-1, 1000, 10, true))
}

func validMutationCheckpointForTest(t *testing.T) mutationCheckpoint {
	t.Helper()
	manifest := datasetManifest{
		RunID: "run-42", Seed: 42, Strategy: "hybrid",
		Users: []benchmarkUser{{ID: 7, Email: "feed-loadtest+run-42-00000@example.invalid"}},
	}
	fingerprint, err := datasetManifestFingerprint(manifest)
	require.NoError(t, err)
	checkpoint, err := finalizeMutationCheckpoint(mutationCheckpoint{
		Version: mutationCheckpointVersion, RunID: manifest.RunID, CreatedAt: time.Unix(100, 0).UTC(),
		Strategy: manifest.Strategy, Seed: manifest.Seed, ManifestFingerprint: fingerprint,
		ToolVersion: mutationCheckpointToolVersion, SubjectBinaryIdentity: "binary-a",
		MySQL: mutationMySQLSnapshot{
			AuthorIDs: []int64{7}, PostIDs: []int64{10}, PostFingerprint: "posts",
			OutboxIDs: []int64{100}, OutboxFingerprint: "outbox",
		},
		Redis: mutationRedisSnapshot{
			Identity: "redis-a",
			Keys:     []mutationRedisKeySnapshot{{Key: "feed:inbox:7", Exists: true, PTTLMillis: -1, Dump: []byte("dump")}},
		},
	})
	require.NoError(t, err)
	return checkpoint
}
