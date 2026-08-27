# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:42:57+08:00
- 采样时长：10.0075102s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 32860 | 32860 | 0 | 0 | 3284.42 | 3.938 | 8.681 | 11.354 | 17.071 | 36.959 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 19886 | 0.605 |
| counter | 1211 | 0.037 |
| mysql | 0 | 0.000 |
| redis | 19771 | 0.602 |
| relation | 1211 | 0.037 |

- Cold compute：3290（0.100 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 22821 |
| l1_stale | 23 |
| l2_fresh | 9573 |
| l2_stale | 443 |
| miss | 0 |

- L1+L2 Fresh ratio：98.58%
- Refresh max：queue=0 active=14 pending=14

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3290 | 3.631 |
| counter | 1211 | 6.384 |
| hydrate | 3290 | 3.599 |
| inbox | 3290 | 3.629 |
| merge_dedup | 3290 | 0.003 |
| relation | 1211 | 7.872 |
| route | 3290 | 5.270 |
| total | 32860 | 1.631 |

## Redis 本轮边界增量

- Commands：58220；input：9661819 bytes；output：22328958 bytes
- Hits/Misses：86869/1479；run hit rate：98.33%
- Evicted/Rejected：0/0；ops/s max：10885；safety epoch：538 -> 538

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 12.647 |
| client:loadtest | cpu_percent_total | 202.348 |
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
| docker:zg-canal | cpu_percent | 1.690 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.930 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.780 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 246.680 |
| docker:zg-kafka | memory_percent | 8.740 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 58.050 |
| docker:zg-zk | memory_percent | 1.420 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23278286.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 44.011 |
| process:counter | cpu_seconds_total | 24.969 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 50245632.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 552.224 |
| process:gateway | cpu_seconds_total | 91.719 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 53587968.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 242.859 |
| process:knowpost | cpu_seconds_total | 261.484 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 99446784.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 70.264 |
| process:relation | cpu_seconds_total | 16.891 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56156160.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36184064.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 125.085 |
| process:user-storage | cpu_seconds_total | 28.281 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 50380800.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205741904.000 |
| redis | connected_clients | 245.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 538.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 693839.000 |
| redis | keyspace_hits | 1003961982.000 |
| redis | keyspace_misses | 5979399.000 |
| redis | net_input_bytes | 42830741419.000 |
| redis | net_output_bytes | 188807179846.000 |
| redis | ops_per_sec | 10885.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48166.000 |
| redis | used_memory_bytes | 116743464.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
