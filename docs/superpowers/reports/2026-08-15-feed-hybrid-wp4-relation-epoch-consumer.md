# Feed Hybrid WP4：Relation Epoch 消费者检查点

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）；真实 Follow/Unfollow 集成验证通过

## 1. 结论

WP4 已将 Relation `following` Outbox 接入 Feed relation epoch：

```text
Relation Follow/Unfollow transaction
  -> outbox
  -> Canal / canal-outbox
  -> knowpost-feed-relation-epoch consumer group
  -> INCR feed:relation:epoch:{fromUserID}
```

消费者使用独立 group，不与 `relation-syncer` 或 KnowPost Counter invalidation 争抢消息。重复、乱序和重投只会额外递增单调 epoch，造成一次额外缓存 miss，不会把旧关系快照重新标记为有效。

该能力仍默认关闭；Docker 通过 `FEED_RELATION_EPOCH_ENABLED=true` 显式启用。WP4 尚未启用 RouteSnapshotCache，因此不会改变当前 Feed 结果或性能路径。

## 2. 实现

新增：

- `services/knowpost/rpc/internal/listener/relation_epoch.go`
- `services/knowpost/rpc/internal/listener/relation_epoch_test.go`
- `services/knowpost/rpc/internal/listener/relation_epoch_integration_test.go`
- `services/knowpost/rpc/app/app_test.go`
- `pkg/kafkax/consumer_test.go`

主要行为：

- 使用 `canalx.ParseFlat` 和 `ExtractOutboxRows` 逐行读取 Outbox。
- 仅处理 `aggregate_type=following` 的 `FollowCreated` / `FollowCanceled`。
- 对合法事件的 `FromUserId` 调用 `FeedEpochStore.BumpRelation`。
- 坏 Canal、坏 payload、未知事件和非法用户使用固定类别 key 限频记录并跳过。
- Redis Bump 失败返回 handler error，禁止提交成功语义。
- Counter 和 Relation epoch listener 独立运行；任一 listener 意外退出都会取消 peer 并让 KnowPost 受控退出，不会静默伪装健康。
- local YAML、Compose 默认值和 Config default 均保持关闭；Docker 配置允许显式 opt-in。

## 3. 审查中发现并修复的 Kafka offset 风险

审查发现原 `pkg/kafkax.RunConsumer` 在 handler 失败后重新 `FetchMessage`。虽然失败消息没有立即 Commit，但后续消息成功后可能提交更大的 offset，从而越过失败消息，永久漏掉 epoch Bump。

共享消费循环已修正为：

1. 每次只 Fetch 一条消息。
2. handler 失败时在同一条已 Fetch 消息上原地退避重试。
3. 只有 handler 成功，或达到有限重试且 DLQ 写入成功后，才 Commit 并 Fetch 下一条。
4. DLQ 写入失败不提交原消息。
5. Context 在 backoff 中取消时返回 `ctx.Err()`，上层将 `context.Canceled` 作为正常关闭；不会形成取消后的 CPU tight loop。

可注入 reader/writer 测试锁定了相同 offset 的处理顺序、Commit 顺序、DLQ 失败不提交以及生产 backoff 的取消退出语义。

## 4. 测试证据

覆盖：

- FollowCreated、FollowCanceled 和单条 Canal 多 Outbox 行。
- 非 following、坏 Canal、坏 payload、未知类型、非法用户跳过。
- 重复和逆序事件执行额外单调 Bump。
- Redis 失败返回 error。
- 独立 topic/group 配置。
- listener error、意外退出和 Context 取消的生命周期监督。
- Kafka 同消息重试期间不 Fetch/Commit 后续 offset。
- poison message 先写 DLQ 再 Commit；DLQ 失败不 Commit。
- Compose opt-in 开关及本地默认关闭配置。

已通过：

```text
go test ./pkg/kafkax -count=10
go test ./pkg/canalx ./pkg/kafkax
go test ./services/knowpost/rpc/internal/feedepoch
go test ./services/knowpost/rpc/internal/listener
go test ./services/knowpost/rpc/app
go test ./services/knowpost/rpc/internal/feed
go test ./services/knowpost/rpc/internal/logic/knowpost
go test ./services/knowpost/rpc/internal/svc
go test ./services/knowpost/cmd/knowpost/internal/config
go test ./services/knowpost/cmd/knowpost
go test ./cmd/loadtest ./deploy/compose
go vet 相关 WP2-WP4 包
git diff --check
```

Windows 当前 `CGO_ENABLED=0` 且 PATH 中没有 `gcc`，因此仍无法构建 Go race runtime；没有声称 `go test -race` 通过。

## 5. Docker 与真实集成验证

使用本地已有依赖和构建层执行：

```text
docker compose -f deploy/compose/docker-compose.dev.yml build --pull=false knowpost
FEED_RELATION_EPOCH_ENABLED=true
docker compose -f deploy/compose/docker-compose.dev.yml up -d --no-deps --force-recreate knowpost
```

结果：

- KnowPost image：`sha256:014363ee5121f87cc01283efa94e108c0eaab9b21f2cc11dd0787e13eba64f1c`
- Container：`healthy`，restart=0，Feed strategy=`hybrid`
- Relation epoch flag：`true`
- Consumer group：`knowpost-feed-relation-epoch`
- Topic：`canal-outbox`，partition 0 current/log-end offset 均为 32777，lag=0

真实测试对开发数据用户 `3124 -> 3125` 执行 Follow 后再 Unfollow：

- 每次写操作都在超时内观察到 `feed:relation:epoch:3124` 严格递增；最终值为 2。
- 原 `relation-syncer` 同步维护 `uf:flws:3124`：Follow 后成员存在，Unfollow 后成员消失。
- 测试结束恢复原始未关注状态，最终 `ZSCORE uf:flws:3124 3125` 为空。
- 完整集成测试耗时 0.25 秒。

构建使用本地缓存，没有重新构建中间件，也没有发生镜像下载问题。后续若 Docker 镜像拉取或构建缓慢，仍按项目约定归类为 Docker registry/mirror 问题，不归因于网络连通性；必要时只保留中间件容器并在宿主机运行项目服务。

## 6. 下一步

进入 WP5：实现完整、版本化的 RouteSnapshotCache。先以测试锁定相同 user/epoch 热命中不再调用 Relation/Counter、epoch 变化强制回源、并发 miss 合并与错误不缓存，再通过 Phase 1 对照门禁验证 Relation RPC/request 至少下降 80%，并验证 QPS 或单位请求 CPU 至少改善 20%。
