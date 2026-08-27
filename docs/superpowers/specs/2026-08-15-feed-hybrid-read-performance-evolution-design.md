# Feed 推拉混合读路径性能演进设计

日期：2026-08-15

状态：已批准并完成既定 RPC 交付（WP1-WP10、WP11 纯读、Cursor、mutation RPC 正式 A/B）；后续 Gateway 性能矩阵按用户决策取消

适用范围：KnowPost 个性化 `GetUserFeed`、Relation 路由依赖、Redis Feed 数据、Gateway 对照入口和 Feed 压测工具

## 1. 决策摘要

本设计采用渐进式分层读路径，不推翻现有推拉混合架构：

1. 先建立可解释的 CPU、分阶段延迟、缓存和依赖调用基线。
2. 将现有“只缓存大 V 分类”的路由缓存升级为完整 Feed 路由快照，消除热路径每次调用 Relation RPC 的成本。
3. 合并 Inbox 与大 V Outbox 的 Redis Pipeline，继续降低 Hybrid 冷路径成本。
4. 只对个性化 Feed 第一页增加独立的 L1 + Redis L2 最终结果缓存，并加入 epoch 校验、Singleflight、TTL 抖动、stale-while-revalidate 和有界刷新。
5. Feed Head 物化仍只在前述优化无法满足目标时进入；Cursor/seek 分页作为独立优化，不依赖 Feed Head，在第 19 节按专用深数据集和严格 A/B 单独实施与验收。

单机 10,000 QPS 是“热点读者、第一页、缓存充分预热”场景的阶段目标，不是随机冷用户的无条件容量承诺。每个阶段必须通过同拓扑、同资源、同数据集的对照压测后才能继续。

## 2. 已验证现状

### 2.1 当前架构

KnowPost 已实现三种可切换策略：

- `push`：作者发布后通过 Kafka 向粉丝 Inbox 扇出。
- `pull`：作者只写 Outbox，读取时按关注作者归并。
- `hybrid`：普通作者走推，大 V 走拉，读取时合并 Inbox 与大 V Outbox。

当前 Hybrid 读路径大致为：

```text
GetUserFeed
  -> Relation.ListFollowing
  -> FeedRouteCache 命中时复用大V分类，否则调用 Counter
  -> Redis Inbox
  -> Redis Pipeline 读取 BigV Outbox
  -> Top-N Merge + Deduplicate
  -> FeedItem L1/L2/DB 批量装载
  -> FeedPage
```

当前路由缓存仍需先调用 Relation 获取完整关注列表，再比较关注快照是否相同，因此只消除了部分 Counter RPC，没有消除 Relation RPC。

### 2.2 已完成优化

- Feed 路由分类缓存与独立 Context 的 Singleflight。
- 大 V Outbox `ZREVRANGE WITHSCORES` Pipeline，每批最多 128 个作者。
- 不支持批量接口时保留最大 16 并发的降级路径。
- 固定容量 Top-N 归并。
- 删除成功请求的高频日志，保留失败日志。
- Counter 故障时不缓存错误降级结果。

### 2.3 性能基线

最终审查修复后的 Hybrid RPC c128、30 秒稳态结果：

| 指标 | 当前基线 |
|---|---:|
| 成功 QPS | 2,396.03 |
| 平均延迟 | 53.40 ms |
| P95 | 68.55 ms |
| P99 | 77.85 ms |
| 失败 | 0 |
| KnowPost CPU 峰值 | 298.59% |
| Relation CPU 峰值 | 237.50% |
| Counter CPU 峰值 | 6.63% |
| Redis 峰值 ops/s | 19,981 |

该结果说明：

- Counter 和重复 Redis 操作已经明显下降。
- KnowPost 与 Relation 成为主要 CPU 消费者。
- 继续增加并发已出现边际收益递减，瓶颈是单位请求成本，不是客户端并发不足。
- 现有 c256 结果采样窗口过短，不能作为容量结论，必须在 Phase 0 重跑。

基线证据：

- `results/feed-loadtest/feedopt-20260814-route-pipeline-nolog/OPTIMIZATION_REPORT.md`
- `results/feed-loadtest/feedopt-20260814-reviewfix/hybrid-rpc-steady-read-c128/report.md`

## 3. 目标与非目标

### 3.1 目标

- 降低每次 Hybrid 读取需要执行的 Relation、Counter 和 Redis 操作数。
- 保留纯推、纯拉、混合三种策略的业务语义和对照能力。
- 建立分层缓存，使热点第一页不再重复执行完整 Feed 计算。
- 在不牺牲删除、可见性和取关收敛要求的前提下，接受 Feed 新内容短暂不可见。
- 为每项优化提供 Feature Flag、自动化测试、指标、压测门禁和回滚路径。
- 在当前单机开发环境中验证热点第一页 10,000 QPS 的可行性。
- 同时报告高基数、低命中场景的真实能力，防止用热点缓存数据冒充通用容量。

### 3.2 非目标

- 不承诺所有读者分布和所有分页都达到 10,000 QPS。
- 不在第一版缓存深分页。
- 不在第一版引入推荐模型、复杂排序或用户商业化分层算法。
- 不在第一版给所有用户预计算 Feed Head。
- 不为页面缓存主动遍历大 V 的全部粉丝。
- 不以对象池、零拷贝、自定义二进制协议等微优化代替 Profile 证据。
- 不重建、清空或替换现有 MySQL、Redis、Kafka、etcd 数据卷。
- 不把本地单机结果直接外推为生产集群容量。

## 4. 一致性与质量属性

### 4.1 分级最终一致性

| 变化 | 正常目标 | 缓存行为 |
|---|---:|---|
| 新帖子发布 | 5 秒内可见 | 不对大 V 粉丝逐个失效；依靠页面 Fresh TTL 和正常刷新 |
| 关注 | 约 1 秒开始反映，5 秒 TTL 兜底 | 用户 relation epoch 变化，旧路由和页面失效 |
| 取关 | 目标 1 秒内消失，5 秒 TTL 兜底 | 用户 relation epoch 变化，禁止返回 epoch 不匹配的 stale 页面 |
| 删除 | 目标 1 秒内消失 | 全局 content-safety epoch 变化，旧页面不可返回 |
| 转私密或审核下架 | 目标 1 秒内消失 | 与删除相同 |
| 修改置顶、标题或其他会影响页面的元数据 | 目标 1 秒内反映 | content-safety epoch 变化 |

正常目标不等于分布式强一致保证。事件系统故障时，关系变化依靠最多 5 秒的路由和页面 TTL 收敛。删除、转私密和下架的 epoch 更新失败时，KnowPost 必须临时旁路页面缓存并执行冷路径，不能继续返回未经校验的旧页面。

### 4.2 优先级

质量属性按以下顺序取舍：

1. Feed 结果正确性和敏感内容撤销。
2. 稳态可用性与可降级性。
3. 可验证的吞吐和尾延迟。
4. 可观测性与可回滚性。
5. 实现复杂度和内存成本。
6. 新内容的绝对实时性。

## 5. 方案比较

| 方案 | 优点 | 代价 | 决定 |
|---|---|---|---|
| 只做 Pipeline、归并、对象池等局部优化 | 代码改动小、风险低 | 冷路径每次仍需计算，预计难以达到 10K | 只保留 Profile 证实有效的部分 |
| 路由快照 + 冷路径优化 + 第一页两级缓存 + 按需物化 | 分阶段收益、可验证、可回滚，兼顾冷热场景 | 需要版本失效、缓存保护和更完整指标 | 采用 |
| 直接全面 Feed Head 物化 | 热读上限最高 | 写放大、失效、分层和运维复杂度最高 | 当前不采用，仅作 Phase 4 条件方案 |

## 6. 目标架构

```text
                       +-----------------------+
                       | Relation outbox/event |
                       +-----------+-----------+
                                   |
                                   v
                         +-------------------+
                         | Relation Epoch    |
                         | Invalidator       |
                         +---------+---------+
                                   | INCR per reader
                                   v
+---------+  GetUserFeed  +--------+---------+       +------------------+
| Gateway | ------------> | KnowPost RPC     | ----> | Redis epoch keys |
+---------+                +--------+---------+       +------------------+
                                   |
                      +------------+-------------+
                      |                          |
                      v                          v
             +----------------+        +----------------+
             | L1 Page Cache  | miss   | L1 Epoch Cache |
             +-------+--------+        +----------------+
                     | miss
                     v
             +----------------+
             | L2 Page Cache  |
             +-------+--------+
                     | miss
                     v
             +----------------------+       +----------------+
             | FeedRouteSnapshot    | ----> | Relation/Counter|
             +----------+-----------+       +----------------+
                        |
                        v
             +----------------------+
             | Redis Pipeline       |
             | Inbox + BigV Outbox  |
             +----------+-----------+
                        |
                        v
             +----------------------+
             | Merge/Dedup/Hydrate  |
             +----------+-----------+
                        |
                        v
                  Fill L2 and L1
```

### 6.1 组件边界

#### FeedRouteSnapshotCache

职责：保存一个读者在短时间内计算 Feed 所需的完整路由输入。

建议值结构：

```go
type FeedRouteSnapshot struct {
    UserID        int64
    RelationEpoch uint64
    Followings    []int64
    BigVAuthors   []int64
    LoadedAt      time.Time
}
```

要求：

- 使用独立 Ristretto 实例，不能再与当前 FeedMine 的 50MB L1 竞争。
- 热命中不得先调用 Relation。
- 冷 miss 才调用 Relation 和 Counter。
- 相同用户和 epoch 的冷 miss 使用 Singleflight。
- 关系事件只递增 epoch 并令快照失效，不增量修改快照；因此事件重复和乱序不会写出错误关系集合。
- 大 V 阈值变化没有关系事件，依靠不超过 5 秒的 Snapshot TTL 重新分类。
- Redis/事件异常时允许直接回源 Relation，不得因路由缓存不可用而使 Feed 不可用。

#### FeedEpochStore

职责：提供低成本的缓存失效代数，不保存 Feed 内容。

第一版包含：

```text
feed:relation:epoch:{userID}  // 关注或取关后递增
feed:content:safety:epoch     // 删除、转私密、下架、页面元数据变化后递增
```

要求：

- relation epoch 由独立失效消费者处理 Relation Outbox 派生事件并写入 Redis。
- 相同 Kafka consumer group 只负责更新 Redis 权威 epoch，不承担向每个 KnowPost L1 广播的职责。
- 每个 KnowPost 实例对 epoch 使用约 1 秒的微型 L1；过期后重新读取 Redis。
- 重复或乱序关系事件只会多递增 epoch，不影响正确性。
- content-safety epoch 在 KnowPost 写事务成功后更新；失败时当前实例进入 Page Cache bypass，直到 epoch 存储恢复。
- epoch key 不设置短 TTL，避免版本回退；测试数据使用独立前缀并支持定向清理。

