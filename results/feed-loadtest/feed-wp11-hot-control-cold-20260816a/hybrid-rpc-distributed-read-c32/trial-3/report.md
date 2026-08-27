# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:24:41+08:00
- 采样时长：1m0.019089s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 184230 | 184230 | 0 | 0 | 3070.16 | 10.297 | 12.254 | 12.944 | 14.832 | 25.618 |

## Redis 本轮边界增量

- Commands：1539468；input：351107749 bytes；output：1585242195 bytes
- Hits/Misses：8664792/108；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27536；safety epoch：362 -> 362

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.598 |
| client:loadtest | cpu_percent_total | 73.570 |
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
| docker:zg-canal | cpu_percent | 1.740 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.120 |
| docker:zg-es | memory_percent | 11.790 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.160 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 201.590 |
| docker:zg-kafka | memory_percent | 8.190 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 54.230 |
| docker:zg-zk | memory_percent | 1.130 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 4166801.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.645 |
| process:counter | cpu_seconds_total | 70.047 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 43126784.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.375 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 38318080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 299.474 |
| process:knowpost | cpu_seconds_total | 2845.750 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 86376448.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 278.424 |
| process:relation | cpu_seconds_total | 2676.484 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 88444928.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.688 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36765696.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.500 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 34897920.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 30847119.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 362.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 700117.000 |
| redis | keyspace_hits | 152624178.000 |
| redis | keyspace_misses | 88251.000 |
| redis | net_input_bytes | 6403673453.000 |
| redis | net_output_bytes | 27919118423.000 |
| redis | ops_per_sec | 27536.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18320.000 |
| redis | used_memory_bytes | 107004128.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
