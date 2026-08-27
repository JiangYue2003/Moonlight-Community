# Feed Hybrid 读取优化复测报告

日期：2026-08-14

## 结论

本轮在不改变推、拉、混合业务路由语义的前提下，完成了四项读路径优化，并在代码审查后补齐三项稳定性保护：

1. 每次读取仍从 Relation 获取最新关注列表；只有关注列表与缓存副本完全一致时，才复用大 V 分类并跳过 Counter RPC。
2. 真实 RedisAdapter 将多个大 V Outbox 的 `ZREVRANGE WITHSCORES` 放入 Pipeline；每批最多 128 个作者、逐批归并到固定大小 Top-N 堆，不支持批量接口的实现继续使用原有 16 并发降级。
3. 删除每次成功分类产生的请求级日志；Relation、Counter、Redis 失败日志保留。
4. Counter 冷缓存 miss 使用按“用户 + 完整关注快照”的 singleflight 合并；共享查询使用独立的 2 秒 Context，每个等待者仍可按自身 Context 退出。Counter 故障时继续按全普通作者降级，但不缓存错误降级结果。

审查修复后的最终 Hybrid RPC c128 30 秒稳态成功吞吐从 1,590.01 QPS 提升到 2,396.03 QPS，提升 50.69%；P95 从 103.53 ms 降低到 68.55 ms，降低 33.79%。全部 71,968 次请求成功。

## 同口径结果

### RPC c128，30 秒稳态

| 指标 | 优化前 | 优化后 | 变化 |
|---|---:|---:|---:|
| 成功 QPS | 1,590.01 | 2,396.03 | +50.69% |
| 平均延迟 | 80.43 ms | 53.40 ms | -33.61% |
| P95 | 103.53 ms | 68.55 ms | -33.79% |
| P99 | 115.58 ms | 77.85 ms | -32.64% |
| 失败 | 0 | 0 | 不变 |
| Redis 峰值 ops/s | 57,294 | 19,981 | -65.13% |
| Counter CPU 峰值 | 90.61% | 6.63% | -92.68% |
| Relation CPU 峰值 | 164.32% | 237.50% | +44.54% |
| KnowPost CPU 峰值 | 274.96% | 298.59% | +8.59% |

Redis 与 Counter 压力显著下降，同时系统处理了更多请求。Relation CPU 上升说明瓶颈已经从 Counter/Redis 重复分类转移到每次请求仍需执行的关注列表 RPC。

### RPC c256，3 秒短窗口

| 指标 | 优化前 | 优化后 | 变化 |
|---|---:|---:|---:|
| 成功 QPS | 1,599.55 | 2,408.24 | +50.56% |
| P95 | 193.74 ms | 142.94 ms | -26.22% |
| 失败 | 0 | 0 | 不变 |

c256 相比优化后 c128 稳态吞吐提升已经很小，但延迟接近翻倍，因此 c256 仍是性能拐点，不建议作为稳定运行档。

### Gateway 分散首页，3 秒短窗口

| 并发 | 优化前 QPS / P95 | 优化后 QPS / P95 | QPS 变化 |
|---:|---:|---:|---:|
| c32 | 953.19 / 42.97 ms | 1,442.75 / 28.44 ms | +51.36% |
| c128 | 869.06 / 180.42 ms | 1,599.85 / 111.20 ms | +84.09% |

后端排队降低后，收益能够穿透 Gateway；但 c128 Gateway 相比 RPC c128 稳态仍有约 32% 的吞吐损耗，HTTP、鉴权和转发仍是下一条独立优化链路。

## 正确性和降级验证

- Hybrid 冒烟通过：普通作者 fanout 为 100/100，大 V 不 fanout，阅读者可见新内容，Kafka 事件与 lag 检查通过。
- 关注列表不变时只调用一次 Counter；列表变化时立即重新分类，并且旧作者不再通过 `AllowsCreator`。
- Counter 故障降级不会污染缓存；恢复后的下一次请求会重新分类。
- 同一用户同一关注快照的并发冷 miss 只触发一次 Counter RPC；不同关注快照不会被错误合并。
- 首个请求取消不会取消共享 Counter 查询或污染仍有效等待者的结果；已取消等待者可立即退出。
- Redis Pipeline 能返回所有请求范围。
- Pipeline 每批最多 128 个作者，避免大关注列表一次性创建并物化全部 Redis 命令。
- 单个 Outbox 发生 `WRONGTYPE` 时，仅该来源失败，其他来源结果继续参与 Feed 合并。
- 不支持批量接口的 RedisClient 仍按原有最大 16 并发读取。
- 请求级成功分类日志从一次 30 秒测试约 70,000 行降为 0。

## 验证命令和状态

- `go test ./services/knowpost/rpc/... ./services/knowpost/cmd/knowpost ./services/knowpost/cmd/knowpost/internal/app -count=1`：通过。
- `go vet ./services/knowpost/rpc/internal/feed ./services/knowpost/rpc/internal/logic/knowpost ./services/knowpost/rpc/internal/svc`：通过。
- `git diff --check`：通过，仅有现有 Windows 行尾提示。
- KnowPost 容器：healthy，restart count 0。
- 最终容器二进制 SHA256：`ee5d516d38f0bbf5c9fdb9c5dd8b44c09e2fdf1222e50996bd926eeab0560cf7`。

## 已知限制

- 当前路由分类缓存 TTL 已从 1 分钟收紧到 5 秒。关注列表变化会立即绕过缓存；作者粉丝数跨越大 V 阈值且关注列表不变时，分类仍最多可能陈旧 5 秒。若生产要求零窗口，需要由 Counter 发布阈值变更事件主动失效，或在迁移窗口双写 Inbox/Outbox。
- 本轮仍然每次请求 Relation RPC。新的资源曲线表明它是下一阶段最优先的瓶颈。
- `-race` 未执行：当前 Windows Go 环境为 `CGO_ENABLED=0` 且无 gcc。普通测试、集成测试、vet、容器冒烟和压力测试已通过。
- KnowPost 全目录存在一项本轮之前就有的配置测试矛盾：`knowpost.yaml` 显式设置 `DisableAPI: true`，而 `TestDecodeKnowpostMergedConfig` 仍期望 `false`。本轮未修改该无关语义。
- 结果来自 Windows + Docker Desktop 单机共享开发环境，不能直接外推生产集群容量。

## 证据索引

- 优化前稳态：`results/feed-loadtest/feedbench-20260814c/hybrid-rpc-steady-read-c128/report.json`
- 第一版优化后稳态：`results/feed-loadtest/feedopt-20260814-route-pipeline-nolog/hybrid-rpc-steady-read-c128/report.json`
- 审查修复后最终稳态：`results/feed-loadtest/feedopt-20260814-reviewfix/hybrid-rpc-steady-read-c128/report.json`
- 优化后 RPC c256：`results/feed-loadtest/feedopt-20260814-route-pipeline-nolog/hybrid-rpc-distributed-read-c256/report.json`
- 优化后 Gateway c32/c128：当前目录的 `hybrid-gateway-distributed-read-c32` 和 `hybrid-gateway-distributed-read-c128`。
- 最终冒烟证据：`results/feed-loadtest/feedopt-20260814-reviewfix/smoke-hybrid.json`。

复测 manifest 是原始数据集的只读测量别名，只改变 `run_id` 以隔离报告目录；邮箱所有权仍属于 `feedbench-20260814c`，因此不得使用这些别名执行 cleanup。
