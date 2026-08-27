# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:33:34+08:00
- 采样时长：10.0111586s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 42321 | 42321 | 0 | 0 | 4228.47 | 14.457 | 21.051 | 25.213 | 31.664 | 65.484 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.057 |
| mysql | 18 | 0.000 |
| redis | 126981 | 3.000 |
| relation | 42321 | 1.000 |

- Cold compute：42321（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 42321 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 42321 | 0.964 |
| counter | 2400 | 2.274 |
| hydrate | 42321 | 1.095 |
| inbox | 42321 | 1.101 |
| merge_dedup | 42321 | 0.004 |
| relation | 42321 | 1.891 |
| route | 42321 | 0.133 |
| total | 42321 | 5.208 |

## Redis 本轮边界增量

- Commands：344846；input：39136715 bytes；output：116078945 bytes
- Hits/Misses：789449/45899；run hit rate：94.51%
- Evicted/Rejected：0/0；ops/s max：35522；safety epoch：530 -> 530

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.838 |
| client:loadtest | cpu_percent_total | 77.414 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.160 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.590 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 306.870 |
| docker:zg-kafka | memory_percent | 9.220 |
| docker:zg-kafka | pids | 152.000 |
| docker:zg-zk | cpu_percent | 31.980 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23254643.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 55.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 22.295 |
| process:counter | cpu_seconds_total | 23.547 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 52252672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 320.051 |
| process:gateway | cpu_seconds_total | 106.625 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 55001088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 304.960 |
| process:knowpost | cpu_seconds_total | 289.656 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 78639104.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 238.469 |
| process:relation | cpu_seconds_total | 220.031 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 61140992.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35987456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 130.114 |
| process:user-storage | cpu_seconds_total | 50.266 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 60518400.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 204865605.000 |
| redis | connected_clients | 267.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 530.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 683220.000 |
| redis | keyspace_hits | 1003045188.000 |
| redis | keyspace_misses | 5944327.000 |
| redis | net_input_bytes | 42713787733.000 |
| redis | net_output_bytes | 188592691472.000 |
| redis | ops_per_sec | 35522.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47603.000 |
| redis | used_memory_bytes | 105660160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
