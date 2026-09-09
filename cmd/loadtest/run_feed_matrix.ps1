[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [string]$RunId,

    [string]$ReportRunId = "",

    [string[]]$Strategies = @("hybrid"),
    [int[]]$Concurrencies = @(32, 64, 128, 256),
    [int[]]$GatewayConcurrencies = @(32, 64, 128, 256),
    [int]$ReadDurationSeconds = 60,
    [int]$WarmupSeconds = 10,
    [int]$Trials = 3,
    [ValidateSet("hot", "distributed", "high")]
    [string]$ReaderCardinality = "hot",
    [ValidateSet("auto", "natural", "cold", "l1-warm", "l2-warm", "expire-together")]
    [string]$CacheState = "auto",
    [ValidateSet("compose-full", "compose-middleware-local-services")]
    [string]$Topology = "compose-full",
    [ValidateSet("control", "treatment", "cursor-ab")]
    [string]$FeaturePreset = "treatment",
    [string]$PprofDir = "",
    [switch]$SkipGateway,
    [switch]$SkipSupplemental,
    [int]$SupplementalReadConcurrency = 128,
    [int]$BurstReadDurationSeconds = 10,
    [int]$SteadyReadDurationSeconds = 30,
    [long]$MaxBenchmarkLogBytes = 67108864
)

$ErrorActionPreference = "Stop"
$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path
$composeFile = Join-Path $repoRoot "deploy\compose\docker-compose.dev.yml"
$configFile = Join-Path $repoRoot "cmd\loadtest\load_test.yaml"
if ([string]::IsNullOrWhiteSpace($ReportRunId)) {
    $ReportRunId = $RunId
}
$ActiveReportRunId = $ReportRunId
if ($Strategies.Count -ne 1) {
	throw "one dataset may only be measured under one Feed strategy; create a fresh RunId per push, pull, or hybrid strategy"
}
$EffectiveCacheState = $CacheState
if ($CacheState -eq "auto") {
	if ($FeaturePreset -eq "control" -or $FeaturePreset -eq "cursor-ab") {
		$EffectiveCacheState = "cold"
	}
	else {
		$EffectiveCacheState = switch ($ReaderCardinality) {
			"hot" { "l1-warm" }
			"distributed" { "l2-warm" }
			"high" { "cold" }
		}
	}
}

function Invoke-Checked {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Command,
        [Parameter(Mandatory = $true)]
        [string[]]$Arguments
    )
    & $Command @Arguments
    if ($LASTEXITCODE -ne 0) {
        throw "$Command failed with exit code $LASTEXITCODE"
    }
}

function Invoke-Loadtest {
    param([string[]]$Arguments)
	Assert-BenchmarkLogsBounded
    $base = @(
        "run", "-buildvcs=false", "./cmd/loadtest", "-f", $configFile, "-run-id", $RunId,
        "-report-run-id", $ActiveReportRunId,
        "-reader-cardinality", $ReaderCardinality,
        "-topology", $Topology,
        "-cache-state", $EffectiveCacheState,
        "-feature-preset", $FeaturePreset
    )
    if (-not [string]::IsNullOrWhiteSpace($PprofDir)) {
        $base += @("-pprof-dir", $PprofDir)
    }
	try {
		Invoke-Checked "go" ($base + $Arguments)
	}
	finally {
		Assert-BenchmarkLogsBounded
	}
}

function Assert-BenchmarkLogsBounded {
	if ($Topology -ne "compose-middleware-local-services") {
		return
	}
	$stateFile = Join-Path $repoRoot ".tmp\feed-local\state.json"
	if (-not (Test-Path -LiteralPath $stateFile)) {
		throw "local topology state file is missing: $stateFile"
	}
	$state = Get-Content -LiteralPath $stateFile -Raw | ConvertFrom-Json
	if ($state.logging.preset -ne "benchmark-error-only") {
		throw "formal local matrix requires logging preset benchmark-error-only; got $($state.logging.preset)"
	}
	foreach ($service in @($state.services)) {
		foreach ($path in @([string]$service.stdout, [string]$service.stderr)) {
			if ([string]::IsNullOrWhiteSpace($path) -or -not (Test-Path -LiteralPath $path)) {
				continue
			}
			$length = (Get-Item -LiteralPath $path).Length
			if ($length -gt $MaxBenchmarkLogBytes) {
				throw "benchmark log exceeded the per-file bound: service=$($service.name) path=$path bytes=$length limit=$MaxBenchmarkLogBytes"
			}
		}
	}
}

