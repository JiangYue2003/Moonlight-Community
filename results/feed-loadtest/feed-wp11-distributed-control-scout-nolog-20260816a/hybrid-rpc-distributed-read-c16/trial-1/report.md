# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:31:27+08:00
- 采样时长：10.0051437s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 62027 | 62027 | 0 | 0 | 6201.76 | 2.601 | 3.586 | 3.853 | 4.788 | 13.773 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.039 |
| mysql | 0 | 0.000 |
| redis | 186081 | 3.000 |
| relation | 62027 | 1.000 |

- Cold compute：62027（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 62027 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62027 | 0.369 |
| counter | 2400 | 0.610 |
| hydrate | 62027 | 0.417 |
| inbox | 62027 | 0.383 |
| merge_dedup | 62027 | 0.004 |
| relation | 62027 | 1.100 |
| route | 62027 | 0.027 |
| total | 62027 | 2.318 |

## Redis 本轮边界增量

- Commands：494171；input：56901923 bytes；output：170106516 bytes
- Hits/Misses：1151324/67222；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：51805；safety epoch：524 -> 524

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.208 |
| client:loadtest | cpu_percent_total | 99.324 |
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
| docker:zg-canal | cpu_percent | 2.960 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.530 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.210 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 225.530 |
| docker:zg-kafka | memory_percent | 8.260 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.090 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22854755.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 24.664 |
| process:counter | cpu_seconds_total | 10.344 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 45727744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.016 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 37175296.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 352.184 |
| process:knowpost | cpu_seconds_total | 75.625 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 62947328.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 294.372 |
| process:relation | cpu_seconds_total | 49.766 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 49819648.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35721216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.125 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 34390016.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 201617299.000 |
| redis | connected_clients | 50.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 524.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677094.000 |
| redis | keyspace_hits | 995681894.000 |
| redis | keyspace_misses | 5515300.000 |
| redis | net_input_bytes | 42345228204.000 |
| redis | net_output_bytes | 187507612015.000 |
| redis | ops_per_sec | 51805.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47476.000 |
| redis | used_memory_bytes | 97460816.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
