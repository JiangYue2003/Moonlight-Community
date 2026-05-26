#!/usr/bin/env bash
# 批量代码生成入口（goctl）
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "== goctl version =="
goctl --version

echo "== generate user-rpc =="
goctl rpc protoc proto/user/user.proto \
  --go_out=services/user/rpc/internal \
  --go-grpc_out=services/user/rpc/internal \
  --zrpc_out=services/user/rpc \
  -m

echo "== generate auth-rpc =="
goctl rpc protoc proto/auth/auth.proto \
  --go_out=services/auth/rpc/internal \
  --go-grpc_out=services/auth/rpc/internal \
  --zrpc_out=services/auth/rpc \
  -m

echo "== generate counter-rpc =="
goctl rpc protoc proto/counter/counter.proto \
  --go_out=services/counter/rpc/internal \
  --go-grpc_out=services/counter/rpc/internal \
  --zrpc_out=services/counter/rpc \
  -m

echo "== generate auth-api =="
goctl api go -api services/auth/api/auth.api -dir services/auth/api

echo "== generate counter-api =="
goctl api go -api services/counter/api/counter.api -dir services/counter/api

echo "== generate models =="
goctl model mysql ddl -src "db/migrations/000001_init_users.up.sql" \
  -dir services/user/rpc/internal/model -c
goctl model mysql ddl -src "db/migrations/000002_init_login_logs.up.sql" \
  -dir services/auth/rpc/internal/model -c

echo "== done =="
