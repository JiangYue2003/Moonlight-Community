# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:48:35+08:00
- 采样时长：1m0.0901226s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4691188 | 4691188 | 0 | 0 | 78185.70 | 0.518 | 0.581 | 0.731 | 1.123 | 25.027 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146167 | 0.031 |
| counter | 10888 | 0.002 |
| mysql | 3 | 0.000 |
| redis | 145204 | 0.031 |
| relation | 10888 | 0.002 |

- Cold compute：22257（0.005 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4611792 |
| l1_stale | 0 |
| l2_fresh | 79396 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=3 pending=3

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22257 | 0.147 |
| counter | 10888 | 0.442 |
| hydrate | 22257 | 0.169 |
| inbox | 22257 | 0.146 |
| merge_dedup | 22257 | 0.005 |
| relation | 10888 | 1.006 |
| route | 22257 | 0.719 |
| total | 4691188 | 0.008 |

## Redis 本轮边界增量

- Commands：457512；input：68708709 bytes；output：167163517 bytes
- Hits/Misses：626203/12747；run hit rate：98.01%
- Evicted/Rejected：0/0；ops/s max：13076；safety epoch：544 -> 544

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 20.610 |
| client:loadtest | cpu_percent_total | 329.765 |
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
| docker:zg-canal | cpu_percent | 4.340 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.820 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.800 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 209.990 |
| docker:zg-kafka | memory_percent | 8.740 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 63.920 |
| docker:zg-zk | memory_percent | 1.660 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23330954.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 34.293 |
| process:counter | cpu_seconds_total | 78.750 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54063104.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 195.672 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 48287744.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 510.617 |
| process:knowpost | cpu_seconds_total | 1221.141 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 110178304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 30.147 |
| process:relation | cpu_seconds_total | 67.297 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56709120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.766 |
| process:search | cpu_seconds_total | 0.312 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36335616.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.559 |
| process:user-storage | cpu_seconds_total | 59.578 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 41742336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 207695346.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 544.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 701212.000 |
| redis | keyspace_hits | 1006440819.000 |
| redis | keyspace_misses | 6045869.000 |
| redis | net_input_bytes | 43115143287.000 |
| redis | net_output_bytes | 189430034077.000 |
| redis | ops_per_sec | 13076.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48554.000 |
| redis | used_memory_bytes | 121571504.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
