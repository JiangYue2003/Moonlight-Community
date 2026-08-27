package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

func captureRunEnvironmentEvidence(
	ctx context.Context,
	cfg benchmarkConfig,
	manifest datasetManifest,
) (map[string]string, error) {
	commit, err := commandOutput(ctx, "git", "rev-parse", "HEAD")
	if err != nil {
		return nil, fmt.Errorf("capture Git commit: %w", err)
	}
	status, err := commandOutput(ctx, "git", "status", "--porcelain", "--untracked-files=normal")
	if err != nil {
		return nil, fmt.Errorf("capture Git dirty state: %w", err)
	}
	manifestBody, err := json.Marshal(manifest)
	if err != nil {
		return nil, fmt.Errorf("fingerprint dataset manifest: %w", err)
	}
	manifestDigest := sha256.Sum256(manifestBody)
	statusDigest := sha256.Sum256([]byte(status))
	patchHash, untrackedFiles, err := captureGitPatchHash(ctx)
	if err != nil {
		return nil, err
	}
	seed := fmt.Sprintf("%d", manifest.Seed)
	if manifest.Seed == 0 {
		seed = "legacy-unknown"
	}
	addresses := []string{
		"gateway=" + cfg.GatewayURL,
		"redis=" + cfg.RedisAddr,
		"knowpost_etcd=" + strings.Join(cfg.KnowPostRpc.Etcd.Hosts, ","),
		"relation_etcd=" + strings.Join(cfg.RelationRpc.Etcd.Hosts, ","),
		"counter_etcd=" + strings.Join(cfg.CounterRpc.Etcd.Hosts, ","),
		"user_etcd=" + strings.Join(cfg.UserRpc.Etcd.Hosts, ","),
		"kafka=" + cfg.Monitor.KafkaContainer + "/" + cfg.Monitor.KafkaBootstrap,
		"prometheus=" + cfg.Monitor.PrometheusURL,
	}
	sort.Strings(addresses)
	resourceLimits, err := captureResourceLimits(ctx, cfg)
	if err != nil {
		return nil, err
	}
	subjectBinaryIdentity, subjectProcessIdentity, err := captureSubjectIdentity(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"git_commit":               strings.TrimSpace(commit),
		"git_dirty":                fmt.Sprintf("%t", strings.TrimSpace(status) != ""),
		"git_status_hash":          hex.EncodeToString(statusDigest[:]),
		"git_patch_hash":           patchHash,
		"git_untracked_files":      fmt.Sprintf("%d", untrackedFiles),
		"dataset_seed":             seed,
		"dataset_fingerprint":      hex.EncodeToString(manifestDigest[:]),
		"gomaxprocs":               fmt.Sprintf("%d", runtime.GOMAXPROCS(0)),
		"logical_cpus":             fmt.Sprintf("%d", runtime.NumCPU()),
		"service_addresses":        strings.Join(addresses, ";"),
		"resource_limits":          resourceLimits,
		"subject_binary_identity":  subjectBinaryIdentity,
		"subject_process_identity": subjectProcessIdentity,
	}, nil
}

func captureGitPatchHash(ctx context.Context) (string, int, error) {
	trackedPatch, err := commandOutput(ctx, "git", "diff", "--binary", "HEAD", "--", ".")
	if err != nil {
		return "", 0, fmt.Errorf("capture tracked Git patch: %w", err)
	}
	untrackedRaw, err := commandOutput(
		ctx,
		"git", "-c", "core.quotePath=false", "ls-files", "-z", "--others", "--exclude-standard",
	)
	if err != nil {
		return "", 0, fmt.Errorf("list untracked Git evidence: %w", err)
	}
	files := make([]string, 0)
	for _, path := range strings.Split(untrackedRaw, "\x00") {
		if path == "" || excludedGitEvidencePath(path) {
			continue
		}
		files = append(files, path)
	}
	sort.Strings(files)
	digest := sha256.New()
	_, _ = io.WriteString(digest, trackedPatch)
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return "", 0, fmt.Errorf("open untracked Git evidence %s: %w", path, err)
		}
		_, _ = io.WriteString(digest, "\x00"+path+"\x00")
		_, copyErr := io.Copy(digest, file)
		closeErr := file.Close()
		if copyErr != nil {
			return "", 0, fmt.Errorf("hash untracked Git evidence %s: %w", path, copyErr)
		}
		if closeErr != nil {
			return "", 0, fmt.Errorf("close untracked Git evidence %s: %w", path, closeErr)
		}
	}
	return hex.EncodeToString(digest.Sum(nil)), len(files), nil
}

func excludedGitEvidencePath(path string) bool {
	normalized := strings.ReplaceAll(path, "\\", "/")
	for _, prefix := range []string{".tmp/", "results/", "logs/"} {
		if strings.HasPrefix(normalized, prefix) {
			return true
		}
	}
	return strings.HasSuffix(strings.ToLower(normalized), ".exe")
}

