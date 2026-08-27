# Feed Hybrid WP11：distributed/high 正式容量报告

日期：2026-08-17

状态：纯读压缩档完成；写入/混合与 Cursor 深分页优化待后续阶段

## 1. 本轮回答的问题

本轮只回答两个数据规模下的单机稳定读能力与 Control/Treatment 取舍：

- `distributed`：约 1,200 个读者；
- `high`：20,000 个读者；
- RPC 代表 Feed 内核上限，Gateway 代表包含 HTTP、JWT 与 Gateway→RPC 的入口能力；两者禁止混成一个“单机 QPS”。

写入、90/10、80/20 没有 checkpoint/restore，仍不具备三轮可比的数据初态，因此不在本报告给出正式容量数字。

## 2. 固定环境与有效性

- Hybrid，page1/size20，60 秒采样，每个正式点独立三轮并取中位数。
- 本地混合拓扑：既有 Kafka/etcd 等中间件容器，Gateway/KnowPost/Relation/Counter/User/Search 本地进程。
- Control 关闭 RouteSnapshot、Combined Pipeline、PageCache 与 epoch consumers；Treatment 全开。
- `benchmark-error-only`：仅保留 error，关闭 access/stat/SQL statement info；Gateway 启动路由输出不随请求增长。
- Gateway high-cardinality 使用本地 RS256 签名 token，仍经过真实 Gateway JWT 验签，避免 20,000 次 bcrypt 登录污染测量前态。
- 所有纳入结果的报告均 `complete=true`、失败=0、timeout=0、无重启、无 Redis eviction/rejection、Kafka lag 恢复、客户端未饱和且指标完整。

Windows Redis 曾因 RDB 后台保存失败触发 `MISCONF`/停止写入。所有受影响旧报告排除；本轮有效报告期间临时关闭 RDB schedule 和 `stop-writes-on-bgsave-error`，原值保存于 `.tmp/feed-local/redis-benchmark-original.json`。该设置只用于完全开发环境压测，不是生产建议。

## 3. 并发选档

为减少长测时间，先做 10 秒 scout，再只对稳定拐点跑 60 秒 × 3：

- distributed：沿用已确认最优档 RPC c64、Gateway c16；
- high Control RPC：c16=5745.83、c32=6307.15、c64=6476.53 QPS；c64 只比 c32 高 2.7%，P95 却从 7.36ms 升到 14.76ms，因此选 c32；
- high Control Gateway：c8=3005.20、c16=2274.40、c32=1884.30 QPS，因此选 c8；
- high Treatment RPC c32 首轮虽有 7768.72 QPS，但触发 `feed_degraded`，KnowPost error-only 日志同期明确记录 `Feed PageCache refresh queue is full`，因此报告无效并停止剩余轮次；c16 scout 有效后才运行正式三轮。

scout 只用于选档，不作为容量结论；短测在多组场景中明显高估 60 秒稳态。

## 4. 正式结果

| 数据规模 | 入口/预设 | C | 三轮 QPS | 中位 QPS | 中位 P95 / P99 | Fresh / cold compute | 资源侧中位证据 |
|---|---|---:|---|---:|---:|---:|---|
| distributed | Control RPC cold | 64 | 5556.24 / 6886.04 / 6797.75 | 6797.75 | 14.06 / 16.80ms | 0% / 1.00 | Redis max 60,204 ops/s；KnowPost CPU max 377.3%；RSS max 71.8MiB |
| distributed | Treatment RPC l2-warm | 64 | 102075.49 / 102312.74 / 99984.97 | 102075.49 | 1.16 / 1.69ms | 100% / ≈0 | Redis max 13,377 ops/s；KnowPost CPU max 643.9%；RSS max 95.4MiB |
| distributed | Control Gateway cold | 16 | 2048.25 / 1090.29 / 1218.99 | 1218.99 | 23.02 / 28.31ms | 0% / 1.00 | Redis max 14,465 ops/s；Gateway CPU max 219.8%；RSS max 48.9MiB |
| distributed | Treatment Gateway l2-warm | 16 | 1608.90 / 1161.09 / 1529.10 | 1529.10 | 26.34 / 33.41ms | 92% / 0.19 | Redis max 13,268 ops/s；Gateway CPU max 465.6%；RSS max 49.3MiB |
| high | Control RPC cold | 32 | 5434.52 / 6191.11 / 6162.47 | 6162.47 | 7.55 / 9.16ms | 0% / 1.00 | Redis max 67,955 ops/s；KnowPost CPU max 388.5%；RSS max 82.5MiB |
| high | Treatment RPC cold | 16 | 9573.70 / 9477.51 / 9436.74 | 9477.51 | 3.73 / 5.77ms | 72% / 0.40 | Redis max 63,908 ops/s；KnowPost CPU max 434.3%；RSS max 225.4MiB |
| high | Control Gateway cold | 8 | 2250.27 / 1402.80 / 1278.31 | 1402.80 | 11.20 / 13.49ms | 0% / 1.00 | Redis max 26,911 ops/s；Gateway CPU max 191.9%；RSS max 48.5MiB |
| high | Treatment Gateway cold | 8 | 2197.33 / 1355.07 / 1407.80 | 1407.80 | 15.16 / 18.62ms | 23% / 0.81 | Redis max 28,132 ops/s；Gateway CPU max 194.5%；RSS max 48.8MiB |

