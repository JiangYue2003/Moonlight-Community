#!/usr/bin/env bash
# DB 迁移辅助脚本
# 用法：scripts/migrate.sh up | down [N] | version | force <version>
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

DB_DSN="${ZG_MIGRATE_DSN:-mysql://root:Zz123456@tcp(127.0.0.1:3306)/zhiguang}"
PATH_DIR="db/migrations"

CMD="${1:-up}"
shift || true

case "$CMD" in
  up)        migrate -path "$PATH_DIR" -database "$DB_DSN" up ;;
  down)      migrate -path "$PATH_DIR" -database "$DB_DSN" down "${1:-1}" ;;
  version)   migrate -path "$PATH_DIR" -database "$DB_DSN" version ;;
  force)     migrate -path "$PATH_DIR" -database "$DB_DSN" force "$1" ;;
  drop)      migrate -path "$PATH_DIR" -database "$DB_DSN" drop ;;
  *)
    echo "Usage: $0 {up|down [N]|version|force <v>|drop}" >&2
    exit 1 ;;
esac
