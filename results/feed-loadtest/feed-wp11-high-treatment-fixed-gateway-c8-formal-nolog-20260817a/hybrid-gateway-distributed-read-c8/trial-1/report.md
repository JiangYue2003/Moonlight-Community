# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-treatment-fixed-gateway-c8-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:39:28+08:00
- 采样时长：1m0.0137571s
- 并发：8
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 131851 | 131851 | 0 | 0 | 2197.33 | 2.715 | 7.083 | 10.913 | 15.693 | 29.031 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 298313 | 2.263 |
| counter | 83145 | 0.631 |
| mysql | 25 | 0.000 |
| redis | 412118 | 3.126 |
| relation | 83145 | 0.631 |

- Cold compute：97154（0.737 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 11165 |
| l1_stale | 27 |
| l2_fresh | 31274 |
| l2_stale | 51432 |
| miss | 37953 |

- L1+L2 Fresh ratio：32.19%
- Refresh max：queue=0 active=10 pending=10

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 97154 | 0.837 |
| counter | 83145 | 1.127 |
| hydrate | 97154 | 0.781 |
| inbox | 97154 | 0.836 |
| merge_dedup | 97154 | 0.002 |
| relation | 83145 | 1.720 |
| route | 97154 | 2.456 |
| total | 131851 | 2.667 |

## Redis 本轮边界增量

- Commands：1559873；input：206210133 bytes；output：260040893 bytes
- Hits/Misses：1925584/210137；run hit rate：90.16%
- Evicted/Rejected：0/0；ops/s max：37511；safety epoch：617 -> 617

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.023 |
| client:loadtest | cpu_percent_total | 48.374 |
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
| docker:zg-canal | cpu_percent | 1.680 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.510 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 8.020 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 204.750 |
| docker:zg-kafka | memory_percent | 8.800 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 58.740 |
| docker:zg-zk | memory_percent | 1.750 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 30214016.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 97.014 |
| process:counter | cpu_seconds_total | 432.906 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54837248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 205.022 |
| process:gateway | cpu_seconds_total | 733.297 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 50429952.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 293.961 |
| process:knowpost | cpu_seconds_total | 3095.453 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 187625472.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 129.352 |
| process:relation | cpu_seconds_total | 652.531 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56897536.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35766272.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 106.814 |
| process:user-storage | cpu_seconds_total | 202.484 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 50348032.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 316676656.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 617.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 798927.000 |
| redis | keyspace_hits | 1163918220.000 |
| redis | keyspace_misses | 19488550.000 |
| redis | net_input_bytes | 56381558778.000 |
| redis | net_output_bytes | 213664491697.000 |
| redis | ops_per_sec | 37511.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 55207.000 |
| redis | used_memory_bytes | 214701528.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
