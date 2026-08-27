# Feed 推拉混合架构综合性能报告

日期：2026-08-17  
范围：个性化 Feed Hybrid 读路径、Cursor 深分页与 mutation RPC；补充早期推/拉/混合探索、公共 Feed 缓存和历史 Gateway 证据  
状态：WP1-WP11 总结；纯读、Cursor 深分页、mutation RPC 正式 A/B 已完成，后续 Gateway 性能矩阵取消

## 1. 执行摘要

本轮演进已经证明：最初的 Hybrid 性能瓶颈不在归并算法，也不是客户端并发不够，而是每次请求都重复执行 Relation/MySQL 路由查询和多段 Redis 网络往返。按相同 Compose 全栈、约 1,200 个分散读者、Hybrid RPC、page1/size20、10 秒预热 + 60 秒采样 × 3 的严格口径：

| 阶段 | 关键变化 | c128 QPS | c128 P95 | 相对 WP2 |
|---|---|---:|---:|---:|
| WP2 Phase 0 | 每请求 Relation + 约 3 次 Redis 往返 | 2,590.31 | 64.343ms | 基线 |
| WP5 RouteSnapshot | 热路由复用，Relation/request 约 1 → 0.001 | 3,896.52 | 49.301ms | QPS +50.43%，P95 -23.38% |
| WP6 Combined Pipeline | Inbox + 首批 BigV 合并，Redis 往返约 3 → 2 | 10,661.31 | 18.112ms | QPS +311.58%，P95 -71.85% |

因此，历史上“Docker 内达到 1 万 QPS”是真实且可重复的：WP6 c128 三轮为 10,638.67 / 10,661.31 / 10,825.93 QPS，失败、超时和 degraded 均为 0。它只代表当时的单机 Feed RPC 内核、约 1,200 个读者、第一页冷计算路径，不等于 Gateway、20,000 个高基数读者或生产集群 SLA。

随后引入的完整页面 L1/L2 Cache、Singleflight 和 SWR 进一步证明了局部性的重要性：

| 工作集与入口 | Control | Treatment | 结果与含义 |
|---|---:|---:|---|
| hot RPC，20 readers，c32 | 3,028.91 QPS / P95 13.401ms | 8,222.49 / 5.309ms | 目标态 QPS +171.47%，P95 -60.38% |
| hot Gateway，20 readers，c32 | 2,379.29 / 17.029ms | 4,930.08 / 9.835ms | 目标态 QPS +107.21%，P95 -42.25% |
| distributed RPC，1,200 readers，c64 | 6,797.75 / 14.06ms | 102,075.49 / 1.16ms | cold Control → l2-warm Treatment，目标态 15.02 倍；不是同缓存状态代码 A/B |
| high RPC，20,000 readers | c32 6,162.47 / 7.55ms | 稳定点 c16 9,477.51 / 3.73ms | 最佳稳定容量 +53.79%；并发不同，不是严格同档 A/B |
| high Gateway，20,000 readers，c8 | 1,402.80 / 11.20ms | 1,407.80 / 15.16ms | QPS 仅 +0.36%，P95 +35.36%，高基数入口没有得到缓存红利 |
| deep-page RPC，c32 | 2,829.27 / 13.900ms | 2,932.98 / 13.990ms | QPS 仅 +3.67%，P95 +0.65%，第一页缓存没有解决深分页 |

综合结论如下：

1. Hybrid 冷路径已完成两次结构性降本，单机 RPC 在固定口径下从约 2.6K 提升到 10.66K QPS。
2. 整页缓存对热点和可复用工作集非常有效，但收益不是常数；读者越分散、TTL 内重复访问越少，Fresh ratio 越低。
3. 当前最明显的剩余问题是 Gateway 真实入口、高基数缓存复用、刷新队列容量，以及 page-number 深分页的重复计算。
4. 继续堆并发不会提升容量。每轮容量曲线都显示越过 knee 后吞吐趋平或下降、P95 急升。
5. 2026-08-27 已补齐写入、90/10、80/20 的 RPC 正式三轮 A/B；Treatment 总吞吐未提升，具体归因仍需单开关或同窗 Profile。

## 2. 报告口径与证据等级

为避免把不同条件下的峰值拼成一个结论，本报告使用四级证据：

| 等级 | 定义 | 可以回答什么 |
|---|---|---|
| A：正式容量 | 固定代码/拓扑/数据/缓存前态；10 秒预热、60 秒采样、独立三轮取中位数；全部完整性门禁通过 | 稳定容量、并发拐点、严格或明确标注条件的对比 |
| B：目标状态对比 | 两边代码或缓存前态不同，但变量和差异被完整记录 | 产品运行目标的收益；不能冒充单项代码增益 |
| C：功能/短样本/微基准 | 1-10 秒样本、单轮、Go benchmark、功能集成 | 路径是否工作、优化是否有潜力；不能作为容量上限 |
| D：无效实验 | 指标缺失、客户端饱和、服务重启、Redis 异常、Kafka 未恢复、Feed degraded、写数据初态漂移等 | 只用于定位过载边界或工具问题，不进入容量结论 |

