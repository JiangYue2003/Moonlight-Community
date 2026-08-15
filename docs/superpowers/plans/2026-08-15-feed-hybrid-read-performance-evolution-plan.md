# Feed 推拉混合读路径性能演进实施计划

日期：2026-08-15

依据：`docs/superpowers/specs/2026-08-15-feed-hybrid-read-performance-evolution-design.md`

状态：待执行批准

## 1. 执行规则

本计划只实施 Spec 的 Phase 0 至 Phase 3。Feed Head 全面物化和 Cursor 分页属于 Phase 4，未经新的设计批准不得开始。

执行时必须遵守：

- 每个工作包先写失败测试或建立性能基线，再做最小实现。
- 每个工作包单独运行定向测试，再运行相关包回归。
- 达不到阶段门禁时停止推进，保留数据并分析，不用后续缓存掩盖前一阶段失败。
- 所有新能力默认关闭，通过 Feature Flag 逐项启用。
- 同一性能对照必须使用同一拓扑、资源、数据集、随机种子和日志级别。
- Docker 镜像构建或拉取异常按 registry/mirror 问题处理，不归因于网络连通性。
- Docker 受阻时保留已构建中间件容器，项目服务改为本地运行；切换拓扑后重建 Phase 0 基线。
- 不清理、不重建 MySQL、Redis、Kafka、etcd 等现有数据卷。
- 不使用 `git reset --hard`、`git checkout --` 或其他会覆盖工作区改动的命令。

当前工作区已有大量未提交修改，开始每个工作包前必须保存：

```powershell
git status --short
git diff -- <本工作包涉及的已有文件>
```

已有脏文件不能直接整文件提交。只有以下两种情况可以创建实现提交：

1. 文件在工作包开始前是干净的；或
2. 用户明确确认该文件的既有改动可以一起提交。

否则保留未提交改动，并在阶段报告中列出本轮新增 diff。新文件可以独立提交。

## 2. 总体工作包

| 工作包 | 目标 | 主要门禁 |
|---|---|---|
| WP0 | 锁定环境和未优化基线 | 环境可复现、冒烟通过 |
| WP1 | 开发环境 pprof | KnowPost/Relation Profile 可采集 |
| WP2 | Feed 分阶段指标 | 仪表开销不超过 5% |
| WP3 | Relation/Content EpochStore | 原子递增、故障旁路测试通过 |
| WP4 | Relation epoch 消费者 | 事件重复/乱序只造成额外失效 |
| WP5 | 完整 RouteSnapshot | Relation RPC/request 降低至少 80% |
| WP6 | Inbox + BigV 合并 Pipeline | 冷路径 P95 降低 15% 或 Redis commands/request 降低 20% |
| WP7 | 第一页 L1/L2 Cache | Fresh 命中正确、冷路径可回退 |
| WP8 | Content Safety 失效 | 删除/私密/下架不返回旧页面 |
| WP9 | Singleflight + SWR + 有界刷新 | 集中失效无协程/依赖雪崩 |
| WP10 | 压测工具与 Docker 本地混合手册 | 三种读者基数、拓扑可记录 |
| WP11 | 正式分阶段压测与总结 | 热点 10K 或明确下一瓶颈 |

## 3. WP0：环境、测试与基线锁定

### 3.1 检查项

只读检查：

```powershell
git status --short
docker compose -f deploy/compose/docker-compose.dev.yml config
docker compose -f deploy/compose/docker-compose.dev.yml ps
go version
docker version
```

记录：

- Git HEAD、工作区变更摘要。
- 运行拓扑：`compose-full` 或 `compose-middleware-local-services`。
- Go、Docker、CPU、内存、GOMAXPROCS。
- Feed Strategy、go-zero Mode 和日志级别。
- MySQL、Redis、Kafka、etcd 地址。
- 当前测试数据 manifest 和随机种子。

### 3.2 回归命令

```powershell
go test ./pkg/cachex/... ./pkg/kafkax/...
go test ./services/knowpost/...
go test ./services/relation/...
go test ./services/gateway/...
go test ./cmd/loadtest/...
go test ./deploy/compose/...
```

若全量包受无关服务或外部依赖影响，必须同时保存失败输出，并保证上述 Feed 直接相关包通过后再继续。

