# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:49:42+08:00
- 采样时长：1m0.0742567s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 172949 | 172949 | 0 | 0 | 2879.57 | 73.814 | 152.944 | 184.640 | 259.215 | 521.135 |

## Redis 本轮边界增量

- Commands：1449003；input：414792014 bytes；output：1895881505 bytes
- Hits/Misses：10555788/191；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26382；safety epoch：382 -> 382

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.815 |
| client:loadtest | cpu_percent_total | 61.044 |
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
| docker:zg-es | cpu_percent | 0.540 |
| docker:zg-es | memory_percent | 11.820 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.970 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 219.160 |
| docker:zg-kafka | memory_percent | 8.340 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 68.480 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 8382067.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 9.292 |
| process:counter | cpu_seconds_total | 155.875 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44470272.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.547 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37273600.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 283.447 |
| process:knowpost | cpu_seconds_total | 6610.172 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 112209920.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 285.170 |
| process:relation | cpu_seconds_total | 6122.531 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 108949504.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36745216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1.062 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34938880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 66138296.000 |
| redis | connected_clients | 238.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 382.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699012.000 |
| redis | keyspace_hits | 382557378.000 |
| redis | keyspace_misses | 91670.000 |
| redis | net_input_bytes | 15555201830.000 |
| redis | net_output_bytes | 69519927976.000 |
| redis | ops_per_sec | 26382.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 19821.000 |
| redis | used_memory_bytes | 106815272.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
