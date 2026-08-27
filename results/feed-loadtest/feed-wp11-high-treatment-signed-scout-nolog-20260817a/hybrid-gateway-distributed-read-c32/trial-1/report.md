# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-signed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:31:27+08:00
- 采样时长：10.0104754s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 17369 | 17369 | 0 | 0 | 1735.26 | 19.985 | 32.468 | 36.071 | 44.600 | 69.986 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 23366 | 1.345 |
| counter | 12779 | 0.736 |
| mysql | 26 | 0.001 |
| redis | 58098 | 3.345 |
| relation | 12779 | 0.736 |

- Cold compute：13968（0.804 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 1169 |
| l1_stale | 1 |
| l2_fresh | 2808 |
| l2_stale | 1785 |
| miss | 11606 |

- L1+L2 Fresh ratio：22.90%
- Refresh max：queue=0 active=6 pending=6

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 13968 | 3.250 |
| counter | 12779 | 3.663 |
| hydrate | 13968 | 3.074 |
| inbox | 13968 | 3.250 |
| merge_dedup | 13968 | 0.003 |
| relation | 12779 | 4.220 |
| route | 13968 | 7.234 |
| total | 17369 | 17.164 |

## Redis 本轮边界增量

- Commands：218436；input：29060229 bytes；output：31128623 bytes
- Hits/Misses：271844/37291；run hit rate：87.94%
- Evicted/Rejected：0/0；ops/s max：24158；safety epoch：563 -> 563

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.995 |
| client:loadtest | cpu_percent_total | 47.919 |
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
| docker:zg-canal | cpu_percent | 1.970 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.490 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.630 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 48.230 |
| docker:zg-kafka | memory_percent | 8.330 |
| docker:zg-kafka | pids | 97.000 |
| docker:zg-zk | cpu_percent | 0.120 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23632530.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 81.259 |
| process:counter | cpu_seconds_total | 452.391 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54149120.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 179.544 |
| process:gateway | cpu_seconds_total | 1819.734 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52236288.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 311.880 |
| process:knowpost | cpu_seconds_total | 3849.188 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 174489600.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 141.003 |
| process:relation | cpu_seconds_total | 487.578 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 61734912.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36175872.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 85.839 |
| process:user-storage | cpu_seconds_total | 448.203 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 53747712.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 217160542.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 563.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 754527.000 |
| redis | keyspace_hits | 1016702407.000 |
| redis | keyspace_misses | 6703024.000 |
| redis | net_input_bytes | 44358437732.000 |
| redis | net_output_bytes | 191477008953.000 |
| redis | ops_per_sec | 24158.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51076.000 |
| redis | used_memory_bytes | 171949208.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
