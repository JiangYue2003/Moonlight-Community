# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-treatment-fixed-gateway-c8-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:41:48+08:00
- 采样时长：1m0.013446s
- 并发：8
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 84477 | 84477 | 0 | 0 | 1407.80 | 4.188 | 13.270 | 15.158 | 18.625 | 30.540 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 177321 | 2.099 |
| counter | 61749 | 0.731 |
| mysql | 23 | 0.000 |
| redis | 285818 | 3.383 |
| relation | 61749 | 0.731 |

- Cold compute：68721（0.813 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4797 |
| l1_stale | 22 |
| l2_fresh | 14605 |
| l2_stale | 28869 |
| miss | 36184 |

- L1+L2 Fresh ratio：22.97%
- Refresh max：queue=0 active=8 pending=8

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 68721 | 1.045 |
| counter | 61749 | 1.412 |
| hydrate | 68721 | 0.979 |
| inbox | 68721 | 1.044 |
| merge_dedup | 68721 | 0.002 |
| relation | 61749 | 2.121 |
| route | 68721 | 3.195 |
| total | 84477 | 4.488 |

## Redis 本轮边界增量

- Commands：1126040；input：146912649 bytes；output：172011315 bytes
- Hits/Misses：1355904/160940；run hit rate：89.39%
- Evicted/Rejected：0/0；ops/s max：26745；safety epoch：619 -> 619

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.364 |
| client:loadtest | cpu_percent_total | 37.830 |
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
| docker:zg-canal | cpu_percent | 1.620 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.330 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.560 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 175.320 |
| docker:zg-kafka | memory_percent | 8.840 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 53.600 |
| docker:zg-zk | memory_percent | 1.590 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 30348035.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 96.785 |
| process:counter | cpu_seconds_total | 536.578 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54681600.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 184.024 |
| process:gateway | cpu_seconds_total | 933.234 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 51576832.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.926 |
| process:knowpost | cpu_seconds_total | 3424.359 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 167317504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 136.040 |
| process:relation | cpu_seconds_total | 803.578 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56672256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.609 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35651584.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 77.272 |
| process:user-storage | cpu_seconds_total | 283.219 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 49250304.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 319125609.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 619.000 |
| redis | hit_rate | 0.983 |
| redis | keys | 797737.000 |
| redis | keyspace_hits | 1166845586.000 |
| redis | keyspace_misses | 19837594.000 |
| redis | net_input_bytes | 56699906253.000 |
| redis | net_output_bytes | 214034915187.000 |
| redis | ops_per_sec | 26745.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 55347.000 |
| redis | used_memory_bytes | 213036912.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
