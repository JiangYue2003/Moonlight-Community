# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:03:16+08:00
- 采样时长：1m0.0136314s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 73149 | 73149 | 0 | 0 | 1218.99 | 13.105 | 20.471 | 23.021 | 28.315 | 45.350 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12359 | 0.169 |
| mysql | 86 | 0.001 |
| redis | 219533 | 3.001 |
| relation | 73149 | 1.000 |

- Cold compute：73149（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 73149 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 73149 | 2.126 |
| counter | 12359 | 3.012 |
| hydrate | 73149 | 2.283 |
| inbox | 73149 | 2.302 |
| merge_dedup | 73149 | 0.004 |
| relation | 73149 | 3.974 |
| route | 73149 | 0.515 |
| total | 73149 | 11.222 |

## Redis 本轮边界增量

- Commands：662467；input：70220331 bytes；output：202742140 bytes
- Hits/Misses：1417590/79359；run hit rate：94.70%
- Evicted/Rejected：0/0；ops/s max：14465；safety epoch：588 -> 588

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.919 |
| client:loadtest | cpu_percent_total | 46.708 |
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
| docker:zg-canal | cpu_percent | 2.270 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.150 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.590 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 181.810 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 49.920 |
| docker:zg-zk | memory_percent | 1.790 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27131960.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 51.909 |
| process:counter | cpu_seconds_total | 102.719 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53616640.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 219.763 |
| process:gateway | cpu_seconds_total | 424.734 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51933184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 259.429 |
| process:knowpost | cpu_seconds_total | 1209.438 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 70705152.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 216.546 |
| process:relation | cpu_seconds_total | 990.375 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57278464.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 0.484 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36061184.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 79.914 |
| process:user-storage | cpu_seconds_total | 155.969 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 49319936.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 270211267.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 588.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 726697.000 |
| redis | keyspace_hits | 1099305954.000 |
| redis | keyspace_misses | 13002209.000 |
| redis | net_input_bytes | 51194049534.000 |
| redis | net_output_bytes | 204289332760.000 |
| redis | ops_per_sec | 14465.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53035.000 |
| redis | used_memory_bytes | 139947984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
