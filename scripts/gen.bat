@echo off
setlocal enabledelayedexpansion

cd /d "%~dp0\.."

echo == goctl version ==
goctl --version

echo == generate user-rpc ==
goctl rpc protoc proto/user/user.proto ^
  --go_out=services/user/rpc ^
  --go-grpc_out=services/user/rpc ^
  --zrpc_out=services/user/rpc ^
  -m
if errorlevel 1 goto :err
rem The context owns runtime implementation; retain only generated protobuf/client contracts.
if exist services\user\rpc\internal rmdir /s /q services\user\rpc\internal
if exist services\user\rpc\etc rmdir /s /q services\user\rpc\etc
if exist services\user\rpc\user.go del /q services\user\rpc\user.go

echo == generate auth-rpc ==
goctl rpc protoc proto/auth/auth.proto ^
  --go_out=services/auth/rpc/internal ^
  --go-grpc_out=services/auth/rpc/internal ^
  --zrpc_out=services/auth/rpc ^
  -m
if errorlevel 1 goto :err

echo == generate counter-rpc ==
goctl rpc protoc proto/counter/counter.proto ^
	--go_out=services/counter/rpc ^
	--go-grpc_out=services/counter/rpc ^
	--zrpc_out=services/counter/rpc ^
	-m
if errorlevel 1 goto :err
rem Counter owns the runtime implementation outside generated public contracts.
if exist services\counter\rpc\internal rmdir /s /q services\counter\rpc\internal
if exist services\counter\rpc\etc rmdir /s /q services\counter\rpc\etc
if exist services\counter\rpc\counter.go del /q services\counter\rpc\counter.go

echo == generate knowpost protobuf ==
protoc --go_out=services/knowpost/rpc ^
  --go-grpc_out=services/knowpost/rpc ^
  proto/knowpost/knowpost.proto
if errorlevel 1 goto :err

echo == generate models ==
goctl model mysql ddl -src "db/migrations/000001_init_users.up.sql" ^
  -dir services/user/internal/adapter/model -c
if errorlevel 1 goto :err
goctl model mysql ddl -src "db/migrations/000002_init_login_logs.up.sql" ^
  -dir services/user/internal/adapter/model_auth -c
if errorlevel 1 goto :err

echo == done ==
exit /b 0

:err
echo generation failed
exit /b 1
