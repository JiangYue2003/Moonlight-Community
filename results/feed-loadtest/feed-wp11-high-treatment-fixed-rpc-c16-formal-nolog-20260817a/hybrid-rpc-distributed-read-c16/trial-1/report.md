# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-fixed-rpc-c16-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:35:23+08:00
- 采样时长：1m0.0325271s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 574447 | 574447 | 0 | 0 | 9573.70 | 1.625 | 2.886 | 3.714 | 5.728 | 17.037 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1035460 | 1.803 |
| counter | 148129 | 0.258 |
| mysql | 29 | 0.000 |
| redis | 1097786 | 1.911 |
| relation | 148129 | 0.258 |

- Cold compute：228612（0.398 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 161767 |
| l1_stale | 425 |
| l2_fresh | 253403 |
| l2_stale | 137944 |
| miss | 20908 |

- L1+L2 Fresh ratio：72.27%
- Refresh max：queue=0 active=26 pending=27

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 228612 | 0.965 |
| counter | 148129 | 1.372 |
| hydrate | 228612 | 0.960 |
| inbox | 228612 | 0.964 |
| merge_dedup | 228612 | 0.002 |
| relation | 148129 | 1.568 |
| route | 228612 | 1.919 |
| total | 574447 | 1.449 |

## Redis 本轮边界增量

- Commands：3507498；input：481401651 bytes；output：768810561 bytes
- Hits/Misses：4606035/378427；run hit rate：92.41%
- Evicted/Rejected：0/0；ops/s max：63908；safety epoch：614 -> 614

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.522 |
| client:loadtest | cpu_percent_total | 120.351 |
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
| docker:zg-canal | cpu_percent | 2.860 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.130 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.040 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 206.440 |
| docker:zg-kafka | memory_percent | 8.830 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 53.480 |
| docker:zg-zk | memory_percent | 1.780 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29793262.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 99.273 |
| process:counter | cpu_seconds_total | 284.578 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 55259136.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 618.328 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 45789184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 434.264 |
| process:knowpost | cpu_seconds_total | 2411.453 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 236367872.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 189.811 |
| process:relation | cpu_seconds_total | 418.766 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56762368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.438 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35700736.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 152.094 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 41349120.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 307225405.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 614.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 795900.000 |
| redis | keyspace_hits | 1151750794.000 |
| redis | keyspace_misses | 18421670.000 |
| redis | net_input_bytes | 55099511242.000 |
| redis | net_output_bytes | 211713453304.000 |
| redis | ops_per_sec | 63908.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54962.000 |
| redis | used_memory_bytes | 213758608.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
