[CmdletBinding(SupportsShouldProcess = $true)]
param(
    [string]$RepoRoot = "",
    [string]$ComposeFile = "",
    [string]$StateFile = "",
    [ValidateSet("push", "pull", "hybrid")]
    [string]$Strategy = "hybrid",
    [ValidateSet("control", "treatment", "cursor-ab")]
    [string]$FeaturePreset = "treatment",
    [ValidateSet("development", "benchmark-error-only")]
    [string]$LoggingPreset = "benchmark-error-only"
)

$ErrorActionPreference = "Stop"

if ([string]::IsNullOrWhiteSpace($RepoRoot)) {
    $RepoRoot = Join-Path $PSScriptRoot ".."
}
$RepoRoot = (Resolve-Path -LiteralPath $RepoRoot).Path
if ([string]::IsNullOrWhiteSpace($ComposeFile)) {
    $ComposeFile = Join-Path $RepoRoot "deploy\compose\docker-compose.dev.yml"
}
$ComposeFile = [System.IO.Path]::GetFullPath($ComposeFile)
$runtimeRoot = Join-Path $RepoRoot ".tmp\feed-local"
if ([string]::IsNullOrWhiteSpace($StateFile)) {
    $StateFile = Join-Path $runtimeRoot "state.json"
}
$StateFile = [System.IO.Path]::GetFullPath($StateFile)
$binDir = Join-Path $runtimeRoot "bin"
$configDir = Join-Path $runtimeRoot "config"
$logDir = Join-Path $runtimeRoot "logs"

$middleware = @("etcd", "zookeeper", "kafka", "canal-server", "elasticsearch")
$composeProjectServices = @("gateway", "user", "storage", "counter", "counter-aggregator", "knowpost", "relation", "search")
$services = @(
    [pscustomobject]@{ Name = "user"; Package = "./services/user/cmd/user"; Config = "services/user/cmd/user/etc/user-docker.yaml"; Ports = @(20002) },
    [pscustomobject]@{ Name = "storage"; Package = "./services/storage/cmd/storage"; Config = "services/storage/cmd/storage/etc/storage-docker.yaml"; Ports = @(20013) },
    [pscustomobject]@{ Name = "counter"; Package = "./services/counter/cmd/counter"; Config = "services/counter/cmd/counter/etc/counter-docker.yaml"; Ports = @(20003, 19103) },
    [pscustomobject]@{ Name = "counter-aggregator"; Package = "./services/counter/aggregator"; Config = "services/counter/aggregator/etc/aggregator.yaml"; Ports = @() },
    [pscustomobject]@{ Name = "relation"; Package = "./services/relation/cmd/relation"; Config = "services/relation/cmd/relation/etc/relation-docker.yaml"; Ports = @(20006, 16066, 19105) },
    [pscustomobject]@{ Name = "knowpost"; Package = "./services/knowpost/cmd/knowpost"; Config = "services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml"; Ports = @(20004, 16064, 19104) },
    [pscustomobject]@{ Name = "search"; Package = "./services/search/cmd/search"; Config = "services/search/cmd/search/etc/search-docker.yaml"; Ports = @(20017, 19107) },
    [pscustomobject]@{ Name = "gateway"; Package = "./services/gateway"; Config = "services/gateway/etc/gateway-docker.yaml"; Ports = @(18080) }
)

function Invoke-Checked {
    param([string]$Command, [string[]]$Arguments)
    & $Command @Arguments
    if ($LASTEXITCODE -ne 0) {
        throw "$Command failed with exit code $LASTEXITCODE"
    }
}

function Test-TcpPort {
    param([int]$Port, [int]$TimeoutMilliseconds = 1000)
    $client = [System.Net.Sockets.TcpClient]::new()
    try {
        $pending = $client.BeginConnect("127.0.0.1", $Port, $null, $null)
        if (-not $pending.AsyncWaitHandle.WaitOne($TimeoutMilliseconds, $false)) {
            return $false
        }
        $client.EndConnect($pending)
        return $true
    }
    catch {
        return $false
    }
    finally {
        $client.Dispose()
    }
}

function Get-PortOwner {
    param([int]$Port)
    $listener = Get-NetTCPConnection -State Listen -LocalPort $Port -ErrorAction SilentlyContinue | Select-Object -First 1
    if ($null -eq $listener) {
        return 0
    }
    return [int]$listener.OwningProcess
}

function Assert-PortsAvailable {
    param([int[]]$Ports)
    foreach ($port in $Ports) {
        $owner = Get-PortOwner -Port $port
        if ($owner -ne 0) {
            throw "local service port $port is already owned by PID $owner; no unrelated process will be stopped"
        }
    }
}