CPU 为进程总 CPU 百分比，100% 约等于占用一个逻辑核；表中为三轮“采样最大值”的中位数，不是平均 CPU。

## 5. 如何解释提升

### 5.1 distributed RPC：缓存命中时读路径被真正缩短

目标态中位数为 Control 的 15.02 倍，Redis 峰值 ops 反而从约 60K 降到约 13K。原因不是 Redis 更快，而是绝大多数请求在 L1/L2 完整页面返回，跳过 Relation、Counter、Inbox/BigV、hydrate、merge/dedup。这里比较的是 `cold Control → l2-warm Treatment` 的产品目标运行态，不是同缓存状态的纯代码增益。

### 5.2 high RPC：能接近 10K，但容量边界提前出现

Treatment c16 稳定在 9.48K QPS，三轮变异很小；c32 会使刷新/依赖路径降级，不能算有效容量。20,000 读者导致单用户复用降低，Fresh 只有约 72%，每个成功请求仍产生约 0.40 次冷计算，因此内存、Redis 和依赖成本没有像 distributed 热态那样消失。

### 5.3 Gateway：当前主要问题不是日志

关闭请求日志后，Gateway stdout 不再随请求量增长，但 high Gateway 的 Treatment 仍与 Control 基本持平，P95 还从 11.20ms 回退到 15.16ms。三轮中 Fresh 从首轮较高水平下降、cold compute 上升，吞吐同步下降；这说明主要瓶颈是高基数下的缓存热度/TTL、刷新与 Relation/Redis 冷路径恢复，再叠加 JWT、HTTP 编解码和 Gateway→RPC hop，而不是日志持久化。

## 6. 与历史 10K/Docker 数字的关系

历史 WP6 Docker 全服务 RPC、约 1,200 读者、不同实现与缓存前态，在 c128 得到约 10,661 QPS；它是当时口径的有效数字，但不能直接和本轮 Gateway/high-cardinality 混比。本轮同样证明：

- Feed 内核并未“失去 10K 能力”：distributed Treatment RPC 已超过 100K，high Treatment RPC 稳定约 9.5K；
- 无法复现的通常是把 Docker/本地、RPC/Gateway、1,200/20,000 读者、cold/warm、短测/60 秒稳态混为同一个数字；
- Docker 是否承载项目服务不是性能高低的唯一解释，真实决定因素是每请求经过的链路、缓存复用和资源口径。

## 7. 后续工作

1. 为 Feed API 增加稳定 Cursor/seek 分页，深分页不再反复放大 `page*size` 候选和 Top-N；单独建立首页、连续翻页、相同时间戳 tie-break、内容删除/插入与旧 page 参数兼容测试。
2. Profile high Gateway 的 JWT/HTTP、Gateway→RPC、Relation 与 Feed cold compute，占比明确后再决定 TTL 分层、热点 reader admission 或 Feed Head 物化。
3. 将 Redis benchmark persistence mode 加入 runtime state/report compatibility fingerprint；正式结束全部性能工作后按保存文件恢复原配置。
4. 实现 mutation checkpoint/restore 后再运行 publish、90/10、80/20；在此之前任何写/混合数字仍是 non-comparable probe。

## 8. 原始报告命名空间

- `feed-wp11-distributed-control-fixed-formal-nolog-20260817a`
- `feed-wp11-distributed-control-fixed-gateway-formal-nolog-20260817a`
- `feed-wp11-distributed-treatment-fixed-rpc-formal-nolog-20260817a`
- `feed-wp11-distributed-treatment-fixed-gateway-formal-nolog-20260817a`
- `feed-wp11-high-control-fixed-rpc-formal-nolog-20260817a`
- `feed-wp11-high-control-fixed-gateway-formal-nolog-20260817a`
- `feed-wp11-high-treatment-fixed-rpc-c16-formal-nolog-20260817a`
- `feed-wp11-high-treatment-fixed-gateway-c8-formal-nolog-20260817a`

边界/无效证据：`feed-wp11-high-treatment-fixed-rpc-formal-nolog-20260817a`（c32，`feed_degraded`）。
