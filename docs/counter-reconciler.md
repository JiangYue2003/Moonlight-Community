# Counter 对账机制（Reconciler）技术文档

## 1. 模块概述

`services/counter/reconciler`（计数对账 worker）是点赞/收藏计数体系里专门负责"纠错"的一个独立后台进程,不对外提供任何 RPC/HTTP 接口,只有一个内部循环:定时扫描 Redis 里所有的计数缓存 key,拿它们的值跟一份被认定为"绝对事实"的数据源做比对,一旦偏差超过阈值就强制覆盖。它存在的前提是:整个点赞/收藏链路(见 `services/counter/rpc`、`services/counter/aggregator`)完全建立在 Redis 之上,没有 MySQL 兜底,而这条链路里有一段是经过 Kafka 异步传递、允许消息丢失或乱序的,所以需要一个独立机制把因为丢消息造成的计数漂移找出来并修正,而不是假设这条链路永远可靠。

要理解为什么需要对账,得先理解点赞/收藏计数在 Redis 里被拆成了三层结构,它们各自的可靠性不同:

- **位图层(Bitmap,事实层)**:`bm:{metric}:{etype}:{eid}:{chunk}`（用户行为位图,一个用户对一个实体的点赞状态用一个 bit 表示）。写入路径是 `Toggle` RPC 直接用 Lua 脚本 `SETBIT`,同步执行、不经过 Kafka,天然幂等(同一个 bit 反复置 1 不会产生累计误差),理论上永远准确,是整套体系里唯一可以拿来做"真值"校验源的东西。
- **聚合桶层(Aggregation Bucket,过渡层)**:`agg:v1:{etype}:{eid}`（Redis Hash,按 delta 累加的临时聚合缓冲区）。写入路径是 `counter-aggregator` 消费 Kafka `counter-events` 主题后 `HIncrBy` 累加,这一步依赖 Kafka 投递成功,如果消息丢失,这个桶里就会少记一次 +1/-1。
- **SDS 缓存层(对外读取层)**:`cnt:v1:{etype}:{eid}`（紧凑二进制编码的计数结构,`GetCounts`/搜索结果里显示的点赞数直接来自这里）。写入路径是 `counter-aggregator` 定时把聚合桶里的 delta 折算进来,如果聚合桶本身已经漂移,这一层自然继承同样的误差。

对账要解决的就是:**位图层永远准,SDS 层可能因为中间的 Kafka/聚合环节出问题而不准**,reconciler 用位图重新算出真实值,写回 SDS,把误差归零。

## 2. 核心抽象与实现要素

- **`Worker`**（对账主循环的载体,`services/counter/reconciler/internal/worker/worker.go`）:内部持有一个 Redis 客户端和配置,`Run` 方法用 `time.Ticker` 按固定小时间隔触发一轮全量对账,不是持续扫描,是周期性批处理任务。
- **`reconcile`**（一轮完整对账的入口）:用 `SCAN cnt:v1:*`（游标式扫描,避免 `KEYS` 阻塞 Redis）遍历所有当前存在计数的实体,对每个 key 调 `reconcileKey`,批次之间可插入 `BatchIntervalMs`（每批扫描后的休眠间隔）限速,避免对账任务本身对 Redis 产生压力尖峰。
- **`reconcileKey`**（单个实体的对账核心逻辑)：从 SDS key 里解析出 `etype`/`eid`（实体类型和实体 ID,如 `knowpost`/`331127866555371520`),对 `schema.Supported`（当前支持对账的指标列表,即 `like`/`fav`)逐个指标执行"计算偏差→判断是否需要重建"。
- **`scanBitmapTotal`**（真值计算)：用 `SCAN bm:{metric}:{etype}:{eid}:*` 找出该实体下所有分片位图(用户量大时一个实体的点赞位图会按 `ChunkSize=32768` 分片存储),Redis Pipeline 批量执行 `BITCOUNT`（统计位图里置 1 的 bit 数,即真实点赞人数)后求和,得到这个指标的绝对真值。
- **`diff`/`pct`**（偏差判定的两个维度)：`diff` 是 SDS 值与位图真值的绝对差,`pct` 是这个差值相对 SDS 当前值的百分比。两者任一超过配置阈值(`ThresholdAbsolute`/`ThresholdPercent`)就判定为需要重建——绝对值阈值兜底小基数场景(比如真实只差 3 个赞但占比很高的冷门内容),百分比阈值兜底大基数场景(热门内容差几十个赞占比很小但绝对数量已经不可忽视)。
- **`rebuildField`**（修复动作)：直接用位图算出的真值以大端(BigEndian)方式覆写 SDS 结构里对应 `metric` 的那 4 个字节(`FieldSize=4`),是整字段的**覆盖替换**而不是增量修正——因为一旦确认 SDS 不可信,就不再信任它的任何历史累加过程,直接推正确答案。
- **`ScanConf`**（对账节奏与灵敏度配置,`services/counter/reconciler/internal/config/config.go`)：`IntervalHours`（对账轮询间隔,默认1小时)、`BatchSize`/`BatchIntervalMs`（scan 分批大小与批间限速)、`ThresholdAbsolute`/`ThresholdPercent`（触发重建的两个阈值,默认绝对差 100、百分比 1%)。

