# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:10:50+08:00
- 采样时长：10.0104686s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 63355 | 63355 | 0 | 0 | 6331.53 | 3.746 | 11.369 | 14.439 | 20.393 | 43.048 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 94649 | 1.494 |
| counter | 26212 | 0.414 |
| mysql | 36 | 0.001 |
| redis | 152048 | 2.400 |
| relation | 26212 | 0.414 |

- Cold compute：33860（0.534 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 12797 |
| l1_stale | 33 |
| l2_fresh | 21555 |
| l2_stale | 9762 |
| miss | 19208 |

- L1+L2 Fresh ratio：54.22%
- Refresh max：queue=1 active=32 pending=33

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 33860 | 1.615 |
| counter | 26212 | 1.871 |
| hydrate | 33860 | 1.536 |
| inbox | 33860 | 1.614 |
| merge_dedup | 33860 | 0.002 |
| relation | 26212 | 2.106 |
| route | 33860 | 3.096 |
| total | 63355 | 4.799 |

## Redis 本轮边界增量

- Commands：521715；input：70770244 bytes；output：91937045 bytes
- Hits/Misses：667159/76533；run hit rate：89.71%
- Evicted/Rejected：0/0；ops/s max：60993；safety epoch：558 -> 558

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.439 |
| client:loadtest | cpu_percent_total | 103.017 |
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
| docker:zg-canal | cpu_percent | 2.460 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.760 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.230 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 58.130 |
| docker:zg-kafka | memory_percent | 8.270 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.260 |
| docker:zg-zk | memory_percent | 1.460 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23519656.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 37.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 96.782 |
| process:counter | cpu_seconds_total | 348.578 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53669888.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1763.609 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52359168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 433.196 |
| process:knowpost | cpu_seconds_total | 3578.984 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 224927744.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 177.041 |
| process:relation | cpu_seconds_total | 418.469 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56086528.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36397056.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 422.203 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 54501376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 213470006.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 558.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 746788.000 |
| redis | keyspace_hits | 1013882259.000 |
| redis | keyspace_misses | 6377931.000 |
| redis | net_input_bytes | 43965287285.000 |
| redis | net_output_bytes | 191036017077.000 |
| redis | ops_per_sec | 60993.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49839.000 |
| redis | used_memory_bytes | 165826400.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
