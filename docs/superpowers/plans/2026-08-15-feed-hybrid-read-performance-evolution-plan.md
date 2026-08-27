# Feed 推拉混合读路径性能演进实施计划

日期：2026-08-15

依据：`docs/superpowers/specs/2026-08-15-feed-hybrid-read-performance-evolution-design.md`

状态：WP1-WP11 已完成既定 RPC 交付（纯读、Cursor、mutation RPC 正式 A/B）；后续 Gateway 性能矩阵按用户决策取消

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
| WP6 | Inbox + BigV 合并 Pipeline | 冷路径 P95 降低 15% 或 Redis round trips/request 降低 20% |
| WP7 | 第一页 L1/L2 Cache | Fresh 命中正确、冷路径可回退 |
| WP8 | Content Safety 失效 | 删除/私密/下架不返回旧页面 |
| WP9 | Singleflight + SWR + 有界刷新 | 集中失效无协程/依赖雪崩 |
| WP10 | 压测工具与 Docker 本地混合手册 | 已完成；终审 Ready，Critical 0 / Important 0 |
| WP11 | 正式分阶段压测与总结 | 完成：纯读与 Cursor 结果已落盘；mutation RPC-only Control/Treatment 共 18 轮正式 A/B 完成，Gateway 后续矩阵取消 |

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
- `services/knowpost/cmd/knowpost/internal/config/config_test.go`
- `services/knowpost/cmd/knowpost/main.go`
- `services/knowpost/cmd/knowpost/main_test.go`
- `services/relation/cmd/relation/internal/config/config.go`
- `services/relation/cmd/relation/internal/config/config_test.go`
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
- CPU profile 等长请求进行中取消 Context 时，请求和组件都能及时结束且不转成进程错误退出。
- 监听失败返回错误并由 merged runner 触发受控退出。
- 组件显式使用独立 `http.ServeMux`，任何实际 listener 都不得把 `http.DefaultServeMux` 作为 Handler。

配置测试：

- KnowPost 本地端口为 `127.0.0.1:6064`。
- Relation 本地端口为 `127.0.0.1:6066`。
- Docker 内部仍监听 `6064/6066`，宿主机回环映射使用 `16064/16066`。
- Windows/Hyper-V 当前排除 TCP `6034-6133`，因此 Docker 不使用 `6064/6066` 作为宿主机发布端口；可用 `netsh interface ipv4 show excludedportrange protocol=tcp` 复验。
- 关闭开关时不暴露端口。

### 4.3 最小实现

- `debughttp.Server` 实现 `Name() string` 和 `Run(ctx) error`，可直接作为两个 merged runner 的 Component。
- 使用 Go 标准库 `net/http/pprof` 注册到私有 mux。
- `net/http/pprof` 包初始化会向 `http.DefaultServeMux` 注册标准路由，因此安全门禁不表述为“全局 mux 没有注册项”，而是：本组件及仓库其他实际 HTTP listener 均不服务默认 mux；新增 `Handler=nil` 或默认 mux listener 必须被审查阻止。
- 不把 pprof 注册到 gRPC 9004/9006 或 Prometheus 9104/9106。
- Compose 只增加回环地址端口映射，不改现有健康检查。

### 4.4 验证

```powershell
go test ./pkg/debughttp/... ./services/knowpost/cmd/knowpost/... ./services/relation/cmd/relation/... ./deploy/compose/...
go tool pprof -top http://127.0.0.1:16064/debug/pprof/profile?seconds=10
go tool pprof -top http://127.0.0.1:16066/debug/pprof/profile?seconds=10
```

本地进程拓扑继续使用 `6064/6066`；Compose 全栈拓扑使用宿主机 `16064/16066`。不得在同一份报告中混淆二者。

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

### 8.5 执行结果（2026-08-15）

状态：完成；最终复审 Ready（Critical 0 / Important 0）；允许进入 WP6。

- Relation RPC/request 从约 1.0 降至 `0.0010~0.0011`，约下降 99.9%。
- c32/c64/c128/c256 QPS 中位数较 WP2 严格基线分别提升 83.66%、70.98%、50.43%、39.04%。
- 12/12 trial 完整，Feed 失败/timeout=0，degraded=false，Kafka lag=0。
- 峰值无错 QPS `3978.74`（c64），稳定并发 c64，拐点 c128。
- 正式压测后 KnowPost healthy、restart=0，定向错误日志 0 条。
- Windows 缺少 gcc，未声称 `go test -race` 通过；确定性并发单测、重复测试和高并发正式压测通过。
- 详细报告：`docs/superpowers/reports/2026-08-15-feed-hybrid-wp5-route-snapshot.md`。

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
- Redis dependency round trips/request 至少下降 20%。

