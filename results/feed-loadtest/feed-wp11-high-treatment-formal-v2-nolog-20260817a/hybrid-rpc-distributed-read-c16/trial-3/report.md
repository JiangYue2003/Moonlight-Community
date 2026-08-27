# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:40:28+08:00
- 采样时长：1m0.0305282s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 538836 | 538836 | 0 | 0 | 8980.27 | 2.064 | 3.151 | 3.820 | 5.942 | 35.391 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 997741 | 1.852 |
| counter | 145987 | 0.271 |
| mysql | 19 | 0.000 |
| redis | 1060877 | 1.969 |
| relation | 145987 | 0.271 |

- Cold compute：222558（0.413 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 144894 |
| l1_stale | 435 |
| l2_fresh | 234815 |
| l2_stale | 137516 |
| miss | 21176 |

- L1+L2 Fresh ratio：70.47%
- Refresh max：queue=0 active=26 pending=26

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 222558 | 1.018 |
| counter | 145987 | 1.425 |
| hydrate | 222558 | 1.004 |
| inbox | 222558 | 1.017 |
| merge_dedup | 222558 | 0.002 |
| relation | 145987 | 1.622 |
| route | 222558 | 2.014 |
| total | 538836 | 1.559 |

## Redis 本轮边界增量

- Commands：3410712；input：468159124 bytes；output：741404399 bytes
- Hits/Misses：4478242/370877；run hit rate：92.35%
- Evicted/Rejected：0/0；ops/s max：62627；safety epoch：571 -> 571

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.335 |
| client:loadtest | cpu_percent_total | 117.362 |
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
| docker:zg-canal | cpu_percent | 2.830 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.090 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.200 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 197.230 |
| docker:zg-kafka | memory_percent | 8.780 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 64.760 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24662368.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 113.137 |
| process:counter | cpu_seconds_total | 777.938 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 55009280.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1840.953 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45228032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 440.583 |
| process:knowpost | cpu_seconds_total | 5481.375 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 237735936.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 140.687 |
| process:relation | cpu_seconds_total | 1005.922 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62492672.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 1.562 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36200448.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.791 |
| process:user-storage | cpu_seconds_total | 457.016 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43593728.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 240815680.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 571.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 814936.000 |
| redis | keyspace_hits | 1047341934.000 |
| redis | keyspace_misses | 9317080.000 |
| redis | net_input_bytes | 47576351439.000 |
| redis | net_output_bytes | 196478420606.000 |
| redis | ops_per_sec | 62627.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51667.000 |
| redis | used_memory_bytes | 234060632.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
