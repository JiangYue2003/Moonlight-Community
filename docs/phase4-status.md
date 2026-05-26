# 阶段4 实施快照（搜索 + LLM + RAG）

> 项目分析：[`project-analysis.md`](./project-analysis.md)
> 阶段1 状态：[`phase1-status.md`](./phase1-status.md)
> 阶段2 状态：[`phase2-status.md`](./phase2-status.md)
> 阶段3 状态：[`phase3-status.md`](./phase3-status.md)

## ✅ 已完成（端到端：knowpost.publish → outbox → Canal binlog → Kafka → search-indexer/rag-indexer → ES 倒排+向量；llm-api → DeepSeek 流式）

### 新增 pkg 库（3 个）
- `pkg/esx` ES 9.x REST 客户端薄封装：`Client / EnsureIndex / Index / Update / Delete / Get / Bulk / DeleteByQuery / Search / Suggest / KnnSearch / Count`
  - mapping JSON `//go:embed`：`mapping/content.json` `mapping/rag.json`
  - 不引官方 SDK，net/http + json，httptest 直测
- `pkg/llmx` LLM 客户端
  - `ChatClient.Complete / Stream`：DeepSeek（OpenAI 兼容 `/v1/chat/completions`），SSE `[DONE]` 关闭 channel
  - `EmbedClient.Embed`：阿里通义 `text-embedding-v3`（1024 维），自动 25 条/批分组
- `pkg/textx` Markdown 切块 + 中文清洗
  - `SplitByHeader`（ATX 标题）、`Chunk`（rune 滑窗，支持 size/overlap）
  - `NormalizeNFKC / CollapseWhitespace / StripWrappingQuotes / StripTrailingPunct / TruncateRunes`
  - `DescriptionPostProcess` 描述生成专用清洗管线（NFKC → 折空白 → 去引号 → 取首段 → 去末标点 → 截断）

### knowpost 扩展（事件主干）
- 新增 `services/knowpost/shared/event/event.go`：`KnowPostEvent` + 4 类型常量（Created / Published / Updated / Deleted）
- `services/knowpost/shared/model.KnowPostsModel.UpdateInTx`：sess 版本，可在外部事务内 update
- `services/knowpost/rpc/internal/logic/knowpost/`：
  - `publishlogic.go`：UPDATE + 写 outbox（KnowPostPublished）同事务
  - `deletelogic.go`：UPDATE 软删 + 写 outbox（KnowPostDeleted）同事务
  - `outbox_helper.go`：`updateAndEmitOutbox` 统一封装，仅 published 状态广播 KnowPostUpdated
  - `patchmetadatalogic / updatevisibilitylogic / updatetoplogic` 全部接入

### 新增 / 补丁服务（4 + 1）
| 服务 | 端口 | 角色 |
|---|---|---|
| search-api | :8007 | HTTP；多字段检索 + Suggest + cursor 分页；JWT 可选 |
| search-indexer | (worker) | 消费 canal-outbox（aggregate_type=knowpost）→ ES UPSERT + 启动 backfill |
| llm-api | :8008 | HTTP；POST /describe（DeepSeek 同步） + GET /qa/stream（SSE 流式 RAG） |
| rag-indexer | (worker) | 消费 canal-outbox → 切块 / 通义嵌入 / ES dense_vector 写入；指纹去重 |
| counter-rpc 补丁 | :9003 | GetCounts 重建路径用 Redis SCAN 替代"前 64 chunk"硬编码 |

### 部署
- `deploy/compose/docker-compose.dev.yml`：新增 `elasticsearch:9.0.3` 容器（healthcheck）+ `kibana:9.0.3`（profile=tools 不默认起）；新增 `zg_es_data` volume
- `deploy/compose/es-plugins/analysis-ik/`：留空目录 + README（用户按版本下载 IK 解压进去；版本必须与 ES 9.0.3 严格一致）
- ES 索引初始化由服务启动时 `EnsureIndex` 完成（不走 SQL migration）

