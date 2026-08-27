# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:30:07+08:00
- 采样时长：1m0.0228347s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 176924 | 176924 | 0 | 0 | 2948.23 | 10.546 | 12.969 | 14.122 | 18.595 | 94.640 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 220 | 0.001 |
| mysql | 176924 | 1.000 |
| redis | 176924 | 1.000 |
| relation | 220 | 0.001 |

- Cold compute：176924（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 176924 |
| l1_fresh | 0 |
| l1_stale | 0 |
| l2_fresh | 0 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 176924 | 0.164 |
| counter | 220 | 0.689 |
| hydrate | 176924 | 5.737 |
| inbox | 176924 | 0.162 |
| merge_dedup | 176924 | 0.015 |
| relation | 220 | 1.498 |
| route | 176924 | 0.008 |
| total | 176924 | 5.948 |

## Redis 本轮边界增量

- Commands：1121124；input：80709876 bytes；output：466606197 bytes
- Hits/Misses：1068314/220；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：20439；safety epoch：489 -> 489

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.949 |
| client:loadtest | cpu_percent_total | 63.179 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.690 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.590 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 186.020 |
| docker:zg-kafka | memory_percent | 8.620 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 60.780 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20603495.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 54.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 9.351 |
| process:counter | cpu_seconds_total | 98.734 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44474368.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37056512.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 481.881 |
| process:knowpost | cpu_seconds_total | 4256.953 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 90628096.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.089 |
| process:relation | cpu_seconds_total | 4.766 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 47992832.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.544 |
| process:search | cpu_seconds_total | 1.000 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37101568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.750 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34795520.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 181606601.000 |
| redis | connected_clients | 71.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 489.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677179.000 |
| redis | keyspace_hits | 980319942.000 |
| redis | keyspace_misses | 5372865.000 |
| redis | net_input_bytes | 40855933505.000 |
| redis | net_output_bytes | 181402696045.000 |
| redis | ops_per_sec | 20439.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 40246.000 |
| redis | used_memory_bytes | 98848240.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
