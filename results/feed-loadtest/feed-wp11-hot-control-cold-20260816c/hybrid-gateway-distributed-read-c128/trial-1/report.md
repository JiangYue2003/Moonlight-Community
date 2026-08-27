# Feed 压测报告：hybrid / gateway / distributed-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:12:14+08:00
- 采样时长：1m0.0476986s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 148455 | 148455 | 0 | 0 | 2472.74 | 50.969 | 59.608 | 63.110 | 73.891 | 114.908 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 110 | 0.001 |
| redis | 445475 | 3.001 |
| relation | 148455 | 1.000 |

- Cold compute：148455（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 148455 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 148455 | 0.261 |
| counter | 240 | 0.937 |
| hydrate | 148455 | 0.377 |
| inbox | 148455 | 0.287 |
| merge_dedup | 148455 | 0.013 |
| relation | 148455 | 49.064 |
| route | 148455 | 0.004 |
| total | 148455 | 50.031 |

## Redis 本轮边界增量

- Commands：1246135；input：283276612 bytes；output：1292984495 bytes
- Hits/Misses：6834910/148565；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22607；safety epoch：445 -> 445

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.815 |
| client:loadtest | cpu_percent_total | 61.045 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.060 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.700 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 260.440 |
| docker:zg-kafka | memory_percent | 8.540 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 74.810 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 88.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19639065.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.048 |
| process:counter | cpu_seconds_total | 697.391 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44748800.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 239.554 |
| process:gateway | cpu_seconds_total | 2558.281 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 60719104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 265.883 |
| process:knowpost | cpu_seconds_total | 9982.109 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 108531712.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 248.768 |
| process:relation | cpu_seconds_total | 9076.828 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 93564928.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.810 |
| process:search | cpu_seconds_total | 5.984 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37281792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 140.775 |
| process:user-storage | cpu_seconds_total | 1345.578 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 66150400.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 170503181.000 |
| redis | connected_clients | 153.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 445.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676577.000 |
| redis | keyspace_hits | 952586038.000 |
| redis | keyspace_misses | 4822933.000 |
| redis | net_input_bytes | 39340149518.000 |
| redis | net_output_bytes | 175439007152.000 |
| redis | ops_per_sec | 22607.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35573.000 |
| redis | used_memory_bytes | 99763192.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
