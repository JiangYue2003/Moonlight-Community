# Feed 压测报告：hybrid / gateway / publish-c8

- Run ID：`feed-wp11-hot-treatment-scout-gateway-publish-20260818`
- 开始时间：2026-08-18T22:40:58+08:00
- 采样时长：15.9799755s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 86 | 86 | 0 | 0 | 5.38 | 1464.455 | 1623.838 | 1646.916 | 1682.391 | 1682.391 |
| publish_draft | 86 | 86 | 0 | 0 | 5.38 | 4.763 | 7.020 | 7.558 | 8.145 | 8.145 |
| publish_metadata | 86 | 86 | 0 | 0 | 5.38 | 502.110 | 580.512 | 585.110 | 610.528 | 610.528 |
| publish_confirm | 86 | 86 | 0 | 0 | 5.38 | 481.143 | 592.104 | 608.082 | 643.149 | 643.149 |
| publish_commit | 86 | 86 | 0 | 0 | 5.38 | 476.906 | 535.340 | 538.445 | 559.012 | 559.012 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 0 |
| l2_fresh | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：265947；input：20787980 bytes；output：4767446 bytes
- Hits/Misses：212/1549；run hit rate：12.04%
- Evicted/Rejected：0/0；ops/s max：17505；safety epoch：7151 -> 7323

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.159 |
| client:loadtest | cpu_percent_total | 2.542 |
| client:loadtest | logical_cpus | 16.000 |
| docker-state:zg-canal | health_configured | 1.000 |
| docker-state:zg-canal | healthy | 1.000 |
| docker-state:zg-canal | restart_count | 0.000 |
| docker-state:zg-canal | running | 1.000 |
| docker-state:zg-es | health_configured | 1.000 |
| docker-state:zg-es | healthy | 1.000 |
| docker-state:zg-es | restart_count | 0.000 |
| docker-state:zg-es | running | 1.000 |
| docker-state:zg-etcd | health_configured | 1.000 |
| docker-state:zg-etcd | healthy | 1.000 |
| docker-state:zg-etcd | restart_count | 0.000 |
| docker-state:zg-etcd | running | 1.000 |
| docker-state:zg-kafka | health_configured | 1.000 |
| docker-state:zg-kafka | healthy | 1.000 |
| docker-state:zg-kafka | restart_count | 0.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 1.940 |
| docker:zg-canal | memory_percent | 3.800 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.740 |
| docker:zg-es | memory_percent | 11.680 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.230 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 126.510 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 40.450 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 5439.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | log_end_offset_total | 5439.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 30037.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.650 |
| process:counter | cpu_seconds_total | 489.172 |
| process:counter | pid | 6424.000 |
| process:counter | process_start_ms | 1787036967888.000 |
| process:counter | rss_bytes | 43282432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 8.778 |
| process:gateway | cpu_seconds_total | 5.391 |
| process:gateway | pid | 28656.000 |
| process:gateway | process_start_ms | 1787036989130.000 |
| process:gateway | rss_bytes | 43167744.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 92.962 |
| process:knowpost | cpu_seconds_total | 622.047 |
| process:knowpost | pid | 28324.000 |
| process:knowpost | process_start_ms | 1787036977976.000 |
| process:knowpost | rss_bytes | 74555392.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 15.484 |
| process:relation | pid | 27824.000 |
| process:relation | process_start_ms | 1787036972515.000 |
| process:relation | rss_bytes | 46628864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.798 |
| process:search | cpu_seconds_total | 13.781 |
| process:search | pid | 23244.000 |
| process:search | process_start_ms | 1787036983061.000 |
| process:search | rss_bytes | 43061248.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.394 |
| process:user-storage | cpu_seconds_total | 11.031 |
| process:user-storage | pid | 14844.000 |
| process:user-storage | process_start_ms | 1787036962489.000 |
| process:user-storage | rss_bytes | 40210432.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 17517847.000 |
| redis | connected_clients | 89.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 7323.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 592330.000 |
| redis | keyspace_hits | 406701.000 |
| redis | keyspace_misses | 176440.000 |
| redis | net_input_bytes | 1274266291.000 |
| redis | net_output_bytes | 350396989.000 |
| redis | ops_per_sec | 17505.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 44459.000 |
| redis | used_memory_bytes | 96297696.000 |

## 停止施压后的恢复

- Kafka drain：5.2530853s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：86
- 测量前恢复：complete=true；耗时=4.9025519s；删除帖子/Outbox=0/0；safety epoch=7101
- 预热后恢复：complete=true；耗时=6.2084659s；删除帖子/Outbox=24/48；safety epoch=7150
- 测量后恢复：complete=true；耗时=4.918453s；删除帖子/Outbox=86/172；safety epoch=7324

## 说明

- SLA values are reference lines, not pass/fail gates.
