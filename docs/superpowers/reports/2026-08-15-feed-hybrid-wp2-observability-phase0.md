# Feed Hybrid WP2：可观测性与 Phase 0 正式基线

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）

## 1. 结论

WP2 已为 Hybrid 个性化 Feed 冷路径建立固定低基数指标，并由压测工具按运行前后累计值计算 delta。Observer 开启后的 c128 三轮中位数相对关闭组：

- 成功 QPS 回退 `1.26%`。
- P95 回退 `1.80%`。
- 均低于 `5%` 仪表开销门禁，且六轮均零失败。

严格复跑确认当前分散读者、冷计算 Hybrid RPC 单机稳定工作档为 `c128`；`c256` 是吞吐增幅首次低于 10% 且 P95 明显上升的拐点，吞吐中位数 `2,793.47 QPS`；`c512` 进入平台区，吞吐仅为 `2,801.37 QPS`，P95 从 `123.603ms` 放大到 `382.092ms`。因此当前冷路径不能靠继续增加并发逼近 10K QPS。

指标同时证明：现有 Route Cache 已把 Counter 调用压至约 `0.0016/request`，但每次读取仍有 `1` 次 Relation RPC、约 `3.01` 次 Redis 调用和 `1` 次完整冷计算。下一阶段应先实施完整 RouteSnapshot/Relation Epoch，而不是继续调整客户端并发。

## 2. 实现范围

KnowPost 新增并接入以下指标：

```text
zhiguang_knowpost_feed_stage_total{stage,outcome}
zhiguang_knowpost_feed_stage_duration_ms{stage,outcome}
zhiguang_knowpost_feed_dependency_calls_total{dependency,operation,outcome}
zhiguang_knowpost_feed_cold_compute_total{outcome}
```

固定 stage：`relation`、`counter`、`route`、`inbox`、`bigv_pipeline`、`merge_dedup`、`hydrate`、`total`。固定 outcome：`success`、`error`、`timeout`、`canceled`、`unknown`。指标不包含 userID、postID、runID 或 cache key。

实现特性：

- `Feed.Observability.Enabled=false` 默认关闭；关闭时 ServiceContext 注入 `nil` observer。
- Observer 错误或 panic 不影响 Feed 请求。
- Relation、Counter、Redis、MySQL 调用按固定 dependency/operation 计数。
- Prometheus 从 KnowPost 容器内部 `127.0.0.1:9104/metrics` 采集，不额外向宿主机暴露端口。
- 压测报告强制要求初末 Prometheus 与 Docker state 抓取成功，并校验 `process_start_time_seconds`、容器 ID、restart count 与 health；进程或容器跨代、counter reset 时报告标记 incomplete。
- stage/dependency 按 outcome 分栏保留；任意非成功 outcome 会标记 `degraded` 并使报告 incomplete，避免 Redis 降级为空仍被算作成功证据。
- 第一页正式读压测要求 Feed 非空，且每条内容的 creator 必须属于 manifest 作者集合；深分页仍允许合法空页。
- 支持 `-warmup`、`-trials` 和 `-report-run-id`；每轮单独落盘，comparison 要求每档具有准确、唯一且完整的预期轮次，遇到证据缺口即停止容量外推。
- 单轮和总表均输出 P90/P95/P99/Max；`compare` 明确采用 `-report-run-id` 指定的结果命名空间。
- `run` 阶段关闭 load generator 的 go-zero 日志 writer，防止大于 500ms 的慢调用格式化并打印完整 Feed 响应；请求失败仍由 recorder 统计。

## 3. 环境与方法

