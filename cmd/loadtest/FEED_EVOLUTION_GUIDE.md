# Feed 混合架构演进压测指南

本指南用于验证 RouteSnapshot、Redis Combined Pipeline、整页 L1/L2 Cache、Singleflight 和 SWR。容量结论只在相同 Git 状态、数据集、读者基数、缓存状态、运行拓扑和资源限制下可比较。

## 1. 固定测试口径

正式矩阵使用：

- 读者基数：`hot`（20）、`distributed`（约 1,200）、`high`（20,000）。
- 完整探索矩阵并发为 32、64、128、256；长测成本受限时允许先对候选档做 10 秒 scout，再只对吞吐/延迟拐点与相邻一档进行正式验证。
- scout 只用于选档。正式容量点仍必须预热后采样 60 秒、独立重复三轮；短测数字不得进入容量结论。
- Feed：Hybrid、RPC、第一页、每页 20 条；Gateway 作为单独入口对照。
- 缓存状态必须显式记录，不能把冷读与热读混为一个容量数字。
- SLA 数值是参考线；错误、超时、正确性失败、进程重启、指标缺失或压测客户端 CPU 饱和会令报告不可用于容量结论。

`10,000 QPS` 只对应 `hot + l1-warm + Hybrid RPC + page1/size20` 的阶段目标，不代表 20,000 个随机冷读者也能达到同一吞吐。

正式本地混合拓扑必须使用 `benchmark-error-only` 日志预设：保留 error，关闭 go-zero stat、RPC Stat middleware、Gateway access log 与 SQL statement info。它不是完全无日志；真实错误仍写入 stderr，并会用于发现 Redis/Kafka/依赖故障。

## 2. 创建三种独立数据集

每种 cardinality、每种 Feed strategy 必须使用独立 Run ID 和 manifest。`setup` 会先核对正在运行的真实策略，并把 strategy 与随机种子写入 manifest；后续 `run` 遇到不一致会拒绝执行：

```powershell
$hot = "feed-hot-$(Get-Date -Format yyyyMMdd-HHmmss)"
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase setup `
  -run-id $hot -reader-cardinality hot -strategy hybrid -feature-preset treatment

$distributed = "feed-distributed-$(Get-Date -Format yyyyMMdd-HHmmss)"
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase setup `
  -run-id $distributed -reader-cardinality distributed -strategy hybrid -feature-preset treatment

$high = "feed-high-$(Get-Date -Format yyyyMMdd-HHmmss)"
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase setup `
  -run-id $high -reader-cardinality high -strategy hybrid -feature-preset treatment
```

`run` 会核对 CLI cardinality、真实读者数、随机种子和 Feed strategy。Push、Pull、Hybrid 不能复用同一个数据集，因为初始 Feed ZSET 的物化布局不同。

## 3. 缓存状态

`-cache-state` 支持：

| 状态 | 测量前动作 | 用途 |
|---|---|---|
| `natural` | 不干预 | 观察自然流量状态 |
| `cold` | 递增全局 safety epoch，等待 L1 epoch 收敛 | 强制 Page Cache miss |
| `l1-warm` | cold 后逐读者预热一次 | 热点上限与 10K 目标 |
| `l2-warm` | 预热后等待 L1 Fresh 过期 | 单独测 L2 命中成本 |
| `expire-together` | 预热后等待 5.2 秒 | 验证集中失效、SWR 与刷新队列 |

epoch 递增只用于完全开发环境的可审计测试控制。它不会扫描或删除 Redis；短暂不可见/不一致是 Feed 已批准的一致性取舍。

Warm 状态会 fail closed：`l1-warm` 只允许至多 20 个工作集读者，并要求 warm + 指标边界不超过 400ms；`l2-warm` 要求不超过 1.7s；`expire-together` 要求 warm + 5.2s 等待 + 指标边界不超过 12s，避免早期 key 已经离开 13s 的最短 stale 保留窗。矩阵的 `-CacheState auto` 在 treatment 下映射为 hot=`l1-warm`、distributed=`l2-warm`、high=`cold`；control 因 Page Cache 关闭固定映射为 `cold`。Go 工具也拒绝把 control 标成 L1/L2/expire。

## 4. 正式单场景与矩阵

热点三轮示例：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase run -run-id $hot -reader-cardinality hot `
  -topology compose-full -cache-state l1-warm `
  -feature-preset treatment `
  -strategy hybrid -entry rpc -scenario hot-read `
  -concurrency 128 -requests 0 -duration 60s -warmup 10s -trials 3 `
  -pprof-dir results/feed-loadtest/$hot/pprof
