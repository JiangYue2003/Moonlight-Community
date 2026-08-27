# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:43:50+08:00
- 采样时长：10.077418s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 13332 | 13332 | 0 | 0 | 1323.11 | 40.291 | 100.519 | 107.313 | 144.700 | 229.352 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 11857 | 0.889 |
| counter | 1252 | 0.094 |
| mysql | 0 | 0.000 |
| redis | 10718 | 0.804 |
| relation | 1260 | 0.095 |

- Cold compute：1400（0.105 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4030 |
| l1_stale | 2270 |
| l2_fresh | 3998 |
| l2_stale | 3034 |
| miss | 0 |

- L1+L2 Fresh ratio：60.22%
- Refresh max：queue=601 active=32 pending=633

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1406 | 39.406 |
| counter | 1252 | 42.743 |
| hydrate | 1400 | 39.831 |
| inbox | 1406 | 39.403 |
| merge_dedup | 1406 | 0.002 |
| relation | 1260 | 45.762 |
| route | 1411 | 94.796 |
| total | 13332 | 39.073 |

## Redis 本轮边界增量

- Commands：33636；input：4475404 bytes；output：12752560 bytes
- Hits/Misses：47548/1396；run hit rate：97.15%
- Evicted/Rejected：0/0；ops/s max：6852；safety epoch：540 -> 540

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.547 |
| client:loadtest | cpu_percent_total | 136.754 |
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
| docker:zg-canal | cpu_percent | 0.610 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.750 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.780 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 149.460 |
| docker:zg-kafka | memory_percent | 8.960 |
| docker:zg-kafka | pids | 143.000 |
| docker:zg-zk | cpu_percent | 37.210 |
| docker:zg-zk | memory_percent | 1.420 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23287143.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 49.213 |
| process:counter | cpu_seconds_total | 38.656 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 50454528.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 436.577 |
| process:gateway | cpu_seconds_total | 183.906 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 56426496.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 280.804 |
| process:knowpost | cpu_seconds_total | 332.406 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 104595456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 78.315 |
| process:relation | cpu_seconds_total | 36.922 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 55451648.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.156 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36298752.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 95.642 |
| process:user-storage | cpu_seconds_total | 53.766 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 48480256.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205939347.000 |
| redis | connected_clients | 245.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 540.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 698302.000 |
| redis | keyspace_hits | 1004206416.000 |
| redis | keyspace_misses | 5989384.000 |
| redis | net_input_bytes | 42859192419.000 |
| redis | net_output_bytes | 188857553120.000 |
| redis | ops_per_sec | 6852.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48220.000 |
| redis | used_memory_bytes | 122251080.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
