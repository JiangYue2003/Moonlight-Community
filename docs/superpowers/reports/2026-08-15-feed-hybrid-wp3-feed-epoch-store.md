# Feed Hybrid WP3：FeedEpochStore 检查点

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）

## 1. 结论

WP3 已建立 Redis 权威、短时本地缓存的 Feed 失效代数，为后续 RouteSnapshot 和 PageCache 提供版本 key：

```text
feed:relation:epoch:{userID}
feed:content:safety:epoch
```

Redis key 使用 String + `INCR`，不设置 TTL；不存在的 key 读取为 0。Redis 读取或递增失败会返回错误，不会伪造 epoch。

## 2. 实现

新增 `services/knowpost/rpc/internal/feedepoch`：

- `Store` 提供 Relation/Safety 读、Bump、Safety pending 标记和补偿刷新。
- Relation 与 Safety L1 TTL 默认均为 1 秒。
- L1 使用独立 Ristretto：`NumCounters=200000`、`MaxCost=4MB`，不与 FeedMine、RouteSnapshot 或 PageCache 共用。
- `KeyPrefix` 可改为带 run_id 的测试前缀，支持定向清理。
- 非法 `userID <= 0` 在访问 Redis 前返回参数错误。

Relation 并发协调：

- L1 热命中无锁返回。
- 256 个分片 mutex 只排序同分片的 cold fill/Bump，不再使用全 Store 锁。
- same-key singleflight 合并 L1 过期瞬间的 Redis GET。
- 共享查询不继承首个 waiter 的取消信号，但由 2 秒内部超时限制；各 waiter 仍响应自己的 Context。
- 同 key cold fill 与 Bump 使用相同分片锁，旧 Redis 结果不能在 Bump 后覆盖新 epoch。

Safety pending：

- Bump 失败递增 pending generation。
- 多次失败可由一次后续成功 Bump 清除，因为任意一个新 epoch 都能废弃全部旧页面。
- 成功 Bump/Flush 只 CAS 清除操作开始前观察到的 generation；操作期间新增的 `MarkSafetyPending` 不会被误清除。
- 成功写 Redis 后等待低频 Ristretto admission，使当前实例在 Bump 返回后立即可见。

WP3 只完成 Store、配置和 ServiceContext 注入，没有连接 Kafka consumer、RouteSnapshot 或页面缓存。

## 3. 测试证据

覆盖：

- Relation/Safety miss 返回 0，L1 到期后重新读取 Redis。
- 64 个并发 Relation INCR 全部保留，Bump 后立即可见。
- 同 key 32 个并发 cold miss 只触发一次 Redis GET。
- 用户 A 阻塞 GET/INCR 均不阻塞用户 B 的 INCR/GET。
- Redis GET 已捕获旧值时，同 key Bump 必须等待；最终旧 fill 不能覆盖新 epoch。
- Relation/Safety Redis key 均保持 persistent，无 TTL。
- Safety 三次失败合并为一次补偿 Bump；补偿完成前 pending 保持 true。
- 阻塞中的成功 Bump 不会清除更晚发生的 pending generation。
- Redis GET/INCR 故障不返回伪造版本。
- 配置解码确认独立 Epoch L1 预算和 1 秒 TTL。

并发测试连续运行 20 次通过。

## 4. 验证

已通过：

```text
go test ./services/knowpost/rpc/internal/feedepoch -count=20
go test ./services/knowpost/rpc/internal/feed
go test ./services/knowpost/rpc/internal/logic/knowpost
go test ./services/knowpost/rpc/internal/svc
go test ./services/knowpost/cmd/knowpost/internal/config
go test ./services/knowpost/cmd/knowpost
go test ./cmd/loadtest
go test ./deploy/compose
go vet 相关 WP2/WP3 包
git diff --check
```

Windows 当前 `CGO_ENABLED=0`；临时启用后仍因 PATH 中没有 `gcc` 而无法构建 race runtime，因此没有声称 `go test -race` 通过。

Docker 使用本地已有层执行 `docker compose build --pull=false knowpost` 成功；KnowPost 替换后 `healthy`、restart=0，Hybrid smoke 通过：reader=3099、normal fanout=100、BigV fanout=0、Kafka events=1。构建过程没有镜像下载故障；若后续出现镜像拉取或构建缓慢，仍按项目约定归类为 Docker registry/mirror 问题，而不是网络连通性问题。

## 5. 下一步

进入 WP4：新增独立 consumer group `knowpost-feed-relation-epoch`，消费 Relation `following` Outbox，对 `FromUserId` 执行 `BumpRelation`。重复、乱序事件只允许造成额外 Bump；Redis 写失败必须返回 error 触发 Kafka 重试。
