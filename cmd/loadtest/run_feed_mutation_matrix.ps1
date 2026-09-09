[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [string]$RunId,

    [string]$ReportRunId = "",
    [string]$ManifestPath = "",
    [string]$CheckpointPath = "",
    [ValidateSet("push", "pull", "hybrid")]
    [string]$Strategy = "hybrid",
    [ValidateSet("publish", "mixed-90-10", "mixed-80-20")]
    [string[]]$Scenarios = @("publish", "mixed-90-10", "mixed-80-20"),
    [int[]]$Concurrencies = @(32, 64, 128, 256),
    [int[]]$GatewayConcurrencies = @(32, 64, 128, 256),
    [int]$DurationSeconds = 60,
    [int]$WarmupSeconds = 10,
    [int]$Trials = 3,
    [ValidateSet("hot", "distributed", "high")]
    [string]$ReaderCardinality = "hot",
    [ValidateSet("natural", "cold", "l1-warm", "l2-warm", "expire-together")]
    [string]$CacheState = "cold",
    [ValidateSet("compose-full", "compose-middleware-local-services")]
    [string]$Topology = "compose-full",
    [ValidateSet("control", "treatment")]
    [string]$FeaturePreset = "treatment",
    [ValidateSet("login", "signed")]
    [string]$GatewayAuthMode = "signed",
    [switch]$CreateCheckpoint,
    [switch]$SkipGateway,
    [long]$MaxBenchmarkLogBytes = 67108864
)

$ErrorActionPreference = "Stop"
$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path
$configFile = Join-Path $repoRoot "cmd\loadtest\load_test.yaml"
if ([string]::IsNullOrWhiteSpace($ReportRunId)) {
    $ReportRunId = "$RunId-$FeaturePreset-mutation-formal"
}
if ([string]::IsNullOrWhiteSpace($ManifestPath)) {
    $ManifestPath = Join-Path $repoRoot "results\feed-loadtest\manifests\$RunId.json"
}
elseif (-not [System.IO.Path]::IsPathRooted($ManifestPath)) {
    $ManifestPath = Join-Path $repoRoot $ManifestPath
}
$ManifestPath = [System.IO.Path]::GetFullPath($ManifestPath)
if ([string]::IsNullOrWhiteSpace($CheckpointPath)) {
    $CheckpointPath = Join-Path $repoRoot "results\feed-loadtest\checkpoints\$RunId.json"
}
elseif (-not [System.IO.Path]::IsPathRooted($CheckpointPath)) {
    $CheckpointPath = Join-Path $repoRoot $CheckpointPath
}
$CheckpointPath = [System.IO.Path]::GetFullPath($CheckpointPath)

if (-not (Test-Path -LiteralPath $ManifestPath)) {
    throw "manifest does not exist: $ManifestPath"
}
if ($DurationSeconds -le 0 -or $WarmupSeconds -le 0 -or $Trials -le 0) {
    throw "formal mutation matrix requires positive duration, warmup, and trials"
}
if ($Concurrencies.Count -eq 0 -or ((-not $SkipGateway) -and $GatewayConcurrencies.Count -eq 0)) {
    throw "formal mutation matrix requires at least one selected concurrency per enabled entry"
}
foreach ($concurrency in @($Concurrencies + $GatewayConcurrencies)) {
    if ($concurrency -le 0) {
        throw "concurrency must be positive: $concurrency"
    }
}
if ($FeaturePreset -eq "control" -and $CacheState -ne "cold" -and $CacheState -ne "natural") {
    throw "control formal mutation matrix cannot claim a warm Page Cache state"
}
if ($ReportRunId -notmatch '^[A-Za-z0-9._-]+$') {
    throw "ReportRunId must contain only letters, digits, dot, underscore, or hyphen"
}

function Assert-UniqueValues {
    param(
        [Parameter(Mandatory = $true)]
        [object[]]$Values,
        [Parameter(Mandatory = $true)]
        [string]$Name
    )
    if ($Values.Count -eq 0) {
        throw "$Name must not be empty"
    }
    $seen = @{}
    foreach ($value in $Values) {
        $key = ([string]$value).ToLowerInvariant()
        if ($seen.ContainsKey($key)) {
            throw "$Name contains duplicate value: $value"
        }
        $seen[$key] = $true
    }
}

