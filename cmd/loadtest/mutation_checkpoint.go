package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	mutationCheckpointVersion     = 1
	mutationCheckpointToolVersion = "feed-loadtest-mutation-v1"
)

type mutationCheckpoint struct {
	Version               int                   `json:"version"`
	ID                    string                `json:"id"`
	RunID                 string                `json:"run_id"`
	CreatedAt             time.Time             `json:"created_at"`
	Strategy              string                `json:"strategy"`
	Seed                  int64                 `json:"seed"`
	ManifestFingerprint   string                `json:"manifest_fingerprint"`
	ToolVersion           string                `json:"tool_version"`
	SubjectBinaryIdentity string                `json:"subject_binary_identity"`
	BaselineFingerprint   string                `json:"baseline_fingerprint"`
	MySQL                 mutationMySQLSnapshot `json:"mysql"`
	Redis                 mutationRedisSnapshot `json:"redis"`
}

type mutationMySQLSnapshot struct {
	AuthorIDs         []int64 `json:"author_ids"`
	PostIDs           []int64 `json:"post_ids"`
	PostFingerprint   string  `json:"post_fingerprint"`
	OutboxIDs         []int64 `json:"outbox_ids"`
	OutboxFingerprint string  `json:"outbox_fingerprint"`
}

type mutationRedisSnapshot struct {
	Identity    string                     `json:"identity"`
	Fingerprint string                     `json:"fingerprint"`
	Keys        []mutationRedisKeySnapshot `json:"keys"`
}

type mutationRedisKeySnapshot struct {
	Key        string `json:"key"`
	Exists     bool   `json:"exists"`
	PTTLMillis int64  `json:"pttl_ms"`
	Dump       []byte `json:"dump,omitempty"`
	DumpSHA256 string `json:"dump_sha256,omitempty"`
}

type mutationRedisKeyIdentity struct {
	Key        string `json:"key"`
	Exists     bool   `json:"exists"`
	PTTLMillis int64  `json:"pttl_ms"`
	DumpSHA256 string `json:"dump_sha256,omitempty"`
}

type mutationBaselineIdentity struct {
	MySQL mutationMySQLSnapshot `json:"mysql"`
	Redis struct {
		Identity string                     `json:"identity"`
		Keys     []mutationRedisKeyIdentity `json:"keys"`
	} `json:"redis"`
}

type mutationCheckpointIdentity struct {
	Version               int                        `json:"version"`
	RunID                 string                     `json:"run_id"`
	Strategy              string                     `json:"strategy"`
	Seed                  int64                      `json:"seed"`
	ManifestFingerprint   string                     `json:"manifest_fingerprint"`
	ToolVersion           string                     `json:"tool_version"`
	SubjectBinaryIdentity string                     `json:"subject_binary_identity"`
	BaselineFingerprint   string                     `json:"baseline_fingerprint"`
	MySQL                 mutationMySQLSnapshot      `json:"mysql"`
	RedisIdentity         string                     `json:"redis_identity"`
	RedisKeys             []mutationRedisKeyIdentity `json:"redis_keys"`
}

func datasetManifestFingerprint(manifest datasetManifest) (string, error) {
	body, err := json.Marshal(manifest)
	if err != nil {
		return "", fmt.Errorf("marshal dataset manifest fingerprint: %w", err)
	}
	digest := sha256.Sum256(body)
	return hex.EncodeToString(digest[:]), nil
}