round trip 以一次单命令调用或一次 `Pipeline.Exec` 计数；命令条数不会因 Pipeline 合并而减少，
不得把该指标表述为 Redis commands/request。正式报告还需对照 Redis `ops/sec`。

### 9.5 执行结果（2026-08-15）

状态：完成；最终复审 Ready（Critical 0 / Important 0）；允许进入可选 WP7。

- 第一批 Inbox + BigV 合并、跨批次 Top-N 等价、部分错误偏移隔离、Context 取消和无 batch 接口回退均有自动化覆盖。
- 普通 Feed 测试、integration build tag 测试、定向重复测试和相关 `go vet` 通过。
- 12/12 正式 trial 完整，Feed 失败/timeout=0，degraded=false，Kafka lag=0。
- Redis dependency round trips/request 从约 3.00 降至约 2.00，下降约 33%；同时保留 Redis ops/s 证明命令负载没有被隐藏。
- c128 三轮中位数 `10661.31 QPS`、P95 `18.112 ms`；稳定点 c64 为 `9866.82 QPS`、P95 `9.171 ms`。
- c32/c64/c128/c256 QPS 较 WP5 分别提升 123.82%、147.99%、173.61%、172.58%，P95 下降 54.46%~63.26%。
- Phase 2 在冷路径上已达到 10K；WP7 变为热点第一页的可选增益阶段，必须保留 WP6 作为回退和对照基线。
- 详细报告：`docs/superpowers/reports/2026-08-15-feed-hybrid-wp6-combined-pipeline.md`。

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

### 10.6 执行结果（2026-08-15）

状态：完成；最终复审 Ready（Critical 0 / Important 0）；PageCache 保持 off，允许进入 WP8。

- 完整 epoch key、仅 Hybrid page1/size20 准入、Fresh L1/L2、Protobuf L2、cost 与故障回退测试通过。
- Epoch error、SafetyPending，以及 cache lookup 中途 pending/relation bump/safety bump 均丢弃旧页并回 WP6。
- L2 get/decode/encode/set 有独立低基数指标和按 operation 限频错误报告。
- loader 使用 Cache 提供的 Context；真实 Ristretto+miniredis 与 64 路只读 marshal 测试通过。
- 普通/integration 回归、50 次定向重复、相关 `go vet` 通过；Windows 缺 gcc，未声称 race 门禁通过。
- Docker `off -> l2 -> l1-l2` 启动与 Hybrid 冒烟通过；40 个真实 page1/size20 RPC 全成功，观察到 `l1_fresh` 与 `miss`。
- 最新 KnowPost 镜像 `sha256:203349ab6e77343a5eb328149da87ac5e664075be9631556813696d84d3f6fb6`；最终运行态 PageCache=off、healthy、restart=0。
- 详细报告：`docs/superpowers/reports/2026-08-15-feed-hybrid-wp7-page-cache-foundation.md`。

## 11. WP8：Content Safety Epoch 写路径

### 11.1 计划文件

新增：

- `services/knowpost/rpc/internal/listener/content_safety_epoch.go`
- `services/knowpost/rpc/internal/listener/content_safety_epoch_test.go`
- `services/knowpost/rpc/internal/logic/knowpost/cache_safety_test.go`
- `cmd/loadtest/content_safety_integration_test.go`

修改：

- `services/knowpost/rpc/internal/logic/knowpost/cache_invalidation_helper.go`
- `services/knowpost/rpc/internal/logic/knowpost/cache_invalidation_helper_test.go`
- `services/knowpost/rpc/internal/logic/knowpost/patchmetadatalogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/updatetoplogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/updatevisibilitylogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/deletelogic.go`
- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go`
- `services/knowpost/rpc/internal/cache/keys.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- KnowPost Config、应用生命周期、YAML、Compose、loadtest 环境快照。
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

### 11.5 执行结果（2026-08-15）

状态：完成；二次复审 Ready（Critical 0 / Important 0）；允许进入 WP9。

