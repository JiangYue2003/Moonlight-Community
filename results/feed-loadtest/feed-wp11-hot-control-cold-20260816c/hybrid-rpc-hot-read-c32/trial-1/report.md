# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:04:42+08:00
- 采样时长：1m0.0187097s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182516 | 182516 | 0 | 0 | 3041.62 | 10.376 | 12.438 | 13.290 | 15.419 | 32.439 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 101 | 0.001 |
| redis | 547649 | 3.001 |
| relation | 182516 | 1.000 |

- Cold compute：182516（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182516 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182516 | 0.196 |
| counter | 12 | 1.893 |
| hydrate | 182516 | 0.288 |
| inbox | 182516 | 0.215 |
| merge_dedup | 182516 | 0.011 |
| relation | 182516 | 9.332 |
| route | 182516 | 0.003 |
| total | 182516 | 10.066 |

## Redis 本轮边界增量

- Commands：1519749；input：347699187 bytes；output：1589511881 bytes
- Hits/Misses：8578541/101；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26996；safety epoch：392 -> 392

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.221 |
| client:loadtest | cpu_percent_total | 67.531 |
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
| docker:zg-canal | cpu_percent | 2.450 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.750 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.380 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 216.450 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 55.400 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9338106.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.829 |
| process:counter | cpu_seconds_total | 13.391 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40083456.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.172 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37011456.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 300.000 |
| process:knowpost | cpu_seconds_total | 216.953 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 75522048.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.952 |
| process:relation | cpu_seconds_total | 193.172 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 66326528.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.188 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37588992.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.628 |
| process:user-storage | cpu_seconds_total | 0.125 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35110912.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 74670112.000 |
| redis | connected_clients | 76.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 392.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698915.000 |
| redis | keyspace_hits | 430196258.000 |
| redis | keyspace_misses | 94141.000 |
| redis | net_input_bytes | 17511762599.000 |
| redis | net_output_bytes | 78232619702.000 |
| redis | ops_per_sec | 26996.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20721.000 |
| redis | used_memory_bytes | 103486736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
