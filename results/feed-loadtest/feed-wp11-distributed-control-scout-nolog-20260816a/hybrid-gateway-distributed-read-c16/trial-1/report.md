# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:32:45+08:00
- 采样时长：10.0054151s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 42532 | 42532 | 0 | 0 | 4252.04 | 3.697 | 4.873 | 5.380 | 6.625 | 11.876 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.056 |
| mysql | 31 | 0.001 |
| redis | 127627 | 3.001 |
| relation | 42532 | 1.000 |

- Cold compute：42532（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 42532 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 42532 | 0.236 |
| counter | 2400 | 0.565 |
| hydrate | 42532 | 0.279 |
| inbox | 42532 | 0.239 |
| merge_dedup | 42532 | 0.004 |
| relation | 42532 | 0.969 |
| route | 42532 | 0.036 |
| total | 42532 | 1.783 |

## Redis 本轮边界增量

- Commands：346539；input：39339099 bytes；output：116648104 bytes
- Hits/Misses：793256/46164；run hit rate：94.50%
- Evicted/Rejected：0/0；ops/s max：38441；safety epoch：528 -> 528

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.343 |
| client:loadtest | cpu_percent_total | 69.494 |
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
| docker:zg-canal | cpu_percent | 0.340 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.920 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 290.940 |
| docker:zg-kafka | memory_percent | 9.130 |
| docker:zg-kafka | pids | 142.000 |
| docker:zg-zk | cpu_percent | 0.080 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23148141.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 20.590 |
| process:counter | cpu_seconds_total | 18.578 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 52428800.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 315.979 |
| process:gateway | cpu_seconds_total | 35.578 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 51335168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.662 |
| process:knowpost | cpu_seconds_total | 224.984 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 80384000.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 214.912 |
| process:relation | cpu_seconds_total | 170.078 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 61091840.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35848192.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 138.379 |
| process:user-storage | cpu_seconds_total | 16.906 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 48373760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 203978591.000 |
| redis | connected_clients | 267.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 528.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 680770.000 |
| redis | keyspace_hits | 1001093834.000 |
| redis | keyspace_misses | 5831672.000 |
| redis | net_input_bytes | 42615172339.000 |
| redis | net_output_bytes | 188305260786.000 |
| redis | ops_per_sec | 38441.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47554.000 |
| redis | used_memory_bytes | 102866296.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
