# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:26:22+08:00
- 采样时长：1m0.061895s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 484088 | 484088 | 0 | 0 | 8063.65 | 31.406 | 39.141 | 41.039 | 46.921 | 96.314 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2429 | 0.005 |
| counter | 181 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2402 | 0.005 |
| relation | 181 | 0.000 |

- Cold compute：363（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 482748 |
| l1_stale | 0 |
| l2_fresh | 1340 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 363 | 0.153 |
| counter | 181 | 0.665 |
| hydrate | 363 | 0.189 |
| inbox | 363 | 0.153 |
| merge_dedup | 363 | 0.014 |
| relation | 181 | 1.442 |
| route | 363 | 1.069 |
| total | 484088 | 0.005 |

## Redis 本轮边界增量

- Commands：62842；input：5575463 bytes；output：6964389 bytes
- Hits/Misses：23865/181；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1322；safety epoch：486 -> 486

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.974 |
| client:loadtest | cpu_percent_total | 127.577 |
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
| docker:zg-canal | cpu_percent | 2.020 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.280 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.270 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 122.030 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 5.530 |
| docker:zg-zk | memory_percent | 1.370 |
| docker:zg-zk | pids | 101.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20184672.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.814 |
| process:counter | cpu_seconds_total | 86.578 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43692032.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 0.531 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37306368.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 241.787 |
| process:knowpost | cpu_seconds_total | 3499.641 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 122748928.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 3.734 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 47616000.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.875 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37089280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.782 |
| process:user-storage | cpu_seconds_total | 0.703 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34959360.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178872647.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 486.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677168.000 |
| redis | keyspace_hits | 977766998.000 |
| redis | keyspace_misses | 5371100.000 |
| redis | net_input_bytes | 40657537001.000 |
| redis | net_output_bytes | 180292639893.000 |
| redis | ops_per_sec | 1322.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 40021.000 |
| redis | used_memory_bytes | 98193600.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
