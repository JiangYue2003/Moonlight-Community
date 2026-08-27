# Feed Cursor/seek 深分页正式 A/B 报告

日期：2026-08-17

结论：Cursor 的性能必要性已经在真实 1,500 候选深数据集上得到证明。RPC 与 Gateway 的 page20、page50 以及连续滚动 50 页均显著减少 Redis 返回成员、详情装载、每请求 CPU、内存分配和 GC；所有正式有效轮次均为 0 错误、0 超时、0 重复、0 oracle 不一致。建议保持功能开关可回滚，先在 RPC 和少量 Gateway 流量灰度启用，不直接全量默认开启。原因不是深页收益不足，而是缺少改造前同 manifest 的 page1 二进制基线，且本机 Gateway 浅页吞吐存在较大调度噪声。

## 1. 为什么旧 deep-page 结果不能回答问题

旧 WP11 distributed/high 数据中，每个读者只有约 8–14 条内容；`page=5,size=20` 多数是空页。它只能说明深页不命中首页缓存，不能测出 rank/offset 对候选读取、归并、详情装载和 GC 的线性放大。本次建立独立 `feed-cursor-deep-v1`：

| 维度 | 实际值 |
|---|---:|
| 读者 | 20 |
| 普通作者 | 20 × 50 帖 |
| 大 V | 5 × 100 帖 |
| 每读者 Inbox / BigV 候选 | 1,000 / 500 |
| 原始帖子 / 去重 oracle | 1,500 / 1,499 |
| 同秒组 | 120 条 |
| 跨来源重复帖子 | 1 条 |

Manifest 只有在 Kafka lag=0、Redis Inbox/Outbox、MySQL ownership/status/visible/publish_time、全序 oracle 和 SHA-256 指纹全部一致后才置 `ready=true`。最终 Redis、MySQL、oracle 指纹分别为：

- Redis：`449bb3422bbc5fedb298afc9b337ac59e609ef33be3c8445c1563473e673623e`
- MySQL：`2f6437dbe84c3b28dd27acafae0def625df314391a9902933cbdee18e7e032e2`
- Oracle：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## 2. 测试口径

- 拓扑：Docker 只运行 etcd、Kafka、ZooKeeper、Canal、Elasticsearch 等既有中间件；Gateway、KnowPost、Relation、Counter、User、Search 为宿主机本地进程。Redis/MySQL 使用本机实例。
- 策略：Hybrid；RouteSnapshot、关系/安全 epoch、Combined Pipeline、Cursor、可观测性开启；PageCache 关闭。
- 日志：`benchmark-error-only`，关闭 Gateway access log、go-zero stat、RPC stat 和 SQL statement info，保留真实 error。
- 数据状态：静态发布/公开数据；压测期间不做写入 mutation；Redis eviction/rejection 为 0，Kafka lag 为 0。
- 入口：RPC 与 Gateway 分开；Gateway 使用已有真实用户 ID 的签名 Token，不并发 Login。
- 正式采样：c16 稳定档和 c32 相邻档；每轮预热 10 秒、采样 60 秒、三轮，使用中位数并同时保留 min/max。
- 固定页：Control 直接请求 `page=N,size=20`；Treatment 在采样前构造并校验前序 Cursor，只计目标页请求。
- 连续滚动：`sequential-page-50` 与 `sequential-50` 都把第 1–50 页的每次请求计入，二者在同一 Git patch、同一业务二进制和同一进程下配对执行。
- 完整性：进程身份、被测二进制 SHA、Git/dirty patch、Redis identity、Docker 健康、Kafka/MySQL/Redis/Prometheus、客户端 CPU、运行时 Flag 和安全 epoch 任一缺失即 `complete=false`。

侦察结果显示 c16 是稳定档。RPC 首页 c32 相比 c16只增加约 3% 吞吐，P95 约翻倍；深页和 Gateway 中 c32 多数不增反降。因此正式结论不向 c64 继续扩展。

## 3. 固定目标页正式结果

下表使用补齐 Runtime 遥测后的 c16 三轮中位数；括号内为三轮 QPS min–max。

