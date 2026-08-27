# Docker 环境快速指南

## 🚀 快速开始

### Windows 用户

```powershell
# 方式一：使用启动脚本（推荐）
.\scripts\docker-start.ps1

# 方式二：手动启动测试环境
docker compose -f deploy\compose\docker-compose.test.yml up -d --build
```

### Linux/Mac 用户

```bash
# 方式一：使用启动脚本（推荐）
bash scripts/docker-start.sh

# 方式二：手动启动测试环境
docker compose -f deploy/compose/docker-compose.test.yml up -d --build
```

## 📦 部署模式

### 测试模式（推荐首次使用）

包含最小化的服务集合，用于快速验证 Etcd 服务发现功能：

- **基础设施**：MySQL、Redis、Etcd
- **核心服务**：Gateway、User-Storage、Counter

```bash
docker compose -f deploy/compose/docker-compose.test.yml up -d --build
```

### 完整模式

包含所有服务和基础设施：

```bash
docker compose -f deploy/compose/docker-compose.full.yml up -d --build
```

## ✅ 验证部署

### 1. 检查服务状态

```bash
docker compose -f deploy/compose/docker-compose.test.yml ps
```

所有服务应该处于 `Up` 状态。

### 2. 检查 Etcd 服务注册

```bash
# 查看所有注册的服务
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix ""

# 查看特定服务（例如 counter.rpc）
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix "counter.rpc"
```

### 3. 测试 API 调用

```bash
# 测试 Gateway 通过服务发现调用 Counter
curl http://localhost:8080/api/v1/counter/knowpost/test

# 预期返回
{"counts":{"fav":0,"like":0},"entityId":"test","entityType":"knowpost"}
```

### 4. 查看服务日志

```bash
# 查看 Gateway 日志
docker compose -f deploy/compose/docker-compose.test.yml logs -f gateway

# 查看 Counter 日志
docker compose -f deploy/compose/docker-compose.test.yml logs -f counter

# 查看所有服务日志
docker compose -f deploy/compose/docker-compose.test.yml logs -f
```

## 🧪 测试多实例和高可用

### 启动第二个 Counter 实例

```bash
# 扩展 Counter 服务到 2 个实例
docker compose -f deploy/compose/docker-compose.test.yml up -d --scale counter=2

# 查看 Etcd 中的注册信息（应该看到 2 个实例）
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix "counter.rpc"

# 多次请求验证负载均衡
for i in {1..10}; do curl -s http://localhost:8080/api/v1/counter/knowpost/$i | grep entityId; done
```

### 测试故障切换

```bash
# 停止一个 Counter 实例
docker stop zg-counter

# 继续请求，服务应该仍然正常
curl http://localhost:8080/api/v1/counter/knowpost/test

# 查看 Etcd，停止的实例应该自动注销
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix "counter.rpc"
```

## 🛠️ 常用命令

### 启动服务

```bash
# 启动所有服务
docker compose -f deploy/compose/docker-compose.test.yml up -d

# 重新构建并启动
docker compose -f deploy/compose/docker-compose.test.yml up -d --build

# 启动特定服务
docker compose -f deploy/compose/docker-compose.test.yml up -d gateway
```

### 停止服务

```bash
# 停止所有服务（保留数据）
docker compose -f deploy/compose/docker-compose.test.yml down

# 停止并删除所有数据（慎用！）
docker compose -f deploy/compose/docker-compose.test.yml down -v

# 停止特定服务
docker compose -f deploy/compose/docker-compose.test.yml stop gateway
```

### 查看状态

```bash
# 查看服务状态
docker compose -f deploy/compose/docker-compose.test.yml ps

# 查看资源使用
docker stats

# 查看服务健康状态
docker inspect zg-mysql | grep -A 10 Health
```

### 查看日志

```bash
# 实时查看日志
docker compose -f deploy/compose/docker-compose.test.yml logs -f

# 查看最近 100 行日志
docker compose -f deploy/compose/docker-compose.test.yml logs --tail=100

# 查看特定服务日志
docker compose -f deploy/compose/docker-compose.test.yml logs -f gateway counter
```

