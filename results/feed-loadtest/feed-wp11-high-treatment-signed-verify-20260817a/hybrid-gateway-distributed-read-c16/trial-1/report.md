# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-signed-verify-20260817a`
- 开始时间：2026-08-17T00:30:10+08:00
- 采样时长：5.0076376s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 16066 | 16066 | 0 | 0 | 3208.98 | 5.373 | 8.050 | 8.834 | 10.604 | 30.166 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 16515 | 1.028 |
| counter | 11013 | 0.685 |
| mysql | 28 | 0.002 |
| redis | 49572 | 3.086 |
| relation | 11013 | 0.685 |

- Cold compute：11733（0.730 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 1710 |
| l1_stale | 1 |
| l2_fresh | 3090 |
| l2_stale | 242 |
| miss | 11023 |

- L1+L2 Fresh ratio：29.88%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 11733 | 0.671 |
| counter | 11013 | 0.938 |
| hydrate | 11733 | 0.625 |
| inbox | 11733 | 0.669 |
| merge_dedup | 11733 | 0.003 |
| relation | 11013 | 1.431 |
| route | 11733 | 2.246 |
| total | 16066 | 4.061 |

## Redis 本轮边界增量

- Commands：188582；input：24689512 bytes；output：25797823 bytes
- Hits/Misses：229778/32925；run hit rate：87.47%
- Evicted/Rejected：0/0；ops/s max：40246；safety epoch：561 -> 561

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.569 |
| client:loadtest | cpu_percent_total | 57.100 |
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
| docker:zg-canal | cpu_percent | 2.590 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.490 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 0.790 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 83.010 |
| docker:zg-kafka | memory_percent | 8.310 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.240 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23595996.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 74.665 |
| process:counter | cpu_seconds_total | 429.672 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53813248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 215.672 |
| process:gateway | cpu_seconds_total | 1777.312 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50638848.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 315.952 |
| process:knowpost | cpu_seconds_total | 3775.422 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 174919680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 138.327 |
| process:relation | cpu_seconds_total | 457.094 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 61726720.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.328 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36188160.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 102.406 |
| process:user-storage | cpu_seconds_total | 429.469 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 53342208.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 216459164.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 561.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 731811.000 |
| redis | keyspace_hits | 1015920899.000 |
| redis | keyspace_misses | 6599622.000 |
| redis | net_input_bytes | 44270032695.000 |
| redis | net_output_bytes | 191383005344.000 |
| redis | ops_per_sec | 40246.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 50994.000 |
| redis | used_memory_bytes | 147711416.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
