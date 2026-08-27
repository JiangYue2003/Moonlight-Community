# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:08:35+08:00
- 采样时长：10.0077623s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 63091 | 63091 | 0 | 0 | 6307.15 | 4.873 | 6.630 | 7.357 | 8.874 | 13.277 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 28148 | 0.446 |
| mysql | 0 | 0.000 |
| redis | 189273 | 3.000 |
| relation | 63091 | 1.000 |

- Cold compute：63091（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 63091 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 63091 | 0.789 |
| counter | 28148 | 1.144 |
| hydrate | 63091 | 0.858 |
| inbox | 63091 | 0.878 |
| merge_dedup | 63091 | 0.003 |
| relation | 63091 | 1.748 |
| route | 63091 | 0.519 |
| total | 63091 | 4.814 |

## Redis 本轮边界增量

- Commands：636068；input：53293510 bytes；output：118441397 bytes
- Hits/Misses：994470/120812；run hit rate：89.17%
- Evicted/Rejected：0/0；ops/s max：74237；safety epoch：593 -> 593

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.504 |
| client:loadtest | cpu_percent_total | 88.057 |
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
| docker:zg-canal | cpu_percent | 1.740 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.560 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.050 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 135.130 |
| docker:zg-kafka | memory_percent | 8.550 |
| docker:zg-kafka | pids | 115.000 |
| docker:zg-zk | cpu_percent | 0.130 |
| docker:zg-zk | memory_percent | 1.540 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27506476.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 10.000 |
| process:counter | cpu_percent | 121.334 |
| process:counter | cpu_seconds_total | 158.844 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54812672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.780 |
| process:gateway | cpu_seconds_total | 426.219 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 46194688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 350.040 |
| process:knowpost | cpu_seconds_total | 1432.234 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 85291008.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 275.210 |
| process:relation | cpu_seconds_total | 1148.375 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58413056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35971072.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 156.047 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 43098112.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 274202371.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 593.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 718524.000 |
| redis | keyspace_hits | 1105192069.000 |
| redis | keyspace_misses | 13719419.000 |
| redis | net_input_bytes | 51526041178.000 |
| redis | net_output_bytes | 204995839865.000 |
| redis | ops_per_sec | 74237.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53304.000 |
| redis | used_memory_bytes | 133277680.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
