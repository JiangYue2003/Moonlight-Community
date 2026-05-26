param(
  [switch]$WithDocker,
  [switch]$RunMigrate,
  [switch]$IncludeMaintenance
)

$ErrorActionPreference = "Stop"

$Root = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
$LogsDir = Join-Path $Root "logs/dev"
$PidFile = Join-Path $LogsDir "pids.json"

function Write-Info([string]$msg) {
  Write-Host "[INFO] $msg" -ForegroundColor Cyan
}

function Write-WarnMsg([string]$msg) {
  Write-Host "[WARN] $msg" -ForegroundColor Yellow
}

function Test-TcpPort([string]$HostName, [int]$Port, [int]$TimeoutMs = 1200) {
  $client = New-Object System.Net.Sockets.TcpClient
  try {
    $ar = $client.BeginConnect($HostName, $Port, $null, $null)
    if (-not $ar.AsyncWaitHandle.WaitOne($TimeoutMs, $false)) {
      return $false
    }
    $client.EndConnect($ar)
    return $true
  } catch {
    return $false
  } finally {
    $client.Close()
  }
}

function Stop-ProcessOnPort([int]$Port) {
  $lines = netstat -ano | Select-String ":$Port\s+.*LISTENING\s+(\d+)$"
  foreach ($line in $lines) {
    $text = $line.ToString().Trim()
    if ($text -match "(\d+)$") {
      $procId = [int]$Matches[1]
      if ($procId -gt 0 -and $procId -ne $PID) {
        try {
          Stop-Process -Id $procId -Force -ErrorAction Stop
          Write-WarnMsg "Stopped process on port $Port (PID=$procId)"
        } catch {
          Write-WarnMsg "Failed to stop PID=$procId on port $Port"
        }
      }
    }
  }
}

function Start-GoService([string]$ServicePath, [string]$LogName, [string]$ExtraArgs = "") {
  $serviceEntry = Join-Path $Root $ServicePath
  if (-not (Test-Path $serviceEntry)) {
    throw "Service path not found: $serviceEntry"
  }

  $logPath = Join-Path $LogsDir $LogName
  $runCmd = "go run ./$ServicePath"
  if (-not [string]::IsNullOrWhiteSpace($ExtraArgs)) {
    $runCmd = "$runCmd $ExtraArgs"
  }
  $cmd = "Set-Location `\"$Root`\"; $runCmd *>> `\"$logPath`\""

  $proc = Start-Process `
    -FilePath "powershell" `
    -ArgumentList @("-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", $cmd) `
    -WindowStyle Hidden `
    -PassThru

  return [PSCustomObject]@{
    service = $ServicePath
    pid = $proc.Id
    log = $logPath
    startedAt = (Get-Date).ToString("s")
  }
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
  throw "go command is not found in PATH"
}

if ($WithDocker) {
  Write-Info "Starting docker dependencies..."
  docker compose -f (Join-Path $Root "deploy/compose/docker-compose.dev.yml") up -d
}

if (-not (Test-TcpPort "127.0.0.1" 3306)) {
  throw "MySQL is not reachable at 127.0.0.1:3306"
}
if (-not (Test-TcpPort "127.0.0.1" 6379)) {
  throw "Redis is not reachable at 127.0.0.1:6379"
}

if ($RunMigrate) {
  Write-Info "Running DB migrations..."
  & (Join-Path $Root "scripts/migrate.bat") up
}

if (-not (Test-Path $LogsDir)) {
  New-Item -ItemType Directory -Path $LogsDir | Out-Null
}

foreach ($p in @(8001,8002,8003,8004,8005,8006,8007,8008,8080,8011,9002,9003,9004,9006,9013,9017,9018,9102,9103,9104,9105,9111,9112,9113,9114,9115,9116,9117,9118,9119)) {
  Stop-ProcessOnPort -Port $p
}

if (Test-Path $PidFile) {
  Write-WarnMsg "Existing PID file found, trying to stop previous processes..."
  try {
    & (Join-Path $Root "scripts/stop-all.ps1")
  } catch {
    Write-WarnMsg "stop-all.ps1 failed, continue and overwrite PID file"
  }
}

if (-not (Test-Path (Join-Path $Root "certs/jwt_private.pem")) -or -not (Test-Path (Join-Path $Root "certs/jwt_public.pem"))) {
  Write-WarnMsg "JWT key files are missing under certs/"
}

if (-not $env:DEEPSEEK_API_KEY) {
  Write-WarnMsg "DEEPSEEK_API_KEY is not set"
}
if (-not $env:DASHSCOPE_API_KEY) {
  Write-WarnMsg "DASHSCOPE_API_KEY is not set"
}

$coreServices = @(
  @{ path = "services/gateway"; log = "gateway.log"; args = "-f services/gateway/etc/gateway.yaml" },
  @{ path = "services/user/rpc"; log = "user-rpc.log"; args = "-f services/user/rpc/etc/user.yaml" },
  @{ path = "services/storage/rpc"; log = "storage-rpc.log"; args = "-f services/storage/rpc/etc/storage.yaml" },
  @{ path = "services/search/rpc"; log = "search-rpc.log"; args = "-f services/search/rpc/etc/search.yaml" },
  @{ path = "services/llm/rpc"; log = "llm-rpc.log"; args = "-f services/llm/rpc/etc/llm.yaml" },
  @{ path = "services/counter/cmd/counter"; log = "counter-service.log"; args = "-f services/counter/cmd/counter/etc/counter.yaml" },
  @{ path = "services/knowpost/cmd/knowpost"; log = "knowpost-service.log"; args = "-f services/knowpost/cmd/knowpost/etc/knowpost.yaml" },
  @{ path = "services/relation/cmd/relation"; log = "relation-service.log"; args = "-f services/relation/cmd/relation/etc/relation.yaml" },
  @{ path = "services/search/cmd/search"; log = "search-service.log"; args = "-f services/search/cmd/search/etc/search.yaml" },
  @{ path = "services/llm/cmd/llm"; log = "llm-service.log"; args = "-f services/llm/cmd/llm/etc/llm.yaml" },
  @{ path = "services/agent/cmd/agent"; log = "agent-service.log"; args = "-f services/agent/cmd/agent/etc/agent.yaml" }
)

$maintenanceServices = @(
  @{ path = "services/outbox/gc"; log = "outbox-gc.log"; args = "services/outbox/gc/etc/gc.yaml" },
  @{ path = "services/counter/reconciler"; log = "counter-reconciler.log"; args = "services/counter/reconciler/etc/reconciler.yaml" }
)

$targets = @()
$targets += $coreServices
if ($IncludeMaintenance) {
  $targets += $maintenanceServices
}

Write-Info "Starting services in background..."
$started = @()
foreach ($s in $targets) {
  $item = Start-GoService -ServicePath $s.path -LogName $s.log -ExtraArgs $s.args
  $started += $item
  Write-Info ("Started {0} (PID={1})" -f $item.service, $item.pid)
  Start-Sleep -Milliseconds 250
}

$started | ConvertTo-Json -Depth 3 | Set-Content -Path $PidFile -Encoding UTF8

Write-Host ""
Write-Info "All start commands have been dispatched."
Write-Info "PID file: $PidFile"
Write-Info "Logs dir: $LogsDir"
Write-Host "Tail sample: Get-Content -Wait `\"$LogsDir\gateway.log`\""