本报告中的 WP2、WP5、WP6 属于 A 级且拓扑、数据、入口、并发档一致，可做严格阶段对比。WP11 hot/distributed/high 的正式轮次也是 A 级容量证据，但 Control/Treatment 若使用不同缓存前态，明确标为 B 级目标状态对比。

2026-07-03 的公共 Feed 实验测试的是 `/api/v1/knowposts/feed` 多级缓存，不是个性化 Hybrid Feed；2026-07-29 的推拉测试以功能和 Redis 微基准为主。它们作为架构演进背景保留，不与后续 Hybrid QPS 直接相减。

## 3. 架构从哪里开始，改成了什么

### 3.1 初始 Hybrid 读路径

```text
GetUserFeed
  -> Relation.ListFollowing
  -> FeedRouteCache：复用大 V 分类，但仍先获取完整关注列表
  -> Redis Inbox
  -> Redis Pipeline：逐批读取 BigV Outbox
  -> Top-N merge + deduplicate + balance
  -> Redis MGET / DB：装载 FeedItem
  -> response
```

初始路径已经有大 V 分类缓存、Singleflight、BigV Pipeline 和固定容量 Top-N，但存在两个主问题：

- FeedRouteCache 不能消除 Relation RPC；Relation 仍在每个请求进入 MySQL/sqlx 查询和反射映射。
- Inbox、BigV Outbox、FeedItem hydrate 至少形成约三段 Redis dependency round trips。

### 3.2 演进后的读路径

```text
GetUserFeed
  -> 读取 relation epoch / safety epoch
  -> PageCache L1
       fresh hit -> 直接返回完整第一页
       stale hit -> 合法时返回 stale，并投递有界异步刷新
       miss      -> PageCache L2
                     hit  -> 回填 L1
                     miss -> 同 full key Singleflight 冷计算
                                -> RouteSnapshot L1
                                     hit  -> 复用 followings / BigV / set
                                     miss -> Relation + Counter，版本化回填
                                -> Pipeline #1：Inbox + 首批 BigV Outbox
                                -> Pipeline #2..M：剩余 BigV
                                -> fixed Top-N / dedup / balance
                                -> FeedItem MGET / DB
                                -> Protobuf Page 回填 L2/L1
```

写侧和一致性配套包括：

- Follow/Unfollow 经 outbox/Kafka 驱动 relation epoch 增长，旧 RouteSnapshot 自然失效。
- 删除、转私密、下架、置顶和敏感元数据变化同步 bump content-safety epoch，并用 Kafka outbox 补偿。
- 发布新帖允许短 TTL 内短暂不可见，这是 Feed 已批准的最终一致性取舍。
- epoch 不可信或 SafetyPending 时旁路页面缓存；敏感状态不会为了性能返回无法验证的 stale 页面。
- Singleflight 合并同 key 冷回源；SWR 使用 32 个刷新 worker、1024 队列和 2 秒 loader timeout，过载时可观测并判 degraded。

## 4. 测试要回答的问题

压测不是只问“最高 QPS”，而是拆成六个问题：

1. Feed 内核不经过 Gateway 时，单机读路径的容量和拐点在哪里？
2. 真实 Gateway 入口加入 JWT、HTTP 编解码和 RPC hop 后损耗多少？
3. 20、约 1,200、20,000 个读者下，缓存局部性如何改变收益？
4. cold、L1 warm、L2 warm、集中失效和自然状态下，冷计算与刷新是否可控？
5. 并发提升时，吞吐、尾延迟、Redis ops、CPU/RSS 和依赖调用如何共同变化？
6. 优化是否破坏 Follow/Unfollow、删除/私密、部分 Redis 失败、Kafka 重试和上下文取消语义？

## 5. 压测设计

### 5.1 两种入口

| 入口 | 链路 | 用途 |
|---|---|---|
| RPC | Load Generator → KnowPost gRPC | 隔离 Feed 计算内核，寻找读路径上限 |
| Gateway | Load Generator → HTTP/JWT → Gateway → KnowPost RPC | 测真实用户入口，包含验签、HTTP/JSON 和跨服务 hop |

RPC 和 Gateway 必须分别报告。把 RPC 10K 说成 Gateway 10K，或把 Gateway 2K 说成 Feed 内核退化，都会误判。

### 5.2 三档读者基数

| Cardinality | Readers | 目的 |
|---|---:|---|
| hot | 20 | 验证热点第一页、L1 命中和单飞上限 |
| distributed | 约 1,200 | 代表常规分散用户，验证 L2 页面复用和冷路径容量 |
| high | 20,000 | 验证低复用、高 key 基数、L1 预算和刷新队列边界 |

每个数据集由固定 seed 和 manifest 生成。Push、Pull、Hybrid 不能共用 manifest，因为各策略的 Inbox/Outbox 物化布局不同。

### 5.3 场景与缓存前态

| 场景/状态 | 测量前动作 | 验证目标 |
|---|---|---|
| hot-read | 在小读者集合重复读取 | 热 key 上限 |
| distributed-read | 在目标读者集合轮转 | 常规/高基数工作集 |
| deep-page | 请求 page > 1 | page-number 分页的候选放大成本 |
| cold | bump safety epoch 并等待 epoch L1 收敛 | 强制整页 miss |
| l1-warm | cold 后逐读者预热，紧接测量 | L1 完整页面上限 |
| l2-warm | 预热后等待 L1 Fresh 过期 | L2 命中成本 |
| expire-together | 预热后等待 5.2 秒 | 集中失效、SWR、刷新队列 |
| natural | 不干预 | 观察自然命中，但不适合严格归因 |

