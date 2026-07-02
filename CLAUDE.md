# TODO

## 1. [已修复] knowpost 增量更新覆盖丢数据

**现象**：`PatchMetadata` 设置 title/description/tags 后，若紧接着调用 `ConfirmContent`（或 `UpdateTop`/`UpdateVisibility`/`Publish` 等其它整行覆盖写接口），刚 patch 的字段会被冲回空值。

**根因**：
- `PatchMetadata` 走 `KnowPostsModel.UpdateInTx`（外部事务里的原生 `session.ExecCtx`），这条路径绕过了 go-zero `CachedConn`，不会清除模型级缓存 key（`cache:knowPosts:id:{id}`）。
- `invalidateKnowPostCaches`（`services/knowpost/rpc/internal/logic/knowpost/cache_invalidation_helper.go`）只清应用层的详情缓存/Feed缓存，清不到这个模型缓存 key。
- 后续任何调用 `findOwnedRow` → `FindOne` 的接口（`confirmcontentlogic.go`、`updatetoplogic.go`、`updatevisibilitylogic.go` 等）会读到缓存里的旧行，只改自己关心的字段后，用 `model.Update()`/`UpdateInTx` 整行覆盖写回数据库，把其它字段的新值冲掉。

**复现步骤**：
1. `POST /knowposts/drafts` 建草稿
2. `PATCH /knowposts/{id}` 设置 title/description/tags
3. `POST /knowposts/{id}/content/confirm` 确认内容
4. 查库：title/description/tags 变回 `null`

**涉及文件**：
- `services/knowpost/rpc/internal/logic/knowpost/patchmetadatalogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/confirmcontentlogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/outbox_helper.go`
- `services/knowpost/rpc/internal/logic/knowpost/publishlogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/deletelogic.go`
- `services/knowpost/shared/model/knowpostsmodel.go`（`UpdateInTx`、新增 `InvalidateCache`）

**修复方案（已落地）**：
- `KnowPostsModel` 接口新增 `InvalidateCache(ctx, id) error`，内部调用 `CachedConn.DelCacheCtx` 清模型缓存 key。
- 所有走 `UpdateInTx` 的写路径（`outbox_helper.go` 的 `updateAndEmitOutbox`，覆盖 patch/top/visibility 三处调用；`publishlogic.go`；`deletelogic.go`）在事务提交成功后立即调用 `InvalidateCache`，不放进事务回调内部（避免回滚场景下缓存被误删）。
- `services/knowpost/rpc/internal/logic/knowpost/bench_helpers_test.go`、`getdetaillogic_test.go` 的 mock 补了 `InvalidateCache` 方法。
- 已跑 `go build ./...`、`go vet ./services/knowpost/...`、`go test ./services/knowpost/...`（knowpost 相关包全绿，唯一失败的 `TestDecodeKnowpostMergedConfig` 是无关的预置问题），并用 curl 端到端复现原 bug 流程验证：patch → confirm content → publish，title/description/tags 全程保留，数据库确认无误。

## 2. [已修复] canal-server 未将 binlog 事件转发到 Kafka

**现象**：`db/migrations` 建的 `canal` 账号权限正常，canal-server 能连上 MySQL 并建立 `Binlog Dump` 连接（`SHOW PROCESSLIST` 可见），MySQL 侧 `outbox` 表的 INSERT 也确实发生了，但 Kafka `canal-outbox` topic 里始终没有消息，最终报 `TimeoutException: Topic canal-outbox not present in metadata after 60000 ms`。

**排查过程中走过的弯路**（均已证明与本问题无关，记录下来避免下次重复排查）：
- 怀疑插件 SPI 加载路径问题（`/plugin`、`/canal/plugin` 绝对路径），建了软链接验证，重启后问题依旧。
- 怀疑 `canal.instance.global.mode = manager` 导致本地 instance.properties 完全不加载 —— 这个问题真实存在（改成 `spring` 后 `example` 实例才真正启动），但只是让 canal 从"完全不跑"变成"跑起来但 MQ 发不出去"，不是本次超时报错的根因。
- 怀疑 kafka-clients 2.4.0（canal 打包版本）与 Kafka 3.7.0 broker 协议不兼容 —— 开 DEBUG 日志后证明 ApiVersions 协商、metadata 请求/响应完全正常，producer 确实拿到了 `canal-outbox` 的 metadata（`leader=1001, replicas=[1001], isr=[1001]`），排除协议兼容性问题。

**真正根因**：`canal.properties` 与 `instance.properties` 里都配了 `canal.mq.partitionsNum = 6`（配合 `canal.mq.partitionHash = zhiguang.outbox:id` 按 id 哈希分派到 0-5 号分区），但 Kafka 里 `canal-outbox` 这个 topic 是靠 `auto.create.topics.enable` 自动创建的，默认只有 **1 个分区**。`CanalKafkaProducer.send()` 用哈希算出的分区号直接构造 `ProducerRecord`，一旦这个分区号（0 以外的值）在实际 topic 里不存在，`KafkaProducer.waitOnMetadata` 就会陷入死循环：每 ~100ms 反复重新请求 metadata，等待这个分区出现，但它永远不会出现，最终超时报出这个看似"连不上 Kafka"的错误（DEBUG 日志证实：60 秒内发了 2260 次完全相同的 metadata 请求）。

**定位方法**：canal-server 顶层的 `logs/canal/canal.log` 只有生命周期日志，真正的异常堆栈在 `logs/example/example.log`（`example` 是 destination 名）。把 `logback.xml` 里 `com.alibaba.otter.canal.connector.kafka` 和新增的 `org.apache.kafka` logger 临时调到 `DEBUG`，能看到完整的 metadata 请求/响应循环。

**涉及文件**：
- `deploy/compose/canal-conf/canal.properties`（`canal.mq.partitionsNum`）
- `deploy/compose/canal-conf/instance.properties`（`canal.instance.global.mode`、`canal.instance.master.address`、`canal.instance.tsdb.enable`、`canal.mq.partitionsNum`）

**修复方案（已落地）**：
- `canal.mq.partitionsNum` 从 `6` 改为 `1`，匹配 `canal-outbox` 实际分区数（两个文件都改，保持一致）。
- 顺手修正的另外三项配置（发现于排查过程，都是必要前提，不改的话根本走不到分区数这一步）：
  - `canal.instance.global.mode`：`manager` → `spring`（否则 `instance.properties` 完全不生效，`example` destination 不会启动）。
  - `canal.instance.master.address`：`mysql:3306` → `host.docker.internal:3306`（compose 里没有 `mysql` 这个服务，MySQL 跑在宿主机）。
  - `canal.instance.tsdb.enable`：`true` → `false`（没配 tsdb 的 h2 jdbc 参数，开着会在启动时刷一堆 `DruidDataSource` 空指针报错，虽不影响主流程但很吵）。
- 验证：重启 canal 后，之前积压的 8 条真实 outbox 事件（`KnowPostUpdated`/`KnowPostPublished`/`FollowCreated`）全部成功发到 `canal-outbox`；`relation-syncer` 消费后 `followers` 计数从 0 变成 1；`search-indexer` 消费后通过 `reindex` 接口重建索引，`gateway` 搜索接口能查到发布的知文并高亮命中。
- 顺带清理：过程中发现 Redis 里有两条帖子（`331127866555371520`、`331129668248014848`）残留着问题1修复之前的脏缓存（`knowpost:detail:*:v1`、`cache:knowPosts:id:*`），已手动 `redis-cli del` 清除 —— 这是测试数据污染，不是新 bug。
