# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:16:22+08:00
- 采样时长：1m0.0284146s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 481783 | 481783 | 0 | 0 | 8029.37 | 3.815 | 5.179 | 5.423 | 6.334 | 16.526 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2463 | 0.005 |
| counter | 183 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2422 | 0.005 |
| relation | 183 | 0.000 |

- Cold compute：367（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 480421 |
| l1_stale | 0 |
| l2_fresh | 1362 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 367 | 0.148 |
| counter | 183 | 0.630 |
| hydrate | 367 | 0.198 |
| inbox | 367 | 0.144 |
| merge_dedup | 367 | 0.010 |
| relation | 183 | 1.460 |
| route | 367 | 1.058 |
| total | 481783 | 0.005 |

## Redis 本轮边界增量

- Commands：62935；input：5595671 bytes；output：7016831 bytes
- Hits/Misses：24108/183；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1426；safety epoch：478 -> 478

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.957 |
| client:loadtest | cpu_percent_total | 127.310 |
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
| docker:zg-canal | cpu_percent | 2.080 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.590 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.930 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.290 |
| docker:zg-kafka | memory_percent | 8.210 |
| docker:zg-kafka | pids | 112.000 |
| docker:zg-zk | cpu_percent | 3.350 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20181790.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.502 |
| process:counter | cpu_seconds_total | 57.750 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43184128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.438 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37933056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 232.378 |
| process:knowpost | cpu_seconds_total | 2294.406 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 120770560.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 1.328 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46137344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.625 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.547 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34570240.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178242296.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 478.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677055.000 |
| redis | keyspace_hits | 977530254.000 |
| redis | keyspace_misses | 5366815.000 |
| redis | net_input_bytes | 40601179518.000 |
| redis | net_output_bytes | 180224777258.000 |
| redis | ops_per_sec | 1426.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39421.000 |
| redis | used_memory_bytes | 98017104.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
