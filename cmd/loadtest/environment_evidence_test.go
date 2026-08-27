package main

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/require"
)

func TestCaptureGitPatchHashHandlesUnicodeUntrackedPath(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}
	repo := t.TempDir()
	runGitForEvidenceTest(t, repo, "init")
	runGitForEvidenceTest(t, repo, "config", "user.email", "feed-loadtest@example.invalid")
	runGitForEvidenceTest(t, repo, "config", "user.name", "Feed Loadtest")
	runGitForEvidenceTest(t, repo, "config", "core.quotePath", "true")
	require.NoError(t, os.WriteFile(filepath.Join(repo, "tracked.txt"), []byte("tracked"), 0o600))
	runGitForEvidenceTest(t, repo, "add", "tracked.txt")
	runGitForEvidenceTest(t, repo, "commit", "-m", "initial")

	unicodePath := filepath.Join(repo, "docs", "API接口文档.md")
	require.NoError(t, os.MkdirAll(filepath.Dir(unicodePath), 0o700))
	require.NoError(t, os.WriteFile(unicodePath, []byte("untracked"), 0o600))
	t.Chdir(repo)

	digest, count, err := captureGitPatchHash(context.Background())

	require.NoError(t, err)
	require.Len(t, digest, 64)
	require.Equal(t, 1, count)
}

func runGitForEvidenceTest(t *testing.T, dir string, args ...string) {
	t.Helper()
	command := exec.Command("git", args...)
	command.Dir = dir
	output, err := command.CombinedOutput()
	require.NoError(t, err, "%s", output)
}

func TestCaptureLocalSubjectIdentityIncludesExecutableHashAndStartTime(t *testing.T) {
	current, err := process.NewProcess(int32(os.Getpid()))
	require.NoError(t, err)
	executable, err := current.Exe()
	require.NoError(t, err)
	createdAt, err := current.CreateTime()
	require.NoError(t, err)
	state := localProcessState{Topology: "compose-middleware-local-services", Services: []localProcessRecord{{
		Name: "knowpost", PID: int32(os.Getpid()), Executable: executable, ProcessStartUnixMS: createdAt,
	}}}
	body, err := json.Marshal(state)
	require.NoError(t, err)
	statePath := filepath.Join(t.TempDir(), "state.json")
	require.NoError(t, os.WriteFile(statePath, body, 0o600))

	binaryIdentity, processIdentity, err := captureSubjectIdentity(t.Context(), benchmarkConfig{Execution: executionConfig{
		Topology: "compose-middleware-local-services", LocalStateFile: statePath,
	}})

	require.NoError(t, err)
	require.Contains(t, binaryIdentity, "knowpost|sha256=")
	require.Contains(t, processIdentity, "knowpost|pid=")
	require.Contains(t, processIdentity, "start_ms=")
}

func TestEnvironmentCompatibilityFingerprintIncludesSubjectIdentity(t *testing.T) {
	first := environmentCompatibilityFingerprint(map[string]string{"subject_binary_identity": "image-a"})
	second := environmentCompatibilityFingerprint(map[string]string{"subject_binary_identity": "image-b"})

	require.NotEqual(t, first, second)
}

func TestEnvironmentCompatibilityFingerprintIncludesLoggingPreset(t *testing.T) {
	first := environmentCompatibilityFingerprint(map[string]string{"logging_preset": "development"})
	second := environmentCompatibilityFingerprint(map[string]string{"logging_preset": "benchmark-error-only"})

	require.NotEqual(t, first, second)
}

func TestEnvironmentCompatibilityFingerprintIncludesMutationCheckpoint(t *testing.T) {
	first := environmentCompatibilityFingerprint(map[string]string{
		"mutation_checkpoint_id": "checkpoint-a", "mutation_baseline_fingerprint": "baseline-a",
	})
	second := environmentCompatibilityFingerprint(map[string]string{
		"mutation_checkpoint_id": "checkpoint-b", "mutation_baseline_fingerprint": "baseline-a",
	})
	third := environmentCompatibilityFingerprint(map[string]string{
		"mutation_checkpoint_id": "checkpoint-a", "mutation_baseline_fingerprint": "baseline-b",
	})

	require.NotEqual(t, first, second)
	require.NotEqual(t, first, third)
}
