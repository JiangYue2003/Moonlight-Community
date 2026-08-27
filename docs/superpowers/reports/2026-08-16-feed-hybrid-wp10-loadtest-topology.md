# Feed Hybrid WP10：压测口径与拓扑工具报告

日期：2026-08-16  
状态：完成；终审 Ready（Critical 0 / Important 0）；正式容量矩阵属于 WP11 且已按用户要求暂停

## 1. 交付范围

- `hot`、`distributed`、`high` 三种读者基数，对应 20、1,200、20,000 名读者。
- `cold`、`l1-warm`、`l2-warm`、`expire-together` 四种显式 Page Cache 状态，以及不干预的 `natural`。
- 报告记录 cardinality、reader count、cache state、topology、Feature Flags、pprof 目录。
- 报告增加 Page Cache source/outcome、Fresh ratio、SWR queue/active/pending、Redis safety epoch。
- 压测客户端进程 CPU 计时；按逻辑 CPU 数归一化，达到默认 90% 时写入 `client_cpu_saturated`，该轮不能作为服务端容量结论。
- comparison 按 cardinality/cache state/topology 分组，同一命名空间混入两种 topology 时 fail closed。
- `run_feed_matrix.ps1` 的正式读矩阵固定 c32/c64/c128/c256、10 秒预热、60 秒采样、三轮；同 manifest 用独立 `ReportRunId`，避免 control/treatment 覆盖。
- 写/混合场景在 trial checkpoint/restore 完成前只允许单轮、单并发 `noncomparable-mutation-probe`，不会混入正式容量命名空间。
- Docker 镜像源回退脚本：中间件 `--no-build --pull never`，项目服务本地构建并隐藏运行；状态文件记录精确 PID、start time、exe、端口、配置和日志，启停均校验进程身份。
- 报告保存 Git patch 内容哈希、manifest/seed、Docker image ID 或本地 exe SHA、资源/地址/工作负载兼容指纹；comparison 遇到不一致证据 fail closed。
- 纯读窗口要求 safety epoch 起止一致；外部敏感 mutation 导致 epoch 变化时报告以 `redis_safety_epoch_changed` fail closed，写/混合场景不应用该门禁。

## 2. 缓存状态控制

每个非 natural trial 先递增开发 Redis 的 `feed:content:safety:epoch`，等待 1.1 秒跨过 epoch L1 TTL；warm 状态再按受控并发逐读者读取第一页。`l2-warm` 额外等待 1.1 秒，`expire-together` 额外等待 5.2 秒。L1、L2 与 expire-together 都校验 warm + 监控边界预算；集中失效最晚 12 秒进入测量，避免早期 key 已离开最短 13 秒 stale 保留窗。

测量前的普通 warmup 在显式 cache-state 准备之前执行，因此最终进入测量窗口的状态由 cache-state 决定，不受上一轮残留状态命名污染。

## 3. 拓扑隔离

- `compose-full`：策略与开关通过 `docker inspect zg-knowpost` 验证，Prometheus 通过容器内地址采集。
- `compose-middleware-local-services`：策略与开关通过 `.tmp/feed-local/state.json` 验证。当前 Windows 排除范围覆盖 6064/6066 与 9102–9107，本地 pprof 映射到 16064/16066、KnowPost Prometheus 映射到 `127.0.0.1:19104`；脚本执行 `netsh` 检查后才启动。Docker 采样覆盖中间件，本地项目进程按 PID+exe+start time 连续采集 CPU/RSS。
- comparison 遇到同一报告命名空间的 topology 不一致会直接报错，不计算提升百分比。
- Docker 镜像构建/下载问题在脚本和文档中明确归类为 registry/mirror source 问题，不归因于网络；脚本不执行 pull、不清缓存/镜像/卷。

## 4. 验证证据

已通过：

```text
go test ./cmd/loadtest/... ./deploy/compose/... -count=1
go vet ./cmd/loadtest/... ./deploy/compose/...
git diff --check
./scripts/start-feed-local-services.ps1 -WhatIf
./scripts/stop-feed-local-services.ps1 -WhatIf
```

dry-run 后确认 `.tmp/feed-local/state.json` 不存在、本地 Feed 服务进程数为 0。

短功能压测报告：

```text
results/feed-loadtest/feed-wp10-functional-20260816/
  hybrid-rpc-hot-read-c32/trial-1/report.json
  hybrid-rpc-hot-read-c32/trial-1/report.md
  comparison.json
  comparison.csv
  comparison.md
```

口径：`hot + l1-warm + compose-full + Hybrid RPC + c32 + 1s warmup + 5s sample + 1 trial`。

| 指标 | 结果 |
|---|---:|
| 成功/总请求 | 162,586 / 162,586 |
| 成功 QPS | 32,513.61 |
| P95 | 1.618ms |
| P99 | 2.173ms |
| L1+L2 Fresh ratio | 99.82% |
| Cold compute | 2 |
| 客户端归一化 CPU | 13.67% |
| KnowPost CPU 峰值 | 404.40% |
| Kafka lag 峰值 | 0 |
| 报告完整 | true |

该短样本只验证工具、指标和优化热路径，不作为单机容量结论。正式结论必须等待 WP11 的 60 秒 × 三轮矩阵。

## 5. 已知环境约束

Windows 当前 `CGO_ENABLED=0`，`go test -race` 会明确失败为 `-race requires cgo`；本报告不声称 Race 已通过。正式交付继续用并发定向测试和代码审查补强，并保留该环境限制。

## 6. 终审结论

- Verdict：Ready。
- Critical：0。
- Important：0。
- 最新稳定态验证：`go test ./cmd/loadtest/... ./deploy/compose/... -count=3`、`go vet`、三份 PowerShell 语法、start/stop `-WhatIf` 与 `git diff --check` 均通过；Git 只提示 Windows CRLF 行尾。
- 正式三档 control/treatment 读矩阵及 publish/mixed checkpoint/restore 不属于 WP10 完成证据，已完整固化在设计 Spec §18，后续明确继续时才执行。
