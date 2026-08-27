# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:50:57+08:00
- 采样时长：1m0.0744088s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167043 | 167043 | 0 | 0 | 2781.21 | 76.712 | 157.095 | 190.769 | 268.982 | 625.505 |

## Redis 本轮边界增量

- Commands：1401721；input：400768292 bytes；output：1831191235 bytes
- Hits/Misses：10195556/157；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25402；safety epoch：383 -> 383

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.186 |
| client:loadtest | cpu_percent_total | 66.974 |
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
| docker:zg-canal | cpu_percent | 0.270 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.630 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.580 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 200.370 |
| docker:zg-kafka | memory_percent | 8.380 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 53.360 |
| docker:zg-zk | memory_percent | 1.210 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 8580775.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 9.538 |
| process:counter | cpu_seconds_total | 160.125 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44236800.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.547 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37273600.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 302.141 |
| process:knowpost | cpu_seconds_total | 6793.734 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 111878144.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 283.764 |
| process:relation | cpu_seconds_total | 6290.891 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 113901568.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36745216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.547 |
| process:user-storage | cpu_seconds_total | 1.094 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34938880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 67806309.000 |
| redis | connected_clients | 238.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 383.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698976.000 |
| redis | keyspace_hits | 394662672.000 |
| redis | keyspace_misses | 91850.000 |
| redis | net_input_bytes | 16031297124.000 |
| redis | net_output_bytes | 71694175841.000 |
| redis | ops_per_sec | 25402.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19896.000 |
| redis | used_memory_bytes | 107098848.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