### 3.3 冒烟

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase smoke -run-id <run-id> -strategy hybrid
```

验证：

- 普通作者内容来自 Inbox。
- 大 V 内容来自 Outbox。
- Feed 无重复，作者和可见性过滤正确。
- Kafka lag 最终清零。

WP0 不修改业务代码。

## 4. WP1：独立开发调试 HTTP 与 pprof

### 4.1 计划文件

新增：

- `pkg/debughttp/server.go`
- `pkg/debughttp/server_test.go`

修改：

- `services/knowpost/cmd/knowpost/internal/config/config.go`
- `services/knowpost/cmd/knowpost/internal/app/components.go`
- `services/knowpost/cmd/knowpost/main.go`
- `services/knowpost/cmd/knowpost/main_test.go`
- `services/relation/cmd/relation/internal/config/config.go`
- `services/relation/cmd/relation/internal/config/config_test.go`
- `services/relation/cmd/relation/internal/app/components.go`
- `services/relation/cmd/relation/main.go`
- `services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml`
- `services/relation/cmd/relation/etc/relation.yaml`
- `services/relation/cmd/relation/etc/relation-docker.yaml`
- `deploy/compose/docker-compose.dev.yml`
- `deploy/compose/core_stack_contract_test.go`

### 4.2 测试先行

为 `pkg/debughttp` 写测试：

- `Enabled=false` 不监听端口。
- 启用后 `/debug/pprof/`、`/debug/pprof/goroutine` 可访问。
- Context 取消后服务关闭。
- 监听失败返回错误并由 merged runner 触发受控退出。
- 使用独立 `http.ServeMux`，不污染默认 mux。

配置测试：

- KnowPost 本地端口为 `127.0.0.1:6064`。
- Relation 本地端口为 `127.0.0.1:6066`。
- Docker 内部监听地址可被容器访问，但宿主机只映射到 `127.0.0.1`。
- 关闭开关时不暴露端口。

### 4.3 最小实现

- `debughttp.Server` 实现 `Name() string` 和 `Run(ctx) error`，可直接作为两个 merged runner 的 Component。
- 使用 Go 标准库 `net/http/pprof` 注册到私有 mux。
- 不把 pprof 注册到 gRPC 9004/9006 或 Prometheus 9104/9106。
- Compose 只增加回环地址端口映射，不改现有健康检查。

### 4.4 验证

```powershell
go test ./pkg/debughttp/... ./services/knowpost/cmd/knowpost/... ./services/relation/cmd/relation/... ./deploy/compose/...
go tool pprof -top http://127.0.0.1:6064/debug/pprof/profile?seconds=10
go tool pprof -top http://127.0.0.1:6066/debug/pprof/profile?seconds=10
```

检查点：不启用时行为完全不变；启用时只在开发环境可达。

## 5. WP2：Feed 分阶段指标与 Profile 基线

### 5.1 计划文件

新增：

- `services/knowpost/rpc/internal/feed/observer.go`
- `services/knowpost/rpc/internal/feed/observer_test.go`

修改：

- `services/knowpost/rpc/internal/feed/reader.go`
- `services/knowpost/rpc/internal/feed/redis_adapter.go`
- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- `cmd/loadtest/metrics.go`
- `cmd/loadtest/report.go`
- `cmd/loadtest/report_test.go`

### 5.2 测试先行

增加可注入 `FeedObserver`，默认 No-op。测试必须证明：

- Relation、Counter、Route、Inbox、BigV Pipeline、Merge/Dedup、FeedItem hydrate、整页总耗时分别记录。
- 成功、错误、超时 outcome 不混淆。
- 指标 label 不包含 userID、postID、runID 等高基数字段。
- Observer 自身 panic 不得影响请求；推荐实现避免产生 panic，而不是在热路径 recover。
- 关闭观察器时不额外分配请求级 map。

### 5.3 最小实现

- 使用固定枚举 stage/outcome。
- 记录调用次数和 duration，不打印热路径成功日志。
- 压测报告增加每成功请求的 Relation/Counter/Redis 调用与冷计算次数；无法获取时明确标记 missing。
- 不在本工作包加入页面缓存指标，后续组件复用同一 observer 接口扩展固定枚举。

### 5.4 验证和 Phase 0 正式基线

```powershell
go test ./services/knowpost/rpc/internal/feed/... ./services/knowpost/rpc/internal/logic/knowpost/... ./cmd/loadtest/...
go test -race ./services/knowpost/rpc/internal/feed/... ./services/knowpost/rpc/internal/logic/knowpost/...
```

正式重跑：

- 并发：32、64、128、256。
- 每档：10 秒预热、60 秒采样、三次重复。
- 场景：Hybrid RPC steady/distributed read。
- 采集 KnowPost 与 Relation CPU、heap、alloc、block、mutex Profile。

门禁：指标开启后的 QPS 或 P95 回退不超过 5%。新结果写入新的 `run_id`，不得覆盖 2026-08-14 基线。

## 6. WP3：FeedEpochStore

### 6.1 计划文件

新增：

- `services/knowpost/rpc/internal/feedepoch/store.go`
- `services/knowpost/rpc/internal/feedepoch/store_test.go`
- `services/knowpost/rpc/internal/feedepoch/keys.go`

修改：

- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- `services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml`

### 6.2 接口

```go
type Store interface {
    Relation(ctx context.Context, userID int64) (uint64, error)
    BumpRelation(ctx context.Context, userID int64) (uint64, error)
    Safety(ctx context.Context) (uint64, error)
    BumpSafety(ctx context.Context) (uint64, error)
    MarkSafetyPending()
    SafetyPending() bool
    FlushPendingSafety(ctx context.Context) error
}
```

### 6.3 测试先行

- 不存在的 epoch 返回 0。
- `INCR` 单调递增，并发不丢增量。
- Relation epoch 使用约 1 秒独立 L1；Bump 后当前进程立即可见。
- Safety Bump 失败时设置 pending/bypass。
- 多次失败可以合并为一次补偿 Bump，因为一次新 epoch 足以废弃所有旧页面。
- 补偿成功前 `SafetyPending=true`，成功后清除。
- Redis 故障不返回伪造的新 epoch。
- key 使用测试前缀时可按 `run_id` 定向清理。

### 6.4 最小实现

- 权威 epoch 存 Redis String，不设置短 TTL。
- Epoch L1 使用独立 Ristretto 预算，不与 FeedMine、RouteSnapshot、PageCache 共用。
- Current Relation/Safety 的 Redis miss 视为 0。
- 非法 userID 直接返回参数错误。
- 不在本工作包接 Kafka 或页面缓存。

### 6.5 验证

```powershell
go test ./services/knowpost/rpc/internal/feedepoch/... ./services/knowpost/rpc/internal/svc/...
go test -race ./services/knowpost/rpc/internal/feedepoch/...
```

## 7. WP4：Relation Epoch 消费者

### 7.1 计划文件

新增：

- `services/knowpost/rpc/internal/listener/relation_epoch.go`
- `services/knowpost/rpc/internal/listener/relation_epoch_test.go`

修改：

- `services/knowpost/rpc/internal/listener/invalidation.go`
- `services/knowpost/rpc/app/app.go`
- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml`

