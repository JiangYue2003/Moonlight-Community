# 可靠性设计：Outbox 幂等、Redis 主从重建、Kafka 高可用

本文档针对三个具体的可靠性问题给出基于当前代码/配置的现状说明,并在代码没有覆盖到的地方(Redis 主从重建、Kafka 多实例)给出通用方案与本项目落地时需要补的东西——这两块目前项目的 dev compose 里都是单点部署,文档里会明确标注哪些是"代码已经做了"、哪些是"需要额外基建但代码接口已经预留"。

## 1. outbox 被 Kafka 重复投递,如何保证幂等

### 现状:三层幂等保护叠加

Kafka 本身的消费模型是**至少一次(at-least-once)**：`pkg/kafkax/consumer.go` 的 `RunConsumer` 采用手动提交 offset(`CommitInterval: 0`),只有 `handler` 返回 `nil` 才提交(`consumer.go:124-130`)。这意味着只要处理成功但提交前进程崩溃、或者 consumer group 发生 rebalance,同一条消息就会被重新投递。项目没有寄望于"Kafka 不重复投递",而是在消费端做了幂等,具体分三层:

**第一层:Redis SETNX 去重(消息级幂等)**

四个消费方(`search-indexer`、`relation-syncer`、`agent-indexer`、`llm-ragindexer`)都实现了同一套 `Dedup` 结构(如 `services/search/indexer/internal/processor/dedup.go`):

```go
func (d *Dedup) Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
    return d.rdb.SetNX(ctx, key, "1", ttl).Result()
}
```

去重 key 以 **outbox 表的自身主键 `id`** 拼接事件类型构造(而不是聚合根 id),例如 `dedup:idx:{type}:{outboxId}`（search-indexer,`processor.go:55`)、`dedup:rel:{type}:{outboxId}`（relation-syncer)。因为 `outbox.id` 是 Snowflake 生成的全局唯一值,同一行 outbox 记录无论被 Canal/Kafka 重复推送多少次,`SETNX` 只会在第一次成功,后续重复的 `Acquire` 返回 `false`,处理函数直接 `return nil`——注意是返回成功(`nil`),不是失败,这一点很关键:如果对重复消息返回 error,反而会触发 `kafkax` 的重试机制,造成无意义的退避等待。

TTL 的选择(`Dedup.TtlSeconds`,relation-syncer 配置为 `86400`(1天),search-indexer 为 `600`(10分钟))决定了这层去重的**时间窗口**而非永久幂等——理论上如果同一条消息在 TTL 过期之后才被重新投递(比如 consumer 下线超过一天后重新上线,offset 没提交导致 Kafka 重新推送积压的老消息),去重会失效。这是一个需要注意的边界:去重窗口要大于"最坏情况下 offset 提交延迟 + consumer 下线恢复时间",目前 relation-syncer 给到 1 天是相对保守的取值,search-indexer 只有 10 分钟,如果 search-indexer 曾经长时间下线过再恢复,存在重复处理的风险(不过 search 这条链路的重复处理本身是幂等的,见下文第二层,所以影响有限)。

**第二层:业务逻辑本身的幂等设计(即使去重失效也不产生脏数据)**

即便第一层去重失效导致同一事件被处理了两次,下游的具体写操作本身也大多设计成幂等:

- `search-indexer` 的 `upsert`/`softDelete` 都是**用实体 ID 做覆盖写**（`Es.Index(ctx, index, postId, doc)` 是全量覆盖,`Es.Update` 是部分字段覆盖),重复执行结果不变。
- `relation-syncer` 的 `FollowerModel.UpsertActive` 走 `INSERT ... ON DUPLICATE KEY UPDATE rel_status=1`（唯一索引 `uk_to_from` 保证同一对用户关系只有一行),重复执行只是把 `rel_status` 反复置 1,不会插入第二行。
- 反例(需要额外小心的地方):`FollowHandler.HandleCreated` 里的 `UserCounterRpc.UserIncrement(Delta:+1)` 是一次**增量**操作,不是覆盖写——如果去重失效导致 `HandleCreated` 被完整重跑一次,双方的关注数/粉丝数会被多加一次,这一步**依赖第一层 SETNX 去重完全生效**,自身没有二次兜底。这是当前设计里"业务幂等"覆盖不到的唯一缺口,好在这类计数偏差最终会被 `counter-reconciler`(如果 usercounter 也接入了同类对账机制)或人工核对发现。

