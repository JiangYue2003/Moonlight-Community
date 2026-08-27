# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:08:17+08:00
- 采样时长：10.0058791s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 57471 | 57471 | 0 | 0 | 5745.83 | 2.662 | 3.902 | 4.372 | 5.373 | 8.566 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 27005 | 0.470 |
| mysql | 0 | 0.000 |
| redis | 172413 | 3.000 |
| relation | 57471 | 1.000 |

- Cold compute：57471（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 57471 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 57471 | 0.357 |
| counter | 27005 | 0.651 |
| hydrate | 57471 | 0.388 |
| inbox | 57471 | 0.386 |
| merge_dedup | 57471 | 0.002 |
| relation | 57471 | 1.068 |
| route | 57471 | 0.315 |
| total | 57471 | 2.534 |

## Redis 本轮边界增量

- Commands：586602；input：48771987 bytes；output：108077834 bytes
- Hits/Misses：912435/110061；run hit rate：89.24%
- Evicted/Rejected：0/0；ops/s max：62508；safety epoch：592 -> 592

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.993 |
| client:loadtest | cpu_percent_total | 95.881 |
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
| docker:zg-canal | cpu_percent | 0.280 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.050 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.410 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 253.580 |
| docker:zg-kafka | memory_percent | 8.790 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 0.240 |
| docker:zg-zk | memory_percent | 1.540 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27429086.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 125.197 |
| process:counter | cpu_seconds_total | 148.766 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54267904.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.787 |
| process:gateway | cpu_seconds_total | 426.125 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 46215168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 361.499 |
| process:knowpost | cpu_seconds_total | 1390.734 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 83865600.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 261.401 |
| process:relation | cpu_seconds_total | 1118.703 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58064896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35930112.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 156.047 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 45318144.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 273410282.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 592.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 718921.000 |
| redis | keyspace_hits | 1103966783.000 |
| redis | keyspace_misses | 13571198.000 |
| redis | net_input_bytes | 51460105311.000 |
| redis | net_output_bytes | 204850359934.000 |
| redis | ops_per_sec | 62508.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53286.000 |
| redis | used_memory_bytes | 132229400.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
