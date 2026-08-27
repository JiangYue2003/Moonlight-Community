# Feed 压测报告：hybrid / rpc / deep-page-c64

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:42:11+08:00
- 采样时长：1m0.0229228s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 174667 | 174667 | 0 | 0 | 2910.57 | 20.667 | 24.866 | 27.068 | 90.107 | 134.290 |

## Redis 本轮边界增量

- Commands：1462780；input：418867642 bytes；output：1914706451 bytes
- Hits/Misses：10660613/164；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26021；safety epoch：376 -> 376

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.077 |
| client:loadtest | cpu_percent_total | 65.235 |
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
| docker:zg-canal | cpu_percent | 2.530 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.390 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 5.610 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 206.760 |
| docker:zg-kafka | memory_percent | 8.310 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 53.330 |
| docker:zg-zk | memory_percent | 1.170 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 7138544.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 70.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.594 |
| process:counter | cpu_seconds_total | 131.062 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44380160.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.772 |
| process:gateway | cpu_seconds_total | 0.531 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37228544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 297.910 |
| process:knowpost | cpu_seconds_total | 5476.812 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 95252480.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 257.833 |
| process:relation | cpu_seconds_total | 5092.766 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 90791936.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.172 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36732928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.938 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35086336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 55721910.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 376.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699198.000 |
| redis | keyspace_hits | 306815267.000 |
| redis | keyspace_misses | 90483.000 |
| redis | net_input_bytes | 12577590534.000 |
| redis | net_output_bytes | 55915913013.000 |
| redis | ops_per_sec | 26021.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19370.000 |
| redis | used_memory_bytes | 106920856.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