#### FeedPageCache

职责：缓存已经完成路由、归并、详情装载的个性化 `FeedPage`。

第一版只接纳：

```text
page == 1
size == 20
strategy == hybrid
```

其他页码、尺寸和策略旁路页面缓存，继续使用原冷路径。这一限制减少 key 数、失效范围和错误面，后续只能根据命中率数据扩大。

建议逻辑 key：

```text
feed:page:v1:{strategyVersion}:{userID}:r{relationEpoch}:s{safetyEpoch}:p1:n20
```

V1 不加入“每次发布都变化的 user feed generation”。普通作者可以借现有 fanout 顺便更新粉丝版本，但大 V 发布若更新所有粉丝版本会重新制造写放大；新内容可见性统一由短 TTL 控制。Feed Head 物化阶段可以为已分层的热点用户引入独立 generation。

缓存形式：

- L1：独立 Ristretto，保存构建完成后不再修改的 `*FeedPage`。
- L1 cost：至少使用 `proto.Size(page)`，并在容量设计中加入对象和 key 开销系数。
- L2：Redis 保存 Protobuf bytes，不使用 JSON。
- 返回后的调用链不得修改缓存对象；测试必须验证不可变约束。若现有调用方会修改，则返回 `proto.Clone`，并用基准测试确认成本。

初始时间参数：

| 参数 | 初始值 | 说明 |
|---|---:|---|
| L1 Fresh TTL | 500ms 至 1s | 降低进程内不一致窗口 |
| L2 Fresh TTL | 3s 至 5s | 保证正常新鲜度目标 |
| TTL jitter | ±20% | 防止同批 key 同时失效 |
| Stale window | 最多 10s | 仅依赖异常时降级；epoch 不匹配禁止返回 |
| Refresh worker | 32 | 初始有界并发，压测后调整 |
| Refresh queue | 1,024 | 队列满时不得创建无界 goroutine |

#### PageRefreshCoordinator

职责：防止击穿和刷新风暴。

- 同一完整 page key 的进程内并发 miss 使用 Singleflight。
- Singleflight 只合并相同用户页面，不宣称能够合并大 V 的所有粉丝。
- Fresh 即将过期时允许异步刷新。
- Stale 且 epoch 匹配时可返回旧页面并投递刷新。
- epoch 不匹配、超过 Stale window 或敏感旁路开启时必须同步走冷路径。
- 刷新任务使用独立超时 Context，不继承已经取消的请求 Context。
- 队列已满时记录指标，并根据页面状态选择同步冷路径或返回仍合法的 stale；禁止无限启动 goroutine。
- 页面已过期且同 key 刷新仍在队列时，前台在 key 协调器内原子接管；worker 后续必须跳过已失效任务。
- 页面已过期且同 key 刷新已经运行时，前台等待其有界结果；刷新失败后回退原 Hybrid 冷路径，刷新器不得成为可用性单点。
- 多实例扩展前再评估 Redis 租约，单实例阶段不提前引入分布式锁。

## 7. 关键数据流

### 7.1 热点第一页命中

1. 根据用户、策略、页码和尺寸判断是否允许页面缓存。
2. 从微型 L1 获取 relation epoch 和 content-safety epoch；过期时从 Redis 刷新。
3. 构建完整 page key。
4. L1 Fresh 命中，直接返回不可变 `FeedPage`。
5. 全程不调用 Relation、Counter、Inbox、BigV Outbox 或 FeedItem DB。

### 7.2 页面缓存 miss

1. L1 miss 后，以完整 page key 进入进程内协调器；同 key 冷请求共享一次 L2 查询和一次 loader，不同 key 独立执行。
2. 协调器 leader 读取 Redis L2；命中且记录仍在 Fresh/Stale 窗口内时，反序列化、回填 L1 并返回。
3. L2 miss、损坏或过期后读取 FeedRouteSnapshot；若同 key 已有刷新，则按其 queued/running 状态接管或等待。
4. 刷新失败、超时或不存在可复用结果时，执行受独立 2 秒 Context 约束的 Hybrid 冷 loader。
5. Snapshot miss 才调用 Relation 和 Counter。
6. 使用一次受控 Pipeline 读取 Inbox 和大 V Outbox。
7. 完成 Top-N、去重和 FeedItem 批量装载。
8. 先写 L2，再写 L1；写缓存失败不影响本次正确结果返回。

### 7.3 关注或取关

1. Relation 按现有事务写 following 和 outbox。
2. Relation Syncer 继续维护既有 ZSet 等派生数据。
3. Feed relation invalidator 以独立 group 消费对应领域事件。
4. 对 `feed:relation:epoch:{fromUserID}` 执行原子 `INCR`。
5. KnowPost 最迟在 epoch L1 过期后使用新 epoch；旧 RouteSnapshot 和 FeedPage 不再命中。
6. 若事件处理延迟，路由和页面 TTL 在 5 秒内触发回源兜底。

### 7.4 发布新帖

1. 普通作者继续走现有 Kafka fanout，写粉丝 Inbox。
2. 大 V 继续只写 Outbox。
3. 不遍历大 V 粉丝，不对所有页面执行通配符删除，不更新所有粉丝 generation。
4. 旧第一页在 Fresh TTL 内允许短暂不可见新内容，正常不超过约 5 秒。

### 7.5 删除、转私密或下架

1. Delete、UpdateVisibility、UpdateTop 和 PatchMetadata 在数据库事务提交后执行既有缓存失效，再同步原子递增全局 `feed:content:safety:epoch`；Publish 不递增该 epoch。
2. 同一 KnowPost 事务的 `canal-outbox` 事件由独立 consumer group `knowpost-feed-content-safety-epoch` 消费；Updated/Deleted 再次递增 epoch，Published 跳过。同步与异步双 bump 是有意的单调失效，异步链路负责跨实例和进程重启后的持久补偿。
3. PageCache 开启时，个人 FeedItem 使用 `feed:item:personal:s{safetyEpoch}:{postID}`。旧无版本 key 和旧 safety 版本 key 均不可用于重建新页面。
4. 页面返回前再次校验 pending、relation epoch 和 safety epoch；lookup 期间任一状态变化或读取失败都丢弃结果。
5. 同步 bump 失败时不篡改已提交的业务结果，而是设置进程内 `SafetyPending`；后续读取先尝试补偿。补偿成功前，或 epoch 无法可信读取时，同时 bypass PageCache 和个人 FeedItem cache，直接由数据库冷路径过滤可见性。
6. 非当前实例的 pending 由持久 outbox consumer 最终推进 Redis epoch；PageCache 非 `off` 时配置强制要求 `SafetyConsumerEnabled=true`，避免只依赖进程内标记。

全局 safety epoch 会使全部第一页缓存失效；同步写路径与 outbox 补偿通常使一次敏感变化推进两次 epoch，但不会改变正确性。删除、转私密、下架和置顶变化相对发布低频，因此先用这种额外冷却换取可证明的撤销正确性；只有监控证明敏感写频率导致持续冷流量时才拆成作者或帖子级版本。

## 8. 分阶段实施与门禁

### Phase 0：可观测性和长窗口基线

实施：

- 为 KnowPost 和 Relation 增加仅开发环境启用、只绑定本机的独立 pprof HTTP 端口。
- 不在 gRPC 9004/9006 或 Prometheus 端口上假设 `/debug/pprof` 已存在。
- 增加 Relation、Counter、RouteSnapshot、Inbox、BigV Outbox、Merge、Dedup、FeedItem hydrate、序列化分阶段耗时。
- 增加依赖调用次数、缓存 outcome、Singleflight、刷新队列、stale 返回和 epoch 延迟指标。
- 重跑 c32/c64/c128/c256：预热 10 秒、正式采样 60 秒、重复三次。

门禁：

- Feed 正确性零失败。
- 仪表代码引入的 QPS 或 P95 回退不超过 5%。
- 形成 KnowPost 和 Relation CPU、heap、alloc、block、mutex Profile。
- 明确新的稳定吞吐、拐点和三轮波动范围。

### Phase 1：完整路由快照

实施：

- 建立独立 RouteSnapshot L1 和配置。
- 加入 relation epoch store、事件 invalidator 和 5 秒 TTL 兜底。
- 热命中跳过 Relation 和 Counter。
- 保留冷 miss Singleflight、独立 2 秒 Context 和错误不缓存语义。

门禁：

- 热路由场景 `Relation RPC / Feed request` 至少下降 80%。
- 相同拓扑和负载下，QPS 至少提升 20%，或单位成功请求 CPU 至少下降 20%。
- Follow/Unfollow、事件重复、乱序、延迟和丢失恢复测试通过。
- 无 Feed 重复、越权作者或永久漏内容。

任一门禁失败则停止进入下一阶段，保留指标并回滚开关分析原因。

Phase 1 实测检查点（2026-08-15）：

- RouteSnapshot、RelationEpoch 和观测开关在真实 Docker 运行态显式启用，Follow/Unfollow 集成与 Hybrid 冒烟通过。
- c32/c64/c128/c256 各 3 次、10 秒预热与 60 秒正式采样全部完成，失败/timeout=0，degraded=false。
- Relation RPC/request 从约 1.0 降到 `0.0010~0.0011`，下降约 99.9%。
- 四档 QPS 中位数较 Phase 0 严格基线提升 `83.66% / 70.98% / 50.43% / 39.04%`。
- 峰值无错吞吐 `3978.74 QPS`（c64，P95 `22.054 ms`）；稳定并发 c64，拐点 c128。
- Phase 1 门禁通过。详细证据见 `docs/superpowers/reports/2026-08-15-feed-hybrid-wp5-route-snapshot.md`。

### Phase 2：Hybrid 冷路径 Pipeline

实施：

- 将 Inbox 与 BigV Outbox 放入同一受控 Pipeline。
- 保留每批最大 128、部分错误隔离和不支持批量接口时的兼容路径。
- 只根据 Phase 0 Profile 决定是否优化归并、切片分配或对象复用。
- 不在没有证据时引入动态并发控制、sync.Pool 或自定义协议。

门禁：

- 冷路径 P95 至少下降 15%，或 Redis dependency round trips/request 至少下降 20%。
- 冷路径 QPS、内存和 GC 不得恶化超过 10%。
- Pipeline 部分失败、超时和 Context 取消测试通过。

