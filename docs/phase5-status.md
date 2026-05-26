# 阶段5 实施快照（生产化运维）

> 阶段1 状态：[`phase1-status.md`](./phase1-status.md)
> 阶段2 状态：[`phase2-status.md`](./phase2-status.md)
> 阶段3 状态：[`phase3-status.md`](./phase3-status.md)
> 阶段4 状态：[`phase4-status.md`](./phase4-status.md)

## ✅ 已完成

### 新增服务（2 个 worker）

| 服务 | 路径 | 功能 |
|---|---|---|
| outbox-gc | `services/outbox/gc/` | 每小时分批删除 7 天前的 outbox 记录，防止表无限膨胀 |
| counter-reconciler | `services/counter/reconciler/` | 每小时 SCAN SDS key，BITCOUNT 对比，偏差超阈值触发重建 |

### 改动现有代码

| 改动 | 文件 | 说明 |
|---|---|---|
| Kafka DLQ | `pkg/kafkax/consumer.go` | 新增 `MaxRetries` / `DlqTopic` 字段；超过重试次数写 `{topic}-dlq` + ERROR 日志 |
| Prometheus | `services/*/etc/*.yaml`（13 个） | 所有 API/RPC 服务加 `Prometheus: Host/Port/Path` 段，go-zero 内置零代码暴露 `/metrics` |
| LLM 限流 | `services/llm/api/` | per-user token bucket（Redis），`/describe` 5 burst/1 rps，`/qa/stream` 3 burst/1 rps；超限推 SSE 错误 + [DONE] |
| dedup TTL | `services/relation/syncer/etc/syncer.yaml` | `TtlSeconds: 600` → `86400`（P2-6 修复） |
| esx basicAuth | `pkg/esx/client.go` | 删除 30 行手写 base64，改用 `req.SetBasicAuth`（P3-11 修复） |

### 测试覆盖（新增）

| 包 | 用例数 | 关键断言 |
|---|---:|---|
| `services/outbox/gc/internal/worker` | 3 | 7天前记录被删 / 7天内保留 / 分批（多次 DELETE，非一次全删） |
| `services/counter/reconciler/internal/worker` | 4 | SDS==bitmap 不重建 / 偏差>阈值重建 / 偏差<阈值不重建 / SCAN 0 key 正常退出 |

### 编译/校验

- `go fmt ./...` ✅
- `go vet ./...` ✅
- `go build ./...` ✅
- `go test ./...` ✅ **0 FAIL**（34 个测试包全绿）

---

## 🚀 端到端启动（阶段5 新增进程）

```bash
# 在阶段4 已有 17 进程基础上 +2 = 19 进程
go run ./services/outbox/gc etc/gc.yaml           # 新
go run ./services/counter/reconciler etc/reconciler.yaml  # 新
```

---

## ⚠️ 阶段5 显式不做

- 内容审核（敏感词 / NSFW）
- 链路追踪（OpenTelemetry）
- ES 集群 / 快照
- Canal HA
- LLM Trace / Token 计费明细

---

## 📊 全阶段对比

| 维度 | 阶段1 | 阶段2 | 阶段3 | 阶段4 | 阶段5 |
|---|---|---|---|---|---|
| 服务 | 5+1w | 8+1w | 11+2w | 13+4w | **13+6w（+2 worker）** |
| pkg | 8 | 13 | 16 | 19 | 19（无新增，有改动） |
| 测试包 | 6 | 20 | 25 | 34 | **34（+7 用例）** |
| 原项目进度 | 6/18 | 9/18 | 12/18 | 15/18 | **17/18** |

剩余 1/18：内容审核（原 Java 也未实现，可选）。