func finalizeMutationCheckpoint(input mutationCheckpoint) (mutationCheckpoint, error) {
	checkpoint := input
	checkpoint.ID = ""
	checkpoint.BaselineFingerprint = ""
	checkpoint.Strategy = strings.ToLower(strings.TrimSpace(checkpoint.Strategy))
	checkpoint.MySQL.AuthorIDs = sortedUniqueInt64s(checkpoint.MySQL.AuthorIDs)
	checkpoint.MySQL.PostIDs = sortedUniqueInt64s(checkpoint.MySQL.PostIDs)
	checkpoint.MySQL.OutboxIDs = sortedUniqueInt64s(checkpoint.MySQL.OutboxIDs)
	redisSnapshot, err := finalizeMutationRedisSnapshot(checkpoint.Redis)
	if err != nil {
		return mutationCheckpoint{}, err
	}
	checkpoint.Redis = redisSnapshot
	if err := validateMutationCheckpointShape(checkpoint); err != nil {
		return mutationCheckpoint{}, err
	}

	redisKeys := make([]mutationRedisKeyIdentity, len(checkpoint.Redis.Keys))
	for index := range checkpoint.Redis.Keys {
		key := checkpoint.Redis.Keys[index]
		redisKeys[index] = mutationRedisKeyIdentity{
			Key: key.Key, Exists: key.Exists, PTTLMillis: key.PTTLMillis, DumpSHA256: key.DumpSHA256,
		}
	}

	baseline := mutationBaselineIdentity{MySQL: checkpoint.MySQL}
	baseline.Redis.Identity = checkpoint.Redis.Identity
	baseline.Redis.Keys = redisKeys
	baselineFingerprint, err := hashCanonicalJSON(baseline)
	if err != nil {
		return mutationCheckpoint{}, fmt.Errorf("fingerprint mutation baseline: %w", err)
	}
	checkpoint.BaselineFingerprint = baselineFingerprint
	checkpoint.ID, err = hashCanonicalJSON(mutationCheckpointIdentity{
		Version: checkpoint.Version, RunID: checkpoint.RunID, Strategy: checkpoint.Strategy,
		Seed: checkpoint.Seed, ManifestFingerprint: checkpoint.ManifestFingerprint,
		ToolVersion:           checkpoint.ToolVersion,
		SubjectBinaryIdentity: checkpoint.SubjectBinaryIdentity,
		BaselineFingerprint:   checkpoint.BaselineFingerprint, MySQL: checkpoint.MySQL,
		RedisIdentity: checkpoint.Redis.Identity, RedisKeys: redisKeys,
	})
	if err != nil {
		return mutationCheckpoint{}, fmt.Errorf("fingerprint mutation checkpoint: %w", err)
	}
	return checkpoint, nil
}

func finalizeMutationRedisSnapshot(input mutationRedisSnapshot) (mutationRedisSnapshot, error) {
	snapshot := input
	snapshot.Fingerprint = ""
	snapshot.Keys = cloneMutationRedisKeys(snapshot.Keys)
	sort.Slice(snapshot.Keys, func(i, j int) bool { return snapshot.Keys[i].Key < snapshot.Keys[j].Key })
	seen := make(map[string]struct{}, len(snapshot.Keys))
	identities := make([]mutationRedisKeyIdentity, len(snapshot.Keys))
	for index := range snapshot.Keys {
		key := &snapshot.Keys[index]
		if strings.TrimSpace(key.Key) == "" {
			return mutationRedisSnapshot{}, fmt.Errorf("mutation checkpoint contains an empty Redis key")
		}
		if _, exists := seen[key.Key]; exists {
			return mutationRedisSnapshot{}, fmt.Errorf("mutation checkpoint contains duplicate Redis key %q", key.Key)
		}
		seen[key.Key] = struct{}{}
		if key.Exists {
			if len(key.Dump) == 0 {
				return mutationRedisSnapshot{}, fmt.Errorf("mutation checkpoint Redis key %q exists without a DUMP payload", key.Key)
			}
			digest := sha256.Sum256(key.Dump)
			key.DumpSHA256 = hex.EncodeToString(digest[:])
		} else {
			key.Dump = nil
			key.DumpSHA256 = ""
		}
		identities[index] = mutationRedisKeyIdentity{
			Key: key.Key, Exists: key.Exists, PTTLMillis: key.PTTLMillis, DumpSHA256: key.DumpSHA256,
		}
	}
	fingerprint, err := hashCanonicalJSON(struct {
		Identity string                     `json:"identity"`
		Keys     []mutationRedisKeyIdentity `json:"keys"`
	}{Identity: snapshot.Identity, Keys: identities})
	if err != nil {
		return mutationRedisSnapshot{}, fmt.Errorf("fingerprint mutation Redis baseline: %w", err)
	}
	snapshot.Fingerprint = fingerprint
	return snapshot, nil
}

