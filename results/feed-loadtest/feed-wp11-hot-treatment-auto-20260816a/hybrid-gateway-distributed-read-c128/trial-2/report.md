# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:12:51+08:00
- 采样时长：1m0.0338255s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 293521 | 293521 | 0 | 0 | 4890.87 | 24.666 | 34.480 | 46.780 | 84.332 | 176.191 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2435 | 0.008 |
| counter | 181 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2402 | 0.008 |
| relation | 181 | 0.001 |

- Cold compute：368（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 292190 |
| l1_stale | 0 |
| l2_fresh | 1331 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 368 | 0.238 |
| counter | 181 | 0.872 |
| hydrate | 368 | 0.338 |
| inbox | 368 | 0.232 |
| merge_dedup | 368 | 0.015 |
| relation | 181 | 1.978 |
| route | 368 | 1.424 |
| total | 293521 | 0.008 |

## Redis 本轮边界增量

- Commands：62987；input：5605712 bytes；output：6983135 bytes
- Hits/Misses：24080/181；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1343；safety epoch：517 -> 517

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.977 |
| client:loadtest | cpu_percent_total | 111.630 |
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
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.780 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.730 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 203.120 |
| docker:zg-kafka | memory_percent | 8.650 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 75.050 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22733967.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 13.902 |
| process:counter | cpu_seconds_total | 244.047 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 45637632.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 422.834 |
| process:gateway | cpu_seconds_total | 4736.922 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 61440000.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 204.224 |
| process:knowpost | cpu_seconds_total | 9463.422 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 132427776.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.865 |
| process:relation | cpu_seconds_total | 12.016 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48164864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 2.219 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 225.089 |
| process:user-storage | cpu_seconds_total | 2373.266 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 66695168.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196793912.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 517.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677480.000 |
| redis | keyspace_hits | 993406268.000 |
| redis | keyspace_misses | 5383559.000 |
| redis | net_input_bytes | 41961014844.000 |
| redis | net_output_bytes | 187095306629.000 |
| redis | ops_per_sec | 1343.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42810.000 |
| redis | used_memory_bytes | 100618304.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