### 7.2 设计约束

- 消费 `canal-outbox`，使用独立 group：`knowpost-feed-relation-epoch`。
- 同 group 的多个 KnowPost 实例只消费一次事件并更新 Redis 权威 epoch；不把 Kafka group 当广播。
- 使用 `pkg/canalx.ParseFlat` 和 `ExtractOutboxRows`，只处理 `aggregate_type=following`。
- 解析 `services/relation/shared/event.RelationEvent`，对 `FromUserId` 执行 BumpRelation。
- 重复、乱序和 offset commit 重试允许多 Bump；只造成额外缓存 miss，不造成错误关系快照。
- 坏 Canal 消息和坏 payload 记录限频错误并跳过，Redis 写失败返回 error 让 Kafka 重试。
- Counter consumer 与 Relation epoch consumer 使用独立 goroutine和独立 group；一个退出不能静默伪装成另一个仍健康。

### 7.3 测试先行

- FollowCreated/FollowCanceled 都递增 from-user epoch。
- 非 following outbox 不处理。
- 一条 Canal 消息包含多行 outbox 时逐行处理。
- 重复和逆序消息最终 epoch 只增不减。
- Redis 失败返回 error，不提交成功语义。
- Context 取消时消费者正常退出。

### 7.4 验证

```powershell
go test ./services/knowpost/rpc/internal/listener/... ./services/knowpost/rpc/app/...
```