function Wait-ServicePorts {
    param([string]$Service, [System.Diagnostics.Process]$Process, [int[]]$Ports, [int]$TimeoutSeconds = 60)
    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $deadline) {
        $Process.Refresh()
        if ($Process.HasExited) {
            throw "$Service exited before readiness with code $($Process.ExitCode)"
        }
        $owned = @($Ports | Where-Object { (Get-PortOwner -Port $_) -eq $Process.Id })
        if ($owned.Count -eq $Ports.Count) {
            return
        }
        Start-Sleep -Milliseconds 250
    }
    $observed = @($Ports | ForEach-Object { "$_=$(Get-PortOwner -Port $_)" }) -join ","
    throw "$Service PID $($Process.Id) did not own all expected ports within $TimeoutSeconds seconds; owners: $observed"
}

function Assert-PortsNotExcluded {
    param([int[]]$Ports)
    $raw = & netsh interface ipv4 show excludedportrange protocol=tcp
    if ($LASTEXITCODE -ne 0) {
        throw "cannot read Windows TCP excluded port ranges"
    }
    $ranges = foreach ($line in $raw) {
        if ($line -match '^\s*(\d+)\s+(\d+)\s*\*?\s*$') {
            [pscustomobject]@{ Start = [int]$Matches[1]; End = [int]$Matches[2] }
        }
    }
    foreach ($port in $Ports) {
        foreach ($range in $ranges) {
            if ($port -ge $range.Start -and $port -le $range.End) {
                throw "local service port $port is excluded by Windows range $($range.Start)-$($range.End); choose and record a different local port"
            }
        }
    }
}

function Convert-LocalConfig {
    param([string]$Source, [string]$Destination, [string]$Service)
    $body = Get-Content -LiteralPath $Source -Raw
    $replacements = [ordered]@{
        "etcd:2379" = "127.0.0.1:12379"
        "mysql:3306" = "127.0.0.1:3306"
        "redis:6379" = "127.0.0.1:6379"
        "kafka:29092" = "127.0.0.1:9092"
        "user:9002" = "127.0.0.1:20002"
        "storage:9013" = "127.0.0.1:20013"
        "counter:9003" = "127.0.0.1:20003"
        "knowpost:9004" = "127.0.0.1:20004"
        "relation:9006" = "127.0.0.1:20006"
        "search:9017" = "127.0.0.1:20017"
        "http://elasticsearch:9200" = "http://127.0.0.1:9200"
        "0.0.0.0:9002" = "127.0.0.1:20002"
        "0.0.0.0:9013" = "127.0.0.1:20013"
        "0.0.0.0:9003" = "127.0.0.1:20003"
        "0.0.0.0:9004" = "127.0.0.1:20004"
        "0.0.0.0:9006" = "127.0.0.1:20006"
        "0.0.0.0:9017" = "127.0.0.1:20017"
        "0.0.0.0:6064" = "127.0.0.1:16064"
        "0.0.0.0:6066" = "127.0.0.1:16066"
        "Port: 8080" = "Port: 18080"
        "Port: 9102" = "Port: 19102"
        "Port: 9103" = "Port: 19103"
        "Port: 9104" = "Port: 19104"
        "Port: 9105" = "Port: 19105"
        "Port: 9106" = "Port: 19106"
        "Port: 9107" = "Port: 19107"
    }
    foreach ($entry in $replacements.GetEnumerator()) {
        $body = $body.Replace($entry.Key, $entry.Value)
    }
	if ($LoggingPreset -eq "benchmark-error-only") {
		# Generated configs only: retain error output while eliminating the
		# request-level info/stat stream from the measured service processes.
		$body = [regex]::Replace($body, '(?m)^[ \t]*Level:[ \t]*(debug|info|error|severe)[ \t]*\r?$', '')
		$body = [regex]::Replace($body, '(?m)^[ \t]*Stat:[ \t]*(true|false)[ \t]*\r?$', '')
		$body = [regex]::Replace(
			$body,
			'(?m)^([ \t]*)Mode:[ \t]*console[ \t]*\r?$',
			'${1}Mode: console' + "`r`n" + '${1}Level: error' + "`r`n" + '${1}Stat: false'
		)
		# zrpc's stat interceptor serializes every request before log-level
		# filtering. Disable it in generated benchmark configs so the measured
		# path does not pay request JSON/log formatting costs.
		if ($body -match '(?m)^Rpc:[ \t]*\r?$') {
			$body = [regex]::Replace(
				$body,
				'(?m)^Rpc:[ \t]*\r?$',
				"Rpc:`r`n  Middlewares:`r`n    Stat: false"
			)
		} elseif ($body -notmatch '(?m)^Middlewares:[ \t]*\r?$') {
			$body = $body.TrimEnd() + "`r`n`r`nMiddlewares:`r`n  Stat: false`r`n"
		}
		if ($Service -eq "gateway") {
			if ($body -match '(?m)^Log:[ \t]*\r?$' -or $body -match '(?m)^HTTPAccessLog:[ \t]*') {
				throw "gateway source config already defines Log or HTTPAccessLog; benchmark logging conversion must be updated explicitly"
			}
			$body = $body.TrimEnd() + "`r`n`r`nLog:`r`n  Mode: console`r`n  Level: error`r`n  Stat: false`r`n`r`nHTTPAccessLog: false`r`n"
		}
	}
    Set-Content -LiteralPath $Destination -Value $body -Encoding utf8NoBOM
}

