# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-fixed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:55:14+08:00
- 采样时长：10.0130642s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 67225 | 67225 | 0 | 0 | 6717.94 | 9.258 | 12.853 | 14.092 | 16.867 | 31.694 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.036 |
| mysql | 0 | 0.000 |
| redis | 201675 | 3.000 |
| relation | 67225 | 1.000 |

- Cold compute：67225（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 67225 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 67225 | 1.798 |
| counter | 2400 | 2.536 |
| hydrate | 67225 | 2.168 |
| inbox | 67225 | 2.245 |
| merge_dedup | 67225 | 0.004 |
| relation | 67225 | 2.880 |
| route | 67225 | 0.095 |
| total | 67225 | 9.209 |

## Redis 本轮边界增量

- Commands：529783；input：61307886 bytes；output：184193524 bytes
- Hits/Misses：1246319/72862；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：56352；safety epoch：580 -> 580

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.920 |
| client:loadtest | cpu_percent_total | 94.720 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.720 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.330 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 240.990 |
| docker:zg-kafka | memory_percent | 8.790 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 75.060 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25509970.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 61.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 24.307 |
| process:counter | cpu_seconds_total | 7.562 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 51625984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.531 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 37437440.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 332.370 |
| process:knowpost | cpu_seconds_total | 79.781 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 74637312.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 269.597 |
| process:relation | cpu_seconds_total | 62.266 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 55468032.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35483648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.047 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 34447360.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 256939411.000 |
| redis | connected_clients | 168.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 580.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 748916.000 |
| redis | keyspace_hits | 1069019934.000 |
| redis | keyspace_misses | 11246000.000 |
| redis | net_input_bytes | 49690033369.000 |
| redis | net_output_bytes | 199834985229.000 |
| redis | ops_per_sec | 56352.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52503.000 |
| redis | used_memory_bytes | 164742648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
