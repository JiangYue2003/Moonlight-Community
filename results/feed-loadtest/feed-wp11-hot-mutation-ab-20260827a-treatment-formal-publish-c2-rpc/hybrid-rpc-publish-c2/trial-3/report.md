# Feed 压测报告：hybrid / rpc / publish-c2

- Run ID：`feed-wp11-hot-mutation-ab-20260827a-treatment-formal-publish-c2-rpc`
- 开始时间：2026-08-27T06:10:32+08:00
- 采样时长：1m0.1855558s
- 并发：2
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 382 | 382 | 0 | 0 | 6.35 | 308.161 | 359.986 | 384.414 | 553.480 | 569.053 |
| publish_draft | 382 | 382 | 0 | 0 | 6.35 | 3.942 | 5.250 | 5.695 | 182.243 | 262.777 |
| publish_metadata | 382 | 382 | 0 | 0 | 6.35 | 100.125 | 122.209 | 125.260 | 142.650 | 168.332 |
| publish_confirm | 382 | 382 | 0 | 0 | 6.35 | 99.016 | 124.810 | 130.522 | 151.283 | 327.175 |
| publish_commit | 382 | 382 | 0 | 0 | 6.35 | 100.092 | 120.293 | 128.011 | 159.456 | 179.146 |

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
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

## Redis 本轮边界增量

- Commands：945793；input：74670006 bytes；output：16833609 bytes
- Hits/Misses：1748/6973；run hit rate：20.04%
- Evicted/Rejected：0/0；ops/s max：17316；safety epoch：38525 -> 39289

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.154 |
| client:loadtest | cpu_percent_total | 2.466 |
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
| docker:zg-canal | cpu_percent | 2.230 |
| docker:zg-canal | memory_percent | 4.920 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.490 |
| docker:zg-es | memory_percent | 13.020 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.550 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 141.980 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 47.080 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 23612.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 23612.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 200184.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 5.422 |
| process:counter | cpu_seconds_total | 8.875 |
| process:counter | pid | 20740.000 |
| process:counter | process_start_ms | 1787781931297.000 |
| process:counter | rss_bytes | 42741760.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.094 |
| process:gateway | pid | 39628.000 |
| process:gateway | process_start_ms | 1787781953868.000 |
| process:gateway | rss_bytes | 37302272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 75.924 |
| process:knowpost | cpu_seconds_total | 133.781 |
| process:knowpost | pid | 39528.000 |
| process:knowpost | process_start_ms | 1787781942481.000 |
| process:knowpost | rss_bytes | 66703360.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.093 |
| process:relation | cpu_seconds_total | 1.844 |
| process:relation | pid | 21680.000 |
| process:relation | process_start_ms | 1787781936365.000 |
| process:relation | rss_bytes | 44646400.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.100 |
| process:search | cpu_seconds_total | 2.453 |
| process:search | pid | 27196.000 |
| process:search | process_start_ms | 1787781948810.000 |
| process:search | rss_bytes | 42360832.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.874 |
| process:user-storage | cpu_seconds_total | 0.328 |
| process:user-storage | pid | 34404.000 |
| process:user-storage | process_start_ms | 1787781926436.000 |
| process:user-storage | rss_bytes | 34983936.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19442262.000 |
| redis | connected_clients | 21.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 39289.000 |
| redis | hit_rate | 0.860 |
| redis | keys | 573328.000 |
| redis | keyspace_hits | 1386136.000 |
| redis | keyspace_misses | 231541.000 |
| redis | net_input_bytes | 1550309043.000 |
| redis | net_output_bytes | 585839611.000 |
| redis | ops_per_sec | 17316.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 12850.000 |
| redis | used_memory_bytes | 79432536.000 |

## 停止施压后的恢复

- Kafka drain：5.2989918s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`43dec82fd77f1173d630fef98b1ac6f2ab4aececde9f5f3e04fb47d1062dd24f`；baseline：`9262e56ae5c3e91bf1a5b3f9d13dccf7e35fd05b5b73c3602bfa86625066dc0b`
- 创建时间：2026-08-26T21:38:29.6034962Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：382
- 测量前恢复：complete=true；耗时=4.9587329s；删除帖子/Outbox=0/0；safety epoch=38391
- 预热后恢复：complete=true；耗时=6.3493487s；删除帖子/Outbox=66/132；safety epoch=38524
- 测量后恢复：complete=true；耗时=4.9861376s；删除帖子/Outbox=382/764；safety epoch=39290

## 说明

- SLA values are reference lines, not pass/fail gates.
