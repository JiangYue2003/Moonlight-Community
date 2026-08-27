# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:51:37+08:00
- 采样时长：1m0.1631643s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5917029 | 5917029 | 0 | 0 | 98615.95 | 0.527 | 1.073 | 1.160 | 1.667 | 18.786 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146686 | 0.025 |
| counter | 10967 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 145422 | 0.025 |
| relation | 10967 | 0.002 |

- Cold compute：22260（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 5837123 |
| l1_stale | 0 |
| l2_fresh | 79906 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=2 pending=3

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22260 | 0.170 |
| counter | 10967 | 0.510 |
| hydrate | 22260 | 0.188 |
| inbox | 22260 | 0.168 |
| merge_dedup | 22260 | 0.005 |
| relation | 10967 | 1.150 |
| route | 22260 | 0.828 |
| total | 5917029 | 0.008 |

## Redis 本轮边界增量

- Commands：458849；input：68770380 bytes；output：167427639 bytes
- Hits/Misses：627406/12819；run hit rate：98.00%
- Evicted/Rejected：0/0；ops/s max：12521；safety epoch：546 -> 546

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 23.958 |
| client:loadtest | cpu_percent_total | 383.333 |
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
| docker:zg-canal | cpu_percent | 4.250 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.730 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.620 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 237.350 |
| docker:zg-kafka | memory_percent | 8.720 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 59.540 |
| docker:zg-zk | memory_percent | 1.600 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23357729.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 26.292 |
| process:counter | cpu_seconds_total | 99.641 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54329344.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 195.688 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 46964736.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 608.281 |
| process:knowpost | cpu_seconds_total | 1900.281 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 105799680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 33.252 |
| process:relation | cpu_seconds_total | 87.688 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56434688.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.406 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36261888.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 59.609 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 41615360.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 208791620.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 546.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 701513.000 |
| redis | keyspace_hits | 1007881625.000 |
| redis | keyspace_misses | 6080924.000 |
| redis | net_input_bytes | 43277273553.000 |
| redis | net_output_bytes | 189803768290.000 |
| redis | ops_per_sec | 12521.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48736.000 |
| redis | used_memory_bytes | 123471640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