Assert-UniqueValues -Values @($Scenarios) -Name "Scenarios"
Assert-UniqueValues -Values @($Concurrencies) -Name "Concurrencies"
if (-not $SkipGateway) {
    Assert-UniqueValues -Values @($GatewayConcurrencies) -Name "GatewayConcurrencies"
}

$reportRoot = Join-Path $repoRoot "results\feed-loadtest\$ReportRunId"

function Assert-FreshReportNamespace {
    if (-not (Test-Path -LiteralPath $reportRoot)) {
        return
    }
    $existing = Get-ChildItem -LiteralPath $reportRoot -Force | Select-Object -First 1
    if ($null -ne $existing) {
        throw "formal mutation report namespace is not empty; choose a new ReportRunId: $reportRoot"
    }
}

function Get-ExpectedMatrixCells {
    $cells = @()
    foreach ($scenario in $Scenarios) {
        foreach ($concurrency in $Concurrencies) {
            $cells += [pscustomobject]@{ entry = "rpc"; scenario = $scenario; concurrency = $concurrency }
        }
        if (-not $SkipGateway) {
            foreach ($concurrency in $GatewayConcurrencies) {
                $cells += [pscustomobject]@{ entry = "gateway"; scenario = $scenario; concurrency = $concurrency }
            }
        }
    }
    return @($cells)
}

function Write-ExpectedMatrixManifest {
    New-Item -ItemType Directory -Path $reportRoot -Force | Out-Null
    $body = [ordered]@{
        version = 1
        generated_at = (Get-Date).ToUniversalTime().ToString("o")
        dataset_run_id = $RunId
        report_run_id = $ReportRunId
        checkpoint_path = $CheckpointPath
        strategy = $Strategy
        feature_preset = $FeaturePreset
        reader_cardinality = $ReaderCardinality
        cache_state = $CacheState
        topology = $Topology
        warmup_seconds = $WarmupSeconds
        duration_seconds = $DurationSeconds
        trials = $Trials
        cells = @(Get-ExpectedMatrixCells)
    }
    $body | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $reportRoot "expected-matrix.json") -Encoding utf8NoBOM
}

