# Feed 压测报告：hybrid / gateway / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:56:16+08:00
- 采样时长：1m0.0304021s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 143657 | 143657 | 0 | 0 | 2393.49 | 25.291 | 30.323 | 33.630 | 93.204 | 125.667 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 140 | 0.001 |
| redis | 431111 | 3.001 |
| relation | 143657 | 1.000 |

- Cold compute：143657（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 143657 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 143657 | 0.246 |
| counter | 12 | 1.884 |
| hydrate | 143657 | 0.365 |
| inbox | 143657 | 0.271 |
| merge_dedup | 143657 | 0.013 |
| relation | 143657 | 24.037 |
| route | 143657 | 0.003 |
| total | 143657 | 24.959 |

## Redis 本轮边界增量

- Commands：1205821；input：274357086 bytes；output：1251151047 bytes
- Hits/Misses：6608472/143797；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：21549；safety epoch：433 -> 433

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.613 |
| client:loadtest | cpu_percent_total | 57.809 |
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
| docker:zg-canal | cpu_percent | 3.160 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.290 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.780 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 254.460 |
| docker:zg-kafka | memory_percent | 8.480 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 69.290 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17643032.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 51.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.822 |
| process:counter | cpu_seconds_total | 207.422 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42774528.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 234.274 |
| process:gateway | cpu_seconds_total | 862.203 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 54362112.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 278.553 |
| process:knowpost | cpu_seconds_total | 7862.625 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 102416384.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 260.998 |
| process:relation | cpu_seconds_total | 7283.141 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 94613504.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.318 |
| process:search | cpu_seconds_total | 1.469 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37064704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 129.044 |
| process:user-storage | cpu_seconds_total | 455.281 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 60919808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 144210546.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 433.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 688868.000 |
| redis | keyspace_hits | 860920524.000 |
| redis | keyspace_misses | 2828653.000 |
| redis | net_input_bytes | 34863775047.000 |
| redis | net_output_bytes | 157908499575.000 |
| redis | ops_per_sec | 21549.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23815.000 |
| redis | used_memory_bytes | 106179176.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
