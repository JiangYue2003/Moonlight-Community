# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:49:17+08:00
- 采样时长：1m0.0131889s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 61278 | 61278 | 0 | 0 | 1021.15 | 10.160 | 33.040 | 36.850 | 44.389 | 73.112 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 117002 | 1.909 |
| counter | 48674 | 0.794 |
| mysql | 25 | 0.000 |
| redis | 216956 | 3.541 |
| relation | 48674 | 0.794 |

- Cold compute：52723（0.860 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 2456 |
| l1_stale | 23 |
| l2_fresh | 8061 |
| l2_stale | 17386 |
| miss | 33352 |

- L1+L2 Fresh ratio：17.16%
- Refresh max：queue=0 active=12 pending=12

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52723 | 2.850 |
| counter | 48674 | 3.470 |
| hydrate | 52723 | 2.727 |
| inbox | 52723 | 2.849 |
| merge_dedup | 52723 | 0.002 |
| relation | 48674 | 4.544 |
| route | 52723 | 7.421 |
| total | 61278 | 13.801 |

## Redis 本轮边界增量

- Commands：835371；input：110461735 bytes；output：124889866 bytes
- Hits/Misses：1036438/130347；run hit rate：88.83%
- Evicted/Rejected：0/0；ops/s max：18927；safety epoch：578 -> 578

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.294 |
| client:loadtest | cpu_percent_total | 36.711 |
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
| docker:zg-canal | cpu_percent | 2.580 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.660 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.760 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 161.210 |
| docker:zg-kafka | memory_percent | 8.800 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 56.250 |
| docker:zg-zk | memory_percent | 1.650 |
| docker:zg-zk | pids | 87.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25350527.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 96.631 |
| process:counter | cpu_seconds_total | 1097.953 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54571008.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 176.663 |
| process:gateway | cpu_seconds_total | 2276.844 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 51310592.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 286.694 |
| process:knowpost | cpu_seconds_total | 6858.000 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 148127744.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 153.243 |
| process:relation | cpu_seconds_total | 1521.750 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62083072.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.540 |
| process:search | cpu_seconds_total | 1.734 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36225024.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 68.160 |
| process:user-storage | cpu_seconds_total | 625.344 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 54272000.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 255475445.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 578.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 796915.000 |
| redis | keyspace_hits | 1066056890.000 |
| redis | keyspace_misses | 11072979.000 |
| redis | net_input_bytes | 49530331383.000 |
| redis | net_output_bytes | 199394260736.000 |
| redis | ops_per_sec | 18927.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52196.000 |
| redis | used_memory_bytes | 214881224.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
