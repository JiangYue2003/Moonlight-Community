# AGENT Project Snapshot (AI-Oriented)

## 0.Meta
- project_root: `services/agent`
- doc_version: `1.0`
- last_verified_date: `2026-05-17`
- target_reader: `LLM/AI Agent`
- source_of_truth:
  - `services/agent/api/internal/*`
  - `services/agent/indexer/internal/*`
  - `services/agent/shared/*`
  - `services/agent/api/etc/agent-api.yaml`
  - `services/agent/indexer/etc/agent-indexer.yaml`

## 1.Runtime Topology
- service_nodes:
  - `agent-api`:
    - role: `在线会话服务`
    - protocol: `HTTP + SSE`
    - listen: `:8011`
    - metrics: `Prometheus /metrics on :9101`
  - `agent-indexer`:
    - role: `异步知识索引服务`
    - protocol: `Kafka consumer + RPC + ES/Milvus writer`
- storage_dependencies:
  - `MySQL`: `会话/反馈/工具审计/长期记忆事实/索引状态`
  - `Redis`: `会话短期记忆+摘要+限流+模型熔断状态+去重`
  - `Elasticsearch`: `知识切块索引(文本+向量字段 embedding)`
  - `Milvus`: `知识向量库 + 记忆事实向量库`
- model_dependencies:
  - `DeepSeek`: `ChatLite + ChatPro`
  - `DashScope`: `Embedding`

## 2.Public Interfaces (agent-api)
- auth_mode: `全部接口挂载 AuthRpc middleware`
- base_path: `/api/v1/agent`

### 2.1 POST `/session`
- request_json:
  - `title` (string, optional)
- response_json:
  - `sessionId` (string)
- side_effect:
  - insert `agent_sessions`

### 2.2 GET `/chat/stream` (SSE)
- query_params:
  - `sessionId` (string, required)
  - `question` (string, required)
  - `topK` (int, optional)
- stream_events:
  - `tool_call`: `{"tool":"hybrid_retrieve","sessionId":"..."}`
  - `tool_result`: `{"tool":"hybrid_retrieve","citations":N,"traceId":"..."}`
  - `route` (conditional, ModelRoute.EmitRouteEvent=true):
    - `{"traceId","model","reason","scenario","fallback",...}`
  - `token`: `{"traceId","content"}`
  - `citation`: `{"postId","chunkId","source"}`
  - `final`: `{"traceId","answer"}`
  - terminal: `data: [DONE]`
- side_effect:
  - append history to Redis list
  - session compact + summary refresh
  - update `agent_sessions.updated_at`

### 2.3 GET `/history`
- query_params:
  - `sessionId` (string, required)
  - `limit` (int, optional, capped)
- response_json:
  - `items[]`:
    - `role`
    - `content`
    - `createdAt`
- data_source: `Redis session list`

### 2.4 POST `/feedback`
- request_json:
  - `sessionId` (string, required)
  - `traceId` (string, required)
  - `score` (int, required, allowed: -1/0/1)
  - `comment` (string, optional)
- response_json:
  - `accepted` (bool)
- side_effect:
  - insert `agent_feedback`

### 2.5 POST `/memory/pin`
- request_json:
  - `sessionId` (string, required)
  - `content` (string, required)
  - `tag` (string, optional)
- response_json:
  - `pinned` (bool)
- side_effect:
  - insert `agent_memory_pin`
  - async fact extraction + upsert to memory stores

## 3.Current Implemented Functions Matrix

### 3.1 会话与对话
- implemented: `YES`
- capabilities:
  - 新建会话
  - SSE 流式回复
  - 会话历史读取
  - 反馈回传
  - 用户级会话权限校验

### 3.2 检索增强问答 (RAG-lite orchestration)
- implemented: `YES`
- retrieval_channels:
  - `milvus vector` (语义召回)
  - `es vector knn` (语义召回)
  - `es bm25` (关键词召回)
  - `memory facts` (长期记忆召回)
  - `graph` (Neo4j channel placeholder, feature flag)
- fusion:
  - `RRF` (`k = Agent.RRFK`, default 60)
- ranking_policy:
  - multi-channel independent ranking -> RRF fused ranking -> topK truncate
- citations:
  - 返回 `postId/chunkId/source`

### 3.3 意图识别与检索路由
- implemented: `YES (rule-based)`
- intent_labels:
  - `exact` / `semantic` / `relation`