func environmentCompatibilityFingerprint(environment map[string]string) string {
	keys := []string{
		"git_commit", "git_patch_hash", "dataset_run_id", "dataset_seed", "dataset_fingerprint",
		"reader_cardinality", "reader_count", "topology", "cache_state", "feature_preset", "gateway_auth_mode",
		"gomaxprocs", "logical_cpus", "resource_limits", "service_addresses", "client_host",
		"subject_binary_identity",
		"go_version", "request_timeout", "page", "size", "configured_duration", "configured_requests",
		"warmup", "effective_strategy", "observability_enabled", "route_snapshot_enabled",
		"relation_epoch_consumer_enabled", "safety_epoch_consumer_enabled", "combined_pipeline_enabled",
		"cursor_pagination_enabled", "page_cache_mode",
		"logging_preset", "logging_level", "logging_stat", "gateway_access_log", "rpc_stat_middleware",
		"sql_statement_info", "logging_persistence",
		"feed_metrics_primed",
		"mutation_checkpoint_id", "mutation_baseline_fingerprint", "mutation_checkpoint_created_at",
		"mutation_checkpoint_tool_version",
	}
	digest := sha256.New()
	for _, key := range keys {
		_, _ = io.WriteString(digest, key+"="+environment[key]+"\n")
	}
	return hex.EncodeToString(digest.Sum(nil))
}

func captureSubjectIdentity(ctx context.Context, cfg benchmarkConfig) (string, string, error) {
	if cfg.Execution.Topology == "compose-middleware-local-services" {
		body, err := os.ReadFile(cfg.Execution.LocalStateFile)
		if err != nil {
			return "", "", fmt.Errorf("read local subject state %s: %w", cfg.Execution.LocalStateFile, err)
		}
		var state localProcessState
		if err := json.Unmarshal(body, &state); err != nil {
			return "", "", fmt.Errorf("decode local subject state %s: %w", cfg.Execution.LocalStateFile, err)
		}
		if state.Topology != "compose-middleware-local-services" || len(state.Services) == 0 {
			return "", "", fmt.Errorf("local subject state %s is incomplete", cfg.Execution.LocalStateFile)
		}
		binaries := make([]string, 0, len(state.Services))
		processes := make([]string, 0, len(state.Services))
		for _, service := range state.Services {
			digest, err := hashFile(service.Executable)
			if err != nil {
				return "", "", fmt.Errorf("hash local subject %s: %w", service.Name, err)
			}
			binaries = append(binaries, fmt.Sprintf("%s|sha256=%s", service.Name, digest))
			processes = append(processes, fmt.Sprintf(
				"%s|pid=%d|start_ms=%d",
				service.Name, service.PID, service.ProcessStartUnixMS,
			))
		}
		sort.Strings(binaries)
		sort.Strings(processes)
		return strings.Join(binaries, ";"), strings.Join(processes, ";"), nil
	}
	if len(cfg.Monitor.DockerContainers) == 0 {
		return "", "", fmt.Errorf("capture Docker subject identity: no containers configured")
	}
	args := []string{"inspect"}
	args = append(args, cfg.Monitor.DockerContainers...)
	raw, err := commandOutput(ctx, "docker", args...)
	if err != nil {
		return "", "", fmt.Errorf("capture Docker subject identity: %w", err)
	}
	var inspected []struct {
		ID    string `json:"Id"`
		Name  string `json:"Name"`
		Image string `json:"Image"`
		State struct {
			StartedAt string `json:"StartedAt"`
		} `json:"State"`
	}
	if err := json.Unmarshal([]byte(raw), &inspected); err != nil {
		return "", "", fmt.Errorf("decode Docker subject identity: %w", err)
	}
	if len(inspected) != len(cfg.Monitor.DockerContainers) {
		return "", "", fmt.Errorf("capture Docker subject identity: got %d identities for %d containers", len(inspected), len(cfg.Monitor.DockerContainers))
	}
	binaries := make([]string, 0, len(inspected))
	processes := make([]string, 0, len(inspected))
	for _, item := range inspected {
		name := strings.TrimPrefix(item.Name, "/")
		if name == "" || item.Image == "" || item.ID == "" || item.State.StartedAt == "" {
			return "", "", fmt.Errorf("capture Docker subject identity: incomplete container %q", name)
		}
		binaries = append(binaries, fmt.Sprintf("%s|image_id=%s", name, item.Image))
		processes = append(processes, fmt.Sprintf("%s|container_id=%s|started_at=%s", name, item.ID, item.State.StartedAt))
	}
	sort.Strings(binaries)
	sort.Strings(processes)
	return strings.Join(binaries, ";"), strings.Join(processes, ";"), nil
}

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	digest := sha256.New()
	if _, err := io.Copy(digest, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(digest.Sum(nil)), nil
}

func captureResourceLimits(ctx context.Context, cfg benchmarkConfig) (string, error) {
	if cfg.Execution.Topology == "compose-middleware-local-services" {
		return fmt.Sprintf("local-processes:host;logical_cpus=%d;gomaxprocs=%d", runtime.NumCPU(), runtime.GOMAXPROCS(0)), nil
	}
	if len(cfg.Monitor.DockerContainers) == 0 {
		return "", fmt.Errorf("capture Docker resource limits: no containers configured")
	}
	args := []string{
		"inspect", "--format",
		"{{.Name}}|nano_cpus={{.HostConfig.NanoCpus}}|cpu_quota={{.HostConfig.CpuQuota}}|cpu_period={{.HostConfig.CpuPeriod}}|memory={{.HostConfig.Memory}}",
	}
	args = append(args, cfg.Monitor.DockerContainers...)
	raw, err := commandOutput(ctx, "docker", args...)
	if err != nil {
		return "", fmt.Errorf("capture Docker resource limits: %w", err)
	}
	lines := strings.Fields(strings.TrimSpace(raw))
	sort.Strings(lines)
	return strings.Join(lines, ";"), nil
}

func commandOutput(ctx context.Context, name string, args ...string) (string, error) {
	raw, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %s: %w: %s", name, strings.Join(args, " "), err, strings.TrimSpace(string(raw)))
	}
	return string(raw), nil
}