Write-Host "Feed local-service topology plan"
Write-Host "  topology: compose-middleware-local-services"
Write-Host "  feature preset: $FeaturePreset"
Write-Host "  logging preset: $LoggingPreset"
Write-Host "  middleware (existing images only): $($middleware -join ', ')"
Write-Host "  local project services: $($services.Name -join ', ')"
Write-Host "  local runtime: Gateway=18080 User=20002 Storage=20013 Counter=20003 KnowPost=20004 Relation=20006 Search=20017"
Write-Host "  local debug: KnowPost=16064 Relation=16066; KnowPost Prometheus=19104"
Write-Host "  state: $StateFile"
Write-Host "  note: an image lookup failure is a Docker registry/mirror-source issue; this script never pulls or builds middleware images"

if ($WhatIfPreference) {
    Write-Host "WhatIf: no container, process, directory, or state-file changes were made."
    return
}

if (Test-Path -LiteralPath $StateFile) {
    throw "state file already exists: $StateFile; stop the recorded services before starting another set"
}
foreach ($required in @($ComposeFile, (Join-Path $RepoRoot "go.mod"))) {
    if (-not (Test-Path -LiteralPath $required)) {
        throw "required path does not exist: $required"
    }
}
foreach ($port in @(3306, 6379)) {
    if (-not (Test-TcpPort -Port $port)) {
        throw "required host dependency is not reachable at 127.0.0.1:$port"
    }
}
Assert-PortsNotExcluded -Ports @(18080, 20002, 20003, 20004, 20006, 20013, 20017, 16064, 16066, 19102, 19103, 19104, 19105, 19106, 19107)

New-Item -ItemType Directory -Force -Path $binDir, $configDir, $logDir | Out-Null

$env:POD_IP = "127.0.0.1"
$env:FEED_STRATEGY = $Strategy
$env:FEED_OBSERVABILITY = "true"
$featureEnabled = if ($FeaturePreset -eq "control") { "false" } else { "true" }
$env:FEED_RELATION_EPOCH_ENABLED = $featureEnabled
$env:FEED_CONTENT_SAFETY_EPOCH_ENABLED = $featureEnabled
$env:FEED_ROUTE_SNAPSHOT_ENABLED = $featureEnabled
$env:FEED_COMBINED_PIPELINE_ENABLED = $featureEnabled
$env:FEED_CURSOR_PAGINATION_ENABLED = $featureEnabled
$env:FEED_PAGE_CACHE_MODE = if ($FeaturePreset -eq "treatment") { "l1-l2" } else { "off" }

# Project containers are stopped by Compose name; arbitrary port owners are
# never killed. Middleware volumes and already-running middleware are retained.
Invoke-Checked "docker" (@("compose", "-f", $ComposeFile, "stop") + $composeProjectServices)
Invoke-Checked "docker" (@("compose", "-f", $ComposeFile, "up", "-d", "--no-build", "--pull", "never", "--wait", "--wait-timeout", "180") + $middleware)
$allServicePorts = @($services | ForEach-Object { $_.Ports } | ForEach-Object { $_ })
Assert-PortsAvailable -Ports $allServicePorts

