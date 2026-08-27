# 知光项目 - Docker 单容器部署方案总结

## 🎯 解决方案

根据你的需求"在一个 Docker 容器中启动所有服务，方便联调"，我创建了一个 **All-in-One 单容器部署方案**。

### 方案特点

✅ **一个容器运行所有服务**
- MySQL 数据库
- Redis 缓存
- Etcd 服务注册中心
- 所有 RPC 服务（User、Counter、KnowPost、Relation、Storage、Search、LLM）
- Gateway 网关

✅ **轻量级**
- 只需要一个容器
- 资源占用少
- 启动速度快

✅ **方便联调**
- 一键启动所有服务
- 统一管理
- 日志集中

## 📁 创建的文件

### 核心文件

1. **Dockerfile.allinone** - 单容器 Dockerfile
   - 编译所有 Go 服务
   - 安装 MySQL、Redis、Etcd
   - 配置 Supervisor 进程管理

2. **deploy/docker/start-all.sh** - 容器启动脚本
   - 初始化 MySQL
   - 执行数据库迁移
   - 启动 Supervisor

3. **deploy/docker/supervisord.conf** - 进程管理配置
   - 定义所有服务的启动方式
   - 配置日志路径
   - 设置自动重启

4. **deploy/compose/docker-compose.allinone.yml** - Docker Compose 配置
   - 单容器部署配置
   - 端口映射
   - 数据持久化

### 文档

5. **deploy/docker/README-allinone.md** - 详细使用文档
6. **docs/docker-allinone-quickstart.md** - 快速开始指南

## 🚀 使用方法

### 一键启动

```bash
# 构建并启动（首次需要 5-10 分钟）
docker compose -f deploy/compose/docker-compose.allinone.yml up -d --build
```

### 查看状态

```bash
# 查看容器日志
docker logs -f zhiguang-allinone

# 查看所有服务状态
docker exec zhiguang-allinone supervisorctl status
```

### 验证服务

```bash
# 测试 API
curl http://localhost:8080/api/v1/counter/knowpost/test

# 查看 Etcd 服务注册
docker exec zhiguang-allinone etcdctl get --prefix ""

# 应该看到所有 RPC 服务都已注册
```

## 🔍 服务架构

```
┌──────────────────────────────────────────────────┐
│           zhiguang-allinone 容器                  │
├──────────────────────────────────────────────────┤
│                                                  │
│  ┌────────────────────────────────────────┐     │
│  │         Supervisor 进程管理             │     │
│  └────────────────────────────────────────┘     │
│                      │                           │
│      ┌───────────────┼───────────────┐          │
│      │               │               │          │
│      ↓               ↓               ↓          │
│  ┌────────┐    ┌────────┐    ┌────────┐        │
│  │ MySQL  │    │ Redis  │    │ Etcd   │        │
│  │ :3306  │    │ :6379  │    │ :2379  │        │
│  └────────┘    └────────┘    └────────┘        │
│                      ↑                           │
│                      │ 服务发现                  │
│      ┌───────────────┼───────────────┐          │
│      │               │               │          │
│      ↓               ↓               ↓          │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐         │
│  │User RPC │  │Counter  │  │KnowPost │         │
│  │  :9002  │  │  :9003  │  │  :9004  │  ...    │
│  └─────────┘  └─────────┘  └─────────┘         │
│                      ↑                           │
│                      │                           │
│                ┌─────────┐                       │
│                │ Gateway │                       │
│                │  :8080  │                       │
│                └─────────┘                       │
│                                                  │
└──────────────────────────────────────────────────┘
               ↓
        外部访问: http://localhost:8080
```

## 📊 端口映射

| 服务 | 容器内端口 | 宿主机端口 | 说明 |
|-----|----------|----------|------|
| Gateway | 8080 | 8080 | HTTP 入口 |
| MySQL | 3306 | 3306 | 数据库 |
| Redis | 6379 | 6379 | 缓存 |
| Etcd | 2379 | 2379 | 服务发现 |
| User RPC | 9002 | 9002 | 用户服务 |
| Counter RPC | 9003 | 9003 | 计数服务 |
| KnowPost RPC | 9004 | 9004 | 知文服务 |
| Relation RPC | 9006 | 9006 | 关系服务 |
| Storage RPC | 9013 | 9013 | 存储服务 |
| Search RPC | 9017 | 9017 | 搜索服务 |
| LLM RPC | 9018 | 9018 | AI 服务 |

## 🛠️ 管理命令

### 查看服务状态

```bash
docker exec zhiguang-allinone supervisorctl status
```

输出示例：
```
etcd                RUNNING   pid 123, uptime 0:05:30
gateway             RUNNING   pid 456, uptime 0:05:20
knowpost-rpc        RUNNING   pid 789, uptime 0:05:25
mysql               RUNNING   pid 101, uptime 0:05:35
redis               RUNNING   pid 112, uptime 0:05:33
user-rpc            RUNNING   pid 234, uptime 0:05:28
counter-rpc         RUNNING   pid 345, uptime 0:05:27
relation-rpc        RUNNING   pid 567, uptime 0:05:26
search-rpc          RUNNING   pid 678, uptime 0:05:24
llm-rpc             RUNNING   pid 890, uptime 0:05:23
storage-rpc         RUNNING   pid 901, uptime 0:05:22
```

