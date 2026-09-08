# 服务发现快速开始指南

本文档提供基于 Etcd 服务发现的快速启动步骤。

## 前置条件

1. **Etcd 服务运行**
   ```bash
   # 检查 Etcd 是否运行
   curl http://127.0.0.1:2379/version
   
   # 如果未运行，启动 Etcd (以默认配置为例)
   etcd
   ```

2. **其他依赖服务**
   - MySQL (3306)
   - Redis (6379)
   - Kafka (9092) - 可选，部分功能需要
   - Elasticsearch (9200) - 可选，搜索功能需要

## 启动服务

### 方式一：使用现有启动脚本（推荐）

已有的 `scripts/start-all.ps1` 脚本无需修改，可以直接使用：

```powershell
# Windows
scripts\start-all.ps1
```

所有服务会自动注册到 Etcd。

### 方式二：手动启动单个服务

从项目根目录启动：

```bash
# 启动 user-rpc
./bin/user-rpc.exe -f services/user/cmd/user/etc/user.yaml

# 启动 counter-rpc
./bin/counter-rpc.exe -f services/counter/rpc/etc/counter.yaml

# 启动 knowpost-rpc
./bin/knowpost-rpc.exe -f services/knowpost/rpc/etc/knowpost.yaml

# 启动 Gateway
./bin/gateway.exe -f services/gateway/etc/gateway.yaml
```

### 方式三：启动多实例测试高可用

以 counter-rpc 为例，启动第二个实例：

1. 复制并修改配置文件：
```yaml
# counter-rpc-2.yaml
Name: counter.rpc
ListenOn: 0.0.0.0:9023  # 使用不同端口

Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: counter.rpc  # 使用相同的 Key

# ... 其他配置保持不变
```

2. 启动第二个实例：
```bash
./bin/counter-rpc.exe -f counter-rpc-2.yaml
```

两个实例会自动注册到同一个服务名下，Gateway 会在它们之间负载均衡。

## 验证服务注册

### 查看已注册的服务

```bash
# 查看 user.rpc 的注册信息
curl -s http://127.0.0.1:2379/v3/kv/range -X POST \
  -d '{"key":"dXNlci5ycGM=","range_end":"dXNlci5ycGN9"}' \
  -H "Content-Type: application/json"

# 查看 counter.rpc 的注册信息
curl -s http://127.0.0.1:2379/v3/kv/range -X POST \
  -d '{"key":"Y291bnRlci5ycGM=","range_end":"Y291bnRlci5ycGN9"}' \
  -H "Content-Type: application/json"
```

注：key 是 base64 编码的服务名。

### 测试服务调用

```bash
# 测试 Gateway 能否发现并调用后端服务
curl http://127.0.0.1:8080/api/v1/counter/knowpost/123
```

预期返回：
```json
{"counts":{"fav":0,"like":0},"entityId":"123","entityType":"knowpost"}
```

## 高可用测试

### 测试实例故障切换

1. 启动两个 counter-rpc 实例（参考方式三）
2. 测试服务正常：
   ```bash
   curl http://127.0.0.1:8080/api/v1/counter/knowpost/test
   ```
3. 停止第一个实例：
   ```bash
   kill <第一个实例的PID>
   ```
4. 再次测试，服务应该继续正常工作
5. 查看 Etcd，停止的实例应该已自动注销

## 常见问题

### Q1: 服务启动报错 "failed to connect to etcd"

**原因**：Etcd 服务未启动或地址配置错误

**解决**：
- 检查 Etcd 是否运行：`curl http://127.0.0.1:2379/version`
- 检查配置文件中的 Etcd 地址

### Q2: Gateway 调用 RPC 服务超时

**原因**：RPC 服务未启动或未成功注册

**解决**：
- 检查 RPC 服务进程是否运行
- 查看 RPC 服务日志，确认是否成功注册到 Etcd
- 使用 Etcd API 查询服务是否已注册

### Q3: 多实例启动失败 "bind: address already in use"

**原因**：端口冲突

**解决**：
- 每个实例必须使用不同的 `ListenOn` 端口
- 每个实例必须使用不同的 `Prometheus.Port`
- `Etcd.Key` 保持相同（这样才能负载均衡）

### Q4: 服务停止后 Etcd 中的注册信息未清理

**原因**：服务异常退出（如 kill -9）可能导致清理延迟

**解决**：
- 正常情况下，租约过期后会自动清理（默认约 30-60 秒）
- 如需立即清理，可以重启 Etcd 或手动删除 key

### Q5: 生产环境配置建议

**Etcd 集群**：
```yaml
Etcd:
  Hosts:
  - etcd-node1:2379
  - etcd-node2:2379
  - etcd-node3:2379
  Key: user.rpc
```

**服务监听地址**：
- 开发环境：`0.0.0.0:端口` (绑定所有网卡)
- 生产环境：考虑使用具体 IP 或仅内网 IP

**监控**：
- 监控 Etcd 连接状态
- 监控服务注册/注销事件
- 配置告警规则

## 下一步

- 阅读完整文档：[docs/etcd-service-discovery.md](./etcd-service-discovery.md)
- 了解高可用演进规划：继续实施 Redis Sentinel、Kafka 集群等改进
- 配置生产环境的 Etcd 集群
- 设置监控和告警
