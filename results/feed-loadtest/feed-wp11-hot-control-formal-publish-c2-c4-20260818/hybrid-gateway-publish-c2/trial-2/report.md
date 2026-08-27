# Feed 压测报告：hybrid / gateway / publish-c2

- Run ID：`feed-wp11-hot-control-formal-publish-c2-c4-20260818`
- 开始时间：2026-08-18T23:26:45+08:00
- 采样时长：1m0.2544676s
- 并发：2
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 282 | 282 | 0 | 0 | 4.68 | 413.310 | 494.551 | 520.583 | 614.298 | 665.105 |
| publish_draft | 282 | 282 | 0 | 0 | 4.68 | 4.384 | 5.854 | 6.218 | 7.401 | 180.337 |
| publish_metadata | 282 | 282 | 0 | 0 | 4.68 | 135.312 | 170.169 | 176.978 | 205.183 | 213.743 |
| publish_confirm | 282 | 282 | 0 | 0 | 4.68 | 134.931 | 165.990 | 182.976 | 342.074 | 365.106 |
| publish_commit | 282 | 282 | 0 | 0 | 4.68 | 134.696 | 161.429 | 174.523 | 201.209 | 219.055 |

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

- Commands：893722；input：69839496 bytes；output：15990626 bytes
- Hits/Misses：665/5436；run hit rate：10.90%
- Evicted/Rejected：0/0；ops/s max：16264；safety epoch：14086 -> 14368

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.091 |
| client:loadtest | cpu_percent_total | 1.452 |
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
| docker:zg-canal | cpu_percent | 3.130 |
| docker:zg-canal | memory_percent | 4.900 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.490 |
| docker:zg-es | memory_percent | 12.160 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.230 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 20.000 |
| docker:zg-kafka | cpu_percent | 138.140 |
| docker:zg-kafka | memory_percent | 7.640 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 6.890 |
| docker:zg-zk | memory_percent | 1.320 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 10071.000 |
| kafka | lag_max | 5.000 |
| kafka | lag_total | 5.000 |
| kafka | log_end_offset_total | 10071.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 139138.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 8.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 15.503 |
| process:counter | cpu_seconds_total | 54.625 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45879296.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.422 |
| process:gateway | cpu_seconds_total | 8.406 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 46911488.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 71.244 |
| process:knowpost | cpu_seconds_total | 677.453 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 71290880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.870 |
| process:relation | cpu_seconds_total | 16.266 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 48574464.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.622 |
| process:search | cpu_seconds_total | 10.719 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 43560960.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.873 |
| process:user-storage | cpu_seconds_total | 5.484 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 41619456.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 37780478.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 14368.000 |
| redis | hit_rate | 0.725 |
| redis | keys | 594996.000 |
| redis | keyspace_hits | 1423025.000 |
| redis | keyspace_misses | 544113.000 |
| redis | net_input_bytes | 2859474188.000 |
| redis | net_output_bytes | 824267647.000 |
| redis | ops_per_sec | 16264.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47251.000 |
| redis | used_memory_bytes | 96033408.000 |

## 停止施压后的恢复

- Kafka drain：6.5046338s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：282
- 测量前恢复：complete=true；耗时=4.9431743s；删除帖子/Outbox=0/0；safety epoch=14032
- 预热后恢复：complete=true；耗时=6.3769176s；删除帖子/Outbox=52/104；safety epoch=14085
- 测量后恢复：complete=true；耗时=4.9400016s；删除帖子/Outbox=282/564；safety epoch=14369

## 说明

- SLA values are reference lines, not pass/fail gates.
