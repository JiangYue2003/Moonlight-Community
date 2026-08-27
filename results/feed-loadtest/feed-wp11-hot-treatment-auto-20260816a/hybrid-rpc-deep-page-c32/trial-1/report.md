# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:28:53+08:00
- 采样时长：1m0.0174345s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 175995 | 175995 | 0 | 0 | 2932.98 | 10.560 | 12.710 | 13.590 | 18.827 | 114.332 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 220 | 0.001 |
| mysql | 175995 | 1.000 |
| redis | 175995 | 1.000 |
| relation | 220 | 0.001 |

- Cold compute：175995（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 175995 |
| l1_fresh | 0 |
| l1_stale | 0 |
| l2_fresh | 0 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 175995 | 0.159 |
| counter | 220 | 0.777 |
| hydrate | 175995 | 5.769 |
| inbox | 175995 | 0.157 |
| merge_dedup | 175995 | 0.014 |
| relation | 220 | 1.496 |
| route | 175995 | 0.008 |
| total | 175995 | 5.974 |

## Redis 本轮边界增量

- Commands：1115574；input：80307602 bytes；output：464164912 bytes
- Hits/Misses：1062740/220；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：20932；safety epoch：488 -> 488

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.832 |
| client:loadtest | cpu_percent_total | 61.310 |
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
| docker:zg-canal | cpu_percent | 2.180 |
| docker:zg-canal | memory_percent | 7.850 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 3.760 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.040 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 192.860 |
| docker:zg-kafka | memory_percent | 8.610 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 55.390 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20393817.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 51.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.055 |
| process:counter | cpu_seconds_total | 94.547 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 42991616.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37048320.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 498.287 |
| process:knowpost | cpu_seconds_total | 3951.828 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 91123712.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.870 |
| process:relation | cpu_seconds_total | 4.344 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46850048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.953 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37097472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.734 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34783232.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 180277089.000 |
| redis | connected_clients | 67.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 488.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677182.000 |
| redis | keyspace_hits | 979055630.000 |
| redis | keyspace_misses | 5372165.000 |
| redis | net_input_bytes | 40760103995.000 |
| redis | net_output_bytes | 180850757571.000 |
| redis | ops_per_sec | 20932.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 40172.000 |
| redis | used_memory_bytes | 98389672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
