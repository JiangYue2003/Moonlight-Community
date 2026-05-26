# DB Migrations

使用 [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 管理数据库变更。

## 安装 migrate CLI

```bash
# Windows (scoop)
scoop install migrate

# Mac
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
```

## 命名规范

`{version}_{description}.{up|down}.sql`

版本号必须递增（建议 6 位），同一版本必须有 up/down 两份。

## 常用命令

```bash
# 应用所有未执行的迁移
migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/zhiguang" up

# 回滚最近 N 个迁移
migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/zhiguang" down 1

# 查看当前版本
migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/zhiguang" version
```

亦可使用 `scripts/migrate.sh` 或 `scripts/migrate.bat`。

## 当前迁移文件

| 版本 | 描述 |
|------|------|
| 000001 | 创建 users 表 |
| 000002 | 创建 login_logs 表 |
| 000003 | 创建 know_posts 表（依赖 users） |
| 000004 | 创建 outbox 表 |
| 000005 | 创建 following + follower 表 |
