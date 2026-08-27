# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-signed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:31:10+08:00
- 采样时长：10.0081786s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 23514 | 23514 | 0 | 0 | 2349.88 | 6.476 | 12.518 | 14.345 | 18.166 | 36.567 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 33245 | 1.414 |
| counter | 15800 | 0.672 |
| mysql | 31 | 0.001 |
| redis | 74708 | 3.177 |
| relation | 15800 | 0.672 |

- Cold compute：17740（0.754 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 2037 |
| l1_stale | 4 |
| l2_fresh | 4789 |
| l2_stale | 2852 |
| miss | 13832 |

- L1+L2 Fresh ratio：29.03%
- Refresh max：queue=0 active=4 pending=4

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17740 | 1.165 |
| counter | 15800 | 1.491 |
| hydrate | 17740 | 1.117 |
| inbox | 17740 | 1.164 |
| merge_dedup | 17740 | 0.002 |
| relation | 15800 | 1.972 |
| route | 17740 | 3.106 |
| total | 23514 | 5.844 |

## Redis 本轮边界增量

- Commands：282205；input：37346192 bytes；output：41156032 bytes
- Hits/Misses：346135/46030；run hit rate：88.26%
- Evicted/Rejected：0/0；ops/s max：34299；safety epoch：562 -> 562

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.201 |
| client:loadtest | cpu_percent_total | 51.208 |
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
| docker:zg-canal | cpu_percent | 0.120 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.500 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.990 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 196.080 |
| docker:zg-kafka | memory_percent | 8.830 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 64.060 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23616715.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 85.401 |
| process:counter | cpu_seconds_total | 442.406 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53846016.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 198.864 |
| process:gateway | cpu_seconds_total | 1799.094 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50651136.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 310.290 |
| process:knowpost | cpu_seconds_total | 3813.484 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 172412928.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 140.117 |
| process:relation | cpu_seconds_total | 472.203 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 60821504.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36175872.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 85.891 |
| process:user-storage | cpu_seconds_total | 438.859 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 53002240.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 216879645.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 562.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 743676.000 |
| redis | keyspace_hits | 1016361145.000 |
| redis | keyspace_misses | 6659701.000 |
| redis | net_input_bytes | 44321770544.000 |
| redis | net_output_bytes | 191435070231.000 |
| redis | ops_per_sec | 34299.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51059.000 |
| redis | used_memory_bytes | 159980096.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
