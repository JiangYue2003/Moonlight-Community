# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:52:16+08:00
- 采样时长：1m0.0226977s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 147810 | 147810 | 0 | 0 | 2463.02 | 12.631 | 15.568 | 16.820 | 20.217 | 34.500 |

## Redis 本轮边界增量

- Commands：1242122；input：282388196 bytes；output：1271978216 bytes
- Hits/Misses：6947338/122；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：22525；safety epoch：384 -> 384

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.597 |
| client:loadtest | cpu_percent_total | 57.556 |
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
| docker:zg-canal | cpu_percent | 0.430 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.570 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 7.770 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 268.010 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 79.910 |
| docker:zg-zk | memory_percent | 1.150 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 8755878.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 62.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.924 |
| process:counter | cpu_seconds_total | 165.438 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44265472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 249.473 |
| process:gateway | cpu_seconds_total | 146.047 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 51429376.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 262.051 |
| process:knowpost | cpu_seconds_total | 6961.172 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 85487616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 249.108 |
| process:relation | cpu_seconds_total | 6445.047 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 96219136.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.391 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36745216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 131.542 |
| process:user-storage | cpu_seconds_total | 78.781 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 54738944.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 69282593.000 |
| redis | connected_clients | 248.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 384.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699088.000 |
| redis | keyspace_hits | 402873544.000 |
| redis | keyspace_misses | 92081.000 |
| redis | net_input_bytes | 16365665685.000 |
| redis | net_output_bytes | 73197637522.000 |
| redis | ops_per_sec | 22525.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19975.000 |
| redis | used_memory_bytes | 107229392.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