## 3. 什么情况下需要对账

对账机制针对的是**SDS 缓存层与位图事实层之间产生偏差**的场景,具体拆解会在这条链路的哪些环节产生偏差:

- **Kafka 消息丢失**:`Toggle` RPC 在位图翻转成功后异步发一条 `CounterEvent` 到 Kafka `counter-events` 主题,如果这次 `Publish` 调用失败(代码里`outbox_helper`及`togglelogic.go`对 Kafka 发送失败的处理策略是**只记日志、不回滚位图操作**——位图已经落地的状态是权威的,不能因为消息队列的问题去撤销一次已经成功的用户点赞动作),这条事件就永久丢失,对应的 +1/-1 永远不会被 `counter-aggregator` 消费到,SDS 里就会少算或多算这一次。
- **Aggregator 消费/落盘过程中断**:`counter-aggregator` 消费 Kafka 后先 `HIncrBy` 进聚合桶,再由 `RunFlusher` 定时把聚合桶的 delta 折算进 SDS。这中间涉及两次写(累加进聚合桶、从聚合桶折算进 SDS 再扣减)。如果进程在两次写之间崩溃、或者 Redis 主从切换导致某次写没有持久化成功,聚合桶或 SDS 都可能停留在中间状态,产生偏差。
- **多副本环境下 flush 逻辑异常**:`RunFlusher` 靠 `Locks.TryAcquire`（分布式锁选主)保证多个 aggregator 副本不会同时 flush 同一个聚合桶,如果锁续期/释放逻辑出现边界问题(比如锁提前过期导致两个副本同时执行了同一批 flush),理论上可能造成 delta 被重复折算或遗漏折算。
- **人工干预或历史数据迁移**:如果运维人员曾经手动修改过 Redis 里的 SDS 值(调试、清脏数据等),或者从旧系统(代码注释多处提到"与 Java 版对齐"，暗示这是从存量 Java 系统迁移而来)导入过历史计数,这些操作都不经过位图,自然会导致两层不一致。

判断"要不要对账"不依赖具体某次故障是否发生,而是**周期性无条件全量扫描**——reconciler 假设上述任何一种偏差都可能随时无声发生且无法被上游主动感知到(Kafka 发送失败只记日志,没有告警或重试机制去补发),所以采用"定期体检"而不是"出问题了再修"的策略,`IntervalHours` 决定了一旦发生偏差,最长要等多久才会被自动纠正。

## 4. 怎么对账：流程剖析

对账分成"计算偏差"和"判定重建"两个阶段,每个 SDS key、每个 metric 都独立走一遍这套流程。

```mermaid
flowchart TB
    A[Ticker 到达 IntervalHours] --> B[SCAN cnt:v1:* 遍历所有实体]
    B --> C[取一个SDS key]
    C --> D[GET该key原始字节并sds.Decode解析成5个字段]
    D --> E{遍历Supported指标<br/>like/fav}
    E --> F[SCAN bm:metric:etype:eid:* 找出所有分片位图]
    F --> G[Pipeline BITCOUNT 批量统计各分片置1位数]
    G --> H[求和得到bitmapTotal 真实值]
    H --> I[diff = abs(sdsVal - bitmapTotal)]
    I --> J[pct = diff / max(sdsVal,1) * 100]
    J --> K{diff > ThresholdAbsolute<br/>或 pct > ThresholdPercent?}
    K -->|否| L[记录日志，跳过]
    K -->|是| M[rebuildField: 用bitmapTotal覆写SDS对应字段]
    M --> N[记录日志: rebuild=true]
    L --> O{还有下一个指标?}
    N --> O
    O -->|是| E
    O -->|否| P{还有下一个key?}
    P -->|是| C
    P -->|否| Q[本轮对账结束，等待下一次Ticker]
```