显式 warm 会校验准备时长，避免第一批 key 在真正采样前已过期。Control 在 PageCache 关闭时只能标为 cold；Treatment auto 映射为 hot=l1-warm、distributed=l2-warm、high=cold。

### 5.4 并发、时长和容量定义

完整探索档为 c32/c64/c128/c256；必要时追加 c512。为了缩短后期长测：

1. 先用 10 秒 scout 找候选点；
2. 只对拐点或相邻档执行 10 秒预热 + 60 秒采样 × 3；
3. 三轮取中位数，短 scout 不进入容量结论；
4. 首轮若 degraded 或报告不完整，立即停止该档剩余 trial，下降一档验证。

`stable c` 是吞吐增幅首次低于 10% 且 P95 明显上升之前的工作档；`knee c` 是对应拐点。最高无错误 QPS 只表示被测上界，不自动成为推荐并发。

SLA 数字只作为参考。错误、超时、错误内容、降级或证据缺失比是否“达到某条 SLA 线”优先级更高。

### 5.5 Control / Treatment

- Control：关闭 RelationEpoch consumer、RouteSnapshot、Combined Pipeline、PageCache 等演进开关，作为旧路径基线。
- Treatment：开启版本化路由、合并 Pipeline、页面缓存、安全 epoch、Singleflight/SWR。
- 严格代码 A/B 应保持两边均为 cold、同入口、同并发、同拓扑和同 manifest。
- cold Control → warm Treatment 只回答“目标运行态能达到什么”，不能把全部差异归功于某一个函数改动。

WP5 和 WP6 使用逐阶段单变量方式，因此可以分别归因 RouteSnapshot 和 Combined Pipeline。WP11 的目标态矩阵更接近产品运行方式，但必须连同缓存前态一起解释。

## 6. 压测是怎么执行的

一次正式轮次按以下顺序进行：

1. 固定 Git commit 和未提交 patch 内容哈希，记录本地 exe SHA 或 Docker image ID。
2. 验证 strategy、FeaturePreset、服务地址、资源口径、GOMAXPROCS、manifest、seed 和真实 reader 数。
3. 验证进程 PID/exe/start time 或容器 ID/restart count，防止测到旧二进制或跨代进程。
4. 执行普通 warmup，但不计入采样。
5. 准备显式 cache-state，并校验 L1/L2/stale 时间预算。
6. 采集测量前 Prometheus、Redis identity/counters、Kafka lag、MySQL 和进程/容器快照。
7. 按固定并发连续施压 60 秒，记录成功延迟分布和客户端 CPU。
8. 停止施压后等待 Kafka lag 恢复，采集末态并计算本轮 counter delta。
9. 检查 safety epoch 是否在纯读窗口变化、Redis 是否重启/evict/reject、服务是否重启、Feed 是否 degraded。
10. 只有 `complete=true` 的三轮才进入 comparison；按中位数汇总 QPS 和 P95/P99。

典型正式命令如下：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase run `
  -run-id $distributed -report-run-id "$distributed-treatment-formal" `
  -reader-cardinality distributed `
  -topology compose-middleware-local-services `
  -cache-state l2-warm -feature-preset treatment `
  -strategy hybrid -entry rpc -scenario distributed-read `
  -concurrency 64 -requests 0 -duration 60s -warmup 10s -trials 3
