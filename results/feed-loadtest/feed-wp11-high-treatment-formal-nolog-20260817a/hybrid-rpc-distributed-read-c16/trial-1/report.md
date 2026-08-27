# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:32:53+08:00
- 采样时长：1m0.0276102s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 376827 | 376827 | 0 | 0 | 6280.09 | 2.115 | 4.938 | 6.399 | 12.246 | 39.583 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 781464 | 2.074 |
| counter | 133281 | 0.354 |
| mysql | 27 | 0.000 |
| redis | 856484 | 2.273 |
| relation | 133281 | 0.354 |

- Cold compute：186819（0.496 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 80286 |
| l1_stale | 321 |
| l2_fresh | 148518 |
| l2_stale | 122588 |
| miss | 25114 |

- L1+L2 Fresh ratio：60.72%
- Refresh max：queue=0 active=28 pending=28

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 186819 | 1.311 |
| counter | 133281 | 1.688 |
| hydrate | 186819 | 1.261 |
| inbox | 186819 | 1.310 |
| merge_dedup | 186819 | 0.002 |
| relation | 133281 | 1.908 |
| route | 186819 | 2.581 |
| total | 376827 | 2.307 |

## Redis 本轮边界增量

- Commands：2860044；input：391216511 bytes；output：588783598 bytes
- Hits/Misses：3741229/329448；run hit rate：91.91%
- Evicted/Rejected：0/0；ops/s max：61718；safety epoch：565 -> 565

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.250 |
| client:loadtest | cpu_percent_total | 100.006 |
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
| docker:zg-canal | cpu_percent | 3.550 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.620 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.250 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 195.430 |
| docker:zg-kafka | memory_percent | 8.790 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 59.290 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23791953.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 88.146 |
| process:counter | cpu_seconds_total | 514.594 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54431744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 1840.859 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54108160.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 412.694 |
| process:knowpost | cpu_seconds_total | 4118.703 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 228274176.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 156.195 |
| process:relation | cpu_seconds_total | 585.641 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62140416.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.391 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36188160.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 456.734 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 49823744.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 220558512.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 565.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 779036.000 |
| redis | keyspace_hits | 1021010656.000 |
| redis | keyspace_misses | 7108214.000 |
| redis | net_input_bytes | 44815627100.000 |
| redis | net_output_bytes | 192135230737.000 |
| redis | ops_per_sec | 61718.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51213.000 |
| redis | used_memory_bytes | 196397824.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