- 拓扑：Compose 全栈，单个 KnowPost、Relation、Counter、Redis、MySQL、Kafka、etcd。
- 策略/入口/场景：`hybrid / rpc / distributed-read`。
- 数据集：`feedbench-20260814c`，报告使用独立 result namespace，不修改 manifest。
- 正式严格结果 ID：`feed-wp2-phase0-strict-20260815`。
- 并发：`32 / 64 / 128 / 256`，因 c256 尚未封住拐点追加 `512`。
- 每档：10 秒不计入统计的预热，60 秒采样，重复 3 次。
- 请求：page=1、size=20，成功延迟单独统计；SLA 仅为参考线，不作为通过/失败条件。
- Observer 开销 A/B：同一镜像、同一数据、同一 c128，5 秒预热 + 20 秒采样，重复 3 次；唯一变量为 `FEED_OBSERVABILITY`。

Docker 使用本地已有层执行 `docker compose build --pull=false knowpost`，构建和启动成功。后续如果镜像获取或构建变慢，按项目约定归类为 Docker registry/mirror 问题，不归因于网络连通性；可只保留 Kafka/MySQL/Redis/etcd 等现有中间件容器并在本地运行项目服务，但切换拓扑后必须重建基线。

## 4. Observer 开销门禁

| Observer | QPS 中位数 | QPS min-max | P95 中位数 | P95 min-max | 失败 |
|---|---:|---:|---:|---:|---:|
| off | 2,636.07 | 2,635.39–2,640.07 | 62.935ms | 62.709–63.564ms | 0 |
| on | 2,602.77 | 2,600.83–2,630.54 | 64.066ms | 62.454–64.123ms | 0 |

结果：QPS 回退 `1.26%`，P95 回退 `1.80%`，WP2 `<=5%` 门禁通过。

证据：

- `results/feed-loadtest/feed-wp2-observer-off-20260815/comparison.md`
- `results/feed-loadtest/feed-wp2-observer-on-20260815/comparison.md`

关闭组预期缺少 Feed metrics，因此单轮报告标记 incomplete；该标记只表达“不可用于分阶段诊断”，不表示请求失败。

## 5. Phase 0 容量曲线

| c | QPS median | QPS min-max | spread | P95 median | P95 min-max | P99 代表范围 | 失败 |
|---:|---:|---:|---:|---:|---:|---:|---:|
| 32 | 1,945.29 | 1,936.89–1,971.36 | 1.77% | 20.847ms | 20.547–21.305ms | 23.044–24.874ms | 0 |
| 64 | 2,327.03 | 2,318.01–2,333.48 | 0.66% | 35.336ms | 35.302–35.703ms | 39.950–40.730ms | 0 |
| 128 | 2,590.31 | 2,578.12–2,609.04 | 1.19% | 64.343ms | 63.659–65.004ms | 72.582–74.692ms | 0 |
| 256 | 2,793.47 | 2,787.75–2,808.98 | 0.76% | 123.603ms | 121.805–125.255ms | 144.012–148.977ms | 0 |
| 512 | 2,801.37 | 2,779.88–2,805.25 | 0.91% | 382.092ms | 378.928–383.916ms | 536.940–544.774ms | 0 |

容量判定：

- 报告口径下峰值无错误成功 QPS：`2,801.37`（c512 三轮中位数），但该点已处于高延迟平台区，不是推荐工作点。
- stable concurrency：`128`。
- knee concurrency：`256`。
- c512 吞吐相对 c256 仅上升约 `0.28%`，P95 上升约 `209.1%`。
- c512 无日志污染复验仍得到同一结论，排除压测客户端控制台输出造成的假拐点。
- 15/15 个正式轮次均 `complete=true`、零请求失败、无 degraded outcome、无初末采样缺失、无进程/容器跨代。

总表：`results/feed-loadtest/feed-wp2-phase0-strict-20260815/comparison.md`。

## 6. 每请求成本与阶段耗时

c128 trial 2（156,623 次成功读取）代表值：

| 指标 | 每成功读取 |
|---|---:|
| Relation `list_followings` | 1.0000 |
| Counter `batch_follower_counts` | 0.0015 |
| Redis Inbox read | 1.0000 |
| Redis BigV pipeline | 1.0000 |
| Redis FeedItem MGet | 1.0000 |
| MySQL FeedItem DB | 0.0099 |
| Redis FeedItem cache write | 0.0099 |
| Cold compute | 1.0000 |