- trigger_rules:
  - 编号/代码/URL -> `exact`
  - “谁+负责/关系” -> `relation`
  - default -> `semantic`
- routing_effect:
  - `relation` 且图谱通道可用时启用 graph channel

### 3.4 模型路由与熔断/回退
- implemented: `YES`
- models:
  - `lite`: deepseek-chat
  - `pro`: deepseek-reasoner
- scenario:
  - `chat`
  - `fact_extract`
- route_strategy:
  - 基于问题长度/提示长度/召回量/摘要长度/会话消息数/Pin文本长度进行 rule-based 选型
- circuit_breaker:
  - 使用 Redis 计数 `agent:modelroute:pro:fail`
  - threshold + window 控制 pro 健康度
- fallback:
  - pro 失败可回退 lite (`RetryOnProFail`)

### 3.5 长期记忆
- implemented: `YES`
- flow:
  - `memory/pin` -> 异步 LLM 三元组抽取 -> 置信度过滤 -> MySQL facts upsert -> embedding -> Milvus facts upsert
- confidence_gate:
  - facts `confidence >= 0.5`

### 3.6 知识入库 (收藏驱动)
- implemented: `YES`
- event_source:
  - Kafka topic: `counter-events`
- trigger:
  - `entityType=knowpost && metric=fav`
  - `delta>0` 上架索引
  - `delta<0` 取消收藏删除索引
- indexing:
  - KnowPostRpc 获取详情 -> HTTP 获取正文 -> 切块 -> embedding -> ES bulk upsert + Milvus upsert
- compensation:
  - 定时 ReconcileAndRetry
  - CounterRpc `IsMarked` 校验收藏真值
  - retry backoff + failed 状态恢复

## 4.Technical Map By Function

### 4.1 Retrieval Function 技术构成
- `召回`: Milvus + ES KNN + ES BM25 + Memory Facts (+Graph 可选)
- `多路并发`: goroutine 并发执行通道
- `融合`: RRF
- `重排`: 融合后全局分数排序，取 `TopN`
- `隔离`: 每路检索都强制 `user_id` filter
- `工具化`: 检索步骤通过 Tool Registry 注册/执行

### 4.2 Streaming QA 技术构成
- `协议`: SSE
- `流程`: retrieve -> model route -> stream token -> citations -> final
- `流程(升级)`: shouldPlan -> (可选 planner 拆解) -> hybrid_retrieve(子问题) -> 去重压缩 -> model route -> stream token -> citations -> final
- `记忆`: Redis 消息窗 + 摘要压缩
- `限流`: token-bucket (per-user)

### 4.3 Memory 技术构成
- `短期记忆`: Redis list + summary key (TTL)
- `长期记忆`: MySQL fact table + Milvus fact vectors
- `写入模式`: API 异步触发，不阻塞主响应

### 4.4 Indexing 技术构成
- `解耦`: Kafka 事件驱动
- `去重`: Redis dedup key + TTL
- `切块`: `textx.SplitByHeader` + `textx.Chunk(size, overlap)`
- `双写`: ES + Milvus
- `一致性`: 删除旧版本 -> 写入新版本，失败进入重试

## 5.Observability Specification

### 5.1 埋点位置 (当前已实现)
- `ModelRouter.Decide`:
  - route hit
  - pro circuit open
- `ChatStreamLogic.Run`:
  - model stream call duration/tokens/cost/status
  - fallback record
- `MemoryPinLogic.extractFactsAsync`:
  - model generate duration/tokens/cost/status
  - fallback record

### 5.2 Prometheus 指标暴露
- expose:
  - host/port/path: `agent-api.yaml -> Prometheus`
  - default: `0.0.0.0:9101/metrics`
- metrics:
  - `zhiguang_agent_model_route_total{scenario,model,reason}`
  - `zhiguang_agent_model_fallback_total{scenario,from_model,to_model,reason}`
  - `zhiguang_agent_model_pro_circuit_open_total{reason}`
  - `zhiguang_agent_model_call_duration_ms{scenario,model,method,status}`
  - `zhiguang_agent_model_tokens_total{scenario,model,token_type,source}`
  - `zhiguang_agent_model_cost_total{scenario,model,currency,source}`
- token_source:
  - `usage` (provider usage meta)
  - `estimate` (字符估算 fallback)

