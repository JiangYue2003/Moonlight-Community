# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:54:42+08:00
- 采样时长：1m0.0325725s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 77886 | 77886 | 0 | 0 | 1297.52 | 9.248 | 25.528 | 29.313 | 37.587 | 78.870 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 90939 | 1.168 |
| counter | 8967 | 0.115 |
| mysql | 10 | 0.000 |
| redis | 90446 | 1.161 |
| relation | 8967 | 0.115 |

- Cold compute：16911（0.217 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 37370 |
| l1_stale | 309 |
| l2_fresh | 32802 |
| l2_stale | 7404 |
| miss | 1 |

- L1+L2 Fresh ratio：90.10%
- Refresh max：queue=0 active=16 pending=16

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 16915 | 7.549 |
| counter | 8967 | 9.424 |
| hydrate | 16910 | 7.500 |
| inbox | 16915 | 7.548 |
| merge_dedup | 16915 | 0.004 |
| relation | 8967 | 11.843 |
| route | 16915 | 11.299 |
| total | 77886 | 7.104 |

## Redis 本轮边界增量

- Commands：280496；input：47911187 bytes；output：100347509 bytes
- Hits/Misses：442850/10384；run hit rate：97.71%
- Evicted/Rejected：0/0；ops/s max：9630；safety epoch：548 -> 548

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.543 |
| client:loadtest | cpu_percent_total | 104.683 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.090 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.670 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 26.000 |
| docker:zg-kafka | cpu_percent | 224.000 |
| docker:zg-kafka | memory_percent | 8.750 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 55.400 |
| docker:zg-zk | memory_percent | 1.440 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23383823.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 58.328 |
| process:counter | cpu_seconds_total | 139.547 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53800960.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 455.869 |
| process:gateway | cpu_seconds_total | 410.781 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 53760000.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 288.159 |
| process:knowpost | cpu_seconds_total | 2413.000 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 100413440.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 95.198 |
| process:relation | cpu_seconds_total | 136.969 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56598528.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.768 |
| process:search | cpu_seconds_total | 0.469 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36311040.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.806 |
| process:user-storage | cpu_seconds_total | 113.250 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 52256768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 209716882.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 548.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 702763.000 |
| redis | keyspace_hits | 1009135182.000 |
| redis | keyspace_misses | 6113801.000 |
| redis | net_input_bytes | 43418790919.000 |
| redis | net_output_bytes | 190108680620.000 |
| redis | ops_per_sec | 9630.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48921.000 |
| redis | used_memory_bytes | 123516024.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