| 入口/目标页 | 分页 | QPS | P95 / P99 | Redis 成员/请求 | 详情 ID/请求 | 分配 KB/请求 | GC 次数/千请求 |
|---|---|---:|---:|---:|---:|---:|---:|
| RPC page20 | 页码 | 1,178.8（1,172.9–1,183.4） | 22.67 / 27.03 ms | 921 | 420 | 721.9 | 77.14 |
| RPC page20 | Cursor | 4,848.8（4,802.5–4,849.2） | 5.38 / 6.80 ms | 67 | 21 | 64.5 | 7.27 |
| RPC page50 | 页码 | 592.4（589.2–595.5） | 35.66 / 41.02 ms | 1,500 | 1,020 | 1,667.7 | 133.47 |
| RPC page50 | Cursor | 7,621.7（7,581.0–7,736.3） | 3.36 / 4.35 ms | 24 | 21 | 58.9 | 6.52 |
| Gateway page20 | 页码 | 1,045.2（1,030.9–1,082.5） | 25.41 / 31.22 ms | 921 | 420 | 722.4 | 77.71 |
| Gateway page20 | Cursor | 1,698.2（1,233.0–1,788.9） | 15.72 / 20.34 ms | 67 | 21 | 64.2 | 6.53 |
| Gateway page50 | 页码 | 489.2（483.7–489.5） | 49.02 / 57.09 ms | 1,500 | 1,020 | 1,669.1 | 133.58 |
| Gateway page50 | Cursor | 2,180.4（1,575.9–2,215.4） | 11.37 / 14.29 ms | 24 | 21 | 58.6 | 6.04 |

相对改进：

| 入口/目标页 | QPS | P95 | Redis 成员 | 详情 ID | 分配 | KnowPost CPU/请求 |
|---|---:|---:|---:|---:|---:|---:|
| RPC page20 | +311% | -76% | -92.7% | -95.0% | -91.1% | 1.76→0.41 ms |
| RPC page50 | +1,187% | -90.6% | -98.4% | -97.9% | -96.5% | 3.36→0.39 ms |
| Gateway page20 | +62% | -38% | -92.7% | -95.0% | -91.1% | 1.89→0.86 ms |
| Gateway page50 | +346% | -76.8% | -98.4% | -97.9% | -96.5% | 3.51→0.90 ms |

KnowPost 峰值 RSS 在 RPC page20/page50 中分别由约 73/82 MB 降至约 68/69 MB。Cursor 因单位时间处理的请求更多，Redis 总 ops/s 和服务总 CPU 可能更高；这不是单位请求回退。更准确的口径是 CPU/request、returned members/request 和 allocation/request，它们均大幅下降。精确 seek 每请求 Redis 命令数约由 7–8 增至 14–15，因为需要多范围 Pipeline 和边界 tie 补齐；网络命令数增加换取返回 payload 从 921/1,500 降到 67/24。

## 4. 连续滚动 50 页严格配对结果

该场景统计每一次真实页面请求；一条完整滚动链包含 50 页。两侧在同一 patch 下运行，12/12 报告完整。

| 入口 | 分页 | 页请求 QPS | 完整链/60s（中位） | P95 / P99 | 成员/页 | 详情 ID/页 | 分配 KB/页 | GC/千页 |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| RPC | 页码 | 991.0 | 1,200 | 28.02 / 33.66 ms | 1,018.9 | 530.0 | 922.3 | 87.3 |
| RPC | Cursor | 4,974.6 | 5,975 | 5.42 / 7.01 ms | 73.0 | 21.4 | 65.5 | 7.7 |
| Gateway | 页码 | 801.8 | 974 | 36.98 / 46.26 ms | 1,018.9 | 530.0 | 922.8 | 89.7 |
| Gateway | Cursor | 1,797.0 | 2,166 | 14.38 / 18.53 ms | 73.0 | 21.4 | 65.2 | 6.9 |

RPC 中 Cursor 的页吞吐约提升 5.0 倍，P95 降低约 81%；Gateway 中约提升 2.24 倍，P95 降低约 61%。每页返回成员降低 92.8%，详情装载降低 96.0%，分配降低约 92.9%。KnowPost CPU/request 在 RPC 中由约 2.10 ms 降到 0.42 ms，在 Gateway 入口的后端部分由约 2.20 ms 降到 0.85 ms。

## 5. 正确性与边界

