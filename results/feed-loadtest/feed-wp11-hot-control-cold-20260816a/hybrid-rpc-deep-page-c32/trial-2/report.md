# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:38:26+08:00
- 采样时长：1m0.02234s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 179231 | 179231 | 0 | 0 | 2986.78 | 10.307 | 12.499 | 13.425 | 16.909 | 112.711 |

## Redis 本轮边界增量

- Commands：1499175；input：429686241 bytes；output：1964701539 bytes
- Hits/Misses：10939037/143；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27062；safety epoch：373 -> 373

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.027 |
| client:loadtest | cpu_percent_total | 64.429 |
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
| docker:zg-canal | cpu_percent | 1.660 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.990 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.720 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 181.220 |
| docker:zg-kafka | memory_percent | 8.190 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 51.920 |
| docker:zg-zk | memory_percent | 1.080 |
| docker:zg-zk | pids | 93.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6513816.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 10.709 |
| process:counter | cpu_seconds_total | 118.266 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 43794432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.793 |
| process:gateway | cpu_seconds_total | 0.469 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37257216.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 314.823 |
| process:knowpost | cpu_seconds_total | 4916.531 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 86110208.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 268.320 |
| process:relation | cpu_seconds_total | 4588.469 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 88506368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.793 |
| process:search | cpu_seconds_total | 1.109 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36708352.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.780 |
| process:user-storage | cpu_seconds_total | 0.875 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35045376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 50489111.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 373.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699333.000 |
| redis | keyspace_hits | 268755366.000 |
| redis | keyspace_misses | 89886.000 |
| redis | net_input_bytes | 11081403155.000 |
| redis | net_output_bytes | 49079894071.000 |
| redis | ops_per_sec | 27062.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19145.000 |
| redis | used_memory_bytes | 106937688.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
