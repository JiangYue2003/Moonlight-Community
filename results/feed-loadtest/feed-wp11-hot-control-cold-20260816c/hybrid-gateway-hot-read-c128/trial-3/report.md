# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T17:00:02+08:00
- 采样时长：1m0.0435735s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 141302 | 141302 | 0 | 0 | 2353.71 | 52.751 | 63.386 | 68.651 | 101.778 | 176.750 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 166 | 0.001 |
| redis | 424072 | 3.001 |
| relation | 141302 | 1.000 |

- Cold compute：141302（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 141302 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 141302 | 0.802 |
| counter | 12 | 2.567 |
| hydrate | 141302 | 0.986 |
| inbox | 141302 | 0.917 |
| merge_dedup | 141302 | 0.013 |
| relation | 141302 | 49.592 |
| route | 141302 | 0.003 |
| total | 141302 | 52.337 |

## Redis 本轮边界增量

- Commands：1185989；input：269859368 bytes；output：1230636613 bytes
- Hits/Misses：6500116/141468；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：21665；safety epoch：436 -> 436

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.733 |
| client:loadtest | cpu_percent_total | 59.722 |
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
| docker:zg-canal | cpu_percent | 0.210 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.810 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.920 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 290.450 |
| docker:zg-kafka | memory_percent | 8.470 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 76.950 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18150091.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 9.290 |
| process:counter | cpu_seconds_total | 221.031 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42532864.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 254.119 |
| process:gateway | cpu_seconds_total | 1304.531 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 59793408.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.668 |
| process:knowpost | cpu_seconds_total | 8361.469 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 107909120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 242.533 |
| process:relation | cpu_seconds_total | 7742.297 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90619904.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 1.625 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37081088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 131.143 |
| process:user-storage | cpu_seconds_total | 683.750 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 56610816.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 148467547.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 436.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 685597.000 |
| redis | keyspace_hits | 884186021.000 |
| redis | keyspace_misses | 3334969.000 |
| redis | net_input_bytes | 35830507245.000 |
| redis | net_output_bytes | 162313456134.000 |
| redis | ops_per_sec | 21665.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24041.000 |
| redis | used_memory_bytes | 109207776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
