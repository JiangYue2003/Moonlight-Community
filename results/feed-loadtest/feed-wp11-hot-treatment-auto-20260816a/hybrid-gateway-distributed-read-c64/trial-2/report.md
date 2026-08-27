# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:09:04+08:00
- 采样时长：1m0.0286742s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 307458 | 307458 | 0 | 0 | 5123.71 | 12.245 | 16.092 | 20.030 | 26.733 | 67.610 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2426 | 0.008 |
| counter | 182 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2397 | 0.008 |
| relation | 182 | 0.001 |

- Cold compute：366（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 306130 |
| l1_stale | 0 |
| l2_fresh | 1328 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 366 | 0.208 |
| counter | 182 | 0.865 |
| hydrate | 366 | 0.311 |
| inbox | 366 | 0.205 |
| merge_dedup | 366 | 0.026 |
| relation | 182 | 1.971 |
| route | 366 | 1.423 |
| total | 307458 | 0.008 |

## Redis 本轮边界增量

- Commands：62938；input：5592612 bytes；output：6966025 bytes
- Hits/Misses：24014/182；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1301；safety epoch：514 -> 514

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.049 |
| client:loadtest | cpu_percent_total | 112.785 |
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
| docker:zg-canal | cpu_percent | 2.960 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.140 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.120 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 258.170 |
| docker:zg-kafka | memory_percent | 8.610 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 72.210 |
| docker:zg-zk | memory_percent | 1.510 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22732921.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 11.610 |
| process:counter | cpu_seconds_total | 228.500 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 45223936.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 464.396 |
| process:gateway | cpu_seconds_total | 3919.750 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 56360960.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 200.953 |
| process:knowpost | cpu_seconds_total | 9080.078 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 120946688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.553 |
| process:relation | cpu_seconds_total | 10.906 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48377856.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 2.031 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 230.650 |
| process:user-storage | cpu_seconds_total | 1968.641 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 65929216.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196556638.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 514.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677384.000 |
| redis | keyspace_hits | 993317361.000 |
| redis | keyspace_misses | 5382434.000 |
| redis | net_input_bytes | 41939936460.000 |
| redis | net_output_bytes | 187069819812.000 |
| redis | ops_per_sec | 1301.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42583.000 |
| redis | used_memory_bytes | 100595504.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