function Set-FeedStrategy {
    param([string]$Strategy)
    if ($Topology -eq "compose-middleware-local-services") {
        $stateFile = Join-Path $repoRoot ".tmp\feed-local\state.json"
        if (-not (Test-Path -LiteralPath $stateFile)) {
            throw "local topology state file is missing: $stateFile"
        }
        $state = Get-Content -LiteralPath $stateFile -Raw | ConvertFrom-Json
        if ($state.strategy -ne $Strategy) {
            throw "local KnowPost strategy is $($state.strategy), requested $Strategy; restart local services with the requested strategy"
        }
		if ($state.feature_preset -ne $FeaturePreset) {
			throw "local KnowPost feature preset is $($state.feature_preset), requested $FeaturePreset; restart local services with the requested preset"
		}
		if ($state.logging.preset -ne "benchmark-error-only") {
			throw "local logging preset is $($state.logging.preset), formal matrix requires benchmark-error-only"
		}
        return
    }
    $env:FEED_STRATEGY = $Strategy
	$enabled = if ($FeaturePreset -eq "control") { "false" } else { "true" }
	$env:FEED_OBSERVABILITY = "true"
	$env:FEED_RELATION_EPOCH_ENABLED = $enabled
	$env:FEED_CONTENT_SAFETY_EPOCH_ENABLED = $enabled
	$env:FEED_ROUTE_SNAPSHOT_ENABLED = $enabled
	$env:FEED_COMBINED_PIPELINE_ENABLED = $enabled
	$env:FEED_CURSOR_PAGINATION_ENABLED = $enabled
	$env:FEED_PAGE_CACHE_MODE = if ($FeaturePreset -eq "treatment") { "l1-l2" } else { "off" }
    Invoke-Checked "docker" @(
        "compose", "-f", $composeFile,
        "up", "-d", "--no-deps", "--force-recreate",
        "--wait", "--wait-timeout", "60", "knowpost"
    )
    $configured = docker inspect zg-knowpost --format '{{range .Config.Env}}{{println .}}{{end}}' |
        Select-String "^FEED_STRATEGY=$Strategy$"
    if (-not $configured) {
        throw "KnowPost did not start with FEED_STRATEGY=$Strategy"
    }
}

function Get-KafkaLag {
    $raw = & docker exec zg-kafka /opt/bitnami/kafka/bin/kafka-consumer-groups.sh `
        --bootstrap-server kafka:29092 `
        --group feed-fanout-group `
        --describe 2>&1
    if ($LASTEXITCODE -ne 0) {
        throw "Kafka lag query failed: $raw"
    }
    $total = 0
    $partitions = 0
    foreach ($line in $raw) {
        $fields = ($line -split '\s+') | Where-Object { $_ }
        if ($fields.Count -ge 6 -and $fields[0] -eq "feed-fanout-group") {
            $lag = 0
            if ([int]::TryParse($fields[5], [ref]$lag)) {
                $total += $lag
                $partitions++
            }
        }
    }
    if ($partitions -eq 0) {
        throw "Kafka lag output contained no feed-fanout-group partition rows"
    }
    return $total
}

function Wait-KafkaDrain {
    param([int]$TimeoutSeconds = 180)
    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    do {
        $lag = Get-KafkaLag
        if ($lag -eq 0) {
            return
        }
        Start-Sleep -Seconds 1
    } while ((Get-Date) -lt $deadline)
    throw "Kafka feed-fanout-group lag did not drain within $TimeoutSeconds seconds; last lag=$lag"
}

function Invoke-ReadMatrix {
    param(
        [string]$Strategy,
        [string]$Entry,
        [int[]]$Steps,
        [string[]]$Scenarios
    )
    foreach ($scenario in $Scenarios) {
        foreach ($concurrency in $Steps) {
            Invoke-Loadtest @(
                "-phase", "run", "-strategy", $Strategy, "-entry", $Entry,
                "-scenario", $scenario, "-concurrency", "$concurrency",
                "-requests", "0", "-duration", "${ReadDurationSeconds}s",
                "-warmup", "${WarmupSeconds}s", "-trials", "$Trials"
            )
        }
    }
}

function Invoke-SupplementalScenarios {
    param([string]$Strategy)

    Write-Host "=== Audited burst/steady/recovery supplement: $Strategy ==="
    Set-FeedStrategy $Strategy
    Invoke-Loadtest @("-phase", "smoke", "-strategy", $Strategy)

    Invoke-Loadtest @(
        "-phase", "run", "-strategy", $Strategy, "-entry", "rpc",
        "-scenario", "burst-read", "-concurrency", "$SupplementalReadConcurrency",
        "-requests", "0", "-duration", "${BurstReadDurationSeconds}s",
        "-warmup", "${WarmupSeconds}s", "-trials", "$Trials"
    )
    Invoke-Loadtest @(
        "-phase", "run", "-strategy", $Strategy, "-entry", "rpc",
        "-scenario", "steady-read", "-concurrency", "$SupplementalReadConcurrency",
        "-requests", "0", "-duration", "${SteadyReadDurationSeconds}s",
        "-warmup", "${WarmupSeconds}s", "-trials", "$Trials"
    )
}

Push-Location $repoRoot
try {
    foreach ($strategy in $Strategies) {
        Write-Host "=== Feed strategy: $strategy ==="
        Set-FeedStrategy $strategy
        Wait-KafkaDrain
        Invoke-Loadtest @("-phase", "smoke", "-strategy", $strategy)

        Invoke-ReadMatrix $strategy "rpc" $Concurrencies @("hot-read", "distributed-read", "deep-page")

        if (-not $SkipGateway) {
            Invoke-ReadMatrix $strategy "gateway" $GatewayConcurrencies @("hot-read", "distributed-read")
        }

        if (-not $SkipSupplemental) {
            Invoke-SupplementalScenarios $strategy
        }

    }
}
finally {
    if ($Topology -eq "compose-full") {
        Write-Host "=== Restoring FEED_STRATEGY=hybrid ==="
        Set-FeedStrategy "hybrid"
    }
    Pop-Location
}
