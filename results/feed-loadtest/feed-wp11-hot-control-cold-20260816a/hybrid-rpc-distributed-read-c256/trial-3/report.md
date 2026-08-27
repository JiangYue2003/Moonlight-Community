# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:35:56+08:00
- 采样时长：1m0.0806013s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182205 | 182205 | 0 | 0 | 3033.41 | 71.816 | 139.021 | 168.762 | 239.508 | 514.480 |

## Redis 本轮边界增量

- Commands：1523519；input：347316778 bytes；output：1567834071 bytes
- Hits/Misses：8569597/128；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27788；safety epoch：371 -> 371

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.420 |
| client:loadtest | cpu_percent_total | 70.712 |
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
| docker:zg-canal | cpu_percent | 2.320 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 80.000 |
| docker:zg-es | cpu_percent | 3.500 |
| docker:zg-es | memory_percent | 11.810 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.080 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 219.780 |
| docker:zg-kafka | memory_percent | 8.280 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 54.010 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 6090075.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.816 |
| process:counter | cpu_seconds_total | 110.016 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44572672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.453 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 37670912.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 302.819 |
| process:knowpost | cpu_seconds_total | 4531.031 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 108871680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 275.305 |
| process:relation | cpu_seconds_total | 4251.922 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 103817216.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.321 |
| process:search | cpu_seconds_total | 1.047 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36831232.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.770 |
| process:user-storage | cpu_seconds_total | 0.797 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 35495936.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 46941155.000 |
| redis | connected_clients | 236.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 371.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699425.000 |
| redis | keyspace_hits | 242932328.000 |
| redis | keyspace_misses | 89486.000 |
| redis | net_input_bytes | 10066417515.000 |
| redis | net_output_bytes | 44441773053.000 |
| redis | ops_per_sec | 27788.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 18995.000 |
| redis | used_memory_bytes | 106986448.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
