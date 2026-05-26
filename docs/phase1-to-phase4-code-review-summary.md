# 阶段1-4 全量 Code Review 总结

> 最后更新：2026-05-16（已完成 P0 全部、P1 部分、EINO 迁移）
> 覆盖范围：19 个 pkg + 13 服务 + 4 worker（阶段1-4 全部实现）

---

## ✅ 已修复

### P0-1. `esx.Client.rr` 数据竞争 ✅ 已修复
- **文件**：`pkg/esx/client.go`
- **修复**：`c.rr++` → `atomic.AddUint32(&c.rr, 1)`，加 `sync/atomic` import。

### P0-2. 撤稿后 ES 不软删（数据不一致）✅ 已修复
- **文件**：`services/knowpost/rpc/internal/logic/knowpost/outbox_helper.go`
- **修复**：移除 `status != "published"` 时跳过 outbox 的分支，无论 status 如何都写 `KnowPostUpdated` outbox。
  search-indexer 收到事件后检查 status/visible，对非 published 帖子执行 SoftDelete，触发路径补全。

### P1-4. `kafkax` 消费失败无退避，CPU 空转 ✅ 已修复
- **文件**：`pkg/kafkax/consumer.go`
- **修复**：加指数退避（100ms × 2^retries，上限 30s），失败时记录错误日志。成功后重置退避计数。

### P1-5. SSE 多行 `data:` 解析不符合规范 ✅ 随 EINO 迁移消除
- **原文件**：`pkg/llmx/chat_stream.go`（已删除）
- **说明**：`pkg/llmx` 旧实现整包替换为 EINO DeepSeek client，SSE 解析由框架内部处理，问题自然消除。

### EINO 迁移 ✅ 已完成
- **删除**：`pkg/llmx/chat.go` / `chat_stream.go` / `embedding.go` / `types.go` / `chat_test.go`
- **新增**：`pkg/llmx/embed_adapter.go`（`EmbedFloat32`：float64→float32 适配，3 行核心逻辑）
- **改写**：
  - `services/llm/api/internal/svc/servicecontext.go`：Chat 换 `einodeepseek.NewChatModel`，Embed 换 `einodashscope.NewEmbedder`
  - `services/llm/api/internal/logic/describelogic.go`：`Chat.Complete` → `Chat.Generate`，用 `schema.Message`
  - `services/llm/api/internal/logic/qastreamlogic.go`：`for ev := range ch` → `StreamReader.Recv() + io.EOF`
  - `services/llm/ragindexer/internal/svc/servicecontext.go`：Embed 换 EINO DashScope
  - `services/llm/ragindexer/internal/processor/processor.go`：`Embed.Embed` → `llmx.EmbedFloat32`
  - `services/llm/api/internal/logic/logic_test.go`：mock 换为 EINO 接口 + deepseek-go 标准响应格式

---

## 🔴 P1 — 高优先级（待修复）

### 3. MySQL 密码明文硬编码进 git
- **文件**：`services/auth/rpc/etc/auth.yaml:12`、`services/knowpost/rpc/etc/knowpost.yaml:7`、
  `services/relation/rpc/etc/relation.yaml:7`、`services/relation/syncer/etc/syncer.yaml:2`、
  `services/storage/api/etc/storage-api.yaml:13`、`services/user/rpc/etc/user.yaml:7`
- **问题**：`DataSource: root:root@tcp(...)` 明文密码进版本库；OSS `AccessKeyId: replace-me` 同理。
- **修复**：改为 `${DB_DSN}` 环境变量占位，或在 `.gitignore` 中排除 `etc/*.yaml`，
  改用 `etc/*.yaml.example` 作为模板。

---

## 🟡 P2 — 中优先级（待修复）

### 6. relation-syncer dedup TTL 太短，重投后重复计数
- **文件**：`services/relation/syncer/internal/processor/processor.go:57`（TTL 600s）
- **问题**：Kafka consumer group rebalance 或 offset reset 超过 10 分钟后，
  SETNX 已过期，`usercounter.Increment` 会被重复调用，关注/粉丝计数永久偏高。
- **修复**：TTL 改为 24h（`86400s`），或改用幂等 upsert 替代 Increment（根本解法）。

### 7. rag-indexer 无测试文件
- **文件**：`services/llm/ragindexer/internal/processor/`（processor.go / fingerprint.go / backfill.go）
- **问题**：阶段4 唯一没有测试覆盖的 processor，指纹去重、切块、嵌入分批等关键逻辑无验证。
- **修复**：参照 `services/search/indexer/internal/processor/processor_test.go` 补充测试，
  重点覆盖：首次 Upsert / 同 sha256 跳过 / sha256 变化重建 / status=draft 删 chunk / 60 chunk 分批。

### 8. `esx.Bulk` 错误码语义混淆
- **文件**：`pkg/esx/doc.go:99`
- **问题**：`&Error{Status: http.StatusOK, Body: "bulk has item errors: ..."}` 把 item-level 错误
  标成 HTTP 200，`IsNotFound` 等调用方会误判。
- **修复**：定义独立 `BulkError` 类型，或用 `Status: -1` 表示"非 HTTP 层错误"。

### 9. `llmx` 跨批索引断言是伪测试（随 EINO 迁移已无此代码，可关闭）
- **说明**：原 `pkg/llmx/chat_test.go:164-167` 已随旧实现删除，此条自动关闭。✅

---

## 🟡 P3 — 低优先级（待修复）

