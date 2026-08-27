# Feed 压测报告：hybrid / gateway / distributed-read-c256

- Run ID：`feed-wp11-hot-treatment-diagnostic-c256-20260816b`
- 开始时间：2026-08-16T22:18:13+08:00
- 采样时长：10.040554s
- 并发：256
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 50566 | 50429 | 137 | 0 | 5024.13 | 49.934 | 62.370 | 87.917 | 115.992 | 255.582 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 383 | 0.008 |
| counter | 20 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 377 | 0.007 |
| relation | 20 | 0.000 |

- Cold compute：54（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 50208 |
| l1_stale | 0 |
| l2_fresh | 221 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 54 | 0.179 |
| counter | 20 | 0.910 |
| hydrate | 54 | 0.304 |
| inbox | 54 | 0.179 |
| merge_dedup | 54 | 0.019 |
| relation | 20 | 2.041 |
| route | 54 | 1.103 |
| total | 50429 | 0.009 |

## Redis 本轮边界增量

- Commands：10161；input：894001 bytes；output：1093446 bytes
- Hits/Misses：3405/20；run hit rate：99.42%
- Evicted/Rejected：0/0；ops/s max：1399；safety epoch：520 -> 520

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.139 |
| client:loadtest | cpu_percent_total | 114.224 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.060 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 105.630 |
| docker:zg-kafka | memory_percent | 8.170 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 61.930 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22734499.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 33.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.971 |
| process:counter | cpu_seconds_total | 257.438 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44036096.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 419.824 |
| process:gateway | cpu_seconds_total | 5071.906 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 68673536.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 187.894 |
| process:knowpost | cpu_seconds_total | 9632.859 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 101302272.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.772 |
| process:relation | cpu_seconds_total | 12.453 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 47153152.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.391 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37023744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 210.072 |
| process:user-storage | cpu_seconds_total | 2540.281 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 53194752.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 197049630.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 520.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677537.000 |
| redis | keyspace_hits | 993440775.000 |
| redis | keyspace_misses | 5384526.000 |
| redis | net_input_bytes | 41980823107.000 |
| redis | net_output_bytes | 187108328227.000 |
| redis | ops_per_sec | 1399.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43082.000 |
| redis | used_memory_bytes | 100869008.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
