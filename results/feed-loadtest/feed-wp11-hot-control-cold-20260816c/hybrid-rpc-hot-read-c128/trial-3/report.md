# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:14:45+08:00
- 采样时长：1m0.0383485s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182089 | 182089 | 0 | 0 | 3033.51 | 42.241 | 48.585 | 51.170 | 58.783 | 91.834 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 102 | 0.001 |
| redis | 546369 | 3.001 |
| relation | 182089 | 1.000 |

- Cold compute：182089（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182089 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182089 | 0.215 |
| counter | 12 | 1.777 |
| hydrate | 182089 | 0.309 |
| inbox | 182089 | 0.240 |
| merge_dedup | 182089 | 0.011 |
| relation | 182089 | 40.931 |
| route | 182089 | 0.003 |
| total | 182089 | 41.730 |

## Redis 本轮边界增量

- Commands：1516280；input：346893958 bytes；output：1585795004 bytes
- Hits/Misses：8558465/108；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27870；safety epoch：400 -> 400

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.686 |
| client:loadtest | cpu_percent_total | 74.978 |
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
| docker:zg-canal | cpu_percent | 3.000 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.270 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.470 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 180.360 |
| docker:zg-kafka | memory_percent | 8.390 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 59.410 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11068186.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.059 |
| process:counter | cpu_seconds_total | 50.516 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40751104.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.320 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37384192.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 301.107 |
| process:knowpost | cpu_seconds_total | 1719.109 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 97038336.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 289.951 |
| process:relation | cpu_seconds_total | 1624.672 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 89845760.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.328 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37138432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.344 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35405824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 89089216.000 |
| redis | connected_clients | 186.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 400.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698821.000 |
| redis | keyspace_hits | 511393630.000 |
| redis | keyspace_misses | 95201.000 |
| redis | net_input_bytes | 20805234133.000 |
| redis | net_output_bytes | 93278223458.000 |
| redis | ops_per_sec | 27870.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21324.000 |
| redis | used_memory_bytes | 107214992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
