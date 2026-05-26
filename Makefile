.PHONY: gen build tidy test migrate-up migrate-down up down

ROOT := $(CURDIR)

# 代码生成
gen:
	@bash scripts/gen.sh

# 整理依赖
tidy:
	go mod tidy

# 全量构建
build:
	go build ./...

# 运行单元测试
test:
	go test ./...

# 数据库迁移
migrate-up:
	@bash scripts/migrate.sh up

migrate-down:
	@bash scripts/migrate.sh down

# 启动本地依赖
up:
	docker compose -f deploy/compose/docker-compose.dev.yml up -d

down:
	docker compose -f deploy/compose/docker-compose.dev.yml down
