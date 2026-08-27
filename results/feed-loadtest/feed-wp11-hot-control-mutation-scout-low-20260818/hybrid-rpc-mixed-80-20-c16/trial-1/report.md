# Feed 压测报告：hybrid / rpc / mixed-80-20-c16

- Run ID：`feed-wp11-hot-control-mutation-scout-low-20260818`
- 开始时间：2026-08-18T23:02:36+08:00
- 采样时长：16.1164225s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 348 | 348 | 0 | 0 | 21.59 | 6.957 | 9.555 | 10.359 | 13.289 | 14.195 |
| publish_total | 86 | 86 | 0 | 0 | 5.34 | 2967.494 | 3181.781 | 3192.447 | 3243.726 | 3243.726 |
| publish_draft | 86 | 86 | 0 | 0 | 5.34 | 4.745 | 6.109 | 6.557 | 7.385 | 7.385 |
| publish_metadata | 86 | 86 | 0 | 0 | 5.34 | 1024.125 | 1162.782 | 1176.627 | 1231.596 | 1231.596 |
| publish_confirm | 86 | 86 | 0 | 0 | 5.34 | 944.297 | 1071.873 | 1106.225 | 1118.664 | 1118.664 |
| publish_commit | 86 | 86 | 0 | 0 | 5.34 | 958.392 | 1130.796 | 1149.301 | 1238.585 | 1238.585 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 60 | 0.172 |
| mysql | 37 | 0.106 |
| redis | 1081 | 3.106 |
| relation | 348 | 1.000 |

- Cold compute：348（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 13920 | 40.000 |
| merge_candidates | 16863 | 48.457 |
| redis_commands | 2088 | 6.000 |
| redis_members | 16863 | 48.457 |
| redis_roundtrips | 696 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 348 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 348 | 1.338 |
| counter | 60 | 1.349 |
| hydrate | 348 | 1.635 |
| inbox | 348 | 1.333 |
| merge_dedup | 348 | 0.018 |
| relation | 348 | 2.217 |
| route | 348 | 0.251 |
| total | 348 | 6.810 |

## Redis 本轮边界增量

- Commands：273107；input：21712079 bytes；output：7713326 bytes
- Hits/Misses：16941/2715；run hit rate：86.19%
- Evicted/Rejected：0/0；ops/s max：18578；safety epoch：10358 -> 10444

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.206 |
| client:loadtest | cpu_percent_total | 3.296 |
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
| docker:zg-canal | cpu_percent | 4.590 |
| docker:zg-canal | memory_percent | 4.660 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.910 |
| docker:zg-es | memory_percent | 11.870 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.910 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 14.000 |
| docker:zg-kafka | cpu_percent | 159.900 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 41.490 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 7098.000 |
| kafka | lag_max | 13.000 |
| kafka | lag_total | 13.000 |
| kafka | log_end_offset_total | 7098.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 71768.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.197 |
| process:counter | cpu_seconds_total | 17.156 |
| process:counter | pid | 24444.000 |
| process:counter | process_start_ms | 1787064665305.000 |
| process:counter | rss_bytes | 45395968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.141 |
| process:gateway | pid | 29440.000 |
| process:gateway | process_start_ms | 1787064684282.000 |
| process:gateway | rss_bytes | 38334464.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 77.459 |
| process:knowpost | cpu_seconds_total | 174.531 |
| process:knowpost | pid | 9704.000 |
| process:knowpost | process_start_ms | 1787064675235.000 |
| process:knowpost | rss_bytes | 69648384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.422 |
| process:relation | cpu_seconds_total | 5.391 |
| process:relation | pid | 26956.000 |
| process:relation | process_start_ms | 1787064669914.000 |
| process:relation | rss_bytes | 47939584.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.798 |
| process:search | cpu_seconds_total | 2.422 |
| process:search | pid | 11296.000 |
| process:search | process_start_ms | 1787064680243.000 |
| process:search | rss_bytes | 42258432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.203 |
| process:user-storage | pid | 30296.000 |
| process:user-storage | process_start_ms | 1787064660932.000 |
| process:user-storage | rss_bytes | 35389440.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 25043921.000 |
| redis | connected_clients | 70.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 10444.000 |
| redis | hit_rate | 0.731 |
| redis | keys | 592691.000 |
| redis | keyspace_hits | 960725.000 |
| redis | keyspace_misses | 353294.000 |
| redis | net_input_bytes | 1866413611.000 |
| redis | net_output_bytes | 547684497.000 |
| redis | ops_per_sec | 18578.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 45758.000 |
| redis | used_memory_bytes | 95882232.000 |

## 停止施压后的恢复

- Kafka drain：6.6615296s
- 完成：true
- 最终 lag：0

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：86
- 测量前恢复：complete=true；耗时=4.8758018s；删除帖子/Outbox=0/0；safety epoch=10324
- 预热后恢复：complete=true；耗时=7.5193218s；删除帖子/Outbox=32/64；safety epoch=10357
- 测量后恢复：complete=true；耗时=4.9029438s；删除帖子/Outbox=86/172；safety epoch=10445

## 说明

- SLA values are reference lines, not pass/fail gates.
