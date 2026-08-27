# Feed WP11 写入/混合负载 RPC 正式 A/B 报告

日期：2026-08-27  
状态：RPC 正式矩阵完成；Gateway 后续性能测试按用户决策取消

## 1. 执行结论

本轮使用同一份全新 checkpoint，完成 Control 与 Treatment 各 9 个 RPC trial：publish c2、90/10 c4、80/20 c4 均为 10 秒预热、60 秒采样、三轮重复。18/18 份报告均 `complete=true`，业务失败和超时为 0，试前、预热后、试后恢复以及 Kafka 排空全部成功。

Treatment 没有提高写入/混合负载的端到端吞吐：

| RPC 场景 | 指标 | Control 中位数 | Treatment 中位数 | 变化 |
|---|---|---:|---:|---:|
| publish c2 | 完整发布 QPS | 6.61 | 6.33 | -4.33% |
| publish c2 | 发布 P95 | 398.19ms | 388.98ms | -2.31% |
| 90/10 c4 | 读 QPS | 58.66 | 55.23 | -5.85% |
| 90/10 c4 | 读 P95 | 3.23ms | 2.64ms | -18.43% |
| 90/10 c4 | 发布 QPS | 6.52 | 6.14 | -5.85% |
| 90/10 c4 | 发布 P95 | 729.46ms | 779.37ms | +6.84% |
| 90/10 c4 | 总操作 QPS | 65.18 | 61.37 | -5.85% |
| 80/20 c4 | 读 QPS | 27.41 | 24.04 | -12.28% |
| 80/20 c4 | 读 P95 | 3.21ms | 3.19ms | -0.81% |
| 80/20 c4 | 发布 QPS | 6.85 | 6.01 | -12.28% |
| 80/20 c4 | 发布 P95 | 712.67ms | 828.74ms | +16.29% |
| 80/20 c4 | 总操作 QPS | 34.26 | 30.05 | -12.28% |

因此，本轮不能宣称 Treatment 提升了混合吞吐。能够确认的是：Treatment 减少了读依赖调用，并改善 90/10 的读尾延迟；同时，完整发布耗时和较低 Fresh 命中与总吞吐回退共同出现。由于本次一次开启整套 Treatment，具体回退来源仍需单开关 A/B 或同窗 Profile 才能归因。

## 2. 严格 A/B 口径

| 项目 | 固定值 |
|---|---|
| RunID | `feed-wp11-hot-mutation-ab-20260827a` |
| checkpoint ID | `43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f` |
| baseline fingerprint | `9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b` |
| 数据集 | 20 个热点读者、20 个普通作者、5 个大 V、50 个基线帖子、1,250 个 Feed Redis key |
| 策略与页面 | Hybrid、page1/size20、cold |
| 拓扑 | Docker 中间件 + 宿主机业务服务 |
| 机器口径 | 16 个逻辑处理器，GOMAXPROCS=16 |
| 日志 | `benchmark-error-only`，保留 error，关闭高频 access/stat/SQL info |
| 采样 | 10 秒预热 + 60 秒正式采样 × 3 |
| Control | RouteSnapshot、Combined Pipeline、Cursor、PageCache、epoch consumer 关闭 |
| Treatment | 上述目标态能力全部开启，PageCache L1/L2 开启 |

18 份有效 RPC 报告的 checkpoint ID、baseline fingerprint、manifest fingerprint、Git commit、patch hash 和业务服务二进制身份各只有一个唯一值。Control/Treatment 的功能开关不同是 A/B 自变量，其余比较条件一致。

## 3. 数据隔离与完整性

每个 trial 固定执行：初始恢复、10 秒 mutation 预热、预热后恢复、60 秒正式采样、Kafka 稳定排空、试后恢复、指纹校验。结果如下：

- Control 9/9、Treatment 9/9 报告完整，所有业务阶段失败和超时为 0。
- 18/18 初始恢复、预热后恢复、试后恢复成功。
- 18/18 Kafka 最终 lag=0，恢复流程完整。
- 试后恢复中位耗时约 5 秒；Kafka 排空中位耗时约 5–6 秒。
- Redis eviction 和 rejected connection 为 0。
- 测量产生的帖子 ID 全部被纳入精确恢复范围；下一轮重新从同一逻辑 Feed 初态开始。

这使本轮不同于早期不可比较的 mutation probe：帖子、Outbox、Feed ZSET 和缓存 epoch 不会随 trial 累积。

## 4. 已证实现象与待验证解释

### 4.1 读路径的重复工作确实减少

| 场景 | 指标/每次成功读取 | Control | Treatment | 变化 |
|---|---|---:|---:|---:|
| 90/10 | Relation 调用 | 1.0000 | 0.0699 | -93.01% |
| 90/10 | Redis 依赖调用 | 3.0489 | 1.9208 | -37.00% |
| 90/10 | cold compute | 1.0000 | 0.4641 | -53.59% |
| 90/10 | PageCache Fresh | 0% | 50.40% | +50.40pp |
| 80/20 | Relation 调用 | 1.0000 | 0.1484 | -85.16% |
| 80/20 | Redis 依赖调用 | 3.1157 | 2.9452 | -5.47% |
| 80/20 | cold compute | 1.0000 | 0.7015 | -29.85% |
| 80/20 | PageCache Fresh | 0% | 27.18% | +27.18pp |

