# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:12:15+08:00
- 采样时长：1m0.0456245s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182850 | 182850 | 0 | 0 | 3045.85 | 42.048 | 48.319 | 50.893 | 58.461 | 95.609 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 115 | 0.001 |
| redis | 548665 | 3.001 |
| relation | 182850 | 1.000 |

- Cold compute：182850（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182850 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182850 | 0.207 |
| counter | 12 | 1.883 |
| hydrate | 182850 | 0.300 |
| inbox | 182850 | 0.227 |
| merge_dedup | 182850 | 0.011 |
| relation | 182850 | 40.795 |
| route | 182850 | 0.003 |
| total | 182850 | 41.563 |

## Redis 本轮边界增量

- Commands：1522143；input：348309592 bytes；output：1592409542 bytes
- Hits/Misses：8594224/116；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27616；safety epoch：398 -> 398

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.822 |
| client:loadtest | cpu_percent_total | 77.155 |
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
| docker:zg-es | cpu_percent | 2.760 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.620 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 202.520 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 55.110 |
| docker:zg-zk | memory_percent | 1.220 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 10635209.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.694 |
| process:counter | cpu_seconds_total | 41.375 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40722432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.219 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38039552.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 303.954 |
| process:knowpost | cpu_seconds_total | 1342.781 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 97972224.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 284.685 |
| process:relation | cpu_seconds_total | 1270.594 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 84987904.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.312 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37101568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.250 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35311616.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 85481289.000 |
| redis | connected_clients | 177.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 398.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698844.000 |
| redis | keyspace_hits | 491073545.000 |
| redis | keyspace_misses | 94950.000 |
| redis | net_input_bytes | 19981066079.000 |
| redis | net_output_bytes | 89512995466.000 |
| redis | ops_per_sec | 27616.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21174.000 |
| redis | used_memory_bytes | 106170384.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