这里的 round trip 按一次单命令调用或一次 `Pipeline.Exec` 计数；合并 Pipeline 不会减少
`ZREVRANGE` 命令条数，禁止把网络往返下降误写为 Redis command 数量下降。报告同时保留
Redis `ops/sec`，用于证明命令负载没有被指标口径隐藏。

Phase 2 实测检查点（2026-08-15）：

- CombinedPipeline 在真实 Docker 运行态显式启用，Hybrid 冒烟、RelationEpoch 集成和部分错误隔离测试通过。
- c32/c64/c128/c256 各 3 次、10 秒预热与 60 秒正式采样全部完成，失败/timeout=0，degraded=false。
- Redis dependency round trips/request 从约 3.00 降到约 2.00，下降约 33%；Redis 命令数没有被误报为下降。
- 四档 QPS 中位数较 Phase 1 分别提升 `123.82% / 147.99% / 173.61% / 172.58%`，P95 下降 `54.46% / 58.42% / 63.26% / 62.82%`。
- c128 三轮中位数 `10661.31 QPS`、P95 `18.112 ms`，在当前单机 RPC、约 1,200 分散读者口径下重复达到 10K。
- 稳定并发仍建议 c64（`9866.82 QPS`、P95 `9.171 ms`）；c128 为吞吐拐点，c256 已出现吞吐回落和排队。
- Redis 峰值约 80K~81K ops/s；下一瓶颈优先检查 Redis 命令处理、FeedItem MGET 和 KnowPost CPU。
- Phase 2 门禁通过。详细证据见 `docs/superpowers/reports/2026-08-15-feed-hybrid-wp6-combined-pipeline.md`。

### Phase 3：第一页最终结果缓存

定位修正：Phase 2 已在不依赖整页缓存的冷路径上达到当前单机 RPC 10K 目标，因此 Phase 3 不再是达标必需项。它作为热点第一页的可选演进继续保留，目标是降低单位请求成本和尾延迟；不得用缓存友好数据集覆盖 WP6 冷路径基线，也不得因追求更高峰值放宽一致性与高基数回退门禁。

WP7 实测检查点（2026-08-15）：

- Fresh L1/L2、完整 epoch key、Protobuf L2、故障回退、loader Context 与默认关闭开关已实现。
- PageCache 返回前二次验证 SafetyPending、relation epoch 和 safety epoch；lookup 中途发生变化时丢弃旧页并回 WP6。
- L2 get/decode/encode/set 独立观测并按 operation 限频错误日志，不把 Redis/codec 故障隐藏为普通 miss。
- 普通、integration、50 次重复测试和复审通过，最终 Ready（Critical 0 / Important 0）。
- Docker `off -> l2 -> l1-l2` 启动/冒烟通过；40 个 page1/size20 真实请求零失败并观察到 Fresh L1 命中。
- 容器已恢复 `PageCache=off`。在 WP8 content-safety 写路径完成前不得默认启用。
- 详细证据见 `docs/superpowers/reports/2026-08-15-feed-hybrid-wp7-page-cache-foundation.md`。

WP8 实测检查点（2026-08-15）：

- 四条敏感写路径在 commit 后同步 bump；commit 后既有缓存删除失败仍执行 bump 并保留原错误，bump 失败则置 pending 且不改变已提交业务结果。
- PageCache 模式下个人 FeedItem 按 safety epoch 隔离；pending、epoch 错误、准入失败以及返回前 epoch 变化均禁止复用个人 FeedItem cache。
- 独立 Kafka consumer 对 Updated/Deleted 提供持久补偿，Published 明确跳过；PageCache 开启时强制要求 consumer 同时开启。
- 真实开发栈中转私密后约 `95.42ms`、删除后约 `92.51ms` 不可见，且两次均观察到 safety epoch `N -> N+2`；consumer group lag 为 0。
- 自动化回归、相关 `go vet`、Docker 健康检查与二次复审通过，最终 Ready（Critical 0 / Important 0）。
- 详细证据见 `docs/superpowers/reports/2026-08-15-feed-hybrid-wp8-content-safety-epoch.md`。

实施：

- 建立独立 Page L1、Redis L2、epoch 校验、Singleflight、TTL jitter、SWR 和有界刷新。
- 只接纳 Hybrid `page=1,size=20`。
- 加入删除、可见性、置顶和元数据变化后的 content-safety epoch。
- 所有功能可通过配置关闭并即时回到冷路径。

门禁分为三类：

#### 热点目标

- 20 个热点读者、第一页、size=20、缓存预热。
- RPC c64/c128/c256 阶梯中至少一档稳定达到 10,000 成功 QPS。
- 每档预热 10 秒、采样 60 秒、重复三次；以三次中位数为结论。
- P95 < 50ms，P99 < 100ms，错误率 < 0.1%，正确性失败为 0。
- Page L1+L2 Fresh 命中率不低于 97%。
- 无容器重启、Redis eviction/rejected connection 或持续资源积压。

#### 分散目标

- 当前约 1,200 读者数据集下，相对 Phase 0 同场景 QPS 至少提升 30%。
- 20,000 高基数读者场景不要求 10K，但冷/低命中吞吐和 P95 不得回退超过 10%。
- 报告必须同时展示热点、分散和高基数结果，不能只展示最好结果。

#### 一致性目标

- 新发布内容正常在 5 秒内进入 Feed。
- 取关、删除、转私密和下架按第 4 节目标收敛。
- 缓存集中失效时，无无界 goroutine、刷新队列无限增长或依赖雪崩。
- Redis、事件消费者和 Relation 故障恢复测试通过。

如果热点命中率不足 97%，首先分析读者基数、TTL、key 维度和内存淘汰，不直接扩大缓存或延长敏感数据 TTL。

### Phase 4：条件式 Feed Head 物化

进入条件：Phase 3 未达到热点目标，或真实/模拟访问分布显示页面缓存命中率不足且冷计算仍是主要瓶颈。

实施边界：

- 热点用户：异步物化 Feed Head。
- 活跃用户：页面缓存加增量归并。
- 冷用户：保留现有 Hybrid 冷计算。
- 普通作者继续 fanout，大 V 只更新 Outbox。
- 只为热点或活跃读者重建 Head，不为大 V 全部粉丝同步构建。
- 用户分层阈值由访问频率、命中率、物化成本和内存实测确定，不预设固定百分比。
- Cursor 分页不再绑定 Feed Head：其独立增量设计已在第 19 节获批，必须独立实现、独立开关、独立 A/B，不与 Feed Head 同批上线。

当前 Spec 仍不授权直接实现全面 Feed Head 物化；第 19 节只授权 Cursor 深分页及其验证工具，不扩大物化范围。

## 9. 压测方法

### 9.1 固定变量

每组 A/B 对照必须固定：

- Git commit 和未提交补丁摘要。
- Feed 策略、入口、读者数据集和随机种子。
- 容器/本地进程拓扑。
- CPU、内存、GOMAXPROCS 和 Docker 资源限制。
- Redis、MySQL、Kafka、etcd 数据状态。
- 并发、预热时间、采样时间、客户端主机和超时。
- 日志级别和 go-zero Mode。

任何拓扑、资源或 Mode 变化都必须重建基线，不能与原结果直接计算提升百分比。

### 9.2 测试矩阵

| 维度 | 取值 |
|---|---|
| 入口 | RPC、Gateway |
| 策略 | 以 Hybrid 为主；Push/Pull 做回归对照 |
| 读者基数 | 20、约 1,200、20,000 |
| 缓存状态 | 冷启动、L2 热、L1 热、集中失效 |
| 并发 | 32、64、128、256；必要时继续倍增 |
| 负载 | 纯读取、90/10 读写混合、突发、恢复观察 |
| 分页 | page1/size20 主目标，深分页旁路回归 |
| 故障 | Redis 延迟/不可用、Relation 延迟、消费者暂停与恢复 |
| 正确性 | 发布、关注、取关、删除、转私密、下架、重复内容 |

### 9.3 统计规则

- 每个正式阶段至少 60 秒并重复三次。
- 报告三次中位数、最小值、最大值和变异范围。
- 吞吐使用成功 QPS，延迟只统计成功请求，同时单列失败和超时。
- 记录 P50/P90/P95/P99/Max，不用估算分位数。
- 记录缓存命中率时区分 L1 Fresh、L2 Fresh、Stale、Miss、Bypass。
- 记录每成功请求的 Relation RPC、Counter RPC、Redis commands 和冷计算次数。
- 记录 Docker/进程 CPU、RSS、GC、goroutine、Redis ops、Kafka lag 和 MySQL 指标。
- 压测客户端自身 CPU 饱和时，该轮报告标记为无效或改为多客户端，不把客户端上限当服务端上限。

### 9.4 可变数据场景隔离

- 纯读容量矩阵与 publish/mixed 矩阵使用不同报告命名空间；纯读矩阵必须先执行，且期间禁止 seed/publish 改变 Feed 初态。
- 一个 manifest 只绑定一种 Feed strategy；setup 记录 strategy 与 seed，run 不匹配时 fail closed。Push/Pull/Hybrid 各建独立数据集。
- publish、90/10 与 80/20 会修改 MySQL、Outbox、Kafka 和 Redis Feed ZSET。每个 concurrency/trial 必须从同一个可审计 checkpoint 恢复，并在恢复后确认 Kafka lag=0、Feed key 指纹与 MySQL 基线一致，才能进入正式 60 秒采样。
- trial 级 checkpoint/restore 已于 2026-08-18 实现；旧 `noncomparable-mutation-probe` 入口已移除。任何 mutation run 缺少 checkpoint 或精确 RunID 确认都会 fail closed；只有 60 秒 × 三轮正式报告可形成容量曲线与提升百分比。

## 10. 可观测性

指标名以项目现有规范为准，但必须能表达以下语义，且禁止把 user ID 放入指标 label：

```text
feed_request_duration{stage,outcome}
feed_dependency_calls{dependency,outcome}
feed_route_cache_requests{level,outcome}
feed_page_cache_requests{level,outcome}
feed_page_refresh_total{outcome}
feed_page_refresh_queue_depth
feed_singleflight_waiters{kind}
feed_epoch_read_total{kind,outcome}
feed_epoch_lag_seconds{kind}
feed_stale_served_total{reason}
feed_cache_bypass_total{reason}
feed_cold_compute_total{outcome}
```

日志要求：

- 热命中不打印请求级成功日志。
- epoch 更新失败、缓存解码失败、刷新队列满、stale 降级和 bypass 状态变化必须限频记录。
- 错误日志携带 trace ID、阶段和错误分类，不打印完整 Feed 内容。