90/10 的读 P95 从 3.23ms 降至 2.64ms，与依赖调用减少一致。Relation CPU 采样峰值中位数从 9.3% 降至 3.9%。80/20 的写入更密集，Fresh 命中和 Redis 降幅都更小，所以读 P95 基本持平。

### 4.2 固定比例调度会放大发布耗时的影响

90/10 和 80/20 使用一个全局有界并发池，读写按固定比例交错。一次“发布”是创建草稿、修改元数据、内容确认、正式发布四个串行 RPC，不是一次轻量写请求。慢发布占用 worker 时，读 QPS 与发布 QPS会按固定比例一起下降。

Treatment 的完整发布平均耗时在 90/10 中增加约 8.4%，在 80/20 中增加约 14.7%。对应总吞吐分别下降 5.85% 和 12.28%。发布 c2 单测中平均耗时只增加约 4.5%。这些数据证明回退主要出现在混合窗口，并与发布耗时增加相关；缓存写回/失效、异步事件处理或共享依赖竞争只是待验证候选，当前 bundled A/B 不能区分其贡献。

### 4.3 页面缓存无法在高 mutation 率下保持热度

90/10 每次发布之间约有 9 次读取，Treatment 得到约 50% Fresh；80/20 每次发布之间只有约 4 次读取，Fresh 降至约 27%。这证明写比例提高时页面复用明显下降，并与整页缓存收益收缩一致；是否由主动失效、TTL、回填成本或调度顺序分别造成，需要进一步拆分验证。

Redis 峰值仍约 19K ops/s，未随 Relation 调用下降而同步下降。该现象表明总 Redis 工作还包含显著的非 Relation 读工作，但仅凭峰值不能在写入、缓存回填、Feed 扇出等来源之间归因。

## 5. 资源证据

CPU 为三轮“采样最大值”的中位数，100% 约等于一个逻辑核。

| 场景 | 预设 | KnowPost CPU | Relation CPU | Redis max ops/s | MySQL running | Kafka lag max |
|---|---|---:|---:|---:|---:|---:|
| publish c2 | Control | 75.1% | 3.1% | 17,074 | 4 | 7 |
| publish c2 | Treatment | 68.2% | 3.1% | 17,316 | 5 | 7 |
| 90/10 c4 | Control | 82.9% | 9.3% | 18,965 | 4 | 8 |
| 90/10 c4 | Treatment | 82.9% | 3.9% | 19,307 | 7 | 8 |
| 80/20 c4 | Control | 90.0% | 8.5% | 19,856 | 5 | 8 |
| 80/20 c4 | Treatment | 84.1% | 3.9% | 19,118 | 5 | 7 |

Treatment 没有造成 CPU 饱和或最终消息积压，Relation 成本明显下降。当前证据把完整发布慢阶段列为下一优先分析对象，但没有隔离证明 Redis、MySQL 或某个 Treatment 开关是总容量瓶颈；需要 Profile 和单开关 A/B 后再下因果结论。

## 6. Gateway 范围调整

2026-08-27 用户明确决定停止后续 Gateway 性能测试，因为此前纯读矩阵已经确认 Gateway/JWT/HTTP/RPC 转发是独立瓶颈。本报告的正式 A/B 因此只接受 RPC：

- Control RPC：9 份有效报告。
- Treatment RPC：9 份有效报告。
- Control publish 命名空间内早先完成的 3 份 Gateway 报告只保留为历史旁证，不进入本 A/B。
- Control mixed 的 `expected-matrix.json` 是调整范围前生成的，仍列出 Gateway 格子；这些格子属于明确取消，不是 RPC 缺测。
- Treatment 使用 `-SkipGateway` 生成，预期矩阵与实际 RPC 范围一致。

后续 Feed 内核容量、回归和优化验证均优先使用 RPC 直连。已有 Gateway 报告继续用于说明入口瓶颈，但不再重复消耗正式矩阵时间。

## 7. 工程结论与下一步

1. 保留 RouteSnapshot 和 Combined Pipeline：它们在 mutation 下仍将 Relation/request 降低 85%–93%，机制有效。
2. PageCache 不应被描述为“任何负载都提升吞吐”。它适合热点、读多写少的第一页；写入比例升高后，需要按命中率或 mutation 率做准入/旁路，避免低复用页面反复编码和回填。
3. 下一性能主线应转向完整发布的 metadata、confirm、commit 三个慢阶段，拆分数据库事务、Outbox、同步依赖和连接池等待。
4. 后续正式性能验证只跑 RPC 稳定档及一个必要相邻档，不再跑 Gateway 矩阵。
5. 本轮证明的是“读优化在写竞争下的真实边界”，不是 Treatment 失败或数据无效。负收益同样是正式结论，不能通过挑选单轮或热缓存结果掩盖。

## 8. 原始证据

- [Control publish comparison（含范围调整前的历史 Gateway 行；本报告只取 RPC）](../../../results/feed-loadtest/feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2/comparison.md)
- [Control mixed](../../../results/feed-loadtest/feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4/comparison.md)
- [Treatment publish RPC](../../../results/feed-loadtest/feed-wp11-hot-mutation-ab-20260827a-treatment-formal-publish-c2-rpc/comparison.md)
- [Treatment mixed RPC](../../../results/feed-loadtest/feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc/comparison.md)
- [Manifest](../../../results/feed-loadtest/manifests/feed-wp11-hot-mutation-ab-20260827a.json)
- [Checkpoint](../../../results/feed-loadtest/checkpoints/feed-wp11-hot-mutation-ab-20260827a.json)
