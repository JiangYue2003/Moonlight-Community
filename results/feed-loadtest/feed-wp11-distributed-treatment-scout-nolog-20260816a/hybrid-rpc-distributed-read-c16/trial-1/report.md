# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:41:33+08:00
- 采样时长：10.0097506s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 353891 | 353891 | 0 | 0 | 35388.38 | 0.518 | 0.944 | 1.071 | 1.662 | 23.718 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 23992 | 0.068 |
| counter | 1314 | 0.004 |
| mysql | 0 | 0.000 |
| redis | 23903 | 0.068 |
| relation | 1314 | 0.004 |

- Cold compute：3607（0.010 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 340720 |
| l1_stale | 0 |
| l2_fresh | 13171 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=3 pending=3

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3607 | 0.266 |
| counter | 1314 | 0.566 |
| hydrate | 3607 | 0.295 |
| inbox | 3607 | 0.266 |
| merge_dedup | 3607 | 0.004 |
| relation | 1314 | 1.131 |
| route | 3607 | 0.626 |
| total | 353891 | 0.026 |

## Redis 本轮边界增量

- Commands：71300；input：11054402 bytes；output：27486323 bytes
- Hits/Misses：99441/1615；run hit rate：98.40%
- Evicted/Rejected：0/0；ops/s max：13361；safety epoch：534 -> 534

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 15.024 |
| client:loadtest | cpu_percent_total | 240.391 |
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
| docker:zg-canal | cpu_percent | 2.040 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.030 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.310 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 56.300 |
| docker:zg-kafka | memory_percent | 8.260 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 41.020 |
| docker:zg-zk | memory_percent | 1.660 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23262428.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 24.740 |
| process:counter | cpu_seconds_total | 13.109 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 48152576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 35.234 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50704384.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 407.444 |
| process:knowpost | cpu_seconds_total | 64.594 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 88174592.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 35.564 |
| process:relation | cpu_seconds_total | 5.953 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 50929664.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36159488.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 11.094 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 47267840.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205249870.000 |
| redis | connected_clients | 132.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 534.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 688047.000 |
| redis | keyspace_hits | 1003313324.000 |
| redis | keyspace_misses | 5957000.000 |
| redis | net_input_bytes | 42757081657.000 |
| redis | net_output_bytes | 188653930186.000 |
| redis | ops_per_sec | 13361.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48082.000 |
| redis | used_memory_bytes | 105753248.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