### 执行命令

```bash
# 进入容器
docker compose -f deploy/compose/docker-compose.test.yml exec gateway sh

# 执行数据库查询
docker compose -f deploy/compose/docker-compose.test.yml exec mysql \
  mysql -uroot -pZz123456 zhiguang -e "SHOW TABLES;"

# 查看 Redis 数据
docker compose -f deploy/compose/docker-compose.test.yml exec redis redis-cli keys "*"

# 查看 Etcd 数据
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix ""
```

## ⚙️ 配置说明

### 配置文件位置

- **本地开发**：`services/*/etc/*.yaml`
- **Docker 环境**：`services/*/etc/*-docker.yaml`

### 主要差异

| 配置项 | 本地开发 | Docker 环境 |
|-------|---------|------------|
| Etcd 地址 | 127.0.0.1:2379 | etcd:2379 |
| MySQL 地址 | 127.0.0.1:3306 | mysql:3306 |
| Redis 地址 | 127.0.0.1:6379 | redis:6379 |
| Kafka 地址 | 127.0.0.1:9092 | kafka:29092 |

### 修改配置

如果需要修改配置：

1. 编辑 `services/*/etc/*-docker.yaml` 文件
2. 重启服务：
   ```bash
   docker compose -f deploy/compose/docker-compose.test.yml restart gateway
   ```

## 🐛 故障排查

### 问题 1：服务启动失败

```bash
# 查看详细日志
docker compose -f deploy/compose/docker-compose.test.yml logs gateway

# 检查依赖服务是否健康
docker compose -f deploy/compose/docker-compose.test.yml ps
```

### 问题 2：无法连接到 Etcd

```bash
# 检查 Etcd 健康状态
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl endpoint health

# 从应用容器测试连接
docker compose -f deploy/compose/docker-compose.test.yml exec gateway ping etcd
```

### 问题 3：数据库连接失败

```bash
# 检查 MySQL 是否就绪
docker compose -f deploy/compose/docker-compose.test.yml exec mysql \
  mysql -uroot -pZz123456 -e "SELECT 1"

# 查看 MySQL 日志
docker compose -f deploy/compose/docker-compose.test.yml logs mysql
```

### 问题 4：构建失败

```bash
# 清理缓存重新构建
docker compose -f deploy/compose/docker-compose.test.yml build --no-cache

# 如果 go mod download 慢，设置代理
export GOPROXY=https://goproxy.cn,direct
docker compose -f deploy/compose/docker-compose.test.yml build
```

### 问题 5：端口冲突

如果本地已经运行了某些服务（如 MySQL、Redis），可以：

1. 停止本地服务
2. 或修改 docker-compose 文件中的端口映射

## 📊 监控和维护

### 查看资源使用

```bash
# 实时资源使用
docker stats

# 查看磁盘使用
docker system df

# 查看网络
docker network ls
docker network inspect compose_default
```

### 数据备份

```bash
# 备份 MySQL 数据
docker compose -f deploy/compose/docker-compose.test.yml exec mysql \
  mysqldump -uroot -pZz123456 zhiguang > backup.sql

# 备份 Redis 数据
docker compose -f deploy/compose/docker-compose.test.yml exec redis \
  redis-cli SAVE
docker cp zg-redis:/data/dump.rdb ./redis-backup.rdb
```

### 清理资源

```bash
# 清理未使用的镜像
docker image prune -a

# 清理未使用的容器
docker container prune

# 清理未使用的卷（慎用！）
docker volume prune

# 清理所有未使用的资源
docker system prune -a
```

## 🎯 下一步

- 阅读完整部署文档：[docs/docker-deployment.md](../../docs/docker-deployment.md)
- 了解服务发现原理：[docs/etcd-service-discovery.md](../../docs/etcd-service-discovery.md)
- 配置生产环境：[docs/production-deployment.md](../../docs/production-deployment.md)

## 📞 需要帮助？

- 查看日志获取错误详情
- 检查 GitHub Issues
- 阅读项目文档