### 测试覆盖（新增）
- `pkg/esx`：8 用例（EnsureIndex 幂等 / Search hits 解析 / Suggest 去重 / KnnSearch body 组装 / 4xx 透出 / Delete-NotFound 幂等 / Count-NotFound=0 / mapping embed）
- `pkg/llmx`：6 用例（Complete OK / Complete 4xx / Stream tokens-DONE / Stream 4xx / Stream ctx 取消 / Embed 自动分批 60→3 次）
- `pkg/textx`：10 用例（SplitByHeader 三态 / Chunk rune-safe / NFKC / 折空白 / 去引号四种成对 / 去末标点 / 截断 / DescriptionPostProcess 全管线）
- `services/search/api/internal/query`：5 用例（Build 空 tags / 有 tags+after / cursor encode-decode / 空 cursor / 坏 cursor）
- `services/search/api/internal/logic`：5 用例（空 q / highlight snippet / description fallback / 坏 cursor 回退 / Suggest 空前缀 + OK）
- `services/search/indexer/internal/processor`：5 用例（published+public Upsert / draft 转 SoftDelete / SETNX 去重 / 非 knowpost 跳过 / 坏 JSON 跳过）
- `services/llm/api/internal/logic`：5 用例（Describe 后处理 / Describe 空 body / Qa 0 命中 / Qa token 流 + [DONE] / Qa filterByPostId / Qa 参数校验）
- `services/counter/rpc/.../getcountslogic`：新增 1 用例（SCAN 跨 chunk=0/100/10000 全量求和）

**累计阶段1+2+3+4 测试：200+ 用例 / 34 个测试包，全部通过。**

### 编译/校验
- `go fmt ./...` ✅ 全部归一
- `go vet ./...` ✅ 无告警
- `go build ./...` ✅ 全绿
- `go test ./...` ✅ **0 FAIL**

---

## ⚠️ 阶段4 显式不做（移交阶段5）
- 多语言搜索（仅中文 IK）
- 同义词词典 / 自定义停用词
- 图片向量召回（CLIP）
- LLM 函数调用 / Tool Use
- RAG re-ranking（如 bge-reranker）
- LLM 答案"引用片段"反向定位
- ES 集群 / 快照 / 跨 region
- LLM 限流 / 配额（按 user / 按 IP）
- 内容审核（敏感词 / NSFW）
- LLM Trace / Token 计费明细
- counter 跨副本对账 / SDS GC

---

## 🚀 端到端启动

```bash
# 0. 准备 IK 插件（首次）：把 elasticsearch-analysis-ik-9.0.3 解压到 deploy/compose/es-plugins/analysis-ik/

make up                                  # mysql/redis/kafka/zk/canal-server/elasticsearch
sleep 30                                 # ES + canal 首启较慢
curl -s localhost:9200/_cluster/health   # status=green/yellow
curl -s localhost:9200/_cat/plugins      # 含 analysis-ik

make migrate-up
go fmt ./... && go vet ./... && go build ./... && go test ./...

# 必须设置（DeepSeek + 通义 API key）
export DEEPSEEK_API_KEY=sk-xxxxx
export DASHSCOPE_API_KEY=sk-xxxxx

# 在阶段3 已有 13 进程基础上 +4 = 17 进程
go run ./services/user/rpc                # :9002
go run ./services/auth/rpc                # :9001
go run ./services/auth/api                # :8001
go run ./services/counter/rpc             # :9003 (含 SDS-SCAN 重建)
go run ./services/counter/api             # :8002
go run ./services/counter/aggregator      # worker
go run ./services/usercounter/rpc         # :9005
go run ./services/knowpost/rpc            # :9004
go run ./services/knowpost/api            # :8005
go run ./services/storage/api             # :8003
go run ./services/profile/api             # :8004
go run ./services/relation/rpc            # :9006
go run ./services/relation/api            # :8006
go run ./services/relation/syncer         # worker
go run ./services/search/api              # :8007  ← 新
go run ./services/search/indexer          # worker ← 新
go run ./services/llm/api                 # :8008  ← 新
go run ./services/llm/ragindexer          # worker ← 新
```

详见 `plans/refactored-dazzling-sunset.md` §11 / §12 验证脚本。

---

## 📊 阶段对比

| 维度 | 阶段1 | 阶段2 | 阶段3 | 阶段4 |
|---|---|---|---|---|
| 服务 | 5 + 1 worker | 8 + 1 worker | 11 + 2 worker | **13 + 4 worker（+2 新 + 2 worker + 1 补丁）** |
| pkg | 8 | 13 | 16 | **19（+esx + llmx + textx）** |
| proto | 3 | 5 | 6 | 6（阶段4 不新增 proto） |
| 测试包 | 6 | 20 | 25 | **34（+9）** |
| 测试用例 | 41 | 80+ | 100+ | **200+** |
| 模块覆盖 | auth / user / counter | + knowpost / storage / profile / usercounter | + relation / canal-bridge | **+ search / llm / rag** |
| 原项目 18 服务进度 | 6/18 | 9/18 | 12/18 | **15/18** |

剩余阶段5：outbox GC / 计数对账 / counter-aggregator HA / Canal HA / 监控 / 审核。