### 10. `esx` 的 `refresh=wait_for` 硬编码
- **文件**：`pkg/esx/doc.go:11/22/33/87`
- **问题**：每次写都等 ES 刷新，批量回填（backfill）场景吞吐被严重拖慢。
- **修复**：给 `Index/Update/Delete/Bulk` 加可选 `RefreshPolicy string` 参数（`""`/`"true"`/`"wait_for"`）。

### 11. `esx.Client` 自己实现 base64
- **文件**：`pkg/esx/client.go`
- **问题**：30 行手写 base64，`req.SetBasicAuth(user, pass)` 一行等价。
- **修复**：删掉 `basicAuth` 函数，改用 `req.SetBasicAuth`。

### 12. `llmx.ChatClient.Stream` 每次 new `http.Client`（随 EINO 迁移已消除）✅
- **说明**：旧 `chat_stream.go` 已删除，EINO ChatModel 内部复用连接池，此条自动关闭。

### 13. `llmx.StreamEvent.Done` 是死代码（随 EINO 迁移已消除）✅
- **说明**：`types.go` 已删除，此条自动关闭。

### 14. `relation-rpc` 裸信任 `from_user_id` 入参
- **文件**：`services/relation/rpc/internal/logic/relation/followlogic.go:38`
- **问题**：`from_user_id` 由 relation-api 从 JWT ctxdata 注入，但 proto 字段是公开的，
  任何能访问 gRPC 端口的调用方都能伪造。
- **修复**：在 relation-api handler 层强制用 ctxdata 覆盖 `from_user_id`，不信任前端传值；
  或在 rpc 层加注释说明信任边界。

### 15. `search-api` size 超限静默降级
- **文件**：`services/search/api/internal/logic/searchlogic.go:36`
- **问题**：`size > 50` 时静默改为 20，调用方无感知，难以调试。
- **修复**：返回 400 Bad Request，或在 handler 层校验并返回明确错误。

### 16. `DescriptionPostProcess.Apply` 顺序与注释不符
- **文件**：`pkg/textx/clean.go:99-109`
- **问题**：实际执行顺序是 NFKC → 取首段 → 折空白 → 去引号 → 去末标点 → 截断，
  但顶部注释描述的顺序不同，会误导维护者。
- **修复**：更新注释与实际执行顺序一致。

---

## 🟢 亮点（设计良好，值得保留）

| 亮点 | 文件 |
|---|---|
| RS256 JWT，token_type 强校验防 access/refresh 互换，jti 支持白名单旋转 | `pkg/jwtx/rs256.go` |
| lockx 用 redsync，WithTries(1) 立即让步，区分"锁被占"与"Redis 错误" | `pkg/lockx/redsync.go` |
| 三级缓存 + singleflight + 空值哨兵 + TTL 抖动 + 热点延长，设计完整 | `pkg/cachex/` |
| Outbox 事务边界：following + outbox 同事务，写路径不维护 ZSet/follower/usercounter | `relation/rpc/logic/followlogic.go` |
| kafkax 手动提交 + 指数退避：handler 成功才 commit offset，失败有退避上限 30s | `pkg/kafkax/consumer.go` |
| textx.Chunk rune 安全滑窗，overlap >= size 时收敛防死循环 | `pkg/textx/chunk.go` |
| esx mapping //go:embed，部署不丢文件，mapping 版本与代码同步 | `pkg/esx/mapping/` |
| bcrypt cost=12，密码哈希安全合理 | `services/auth/rpc/etc/auth.yaml` |
| Canal MQ 模式，HA/重连/position 由 Canal Server 兜底，Go 侧只消费 Kafka | `deploy/compose/canal-conf/` |
| search-indexer 软删而非物理删，保留文档可恢复性 | `search/indexer/processor/processor.go` |
| counter-rpc 重建用 Redis SCAN 全量枚举，覆盖任意 entityId（阶段4 清技术债） | `counter/rpc/logic/getcountslogic.go` |
| EINO 统一 LLM 接口，SSE 解析、连接池、多 provider 切换由框架保障 | `pkg/llmx/embed_adapter.go` + `services/llm/` |

---

## 修复优先级速查

```
✅ 已完成：
  P0-1. pkg/esx/client.go          — rr 数据竞争 → atomic
  P0-2. outbox_helper.go            — 撤稿后 ES 不软删
  P1-4. pkg/kafkax/consumer.go      — 失败无退避 → 指数退避（上限 30s）
  P1-5. pkg/llmx/chat_stream.go     — SSE 解析 → EINO 迁移消除
  EINO. pkg/llmx + llm-api + rag-indexer → EINO ChatModel + Embedder

待修复：
  P1-3. services/*/etc/*.yaml        — MySQL 密码明文 → 环境变量
  P2-6. relation/syncer/processor.go — dedup TTL 600s → 86400s
  P2-7. llm/ragindexer/processor/    — 补充测试文件
  P2-8. pkg/esx/doc.go               — Bulk 错误类型 → BulkError
  P3-10~16. 见上方各条
```

---

## 阶段进度回顾

| 阶段 | 服务数 | pkg 数 | 测试包 | 测试用例 | 原项目进度 |
|---|---|---|---|---|---|
| 阶段1 | 5 + 1 worker | 8 | 6 | 41 | 6/18 |
| 阶段2 | 8 + 1 worker | 13 | 20 | 80+ | 9/18 |
| 阶段3 | 11 + 2 worker | 16 | 25 | 100+ | 12/18 |
| 阶段4 | 13 + 4 worker | 19 | 34 | 200+ | **15/18** |
| 修复轮 | — | — | 34 | 200+ | 15/18（P0+P1-4+EINO 落地） |

剩余阶段5：outbox GC / 计数对账 / counter-aggregator HA / Canal HA / 监控 / 审核。
