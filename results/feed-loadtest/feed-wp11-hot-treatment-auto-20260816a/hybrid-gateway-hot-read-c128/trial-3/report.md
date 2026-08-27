# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:54:00+08:00
- 采样时长：1m0.0326549s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 277663 | 277663 | 0 | 0 | 4626.71 | 24.759 | 36.225 | 50.471 | 119.320 | 220.236 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 158 | 0.001 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 120 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：18（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 277559 |
| l1_stale | 0 |
| l2_fresh | 104 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.298 |
| counter | 9 | 1.993 |
| hydrate | 18 | 0.327 |
| inbox | 18 | 0.298 |
| merge_dedup | 18 | 0.000 |
| relation | 9 | 2.263 |
| route | 18 | 2.128 |
| total | 277663 | 0.007 |

## Redis 本轮边界增量

- Commands：53531；input：3826196 bytes；output：1433241 bytes
- Hits/Misses：1327/9；run hit rate：99.33%
- Evicted/Rejected：0/0；ops/s max：1101；safety epoch：508 -> 508

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.386 |
| client:loadtest | cpu_percent_total | 102.184 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.720 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.490 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 199.590 |
| docker:zg-kafka | memory_percent | 8.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 67.870 |
| docker:zg-zk | memory_percent | 1.440 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22730806.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 14.645 |
| process:counter | cpu_seconds_total | 181.547 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43053056.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 438.538 |
| process:gateway | cpu_seconds_total | 2381.578 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 60395520.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 187.884 |
| process:knowpost | cpu_seconds_total | 8320.375 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 111984640.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 9.078 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 44490752.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.544 |
| process:search | cpu_seconds_total | 1.641 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36978688.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 201.912 |
| process:user-storage | cpu_seconds_total | 1186.766 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 65658880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195689300.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 508.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677255.000 |
| redis | keyspace_hits | 993152814.000 |
| redis | keyspace_misses | 5378985.000 |
| redis | net_input_bytes | 41870493402.000 |
| redis | net_output_bytes | 187014582974.000 |
| redis | ops_per_sec | 1101.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41679.000 |
| redis | used_memory_bytes | 100053544.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
