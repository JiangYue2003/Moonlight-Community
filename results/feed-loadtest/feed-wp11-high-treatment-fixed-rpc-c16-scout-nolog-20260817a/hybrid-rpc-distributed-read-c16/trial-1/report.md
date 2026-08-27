# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-fixed-rpc-c16-scout-nolog-20260817a`
- 开始时间：2026-08-17T01:34:28+08:00
- 采样时长：10.0075139s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 76934 | 76934 | 0 | 0 | 7691.04 | 1.579 | 4.782 | 5.782 | 7.697 | 16.418 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 110106 | 1.431 |
| counter | 27006 | 0.351 |
| mysql | 27 | 0.000 |
| redis | 168759 | 2.194 |
| relation | 27006 | 0.351 |

- Cold compute：36732（0.477 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 18318 |
| l1_stale | 32 |
| l2_fresh | 28791 |
| l2_stale | 10208 |
| miss | 19585 |

- L1+L2 Fresh ratio：61.23%
- Refresh max：queue=0 active=9 pending=12

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 36732 | 0.662 |
| counter | 27006 | 0.888 |
| hydrate | 36732 | 0.634 |
| inbox | 36732 | 0.661 |
| merge_dedup | 36732 | 0.002 |
| relation | 27006 | 1.259 |
| route | 36732 | 1.594 |
| total | 76934 | 1.861 |

## Redis 本轮边界增量

- Commands：570330；input：77285260 bytes；output：103900940 bytes
- Hits/Misses：724812/80324；run hit rate：90.02%
- Evicted/Rejected：0/0；ops/s max：63567；safety epoch：613 -> 613

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.124 |
| client:loadtest | cpu_percent_total | 113.977 |
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
| docker:zg-canal | cpu_percent | 0.110 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.550 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.540 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 214.340 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 44.010 |
| docker:zg-zk | memory_percent | 1.580 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29627798.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 86.662 |
| process:counter | cpu_seconds_total | 234.938 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54538240.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 618.281 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 45719552.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 388.756 |
| process:knowpost | cpu_seconds_total | 2154.609 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 229584896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 161.691 |
| process:relation | cpu_seconds_total | 338.188 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56582144.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.391 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35627008.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 152.047 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 42782720.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 303340805.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 613.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 787764.000 |
| redis | keyspace_hits | 1146748314.000 |
| redis | keyspace_misses | 17990601.000 |
| redis | net_input_bytes | 54571869569.000 |
| redis | net_output_bytes | 210893743333.000 |
| redis | ops_per_sec | 63567.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54857.000 |
| redis | used_memory_bytes | 205857056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
