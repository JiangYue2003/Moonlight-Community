# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:10:58+08:00
- 采样时长：1m0.0270353s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 150063 | 150063 | 0 | 0 | 2500.40 | 25.063 | 29.602 | 31.621 | 36.922 | 64.272 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 104 | 0.001 |
| redis | 450293 | 3.001 |
| relation | 150063 | 1.000 |

- Cold compute：150063（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 150063 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 150063 | 0.218 |
| counter | 240 | 1.052 |
| hydrate | 150063 | 0.328 |
| inbox | 150063 | 0.237 |
| merge_dedup | 150063 | 0.014 |
| relation | 150063 | 23.087 |
| route | 150063 | 0.005 |
| total | 150063 | 23.913 |

## Redis 本轮边界增量

- Commands：1258926；input：286296451 bytes；output：1306974977 bytes
- Hits/Misses：6908884/150167；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22460；safety epoch：444 -> 444

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.875 |
| client:loadtest | cpu_percent_total | 62.003 |
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
| docker:zg-canal | cpu_percent | 3.410 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.730 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.560 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 270.010 |
| docker:zg-kafka | memory_percent | 8.490 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 67.430 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 101.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19463117.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 50.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.904 |
| process:counter | cpu_seconds_total | 693.359 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44642304.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 251.540 |
| process:gateway | cpu_seconds_total | 2411.031 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 55730176.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 268.342 |
| process:knowpost | cpu_seconds_total | 9814.672 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 103407616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 266.349 |
| process:relation | cpu_seconds_total | 8916.734 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87904256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 5.953 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37281792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 126.047 |
| process:user-storage | cpu_seconds_total | 1267.125 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 58675200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 169024820.000 |
| redis | connected_clients | 153.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 444.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676532.000 |
| redis | keyspace_hits | 944503628.000 |
| redis | keyspace_misses | 4647254.000 |
| redis | net_input_bytes | 39004846879.000 |
| redis | net_output_bytes | 173909966281.000 |
| redis | ops_per_sec | 22460.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35497.000 |
| redis | used_memory_bytes | 99185872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
