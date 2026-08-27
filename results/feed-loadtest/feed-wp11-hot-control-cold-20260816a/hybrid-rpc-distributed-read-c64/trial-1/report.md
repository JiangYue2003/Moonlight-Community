# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:25:57+08:00
- 采样时长：1m0.0363326s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 184144 | 184144 | 0 | 0 | 3067.94 | 20.759 | 24.345 | 25.668 | 28.880 | 44.540 |

## Redis 本轮边界增量

- Commands：1538776；input：350945653 bytes；output：1584504129 bytes
- Hits/Misses：8660754/104；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27631；safety epoch：363 -> 363

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.691 |
| client:loadtest | cpu_percent_total | 75.059 |
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
| docker:zg-canal | cpu_percent | 3.020 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.810 |
| docker:zg-es | memory_percent | 11.790 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.700 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 226.170 |
| docker:zg-kafka | memory_percent | 8.190 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 53.860 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4383679.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 70.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.053 |
| process:counter | cpu_seconds_total | 74.547 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 43319296.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.375 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38318080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 318.006 |
| process:knowpost | cpu_seconds_total | 3035.938 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 105074688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 285.477 |
| process:relation | cpu_seconds_total | 2855.000 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 93470720.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.719 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36769792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.546 |
| process:user-storage | cpu_seconds_total | 0.547 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34934784.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 32661814.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 363.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 700002.000 |
| redis | keyspace_hits | 162807332.000 |
| redis | keyspace_misses | 88392.000 |
| redis | net_input_bytes | 6816693480.000 |
| redis | net_output_bytes | 29782242555.000 |
| redis | ops_per_sec | 27631.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18396.000 |
| redis | used_memory_bytes | 106945176.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