**第三层:Canal 分区哈希保证同一聚合根的事件顺序不乱**

`deploy/compose/canal-conf/canal.properties`(`canal.mq.partitionHash = zhiguang.outbox:id`)配置了 Canal 按 `outbox.id` 做分区哈希投递到 Kafka——这保证的不是幂等,而是**同一行 outbox 记录的产生顺序**在 Kafka 分区内是有序的,避免"删除事件先于更新事件被消费"这类乱序问题。当前 `canal.mq.partitionsNum=1`(单分区),配置文件里特别注释了"必须与 canal-outbox 在 Kafka 里的实际分区数一致"——`docker-compose.dev.yml` 里 Kafka 开了 `AUTO_CREATE_TOPICS_ENABLE=true`,topic 会以默认分区数(通常是 1)自动创建,这就是为什么 CLAUDE.md 里记录过"分区数不一致导致 canal-server 卡死在 metadata 请求循环"的历史 bug——分区数必须两边手动对齐,不能依赖自动创建的默认值。

### 结论

outbox 幂等的核心不是数据库层面的设计(outbox 表本身没有"已消费"标记,也不需要),而是消费端用 `SETNX` 卡住重复消息,业务写路径尽量设计成基于唯一键的覆盖写或 upsert,把真正"不可幂等"的增量操作(如计数 +1/-1)完全依赖第一层去重兜底。

---

## 2. Redis 挂了,如何用主从方案重建

### 需要先明确:当前代码/配置里 Redis 是什么形态

`pkg/redisx/client.go` 的 `Config.Type` 只支持 `node`(单机/由外部保证高可用的地址)和 `cluster`(Redis Cluster)两种,`docker-compose.dev.yml` 明确注释"Redis(6379)由宿主机本地进程提供,不在此 compose 中启动"——也就是说**当前项目里 Redis 本身的高可用形态没有在代码/编排层面落地**,`redisx.New` 只是把配置里的地址包给 `go-redis` 的 `UniversalClient`,主从切换、哨兵(Sentinel)选主这些工作被假定由外部的 Redis 部署方案(比如云厂商托管实例,或者自建的 Sentinel/Cluster)负责,应用代码不感知,只需要保证连接的地址在故障切换后依然可用即可(`go-redis` 的 `UniversalClient` 本身支持 Sentinel 地址格式,但 `redisx.Config` 目前没有暴露 Sentinel 专属的 `MasterName` 等参数,如果要接入 Sentinel 需要扩展这个结构)。

所以这个问题要拆成两层回答:**Redis 服务本身怎么做主从故障切换**（运维/中间件层的事)和**切换后/数据丢失后,应用侧的缓存内容要怎么重建**（这才是代码可以回答的部分,也是更关键的部分)。

### 2.1 主从故障切换(基础设施层,现状缺失)

标准方案是 Redis Sentinel 或 Redis Cluster:

- **Sentinel 方案**:1 主 N 从 + 至少 3 个 Sentinel 节点监控。主节点心跳超时后,Sentinel 集群投票选出新主,客户端(`go-redis` 支持 `NewFailoverClient`)通过 Sentinel 感知新主地址并自动重连,业务代码不需要修改连接字符串。切换期间(通常几秒到十几秒)主库不可写,读请求如果配置了从库读可以不受影响。
- **Cluster 方案**:数据按 slot 分片到多个主节点,每个主节点带从节点。单个分片主节点故障后集群内自动 failover 到该分片的从节点,只影响该分片对应的 slot,不是全局不可用。`redisx.Config.Type=cluster` 已经支持这种形态,但需要注意前面提到的:代码里用到 `SCAN`(如 reconciler、cache_invalidation_helper 的按前缀清理)在 Cluster 模式下只能扫本地分片,如果 key 分布在多个分片,需要对每个节点分别 `SCAN`——`go-redis` 的 `UniversalClient` 在 cluster 模式下的 `Scan` 默认只覆盖一个节点,这是接入 Cluster 前需要额外验证的兼容点。

