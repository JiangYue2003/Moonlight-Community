# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:14:44+08:00
- 采样时长：1m0.0485101s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 147780 | 147780 | 0 | 0 | 2461.42 | 51.140 | 60.085 | 63.766 | 72.737 | 105.325 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 120 | 0.001 |
| redis | 443460 | 3.001 |
| relation | 147780 | 1.000 |

- Cold compute：147780（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 147780 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 147780 | 0.282 |
| counter | 240 | 1.380 |
| hydrate | 147780 | 0.398 |
| inbox | 147780 | 0.315 |
| merge_dedup | 147780 | 0.013 |
| relation | 147780 | 49.183 |
| route | 147780 | 0.005 |
| total | 147780 | 50.222 |

## Redis 本轮边界增量

- Commands：1240860；input：282016690 bytes；output：1287112030 bytes
- Hits/Misses：6803850/147900；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22326；safety epoch：447 -> 447

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.801 |
| client:loadtest | cpu_percent_total | 60.810 |
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
| docker:zg-canal | cpu_percent | 3.300 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.100 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 8.450 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 289.040 |
| docker:zg-kafka | memory_percent | 8.510 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 83.800 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19990525.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.031 |
| process:counter | cpu_seconds_total | 705.828 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44523520.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 257.518 |
| process:gateway | cpu_seconds_total | 2860.016 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 60817408.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.539 |
| process:knowpost | cpu_seconds_total | 10324.781 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 108871680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 253.268 |
| process:relation | cpu_seconds_total | 9397.766 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 91451392.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.778 |
| process:search | cpu_seconds_total | 6.047 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37298176.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 126.521 |
| process:user-storage | cpu_seconds_total | 1502.312 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 56721408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 173455250.000 |
| redis | connected_clients | 153.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 447.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676577.000 |
| redis | keyspace_hits | 968732530.000 |
| redis | keyspace_misses | 5173914.000 |
| redis | net_input_bytes | 40009916706.000 |
| redis | net_output_bytes | 178493590829.000 |
| redis | ops_per_sec | 22326.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35723.000 |
| redis | used_memory_bytes | 100131416.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
