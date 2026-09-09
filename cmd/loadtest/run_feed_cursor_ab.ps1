[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [string]$RunId,

    [string]$ReportRunId = "",

    [ValidateSet("scout", "formal")]
    [string]$Mode = "scout",

    [ValidateSet("compose-full", "compose-middleware-local-services")]
    [string]$Topology = "compose-middleware-local-services",

    [ValidateSet("rpc", "gateway")]
    [string[]]$Entries = @("rpc", "gateway"),

    [ValidateSet("login", "signed")]
    [string]$GatewayAuthMode = "signed",

    [ValidateSet("page1", "page5", "cursor-page5", "page20", "cursor-page20", "page50", "cursor-page50", "sequential-page-50", "sequential-50", "same-second-cursor")]
    [string[]]$Scenarios = @("page1", "page5", "cursor-page5", "page20", "cursor-page20", "page50", "cursor-page50", "sequential-page-50", "sequential-50", "same-second-cursor"),

    [int[]]$Concurrencies = @(),
    [int]$DurationSeconds = 0,
    [int]$WarmupSeconds = -1,
    [int]$Trials = 0
)

$ErrorActionPreference = "Stop"
$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path
$configFile = Join-Path $repoRoot "cmd\loadtest\load_test.yaml"
$manifestFile = Join-Path $repoRoot "results\feed-loadtest\manifests\$RunId.json"
$binaryDir = Join-Path $repoRoot ".tmp\feed-cursor-ab"
$binary = Join-Path $binaryDir "feed-loadtest.exe"

if ([string]::IsNullOrWhiteSpace($ReportRunId)) {
    $ReportRunId = "$RunId-cursor-ab-$Mode"
}
if (-not (Test-Path -LiteralPath $manifestFile)) {
    throw "cursor-deep manifest is missing: $manifestFile"
}
if ($Concurrencies.Count -eq 0) {
    if ($Mode -eq "scout") {
        $Concurrencies = @(16, 32, 64)
    }
    else {
        throw "formal mode requires one stable concurrency and optionally one adjacent concurrency"
    }
}
if ($DurationSeconds -le 0) {
    $DurationSeconds = if ($Mode -eq "scout") { 10 } else { 60 }
}
if ($WarmupSeconds -lt 0) {
    $WarmupSeconds = if ($Mode -eq "scout") { 3 } else { 10 }
}
if ($Trials -le 0) {
    $Trials = if ($Mode -eq "scout") { 1 } else { 3 }
}
if ($Mode -eq "formal" -and $Concurrencies.Count -gt 2) {
    throw "formal mode accepts only the selected stable concurrency and one adjacent concurrency"
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

function Assert-CompleteReports {
    param(
        [string]$Entry,
        [string]$Scenario,
        [int]$Concurrency
    )
    for ($trial = 1; $trial -le $Trials; $trial++) {
        $reportPath = Join-Path $repoRoot "results\feed-loadtest\$ReportRunId\hybrid-$Entry-$Scenario-c$Concurrency\trial-$trial\report.json"
        if (-not (Test-Path -LiteralPath $reportPath)) {
            throw "cursor-ab report is missing: $reportPath"
        }
        $report = Get-Content -LiteralPath $reportPath -Raw | ConvertFrom-Json
        if (-not [bool]$report.complete) {
            $missing = @($report.missing_metrics) -join ","
            throw "cursor-ab report is incomplete; stop before a higher load: entry=$Entry scenario=$Scenario concurrency=$Concurrency trial=$trial missing=$missing"
        }
    }
}

New-Item -ItemType Directory -Force -Path $binaryDir | Out-Null
Push-Location $repoRoot
try {
    Invoke-Checked "go" @("build", "-buildvcs=false", "-o", $binary, "./cmd/loadtest")
    foreach ($entry in $Entries) {
        foreach ($scenario in $Scenarios) {
            foreach ($concurrency in $Concurrencies) {
                Write-Host "cursor-ab: mode=$Mode entry=$entry scenario=$scenario concurrency=$concurrency duration=${DurationSeconds}s trials=$Trials"
                Invoke-Checked $binary @(
                    "-f", $configFile,
                    "-phase", "run",
                    "-run-id", $RunId,
                    "-report-run-id", $ReportRunId,
                    "-manifest", $manifestFile,
                    "-reader-cardinality", "cursor-deep",
                    "-topology", $Topology,
                    "-cache-state", "cold",
                    "-feature-preset", "cursor-ab",
                    "-strategy", "hybrid",
                    "-entry", $entry,
                    "-gateway-auth-mode", $GatewayAuthMode,
                    "-scenario", $scenario,
                    "-concurrency", "$concurrency",
                    "-requests", "0",
                    "-duration", "${DurationSeconds}s",
                    "-warmup", "${WarmupSeconds}s",
                    "-trials", "$Trials"
                )
                Assert-CompleteReports -Entry $entry -Scenario $scenario -Concurrency $concurrency
            }
        }
    }
    Invoke-Checked $binary @(
        "-f", $configFile,
        "-phase", "compare",
        "-run-id", $RunId,
        "-report-run-id", $ReportRunId,
        "-reader-cardinality", "cursor-deep",
        "-topology", $Topology,
        "-cache-state", "cold",
        "-feature-preset", "cursor-ab"
    )
}
finally {
    Pop-Location
}
