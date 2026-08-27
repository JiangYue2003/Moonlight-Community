# Feed 三策略压测复现指南

> 当前性能演进的正式 20/1,200/20,000 读者、缓存状态、客户端 CPU 门禁和 Docker 镜像源降级流程，以 [FEED_EVOLUTION_GUIDE.md](./FEED_EVOLUTION_GUIDE.md) 为准。本文只描述单策略正式读矩阵；写入/混合使用 [FEED_MUTATION_GUIDE.md](./FEED_MUTATION_GUIDE.md) 和 `run_feed_mutation_matrix.ps1`。

本文对应当前 `cmd/loadtest` 实现。目录中旧版 `README.md`、`QUICKSTART.md` 和 `run_test.*` 使用的是历史 CLI 参数，不适用于本轮三策略矩阵。

## 1. 前置条件

- 从仓库根目录执行命令。
- Docker Compose 文件：`deploy/compose/docker-compose.dev.yml`。
- MySQL 和 Redis 由宿主机提供；etcd、Kafka 和业务服务由 Compose 管理。
- 压测器通过 `127.0.0.1:12379` 的 etcd 发现宿主机可达 RPC 地址。
- 所有 setup/run/reset 操作都会修改开发环境数据。`compare` 只读取已有报告；`snapshot` 只读取 Docker/Kafka 状态。

启动或检查环境：

```powershell
docker compose -f deploy/compose/docker-compose.dev.yml up -d --no-build --pull never --wait --wait-timeout 180
docker compose -f deploy/compose/docker-compose.dev.yml ps
```

若 Docker 构建或镜像下载异常，按 registry/mirror source 问题处理，不归因于网络，也不反复 pull 或清缓存。改用 [FEED_EVOLUTION_GUIDE.md](./FEED_EVOLUTION_GUIDE.md) 中“已有中间件容器 + 本地项目服务”的脚本。

## 2. 创建可复用数据集

每次正式测试使用新的 Run ID：

```powershell
$runId = "feed-hybrid-hot-$(Get-Date -Format yyyyMMdd-HHmmss)"
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase setup `
  -run-id $runId -reader-cardinality hot `
  -strategy hybrid -feature-preset treatment
```

manifest 固定 reader cardinality、随机种子与 Feed strategy。Push/Pull/Hybrid 必须分别 setup，不能在同一个 Run ID 上切换物化布局。manifest 写入：

```text
results/feed-loadtest/manifests/<run-id>.json
```

如果设置 `FEED_LOADTEST_PASSWORD`，后续 Gateway 测试必须继续使用同一个值；未设置时工具使用开发环境默认压测密码。

## 3. 运行完整矩阵

```powershell
$reportRunId = "$runId-treatment-auto"
./cmd/loadtest/run_feed_matrix.ps1 `
  -RunId $runId -ReportRunId $reportRunId `
  -Strategies @("hybrid") `
  -ReaderCardinality hot -CacheState auto `
  -FeaturePreset treatment -Topology compose-full `
  -Concurrencies @(32,64,128,256) `
  -GatewayConcurrencies @(32,64,128,256) `
  -WarmupSeconds 10 -ReadDurationSeconds 60 -Trials 3 `
  -SkipSupplemental
```

脚本要求恰好一个 strategy，并执行：

- 验证真实 strategy、FeaturePreset、manifest strategy/seed/cardinality 与 Kafka lag；
- RPC：hot-read、distributed-read、deep-page；
- Gateway：hot-read、distributed-read，可用 `-SkipGateway` 跳过；
- 可选读突发/steady 补测，可用 `-SkipSupplemental` 跳过；
- `CacheState=auto`：control=`cold`；treatment hot=`l1-warm`、distributed=`l2-warm`、high=`cold`；
- 不执行 reset/seed/publish，因此正式读矩阵各并发与 trial 保持同一 Feed 数据初态；
- 无论成功或失败，`finally` 都将 KnowPost 恢复为 Hybrid。

严格 control/treatment A/B 两边都使用 `-CacheState cold`。treatment `auto` 是额外的目标容量卡，不能与 control cold 直接标成同缓存状态提升。

写入/混合不再由本脚本执行。先创建 checkpoint，再使用独立正式入口；短时示例可把并发/时长/轮次缩小，但仍必须完成恢复：

```powershell
./cmd/loadtest/run_feed_mutation_matrix.ps1 `
  -RunId $runId -ReportRunId "$runId-mutation-formal" `
  -Strategy hybrid -ReaderCardinality hot `
  -FeaturePreset treatment -CacheState cold `
  -Concurrencies @(32) -GatewayConcurrencies @(32) `
  -WarmupSeconds 10 -DurationSeconds 60 -Trials 3
