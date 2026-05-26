# Phase 7 - 服务合并进度（search / relation / counter / knowpost / agent / llm）

## 目标
- 将 `search-api + search-indexer` 合并为单进程服务，减少部署与运维复杂度。
- 将 `relation-api + relation-rpc + relation-syncer` 合并为单进程服务，同时保留 REST 与 gRPC 能力。
- 保持对现有业务逻辑与接口行为兼容，优先做“启动编排层”合并，不改核心业务逻辑。

## 本次已完成

### 1. Search 服务合并（已完成）
- 新增合并入口：`services/search/cmd/search/main.go`
- 新增合并配置：`services/search/cmd/search/etc/search.yaml`
- 新增并发编排层：`services/search/cmd/search/internal/app/runner.go`
- 新增组件适配层：`services/search/cmd/search/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/search/api/app/app.go`
  - `services/search/indexer/app/app.go`

### 2. Relation 服务合并（已完成）
- 新增合并入口：`services/relation/cmd/relation/main.go`
- 新增合并配置：`services/relation/cmd/relation/etc/relation.yaml`
- 新增并发编排层：`services/relation/cmd/relation/internal/app/runner.go`
- 新增组件适配层：`services/relation/cmd/relation/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/relation/api/app/app.go`
  - `services/relation/rpc/app/app.go`
  - `services/relation/syncer/app/app.go`

### 3. 既有测试修复（兼容性修正）
- 修复 `services/search/api/internal/logic/logic_test.go` 中过时字段断言：
  - `Snippet` → `Description`
- 原因：当前搜索返回结构中高亮片段已映射到 `description` 字段。

### 4. 启动脚本适配（已完成）
- `scripts/start-all.ps1` 已切换：
  - 删除单独启动：`services/relation/rpc`
  - 删除单独启动：`services/relation/syncer`
  - 删除单独启动：`services/search/indexer`
  - 新增合并启动：`services/relation/cmd/relation -f services/relation/cmd/relation/etc/relation.yaml`
  - 新增合并启动：`services/search/cmd/search -f services/search/cmd/search/etc/search.yaml`

## 测试结果

> 为避免本机 Go 默认缓存目录权限问题，本次使用 `GOCACHE=F:\zhiguang_be\zhiguang-go\.gocache` 执行测试。

通过命令：
- `go test ./services/search/cmd/search/... ./services/relation/cmd/relation/...`
- `go test ./services/search/... ./services/relation/...`

结果：通过。

## 设计说明（关键点）
- 合并策略是“同进程多组件并发运行”，而非把业务逻辑硬合并到一个包。
- `runner` 负责：
  - 并发拉起多个组件；
  - 任一组件异常退出时取消上下文并触发整体退出；
  - 支持 OS 信号触发优雅停机。
- 各原子服务核心逻辑保持原样，降低改造风险。

## 回滚策略
- 旧入口仍保留：
  - `services/search/api`
  - `services/search/indexer`
  - `services/relation/api`
  - `services/relation/rpc`
  - `services/relation/syncer`
- 如需回滚，仅需恢复 `scripts/start-all.ps1` 的服务清单即可。

## 5. Counter 服务合并（已完成）
- 新增合并入口：`services/counter/cmd/counter/main.go`
- 新增合并配置：`services/counter/cmd/counter/etc/counter.yaml`
- 新增并发编排层：`services/counter/cmd/counter/internal/app/runner.go`
- 新增组件适配层：`services/counter/cmd/counter/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/counter/api/app/app.go`
  - `services/counter/rpc/app/app.go`
  - `services/counter/aggregator/app/app.go`
- 补齐 aggregator 可执行入口：
  - `services/counter/aggregator/aggregator.go`
- 与主目录 docs 对齐修正：
  - `GET /api/v1/counter/{etype}/{eid}` 调整为强制鉴权（Bearer），与 `docs/API接口文档_计数.md` 一致。

## 6. KnowPost 服务合并（已完成）
- 新增合并入口：`services/knowpost/cmd/knowpost/main.go`
- 新增合并配置：`services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- 新增并发编排层：`services/knowpost/cmd/knowpost/internal/app/runner.go`
- 新增组件适配层：`services/knowpost/cmd/knowpost/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/knowpost/api/app/app.go`
  - `services/knowpost/rpc/app/app.go`
- 为保持 `llm` 服务独立，新增 knowpost 侧兼容代理：
  - `POST /api/v1/knowposts/description/suggest`（转发到 llm `describe`）
  - `GET /api/v1/knowposts/{id}/qa/stream`（转发到 llm `qa/stream`）
- 与主目录 docs 对齐修正：
  - QA 流式接口改为公开路由（可选鉴权），匹配 `docs/API接口文档_knowpost.md`。
  - 写接口统一返回 `204 No Content`：`content/confirm`、`patch`、`publish`、`top`、`visibility`、`delete`、`reindex`。
  - 增加文档路径兼容：`POST /api/v1/knowposts/{id}/rag/reindex`（保留 `/:id/reindex` 作为兼容）。

## 7. 启动脚本适配（已更新）
- `scripts/start-all.ps1` 已切换为 merged 入口：
  - `services/counter/cmd/counter -f services/counter/cmd/counter/etc/counter.yaml`
  - `services/knowpost/cmd/knowpost -f services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- `counter/reconciler` 继续作为可选维护组件（`-IncludeMaintenance`）独立启动。

