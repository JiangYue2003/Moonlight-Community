# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:23:35+08:00
- 采样时长：1m0.0258824s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177620 | 177620 | 0 | 0 | 2959.65 | 21.377 | 24.507 | 25.861 | 29.792 | 50.691 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 113 | 0.001 |
| redis | 532973 | 3.001 |
| relation | 177620 | 1.000 |

- Cold compute：177620（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 177620 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 177620 | 0.176 |
| counter | 240 | 0.715 |
| hydrate | 177620 | 0.276 |
| inbox | 177620 | 0.192 |
| merge_dedup | 177620 | 0.012 |
| relation | 177620 | 20.482 |
| route | 177620 | 0.004 |
| total | 177620 | 21.165 |

## Redis 本轮边界增量

- Commands：1486233；input：338642590 bytes；output：1547057540 bytes
- Hits/Misses：8354117/113；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26161；safety epoch：407 -> 407

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.772 |
| client:loadtest | cpu_percent_total | 76.347 |
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
| docker:zg-canal | cpu_percent | 2.330 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.790 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.640 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 168.520 |
| docker:zg-kafka | memory_percent | 8.420 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 53.520 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 12549577.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 9.384 |
| process:counter | cpu_seconds_total | 84.156 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44007424.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.359 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37539840.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 316.128 |
| process:knowpost | cpu_seconds_total | 3045.812 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 96370688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 279.003 |
| process:relation | cpu_seconds_total | 2879.984 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90529792.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.609 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37122048.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.772 |
| process:user-storage | cpu_seconds_total | 0.609 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35278848.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 101474121.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 407.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698768.000 |
| redis | keyspace_hits | 580936838.000 |
| redis | keyspace_misses | 96173.000 |
| redis | net_input_bytes | 23626618729.000 |
| redis | net_output_bytes | 106160401077.000 |
| redis | ops_per_sec | 26161.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21854.000 |
| redis | used_memory_bytes | 106512720.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
