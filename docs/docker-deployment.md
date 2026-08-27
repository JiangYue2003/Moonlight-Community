# Docker 部署说明文档

## 概述

本文档说明如何使用 Docker Compose 部署知光项目的完整环境，包括基础设施和应用服务。

## 架构

### 基础设施层
- **MySQL 8.0**：主数据库
- **Redis 7**：缓存和计数存储
- **Kafka 3.7 + ZooKeeper**：消息队列
- **Etcd 3.5**：服务注册与发现
- **Canal 1.1.7**：MySQL binlog 监听
- **Elasticsearch 9.0**：全文搜索
- **Adminer**：数据库管理工具（可选）

### 应用服务层
- **Gateway** (8080)：统一 HTTP 入口
- **User + Storage**：用户认证和文件存储 RPC
- **Counter**：计数服务（RPC + Aggregator）
- **KnowPost**：知文服务 RPC
- **Relation**：关系服务（RPC + Syncer）
- **Search**：搜索服务（RPC + Indexer）
- **LLM**：AI 服务（RPC + RAG Indexer）
- **Agent** (8011)：知识助手独立入口

## 前置准备

### 1. 安装 Docker 和 Docker Compose

确保已安装：
- Docker 20.10+
- Docker Compose v2.0+

### 2. 环境变量配置

创建 `.env` 文件（可选，用于 API Key）：
```bash
# LLM API Keys（可选）
DEEPSEEK_API_KEY=your_deepseek_key
DASHSCOPE_API_KEY=your_dashscope_key
MILVUS_API_KEY=your_milvus_key
```

### 3. 准备证书文件

确保 `certs/` 目录包含以下文件：
- `jwt_private.pem`
- `jwt_public.pem`

如果没有，可以生成：
```bash
# 生成 RSA 密钥对
openssl genrsa -out certs/jwt_private.pem 2048
openssl rsa -in certs/jwt_private.pem -pubout -out certs/jwt_public.pem
```

## 快速开始

### 方式一：完整部署（所有服务）

```bash
# 1. 构建镜像并启动所有服务
docker compose -f deploy/compose/docker-compose.full.yml up -d --build

# 2. 查看服务状态
docker compose -f deploy/compose/docker-compose.full.yml ps

# 3. 查看日志
docker compose -f deploy/compose/docker-compose.full.yml logs -f gateway

# 4. 停止所有服务
docker compose -f deploy/compose/docker-compose.full.yml down
```

### 方式二：分步部署（推荐首次部署）

#### Step 1: 启动基础设施

```bash
# 仅启动基础设施服务
docker compose -f deploy/compose/docker-compose.full.yml up -d \
  mysql redis zookeeper kafka etcd elasticsearch canal-server
```

等待服务健康检查通过（约 30-60 秒）：
```bash
# 检查 MySQL
docker compose -f deploy/compose/docker-compose.full.yml ps mysql

# 检查 Etcd
docker compose -f deploy/compose/docker-compose.full.yml exec etcd etcdctl endpoint health
```

#### Step 2: 执行数据库迁移

```bash
# 方式 A: 使用 migrate 工具
./scripts/migrate.sh up

# 方式 B: 手动执行 SQL
docker compose -f deploy/compose/docker-compose.full.yml exec mysql \
  mysql -uroot -pZz123456 zhiguang < db/migrations/001_create_users.sql
```

#### Step 3: 启动应用服务

```bash
# 启动所有应用服务
docker compose -f deploy/compose/docker-compose.full.yml up -d \
  gateway user-storage counter knowpost relation search llm agent
```

## 验证部署

### 1. 检查服务注册

```bash
# 查看 Etcd 中注册的服务
docker compose -f deploy/compose/docker-compose.full.yml exec etcd \
  etcdctl get --prefix ""
```

### 2. 测试 API

```bash
# 测试 Gateway 健康检查
curl http://localhost:8080/api/v1/counter/knowpost/test

# 预期返回
# {"counts":{"fav":0,"like":0},"entityId":"test","entityType":"knowpost"}
```

### 3. 查看服务日志

```bash
# 查看 Gateway 日志
docker compose -f deploy/compose/docker-compose.full.yml logs -f gateway

# 查看所有应用日志
docker compose -f deploy/compose/docker-compose.full.yml logs -f gateway counter knowpost
```

## 多实例部署测试

测试服务发现和负载均衡：

```bash
# 启动第二个 counter 实例
docker compose -f deploy/compose/docker-compose.full.yml up -d --scale counter=2

# 查看 Etcd 中的注册信息
docker compose -f deploy/compose/docker-compose.full.yml exec etcd \
  etcdctl get --prefix "counter.rpc"

# 多次请求测试负载均衡
for i in {1..10}; do 
  curl -s http://localhost:8080/api/v1/counter/knowpost/$i | jq .entityId
done
```

