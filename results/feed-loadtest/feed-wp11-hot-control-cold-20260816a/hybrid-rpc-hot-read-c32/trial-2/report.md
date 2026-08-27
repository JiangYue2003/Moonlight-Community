# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:08:28+08:00
- 采样时长：1m0.0203028s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183011 | 183011 | 0 | 0 | 3049.81 | 10.352 | 12.397 | 13.137 | 14.973 | 97.487 |

## Redis 本轮边界增量

- Commands：1524756；input：348701098 bytes；output：1574627613 bytes
- Hits/Misses：8601829/78；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27338；safety epoch：349 -> 349

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.528 |
| client:loadtest | cpu_percent_total | 72.449 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.040 |
| docker:zg-es | memory_percent | 11.770 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.670 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 217.360 |
| docker:zg-kafka | memory_percent | 8.110 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 30.410 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1353362.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 42.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.600 |
| process:counter | cpu_seconds_total | 14.750 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 40271872.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.188 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37044224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 329.416 |
| process:knowpost | cpu_seconds_total | 384.797 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 76234752.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.074 |
| process:relation | cpu_seconds_total | 355.422 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 67526656.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.320 |
| process:search | cpu_seconds_total | 0.219 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 37048320.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.029 |
| process:user-storage | cpu_seconds_total | 0.250 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34463744.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7376404.000 |
| redis | connected_clients | 61.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 349.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 702689.000 |
| redis | keyspace_hits | 20571954.000 |
| redis | keyspace_misses | 86415.000 |
| redis | net_input_bytes | 1047374489.000 |
| redis | net_output_bytes | 3748288582.000 |
| redis | ops_per_sec | 27338.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 17347.000 |
| redis | used_memory_bytes | 103537528.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