func validateMutationCheckpoint(
	manifest datasetManifest,
	strategy string,
	subjectBinaryIdentity string,
	checkpoint mutationCheckpoint,
) error {
	manifestFingerprint, err := datasetManifestFingerprint(manifest)
	if err != nil {
		return err
	}
	if checkpoint.RunID != manifest.RunID {
		return fmt.Errorf("checkpoint run id %q does not match manifest run id %q", checkpoint.RunID, manifest.RunID)
	}
	strategy = strings.ToLower(strings.TrimSpace(strategy))
	if checkpoint.Strategy != strategy {
		return fmt.Errorf("checkpoint strategy %q does not match requested strategy %q", checkpoint.Strategy, strategy)
	}
	if checkpoint.Seed != manifest.Seed {
		return fmt.Errorf("checkpoint seed %d does not match manifest seed %d", checkpoint.Seed, manifest.Seed)
	}
	if checkpoint.ManifestFingerprint != manifestFingerprint {
		return fmt.Errorf("checkpoint manifest fingerprint does not match current manifest")
	}
	if checkpoint.SubjectBinaryIdentity != subjectBinaryIdentity {
		return fmt.Errorf("checkpoint binary identity does not match current subject binary")
	}
	finalized, err := finalizeMutationCheckpoint(checkpoint)
	if err != nil {
		return err
	}
	if checkpoint.ID == "" || finalized.ID != checkpoint.ID || finalized.BaselineFingerprint != checkpoint.BaselineFingerprint {
		return fmt.Errorf("checkpoint identity or baseline fingerprint is invalid")
	}
	return nil
}

func mutationFeedKeys(manifest datasetManifest) []string {
	all := feedKeysForManifest(manifest)
	keys := make([]string, 0, len(all))
	for _, key := range all {
		if strings.HasPrefix(key, "feed:inbox:") || strings.HasPrefix(key, "feed:bigv:") {
			keys = append(keys, key)
		}
	}
	return keys
}

func validateMutationManifest(manifest datasetManifest, confirmation string) error {
	if err := validateCleanupManifest(manifest, confirmation); err != nil {
		return err
	}
	if manifest.Seed == 0 || strings.TrimSpace(manifest.Strategy) == "" {
		return fmt.Errorf("mutation manifest seed and strategy are required")
	}
	users := make(map[int64]struct{}, len(manifest.Users))
	for _, user := range manifest.Users {
		users[user.ID] = struct{}{}
	}
	for _, authorID := range append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...) {
		if _, ok := users[authorID]; !ok {
			return fmt.Errorf("mutation manifest author %d is not an owned manifest user", authorID)
		}
	}
	for _, readerID := range manifest.Readers {
		if _, ok := users[readerID]; !ok {
			return fmt.Errorf("mutation manifest reader %d is not an owned manifest user", readerID)
		}
	}
	return nil
}

func saveMutationCheckpoint(path string, checkpoint mutationCheckpoint) error {
	if err := validateStoredMutationCheckpoint(checkpoint); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create mutation checkpoint directory: %w", err)
	}
	body, err := json.MarshalIndent(checkpoint, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal mutation checkpoint: %w", err)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("mutation checkpoint %s already exists", path)
		}
		return fmt.Errorf("create mutation checkpoint %s: %w", path, err)
	}
	writeErr := error(nil)
	if _, err := file.Write(append(body, '\n')); err != nil {
		writeErr = fmt.Errorf("write mutation checkpoint %s: %w", path, err)
	} else if err := file.Sync(); err != nil {
		writeErr = fmt.Errorf("sync mutation checkpoint %s: %w", path, err)
	}
	if err := file.Close(); writeErr == nil && err != nil {
		writeErr = fmt.Errorf("close mutation checkpoint %s: %w", path, err)
	}
	return writeErr
}

func loadMutationCheckpoint(path string) (mutationCheckpoint, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return mutationCheckpoint{}, fmt.Errorf("read mutation checkpoint %s: %w", path, err)
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	var checkpoint mutationCheckpoint
	if err := decoder.Decode(&checkpoint); err != nil {
		return mutationCheckpoint{}, fmt.Errorf("decode mutation checkpoint %s: %w", path, err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return mutationCheckpoint{}, fmt.Errorf("decode mutation checkpoint %s: trailing JSON value", path)
		}
		return mutationCheckpoint{}, fmt.Errorf("decode mutation checkpoint %s trailing data: %w", path, err)
	}
	if err := validateStoredMutationCheckpoint(checkpoint); err != nil {
		return mutationCheckpoint{}, fmt.Errorf("validate mutation checkpoint %s: %w", path, err)
	}
	return checkpoint, nil
}

