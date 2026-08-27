# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-fixed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:55:33+08:00
- 采样时长：10.0049936s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 30634 | 30634 | 0 | 0 | 3062.35 | 5.024 | 6.861 | 7.489 | 8.812 | 13.401 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.078 |
| mysql | 0 | 0.000 |
| redis | 91902 | 3.000 |
| relation | 30634 | 1.000 |

- Cold compute：30634（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 30634 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 30634 | 0.737 |
| counter | 2400 | 1.125 |
| hydrate | 30634 | 0.821 |
| inbox | 30634 | 0.815 |
| merge_dedup | 30634 | 0.004 |
| relation | 30634 | 1.700 |
| route | 30634 | 0.092 |
| total | 30634 | 4.187 |

## Redis 本轮边界增量

- Commands：259373；input：28836358 bytes；output：84225738 bytes
- Hits/Misses：575436/33189；run hit rate：94.55%
- Evicted/Rejected：0/0；ops/s max：29675；safety epoch：581 -> 581

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.451 |
| client:loadtest | cpu_percent_total | 71.214 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.680 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 0.900 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 229.810 |
| docker:zg-kafka | memory_percent | 8.880 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 36.290 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25548444.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 61.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 27.827 |
| process:counter | cpu_seconds_total | 10.625 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 52932608.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 241.484 |
| process:gateway | cpu_seconds_total | 26.641 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50683904.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 259.285 |
| process:knowpost | cpu_seconds_total | 110.062 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 71319552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 209.347 |
| process:relation | cpu_seconds_total | 85.922 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 53600256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35491840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 100.495 |
| process:user-storage | cpu_seconds_total | 10.969 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 43884544.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 257276786.000 |
| redis | connected_clients | 168.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 581.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 747637.000 |
| redis | keyspace_hits | 1069746336.000 |
| redis | keyspace_misses | 11287673.000 |
| redis | net_input_bytes | 49726921080.000 |
| redis | net_output_bytes | 199941051216.000 |
| redis | ops_per_sec | 29675.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52522.000 |
| redis | used_memory_bytes | 160624560.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