pprof 要求：

- 仅开发/压测配置启用。
- 绑定 `127.0.0.1` 或 Compose 仅映射到宿主机回环地址。
- 使用独立 HTTP 端口，不与 gRPC 或 Prometheus 端口混用。
- debug HTTP 显式使用私有 mux；禁止任何实际 listener 使用 `http.DefaultServeMux` 或 `Handler=nil`，避免标准库 pprof 的包初始化注册被其他端口意外暴露。
- 容器内使用 `6064/6066`；当前 Windows/Hyper-V 环境排除 TCP `6034-6133`，Compose 宿主机回环端口固定为 `16064/16066`。
- 正式报告同时采集 KnowPost 与 Relation，避免只优化表面服务。

## 11. 配置、发布与回滚

建议配置开关：

```yaml
Feed:
  Strategy: hybrid
  Epoch:
    KeyPrefix: feed
    RelationL1TTL: 1s
    SafetyL1TTL: 1s
    RelationConsumerEnabled: false
    SafetyConsumerEnabled: false # PageCache 非 off 时必须为 true
  RouteSnapshot:
    Enabled: false
    TTL: 5s
    EpochL1TTL: 1s
  CombinedPipeline:
    Enabled: false
    BatchSize: 128
  PageCache:
    Mode: off # off | l2 | l1-l2
    Page: 1
    Size: 20
    L1FreshTTL: 800ms # ±20% 后仍在 500ms~1s 硬边界内
    L2FreshTTL: 4s    # ±20% 后仍在 3s~5s 硬边界内
    StaleTTL: 10s
    JitterPercent: 20
    RefreshWorkers: 32
    RefreshQueue: 1024
    LoaderTimeout: 2s
  FeedHead:
    Enabled: false
```

具体字段需遵循当前 go-zero 配置解析方式；时长字段若不能直接解析 `time.Duration`，实施计划中统一改为毫秒字段，禁止在不同配置文件里混用单位。

发布顺序：

1. 先上线指标和 pprof，所有性能开关关闭。
2. 单独启用 RouteSnapshot，完成 Phase 1 验证。
3. 单独启用 CombinedPipeline，完成 Phase 2 验证。
4. 先启用 Content Safety consumer 并验证 consumer group 健康，再按 L2、L1、SWR 的顺序启用 PageCache。
5. 每步出现正确性异常、错误率上升或依赖雪崩，立即关闭当前开关回退冷路径。

页面缓存不是可用性的单点：L1、L2、epoch 或刷新器任一失败时，默认回退原 Hybrid 冷路径。只有敏感 epoch 无法确认时，页面缓存 fail closed，冷路径继续服务。

## 12. Docker 与本地混合运行约束

### 12.1 问题定性

开发过程中如果 Docker 镜像无法正确构建、拉取缓慢或拉取失败，按 Docker registry/mirror 镜像源问题处理，不把它归因于项目网络连通性，也不重复进行无关的通用网络故障排查。

允许记录并诊断：

- 具体失败的 registry、repository、tag 和 digest。
- Docker daemon 当前 registry mirror 配置。
- 本地是否已经存在可复用镜像和构建缓存。
- 镜像 manifest、平台架构或上游 tag 是否变化。

不允许默认采取：

- 反复执行全量 `docker compose pull`。
- 因镜像源问题清空 BuildKit 缓存、镜像、容器或数据卷。
- 重建 MySQL、Redis、Kafka、etcd 等已可用依赖。
- 把镜像拉取错误写成“网络不通”结论。

### 12.2 首选拓扑：Compose 全栈

在现有应用镜像可以利用本地缓存正常构建时：

- 继续使用 `deploy/compose/docker-compose.dev.yml`。
- 只重建发生代码变化的项目服务。
- 中间件使用已经构建/拉取的本地镜像和现有数据卷。
- 不执行无必要的 `pull`、全栈 rebuild 或依赖重建。

### 12.3 降级拓扑：Compose 中间件 + 本地项目服务

出现 registry/mirror 问题且短时间无法绕过时：

```text
本地进程：Gateway、KnowPost、Relation、Counter，以及本轮涉及的其他项目服务
Compose：Kafka、ZooKeeper、etcd、Elasticsearch 等已经可运行的中间件
宿主机既有服务：MySQL、Redis（按当前开发栈配置）
```

要求：

- Compose 中间件使用 `--no-build` 思路启动已存在服务，不触发应用镜像构建。
- 项目服务使用本地配置运行，依赖地址指向宿主机已暴露端口。
- etcd 注册地址必须是其他本地进程能够访问的宿主机地址，不能注册容器内部地址。
- Kafka bootstrap 地址使用宿主机映射地址；MySQL、Redis 使用当前本地开发地址。
- JWT 证书、对象存储、环境变量和 Feed Strategy 与 Compose 基线保持一致。
- 本地运行日志和 PID 必须单独记录，压测结束后只停止本轮启动的进程。
- pprof 端口只绑定本机，避免与已有服务冲突。
- 在 Windows 上发布新端口前执行 `netsh interface ipv4 show excludedportrange protocol=tcp`；若端口位于排除范围，修改宿主机映射而不改容器内监听端口，并把映射写入测试报告。

### 12.4 压测可比性

Compose 全栈与“Compose 中间件 + 本地项目服务”是两种不同测试拓扑：

- 不直接比较两种拓扑的 QPS 提升比例。
- 切换拓扑后先执行冒烟，再重跑 Phase 0 基线。
- 每份报告必须记录 topology、进程启动方式、服务地址、资源限制和 Git 状态。
- 同一优化前后对照必须使用同一拓扑。

因此 Docker 镜像源问题不会阻塞代码开发和功能验证，但会触发拓扑重基线，保证性能结论仍然可信。

## 13. 测试策略

### 13.1 单元测试

- RouteSnapshot 命中不调用 Relation/Counter。
- RouteSnapshot TTL、epoch 变化、重复事件和乱序事件。
- Singleflight 独立 Context 和每个 waiter 取消。
- Pipeline Inbox + BigV 结果顺序、部分错误和批次边界。
- Page key 的用户、epoch、策略、页码和尺寸隔离。
- L1/L2 Fresh、Stale、Miss、Bypass 状态机。
- epoch 不匹配时绝不返回 stale。
- 刷新队列满时无无界 goroutine。
- 缓存对象不可变约束或 Clone 路径。
- content-safety epoch 更新失败后 Page Cache bypass。

### 13.2 集成测试

- Redis epoch 原子递增和进程重启后不回退。
- Relation Outbox 事件到 Redis epoch 的链路。
- Follow/Unfollow 后 RouteSnapshot 和 Page Cache 收敛。
- 发布后普通作者 Inbox、大 V Outbox 和页面 TTL 新鲜度。
- 删除、转私密、下架后旧 L1/L2/Stale 页面不可见。
- Redis 故障时回退冷路径，恢复后重新进入缓存路径。
- 多个并发相同 page key 只产生一次冷计算。
- 多个不同用户 page key 不错误宣称被 Singleflight 合并。

### 13.3 回归与压测

- `go test` 覆盖 KnowPost、Relation、Gateway、cachex 和 loadtest。
- Race 测试覆盖新增缓存状态、epoch L1 和刷新协调器。
- 组件 Benchmark 只解释局部成本，不作为端到端容量结论。
- 正式容量结论必须来自第 9 节矩阵和原始 JSON/CSV 报告。

## 14. 风险与反挑战

| 风险 | 后果 | 缓解 |
|---|---|---|
| 10K 只在少量热点用户成立 | 被误读为通用容量 | 强制同时报告 20、1,200、20,000 用户结果 |
| Relation 事件延迟或丢失 | 路由快照短暂陈旧 | Redis epoch、5 秒 TTL、冷路径回源、lag 指标 |
| safety epoch 全局变化 | 所有页面瞬时 miss | 变化低频、TTL jitter、SWR 不跨 epoch、刷新有界 |
| 100k 粉丝页面同时过期 | 不同 key 无法被 Singleflight 合并 | jitter、SWR、刷新池、Outbox 热数据缓存、过载保护 |
| Ristretto 容量估计偏低 | 淘汰率高、命中不足 | `proto.Size` 加开销系数、独立缓存、实测 RSS/GC |
| 缓存 `*FeedPage` 被修改 | 跨请求数据污染或竞态 | 不可变约束、Race 测试、必要时 Clone |
| Redis L2 热点 | 网络或单 key 压力 | L1、Pipeline、指标；实测后再决定分片 |
| 页面缓存掩盖冷路径退化 | 缓存失效时雪崩 | 高基数冷测试和集中失效测试作为硬门禁 |
| 开发 Mode 无负载保护 | 容量测试和生产行为不同 | 容量测试保留 dev 口径；另做 pro 保护性对照，不混合结论 |
| Docker 镜像源异常 | 构建受阻 | 使用已构建中间件、本地运行项目服务、切换后重基线 |

如果用户增长 100 倍，首先失效的不是归并算法，而是：

1. 高基数读者导致 L1 页面命中率下降，冷计算比例上升。
2. Relation epoch 与页面 L2 Redis 流量增加。
3. KnowPost 冷路径和 Relation 再次饱和。

对应演进顺序是热点 Feed Head、Redis 水平扩展、KnowPost 横向扩容，而不是继续无限增加单机协程或 Pipeline 批次。

## 15. ADR

### ADR-001：采用渐进式分层优化，不直接全面物化

- 背景：当前已有可工作的 Hybrid 冷路径和 2.396K QPS 基线。
- 决定：依次实施可观测性、路由快照、Pipeline、第一页缓存。
- 理由：每阶段可独立验证和回滚，能区分真实收益来源。
- 代价：到达最终目标的步骤更多，短期保留冷热两条路径。

### ADR-002：Feed 新内容采用短 TTL 最终一致性

- 背景：大 V 发帖不能遍历全部粉丝失效页面。
- 决定：正常允许最多约 5 秒不可见，不给全部粉丝更新 generation。
- 理由：保留 Hybrid 避免写放大的核心价值。
- 代价：新内容不是读己之写级别的立即可见。

### ADR-003：关系事件只失效，不增量修改快照

- 背景：现有 RelationEvent 没有严格的每用户事务版本，事件可能重复、延迟或乱序。
- 决定：事件只递增 Redis relation epoch，快照 miss 后完整回源。
- 理由：重复和乱序只造成额外 miss，不会构造错误关注集合。
- 代价：关系变化后的第一次读取需要冷回源。