Redis dependency 汇总约 `3.0099/request`。Counter Route Cache 已有效，但 Relation 和三段 Redis 冷路径仍按请求执行。

| stage | mean |
|---|---:|
| relation | 22.431ms |
| hydrate | 9.533ms |
| inbox | 7.730ms |
| bigv_pipeline | 6.898ms |
| counter（仅 miss） | 8.478ms |
| route | 0.030ms |
| merge/dedup | 0.029ms |
| total | 46.701ms |

这些 duration 存在嵌套，不能直接相加当作总耗时；`total` 是请求级权威值。

## 7. Profile 结论

c128 trial 1 同窗采集 KnowPost 与 Relation 30 秒 CPU Profile，并保存 heap、allocs、block、mutex 快照：

`results/feed-loadtest/feed-wp2-phase0-enabled-20260815/profiles/c128-trial1/`

Profile 在与严格复跑相同的代码、数据、镜像和拓扑上采集；严格复跑只加强报告证据门禁，没有改变业务读路径，因此保留该同窗 Profile 作为热点排序证据。

KnowPost：

- `GetUserFeed` 累计覆盖约 `50.7%` CPU samples。
- `loadFeedItems` 约 `24.2%`，Redis `MGet`、reply 读取和 `encoding/json.Unmarshal` 是主要 hydrate 成本。
- `FeedReadSnapshot.GetFeed` 约 `17.1%`，Inbox/BigV Redis IO 为主。
- 累计 alloc profile 中 `loadFeedItems`、JSON 解码、`Deduplicate`、MGet reply 和 BigV pipeline 是主要分配源；该快照是进程累计值，只用于排序热点，不能直接换算本轮 bytes/request。

Relation：

- `ListFollowing -> PageActive` 累计覆盖约 `45%` CPU samples，当前分散读者路径频繁进入 MySQL/sqlx。
- sqlx struct 字段展开与映射、`summarizeUserIDs` 及字符串切分是主要分配源。
- 这解释了为何移除热路径 Relation RPC 的收益不仅是一次网络 RTT，还会消除 MySQL 查询、反射映射和响应分配。

block/mutex profile 文件可读取但没有显著样本，当前证据不支持把锁竞争列为 P0 瓶颈。

## 8. 验证

已通过：

```text
go test ./services/knowpost/rpc/internal/feed
go test ./services/knowpost/rpc/internal/logic/knowpost
go test ./services/knowpost/rpc/internal/svc
go test ./services/knowpost/cmd/knowpost
go test ./services/relation/cmd/relation
go test ./pkg/debughttp
go test ./cmd/loadtest
go test ./deploy/compose
go vet 相关 WP2 包
git diff --check
```

另执行严格容器校验 `feed-wp2-strict-validation-20260815`：c7、10 秒、10,579 次成功读取，报告完整，进程启动时间及 KnowPost 容器 ID 均只有一个唯一值。

复审额外指出 trial 集合除“数量正确且唯一”外还应精确为 `1..N`；已补测试与门禁，`{2,3,4}` 不再能冒充三轮完整证据。

Windows 当前 Go 环境为 `CGO_ENABLED=0` 且没有 C compiler，`go test -race` 无法执行；这属于工具链限制，不以普通测试替代 race 结论。

## 9. 下一步门禁

进入 WP3/WP4：FeedEpochStore 与 Relation Epoch 消费者，为 WP5 完整 RouteSnapshot 提供一致性版本。

WP5 的量化门禁保持不变：

- 热路由 `Relation RPC / Feed request` 至少下降 80%。
- 相同拓扑和负载下 QPS 至少提升 20%，或单位成功请求 CPU 至少下降 20%。
- Follow/Unfollow、事件重复/乱序/延迟与 TTL 兜底测试通过。

在 WP5 通过前，不用页面缓存掩盖 Relation 冷路径问题；也不基于当前证据实施 sync.Pool、动态并发或自定义协议。
