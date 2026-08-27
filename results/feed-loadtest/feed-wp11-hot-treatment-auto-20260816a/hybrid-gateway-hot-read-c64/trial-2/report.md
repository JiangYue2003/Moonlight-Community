# Feed 压测报告：hybrid / gateway / hot-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:48:58+08:00
- 采样时长：1m0.026487s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 283053 | 283053 | 0 | 0 | 4717.05 | 12.164 | 17.430 | 21.917 | 87.410 | 142.474 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 169 | 0.001 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 120 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：18（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 282938 |
| l1_stale | 0 |
| l2_fresh | 115 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.263 |
| counter | 9 | 1.951 |
| hydrate | 18 | 0.284 |
| inbox | 18 | 0.263 |
| merge_dedup | 18 | 0.000 |
| relation | 9 | 3.513 |
| route | 18 | 2.732 |
| total | 283053 | 0.007 |

## Redis 本轮边界增量

- Commands：53242；input：3806547 bytes；output：1428378 bytes
- Hits/Misses：1327/9；run hit rate：99.33%
- Evicted/Rejected：0/0；ops/s max：1095；safety epoch：504 -> 504

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.535 |
| client:loadtest | cpu_percent_total | 104.563 |
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
| docker:zg-canal | cpu_percent | 2.610 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.470 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 157.000 |
| docker:zg-etcd | cpu_percent | 7.220 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 221.800 |
| docker:zg-kafka | memory_percent | 8.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 66.190 |
| docker:zg-zk | memory_percent | 1.390 |
| docker:zg-zk | pids | 93.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22730424.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 12.454 |
| process:counter | cpu_seconds_total | 163.109 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43032576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 437.318 |
| process:gateway | cpu_seconds_total | 1307.531 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 55054336.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 191.063 |
| process:knowpost | cpu_seconds_total | 7833.625 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 113135616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 8.938 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 45076480.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 1.453 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36974592.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 203.434 |
| process:user-storage | cpu_seconds_total | 650.938 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 68055040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195421468.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 504.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677239.000 |
| redis | keyspace_hits | 993146412.000 |
| redis | keyspace_misses | 5378723.000 |
| redis | net_input_bytes | 41851312985.000 |
| redis | net_output_bytes | 187007529909.000 |
| redis | ops_per_sec | 1095.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41377.000 |
| redis | used_memory_bytes | 100122056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
