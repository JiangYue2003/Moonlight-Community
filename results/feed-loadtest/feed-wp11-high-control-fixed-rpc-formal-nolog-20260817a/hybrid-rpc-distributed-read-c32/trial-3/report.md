# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:14:34+08:00
- 采样时长：1m0.0292721s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 369762 | 369762 | 0 | 0 | 6162.47 | 4.964 | 6.873 | 7.553 | 9.161 | 17.336 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 148612 | 0.402 |
| mysql | 86 | 0.000 |
| redis | 1109372 | 3.000 |
| relation | 369762 | 1.000 |

- Cold compute：369762（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 369762 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 369762 | 0.814 |
| counter | 148612 | 1.186 |
| hydrate | 369762 | 0.899 |
| inbox | 369762 | 0.923 |
| merge_dedup | 369762 | 0.002 |
| relation | 369762 | 1.786 |
| route | 369762 | 0.485 |
| total | 369762 | 4.927 |

## Redis 本轮边界增量

- Commands：3649325；input：310165427 bytes；output：691937118 bytes
- Hits/Misses：5749276/708315；run hit rate：89.03%
- Evicted/Rejected：0/0；ops/s max：66305；safety epoch：600 -> 600

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.367 |
| client:loadtest | cpu_percent_total | 85.870 |
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
| docker:zg-canal | cpu_percent | 3.620 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.530 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.100 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 257.910 |
| docker:zg-kafka | memory_percent | 8.770 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 64.500 |
| docker:zg-zk | memory_percent | 1.800 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 28834125.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 9.000 |
| process:counter | cpu_percent | 101.609 |
| process:counter | cpu_seconds_total | 346.547 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54730752.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.068 |
| process:gateway | cpu_seconds_total | 495.156 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 48619520.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 388.545 |
| process:knowpost | cpu_seconds_total | 2233.375 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 88141824.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 289.939 |
| process:relation | cpu_seconds_total | 1739.141 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 59936768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.777 |
| process:search | cpu_seconds_total | 0.734 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36057088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.788 |
| process:user-storage | cpu_seconds_total | 183.984 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 46432256.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 287715350.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 600.000 |
| redis | hit_rate | 0.986 |
| redis | keys | 711882.000 |
| redis | keyspace_hits | 1126042210.000 |
| redis | keyspace_misses | 16261317.000 |
| redis | net_input_bytes | 52658979867.000 |
| redis | net_output_bytes | 207488199102.000 |
| redis | ops_per_sec | 66305.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53713.000 |
| redis | used_memory_bytes | 126367200.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
