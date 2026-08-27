# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-gateway-logverify-20260816a`
- 开始时间：2026-08-16T23:39:59+08:00
- 采样时长：5.0116913s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 16700 | 16700 | 0 | 0 | 3332.94 | 4.059 | 8.193 | 10.241 | 15.141 | 33.597 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 9599 | 0.575 |
| counter | 309 | 0.019 |
| mysql | 0 | 0.000 |
| redis | 9556 | 0.572 |
| relation | 309 | 0.019 |

- Cold compute：1493（0.089 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 11576 |
| l1_stale | 4 |
| l2_fresh | 4970 |
| l2_stale | 150 |
| miss | 0 |

- L1+L2 Fresh ratio：99.08%
- Refresh max：queue=0 active=4 pending=4

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1493 | 2.996 |
| counter | 309 | 6.033 |
| hydrate | 1493 | 3.080 |
| inbox | 1493 | 2.994 |
| merge_dedup | 1493 | 0.005 |
| relation | 309 | 7.360 |
| route | 1493 | 2.784 |
| total | 16700 | 1.387 |

## Redis 本轮边界增量

- Commands：25732；input：4400865 bytes；output：10880259 bytes
- Hits/Misses：39074/427；run hit rate：98.92%
- Evicted/Rejected：0/0；ops/s max：9197；safety epoch：533 -> 533

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 12.783 |
| client:loadtest | cpu_percent_total | 204.522 |
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
| docker:zg-canal | cpu_percent | 0.100 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.480 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 2.180 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 48.410 |
| docker:zg-kafka | memory_percent | 8.260 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 37.700 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23258685.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 49.745 |
| process:counter | cpu_seconds_total | 5.688 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 45182976.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 532.924 |
| process:gateway | cpu_seconds_total | 35.234 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52494336.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 243.750 |
| process:knowpost | cpu_seconds_total | 20.859 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 85057536.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 75.861 |
| process:relation | cpu_seconds_total | 3.188 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 48046080.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.062 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 35717120.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 106.406 |
| process:user-storage | cpu_seconds_total | 11.062 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 48435200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205052135.000 |
| redis | connected_clients | 122.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 533.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 686931.000 |
| redis | keyspace_hits | 1003150569.000 |
| redis | keyspace_misses | 5950107.000 |
| redis | net_input_bytes | 42733306201.000 |
| redis | net_output_bytes | 188615618543.000 |
| redis | ops_per_sec | 9197.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47983.000 |
| redis | used_memory_bytes | 104991736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