### ADR-004：敏感内容使用全局 content-safety epoch

- 背景：完整页面缓存可能继续包含已删除或已转私密的帖子。
- 决定：敏感事务提交后同步递增全局 epoch，同时由独立 `canal-outbox` consumer 持久补偿；个人 FeedItem key 携带 safety epoch。pending 或 epoch 不可信时 fail closed，页面与个人 FeedItem cache 均 bypass。
- 理由：同步路径提供低延迟，outbox 避免进程崩溃或跨实例导致补偿永久丢失，版本化 FeedItem 阻止旧详情重建新页面。
- 代价：正常敏感变化通常双 bump 并造成全局第一页缓存冷却；需监控敏感 mutation 频率和额外冷流量。

### ADR-005：第一页缓存独立 L1，L2 使用 Protobuf

- 背景：当前 FeedMine L1 只有 50MB 且已被 RouteCache 复用，缓存对象的 Go 内存成本高于序列化大小。
- 决定：RouteSnapshot 和 FeedPage 使用独立预算；L2 存 Protobuf bytes。
- 理由：避免互相淘汰，降低 L2 存储与解码成本。
- 代价：增加配置和内存监控要求。

### ADR-006：Docker 镜像问题不阻塞项目服务本地运行

- 背景：Docker registry/mirror 可能造成镜像拉取或构建缓慢，但项目网络连通性正常。
- 决定：保留已构建中间件容器，项目服务可切换为本地进程。
- 理由：继续开发和验证，不对无关网络问题做错误归因。
- 代价：切换拓扑后必须重新建立容量基线。

## 16. Definition of Done

本设计完成的判定条件：

- Phase 0 至 Phase 3 的代码、配置、测试、指标和 Feature Flag 已实现。
- 每阶段均有优化前后同拓扑报告、原始 JSON/CSV 和环境快照。
- Follow/Unfollow、发布、删除、转私密和下架正确性测试全部通过。
- 缓存和事件故障能够自动回退，未出现无界 goroutine 或不可恢复缓存状态。
- 热点第一页达到 10K 目标，或者报告以证据说明未达到的第一瓶颈和下一步选择。
- 分散和高基数结果被完整披露，不把热点结果外推为通用容量。
- Docker 全栈或本地混合拓扑均有可执行说明；同一对照不混用拓扑。
- Feed Head 物化未经独立设计批准不会开始实现；Cursor 只按第 19 节已批准边界实施。

## 17. 实施计划边界

下一步实施计划应按以下工作包拆分，每个工作包独立测试、独立审查：

1. P0：pprof、阶段指标、长窗口基线。
2. P1：EpochStore 与 RelationInvalidator。
3. P2：独立 RouteSnapshotCache。
4. P3：Inbox + BigV 合并 Pipeline。
5. P4：PageCache key/value、L1/L2 和 epoch 校验。
6. P5：Singleflight、SWR、有界刷新和故障降级。
7. P6：正确性矩阵、热点/分散/高基数压测与报告。
8. P7：Docker 全栈与本地混合运行手册验证。

每个工作包必须先补充失败测试或基准，再实现最小改动；不得把 Phase 4 Feed Head 物化混入上述计划。

## 18. WP11 正式容量与延后工作

状态：纯读部分已于 2026-08-17 完成约 1,200 / 20,000 读者两档的 Control/Treatment；Cursor 正式 A/B 已完成。写入/混合数据隔离已于 2026-08-18 按 18.2 实现，2026-08-27 在新鲜 checkpoint 上完成 RPC-only 正式 Control/Treatment。后续 Gateway 性能矩阵因已知入口瓶颈按用户决策取消。

### 18.1 三档完整 control/treatment 纯读矩阵

目标是给出单机读容量、拐点和优化收益，而不是继续修改 Feed 代码。

数据与命名空间：

- 为 `hot=20`、`distributed=1,200`、`high=20,000` 分别创建 Hybrid manifest；manifest 固定 strategy、seed、reader 数和初始帖子。
- 同一 cardinality 的 control/treatment 复用同一个 manifest，但使用不同 `ReportRunId`；Push/Pull 回归必须使用各自新建的 manifest，不能切换同一数据集的物化布局。
- 每张卡至少保留三个结果命名空间：`control-cold`、`treatment-cold`、`treatment-target`。前两者是严格同缓存状态 A/B；`treatment-target` 使用目标缓存状态验证最终容量，不能与 control cold 直接宣称为同口径提升。
- `treatment-target` 默认映射：hot=`l1-warm`、distributed=`l2-warm`、high=`cold`。若 warm 时间窗门禁失败，该轮停止并报告 working-set/TTL 不可满足，不降级后仍保留原标签。

固定矩阵：

| 维度 | 取值 |
|---|---|
| 入口 | RPC、Gateway（分别报告） |
| 场景 | hot-read、distributed-read、deep-page；突发/steady 作为独立补充 |
| 并发 | 32、64、128、256；若 c256 仍线性增长再单独批准倍增 |
| 采样 | 10 秒预热 + 60 秒正式采样 |
| 重复 | 每个点 3 次，报告中位数/min/max/变异范围 |
| 页面 | page1/size20；deep-page 单列 |
| 拓扑 | 一次矩阵内固定 `compose-full` 或 `compose-middleware-local-services` |

每轮开始前必须验证 Git patch hash、被测 image ID/本地 exe SHA、资源限制、GOMAXPROCS、服务地址、manifest 指纹、strategy、FeaturePreset、Kafka lag 和 Redis identity；任何不一致都 fail closed。每轮结束必须满足：报告 `complete=true`、失败/timeout=0、客户端未饱和、服务无重启、Redis eviction/rejection=0、Kafka 恢复为 lag=0、Feed 指标覆盖完整。

WP11 还应把 go-zero Mode 与日志级别作为独立环境字段记录，而不只依赖 Git/image/exe 身份间接约束；SWR queue/active/pending 当前是 1 秒周期的 sampled max，报告必须按“采样峰值”解释。Redis INFO 与 safety epoch GET 的少量监控自流量也写入测试说明。

结果必须回答：

- 每个 cardinality 的稳定并发、吞吐峰值、膝点、P95/P99 和 CPU/RSS/Redis ops。
- `control-cold -> treatment-cold` 的严格 A/B 收益及冷路径是否回退。
- treatment-target 是否在热点卡稳定达到 10K，Fresh ratio 是否达到 97%。
- 1,200 读者是否相对同拓扑 Phase 0 提升至少 30%；20,000 读者是否回退超过 10%。
- Gateway 开销与 RPC 单机上限分别是多少，禁止把两种入口混成一个数字。

原始全并发矩阵预计自动运行时间约 7–9 小时。2026-08-17 执行时经用户批准改为“10 秒 scout 找拐点，正式阶段只跑稳定最优档 60 秒 × 3”，从而避免对已确认过载的并发重复长测。该压缩不改变每个正式点的采样时长、三轮中位数和完整性门禁。

### 18.2 publish/mixed trial 级数据隔离

原 `IncludeMutationProbes` 非隔离入口已移除；所有 publish、90/10、80/20 运行现在都必须提供 checkpoint 和精确 `confirm-mutation=RunID`。实现目标不是“尽量清理”，而是能证明每个 warmup 和 measured trial 从相同的逻辑 Feed 数据状态开始。

Checkpoint 内容：

- manifest hash、strategy、seed、作者/读者集合和基线帖子 ID。
- benchmark 作者在 MySQL `know_posts` 与相关 `outbox` 的精确基线集合及内容指纹。
- manifest 范围内所有 Inbox/BigV ZSET 的 Redis `DUMP`/内容摘要、key 数、成员数和总指纹；备份 key 必须带 run ID，禁止扫描或覆盖非本 run 数据。
- Kafka `feed-fanout-group` 的分区 offset/lag 证据；Kafka 历史不回滚，但 restore 前后必须 lag=0，且不存在旧 mutation 在途。
- content-safety epoch 只允许单调递增，禁止回写旧值；restore 后 bump 新 epoch，使所有旧 L1/L2 页面不可命中。
- checkpoint ID、创建时间、工具版本和二进制身份写入每份 mutation 报告。

每个 trial 的顺序固定为：

1. 等待 Kafka lag=0，验证 checkpoint 与当前 manifest/strategy/binary 匹配。
2. 删除 benchmark 作者在基线之后产生的帖子及其 Outbox/精确 Redis 派生 key；不得按宽泛 pattern 删除。
3. 从 checkpoint 恢复 manifest 范围内的 Inbox/BigV ZSET，并校验恢复后的完整指纹。
4. bump safety epoch，等待 epoch L1 收敛，再执行对应 cache-state 准备。
5. mutation warmup 完成并等待 Kafka lag=0 后，再执行一次完整 restore；warmup 产生的数据不能进入正式测量初态。
6. 执行 60 秒 measured trial，记录生成 post IDs、Kafka 恢复时间和所有资源指标。
7. 报告落盘后再次 restore，为下一 trial 提供相同初态；restore 失败立即停止整个矩阵。

需要实现的代码与测试：

- 新增 manifest-owned checkpoint schema、create/verify/restore 命令和报告字段；所有写操作要求 RunID 精确确认。
- 单元测试覆盖 key 映射、指纹稳定性、外部 key 保护、缺失/损坏 checkpoint、重复 restore 幂等、epoch 不回退。
- Redis/MySQL 集成测试覆盖“warmup 写入 -> restore -> 指纹等于基线”“measured 写入 -> restore -> 下一轮同初态”。
- Kafka 集成测试覆盖 restore 前 lag 未清零时拒绝执行，以及 drain 后不再处理上一轮消息。
- 故障注入覆盖 MySQL 事务失败、Redis 部分恢复失败和进程中断；失败时 checkpoint 保留、报告 incomplete，不继续下一 trial。
- comparison 只有在 checkpoint ID、基线指纹、strategy、seed、工作负载时长和资源证据一致时才接受三轮及容量曲线。

2026-08-18 实现与验证结果：

