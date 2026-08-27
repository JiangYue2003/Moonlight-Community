# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:33:10+08:00
- 采样时长：10.0083581s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 43821 | 43821 | 0 | 0 | 4379.83 | 7.042 | 9.921 | 11.411 | 14.851 | 27.939 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.055 |
| mysql | 0 | 0.000 |
| redis | 131463 | 3.000 |
| relation | 43821 | 1.000 |

- Cold compute：43821（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 43821 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 43821 | 0.336 |
| counter | 2400 | 0.743 |
| hydrate | 43821 | 0.394 |
| inbox | 43821 | 0.360 |
| merge_dedup | 43821 | 0.004 |
| relation | 43821 | 1.109 |
| route | 43821 | 0.045 |
| total | 43821 | 2.269 |

## Redis 本轮边界增量

- Commands：356325；input：40499322 bytes；output：120228464 bytes
- Hits/Misses：817152/47512；run hit rate：94.51%
- Evicted/Rejected：0/0；ops/s max：39056；safety epoch：529 -> 529

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.888 |
| client:loadtest | cpu_percent_total | 78.216 |
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
| docker:zg-es | cpu_percent | 2.220 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.950 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 367.470 |
| docker:zg-kafka | memory_percent | 9.340 |
| docker:zg-kafka | pids | 146.000 |
| docker:zg-zk | cpu_percent | 0.240 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23202062.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 63.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 28.643 |
| process:counter | cpu_seconds_total | 21.125 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 52391936.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 316.090 |
| process:gateway | cpu_seconds_total | 71.469 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 52576256.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 260.638 |
| process:knowpost | cpu_seconds_total | 256.156 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 78589952.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 216.412 |
| process:relation | cpu_seconds_total | 194.062 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 61583360.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35987456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 116.908 |
| process:user-storage | cpu_seconds_total | 33.141 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 53583872.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 204427915.000 |
| redis | connected_clients | 267.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 529.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 681994.000 |
| redis | keyspace_hits | 1002083141.000 |
| redis | keyspace_misses | 5888785.000 |
| redis | net_input_bytes | 42665161020.000 |
| redis | net_output_bytes | 188451021629.000 |
| redis | ops_per_sec | 39056.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47579.000 |
| redis | used_memory_bytes | 103689864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
