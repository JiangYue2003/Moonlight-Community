# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:52:30+08:00
- 采样时长：1m0.0189992s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 150806 | 150806 | 0 | 0 | 2513.10 | 12.405 | 14.881 | 16.021 | 19.382 | 39.535 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 113 | 0.001 |
| redis | 452531 | 3.001 |
| relation | 150806 | 1.000 |

- Cold compute：150806（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 150806 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 150806 | 0.219 |
| counter | 12 | 2.078 |
| hydrate | 150806 | 0.329 |
| inbox | 150806 | 0.237 |
| merge_dedup | 150806 | 0.014 |
| relation | 150806 | 10.260 |
| route | 150806 | 0.003 |
| total | 150806 | 11.086 |

## Redis 本轮边界增量

- Commands：1264192；input：287892874 bytes；output：1313379361 bytes
- Hits/Misses：6937337/150935；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22332；safety epoch：430 -> 430

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.723 |
| client:loadtest | cpu_percent_total | 59.564 |
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
| docker:zg-canal | cpu_percent | 1.010 |
| docker:zg-canal | memory_percent | 7.850 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.690 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 7.460 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 257.920 |
| docker:zg-kafka | memory_percent | 8.480 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 76.150 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17130110.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.758 |
| process:counter | cpu_seconds_total | 194.078 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43421696.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 253.654 |
| process:gateway | cpu_seconds_total | 428.875 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 51785728.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 258.208 |
| process:knowpost | cpu_seconds_total | 7370.891 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 93466624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 244.591 |
| process:relation | cpu_seconds_total | 6827.422 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 85803008.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.549 |
| process:search | cpu_seconds_total | 1.219 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37068800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 120.834 |
| process:user-storage | cpu_seconds_total | 227.109 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 66080768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 139900954.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 430.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 693191.000 |
| redis | keyspace_hits | 837370119.000 |
| redis | keyspace_misses | 2316270.000 |
| redis | net_input_bytes | 33885194741.000 |
| redis | net_output_bytes | 153449593875.000 |
| redis | ops_per_sec | 22332.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23589.000 |
| redis | used_memory_bytes | 107022384.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
