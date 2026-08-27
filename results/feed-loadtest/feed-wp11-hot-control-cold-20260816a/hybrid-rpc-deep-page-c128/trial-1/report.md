# Feed 压测报告：hybrid / rpc / deep-page-c128

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:44:41+08:00
- 采样时长：1m0.040647s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 178993 | 178993 | 0 | 0 | 2981.82 | 42.360 | 50.271 | 53.434 | 65.230 | 145.863 |

## Redis 本轮边界增量

- Commands：1497397；input：429133862 bytes；output：1962092304 bytes
- Hits/Misses：10924494/169；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26688；safety epoch：378 -> 378

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.104 |
| client:loadtest | cpu_percent_total | 65.659 |
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
| docker:zg-canal | cpu_percent | 3.050 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.990 |
| docker:zg-es | memory_percent | 11.820 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.570 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 227.520 |
| docker:zg-kafka | memory_percent | 8.340 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 59.460 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 7562930.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 9.284 |
| process:counter | cpu_seconds_total | 139.125 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 45232128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.531 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37257216.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 306.689 |
| process:knowpost | cpu_seconds_total | 5858.891 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 99717120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 286.731 |
| process:relation | cpu_seconds_total | 5441.500 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 94117888.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.219 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36732928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.953 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35106816.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 59273203.000 |
| redis | connected_clients | 238.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 378.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699136.000 |
| redis | keyspace_hits | 332667201.000 |
| redis | keyspace_misses | 90868.000 |
| redis | net_input_bytes | 13593654652.000 |
| redis | net_output_bytes | 60559188882.000 |
| redis | ops_per_sec | 26688.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19520.000 |
| redis | used_memory_bytes | 107532376.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