## 常见问题

### Q1: 服务启动失败 "connection refused"

**原因**：依赖服务尚未就绪

**解决**：
1. 检查基础设施服务状态
2. 等待健康检查通过
3. 按顺序启动：基础设施 → 应用服务

### Q2: MySQL 连接失败

**原因**：数据库未初始化或迁移未执行

**解决**：
```bash
# 检查 MySQL 容器日志
docker compose -f deploy/compose/docker-compose.full.yml logs mysql

# 手动连接测试
docker compose -f deploy/compose/docker-compose.full.yml exec mysql \
  mysql -uroot -pZz123456 -e "SHOW DATABASES;"
```

### Q3: Etcd 服务发现不工作

**原因**：Etcd 未启动或网络配置错误

**解决**：
```bash
# 检查 Etcd 健康状态
docker compose -f deploy/compose/docker-compose.full.yml exec etcd \
  etcdctl endpoint health

# 检查网络连接
docker compose -f deploy/compose/docker-compose.full.yml exec gateway \
  ping etcd
```

### Q4: 服务无法互相访问

**原因**：Docker 网络配置问题

**解决**：
```bash
# 检查网络
docker network ls
docker network inspect zhiguang-go_default

# 重建网络
docker compose -f deploy/compose/docker-compose.full.yml down
docker compose -f deploy/compose/docker-compose.full.yml up -d
```

### Q5: 镜像构建失败

**原因**：依赖下载或编译错误

**解决**：
```bash
# 使用国内镜像加速
export GOPROXY=https://goproxy.cn,direct

# 清理并重新构建
docker compose -f deploy/compose/docker-compose.full.yml build --no-cache
```

## 端口映射

| 服务 | 容器端口 | 宿主端口 | 说明 |
|------|---------|---------|------|
| Gateway | 8080 | 8080 | 统一 HTTP 入口 |
| Agent | 8011 | 8011 | 知识助手入口 |
| MySQL | 3306 | 3306 | 数据库 |
| Redis | 6379 | 6379 | 缓存 |
| Kafka | 9092/29092 | 9092 | 消息队列 |
| ZooKeeper | 2181 | 2181 | Kafka 协调 |
| Etcd | 2379/2380 | 2379/2380 | 服务发现 |
| Elasticsearch | 9200 | 9200 | 搜索引擎 |
| Canal | 11111 | 11111 | Binlog 监听 |
| Adminer | 8080 | 8090 | 数据库管理 |

## 性能优化

### 1. 资源限制

在生产环境中添加资源限制：
```yaml
services:
  gateway:
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 512M
        reservations:
          cpus: '0.5'
          memory: 256M
```

### 2. 持久化数据备份

定期备份 volumes：
```bash
# 备份 MySQL 数据
docker run --rm -v zhiguang-go_zg_mysql_data:/data -v $(pwd):/backup \
  alpine tar czf /backup/mysql-backup-$(date +%Y%m%d).tar.gz /data

# 恢复
docker run --rm -v zhiguang-go_zg_mysql_data:/data -v $(pwd):/backup \
  alpine tar xzf /backup/mysql-backup-20260729.tar.gz -C /
```

### 3. 日志管理

配置日志轮转：
```yaml
services:
  gateway:
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
```

## 监控和维护

### 查看资源使用

```bash
# 查看容器资源使用
docker stats

# 查看特定服务
docker stats zg-gateway zg-counter
```

### 清理资源

```bash
# 停止并删除所有容器和网络
docker compose -f deploy/compose/docker-compose.full.yml down

# 同时删除 volumes（慎用！会删除数据）
docker compose -f deploy/compose/docker-compose.full.yml down -v

# 清理未使用的镜像
docker image prune -a
```

## 生产环境建议

1. **使用外部数据库**：不要在容器中运行生产数据库
2. **Etcd 集群**：部署 3 或 5 节点 Etcd 集群
3. **负载均衡**：在 Gateway 前加 Nginx/HAProxy
4. **监控告警**：集成 Prometheus + Grafana
5. **日志收集**：集成 ELK/Loki
6. **密钥管理**：使用 Docker Secrets 或外部密钥管理服务
7. **网络隔离**：使用自定义网络和网络策略

## 下一步

- 配置 Etcd 集群：[docs/etcd-cluster-setup.md](./etcd-cluster-setup.md)
- 配置监控告警：[docs/monitoring-setup.md](./monitoring-setup.md)
- 生产环境部署：[docs/production-deployment.md](./production-deployment.md)
