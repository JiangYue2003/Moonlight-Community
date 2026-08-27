# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:37:11+08:00
- 采样时长：1m0.0385202s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177941 | 177941 | 0 | 0 | 2964.39 | 10.327 | 12.693 | 13.723 | 24.807 | 99.098 |

## Redis 本轮边界增量

- Commands：1489271；input：426651832 bytes；output：1950581618 bytes
- Hits/Misses：10860364/127；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27160；safety epoch：372 -> 372

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.153 |
| client:loadtest | cpu_percent_total | 66.442 |
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
| docker:zg-canal | cpu_percent | 2.990 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.880 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 5.380 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 226.090 |
| docker:zg-kafka | memory_percent | 8.280 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 58.430 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6301573.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 61.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.042 |
| process:counter | cpu_seconds_total | 114.000 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44736512.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.453 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37257216.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 313.429 |
| process:knowpost | cpu_seconds_total | 4724.859 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 85917696.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 265.613 |
| process:relation | cpu_seconds_total | 4420.594 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 88641536.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.062 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36700160.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.812 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35061760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 48712217.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 372.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699383.000 |
| redis | keyspace_hits | 255821639.000 |
| redis | keyspace_misses | 89690.000 |
| redis | net_input_bytes | 10573046912.000 |
| redis | net_output_bytes | 46756845531.000 |
| redis | ops_per_sec | 27160.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19070.000 |
| redis | used_memory_bytes | 106945408.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
