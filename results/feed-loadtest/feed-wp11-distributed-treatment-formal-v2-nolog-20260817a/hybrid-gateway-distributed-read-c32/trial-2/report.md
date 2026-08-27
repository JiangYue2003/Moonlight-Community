# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:07:12+08:00
- 采样时长：1m0.0313438s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 69268 | 69268 | 0 | 0 | 1153.99 | 23.756 | 57.403 | 64.079 | 81.474 | 153.670 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 86749 | 1.252 |
| counter | 8840 | 0.128 |
| mysql | 8 | 0.000 |
| redis | 85476 | 1.234 |
| relation | 8840 | 0.128 |

- Cold compute：16272（0.235 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 30523 |
| l1_stale | 809 |
| l2_fresh | 29926 |
| l2_stale | 8009 |
| miss | 1 |

- L1+L2 Fresh ratio：87.27%
- Refresh max：queue=68 active=32 pending=100

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 16272 | 21.102 |
| counter | 8840 | 22.724 |
| hydrate | 16272 | 20.592 |
| inbox | 16272 | 21.101 |
| merge_dedup | 16272 | 0.004 |
| relation | 8840 | 26.165 |
| route | 16272 | 26.725 |
| total | 69268 | 21.543 |

## Redis 本轮边界增量

- Commands：262681；input：45571961 bytes；output：94443304 bytes
- Hits/Misses：425553/10200；run hit rate：97.66%
- Evicted/Rejected：0/0；ops/s max：13362；safety epoch：555 -> 555

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.847 |
| client:loadtest | cpu_percent_total | 109.552 |
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
| docker:zg-canal | cpu_percent | 2.250 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.470 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.490 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 166.240 |
| docker:zg-kafka | memory_percent | 8.740 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 40.240 |
| docker:zg-zk | memory_percent | 1.710 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23449573.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 42.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 52.432 |
| process:counter | cpu_seconds_total | 298.000 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54611968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 458.530 |
| process:gateway | cpu_seconds_total | 1550.953 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54611968.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 310.716 |
| process:knowpost | cpu_seconds_total | 3314.500 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 101298176.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 80.190 |
| process:relation | cpu_seconds_total | 341.594 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57602048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.767 |
| process:search | cpu_seconds_total | 0.734 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36364288.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 147.960 |
| process:user-storage | cpu_seconds_total | 376.469 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 56307712.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 211944570.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 555.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 706750.000 |
| redis | keyspace_hits | 1011954748.000 |
| redis | keyspace_misses | 6197972.000 |
| redis | net_input_bytes | 43754393147.000 |
| redis | net_output_bytes | 190732222672.000 |
| redis | ops_per_sec | 13362.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49671.000 |
| redis | used_memory_bytes | 125859064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
