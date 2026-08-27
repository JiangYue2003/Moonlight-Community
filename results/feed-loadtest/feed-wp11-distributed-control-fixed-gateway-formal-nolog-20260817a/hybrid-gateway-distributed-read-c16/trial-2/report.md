# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:02:05+08:00
- 采样时长：1m0.0151625s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 65429 | 65429 | 0 | 0 | 1090.29 | 15.956 | 25.300 | 28.155 | 34.168 | 54.667 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12220 | 0.187 |
| mysql | 78 | 0.001 |
| redis | 196365 | 3.001 |
| relation | 65429 | 1.000 |

- Cold compute：65429（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 65429 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 65429 | 2.419 |
| counter | 12220 | 3.324 |
| hydrate | 65429 | 2.590 |
| inbox | 65429 | 2.614 |
| merge_dedup | 65429 | 0.004 |
| relation | 65429 | 4.336 |
| route | 65429 | 0.628 |
| total | 65429 | 12.610 |

## Redis 本轮边界增量

- Commands：600441；input：63044655 bytes；output：181492236 bytes
- Hits/Misses：1274862/70984；run hit rate：94.73%
- Evicted/Rejected：0/0；ops/s max：14006；safety epoch：587 -> 587

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.813 |
| client:loadtest | cpu_percent_total | 45.015 |
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
| docker:zg-canal | cpu_percent | 3.360 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.800 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.880 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 163.510 |
| docker:zg-kafka | memory_percent | 8.800 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 20.920 |
| docker:zg-zk | memory_percent | 1.730 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27054232.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 67.277 |
| process:counter | cpu_seconds_total | 80.812 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53624832.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 207.024 |
| process:gateway | cpu_seconds_total | 307.312 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51273728.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 255.326 |
| process:knowpost | cpu_seconds_total | 1060.172 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 70668288.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 196.617 |
| process:relation | cpu_seconds_total | 871.828 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57532416.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.453 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36012032.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 78.816 |
| process:user-storage | cpu_seconds_total | 113.609 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 52150272.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 269498117.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 587.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 729069.000 |
| redis | keyspace_hits | 1097801421.000 |
| redis | keyspace_misses | 12918080.000 |
| redis | net_input_bytes | 51118937288.000 |
| redis | net_output_bytes | 204074224124.000 |
| redis | ops_per_sec | 14006.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52964.000 |
| redis | used_memory_bytes | 142442528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
