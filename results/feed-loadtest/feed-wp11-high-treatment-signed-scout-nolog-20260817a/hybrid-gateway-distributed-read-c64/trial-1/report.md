# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-high-treatment-signed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:31:45+08:00
- 采样时长：10.0440097s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 13799 | 13799 | 0 | 0 | 1374.00 | 51.823 | 80.781 | 88.109 | 106.912 | 158.373 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 17965 | 1.302 |
| counter | 10442 | 0.757 |
| mysql | 22 | 0.002 |
| redis | 47658 | 3.454 |
| relation | 10442 | 0.757 |

- Cold compute：11538（0.836 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 727 |
| l1_stale | 6 |
| l2_fresh | 1837 |
| l2_stale | 1285 |
| miss | 9944 |

- L1+L2 Fresh ratio：18.58%
- Refresh max：queue=0 active=25 pending=25

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 11538 | 8.106 |
| counter | 10442 | 8.991 |
| hydrate | 11538 | 7.951 |
| inbox | 11538 | 8.105 |
| merge_dedup | 11538 | 0.002 |
| relation | 10442 | 9.397 |
| route | 11538 | 16.664 |
| total | 13799 | 44.828 |

## Redis 本轮边界增量

- Commands：177648；input：23835633 bytes；output：25102796 bytes
- Hits/Misses：223048/31031；run hit rate：87.79%
- Evicted/Rejected：0/0；ops/s max：18626；safety epoch：564 -> 564

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.732 |
| client:loadtest | cpu_percent_total | 43.714 |
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
| docker:zg-canal | cpu_percent | 2.700 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.170 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.510 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 152.420 |
| docker:zg-kafka | memory_percent | 8.730 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 56.190 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23645030.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 76.747 |
| process:counter | cpu_seconds_total | 461.328 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54202368.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 176.751 |
| process:gateway | cpu_seconds_total | 1840.844 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54480896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 293.188 |
| process:knowpost | cpu_seconds_total | 3884.984 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 155979776.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 161.628 |
| process:relation | cpu_seconds_total | 503.141 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 61616128.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.359 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36175872.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 71.921 |
| process:user-storage | cpu_seconds_total | 456.703 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 55410688.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 217389448.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 564.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 763099.000 |
| redis | keyspace_hits | 1016975545.000 |
| redis | keyspace_misses | 6738043.000 |
| redis | net_input_bytes | 44387862003.000 |
| redis | net_output_bytes | 191510890365.000 |
| redis | ops_per_sec | 18626.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51094.000 |
| redis | used_memory_bytes | 182384608.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
