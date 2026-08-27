# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:24:02+08:00
- 采样时长：1m0.1588743s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6138808 | 6138808 | 0 | 0 | 102312.74 | 0.528 | 1.071 | 1.151 | 1.668 | 17.346 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 147009 | 0.024 |
| counter | 10901 | 0.002 |
| mysql | 3 | 0.000 |
| redis | 145690 | 0.024 |
| relation | 10901 | 0.002 |

- Cold compute：22291（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 6058672 |
| l1_stale | 0 |
| l2_fresh | 80136 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=2 pending=2

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22291 | 0.175 |
| counter | 10901 | 0.516 |
| hydrate | 22291 | 0.201 |
| inbox | 22291 | 0.174 |
| merge_dedup | 22291 | 0.005 |
| relation | 10901 | 1.164 |
| route | 22291 | 0.833 |
| total | 6138808 | 0.008 |

## Redis 本轮边界增量

- Commands：461644；input：69044449 bytes；output：167804339 bytes
- Hits/Misses：627644/12766；run hit rate：98.01%
- Evicted/Rejected：0/0；ops/s max：12617；safety epoch：605 -> 605

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.525 |
| client:loadtest | cpu_percent_total | 392.398 |
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
| docker:zg-canal | cpu_percent | 2.900 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.690 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.920 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 273.210 |
| docker:zg-kafka | memory_percent | 8.840 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 68.140 |
| docker:zg-zk | memory_percent | 1.690 |
| docker:zg-zk | pids | 101.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29188134.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 27.827 |
| process:counter | cpu_seconds_total | 22.359 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 53411840.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.047 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 37294080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 643.928 |
| process:knowpost | cpu_seconds_total | 707.094 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 100048896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 30.909 |
| process:relation | cpu_seconds_total | 18.906 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 55726080.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.078 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35561472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 0.234 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 34422784.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 292850996.000 |
| redis | connected_clients | 214.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 605.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 708753.000 |
| redis | keyspace_hits | 1133042680.000 |
| redis | keyspace_misses | 16922620.000 |
| redis | net_input_bytes | 53132413358.000 |
| redis | net_output_bytes | 208493322222.000 |
| redis | ops_per_sec | 12617.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54281.000 |
| redis | used_memory_bytes | 124102112.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