集成验证：执行 Follow/Unfollow，轮询 `feed:relation:epoch:{userID}` 变化，同时确认原 Relation Syncer 继续维护 ZSet。

## 8. WP5：完整 RouteSnapshotCache

### 8.1 计划文件

新增：

- `services/knowpost/rpc/internal/feed/route_snapshot.go`
- `services/knowpost/rpc/internal/feed/route_snapshot_test.go`

修改：

- `services/knowpost/rpc/internal/feed/reader.go`
- `services/knowpost/rpc/internal/feed/reader_test.go`
- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- `services/knowpost/rpc/internal/svc/feed_strategy_test.go`
- 两份 KnowPost merged 配置文件。

### 8.2 核心行为

`Prepare` 改为：

1. 读取用户 relation epoch。
2. 用 `feed:route:v2:{userID}:e{epoch}` 查询独立 RouteSnapshot L1。
3. 命中时直接构建 `FeedReadSnapshot`，不调用 Relation/Counter。
4. miss 时使用完整 route key 进入 Singleflight。
5. 双检后用独立 2 秒 Context 调 Relation 和 Counter。
6. 只在 Relation 和分类数据有效时缓存完整 snapshot。
7. EpochStore 故障时旁路 RouteCache并回源，不使用无法验证的旧快照。

Snapshot 内部的 following set 和 BigV slice 构建后只读。测试若发现调用方会修改，改为 clone；不预先复制每请求 map。

### 8.3 测试先行

- 同 user/epoch 第二次 Prepare 不调用 Relation/Counter。
- epoch 变化强制回源。
- TTL 到期后重新分类大 V 阈值。
- 并发相同 user/epoch 只触发一次 Relation/Counter。
- 第一个 waiter 取消不取消共享查询；其他 waiter 可成功。
- Relation/Counter 错误不缓存。
- Push/Pull 策略保持原语义。
- RouteSnapshot 使用独立 L1，ServiceContext 不再传入 `l1FeedMine`。
- Race 测试验证只读 snapshot 无数据竞争。

### 8.4 验证与门禁

```powershell
go test ./services/knowpost/rpc/internal/feed/... ./services/knowpost/rpc/internal/svc/...
go test -race ./services/knowpost/rpc/internal/feed/...
```

运行 Phase 1 c32/c64/c128/c256 对照。必须满足：

- 热场景 Relation RPC/request 至少下降 80%。
- QPS 提升至少 20%，或单位成功请求 CPU 下降至少 20%。
- 正确性零失败。

不满足即暂停，不进入 WP6。

## 9. WP6：Inbox + BigV Outbox 合并 Pipeline

### 9.1 计划文件

修改：

- `services/knowpost/rpc/internal/feed/reader.go`
- `services/knowpost/rpc/internal/feed/reader_test.go`
- `services/knowpost/rpc/internal/feed/redis_adapter.go`
- `services/knowpost/rpc/internal/feed/redis_adapter_test.go`

### 9.2 测试先行

- 第一批 Pipeline 同时包含 Inbox 和受批次上限保护的大 V Outbox。
- 只有 Inbox、没有大 V 时仍返回正确结果。
- Inbox 请求失败时降级为空，但 BigV 结果仍可返回。
- 某个 BigV 失败只丢弃该路结果。
- 超过批次上限时请求被正确分批，最终 Top-N 与旧实现一致。
- 不支持 batch 接口的 mock/实现继续走原读路径。
- Context 取消时不继续提交后续批次。

### 9.3 最小实现

- 把 `readInbox` 与 `pullFromBigVsBatch` 的第一轮 Redis 请求合并。
- 保持每轮请求数有上限；Inbox 占用第一批的一个位置。
- 继续边读边维护固定大小 Top-N，不收集所有大 V 帖子。
- 不在本工作包引入 sync.Pool、动态并发或新依赖。

### 9.4 验证与门禁

```powershell
go test ./services/knowpost/rpc/internal/feed/...
go test -tags=integration ./services/knowpost/rpc/internal/feed/...
```

满足以下之一并且无其他回退：

- 冷路径 P95 至少下降 15%；或
- Redis commands/request 至少下降 20%。

## 10. WP7：第一页 L1/L2 PageCache 基础路径

### 10.1 计划文件

