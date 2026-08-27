# Docker 环境部署总结

## 已完成的工作

### 1. Dockerfile 创建
- ✅ 多阶段构建，优化镜像大小
- ✅ 编译所有核心服务（gateway、user、counter、knowpost、relation、search、llm、agent）
- ✅ 最小化运行镜像（基于 alpine）

### 2. Docker Compose 配置
- ✅ **完整部署配置**：`docker-compose.full.yml`（所有服务 + 基础设施）
- ✅ **测试部署配置**：`docker-compose.test.yml`（核心服务快速测试）
- ✅ **开发依赖配置**：`docker-compose.dev.yml`（已存在，仅基础设施）

### 3. 服务配置文件
为每个服务创建了 Docker 专用配置（`*-docker.yaml`），主要修改：
- Etcd 地址：`127.0.0.1:2379` → `etcd:2379`
- MySQL 地址：`127.0.0.1:3306` → `mysql:3306`
- Redis 地址：`127.0.0.1:6379` → `redis:6379`
- Kafka 地址：`127.0.0.1:9092` → `kafka:29092`

已创建的配置文件：
- `services/gateway/etc/gateway-docker.yaml`
- `services/counter/cmd/counter/etc/counter-docker.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml`
- `services/relation/cmd/relation/etc/relation-docker.yaml`
- `services/search/cmd/search/etc/search-docker.yaml`
- `services/llm/cmd/llm/etc/llm-docker.yaml`
- `services/user/cmd/user/etc/user-docker.yaml`
- `services/agent/cmd/agent/etc/agent-docker.yaml`

### 4. 辅助脚本
- ✅ `scripts/docker-start.sh`（Linux/Mac 启动脚本）
- ✅ `scripts/docker-start.ps1`（Windows PowerShell 启动脚本）

### 5. 文档
- ✅ `docs/docker-deployment.md`（完整部署文档）
- ✅ `deploy/compose/README.md`（快速使用指南）
- ✅ `.dockerignore`（优化构建速度）

## 架构说明

### 基础设施层
```
┌─────────────────────────────────────────────┐
│          基础设施容器                          │
├─────────────────────────────────────────────┤
│ MySQL (3306)     │ Redis (6379)             │
│ Etcd (2379)      │ Kafka (9092)             │
│ ZooKeeper (2181) │ Elasticsearch (9200)     │
│ Canal (11111)    │ Adminer (8090)           │
└─────────────────────────────────────────────┘
```

### 应用服务层
```
┌─────────────────────────────────────────────┐
│           应用服务容器                         │
├─────────────────────────────────────────────┤
│                                              │
│  Gateway (8080) ─────┐                      │
│                      │                      │
│                      ├─→ User-Storage RPC   │
│                      ├─→ Counter RPC        │
│                      ├─→ KnowPost RPC       │
│                      ├─→ Relation RPC       │
│                      ├─→ Search RPC         │
│                      └─→ LLM RPC            │
│                                              │
│  Agent (8011) ───────→ 独立入口              │
│                                              │
└─────────────────────────────────────────────┘
                     ↓↑
              通过 Etcd 服务发现
```

### 服务发现流程
```
1. RPC 服务启动
   └→ 注册到 Etcd (例如: counter.rpc/xxx → 10.0.0.5:9003)

2. Gateway 启动
   └→ 连接 Etcd
   └→ Watch 所有需要的服务 key

3. 用户请求到达
   └→ Gateway 查询 Etcd 获取可用实例列表
   └→ 负载均衡选择一个实例
   └→ 调用选中的实例

4. RPC 实例下线
   └→ Etcd 租约过期自动注销
   └→ Gateway Watch 到变化，更新实例列表
   └→ 后续请求不再路由到下线实例
```

## 使用方式

### 快速启动（测试模式）

```bash
# Windows
.\scripts\docker-start.ps1

# Linux/Mac
bash scripts/docker-start.sh

# 或手动启动
docker compose -f deploy/compose/docker-compose.test.yml up -d --build
```

### 验证部署

```bash
# 1. 检查服务状态
docker compose -f deploy/compose/docker-compose.test.yml ps

# 2. 查看 Etcd 服务注册
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix ""

# 3. 测试 API
curl http://localhost:8080/api/v1/counter/knowpost/test
```

### 多实例测试

```bash
# 启动 2 个 Counter 实例
docker compose -f deploy/compose/docker-compose.test.yml up -d --scale counter=2

# 验证负载均衡
for i in {1..10}; do 
  curl -s http://localhost:8080/api/v1/counter/knowpost/$i | jq .entityId
done
```

## 与本地开发的对比

| 项目 | 本地开发 | Docker 环境 |
|-----|---------|-----------|
| 服务地址 | 127.0.0.1 | 容器名（etcd, mysql, redis） |
| 依赖管理 | 手动启动各个服务 | Compose 自动管理依赖关系 |
| 多实例 | 手动修改端口 | `--scale` 一键扩展 |
| 网络隔离 | 共享宿主机网络 | Docker 内部网络隔离 |
| 数据持久化 | 本地文件系统 | Docker Volumes |
| 环境清理 | 手动停止进程 | `docker compose down` |

## 优势

1. **一键部署**：无需手动启动多个服务
2. **环境隔离**：不污染本地环境
3. **依赖管理**：自动处理服务启动顺序和依赖关系
4. **易于重置**：`docker compose down -v` 清理所有数据
5. **多实例测试**：轻松测试负载均衡和高可用
6. **版本一致**：确保团队使用相同的服务版本

## 注意事项

1. **首次构建**：需要下载 Go 镜像和依赖，约需 5-10 分钟
2. **资源占用**：完整部署需要约 4GB 内存
3. **端口冲突**：确保本地没有占用相同端口的服务
4. **数据持久化**：volumes 数据在 `docker compose down` 后保留，需要 `-v` 参数才会删除
5. **配置修改**：修改配置后需要重启容器

## 故障排查

### 构建失败
```bash
# 使用国内镜像加速
export GOPROXY=https://goproxy.cn,direct

# 清理缓存重新构建
docker compose -f deploy/compose/docker-compose.test.yml build --no-cache
```

### 服务无法互相访问
```bash
# 检查 Docker 网络
docker network ls
docker network inspect compose_default

# 从容器内测试连接
docker compose -f deploy/compose/docker-compose.test.yml exec gateway ping etcd
```

### Etcd 服务发现不工作
```bash
# 检查 Etcd 健康状态
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl endpoint health

# 查看注册信息
docker compose -f deploy/compose/docker-compose.test.yml exec etcd etcdctl get --prefix ""

# 查看服务日志
docker compose -f deploy/compose/docker-compose.test.yml logs gateway counter
```

## 下一步

1. **等待构建完成**：首次构建需要时间，请耐心等待
2. **验证服务发现**：按照上述步骤测试服务注册和调用
3. **测试多实例**：使用 `--scale` 测试负载均衡
4. **测试高可用**：停止一个实例，验证流量自动切换
5. **完整部署**：测试通过后，使用 `docker-compose.full.yml` 部署所有服务

## 当前构建状态

Docker 镜像正在构建中，主要步骤：
1. ⏳ 下载 Go 1.23 基础镜像（约 74MB）
2. 等待：编译所有 Go 服务
3. 等待：创建最终运行镜像

构建完成后会自动启动所有服务。

---

**创建时间**：2026-07-29  
**文档版本**：1.0
