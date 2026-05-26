@echo off
setlocal enabledelayedexpansion

cd /d "%~dp0\.."

echo == goctl version ==
goctl --version

echo == generate user-rpc ==
goctl rpc protoc proto/user/user.proto ^
  --go_out=services/user/rpc/internal ^
  --go-grpc_out=services/user/rpc/internal ^
  --zrpc_out=services/user/rpc ^
  -m
if errorlevel 1 goto :err

echo == generate auth-rpc ==
goctl rpc protoc proto/auth/auth.proto ^
  --go_out=services/auth/rpc/internal ^
  --go-grpc_out=services/auth/rpc/internal ^
  --zrpc_out=services/auth/rpc ^
  -m
if errorlevel 1 goto :err

echo == generate counter-rpc ==
goctl rpc protoc proto/counter/counter.proto ^
  --go_out=services/counter/rpc/internal ^
  --go-grpc_out=services/counter/rpc/internal ^
  --zrpc_out=services/counter/rpc ^
  -m
if errorlevel 1 goto :err

echo == generate auth-api ==
goctl api go -api services/auth/api/auth.api -dir services/auth/api
if errorlevel 1 goto :err

echo == generate counter-api ==
goctl api go -api services/counter/api/counter.api -dir services/counter/api
if errorlevel 1 goto :err

echo == generate models ==
goctl model mysql ddl -src "db/migrations/000001_init_users.up.sql" ^
  -dir services/user/rpc/internal/model -c
goctl model mysql ddl -src "db/migrations/000002_init_login_logs.up.sql" ^
  -dir services/auth/rpc/internal/model -c

echo == done ==
exit /b 0

:err
echo generation failed
exit /b 1