```

Compose 镜像源出现构建/下载问题时，本项目按既定约束归类为 Docker registry/mirror source 问题，不归因于网络。后期使用“既有中间件容器 + 本地项目服务”拓扑，并为该拓扑独立建立 Control/Treatment；不同拓扑的绝对值不做严格 A/B。

正式本地混合拓扑使用 `benchmark-error-only`：保留 error，关闭 go-zero stat、RPC Stat、Gateway access log 和 SQL statement info。关闭高频请求日志后 Gateway stdout 保持约 6.7KB，不再随请求量增长。

### 6.1 采集的指标

- 请求：success/failed/timeout、QPS、P50/P90/P95/P99/Max。
- Feed stages：relation、counter、route、inbox、bigv pipeline、merge/dedup、hydrate、total。
- 依赖：每成功请求 Relation/Counter/Redis/MySQL 调用数、cold compute。
- PageCache：L1/L2 Fresh、Stale、Miss、Fresh ratio。
- SWR：queue、active、pending 与 refresh outcome。
- 系统：KnowPost/Gateway/Relation/Counter CPU、RSS、启动时间和重启；Redis ops/delta/identity/eviction/rejection；MySQL threads；Kafka lag。
- 负载端：客户端进程 CPU；归一化达到默认 90% 即判客户端饱和。

### 6.2 报告失效门禁

以下任一情况都会令轮次退出正式容量统计：

- 请求失败、超时、内容正确性失败或 Feed stage 非 success。
- `feed_degraded`、refresh queue full 等依赖降级。
- 客户端 CPU 饱和。
- Prometheus、Redis、Kafka、MySQL 或进程/容器指标缺失。
- Redis identity 变化、eviction/rejection、RDB `MISCONF`。
- 进程/容器重启或实际二进制身份变化。
- 纯读期间 safety epoch 改变。
- Kafka 在恢复窗口内没有回到 lag=0。
- Git/manifest/seed/拓扑/Feature Flags/缓存状态等兼容指纹不一致。

## 7. 性能演进时间线

### 7.1 早期公共 Feed：多级缓存价值验证

2026-07-03 的公共 Feed Gateway 实验使用 5,000 条帖子和 k6 ramping-vus，应用层瞬时 500 VU 峰值窗口为 2,112.3 QPS、失败率 0.47%，直连 MySQL 对照为 963.2 QPS、失败率约 1.04%；稳态数据库查询约 2 QPS。它证明 L1/L2 + Singleflight 能将用户 QPS 与 DB QPS 解耦，但不是个性化 Hybrid 的基线，也不是零错误容量点。

### 7.2 早期推拉功能和微基准

2026-07-29 验证了阈值 1,000 的语义：普通作者发帖异步 fanout，大 V 只写 Outbox，读取时合并 Inbox + BigV Outbox。7 个集成用例全部通过。

代表性 Redis 微基准：

- 1,000 粉丝批量扇出约 35ms；
- ZAdd + Trim 约 63.68μs，约 15.7K ops/s；
- Pipeline 批量写 100 条约 250.8μs。

这些数字是组件微基准，不能推导端到端 Feed QPS。早期文档中的“生产就绪/Twitter 级别”结论没有正式容量、故障和高基数证据支撑，本报告不沿用该表述。

同一时期的三策略探索矩阵提供了方向性证据：RPC distributed-read 峰值无错 QPS 约为 Push 2,915.53、Hybrid 1,599.55、Pull 621.91。它符合 fanout-on-write 的读放大取舍：Push 读最轻，Pull 读最重，旧 Hybrid 位于中间。由于当时采样、重复和 mutation 数据隔离尚未达到后续门禁，这组结果不用于当前容量承诺。

### 7.3 初始 Hybrid 基线与 Profile

优化前的 30 秒 Hybrid RPC c128 稳态记录为 2,396.03 QPS、P95 68.55ms、P99 77.85ms，KnowPost CPU 298.59%、Relation CPU 237.50%、Redis 19,981 ops/s。

WP1 同口径 Profile 检查点为 2,636.79 QPS、P95 62.435ms。CPU Profile 指向：

- KnowPost `GetUserFeed` 约 49%-51%；`loadFeedItems` 约 24%；Hybrid Redis 读约 16%-17%。
- Relation `ListFollowing -> PageActive` 约 45%-52%，主要是 MySQL/sqlx、行映射、反射和响应分配。
- block/mutex 没有显著样本，不支持把锁竞争列为 P0。

### 7.4 WP2：严格基线与并发拐点

Observer 开启只造成 QPS -1.26%、P95 +1.80%，低于 5% 仪表开销门禁。

| C | QPS median | P95 | 结论 |
|---:|---:|---:|---|
| 32 | 1,945.29 | 20.847ms | 低并发 |
| 64 | 2,327.03 | 35.336ms | 吞吐继续增长 |
| 128 | 2,590.31 | 64.343ms | 推荐稳定档 |
| 256 | 2,793.47 | 123.603ms | knee，边际收益下降 |
| 512 | 2,801.37 | 382.092ms | 仅比 c256 +0.28%，P95 +209.1% |

c128 每成功请求约有 1 次 Relation、1 次 Inbox Redis、1 次 BigV Pipeline、1 次 FeedItem MGET，Redis dependency 合计约 3.01/request。Relation stage mean 22.431ms，是第一优先级。

### 7.5 WP5：RouteSnapshot 消除 Relation 热路径

版本化 RouteSnapshot key 包含 user 和 relation epoch；TTL 最长 5 秒，冷 miss 使用同 key Singleflight，Ristretto 写入在释放 singleflight 前等待可见。

| C | WP2 QPS | WP5 QPS | QPS 变化 | WP2 P95 | WP5 P95 |
|---:|---:|---:|---:|---:|---:|
| 32 | 1,945.29 | 3,572.68 | +83.66% | 20.847ms | 12.050ms |
| 64 | 2,327.03 | 3,978.74 | +70.98% | 35.336ms | 22.054ms |
| 128 | 2,590.31 | 3,896.52 | +50.43% | 64.342ms | 49.301ms |
| 256 | 2,793.47 | 3,884.14 | +39.04% | 123.603ms | 88.936ms |

Relation calls/request 降低约 99.9%。新峰值位于 c64，说明 Relation/MySQL 不再是热路径主瓶颈；每请求约 3 次 Redis 往返和 FeedItem hydrate 成为下一约束。

### 7.6 WP6：Combined Pipeline 达到可重复 10K

Inbox 和第一批 BigV Outbox 合并进一个 Pipeline，剩余 BigV 继续有界分批；固定 Top-N 跨批次维护。优化降低的是网络往返，不是 Redis command 数。

| C | WP5 QPS | WP6 QPS | QPS 变化 | WP5 P95 | WP6 P95 |
|---:|---:|---:|---:|---:|---:|
| 32 | 3,572.68 | 7,996.33 | +123.82% | 12.050ms | 5.487ms |
| 64 | 3,978.74 | 9,866.82 | +147.99% | 22.054ms | 9.171ms |
| 128 | 3,896.52 | 10,661.31 | +173.61% | 49.301ms | 18.112ms |
| 256 | 3,884.14 | 10,587.21 | +172.58% | 88.936ms | 33.068ms |

相对 WP2，c128 QPS +311.58%、P95 -71.85%。若使用更早的 30 秒 2,396.03 QPS 记录作辅助参考，则到 10,661.31 是 4.45 倍、+344.96%；由于窗口和观测口径不同，不作为正式阶段增益。

c128 Redis 峰值约 79.6K-80.5K ops/s，KnowPost CPU 峰值约 488%-533%。c256 Redis 达约 81K ops/s，吞吐反而下降、P95 上升。因此 10K 不是“继续加并发”得到的，而是降低单位请求 RTT 得到的。

### 7.7 WP7-WP9：页面缓存、一致性和防击穿

- WP7 建立 page1/size20 的 L1/L2 完整页面缓存，L2 使用 Protobuf，key 包含 strategy、user、relation epoch、safety epoch、page、size。
- WP8 为删除、私密、下架等敏感变化建立同步 safety epoch 和 outbox 补偿；真实验证中私密内容约 95.4ms、删除内容约 92.5ms 从 Feed 不可见。
- WP9 引入同 full key Singleflight、SWR、有界 worker/queue 和 epoch revalidation。

WP9 的 400 请求热点功能样本达到 10,694.99 QPS、P95 1.3515ms；SWR 40 请求 P95 0.804ms。它们只证明功能热路径，不是持续容量。

### 7.8 WP10：压测工具与 32K 短样本

WP10 把三档 cardinality、缓存状态、客户端 CPU、进程身份、Redis identity、Git patch、服务二进制、safety epoch 和指标完整性变成 fail-closed 报告字段。

5 秒 hot+l1-warm+RPC+c32 功能样本为 32,513.61 QPS、P95 1.618ms、Fresh 99.82%。这个数没有消失；它是极小工作集、已热缓存、5 秒窗口的瞬时热路证据。正式 60 秒 × 3 不能用它替代，因为长期运行会包含 TTL、刷新、调度、GC 和共享资源波动。

### 7.9 WP11：三种局部性下的正式容量

#### hot：页面缓存收益明确

同一本地混合拓扑、c32、60 秒 × 3：

- RPC hot-read：3,028.91 → 8,222.49 QPS，+171.47%；P95 13.401 → 5.309ms，-60.38%。
- Gateway hot-read：2,379.29 → 4,930.08 QPS，+107.21%；P95 17.029 → 9.835ms，-42.25%。
- RPC deep-page：2,829.27 → 2,932.98 QPS，+3.67%；P95 基本不变。

hot Control 为 cold，Treatment 为 l1-warm，所以前两项是目标运行态对比，不是单项代码 A/B。deep-page 不走第一页页面缓存，直接暴露了 page-number 路径问题。

#### distributed：L2 命中能把计算路径近乎消除

RPC c64 的 cold Control 为 6,797.75 QPS，l2-warm Treatment 为 102,075.49 QPS，P95 14.06 → 1.16ms。Fresh=100%，cold compute≈0；Redis 峰值从约 60K 降到 13K ops/s，即使 QPS 提升 15.02 倍。

这说明收益来自跳过 Relation、Inbox/BigV、merge、hydrate 的整个页面计算，而不是 Redis 单条命令提速。它是目标态上限，不是 cold-vs-cold 单变量比较。

Gateway c16 只从 1,218.99 增至 1,529.10 QPS（+25.44%），P95 反而从 23.02 增至 26.34ms。Fresh 约 92%，但 Gateway/JWT/HTTP 与共享资源已经限制入口吞吐。

这组 Gateway 三轮波动明显：Control 为 2,048.25 / 1,090.29 / 1,218.99 QPS，Treatment 为 1,608.90 / 1,161.09 / 1,529.10 QPS。三轮都通过完整性门禁，所以本报告保留中位数，但它同时说明本地混合拓扑的 Gateway 对共享主机调度和缓存恢复很敏感；在做同窗 Profile 与资源隔离复验前，不应把 25.44% 当成稳定的线性增益。

#### high：低复用使收益收缩并出现刷新过载

- RPC Control c32：6,162.47 QPS，P95 7.55ms。
- RPC Treatment 稳定点 c16：9,477.51 QPS，P95 3.73ms，Fresh 约 72%，cold compute 0.40/request。
- Treatment c32 首轮 7,768.72 QPS，但出现 `feed_degraded` 和 `Feed PageCache refresh queue is full`，该轮无效并停止剩余 trial。
- Gateway c8：1,402.80 → 1,407.80 QPS，吞吐近乎不变，P95 11.20 → 15.16ms。

20,000 个读者让单用户在 TTL 内的重复访问减少，L1/L2 页不再持续 Fresh；恢复冷路径和异步刷新同时消耗 CPU、Redis 和内存。Treatment RPC RSS 中位峰值约 225.4MiB，对照约 82.5MiB，也说明高基数页面缓存有真实内存成本。

high Gateway 同样有较大轮间波动：Control 为 2,250.27 / 1,402.80 / 1,278.31 QPS，Treatment 为 2,197.33 / 1,355.07 / 1,407.80 QPS。两组中位数接近且尾延迟回退，因此当前可靠结论是“未证明有吞吐提升”，而不是根据某一个首轮峰值声称 Gateway 达到 2.2K。

WP11 正式点的资源侧中位证据如下。CPU 是三轮“采样最大值”的中位数，100% 约等于一个逻辑核，不是全程平均 CPU：

| 数据规模/入口 | 预设 | Redis max ops/s | 主要进程 CPU max | 主要进程 RSS max |
|---|---|---:|---:|---:|
| distributed RPC | Control | 60,204 | KnowPost 377.3% | 71.8MiB |
| distributed RPC | Treatment | 13,377 | KnowPost 643.9% | 95.4MiB |
| distributed Gateway | Control | 14,465 | Gateway 219.8% | 48.9MiB |
| distributed Gateway | Treatment | 13,268 | Gateway 465.6% | 49.3MiB |
| high RPC | Control | 67,955 | KnowPost 388.5% | 82.5MiB |
| high RPC | Treatment | 63,908 | KnowPost 434.3% | 225.4MiB |
| high Gateway | Control | 26,911 | Gateway 191.9% | 48.5MiB |
| high Gateway | Treatment | 28,132 | Gateway 194.5% | 48.8MiB |

distributed Treatment RPC 在 QPS 大幅上升时 Redis ops 下降但 KnowPost CPU 上升，说明请求主体已经变成极短的缓存查找、Protobuf/gRPC 返回与调度开销；high Treatment 则没有消除 Redis 工作，且用更多内存换取了有限 Fresh 命中。

### 7.10 WP11：mutation RPC 正式 A/B

2026-08-27 在全新 checkpoint 上完成 RPC-only publish c2、90/10 c4、80/20 c4 的 Control/Treatment，双方各 9 轮，均为 10 秒预热 + 60 秒采样 × 3。18/18 报告完整，失败/超时为 0，所有试后 restore 和 Kafka drain 成功。

- publish：6.61 → 6.33 完整流程 QPS（-4.33%），P95 398.19 → 388.98ms（-2.31%）。
- 90/10：总操作 QPS 65.18 → 61.37（-5.85%）；读 P95 3.23 → 2.64ms（-18.43%），发布 P95 729.46 → 779.37ms（+6.84%）。
- 80/20：总操作 QPS 34.26 → 30.05（-12.28%）；读 P95 3.21 → 3.19ms（-0.81%），发布 P95 712.67 → 828.74ms（+16.29%）。

Treatment 将 Relation/request 降低 85%–93%，90/10 的 Redis dependency/request 下降约 37%，说明读侧重复工作减少。但 PageCache Fresh 只有约 50% 和 27%，完整发布平均耗时也上升，读侧节省没有转化为端到端吞吐增长。正式结论是“读依赖和部分读尾延迟改善，但 mutation 总吞吐回退”；缓存失效、回填、异步事件或共享依赖竞争的各自贡献尚未被 bundled A/B 隔离。

用户已决定不再追加 Gateway 性能矩阵。旧 Gateway 数据继续证明入口层存在独立瓶颈，但本次严格 A/B 只包含 18 份 RPC 报告。完整报告见 [WP11 mutation RPC A/B](2026-08-27-feed-wp11-mutation-rpc-ab.md)。

## 8. 性能为什么提升

### 8.1 消除 Relation RPC，而不是优化一个小函数

RouteSnapshot 将每请求 Relation 调用从约 1 降到约 0.001。收益包含一次 RPC RTT，以及 Relation 内部 MySQL 查询、sqlx 反射映射、切片/响应分配；因此在 c32-c256 全档得到 39%-84% QPS 增长。

### 8.2 Pipeline 消除网络往返，而不是减少 Redis 命令

Combined Pipeline 把 Inbox 和首批 BigV 合并，一次请求的 Redis dependency round trips 约 3 → 2。Windows + Docker 本地链路对 RTT 很敏感，因此 33% 往返减少带来 124%-174% QPS 增长。Redis ops/s 并未同比下降，反而随总 QPS 升至约 80K，这与实现机制一致。

### 8.3 完整页面缓存跳过整个计算 DAG

Fresh hit 不再执行 Relation、Counter、Inbox/BigV、merge/dedup 和 hydrate。distributed l2-warm 中，QPS 15 倍时 Redis ops 反而降低，证明它不是局部 micro-optimization，而是把高成本计算从请求路径删除。

### 8.4 Singleflight/SWR 的主要价值是尾延迟与稳定性

Singleflight 合并相同 key 冷 miss，SWR 允许在 epoch 合法时返回短 stale 并异步刷新，有界队列防止无限 goroutine。它们不保证所有场景提高平均 QPS，但能限制击穿、暴露刷新过载并维持一致性边界。

### 8.5 日志优化去除了测量污染，但没有消除 Gateway 主瓶颈

关闭 access/stat/SQL info 后不再发生每请求持久化日志，避免磁盘和格式化开销污染压测。high Gateway 仍基本不提升，说明剩余问题主要在 JWT/HTTP/RPC hop、低缓存复用和冷路径恢复，而不是日志 I/O。

## 9. 性能为什么下降或没有提升

### 9.1 超过并发拐点后进入排队区

WP2 c256→c512 吞吐只增 0.28%，P95 增 209.1%；WP6 c128→c256 吞吐下降，P95 继续增加；high Treatment c32 触发刷新队列满。提高并发只增加排队、Context 和内存压力，不会创造服务能力。

### 9.2 高基数削弱缓存局部性

20,000 readers 在相同时间窗内分散请求，单 key 很难在 L1/L2 Fresh TTL 内再次访问。结果是 Fresh 从 distributed 的 100% 降到 high RPC 的约 72%、high Gateway 的约 23%，cold compute 分别上升到 0.40 和 0.81/request。

### 9.3 Gateway 链路有额外固定成本

Gateway 包含 JWT 验签、HTTP 解析与 JSON 编解码、一次 Gateway→KnowPost RPC hop，并与其他本地服务共享 CPU。高基数 Gateway 约 1.4K QPS 的平台说明，即使 Feed 内核更快，入口层仍可能成为独立瓶颈，需要同窗 Profile 后再优化。

### 9.4 深分页重复扩大候选集

当前接口使用 page/size。深页通常需要读取和归并前 `page*size` 个候选，再丢弃前页；PageCache 只覆盖 page1/size20。正式 deep-page c32 只有 +3.67% QPS，说明整页缓存没有触及该成本。Cursor/seek 分页是下一阶段正确方向。

### 9.5 短样本高估长期容量

5 秒热样本 32.5K QPS 主要处于刚预热、几乎 100% Fresh 的窗口；60 秒正式运行会跨越 TTL、刷新周期、GC 和资源波动。短样本适合验证路径，不能替代三轮稳态容量。

### 9.6 不同拓扑不能直接相减

WP6 是 Compose 全栈；WP11 后期是中间件 Compose + 项目服务本地运行。后者由 Docker 镜像源问题触发，但重新建立了自身 Control/Treatment。容器与本地绝对 QPS 差异不能直接归因于 Docker 本身。

## 10. 测试中排除的异常与经验

### 10.1 Redis RDB `MISCONF`

Windows Redis 后台保存失败曾触发 `MISCONF` 并停止写入，污染 cache-state。所有受影响报告均作废。有效 WP11 轮次在完全开发环境临时关闭 RDB schedule 和 `stop-writes-on-bgsave-error`；原配置保存于 `.tmp/feed-local/redis-benchmark-original.json`，全部性能工作结束后必须恢复。

### 10.2 Docker 镜像源问题

无法可靠重建/下载部分镜像被明确归为 Docker registry/mirror source 问题，不是网络故障。后续复用已构建的 Kafka/etcd 等中间件容器，项目服务本地编译运行，并记录 PID、启动时间、exe SHA 和地址。切换拓扑后没有复用旧基线。

### 10.3 Docker API / 指标缺失

部分 scout 因沙箱无法访问 Docker API，导致 docker/kafka metrics 缺失，报告 incomplete。这些轮次没有进入容量表；以具备采集权限的复跑替代。

### 10.4 Gateway token 数据准备污染

20,000 readers 如果逐个登录，会把 bcrypt、User RPC 和日志开销引入测量前态。后期使用本地 RS256 私钥批量签发测试 token，但请求仍经过真实 Gateway JWT 验签，因此只消除了 setup 污染，没有绕过被测认证路径。

### 10.5 写入/混合数据隔离已补齐

早期 publish、90/10、80/20 会持续增加 MySQL 帖子、Outbox 和 Feed ZSET，只等待 Kafka lag=0 不能恢复初态，因此旧数字仍是 non-comparable probe。2026-08-18 后实现了 trial 级 MySQL/Redis/Kafka/epoch checkpoint/restore，2026-08-27 新鲜基线上的 18 轮 RPC A/B 才是正式 mutation 结论。

## 11. 当前性能结论

### 11.1 已经被证明的能力

- Hybrid 冷计算内核：Compose 全栈、约 1,200 readers、RPC、page1/size20，在 c128 可重复达到 10,661 QPS，P95 18.112ms。
- 热点完整页：本地混合拓扑、20 readers、RPC c32 稳态约 8.22K QPS，Gateway c32 约 4.93K QPS。
- 分散 L2 完全命中目标态：RPC c64 约 102K QPS，P95 1.16ms；这是缓存上限，不是冷读承诺。
- 20,000 readers：RPC 稳定点约 9.48K QPS；Gateway 约 1.4K QPS，当前优化几乎没有带来吞吐增益。
- mutation RPC：publish c2 约 6.3–6.6 个完整四阶段流程/秒；90/10 与 80/20 Treatment 总吞吐分别回退 5.85%、12.29%，但 90/10 读 P95 改善 18.44%。

### 11.2 不能宣称的内容

- 不能宣称 Gateway 已达到 10K。
- 不能把 32.5K 五秒功能样本或 102K L2-warm 上限当成通用单机容量。
- 不能把不同拓扑、入口、读者基数、缓存前态和并发档的数字做严格百分比对比。
- 不能宣称 Treatment 提升了写入或混合总吞吐；正式 RPC A/B 显示 -4% 至 -12% 的吞吐变化。
- 不能从本地单机数字外推生产集群 SLA。

此外，当前 Windows Go 环境为 `CGO_ENABLED=0` 且没有 C 编译器，无法运行 `go test -race`。并发语义由确定性单测、重复测试、真实 Redis/Kafka 集成和高并发长测覆盖，但本报告不声称 Race Detector 已通过。

## 12. 下一阶段建议与验证顺序

1. 实现稳定 Cursor/seek 分页：使用 `(publish_time, post_id)` 等全序游标，验证同时间戳 tie-break、插入/删除、重复/遗漏、旧 page API 兼容；重新跑首页与连续翻页 c32 及相邻档。
2. Profile 完整发布的 metadata、confirm、commit：拆分数据库事务、Outbox、同步依赖和连接池等待；后续容量复验只跑 RPC 稳定档及一个必要相邻档。
3. 处理 high-cardinality refresh admission：按证据评估 TTL 分层、热点 reader admission、刷新队列/worker 配置或有条件 Feed Head 物化；以 c32 不再 degraded 为门禁，而不是只扩大队列。
4. 根据 reader mutation 率和近期 Fresh 命中决定 PageCache 准入或旁路，避免 80/20 下低复用页面反复编码、回填和失效。
5. 恢复 Redis persistence 原配置，并执行 SET/GET/DEL 与服务冒烟。Gateway 后续性能矩阵不再执行。

## 13. 原始证据索引

设计与方法：

- [全面压测设计](../specs/2026-08-14-feed-performance-load-test-design.md)
- [Hybrid 读路径演进设计](../specs/2026-08-15-feed-hybrid-read-performance-evolution-design.md)
- [压测执行指南](../../../cmd/loadtest/FEED_EVOLUTION_GUIDE.md)

阶段报告：

- [WP1 pprof 检查点](2026-08-15-feed-hybrid-wp1-pprof-checkpoint.md)
- [WP2 Phase 0 正式基线](2026-08-15-feed-hybrid-wp2-observability-phase0.md)
- [WP3 Feed Epoch Store](2026-08-15-feed-hybrid-wp3-feed-epoch-store.md)
- [WP4 Relation Epoch Consumer](2026-08-15-feed-hybrid-wp4-relation-epoch-consumer.md)
- [WP5 RouteSnapshot](2026-08-15-feed-hybrid-wp5-route-snapshot.md)
- [WP6 Combined Pipeline](2026-08-15-feed-hybrid-wp6-combined-pipeline.md)
- [WP7 PageCache Foundation](2026-08-15-feed-hybrid-wp7-page-cache-foundation.md)
- [WP8 Content Safety Epoch](2026-08-15-feed-hybrid-wp8-content-safety-epoch.md)
- [WP9 Singleflight/SWR](2026-08-16-feed-hybrid-wp9-singleflight-swr.md)
- [WP11 mutation RPC A/B](2026-08-27-feed-wp11-mutation-rpc-ab.md)
- [WP10 压测工具与拓扑](2026-08-16-feed-hybrid-wp10-loadtest-topology.md)
- [WP11 distributed/high 正式容量](2026-08-17-feed-hybrid-wp11-capacity.md)

历史背景：

- [公共 Feed 多级缓存压测](../../loadtest-knowpost-feed.md)
- [早期推拉结合功能与微基准](../../feed-push-pull-test-report.md)
- [早期三策略探索汇总](../../../results/feed-loadtest/feedbench-20260814c/comparison.md)

正式结果命名空间：

- `feed-wp2-phase0-strict-20260815`
- `feed-wp5-route-enabled-20260815`
- `feed-wp6-combined-enabled-20260815`
- `feed-wp11-hot-control-cold-20260816c`
- `feed-wp11-hot-treatment-auto-20260816a`
- `feed-wp11-distributed-control-fixed-formal-nolog-20260817a`
- `feed-wp11-distributed-treatment-fixed-rpc-formal-nolog-20260817a`
- `feed-wp11-high-control-fixed-rpc-formal-nolog-20260817a`
- `feed-wp11-high-treatment-fixed-rpc-c16-formal-nolog-20260817a`
- `feed-wp11-high-control-fixed-gateway-formal-nolog-20260817a`
- `feed-wp11-high-treatment-fixed-gateway-c8-formal-nolog-20260817a`

## 14. 一句话总结

这次性能演进的核心不是“把 Go 代码写得更快”，而是依次删除三类重复工作：每请求关系查询、多次 Redis RTT、重复生成完整第一页；前两项把可重复冷路径从约 2.6K 推到 10.66K QPS，第三项让高复用场景进一步数量级提升，同时也通过 20,000-reader 和 Gateway 长测暴露出缓存局部性、刷新容量与深分页才是下一阶段真正需要解决的问题。
