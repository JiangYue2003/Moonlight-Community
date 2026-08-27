# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:15:07+08:00
- 采样时长：1m0.0284356s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 489599 | 489599 | 0 | 0 | 8159.55 | 3.748 | 4.883 | 5.306 | 6.152 | 13.488 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2447 | 0.005 |
| counter | 181 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2409 | 0.005 |
| relation | 181 | 0.000 |

- Cold compute：364（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 488244 |
| l1_stale | 0 |
| l2_fresh | 1355 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 364 | 0.138 |
| counter | 181 | 0.679 |
| hydrate | 364 | 0.225 |
| inbox | 364 | 0.138 |
| merge_dedup | 364 | 0.014 |
| relation | 181 | 1.403 |
| route | 364 | 1.045 |
| total | 489599 | 0.005 |

## Redis 本轮边界增量

- Commands：62846；input：5579359 bytes；output：6980471 bytes
- Hits/Misses：23915/181；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1295；safety epoch：477 -> 477

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.794 |
| client:loadtest | cpu_percent_total | 124.707 |
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
| docker:zg-canal | cpu_percent | 1.750 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.450 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.850 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 136.040 |
| docker:zg-kafka | memory_percent | 8.420 |
| docker:zg-kafka | pids | 120.000 |
| docker:zg-zk | cpu_percent | 42.100 |
| docker:zg-zk | memory_percent | 1.460 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20181431.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.734 |
| process:counter | cpu_seconds_total | 54.312 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 42872832.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.438 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37924864.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 239.129 |
| process:knowpost | cpu_seconds_total | 2143.078 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 142958592.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.872 |
| process:relation | cpu_seconds_total | 1.156 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46419968.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 0.594 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.516 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34566144.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178162520.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 477.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677038.000 |
| redis | keyspace_hits | 977500435.000 |
| redis | keyspace_misses | 5366300.000 |
| redis | net_input_bytes | 40594069133.000 |
| redis | net_output_bytes | 180216229885.000 |
| redis | ops_per_sec | 1295.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39346.000 |
| redis | used_memory_bytes | 97692288.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
