# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:31:44+08:00
- 采样时长：10.0071944s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 65995 | 65995 | 0 | 0 | 6597.44 | 4.779 | 6.223 | 6.829 | 8.105 | 15.211 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.036 |
| mysql | 0 | 0.000 |
| redis | 197985 | 3.000 |
| relation | 65995 | 1.000 |

- Cold compute：65995（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 65995 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 65995 | 0.785 |
| counter | 2400 | 1.323 |
| hydrate | 65995 | 0.949 |
| inbox | 65995 | 0.962 |
| merge_dedup | 65995 | 0.004 |
| relation | 65995 | 1.803 |
| route | 65995 | 0.052 |
| total | 65995 | 4.574 |

## Redis 本轮边界增量

- Commands：524284；input：60472118 bytes；output：180920355 bytes
- Hits/Misses：1223901/71514；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：55357；safety epoch：525 -> 525

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.992 |
| client:loadtest | cpu_percent_total | 95.869 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.870 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.800 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 78.490 |
| docker:zg-kafka | memory_percent | 8.310 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 0.140 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22934923.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 34.151 |
| process:counter | cpu_seconds_total | 12.562 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 45805568.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.016 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 37203968.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 351.854 |
| process:knowpost | cpu_seconds_total | 115.797 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 66457600.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 289.689 |
| process:relation | cpu_seconds_total | 82.547 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 53071872.000 |
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
| process:user-storage | rss_bytes | 34471936.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 202262604.000 |
| redis | connected_clients | 83.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 525.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 677094.000 |
| redis | keyspace_hits | 997172568.000 |
| redis | keyspace_misses | 5602149.000 |
| redis | net_input_bytes | 42419093136.000 |
| redis | net_output_bytes | 187727552851.000 |
| redis | ops_per_sec | 55357.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47493.000 |
| redis | used_memory_bytes | 99134456.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
