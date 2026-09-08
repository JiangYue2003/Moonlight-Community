param(
  [switch]$WithDocker,
  [switch]$RunMigrate,
  [string]$DependencyProfile = "",
  [switch]$ValidateOnly
)

$ErrorActionPreference = "Stop"

$Root = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
$TopologyFile = Join-Path $Root "deploy/topology/services.json"
$LogsDir = Join-Path $Root "var/log/dev"
$RunDir = Join-Path $Root "var/run/dev"
$BinDir = Join-Path $Root "var/bin/dev"
$CacheDir = Join-Path $Root "var/cache/go-build"
$PidFile = Join-Path $RunDir "pids.json"
$LegacyPidFile = Join-Path $Root "logs/dev/pids.json"

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

function Get-PortOwner([int]$Port) {
  $line = netstat -ano | Select-String ":$Port\s+.*LISTENING\s+(\d+)$" | Select-Object -First 1
  if ($null -eq $line) {
    return 0
  }
  $text = $line.ToString().Trim()
  if ($text -match "(\d+)$") {
    return [int]$Matches[1]
  }
  return 0
}

function Assert-PortsAvailable([int[]]$Ports) {
  foreach ($port in $Ports) {
    $owner = Get-PortOwner -Port $port
    if ($owner -ne 0) {
      throw "service port $port is already owned by PID $owner; stop the recorded stack or choose another port"
    }
  }
}

function Wait-ServicePorts([string]$ServiceId, [System.Diagnostics.Process]$Process, [int[]]$Ports, [int]$TimeoutSeconds = 60) {
  $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
  while ((Get-Date) -lt $deadline) {
    $Process.Refresh()
    if ($Process.HasExited) {
      throw "$ServiceId exited before readiness with code $($Process.ExitCode)"
    }
    $ready = @($Ports | Where-Object { (Get-PortOwner -Port $_) -eq $Process.Id })
    if ($ready.Count -eq $Ports.Count) {
      return
    }
    Start-Sleep -Milliseconds 250
  }
  $owners = @($Ports | ForEach-Object { "$_=$(Get-PortOwner -Port $_)" }) -join ", "
  throw "$ServiceId did not own all expected ports within $TimeoutSeconds seconds; owners: $owners"
}

