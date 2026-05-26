# 知光平台 (ZhiGuang) - Go 重构项目参考文档

> 文档说明：本文档描述的是当前 Go 重构项目的实际运行形态，已经包含服务合并与统一网关收缩结果。  
> 原 Java 技术栈说明保留用于对比。  
> 最后更新：2026-05-25

## 一、项目总览

### 1.1 项目定位

知光平台是一个知识获取与分享社区，支持：

- 发布知文（图文 / Markdown）
- 点赞、收藏、关注
- 首页 Feed 与详情页展示
- OSS 直传
- AI 生成摘要
- 基于知文内容的 RAG 问答
- 独立的个人知识助手 Agent

### 1.2 Go 重构技术栈

| 组件 | Go 重构方案 | 对应 Java 方案 |
|------|------------|---------------|
| 语言 | Go 1.25.x | Java 21 |
| 框架 | go-zero + Gin Gateway | Spring Boot 3.2.4 + Spring Security |
| ORM | go-zero sqlx（手写 SQL） | MyBatis |
| 数据库 | MySQL 8.0 | MySQL 8.0 |
| 缓存 | go-redis/v9 + ristretto | Redis + Caffeine |
| 消息队列 | segmentio/kafka-go | Spring Kafka |
| 搜索引擎 | Elasticsearch 9.x | Elasticsearch 9.x |
| 对象存储 | 阿里云 OSS | 阿里云 OSS SDK Java |
| AI/LLM | EINO（DeepSeek + DashScope） | Spring AI + DeepSeek |
| 向量能力 | ES dense_vector + Milvus（Agent） | ES Vector Store |
| 图谱/关系 | Neo4j 预留，当前 Agent 默认关闭 | 无 |
| Binlog 订阅 | Canal Server MQ 模式 | Canal Java Client |
| 分布式锁 | redsync | Redisson |
| JWT | RS256（golang-jwt/jwt/v5） | Spring Security + Nimbus JWT |

### 1.3 Go 微服务架构（当前状态：12 个主进程）

原 Java 项目为单体应用。当前 Go 重构版已进入“统一外部入口 + 内部 RPC / Worker”阶段。

当前主启动拓扑：

```text
统一外部 HTTP 入口（2 个）：
  gateway              :8080   非 agent 业务统一入口
  agent-api            :8011   Agent 独立入口（暂不接入 gateway）

内部 RPC（7 个）：
  user-rpc             :9002   用户 + 认证
  counter-rpc          :9003   内容计数 + 用户计数
  knowpost-rpc         :9004   知文核心
  relation-rpc         :9006   关注关系
  storage-rpc          :9013   OSS 预签名
  search-rpc           :9017   搜索 / suggest
  llm-rpc              :9018   AI 描述 / RAG SSE 源

merged service / worker（3 个主业务进程 + 3 个内嵌 worker + 1 个独立 agent 进程）：
  counter/cmd/counter  内含 counter-rpc 相关业务聚合 + aggregator
  knowpost/cmd/knowpost 内含 knowpost-rpc
  relation/cmd/relation 内含 relation-rpc + syncer
  search/cmd/search    内含 search-indexer（对外 API 默认关闭）
  llm/cmd/llm          内含 ragindexer（对外 API 默认关闭）
  agent/cmd/agent      内含 agent-api + agent-indexer

可选维护进程：
  counter/reconciler
  outbox/gc
```

关键变化：

- `auth/profile/storage/knowpost/relation/counter/search/llm` 的外部 HTTP 入口已统一收敛到 `gateway`
- `agent` 仍保持独立入口，后续可再做独立鉴权或并入统一网关
- `counter/knowpost/relation/search/llm` merged service 默认 `DisableAPI: true`
- 旧 `api` 目录仍保留在仓库中，作为兼容代码与回滚面，不再是当前正式外部入口

### 1.4 当前外部路径约定

对前端可见的主要路径：

