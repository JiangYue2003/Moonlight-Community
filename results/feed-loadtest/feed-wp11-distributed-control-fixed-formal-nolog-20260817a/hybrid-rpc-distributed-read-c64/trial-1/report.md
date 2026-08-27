# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-fixed-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:56:45+08:00
- 采样时长：1m0.0266818s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 333401 | 333401 | 0 | 0 | 5556.24 | 9.597 | 20.856 | 25.444 | 35.650 | 86.785 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 14352 | 0.043 |
| mysql | 176 | 0.001 |
| redis | 1000379 | 3.001 |
| relation | 333401 | 1.000 |

- Cold compute：333401（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 333401 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 333401 | 2.327 |
| counter | 14352 | 3.445 |
| hydrate | 333401 | 2.626 |
| inbox | 333401 | 2.701 |
| merge_dedup | 333401 | 0.004 |
| relation | 333401 | 3.383 |
| route | 333401 | 0.153 |
| total | 333401 | 11.211 |

## Redis 本轮边界增量

- Commands：2643674；input：304761959 bytes；output：915018053 bytes
- Hits/Misses：6201626/361300；run hit rate：94.49%
- Evicted/Rejected：0/0；ops/s max：58358；safety epoch：583 -> 583

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.322 |
| client:loadtest | cpu_percent_total | 85.144 |
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
| docker:zg-canal | cpu_percent | 0.230 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.760 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 187.730 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 57.140 |
| docker:zg-zk | memory_percent | 1.540 |
| docker:zg-zk | pids | 88.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25948425.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 11.000 |
| process:counter | cpu_percent | 38.648 |
| process:counter | cpu_seconds_total | 24.594 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53784576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.546 |
| process:gateway | cpu_seconds_total | 52.453 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51564544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 348.189 |
| process:knowpost | cpu_seconds_total | 341.266 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 75038720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 308.512 |
| process:relation | cpu_seconds_total | 281.797 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 59543552.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.188 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35946496.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 21.078 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 46686208.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 260517471.000 |
| redis | connected_clients | 185.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 583.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 743289.000 |
| redis | keyspace_hits | 1077186857.000 |
| redis | keyspace_misses | 11720851.000 |
| redis | net_input_bytes | 50096690596.000 |
| redis | net_output_bytes | 201038494769.000 |
| redis | ops_per_sec | 58358.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52644.000 |
| redis | used_memory_bytes | 159543088.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
