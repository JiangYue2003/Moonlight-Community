# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:43:23+08:00
- 采样时长：10.0278398s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 16882 | 16882 | 0 | 0 | 1683.69 | 12.214 | 42.469 | 46.473 | 56.621 | 98.650 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 16359 | 0.969 |
| counter | 1506 | 0.089 |
| mysql | 0 | 0.000 |
| redis | 16094 | 0.953 |
| relation | 1507 | 0.089 |

- Cold compute：2843（0.168 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 8851 |
| l1_stale | 203 |
| l2_fresh | 6771 |
| l2_stale | 1057 |
| miss | 0 |

- L1+L2 Fresh ratio：92.54%
- Refresh max：queue=6 active=32 pending=38

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2847 | 17.104 |
| counter | 1506 | 18.249 |
| hydrate | 2843 | 16.658 |
| inbox | 2847 | 17.103 |
| merge_dedup | 2847 | 0.003 |
| relation | 1507 | 20.246 |
| route | 2850 | 20.666 |
| total | 16882 | 13.334 |

## Redis 本轮边界增量

- Commands：48029；input：8105501 bytes；output：18057379 bytes
- Hits/Misses：76554/1727；run hit rate：97.79%
- Evicted/Rejected：0/0；ops/s max：7449；safety epoch：539 -> 539

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 9.047 |
| client:loadtest | cpu_percent_total | 144.753 |
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
| docker:zg-canal | cpu_percent | 2.610 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.260 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.420 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 68.230 |
| docker:zg-kafka | memory_percent | 8.260 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 9.900 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23282952.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 44.720 |
| process:counter | cpu_seconds_total | 31.453 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 50618368.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 422.754 |
| process:gateway | cpu_seconds_total | 136.062 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 53604352.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 267.539 |
| process:knowpost | cpu_seconds_total | 296.453 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 98938880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 71.320 |
| process:relation | cpu_seconds_total | 26.531 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 54530048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.769 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36171776.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 84.907 |
| process:user-storage | cpu_seconds_total | 40.344 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 51609600.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205848833.000 |
| redis | connected_clients | 245.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 539.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 696080.000 |
| redis | keyspace_hits | 1004100264.000 |
| redis | keyspace_misses | 5984570.000 |
| redis | net_input_bytes | 42846883545.000 |
| redis | net_output_bytes | 188835297275.000 |
| redis | ops_per_sec | 7449.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48192.000 |
| redis | used_memory_bytes | 119348680.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
