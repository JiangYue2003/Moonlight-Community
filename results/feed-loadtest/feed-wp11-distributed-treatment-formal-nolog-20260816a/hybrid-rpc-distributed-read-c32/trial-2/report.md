# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:47:23+08:00
- 采样时长：1m0.1037642s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4711374 | 4711374 | 0 | 0 | 78522.41 | 0.518 | 0.582 | 0.736 | 1.128 | 18.137 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146070 | 0.031 |
| counter | 10855 | 0.002 |
| mysql | 2 | 0.000 |
| redis | 145139 | 0.031 |
| relation | 10855 | 0.002 |

- Cold compute：22252（0.005 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4632060 |
| l1_stale | 0 |
| l2_fresh | 79314 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=3 pending=3

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22252 | 0.147 |
| counter | 10855 | 0.445 |
| hydrate | 22252 | 0.168 |
| inbox | 22252 | 0.146 |
| merge_dedup | 22252 | 0.004 |
| relation | 10855 | 1.005 |
| route | 22252 | 0.718 |
| total | 4711374 | 0.008 |

## Redis 本轮边界增量

- Commands：457116；input：68688621 bytes；output：167065346 bytes
- Hits/Misses：625941/12709；run hit rate：98.01%
- Evicted/Rejected：0/0；ops/s max：13622；safety epoch：543 -> 543

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 20.510 |
| client:loadtest | cpu_percent_total | 328.156 |
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
| docker:zg-canal | cpu_percent | 2.740 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.820 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.960 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 227.060 |
| docker:zg-kafka | memory_percent | 8.830 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 60.330 |
| docker:zg-zk | memory_percent | 1.660 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23317598.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 22.601 |
| process:counter | cpu_seconds_total | 68.844 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54243328.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 195.656 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54370304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 528.667 |
| process:knowpost | cpu_seconds_total | 932.969 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 105189376.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 30.158 |
| process:relation | cpu_seconds_total | 58.906 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56287232.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.772 |
| process:search | cpu_seconds_total | 0.281 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36343808.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.543 |
| process:user-storage | cpu_seconds_total | 59.469 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 41664512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 207166379.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 543.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 700890.000 |
| redis | keyspace_hits | 1005720573.000 |
| redis | keyspace_misses | 6028996.000 |
| redis | net_input_bytes | 43035441690.000 |
| redis | net_output_bytes | 189243107068.000 |
| redis | ops_per_sec | 13622.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48482.000 |
| redis | used_memory_bytes | 121289744.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
