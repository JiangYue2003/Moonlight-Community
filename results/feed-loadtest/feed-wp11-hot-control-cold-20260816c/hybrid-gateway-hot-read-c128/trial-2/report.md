# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:58:47+08:00
- 采样时长：1m0.039953s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 143431 | 143431 | 0 | 0 | 2389.34 | 51.485 | 61.365 | 67.923 | 122.964 | 167.778 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 181 | 0.001 |
| redis | 430474 | 3.001 |
| relation | 143431 | 1.000 |

- Cold compute：143431（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 143431 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 143431 | 0.557 |
| counter | 12 | 1.808 |
| hydrate | 143431 | 0.709 |
| inbox | 143431 | 0.629 |
| merge_dedup | 143431 | 0.012 |
| relation | 143431 | 49.707 |
| route | 143431 | 0.003 |
| total | 143431 | 51.642 |

## Redis 本轮边界增量

- Commands：1202945；input：273864163 bytes；output：1249155815 bytes
- Hits/Misses：6598028/143618；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：21944；safety epoch：435 -> 435

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.621 |
| client:loadtest | cpu_percent_total | 57.930 |
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
| docker:zg-canal | cpu_percent | 3.440 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.050 |
| docker:zg-es | memory_percent | 11.860 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 8.690 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 303.690 |
| docker:zg-kafka | memory_percent | 8.500 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 69.280 |
| docker:zg-zk | memory_percent | 1.060 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17983235.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 76.000 |
| mysql | threads_running | 8.000 |
| process:counter | cpu_percent | 11.586 |
| process:counter | cpu_seconds_total | 216.750 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42786816.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 239.627 |
| process:gateway | cpu_seconds_total | 1156.578 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 59912192.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 265.095 |
| process:knowpost | cpu_seconds_total | 8192.844 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 108363776.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 254.275 |
| process:relation | cpu_seconds_total | 7590.156 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 93048832.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.562 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37072896.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 137.727 |
| process:user-storage | cpu_seconds_total | 608.625 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 66617344.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 147066091.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 435.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 686565.000 |
| redis | keyspace_hits | 876528827.000 |
| redis | keyspace_misses | 3168319.000 |
| redis | net_input_bytes | 35512303809.000 |
| redis | net_output_bytes | 160863679018.000 |
| redis | ops_per_sec | 21944.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23966.000 |
| redis | used_memory_bytes | 108992688.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
