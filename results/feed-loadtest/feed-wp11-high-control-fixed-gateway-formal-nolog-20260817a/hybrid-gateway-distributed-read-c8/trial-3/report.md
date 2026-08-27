# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:18:39+08:00
- 采样时长：1m0.0074258s
- 并发：8
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 76701 | 76701 | 0 | 0 | 1278.31 | 6.589 | 10.126 | 11.205 | 13.487 | 21.735 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 58442 | 0.762 |
| mysql | 46 | 0.001 |
| redis | 230149 | 3.001 |
| relation | 76701 | 1.000 |

- Cold compute：76701（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 76701 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 76701 | 0.642 |
| counter | 58442 | 1.155 |
| hydrate | 76701 | 0.669 |
| inbox | 76701 | 0.678 |
| merge_dedup | 76701 | 0.002 |
| relation | 76701 | 1.963 |
| route | 76701 | 0.894 |
| total | 76701 | 4.867 |

## Redis 本轮边界增量

- Commands：939561；input：71536294 bytes；output：148193740 bytes
- Hits/Misses：1326446/146940；run hit rate：90.03%
- Evicted/Rejected：0/0；ops/s max：19323；safety epoch：603 -> 603

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.327 |
| client:loadtest | cpu_percent_total | 37.235 |
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
| docker:zg-canal | cpu_percent | 3.030 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 2.110 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.100 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 157.370 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 48.230 |
| docker:zg-zk | memory_percent | 1.720 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29161308.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 95.894 |
| process:counter | cpu_seconds_total | 507.453 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54718464.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 184.827 |
| process:gateway | cpu_seconds_total | 812.984 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50860032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 260.028 |
| process:knowpost | cpu_seconds_total | 2707.125 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 79294464.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 180.961 |
| process:relation | cpu_seconds_total | 2050.312 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58413056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 0.891 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36089856.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 80.608 |
| process:user-storage | cpu_seconds_total | 320.188 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 50221056.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 291617847.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 603.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 708457.000 |
| redis | keyspace_hits | 1131600533.000 |
| redis | keyspace_misses | 16887443.000 |
| redis | net_input_bytes | 52960522824.000 |
| redis | net_output_bytes | 208116575099.000 |
| redis | ops_per_sec | 19323.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53958.000 |
| redis | used_memory_bytes | 121251064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