## 8. Agent 服务合并（已完成）
- 新增合并入口：`services/agent/cmd/agent/main.go`
- 新增合并配置：`services/agent/cmd/agent/etc/agent.yaml`
- 新增并发编排层：`services/agent/cmd/agent/internal/app/runner.go`
- 新增组件适配层：`services/agent/cmd/agent/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/agent/api/app/app.go`
  - `services/agent/indexer/app/app.go`
- 接口/数据流适配结论：
  - `agent` 对外 `/api/v1/agent/*` 路由未改。
  - 与 `counter-events` 消费、`knowpost-rpc/counter-rpc` 调用、MySQL/Redis/ES/Milvus 写入链路保持一致。

## 9. LLM 服务合并（已完成）
- 新增合并入口：`services/llm/cmd/llm/main.go`
- 新增合并配置：`services/llm/cmd/llm/etc/llm.yaml`
- 新增并发编排层：`services/llm/cmd/llm/internal/app/runner.go`
- 新增组件适配层：`services/llm/cmd/llm/internal/app/components.go`
- 为解决 Go `internal` 包可见性限制，新增可复用启动封装：
  - `services/llm/api/app/app.go`
  - `services/llm/ragindexer/app/app.go`
- 接口/数据流适配结论：
  - `llm-api` 路由保持：`/api/v1/llm/describe`、`/api/v1/llm/qa/stream`。
  - knowpost 兼容路由保持：`/api/v1/knowposts/description/suggest`、`/api/v1/knowposts/:id/qa/stream`。
  - `rag-indexer` 的 `canal-outbox -> vector index` 链路保持不变。

## 10. 启动脚本适配（本次更新）
- `scripts/start-all.ps1` 继续沿用 merged 思路，新增切换：
  - 删除单独启动：`services/agent/api`
  - 删除单独启动：`services/agent/indexer`
  - 删除单独启动：`services/llm/api`
  - 删除单独启动：`services/llm/ragindexer`
  - 新增合并启动：`services/agent/cmd/agent -f services/agent/cmd/agent/etc/agent.yaml`
  - 新增合并启动：`services/llm/cmd/llm -f services/llm/cmd/llm/etc/llm.yaml`

## 11. 测试与构建验证（本次）
为规避本机 Go telemetry 上传权限提示，执行时统一设置：
- `GOCACHE=F:\zhiguang_be\zhiguang-go\.gocache`
- `GOTELEMETRY=off`

执行命令：
- `go test ./services/search/... ./services/relation/... ./services/counter/... ./services/knowpost/... ./services/agent/... ./services/llm/...`
- `go build ./services/agent/... ./services/llm/...`

结果：
- 通过（若出现 telemetry Access is denied 提示，不影响测试与构建结果）。

## 后续服务合并进度
- [x] counter（api/rpc/aggregator，reconciler 独立）
- [x] knowpost（api/rpc，llm 独立）
- [x] agent（api/indexer）
- [x] llm（api/ragindexer）
- [ ] 其他可选维护组件整合

## 12. Phase 7.1 网关收缩（进行中）
- 新增统一外部入口：`services/gateway`
- 统一监听端口：`:8080`
- 前端既有相对路径 `/api/v1/*` 保持不变，Vite 代理默认 `http://localhost:8080`，无需改前端开发代理
- 当前纳入 gateway 的外部链路：
  - `auth/profile/storage/knowpost/relation/counter/search/llm`
- 当前明确排除：
  - `agent` 保持独立，不接入新 gateway
- 新增内部 RPC 服务：
  - `storage-rpc` `:9013`
  - `search-rpc` `:9017`
  - `llm-rpc` `:9018`
- merged service 已增加 `DisableAPI` 开关：
  - `counter/knowpost/relation/search/llm` 默认 `DisableAPI: true`
  - 目的：保留内部 RPC/worker，但关闭旧的对外 HTTP 入口，避免双入口长期并存
- 启动脚本已切换为：
  - 启动 `gateway`
  - 启动 `storage-rpc`
  - 不再启动 `auth-api/profile-api`
  - merged service 继续启动，但其内置 API 默认关闭

## 13. 当前状态
- `storage-rpc/search-rpc/llm-rpc` 已补齐 proto、client、server、logic 与基础测试
- `services/gateway` 已完成第一版骨架与主路由映射
- `go build ./services/gateway` 已通过
- 后续仍需继续做更高覆盖率的接口级测试与文档同步

## 已知风险与建议
- 合并后单进程内组件相互影响（某一组件 panic 可能导致整体退出），建议后续引入更细粒度隔离和重启策略。
- 建议在下一阶段增加组件级健康检查与启动依赖探测（Kafka/MySQL/Redis/ES readiness）。