要把这一层补起来,需要:引入 Sentinel/Cluster 部署 + 扩展 `redisx.Config` 支持 Sentinel 的 `MasterName`/多地址,以及验证所有 `SCAN` 类调用在目标拓扑下的行为。

### 2.2 应用侧缓存重建(代码可以回答的部分)

Redis 故障切换或数据丢失后,项目里存的东西按"能否重建"分成三类,重建路径完全不同:

**A 类:可以从 MySQL 重新推导的缓存(可完全重建,无数据损失)**

- `know_posts` 的详情缓存(`cache:knowPosts:id:*`,go-zero `CachedConn` 管理)、`DetailCache`/`L1FeedItem`/Feed 分页缓存——这些都是 MySQL `know_posts` 表数据的只读投影,Redis 数据丢失后,下一次读请求会直接 miss 缓存回源 MySQL,`cachex`(三层缓存:L1 ristretto + L2 Redis + loader)本身就是为"缓存未命中"设计的,不需要人工干预,回源过程自然重建缓存,只是短时间内 MySQL 负载会因为缓存穿透而上升(冷启动效应)。
- `following`/`follower` 的双向 ZSet 缓存(`uf:flws:{userId}`/`uf:fans:{userId}`)——本质是 `following`/`follower` 两张 MySQL 表里 `rel_status=1` 的行的一份物化视图。**重建流程**:写一个一次性脚本(或复用 `relation-syncer` 里已有的 `zsetAdd` 逻辑),按 `from_user_id`(或 `to_user_id`)分组扫全表 `SELECT from_user_id, to_user_id, created_at FROM following WHERE rel_status=1`,对每个用户批量 `ZAdd` 重建对应 key。由于 ZSet 只保留最近 `MaxMembers` 条(截断策略),重建时按 `created_at DESC` 取前 N 条即可,不需要全量导入。这类缓存丢失不影响正确性,只是丢失期间的关注列表查询会退化(如果没有直接查库的兜底路径,需要临时加一个)。

**B 类:有 MySQL 落地但"事件已经过去"、只能补数据不能补时序的缓存**

- 无(在当前架构下,following/know_posts 相关缓存都是纯投影,不存在这种中间态)。

**C 类:完全没有 MySQL 兜底,Redis 是唯一事实来源(不可逆,重建=清零)**

- **counter 域的位图(`bm:*`)和 SDS(`cnt:v1:*`)**——这是本项目里唯一的高风险点。`counter-reconciler` 的整套设计前提是"位图永远正确",因为位图本身就是事实来源,没有更上游的数据库记录每一次点赞的历史。如果 Redis 因为故障丢失了这部分数据(比如主从切换时从库数据落后、或者两者都没做持久化就重启),**这些点赞/收藏计数是无法重建的**,只能清零重新开始积累,或者如果有能力从 ES/日志等旁路数据里反推近似值做人工修正,但代码里没有这样的机制。
- **refresh token(`services/user/internal/adapter/token` 存 Redis)**——丢失后受影响用户的 refresh token 失效,需要重新登录,不算数据损失,只是体验影响。
- **验证码(`services/user/internal/adapter/verification`)**——丢失后进行中的验证码流程失效,用户重新发一次验证码即可,天然是短 TTL 数据,影响窗口很小。

### 2.3 针对本项目的具体重建流程建议

如果 Redis 整体不可用后重新拉起一个空实例,推荐的恢复顺序:

