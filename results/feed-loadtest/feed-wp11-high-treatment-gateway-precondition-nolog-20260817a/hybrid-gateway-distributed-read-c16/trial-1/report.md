# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-gateway-precondition-nolog-20260817a`
- 开始时间：2026-08-17T00:44:58+08:00
- 采样时长：1m0.0210329s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 130157 | 130157 | 0 | 0 | 2168.88 | 5.819 | 17.671 | 21.797 | 28.301 | 48.434 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 291114 | 2.237 |
| counter | 81938 | 0.630 |
| mysql | 31 | 0.000 |
| redis | 406344 | 3.122 |
| relation | 81938 | 0.630 |

- Cold compute：95807（0.736 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 11128 |
| l1_stale | 54 |
| l2_fresh | 30956 |
| l2_stale | 49540 |
| miss | 38479 |

- L1+L2 Fresh ratio：32.33%
- Refresh max：queue=0 active=19 pending=19

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 95807 | 2.016 |
| counter | 81938 | 2.409 |
| hydrate | 95807 | 1.892 |
| inbox | 95807 | 2.015 |
| merge_dedup | 95807 | 0.002 |
| relation | 81938 | 2.968 |
| route | 95807 | 4.619 |
| total | 130157 | 6.265 |

## Redis 本轮边界增量

- Commands：1494750；input：200312946 bytes；output：254696274 bytes
- Hits/Misses：1897586/208247；run hit rate：90.11%
- Evicted/Rejected：0/0；ops/s max：40027；safety epoch：575 -> 575

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.306 |
| client:loadtest | cpu_percent_total | 52.898 |
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
| docker:zg-canal | cpu_percent | 4.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.980 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.500 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 202.800 |
| docker:zg-kafka | memory_percent | 8.780 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 77.680 |
| docker:zg-zk | memory_percent | 1.730 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25148503.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 95.187 |
| process:counter | cpu_seconds_total | 940.453 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54562816.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 224.225 |
| process:gateway | cpu_seconds_total | 1957.547 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 51871744.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 323.704 |
| process:knowpost | cpu_seconds_total | 6321.906 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 189661184.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 150.315 |
| process:relation | cpu_seconds_total | 1272.391 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62521344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.672 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36212736.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 98.009 |
| process:user-storage | cpu_seconds_total | 503.266 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 60456960.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 251843391.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 575.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 808494.000 |
| redis | keyspace_hits | 1061632568.000 |
| redis | keyspace_misses | 10546285.000 |
| redis | net_input_bytes | 49053966599.000 |
| redis | net_output_bytes | 198834381216.000 |
| redis | ops_per_sec | 40027.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51937.000 |
| redis | used_memory_bytes | 227144200.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
