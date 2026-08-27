# Feed Hybrid WP6：Inbox + BigV 合并 Pipeline 检查点

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）；Phase 2 正确性、性能与稳定性门禁通过

## 1. 结论

WP6 在不改变 Hybrid 推拉语义的前提下，把 Inbox 与第一批 BigV Outbox 合并到同一个受控 Redis Pipeline：

```text
GetUserFeed
  -> RouteSnapshot
  -> Pipeline #1: Inbox + BigV Outbox[0:N]
  -> Pipeline #2..M: BigV Outbox[N:]
  -> fixed Top-N merge / dedup / balance / FeedItem MGET
```

正式对照中，每个成功请求的 Redis dependency round trips 从约 `3.00` 降到约 `2.00`，下降约 33%；c128 三轮中位数达到 `10661.31 QPS`，P95 `18.112 ms`，在相同单机 Docker 拓扑、同一数据集和零错误口径下达成 10K 目标。

本轮推荐稳定并发仍为 c64：`9866.82 QPS`、P95 `9.171 ms`。c128 是吞吐拐点和 10K 可达点；c256 吞吐回落为 `10587.21 QPS`、P95 上升到 `33.068 ms`，不应通过继续增加并发换取容量数字。

Pipeline 降低的是网络往返，不是 Redis 命令条数。随着 QPS 增长，Redis 峰值已接近 `80K~81K ops/s`；下一瓶颈更可能位于 Redis 命令处理、FeedItem MGET 与 KnowPost CPU，而不是 Relation RPC。

## 2. 实现边界

主要修改：

- `services/knowpost/rpc/internal/feed/reader.go`
- `services/knowpost/rpc/internal/feed/reader_test.go`
- `services/knowpost/rpc/internal/feed/redis_adapter.go`
- `services/knowpost/rpc/internal/feed/observer.go`
- KnowPost Config、两份 YAML 与 Compose 环境变量。
- `cmd/loadtest` 运行态开关校验、快照和报告字段。

核心语义：

- 仅在 `Feed.CombinedPipeline.Enabled=true` 时启用；默认关闭可即时回到旧路径。
- `BatchSize` 默认 128，配置只接受 2 到 128；Inbox 占第一批的一个槽位。
- Redis 实现不支持 batch 接口时自动走原 Inbox + BigV 路径。
- 每个 Pipeline 的返回按请求偏移独立解析；Inbox 或任一 BigV 失败不会丢弃其他成功结果。
- 超过首批时继续分批，并在所有批次间维护固定大小 Top-N，不收集完整结果集。
- Context 取消后不提交后续批次。
- 新增观测操作 `inbox_bigv_pipeline`；一次 `Pipeline.Exec` 计为一次 dependency round trip。
- 本阶段没有引入 sync.Pool、动态并发、新协议或第三方缓存依赖。

## 3. 自动化验证与复审

测试覆盖：

- 默认关闭与显式开启两条路径。
- 第一批请求顺序、Inbox 占位和可配置批次上限。
- 129 个 BigV 跨两个非空批次时，与旧路径的最终 Top-N、时间排序和同时间 ID 降序完全等价。
- `[Inbox 成功, BigV1 失败, BigV2 成功]` 等部分错误偏移隔离。
- Inbox 失败保留 BigV，BigV 失败保留 Inbox/其他 BigV。
- 不支持 batch 接口时的端到端兼容回退。
- Context 取消停止后续批次。
- 真实 RedisAdapter + miniredis 的 Pipeline 行为。
- Combined、后续 BigV 和 Inbox 指标的精确 dependency 次数。
- 集成测试只清理固定测试 key，不扫描或删除开发环境其他 Feed 数据。

已通过：

```text
go test ./services/knowpost/rpc/internal/feed/...
go test -tags=integration ./services/knowpost/rpc/internal/feed/...
go test 相关 WP2-WP6 包
go vet 相关 WP2-WP6 包
git diff --check
```

Combined/observer 定向测试额外重复 50 次通过，真实 RedisAdapter 定向集成重复 10 次通过。复审中提出的跨批次 Top-N、无 batch 接口回退和偏移隔离覆盖均已补齐；最终结论为 Ready，Critical 0 / Important 0。

Windows 当前 `CGO_ENABLED=0` 且 PATH 中没有 `gcc`，无法构建 Go race runtime，因此没有声称 `go test -race` 通过；并发语义由确定性单测、重复测试和正式高并发压测共同验证。

## 4. Docker 与真实运行验证

只复用现有中间件和本地构建层重建 KnowPost，未重新拉取依赖镜像：

```text
FEED_STRATEGY=hybrid
FEED_OBSERVABILITY=true
FEED_RELATION_EPOCH_ENABLED=true
FEED_ROUTE_SNAPSHOT_ENABLED=true
FEED_COMBINED_PIPELINE_ENABLED=true
```

结果：

- KnowPost image：`sha256:eb85f7673203f74efbd852b6ec61013532a6be3a446bb0254b5eb1fb52978311`
- 容器 `healthy`，restart=0；运行态明确记录上述五个开关。
- Hybrid 冒烟和 RelationEpoch 真实集成通过。
- 35,000 请求小样本验证成功 34,959、失败 0，`3495.15 QPS`、P95 `2.644 ms`。
- 小样本观测为 `inbox_bigv_pipeline=1`、`feed_item_mget=1`、`inbox_read=0`，合计约 2 次 dependency round trips/request。
- 正式压测后最近 50,000 行日志中定向扫描 error/fatal/panic/Feed 优化错误为 0。