- 新增版本化 checkpoint schema 及 `create-mutation-checkpoint`、`verify-mutation-checkpoint`、`restore-mutation-checkpoint`；文件独占创建，不覆盖旧基线，并绑定 manifest 指纹、strategy、seed、Redis run identity、工具版本与本地 exe/容器身份。
- 每轮严格执行“初始 restore -> 可选 warmup -> warmup 后 restore -> cache-state 准备 -> measured trial -> 先落盘 pending/incomplete 报告 -> measured 后 restore -> 完整报告”；进程中断或恢复失败不会留下可误用的 complete 报告，矩阵立即停止。
- MySQL 只删除 benchmark 专属作者基线之后的精确帖子/Outbox ID；Redis 只恢复 manifest 可枚举的 Inbox/BigV key，并只删除新帖子可推导的缓存/processing key，不使用 pattern/SCAN。
- 测量报告保存全部成功发布 post IDs；测量后 restore 必须证明这些 ID 全部位于删除集合，否则记 `mutation_restore_scope` 并判无效。checkpoint ID、baseline fingerprint、创建时间、工具版本进入 compatibility fingerprint，比较器拒绝跨基线汇总。
- 定位并修复 go-zero `kq.Pusher` 默认 1 秒异步 flush 导致的“瞬时 lag=0 假排空”：mutation drain 现在要求 lag=0 且 log-end offset 连续稳定 2 秒；延迟到达的消息会重置窗口。
- Redis TTL 校验区分两个时点：restore 后立即校验要求绝对误差不超过容差；独立 `verify-mutation-checkpoint` 允许 TTL 自 checkpoint 创建后自然减少，但键缺失、内容变化、永不过期属性变化或 TTL 被重新延长仍 fail closed。
- 90/10 与 80/20 使用一个全局有界并发池，按操作序号均匀交错写请求；固定请求数比例精确，时长模式尾部误差有界，不再按两个独立 worker 池近似比例。
- 真实短测均在 `compose-middleware-local-services`、Hybrid、treatment、cold 下通过：RPC publish c2（12/12 发布并删除）、RPC 90/10 c4（175 读/19 写，实际 90.21%）、RPC 80/20 c4（64 读/16 写，80.00%）、Gateway 90/10 c2（90 读/10 写，90.00%）；四轮均 `complete=true`，恢复后抽查 Inbox 回到 48 成员。它们只验证闭环，不是容量数字。
- 终审后补强报告审计门禁：首次 restore 前先落 `mutation_preparation_pending`；任何准备失败会覆盖为永久 incomplete；临时 pending 报告首次写入失败后即使最终写入成功也保留 `mutation_provisional_report`，不得升级 complete。
- 正式矩阵只接受全新非空命名空间，先保存 `expected-matrix.json`，结束时逐格核对 entry/scenario/concurrency/trial、期望轮次、报告完整性和 after-trial restore，并拒绝缺格、重复参数、额外陈旧报告；最后再次独立 verify checkpoint 后才运行 comparison。
- 补强后的正式脚本端到端短演练 `feed-wp11-hot-mutation-script-smoke-v2-20260818` 已通过：RPC 90/10 c2、1 秒预热、2 秒采样、1 轮，90 读/10 写、失败和超时为 0、10 个成功发布 ID 全部进入删除集合，测量后 restore `complete=true`，脚本末尾 checkpoint 独立验收再次通过。该演练只证明脚本编排和 fail-closed 报告链路，不进入容量对比。
- 正式入口为 `cmd/loadtest/run_feed_mutation_matrix.ps1`。脚本支持 publish、90/10、80/20 与 RPC/Gateway，但 2026-08-27 的最终批准范围固定为 RPC-only：publish c2、90/10 c4、80/20 c4，10 秒预热 + 60 秒采样 × 3；Control 与 Treatment 使用不同 ReportRunId、相同 checkpoint。后续 Gateway 性能矩阵取消。

### 18.2.1 2026-08-27 mutation RPC 正式结果

全新 RunID 为 `feed-wp11-hot-mutation-ab-20260827a`，checkpoint ID 为 `43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`。Control/Treatment 各 9 轮、合计 18 轮均 `complete=true`，失败/超时为 0，初始、预热后、试后 restore 与 Kafka drain 全部通过；两侧 manifest、checkpoint、baseline、Git patch 和业务二进制身份一致。

| RPC 场景 | Control QPS / P95 | Treatment QPS / P95 | 结论 |
|---|---:|---:|---|
| publish c2 | 6.61 / 398.19ms | 6.33 / 388.98ms | QPS -4.33%，P95 -2.31% |
| 90/10 c4 读取 | 58.66 / 3.23ms | 55.23 / 2.64ms | QPS -5.85%，P95 -18.43% |
| 90/10 c4 发布 | 6.52 / 729.46ms | 6.14 / 779.37ms | QPS -5.85%，P95 +6.84% |
| 80/20 c4 读取 | 27.41 / 3.21ms | 24.04 / 3.19ms | QPS -12.28%，P95 -0.81% |
| 80/20 c4 发布 | 6.85 / 712.67ms | 6.01 / 828.74ms | QPS -12.28%，P95 +16.29% |

Treatment 在 90/10、80/20 中将 Relation/request 分别从 1 降至 0.0699、0.1484，但页面 Fresh 仅为 50.40%、27.18%，并且完整发布平均耗时上升。固定比例 worker 池使读写吞吐联动，因此不得宣称 mutation 总吞吐提升；由于 A/B 一次开启整套 Treatment，持续失效、回填和共享依赖竞争只能列为候选解释，需单开关或 Profile 验证。完整证据见 `docs/superpowers/reports/2026-08-27-feed-wp11-mutation-rpc-ab.md`。

### 18.3 2026-08-17 压缩档纯读执行结果

固定条件：

- 拓扑为 `compose-middleware-local-services`：MySQL/Redis 使用既有宿主服务，Kafka/etcd 等使用既有中间件容器，Gateway/KnowPost/Relation/Counter/User/Search 本地运行。
- Hybrid、page1/size20、60 秒采样、三轮中位数；Gateway 使用本地签发但仍经过真实 JWT 验签的 RS256 token。
- 日志预设固定为 `benchmark-error-only`：go-zero level=error、stat=false、RPC Stat middleware=false、Gateway access log=false、SQL statement info=false。错误日志仍保留，不能把该预设描述为“完全无日志”。
- Windows Redis 在此前长压中因 RDB 后台保存失败触发 `stop-writes-on-bgsave-error`，受影响旧报告全部排除。有效报告期间临时设置 `save ""`、`stop-writes-on-bgsave-error no`；原配置证据保存在 `.tmp/feed-local/redis-benchmark-original.json`。这是开发压测条件，不是生产配置建议。
- 所有下表报告均 `complete=true`、请求失败/timeout=0、无 `feed_degraded`。唯一例外是 high Treatment RPC c32 的边界轮：7768.72 QPS / P95 10.76ms，但 `feed_degraded=true`，KnowPost 同期记录 refresh queue full，因此明确判无效并停止剩余 trial。

| Cardinality | 入口 | Control 稳定档 | Control 中位 QPS / P95 | Treatment 稳定档 | Treatment 中位 QPS / P95 | 解释 |
|---|---|---:|---:|---:|---:|---|
| distributed≈1,200 | RPC | c64 cold | 6,797.75 / 14.06ms | c64 l2-warm | 102,075.49 / 1.16ms | 目标运行态 15.02×；Fresh=100%，但不是同缓存状态严格 A/B |
| distributed≈1,200 | Gateway | c16 cold | 1,218.99 / 23.02ms | c16 l2-warm | 1,529.10 / 26.34ms | 吞吐约 1.25×，P95 反而变差；Fresh 中位约 92%，冷计算约 0.19/请求 |
| high=20,000 | RPC | c32 cold | 6,162.47 / 7.55ms | c16 cold | 9,477.51 / 3.73ms | 各自稳定最优档约 1.54×；Treatment c32 已触发降级，不能宣称同并发收益 |
| high=20,000 | Gateway | c8 cold | 1,402.80 / 11.20ms | c8 cold | 1,407.80 / 15.16ms | 同并发同 cache-state，吞吐约 1.00×且 P95 回退；Fresh 中位约 23%，冷计算约 0.81/请求 |

结论与下一步：

1. 整页缓存对能够维持热点的 RPC 入口非常有效，distributed 卡稳定超过 100K QPS；“之前出现 10K、后来无法复现”主要来自拓扑、缓存热度、日志、Redis 持久化错误和测试口径混用，不能把任一单轮数字当作统一单机容量。
2. 20,000 随机读者的 Gateway 链路无法在当前 800ms L1 / 4s L2 Fresh 窗口内维持页面热点。Fresh 下降会带动 cold compute 和 Relation/Redis 依赖恢复，吞吐出现稳态漂移；继续堆并发只增加排队。
3. 当前最高优先级不是内存池或更大 pipeline，而是 Cursor/seek 分页、按 cardinality/访问频率调整缓存策略，以及拆分 Gateway/鉴权/HTTP 与 Feed 内核的 Profile。深分页优化必须独立建立 page/cursor 正确性和性能基线，不能拿 page1 缓存数字外推。
4. 本段纯读完成时，写入、90/10、80/20 尚未形成正式容量结论；18.2 的 checkpoint/restore 和 RPC 正式 A/B 已于 2026-08-27 后续完成，结果见 18.2.1。

完整报告见 `docs/superpowers/reports/2026-08-17-feed-hybrid-wp11-capacity.md`，原始数据位于对应的 `results/feed-loadtest/feed-wp11-*-fixed-*-formal-nolog-20260817a` 命名空间。

## 19. Cursor/seek 深分页增量设计

状态：2026-08-17 已批准并完成 C0-C5。Cursor 分页、专用深分页数据集、RPC/Gateway 正确性和正式 A/B 已落地；本节不授权 Feed Head 物化。Cursor 工作当时不解除 18.2 的数据隔离门禁，该门禁与 mutation RPC 正式 A/B 已于 2026-08-27 后续完成。

### 19.1 证据边界与问题定义

当前页码路径会先计算：

```text
start = (page - 1) * size
end = start + size
candidateLimit = end + size
```

随后从 Inbox 和每个大 V Outbox 重新读取头部候选，执行归并、去重、可见性过滤、详情装载，最后丢弃前 `start` 条。其读取和计算成本随 `page * size` 增长；当前实现最多扩大到 5,000 个候选，深页还会重复处理已经在前页处理过的内容。

现有 WP11 distributed/high manifest 不足以验证这一问题：每个读者只有约 8–14 条种子内容，而原 `deep-page` 至少请求 `page=5,size=20`，多数请求得到的是空页。该结果只能证明“深页不命中第一页 PageCache”，不能证明真实深数据下的页码放大程度，也不能作为 Cursor 收益数字。Cursor 上线前必须使用 19.9 的专用深数据集重新建立基线。

### 19.2 目标与非目标

目标：