function Start-GoService($Service) {
  $executable = Join-Path $BinDir ($Service.id + ".exe")
  & go build -buildvcs=false -trimpath -o $executable ("./" + $Service.run.package)
  if ($LASTEXITCODE -ne 0) {
    throw "go build failed for $($Service.id)"
  }

  $stdoutPath = Join-Path $LogsDir $Service.logFile
  $stderrPath = Join-Path $LogsDir ([System.IO.Path]::GetFileNameWithoutExtension($Service.logFile) + ".stderr.log")
  $proc = Start-Process `
    -FilePath $executable `
    -ArgumentList @($Service.run.args) `
    -WorkingDirectory $Root `
    -WindowStyle Hidden `
    -RedirectStandardOutput $stdoutPath `
    -RedirectStandardError $stderrPath `
    -PassThru

  return [PSCustomObject]@{
    service = $Service.id
    pid = $proc.Id
    executable = $executable
    processStartUnixMs = ([DateTimeOffset]$proc.StartTime.ToUniversalTime()).ToUnixTimeMilliseconds()
    stdout = $stdoutPath
    stderr = $stderrPath
    ports = @($Service.ports | Where-Object { $_.managedOnStart } | ForEach-Object { $_.number })
    startedAt = (Get-Date).ToString("s")
  }
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
  throw "go command is not found in PATH"
}

foreach ($dir in @($LogsDir, $RunDir, $BinDir, $CacheDir)) {
  if (-not (Test-Path $dir)) {
    New-Item -ItemType Directory -Path $dir | Out-Null
  }
}
$env:GOCACHE = $CacheDir

if (-not (Test-Path -LiteralPath $TopologyFile)) {
  throw "Topology manifest not found: $TopologyFile"
}
$topology = Get-Content -LiteralPath $TopologyFile -Raw | ConvertFrom-Json
Push-Location $Root
try {
  & go run ./deploy/topology/cmd/validate -manifest $TopologyFile -repo-root $Root
  if ($LASTEXITCODE -ne 0) {
    throw "Topology validation failed"
  }
} finally {
  Pop-Location
}
$coreServices = @($topology.services | Where-Object { $_.defaultLocal })
if ($coreServices.Count -eq 0) {
  throw "Topology has no default local services"
}

if ([string]::IsNullOrWhiteSpace($DependencyProfile)) {
  $DependencyProfile = if ($WithDocker) {
    [string]$topology.dockerDependencyProfile
  } else {
    [string]$topology.defaultDependencyProfile
  }
}
$profile = @($topology.dependencyProfiles | Where-Object { $_.id -eq $DependencyProfile })
if ($profile.Count -ne 1) {
  $availableProfiles = @($topology.dependencyProfiles.id) -join ", "
  throw "Unknown dependency profile '$DependencyProfile'. Available profiles: $availableProfiles"
}
$profile = $profile[0]
foreach ($service in $coreServices) {
  if ($DependencyProfile -notin @($service.dependencyProfiles)) {
    throw "Service '$($service.id)' does not support dependency profile '$DependencyProfile'"
  }
}

Write-Info "Topology manifest: $TopologyFile"
Write-Info "Dependency profile: $DependencyProfile"
Write-Info ("Default local services: {0}" -f (@($coreServices.id) -join ", "))
if ($ValidateOnly) {
  Write-Info "Topology and startup plan are valid; no services were started."
  return
}

if ($WithDocker) {
  if ([string]::IsNullOrWhiteSpace([string]$profile.composeFile) -or @($profile.composeServices).Count -eq 0) {
    throw "Dependency profile '$DependencyProfile' does not define Compose dependencies"
  }
  $composeFile = Join-Path $Root ([string]$profile.composeFile)
  $dependencies = @($profile.composeServices)
  Write-Info "Starting docker dependencies: $($dependencies -join ', ')"
  docker compose -f $composeFile up -d --wait --wait-timeout 180 @dependencies
  if ($LASTEXITCODE -ne 0) {
    throw "Docker dependencies failed to start"
  }
}

foreach ($endpoint in @($profile.requiredEndpoints)) {
  if (-not (Test-TcpPort $endpoint.host $endpoint.port)) {
    throw "$($endpoint.name) is not reachable at $($endpoint.host):$($endpoint.port)"
  }
}

if ($RunMigrate) {
  Write-Info "Running DB migrations..."
  & (Join-Path $Root "scripts/migrate.bat") up
}

if ((Test-Path $PidFile) -or (Test-Path $LegacyPidFile)) {
  Write-WarnMsg "Existing PID file found, trying to stop previous processes..."
  & (Join-Path $Root "scripts/stop-all.ps1")
}

$managedPorts = @($coreServices.ports | Where-Object { $_.managedOnStart } | ForEach-Object { $_.number } | Sort-Object -Unique)
Assert-PortsAvailable -Ports $managedPorts

if (-not (Test-Path (Join-Path $Root "certs/jwt_private.pem")) -or -not (Test-Path (Join-Path $Root "certs/jwt_public.pem"))) {
  Write-WarnMsg "JWT key files are missing under certs/"
}

Write-Info "Starting services in background..."
$started = @()
try {
  foreach ($s in $coreServices) {
    $servicePorts = @($s.ports | Where-Object { $_.managedOnStart } | ForEach-Object { $_.number })
    Assert-PortsAvailable -Ports $servicePorts
    $item = Start-GoService -Service $s
    $started += $item
    $started | ConvertTo-Json -Depth 4 | Set-Content -LiteralPath $PidFile -Encoding UTF8
    $process = Get-Process -Id $item.pid -ErrorAction Stop
    Wait-ServicePorts -ServiceId $s.id -Process $process -Ports $servicePorts
    Write-Info ("Started {0} (PID={1})" -f $item.service, $item.pid)
  }
} catch {
  $startupError = $_
  if (Test-Path -LiteralPath $PidFile) {
    try {
      & (Join-Path $Root "scripts/stop-all.ps1") -PidFile $PidFile
    } catch {
      Write-WarnMsg "Startup cleanup failed: $_"
    }
  }
  throw $startupError
}

Write-Host ""
Write-Info "All start commands have been dispatched."
Write-Info "PID file: $PidFile"
Write-Info "Logs dir: $LogsDir"
Write-Host "Tail sample: Get-Content -Wait `"$LogsDir\gateway.log`""
