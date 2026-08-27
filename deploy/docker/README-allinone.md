# All-in-One 单容器部署指南

## 概述

这是一个单容器部署方案，所有服务（MySQL、Redis、Etcd、各个 RPC 服务、Gateway）都运行在一个 Docker 容器内。

**优势**：
- ✅ 轻量级，只需要一个容器
- ✅ 启动快速，适合开发和联调
- ✅ 资源占用少
- ✅ 配置简单，易于管理

**适用场景**：
- 本地开发调试
- 前后端联调
- 快速演示
- 功能测试

**不适用于**：
- 生产环境（生产环境请使用多容器方案）
- 高并发压测
- 多实例部署

## 快速开始

### 1. 构建并启动

```bash
# 构建并启动单容器
docker compose -f deploy/compose/docker-compose.allinone.yml up -d --build
```

首次构建需要 5-10 分钟，后续启动只需几秒。

### 2. 查看启动日志

```bash
# 查看容器日志
docker logs -f zhiguang-allinone

# 查看特定服务日志
docker exec zhiguang-allinone tail -f /var/log/supervisor/gateway.log
docker exec zhiguang-allinone tail -f /var/log/supervisor/knowpost-rpc.log
```

### 3. 验证服务

```bash
# 测试 API
curl http://localhost:8080/api/v1/counter/knowpost/test

# 预期返回
{"counts":{"fav":0,"like":0},"entityId":"test","entityType":"knowpost"}
```

### 4. 查看服务状态

```bash
# 查看所有服务状态
docker exec zhiguang-allinone supervisorctl status

# 输出示例：
# etcd                             RUNNING   pid 123, uptime 0:05:30
# gateway                          RUNNING   pid 456, uptime 0:05:20
# knowpost-rpc                     RUNNING   pid 789, uptime 0:05:25
# mysql                            RUNNING   pid 101, uptime 0:05:35
# redis                            RUNNING   pid 112, uptime 0:05:33
# user-rpc                         RUNNING   pid 234, uptime 0:05:28
# ...
```

## 管理服务

### 重启服务

```bash
# 重启所有服务
docker restart zhiguang-allinone

# 重启特定服务
docker exec zhiguang-allinone supervisorctl restart gateway
docker exec zhiguang-allinone supervisorctl restart knowpost-rpc
```

### 停止/启动特定服务

```bash
# 停止服务
docker exec zhiguang-allinone supervisorctl stop gateway

# 启动服务
docker exec zhiguang-allinone supervisorctl start gateway
```

### 查看日志

```bash
# 实时查看 Gateway 日志
docker exec zhiguang-allinone tail -f /var/log/supervisor/gateway.log

# 查看最近 100 行
docker exec zhiguang-allinone tail -100 /var/log/supervisor/knowpost-rpc.log

# 查看所有服务日志列表
docker exec zhiguang-allinone ls -la /var/log/supervisor/
```

## 访问内部服务

### 数据库

```bash
# 连接 MySQL
docker exec -it zhiguang-allinone mysql -uroot -pZz123456 zhiguang

# 执行 SQL
docker exec zhiguang-allinone mysql -uroot -pZz123456 zhiguang -e "SHOW TABLES;"
```

### Redis

```bash
# 连接 Redis
docker exec -it zhiguang-allinone redis-cli

# 查看所有 key
docker exec zhiguang-allinone redis-cli KEYS "*"
```

### Etcd

```bash
# 查看服务注册信息
docker exec zhiguang-allinone etcdctl get --prefix ""

# 查看特定服务
docker exec zhiguang-allinone etcdctl get --prefix "counter.rpc"
```

## 验证 Etcd 服务发现

### 1. 查看注册的服务

```bash
docker exec zhiguang-allinone etcdctl get --prefix "" --keys-only
```

应该看到类似：
```
user.rpc/7587896510965527045
counter.rpc/7587896510965527046
knowpost.rpc/7587896510965527047
...
```

### 2. 查看具体服务实例

```bash
# 查看 counter.rpc 的注册信息
docker exec zhiguang-allinone etcdctl get --prefix "counter.rpc"

# 输出示例：
# counter.rpc/7587896510965527046
# 127.0.0.1:9003
```

### 3. 测试服务调用

```bash
# Gateway 通过 Etcd 发现并调用 Counter RPC
curl http://localhost:8080/api/v1/counter/knowpost/123

# 成功返回说明服务发现正常工作
```

