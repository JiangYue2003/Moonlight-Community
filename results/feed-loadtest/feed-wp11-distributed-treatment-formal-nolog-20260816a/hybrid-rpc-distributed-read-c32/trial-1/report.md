# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:46:11+08:00
- 采样时长：1m0.0879847s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4717311 | 4717311 | 0 | 0 | 78621.26 | 0.518 | 0.578 | 0.723 | 1.120 | 16.091 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146066 | 0.031 |
| counter | 10825 | 0.002 |
| mysql | 3 | 0.000 |
| redis | 145202 | 0.031 |
| relation | 10825 | 0.002 |

- Cold compute：22264（0.005 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4638037 |
| l1_stale | 0 |
| l2_fresh | 79274 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=5 pending=5

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22264 | 0.144 |
| counter | 10825 | 0.441 |
| hydrate | 22264 | 0.166 |
| inbox | 22264 | 0.142 |
| merge_dedup | 22264 | 0.005 |
| relation | 10825 | 0.979 |
| route | 22264 | 0.701 |
| total | 4717311 | 0.008 |

## Redis 本轮边界增量

- Commands：457016；input：68707532 bytes；output：167124378 bytes
- Hits/Misses：626097/12677；run hit rate：98.02%
- Evicted/Rejected：0/0；ops/s max：14216；safety epoch：542 -> 542

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 20.835 |
| client:loadtest | cpu_percent_total | 333.365 |
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
| docker:zg-canal | cpu_percent | 3.220 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.110 |
| docker:zg-es | memory_percent | 12.030 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 6.090 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 230.140 |
| docker:zg-kafka | memory_percent | 8.750 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 60.930 |
| docker:zg-zk | memory_percent | 1.660 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23304236.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 48.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 27.832 |
| process:counter | cpu_seconds_total | 59.219 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53272576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 195.656 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 56848384.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 538.315 |
| process:knowpost | cpu_seconds_total | 639.031 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 105185280.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 43.303 |
| process:relation | cpu_seconds_total | 49.906 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56111104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 0.266 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36343808.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 59.422 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 45686784.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 206637537.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 542.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 700472.000 |
| redis | keyspace_hits | 1005000083.000 |
| redis | keyspace_misses | 6012278.000 |
| redis | net_input_bytes | 42955811064.000 |
| redis | net_output_bytes | 189056263306.000 |
| redis | ops_per_sec | 14216.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48410.000 |
| redis | used_memory_bytes | 120938160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
