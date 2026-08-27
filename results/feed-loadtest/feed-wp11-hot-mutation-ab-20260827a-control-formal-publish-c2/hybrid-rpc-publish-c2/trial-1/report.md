# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2`
- 开始时间：2026-08-27T05:39:58+08:00
- 采样时长：1m0.2332415s
- 并发：2
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 380 | 380 | 0 | 0 | 6.31 | 300.749 | 387.332 | 442.305 | 597.141 | 695.421 |
| publish_draft | 380 | 380 | 0 | 0 | 6.31 | 4.060 | 5.292 | 5.678 | 6.150 | 14.616 |
| publish_metadata | 380 | 380 | 0 | 0 | 6.31 | 99.792 | 130.853 | 153.607 | 204.400 | 233.309 |
| publish_confirm | 380 | 380 | 0 | 0 | 6.31 | 94.216 | 128.601 | 158.431 | 214.495 | 252.373 |
| publish_commit | 380 | 380 | 0 | 0 | 6.31 | 99.193 | 131.837 | 141.722 | 218.471 | 234.373 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 0 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 0 | 0.000 |
| relation | 0 | 0.000 |

- Cold compute：0（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：865152；input：68258651 bytes；output：15321116 bytes
- Hits/Misses：1741/6228；run hit rate：21.85%
- Evicted/Rejected：0/0；ops/s max：16936；safety epoch：24197 -> 24577

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.183 |
| client:loadtest | cpu_percent_total | 2.931 |
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
| docker-state:zg-kafka | restart_count | 4.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 4.170 |
| docker:zg-canal | memory_percent | 3.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 19.330 |
| docker:zg-es | memory_percent | 12.470 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.960 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 170.860 |
| docker:zg-kafka | memory_percent | 7.420 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.140 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 18056.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 18056.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 68636.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.653 |
| process:counter | cpu_seconds_total | 11.844 |
| process:counter | pid | 36928.000 |
| process:counter | process_start_ms | 1787780172149.000 |
| process:counter | rss_bytes | 43147264.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 28216.000 |
| process:gateway | process_start_ms | 1787780195222.000 |
| process:gateway | rss_bytes | 36679680.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.869 |
| process:knowpost | cpu_seconds_total | 57.578 |
| process:knowpost | pid | 33572.000 |
| process:knowpost | process_start_ms | 1787780183823.000 |
| process:knowpost | rss_bytes | 61116416.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.077 |
| process:relation | cpu_seconds_total | 16.094 |
| process:relation | pid | 32348.000 |
| process:relation | process_start_ms | 1787780177514.000 |
| process:relation | rss_bytes | 54775808.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 4.652 |
| process:search | cpu_seconds_total | 1.328 |
| process:search | pid | 37640.000 |
| process:search | process_start_ms | 1787780190093.000 |
| process:search | rss_bytes | 41951232.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.969 |
| process:user-storage | pid | 35348.000 |
| process:user-storage | process_start_ms | 1787780167149.000 |
| process:user-storage | rss_bytes | 44916736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1465106.000 |
| redis | connected_clients | 75.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 24577.000 |
| redis | hit_rate | 0.745 |
| redis | keys | 566850.000 |
| redis | keyspace_hits | 36971.000 |
| redis | keyspace_misses | 18299.000 |
| redis | net_input_bytes | 109715486.000 |
| redis | net_output_bytes | 24549911.000 |
| redis | ops_per_sec | 16936.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 11016.000 |
| redis | used_memory_bytes | 79871344.000 |

## 停止施压后的恢复

- Kafka drain：5.2971127s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：380
- 测量前恢复：complete=true；耗时=5.1870722s；删除帖子/Outbox=0/0；safety epoch=24121
- 预热后恢复：complete=true；耗时=6.4874719s；删除帖子/Outbox=74/148；safety epoch=24196
- 测量后恢复：complete=true；耗时=5.0037273s；删除帖子/Outbox=380/760；safety epoch=24578

## 说明

- SLA values are reference lines, not pass/fail gates.
