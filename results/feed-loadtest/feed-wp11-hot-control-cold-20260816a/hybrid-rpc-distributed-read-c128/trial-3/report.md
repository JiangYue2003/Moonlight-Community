# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:32:11+08:00
- 采样时长：1m0.0488122s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177859 | 177859 | 0 | 0 | 2962.54 | 42.558 | 51.088 | 55.065 | 63.693 | 114.276 |

## Redis 本轮边界增量

- Commands：1488367；input：339106052 bytes；output：1530470325 bytes
- Hits/Misses：8365368/95；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27068；safety epoch：368 -> 368

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.531 |
| client:loadtest | cpu_percent_total | 72.493 |
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
| docker:zg-canal | cpu_percent | 2.650 |
| docker:zg-canal | memory_percent | 7.860 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 2.410 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 5.770 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 189.270 |
| docker:zg-kafka | memory_percent | 8.230 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 60.400 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 85.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 5458868.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 8.000 |
| process:counter | cpu_percent | 9.964 |
| process:counter | cpu_seconds_total | 96.625 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44916736.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.779 |
| process:gateway | cpu_seconds_total | 0.453 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38539264.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 305.134 |
| process:knowpost | cpu_seconds_total | 3975.938 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 100184064.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 285.095 |
| process:relation | cpu_seconds_total | 3737.656 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 94146560.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.764 |
| process:search | cpu_seconds_total | 0.844 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36814848.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.544 |
| process:user-storage | cpu_seconds_total | 0.719 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35340288.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 41655394.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 368.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699596.000 |
| redis | keyspace_hits | 213296810.000 |
| redis | keyspace_misses | 89054.000 |
| redis | net_input_bytes | 8864185627.000 |
| redis | net_output_bytes | 39019717301.000 |
| redis | ops_per_sec | 27068.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18770.000 |
| redis | used_memory_bytes | 107337616.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
