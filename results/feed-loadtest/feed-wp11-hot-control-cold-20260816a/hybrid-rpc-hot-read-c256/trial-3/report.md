# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:20:56+08:00
- 采样时长：1m0.0724181s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182615 | 182615 | 0 | 0 | 3040.64 | 71.627 | 138.564 | 167.639 | 236.723 | 517.660 |

## Redis 本轮边界增量

- Commands：1520952；input：347912477 bytes；output：1571207902 bytes
- Hits/Misses：8583203/92；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27023；safety epoch：359 -> 359

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.505 |
| client:loadtest | cpu_percent_total | 72.074 |
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
| docker:zg-es | cpu_percent | 5.480 |
| docker:zg-es | memory_percent | 11.790 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.580 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 198.990 |
| docker:zg-kafka | memory_percent | 8.170 |
| docker:zg-kafka | pids | 128.000 |
| docker:zg-zk | cpu_percent | 54.050 |
| docker:zg-zk | memory_percent | 0.970 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3515171.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.055 |
| process:counter | cpu_seconds_total | 57.891 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 41537536.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.359 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38412288.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 303.033 |
| process:knowpost | cpu_seconds_total | 2272.141 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 107565056.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 298.568 |
| process:relation | cpu_seconds_total | 2146.094 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 104247296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.018 |
| process:search | cpu_seconds_total | 0.609 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36753408.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.453 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34795520.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 25396256.000 |
| redis | connected_clients | 230.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 359.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 700501.000 |
| redis | keyspace_hits | 122017975.000 |
| redis | keyspace_misses | 87808.000 |
| redis | net_input_bytes | 5162528768.000 |
| redis | net_output_bytes | 22319389522.000 |
| redis | ops_per_sec | 27023.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18095.000 |
| redis | used_memory_bytes | 106931592.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