1. **先接入新 Redis,不做任何数据预热**——让服务直接跑,依赖 A 类缓存的自然回源能力,观察 MySQL 负载(如果预估回源流量过大,可以临时调高 MySQL 连接池或限流保护)。
2. **批量重建 following/follower 的 ZSet 缓存**——写一次性脚本按上文 A 类描述的逻辑跑一遍,这是少数值得主动重建而不是被动等待回源的缓存(因为关注列表的查询模式是"点开某人主页立即要看关注列表",被动回源体验上会有一次明显的冷启动延迟,主动预热成本很低)。
3. **确认 counter 位图/SDS 数据丢失范围**——如果只是部分 key 因为过期/淘汰丢失(不是整实例数据全丢),`counter-reconciler` 会在下一轮扫描时发现 SDS 与位图的不一致(此时位图也可能是 0,SDS 也可能是 0,如果两者都归零反而"看起来一致",这是 reconciler 检测不出来的场景——它只能发现"两者不一致",发现不了"两者都错但恰好相等")。**这种情况下无法通过 reconciler 自动恢复,只能接受数据丢失或人工从其它旁路(如日志、ES 里记录的历史 like_count 快照)估算后手动写回**。
4. **重新触发一次 Elasticsearch 全量 reindex**(如果 ES 索引也受影响,或者为了保证与刚重建的 counter 数据一致)——当前 `Reindex` RPC 是 stub,需要靠重新消费 Kafka 历史消息或写一次性脚本遍历 `know_posts` 表调用 `upsert` 逻辑。

---

## 3. Kafka 挂了怎么办,依靠多实例是否足够

"依靠多实例"是正确方向,但要分清楚"多实例"具体指的是哪一层,以及当前配置离这个目标还差多远。

### 3.1 现状:dev 环境是单 broker,没有复制

`docker-compose.dev.yml` 里只起了一个 `kafka` 容器,没有配置 `KAFKA_CFG_DEFAULT_REPLICATION_FACTOR` 或多 broker,也没有 `min.insync.replicas` 配置——这是**开发环境**的配置,不代表生产可用的高可用形态。生产环境如果只用这套配置原样部署,Kafka 单点故障时全部依赖它的写入(outbox 事件分发、counter-events)会中断,这是当前配置的明确缺口。

### 3.2 生产环境"多实例"具体要做的三件事

**第一,broker 多副本(这是"多实例"的核心,你的方向是对的)**

至少 3 个 broker 节点组成集群,每个 topic 设置 `replication.factor=3`,配合 `min.insync.replicas=2`(至少 2 个副本确认写入才算成功)。这样单个 broker 宕机,分区的 leader 会自动切换到其它副本上的节点,不丢数据、生产消费基本不中断(除了 leader 切换瞬间的短暂重试)。这一层完全是 Kafka 集群自身的能力,应用代码不需要改动,但**当前的分区数配置需要重新考虑**:目前 `canal-outbox` topic 是 `partitionsNum=1`(单分区),多 broker 部署下单分区意味着所有写入还是落在一个 leader 上,没有利用到多 broker 的吞吐能力,只是获得了故障切换的可靠性。如果要提升吞吐还需要增加分区数——但增加分区数会打破当前"按 `outbox.id` 哈希保证同聚合根事件有序"的假设吗?**不会**,因为 `partitionHash = zhiguang.outbox:id` 是按 outbox 行自身的 id 哈希,而不是按聚合根 id,如果要保证同一个 `aggregate_id`(比如同一篇知文)的多次更新事件严格按顺序被同一个 consumer 处理,分区键需要改成按 `aggregate_id` 哈希,而不是当前的按 outbox 自增 id 哈希——这是当前实现里一个容易被忽略的点:多分区场景下,同一篇知文的 `KnowPostUpdated` 和后续的 `KnowPostPublished` 理论上可能被分到不同分区,不再保证消费顺序。**这是在真正上多分区之前必须先解决的问题**,否则会引入乱序处理的新 bug。

**第二,Producer 侧的可靠性配置(代码已经做对了这部分)**

