# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:28:27+08:00
- 采样时长：1m0.0254752s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183946 | 183946 | 0 | 0 | 3065.11 | 20.768 | 24.396 | 25.831 | 29.126 | 44.304 |

## Redis 本轮边界增量

- Commands：1537148；input：350571193 bytes；output：1582797101 bytes
- Hits/Misses：8651432/120；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27498；safety epoch：365 -> 365

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.661 |
| client:loadtest | cpu_percent_total | 74.578 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.780 |
| docker:zg-es | memory_percent | 11.790 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.720 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 180.750 |
| docker:zg-kafka | memory_percent | 8.220 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 54.590 |
| docker:zg-zk | memory_percent | 1.220 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4817911.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 70.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.046 |
| process:counter | cpu_seconds_total | 83.453 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44597248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.438 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37564416.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 303.417 |
| process:knowpost | cpu_seconds_total | 3416.766 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 96038912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 273.146 |
| process:relation | cpu_seconds_total | 3207.719 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 90673152.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.750 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36777984.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.656 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35033088.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 36292326.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 365.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699820.000 |
| redis | keyspace_hits | 183199053.000 |
| redis | keyspace_misses | 88676.000 |
| redis | net_input_bytes | 7643534161.000 |
| redis | net_output_bytes | 33513100920.000 |
| redis | ops_per_sec | 27498.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18546.000 |
| redis | used_memory_bytes | 106794088.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
