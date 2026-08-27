# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:39:41+08:00
- 采样时长：1m0.0213973s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 179183 | 179183 | 0 | 0 | 2985.95 | 10.310 | 12.510 | 13.519 | 17.169 | 104.914 |

## Redis 本轮边界增量

- Commands：1498782；input：429577396 bytes；output：1964169510 bytes
- Hits/Misses：10936074/179；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27195；safety epoch：374 -> 374

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.089 |
| client:loadtest | cpu_percent_total | 65.419 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.100 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.750 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 184.870 |
| docker:zg-kafka | memory_percent | 8.290 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 59.680 |
| docker:zg-zk | memory_percent | 1.190 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6725036.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 41.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.308 |
| process:counter | cpu_seconds_total | 122.359 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44081152.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.547 |
| process:gateway | cpu_seconds_total | 0.500 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37285888.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 313.900 |
| process:knowpost | cpu_seconds_total | 5106.656 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 86016000.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 276.978 |
| process:relation | cpu_seconds_total | 4758.672 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 88387584.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.777 |
| process:search | cpu_seconds_total | 1.156 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36732928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.922 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35061760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 52257267.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 374.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699286.000 |
| redis | keyspace_hits | 281626622.000 |
| redis | keyspace_misses | 90089.000 |
| redis | net_input_bytes | 11587292201.000 |
| redis | net_output_bytes | 51391713960.000 |
| redis | ops_per_sec | 27195.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19220.000 |
| redis | used_memory_bytes | 107056336.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
