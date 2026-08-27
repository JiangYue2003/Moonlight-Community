# Feed 压测报告：hybrid / gateway / hot-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:47:44+08:00
- 采样时长：1m0.0389101s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 288765 | 288765 | 0 | 0 | 4811.63 | 12.235 | 17.053 | 21.265 | 38.707 | 198.531 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 162 | 0.001 |
| counter | 8 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 118 | 0.000 |
| relation | 8 | 0.000 |

- Cold compute：17（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 288654 |
| l1_stale | 0 |
| l2_fresh | 111 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17 | 0.215 |
| counter | 8 | 2.310 |
| hydrate | 17 | 0.397 |
| inbox | 17 | 0.215 |
| merge_dedup | 17 | 0.000 |
| relation | 8 | 2.320 |
| route | 17 | 2.179 |
| total | 288765 | 0.007 |

## Redis 本轮边界增量

- Commands：53365；input：3812376 bytes；output：1423773 bytes
- Hits/Misses：1258/8；run hit rate：99.37%
- Evicted/Rejected：0/0；ops/s max：1118；safety epoch：503 -> 503

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.622 |
| client:loadtest | cpu_percent_total | 105.947 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.750 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.390 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 241.180 |
| docker:zg-kafka | memory_percent | 8.630 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 60.210 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22730338.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 6.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.835 |
| process:counter | cpu_seconds_total | 157.781 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43020288.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 454.602 |
| process:gateway | cpu_seconds_total | 1045.469 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 55066624.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 204.566 |
| process:knowpost | cpu_seconds_total | 7715.281 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 113659904.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.333 |
| process:relation | cpu_seconds_total | 8.844 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 45080576.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.055 |
| process:search | cpu_seconds_total | 1.391 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36970496.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 225.175 |
| process:user-storage | cpu_seconds_total | 523.875 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 67485696.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195355566.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 503.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677239.000 |
| redis | keyspace_hits | 993144780.000 |
| redis | keyspace_misses | 5378658.000 |
| redis | net_input_bytes | 41846590173.000 |
| redis | net_output_bytes | 187005786356.000 |
| redis | ops_per_sec | 1118.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41303.000 |
| redis | used_memory_bytes | 100140776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
