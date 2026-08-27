# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:29:52+08:00
- 采样时长：1m0.0409596s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 175678 | 175678 | 0 | 0 | 2926.60 | 43.354 | 49.077 | 51.611 | 61.555 | 99.377 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 118 | 0.001 |
| redis | 527152 | 3.001 |
| relation | 175678 | 1.000 |

- Cold compute：175678（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 175678 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 175678 | 0.183 |
| counter | 240 | 0.737 |
| hydrate | 175678 | 0.284 |
| inbox | 175678 | 0.202 |
| merge_dedup | 175678 | 0.012 |
| relation | 175678 | 42.552 |
| route | 175678 | 0.004 |
| total | 175678 | 43.259 |

## Redis 本轮边界增量

- Commands：1470706；input：334990746 bytes；output：1530157793 bytes
- Hits/Misses：8262834/122；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26335；safety epoch：412 -> 412

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.775 |
| client:loadtest | cpu_percent_total | 76.406 |
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
| docker:zg-canal | cpu_percent | 2.650 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.340 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.800 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 189.580 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 56.170 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 13592765.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 12.288 |
| process:counter | cpu_seconds_total | 108.000 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44224512.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.545 |
| process:gateway | cpu_seconds_total | 0.484 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37777408.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 292.948 |
| process:knowpost | cpu_seconds_total | 4002.859 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 102088704.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.966 |
| process:relation | cpu_seconds_total | 3777.766 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 94195712.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.809 |
| process:search | cpu_seconds_total | 0.750 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37175296.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.783 |
| process:user-storage | cpu_seconds_total | 0.750 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35438592.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 110213584.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 412.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698738.000 |
| redis | keyspace_hits | 629919215.000 |
| redis | keyspace_misses | 96875.000 |
| redis | net_input_bytes | 25613925383.000 |
| redis | net_output_bytes | 115231562958.000 |
| redis | ops_per_sec | 26335.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22231.000 |
| redis | used_memory_bytes | 107572376.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
