# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:11:36+08:00
- 采样时长：1m0.0464424s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 295563 | 295563 | 0 | 0 | 4923.98 | 24.880 | 35.160 | 45.209 | 66.224 | 160.621 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2425 | 0.008 |
| counter | 178 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2394 | 0.008 |
| relation | 178 | 0.001 |

- Cold compute：362（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 294224 |
| l1_stale | 0 |
| l2_fresh | 1339 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 362 | 0.254 |
| counter | 178 | 0.873 |
| hydrate | 362 | 0.325 |
| inbox | 362 | 0.254 |
| merge_dedup | 362 | 0.015 |
| relation | 178 | 2.185 |
| route | 362 | 1.523 |
| total | 295563 | 0.008 |

## Redis 本轮边界增量

- Commands：62871；input：5576494 bytes；output：6945973 bytes
- Hits/Misses：23739/178；run hit rate：99.26%
- Evicted/Rejected：0/0；ops/s max：1301；safety epoch：516 -> 516

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.066 |
| client:loadtest | cpu_percent_total | 113.064 |
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
| docker:zg-canal | cpu_percent | 2.720 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.620 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.380 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 200.820 |
| docker:zg-kafka | memory_percent | 8.640 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 66.320 |
| docker:zg-zk | memory_percent | 1.530 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22733633.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 13.199 |
| process:counter | cpu_seconds_total | 239.219 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 45072384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 442.884 |
| process:gateway | cpu_seconds_total | 4467.781 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 60563456.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 196.938 |
| process:knowpost | cpu_seconds_total | 9338.125 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 115302400.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.089 |
| process:relation | cpu_seconds_total | 11.625 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48459776.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 2.188 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37011456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 218.166 |
| process:user-storage | cpu_seconds_total | 2239.812 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 66998272.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196715169.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 516.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677466.000 |
| redis | keyspace_hits | 993376595.000 |
| redis | keyspace_misses | 5383211.000 |
| redis | net_input_bytes | 41954015814.000 |
| redis | net_output_bytes | 187086829723.000 |
| redis | ops_per_sec | 1301.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42735.000 |
| redis | used_memory_bytes | 100644576.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