- Delete、UpdateVisibility、UpdateTop、PatchMetadata 在 commit 后同步 bump；Publish 不触发全局 safety 冷却。
- commit 后既有缓存失效失败仍推进 safety 并返回原错误；safety bump 失败不篡改已提交业务结果，而是合并 pending 并由后续读尝试补偿。
- PageCache 路径的个人 FeedItem key 携带 safety epoch；pending、epoch 错误、cache lookup 中途版本变化或 PageCache 不准入时，个人 FeedItem cache 与 PageCache 一同 bypass。
- 新增独立 `canal-outbox` consumer group；Updated/Deleted 持久 bump，Published 跳过。PageCache 非 off 时配置强制要求 `SafetyConsumerEnabled=true`。
- 真实开发栈集成测试：转私密 `95.4161ms`、删除 `92.5074ms` 内从 Feed 消失；两次均观察到 epoch `N -> N+2`，证明同步和异步补偿均生效。
- 普通回归、integration 编译、相关 `go vet` 和 `git diff --check` 通过；Windows `CGO_ENABLED=0` 且无 gcc，未声称 race 门禁通过。
- Docker 镜像 `sha256:f604129f626e4259ac1bc76b39aac71a1e842ad82f6c725624608d005a457cca`；KnowPost healthy、restart=0，safety consumer group lag=0。
- 小样本运行态报告 `results/feed-loadtest/feed-wp8-safety-validation-20260815` 仅证明开关、路径和指标，不作为容量结论。
- 详细报告：`docs/superpowers/reports/2026-08-15-feed-hybrid-wp8-content-safety-epoch.md`。

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

### 12.5 执行结果（2026-08-16）

状态：完成；终审 Ready（Critical 0 / Important 0）；允许进入 WP10。

- 相同完整 page key 的冷 burst 合并 L2 查询和 loader；不同 key 不互相阻塞，每个 waiter 使用自己的请求 Context。
- Fresh 最后 20% 与合法 Stale 均可立即返回并投递刷新；worker=`32`、queue=`1024`、loader timeout=`2s`，队列溢出不派生 goroutine。
- queued refresh 可由已过期的前台请求原子接管；running refresh 失败后回退新的有界 Hybrid cold loader，不让刷新协调器成为可用性单点。
- L1/L2 nominal TTL 为 `800ms/4s`，在 ±20% jitter 后分别严格落在 `500ms~1s` 与 `3s~5s`；Stale window 不超过 `10s`。
- queue depth、active workers、pending keys 已接入无用户标签 Gauge；queue full 与 refresh load error 分操作限频报告。
- PageCache 定向包连续 50 轮、KnowPost/Cache/Kafka/Loadtest/Compose 相关回归、`go vet` 与 `git diff --check` 通过。Windows `CGO_ENABLED=0`，`go test -race` 明确报 `-race requires cgo`，未声称 race 门禁通过。
- 最终镜像 `sha256:5a4c1e4832048c71c35fe3eb5d9e133eb7f57091174f31abcbb8d9ff4a048ef5`；KnowPost healthy、restart=0，全部 Feed 开关启用，三个 Kafka group lag=0。
- 真实 SWR 定向样本 40/40 成功、P95 `0.804ms`，观察到 enqueue/load 各 1 次，结束时 queue/active/pending 均为 0。
- c16、400 请求热读小样本 400/400 成功、QPS `10694.99`、P95 `1.3515ms`、cold compute=1、singleflight shared=15。该结果只证明运行路径，不替代 WP10 的固定时长三轮容量结论。
- 内容安全回归：转私密 `95.3161ms`、删除 `93.2795ms` 内从已预热 Feed 消失，safety epoch 分别 `26->28`、`30->32`。
- 详细报告：`docs/superpowers/reports/2026-08-16-feed-hybrid-wp9-singleflight-swr.md`。

## 13. WP10：压测工具、三种基数与 Docker 混合拓扑

### 13.1 计划文件

修改：

- `cmd/loadtest/load_test.yaml`
- `cmd/loadtest/config_test.go`
- `cmd/loadtest/cache_state.go`
- `cmd/loadtest/cache_state_test.go`
- `cmd/loadtest/client_cpu.go`
- `cmd/loadtest/client_cpu_test.go`
- `cmd/loadtest/comparison.go`
- `cmd/loadtest/comparison_test.go`
- `cmd/loadtest/dataset.go`
- `cmd/loadtest/dataset_test.go`
- `cmd/loadtest/environment_evidence.go`
- `cmd/loadtest/environment_evidence_test.go`
- `cmd/loadtest/local_process_monitor.go`
- `cmd/loadtest/local_process_monitor_test.go`
- `cmd/loadtest/runner.go`
- `cmd/loadtest/monitor.go`
- `cmd/loadtest/monitor_test.go`
- `cmd/loadtest/report.go`
- `cmd/loadtest/report_test.go`
- `cmd/loadtest/runtime_strategy.go`
- `cmd/loadtest/runtime_strategy_test.go`
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
- 正式矩阵默认只跑数据不变的读场景；publish/mixed 在 checkpoint/restore 完成前只允许独立命名空间的单轮 probe，不能生成容量结论。
- manifest 固定 strategy 与 seed；报告固定 Git patch 内容哈希、镜像 ID/本地 exe SHA、资源/地址/工作负载兼容指纹，证据不一致时 comparison fail closed。
- 本地拓扑连续采集项目进程 CPU/RSS，并按 PID+exe+start time 检测替换；停止脚本同样校验三者。

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