```

固定矩阵脚本默认 Hybrid、c32/c64/c128/c256、10 秒预热、60 秒采样、三轮：

```powershell
./cmd/loadtest/run_feed_matrix.ps1 `
  -RunId $hot -ReportRunId "$hot-treatment" -ReaderCardinality hot `
  -CacheState auto -FeaturePreset treatment -Topology compose-full `
  -PprofDir "results/feed-loadtest/$hot/pprof"
```

`run_feed_matrix.ps1` 只运行不会改变数据集的正式读矩阵（RPC、Gateway、hot/distributed/deep-page 与读突发/steady），从而保证不同并发和三轮的 MySQL/Redis Feed 初态一致。脚本拒绝一个 Run ID 同时切换多种 Feed strategy。

写入和读写混合已迁移到独立的 `run_feed_mutation_matrix.ps1`。所有 publish/mixed 运行强制使用 manifest-owned checkpoint 和精确 RunID 确认；每次预热后、测量后都恢复 MySQL 帖子/Outbox 与 Redis Feed key，并要求 Kafka lag/log-end 连续稳定、safety epoch 单调递增、已发布帖子全部进入删除集合。完整命令和判读见 [FEED_MUTATION_GUIDE.md](./FEED_MUTATION_GUIDE.md)。短时 smoke 仍只证明工具闭环，不得冒充 60 秒 × 三轮容量结果。

Control 与 treatment 必须复用同一个 `-RunId`/manifest，并使用不同 `-ReportRunId`，否则同场景路径会覆盖旧报告。严格 A/B 使用两边都为 `cold`；treatment 的 `auto` 热缓存卡用于验证目标容量，不能与 control cold 直接宣称为同缓存状态提升。`$hot-control`、`$hot-treatment-cold`、`$hot-treatment-auto` 分别汇总后再生成最终演进报告。

压缩档示例（先 scout，再只跑已选 c64）：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase run `
  -run-id $distributed -report-run-id "$distributed-treatment-scout" `
  -reader-cardinality distributed -topology compose-middleware-local-services `
  -cache-state l2-warm -feature-preset treatment -strategy hybrid `
  -entry rpc -scenario distributed-read -concurrencies 32,64,128 `
  -requests 0 -duration 10s -warmup 2s -trials 1

go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase run `
  -run-id $distributed -report-run-id "$distributed-treatment-formal" `
  -reader-cardinality distributed -topology compose-middleware-local-services `
  -cache-state l2-warm -feature-preset treatment -strategy hybrid `
  -entry rpc -scenario distributed-read -concurrency 64 `
  -requests 0 -duration 60s -warmup 10s -trials 3
```

若任一正式首轮出现 `feed_degraded`、客户端饱和或其他 hard missing，应停止该档剩余 trial，保留无效报告作为过载边界，再降到相邻低档做一次短有效性探测。

## 5. Profile

当前 Windows 排除范围覆盖 6064/6066 与 9102–9107，因此 Compose 和本地混合拓扑都使用 KnowPost `16064`、Relation `16066`；本地 KnowPost Prometheus 使用 `19104`。启动脚本会读取 `netsh interface ipv4 show excludedportrange protocol=tcp` 并 fail closed。Profile 必须与对应负载窗口重叠，并保存到传入报告的 `-pprof-dir`：

```powershell
go tool pprof -proto -output results/feed-loadtest/$hot/pprof/knowpost-c128.pb.gz `
  "http://127.0.0.1:16064/debug/pprof/profile?seconds=60"
go tool pprof -proto -output results/feed-loadtest/$hot/pprof/relation-c128.pb.gz `
  "http://127.0.0.1:16066/debug/pprof/profile?seconds=60"
