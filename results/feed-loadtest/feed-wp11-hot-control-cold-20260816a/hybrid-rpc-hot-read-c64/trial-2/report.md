# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:12:12+08:00
- 采样时长：1m0.025919s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183908 | 183908 | 0 | 0 | 3064.49 | 20.813 | 24.222 | 25.393 | 28.640 | 46.674 |

## Redis 本轮边界增量

- Commands：1531728；input：350382006 bytes；output：1582327637 bytes
- Hits/Misses：8643943/123；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27238；safety epoch：352 -> 352

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.430 |
| client:loadtest | cpu_percent_total | 70.881 |
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
| docker:zg-canal | cpu_percent | 2.770 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.620 |
| docker:zg-es | memory_percent | 11.770 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 5.390 |
| docker:zg-etcd | memory_percent | 0.250 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 186.670 |
| docker:zg-kafka | memory_percent | 8.130 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 69.070 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 2002419.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 67.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 9.288 |
| process:counter | cpu_seconds_total | 27.469 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 40591360.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.553 |
| process:gateway | cpu_seconds_total | 0.312 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38195200.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 309.795 |
| process:knowpost | cpu_seconds_total | 948.891 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 88043520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 293.308 |
| process:relation | cpu_seconds_total | 891.203 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 78737408.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 0.266 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36651008.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.299 |
| process:user-storage | cpu_seconds_total | 0.359 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34545664.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12787235.000 |
| redis | connected_clients | 101.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 352.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 701850.000 |
| redis | keyspace_hits | 51033614.000 |
| redis | keyspace_misses | 86853.000 |
| redis | net_input_bytes | 2283050330.000 |
| redis | net_output_bytes | 9324717522.000 |
| redis | ops_per_sec | 27238.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 17571.000 |
| redis | used_memory_bytes | 104559968.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
