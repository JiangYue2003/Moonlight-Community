# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-control-fixed-gateway-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:10:28+08:00
- 采样时长：10.0059419s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 22754 | 22754 | 0 | 0 | 2274.40 | 7.019 | 9.676 | 10.533 | 12.371 | 16.418 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 15782 | 0.694 |
| mysql | 0 | 0.000 |
| redis | 68262 | 3.000 |
| relation | 22754 | 1.000 |

- Cold compute：22754（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 22754 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22754 | 0.946 |
| counter | 15782 | 1.380 |
| hydrate | 22754 | 0.989 |
| inbox | 22754 | 1.025 |
| merge_dedup | 22754 | 0.003 |
| relation | 22754 | 1.991 |
| route | 22754 | 0.970 |
| total | 22754 | 5.942 |

## Redis 本轮边界增量

- Commands：263177；input：20436255 bytes；output：43535422 bytes
- Hits/Misses：385457/43626；run hit rate：89.83%
- Evicted/Rejected：0/0；ops/s max：30868；safety epoch：596 -> 596

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.474 |
| client:loadtest | cpu_percent_total | 55.592 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.740 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 0.650 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 121.620 |
| docker:zg-kafka | memory_percent | 8.330 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.220 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27649537.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 107.613 |
| process:counter | cpu_seconds_total | 190.688 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54595584.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 205.967 |
| process:gateway | cpu_seconds_total | 471.656 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50909184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 291.187 |
| process:knowpost | cpu_seconds_total | 1538.719 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 81731584.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 196.675 |
| process:relation | cpu_seconds_total | 1221.797 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57913344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.625 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35987456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 87.354 |
| process:user-storage | cpu_seconds_total | 175.344 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 51183616.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 275817611.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 596.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 716235.000 |
| redis | keyspace_hits | 1107525424.000 |
| redis | keyspace_misses | 13993379.000 |
| redis | net_input_bytes | 51655598642.000 |
| redis | net_output_bytes | 205268060046.000 |
| redis | ops_per_sec | 30868.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53417.000 |
| redis | used_memory_bytes | 129590040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