## 6.Resource Acquisition (What + How)

### 6.1 用户收藏资源
- resource: `收藏行为事件`
- acquire_method:
  - Kafka 消费 `counter-events`
  - event filter: `knowpost/fav`

### 6.2 知识正文资源
- resource: `知文详情 + 内容正文`
- acquire_method:
  - KnowPostRpc `GetDetail(id, viewerId)`
  - HTTP GET `content_url` 拉取正文

### 6.3 收藏真值资源
- resource: `当前是否仍收藏`
- acquire_method:
  - CounterRpc `IsMarked(entityType=knowpost, metric=fav)`
  - 用于补偿/重试阶段一致性校验

### 6.4 用户身份资源
- resource: `user_id`
- acquire_method:
  - AuthRpc middleware 注入上下文
  - 业务逻辑层读取 ctx user_id

### 6.5 模型资源
- resource:
  - `chat model lite/pro`
  - `embedding model`
- acquire_method:
  - EINO DeepSeek chat
  - EINO DashScope embedding

## 7.Microservice Collaboration Contract

### 7.1 上游/旁路服务
- `user-auth-rpc`:
  - purpose: 鉴权
  - impact_boundary: 同步调用，仅认证，不做重逻辑
- `knowpost-rpc`:
  - purpose: 获取知文详情
  - used_by: indexer
  - impact_boundary: 异步链路，不影响主站请求延迟
- `counter-rpc`:
  - purpose: 收藏真值确认
  - used_by: compensation
- `counter-event`:
  - purpose: 驱动索引增删
  - used_by: indexer consumer

### 7.2 性能隔离策略
- 主业务链路与索引链路解耦:
  - 在线问答读取已建索引，不触发实时切块
  - 收藏入库在 `agent-indexer` 异步完成
- 记忆抽取异步化:
  - `memory/pin` 先响应，再后台抽取事实
- 存储层分工:
  - Redis 短期高频
  - ES/Milvus 检索
  - MySQL 状态与审计

## 8.Security and Sandbox Constraints (Current)
- user_scope_guard:
  - 每次工具执行前强制 `user_id > 0`
  - 检索条件强制 `user_id` filter
- tool_sandbox:
  - whitelist-only tool execution
  - strict schema validation
  - reject unknown params
  - tool call audit table: `agent_tool_audit`
- input_guard:
  - query length cap (`MaxQuestionRunes`)
  - blocked prompt-injection patterns (basic blacklist)
  - topK clamp
- data_scope:
  - 仅访问本服务配置中声明的数据源
  - 不存在文件系统任意读取工具

## 9.Config Keys (Operational)
- `Agent.DefaultTopK/MaxTopK/RRFK`
- `Agent.EnableMilvus/EnableGraph`
- `Agent.ToolWhitelist`
- `Agent.ModelRoute.*`
- `Agent.Observability.Enable`
- `Agent.ModelCost.*`
- `Milvus.*`
- `KnowledgeIndex`
- `Prometheus.*`

## 10.Known Gaps / Non-Implemented or Placeholder
- graph retrieval:
  - `Neo4jProvider` 当前为 feature-flag 占位能力，默认关闭
- advanced reranker:
  - 目前无 cross-encoder rerank，仅 RRF 分数融合
- adaptive intent:
  - 当前为规则法，无训练式分类器
- memory lifecycle:
  - 已有 upsert，缺少完善的衰减/冲突消解/批量压缩策略

## 11.Future Plan (Reference, Optional)
- C.3:
  - 增加 cross-encoder reranker
  - 增加 query rewrite/multi-query recall
  - 图谱通道接入 Neo4j 实检索
- C.4:
  - 完整 agent tool runtime policy (tool budget, timeout budget, risk policy)
  - 细粒度安全策略: PII redaction + tenant rule engine
- C.5:
  - 观测面扩展: recall quality, citation hit-rate, per-channel latency/cost

## 12.Quick Read For Other AI Agents
- `如果只看核心`:
  - 在线: `chat/stream` + `orchestrator.hybridRetrieve + RRF + model router + SSE`
  - 离线: `indexer.Handle -> upsertByFavorite/removeByUnfavorite + retry reconcile`
  - 记忆: `memory/pin -> extractFactsAsync -> MySQL + Milvus`
  - 观测: `observability.AgentObservability` 全链路模型路由/调用/成本指标
