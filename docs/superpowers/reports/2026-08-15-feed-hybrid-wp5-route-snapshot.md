# Feed Hybrid WP5：版本化 RouteSnapshot 检查点

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）；Phase 1 性能、正确性与稳定性门禁通过

## 1. 结论

WP5 已把 Feed Hybrid 热读路径中的关系路由计算改为版本化完整快照：

```text
GetUserFeed
  -> GET feed:relation:epoch:{userID}
  -> L1 feed:route:v2:{userID}:e{epoch}
       hit  -> 复用完整 followings / bigVAuthors / followingSet
       miss -> Singleflight -> Relation + Counter -> 同步发布 L1
  -> Inbox + BigV Outbox + Merge/Dedup + Hydrate
```

正式对照中，Relation RPC/request 从约 `1.0` 降到 `0.0010~0.0011`，降低约 `99.9%`；四个并发档 QPS 中位数均比 WP2 严格基线提升至少 39%，正确性零失败。Phase 1 的主要门禁全部通过，可以进入 WP6。

本阶段峰值无错吞吐为 c64 的 `3978.74 QPS`，P95 `22.054 ms`。推荐稳定并发为 c64，拐点为 c128；c128/c256 吞吐不再增长，P95 分别升至 `49.301 ms` 和 `88.936 ms`。这说明 Relation/MySQL 已不再是热路径主瓶颈，下一约束是每请求约 3 次 Redis 操作、KnowPost CPU 和 FeedItem hydrate。

## 2. 实现边界

主要新增：

- `services/knowpost/rpc/internal/feed/route_snapshot.go`
- `services/knowpost/rpc/internal/feed/route_snapshot_test.go`

主要修改：