新增：

- `services/knowpost/rpc/internal/cache/userfeed/key.go`
- `services/knowpost/rpc/internal/cache/userfeed/codec.go`
- `services/knowpost/rpc/internal/cache/userfeed/cache.go`
- `services/knowpost/rpc/internal/cache/userfeed/cache_test.go`

修改：

- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic_test.go`
- 两份 KnowPost merged 配置文件。

### 10.2 API 轮廓

```go
type Loader func(ctx context.Context) (*pb.FeedPage, error)

type Cache interface {
    GetOrLoad(ctx context.Context, req Request, loader Loader) (*pb.FeedPage, Source, error)
    Close() error
}
```

`Request` 包含 userID、strategyVersion、relationEpoch、safetyEpoch、page 和 size。

### 10.3 测试先行

- 只有 Hybrid `page=1,size=20` 进入缓存；其他请求 bypass。
- key 包含 user、strategy、relation epoch、safety epoch、page、size。
- L1 Fresh 命中不访问 Redis、不执行 loader。
- L2 Fresh 命中反序列化 Protobuf、回填 L1、不执行 loader。
- miss 执行 loader，成功后写 L2 和 L1。
- loader 错误、空结果错误或 protobuf 解码失败不污染缓存。
- L2 写失败仍返回本次正确结果。
- EpochStore 错误或 SafetyPending 时 bypass 页面缓存。
- L1 初始 cost 使用 `2*proto.Size(page) + len(key) + 256`，再根据 heap Profile 校准，避免只按序列化大小低估 Go 对象图。
- 缓存返回对象在并发 gRPC marshal 下无竞态；若测试证明调用链会修改，改为 clone。

### 10.4 Logic 重构

- 将当前 `GetUserFeed` 主体移动到私有 `computeUserFeed(ctx,in)`，保持算法不变。
- 外层 `GetUserFeed` 只负责规范参数、读取 epoch、调用 PageCache。
- loader 接收 Cache 提供的 Context，不能闭包绑定已取消的请求 Context。
- 第一版只实现 Fresh L1/L2，不在同一工作包加入 SWR。

### 10.5 验证

```powershell
go test ./services/knowpost/rpc/internal/cache/userfeed/... ./services/knowpost/rpc/internal/logic/knowpost/...
go test -race ./services/knowpost/rpc/internal/cache/userfeed/... ./services/knowpost/rpc/internal/logic/knowpost/...
```

Feature Flag 顺序：`off -> l2 -> l1-l2`。每一步冒烟后再继续。

## 11. WP8：Content Safety Epoch 写路径

### 11.1 计划文件

修改：

- `services/knowpost/rpc/internal/logic/knowpost/cache_invalidation_helper.go`
- `services/knowpost/rpc/internal/logic/knowpost/cache_invalidation_helper_test.go`
- `services/knowpost/rpc/internal/logic/knowpost/patchmetadatalogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/updatetoplogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/updatevisibilitylogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/deletelogic.go`
- 对应测试文件。

### 11.2 顺序

```text
数据库写成功
  -> 现有 Detail / FeedItem / Public / Mine 缓存失效
  -> BumpSafety
  -> 成功后返回
