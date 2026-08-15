# Feed 推拉混合读路径性能演进设计

日期：2026-08-15

状态：已批准，待实施计划

适用范围：KnowPost 个性化 `GetUserFeed`、Relation 路由依赖、Redis Feed 数据、Gateway 对照入口和 Feed 压测工具

## 1. 决策摘要

本设计采用渐进式分层读路径，不推翻现有推拉混合架构：

1. 先建立可解释的 CPU、分阶段延迟、缓存和依赖调用基线。
2. 将现有“只缓存大 V 分类”的路由缓存升级为完整 Feed 路由快照，消除热路径每次调用 Relation RPC 的成本。
3. 合并 Inbox 与大 V Outbox 的 Redis Pipeline，继续降低 Hybrid 冷路径成本。
4. 只对个性化 Feed 第一页增加独立的 L1 + Redis L2 最终结果缓存，并加入 epoch 校验、Singleflight、TTL 抖动、stale-while-revalidate 和有界刷新。
5. 只有前述优化仍不能满足目标时，才为热点、活跃用户引入异步 Feed Head 物化和 Cursor 分页。

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
- 多实例扩展前再评估 Redis 租约，单实例阶段不提前引入分布式锁。

## 7. 关键数据流

### 7.1 热点第一页命中

1. 根据用户、策略、页码和尺寸判断是否允许页面缓存。
2. 从微型 L1 获取 relation epoch 和 content-safety epoch；过期时从 Redis 刷新。
3. 构建完整 page key。
4. L1 Fresh 命中，直接返回不可变 `FeedPage`。
5. 全程不调用 Relation、Counter、Inbox、BigV Outbox 或 FeedItem DB。

### 7.2 页面缓存 miss

1. L1 miss 后读取 Redis L2。
2. L2 命中并且 value 中的 epoch 与当前 epoch 一致，则反序列化、回填 L1 并返回。
3. L2 miss 进入相同 key 的 Singleflight。
4. 双检 L1/L2，仍 miss 后读取 FeedRouteSnapshot。
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

1. 数据写入成功后更新详情和 FeedItem 缓存。
2. 原子递增全局 `feed:content:safety:epoch`。
3. 当前实例同步更新本地 epoch；其他实例在约 1 秒内重新读取 Redis epoch。
4. 所有旧 Page Cache key 因 safety epoch 不匹配而不可使用，包括 Stale 页面。
5. 若 Redis epoch 更新失败，Page Cache 进入 bypass，冷路径负责过滤不可见内容。

全局 safety epoch 会使全部第一页缓存失效，但删除、转私密、下架和置顶变化相对发布低频。该方案用罕见的全局冷却换取简单、可证明的撤销正确性；只有实测出现频繁全局抖动时才拆成作者或帖子级版本。

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

### Phase 2：Hybrid 冷路径 Pipeline

实施：

- 将 Inbox 与 BigV Outbox 放入同一受控 Pipeline。
- 保留每批最大 128、部分错误隔离和不支持批量接口时的兼容路径。
- 只根据 Phase 0 Profile 决定是否优化归并、切片分配或对象复用。
- 不在没有证据时引入动态并发控制、sync.Pool 或自定义协议。

门禁：

- 冷路径 P95 至少下降 15%，或 Redis commands/request 至少下降 20%。
- 冷路径 QPS、内存和 GC 不得恶化超过 10%。
- Pipeline 部分失败、超时和 Context 取消测试通过。

### Phase 3：第一页最终结果缓存

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
- Cursor 分页作为该阶段独立子设计，不与 Phase 3 同时上线。

该阶段必须另写设计增量并重新批准，当前 Spec 不授权直接实现全面物化。

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
- 正式报告同时采集 KnowPost 与 Relation，避免只优化表面服务。

## 11. 配置、发布与回滚

建议配置开关：

```yaml
Feed:
  Strategy: hybrid
  RouteSnapshot:
    Enabled: false
    TTL: 5s
    EpochL1TTL: 1s
  CombinedPipeline:
    Enabled: false
    BatchSize: 128
  PageCache:
    Mode: off # off | l1 | l1-l2
    Page: 1
    Size: 20
    L1FreshTTL: 1s
    L2FreshTTL: 5s
    StaleTTL: 10s
    JitterPercent: 20
    RefreshWorkers: 32
    RefreshQueue: 1024
  FeedHead:
    Enabled: false
```

具体字段需遵循当前 go-zero 配置解析方式；时长字段若不能直接解析 `time.Duration`，实施计划中统一改为毫秒字段，禁止在不同配置文件里混用单位。

发布顺序：

1. 先上线指标和 pprof，所有性能开关关闭。
2. 单独启用 RouteSnapshot，完成 Phase 1 验证。
3. 单独启用 CombinedPipeline，完成 Phase 2 验证。
4. PageCache 先启用 L2，再启用 L1，最后启用 SWR。
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
- 决定：敏感变化递增全局 epoch，所有旧页面失效。
- 理由：实现简单、可证明，敏感变化相对低频。
- 代价：一次敏感变化会造成全局第一页缓存冷却。

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
- Phase 4 未经独立设计批准不会开始实现。

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