具体到单个 `reconcileKey` 调用内部的判定逻辑(对应 `worker.go:77-120`):

1. 从 key 名 `cnt:v1:{etype}:{eid}` 用 `SplitN` 拆出 `etype`/`eid`,`GET` 取出原始字节,`sds.Decode` 解出一个长度为 `SchemaLen=5` 的整型数组(对应预留的 read/like/fav/comment/repost 五个字段,当前只有 `like`/`fav` 被启用)。
2. 对每个受支持的 metric,先用 `schema.IdxOf` 定位它在 SDS 数组里的下标,再调 `scanBitmapTotal` 算出该指标的位图真值——这一步本身也是一次子扫描(`SCAN bm:{metric}:{etype}:{eid}:*` + Pipeline `BITCOUNT`),因为一个高热度实体的用户行为位图可能被分成多个 `ChunkSize=32768` 位一片的分片存储,必须把所有分片的置位数累加才是完整真值。
3. 计算绝对偏差 `diff` 和相对偏差 `pct`,两个阈值任一触发就认为"这个字段不可信了",调用 `rebuildField` 直接用真值覆盖对应的 4 字节(`FieldSize=4`,大端编码,与 `IncrFieldScript`/`DecrFieldScript` 写入时使用的编码格式一致,保证覆盖后依然能被日常读写逻辑正确解析)。
4. 无论是否触发重建,都会打一条 `logx.Infof` 日志,把 `sds`/`bitmap`/`diff`/`pct`/`rebuild` 全部打出来——这既是对账结果的审计记录,也是排查"为什么某个实体计数总是不对"时的第一手线索来源。

## 5. 异常处理与边界设计

- **对账过程中的读失败被隔离到单个 key,不中断整轮扫描**:`reconcile` 里对每个 key 调 `reconcileKey` 出错时只 `logx.Errorf` 记日志,`continue` 处理下一个 key,不会因为某一个实体的位图扫描失败(比如网络抖动)拖垮整轮对账。但反过来说,`scanBitmapTotal`/`w.rdb.Scan` 内部一旦返回 error,`reconcileKey` 会直接 return,当次这个 key 的其余指标也不会被检查,要等下一轮 `IntervalHours` 后才会重试——这是一个"最终会自我修复但当次静默失败"的设计取舍。
- **SDS key 不存在时直接跳过**:`w.rdb.Get` 返回 `goredis.Nil`(key 不存在)时 `reconcileKey` 直接返回 nil,不当作错误处理——这对应"从未被点赞过"或"缓存已过期被淘汰"的正常情况,不需要重建。
- **限速设计防止对账任务本身冲击 Redis**:`BatchSize`(每次 SCAN 返回的 key 数量提示值)和 `BatchIntervalMs`(每批之间的休眠)组合限制了对账任务的 QPS 上限,这是因为一次全量对账要对每个实体、每个 metric 都做一次子扫描+Pipeline,如果没有限速,在实体数量很大时会对 Redis 产生显著的额外负载,与线上读写流量抢资源。
- **覆盖式修复而非增量修复,且只信任位图**:`rebuildField` 采用整字段覆盖而不是"计算出差值再增量修正"——因为一旦 SDS 和位图出现偏差,SDS 自身的历史累加路径已经不可信,唯一可信的是位图这一份事实来源,所以选择完全推倒重来而不是在不可信的基础上再做一次增量运算(增量运算如果算错方向,反而会让偏差扩大而不是收敛)。
- **不处理位图层本身的数据丢失**:整套对账机制的前提假设是"位图永远准确",如果位图这一层因为 Redis 数据丢失(比如未持久化就重启、或者主从切换丢数据)而本身不准,reconciler 是没有能力发现或修复的——因为它没有更上游的真值源(比如 MySQL)可以比对。这是当前架构里唯一没有兜底的一层风险,如果未来出现"点赞数全局性异常归零"这类问题,基本可以判断是位图层数据丢失,而不是 SDS 偏差,reconciler 对此无能为力。
- **只对账已支持的 metric,预留字段被忽略**:`schema.Supported` 目前只包含 `like`/`fav`,`schema.go` 里预留的 `IdxRead`/`IdxComment`/`IdxRepost` 三个字段不会被对账检查——如果未来启用这些指标但忘记加入 `Supported` 列表,即使产生了偏差,reconciler 也不会发现,这是新增计数指标时需要注意补齐的一处配置点。
