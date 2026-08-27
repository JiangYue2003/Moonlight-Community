# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:01:31+08:00
- 采样时长：1m0.0242135s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 304930 | 304930 | 0 | 0 | 5081.87 | 6.118 | 8.144 | 9.301 | 12.294 | 42.193 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2472 | 0.008 |
| counter | 184 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2417 | 0.008 |
| relation | 184 | 0.001 |

- Cold compute：370（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 303568 |
| l1_stale | 0 |
| l2_fresh | 1362 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 370 | 0.251 |
| counter | 184 | 0.864 |
| hydrate | 370 | 0.321 |
| inbox | 370 | 0.249 |
| merge_dedup | 370 | 0.016 |
| relation | 184 | 2.056 |
| route | 370 | 1.468 |
| total | 304930 | 0.008 |

## Redis 本轮边界增量

- Commands：63030；input：5612952 bytes；output：7018636 bytes
- Hits/Misses：24256/184；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1397；safety epoch：511 -> 511

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.069 |
| client:loadtest | cpu_percent_total | 113.105 |
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
| docker:zg-canal | cpu_percent | 2.870 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.220 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.190 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 218.910 |
| docker:zg-kafka | memory_percent | 8.610 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 70.330 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22731944.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.828 |
| process:counter | cpu_seconds_total | 205.156 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44896256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 460.674 |
| process:gateway | cpu_seconds_total | 3232.625 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 54693888.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 199.550 |
| process:knowpost | cpu_seconds_total | 8739.344 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 133144576.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 10.078 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 49369088.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.318 |
| process:search | cpu_seconds_total | 1.859 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37003264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 217.038 |
| process:user-storage | cpu_seconds_total | 1618.953 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 87293952.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196125553.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 511.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677335.000 |
| redis | keyspace_hits | 993242257.000 |
| redis | keyspace_misses | 5380850.000 |
| redis | net_input_bytes | 41905768542.000 |
| redis | net_output_bytes | 187043976337.000 |
| redis | ops_per_sec | 1397.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42130.000 |
| redis | used_memory_bytes | 100556480.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
