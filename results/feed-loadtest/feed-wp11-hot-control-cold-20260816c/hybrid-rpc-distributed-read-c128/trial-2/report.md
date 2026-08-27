# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:28:36+08:00
- 采样时长：1m0.03933s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177755 | 177755 | 0 | 0 | 2961.24 | 42.839 | 48.643 | 51.017 | 61.297 | 92.437 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 116 | 0.001 |
| redis | 533381 | 3.001 |
| relation | 177755 | 1.000 |

- Cold compute：177755（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 177755 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 177755 | 0.170 |
| counter | 240 | 0.700 |
| hydrate | 177755 | 0.271 |
| inbox | 177755 | 0.186 |
| merge_dedup | 177755 | 0.012 |
| relation | 177755 | 42.094 |
| route | 177755 | 0.004 |
| total | 177755 | 42.760 |

## Redis 本轮边界增量

- Commands：1487322；input：338899585 bytes；output：1548231778 bytes
- Hits/Misses：8360453/122；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26226；safety epoch：411 -> 411

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.497 |
| client:loadtest | cpu_percent_total | 71.958 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.840 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.840 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 193.600 |
| docker:zg-kafka | memory_percent | 8.430 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 28.370 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 13385225.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.063 |
| process:counter | cpu_seconds_total | 103.109 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44257280.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.797 |
| process:gateway | cpu_seconds_total | 0.453 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38608896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 299.278 |
| process:knowpost | cpu_seconds_total | 3811.094 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 100880384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 281.996 |
| process:relation | cpu_seconds_total | 3599.922 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 94060544.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 0.734 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37142528.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.688 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35405824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 108473952.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 411.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698743.000 |
| redis | keyspace_hits | 620175474.000 |
| redis | keyspace_misses | 96734.000 |
| redis | net_input_bytes | 25218535813.000 |
| redis | net_output_bytes | 113427106651.000 |
| redis | ops_per_sec | 26226.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22155.000 |
| redis | used_memory_bytes | 106672776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