### 13.5 完成证据与终审

- `go test ./cmd/loadtest/... ./deploy/compose/... -count=3` 通过。
- `go vet ./cmd/loadtest/... ./deploy/compose/...` 通过。
- `run_feed_matrix.ps1`、start/stop 本地服务脚本 PowerShell 语法检查通过。
- start/stop `-WhatIf` 无容器、进程或状态文件副作用；`git diff --check` 通过，仅有 Windows CRLF 提示。
- 终审 Verdict：Ready；Critical 0，Important 0。三项 Minor（SWR sampled max 命名、显式 Mode/日志字段、Redis 监控自流量说明）已记入 WP11 Spec，不阻塞 WP10。

## 14. WP11：正式阶段压测与交付报告

WP11 分两段执行：先完成不改变数据集的三档纯读 control/treatment 容量矩阵；再实现并验证 mutation checkpoint/restore，随后才运行 publish、90/10、80/20 的 60 秒 × 三轮正式矩阵。checkpoint 至少覆盖 benchmark 作者帖子/Outbox、Redis Feed keys、Kafka drain、Page Cache safety epoch 和恢复前后数据指纹。未完成该门禁时，mutation probe 不计入正式交付结论。

执行进度（更新至 2026-08-27）：

- 已按用户批准的压缩档策略完成约 1,200 与 20,000 读者的 RPC/Gateway Control/Treatment：先用 10 秒 scout 确定容量拐点，再仅对稳定最优档执行 60 秒 × 3。
- distributed：RPC c64 从 6,797.75 提升至目标态 102,075.49 QPS；Gateway c16 从 1,218.99 提升至 1,529.10 QPS，但 P95 从 23.02ms 回退到 26.34ms。
- high：Control RPC c32 为 6,162.47 QPS；Treatment c32 触发 `feed_degraded` 判无效，稳定档 c16 为 9,477.51 QPS。Gateway c8 为 1,402.80 → 1,407.80 QPS，基本持平且 P95 回退。
- 压测日志预设、签名 token、Redis RDB 失败隔离和资源证据已实现；受 Redis `MISCONF` 影响的旧报告不进入结论。
- Cursor 深分页专用数据集、实现与 RPC/Gateway 正式 A/B 已完成，结果按第 19 节交付。
- trial 级 mutation checkpoint/restore 已实现：覆盖 MySQL 帖子/Outbox、精确 Redis Feed key DUMP、Kafka 稳定排空、safety epoch、生成帖子 ID 和恢复范围校验；RPC publish、RPC 90/10、RPC 80/20、Gateway 90/10 短时闭环均 `complete=true`。
- 独立正式矩阵脚本经终审补强后完成短时端到端验收：RPC 90/10 c2 得到 90 读/10 写、零失败/超时，10 个成功发布 ID 全部删除，restore 与脚本末尾 checkpoint verify 均通过；TTL 自然老化与即时恢复严格校验已分离。正式脚本现强制新鲜结果命名空间、预期矩阵 manifest、逐格/总数验收，以及从首次 restore 前开始的 incomplete 审计记录。
- 已在全新 checkpoint `43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f` 上完成 RPC-only 正式矩阵：Control/Treatment 各 9 轮，publish c2、90/10 c4、80/20 c4 均为 10 秒预热 + 60 秒采样 × 3；18/18 完整，失败/超时为 0，所有试后恢复成功。
- Treatment 的完整发布、90/10、80/20 总吞吐分别变化 -4.33%、-5.85%、-12.28%；90/10 读 P95 改善 18.43%。发布耗时增加、低 Fresh 命中与吞吐回退共同出现，但 bundled A/B 不对单个开关做因果归因；正式结论见 `docs/superpowers/reports/2026-08-27-feed-wp11-mutation-rpc-ab.md`。
- 后续 Gateway 性能测试按用户决策取消。已有 Gateway 数据保留为已知入口瓶颈旁证，不进入本次 mutation 严格 A/B，也不再作为 WP11 未完成项。

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
- Gateway 历史对照保留且不与 RPC 10K 目标混为同一口径；2026-08-27 后不再追加 Gateway 性能矩阵。

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
