# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:09:43+08:00
- 采样时长：1m0.039489s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 146011 | 146011 | 0 | 0 | 2432.45 | 25.855 | 30.864 | 33.002 | 38.483 | 63.194 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 112 | 0.001 |
| redis | 438145 | 3.001 |
| relation | 146011 | 1.000 |

- Cold compute：146011（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 146011 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 146011 | 0.256 |
| counter | 240 | 0.994 |
| hydrate | 146011 | 0.375 |
| inbox | 146011 | 0.290 |
| merge_dedup | 146011 | 0.013 |
| relation | 146011 | 23.589 |
| route | 146011 | 0.005 |
| total | 146011 | 24.554 |

## Redis 本轮边界增量

- Commands：1226770；input：278690346 bytes；output：1271729063 bytes
- Hits/Misses：6722484/146123；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22053；safety epoch：443 -> 443

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.759 |
| client:loadtest | cpu_percent_total | 60.143 |
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
| docker:zg-canal | cpu_percent | 2.780 |
| docker:zg-canal | memory_percent | 7.850 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.730 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 8.130 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 303.410 |
| docker:zg-kafka | memory_percent | 8.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 88.710 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19285487.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 47.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 11.039 |
| process:counter | cpu_seconds_total | 688.891 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44384256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 237.357 |
| process:gateway | cpu_seconds_total | 2257.812 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 56725504.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 261.024 |
| process:knowpost | cpu_seconds_total | 9644.297 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 106717184.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 248.177 |
| process:relation | cpu_seconds_total | 8757.062 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 88788992.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.055 |
| process:search | cpu_seconds_total | 5.938 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37257216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 130.481 |
| process:user-storage | cpu_seconds_total | 1190.484 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 59101184.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 167533643.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 443.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676533.000 |
| redis | keyspace_hits | 936338449.000 |
| redis | keyspace_misses | 4469776.000 |
| redis | net_input_bytes | 38666255584.000 |
| redis | net_output_bytes | 172365279585.000 |
| redis | ops_per_sec | 22053.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35422.000 |
| redis | used_memory_bytes | 100045392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
