# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:28:04+08:00
- 采样时长：1m0.015062s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 69677 | 69677 | 0 | 0 | 1161.09 | 10.403 | 29.344 | 33.986 | 43.022 | 96.992 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 87121 | 1.250 |
| counter | 8955 | 0.129 |
| mysql | 4 | 0.000 |
| redis | 86604 | 1.243 |
| relation | 8955 | 0.129 |

- Cold compute：16369（0.235 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 31341 |
| l1_stale | 322 |
| l2_fresh | 30088 |
| l2_stale | 7926 |
| miss | 0 |

- L1+L2 Fresh ratio：88.16%
- Refresh max：queue=0 active=22 pending=22

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 16369 | 9.364 |
| counter | 8955 | 10.829 |
| hydrate | 16369 | 9.180 |
| inbox | 16369 | 9.363 |
| merge_dedup | 16369 | 0.004 |
| relation | 8955 | 13.054 |
| route | 16369 | 13.100 |
| total | 69677 | 9.107 |

## Redis 本轮边界增量

- Commands：269833；input：46175386 bytes；output：95836540 bytes
- Hits/Misses：428581/10325；run hit rate：97.65%
- Evicted/Rejected：0/0；ops/s max：13268；safety epoch：608 -> 608

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.969 |
| client:loadtest | cpu_percent_total | 95.497 |
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
| docker:zg-canal | cpu_percent | 2.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.660 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.740 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 166.550 |
| docker:zg-kafka | memory_percent | 8.780 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 52.100 |
| docker:zg-zk | memory_percent | 1.710 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29224466.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 50.759 |
| process:counter | cpu_seconds_total | 83.906 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54370304.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 465.627 |
| process:gateway | cpu_seconds_total | 413.281 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 51666944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 258.252 |
| process:knowpost | cpu_seconds_total | 1361.328 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 93159424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 84.430 |
| process:relation | cpu_seconds_total | 102.750 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56545280.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.780 |
| process:search | cpu_seconds_total | 0.203 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35573760.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 105.226 |
| process:user-storage | cpu_seconds_total | 101.406 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 47542272.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 294100403.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 608.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 709629.000 |
| redis | keyspace_hits | 1134814392.000 |
| redis | keyspace_misses | 16970183.000 |
| redis | net_input_bytes | 53329726629.000 |
| redis | net_output_bytes | 208912535943.000 |
| redis | ops_per_sec | 13268.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54523.000 |
| redis | used_memory_bytes | 127516968.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