func validateStoredMutationCheckpoint(checkpoint mutationCheckpoint) error {
	finalized, err := finalizeMutationCheckpoint(checkpoint)
	if err != nil {
		return err
	}
	if checkpoint.ID == "" || checkpoint.ID != finalized.ID {
		return fmt.Errorf("mutation checkpoint identity is invalid")
	}
	if checkpoint.BaselineFingerprint == "" || checkpoint.BaselineFingerprint != finalized.BaselineFingerprint {
		return fmt.Errorf("mutation checkpoint baseline fingerprint is invalid")
	}
	if checkpoint.Redis.Fingerprint == "" || checkpoint.Redis.Fingerprint != finalized.Redis.Fingerprint {
		return fmt.Errorf("mutation checkpoint Redis fingerprint is invalid")
	}
	for index := range checkpoint.Redis.Keys {
		if checkpoint.Redis.Keys[index].DumpSHA256 != finalized.Redis.Keys[index].DumpSHA256 {
			return fmt.Errorf("mutation checkpoint Redis DUMP fingerprint is invalid for key %q", checkpoint.Redis.Keys[index].Key)
		}
	}
	return nil
}

func validateMutationCheckpointShape(checkpoint mutationCheckpoint) error {
	if checkpoint.Version != mutationCheckpointVersion {
		return fmt.Errorf("mutation checkpoint version %d is not supported", checkpoint.Version)
	}
	if strings.TrimSpace(checkpoint.RunID) == "" || checkpoint.Seed == 0 || checkpoint.Strategy == "" {
		return fmt.Errorf("mutation checkpoint run id, strategy, and seed are required")
	}
	if checkpoint.ManifestFingerprint == "" || checkpoint.SubjectBinaryIdentity == "" {
		return fmt.Errorf("mutation checkpoint manifest and binary identity are required")
	}
	if checkpoint.ToolVersion != mutationCheckpointToolVersion {
		return fmt.Errorf("mutation checkpoint tool version %q is not supported", checkpoint.ToolVersion)
	}
	if checkpoint.MySQL.PostFingerprint == "" || checkpoint.MySQL.OutboxFingerprint == "" {
		return fmt.Errorf("mutation checkpoint MySQL fingerprints are required")
	}
	if checkpoint.Redis.Identity == "" {
		return fmt.Errorf("mutation checkpoint Redis identity is required")
	}
	seenKeys := make(map[string]struct{}, len(checkpoint.Redis.Keys))
	for _, key := range checkpoint.Redis.Keys {
		if strings.TrimSpace(key.Key) == "" {
			return fmt.Errorf("mutation checkpoint contains an empty Redis key")
		}
		if _, exists := seenKeys[key.Key]; exists {
			return fmt.Errorf("mutation checkpoint contains duplicate Redis key %q", key.Key)
		}
		seenKeys[key.Key] = struct{}{}
		if key.Exists && len(key.Dump) == 0 {
			return fmt.Errorf("mutation checkpoint Redis key %q exists without a DUMP payload", key.Key)
		}
	}
	return nil
}

func sortedUniqueInt64s(values []int64) []int64 {
	result := append([]int64(nil), values...)
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	if len(result) < 2 {
		return result
	}
	write := 1
	for read := 1; read < len(result); read++ {
		if result[read] == result[write-1] {
			continue
		}
		result[write] = result[read]
		write++
	}
	return result[:write]
}

func cloneMutationRedisKeys(values []mutationRedisKeySnapshot) []mutationRedisKeySnapshot {
	result := make([]mutationRedisKeySnapshot, len(values))
	for index, value := range values {
		result[index] = value
		result[index].Dump = append([]byte(nil), value.Dump...)
	}
	return result
}

func hashCanonicalJSON(value any) (string, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	digest := sha256.Sum256(body)
	return hex.EncodeToString(digest[:]), nil
}
