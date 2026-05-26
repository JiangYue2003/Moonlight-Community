@echo off
setlocal

cd /d "%~dp0\.."

if "%ZG_MIGRATE_DSN%"=="" (
  set "ZG_MIGRATE_DSN=mysql://root:Zz123456@tcp(127.0.0.1:3306)/zhiguang"
)
set "PATH_DIR=db/migrations"

if "%~1"=="" (
  set "CMD=up"
) else (
  set "CMD=%~1"
)

if /i "%CMD%"=="up"      goto :up
if /i "%CMD%"=="down"    goto :down
if /i "%CMD%"=="version" goto :version
if /i "%CMD%"=="force"   goto :force
if /i "%CMD%"=="drop"    goto :drop

echo Usage: %~n0 {up^|down [N]^|version^|force ^<v^>^|drop}
exit /b 1

:up
migrate -path "%PATH_DIR%" -database "%ZG_MIGRATE_DSN%" up
exit /b %errorlevel%

:down
set "N=%~2"
if "%N%"=="" set "N=1"
migrate -path "%PATH_DIR%" -database "%ZG_MIGRATE_DSN%" down %N%
exit /b %errorlevel%

:version
migrate -path "%PATH_DIR%" -database "%ZG_MIGRATE_DSN%" version
exit /b %errorlevel%

:force
migrate -path "%PATH_DIR%" -database "%ZG_MIGRATE_DSN%" force %~2
exit /b %errorlevel%

:drop
migrate -path "%PATH_DIR%" -database "%ZG_MIGRATE_DSN%" drop
exit /b %errorlevel%