`pkg/kafkax/producer.go` 的 `RequiredAcks: kafka.RequireAll`(等价于 Kafka 的 `acks=all`,要求所有同步副本都确认才算发送成功)——这个配置本身是正确的,配合上面的多副本 broker,能保证消息一旦发送成功,即使 leader 立刻宕机,数据也在其它副本上不丢。目前的风险点是:`Publish` 失败时的处理策略在不同调用点不一致——`counter` 的 `Toggle` RPC 对 Kafka 发送失败**只记日志不重试**(位图已经落地,不能因为消息队列问题回滚用户的点赞动作),这意味着如果 Kafka 集群短暂不可用(即使有多副本,集群整体重启或网络分区期间仍可能短暂拒绝写入),这段时间产生的点赞事件会永久丢失,只能靠 `counter-reconciler` 事后从位图重新算出正确的 SDS 值来补偿——这也是为什么"位图是事实来源"这个设计在 Kafka 不可靠的场景下显得更重要:**counter 这条链路把 Kafka 当作"尽力而为的异步通知",不当作可靠传输**,所以 Kafka 短暂故障不会丢失最终一致的数据,只会延迟 SDS 更新。而 outbox→canal→Kafka 这条链路本身的可靠性不依赖 Producer 重试,因为写 outbox 行是在 MySQL 事务里完成的,Canal 是持续 tail binlog(不是消费一次性的消息),即使 Kafka 短暂不可用,Canal 会在 Kafka 恢复后继续从上次的 binlog 位点转发,不会漏发 outbox 事件(binlog 是持久化的、可重放的,这是 outbox+CDC 模式相对于"业务代码直接发 Kafka"更可靠的核心原因)。

**第三,Consumer 侧的重试与死信队列(代码已经做了,是消费侧对"投递可能失败/处理可能失败"的兜底)**

`kafkax.RunConsumer` 的手动 offset 提交 + 指数退避重试(`backoffBase=100ms` 到 `backoffMax=30s`)+ 超过 `MaxRetries` 后写入死信队列(`DlqTopic`)——这套机制处理的是"consumer 处理消息失败"(比如下游 RPC 超时、ES 写入失败),跟"Kafka broker 本身挂了"是两个不同的故障维度:如果是 broker 集群不可用,`FetchMessage`/`CommitMessages` 本身会返回网络错误,`RunConsumer` 目前对这类错误的处理是**直接返回并结束整个消费循环**(`consumer.go:74-79`,非 `context.Canceled`/`io.EOF` 的错误会导致函数返回,不会自动重连)——这意味着如果 Kafka 集群短暂抖动导致连接中断,当前的 consumer 进程会退出,需要依赖外部的进程管理(如 k8s 的 Pod 重启策略、systemd 的 Restart=always)重新拉起进程才能恢复消费,而不是在 `RunConsumer` 内部自动重连重试。**这是当前实现的一个可以改进的点**:生产环境如果 Kafka 集群做了多 broker 高可用,更希望 consumer 在 broker 级别的短暂故障时能自己重试重连,而不是整个进程退出等外部重启,可以考虑在 `RunConsumer` 外包一层重连循环,或者依赖 k8s 层面的自动重启作为等效方案(如果部署在 k8s 上,这条也可以不改代码,直接靠 Pod 重启策略兜底,只是恢复速度取决于重启探针的检测间隔)。

### 3.3 结论

"依靠多实例"的思路完全正确,但完整的 Kafka 高可用要包含:broker 多副本(集群层,当前 dev 配置缺失,生产必须补)+ 恰当的分区数与分区键选择(当前单分区且按 outbox 自增 id 哈希,扩分区前需要先把分区键换成按聚合根 id 哈希以保序)+ Producer 端 `acks=all`(代码已具备)+ Consumer 端对处理失败的重试/死信(代码已具备)+ Consumer 端对连接层故障的自动重连(当前依赖外部进程管理兜底,可以考虑内建重连作为增强)。另外要接受一个现实:即使 Kafka 100% 高可用,业务代码里仍有主动选择"不可靠传输"的地方(counter 的 Toggle 事件发送失败只记日志),这不是 Kafka 的问题,是这条链路的设计取舍——用位图事实层+定期对账换取"点赞响应不被消息队列拖慢或阻塞"的性能收益,理解这一点后,Kafka 高可用建设的优先级应该放在 outbox→索引/关系同步这类**没有独立事实来源兜底**的链路上(search-indexer、relation-syncer),这些链路一旦丢消息,除非重新做一次全量 reindex/repair,数据是真的会长期不一致,不像 counter 有 reconciler 自动收敛。
