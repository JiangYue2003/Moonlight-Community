# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:10:59+08:00
- 采样时长：1m0.0275963s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 184602 | 184602 | 0 | 0 | 3076.01 | 20.625 | 24.363 | 25.938 | 30.088 | 50.138 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 125 | 0.001 |
| redis | 553931 | 3.001 |
| relation | 184602 | 1.000 |

- Cold compute：184602（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 184602 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 184602 | 0.200 |
| counter | 12 | 1.805 |
| hydrate | 184602 | 0.291 |
| inbox | 184602 | 0.219 |
| merge_dedup | 184602 | 0.011 |
| relation | 184602 | 19.610 |
| route | 184602 | 0.003 |
| total | 184602 | 20.356 |

## Redis 本轮边界增量

- Commands：1536402；input：351627070 bytes；output：1607659815 bytes
- Hits/Misses：8676558/126；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27892；safety epoch：397 -> 397

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.453 |
| client:loadtest | cpu_percent_total | 71.243 |
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
| docker:zg-canal | cpu_percent | 2.750 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.120 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.870 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 222.190 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 51.990 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 10418561.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.822 |
| process:counter | cpu_seconds_total | 36.172 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40632320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.219 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38019072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 302.055 |
| process:knowpost | cpu_seconds_total | 1153.781 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 92909568.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 279.336 |
| process:relation | cpu_seconds_total | 1091.094 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 73895936.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.297 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37101568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.778 |
| process:user-storage | cpu_seconds_total | 0.234 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35229696.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 83675956.000 |
| redis | connected_clients | 135.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 397.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698853.000 |
| redis | keyspace_hits | 480908039.000 |
| redis | keyspace_misses | 94813.000 |
| redis | net_input_bytes | 19568730963.000 |
| redis | net_output_bytes | 87629354818.000 |
| redis | ops_per_sec | 27892.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21098.000 |
| redis | used_memory_bytes | 104598272.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