```

正式脚本要求已有 checkpoint，并把 checkpoint ID、基线指纹、每轮发布/删除帖子 ID、Kafka 稳定恢复和最终 restore 写入报告；任一恢复或指标门禁失败都会停止后续 trial。

## 4. 单场景命令

验证当前策略可见性：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase smoke -run-id $runId -strategy hybrid
```

RPC 分散读，c128，固定 30 秒：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase run -run-id $runId -strategy hybrid -entry rpc `
  -feature-preset treatment -cache-state cold `
  -scenario distributed-read -concurrency 128 -requests 0 -duration 30s
```

Gateway 80/20 短时门禁。必须使用独立 `-report-run-id`、已有 checkpoint 和精确 RunID 确认：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase run -run-id $runId -report-run-id "$runId-mutation-smoke" `
  -strategy hybrid -feature-preset treatment -cache-state cold -entry gateway `
  -scenario mixed-80-20 -concurrency 4 -requests 0 -duration 3s -warmup 1s `
  -checkpoint "results/feed-loadtest/checkpoints/$runId.json" `
  -confirm-mutation $runId
```

场景枚举：

```text
hot-read, distributed-read, deep-page, burst-read, steady-read,
publish, burst-publish, mixed-90-10, mixed-80-20
```

`burst-read` 和 `steady-read` 与分散读使用同一目标分布，区别由 `-duration` 控制；本轮补测分别使用 10 秒和 30 秒。`burst-publish` 与完整 `publish` 工作流一致，主要用于在持续读之后单独标记写突发。所有包含发布的场景都会在停止施压后继续等待 Kafka lag 清零，并把恢复耗时和最终 lag 写入报告。

入口枚举：`rpc`、`gateway`。策略标签枚举：`push`、`pull`、`hybrid`；标签必须与 KnowPost 容器实际 `FEED_STRATEGY` 一致。

## 5. 汇总已有报告

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase compare `
  -run-id $runId -report-run-id $reportRunId
```

输出位于：

```text
results/feed-loadtest/<report-run-id>/comparison.json
results/feed-loadtest/<report-run-id>/comparison.csv
results/feed-loadtest/<report-run-id>/comparison.md
```

每个单场景目录还包含 `report.json`、`stages.csv`、`report.md`。

保存当前容器健康、restart count、实际策略和 Kafka offset/lag 快照：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase snapshot -run-id $runId
```

输出为 `results/feed-loadtest/<run-id>/environment-snapshot.json`。单轮报告会采集 Feed Prometheus、Page Cache/SWR、Docker、Kafka、Redis、MySQL 和客户端 CPU；环境快照中的 `missing_data` 仍明确表示它本身不替代负载窗口内的关联日志。

## 6. 精确重置与数据保留

仅删除 manifest 可证明归属的 Inbox、Outbox 和 Feed processing Redis keys：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml -phase reset-feed -run-id $runId
```

`reset-feed` 不扫描 Redis，不删除其他用户数据，也不删除 MySQL 用户、关注关系和帖子。

先执行只读 dry-run。它会等待 Kafka lag=0、全量核对每个 manifest 用户的 ID/email，并运行 SQL 选择和 Redis key 推导，但回滚事务且不发送删除命令：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase cleanup -run-id $runId -confirm-cleanup $runId -cleanup-dry-run
```

预览写入 `cleanup-preview.json`。报告复核完成且预览数量正确后，才执行全量 manifest 范围清理：

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml `
  -phase cleanup -run-id $runId -confirm-cleanup $runId
```

真实 `cleanup` 会在同样的 fail-closed 校验后，于一个 MySQL 事务内删除对应 Outbox、登录日志、关系、帖子和用户，最后删除可精确推导的 Redis keys。报告和 manifest 会保留；Elasticsearch 文档、Kafka 历史和进程内 L1 不在清理范围，清理后应重启应用容器。

## 7. 最终状态检查

```powershell
docker inspect zg-knowpost --format '{{range .Config.Env}}{{println .}}{{end}}' |
  Select-String '^FEED_STRATEGY=hybrid$'

docker exec zg-kafka /opt/bitnami/kafka/bin/kafka-consumer-groups.sh `
  --bootstrap-server kafka:29092 `
  --group feed-fanout-group `
  --describe
```

预期 KnowPost 为 healthy、策略为 Hybrid、Kafka `LAG` 为 0。