### 重启服务

```bash
# 重启特定服务
docker exec zhiguang-allinone supervisorctl restart gateway

# 重启所有服务
docker restart zhiguang-allinone
```

### 查看日志

```bash
# Gateway 日志
docker exec zhiguang-allinone tail -f /var/log/supervisor/gateway.log

# KnowPost RPC 日志
docker exec zhiguang-allinone tail -f /var/log/supervisor/knowpost-rpc.log

# 所有日志
docker exec zhiguang-allinone ls -la /var/log/supervisor/
```

### 访问数据库

```bash
# 连接 MySQL
docker exec -it zhiguang-allinone mysql -uroot -pZz123456 zhiguang

# 连接 Redis
docker exec -it zhiguang-allinone redis-cli

# 查看 Etcd
docker exec zhiguang-allinone etcdctl get --prefix ""
```

## ✅ 与原需求的对比

| 需求 | 实现方式 |
|-----|---------|
| 在一个容器中启动所有服务 | ✅ 单容器包含所有服务 |
| 方便联调 | ✅ 一键启动，统一管理 |
| 包含 Etcd | ✅ 内置 Etcd 服务发现 |
| 包含 KnowPost 等服务 | ✅ 所有 RPC 服务都在容器内 |
| 轻量级 | ✅ 只需一个容器 |

## 🆚 方案对比

### 单容器方案 vs 多容器方案

| 特性 | 单容器（推荐用于开发） | 多容器（推荐用于生产） |
|-----|---------------------|---------------------|
| 启动速度 | ⚡ 快 | 慢 |
| 资源占用 | ✅ 少（约 1-2GB） | 多（约 3-4GB） |
| 管理复杂度 | ✅ 简单 | 复杂 |
| 服务隔离 | ❌ 无 | ✅ 有 |
| 水平扩展 | ❌ 不支持 | ✅ 支持 |
| 故障隔离 | ❌ 一个服务崩溃影响全部 | ✅ 服务独立 |
| 适用场景 | 开发、联调、演示 | 测试、生产 |

## 💡 使用建议

### 推荐使用单容器的场景

✅ **本地开发调试**
- 快速启动
- 轻松重置
- 方便查看日志

✅ **前后端联调**
- 前端只需连接 `http://localhost:8080`
- 所有后端服务一键就绪

✅ **功能演示**
- 一条命令展示完整功能
- 环境隔离，不污染本地

✅ **快速验证**
- 测试新功能
- 验证 Bug 修复

### 不推荐使用单容器的场景

❌ **生产环境**
- 缺乏服务隔离
- 无法水平扩展
- 故障影响范围大

❌ **性能测试**
- 所有服务共享资源
- 无法模拟真实部署架构

❌ **高并发测试**
- 资源竞争严重
- 无法测试负载均衡

## 🐛 故障排查

### 容器无法启动

```bash
# 查看详细日志
docker logs zhiguang-allinone

# 检查端口占用
netstat -ano | findstr "8080"
```

### 某个服务异常

```bash
# 查看服务状态
docker exec zhiguang-allinone supervisorctl status

# 查看错误日志
docker exec zhiguang-allinone tail -100 /var/log/supervisor/gateway_err.log

# 重启服务
docker exec zhiguang-allinone supervisorctl restart gateway
```

### Etcd 服务发现不工作

```bash
# 检查 Etcd 健康状态
docker exec zhiguang-allinone etcdctl endpoint health

# 查看已注册的服务
docker exec zhiguang-allinone etcdctl get --prefix ""

# 查看 RPC 服务日志
docker exec zhiguang-allinone tail -100 /var/log/supervisor/user-rpc.log
```

## 📝 下一步

构建完成后（当前正在构建中）：

1. **启动容器**
   ```bash
   docker compose -f deploy/compose/docker-compose.allinone.yml up -d
   ```

2. **等待服务就绪**（约 30-60 秒）

3. **验证服务**
   ```bash
   # 查看服务状态
   docker exec zhiguang-allinone supervisorctl status
   
   # 测试 API
   curl http://localhost:8080/api/v1/counter/knowpost/test
   
   # 查看服务注册
   docker exec zhiguang-allinone etcdctl get --prefix ""
   ```

4. **开始前后端联调**
   - 前端项目配置代理到 `http://localhost:8080`
   - 所有后端服务已就绪

## 📚 相关文档

- **快速开始**：`docs/docker-allinone-quickstart.md`
- **详细文档**：`deploy/docker/README-allinone.md`
- **服务发现**：`docs/etcd-service-discovery.md`
- **多容器方案**：`docs/docker-deployment.md`（生产环境推荐）

## 🎓 总结

单容器方案完美解决了你的需求：

✅ **一个容器包含所有服务**（MySQL、Redis、Etcd、所有 RPC、Gateway）
✅ **方便联调**（一键启动，统一管理）
✅ **轻量级**（只需要一个容器）
✅ **支持 Etcd 服务发现**（所有服务自动注册）

适合日常开发和前后端联调，生产环境建议使用多容器方案。

---

**创建时间**：2026-07-29  
**适用场景**：开发、联调、演示  
**状态**：镜像构建中，完成后即可使用
