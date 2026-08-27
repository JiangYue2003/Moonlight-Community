# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:46:57+08:00
- 采样时长：1m0.0164148s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 114603 | 114603 | 0 | 0 | 1909.82 | 5.652 | 17.773 | 19.997 | 24.561 | 52.337 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 257550 | 2.247 |
| counter | 76415 | 0.667 |
| mysql | 22 | 0.000 |
| redis | 369528 | 3.224 |
| relation | 76415 | 0.667 |

- Cold compute：87787（0.766 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 8332 |
| l1_stale | 60 |
| l2_fresh | 24577 |
| l2_stale | 44248 |
| miss | 37386 |

- L1+L2 Fresh ratio：28.72%
- Refresh max：queue=0 active=18 pending=18

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 87787 | 2.078 |
| counter | 76415 | 2.529 |
| hydrate | 87787 | 1.973 |
| inbox | 87787 | 2.077 |
| merge_dedup | 87787 | 0.002 |
| relation | 76415 | 3.046 |
| route | 87787 | 4.873 |
| total | 114603 | 7.242 |

## Redis 本轮边界增量

- Commands：1371550；input：183395166 bytes；output：229334890 bytes
- Hits/Misses：1737474/194228；run hit rate：89.95%
- Evicted/Rejected：0/0；ops/s max：26782；safety epoch：576 -> 576

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.005 |
| client:loadtest | cpu_percent_total | 48.086 |
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
| docker:zg-canal | cpu_percent | 2.770 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.050 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.310 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 192.270 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 59.970 |
| docker:zg-zk | memory_percent | 1.650 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25231943.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 99.843 |
| process:counter | cpu_seconds_total | 995.797 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54472704.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 198.306 |
| process:gateway | cpu_seconds_total | 2072.781 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 51511296.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 305.775 |
| process:knowpost | cpu_seconds_total | 6507.656 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 172994560.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 138.337 |
| process:relation | cpu_seconds_total | 1353.375 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62337024.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.688 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36225024.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.104 |
| process:user-storage | cpu_seconds_total | 550.391 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 56426496.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 253406801.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 576.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 792825.000 |
| redis | keyspace_hits | 1063506753.000 |
| redis | keyspace_misses | 10760755.000 |
| redis | net_input_bytes | 49257675584.000 |
| redis | net_output_bytes | 199079886529.000 |
| redis | ops_per_sec | 26782.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52056.000 |
| redis | used_memory_bytes | 210989128.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
