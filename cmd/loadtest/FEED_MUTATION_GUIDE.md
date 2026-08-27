# Feed 写入/混合正式压测指南

本文只覆盖会修改数据的 `publish`、`mixed-90-10`、`mixed-80-20`。纯读矩阵继续使用 `run_feed_matrix.ps1`。

## 1. 已实现的数据隔离

每个 measured trial 使用同一个 manifest-owned checkpoint，并固定执行：

1. Kafka lag=0 且 log-end offset 连续稳定 2 秒；
2. 校验 RunID、manifest、strategy、seed、Redis identity 和被测二进制身份；
3. 删除 benchmark 作者基线之后的精确帖子、Outbox 和可推导 Redis 派生 key；
4. 恢复 manifest 下所有 Inbox/BigV ZSET，校验完整指纹，递增 safety epoch；
5. 若配置 warmup，warmup 结束后再次完整恢复；
6. 执行 measured trial，记录成功发布 post IDs 和资源指标；
7. 先写入 `mutation_restore_pending` 的 incomplete 报告，再执行测量后恢复；只有恢复成功且所有已发布 ID 均被删除，报告才升级为 `complete=true`。

任何 checkpoint 缺失、身份不一致、Kafka 未稳定、MySQL/Redis 指纹变化、恢复范围不完整或指标缺失都会停止矩阵。

## 2. 创建和验证 checkpoint

先使用已确认、无 mutation 在途的基线数据集。示例：

```powershell
$runId = "feed-wp11-hot-20260816a"
$manifest = "results/feed-loadtest/manifests/$runId.json"

go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase create-mutation-checkpoint `
  -run-id $runId -manifest $manifest `
  -reader-cardinality hot -strategy hybrid `
  -topology compose-middleware-local-services `
  -feature-preset treatment `
  -confirm-mutation $runId

go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase verify-mutation-checkpoint `
  -run-id $runId -manifest $manifest `
  -reader-cardinality hot -strategy hybrid `
  -topology compose-middleware-local-services `
  -feature-preset treatment
```

默认文件为 `results/feed-loadtest/checkpoints/<run-id>.json`，使用独占创建，存在时拒绝覆盖。若代码或数据基线确实变化，应使用新的 RunID/checkpoint；不要删除旧文件后假装是同一基线。

独立 verify 允许有过期时间的 key 从 checkpoint 时刻起自然减少剩余 TTL，但不允许 TTL 被重新延长；restore 内部的即时验收仍使用严格双向容差。键缺失、内容指纹变化、持久/过期属性变化都会拒绝继续。

## 3. 短时门禁

正式长测前，先分别验证 RPC 和 Gateway。所有 mutation run 都必须传 `-confirm-mutation`：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase run -run-id $runId -manifest $manifest `
  -report-run-id "$runId-mutation-smoke" `
  -reader-cardinality hot -strategy hybrid `
  -topology compose-middleware-local-services `
  -feature-preset treatment -cache-state cold `
  -entry rpc -scenario mixed-90-10 -concurrency 4 `
  -requests 0 -duration 3s -warmup 1s -trials 1 `
  -confirm-mutation $runId
```

通过条件：报告 `complete=true`、失败/超时为 0、实际读写比例符合场景、`published_post_ids` 数量等于 publish success、`after_trial_restore.complete=true`、删除集合覆盖全部已发布 ID，且下一轮 `verify-mutation-checkpoint` 通过。

## 4. 正式矩阵

Control 与 Treatment 复用同一 RunID、manifest、checkpoint，使用不同 ReportRunId。严格 A/B 默认都使用 `cold`：

```powershell
./cmd/loadtest/run_feed_mutation_matrix.ps1 `
  -RunId $runId -ReportRunId "$runId-treatment-mutation-formal" `
  -Strategy hybrid -ReaderCardinality hot `
  -FeaturePreset treatment -CacheState cold `
  -Topology compose-middleware-local-services `
  -Scenarios @("publish","mixed-90-10","mixed-80-20") `
  -Concurrencies @(32,64,128,256) `
  -GatewayConcurrencies @(32,64,128,256) `
  -GatewayAuthMode signed `
  -WarmupSeconds 10 -DurationSeconds 60 -Trials 3
```

脚本启动前会验证 checkpoint，逐场景/并发运行 RPC 和 Gateway，任一 incomplete trial 立即停止，全部完成后生成 comparison。若按已批准的压缩策略只跑最优档及相邻档，直接缩减两个 concurrency 数组，不降低 60 秒和三轮门禁。

`ReportRunId` 必须是新的非空命名空间，脚本拒绝覆盖旧报告。脚本会先写 `expected-matrix.json`，结束时逐格检查入口、场景、并发、`1..N` 轮次、报告完整性和测量后恢复，并拒绝缺失或额外的 `report.json`；随后再次验证 checkpoint，最后才生成 comparison。首次 restore 前已经存在 pending/incomplete 审计报告，准备阶段失败也不会留下无证据空洞。

本地混合拓扑只复用现有 Docker 中间件，业务服务在本地运行；镜像构建/下载异常按 Docker registry/mirror source 问题处理，不归因于网络。

## 5. 报告解释

短时 smoke 只证明 checkpoint/restore 和入口正确，不能报告单机容量。正式结论至少分别给出：

- publish 成功 QPS、四阶段 P95/P99、Kafka 稳定恢复时间；
- 90/10、80/20 的实际完成比例、读 QPS/延迟、写 QPS/延迟；
- Gateway 相对 RPC 的吞吐损耗；
- KnowPost/Gateway/Relation/Redis/MySQL CPU、RSS、Redis ops、Kafka lag；
- checkpoint ID、baseline fingerprint、生成/删除 post 数和恢复耗时；
- Control/Treatment 三轮中位数、min/max，以及首个错误、客户端饱和或资源拐点。