| 类型 | 路径前缀 | 实际入口 |
|------|----------|----------|
| 通用业务 | `/api/v1/*` | `gateway:8080` |
| Agent | `/api/v1/agent/*` | `agent-api:8011` |

前端 Vite 开发代理默认指向 `http://localhost:8080`，已与当前网关一致。

### 1.5 pkg 公共库（19 个）

| pkg | 职责 |
|-----|------|
| `pkg/cachex` | 多级缓存、空值哨兵、TTL 抖动 |
| `pkg/canalx` | Canal FlatMessage 解析 |
| `pkg/counterlua` | 位图 / SDS Lua 脚本 |
| `pkg/errorx` | 统一错误码 |
| `pkg/esx` | Elasticsearch REST 客户端 |
| `pkg/hotkey` | 热点探测 |
| `pkg/jwtx` | RS256 JWT |
| `pkg/kafkax` | Kafka consumer / retry / DLQ |
| `pkg/llmx` | EINO embedding 适配 |
| `pkg/lockx` | redsync 分布式锁 |
| `pkg/ossx` | OSS 预签名 / 上传 |
| `pkg/ratelimit` | Redis Lua 令牌桶 |
| `pkg/redisx` | Redis 封装 |
| `pkg/responsex` | 统一响应封装 |
| `pkg/sfx` | singleflight 封装 |
| `pkg/snowflakex` | Snowflake ID |
| `pkg/textx` | 文本清洗 / 切块 |
| `pkg/txx` | 事务辅助 |

## 二、数据库与核心业务对象

数据库主库仍为 `zhiguang`，核心表保持与原 Java 项目兼容：

- `users`
- `know_posts`
- `following`
- `follower`
- `outbox`
- `login_logs`

重构没有改动这些核心业务表的领域含义，重点改在：

- 服务边界
- 外部入口拓扑
- 内部通信方式
- AI / Agent 能力扩展

## 三、模块说明

### 3.1 Auth / User

对外路径：

