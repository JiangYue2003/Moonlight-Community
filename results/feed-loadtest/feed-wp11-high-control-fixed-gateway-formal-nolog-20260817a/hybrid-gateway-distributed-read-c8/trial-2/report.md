# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:17:28+08:00
- 采样时长：1m0.0105849s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 84175 | 84175 | 0 | 0 | 1402.80 | 3.317 | 11.573 | 13.045 | 15.880 | 28.223 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 62336 | 0.741 |
| mysql | 49 | 0.001 |
| redis | 252574 | 3.001 |
| relation | 84175 | 1.000 |

- Cold compute：84175（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 84175 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 84175 | 0.576 |
| counter | 62336 | 1.061 |
| hydrate | 84175 | 0.596 |
| inbox | 84175 | 0.596 |
| merge_dedup | 84175 | 0.002 |
| relation | 84175 | 1.789 |
| route | 84175 | 0.799 |
| total | 84175 | 4.377 |

## Redis 本轮边界增量

- Commands：1016053；input：77810651 bytes；output：162242371 bytes
- Hits/Misses：1446855/161291；run hit rate：89.97%
- Evicted/Rejected：0/0；ops/s max：26911；safety epoch：602 -> 602

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.496 |
| client:loadtest | cpu_percent_total | 39.941 |
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
| docker:zg-canal | cpu_percent | 2.980 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.760 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.610 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 181.370 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 53.110 |
| docker:zg-zk | memory_percent | 1.750 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29079953.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 105.996 |
| process:counter | cpu_seconds_total | 451.266 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54775808.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 191.915 |
| process:gateway | cpu_seconds_total | 713.344 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51216384.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 282.091 |
| process:knowpost | cpu_seconds_total | 2553.453 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 79638528.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 174.749 |
| process:relation | cpu_seconds_total | 1950.172 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57946112.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.844 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36069376.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 89.694 |
| process:user-storage | cpu_seconds_total | 280.188 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 57552896.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 290613394.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 602.000 |
| redis | hit_rate | 0.986 |
| redis | keys | 709383.000 |
| redis | keyspace_hits | 1130194052.000 |
| redis | keyspace_misses | 16731851.000 |
| redis | net_input_bytes | 52884193180.000 |
| redis | net_output_bytes | 207959476972.000 |
| redis | ops_per_sec | 26911.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53887.000 |
| redis | used_memory_bytes | 122088968.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
