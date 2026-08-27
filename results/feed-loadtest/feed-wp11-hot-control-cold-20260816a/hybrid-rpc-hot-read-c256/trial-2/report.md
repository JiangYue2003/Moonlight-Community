# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:19:41+08:00
- 采样时长：1m0.0921405s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182595 | 182595 | 0 | 0 | 3039.39 | 71.675 | 138.488 | 167.332 | 236.328 | 618.462 |

## Redis 本轮边界增量

- Commands：1520907；input：347885641 bytes；output：1571040575 bytes
- Hits/Misses：8582250/105；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27294；safety epoch：358 -> 358

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.578 |
| client:loadtest | cpu_percent_total | 73.247 |
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
| docker:zg-canal | cpu_percent | 2.560 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.150 |
| docker:zg-es | memory_percent | 11.770 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.060 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 184.780 |
| docker:zg-kafka | memory_percent | 8.170 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 59.540 |
| docker:zg-zk | memory_percent | 1.200 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 3299157.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.195 |
| process:counter | cpu_seconds_total | 53.562 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 41312256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.359 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38363136.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 323.066 |
| process:knowpost | cpu_seconds_total | 2084.016 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 106344448.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 296.015 |
| process:relation | cpu_seconds_total | 1966.641 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 104607744.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.562 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36749312.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.422 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34775040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 23595656.000 |
| redis | connected_clients | 230.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 358.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 700650.000 |
| redis | keyspace_hits | 111880000.000 |
| redis | keyspace_misses | 87682.000 |
| redis | net_input_bytes | 4751301246.000 |
| redis | net_output_bytes | 20463502916.000 |
| redis | ops_per_sec | 27294.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18020.000 |
| redis | used_memory_bytes | 107100688.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
