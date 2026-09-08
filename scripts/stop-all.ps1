param(
  [string]$PidFile = ""
)

$ErrorActionPreference = "Stop"

$Root = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
$PidFiles = if ([string]::IsNullOrWhiteSpace($PidFile)) {
  @(
    (Join-Path $Root "var/run/dev/pids.json"),
    (Join-Path $Root "logs/dev/pids.json")
  )
} else {
  @([System.IO.Path]::GetFullPath($PidFile))
}

function Write-Info([string]$msg) {
  Write-Host "[INFO] $msg" -ForegroundColor Cyan
}

function Write-WarnMsg([string]$msg) {
  Write-Host "[WARN] $msg" -ForegroundColor Yellow
}

$existingPidFiles = @($PidFiles | Where-Object { Test-Path -LiteralPath $_ })
if ($existingPidFiles.Count -eq 0) {
  Write-WarnMsg ("PID file not found: {0}" -f ($PidFiles -join ", "))
  exit 0
}

foreach ($path in $existingPidFiles) {
  $items = Get-Content -LiteralPath $path -Raw | ConvertFrom-Json
  $remaining = [System.Collections.Generic.List[object]]::new()
  if ($null -eq $items) {
    Write-WarnMsg "PID file is empty: $path"
  } else {
    foreach ($it in @($items)) {
      $p = Get-Process -Id $it.pid -ErrorAction SilentlyContinue
      if ($null -eq $p) {
        Write-WarnMsg ("Already exited or not found: {0} (PID={1})" -f $it.service, $it.pid)
        continue
      }
      try {
        if ($null -ne $it.executable -and $null -ne $it.processStartUnixMs) {
          $actualExecutable = [System.IO.Path]::GetFullPath($p.Path)
          $expectedExecutable = [System.IO.Path]::GetFullPath([string]$it.executable)
          $actualStartUnixMs = ([DateTimeOffset]$p.StartTime.ToUniversalTime()).ToUnixTimeMilliseconds()
          if ((-not $actualExecutable.Equals($expectedExecutable, [System.StringComparison]::OrdinalIgnoreCase)) -or
              $actualStartUnixMs -ne [int64]$it.processStartUnixMs) {
            throw "process identity does not match the recorded executable and start time"
          }
        } else {
          throw "legacy PID record has no executable identity; refusing to stop it"
        }
        Stop-Process -Id $p.Id -Force -ErrorAction Stop
        Write-Info ("Stopped {0} (PID={1})" -f $it.service, $it.pid)
      } catch {
        $remaining.Add($it)
        Write-WarnMsg ("Refused or failed to stop {0} (PID={1}): {2}" -f $it.service, $it.pid, $_.Exception.Message)
      }
    }
  }

  if ($remaining.Count -eq 0) {
    Remove-Item -LiteralPath $path -Force
    Write-Info "PID file removed: $path"
  } else {
    $remaining | ConvertTo-Json -Depth 4 | Set-Content -LiteralPath $path -Encoding UTF8
    throw "$($remaining.Count) recorded process(es) could not be stopped; PID file retained: $path"
  }
}
