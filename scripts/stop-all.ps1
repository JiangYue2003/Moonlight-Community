$ErrorActionPreference = "Stop"

$Root = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
$PidFile = Join-Path $Root "logs/dev/pids.json"

function Write-Info([string]$msg) {
  Write-Host "[INFO] $msg" -ForegroundColor Cyan
}

function Write-WarnMsg([string]$msg) {
  Write-Host "[WARN] $msg" -ForegroundColor Yellow
}

if (-not (Test-Path $PidFile)) {
  Write-WarnMsg "PID file not found: $PidFile"
  exit 0
}

$items = Get-Content $PidFile -Raw | ConvertFrom-Json
if ($null -eq $items) {
  Write-WarnMsg "PID file is empty, skip"
  exit 0
}

foreach ($it in @($items)) {
  try {
    $p = Get-Process -Id $it.pid -ErrorAction Stop
    Stop-Process -Id $p.Id -Force -ErrorAction Stop
    Write-Info ("Stopped {0} (PID={1})" -f $it.service, $it.pid)
  } catch {
    Write-WarnMsg ("Already exited or not found: {0} (PID={1})" -f $it.service, $it.pid)
  }
}

if (Test-Path $PidFile) {
  Remove-Item $PidFile -Force
  Write-Info "PID file removed."
}
