# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:17:37+08:00
- 采样时长：1m0.037147s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 480980 | 480980 | 0 | 0 | 8015.17 | 7.903 | 9.953 | 10.614 | 12.192 | 26.764 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2438 | 0.005 |
| counter | 180 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2393 | 0.005 |
| relation | 180 | 0.000 |

- Cold compute：361（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 479625 |
| l1_stale | 0 |
| l2_fresh | 1355 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 361 | 0.148 |
| counter | 180 | 0.663 |
| hydrate | 361 | 0.195 |
| inbox | 361 | 0.148 |
| merge_dedup | 361 | 0.014 |
| relation | 180 | 1.427 |
| route | 361 | 1.063 |
| total | 480980 | 0.005 |

## Redis 本轮边界增量

- Commands：63178；input：5593540 bytes；output：6948800 bytes
- Hits/Misses：23746/180；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1314；safety epoch：479 -> 479

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.804 |
| client:loadtest | cpu_percent_total | 124.871 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.570 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.470 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 162.100 |
| docker:zg-kafka | memory_percent | 8.580 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 53.320 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20182154.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.070 |
| process:counter | cpu_seconds_total | 61.375 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43212800.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.453 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37933056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 251.381 |
| process:knowpost | cpu_seconds_total | 2449.750 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 123715584.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.319 |
| process:relation | cpu_seconds_total | 1.609 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46317568.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.766 |
| process:search | cpu_seconds_total | 0.656 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.594 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34680832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178321042.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 479.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677073.000 |
| redis | keyspace_hits | 977559534.000 |
| redis | keyspace_misses | 5367415.000 |
| redis | net_input_bytes | 40608223814.000 |
| redis | net_output_bytes | 180233218365.000 |
| redis | ops_per_sec | 1314.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39496.000 |
| redis | used_memory_bytes | 98004984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
