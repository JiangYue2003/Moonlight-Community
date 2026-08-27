# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816a`
- 开始时间：2026-08-16T15:54:46+08:00
- 采样时长：1m0.0172691s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 144533 | 144533 | 0 | 0 | 2408.58 | 12.520 | 15.552 | 17.261 | 25.620 | 102.261 |

## Redis 本轮边界增量

- Commands：1215906；input：276220594 bytes；output：1243805514 bytes
- Hits/Misses：6793319/122；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：21846；safety epoch：386 -> 386

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.586 |
| client:loadtest | cpu_percent_total | 57.379 |
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
| docker:zg-canal | cpu_percent | 3.080 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.980 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 11.220 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 264.800 |
| docker:zg-kafka | memory_percent | 8.400 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 87.190 |
| docker:zg-zk | memory_percent | 1.220 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3099.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3099.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9104349.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.248 |
| process:counter | cpu_seconds_total | 173.875 |
| process:counter | pid | 24156.000 |
| process:counter | process_start_ms | 1786863853130.000 |
| process:counter | rss_bytes | 44273664.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 255.408 |
| process:gateway | cpu_seconds_total | 441.203 |
| process:gateway | pid | 24784.000 |
| process:gateway | process_start_ms | 1786863868506.000 |
| process:gateway | rss_bytes | 51974144.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 264.594 |
| process:knowpost | cpu_seconds_total | 7299.359 |
| process:knowpost | pid | 30108.000 |
| process:knowpost | process_start_ms | 1786863860928.000 |
| process:knowpost | rss_bytes | 84492288.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 253.229 |
| process:relation | cpu_seconds_total | 6747.062 |
| process:relation | pid | 29116.000 |
| process:relation | process_start_ms | 1786863856781.000 |
| process:relation | rss_bytes | 88850432.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 1.500 |
| process:search | pid | 20388.000 |
| process:search | process_start_ms | 1786863865032.000 |
| process:search | rss_bytes | 36749312.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 132.878 |
| process:user-storage | cpu_seconds_total | 231.469 |
| process:user-storage | pid | 12272.000 |
| process:user-storage | process_start_ms | 1786863849670.000 |
| process:user-storage | rss_bytes | 56131584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 72214806.000 |
| redis | connected_clients | 248.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 386.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 699036.000 |
| redis | keyspace_hits | 419225297.000 |
| redis | keyspace_misses | 92381.000 |
| redis | net_input_bytes | 17030927318.000 |
| redis | net_output_bytes | 76191605546.000 |
| redis | ops_per_sec | 21846.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20125.000 |
| redis | used_memory_bytes | 107305640.000 |

## 缺失指标

- feed_metrics
- environment_evidence

## 说明

- SLA values are reference lines, not pass/fail gates.
