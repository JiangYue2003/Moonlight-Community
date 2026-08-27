# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:08:53+08:00
- 采样时长：10.0104788s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 64802 | 64802 | 0 | 0 | 6476.53 | 9.621 | 13.323 | 14.759 | 17.975 | 24.923 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 28643 | 0.442 |
| mysql | 0 | 0.000 |
| redis | 194406 | 3.000 |
| relation | 64802 | 1.000 |

- Cold compute：64802（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 64802 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 64802 | 1.777 |
| counter | 28643 | 2.440 |
| hydrate | 64802 | 1.907 |
| inbox | 64802 | 1.996 |
| merge_dedup | 64802 | 0.002 |
| relation | 64802 | 2.765 |
| route | 64802 | 1.089 |
| total | 64802 | 9.556 |

## Redis 本轮边界增量

- Commands：646636；input：54318864 bytes；output：121493778 bytes
- Hits/Misses：1020017/124103；run hit rate：89.15%
- Evicted/Rejected：0/0；ops/s max：69390；safety epoch：594 -> 594

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.839 |
| client:loadtest | cpu_percent_total | 77.419 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.560 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.850 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 100.690 |
| docker:zg-kafka | memory_percent | 8.330 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 41.670 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 84.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27584235.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 43.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 98.954 |
| process:counter | cpu_seconds_total | 168.562 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54501376.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.547 |
| process:gateway | cpu_seconds_total | 426.312 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 46194688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 357.671 |
| process:knowpost | cpu_seconds_total | 1474.328 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 89894912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 288.876 |
| process:relation | cpu_seconds_total | 1179.062 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 59944960.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 0.578 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35971072.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 156.047 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 43196416.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 274993783.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 594.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 718121.000 |
| redis | keyspace_hits | 1106423696.000 |
| redis | keyspace_misses | 13868339.000 |
| redis | net_input_bytes | 51591948940.000 |
| redis | net_output_bytes | 205141916618.000 |
| redis | ops_per_sec | 69390.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53322.000 |
| redis | used_memory_bytes | 134652736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
