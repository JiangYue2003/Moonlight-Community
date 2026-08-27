# Feed 压测报告：hybrid / gateway / mixed-80-20-c8

- Run ID：`feed-wp11-hot-control-formal-mixed-c4-c8-20260818`
- 开始时间：2026-08-19T01:19:53+08:00
- 采样时长：1m0.2448116s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 1312 | 1312 | 0 | 0 | 21.78 | 3.711 | 5.378 | 6.030 | 7.500 | 9.148 |
| publish_total | 328 | 328 | 0 | 0 | 5.44 | 1434.614 | 1590.872 | 1642.853 | 1665.566 | 1675.392 |
| publish_draft | 328 | 328 | 0 | 0 | 5.44 | 4.216 | 4.967 | 5.420 | 6.372 | 9.254 |
| publish_metadata | 328 | 328 | 0 | 0 | 5.44 | 490.962 | 592.903 | 625.534 | 654.218 | 660.487 |
| publish_confirm | 328 | 328 | 0 | 0 | 5.44 | 465.952 | 575.425 | 597.678 | 647.107 | 660.290 |
| publish_commit | 328 | 328 | 0 | 0 | 5.44 | 459.700 | 569.695 | 592.253 | 625.600 | 678.558 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 200 | 0.152 |
| mysql | 114 | 0.087 |
| redis | 4050 | 3.087 |
| relation | 1312 | 1.000 |

- Cold compute：1312（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 52480 | 40.000 |
| merge_candidates | 90530 | 69.002 |
| redis_commands | 7872 | 6.000 |
| redis_members | 97222 | 74.102 |
| redis_roundtrips | 2624 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 1312 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1312 | 0.543 |
| counter | 200 | 0.826 |
| hydrate | 1312 | 0.699 |
| inbox | 1312 | 0.553 |
| merge_dedup | 1312 | 0.013 |
| relation | 1312 | 1.251 |
| route | 1312 | 0.133 |
| total | 1312 | 3.219 |

## Redis 本轮边界增量

- Commands：989343；input：78719239 bytes；output：29634664 bytes
- Hits/Misses：65278/8063；run hit rate：89.01%
- Evicted/Rejected：0/0；ops/s max：18711；safety epoch：23008 -> 23336

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.143 |
| client:loadtest | cpu_percent_total | 2.282 |
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
| docker:zg-canal | cpu_percent | 2.800 |
| docker:zg-canal | memory_percent | 5.010 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 16.270 |
| docker:zg-es | memory_percent | 12.300 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.420 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 206.020 |
| docker:zg-kafka | memory_percent | 8.040 |
| docker:zg-kafka | pids | 143.000 |
| docker:zg-zk | cpu_percent | 47.330 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 17092.000 |
| kafka | lag_max | 8.000 |
| kafka | lag_total | 8.000 |
| kafka | log_end_offset_total | 17092.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 334035.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.199 |
| process:counter | cpu_seconds_total | 203.141 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 43704320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.428 |
| process:gateway | cpu_seconds_total | 37.812 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 47304704.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 89.139 |
| process:knowpost | cpu_seconds_total | 2021.781 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 67584000.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.745 |
| process:relation | cpu_seconds_total | 65.766 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47083520.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.323 |
| process:search | cpu_seconds_total | 30.812 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 41639936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.875 |
| process:user-storage | cpu_seconds_total | 23.703 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 40587264.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 69427143.000 |
| redis | connected_clients | 43.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 23336.000 |
| redis | hit_rate | 0.812 |
| redis | keys | 589114.000 |
| redis | keyspace_hits | 4094499.000 |
| redis | keyspace_misses | 947467.000 |
| redis | net_input_bytes | 5365012594.000 |
| redis | net_output_bytes | 1840235921.000 |
| redis | ops_per_sec | 18711.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54038.000 |
| redis | used_memory_bytes | 95403448.000 |

## 停止施压后的恢复

- Kafka drain：5.281357s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：328
- 测量前恢复：complete=true；耗时=4.9363129s；删除帖子/Outbox=0/0；safety epoch=22942
- 预热后恢复：complete=true；耗时=6.302438s；删除帖子/Outbox=64/128；safety epoch=23007
- 测量后恢复：complete=true；耗时=4.8995251s；删除帖子/Outbox=328/656；safety epoch=23337

## 说明

- SLA values are reference lines, not pass/fail gates.