```

Publish 不 BumpSafety，新内容依靠 3 至 5 秒 Fresh TTL。

### 11.3 测试先行

- Delete、UpdateVisibility、UpdateTop、PatchMetadata 成功后各 Bump 一次。
- DB 写失败时不 Bump。
- Bump 失败时写请求仍按原业务结果返回，但设置 SafetyPending，所有页面缓存 bypass。
- 下一次读请求先尝试 FlushPendingSafety；成功前不得命中旧页。
- 多次写期间的 pending 可以合并，但成功补偿后 epoch 必须大于所有旧页携带值。
- Publish 不触发全局 safety 冷却。
- 旧 L1、L2 和未来 stale 页面在 safety epoch 变化后都不可返回。

### 11.4 验证

```powershell
go test ./services/knowpost/rpc/internal/logic/knowpost/... ./services/knowpost/rpc/internal/feedepoch/...
```

集成测试测量删除、转私密和下架到 Feed 不可见的时间，正常目标 1 秒。

## 12. WP9：Singleflight、SWR 与有界刷新

### 12.1 计划文件

新增：

- `services/knowpost/rpc/internal/cache/userfeed/clock.go`
- `services/knowpost/rpc/internal/cache/userfeed/refresh.go`
- `services/knowpost/rpc/internal/cache/userfeed/refresh_test.go`

修改：

- `services/knowpost/rpc/internal/cache/userfeed/cache.go`
- `services/knowpost/rpc/internal/cache/userfeed/cache_test.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`

### 12.2 测试先行

使用 fake clock 和可控 loader，不用长时间 `time.Sleep`：

- 相同完整 page key 的并发 miss 只执行一次 loader。
- 不同 user key 不会被错误合并。
- 首个 waiter 取消不污染共享 loader；每个 waiter 可按自身 Context 返回。
- Fresh 剩余时间进入其 TTL 最后 20% 时只投递一个刷新任务。
- Stale 且 epoch 匹配时可返回并异步刷新。
- epoch 不匹配、SafetyPending 或超过 Stale window 时禁止返回 stale。
- 刷新队列满时不创建额外 goroutine。
- 队列满时：Fresh 继续返回；合法 Stale 继续返回但不刷新；无缓存则同步执行受超时约束的 loader。
- Worker 并发不超过配置值。
- Loader 使用独立、有限时的 background Context。
- `Close` 停止 worker，不泄漏 goroutine。
- L2 payload 使用二进制时间头加 Protobuf body，不使用 JSON。

### 12.3 最小实现

- Fresh TTL：L1 500ms 至 1s，L2 3s 至 5s。
- Jitter：±20%，由测试注入确定性随机源。
- Stale window：最多 10s，只用于依赖异常或刷新窗口。
- 初始 worker=32、queue=1024、refresh timeout=2s，全部配置化。
- 本阶段不加分布式锁；当前目标是一台 KnowPost 实例。

### 12.4 验证

```powershell
go test ./services/knowpost/rpc/internal/cache/userfeed/...
go test -race ./services/knowpost/rpc/internal/cache/userfeed/...
```

集中失效测试必须同时观察 goroutine、refresh queue、Relation QPS 和 Redis ops。

## 13. WP10：压测工具、三种基数与 Docker 混合拓扑

### 13.1 计划文件

修改：

- `cmd/loadtest/load_test.yaml`
- `cmd/loadtest/config_test.go`
- `cmd/loadtest/dataset.go`
- `cmd/loadtest/dataset_test.go`
- `cmd/loadtest/runner.go`
- `cmd/loadtest/monitor.go`
- `cmd/loadtest/report.go`
- `cmd/loadtest/report_test.go`
- `cmd/loadtest/run_feed_matrix.ps1`
- `cmd/loadtest/FEED_MATRIX_GUIDE.md`

新增：

- `cmd/loadtest/FEED_EVOLUTION_GUIDE.md`
- `scripts/start-feed-local-services.ps1`
- `scripts/stop-feed-local-services.ps1`

### 13.2 压测能力

- 读者基数预设：20、约 1,200、20,000。
- 并发预设：32、64、128、256。
- 预热 10 秒、采样 60 秒、重复三次。
- 可显式选择 cache state：cold、l2-warm、l1-warm、expire-together。
- 报告记录 topology：`compose-full` 或 `compose-middleware-local-services`。
- 报告记录 Route/Page cache outcome、冷计算、epoch、刷新队列和 pprof 文件位置。
- 压测前检测客户端 CPU；客户端饱和时报告无效。

### 13.3 Docker 镜像源降级路径

脚本与手册必须体现：

- 不执行全量 `docker compose pull`。
- 不清空镜像、BuildKit 缓存或数据卷。
- Compose 只启动本地已经存在镜像的中间件，例如 etcd、Kafka、ZooKeeper、Canal 和本轮必需依赖。
- Gateway、KnowPost、Relation、Counter、User 等项目服务在本地以隐藏进程运行。
- 本地服务使用宿主机映射端口访问 Kafka/etcd，并把可达的宿主机地址注册到 etcd。
- 脚本记录 PID、stdout、stderr，只停止自己启动的进程。
- `-WhatIf` 或等价 dry-run 能显示将启动/停止的目标。
- 切换拓扑后强制生成新 Phase 0 baseline，工具拒绝把不同 topology 自动算作优化前后对照。

### 13.4 测试

- 配置解析和默认值测试。
- 三种 cardinality 的数据拓扑测试。
- topology 不同的报告拒绝比较。
- 本地脚本 dry-run 不启动进程、不修改容器。
- 清理只处理当前 run_id。

验证：

```powershell
go test ./cmd/loadtest/... ./deploy/compose/...
& ./scripts/start-feed-local-services.ps1 -WhatIf
& ./scripts/stop-feed-local-services.ps1 -WhatIf
```

## 14. WP11：正式阶段压测与交付报告

### 14.1 每阶段固定流程

1. 运行相关单元、Race 和集成测试。
2. 关闭本阶段 Feature Flag，跑同口径 control。
3. 只开启本阶段 Feature Flag，跑 treatment。
4. 按 32/64/128/256 进行 10 秒预热 + 60 秒采样，重复三次。
5. 保存环境、Git、拓扑、配置和原始结果。
6. 对照阶段门禁；通过才进入下一阶段。
7. 恢复 Hybrid 默认策略和安全配置。

### 14.2 Phase 3 最终矩阵

热点目标：

- 20 个热点读者，Hybrid RPC，page1/size20，缓存预热。
- 至少一档稳定达到 10,000 成功 QPS。
- P95 < 50ms，P99 < 100ms，错误率 < 0.1%，正确性失败为 0。
- L1+L2 Fresh 命中率至少 97%。

分散目标：

- 约 1,200 读者，相对 Phase 0 同场景 QPS 至少提升 30%。
- 20,000 读者不要求 10K，但吞吐和 P95 不得比同拓扑 Phase 0 回退超过 10%。

一致性与故障：

- 新发布内容正常 5 秒内可见。
- 取关、删除、转私密、下架按 Spec 收敛。
- Redis、Relation、epoch consumer 暂停与恢复。
- 页面集中失效时无 refresh/goroutine/依赖雪崩。
- Gateway 入口完成对照，但不与 RPC 10K 目标混为同一口径。

### 14.3 交付

新增最终报告：

- `results/feed-loadtest/<run-id>/EVOLUTION_REPORT.md`

更新：

- `cmd/loadtest/FEED_EVOLUTION_GUIDE.md`
- 必要的配置说明和运行命令。

报告必须回答：

- 每个工作包实际减少了哪些 RPC、Redis command 或 CPU。
- 10K 是否达到；若未达到，第一瓶颈是什么。
- 热点、分散、高基数三种结果分别是多少。
- Compose 全栈和本地混合拓扑是否分别重建基线。
- 哪些 Feature Flag 保留启用，哪些回滚。
- 是否有足够证据进入独立的 Phase 4 Feed Head 设计。

## 15. 最终回归清单

```powershell
go test ./pkg/cachex/... ./pkg/kafkax/... ./pkg/debughttp/...
go test ./services/knowpost/...
go test ./services/relation/...
go test ./services/gateway/...
go test ./cmd/loadtest/...
go test ./deploy/compose/...
go test -race ./services/knowpost/rpc/internal/feed/... ./services/knowpost/rpc/internal/feedepoch/... ./services/knowpost/rpc/internal/cache/userfeed/...
git diff --check
```

容器或本地进程冒烟完成后检查：

- 所有涉及服务健康且无重启。
- Feed Strategy 恢复为 `hybrid`。
- pprof 只在开发配置和本机可达。
- PageCache 默认开关符合最终阶段决定。
- Kafka lag 清零。
- Redis 无 eviction/rejected connection。
- 没有测试进程、刷新 worker 或本地服务进程泄漏。
- 没有清理非当前 run_id 的数据。

## 16. 停止条件

遇到以下任一情况立即停止当前阶段：

- Feed 正确性失败、越权内容或敏感内容持续可见。
- 新增缓存无法可靠旁路到冷路径。
- 线程、goroutine、连接数或刷新队列持续增长。
- 增加并发后吞吐不涨且延迟/错误继续上升。
- A/B 对照的拓扑、资源或数据集不一致。
- Docker registry/mirror 问题迫使拓扑变化但尚未重建基线。
- 工作包需要覆盖无法确认归属的既有修改。
- 阶段门禁未通过却只能依赖下一阶段掩盖。

停止不是项目失败。保存证据、恢复 Feature Flag、明确第一瓶颈后，再决定修改本阶段设计或结束目标。