- Runtime 关键点 24/24、连续滚动配对 12/12 报告完整；累计失败、超时、oracle mismatch、重复、Cursor loop、提前结束均为 0。
- 同秒边界 probe：RPC 5,614 次、Gateway 4,085 次请求全部成功；每请求读取 236 个成员，其中 126 个来自 tie group，0 漏项/重复/循环。
- 静态 oracle、同秒数值 postID 排序、跨 Inbox/BigV 去重、插入、删除、转私密、取关、回填、依赖部分失败、Context 取消、非法 Cursor、开关关闭和旧 page 兼容测试全部通过。
- Cursor page50 的 P95 低于 page5，而非随页深增长。page5 落在刻意构造的 120 条同秒组中，Cursor 每请求需回读 236 个成员；page50 只需 24 个。这证明成本主要由来源数、size 和 tie group 决定，而不是由页号线性决定。

## 6. 浅页与 Gateway 噪声

RPC page5 正式中位数为页码 2,338 QPS / P95 11.27 ms，Cursor 2,677 QPS / 9.89 ms，是小幅正收益。Gateway page5 正式中位数则从 1,568 QPS / 17.29 ms 变为 1,233 QPS / 21.48 ms。三组 10 秒正反顺序复验中，页码赢两次、Cursor 赢一次，结果随运行顺序漂移。

因此本报告不声称 Cursor 在浅页必然更快。page5 同秒大组让精确 seek 的固定成本接近页码成本，Gateway 的 JWT、HTTP/JSON、RPC 转发与 Windows 宿主机调度噪声会主导剩余差异。默认启用的主要证据来自 page20/page50 和连续滚动，不来自挑选 page5 的最好单轮。

## 7. 决策与限制

通过的决策规则：

- page50 Redis 返回成员 1,500→24，下降 98.4%，超过 70% 门槛。
- RPC/Gateway page20、page50 P95 均改善超过 30%；P99、CPU/request、分配和 GC 不回退。
- 连续滚动的成员、详情、CPU/request 均下降超过 50%，总耗时和 QPS 不回退。
- page50 Cursor P95 不高于 page5 的 1.5 倍。

尚未直接闭环的门禁是“相对改造前同 manifest 的 page1 二进制基线”。本次改造后的 page1 RPC c16 中位约 5,074 QPS，兼容/回归测试全部通过，但改造前没有保存 `feed-cursor-deep-v1` 的同二进制、同数据正式报告，不能事后伪造。因此决策为：

1. Cursor 接口和性能必要性已证明；
2. 保持 `Feed.CursorPagination.Enabled` 可一键关闭；
3. 先在 RPC 与少量 Gateway 流量灰度，监控非法 Cursor、依赖错误、P99、回填、分配和 GC；
4. 观察稳定后再扩大，不把当前结果解释为“所有页都默认更快”。

Feed 仍是最终一致集合：翻页期间的新帖不会插入旧 Cursor 窗口，删除/转私密/取关会立即影响后续页并可能缩短页面。这是已接受的短暂不一致/不可见取舍，不承诺跨请求快照。

本报告不解除 WP11 mutation 门禁。publish、90/10、80/20 仍必须先实现 trial 级 MySQL/Redis/Kafka/cache epoch checkpoint/restore，不能用本次静态读 A/B 代替。

## 8. 原始证据

- 深数据 manifest：`results/feed-loadtest/manifests/feed-cursor-deep-v1.json`
- RPC 全场景正式矩阵：`results/feed-loadtest/feed-cursor-deep-v1-rpc-formal-20260817a`
- Gateway 全场景正式矩阵：`results/feed-loadtest/feed-cursor-deep-v1-gateway-formal-20260817a`
- Runtime 关键点：`results/feed-loadtest/feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 严格累计 A/B：`results/feed-loadtest/feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 同秒边界：`results/feed-loadtest/feed-cursor-deep-v1-same-second-probe-20260817a`
- Gateway page5 诊断：`results/feed-loadtest/feed-cursor-deep-v1-gateway-page5-recheck-{a,b,c}`

曾因把 `go_memstats_heap_alloc_bytes` Gauge 当作单调 Counter 而产生的无效报告保留在 `feed-cursor-deep-v1-keypoints-runtime-formal-20260817a`，明确不纳入结论；修复后使用新命名空间全量重跑并通过完整性门禁。