function Assert-ExpectedMutationMatrix {
    $cells = @(Get-ExpectedMatrixCells)
    foreach ($cell in $cells) {
        $scenarioName = "$Strategy-$($cell.entry)-$($cell.scenario)-c$($cell.concurrency)"
        for ($trial = 1; $trial -le $Trials; $trial++) {
            $reportPath = Join-Path $reportRoot "$scenarioName\trial-$trial\report.json"
            if (-not (Test-Path -LiteralPath $reportPath)) {
                throw "formal mutation matrix is missing expected report: $reportPath"
            }
            $report = Get-Content -LiteralPath $reportPath -Raw | ConvertFrom-Json
            if (-not [bool]$report.complete) {
                throw "formal mutation report is incomplete: $reportPath missing=$($report.missing_metrics -join ',')"
            }
            if ($report.run_id -ne $ReportRunId -or $report.strategy -ne $Strategy -or
                $report.entry -ne $cell.entry -or $report.scenario -ne "$($cell.scenario)-c$($cell.concurrency)" -or
                [int]$report.concurrency -ne [int]$cell.concurrency -or [int]$report.trial -ne $trial -or
                [int]$report.expected_trials -ne $Trials) {
                throw "formal mutation report identity does not match expected matrix: $reportPath"
            }
            if ($null -eq $report.mutation -or -not [bool]$report.mutation.after_trial_restore.complete) {
                throw "formal mutation report lacks a complete after-trial restore: $reportPath"
            }
        }
    }
    $actualReports = @(Get-ChildItem -LiteralPath $reportRoot -Recurse -Filter "report.json" -File)
    $expectedReports = $cells.Count * $Trials
    if ($actualReports.Count -ne $expectedReports) {
        throw "formal mutation report count mismatch: actual=$($actualReports.Count) expected=$expectedReports"
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

function Assert-BenchmarkLogsBounded {
    if ($Topology -ne "compose-middleware-local-services") {
        return
    }
    $stateFile = Join-Path $repoRoot ".tmp\feed-local\state.json"
    if (-not (Test-Path -LiteralPath $stateFile)) {
        throw "local topology state file is missing: $stateFile"
    }
    $state = Get-Content -LiteralPath $stateFile -Raw | ConvertFrom-Json
    if ($state.strategy -ne $Strategy) {
        throw "local strategy is $($state.strategy), requested $Strategy"
    }
    if ($state.feature_preset -ne $FeaturePreset) {
        throw "local feature preset is $($state.feature_preset), requested $FeaturePreset"
    }
    if ($state.logging.preset -ne "benchmark-error-only") {
        throw "formal mutation matrix requires logging preset benchmark-error-only"
    }
    foreach ($service in @($state.services)) {
        foreach ($path in @([string]$service.stdout, [string]$service.stderr)) {
            if ([string]::IsNullOrWhiteSpace($path) -or -not (Test-Path -LiteralPath $path)) {
                continue
            }
            $length = (Get-Item -LiteralPath $path).Length
            if ($length -gt $MaxBenchmarkLogBytes) {
                throw "benchmark log exceeded bound: service=$($service.name) path=$path bytes=$length limit=$MaxBenchmarkLogBytes"
            }
        }
    }
}

function Invoke-Loadtest {
    param([string[]]$Arguments)
    Assert-BenchmarkLogsBounded
    $base = @(
        "run", "-buildvcs=false", "./cmd/loadtest", "-f", $configFile,
        "-run-id", $RunId, "-report-run-id", $ReportRunId,
        "-manifest", $ManifestPath, "-checkpoint", $CheckpointPath,
        "-confirm-mutation", $RunId,
        "-reader-cardinality", $ReaderCardinality,
        "-strategy", $Strategy, "-topology", $Topology,
        "-cache-state", $CacheState, "-feature-preset", $FeaturePreset,
        "-gateway-auth-mode", $GatewayAuthMode
    )
    try {
        Invoke-Checked "go" ($base + $Arguments)
    }
    finally {
        Assert-BenchmarkLogsBounded
    }
}

function Invoke-MutationEntry {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Entry,
        [Parameter(Mandatory = $true)]
        [int[]]$Steps
    )
    foreach ($scenario in $Scenarios) {
        foreach ($concurrency in $Steps) {
            Write-Host "=== formal mutation: entry=$Entry scenario=$scenario concurrency=$concurrency trials=$Trials ==="
            Invoke-Loadtest @(
                "-phase", "run", "-entry", $Entry,
                "-scenario", $scenario, "-concurrency", "$concurrency",
                "-requests", "0", "-duration", "${DurationSeconds}s",
                "-warmup", "${WarmupSeconds}s", "-trials", "$Trials"
            )
        }
    }
}

Push-Location $repoRoot
try {
    Assert-FreshReportNamespace
    if ($CreateCheckpoint) {
        if (Test-Path -LiteralPath $CheckpointPath) {
            throw "checkpoint already exists; refusing to overwrite: $CheckpointPath"
        }
        Invoke-Loadtest @("-phase", "create-mutation-checkpoint")
    }
    elseif (-not (Test-Path -LiteralPath $CheckpointPath)) {
        throw "checkpoint does not exist; run with -CreateCheckpoint after confirming the baseline: $CheckpointPath"
    }

    Invoke-Loadtest @("-phase", "verify-mutation-checkpoint")
    Write-ExpectedMatrixManifest
    Invoke-MutationEntry -Entry "rpc" -Steps $Concurrencies
    if (-not $SkipGateway) {
        Invoke-MutationEntry -Entry "gateway" -Steps $GatewayConcurrencies
    }
    Invoke-Loadtest @("-phase", "verify-mutation-checkpoint")
    Assert-ExpectedMutationMatrix
    Invoke-Loadtest @("-phase", "compare")
}
finally {
    Pop-Location
}
