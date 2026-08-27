# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-control-fixed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:55:51+08:00
- 采样时长：10.0076647s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 23565 | 23565 | 0 | 0 | 2354.94 | 13.077 | 18.511 | 20.141 | 24.234 | 35.168 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2398 | 0.102 |
| mysql | 59 | 0.003 |
| redis | 70754 | 3.003 |
| relation | 23565 | 1.000 |

- Cold compute：23565（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 23565 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 23565 | 2.464 |
| counter | 2398 | 3.344 |
| hydrate | 23565 | 2.772 |
| inbox | 23565 | 2.858 |
| merge_dedup | 23565 | 0.004 |
| relation | 23565 | 3.890 |
| route | 23565 | 0.346 |
| total | 23565 | 12.354 |

## Redis 本轮边界增量

- Commands：197700；input：21910590 bytes；output：64742648 bytes
- Hits/Misses：445881/25619；run hit rate：94.57%
- Evicted/Rejected：0/0；ops/s max：22623；safety epoch：582 -> 582

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.933 |
| client:loadtest | cpu_percent_total | 62.921 |
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
| docker:zg-canal | cpu_percent | 1.670 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.570 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.400 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 144.970 |
| docker:zg-kafka | memory_percent | 8.840 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 0.610 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25577937.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 30.186 |
| process:counter | cpu_seconds_total | 13.828 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53190656.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 224.534 |
| process:gateway | cpu_seconds_total | 52.125 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 52064256.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 261.876 |
| process:knowpost | cpu_seconds_total | 141.219 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 71233536.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 212.843 |
| process:relation | cpu_seconds_total | 110.594 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 53714944.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35880960.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 88.237 |
| process:user-storage | cpu_seconds_total | 21.062 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 48418816.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 257535556.000 |
| redis | connected_clients | 168.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 582.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 746508.000 |
| redis | keyspace_hits | 1070305296.000 |
| redis | keyspace_misses | 11319555.000 |
| redis | net_input_bytes | 49754934472.000 |
| redis | net_output_bytes | 200021964128.000 |
| redis | ops_per_sec | 22623.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52540.000 |
| redis | used_memory_bytes | 160290000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