- 保留现有 `page/size` 客户端兼容性，同时为个性化 `GetUserFeed` 增加稳定 seek 分页。
- 静态关注关系和静态内容集合下，以 `(Feed ZSet 时间分数, postID)` 全序保证连续翻页无重复、无遗漏。
- 将单页候选读取成本从随页深增长改为主要由 `size`、Feed 来源数和同时间戳 tie group 决定。
- 用同二进制、同数据、同缓存和同资源的页码/Cursor A/B 决定 Cursor 是否值得默认启用，不预设收益结论。

非目标：

- 不提供跨多次请求的数据库快照隔离；Feed 仍是最终一致的动态集合。
- 不为 Cursor 深页增加整页 PageCache，不缓存任意游标组合。
- 不修改 `INBOX_MAX_SIZE=1000`、`BIGV_OUTBOX_MAX_SIZE=100` 或推拉路由策略。
- 不在本工作包中实现 Feed Head 物化、用户分层或 mutation checkpoint/restore。
- 不把公共 Feed、我的发布列表或 Relation 列表分页混入本次 A/B。

### 19.3 API 与兼容契约

Protobuf 只做追加字段：

```proto
message FeedPage {
  repeated FeedItem items = 1;
  bool has_more = 2;
  int32 size = 3;
  int32 page = 4;
  string next_cursor = 5;
}

message GetUserFeedReq {
  int64 user_id = 1;
  int32 page = 2;
  int32 size = 3;
  string cursor = 4;
}
```

Gateway 的 `/api/v1/knowposts/following-feed` 增加可选查询参数 `cursor`，JSON 响应增加 `nextCursor`。规则固定为：

1. `cursor` 为空时保持页码语义；首页仍是 `page=1&size=20`。
2. 功能开关开启后，页码响应在存在下一页时也返回精确 `nextCursor`，便于旧客户端渐进迁移；旧客户端忽略新增字段即可。开关关闭时该字段为空。
3. `cursor` 非空时进入 Cursor 模式并忽略 `page`，响应 `page=0`；`size` 仍限制在 1–100。
4. Cursor 模式有更多内容时返回 `nextCursor`；已到末尾、返回空页或没有更多内容时返回空字符串。
5. Cursor 最大长度 256 字节；Base64URL、结构、版本、时间和帖子 ID 任一非法都返回 gRPC `InvalidArgument`，Gateway 映射为 HTTP 400，禁止静默退化为首页。
6. Cursor 功能开关关闭时，带 Cursor 的请求明确失败，禁止把它当页码请求执行。

Cursor v1 是 Base64URL 编码的版本化载荷：

```text
{ version: 1, sort_time_seconds: int64, post_id: int64 }
```

载荷只包含非敏感排序位置，既不作为身份凭证也不替代当前用户授权。v1 不增加 HMAC 配置：客户端修改一个结构合法的游标只会改变自己 Feed 的读取起点，不会突破关注关系、内容可见性或用户身份边界。解码必须拒绝未知字段、非正整数、未知版本和非规范编码，为以后升级保留版本空间。

### 19.4 排序、不等式与同秒 tie-break

当前 Feed 的规范顺序继续定义为：

```text
A comes before B
  if A.sortTime > B.sortTime
  or A.sortTime == B.sortTime and A.postID > B.postID
```

因此位于游标 `(T, P)` 之后的候选必须严格满足：

```text
candidate.sortTime < T
  or candidate.sortTime == T and candidate.postID < P
```

所有 Feed 写入路径必须继续把整数 Unix 秒写入 Redis ZSet score；帖子 ID 是正整数。Redis 相同 score 下按 member 字符串排序，不能直接当作数值帖子 ID 的 tie-break。Cursor 读取每个 ZSet 时先在第一轮 Pipeline 中读取：

1. 读取 `score == T` 的完整 tie group，在应用层按数值 `postID < P` 过滤。
2. 读取 `score < T` 的前 `candidateLimit` 条。

第二类限量结果也可能截断更早的某个同秒组。如果其返回数等于 `candidateLimit`，必须取最后一条的 score 作为边界，在第二轮 Pipeline 中补齐每个受影响 ZSet 的完整 `score == boundaryScore` 组；返回数不足上限时不需要第二轮。所有结果合并后再按数值全序维护 Top-N。

tie group 不是无界数据：Inbox 最多 1,000 个成员，每个大 V Outbox 最多 100 个成员；测试必须覆盖达到该上界的同秒数据。禁止采用“只补游标所在秒”或“多读固定 20/100 条然后猜测足够”的近似实现，因为 `score < T` 的截断边界同样可能按 Redis 字符串序切开一个同秒组并永久漏帖。两轮方案以至多一次额外 Redis 往返换取无数据迁移的精确性；报告需单列 Cursor Pipeline 轮数和命令数。

### 19.5 组件边界与读路径

实现拆成四个可独立测试的单元：

1. `feedcursor` 编解码器：只负责版本化载荷、规范编码和输入校验，不依赖 Redis、RPC 或业务 Logic。
2. Cursor Redis Reader：以可选窄接口提供批量 `score == T` 与 `score < T` 读取；生产 `RedisAdapter` 实现 Pipeline，现有 Writer `RedisClient` 不因 Cursor 扩展而被迫修改全部测试 fake。
3. `FeedReadSnapshot` Cursor 读取：复用一次请求内已固化的关注列表和大 V 分类，读取 Inbox/BigV、精确 seek、归并、去重并返回带排序键的 `[]Post`。
4. `GetUserFeedLogic`：负责解码 Cursor、可见性过滤、详情批量装载、有限回填以及生成最后一个真实可见条目的 `nextCursor`。

现有 `FeedReader.GetFeed(...page,size)` 和 RPC 方法名保持不变。内部新增返回 `Post` 而非只有 ID 的读取入口，使页码首页也能从 Redis 的真实排序键生成 `nextCursor`；外部旧调用仍可把 `Post` 映射为 ID，不扩散内部结构。

Cursor 冷路径为：

```text
GetUserFeed(cursor,size)
  -> decode and validate cursor
  -> Prepare route snapshot once
  -> pipeline Inbox + BigV cursor-tie ranges and bounded older ranges
  -> when bounded older results are full, pipeline their complete boundary tie ranges
  -> merge + deduplicate + keep candidateLimit
  -> batch hydrate and visibility filter
  -> bounded backfill when filtered rows leave fewer than size+1 visible posts
  -> return first size visible posts and cursor of the last returned Post
```

初始 `candidateLimit=size+1`。若删除、转私密、取消关注或缺失详情导致可见条目不足，但原始候选仍可能有更多，则按现有策略成倍扩大上限并重新读取同一游标窗口，最大不超过 `maxPersonalFeedCandidates`。同一请求只允许执行一次 `Prepare`，不得因回填重复调用 Relation/Counter。

若生产 RedisAdapter 不具备精确 Cursor 接口，Cursor 请求必须失败并记录降级原因，不能静默走 rank/offset；页码路径继续正常工作。查询 Context、超时和取消信号沿用当前请求，不创建无界 goroutine。

Cursor Pipeline 中任一必需范围读取失败都必须使整页失败，客户端可使用同一个 Cursor 安全重试；禁止返回缺少某个 Inbox/BigV 范围的部分页面，否则下一次 Cursor 会越过失败来源中的内容并造成永久遗漏。旧页码路径既有的部分降级语义保持不变，不在本工作包中扩大行为修改。

### 19.6 PageCache、FeedItem Cache 与一致性

- PageCache 继续只接纳 Hybrid `page=1,size=20`；任意非空 Cursor 都旁路整页缓存。
- 第一页缓存值包含 `nextCursor`，其 key 继续受 relation epoch 和 content-safety epoch 约束。
- Cursor A/B 的正式吞吐测试统一关闭 PageCache，避免把首页命中收益混入分页算法收益；RouteSnapshot、合并 Pipeline 和 FeedItem cache 配置在 A/B 两侧必须完全一致。
- Cursor 必须来自 Redis 候选的 `sortTime/postID`，不能用 `FeedItem.publishTime` 反推。数据库时间为毫秒，Feed Writer 当前另取整数秒，两者在秒边界可能不同，混用会产生重复或遗漏。
- Cursor 不固化 relation/safety epoch，也不承诺会话快照。每一页仍按当前关注关系和当前可见性过滤，删除、转私密和取关不会因为旧 Cursor 泄露内容。
- 翻页期间新增且排序位置新于 Cursor 的帖子不会插入后续页，因此不会挤动已经读过的窗口；用户回到首页后才能看到它们。
- 删除或转私密可能使后续页缩短，Logic 应向后回填；关注关系变化可能改变尚未读取的集合。这是已经接受的 Feed 短暂不一致取舍，不伪装成快照一致性。

### 19.7 功能开关、可观测性与回滚

新增 `Feed.CursorPagination.Enabled`，默认关闭。开关关闭时不生成 `nextCursor` 并拒绝 Cursor 请求，旧页码路径完全不变；回滚只需关闭开关，无需迁移或删除 Redis 数据。

至少记录以下低基数指标：

- Cursor 请求数、成功数、非法游标数和精确 Reader 不可用数。
- Cursor 每次请求读取的 Redis 成员数、tie-group 成员数、归并候选数、详情装载 ID 数和回填轮数。
- Cursor decode、Redis seek、merge/dedup、hydrate 和 total 延迟。
- `page` 与 `cursor` 两种模式标签；标签值固定，禁止把游标内容、用户 ID、page 数或 post ID 放入指标标签。

错误日志只保留采样后的非法游标摘要和依赖失败，不打印完整 Cursor，不恢复压测阶段已关闭的高频成功日志。

### 19.8 正确性门禁

生产代码前必须先看到以下自动化测试按预期失败，再实现最小代码使其通过：

1. Codec：v1 round-trip、非规范 Base64、超长、截断、未知字段、未知版本、零/负值和尾随数据。
2. 排序：不同时间、相同时间不同 ID、跨 Inbox/BigV 重复成员和 tie group 上界。
3. 静态 oracle：将完整候选在内存中按规范全序排序，与 Cursor 连续翻页到末尾的结果逐项相等，重复和遗漏均为 0。
4. 插入：读取第一页后插入更晚帖子，后续页不重复、不把新帖插入旧窗口；重新读取首页可看到新帖。
5. 删除/转私密：游标后的候选失效时不返回该内容并正确回填；最后一页 `hasMore=false,nextCursor=""`。
6. 关系变化：取消关注后旧 Cursor 不得返回该作者内容；新增关注允许只影响尚未读取的动态集合。
7. 兼容：旧 page 请求和 JSON 字段保持可用；页码首页新增的 `nextCursor` 对旧客户端透明。
8. Gateway：认证用户 ID 仍来自 Token；Cursor 只透传给当前用户 RPC，非法参数在 RPC 前或 RPC 中得到确定错误。
9. Adapter：真实 Redis 覆盖同 score 范围、exclusive older range、Pipeline 结果偏移、部分命令失败和 Context 取消。
10. 回退：关闭开关或缺少精确 Cursor Reader 时，Cursor fail closed，旧页码路径不受影响。