## 前端联调

前端项目可以直接配置：

```javascript
// vite.config.ts 或 next.config.js
export default {
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  }
}
```

## 数据管理

### 备份数据

```bash
# 备份 MySQL
docker exec zhiguang-allinone mysqldump -uroot -pZz123456 zhiguang > backup-$(date +%Y%m%d).sql

# 备份 Redis
docker exec zhiguang-allinone redis-cli SAVE
docker cp zhiguang-allinone:/var/lib/redis/dump.rdb ./redis-backup.rdb
```

### 恢复数据

```bash
# 恢复 MySQL
docker exec -i zhiguang-allinone mysql -uroot -pZz123456 zhiguang < backup-20260729.sql

# 恢复 Redis
docker cp redis-backup.rdb zhiguang-allinone:/var/lib/redis/dump.rdb
docker exec zhiguang-allinone supervisorctl restart redis
```

### 清空数据重新开始

```bash
# 停止并删除容器（保留镜像）
docker compose -f deploy/compose/docker-compose.allinone.yml down

# 删除数据卷（慎用！）
docker compose -f deploy/compose/docker-compose.allinone.yml down -v

# 重新启动
docker compose -f deploy/compose/docker-compose.allinone.yml up -d --build
```

## 常见问题

### Q1: 容器启动失败

```bash
# 查看详细日志
docker logs zhiguang-allinone

# 检查是否端口冲突
netstat -ano | findstr "8080"
netstat -ano | findstr "3306"
netstat -ano | findstr "6379"
```

### Q2: 某个服务一直重启

```bash
# 查看服务状态
docker exec zhiguang-allinone supervisorctl status

# 查看该服务的错误日志
docker exec zhiguang-allinone tail -100 /var/log/supervisor/gateway_err.log
```

### Q3: 数据库连接失败

```bash
# 检查 MySQL 是否运行
docker exec zhiguang-allinone supervisorctl status mysql

# 测试连接
docker exec zhiguang-allinone mysql -uroot -pZz123456 -e "SELECT 1"

# 查看 MySQL 日志
docker exec zhiguang-allinone tail -100 /var/log/supervisor/mysql_err.log
```

### Q4: Etcd 服务发现不工作

```bash
# 检查 Etcd 是否运行
docker exec zhiguang-allinone supervisorctl status etcd

# 测试 Etcd 连接
docker exec zhiguang-allinone etcdctl endpoint health

# 查看 Gateway 日志，看是否有连接 Etcd 的错误
docker exec zhiguang-allinone tail -100 /var/log/supervisor/gateway.log
```

### Q5: 如何修改配置

配置文件都在容器内的 `/app/services/` 目录，修改后需要重启服务：

```bash
# 进入容器
docker exec -it zhiguang-allinone sh

# 修改配置（容器内）
vi /app/services/gateway/etc/gateway.yaml

# 重启服务
supervisorctl restart gateway
```

或者，修改宿主机的配置文件后重新构建：

```bash
# 编辑配置文件
vim services/gateway/etc/gateway.yaml

# 重新构建并启动
docker compose -f deploy/compose/docker-compose.allinone.yml up -d --build
```

## 性能说明

单容器方案的性能限制：

- **并发能力**：适合开发和小规模测试，不适合压测
- **资源限制**：所有服务共享容器资源
- **隔离性**：服务间没有隔离，一个服务崩溃可能影响其他服务

如果需要：
- 高并发测试 → 使用多容器方案 `docker-compose.test.yml`
- 生产环境 → 使用完整方案 `docker-compose.full.yml` 或 Kubernetes

## 停止和清理

```bash
# 停止容器（保留数据）
docker compose -f deploy/compose/docker-compose.allinone.yml down

# 停止并删除所有数据
docker compose -f deploy/compose/docker-compose.allinone.yml down -v

# 查看资源占用
docker stats zhiguang-allinone
```

## 总结

单容器方案优点：
- ✅ 一键启动，方便快速
- ✅ 资源占用少
- ✅ 适合本地开发和前后端联调

适用场景：
- 开发调试
- 前端联调
- 功能演示

不适用：
- 生产环境
- 性能测试
- 多实例部署

---

**推荐使用场景**：日常开发、前后端联调  
**不推荐场景**：生产部署、压力测试
