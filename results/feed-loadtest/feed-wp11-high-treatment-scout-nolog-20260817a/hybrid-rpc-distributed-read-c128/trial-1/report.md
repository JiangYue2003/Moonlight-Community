# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-high-treatment-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:11:26+08:00
- 采样时长：10.0183891s
- 并发：128
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 153451 | 153451 | 0 | 0 | 15332.14 | 5.822 | 22.428 | 32.800 | 45.747 | 165.591 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 124055 | 0.808 |
| counter | 23806 | 0.155 |
| mysql | 40 | 0.000 |
| redis | 182241 | 1.188 |
| relation | 23813 | 0.155 |

- Cold compute：31679（0.206 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 46256 |
| l1_stale | 19163 |
| l2_fresh | 46219 |
| l2_stale | 21517 |
| miss | 20296 |

- L1+L2 Fresh ratio：60.26%
- Refresh max：queue=1024 active=32 pending=1056

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 31688 | 4.414 |
| counter | 23806 | 4.603 |
| hydrate | 31679 | 4.209 |
| inbox | 31688 | 4.413 |
| merge_dedup | 31688 | 0.003 |
| relation | 23813 | 4.730 |
| route | 31690 | 7.221 |
| total | 153451 | 8.054 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 41341 | 0.000 |

## Redis 本轮边界增量

- Commands：558931；input：70135469 bytes；output：118125662 bytes
- Hits/Misses：697359/73085；run hit rate：90.51%
- Evicted/Rejected：0/0；ops/s max：67539；safety epoch：560 -> 560

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.958 |
| client:loadtest | cpu_percent_total | 143.330 |
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
| docker:zg-canal | cpu_percent | 1.750 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.530 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.560 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 236.550 |
| docker:zg-kafka | memory_percent | 8.740 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 37.610 |
| docker:zg-zk | memory_percent | 1.460 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23577430.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 45.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 107.489 |
| process:counter | cpu_seconds_total | 366.984 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54116352.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1763.625 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 49610752.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 468.399 |
| process:knowpost | cpu_seconds_total | 3684.031 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 255303680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 179.944 |
| process:relation | cpu_seconds_total | 446.828 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 61931520.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36397056.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.592 |
| process:user-storage | cpu_seconds_total | 422.250 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 54501376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 214812867.000 |
| redis | connected_clients | 324.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 560.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 784389.000 |
| redis | keyspace_hits | 1015545150.000 |
| redis | keyspace_misses | 6545743.000 |
| redis | net_input_bytes | 44134447068.000 |
| redis | net_output_bytes | 191315641593.000 |
| redis | ops_per_sec | 67539.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49875.000 |
| redis | used_memory_bytes | 210532664.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
