package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestBenchmarkLoggingConfigConversionPreservesNestedYAML(t *testing.T) {
	pwsh, err := exec.LookPath("pwsh")
	if err != nil {
		t.Skip("pwsh is not available")
	}
	root, err := filepath.Abs(filepath.Join("..", ".."))
	require.NoError(t, err)
	script := filepath.Join(root, "scripts", "start-feed-local-services.ps1")
	tests := []struct {
		service string
		source  string
		check   func(*testing.T, string)
	}{
		{
			service: "counter",
			source:  filepath.Join(root, "services", "counter", "cmd", "counter", "etc", "counter-docker.yaml"),
			check: func(t *testing.T, body string) {
				require.Regexp(t, `(?m)^Prometheus:`, body)
				require.Contains(t, body, "\nMiddlewares:\n  Stat: false")
				require.Contains(t, body, "\n  Level: error")
				require.Contains(t, body, "\n  Stat: false")
			},
		},
		{
			service: "gateway",
			source:  filepath.Join(root, "services", "gateway", "etc", "gateway-docker.yaml"),
			check: func(t *testing.T, body string) {
				require.Contains(t, body, "\nLog:\n  Mode: console\n  Level: error\n  Stat: false")
				require.Contains(t, body, "\nHTTPAccessLog: false")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.service, func(t *testing.T) {
			destination := filepath.Join(t.TempDir(), tt.service+".yaml")
			quote := func(value string) string { return strings.ReplaceAll(value, "'", "''") }
			commandText := strings.Join([]string{
				"$tokens=$null; $parseErrors=$null",
				"$ast=[System.Management.Automation.Language.Parser]::ParseFile('" + quote(script) + "',[ref]$tokens,[ref]$parseErrors)",
				"if($parseErrors.Count -gt 0){throw ($parseErrors -join '; ')}",
				"$fn=$ast.Find({param($node) $node -is [System.Management.Automation.Language.FunctionDefinitionAst] -and $node.Name -eq 'Convert-LocalConfig'},$true)",
				"Invoke-Expression $fn.Extent.Text",
				"$LoggingPreset='benchmark-error-only'",
				"Convert-LocalConfig -Source '" + quote(tt.source) + "' -Destination '" + quote(destination) + "' -Service '" + tt.service + "'",
			}, "; ")
			output, err := exec.Command(pwsh, "-NoProfile", "-NonInteractive", "-Command", commandText).CombinedOutput()
			require.NoError(t, err, "%s", output)
			body, err := os.ReadFile(destination)
			require.NoError(t, err)
			tt.check(t, strings.ReplaceAll(string(body), "\r\n", "\n"))
		})
	}
}

func TestLocalServiceScriptsProtectExistingProcessesAndImages(t *testing.T) {
	root := filepath.Clean(filepath.Join("..", ".."))
	startBody, err := os.ReadFile(filepath.Join(root, "scripts", "start-feed-local-services.ps1"))
	require.NoError(t, err)
	stopBody, err := os.ReadFile(filepath.Join(root, "scripts", "stop-feed-local-services.ps1"))
	require.NoError(t, err)

	start := string(startBody)
	stop := string(stopBody)
	require.Contains(t, start, "SupportsShouldProcess")
	require.Contains(t, start, "$WhatIfPreference")
	require.Contains(t, start, `"--pull", "never"`)
	require.Contains(t, start, `"--no-build"`)
	require.NotContains(t, start, "docker compose pull")
	require.NotContains(t, start, "Stop-ProcessOnPort")
	require.Contains(t, start, "-WindowStyle Hidden")
	require.Contains(t, start, "netsh interface ipv4 show excludedportrange protocol=tcp")
	require.Contains(t, start, "19104")
	require.Contains(t, start, "Get-NetTCPConnection")
	require.Contains(t, start, "OwningProcess")
	require.Contains(t, start, "Process.HasExited")
	require.Contains(t, start, "process_start_unix_ms")
	require.Contains(t, start, "startup cleanup refused")
	require.Contains(t, start, "actualStartUnixMs")
	require.Contains(t, start, "Name = \"user\"; Package = \"./services/user/cmd/user\"; Config = \"services/user/cmd/user/etc/user-docker.yaml\"; Ports = @(20002) }")
	require.Contains(t, start, "Name = \"storage\"; Package = \"./services/storage/cmd/storage\"; Config = \"services/storage/cmd/storage/etc/storage-docker.yaml\"; Ports = @(20013) }")
	require.Contains(t, start, `"0.0.0.0:9002" = "127.0.0.1:20002"`)
	require.Contains(t, start, `"Port: 8080" = "Port: 18080"`)
	require.Contains(t, start, "gateway_http = 18080")
	require.Contains(t, start, "compose-middleware-local-services")
	require.Contains(t, start, "ConvertTo-Json")
	require.Contains(t, start, `ValidateSet("development", "benchmark-error-only")`)
	require.Contains(t, start, `logging = if ($LoggingPreset -eq "benchmark-error-only")`)
	require.Contains(t, start, `Level: error`)
	require.Contains(t, start, `Stat: false`)
	require.Contains(t, start, `Log: false`)

	matrixBody, err := os.ReadFile(filepath.Join(root, "cmd", "loadtest", "run_feed_matrix.ps1"))
	require.NoError(t, err)
	matrix := string(matrixBody)
	require.Contains(t, matrix, `benchmark-error-only`)
	require.Contains(t, matrix, `MaxBenchmarkLogBytes`)
	require.Contains(t, matrix, `Assert-BenchmarkLogsBounded`)

	require.Contains(t, stop, "SupportsShouldProcess")
	require.Contains(t, stop, "$WhatIfPreference")
	require.Contains(t, stop, "Get-Content -LiteralPath $StateFile")
	require.Contains(t, stop, "process_start_unix_ms")
	require.Contains(t, stop, "actualStartUnixMs")
	require.Contains(t, stop, "Stop-Process -Id $processId")
}

func TestApplyExecutionTopologyUsesLocalObservablePorts(t *testing.T) {
	cfg := benchmarkConfig{
		Execution: executionConfig{Topology: "compose-middleware-local-services", LocalStateFile: ".tmp/custom-state.json"},
		Monitor: monitorConfig{
			StrategyContainer: "zg-knowpost", PrometheusContainer: "zg-knowpost",
			PrometheusURL: "http://127.0.0.1:9104/metrics",
		},
	}

	applyExecutionTopology(&cfg)

	require.Equal(t, "http://127.0.0.1:18080", cfg.GatewayURL)
	require.Empty(t, cfg.Monitor.StrategyContainer)
	require.Empty(t, cfg.Monitor.PrometheusContainer)
	require.Equal(t, "http://127.0.0.1:19104/metrics", cfg.Monitor.PrometheusURL)
	require.Equal(t, ".tmp/custom-state.json", cfg.Monitor.LocalStateFile)
	require.Equal(t, cfg.Execution.LocalStateFile, cfg.Monitor.LocalStateFile)
	require.Contains(t, cfg.Monitor.DockerContainers, "zg-kafka")
}

func TestFeedMatrixScriptIncludesAuditedSupplementalSequence(t *testing.T) {
	body, err := os.ReadFile("run_feed_matrix.ps1")
	require.NoError(t, err)
	text := string(body)

	for _, snippet := range []string{
		"Invoke-SupplementalScenarios",
		`"-scenario", "burst-read"`,
		`"-scenario", "steady-read"`,
	} {
		require.Contains(t, text, snippet)
	}
}

func TestFormalMutationMatrixUsesCheckpointAndDurationTrials(t *testing.T) {
	body, err := os.ReadFile("run_feed_mutation_matrix.ps1")
	require.NoError(t, err)
	text := string(body)

	for _, snippet := range []string{
		`ValidateSet("publish", "mixed-90-10", "mixed-80-20")`,
		`-phase", "verify-mutation-checkpoint"`,
		`"-confirm-mutation", $RunId`,
		`"-checkpoint", $CheckpointPath`,
		`"-requests", "0"`,
		`"-duration", "${DurationSeconds}s"`,
		`"-warmup", "${WarmupSeconds}s"`,
		`"-trials", "$Trials"`,
		`Assert-BenchmarkLogsBounded`,
		`Assert-FreshReportNamespace`,
		`Assert-UniqueValues`,
		`Write-ExpectedMatrixManifest`,
		`Assert-ExpectedMutationMatrix`,
	} {
		require.Contains(t, text, snippet)
	}
}

func TestValidateExecutionConfigRejectsWarmCacheLabelsForControl(t *testing.T) {
	base := executionConfig{
		Topology: "compose-full", FeaturePreset: "control", ClientCPULimitPercent: 90,
		CacheWarmConcurrency: 1,
	}
	for _, state := range []string{"l1-warm", "l2-warm", "expire-together"} {
		cfg := base
		cfg.CacheState = state
		require.ErrorContains(t, validateExecutionConfig(cfg), "control")
	}
	base.CacheState = "cold"
	require.NoError(t, validateExecutionConfig(base))

	cursorAB := base
	cursorAB.FeaturePreset = "cursor-ab"
	for _, state := range []string{"l1-warm", "l2-warm", "expire-together"} {
		cursorAB.CacheState = state
		require.ErrorContains(t, validateExecutionConfig(cursorAB), "cursor-ab")
	}
	cursorAB.CacheState = "cold"
	require.NoError(t, validateExecutionConfig(cursorAB))
}

func TestCursorABScriptLocksComparablePaginationContract(t *testing.T) {
	body, err := os.ReadFile("run_feed_cursor_ab.ps1")
	require.NoError(t, err)
	text := string(body)
	for _, snippet := range []string{
		`"-reader-cardinality", "cursor-deep"`,
		`"-cache-state", "cold"`,
		`"-feature-preset", "cursor-ab"`,
		`"-strategy", "hybrid"`,
		`"-requests", "0"`,
		`[string]$GatewayAuthMode = "signed"`,
		`"-gateway-auth-mode", $GatewayAuthMode`,
		`"page50"`,
		`"cursor-page50"`,
		`"sequential-50"`,
		`"same-second-cursor"`,
		`Assert-CompleteReports`,
		`stop before a higher load`,
	} {
		require.Contains(t, text, snippet)
	}
}

func TestLoadBenchmarkConfigFromCheckedInYAML(t *testing.T) {
	t.Setenv("FEED_LOADTEST_PASSWORD", "")
	var cfg benchmarkConfig

	err := conf.Load("load_test.yaml", &cfg, conf.UseEnv())

	require.NoError(t, err)
	require.Equal(t, 1200, cfg.Topology.FollowerUsers)
	require.Equal(t, "hot", cfg.ReaderCardinality)
	require.Equal(t, 1100, cfg.Topology.BigVFollowerCount)
	require.Equal(t, 1, cfg.Setup.WarmupConcurrency)
	require.Equal(t, 4, cfg.Setup.CursorSeedConcurrency)
	require.Equal(t, 5, cfg.Setup.StrategySeedPosts)
	require.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128, 256}, cfg.Load.ConcurrencySteps)
	for name, client := range map[string]struct {
		hosts []string
		key   string
		want  string
	}{
		"knowpost": {cfg.KnowPostRpc.Etcd.Hosts, cfg.KnowPostRpc.Etcd.Key, "knowpost.rpc"},
		"relation": {cfg.RelationRpc.Etcd.Hosts, cfg.RelationRpc.Etcd.Key, "relation.rpc"},
		"counter":  {cfg.CounterRpc.Etcd.Hosts, cfg.CounterRpc.Etcd.Key, "counter.rpc"},
		"user":     {cfg.UserRpc.Etcd.Hosts, cfg.UserRpc.Etcd.Key, "user.rpc"},
		"auth":     {cfg.AuthRpc.Etcd.Hosts, cfg.AuthRpc.Etcd.Key, "user.rpc"},
	} {
		require.Equal(t, []string{"127.0.0.1:12379"}, client.hosts, name)
		require.Equal(t, client.want, client.key, name)
	}
	require.Equal(t, "feed-fanout-group", cfg.Monitor.KafkaGroup)
	require.Equal(t, "zg-knowpost", cfg.Monitor.PrometheusContainer)
	require.Equal(t, "http://127.0.0.1:9104/metrics", cfg.Monitor.PrometheusURL)
	require.Equal(t, []string{"zg-gateway", "zg-knowpost", "zg-relation", "zg-counter"}, cfg.Monitor.DockerContainers)
	require.Equal(t, "compose-full", cfg.Execution.Topology)
	require.Equal(t, ".tmp/feed-local/state.json", cfg.Execution.LocalStateFile)
	require.Equal(t, "treatment", cfg.Execution.FeaturePreset)
	require.Equal(t, "natural", cfg.Execution.CacheState)
	require.Equal(t, 64, cfg.Execution.CacheWarmConcurrency)
	require.Equal(t, 90.0, cfg.Execution.ClientCPULimitPercent)
}

func TestEffectiveReportRunIDPrefersExplicitResultNamespace(t *testing.T) {
	require.Equal(t, "phase0-strict", effectiveReportRunID("dataset-run", "phase0-strict"))
	require.Equal(t, "dataset-run", effectiveReportRunID("dataset-run", ""))
}
