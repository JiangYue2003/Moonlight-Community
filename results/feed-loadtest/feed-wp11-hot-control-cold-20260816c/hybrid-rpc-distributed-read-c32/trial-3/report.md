# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:22:19+08:00
- 采样时长：1m0.0192674s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 178693 | 178693 | 0 | 0 | 2977.89 | 10.592 | 12.308 | 13.133 | 15.239 | 25.206 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 102 | 0.001 |
| redis | 536181 | 3.001 |
| relation | 178693 | 1.000 |

- Cold compute：178693（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 178693 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 178693 | 0.170 |
| counter | 240 | 0.700 |
| hydrate | 178693 | 0.269 |
| inbox | 178693 | 0.184 |
| merge_dedup | 178693 | 0.012 |
| relation | 178693 | 9.625 |
| route | 178693 | 0.004 |
| total | 178693 | 10.286 |

## Redis 本轮边界增量

- Commands：1494815；input：340660878 bytes；output：1556394331 bytes
- Hits/Misses：8404550/111；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26238；safety epoch：406 -> 406

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.725 |
| client:loadtest | cpu_percent_total | 75.601 |
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
| docker:zg-canal | cpu_percent | 2.710 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.380 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 5.240 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 210.700 |
| docker:zg-kafka | memory_percent | 8.370 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 50.240 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 12340263.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.606 |
| process:counter | cpu_seconds_total | 79.578 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44007424.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.344 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38236160.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 295.048 |
| process:knowpost | cpu_seconds_total | 2853.234 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 97984512.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 284.340 |
| process:relation | cpu_seconds_total | 2702.594 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87351296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.547 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37101568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 0.547 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35278848.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 99720375.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 406.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698778.000 |
| redis | keyspace_hits | 571108871.000 |
| redis | keyspace_misses | 96027.000 |
| redis | net_input_bytes | 23227852979.000 |
| redis | net_output_bytes | 104340307774.000 |
| redis | ops_per_sec | 26238.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21778.000 |
| redis | used_memory_bytes | 106596480.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
