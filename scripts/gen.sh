#!/usr/bin/env bash
# 批量代码生成入口（goctl）
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "== goctl version =="
goctl --version

echo "== generate user-rpc =="
goctl rpc protoc proto/user/user.proto \
  --go_out=services/user/rpc \
  --go-grpc_out=services/user/rpc \
  --zrpc_out=services/user/rpc \
  -m
# The context owns runtime implementation; retain only generated protobuf/client contracts.
rm -rf services/user/rpc/internal services/user/rpc/etc
rm -f services/user/rpc/user.go

echo "== generate auth-rpc =="
goctl rpc protoc proto/auth/auth.proto \
  --go_out=services/auth/rpc/internal \
  --go-grpc_out=services/auth/rpc/internal \
  --zrpc_out=services/auth/rpc \
  -m

echo "== generate counter-rpc =="
goctl rpc protoc proto/counter/counter.proto \
	--go_out=services/counter/rpc \
	--go-grpc_out=services/counter/rpc \
	--zrpc_out=services/counter/rpc \
	-m
# Counter owns the runtime implementation outside generated public contracts.
rm -rf services/counter/rpc/internal services/counter/rpc/etc
rm -f services/counter/rpc/counter.go

echo "== generate knowpost protobuf =="
protoc --go_out=services/knowpost/rpc \
  --go-grpc_out=services/knowpost/rpc \
  proto/knowpost/knowpost.proto

echo "== generate models =="
goctl model mysql ddl -src "db/migrations/000001_init_users.up.sql" \
  -dir services/user/internal/adapter/model -c
goctl model mysql ddl -src "db/migrations/000002_init_login_logs.up.sql" \
  -dir services/user/internal/adapter/model_auth -c

echo "== done =="
