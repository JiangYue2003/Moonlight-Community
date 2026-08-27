# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:34:04+08:00
- 采样时长：1m0.0316068s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 556647 | 556647 | 0 | 0 | 9277.09 | 1.707 | 3.091 | 3.753 | 5.856 | 29.175 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1016474 | 1.826 |
| counter | 147205 | 0.264 |
| mysql | 20 | 0.000 |
| redis | 1078969 | 1.938 |
| relation | 147205 | 0.264 |

- Cold compute：225540（0.405 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 153499 |
| l1_stale | 475 |
| l2_fresh | 244276 |
| l2_stale | 137428 |
| miss | 20969 |

- L1+L2 Fresh ratio：71.46%
- Refresh max：queue=0 active=24 pending=24

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 225540 | 0.995 |
| counter | 147205 | 1.399 |
| hydrate | 225540 | 0.982 |
| inbox | 225540 | 0.994 |
| merge_dedup | 225540 | 0.002 |
| relation | 147205 | 1.593 |
| route | 225540 | 1.967 |
| total | 556647 | 1.505 |

## Redis 本轮边界增量

- Commands：3458803；input：474656036 bytes；output：754919741 bytes
- Hits/Misses：4541925/374626；run hit rate：92.38%
- Evicted/Rejected：0/0；ops/s max：64123；safety epoch：566 -> 566

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.296 |
| client:loadtest | cpu_percent_total | 116.735 |
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
| docker:zg-canal | cpu_percent | 2.710 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.420 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.200 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 256.900 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 65.400 |
| docker:zg-zk | memory_percent | 1.490 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23954610.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 105.238 |
| process:counter | cpu_seconds_total | 563.531 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54853632.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1840.891 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50159616.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 443.604 |
| process:knowpost | cpu_seconds_total | 4368.750 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 234397696.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 154.780 |
| process:relation | cpu_seconds_total | 662.219 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62181376.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.406 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36192256.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 456.766 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 49823744.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 224347228.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 566.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 793186.000 |
| redis | keyspace_hits | 1025971724.000 |
| redis | keyspace_misses | 7517347.000 |
| redis | net_input_bytes | 45334243436.000 |
| redis | net_output_bytes | 192958156264.000 |
| redis | ops_per_sec | 64123.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51283.000 |
| redis | used_memory_bytes | 210301032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