$started = [System.Collections.Generic.List[object]]::new()
try {
    Push-Location $RepoRoot
    try {
        foreach ($service in $services) {
			Assert-PortsAvailable -Ports $service.Ports
            $executable = Join-Path $binDir ($service.Name + ".exe")
            Invoke-Checked "go" @("build", "-o", $executable, $service.Package)

            $sourceConfig = Join-Path $RepoRoot $service.Config
            $localConfig = Join-Path $configDir ($service.Name + ".yaml")
            Convert-LocalConfig -Source $sourceConfig -Destination $localConfig -Service $service.Name

            $stdout = Join-Path $logDir ($service.Name + ".stdout.log")
            $stderr = Join-Path $logDir ($service.Name + ".stderr.log")
            $process = Start-Process -FilePath $executable `
                -ArgumentList @("-f", $localConfig) `
                -WorkingDirectory $RepoRoot `
                -WindowStyle Hidden `
                -RedirectStandardOutput $stdout `
                -RedirectStandardError $stderr `
                -PassThru
            $record = [pscustomobject]@{
                name = $service.Name
                pid = $process.Id
                executable = $executable
				process_start_unix_ms = ([DateTimeOffset]$process.StartTime.ToUniversalTime()).ToUnixTimeMilliseconds()
                config = $localConfig
                stdout = $stdout
                stderr = $stderr
				ports = @($service.Ports)
            }
            $started.Add($record)
			Wait-ServicePorts -Service $service.Name -Process $process -Ports $service.Ports
        }
    }
    finally {
        Pop-Location
    }

    $state = [pscustomobject]@{
        topology = "compose-middleware-local-services"
        repo_root = $RepoRoot
        compose_file = $ComposeFile
        strategy = $Strategy
		feature_preset = $FeaturePreset
        feature_flags = [pscustomobject]@{
            observability = $env:FEED_OBSERVABILITY
            relation_epoch = $env:FEED_RELATION_EPOCH_ENABLED
            safety_epoch = $env:FEED_CONTENT_SAFETY_EPOCH_ENABLED
            route_snapshot = $env:FEED_ROUTE_SNAPSHOT_ENABLED
            combined_pipeline = $env:FEED_COMBINED_PIPELINE_ENABLED
            cursor_pagination = $env:FEED_CURSOR_PAGINATION_ENABLED
            page_cache_mode = $env:FEED_PAGE_CACHE_MODE
        }
		logging = if ($LoggingPreset -eq "benchmark-error-only") {
			[pscustomobject]@{
				preset = "benchmark-error-only"
				level = "error"
				stat = $false
				gateway_access_log = $false
				rpc_stat_middleware = $false
				sql_statement_info = $false
				persistence = "error-only-bounded"
			}
		}
		else {
			[pscustomobject]@{
				preset = "development"
				level = "info"
				stat = $true
				gateway_access_log = $true
				rpc_stat_middleware = $true
				sql_statement_info = $true
				persistence = "unbounded-console-files"
			}
		}
        started_at = (Get-Date).ToUniversalTime().ToString("o")
        middleware = $middleware
        local_ports = [pscustomobject]@{
            gateway_http = 18080
            user_rpc = 20002
            storage_rpc = 20013
            counter_rpc = 20003
            knowpost_rpc = 20004
            relation_rpc = 20006
            search_rpc = 20017
            knowpost_pprof = 16064
            relation_pprof = 16066
            knowpost_prometheus = 19104
        }
        services = @($started)
    }
    $state | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath $StateFile -Encoding utf8NoBOM
}
catch {
	$startupError = $_
	$cleanupRefusals = [System.Collections.Generic.List[string]]::new()
    foreach ($record in @($started) | Sort-Object pid -Descending) {
		$recordedPid = [int]$record.pid
		$process = Get-Process -Id $recordedPid -ErrorAction SilentlyContinue
		if ($null -eq $process) {
			continue
		}
		try {
			$actualExecutable = [System.IO.Path]::GetFullPath($process.Path)
			$expectedExecutable = [System.IO.Path]::GetFullPath([string]$record.executable)
			$actualStartUnixMs = ([DateTimeOffset]$process.StartTime.ToUniversalTime()).ToUnixTimeMilliseconds()
			$expectedStartUnixMs = [int64]$record.process_start_unix_ms
		}
		catch {
			$cleanupRefusals.Add("$($record.name) PID $recordedPid identity could not be verified")
			continue
		}
		if ((-not $actualExecutable.Equals($expectedExecutable, [System.StringComparison]::OrdinalIgnoreCase)) -or
			$actualStartUnixMs -ne $expectedStartUnixMs) {
			$cleanupRefusals.Add("$($record.name) PID $recordedPid identity changed")
			continue
		}
		Stop-Process -Id $recordedPid
    }
	if ($cleanupRefusals.Count -gt 0) {
		throw "startup failed: $startupError; startup cleanup refused to stop unverified processes: $($cleanupRefusals -join '; ')"
	}
	throw $startupError
}

Write-Host "Local Feed services are ready. State: $StateFile"