- `services/knowpost/rpc/internal/feed/reader.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- KnowPost Config、两份 YAML 与 Compose 环境变量。
- `cmd/loadtest` 运行态开关校验和报告环境字段。

核心语义：

- 仅在 `Feed.RouteSnapshot.Enabled=true` 且依赖就绪时启用；默认关闭可即时回到旧路径。
- 完整 key 为 `feed:route:v2:{userID}:e{epoch}`，关系 epoch 变化后旧快照自然不可命中。
- RouteSnapshot 使用独立 Ristretto L1，不与 FeedMine 或 Epoch L1 共用预算。
- TTL 默认 5 秒且配置强制不允许超过 5 秒，作为事件延迟或丢失时的一致性兜底。
- 热命中同时跳过 Relation 和 Counter；Push/Pull 行为保持原语义。
- 冷 miss 按完整 key Singleflight 合并；共享查询使用独立 2 秒 Context，第一个 waiter 取消不会取消其他 waiter。
- Relation 错误不缓存；Counter 失败保留既有全普通作者 fallback，但不缓存不完整分类。
- Epoch Redis 错误时旁路不可验证的快照并回源；请求取消不会误入旁路。
- Snapshot 类型和字段保持包私有，发布时 clone 可变 slice，命中后只读复用 map/slice。
- Ristretto `SetWithTTL` 后显式 `Wait`，保证 Singleflight 释放前快照可见；Set buffer 拒绝仅做一次有界重试。

## 3. 审查中发现并修复的问题

复审曾发现两个 Important，均已修复并补回归：

1. 配置只给非正 TTL 设置默认值，未阻止大于 5 秒的 TTL，可能突破 Spec 的一致性上界。现在 ServiceContext 建立外部资源前即拒绝该配置。
2. Ristretto 写入异步，Singleflight 完成到缓存真正可见之间可能出现第二波 Relation/Counter 冷填充。现在 RouteSnapshot 发布显式等待可见，并用真实 L1 的立即二读和 post-flight burst 测试锁定行为。

WP4 同期审查还修复了共享 Kafka 消费器的 offset 越过与取消后 tight loop 风险；这些修复是 RelationEpoch 可可靠驱动 RouteSnapshot 失效的前置条件。

最终复审结论：Ready，Critical 0 / Important 0。

## 4. 自动化验证

RouteSnapshot 测试覆盖：

- 相同 user/epoch 第二次 Prepare 不调用 Relation/Counter。
- feature flag 关闭时保留旧路径。
- key 包含 user 与 epoch；epoch 变化强制回源。
- 5 秒 TTL 后重新执行大 V 阈值分类。
- 32 个并发冷 miss 只触发一次 Relation/Counter。
- 首个 waiter 取消不取消共享查询。
- Relation/Counter/Epoch 错误的缓存与旁路语义。
- Epoch 失败日志限频，请求取消不产生旁路调用。
- Push/Pull 语义不变。
- 真实 Ristretto L1 无手工测试 Wait 的立即二读与 32 请求 post-flight burst。
- 64 个并发热读者共享只读状态。
- Set buffer 拒绝只进行一次有界重试。
- TTL 大于 5 秒时配置拒绝，Route L1 独立预算。

已通过：

```text
go test ./pkg/cachex ./pkg/canalx ./pkg/kafkax
go test ./services/knowpost/rpc/internal/feedepoch
go test ./services/knowpost/rpc/internal/listener
go test ./services/knowpost/rpc/app
go test ./services/knowpost/rpc/internal/feed
go test ./services/knowpost/rpc/internal/logic/knowpost
go test ./services/knowpost/rpc/internal/svc
go test ./services/knowpost/cmd/knowpost/internal/config
go test ./services/knowpost/cmd/knowpost
go test ./cmd/loadtest ./deploy/compose
go vet 相关 WP2-WP5 包
git diff --check
```

真实 L1 可见性测试额外以 `-count=50` 重复通过。Windows 当前 `CGO_ENABLED=0` 且 PATH 中没有 `gcc`，无法构建 Go race runtime，因此没有声称 `go test -race` 通过；并发正确性由确定性单测、重复测试和正式高并发压测共同验证。

## 5. Docker 与真实一致性验证

只复用现有中间件和本地构建层重建 KnowPost：

```text
docker compose -f deploy/compose/docker-compose.dev.yml build --pull=false knowpost
FEED_STRATEGY=hybrid
FEED_OBSERVABILITY=true
FEED_RELATION_EPOCH_ENABLED=true
FEED_ROUTE_SNAPSHOT_ENABLED=true
docker compose -f deploy/compose/docker-compose.dev.yml up -d --no-deps --force-recreate --wait knowpost
```

结果：

- KnowPost image：`sha256:29a9e133a35bcf3b7c598ed58c938cc8605fa41c2d4a926e46b8ae7424cf79fc`
- 容器运行态显式记录四个开关；`healthy`，restart=0。
- RelationEpoch Follow/Unfollow 真实集成测试 0.24 秒通过。
- Hybrid 冒烟通过：normal fanout=100、bigV fanout=0、Kafka event=1、Feed 可读。
- 正式压测后容器仍 `healthy`、restart=0；最近 20 分钟定向扫描 error/fatal/panic/RouteSnapshot/RelationEpoch 失败为 0 条。

本轮构建没有镜像下载问题。后续若 Docker 构建或拉取缓慢，继续按 Docker registry/mirror 问题处理，不归因于网络连通性；必要时保留中间件容器并在宿主机运行项目服务，切换拓扑后重新建立对应基线。

## 6. Phase 1 正式性能结果

环境与口径：

- 数据集：`feedbench-20260814c`
- 入口：RPC；场景：`distributed-read`
- 每档：10 秒预热、60 秒正式采样、3 次重复
- 新报告：`results/feed-loadtest/feed-wp5-route-enabled-20260815`
- 旧基线：`results/feed-loadtest/feed-wp2-phase0-strict-20260815`
- 所有 12 个 trial `complete=true`、失败 0、timeout 0、degraded=false、Kafka lag=0。

| 并发 | WP2 QPS | WP5 QPS | QPS 提升 | WP2 P95 | WP5 P95 | P95 下降 | Relation/request 下降 |
|---:|---:|---:|---:|---:|---:|---:|---:|
| 32 | 1945.29 | 3572.68 | 83.66% | 20.847 ms | 12.050 ms | 42.20% | 99.888% |
| 64 | 2327.03 | 3978.74 | 70.98% | 35.336 ms | 22.054 ms | 37.59% | 99.899% |
| 128 | 2590.31 | 3896.52 | 50.43% | 64.342 ms | 49.301 ms | 23.38% | 99.897% |
| 256 | 2793.47 | 3884.14 | 39.04% | 123.603 ms | 88.936 ms | 28.05% | 99.897% |

WP5 内部三轮波动：

- c32 QPS spread 1.00%。
- c64 QPS spread 0.33%。
- c128 QPS spread 3.27%。
- c256 QPS spread 0.67%。

容量结论：

- 峰值无错 QPS：`3978.74`，c64。
- 推荐稳定并发：c64。
- 新拐点：c128。
- 无错上界（本轮已测）：c256；不等同于推荐运行点。
- c128/c256 吞吐相对 c64 不增长，但 P95 分别约为 c64 的 2.24 倍和 4.03 倍，系统已进入非线性排队区。

## 7. 门禁判定

| 门禁 | 结果 | 判定 |
|---|---|---|
| Relation RPC/request 至少下降 80% | 各档约下降 99.9% | 通过 |
| QPS 提升至少 20%，或单位请求 CPU 下降 20% | 各档 QPS 提升 39.04%~83.66% | 通过 |
| Follow/Unfollow、重复/乱序、延迟/丢失恢复 | 单测、Kafka 语义测试和真实集成通过 | 通过 |
| 无 Feed 错误、超时、降级或永久不一致证据 | 12/12 完整，失败/timeout=0，degraded=false | 通过 |
| 运行稳定性 | healthy，restart=0，定向错误日志 0 | 通过 |

## 8. 下一步

允许进入 WP6：把 Inbox 与第一批 BigV Outbox 合并为一个受控 Redis Pipeline，并以冷路径 P95 至少下降 15%或 Redis commands/request 至少下降 20%作为门禁。

WP5 不足以达到单机 10K。当前每个成功请求仍约执行：

- `redis.inbox_read = 1`
- `redis.bigv_pipeline = 1`
- `redis.feed_item_mget = 1`
- 合计 Redis dependency calls 约 `3.00~3.01/request`

因此 WP6 应先减少 Redis 往返并重新 Profile。若冷路径仍停留在约 4K QPS，再进入 WP7-WP9 的第一页最终结果缓存；不得用 PageCache 掩盖 WP6 自身未达门禁。
