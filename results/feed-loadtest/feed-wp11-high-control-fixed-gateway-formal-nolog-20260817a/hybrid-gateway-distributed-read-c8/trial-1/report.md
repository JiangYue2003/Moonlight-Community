# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:16:18+08:00
- 采样时长：1m0.0119918s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 135023 | 135023 | 0 | 0 | 2250.27 | 3.175 | 5.921 | 6.764 | 8.367 | 13.586 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 87227 | 0.646 |
| mysql | 34 | 0.000 |
| redis | 405103 | 3.000 |
| relation | 135023 | 1.000 |

- Cold compute：135023（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 135023 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 135023 | 0.313 |
| counter | 87227 | 0.643 |
| hydrate | 135023 | 0.327 |
| inbox | 135023 | 0.317 |
| merge_dedup | 135023 | 0.002 |
| relation | 135023 | 1.191 |
| route | 135023 | 0.427 |
| total | 135023 | 2.595 |

## Redis 本轮边界增量

- Commands：1532096；input：120548924 bytes；output：257868751 bytes
- Hits/Misses：2259522/258674；run hit rate：89.73%
- Evicted/Rejected：0/0；ops/s max：34766；safety epoch：601 -> 601

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.038 |
| client:loadtest | cpu_percent_total | 48.610 |
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
| docker:zg-canal | cpu_percent | 3.150 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.110 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.740 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 242.550 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 65.630 |
| docker:zg-zk | memory_percent | 1.750 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 28986219.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 89.846 |
| process:counter | cpu_seconds_total | 396.891 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54398976.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 205.489 |
| process:gateway | cpu_seconds_total | 612.672 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50446336.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 277.113 |
| process:knowpost | cpu_seconds_total | 2397.266 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 80777216.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 191.093 |
| process:relation | cpu_seconds_total | 1850.078 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58384384.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.540 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36069376.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 103.462 |
| process:user-storage | cpu_seconds_total | 237.234 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 52445184.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 289479591.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 601.000 |
| redis | hit_rate | 0.986 |
| redis | keys | 710323.000 |
| redis | keyspace_hits | 1128586148.000 |
| redis | keyspace_misses | 16552512.000 |
| redis | net_input_bytes | 52797386002.000 |
| redis | net_output_bytes | 207779076625.000 |
| redis | ops_per_sec | 34766.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53817.000 |
| redis | used_memory_bytes | 123004696.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
