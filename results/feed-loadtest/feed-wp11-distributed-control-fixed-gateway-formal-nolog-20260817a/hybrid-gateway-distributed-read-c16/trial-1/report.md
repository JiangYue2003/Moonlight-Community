# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:00:55+08:00
- 采样时长：1m0.0157346s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 122909 | 122909 | 0 | 0 | 2048.25 | 7.308 | 12.299 | 14.152 | 17.736 | 28.527 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 13227 | 0.108 |
| mysql | 59 | 0.000 |
| redis | 368786 | 3.000 |
| relation | 122909 | 1.000 |

- Cold compute：122909（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 122909 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 122909 | 1.214 |
| counter | 13227 | 1.820 |
| hydrate | 122909 | 1.341 |
| inbox | 122909 | 1.365 |
| merge_dedup | 122909 | 0.004 |
| relation | 122909 | 2.432 |
| route | 122909 | 0.200 |
| total | 122909 | 6.575 |

## Redis 本轮边界增量

- Commands：1066357；input：116738423 bytes；output：339677414 bytes
- Hits/Misses：2336610/133174；run hit rate：94.61%
- Evicted/Rejected：0/0；ops/s max：24766；safety epoch：586 -> 586

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.598 |
| client:loadtest | cpu_percent_total | 57.563 |
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
| docker:zg-canal | cpu_percent | 2.810 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.290 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.640 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 182.440 |
| docker:zg-kafka | memory_percent | 8.770 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.220 |
| docker:zg-zk | memory_percent | 1.720 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 26981142.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 40.986 |
| process:counter | cpu_seconds_total | 58.578 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53944320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 243.913 |
| process:gateway | cpu_seconds_total | 189.266 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51212288.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 281.271 |
| process:knowpost | cpu_seconds_total | 917.016 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 71696384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 223.666 |
| process:relation | cpu_seconds_total | 760.312 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57507840.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.098 |
| process:search | cpu_seconds_total | 0.453 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35971072.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.995 |
| process:user-storage | cpu_seconds_total | 71.594 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 55750656.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 268822156.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 586.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 731579.000 |
| redis | keyspace_hits | 1096382398.000 |
| redis | keyspace_misses | 12838996.000 |
| redis | net_input_bytes | 51048143906.000 |
| redis | net_output_bytes | 203871886293.000 |
| redis | ops_per_sec | 24766.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52894.000 |
| redis | used_memory_bytes | 144765920.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