### 19.9 专用深分页数据集

新增独立 manifest，例如 `feed-cursor-deep-v1`，不得覆盖或重用 WP11 hot/distributed/high manifest：

| 维度 | 固定值 |
|---|---:|
| 读者 | 20 |
| 普通作者 | 20 |
| 大 V 作者 | 5 |
| 每个普通作者帖子 | 50 |
| 每个大 V 帖子 | 100 |
| 每个读者 Inbox 候选 | 1,000 |
| 每个读者 BigV Outbox 候选 | 500 |
| 每个读者总候选 | 约 1,500 |

20 个读者都关注上述 25 个作者；普通作者的 1,000 条内容进入每个读者 Inbox，大 V 的 500 条内容保留在 5 个 Outbox。容量测试使用静态、可见、已发布内容，避免 mutation 影响吞吐口径；正确性子数据另包含删除、转私密和翻页间插入。

数据时间必须覆盖足够跨度，并至少构造：

- 100 条以上相同秒级 score、不同数值 post ID 的 tie group。
- Inbox 与 BigV Outbox 中重复的 post ID，用于验证去重。
- page1、page5、page20、page50 均有非空结果。
- 连续读取直到候选耗尽时有确定 oracle 和指纹。

Manifest 记录 seed、每类帖子数、ZSet member/score 指纹、MySQL 行指纹和预期全序摘要。setup 完成后必须验证 Inbox/Outbox 容量、DB 行数、首尾排序键和 oracle hash；任何不一致都禁止进入正式 A/B。

### 19.10 页码/Cursor A/B 设计

严格 A/B 使用同一份专用 manifest、同一被测二进制、同一 Git patch、同一进程拓扑、同一资源、同一 Hybrid 策略和相同功能开关。两侧只改变分页方式：

- Control：页码 `page=N,size=20`。
- Treatment：从同一首页获得 Cursor，按 `nextCursor` seek 到等价深度。

PageCache 在两侧都关闭；RouteSnapshot、Combined Pipeline、FeedItem cache、日志级别和 Redis 持久化条件完全相同。RPC 和 Gateway 分开报告，禁止把入口差异算作 Cursor 收益。

固定深度的 page5/page20/page50 比较只测“目标页单次请求”的服务成本。Treatment 所需的前序 Cursor 链在 warmup/准备阶段生成并按读者校验，不计入目标页采样窗口；Control 同样不计入客户端构造页码的准备成本。`sequential-50` 则把首页到第 50 页的全部请求和准备成本都计入，用来比较真实连续滚动的累计成本。两类数字必须分栏报告，禁止用单页 Treatment 延迟冒充完整滚动耗时。

场景：

| 场景 | 验证内容 |
|---|---|
| page1 | 新代码对首页的兼容和非回退 |
| page5 | 浅层翻页成本 |
| page20 | 中等深度放大 |
| page50 | 接近当前保留上限的深页成本 |
| sequential-50 | 单个虚拟用户从首页连续读取 50 页的累计成本与正确性 |
| same-second | 大 tie group 下的正确性和最坏读取量 |

执行先以 c16/c32/c64、每点 10 秒 scout 找稳定档；正式阶段只跑最优档和一个相邻档，预热 10 秒、采样 60 秒、重复三次并报告中位数/min/max。若 scout 已出现超时、错误、客户端饱和、服务降级或资源队列积压，不继续更高并发。A/B 每轮仍需满足 WP10/WP11 的报告完整性、进程身份、Redis identity、零 eviction/rejection 和客户端未饱和门禁。

每个结果必须同时报告：

- 成功 QPS、P50/P95/P99、错误和超时。
- KnowPost/Gateway CPU、RSS、GC、Redis ops，以及 Redis 实际返回成员数。
- 每成功请求的候选数、tie-group 数、归并数、详情装载数和回填轮数。
- 页深与 P95/CPU/request/候选数的变化曲线。
- sequential-50 的总请求时间、总 Redis 返回成员、总详情装载 ID、重复数和遗漏数。

### 19.11 验收与决策规则

硬门禁：

- 静态 oracle、同秒 tie-break、插入、删除、可见性和取消关注测试全部通过，重复/遗漏/越权返回均为 0。
- 旧页码 API、第一页 PageCache、RPC/Gateway 兼容测试和现有 Feed 全量回归通过。
- page1 相对改造前同口径基线的 QPS 或 P95 回退均不超过 5%。
- Cursor page5/page20/page50 的单页候选成本不随页深近似线性增长；page50 相对等价页码请求的 Redis 返回成员或详情装载 ID 至少降低 70%。
- Cursor page50 的 P95 不高于 Cursor page5 的 1.5 倍；超过时必须用 tie group、依赖或资源证据解释，不能直接上线。

“Cursor 性能必要性得到证明”必须在硬门禁通过后再满足以下至少一项：

1. page20 或 page50 的 P95 相对页码方式改善至少 30%，且错误率、P99 和 CPU 不回退。
2. sequential-50 的 CPU/request、Redis 返回成员或详情装载量至少降低 50%，且总耗时和 QPS 不回退。

如果只通过正确性而达不到性能决策规则，代码保持默认关闭，报告结论应是“接口具备但当前规模下尚未证明需要默认启用”，不得选择性展示最好单轮。若通过，先在 RPC 和少量 Gateway 流量启用，再观察非法 Cursor、依赖错误、P99 和回填指标后逐步扩大。

### 19.12 实施工作包

Cursor 阶段按以下顺序执行，每个包独立测试、独立审查：

1. C0：冻结本节 Spec，建立内存 oracle 与深数据 manifest 契约。
2. C1：追加 Proto 字段，实现 Cursor codec、参数错误和生成代码兼容测试。
3. C2：实现精确 Redis seek Pipeline 与 `FeedReadSnapshot` Cursor 路径。
4. C3：接入 `GetUserFeedLogic`、有限回填、`nextCursor`、Gateway 和功能开关。
5. C4：扩展 loadtest 的 Cursor 客户端、连续翻页 oracle、候选指标和报告字段。
6. C5：创建深数据集，先跑正确性/短 scout，再跑选定并发的正式 A/B，形成独立报告。

每个生产行为必须遵循 red-green-refactor：先增加能因缺少该行为而失败的最小测试，确认失败原因正确，再实现最小代码并运行相关回归。C1–C3 未通过正确性门禁前不得开始长时间压测；C5 不与 18.2 的写入/混合矩阵并行执行或共用数据命名空间。

### 19.13 C5 可执行数据契约

`cursor-deep` 使用独立 `setup-cursor-deep`/`seed-cursor-deep` 阶段和 `cursor-ab` FeaturePreset。`cursor-ab` 与 treatment 相同地开启 RouteSnapshot、关系/安全 epoch、Combined Pipeline、Cursor 和可观测性，但把 PageCache 固定为 `off`；其 CacheState 只能是 `cold` 或 `natural`。运行时实际 Flag 与预设不一致时压测器必须 fail closed。

Manifest 在真实发布尚未完成或任一证据缺失时保持 `cursor_deep.ready=false`。发布成功项同时写入 `posts` 和 `post_authors` 后落盘；续跑按每作者已确认数量只补缺口。完成发布后必须依次通过 Kafka lag=0、原始或已归一化 Redis 成员集合、确定性 score 归一化、所有读者相同 oracle、MySQL ownership/status/visible/publish_time 以及 Redis/MySQL SHA-256 指纹，最后才能设置 Ready。客户端超时导致“服务端可能提交但未获得 post ID”时，不承诺自动恢复；额外成员会使 Redis 门禁失败，禁止手工修改 Ready 接纳。

正式执行入口为 `cmd/loadtest/run_feed_cursor_ab.ps1`。Scout 默认 c16/c32/c64、10 秒、单轮；Formal 必须显式传入 Scout 选出的稳定档及至多一个相邻档，默认 60 秒、三轮。脚本固定 Hybrid、PageCache off、cold、duration 模式和同一 manifest，并将 RPC/Gateway 分开写入报告。

### 19.14 C5 执行结果与决策

2026-08-17 使用 `feed-cursor-deep-v1` 完成真实中间件、宿主机业务进程拓扑下的正式验证。数据集为 20 个读者、1,000 个 Inbox 候选、500 个大 V Outbox 候选、1,500 个原始帖子、1,499 个去重 oracle、120 条同秒 tie group 和 1 条跨来源重复。Scout 选择 c16 为稳定档、c32 为相邻过载边界；正式点均为 10 秒预热、60 秒采样、三轮。

关键 c16 三轮中位数：

| 入口/目标页 | 页码 QPS / P95 | Cursor QPS / P95 | Redis 成员 | 详情 ID | 分配 KB/请求 |
|---|---:|---:|---:|---:|---:|
| RPC page20 | 1,178.8 / 22.67ms | 4,848.8 / 5.38ms | 921→67 | 420→21 | 721.9→64.5 |
| RPC page50 | 592.4 / 35.66ms | 7,621.7 / 3.36ms | 1,500→24 | 1,020→21 | 1,667.7→58.9 |
| Gateway page20 | 1,045.2 / 25.41ms | 1,698.2 / 15.72ms | 921→67 | 420→21 | 722.4→64.2 |
| Gateway page50 | 489.2 / 49.02ms | 2,180.4 / 11.37ms | 1,500→24 | 1,020→21 | 1,669.1→58.6 |

同一 patch 下的连续 50 页严格配对中，RPC 页请求 QPS 为 991.0→4,974.6、P95 为 28.02→5.42ms；Gateway 为 801.8→1,797.0、36.98→14.38ms。每页平均成员 1,018.9→73、详情 ID 530→21.4、分配约 923→65KB。正式有效报告和同秒边界 probe 的失败、超时、重复、oracle mismatch、Cursor loop、提前结束均为 0。

因此 19.11 的 page20/page50、工作量和连续滚动性能门槛均通过，Cursor 性能必要性已证明。改造前未保存同 manifest 的 page1 二进制正式基线，且本机 Gateway page5 受 tie group 与宿主机调度影响、收益不确定，所以保持功能开关可回滚，先在 RPC 与少量 Gateway 流量灰度，不直接宣称全页型、全流量默认开启。完整报告见 `docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md`。