```

## 6. Docker 镜像源降级拓扑

Docker 镜像构建或下载缓慢/失败在本项目中按 **registry/mirror source 问题** 处理，不归因于网络连通性。不要反复 pull，不清空 BuildKit、镜像、容器、卷，也不重建 MySQL、Redis、Kafka 或 etcd。

先检查 dry-run：

```powershell
./scripts/start-feed-local-services.ps1 -WhatIf
./scripts/stop-feed-local-services.ps1 -WhatIf
```

再启动已存在镜像的中间件与本地项目服务：

```powershell
./scripts/start-feed-local-services.ps1 -Strategy hybrid -FeaturePreset treatment -LoggingPreset benchmark-error-only
```

脚本使用 Compose `--no-build --pull never`，只保留/启动已存在的 etcd、ZooKeeper、Kafka、Canal、Elasticsearch；Gateway、User/Storage、Counter、Relation、KnowPost、Search 由本地编译后的隐藏进程提供。生成的配置把 etcd 指向 `127.0.0.1:12379`、Kafka 指向 `127.0.0.1:9092`，并把 RPC 注册为宿主机可达地址。

状态写入 `.tmp/feed-local/state.json`，包含 PID、进程启动毫秒、可执行文件、监听端口、配置、stdout/stderr、策略和 Feature Flags。启动成功要求新进程仍存活且拥有全部预期端口；停止脚本只在 PID、启动时间与可执行文件路径都匹配时停止该进程，不停止中间件：

```powershell
./scripts/stop-feed-local-services.ps1
```

本地拓扑运行时必须传：

```text
-topology compose-middleware-local-services
```

压测器会从状态文件验证实际策略和 Feature Flags，直接读取宿主机 KnowPost `127.0.0.1:19104` Prometheus，并连续采集本地 Gateway/KnowPost/Relation/Counter/User/Search 的 CPU、RSS、PID、exe 与启动时间。比较器在本地拓扑使用 `process:knowpost/cpu_percent`，不会把 KnowPost CPU 默认为 0。它会拒绝同一报告命名空间混合 topology；切换拓扑后必须重新执行 smoke 和 Phase 0 baseline。

### 6.1 Windows Redis benchmark persistence

若 error-only 日志出现 `MISCONF Redis is configured to save RDB snapshots, but it's currently unable to persist to disk`，这不是网络或 Docker 镜像源问题。该错误会让 Redis 停止写入并污染 Feed cache-state，相关报告必须作废。

完全开发环境可在压测窗口临时关闭 RDB schedule 与失败写保护，但必须先记录原配置，且不得把该设置带入生产：

```powershell
redis-cli CONFIG GET save
redis-cli CONFIG GET stop-writes-on-bgsave-error
redis-cli CONFIG SET save ""
redis-cli CONFIG SET stop-writes-on-bgsave-error no
```

本项目 2026-08-17 的原值证据保存在 `.tmp/feed-local/redis-benchmark-original.json`。压测报告必须记录 persistence mode；全部性能工作结束后恢复原值并做 SET/GET/DEL 冒烟。若无法证明本轮 persistence mode 与 Redis identity 一致，Control/Treatment 不可比较。

## 7. 报告判读

每个 `report.json/report.md` 包含：

- QPS、P50/P90/P95/P99/Max、错误和超时；
- Relation/Counter/Redis 调用及每成功读调用数、cold compute；
- Page Cache L1/L2 Fresh、Stale、Miss、Fresh ratio；
- SWR queue/active/pending 峰值；
- Redis 本轮命令/网络/命中 delta、eviction/rejection、safety epoch，Kafka lag，MySQL，容器或本地进程 CPU/RSS，以及压测客户端 CPU；
- Git commit、未提交 patch 内容哈希、manifest/seed、被测镜像 ID或本地 exe SHA256、资源限制、服务地址、工作负载参数、cardinality、cache state、topology、Feature Flags 和 pprof 目录。

若 `missing_metrics` 包含 `client_cpu_saturated`，说明负载发生器归一化 CPU 已达到配置阈值（默认 90%），该轮只表示客户端上限，不是服务端单机上限。纯读采样期间若 safety epoch 被外部敏感 mutation/consumer 改变，报告会写入 `redis_safety_epoch_changed` 并失效，避免把整个 Page Cache namespace 的中途切换混入 cold/warm 容量数字。

