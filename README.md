# 知光平台 Go 重构版（`zhiguang-go`）

从 Java / Spring Boot 单体重构而来的 Go 微服务版本。当前后端已经完成两层收敛：

- 业务服务按边界拆分为 `rpc` / `worker` / merged service
- 对外 HTTP 入口收敛为统一 `gateway`

`agent` 目前保持独立入口，不纳入统一网关。

## 当前架构

- 统一外部入口：`services/gateway` `:8080`
- 独立外部入口：`services/agent/cmd/agent` `:8011`
- 核心内部 RPC：
  - `user-rpc` `:9002`
  - `counter-rpc` `:9003`
  - `knowpost-rpc` `:9004`
  - `relation-rpc` `:9006`
  - `storage-rpc` `:9013`
  - `search-rpc` `:9017`
  - `llm-rpc` `:9018`
- merged service：
  - `counter/cmd/counter`
  - `knowpost/cmd/knowpost`
  - `relation/cmd/relation`
  - `search/cmd/search`
  - `llm/cmd/llm`
  - `agent/cmd/agent`

说明：
- `counter/knowpost/relation/search/llm` 的 merged service 默认 `DisableAPI: true`
- 即这些进程继续承载内部 `rpc` / `worker` 能力，但不再作为正式外部 HTTP 入口
- 外部 HTTP 请求统一通过 `gateway`

## 目录结构

```text
zhiguang-go/
├── docs/                  # 项目分析、阶段文档、实施记录
├── proto/                 # protobuf 定义
├── pkg/                   # 通用基础库
├── common/                # 业务公共代码
├── deploy/                # docker / compose / 部署辅助
├── scripts/               # 启动、停止、迁移、生成脚本
└── services/
    ├── gateway/           # 统一对外 HTTP 网关
    ├── user/rpc/          # 用户 + 认证
    ├── storage/rpc/       # OSS 预签名
    ├── counter/           # 计数 rpc + aggregator + reconciler + merged
    ├── knowpost/          # 知文 rpc + merged
    ├── relation/          # 关系 rpc + syncer + merged
    ├── search/            # 搜索 rpc + indexer + merged
    ├── llm/               # AI 描述 / RAG rpc + ragindexer + merged
    ├── agent/             # 个人知识助手，当前独立入口
    └── outbox/gc/         # outbox 清理 worker
```

## 快速启动

### 1. 启动依赖

```bash
docker compose -f deploy/compose/docker-compose.dev.yml up -d
```

### 2. 执行迁移

```bash
./scripts/migrate.sh up
```

Windows:

```powershell
scripts\migrate.bat up
```

### 3. 一键启动后端

Windows:

```powershell
scripts\start-all.ps1
```

当前主启动清单会拉起：

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

### 4. 前端开发代理

前端 `vite.config.ts` 默认代理到：

```text
http://localhost:8080
```

因此新网关已经和当前前端开发方式对齐。

## 关键入口

- 统一业务入口：`http://localhost:8080/api/v1/*`
- Agent 独立入口：`http://localhost:8011/api/v1/agent/*`

## 参考文档

- 总体分析：[docs/project-analysis.md](/F:/zhiguang_be/zhiguang-go/docs/project-analysis.md)
- 服务合并记录：[docs/phase7-merge-service.md](/F:/zhiguang_be/zhiguang-go/docs/phase7-merge-service.md)
- 本阶段网关收缩总结：`docs/phase8-gateway-consolidation.md`
