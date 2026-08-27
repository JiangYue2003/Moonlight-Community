# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:37:24+08:00
- 采样时长：1m0.0203771s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167476 | 167476 | 0 | 0 | 2790.93 | 10.666 | 12.965 | 14.221 | 19.828 | 106.972 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 155 | 0.001 |
| redis | 502583 | 3.001 |
| relation | 167476 | 1.000 |

- Cold compute：167476（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 167476 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167476 | 0.201 |
| counter | 240 | 0.985 |
| hydrate | 167476 | 0.356 |
| inbox | 167476 | 0.221 |
| merge_dedup | 167476 | 0.015 |
| relation | 167476 | 10.123 |
| route | 167476 | 0.004 |
| total | 167476 | 10.949 |

## Redis 本轮边界增量

- Commands：1405064；input：425233212 bytes；output：1974769138 bytes
- Hits/Misses：10891874/156；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25111；safety epoch：418 -> 418

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.053 |
| client:loadtest | cpu_percent_total | 64.848 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.620 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 177.430 |
| docker:zg-kafka | memory_percent | 8.370 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 63.080 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14796262.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 41.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.612 |
| process:counter | cpu_seconds_total | 136.391 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43474944.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36966400.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 334.679 |
| process:knowpost | cpu_seconds_total | 5140.969 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 89485312.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 275.274 |
| process:relation | cpu_seconds_total | 4811.625 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 100655104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.875 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.969 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35188736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 120314490.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 418.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698726.000 |
| redis | keyspace_hits | 697307540.000 |
| redis | keyspace_misses | 97893.000 |
| redis | net_input_bytes | 28289166258.000 |
| redis | net_output_bytes | 127558750006.000 |
| redis | ops_per_sec | 25111.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22683.000 |
| redis | used_memory_bytes | 107208640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