汇总同一命名空间：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase compare `
  -run-id $hot -report-run-id "$hot-treatment"
```

最终结论至少同时给出热点、分散、高基数三张卡片，不能只报告最高 QPS。

2026-08-17 distributed/high 压缩档正式结果与判读见 `docs/superpowers/reports/2026-08-17-feed-hybrid-wp11-capacity.md`。

## 8. Cursor 深分页专用 A/B

Cursor A/B 不复用 hot/distributed/high manifest。先用 `cursor-ab` 预设启动服务：混合架构、RouteSnapshot、关系/安全 epoch、Combined Pipeline 和 Cursor 全部开启，但 PageCache 强制关闭。Docker 镜像源异常时继续使用既有中间件加本地项目服务，不重新构建业务镜像：

```powershell
./scripts/start-feed-local-services.ps1 `
  -Strategy hybrid -FeaturePreset cursor-ab -LoggingPreset benchmark-error-only
```

创建独立数据集。该阶段精确发布 20 个普通作者×50 条、5 个大 V×100 条，等待 Kafka lag 归零，再验证每个读者 1,000 条 Inbox、每个大 V 100 条 Outbox、1,500 条 MySQL 已发布公开记录、同秒 tie 和跨来源重复，最后才把 manifest 标成 `ready=true`：

```powershell
$cursor = "feed-cursor-deep-v1"
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase setup-cursor-deep -run-id $cursor `
  -reader-cardinality cursor-deep `
  -topology compose-middleware-local-services `
  -cache-state cold -feature-preset cursor-ab -strategy hybrid
```

若发布或 Kafka/Redis/MySQL 校验中断，保留 manifest，使用相同参数续跑；工具只补齐已明确缺少的作者配额。若出现“RPC 已提交但客户端未收到响应”造成的未知帖子，Redis 成员集合门禁会失败，此时不能强行把 manifest 改成 Ready，应清理该独立数据集后重建：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase seed-cursor-deep -run-id $cursor `
  -reader-cardinality cursor-deep `
  -topology compose-middleware-local-services `
  -cache-state cold -feature-preset cursor-ab -strategy hybrid
```

先跑 c16/c32/c64、10 秒单轮 scout：

```powershell
./cmd/loadtest/run_feed_cursor_ab.ps1 -RunId $cursor -Mode scout
```

根据完整报告选稳定最优档和一个相邻档，再跑 60 秒、三轮正式点。2026-08-17 的本机实测选择 c16 为稳定档、c32 为相邻过载边界：

```powershell
./cmd/loadtest/run_feed_cursor_ab.ps1 -RunId $cursor -Mode formal `
  -Concurrencies @(16,32)
```

脚本分别执行 RPC/Gateway 的 page1、page5/20/50、对应 Cursor 目标页、`sequential-page-50` 页码累计 Control、`sequential-50` Cursor 累计 Treatment 和 `same-second-cursor`。Gateway 默认使用 `signed` 压测 Token：仅为 manifest 中已有的真实用户 ID 签发短期凭证，避免并发 Login 的 Redis/登录日志成本污染 Feed Gateway 数据；报告会记录认证模式。固定目标页的前序 Cursor 链在采样前生成，不计入目标页耗时；两种 sequential 场景都计入第 1–50 页的所有真实请求。固定目标页和累计滚动结果不得混为同一个性能数字。

Prometheus 资源样本同时保存 Go Runtime 的 GC 次数/暂停总量、累计分配字节和堆存量。`go_memstats_heap_alloc_bytes` 是可下降的 Gauge，不得按单调 Counter 做重启判断；正式报告用 `alloc_bytes_total`、GC count/sum 的初末 delta 计算每请求分配和 GC 成本。

2026-08-17 正式结果见 `docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md`。page20/page50 和连续滚动门槛通过；Gateway page5 在同秒大组与宿主机调度下收益不确定，因此结论是保持开关可回滚、先灰度，而不是声称所有页型都默认更快。