- `POST /api/v1/auth/send-code`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/token/refresh`
- `POST /api/v1/auth/logout`
- `POST /api/v1/auth/password/reset`
- `GET /api/v1/auth/me`

当前实现：

- 外部入口由 `gateway` 暴露
- 内部由 `user-rpc` 同时承载用户与认证能力
- JWT 校验通过 `user-rpc.VerifyToken`

### 3.2 Profile

对外路径：

- `GET /api/v1/profile/me`
- `PATCH /api/v1/profile/`
- `POST /api/v1/profile/avatar`

当前实现：

- 外部入口由 `gateway` 暴露
- 数据修改走 `user-rpc`
- 头像上传由网关直连 OSS，再回写 `user-rpc`

### 3.3 Storage

对外路径：

- `POST /api/v1/storage/presign`

当前实现：

- 外部入口由 `gateway` 暴露
- 内部能力由 `storage-rpc` 提供
- 主要用于知文内容、图片等直传场景

### 3.4 KnowPost

对外路径：

- `POST /api/v1/knowposts/drafts`
- `POST /api/v1/knowposts/{id}/content/confirm`
- `PATCH /api/v1/knowposts/{id}`
- `POST /api/v1/knowposts/{id}/publish`
- `PATCH /api/v1/knowposts/{id}/top`
- `PATCH /api/v1/knowposts/{id}/visibility`
- `DELETE /api/v1/knowposts/{id}`
- `GET /api/v1/knowposts/feed`
- `GET /api/v1/knowposts/mine`
- `GET /api/v1/knowposts/detail/{id}`
- `POST /api/v1/knowposts/description/suggest`
- `GET /api/v1/knowposts/{id}/qa/stream`
- `POST /api/v1/knowposts/{id}/rag/reindex`

当前实现：

- 外部路径由 `gateway` 提供
- 核心业务在 `knowpost-rpc`
- 兼容路由 `/description/suggest` 和 `/:id/qa/stream` 由网关直接桥接 `llm-rpc`

### 3.5 Counter

对外路径：

- `POST /api/v1/action/like`
- `POST /api/v1/action/unlike`
- `POST /api/v1/action/fav`
- `POST /api/v1/action/unfav`
- `GET /api/v1/counter/{etype}/{eid}`

当前实现：

- 外部入口由 `gateway`
- 内部核心由 `counter-rpc`
- 聚合刷新由 `counter/aggregator`
- 对账由 `counter/reconciler`

### 3.6 Relation

对外路径：

- `POST /api/v1/relation/follow`
- `POST /api/v1/relation/unfollow`
- `GET /api/v1/relation/status`
- `GET /api/v1/relation/following`
- `GET /api/v1/relation/followers`
- `GET /api/v1/relation/counter`

当前实现：

- 外部入口由 `gateway`
- 内部核心由 `relation-rpc`
- 最终一致同步由 `relation/syncer`

### 3.7 Search

对外路径：

- `GET /api/v1/search/`
- `GET /api/v1/search/suggest`

当前实现：

- 外部入口由 `gateway`
- 内部查询由 `search-rpc`
- 索引写入由 `search/indexer`
- `search/cmd/search` 进程继续承载 indexer，但其旧 API 默认关闭

### 3.8 LLM

对外路径：

- `POST /api/v1/llm/describe`
- `GET /api/v1/llm/qa/stream`

兼容路径：

- `POST /api/v1/knowposts/description/suggest`
- `GET /api/v1/knowposts/{id}/qa/stream`

当前实现：

- 外部入口由 `gateway`
- 内部能力由 `llm-rpc`
- 向量索引构建由 `llm/ragindexer`
- `llm/cmd/llm` 进程继续承载 ragindexer，但其旧 API 默认关闭

### 3.9 Agent

对外路径：

- `/api/v1/agent/*`

当前实现：

- `agent` 不接入 `gateway`
- 保持独立 `agent-api:8011`
- 独立使用 ES + Milvus + Redis + MySQL
- 已实现个人知识助手与知识索引能力

## 四、服务划分与调用关系

### 4.1 当前服务清单

| 服务名 | 端口 | 角色 | 说明 |
|--------|------|------|------|
| `gateway` | `:8080` | 外部 HTTP | 非 agent 统一入口 |
| `agent-api` | `:8011` | 外部 HTTP | Agent 独立入口 |
| `user-rpc` | `:9002` | 内部 RPC | 用户 + 认证 |
| `counter-rpc` | `:9003` | 内部 RPC | 内容计数 + 用户计数 |
| `knowpost-rpc` | `:9004` | 内部 RPC | 知文 |
| `relation-rpc` | `:9006` | 内部 RPC | 关系 |
| `storage-rpc` | `:9013` | 内部 RPC | OSS 预签名 |
| `search-rpc` | `:9017` | 内部 RPC | 搜索 |
| `llm-rpc` | `:9018` | 内部 RPC | AI 描述 / QA |
| `counter/cmd/counter` | — | merged | 含聚合 worker |
| `knowpost/cmd/knowpost` | — | merged | 知文主业务 |
| `relation/cmd/relation` | — | merged | 含 syncer |
| `search/cmd/search` | — | merged | 含 search-indexer |
| `llm/cmd/llm` | — | merged | 含 ragindexer |
| `agent/cmd/agent` | — | merged | 含 agent indexer |
| `counter/reconciler` | — | worker | 可选维护进程 |
| `outbox/gc` | — | worker | 可选维护进程 |

### 4.2 请求路径

```text
Browser / Frontend
        ↓
   gateway:8080
        ↓
  user-rpc / storage-rpc / knowpost-rpc / relation-rpc / counter-rpc / search-rpc / llm-rpc
        ↓
 MySQL / Redis / OSS / Elasticsearch / Kafka

Agent 前端或客户端
        ↓
  agent-api:8011
        ↓
 Agent 内部组件 / knowpost-rpc / counter-rpc / ES / Milvus / Redis / MySQL
```

### 4.3 共享依赖

- MySQL：`user-rpc / knowpost-rpc / relation-rpc / outbox-gc / agent`
- Redis：全局共享，按 key 前缀隔离
- Kafka：
  - `counter-events`
  - `canal-outbox`
- Elasticsearch：
  - 业务搜索
  - LLM RAG 向量检索
  - Agent ES 检索
- Milvus：仅 Agent 知识向量库
- OSS：内容与头像对象存储

## 五、实际目录布局

```text
zhiguang-go/
├── docs/
├── deploy/
├── scripts/
├── proto/
├── pkg/
├── common/
└── services/
    ├── gateway/
    ├── user/rpc/
    ├── storage/rpc/
    ├── counter/
    │   ├── rpc/
    │   ├── aggregator/
    │   ├── reconciler/
    │   └── cmd/counter/
    ├── knowpost/
    │   ├── rpc/
    │   └── cmd/knowpost/
    ├── relation/
    │   ├── rpc/
    │   ├── syncer/
    │   └── cmd/relation/
    ├── search/
    │   ├── rpc/
    │   ├── indexer/
    │   └── cmd/search/
    ├── llm/
    │   ├── rpc/
    │   ├── ragindexer/
    │   └── cmd/llm/
    ├── agent/
    │   ├── api/
    │   ├── indexer/
    │   └── cmd/agent/
    └── outbox/gc/
```

说明：

- 旧 `auth/api`、`profile/api`、`storage/api`、`search/api`、`llm/api` 等目录仍保留
- 这些目录当前主要用于历史实现、回滚面或被新 merged service / gateway 复用，不再是正式外部入口

## 六、当前重点实现状态

### 6.1 已完成的大结构调整

- `auth + user` 已合并到 `user-rpc`
- `usercounter` 已合并到 `counter-rpc`
- `counter / knowpost / relation / search / llm / agent` 已完成 merged service 改造
- 新增统一外部入口 `gateway`
- `agent` 暂不接入 `gateway`
- 新增 `storage-rpc / search-rpc / llm-rpc`

### 6.2 与原 Java 项目的当前差异

| 差异点 | Java | 当前 Go |
|---|---|---|
| 外部入口 | 单体内统一 | `gateway + 独立 agent` |
| 认证 HTTP | 单体控制器 | gateway → user-rpc |
| 搜索 HTTP | 单体控制器 | gateway → search-rpc |
| LLM HTTP | 单体控制器 | gateway → llm-rpc |
| Agent | 无 | 新增独立模块 |

### 6.3 当前未实现或后续再完善

- 内容审核
- OpenTelemetry 链路追踪
- 更完整的网关接口级测试
- gateway 与 agent 的统一鉴权模型
- 更细粒度的网关可观测性与配额策略

## 七、端到端启动方式

### 7.1 启动依赖

```bash
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d
```

### 7.2 执行迁移

```bash
make migrate-up
```

### 7.3 编译校验

```bash
go fmt ./...
go vet ./...
go test ./...
go build ./...
```

### 7.4 启动主服务

需要环境变量：

- `DEEPSEEK_API_KEY`
- `DASHSCOPE_API_KEY`
- `MILVUS_API_KEY`（如启用 Milvus）

推荐直接使用：

```powershell
scripts\start-all.ps1
```

当前主启动清单包含：

```text
services/gateway
services/user/rpc
services/storage/rpc
services/search/rpc
services/llm/rpc
services/counter/cmd/counter
services/knowpost/cmd/knowpost
services/relation/cmd/relation
services/search/cmd/search
services/llm/cmd/llm
services/agent/cmd/agent
```

可选维护项：

```powershell
scripts\start-all.ps1 -IncludeMaintenance
```

## 八、文档索引

- 服务合并进度：`docs/phase7-merge-service.md`
- 网关收缩总结：`docs/phase8-gateway-consolidation.md`
- Agent 说明：`services/agent/AGENT-README.md`
