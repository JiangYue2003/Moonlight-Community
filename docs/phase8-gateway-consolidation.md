# Phase 8 - 统一网关与入口收缩总结

## 目标

本阶段目标不是继续拆更多服务，而是把已经拆散的外部 HTTP 入口重新收束：

- 对前端提供统一外部入口
- 把业务服务尽量收回到内部 `gRPC` 通信
- 保留 `agent` 的独立入口，不强行并入统一网关
- 在不破坏既有业务链路的前提下，降低“一个项目要记很多 HTTP 端口”的复杂度

## 本阶段完成内容

### 1. 新增统一 Gateway

新增：

- `services/gateway`

职责：

- 作为非 agent 业务统一外部入口
- 对外监听 `:8080`
- 保持前端既有 `/api/v1/*` 路径不变
- 内部将请求转发到对应 `rpc`

当前纳入 Gateway 的外部链路：

- `auth`
- `profile`
- `storage`
- `knowpost`
- `relation`
- `counter`
- `search`
- `llm`

明确不纳入：

- `agent`

### 2. 新增内部 RPC 服务

为支持网关统一转发，本阶段补齐了几个之前缺失的内部 RPC：

- `storage-rpc` `:9013`
- `search-rpc` `:9017`
- `llm-rpc` `:9018`

对应实现已包含：

- proto
- pb / grpc 生成物
- client
- server
- logic
- 基础测试

### 3. merged service 关闭旧外部 API

此前若干 merged service 虽然已经完成“业务合并”，但仍保留旧 HTTP 暴露能力，导致实际是双入口并存。

本阶段为以下 merged service 增加 `DisableAPI`：

- `counter/cmd/counter`
- `knowpost/cmd/knowpost`
- `relation/cmd/relation`
- `search/cmd/search`
- `llm/cmd/llm`

默认值已改为：

```yaml
DisableAPI: true
```

含义：

- 这些进程继续承载内部 `rpc` / `worker`
- 但不再作为正式外部 HTTP 入口

### 4. SSE 流式链路已在 Gateway 中桥接

本阶段没有回退流式能力，而是在网关层保留了 SSE 桥接：

- `GET /api/v1/llm/qa/stream`
- `GET /api/v1/knowposts/{id}/qa/stream`

这样前端仍然可以按既有方式消费流式返回，而内部则由 `llm-rpc` 提供流式源。

### 5. 启动脚本已切换

`scripts/start-all.ps1` 已切换为新的入口拓扑。

现在主启动会拉起：

- `gateway`
- `user-rpc`
- `storage-rpc`
- `search-rpc`
- `llm-rpc`
- `counter/cmd/counter`
- `knowpost/cmd/knowpost`
- `relation/cmd/relation`
- `search/cmd/search`
- `llm/cmd/llm`
- `agent/cmd/agent`

不再把下列目录作为正式外部入口启动：

- `auth-api`
- `profile-api`
- `storage-api`
- merged service 内部旧 API

## 当前结构结果

### 对前端

前端现在主要只需要记住两个入口：

- `http://localhost:8080`：主业务入口
- `http://localhost:8011`：Agent 独立入口

其中前端开发代理默认已经指向 `http://localhost:8080`，因此不需要再改 Vite 默认代理。

### 对后端

后端当前结构变为：

- 外部请求统一进入 `gateway`
- gateway 负责鉴权注入、参数绑定、响应映射、SSE 桥接
- 核心业务逻辑回到内部 `rpc`
- 异步索引、聚合、补偿继续由 worker / merged service 承载

## 本阶段验证结果

已完成验证：

- `storage-rpc/search-rpc/llm-rpc` logic 测试通过
- `go build ./services/gateway` 通过
- `go test ./services/gateway/...` 通过
- `gateway + rpc + merged service` 关键入口构建通过

## 仍保留的设计取舍

### 1. 旧 `api` 目录没有删除

原因：

- 作为回滚面保留
- 一些实现仍可被 merged service 或 gateway 参考
- 现在的目标是“切换正式入口”，不是立即做代码清仓

### 2. Agent 暂不纳入 Gateway

这是明确的边界决策：

- Agent 后续可能做独立鉴权
- Agent 使用的资源与主业务链路差异较大
- 当前先保证主业务入口收敛，不让 Agent 影响这条改造链路

## 后续建议

### P0

- 为 Gateway 补更多接口级测试
- 真实联调验证 `gateway:8080` 下的登录、发帖、搜索、SSE
- 更新更多与当前启动方式相关的文档

### P1

- 进一步明确哪些旧 `api` 目录只保留为历史回滚代码
- 给 Gateway 增加更系统的监控指标
- 评估是否为 Agent 做独立鉴权与独立网关

## 结论

本阶段的核心结果不是“又多了一个服务”，而是：

- 对外入口从多端口分散暴露，收敛为统一 `gateway`
- 内部能力补齐到 `gRPC` 层
- merged service 从“既跑内部又跑旧外部 HTTP”的混合状态，切换为“内部能力优先”
- `agent` 保持独立，不影响主业务入口收敛
