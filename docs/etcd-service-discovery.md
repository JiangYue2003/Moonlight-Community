# Etcd 服务发现改造总结

## 改造概述

将知光项目从硬编码的点对点 RPC 调用改造为基于 Etcd 的服务发现机制，实现服务的动态注册与发现，支持水平扩展和高可用部署。

## 改造范围

### 1. RPC 服务端配置修改

为所有 RPC 服务添加 Etcd 注册配置，修改的服务包括：

- `services/user/rpc/etc/user.yaml`
- `services/counter/rpc/etc/counter.yaml`
- `services/knowpost/rpc/etc/knowpost.yaml`
- `services/relation/rpc/etc/relation.yaml`
- `services/search/rpc/etc/search.yaml` (已有配置)
- `services/storage/rpc/etc/storage.yaml` (已有配置)
- `services/llm/rpc/etc/llm.yaml` (已有配置)

**配置示例**：
```yaml
Name: user.rpc
ListenOn: 0.0.0.0:9002

# Etcd 服务注册配置
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: user.rpc

Mode: dev
```

### 2. RPC 客户端配置修改

将所有 RPC 客户端从硬编码的 `Endpoints` 改为 Etcd 服务发现模式：

**修改前**：
```yaml
CounterRpc:
  Endpoints:
  - 127.0.0.1:9003
  NonBlock: true
```

**修改后**：
```yaml
CounterRpc:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: counter.rpc
  NonBlock: true
```

修改的配置文件包括：
- `services/gateway/etc/gateway.yaml` (所有 RPC 客户端)
- `services/knowpost/rpc/etc/knowpost.yaml` (CounterRpc 和 UserCounterRpc)
- `services/relation/rpc/etc/relation.yaml` (UserRpc)
- `services/search/rpc/etc/search.yaml` (CounterRpc)
- `services/counter/cmd/counter/etc/counter.yaml` (merged service)
- `services/knowpost/cmd/knowpost/etc/knowpost.yaml` (merged service)
- `services/relation/cmd/relation/etc/relation.yaml` (merged service)
- `services/search/cmd/search/etc/search.yaml` (merged service)
- `services/llm/cmd/llm/etc/llm.yaml` (merged service)
- `services/user/cmd/user/etc/user.yaml` (merged service)
- `services/agent/cmd/agent/etc/agent.yaml`

## 验证测试

### 1. 服务注册验证

启动服务后，通过 Etcd API 查询注册信息：

```bash
curl -s http://127.0.0.1:2379/v3/kv/range -X POST \
  -d '{"key":"dXNlci5ycGM=","range_end":"dXNlci5ycGN9"}' \
  -H "Content-Type: application/json"
```

**结果**：成功在 Etcd 中找到服务注册信息，value 为 `10.131.181.181:9002` (服务实际地址)

### 2. 服务发现验证

启动 Gateway 后，测试 RPC 调用：

```bash
curl http://127.0.0.1:8080/api/v1/counter/knowpost/123
```

**结果**：Gateway 成功通过 Etcd 发现 counter-rpc 服务并完成调用，返回正确响应。

### 3. 多实例负载均衡验证

启动两个 counter-rpc 实例（不同端口 9003 和 9023），Etcd 中注册了两个实例：
- `10.131.181.181:9003`
- `10.131.181.181:9023`

连续 10 次请求全部成功，验证了 go-zero 的负载均衡机制正常工作。

### 4. 高可用验证

停止第一个 counter-rpc 实例后：
- 服务继续正常响应请求
- Etcd 中自动清理了停止实例的注册信息（从 count:2 变为 count:1）
- 所有请求自动路由到剩余的健康实例

**结论**：验证了服务发现的高可用特性，单实例故障不影响整体服务。

## 技术要点

### 1. Etcd 服务注册机制

- go-zero 使用 **租约（Lease）** 机制实现服务注册
- 服务启动时自动在 Etcd 中注册，定期续约
- 服务停止或异常退出时，租约过期后自动注销
- 注册 key 格式：`{ServiceKey}/{LeaseID}`
- 注册 value：服务实际监听地址（IP:Port）

### 2. 服务发现与负载均衡

- go-zero 客户端通过 **Watch** 机制监听 Etcd 中服务列表变化
- 内置负载均衡策略（轮询、随机等）
- 服务实例上下线时自动更新可用服务列表
- 不健康实例自动摘除，恢复后自动加回

### 3. 配置兼容性

- go-zero 同时支持 `Endpoints` 和 `Etcd` 两种配置方式
- 如果同时配置，`Etcd` 优先级更高
- 无需修改代码，只需修改配置文件即可完成改造

## 收益

### 1. 水平扩展能力

- 任意 RPC 服务都可以启动多个实例
- 新实例启动后自动加入服务池，无需修改配置
- 支持动态扩容和缩容

### 2. 高可用

- 单个实例故障不影响整体服务
- 故障实例自动摘除，流量自动转移到健康实例
- 实例恢复后自动重新加入

### 3. 运维便利性

- 无需维护服务地址清单
- 服务迁移（IP/端口变更）对客户端透明
- 支持灰度发布和蓝绿部署

### 4. 多机房部署准备

- 为后续多机房、异地多活部署打下基础
- 可以基于 Etcd 实现跨机房的服务发现

## 后续优化方向

### 1. Etcd 集群化

当前使用单节点 Etcd（`127.0.0.1:2379`），生产环境需要部署 Etcd 集群（建议 3 或 5 节点）：

```yaml
Etcd:
  Hosts:
  - 192.168.1.11:2379
  - 192.168.1.12:2379
  - 192.168.1.13:2379
  Key: user.rpc
```

### 2. 健康检查增强

- 配置自定义健康检查逻辑
- 设置合理的租约 TTL（当前使用 go-zero 默认值）
- 监控服务注册状态

### 3. 负载均衡策略优化

- 根据实际业务选择合适的负载均衡算法
- 考虑实例负载情况（CPU、内存、连接数）的加权负载均衡
- 实现同机房优先、跨机房容灾的智能路由

### 4. 配置中心集成

- 利用 Etcd 实现配置中心功能
- 支持配置动态更新，无需重启服务

### 5. 监控与告警

- 监控 Etcd 连接状态
- 监控服务注册/注销事件
- 监控服务实例健康度

## 注意事项

1. **Etcd 依赖**：所有服务启动前必须确保 Etcd 服务可用
2. **网络策略**：确保服务实例之间、服务与 Etcd 之间网络互通
3. **时钟同步**：多实例部署时建议配置 NTP，保证时钟一致
4. **租约时间**：生产环境需根据网络状况调整租约 TTL
5. **Etcd 存储**：Etcd 不适合存储大量数据，仅用于服务注册和轻量配置

## 相关文档

- [go-zero RPC 服务发现文档](https://go-zero.dev/docs/tutorials/service/service-discovery)
- [Etcd 官方文档](https://etcd.io/docs/)
- [项目高可用演进规划](./high-availability-roadmap.md)