一次读取完整 20 分钟 Docker 日志因日志量过大超过 60 秒，随后改用有界 50,000 行扫描；这是日志体量问题，不是网络或服务故障。本轮没有镜像下载问题。后续若 Docker 构建或拉取缓慢，仍按 Docker registry/mirror 问题处理，不能归因于网络连通性；必要时保留中间件容器、项目服务改为宿主机运行，并为新拓扑重新建立基线。

## 5. Phase 2 正式性能结果

环境与口径：

- 数据集：`feedbench-20260814c`
- 入口：RPC；场景：`distributed-read`
- 每档：10 秒预热、60 秒正式采样、3 次重复
- WP6：`results/feed-loadtest/feed-wp6-combined-enabled-20260815`
- WP5 对照：`results/feed-loadtest/feed-wp5-route-enabled-20260815`
- Phase 0 严格基线：`results/feed-loadtest/feed-wp2-phase0-strict-20260815`
- 所有 12 个 trial `complete=true`、失败 0、timeout 0、degraded=false、Kafka lag=0。

| 并发 | WP5 QPS | WP6 QPS | QPS 提升 | WP5 P95 | WP6 P95 | P95 下降 | Redis round trips/request 下降 |
|---:|---:|---:|---:|---:|---:|---:|---:|
| 32 | 3572.68 | 7996.33 | 123.82% | 12.050 ms | 5.487 ms | 54.46% | 33.38% |
| 64 | 3978.74 | 9866.82 | 147.99% | 22.054 ms | 9.171 ms | 58.42% | 33.43% |
| 128 | 3896.52 | 10661.31 | 173.61% | 49.301 ms | 18.112 ms | 63.26% | 33.50% |
| 256 | 3884.14 | 10587.21 | 172.58% | 88.936 ms | 33.068 ms | 62.82% | 33.63% |

WP6 三轮波动：

- c32：`7980.49 / 7996.33 / 7997.20 QPS`，spread 0.21%。
- c64：`9763.14 / 9866.82 / 9950.94 QPS`，spread 1.90%。
- c128：`10638.67 / 10661.31 / 10825.93 QPS`，spread 1.76%。
- c256：`10489.76 / 10587.21 / 10651.73 QPS`，spread 1.53%。

相对最初 WP2 严格基线，c32/c64/c128/c256 QPS 分别提升 `311.06% / 324.01% / 311.58% / 279.00%`，P95 分别下降 `73.68% / 74.05% / 71.85% / 73.25%`。

## 6. 容量与资源解释

- 峰值无错三轮中位数：`10661.31 QPS`，c128。
- 推荐稳定并发：c64。
- 吞吐拐点：c128。
- 无错上界（本轮已测）：c256；不等同于推荐运行点。
- c128 比 c64 仅增加约 8.1% 吞吐，但 P95 接近翻倍；c256 吞吐低于 c128且 P95 继续上升，排队已经出现。
- c128 KnowPost CPU 峰值为 `487.67%~533.18%`，Redis 峰值 `79,643~80,510 ops/s`。
- c256 Redis 峰值约 `80,665~81,160 ops/s`；MySQL `threads_running=3`，Kafka lag=0。
- Relation 与 MySQL 已不是主要热路径瓶颈；下一轮 Profile 应优先检查 Redis 命令处理、FeedItem MGET、Protobuf/gRPC marshal 与 KnowPost CPU。

这里的 10K 是当前单机、RPC、约 1,200 分散读者、`page=1,size=20`、固定数据集与资源拓扑下可重复达到的容量证据，不外推为 Gateway 容量、20,000 高基数读者容量或生产 SLA。

## 7. 门禁判定

| 门禁 | 结果 | 判定 |
|---|---|---|
| 冷路径 P95 至少下降 15%，或 Redis round trips/request 至少下降 20% | P95 下降 54.46%~63.26%；往返下降约 33% | 通过 |
| 冷路径 QPS、内存和 GC 不恶化超过 10% | QPS 提升 123.82%~173.61%，无稳定性退化证据 | 通过 |
| Pipeline 部分失败、超时、Context 取消正确 | 确定性单测和重复测试通过 | 通过 |
| 正确性与运行稳定性 | 12/12 完整，失败/timeout=0，healthy、restart=0 | 通过 |
| 指标没有隐藏 Redis 命令负载 | 同时报告 dependency round trips 与 Redis ops/s | 通过 |

## 8. 下一步

Phase 2 门禁通过，允许进入 WP7。但 WP7 PageCache 已不再是达到单机 10K 的必要条件，而是面向热点第一页进一步降低成本和延迟的可选架构阶段；它必须独立证明：

- Fresh L1/L2 命中正确且故障时可回退到当前已验证的 10K 冷路径。
- 短暂陈旧只发生在已批准的 Feed 可接受窗口内，删除、私密、下架等 safety 变化仍受 epoch 保护。
- 20 个热点、约 1,200 分散读者和 20,000 高基数读者三组结果同时报告，不能只展示缓存友好的峰值。
- 新增内存、失效风暴和缓存击穿成本不会让分散/低命中场景回退超过门禁。

进入 WP7 前先保留本报告作为新的冷路径性能基线；任何 PageCache 结果都必须同时与 WP6 对照。
