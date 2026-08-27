# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:32:03+08:00
- 采样时长：10.008956s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 66676 | 66676 | 0 | 0 | 6664.40 | 9.102 | 12.950 | 13.930 | 16.720 | 31.425 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.036 |
| mysql | 0 | 0.000 |
| redis | 200028 | 3.000 |
| relation | 66676 | 1.000 |

- Cold compute：66676（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 66676 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 66676 | 1.864 |
| counter | 2400 | 2.518 |
| hydrate | 66676 | 2.183 |
| inbox | 66676 | 2.247 |
| merge_dedup | 66676 | 0.004 |
| relation | 66676 | 2.890 |
| route | 66676 | 0.094 |
| total | 66676 | 9.301 |

## Redis 本轮边界增量

- Commands：525748；input：60822854 bytes；output：182702771 bytes
- Hits/Misses：1236302/72255；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：55826；safety epoch：526 -> 526

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.196 |
| client:loadtest | cpu_percent_total | 99.130 |
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
| docker:zg-canal | cpu_percent | 1.720 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.590 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.210 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 232.850 |
| docker:zg-kafka | memory_percent | 8.730 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 35.950 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23015580.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 50.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 23.150 |
| process:counter | cpu_seconds_total | 14.547 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 52244480.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.016 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 37203968.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 323.951 |
| process:knowpost | cpu_seconds_total | 154.156 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 73580544.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 285.512 |
| process:relation | cpu_seconds_total | 114.641 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 57892864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35721216.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.125 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 34680832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 202909684.000 |
| redis | connected_clients | 179.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 526.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 677095.000 |
| redis | keyspace_hits | 998671954.000 |
| redis | keyspace_misses | 5689521.000 |
| redis | net_input_bytes | 42493256837.000 |
| redis | net_output_bytes | 187948767924.000 |
| redis | ops_per_sec | 55826.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47512.000 |
| redis | used_memory_bytes | 103230864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
