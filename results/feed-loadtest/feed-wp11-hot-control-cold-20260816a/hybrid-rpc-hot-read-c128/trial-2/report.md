# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:15:57+08:00
- 采样时长：1m0.0479093s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 181653 | 181653 | 0 | 0 | 3025.85 | 42.228 | 48.961 | 50.738 | 54.951 | 87.944 |

## Redis 本轮边界增量

- Commands：1513461；input：346121603 bytes；output：1562939684 bytes
- Hits/Misses：8537964/117；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26968；safety epoch：355 -> 355

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.528 |
| client:loadtest | cpu_percent_total | 72.442 |
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
| docker:zg-canal | cpu_percent | 2.940 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.830 |
| docker:zg-es | memory_percent | 11.770 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.490 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 203.600 |
| docker:zg-kafka | memory_percent | 8.160 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 58.120 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2651882.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.831 |
| process:counter | cpu_seconds_total | 40.641 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 41140224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.328 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38187008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 311.965 |
| process:knowpost | cpu_seconds_total | 1513.125 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 105177088.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 280.105 |
| process:relation | cpu_seconds_total | 1426.609 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 96976896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.621 |
| process:search | cpu_seconds_total | 0.484 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36634624.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 0.375 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34705408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18201546.000 |
| redis | connected_clients | 179.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 355.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 701165.000 |
| redis | keyspace_hits | 81511213.000 |
| redis | keyspace_misses | 87264.000 |
| redis | net_input_bytes | 3519412442.000 |
| redis | net_output_bytes | 14904080273.000 |
| redis | ops_per_sec | 26968.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 17796.000 |
| redis | used_memory_bytes | 106712264.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
