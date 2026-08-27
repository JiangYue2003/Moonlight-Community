# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:39:17+08:00
- 采样时长：1m0.0306284s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 549715 | 549715 | 0 | 0 | 9161.55 | 1.859 | 3.140 | 3.762 | 5.666 | 37.752 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1008901 | 1.835 |
| counter | 146633 | 0.267 |
| mysql | 20 | 0.000 |
| redis | 1071972 | 1.950 |
| relation | 146633 | 0.267 |

- Cold compute：224383（0.408 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 150151 |
| l1_stale | 437 |
| l2_fresh | 240461 |
| l2_stale | 137512 |
| miss | 21154 |

- L1+L2 Fresh ratio：71.06%
- Refresh max：queue=0 active=27 pending=26

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 224383 | 1.010 |
| counter | 146633 | 1.419 |
| hydrate | 224383 | 0.995 |
| inbox | 224383 | 1.009 |
| merge_dedup | 224383 | 0.002 |
| relation | 146633 | 1.607 |
| route | 224383 | 1.992 |
| total | 549715 | 1.524 |

## Redis 本轮边界增量

- Commands：3439012；input：472071693 bytes；output：749574675 bytes
- Hits/Misses：4516588/373191；run hit rate：92.37%
- Evicted/Rejected：0/0；ops/s max：63401；safety epoch：570 -> 570

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.335 |
| client:loadtest | cpu_percent_total | 117.362 |
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
| docker:zg-canal | cpu_percent | 2.770 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.280 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 8.530 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 207.200 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 65.270 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24500245.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 104.627 |
| process:counter | cpu_seconds_total | 728.609 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54677504.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1840.953 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45789184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 436.772 |
| process:knowpost | cpu_seconds_total | 5230.484 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 234270720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 179.927 |
| process:relation | cpu_seconds_total | 927.516 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62767104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.516 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36184064.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.544 |
| process:user-storage | cpu_seconds_total | 456.969 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43601920.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 237052324.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 570.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 809527.000 |
| redis | keyspace_hits | 1042414255.000 |
| redis | keyspace_misses | 8909788.000 |
| redis | net_input_bytes | 47061104705.000 |
| redis | net_output_bytes | 195663020008.000 |
| redis | ops_per_sec | 63401.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51597.000 |
| redis | used_memory_bytes | 226544352.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
