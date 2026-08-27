[CmdletBinding(SupportsShouldProcess = $true)]
param(
    [string]$RepoRoot = "",
    [string]$StateFile = ""
)

$ErrorActionPreference = "Stop"

if ([string]::IsNullOrWhiteSpace($RepoRoot)) {
    $RepoRoot = Join-Path $PSScriptRoot ".."
}
$RepoRoot = (Resolve-Path -LiteralPath $RepoRoot).Path
if ([string]::IsNullOrWhiteSpace($StateFile)) {
    $StateFile = Join-Path $RepoRoot ".tmp\feed-local\state.json"
}
$StateFile = [System.IO.Path]::GetFullPath($StateFile)

Write-Host "Feed local-service stop plan"
Write-Host "  state: $StateFile"
Write-Host "  scope: only exact PIDs and executable paths recorded by start-feed-local-services.ps1"
Write-Host "  middleware: retained"

if ($WhatIfPreference) {
    Write-Host "WhatIf: no process or file changes were made."
    return
}
if (-not (Test-Path -LiteralPath $StateFile)) {
    throw "state file does not exist: $StateFile"
}

$state = Get-Content -LiteralPath $StateFile -Raw | ConvertFrom-Json
if ($state.topology -ne "compose-middleware-local-services") {
    throw "refusing unknown state topology: $($state.topology)"
}

$records = @($state.services)
[array]::Reverse($records)
foreach ($record in $records) {
    $processId = [int]$record.pid
    $expectedExecutable = [System.IO.Path]::GetFullPath([string]$record.executable)
	$expectedStartUnixMs = [int64]$record.process_start_unix_ms
	if ($expectedStartUnixMs -le 0) {
		throw "state record for $($record.name) has no valid process start time; no process was stopped"
	}
    $process = Get-Process -Id $processId -ErrorAction SilentlyContinue
    if ($null -eq $process) {
        Write-Host "$($record.name): PID $processId already exited"
        continue
    }
    $actualExecutable = ""
    try {
        $actualExecutable = [System.IO.Path]::GetFullPath($process.Path)
    }
    catch {
        throw "cannot verify executable path for PID $processId; no process was stopped"
    }
    if (-not $actualExecutable.Equals($expectedExecutable, [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "PID $processId now belongs to $actualExecutable, expected $expectedExecutable; no mismatched process was stopped"
    }
	$actualStartUnixMs = ([DateTimeOffset]$process.StartTime.ToUniversalTime()).ToUnixTimeMilliseconds()
	if ($actualStartUnixMs -ne $expectedStartUnixMs) {
		throw "PID $processId start time is $actualStartUnixMs, expected $expectedStartUnixMs; no reused process was stopped"
	}
    Stop-Process -Id $processId
    Write-Host "$($record.name): stopped PID $processId"
}

Remove-Item -LiteralPath $StateFile
Write-Host "Recorded local Feed services stopped; middleware was left running."
